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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		header := []string{"Date", "Home", "Goals", "Away", "Goals"}
		var rows [][]string
		for _, v := range response.Body.Matches {
			if v.Season.CurrentMatchday-1 == v.Matchday {
				rows = append(rows, []string{
					fmt.Sprint(v.UtcDate.Format("2006 January 02")),
					v.HomeTeam.Name,
					fmt.Sprint(v.Score.FullTime.HomeTeam),
					v.AwayTeam.Name,
					fmt.Sprint(v.Score.FullTime.AwayTeam),
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
