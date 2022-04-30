package writer

import (
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

type Table struct {
	Header             []string
	Message            string
	Rows               [][]string
	TeamLeaguePosition int
	Output             io.Writer
}

// NewTable returns a Table with a default of os.Stdout as the output
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

func (t *Table) Render() {
	fmt.Println(t.Message)
	w := tablewriter.NewWriter(t.Output)
	w.SetHeader(t.Header)
	w.AppendBulk(t.Rows)
	w.Render()
}

func (t *Table) RenderWithTeamHighlight() {
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
