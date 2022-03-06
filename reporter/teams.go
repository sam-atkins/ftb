package reporter

import (
	"fmt"
	"log"

	"github.com/sam-atkins/ftb/api"
)

// GetTable gets a league table
func GetTeams(league string) {
	endpoint := fmt.Sprintf("competitions/%s/teams", league)
	client := api.NewClient()
	response, responseErr := client.GetTeams(endpoint)
	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}

	// if debug is true then print the below to terminal
	fmt.Printf("Competition ID: %v\n", response.Body.Competition.ID)
	fmt.Printf("Competition Code: %v\n", response.Body.Competition.Code)
	fmt.Printf("Teams: %v\n", response.Body.Teams)
}
