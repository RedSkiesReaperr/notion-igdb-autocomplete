package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/notion"
	"strings"
	"time"

	"github.com/jomei/notionapi"
)

type updateRequestBody struct {
	PageID string `json:"page_id,required"`
	Search string `json:"search,required"`
}

func main() {
	config, err := config.Load("")
	if err != nil {
		log.Fatalf("Unable to load config: %s\n", err)
	} else {
		log.Println("Successfully loaded config!")
	}

	updater := updaterClient{
		GameInfosUrl:  fmt.Sprintf("%s%s", config.UpdaterURL(), "game_infos"),
		TimeToBeatUrl: fmt.Sprintf("%s%s", config.UpdaterURL(), "time_to_beat"),
	}

	notionClient := notion.NewClient(config.NotionAPISecret)
	log.Println("Successfully created Notion client!")

	gamesToUpdate := make(chan updateRequestBody, 20)
	TtbToUpdate := make(chan updateRequestBody, 20)

	runningGameUpdates := make(map[string]notionapi.ObjectID)
	runningTimeToBeatUpdates := make(map[string]notionapi.ObjectID)

	go updater.updateGameInfos(gamesToUpdate, runningGameUpdates)
	go updater.updateTimeToBeat(TtbToUpdate, runningTimeToBeatUpdates)

	log.Println("Looking for pages to update...")
	for range time.Tick(time.Duration(config.WatcherTickDelay) * time.Second) {
		emptyGamesEntries, err := notionClient.Database(config.NotionPageID).GetEmptyGamesEntries()
		if err != nil {
			log.Fatalf("Unable to fetch empty games pages: %s\n", err)
		}

		timeToBeatEntries, err := notionClient.Database(config.NotionPageID).GetEmptyTimeToBeatEntries()
		if err != nil {
			log.Fatalf("Unable to fetch time to beat games pages: %s\n", err)
		}

		go enqueueGameInfosRequests(gamesToUpdate, emptyGamesEntries, runningGameUpdates)
		go enqueueTimeToBeatRequests(TtbToUpdate, timeToBeatEntries, runningTimeToBeatUpdates)
	}
}

type updaterClient struct {
	GameInfosUrl  string
	TimeToBeatUrl string
}

func enqueueGameInfosRequests(c chan<- updateRequestBody, entries []notionapi.Page, runningUpdates map[string]notionapi.ObjectID) {
	titleCleaner := strings.NewReplacer("{{", "", "}}", "")

	for _, entry := range entries {
		id := entry.ID
		title := entry.Properties["Title"].(*notionapi.TitleProperty).Title[0].Text.Content
		cleanTitle := titleCleaner.Replace(title)

		if _, ok := runningUpdates[id.String()]; !ok {
			runningUpdates[id.String()] = id

			c <- updateRequestBody{PageID: id.String(), Search: cleanTitle}
		}
	}
}

func enqueueTimeToBeatRequests(c chan<- updateRequestBody, entries []notionapi.Page, runningUpdates map[string]notionapi.ObjectID) {
	for _, entry := range entries {
		id := entry.ID
		title := entry.Properties["Title"].(*notionapi.TitleProperty).Title[0].Text.Content

		if _, ok := runningUpdates[id.String()]; !ok {
			runningUpdates[id.String()] = id

			c <- updateRequestBody{PageID: id.String(), Search: title}
		}
	}
}

func (uc updaterClient) updateGameInfos(c chan updateRequestBody, runningUpdates map[string]notionapi.ObjectID) {
	for updateRequest := range c {
		if err := uc.call(uc.GameInfosUrl, updateRequest); err != nil {
			log.Printf("UpdateGameInfos failed: %s", err)
		} else {
			log.Printf("Successfully updated page '%s' !", updateRequest.PageID)
		}

		delete(runningUpdates, updateRequest.PageID)
	}
}

func (uc updaterClient) updateTimeToBeat(c chan updateRequestBody, runningUpdates map[string]notionapi.ObjectID) {
	for updateRequest := range c {
		if err := uc.call(uc.TimeToBeatUrl, updateRequest); err != nil {
			log.Printf("updateTimeToBeat failed: %s", err)
		} else {
			log.Printf("Successfully updated page '%s' !", updateRequest.PageID)
		}

		delete(runningUpdates, updateRequest.PageID)
	}
}

func (uc updaterClient) call(url string, updateRequest updateRequestBody) error {
	reqBody, err := json.Marshal(updateRequest)
	if err != nil {
		return fmt.Errorf("marshall error: err='%s', pageID='%s', body='%v'\n", err, updateRequest.PageID, updateRequest)
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("PUT", url, strings.NewReader(string(reqBody)))
	if err != nil {
		return fmt.Errorf("creation error: err='%s', pageID='%s', body='%s'\n", err, updateRequest.PageID, string(reqBody))
	}

	log.Printf("Requesting %s => %s\n", url, reqBody)
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: err='%s', response='%v'\n", err, resp)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid request response: err='%s', response='%v'\n", err, resp)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("cannot update: status='%s', body='%s', pageID='%s'\n", resp.Status, body, updateRequest.PageID)
	}

	return nil
}
