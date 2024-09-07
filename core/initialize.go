package core

import (
	"fmt"
	"log"

	"notion-igdb-autocomplete/igdb"
	"notion-igdb-autocomplete/notion"

	"github.com/RedSkiesReaperr/howlongtobeat"
)

func (c *Core) initializeNotion() error {
	log.Println("Initializing Notion client")
	log.Println(".... Creating")
	client := notion.NewClient(c.Config.NotionAPISecret)
	c.notion = &client

	log.Println(".... Authenticating")
	log.Println(".... Testing")
	if err := c.notion.TestConnection(); err != nil {
		return fmt.Errorf("test failed: %v", err)
	}

	log.Println(".... Success!")
	return nil
}

func (c *Core) initializeIGDB() error {
	log.Println("Initializing IGDB client")
	log.Println(".... Creating")
	log.Println(".... Authenticating")
	client, err := igdb.NewClient(c.Config.IGDBClientID, c.Config.IGDBSecret)
	if err != nil {
		return fmt.Errorf("create failed: %v", err)
	}
	c.igdb = &client

	log.Println(".... Success!")
	return nil
}

func (c *Core) initializeHLTB() error {
	log.Println("Initializing HowLongToBeat client")
	log.Println(".... Creating")
	client, err := howlongtobeat.New()
	if err != nil {
		return fmt.Errorf("create failed: %v", err)
	}

	c.hltb = client

	log.Println(".... Success!")
	return nil
}
