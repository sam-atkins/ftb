package reporter

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

func ResultsCLI(league, team string) {
	fmt.Printf("Results for league: %s\n", league)
	fmt.Printf("Results for team: %s\n", team)
	// cmd should ensure both league and team are not empty but handle here too
	if league == "" && team == "" {
		fmt.Println("No league or team specified")
		os.Exit(1)
	}

	if league != "" {
		handleResultsByLeague(league)
	}

	if team != "" {
		fmt.Println("Not yet implemented")
		os.Exit(1)
		// TODO
		// handleResultsByTeam(team)
	}
}

type results struct {
	// client *api.Client ?
	league   string
	team     string
	endpoint string
	message  string
	header   []string
	rows     [][]string
}

func handleResultsByLeague(league string) {
	r := resultsLeague(league)

	// should return what I want in order to send to the writer: message, header and rows
	r.getResultsByLeague()

	writer.Table(r.header, r.rows)
}

// ResultsByLeague wrapper on the Results struct for league results
func resultsLeague(league string) *results {
	return &results{
		league: league,
	}
}

// resultsByTeam wrapper on the Results struct for team results
func resultsTeam(team string) *results {
	return &results{}
}

func (r *results) getResultsByLeague() *results {
	r.endpoint = buildLeagueURL(r.league)
	r.message = fmt.Sprintf("Results from the %s", r.league)
	r.header = []string{"Date", "Home", "", "", "Away"}

	results, err := fetchResultsByLeague(r.endpoint)
	if err != nil {
		log.Printf("Something went wrong: %s\n", err)
		os.Exit(1)
	}
	r.rows = results

	return r
}

func buildLeagueURL(league string) string {
	return fmt.Sprintf("competitions/%s/matches", league)
}

func fetchResultsByLeague(endpoint string) ([][]string, error) {
	client := api.NewClient()
	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		// log.Printf("Something went wrong with the request: %s\n", responseErr)
		return nil, responseErr
	}
	result := buildResultsByLeagueRows(response)
	return result, nil
}

// make the response struct public as interim step
func buildResultsByLeagueRows(response *api.ApiMatchesResponse) [][]string {
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
	return rows
}

// TODO delete
// ResultsByLeague gets the results for a football league and pretty prints
// the league table to stdout
// func ResultsByLeague(league string) {
// 	endpoint := fmt.Sprintf("competitions/%s/matches", league)
// 	client := api.NewClient()
// 	response, responseErr := client.GetMatches(endpoint)
// 	if responseErr != nil {
// 		log.Printf("Something went wrong with the request: %s\n", responseErr)
// 		return
// 	}

// 	fmt.Printf("Results from the %v\n", response.Body.Competition.Name)

// 	header := []string{"Date", "Home", "", "", "Away"}
// 	var rows [][]string
// 	for _, v := range response.Body.Matches {
// 		if v.Season.CurrentMatchday-1 == v.Matchday && v.Status == "FINISHED" ||
// 			v.Season.CurrentMatchday == v.Matchday && v.Status == "FINISHED" {
// 			rows = append(rows, []string{
// 				fmt.Sprint(v.UtcDate.Local().Format(dateTimeFormat)),
// 				v.HomeTeam.Name,
// 				fmt.Sprint(v.Score.FullTime.HomeTeam),
// 				fmt.Sprint(v.Score.FullTime.AwayTeam),
// 				v.AwayTeam.Name,
// 			})
// 		}
// 	}
// 	writer.Table(header, rows)
// }

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
