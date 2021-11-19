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
	viper.AddConfigPath("../")
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
		fmt.Fprint(w, string(loadTestJson("../examples/matches-FCB.json")))
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

	if got.StatusCode == http.StatusOK {
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
