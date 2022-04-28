package reporter

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// ResultsCLI is the entrypoint for the reporter to get results for a league or team
func ResultsCLI(league, team string) {
	switch {
	case league != "":
		handleResultsByLeague(league)
	case team != "":
		handleResultsByTeam(team)
	default:
		// cmd should ensure both league and team are not empty but handle here too
		fmt.Println("No league or team specified")
		os.Exit(1)
	}
}

type results struct {
	// client *api.Client ?
	league   string
	teamCode string
	teamId   string
	teamName string
	endpoint string
	message  string
	header   []string
	rows     [][]string
}

func handleResultsByLeague(league string) {
	r := resultsLeague(league)
	r.getResultsByLeague()
	// TODO:
	// pass in r.message for printing
	writer.Table(r.header, r.rows)
}

func handleResultsByTeam(team string) {
	r := resultsTeam(team)
	r.getResultsByTeam()
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
	teamCode := strings.ToUpper(team)
	return &results{
		teamCode: teamCode,
	}
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

func fetchResultsByLeague(endpoint string) ([][]string, error) {
	// TODO: move client to struct so in tests can inject test client
	client := api.NewClient()
	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
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

func (r *results) getResultsByTeam() *results {
	_, r.teamName, r.teamId = config.GetTeamInfoFromUserTeamCode(r.teamCode)
	r.endpoint = newTeamURL().teamFinishedMatches(r.teamId, true)
	r.message = fmt.Sprintf("Results for %s", r.teamName)
	r.header = []string{"Date", "Competition", "Home", "", "", "Away"}

	results, err := fetchResultsByTeam(r.endpoint)
	if err != nil {
		log.Printf("Something went wrong: %s\n", err)
		os.Exit(1)
	}
	r.rows = results
	return r
}

func fetchResultsByTeam(endpoint string) ([][]string, error) {
	client := api.NewClient()
	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	result := buildResultsByTeamRows(response)
	return result, nil
}

func buildResultsByTeamRows(response *api.ApiMatchesResponse) [][]string {
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
	return rows
}

// ResultsByTeam fetches results for a team and prints to stdout. Arg matchLimit limits
// the results to the previous three weeks
// TODO: used in the status cmd so cannot delete yet
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
