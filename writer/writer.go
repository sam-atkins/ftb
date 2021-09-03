/*
Copyright © 2021 Sam Atkins <samatkins@hey.com>
MIT License
*/
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
