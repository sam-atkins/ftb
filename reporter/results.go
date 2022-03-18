/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package reporter

import (
	"fmt"
	"log"
	"time"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// ResultsByLeague gets the results for a football league and pretty prints
// the league table to stdout
func ResultsByLeague(league string) {
	endpoint := fmt.Sprintf("competitions/%s/matches", league)
	client := api.NewClient()
	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}

	fmt.Printf("Results from the %v\n", response.Body.Competition.Name)

	header := []string{"Date", "Home", "", "", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		if v.Season.CurrentMatchday-1 == v.Matchday && v.Status == "FINISHED" ||
			v.Season.CurrentMatchday == v.Matchday && v.Status == "FINISHED" {
			rows = append(rows, []string{
				fmt.Sprint(v.UtcDate.Local().Format(dateTimeFormat)),
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
	_, teamName, teamId := config.GetTeamInfoFromUserTeamCode(teamCode)
	endpoint := fmt.Sprintf("teams/%s/matches?status=FINISHED", teamId)
	if matchLimit {
		now := time.Now()
		dateFrom := now.AddDate(0, 0, daysAgo).Format("2006-01-02")
		dateTo := now.AddDate(0, 0, daysAhead).Format("2006-01-02")
		endpoint = fmt.Sprintf("teams/%s/matches?status=FINISHED&dateFrom=%s&dateTo=%s", teamId, dateFrom, dateTo)
	}
	client := api.NewClient()
	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr.Error())
		return
	}

	fmt.Printf("Results for %s\n", teamName)

	header := []string{"Date", "Competition", "Home", "", "", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		rows = append(rows, []string{
			fmt.Sprint(v.UtcDate.Local().Format(dateTimeFormat)),
			v.Competition.Name,
			v.HomeTeam.Name,
			fmt.Sprint(v.Score.FullTime.HomeTeam),
			fmt.Sprint(v.Score.FullTime.AwayTeam),
			v.AwayTeam.Name,
		})
	}
	writer.Table(header, rows)
}
