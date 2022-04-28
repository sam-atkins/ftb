package reporter

import (
	"fmt"
	"log"

	"github.com/sam-atkins/ftb/api"
	"github.com/sam-atkins/ftb/writer"
)

// GetScorers fetches top scorers for a league and prints to stdout
func GetScorers(league string) {
	endpoint := fmt.Sprintf("competitions/%s/scorers", league)
	client := api.NewClient()
	response, responseErr := client.GetScorers(endpoint)
	if responseErr != nil {
		log.Printf("Something went wrong with the request: %s\n", responseErr)
		return
	}

	message := fmt.Sprintf("Top Scorers in the %v\n", response.Body.Competition.Name)
	header := []string{"Name", "Team", "Goals"}
	var rows [][]string
	for _, v := range response.Body.Scorers {
		rows = append(rows, []string{
			v.Player.Name,
			v.Team.Name,
			fmt.Sprint(v.NumberOfGoals),
		})
	}
	writer.NewTable(header, message, rows).RenderTable()
}
