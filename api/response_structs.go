package api

import "time"

type ApiLeagueResponse struct {
	StatusCode int
	Body       LeagueResponse
}

type ApiMatchesResponse struct {
	StatusCode int
	Body       matchesResponse
}

type ApiScorersResponse struct {
	StatusCode int
	Body       scorersResponse
}

type apiTeamsResponse struct {
	StatusCode int
	Body       teamsResponse
}

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
	Standings []standings `json:"standings"`
}

type standings struct {
	Stage string      `json:"stage"`
	Type  string      `json:"type"`
	Group interface{} `json:"group"`
	Table []table     `json:"table"`
}

type table struct {
	Position       int         `json:"position"`
	Team           team        `json:"team"`
	PlayedGames    int         `json:"playedGames"`
	Form           interface{} `json:"form"`
	Won            int         `json:"won"`
	Draw           int         `json:"draw"`
	Lost           int         `json:"lost"`
	Points         int         `json:"points"`
	GoalsFor       int         `json:"goalsFor"`
	GoalsAgainst   int         `json:"goalsAgainst"`
	GoalDifference int         `json:"goalDifference"`
}

type team struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CrestURL string `json:"crestUrl"`
}

type matchesResponse struct {
	Count   int `json:"count"`
	Filters struct {
		Permission string   `json:"permission"`
		Status     []string `json:"status"`
		Limit      int      `json:"limit"`
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
	Matches []matches `json:"matches"`
}

type matches struct {
	ID          int `json:"id"`
	Competition struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Area struct {
			Name      string `json:"name"`
			Code      string `json:"code"`
			EnsignURL string `json:"ensignUrl"`
		} `json:"area"`
	} `json:"competition"`
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
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
		Winner   interface{} `json:"winner"`
		Duration string      `json:"duration"`
		FullTime struct {
			HomeTeam interface{} `json:"homeTeam"`
			AwayTeam interface{} `json:"awayTeam"`
		} `json:"fullTime"`
		HalfTime struct {
			HomeTeam interface{} `json:"homeTeam"`
			AwayTeam interface{} `json:"awayTeam"`
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
	Referees []interface{} `json:"referees"`
}

type scorersResponse struct {
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
	Scorers []scorers `json:"scorers"`
}

type scorers struct {
	Player player `json:"player"`
	Team   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"team"`
	NumberOfGoals int `json:"numberOfGoals"`
}

type player struct {
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
}

type teamsResponse struct {
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
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday interface{} `json:"currentMatchday"`
		AvailableStages []string    `json:"availableStages"`
	} `json:"season"`
	Teams []struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		ShortName   string    `json:"shortName"`
		Tla         string    `json:"tla"`
		CrestURL    string    `json:"crestUrl"`
		Address     string    `json:"address"`
		Phone       string    `json:"phone"`
		Website     string    `json:"website"`
		Email       string    `json:"email"`
		Founded     int       `json:"founded"`
		ClubColors  string    `json:"clubColors"`
		Venue       string    `json:"venue"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"teams"`
}
