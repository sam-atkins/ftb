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
type Table struct {
	Header             []string
	Message            string
	Rows               [][]string
	TeamLeaguePosition int
	Output             io.Writer
}

func NewTable(header []string, message string, rows [][]string) *Table {
	return &Table{
		Header:  header,
		Message: message,
		Rows:    rows,
		Output:  os.Stdout,
	}
}

func NewTableWithPositionHighlight(header []string, message string, rows [][]string, teamLeaguePosition int) *Table {
	return &Table{
		Header:             header,
		Message:            message,
		Rows:               rows,
		Output:             os.Stdout,
		TeamLeaguePosition: teamLeaguePosition,
	}
}

func (t *Table) RenderTable() {
	fmt.Println(t.Message)
	w := tablewriter.NewWriter(t.Output)
	w.SetHeader(t.Header)
	w.AppendBulk(t.Rows)
	w.Render()
}

func (t *Table) RenderTableWithTeamHighlight() {
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
