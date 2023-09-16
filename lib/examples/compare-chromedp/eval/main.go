// Package main ...
package main

import (
	"context"
	"log"

	"github.com/go-rod/rod"
)

// This example shows how we can use Eval to run scripts in the page.
// Note: `this` in the eval function will refer to the element that Eval is
// called  on. This can be useful for things such as blurring elements.
func main() {
	ctx := context.Background()
	res := rod.New().MustConnect(ctx).
		MustPage("https://www.google.com/").
		MustElement(`input`).
		MustEval("() => Object.keys(window)")

	log.Printf("window object keys: %v", res)
}
