/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package broker

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// MatchesByLeague fetches matches for a league and prints to stdout
func MatchesByLeague(league string) {
	endpoint := fmt.Sprintf("competitions/%s/matches", league)
	c := api.NewClient()
	response, responseErr := c.GetMatches(endpoint)

	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
	}

	fmt.Printf("Next match day fixtures in the %v\n", response.Body.Competition.Name)

	header := []string{"Date", "Home", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		if v.Season.CurrentMatchday == v.Matchday {
			rows = append(rows, []string{
				fmt.Sprint(v.UtcDate.Local().Format(dateTimeFormat)),
				v.HomeTeam.Name,
				v.AwayTeam.Name,
			})
		}
	}
	writer.Table(header, rows)
}

// MatchesByTeam fetches matches for a team and prints to stdout. Arg matchLimit limits
// the results to the next three weeks
func MatchesByTeam(teamCode string, matchLimit bool) {
	var teamId string
	var teamName string
	for _, v := range config.TeamConfig {
		if v.Code == strings.ToUpper(teamCode) {
			teamId = v.Id
			teamName = v.Name
		}
	}

	if teamId == "" {
		fmt.Println("Did not recognise that team. These are the available team codes:")
		header := []string{"Team", "Code"}
		teamCodes := config.GetTeamCodes()
		writer.Table(header, teamCodes)
		os.Exit(1)
	}

	endpoint := fmt.Sprintf("teams/%s/matches?status=SCHEDULED", teamId)
	client := api.NewClient()
	if matchLimit {
		now := time.Now()
		dateFrom := now.AddDate(0, 0, daysAgo).Format("2006-01-02")
		dateTo := now.AddDate(0, 0, daysAhead).Format("2006-01-02")
		endpoint = fmt.Sprintf("teams/%s/matches?status=SCHEDULED&dateFrom=%s&dateTo=%s", teamId, dateFrom, dateTo)
	}

	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
	}

	fmt.Printf("Matches for %s\n", teamName)

	header := []string{"Date", "Competition", "Home", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		rows = append(rows, []string{
			fmt.Sprint(v.UtcDate.Local().Format(dateTimeFormat)),
			v.Competition.Name,
			v.HomeTeam.Name,
			v.AwayTeam.Name,
		})
	}
	writer.Table(header, rows)
}
