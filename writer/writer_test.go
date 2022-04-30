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

func TestTableWithPositionHighlight_Render(t *testing.T) {
	type fields struct {
		Header             []string
		Message            string
		Rows               [][]string
		TeamLeaguePosition int
		Output             io.Writer
	}
	tests := []struct {
		name       string
		fields     fields
		tempPath   string
		goldenPath string
	}{
		{
			name: "Render table with highlighted team",
			fields: fields{
				Header:  []string{"Pos", "Team", "Played", "Won", "Draw", "Lost", "+", "-", "GD", "Points"},
				Message: "League table: Bundesliga",
				Rows: [][]string{
					{"1", "FC Bayern München", "31", "24", "3", "4", "92", "30", "62", "75"},
					{"2", "Borussia Dortmund", "31", "20", "3", "8", "77", "46", "31", "63"},
					{"3", "Bayer Leverkusen", "31", "16", "7", "8", "72", "44", "28", "55"},
					{"4", "RB Leipzig", "31", "19", "16", "6", "9", "66", "33", "33", "54"},
				},
				TeamLeaguePosition: 0,
			},
			tempPath:   t.TempDir() + "/team_results.txt",
			goldenPath: "../testdata/team_table_bundesliga.golden",
		},
	}
	for _, tt := range tests {
		file, err := os.OpenFile(tt.tempPath, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			t.Errorf("%s: failed to open file: %v", tt.name, err)
		}
		defer file.Close()
		t.Run(tt.name, func(t *testing.T) {
			tbl := NewTableWithPositionHighlight(tt.fields.Header, tt.fields.Message, tt.fields.Rows, tt.fields.TeamLeaguePosition)
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
