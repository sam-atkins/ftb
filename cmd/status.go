/*
Copyright © 2020 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"fmt"

	"github.com/sam-atkins/ftb/broker"
	"github.com/sam-atkins/ftb/config"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Provides previous and next matches and the league table for a club",
	Long: `Provides previous and next matches and the league table for a club
For example:

ftb status -t fcb

ftb status -t LIV
`,
	Run: func(cmd *cobra.Command, args []string) {
		team, _ := cmd.Flags().GetString("team")
		if team != "" {
			broker.ResultsByTeam(team, true)
			fmt.Println("")
			broker.MatchesByTeam(team, true)
			fmt.Println("")
			broker.GetTableForTeam(team)
			return
		}
		config.CodeNotFound()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringP("team", "t", "", "The team to show results for e.g. FCB, LIV")
}
