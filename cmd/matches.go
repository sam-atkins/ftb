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
		if league == "" && team == "" {
			fmt.Print("No flag provided. Check the below help menu for options.\n\n")
			helpErr := cmd.Help()
			if helpErr != nil {
				os.Exit(1)
			}
			os.Exit(1)
		}
		if league != "" {
			reporter.MatchesCLI(league, team, false)
			return
		}

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
	matchesCmd.Flags().StringVarP(&league, "league", "l", "", "The league to show e.g. PL, BL1")
	matchesCmd.Flags().StringVarP(&team, "team", "t", "", "The team to show results for e.g. FCB, LIV")
}
