package main

import (
	"fmt"
	"log"
	"os"
)

func writeResults(searchResultsURLs []string, output string) {
	file, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("%sError in writeResults()%s: %s", RED, RESET, err)
	}
	defer file.Close()

	for _, value := range searchResultsURLs {
		line := fmt.Sprintf("%s\n", value)
		lineBytes := []byte(line)

		file.Write(lineBytes)
	}
}
