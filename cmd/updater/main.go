package main

import (
	"fmt"
	"log"
	"net/http"
	"notion-igdb-autocomplete/choose"
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/howlongtobeat"
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

	hltbClient := howlongtobeat.NewClient()

	server := gin.Default()
	server.GET("/heartbeat", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "I'm alive!"})
	})

	server.PUT("/game_infos", func(ctx *gin.Context) {
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

		updatedPage, err := notionClient.Page(payload.PageID).UpdateGameInfos(game)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Updated page %s with game %s infos", updatedPage.ID, game.Name)})
	})

	server.PUT("/time_to_beat", func(ctx *gin.Context) {
		var payload body

		err := ctx.ShouldBindJSON(&payload)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		hltbGame, err := searchHowLongToBeatGame(payload.Search, &hltbClient)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		}

		updatedPage, err := notionClient.Page(payload.PageID).UpdateTimeToBeat(hltbGame)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Updated %s(page %s) with time to beat infos", payload.Search, updatedPage.ID)})

		ctx.Status(http.StatusOK)
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
		return &igdb.Game{Name: fmt.Sprintf("Not found (%s)", gameName)}, nil
	}

	return choose.Game(gameName, results), nil
}

func searchHowLongToBeatGame(gameName string, client *howlongtobeat.Client) (*howlongtobeat.Game, error) {
	games, err := client.SearchGame(gameName)
	if err != nil {
		log.Fatalf("cannot find time infos: %s", err)
	}

	if len(games) <= 0 {
		return &howlongtobeat.Game{Name: fmt.Sprintf("Not found (%s)", gameName), CompletionMain: 0, CompletionPlus: 0, CompletionFull: 0}, nil
	}

	return &games[0], nil
}
