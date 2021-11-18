package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type clientAPI interface {
	// GetMatches returns the competition matches
	GetMatches(endpoint string) (*apiMatchesResponse, error)
	// GetScorers returns the tops scorers in a league
	GetScorers(endpoint string) (*apiScorersResponse, error)
	// GetTable returns the league table
	GetTable(endpoint string) (*apiLeagueResponse, error)
}

type client struct {
	baseURL string
	token   string
}

// NewClient is a factory interface to the clientAPI
func NewClient(url ...string) clientAPI {
	token := viper.GetString("TOKEN")
	if token == "" {
		fmt.Print("API token missing from config")
		os.Exit(1)
	}
	baseUrl := "https://api.football-data.org/v2/"
	if len(url) != 0 {
		baseUrl = url[0]
	}
	return client{
		baseURL: baseUrl,
		token:   token,
	}
}

func (c client) GetMatches(endpoint string) (*apiMatchesResponse, error) {
	response, responseErr := c.doRequest(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	defer response.Body.Close()
	var decodedResponse matchesResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	clientResponse := &apiMatchesResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}

func (c client) GetScorers(endpoint string) (*apiScorersResponse, error) {
	response, responseErr := c.doRequest(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	defer response.Body.Close()
	var decodedResponse scorersResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	clientResponse := &apiScorersResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}

func (c client) GetTable(endpoint string) (*apiLeagueResponse, error) {
	response, responseErr := c.doRequest(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	defer response.Body.Close()
	var decodedResponse leagueResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	clientResponse := &apiLeagueResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}

// // doRequest makes a request to the API and returns an HTTP response
func (c client) doRequest(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("X-Auth-Token", c.token)

	var client http.Client
	response, respErr := client.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	if response.StatusCode != 200 {
		fmt.Printf("API request status: %v", response.StatusCode)
		os.Exit(1)
	}

	return response, nil
}
