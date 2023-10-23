package notion

import (
	"context"
	"notion-igdb-autocomplete/igdb"
	"time"

	"github.com/jomei/notionapi"
)

type Page struct {
	apiClient *notionapi.Client
	Id        string
}

func NewPage(id string, client *notionapi.Client) *Page {
	return &Page{
		Id:        id,
		apiClient: client,
	}
}

func (p *Page) Update(game *igdb.Game) (*notionapi.Page, error) {
	request := createUpdateRequest(game)
	pageID := notionapi.PageID(p.Id)

	updatedPage, err := p.apiClient.Page.Update(context.Background(), pageID, &request)
	if err != nil {
		return nil, err
	}

	return updatedPage, nil
}

func createUpdateRequest(game *igdb.Game) notionapi.PageUpdateRequest {
	releaseDate := notionapi.Date(time.Unix(game.ReleaseDate, 0))
	platforms := game.NotionPlatforms()
	genres := game.NotionGenres()
	franchises := game.NotionFranchises()

	request := notionapi.PageUpdateRequest{
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
		request.Properties["Platforms"] = notionapi.MultiSelectProperty{
			Type:        notionapi.PropertyTypeMultiSelect,
			MultiSelect: game.NotionPlatforms(),
		}
	}

	if len(franchises) > 0 {
		request.Properties["Franchises"] = notionapi.MultiSelectProperty{
			Type:        notionapi.PropertyTypeMultiSelect,
			MultiSelect: franchises,
		}
	}

	if len(genres) > 0 {
		request.Properties["Genres"] = notionapi.MultiSelectProperty{
			Type:        notionapi.PropertyTypeMultiSelect,
			MultiSelect: game.NotionGenres(),
		}
	}

	return request
}
