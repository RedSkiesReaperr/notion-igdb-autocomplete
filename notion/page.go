package notion

import (
	"context"
	"notion-igdb-autocomplete/howlongtobeat"
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

func (p *Page) UpdateGameInfos(game *igdb.Game) (*notionapi.Page, error) {
	request := createGameInfosUpdateRequest(game)
	pageID := notionapi.PageID(p.Id)

	updatedPage, err := p.apiClient.Page.Update(context.Background(), pageID, &request)
	if err != nil {
		return nil, err
	}

	return updatedPage, nil
}

func (p *Page) UpdateTimeToBeat(timeInfos *howlongtobeat.Game) (*notionapi.Page, error) {
	request := createTimeToBeatUpdateRequest(timeInfos)
	pageID := notionapi.PageID(p.Id)

	updatedPage, err := p.apiClient.Page.Update(context.Background(), pageID, &request)
	if err != nil {
		return nil, err
	}

	return updatedPage, nil
}

func createGameInfosUpdateRequest(game *igdb.Game) notionapi.PageUpdateRequest {
	releaseDate := notionapi.Date(time.Unix(game.ReleaseDate, 0))

	return notionapi.PageUpdateRequest{
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
			"Franchises": notionapi.MultiSelectProperty{
				Type:        notionapi.PropertyTypeMultiSelect,
				MultiSelect: game.NotionFranchises(),
			},
			"Genres": notionapi.MultiSelectProperty{
				Type:        notionapi.PropertyTypeMultiSelect,
				MultiSelect: game.NotionGenres(),
			},
			"Platforms": notionapi.MultiSelectProperty{
				Type:        notionapi.PropertyTypeMultiSelect,
				MultiSelect: game.NotionPlatforms(),
			},
		},
	}
}

func createTimeToBeatUpdateRequest(timeInfos *howlongtobeat.Game) notionapi.PageUpdateRequest {
	return notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"Time to complete (Main Story)": notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: timeInfos.ReadableCompletion(timeInfos.CompletionMain),
						},
					},
				},
			},
			"Time to complete (Main + Sides)": notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: timeInfos.ReadableCompletion(timeInfos.CompletionPlus),
						},
					},
				},
			},
			"Time to complete (Completionist)": notionapi.RichTextProperty{
				Type: notionapi.PropertyTypeRichText,
				RichText: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: timeInfos.ReadableCompletion(timeInfos.CompletionFull),
						},
					},
				},
			},
		},
	}
}
