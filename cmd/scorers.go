package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/sam-atkins/ftb/reporter"
	"github.com/spf13/cobra"
)

var leagueScorers string

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
		if leagueScorers == "" {
			fmt.Print("No flag provided. Check the below help menu for options.\n\n")
			helpErr := cmd.Help()
			if helpErr != nil {
				os.Exit(1)
			}
			os.Exit(1)
		}
		reporter.ScorersCLI(strings.ToUpper(leagueScorers))
	},
}

func init() {
	rootCmd.AddCommand(scorersCmd)
	scorersCmd.Flags().StringVarP(&leagueScorers, "league", "l", "", "The league to show e.g. PL, BL1")
}
