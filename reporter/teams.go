package reporter

import (
	"fmt"
	"log"
	"os"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/config"
	"gopkg.in/yaml.v2"
)

// GetTeamsConfig adds team config to the team config file based on the leagues
// available in the config package.
// truncateConfigFile will delete the contents of the existing team config file. This is
// useful at the beginning of a new season to ensure the teams are accurate.
// debug switches on additional logging to the console.
func GetTeamsConfig(truncateConfigFile, debug bool) {
	fileName := config.GetTeamConfigPath()
	if truncateConfigFile {
		config.ResetTeamConfigFile(fileName)
	}

	for _, v := range config.LeagueConfig {
		writeTeamDataToConfigFile(v.LeagueCode, fileName, debug)
		fmt.Printf("Added team config for the %s\n", v.LeagueName)
	}
}

func writeTeamDataToConfigFile(league, configFilePath string, debug bool) {
	endpoint := fmt.Sprintf("competitions/%s/teams", league)
	client := api.NewClient()
	response, responseErr := client.GetTeams(endpoint)
	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}

	if debug {
		fmt.Printf("Competition ID: %v\n", response.Body.Competition.ID)
		fmt.Printf("Competition Code: %v\n", response.Body.Competition.Code)
		fmt.Printf("Teams: %v\n", response.Body.Teams)
	}

	yamlData, err := yaml.Marshal(&response.Body.Teams)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	configFile, fileErr := os.OpenFile(configFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		log.Println("Unable to write data into the file")
		return
	}

	if _, err := configFile.Write(yamlData); err != nil {
		configFile.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := configFile.Close(); err != nil {
		log.Fatal(err)
	}
}
