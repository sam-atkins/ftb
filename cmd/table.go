/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"github.com/sam-atkins/ftb/broker"
	"github.com/spf13/cobra"
)

// tableCmd represents the league command
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Prints the league table",
	Long: `Prints the league table per the user input
For example:

	// print the English Premier League table
	ftb table --league PL

	// print the German Bundesliga
	ftb table -l BL1

	// print the Spanish La Liga (Primera Division)
	ftb table --league PD
`,
	Run: func(cmd *cobra.Command, args []string) {
		league, _ := cmd.Flags().GetString("league")
		if league == "" {
			// TODO(sam) add default league to config
			league = "BL1"
		}
		broker.GetTable(league)
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
	tableCmd.Flags().StringP("league", "l", "", "The league to show e.g. PL, BL1")
}
