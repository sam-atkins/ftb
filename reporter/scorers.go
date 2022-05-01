package reporter

import (
	"fmt"
	"log"
	"os"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/writer"
)

// ScorersCLI is the entrypoint for the reporter to get top scorers for a league
func ScorersCLI(league string) {
	s := newScorers(league)
	s.getScorersByLeague()
	writer.NewTable(s.header, s.message, s.rows).Render()
}

type scorers struct {
	endpoint string
	league   string
	message  string
	header   []string
	rows     [][]string
}

func newScorers(league string) *scorers {
	return &scorers{
		league: league,
	}
}

func (s *scorers) getScorersByLeague() *scorers {
	s.endpoint = scorersURL(s.league)
	response, err := fetchScorers(s.endpoint)
	if err != nil {
		log.Printf("Something went wrong with the request: %s\n", err)
		os.Exit(1)
	}
	s.message = fmt.Sprintf("Top Scorers in the %v\n", response.Body.Competition.Name)
	s.header = []string{"Name", "Team", "Goals"}
	s.buildScorersByLeagueRows(response)
	return s
}

func (s *scorers) buildScorersByLeagueRows(response *api.ApiScorersResponse) *scorers {
	var rows [][]string
	for _, v := range response.Body.Scorers {
		rows = append(rows, []string{
			v.Player.Name,
			v.Team.Name,
			fmt.Sprint(v.NumberOfGoals),
		})
	}
	s.rows = rows
	return s
}

func fetchScorers(endpoint string) (*api.ApiScorersResponse, error) {
	client := api.NewClient()
	response, responseErr := client.GetScorers(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}
