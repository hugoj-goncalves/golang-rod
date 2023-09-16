// Package main ...
package main

import (
	"context"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
)

func main() {
	ctx := context.Background()
	page := rod.New().MustConnect(ctx).MustPage()

	// emulate iPhone 7 landscape
	err := page.Emulate(devices.IPhone6or7or8.Landscape())
	if err != nil {
		panic(err)
	}

	page.MustNavigate("https://www.whatsmyua.info/")
	page.MustScreenshot("screenshot1.png")

	// reset
	page.MustEmulate(devices.Clear)

	page.MustSetViewport(1920, 2000, 1, false)
	page.MustNavigate("https://www.whatsmyua.info/?a")
	page.MustScreenshot("screenshot2.png")
}
