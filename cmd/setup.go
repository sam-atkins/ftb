/*
Copyright © 2022 Sam Atkins <samatkins@hey.com>
MIT License
*/
package cmd

import (
	"fmt"

	"github.com/sam-atkins/ftb/reporter"
	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Adds team config",
	Long: `Adds team config for the leagues saved in the config package to the teams
config file. Flags can amend the operation. For example:

ftb setup --reset
	will reset the existing teams config file. This is useful at the beginning of a
	new season to ensure the teams are accurate.

ftb setup --debug
	switches on additional logging to the console.
`,
	Run: func(cmd *cobra.Command, args []string) {
		reset, _ := cmd.Flags().GetBool("reset")
		if reset {
			fmt.Println("Resetting the teams config file")
			reporter.GetTeamsConfig(true, false)
		} else {
			reporter.GetTeamsConfig(false, false)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	setupCmd.Flags().BoolP("debug", "d", false, "Set debug on")
	setupCmd.Flags().BoolP("reset", "r", false, "Reset the teams config file")
}
