package reporter

import (
	"fmt"
	"log"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// GetTable gets a league table
func GetTable(league string) {
	endpoint := fmt.Sprintf("competitions/%s/standings", league)
	client := api.NewClient()
	response, responseErr := client.GetTable(endpoint)
	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}

	fmt.Printf("League table: %v\n", response.Body.Competition.Name)

	header := []string{"Pos", "Team", "Played", "Won", "Draw", "Lost", "+", "-", "GD", "Points"}
	var rows [][]string
	for _, v := range response.Body.Standings[0].Table {
		rows = append(rows, []string{
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
	writer.Table(header, rows)
}

// GetTable gets the league table for the given team
func GetTableForTeam(teamCode string) {
	leagueCode, teamName, _ := config.GetTeamInfoFromUserTeamCode(teamCode)

	endpoint := fmt.Sprintf("competitions/%s/standings", leagueCode)
	client := api.NewClient()
	response, responseErr := client.GetTable(endpoint)
	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}

	fmt.Printf("League table: %v\n", response.Body.Competition.Name)

	var teamIndex int
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

	header := []string{"Pos", "Team", "Played", "Won", "Draw", "Lost", "+", "-", "GD", "Points"}
	writer.TableWithTeamHighlight(teamIndex, header, data)
}
