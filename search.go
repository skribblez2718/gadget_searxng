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

	chromeDpCtx := navigateToSearxNG(chromeDpContext, search_url)
	chromeDpCtx = expandSearchResults(chromeDpCtx, maxPages)

	searchResultUrls := getSearchResultUrls(chromeDpCtx)

	return searchResultUrls
}

func getSearchURL(host string, searchTerm string) string {
	/*
		- url encode the search term
		- return the searxng search url
	*/

	encodedSearchTerm := url.QueryEscape(searchTerm)

	return fmt.Sprintf("https://%s/search?q=%s", host, encodedSearchTerm)
}

func navigateToSearxNG(chromeDpCtx context.Context, search_url string) context.Context {
	/*
		- navigate to https://{host}/search?q={searchTerm}
		- wait for "Next page" button to appear to ensure all results are loaded
	*/

	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.Navigate(search_url),
		chromedp.WaitVisible(MORE_RESULTS_SELECTOR, chromedp.ByQuery),
	); err != nil {
		log.Fatalf("%sError in navigateToSearxNG()%s: %s", RED, RESET, err)
	}

	return chromeDpCtx
}

func expandSearchResults(chromeDpCtx context.Context, maxPages int) context.Context {
	/*
		- get the CSS selector for the "Next page" button
		- check the page to see if the "Next page" button exists
			- if it does not exist, we have reached the final results page
			- else, keeping clicking until maxPages is reached
	*/

	pageCount := maxPages

	var moreResults bool
	for {
		chromeDpCtx, moreResults = checkForMoreResults(chromeDpCtx, moreResults)

		if (!moreResults) || (pageCount == 0) {
			break
		}

		chromeDpCtx = getMoreResults(chromeDpCtx)

		pageCount -= 1
	}

	return chromeDpCtx
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

func getSearchResultUrls(chromeDpCtx context.Context) []string {
	/*
		- get all the <span class="url_il"> elements' InnerText (these contain URLs)
		- iterate through the InnerTexts
			- if InnerText contains a domain, transform to URL format
				- not all InnerTexts may contain a domain
	*/

	resultLinks := getResultLinks(chromeDpCtx)

	var searchResultsURLs []string
	searchResultsURLs = append(searchResultsURLs, resultLinks...)

	return searchResultsURLs
}
