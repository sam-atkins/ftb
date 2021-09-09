/*
Copyright © 2020 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/sam-atkins/ftb/broker"
	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
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

		fmt.Println("Did not recognise that team. These are the available team codes:")
		header := []string{"Team", "Code"}
		teamCodes := config.GetTeamCodes()
		writer.Table(header, teamCodes)
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringP("team", "t", "", "The team to show results for e.g. FCB, LIV")
}
