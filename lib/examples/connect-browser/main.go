// Package main ...
package main

import (
	"context"
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// To manually launch a browser
func main() {
	ctx := context.Background()
	// Launch your local browser first:
	//
	//     chrome --headless --remote-debugging-port=9222
	//
	// Or use docker:
	//
	//     docker run -p 9222:9222 ghcr.io/go-rod/rod chrome --headless --no-sandbox --remote-debugging-port=9222 --remote-debugging-address=0.0.0.0
	//
	u := launcher.MustResolveURL("")

	browser := rod.New().ControlURL(u).MustConnect(ctx)

	fmt.Println(
		browser.MustPage("https://mdn.dev/").MustEval("() => document.title"),
	)
}
