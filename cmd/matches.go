/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
	"github.com/spf13/cobra"
)

// matchesCmd represents the matches command
var matchesCmd = &cobra.Command{
	Use:   "matches",
	Short: "Shows the fixtures for the next match day",
	Long: `Shows the fixtures for the next match day.
For example, to show matches for a league:

ftb matches --league PL
ftb matches -l BL1


For example, to show matches for a club:

ftb results --team FCB
ftb results --team fcb
ftb results -t LIV
ftb results -t liv
`,
	Run: func(cmd *cobra.Command, args []string) {
		league, _ := cmd.Flags().GetString("league")
		if league != "" {
			matchesByLeague(league)
			return
		}

		team, _ := cmd.Flags().GetString("team")
		if team != "" {
			matchesByTeam(team)
			return
		}

		fmt.Print("No flag provided. Check the below help menu for options.\n\n")
		helpErr := cmd.Help()
		if helpErr != nil {
			os.Exit(1)
		}
		os.Exit(1)
	},
}

func matchesByLeague(league string) {
	client := api.Client{}
	endpoint := fmt.Sprintf("competitions/%s/matches", league)

	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
	}

	fmt.Printf("Next match day fixtures in the %v\n", response.Body.Competition.Name)

	header := []string{"Date", "Home", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		if v.Season.CurrentMatchday == v.Matchday {
			rows = append(rows, []string{
				fmt.Sprint(v.UtcDate.Local().Format("2006 January 02 15:04")),
				v.HomeTeam.Name,
				v.AwayTeam.Name,
			})
		}
	}
	writer.Table(header, rows)
}

func matchesByTeam(teamCode string) {
	var teamId string
	var teamName string
	for _, v := range config.TeamConfig {
		if v.Code == strings.ToUpper(teamCode) {
			teamId = v.Id
			teamName = v.Name
		}
	}

	// TODO(sam) handover to fn or cmd that lists team config?
	if teamId == "" {
		fmt.Println("Did not recognise that team")
		os.Exit(1)
	}

	client := api.Client{}
	endpoint := fmt.Sprintf("teams/%s/matches?status=SCHEDULED", teamId)

	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
	}

	fmt.Printf("Results for %s\n", teamName)

	header := []string{"Date", "Competition", "Home", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		rows = append(rows, []string{
			fmt.Sprint(v.UtcDate.Format("2006 January 02")),
			v.Competition.Name,
			v.HomeTeam.Name,
			v.AwayTeam.Name,
		})
	}
	writer.Table(header, rows)
}

func init() {
	rootCmd.AddCommand(matchesCmd)
	matchesCmd.Flags().StringP("league", "l", "", "The league to show e.g. PL, BL1")
	matchesCmd.Flags().StringP("team", "t", "", "The team to show results for e.g. FCB, LIV")
}
