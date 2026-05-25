package notion

import (
	"context"

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

func (p *Page) Update(requester PageUpdateRequester) (*notionapi.Page, error) {
	request := requester.UpdateRequest()
	pageID := notionapi.PageID(p.Id)

	return p.apiClient.Page.Update(context.Background(), pageID, &request)
}
