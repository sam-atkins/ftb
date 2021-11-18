package config

type leagueData struct {
	LeagueCode string
	LeagueName string
}

type teamData struct {
	Id         string
	Name       string
	Code       string
	League     string
	LeagueCode string
}

var LeagueConfig = []leagueData{
	{
		LeagueCode: "BL1",
		LeagueName: "1. Bundesliga",
	},
	{
		LeagueCode: "PL",
		LeagueName: "Premier League",
	},
	{
		LeagueCode: "PD",
		LeagueName: "La Liga",
	},
}

// TeamConfig provides useful info for each team to help with commands and API requests
var TeamConfig = []teamData{
	{
		Id:         "1",
		Name:       "1. FC Köln",
		Code:       "FCK",
		League:     "1. Bundesliga",
		LeagueCode: "BL1",
	},
	{
		Id:         "4",
		Name:       "Borussia Dortmund",
		Code:       "BVB",
		League:     "1. Bundesliga",
		LeagueCode: "BL1",
	},
	{
		Id:         "61",
		Name:       "Chelsea FC",
		Code:       "CHE",
		League:     "Premier League",
		LeagueCode: "PL",
	},
	{
		Id:         "78",
		Name:       "Club Atlético de Madrid",
		Code:       "ATM",
		League:     "La Liga",
		LeagueCode: "PD",
	},
	{
		Id:         "81",
		Name:       "FC Barcelona",
		Code:       "BAR",
		League:     "La Liga",
		LeagueCode: "PD",
	},
	{
		Id:         "5",
		Name:       "FC Bayern München",
		Code:       "FCB",
		League:     "1. Bundesliga",
		LeagueCode: "BL1",
	},
	{
		Id:         "64",
		Name:       "Liverpool FC",
		Code:       "LIV",
		League:     "Premier League",
		LeagueCode: "PL",
	},
	{
		Id:         "86",
		Name:       "Real Madrid CF",
		Code:       "RMA",
		League:     "La Liga",
		LeagueCode: "PD",
	},
	{
		Id:         "73",
		Name:       "Tottenham Hotspur FC",
		Code:       "TOT",
		League:     "Premier League",
		LeagueCode: "PL",
	},
}
