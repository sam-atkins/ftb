package writer

import (
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTable_Render(t *testing.T) {
	type fields struct {
		Header  []string
		Message string
		Rows    [][]string
		Output  io.Writer
	}
	tests := []struct {
		name       string
		fields     fields
		tempPath   string
		goldenPath string
	}{
		{
			name: "Render team results",
			fields: fields{
				Header:  []string{"Date", "Competition", "Home", "", "", "Away"},
				Message: "Results for Liverpool FC",
				Rows: [][]string{
					{"2022 Apr 02 Sat 12:30", "Premier League", "Liverpool FC", "2", "0", "Watford FC"},
					{"2022 Apr 05 Tue 20:00", "UEFA Champions League", "Sport Lisboa e Benfica", "1", "3", "Liverpool FC"},
					{"2022 Apr 10 Sun 16:30", "Premier League", "Manchester City FC", "2", "2", "Liverpool FC"},
					{"2022 Apr 13 Wed 20:00", "UEFA Champions League", "Liverpool FC", "3", "3", "Sport Lisboa e Benfica"},
					{"2022 Apr 19 Tue 20:00", "Premier League", "Liverpool FC", "4", "0", "Manchester United FC"},
					{"2022 Apr 24 Sun 16:30", "Premier League", "Liverpool FC", "2", "0", "Everton FC"},
					{"2022 Apr 27 Wed 20:00", "UEFA Champions League", "Liverpool FC", "2", "0", "Villarreal CF"},
				},
			},
			tempPath:   t.TempDir() + "/team_results.txt",
			goldenPath: "../testdata/team_results.golden",
		},
	}
	for _, tt := range tests {
		file, err := os.OpenFile(tt.tempPath, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			t.Errorf("%s: failed to open file: %v", tt.name, err)
		}
		defer file.Close()
		t.Run(tt.name, func(t *testing.T) {
			tbl := NewTable(tt.fields.Header, tt.fields.Message, tt.fields.Rows)
			tbl.Output = file
			tbl.Render()
			got, err := os.ReadFile(tt.tempPath)
			if err != nil {
				t.Fatal(err)
			}
			want, err := os.ReadFile(tt.goldenPath)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(want, got) {
				t.Fatal(cmp.Diff(want, got))
			}
		})
	}
}
