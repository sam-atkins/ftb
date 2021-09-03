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

// scorersCmd represents the scorers command
var scorersCmd = &cobra.Command{
	Use:   "scorers",
	Short: "Get the top scorers in a league",
	Long: `Get the top scorers in a league.
For example:

// top scorers in the Bundesliga
ftb scorers -l BL1

// top scorers in the Premier League
ftb scorers --league PL
`,
	Run: func(cmd *cobra.Command, args []string) {
		league, _ := cmd.Flags().GetString("league")
		if league == "" {
			// TODO(sam) add default league to config
			league = "BL1"
		}
		client := api.Client{}
		endpoint := fmt.Sprintf("competitions/%s/scorers", league)

		response, responseErr := client.GetScorers(endpoint)
		if responseErr != nil {
			fmt.Printf("Something went wrong with the request %s", responseErr)
		}

		fmt.Printf("Top Scorers in the %v\n", response.Body.Competition.Name)

		header := []string{"Name", "Team", "Goals"}
		var rows [][]string
		for _, v := range response.Body.Scorers {
			rows = append(rows, []string{
				v.Player.Name,
				v.Team.Name,
				fmt.Sprint(v.NumberOfGoals),
			})
		}
		writer.Table(header, rows)
	},
}

func init() {
	rootCmd.AddCommand(scorersCmd)
	scorersCmd.Flags().StringP("league", "l", "", "The league to show e.g. PL, BL1")
}
