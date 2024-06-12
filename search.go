package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
)

func search(host string, searchTerm string, chromeDpContext context.Context, maxPages int) []string {
	/*
		- get the search URL with URL encoded q param
		- navigate to https://{host}/search?q={searchTerm}
		- get more search results by clicking "Next page"
		- get the URLs from the <span class="url_li"> tags in the search results
	*/

	search_url := getSearchURL(host, searchTerm)

	var areResults bool
	chromeDpCtx, areResults := navigateToSearxNG(chromeDpContext, search_url, areResults)

	if areResults {
		searchResults := getSearchResults(chromeDpCtx, maxPages)

		return searchResults
	} else {
		return make([]string, 0)
	}
}

func getSearchURL(host string, searchTerm string) string {
	/*
		- url encode the search term
		- return the searxng search url
	*/

	encodedSearchTerm := url.QueryEscape(searchTerm)

	return fmt.Sprintf("https://%s/search?q=%s", host, encodedSearchTerm)
}

func navigateToSearxNG(chromeDpCtx context.Context, search_url string, areResults bool) (context.Context, bool) {
	/*
		- navigate to https://{host}/search?q={searchTerm}
		- wait for "Next page" button to appear to ensure all results are loaded
	*/

	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.Navigate(search_url),
		chromedp.Evaluate(
			fmt.Sprintf(`Array.from(document.querySelectorAll('%s')).length > 0`, MORE_RESULTS_SELECTOR),
			&areResults,
		),
	); err != nil {
		log.Fatalf("%sError in navigateToSearxNG()%s: %s", RED, RESET, err)
	}

	return chromeDpCtx, areResults
}

func getSearchResults(chromeDpCtx context.Context, maxPages int) []string {
	/*
		- get the CSS selector for the "Next page" button
		- check the page to see if the "Next page" button exists
			- if it does not exist, we have reached the final results page
			- else, keeping clicking until maxPages is reached
	*/

	pageCount := maxPages

	var moreResults bool
	var searchResultsURLs []string
	for {
		chromeDpCtx, moreResults = checkForMoreResults(chromeDpCtx, moreResults)

		if (!moreResults) || (pageCount == 0) {
			break
		}

		resultLinks := getResultLinks(chromeDpCtx)
		searchResultsURLs = append(searchResultsURLs, resultLinks...)
		chromeDpCtx = getMoreResults(chromeDpCtx)

		pageCount -= 1
	}

	return searchResultsURLs
}

func getMoreResults(chromeDpCtx context.Context) context.Context {
	/*
		- if the "Next page" button exists, click it to get more results
		- wait 1 second for new results to load
	*/

	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.Click(MORE_RESULTS_SELECTOR),
	); err != nil {
		log.Fatalf("%sError in getMoreResults()%s: %s", RED, RESET, err)
	}

	time.Sleep(1 * time.Second)

	return chromeDpCtx
}
