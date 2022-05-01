package reporter

import (
	"fmt"
	"log"
	"os"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// consolidate with TeamTableCLI ?
func TableCLI(league string) {
	GetTable(league)
}

func TeamTableCLI(teamCode string) {
	handleGetTeamTable(teamCode)
}

type table struct {
	endpoint     string
	message      string
	header       []string
	leagueCode   string
	rows         [][]string
	tableForTeam bool
	teamCode     string
	teamName     string
	teamIndex    int
}

func newTable(leagueCode string) *table {
	return &table{
		leagueCode:   leagueCode,
		header:       []string{"Pos", "Team", "Played", "Won", "Draw", "Lost", "+", "-", "GD", "Points"},
		tableForTeam: false,
	}
}

func newTeamTable(teamCode string) *table {
	return &table{
		header:       []string{"Pos", "Team", "Played", "Won", "Draw", "Lost", "+", "-", "GD", "Points"},
		tableForTeam: true,
		teamCode:     teamCode,
	}
}

func handleGetTable(leagueCode string) {}

func handleGetTeamTable(teamCode string) {
	t := newTeamTable(teamCode)
	t.getTable()
	writer.NewTableWithPositionHighlight(t.header, t.message, t.rows, t.teamIndex).Render()
}

func (t *table) getTable() *table {
	t.leagueCode, t.teamName, _ = config.GetTeamInfoFromUserTeamCode(t.teamCode)
	t.endpoint = buildLeagueStandingsURL(t.leagueCode)
	response, err := fetchTable(t.endpoint)
	if err != nil {
		log.Printf("Something went wrong with the request: %s\n", err)
		os.Exit(1)
	}
	t.message = fmt.Sprintf("League table: %v\n", response.Body.Competition.Name)
	t.teamIndex, t.rows = buildLeagueTableRows(response, t.teamName)
	return t
}

func fetchTable(endpoint string) (*api.ApiLeagueResponse, error) {
	client := api.NewClient()
	response, err := client.GetTable(endpoint)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// GetTable gets a league table
func GetTable(league string) {
	endpoint := fmt.Sprintf("competitions/%s/standings", league)
	client := api.NewClient()
	response, responseErr := client.GetTable(endpoint)
	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}

	message := fmt.Sprintf("League table: %v\n", response.Body.Competition.Name)

	header := []string{"Pos", "Team", "Played", "Won", "Draw", "Lost", "+", "-", "GD", "Points"}
	_, rows := buildLeagueTableRows(response, "")
	writer.NewTable(header, message, rows).Render()
}

// buildLeagueTableRows builds the rows for the league table and returns the index where
// the team is located. If you don't want to highlight a team, pass in an empty string ""
// and index returned will be -1.
func buildLeagueTableRows(response *api.ApiLeagueResponse, teamName string) (int, [][]string) {
	teamIndex := -1
	var data [][]string
	for i, v := range response.Body.Standings[0].Table {
		if v.Team.Name == teamName {
			teamIndex = i
		}
		data = append(data,
			[]string{
				fmt.Sprint(v.Position),
				v.Team.Name,
				fmt.Sprint(v.PlayedGames),
				fmt.Sprint(v.Won),
				fmt.Sprint(v.Draw),
				fmt.Sprint(v.Lost),
				fmt.Sprint(v.GoalsFor),
				fmt.Sprint(v.GoalsAgainst),
				fmt.Sprint(v.GoalDifference),
				fmt.Sprint(v.Points),
			})
	}
	return teamIndex, data
}
