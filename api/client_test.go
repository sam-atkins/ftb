package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/viper"
)

// test fixtures

func testClient(t *testing.T) (clientAPI, *http.ServeMux, func()) {
	viper.SetConfigName("test_config")
	viper.AddConfigPath("../test_files/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(server.URL)

	return client, mux, func() {
		server.Close()
	}
}

func loadTestJson(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	return content
}

// tests

func Test_client_GetMatches_200(t *testing.T) {
	client, mux, teardown := testClient(t)
	defer teardown()

	endpoint := "/competitions/BL1/matches"
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(loadTestJson("../test_files/matches-FCB.json")))
	})
	wantRes := &apiMatchesResponse{
		StatusCode: http.StatusOK,
		Body: matchesResponse{
			Matches: []matches{
				{
					ID:       329670,
					Status:   "SCHEDULED",
					Matchday: 4,
				},
			},
		},
	}

	got, _ := client.GetMatches(endpoint)
	if statusCode := got.StatusCode; statusCode != wantRes.StatusCode {
		t.Errorf("client.GetMatches() error = %v, wantErr %v", got.StatusCode, wantRes.StatusCode)
		return
	}

	if matchID := got.Body.Matches[0].ID; matchID != wantRes.Body.Matches[0].ID {
		t.Errorf("client.GetMatches() Matches[0].ID = %v, want %v", got.Body.Matches[0].ID, wantRes.Body.Matches[0].ID)
	}
	if matchStatus := got.Body.Matches[0].Status; matchStatus != wantRes.Body.Matches[0].Status {
		t.Errorf("client.GetMatches() Matches[0].Status = %v, want %v", got.Body.Matches[0].Status, wantRes.Body.Matches[0].Status)
	}
	if matchDay := got.Body.Matches[0].Matchday; matchDay != wantRes.Body.Matches[0].Matchday {
		t.Errorf("client.GetMatches() Matches[0].Matchday = %v, want %v", got.Body.Matches[0].Matchday, wantRes.Body.Matches[0].Matchday)
	}
}

func Test_client_GetMatches_400(t *testing.T) {
	client, mux, teardown := testClient(t)
	defer teardown()

	endpoint := "/competitions/PL/matches"
	resBody := `{"message":"Your API token is invalid.","errorCode":400}`
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, resBody)
	})
	wantErr := true
	_, err := client.GetMatches(endpoint)
	if (err != nil) != wantErr {
		t.Errorf("client.GetMatches() error = %v, wantErr %v", err, resBody)
		return
	}
}

func Test_client_GetScorers_200(t *testing.T) {
	client, mux, teardown := testClient(t)
	defer teardown()

	endpoint := "/competitions/BL1/scorers"
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(loadTestJson("../test_files/top-scorers-BL1.json")))
	})
	wantRes := &apiScorersResponse{
		StatusCode: http.StatusOK,
		Body: scorersResponse{
			Scorers: []scorers{
				{
					Player: player{
						ID:   371,
						Name: "Robert Lewandowski",
					},
					NumberOfGoals: 5,
				},
			},
		},
	}

	got, _ := client.GetScorers(endpoint)
	if statusCode := got.StatusCode; statusCode != wantRes.StatusCode {
		t.Errorf("client.GetScorers() error = %v, wantErr %v", got.StatusCode, wantRes.StatusCode)
		return
	}

	if playerID := got.Body.Scorers[0].Player.ID; playerID != wantRes.Body.Scorers[0].Player.ID {
		t.Errorf("client.GetMatches() Scorers[0].Player.ID = %v, want %v", got.Body.Scorers[0].Player.ID, wantRes.Body.Scorers[0].Player.ID)
	}
	if playerName := got.Body.Scorers[0].Player.Name; playerName != wantRes.Body.Scorers[0].Player.Name {
		t.Errorf("client.GetMatches() Scorers[0].Player.Name = %v, want %v", got.Body.Scorers[0].Player.Name, wantRes.Body.Scorers[0].Player.Name)
	}
	if numberOfGoals := got.Body.Scorers[0].NumberOfGoals; numberOfGoals != wantRes.Body.Scorers[0].NumberOfGoals {
		t.Errorf("client.GetMatches() numberOfGoals = %v, want %v", got.Body.Scorers[0].NumberOfGoals, wantRes.Body.Scorers[0].NumberOfGoals)
	}
}

func Test_client_GetScorers_400(t *testing.T) {
	client, mux, teardown := testClient(t)
	defer teardown()

	endpoint := "/competitions/PL/scorers"
	resBody := `{"message":"Your API token is invalid.","errorCode":400}`
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, resBody)
	})
	wantErr := true
	_, err := client.GetScorers(endpoint)
	if (err != nil) != wantErr {
		t.Errorf("client.GetScorers() error = %v, wantErr %v", err, resBody)
		return
	}
}
