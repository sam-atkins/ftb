package config

import (
	"fmt"
	"os"

	"github.com/sam-atkins/ftb/writer"
)

// GetTeamCodes returns the teams and their codes
func GetTeamCodes() [][]string {
	var teamCodes [][]string
	for _, v := range TeamConfig {
		teamCodes = append(teamCodes, []string{
			v.Name,
			v.Code,
			v.LeagueCode,
		},
		)
	}
	return teamCodes
}

// GetLeagueCodes returns the leagues and their codes
func GetLeagueCodes() [][]string {
	var leagueCodes [][]string
	for _, v := range LeagueConfig {
		leagueCodes = append(leagueCodes, []string{
			v.LeagueName,
			v.LeagueCode,
		},
		)
	}
	return leagueCodes
}

// CodeNotFound used when the user enters an unknown flag code. It prints the available
// codes to stdout and exits (1)
func CodeNotFound() {
	fmt.Println("Did not recognise that team. These are the available team codes:")
	header := []string{"Team", "Code"}
	teamCodes := GetTeamCodes()
	writer.Table(header, teamCodes)
	os.Exit(1)
}
