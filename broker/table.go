/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package broker

import (
	"fmt"
	"strings"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// GetTable fetches gets the league table and prints to stdout
func GetTable(league string) {
	client := api.Client{}
	endpoint := fmt.Sprintf("competitions/%s/standings", league)
	response, responseErr := client.GetTable(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
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

// GetTable fetches gets the league table and prints to stdout
func GetTableForTeam(teamCode string) {
	var leagueCode string
	for _, v := range config.TeamConfig {
		if v.Code == strings.ToUpper(teamCode) {
			leagueCode = v.LeagueCode

		}
	}

	client := api.Client{}
	endpoint := fmt.Sprintf("competitions/%s/standings", leagueCode)
	response, responseErr := client.GetTable(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
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
