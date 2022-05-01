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

func Test_buildLeagueStandingsURL(t *testing.T) {
	type args struct {
		leagueCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Build Bundesliga league standings URL",
			args: args{
				leagueCode: "BL1",
			},
			want: "competitions/BL1/standings",
		},
		{
			name: "Build Premier League league standings URL",
			args: args{
				leagueCode: "PL",
			},
			want: "competitions/PL/standings",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildLeagueStandingsURL(tt.args.leagueCode); got != tt.want {
				t.Errorf("buildLeagueStandingsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scorersURL(t *testing.T) {
	type args struct {
		league string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Build Bundesliga scorers URL",
			args: args{
				league: "BL1",
			},
			want: "competitions/BL1/scorers",
		},
		{
			name: "Build Premier League scorers URL",
			args: args{
				league: "PL",
			},
			want: "competitions/PL/scorers",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scorersURL(tt.args.league); got != tt.want {
				t.Errorf("scorersURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teamURL_teamFinishedMatches(t *testing.T) {
	type fields struct {
		now time.Time
	}
	type args struct {
		teamId        string
		matchLimit    bool
		matchComplete bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "Build URL for team finished matches with time limit",
			fields: fields{now: time.Date(2021, time.November, 1, 23, 0, 0, 0, time.UTC)},
			args:   args{teamId: "FCB", matchLimit: true, matchComplete: true},
			want:   "teams/FCB/matches?status=FINISHED&dateFrom=2021-10-04&dateTo=2021-11-29",
		},
		{
			name:   "Build URL for team finished matches without time limit",
			fields: fields{now: time.Date(2021, time.November, 1, 23, 0, 0, 0, time.UTC)},
			args:   args{teamId: "FCB", matchLimit: false, matchComplete: true},
			want:   "teams/FCB/matches?status=FINISHED",
		},
		{
			name:   "Build URL for team scheduled matches URL without time limit",
			fields: fields{now: time.Date(2021, time.November, 1, 23, 0, 0, 0, time.UTC)},
			args:   args{teamId: "FCB", matchLimit: false, matchComplete: false},
			want:   "teams/FCB/matches?status=SCHEDULED",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &teamURL{
				now: tt.fields.now,
			}
			if got := tr.teamMatches(tt.args.teamId, tt.args.matchLimit, tt.args.matchComplete); got != tt.want {
				t.Errorf("teamURL.teamMatches() got = %v, want %v", got, tt.want)
			}
		})
	}
}
