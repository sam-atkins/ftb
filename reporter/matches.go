package reporter

import (
	"fmt"
	"log"
	"os"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
)

// MatchesCLI is the entrypoint for the reporter to get results for a league or team
func MatchesCLI(league, team string, matchLimit bool) {
	switch {
	case league != "":
		handleMatchesByLeague(league, matchLimit)
	case team != "":
		handleMatchesByTeam(team, matchLimit)
	default:
		// cmd should ensure both league and team are not empty but handle here too
		fmt.Println("No league or team specified")
		os.Exit(1)
	}
}

type matches struct {
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

func matchesLeague(league string, matchLimit bool) *matches {
	return &matches{
		league:     league,
		matchLimit: matchLimit,
	}
}

func matchesTeam(teamCode string, matchLimit bool) *matches {
	cfg := config.NewTeamConfig(teamCode)
	return &matches{
		teamCode:   teamCode,
		matchLimit: matchLimit,
		teamId:     cfg.TeamId,
		teamName:   cfg.TeamName,
	}
}

func handleMatchesByLeague(league string, matchLimit bool) {
	m := matchesLeague(league, matchLimit)
	m.getMatchesByLeague()
	writer.NewTable(m.header, m.message, m.rows).Render()
}

func (m *matches) getMatchesByLeague() *matches {
	m.endpoint = buildLeagueURL(m.league)
	m.header = []string{"Date", "Home", "", "", "Away", "Match Status"}
	response, err := fetchMatchesByLeague(m.endpoint)
	if err != nil {
		log.Printf("Something went wrong with the request: %s\n", err)
		os.Exit(1)
	}
	m.message = fmt.Sprintf("Current match day fixtures in the %v\n", response.Body.Competition.Name)
	m.rows = buildRowsForCurrentMatchDayFixtures(response)
	return m
}

func fetchMatchesByLeague(endpoint string) (*api.ApiMatchesResponse, error) {
	client := api.NewClient()
	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func buildRowsForCurrentMatchDayFixtures(response *api.ApiMatchesResponse) [][]string {
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
	return rows
}

func handleMatchesByTeam(team string, matchLimit bool) {
	m := matchesTeam(team, matchLimit)
	m.getMatchesByTeam()
	writer.NewTable(m.header, m.message, m.rows).Render()
}

func (m *matches) getMatchesByTeam() *matches {
	m.endpoint = newTeafmURL().teamMatches(m.teamId, m.matchLimit, false)
	m.header = []string{"Date", "Competition", "Home", "Away"}
	response, err := fetchMatchesByTeam(m.endpoint)
	if err != nil {
		log.Printf("Something went wrong with the request: %s\n", err)
		os.Exit(1)
	}
	m.message = fmt.Sprintf("Matches for %s\n", m.teamName)
	m.rows = buildRowsForTeamFixtures(response)
	return m
}

func fetchMatchesByTeam(endpoint string) (*api.ApiMatchesResponse, error) {
	client := api.NewClient()
	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		return nil, responseErr
	}
	return response, nil
}

func buildRowsForTeamFixtures(response *api.ApiMatchesResponse) [][]string {
	var rows [][]string
	for _, v := range response.Body.Matches {
		rows = append(rows, []string{
			fmt.Sprint(v.UtcDate.Local().Format(dateTimeFormat)),
			v.Competition.Name,
			v.HomeTeam.Name,
			v.AwayTeam.Name,
		})
	}
	return rows
}
