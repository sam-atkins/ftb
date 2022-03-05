package reporter

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func formatFloatOrNil(teamScore interface{}) string {
	if teamScore == nil {
		return "-"
	}
	return fmt.Sprint(teamScore)
}

func convertToTitle(input string) string {
	// regex for only letters
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(input, " ")
	lower := strings.ToLower(processedString)
	return strings.Title(lower)
}
