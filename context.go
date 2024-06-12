package main

import (
	"context"

	"github.com/chromedp/chromedp"
)

func getEmptyContext() (context.Context, context.CancelFunc) {
	return chromedp.NewContext(context.Background())
}

func getAllocContext(ctx context.Context) (context.Context, context.CancelFunc) {
	/*
		- set headless to false to captcha can be solved if necessary
			- alternative may be 2Captcha, but that looks dodgey and costs $$
			- solving a Captcha quick never hurt anyone; the rest is automated
	*/
	return chromedp.NewExecAllocator(ctx, append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", true),
	)...)

}

func getChromeDpContext(allocCtx context.Context) (context.Context, context.CancelFunc) {
	return chromedp.NewContext(allocCtx)
}
