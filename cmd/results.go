/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"fmt"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/writer"
	"github.com/spf13/cobra"
)

// resultsCmd represents the results command
var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Shows the results for the previous match day",
	Long: `Shows the results for the previous match day.
For example:

ftb results --league PL
ftb results -l BL1
`,
	Run: func(cmd *cobra.Command, args []string) {
		league, _ := cmd.Flags().GetString("league")
		if league == "" {
			// TODO(sam) add default league to config
			league = "BL1"
		}

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
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)
	resultsCmd.Flags().StringP("league", "l", "", "The league to show e.g. PL, BL1")
}
