package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"notion-igdb-autocomplete/config"
	"notion-igdb-autocomplete/igdb"
	"sort"
	"time"

	"github.com/agnivade/levenshtein"
	"github.com/gin-gonic/gin"
	"github.com/jomei/notionapi"
)

type body struct {
	PageID string `json:"page_id,required"`
	Search string `json:"search,required"`
}

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Unable to load config: %s\n", err)
	}
	log.Println("Successfully loaded config!")

	igdbClient, err := igdb.NewClient(config.IGDBClientID, config.IGDBSecret)
	if err != nil {
		log.Fatalf("Unable to create IGDB client: %s\n", err)
	}
	log.Println("Successfully created IGDB client!")

	notionClient := notionapi.NewClient(notionapi.Token(config.NotionAPISecret))
	log.Println("Successfully created Notion client!")

	server := gin.Default()
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

		updatedPage, err := updateNotionPage(payload.PageID, game, notionClient)
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
	query := igdb.NewSearchQuery(gameName, []string{"name", "platforms.name", "first_release_date", "franchises.name", "genres.name", "cover.image_id"})
	results, err := client.SearchGame(query)
	if err != nil {
		return nil, err
	}

	if len(results) <= 0 {
		return nil, fmt.Errorf("cannot find game '%s'", gameName)
	}

	return findBestGame(gameName, results), nil
}

func updateNotionPage(pageID string, game *igdb.Game, notionClient *notionapi.Client) (*notionapi.Page, error) {
	releaseDate := notionapi.Date(time.Unix(game.ReleaseDate, 0))
	platforms := game.NotionPlatforms()
	genres := game.NotionGenres()
	franchises := game.NotionFranchises()

	updateReq := notionapi.PageUpdateRequest{
		Cover: &notionapi.Image{
			Type: "external",
			External: &notionapi.FileObject{
				URL: game.CoverURL(),
			},
		},
		Properties: notionapi.Properties{
			"Title": notionapi.TitleProperty{
				Type: notionapi.PropertyTypeTitle,
				Title: []notionapi.RichText{
					{Text: &notionapi.Text{Content: game.Name}},
				},
			},
			"Release date": notionapi.DateProperty{
				Type: notionapi.PropertyTypeDate,

				Date: &notionapi.DateObject{
					Start: &releaseDate,
				},
			},
		},
	}

	if len(platforms) > 0 {
		updateReq.Properties["Platforms"] = notionapi.MultiSelectProperty{
			Type:        notionapi.PropertyTypeMultiSelect,
			MultiSelect: game.NotionPlatforms(),
		}
	}

	if len(franchises) > 0 {
		updateReq.Properties["Franchises"] = notionapi.MultiSelectProperty{
			Type:        notionapi.PropertyTypeMultiSelect,
			MultiSelect: franchises,
		}
	}

	if len(genres) > 0 {
		updateReq.Properties["Genres"] = notionapi.MultiSelectProperty{
			Type:        notionapi.PropertyTypeMultiSelect,
			MultiSelect: game.NotionGenres(),
		}
	}

	page, err := notionClient.Page.Update(context.Background(), notionapi.PageID(pageID), &updateReq)
	if err != nil {
		return nil, err
	}

	return page, nil
}

type ComparedGames []ComparedGame
type ComparedGame struct {
	Game  igdb.Game
	Index int
}

// Implements interface sort.Interface
func (cg ComparedGames) Len() int {
	return len(cg)
}

// Implements interface sort.Interface
func (cg ComparedGames) Swap(i, j int) {
	cg[i], cg[j] = cg[j], cg[i]
}

// Implements interface sort.Interface
func (cg ComparedGames) Less(i, j int) bool {
	return cg[i].Index < cg[j].Index
}

func findBestGame(search string, games igdb.Games) *igdb.Game {
	var comparisons ComparedGames

	for _, game := range games {
		comparisons = append(comparisons, ComparedGame{
			Game:  game,
			Index: levenshtein.ComputeDistance(search, game.Name),
		})
	}

	sort.Sort(comparisons)

	return &comparisons[0].Game
}
