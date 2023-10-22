package igdb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Token struct {
	AccessToken string
	expiresAt   time.Time
	expiresIn   int
}

type authResponse struct {
	AccessToken string `json:"access_token,required"`
	ExpiresIn   int    `json:"expires_in,required"`
	TokenType   string `json:"token_type,required"`
}

func NewToken(clientID string, secret string) (Token, error) {
	newToken := Token{}

	response, err := newToken.retrieve(clientID, secret)
	if err != nil {
		return newToken, err
	}

	expirationTime := time.Now().Local().Add(time.Second * time.Duration(response.ExpiresIn))

	return Token{
		AccessToken: response.AccessToken,
		expiresIn:   response.ExpiresIn,
		expiresAt:   expirationTime,
	}, nil
}

func (*Token) retrieve(clientID string, secret string) (authResponse, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", "https://id.twitch.tv/oauth2/token", nil)
	query := req.URL.Query()

	query.Add("client_id", clientID)
	query.Add("client_secret", secret)
	query.Add("grant_type", "client_credentials")
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return authResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return authResponse{}, err
	}

	if resp.StatusCode != 200 {
		return authResponse{}, fmt.Errorf("%s: %s", resp.Status, body)
	}

	var response authResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return authResponse{}, err
	}

	log.Println("Authentication succeed!")
	return response, nil
}
