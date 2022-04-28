package reporter

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// ResultsCLI is the entrypoint for the reporter to get results for a league or team
func ResultsCLI(league, team string, matchLimit bool) {
	switch {
	case league != "":
		handleResultsByLeague(league)
	case team != "":
		handleResultsByTeam(team, matchLimit)
	default:
		// cmd should ensure both league and team are not empty but handle here too
		fmt.Println("No league or team specified")
		os.Exit(1)
	}
}

type results struct {
	// client *api.Client ?
	endpoint   string
	header     []string
	league     string
	matchLimit bool
	message    string
	rows       [][]string
	teamCode   string
	teamId     string
	teamName   string
}

func handleResultsByLeague(league string) {
	r := resultsLeague(league)
	r.getResultsByLeague()
	writer.NewTable(r.header, r.message, r.rows).Render()
}

func handleResultsByTeam(team string, matchLimit bool) {
	r := resultsTeam(team, matchLimit)
	r.getResultsByTeam()
	w := writer.NewTable(r.header, r.message, r.rows)
	w.Render()
}

// ResultsByLeague wrapper on the Results struct for league results
func resultsLeague(league string) *results {
	return &results{
		league: league,
	}
}

// resultsByTeam wrapper on the Results struct for team results
func resultsTeam(team string, matchLimit bool) *results {
	teamCode := strings.ToUpper(team)
	return &results{
		matchLimit: matchLimit,
		teamCode:   teamCode,
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
	r.endpoint = newTeamURL().teamFinishedMatches(r.teamId, r.matchLimit)
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
