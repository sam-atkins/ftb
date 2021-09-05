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

// resultsCmd represents the results command
var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Shows football match results",
	Long: `Shows football match results for a league's previous match day or for a club.
For example, to show league results:

ftb results --league PL
ftb results -l BL1

For example, to show results for a club:

ftb results --team FCB
ftb results --team fcb
ftb results -t LIV
ftb results -t liv
`,
	Run: func(cmd *cobra.Command, args []string) {
		league, _ := cmd.Flags().GetString("league")
		if league != "" {
			resultsByLeague(league)
			return
		}

		team, _ := cmd.Flags().GetString("team")
		if team != "" {
			resultsByTeam(team)
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

func resultsByLeague(league string) {
	client := api.Client{}
	endpoint := fmt.Sprintf("competitions/%s/matches", league)

	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
	}

	fmt.Printf("Results from the %v\n", response.Body.Competition.Name)

	header := []string{"Date", "Home", "", "", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		if v.Season.CurrentMatchday-1 == v.Matchday {
			rows = append(rows, []string{
				fmt.Sprint(v.UtcDate.Format("2006 January 02")),
				v.HomeTeam.Name,
				fmt.Sprint(v.Score.FullTime.HomeTeam),
				fmt.Sprint(v.Score.FullTime.AwayTeam),
				v.AwayTeam.Name,
			})
		}
	}
	writer.Table(header, rows)
}

func resultsByTeam(teamCode string) {
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
	endpoint := fmt.Sprintf("teams/%s/matches?status=FINISHED", teamId)

	response, responseErr := client.GetMatches(endpoint)
	if responseErr != nil {
		fmt.Printf("Something went wrong with the request %s", responseErr)
	}

	fmt.Printf("Results for %s\n", teamName)

	header := []string{"Date", "Home", "", "", "Away"}
	var rows [][]string
	for _, v := range response.Body.Matches {
		rows = append(rows, []string{
			fmt.Sprint(v.UtcDate.Format("2006 January 02")),
			v.HomeTeam.Name,
			fmt.Sprint(v.Score.FullTime.HomeTeam),
			fmt.Sprint(v.Score.FullTime.AwayTeam),
			v.AwayTeam.Name,
		})
	}
	writer.Table(header, rows)
}

func init() {
	rootCmd.AddCommand(resultsCmd)
	resultsCmd.Flags().StringP("league", "l", "", "The league to show e.g. PL, BL1")
	resultsCmd.Flags().StringP("team", "t", "", "The team to show results for e.g. FCB, LIV")
}
