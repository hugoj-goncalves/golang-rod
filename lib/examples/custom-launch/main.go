// Package main ...
package main

import (
	"context"
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	ctx := context.Background()
	l := launcher.New()

	// For more info: https://pkg.go.dev/github.com/go-rod/rod/lib/launcher
	u := l.MustLaunch(ctx)

	browser := rod.New().ControlURL(u).MustConnect(ctx)

	page := browser.MustPage("http://example.com").MustWaitStable()

	fmt.Println(page.MustInfo().Title)
}
