package howlongtobeat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/corpix/uarand"
)

type Client struct {
	UserAgent string
}

func NewClient() Client {
	return Client{
		UserAgent: uarand.GetRandom(),
	}
}

func (c Client) SearchGame(gameName string) (Games, error) {
	request := NewRequest(gameName)
	games := Games{}
	httpClient := &http.Client{}

	requestBody, err := json.Marshal(request.Body())
	if err != nil {
		return games, fmt.Errorf("invalid request body: %s", err)
	}

	req, _ := http.NewRequest("POST", "https://howlongtobeat.com/api/search", strings.NewReader(string(requestBody)))
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
		"User-Agent":   {c.UserAgent},
		"Referer":      {"https://howlongtobeat.com/"},
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

	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return games, err
	}

	return result.Data, nil
}
