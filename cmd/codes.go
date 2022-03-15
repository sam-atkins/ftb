/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"fmt"

	"github.com/sam-atkins/ftb/config"
	"github.com/sam-atkins/ftb/writer"
	"github.com/spf13/cobra"
)

// codesCmd represents the codes command
var codesCmd = &cobra.Command{
	Use:   "codes",
	Short: "Displays league and team codes",
	Long: `Displays league and team codes. These codes are used as flags to
other commands e.g.
	ftb matches --league PL
	ftb scorers -l BL1
	ftb results --team FCB
	ftb status --team fcb
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("These are the available league codes:")
		headerLeagues := []string{"league", "Code", "Country"}
		leagueCodes := config.GetLeagueCodes()
		writer.Table(headerLeagues, leagueCodes)
		fmt.Println()
		fmt.Println("These are the available team codes:")
		headerClubs := []string{"Team", "Team Code", "Country"}
		teamCodes := config.GetTeamCodesFromConfig()
		writer.Table(headerClubs, teamCodes)
	},
}

func init() {
	rootCmd.AddCommand(codesCmd)
}
