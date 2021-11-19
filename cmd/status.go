/*
Copyright © 2020 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"fmt"

	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/reporter"
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
			reporter.ResultsByTeam(team, true)
			fmt.Println("")
			reporter.MatchesByTeam(team, true)
			fmt.Println("")
			reporter.GetTableForTeam(team)
			return
		}
		config.CodeNotFound()
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringP("team", "t", "", "The team to show results for e.g. FCB, LIV")
}
