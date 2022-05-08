package reporter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

// Test Fixtures
func testServerFixture(t *testing.T) (*httptest.Server, *http.ServeMux, func()) {
	viper.SetConfigName("test_config")
	viper.AddConfigPath("../testdata/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	return server, mux, func() {
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

func Test_table_getTable(t *testing.T) {
	t.Parallel()

	server, mux, teardown := testServerFixture(t)
	defer teardown()

	endpoint := "/competitions/BL1/standings"
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(loadTestJson("../testdata/standings-BL1.json")))
	})

	teamCode := "FCB"
	tb := newTeamTable(teamCode)
	tb.getTable(server.URL)
	got := tb.rows
	want := [][]string{
		{
			"1",
			"FC Bayern München",
			"11",
			"9",
			"1",
			"1",
			"40",
			"11",
			"29",
			"28",
		},
		{
			"2",
			"Borussia Dortmund",
			"11",
			"8",
			"0",
			"3",
			"28",
			"17",
			"11",
			"24",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getTable() got %v, want %v", got, want)
	}
}
