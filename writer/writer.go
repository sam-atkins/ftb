package writer

import (
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

// TODO:
// take arg message and print this before rendering the table
// Table struct, rename
type Tables struct {
	Header             []string
	Message            string
	Rows               [][]string
	TeamLeaguePosition int
	Output             io.Writer
}

func NewTable(header []string, message string, rows [][]string) *Tables {
	return &Tables{
		Header:  header,
		Message: message,
		Rows:    rows,
		Output:  os.Stdout,
	}
}

func NewTableWithPositionHighlight(header []string, message string, rows [][]string, teamLeaguePosition int) *Tables {
	return &Tables{
		Header:             header,
		Message:            message,
		Rows:               rows,
		Output:             os.Stdout,
		TeamLeaguePosition: teamLeaguePosition,
	}
}

func (t *Tables) RenderTable() {
	fmt.Println(t.Message)
	w := tablewriter.NewWriter(t.Output)
	w.SetHeader(t.Header)
	w.AppendBulk(t.Rows)
	w.Render()
}

func (t *Tables) RenderTableWithTeamHighlight() {
	fmt.Println(t.Message)
	w := tablewriter.NewWriter(t.Output)
	w.SetHeader(t.Header)
	w.SetColumnAlignment([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})

	for i, row := range t.Rows {
		if i == t.TeamLeaguePosition {
			w.Rich(row, []tablewriter.Colors{
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
		w.Append(row)
	}
	w.Render()
}

// func to create and return a NewTable
// tablewriter.NewWriter(os.Stdout) is the default, override for tests

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
