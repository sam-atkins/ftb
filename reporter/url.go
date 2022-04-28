package reporter

import (
	"fmt"
	"time"
)

func buildLeagueURL(league string) string {
	return fmt.Sprintf("competitions/%s/matches", league)
}

func buildTeamURL(teamId string, matchLimit bool) string {
	if matchLimit {
		now := time.Now()
		dateFrom := now.AddDate(0, 0, daysAgo).Format("2006-01-02")
		dateTo := now.AddDate(0, 0, daysAhead).Format("2006-01-02")
		return fmt.Sprintf("teams/%s/matches?status=FINISHED&dateFrom=%s&dateTo=%s", teamId, dateFrom, dateTo)
	}
	return fmt.Sprintf("teams/%s/matches?status=FINISHED", teamId)
}
