package notion

import "github.com/jomei/notionapi"

type PageUpdateRequester interface {
	UpdateRequest() notionapi.PageUpdateRequest
}
