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
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Unable to load config: %s\n", err)
	} else {
		log.Println("Successfully loaded config!")
	}

	notionClient := notion.NewClient(config.NotionAPISecret)
	log.Println("Successfully created Notion client!")

	titleCleaner := strings.NewReplacer("{{", "", "}}", "")

	log.Println("Looking for pages to update...")
	for range time.Tick(time.Duration(config.WatcherTickDelay)) {
		entries, err := notionClient.Database(config.NotionPageID).GetEntries()
		if err != nil {
			log.Fatalf("Unable to fetch pages: %s\n", err)
		}

		for _, entry := range entries {
			id := entry.ID
			title := entry.Properties["Title"].(*notionapi.TitleProperty).Title[0].Text.Content
			cleanTitle := titleCleaner.Replace(title)

			err = callUpdater(config.UpdaterURL(), updateRequestBody{PageID: id.String(), Search: cleanTitle})
			if err != nil {
				log.Printf("Unable to update page '%s': %s\n", id, err)
			} else {
				log.Printf("page '%s' successfully updated!", id)
			}
		}
	}
}

func callUpdater(updaterURL string, payload updateRequestBody) error {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("PUT", updaterURL, strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}

	log.Printf("Requesting update: %s\n", reqBody)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s: %s", resp.Status, body)
	}

	return nil
}
