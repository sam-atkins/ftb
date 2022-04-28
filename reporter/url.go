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

type teamURL struct {
	now time.Time
}

// newTeamURL is a wrapper on the teamURL struct for creating a new teamURL
func newTeamURL() *teamURL {
	return &teamURL{now: time.Now()}
}

// teamFinishedMatches builds the URL for a team's finished matches. The matchLimit arg
// adds a query string to limit the results
func (t *teamURL) teamFinishedMatches(teamId string, matchLimit bool) string {
	if matchLimit {
		dateFrom := t.now.AddDate(0, 0, daysAgo).Format("2006-01-02")
		dateTo := t.now.AddDate(0, 0, daysAhead).Format("2006-01-02")
		return fmt.Sprintf("teams/%s/matches?status=FINISHED&dateFrom=%s&dateTo=%s", teamId, dateFrom, dateTo)
	}
	return fmt.Sprintf("teams/%s/matches?status=FINISHED", teamId)
}
