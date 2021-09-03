/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/sam-atkins/ftb/api"
	"github.com/spf13/cobra"
)

// standingsCmd represents the league command
var standingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "Prints the league table",
	Long: `Prints the league table per the user input
For example:

	// print the English Premier League table
	ftb standings --league PL

	// print the German Bundesliga
	ftb standings -l BL1
`,
	Run: func(cmd *cobra.Command, args []string) {
		league, _ := cmd.Flags().GetString("league")
		if league == "" {
			// TODO(sam) add default league to config
			league = "BL1"
		}

		endpoint := fmt.Sprintf("competitions/%s/standings", league)

		client := api.Client{}
		response, responseErr := client.DoRequest(endpoint)
		if responseErr != nil {
			fmt.Printf("Something went wrong with the request %s", responseErr)
		}

		fmt.Printf("Table for League: %v\n", response.Body.Competition.Name)

		var rows [][]string
		for _, v := range response.Body.Standings[0].Table {
			rows = append(rows, []string{
				fmt.Sprint(v.Position),
				v.Team.Name,
				fmt.Sprint(v.PlayedGames),
				fmt.Sprint(v.Won),
				fmt.Sprint(v.Draw),
				fmt.Sprint(v.Lost),
				fmt.Sprint(v.GoalDifference),
				fmt.Sprint(v.Points),
			})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Pos", "Team", "Played", "Won", "Draw", "Lost", "GD", "Points"})
		table.AppendBulk(rows)
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(standingsCmd)
	standingsCmd.Flags().StringP("league", "l", "", "The league to show e.g. PL, BL1")
}
