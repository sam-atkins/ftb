package cmd

import (
	"fmt"
	"os"

	"github.com/sam-atkins/ftb/reporter"
	"github.com/spf13/cobra"
)

var leagueTable string

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
		if leagueTable == "" {
			fmt.Print("No flag provided. Check the below help menu for options.\n\n")
			helpErr := cmd.Help()
			if helpErr != nil {
				os.Exit(1)
			}
			os.Exit(1)
		}
		reporter.TableCLI(leagueTable)
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
	tableCmd.Flags().StringVarP(&leagueTable, "league", "l", "", "The league to show e.g. PL, BL1")
}
