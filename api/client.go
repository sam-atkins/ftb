/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

const BASE_URL = "https://api.football-data.org/v2/"

// Client for the Football-Data API, requires an API Token
type Client struct {
	baseURL string
}

// TODO(sam) rename to APILeagueResponse, how to reuse?
type APIResponse struct {
	StatusCode int
	Body       LeagueResponse
}

func (c *Client) BaseURL() string {
	if c.baseURL == "" {
		return BASE_URL
	}
	return c.baseURL
}

// DoRequest makes a request to the API and returns an APIResponse
func (c *Client) DoRequest(endpoint string) (*APIResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL()+endpoint, nil)
	if err != nil {
		return nil, err
	}

	token := viper.GetString("TOKEN")
	if token == "" {
		fmt.Print("API token missing from config")
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("X-Auth-Token", token)

	var client http.Client
	response, respErr := client.Do(req)
	if respErr != nil {
		return nil, respErr
	}
	defer response.Body.Close()
	var decodedResponse LeagueResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	if response.StatusCode != 200 {
		fmt.Printf("API request status: %v", response.StatusCode)
		os.Exit(1)
	}

	clientResponse := &APIResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}
