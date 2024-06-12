package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func checkForMoreResults(chromeDpCtx context.Context, moreResults bool) (context.Context, bool) {
	/*
	   - check to see if the "Next page" button is on the page
	*/

	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.Evaluate(
			fmt.Sprintf(`Array.from(document.querySelectorAll('%s')).filter(el => el.innerText.trim() === "Next page") !== []`, MORE_RESULTS_SELECTOR),
			&moreResults,
		),
	); err != nil {
		log.Fatalf("%sError in checkForMoreResults()%s: %s", err)
	}

	return chromeDpCtx, moreResults
}

func getResultLinks(chromeDpCtx context.Context) []string {
	/*
		- find all <span url_il>
			- this appears to be the element that holds all links to urls
	*/
	var resultLinks []string
	if err := chromedp.Run(
		chromeDpCtx,
		chromedp.EvaluateAsDevTools(
			fmt.Sprintf(`Array.from(document.querySelectorAll('%s')).map(el => el.innerText)`, RESULT_SELECTOR),
			&resultLinks,
		),
	); err != nil {
		log.Fatalf("%sError in getResultLinks%s: %s", RED, RESET, err)
	}

	return resultLinks
}
