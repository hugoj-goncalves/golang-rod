// Package main ...
package main

import (
	"context"
	"log"
	"strings"

	"github.com/go-rod/rod"
)

// This example demonstrates  how to extract text from a specific element.
func main() {
	ctx := context.Background()
	page := rod.New().MustConnect(ctx).MustPage("https://pkg.go.dev/time")

	res := page.MustElement("#pkg-overview").MustParent().MustText()
	log.Println(strings.TrimSpace(res))
}
