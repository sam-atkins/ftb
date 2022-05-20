package reporter

import (
	"fmt"
	"log"
	"os"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// TableCLI renders a league table
func TableCLI(league string) {
	t := newTable(league)
	t.getTable()
	writer.NewTable(t.header, t.message, t.rows).Render()
}

// TeamTableCLI renders a league table with the specified team highlighted
func TeamTableCLI(teamCode string) {
	t := newTeamTable(teamCode)
	t.getTable()
	writer.NewTableWithPositionHighlight(t.header, t.message, t.rows, t.teamIndex).Render()
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
	cfg := config.NewTeamConfig(teamCode)
	return &table{
		header:       []string{"Pos", "Team", "Played", "Won", "Draw", "Lost", "+", "-", "GD", "Points"},
		leagueCode:   cfg.LeagueCode,
		tableForTeam: true,
		teamCode:     teamCode,
		teamName:     cfg.TeamName,
	}
}

// getTable prepares the data required in order to render the table. The variadic baseUrl
// arg is optional. If not passed in, it will use the default baseUrl set in the API
// Client.
func (t *table) getTable(baseUrl ...string) *table {
	t.endpoint = buildLeagueStandingsURL(t.leagueCode)
	response, err := fetchTable(t.endpoint, baseUrl...)
	if err != nil {
		log.Printf("Something went wrong with the request: %s\n", err)
		os.Exit(1)
	}
	t.message = fmt.Sprintf("League table: %v\n", response.Competition.Name)
	if t.tableForTeam {
		t.buildLeagueTableRows(response, t.teamName)
	} else {
		t.buildLeagueTableRows(response)
	}
	return t
}

func fetchTable(endpoint string, baseUrl ...string) (*api.LeagueResponse, error) {
	client := api.NewClient(baseUrl...)
	response, err := client.GetTable(endpoint)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// buildLeagueTableRows builds the rows for the league table and returns the rows and an
// updated table instance with the index where the team is located. If you don't want to
// highlight a team, do not pass in the teamName arg when calling this method. The index
// returned will set as -1 if the team is not found or the teamName arg is not provided.
func (t *table) buildLeagueTableRows(response *api.LeagueResponse, teamName ...string) *table {
	teamIndex := -1
	var data [][]string
	for i, v := range response.Standings[0].Table {
		if len(teamName) > 0 && v.Team.Name == teamName[0] {
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
	t.rows = data
	t.teamIndex = teamIndex
	return t
}
