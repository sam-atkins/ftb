package reporter

import (
	"testing"
	"time"
)

func Test_buildLeagueURL(t *testing.T) {
	type args struct {
		league string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Build Bundesliga URL",
			args: args{
				league: "BL1",
			},
			want: "competitions/BL1/matches",
		},
		{
			name: "Build Premier League URL",
			args: args{
				league: "PL",
			},
			want: "competitions/PL/matches",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildLeagueURL(tt.args.league); got != tt.want {
				t.Errorf("buildLeagueURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamURL_teamFinishedMatches(t *testing.T) {
	type fields struct {
		now time.Time
	}
	type args struct {
		teamId     string
		matchLimit bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "Build team finished matches URL with time limit",
			fields: fields{now: time.Date(2021, time.November, 1, 23, 0, 0, 0, time.UTC)},
			args:   args{teamId: "FCB", matchLimit: true},
			want:   "teams/FCB/matches?status=FINISHED&dateFrom=2021-10-04&dateTo=2021-11-29",
		},
		{
			name:   "Build team finished matches URL with time limit",
			fields: fields{now: time.Date(2021, time.November, 1, 23, 0, 0, 0, time.UTC)},
			args:   args{teamId: "FCB", matchLimit: false},
			want:   "teams/FCB/matches?status=FINISHED",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &teamURL{
				now: tt.fields.now,
			}
			if got := tr.teamFinishedMatches(tt.args.teamId, tt.args.matchLimit); got != tt.want {
				t.Errorf("teamURL.teamFinishedMatches() got = %v, want %v", got, tt.want)
			}
		})
	}
}
