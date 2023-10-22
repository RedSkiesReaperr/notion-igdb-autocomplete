package igdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	clientID  string
	apiSecret string
	token     Token
}

func NewClient(clientID string, secret string) (Client, error) {
	token, err := NewToken(clientID, secret)
	if err != nil {
		return Client{}, fmt.Errorf("unable to create token: %s", err)
	}

	return Client{
		clientID:  clientID,
		apiSecret: secret,
		token:     token,
	}, nil
}

func (c *Client) SearchGame(query string) (Games, error) {
	games := Games{}
	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.igdb.com/v4/games", strings.NewReader(query))
	req.Header = map[string][]string{
		"Accept":        {"application/json"},
		"Client-ID":     {c.clientID},
		"Authorization": {fmt.Sprintf("Bearer %s", c.token.AccessToken)},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return games, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return games, err
	}

	if resp.StatusCode != 200 {
		return games, fmt.Errorf("%s: %s", resp.Status, body)
	}

	var gameList Games
	err = json.Unmarshal(body, &gameList)
	if err != nil {
		return games, err
	}

	return gameList, nil
}
