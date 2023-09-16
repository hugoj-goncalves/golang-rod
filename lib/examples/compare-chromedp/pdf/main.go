// Package main ...
package main

import (
	"context"
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	ctx := context.Background()
	rod.New().MustConnect(ctx).MustPage("https://www.google.com/").MustWaitLoad().MustPDF("sample.pdf")
	fmt.Println("wrote sample.pdf")
}
