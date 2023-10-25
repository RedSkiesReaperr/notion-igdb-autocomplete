package main

import (
	"fmt"
	"log"
	"net/http"
	"notion-igdb-autocomplete/choose"
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/igdb"
	"notion-igdb-autocomplete/notion"

	"github.com/gin-gonic/gin"
)

type body struct {
	PageID string `json:"page_id,required"`
	Search string `json:"search,required"`
}

func main() {
	config, err := config.Load("")
	if err != nil {
		log.Fatalf("Unable to load config: %s\n", err)
	}
	log.Println("Successfully loaded config!")

	igdbClient, err := igdb.NewClient(config.IGDBClientID, config.IGDBSecret)
	if err != nil {
		log.Fatalf("Unable to create IGDB client: %s\n", err)
	}
	log.Println("Successfully created IGDB client!")

	notionClient := notion.NewClient(config.NotionAPISecret)
	log.Println("Successfully created Notion client!")

	server := gin.Default()
	server.GET("/heartbeat", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "I'm alive!"})
	})

	server.PUT("/", func(ctx *gin.Context) {
		var payload body

		err := ctx.ShouldBindJSON(&payload)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		game, err := searchIgdbGame(payload.Search, &igdbClient)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		updatedPage, err := notionClient.Page(payload.PageID).Update(game)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Updated page %s with game %s infos", updatedPage.ID, game.Name)})
	})

	err = server.Run(fmt.Sprintf("0.0.0.0:%d", config.UpdaterPort))
	if err != nil {
		log.Fatalf("Unable to start server: %s\n", err)
	}
}

func searchIgdbGame(gameName string, client *igdb.Client) (*igdb.Game, error) {
	query := igdb.NewSearchQuery(gameName, "name", "platforms.name", "first_release_date", "franchises.name", "genres.name", "cover.image_id")
	results, err := client.SearchGame(query)
	if err != nil {
		return nil, err
	}

	if len(results) <= 0 {
		return nil, fmt.Errorf("cannot find game '%s'", gameName)
	}

	return choose.Game(gameName, results), nil
}
