package cmd

import (
	"fmt"
	"os"

	"github.com/sam-atkins/ftb/reporter"
	"github.com/spf13/cobra"
)

// matchesCmd represents the matches command
var matchesCmd = &cobra.Command{
	Use:   "matches",
	Short: "Shows the fixtures for the next match day",
	Long: `Shows the fixtures for the next match day.
For example, to show matches for a league:

ftb matches --league PL
ftb matches -l BL1

For example, to show matches for a club:

ftb results --team FCB
ftb results -t liv
`,
	Run: func(cmd *cobra.Command, args []string) {
		league, _ := cmd.Flags().GetString("league")
		if league != "" {
			reporter.MatchesByLeague(league)
			return
		}

		team, _ := cmd.Flags().GetString("team")
		if team != "" {
			reporter.MatchesByTeam(team, false)
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

func init() {
	rootCmd.AddCommand(matchesCmd)
	matchesCmd.Flags().StringP("league", "l", "", "The league to show e.g. PL, BL1")
	matchesCmd.Flags().StringP("team", "t", "", "The team to show results for e.g. FCB, LIV")
}
