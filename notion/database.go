package notion

import (
	"context"

	"github.com/jomei/notionapi"
)

type Database struct {
	apiClient *notionapi.Client
	Id        string
}

func NewDatabase(id string, client *notionapi.Client) *Database {
	return &Database{
		Id:        id,
		apiClient: client,
	}
}

func (d *Database) GetEntries() ([]notionapi.Page, error) {
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

	result, err := d.apiClient.Database.Query(context.Background(), notionapi.DatabaseID(d.Id), query)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}
