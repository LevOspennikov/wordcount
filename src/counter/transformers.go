package counter

import (
	"log"
	"regexp"
)

func TransformerWords(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(str, "")
	return processedString
}