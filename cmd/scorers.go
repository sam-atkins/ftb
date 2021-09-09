/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"github.com/sam-atkins/ftb/broker"
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
		broker.GetScorers(league)
	},
}

func init() {
	rootCmd.AddCommand(scorersCmd)
	scorersCmd.Flags().StringP("league", "l", "", "The league to show e.g. PL, BL1")
}
