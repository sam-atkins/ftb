package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sam-atkins/httpc"
	"github.com/spf13/viper"
)

// type ClientAPI interface {
type ClientAPI interface {
	// GetMatches returns the competition matches
	GetMatches(endpoint string) (*ApiMatchesResponse, error)
	// GetScorers returns the tops scorers in a league
	GetScorers(endpoint string) (*ApiScorersResponse, error)
	// GetTable returns the league table
	GetTable(endpoint string) (*LeagueResponse, error)
	// GetTeams returns all teams in a competition
	GetTeams(endpoint string) (*apiTeamsResponse, error)
}

type client struct {
	baseURL string
	token   string
}

// NewClient is a factory interface to the clientAPI
func NewClient(url ...string) ClientAPI {
	token := viper.GetString("TOKEN")
	if token == "" {
		fmt.Print("API token missing from config")
		os.Exit(1)
	}
	baseUrl := "https://api.football-data.org/v2"
	if len(url) != 0 {
		baseUrl = url[0]
	}
	return client{
		baseURL: baseUrl,
		token:   token,
	}
}

func (c client) GetMatches(endpoint string) (*ApiMatchesResponse, error) {
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

	clientResponse := &ApiMatchesResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}

func (c client) GetScorers(endpoint string) (*ApiScorersResponse, error) {
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

	clientResponse := &ApiScorersResponse{
		StatusCode: response.StatusCode,
		Body:       decodedResponse,
	}

	return clientResponse, nil
}

func (c client) GetTable(endpoint string) (*LeagueResponse, error) {
	headers := map[string]string{"X-Auth-Token": c.token}
	var leagueRes LeagueResponse
	err := httpc.GetJson(c.baseURL + endpoint).AddHeaders(headers).Load(&leagueRes)
	if err != nil {
		return nil, err
	}
	return &leagueRes, nil
}

func (c client) GetTeams(endpoint string) (*apiTeamsResponse, error) {
	response, responseErr := c.doRequest(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	defer response.Body.Close()
	var decodedResponse teamsResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&decodedResponse)
	if decodeErr != nil {
		return nil, decodeErr
	}

	editedResponse := editTeamsResponse(decodedResponse)

	clientResponse := &apiTeamsResponse{
		StatusCode: response.StatusCode,
		Body:       editedResponse,
	}

	return clientResponse, nil
}

// doRequest is a helper function which makes a HTTP request
func (c client) doRequest(endpoint string) (*http.Response, error) {
	headers := map[string]string{"X-Auth-Token": c.token}
	response, respErr := httpc.GetJson(c.baseURL + endpoint).AddHeaders(headers).Do()
	if respErr != nil {
		return nil, respErr
	}

	if response.StatusCode != http.StatusOK {
		resBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		return nil, errors.New(string(resBody))
	}

	return response, nil
}

// This edits the teamsResponse. It removes the short code (Tla) collision between FC
// Bayern and FC Barcelona. Both have "FCB" as their code.
func editTeamsResponse(response teamsResponse) teamsResponse {
	for i, v := range response.Teams {
		if v.Tla == "FCB" && v.Name == "FC Barcelona" {
			response.Teams[i].Tla = "BAR"
		}
	}

	return response
}
