package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"notion-igdb-autocomplete/config"
	"strings"
	"time"

	"github.com/jomei/notionapi"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Unable to load config: %s\n", err)
	} else {
		log.Println("Successfully loaded config!")
	}

	titleCleaner := strings.NewReplacer("{{", "", "}}", "")

	log.Println("Looking for pages to update...")
	for range time.Tick(time.Duration(config.WatcherTickDelay)) {
		pages, err := fetchPages(config.NotionAPISecret, config.NotionPageID)
		if err != nil {
			log.Fatalf("Unable to fetch pages: %s\n", err)
		}

		for _, obj := range pages {
			id := obj.ID
			title := obj.Properties["Title"].(*notionapi.TitleProperty).Title[0].Text.Content
			cleanTitle := titleCleaner.Replace(title)

			err = callUpdater(config.UpdaterURL(), updaterBody{PageID: id.String(), Search: cleanTitle})
			if err != nil {
				log.Printf("Unable to update page '%s': %s\n", id, err)
			} else {
				log.Printf("page '%s' successfully updated!", id)
			}
		}
	}
}

func fetchPages(apiSecret string, databaseID string) ([]notionapi.Page, error) {
	notion := notionapi.NewClient(notionapi.Token(apiSecret))
	query := &notionapi.DatabaseQueryRequest{
		Filter: &notionapi.AndCompoundFilter{
			notionapi.PropertyFilter{
				Property: "Title",
				RichText: &notionapi.TextFilterCondition{
					StartsWith: "{{",
				},
			},
			notionapi.PropertyFilter{
				Property: "Title",
				RichText: &notionapi.TextFilterCondition{
					EndsWith: "}}",
				},
			},
		},
	}

	result, err := notion.Database.Query(context.Background(), notionapi.DatabaseID(databaseID), query)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}

type updaterBody struct {
	PageID string `json:"page_id,required"`
	Search string `json:"search,required"`
}

func callUpdater(updaterURL string, payload updaterBody) error {
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
