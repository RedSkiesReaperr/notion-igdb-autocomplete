package notion

import (
	"context"
	"github.com/jomei/notionapi"
)

type Client struct {
	APIClient notionapi.Client
}

func NewClient(apiSecret string) Client {
	return Client{
		APIClient: *notionapi.NewClient(notionapi.Token(apiSecret)),
	}
}

func (c *Client) Page(id string) *Page {
	return NewPage(id, &c.APIClient)
}

func (c *Client) Database(id string) *Database {
	return NewDatabase(id, &c.APIClient)
}

func (c *Client) TestConnection() error {
	_, err := c.APIClient.User.Me(context.Background())

	return err
}
