package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/sam-atkins/ftb/writer"
	"gopkg.in/yaml.v2"
)

var teamConfigFile = ".config/ftb/teams.yaml"

type leagueData struct {
	LeagueCode string
	LeagueName string
	Country    string
}

var LeagueConfig = []leagueData{
	{
		LeagueCode: "BL1",
		LeagueName: "1. Bundesliga",
		Country:    "Germany",
	},
	{
		LeagueCode: "PL",
		LeagueName: "Premier League",
		Country:    "England",
	},
	{
		LeagueCode: "PD",
		LeagueName: "La Liga",
		Country:    "Spain",
	},
}

// GetLeagueCodes returns the leagues and their codes
func GetLeagueCodes() [][]string {
	var leagueCodes [][]string
	for _, v := range LeagueConfig {
		leagueCodes = append(leagueCodes, []string{
			v.LeagueName,
			v.LeagueCode,
			v.Country,
		},
		)
	}
	return leagueCodes
}

type teamConfig []struct {
	ID   int `yaml:"id"`
	Area struct {
		ID   int    `yaml:"id"`
		Name string `yaml:"name"`
	} `yaml:"area"`
	Name        string    `yaml:"name"`
	Shortname   string    `yaml:"shortname"`
	Tla         string    `yaml:"tla"`
	Cresturl    string    `yaml:"cresturl"`
	Address     string    `yaml:"address"`
	Phone       string    `yaml:"phone"`
	Website     string    `yaml:"website"`
	Email       string    `yaml:"email"`
	Founded     int       `yaml:"founded"`
	Clubcolors  string    `yaml:"clubcolors"`
	Venue       string    `yaml:"venue"`
	Lastupdated time.Time `yaml:"lastupdated"`
}

func (c *teamConfig) parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

// GetTeamCodesForWriter returns the teams and their codes
func GetTeamCodesForWriter() [][]string {
	teamCfg, err := readTeamsCodesFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	var teamCodes [][]string
	for _, v := range teamCfg {
		teamCodes = append(teamCodes, []string{
			v.Name,
			v.Tla,
			v.Area.Name,
		},
		)
	}
	return teamCodes
}

// GetTeamInfoFromUserTeamCode verifies the arg userTeamcode and then provides the team's
// league code, team name and ID
func GetTeamInfoFromUserTeamCode(userTeamCode string) (leagueCode string, teamName string, teamId string) {
	teamCfg, err := readTeamsCodesFromConfig()
	if err != nil {
		log.Fatal(err)
	}

	var teamCountry string
	for _, v := range teamCfg {
		if v.Tla == strings.ToUpper(userTeamCode) {
			teamName = v.Name
			teamId = strconv.Itoa(v.ID)
			teamCountry = v.Area.Name
		}
		for _, v := range LeagueConfig {
			if teamCountry == v.Country {
				leagueCode = v.LeagueCode
			}
		}
	}

	if leagueCode == "" && teamName == "" {
		CodeNotFound()
	}
	return leagueCode, teamName, teamId
}

// CodeNotFound used when the user enters an unknown flag code. It prints the available
// codes to stdout and exits (1)
func CodeNotFound() {
	header := []string{"Team", "Code"}
	message := fmt.Sprintf("Did not recognise that team. These are the available team codes:")
	teamCodes := GetTeamCodesForWriter()
	writer.NewTable(header, message, teamCodes).RenderTable()
	os.Exit(1)
}

// GetTeamConfigPath returns the path to the teams.yml config file
func GetTeamConfigPath() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return filepath.Join(home, teamConfigFile)
}

// ResetTeamConfigFile truncates the team config yaml file
func ResetTeamConfigFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("could not open file %q for truncation: %v", filename, err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("could not close file handler for %q after truncation: %v", filename, err)
	}
	return nil
}

func readTeamsCodesFromConfig() (teamConfig, error) {
	fileName := GetTeamConfigPath()
	data, fileErr := ioutil.ReadFile(fileName)
	if fileErr != nil {
		return nil, fileErr
	}

	var teamCfg teamConfig
	if parseErr := teamCfg.parse(data); parseErr != nil {
		return nil, parseErr
	}

	return teamCfg, nil
}
