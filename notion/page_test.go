package notion

import (
	"testing"

	"github.com/jomei/notionapi"
)

var testClient = notionapi.NewClient(notionapi.Token("apiSecret"))

func TestNewPage(t *testing.T) {
	pageID := "pageID1"
	page := NewPage(pageID, testClient)

	if page.Id != pageID {
		t.Errorf("Unexpected page.Id. expected='%s', got='%s'\n", pageID, page.Id)
	}

	if page.apiClient != testClient {
		t.Errorf("Unexpected page.apiClient. expected='%v', got='%v'\n", testClient, page.apiClient)
	}
}
