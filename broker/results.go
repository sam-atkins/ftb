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

func ResultsByLeague(league string) {
	client := api.Client{}
	endpoint := fmt.Sprintf("competitions/%s/matches", league)

	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
	}

	fmt.Printf("Results from the %v\n", response.Body.Competition.Name)

	header := []string{"Date", "Home", "", "", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		if v.Season.CurrentMatchday-1 == v.Matchday {
			rows = append(rows, []string{
				fmt.Sprint(v.UtcDate.Format("2006 January 02")),
				v.HomeTeam.Name,
				fmt.Sprint(v.Score.FullTime.HomeTeam),
				fmt.Sprint(v.Score.FullTime.AwayTeam),
				v.AwayTeam.Name,
			})
		}
	}
	writer.Table(header, rows)
}

// ResultsByTeam fetches results for a team and prints to stdout. Arg matchLimit limits
// the results to the previous three weeks
func ResultsByTeam(teamCode string, matchLimit bool) {
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

	client := api.Client{}
	endpoint := fmt.Sprintf("teams/%s/matches?status=FINISHED", teamId)
	if matchLimit {
		now := time.Now()
		dateFrom := now.AddDate(0, 0, -21).Format("2006-01-02")
		dateTo := now.AddDate(0, 0, 21).Format("2006-01-02")
		endpoint = fmt.Sprintf("teams/%s/matches?status=FINISHED&dateFrom=%s&dateTo=%s", teamId, dateFrom, dateTo)
	}

	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
	}

	fmt.Printf("Results for %s\n", teamName)

	header := []string{"Date", "Home", "", "", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		rows = append(rows, []string{
			fmt.Sprint(v.UtcDate.Format("2006 January 02")),
			v.HomeTeam.Name,
			fmt.Sprint(v.Score.FullTime.HomeTeam),
			fmt.Sprint(v.Score.FullTime.AwayTeam),
			v.AwayTeam.Name,
		})
	}
	writer.Table(header, rows)
}
