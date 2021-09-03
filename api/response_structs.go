package api

import "time"

type LeagueResponse struct {
	Filters struct {
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"season"`
	Standings []struct {
		Stage string      `json:"stage"`
		Type  string      `json:"type"`
		Group interface{} `json:"group"`
		Table []struct {
			Position int `json:"position"`
			Team     struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				CrestURL string `json:"crestUrl"`
			} `json:"team"`
			PlayedGames    int         `json:"playedGames"`
			Form           interface{} `json:"form"`
			Won            int         `json:"won"`
			Draw           int         `json:"draw"`
			Lost           int         `json:"lost"`
			Points         int         `json:"points"`
			GoalsFor       int         `json:"goalsFor"`
			GoalsAgainst   int         `json:"goalsAgainst"`
			GoalDifference int         `json:"goalDifference"`
		} `json:"table"`
	} `json:"standings"`
}

type ScorersResponse struct {
	Count   int `json:"count"`
	Filters struct {
		Limit int `json:"limit"`
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"season"`
	Scorers []struct {
		Player struct {
			ID             int         `json:"id"`
			Name           string      `json:"name"`
			FirstName      string      `json:"firstName"`
			LastName       interface{} `json:"lastName"`
			DateOfBirth    string      `json:"dateOfBirth"`
			CountryOfBirth string      `json:"countryOfBirth"`
			Nationality    string      `json:"nationality"`
			Position       string      `json:"position"`
			ShirtNumber    interface{} `json:"shirtNumber"`
			LastUpdated    time.Time   `json:"lastUpdated"`
		} `json:"player"`
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
		NumberOfGoals int `json:"numberOfGoals"`
	} `json:"scorers"`
}

type CompetitionResponse struct {
	Count   int `json:"count"`
	Filters struct {
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Matches []struct {
		ID     int `json:"id"`
		Season struct {
			ID              int    `json:"id"`
			StartDate       string `json:"startDate"`
			EndDate         string `json:"endDate"`
			CurrentMatchday int    `json:"currentMatchday"`
		} `json:"season"`
		UtcDate     time.Time   `json:"utcDate"`
		Status      string      `json:"status"`
		Matchday    int         `json:"matchday"`
		Stage       string      `json:"stage"`
		Group       interface{} `json:"group"`
		LastUpdated time.Time   `json:"lastUpdated"`
		Odds        struct {
			Msg string `json:"msg"`
		} `json:"odds"`
		Score struct {
			Winner   string `json:"winner"`
			Duration string `json:"duration"`
			FullTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"fullTime"`
			HalfTime struct {
				HomeTeam int `json:"homeTeam"`
				AwayTeam int `json:"awayTeam"`
			} `json:"halfTime"`
			ExtraTime struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"extraTime"`
			Penalties struct {
				HomeTeam interface{} `json:"homeTeam"`
				AwayTeam interface{} `json:"awayTeam"`
			} `json:"penalties"`
		} `json:"score"`
		HomeTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"awayTeam"`
		Referees []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Role        string `json:"role"`
			Nationality string `json:"nationality"`
		} `json:"referees"`
	} `json:"matches"`
}
