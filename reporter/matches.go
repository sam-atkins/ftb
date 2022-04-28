package reporter

import (
	"fmt"
	"log"
	"time"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// MatchesByLeague fetches matches for a league and prints to stdout
func MatchesByLeague(league string) {
	endpoint := fmt.Sprintf("competitions/%s/matches", league)
	client := api.NewClient()
	response, responseErr := client.GetMatches(endpoint)

	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}

	header := []string{"Date", "Home", "", "", "Away", "Match Status"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		if v.Season.CurrentMatchday == v.Matchday {
			rows = append(rows, []string{
				fmt.Sprint(v.UtcDate.Local().Format(dateTimeFormat)),
				v.HomeTeam.Name,
				formatFloatOrNil(v.Score.FullTime.HomeTeam),
				formatFloatOrNil(v.Score.FullTime.AwayTeam),
				v.AwayTeam.Name,
				convertToTitle(v.Status),
			})
		}
	}
	message := fmt.Sprintf("Current match day fixtures in the %v\n", response.Body.Competition.Name)
	writer.NewTable(header, message, rows).Render()
}

// MatchesByTeam fetches matches for a team and prints to stdout. Arg matchLimit limits
// the results to the next three weeks
func MatchesByTeam(teamCode string, matchLimit bool) {
	_, teamName, teamId := config.GetTeamInfoFromUserTeamCode(teamCode)
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
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}
	message := fmt.Sprintf("Matches for %s\n", teamName)
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
	writer.NewTable(header, message, rows).Render()
}
