package writer

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Table writes a table to standard out
func Table(header []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(rows)
	table.Render()
}

// TableWithTeamHighlight prints a league table to stdout, highlighting the team
// at position teamIndex
func TableWithTeamHighlight(teamIndex int, header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetColumnAlignment([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})

	for i, row := range data {
		if i == teamIndex {
			table.Rich(row, []tablewriter.Colors{
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
				{tablewriter.Bold, tablewriter.FgWhiteColor},
			})
			// NOTE: important to `continue` here otherwise the highlighted team is
			// appended twice
			continue
		}
		table.Append(row)
	}
	table.Render()
}
