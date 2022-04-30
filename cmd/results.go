package cmd

import (
	"fmt"
	"os"

	"github.com/sam-atkins/ftb/reporter"
	"github.com/spf13/cobra"
)

// resultsCmd represents the results command
var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Shows football match results",
	Long: `Shows football match results for a league's previous match day or for a club.
For example, to show league results:

ftb results --league PL
ftb results -l BL1

For example, to show results for a club:

ftb results --team FCB
ftb results --team fcb
ftb results -t LIV
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
		reporter.ResultsCLI(league, team, false)
	},
}

func init() {
	rootCmd.AddCommand(resultsCmd)
	resultsCmd.Flags().StringVarP(&league, "league", "l", "", "The league to show e.g. PL, BL1")
	resultsCmd.Flags().StringVarP(&team, "team", "t", "", "The team to show results for e.g. FCB, LIV")
}
