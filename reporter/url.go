package reporter

import (
	"fmt"
	"time"
)

const (
	daysAgo   = -28
	daysAhead = 28
)

func buildLeagueURL(league string) string {
	return fmt.Sprintf("competitions/%s/matches", league)
}

func buildLeagueStandingsURL(leagueCode string) string {
	return fmt.Sprintf("competitions/%s/standings", leagueCode)
}

func scorersURL(league string) string {
	return fmt.Sprintf("competitions/%s/scorers", league)
}

type teamURL struct {
	now time.Time
}

// newTeamURL is a wrapper on the teamURL struct for creating a new teamURL
func newTeamURL() *teamURL {
	return &teamURL{now: time.Now()}
}

// teamMatches builds the URL for a team's finished or scheduled matches. The matchLimit
// arg adds a query string to limit the results. MatchComplete adds to the query string
// to filter results by finished or scheduled matches.
func (t *teamURL) teamMatches(teamId string, matchLimit bool, matchComplete bool) string {
	matchStatus := "SCHEDULED"
	if matchComplete {
		matchStatus = "FINISHED"
	}
	if matchLimit {
		dateFrom := t.now.AddDate(0, 0, daysAgo).Format("2006-01-02")
		dateTo := t.now.AddDate(0, 0, daysAhead).Format("2006-01-02")
		return fmt.Sprintf("teams/%s/matches?status=%s&dateFrom=%s&dateTo=%s", teamId, matchStatus, dateFrom, dateTo)
	}
	return fmt.Sprintf("teams/%s/matches?status=%s", teamId, matchStatus)
}
