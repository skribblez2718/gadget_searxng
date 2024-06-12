package main

import (
	"flag"
	"fmt"
	"os"
)

func getArgs() (string, string, string, int) {
	flag.Usage = displayUsage

	var (
		host       string
		searchTerm string
		output     string
		maxPages   int
	)

	flag.StringVar(&host, "host", "", "SearxNG hostname. Example: my.searxng.local")
	flag.StringVar(&searchTerm, "searchTerm", "", "Searxng search term")
	flag.StringVar(&output, "output", "", "Path to file to write results to")
	flag.IntVar(&maxPages, "maxPages", -1, "Maximum number of result pages to search")

	flag.Parse()

	if searchTerm == "" || output == "" {
		flag.Usage()
		os.Exit(1)
	}

	return host, searchTerm, output, maxPages
}

func displayUsage() {
	usageMessage := "Description:\n"
	usageMessage += "\tUses ChromeDP to automate extracting URLs from SearxNG search results.\n"
	usageMessage += "Usage:\n"
	usageMessage += "\tgadget_searxng --host <searxng_host> --searchTerm \"<search_term>\" --output \"<output_file>\" [--maxPages <int>]\n"

	fmt.Print(usageMessage)
}
