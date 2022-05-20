package cmd

import (
	"fmt"
	"os"

	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/reporter"
	"github.com/spf13/cobra"
)

var teamStatus string

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
		if teamStatus == "" {
			config.CodeNotFound()
		}
		reporter.ResultsCLI("", teamStatus, true)
		fmt.Println("")
		reporter.MatchesCLI("", teamStatus, false)
		fmt.Println("")
		reporter.TeamTableCLI(teamStatus)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().StringVarP(&teamStatus, "team", "t", "", "The team to show results for e.g. FCB, LIV")
}
