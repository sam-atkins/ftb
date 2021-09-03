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

type APILeagueResponse struct {
	StatusCode int
	Body       LeagueResponse
}

type APIResultsResponse struct {
	StatusCode int
	Body       CompetitionResponse
}

type APIScorersResponse struct {
	StatusCode int
	Body       ScorersResponse
}

func (c *Client) BaseURL() string {
	if c.baseURL == "" {
		return BASE_URL
	}
	return c.baseURL
}

// GetResults returns the competition matches, decoded against the CompetitionResponse
// struct
func (c *Client) GetResults(endpoint string) (*APIResultsResponse, error) {
	response, responseErr := c.doRequest(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	defer response.Body.Close()
	var decodedResponse CompetitionResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	if response.StatusCode != 200 {
		fmt.Printf("API request status: %v", response.StatusCode)
		os.Exit(1)
	}

	clientResponse := &APIResultsResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}

// GetScorers returns the tops scorers, decoded against the ScorersResponse struct
func (c *Client) GetScorers(endpoint string) (*APIScorersResponse, error) {
	response, responseErr := c.doRequest(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	defer response.Body.Close()
	var decodedResponse ScorersResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	if response.StatusCode != 200 {
		fmt.Printf("API request status: %v", response.StatusCode)
		os.Exit(1)
	}

	clientResponse := &APIScorersResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}

// GetTable returns the league table, decoded against the LeagueResponse struct
func (c *Client) GetTable(endpoint string) (*APILeagueResponse, error) {
	response, responseErr := c.doRequest(endpoint)
	if responseErr != nil {
		return nil, responseErr
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

	clientResponse := &APILeagueResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}

// doRequest makes a request to the API and returns an HTTP response
func (c *Client) doRequest(endpoint string) (*http.Response, error) {
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
	return response, nil
}
