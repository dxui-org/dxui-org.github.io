package main

import (
	"log"

	"github.com/dxui-org/dxui"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	font := dxui.FontBytes("application-font", goregular.TTF)
	app := dxui.NewApp(dxui.AppOptions{
		Title:                     "Application font",
		Width:                     640,
		Height:                    400,
		Background:                dxui.RGBA(248, 250, 252, 255),
		Fonts:                     []dxui.Font{font},
		DefaultFont:               "application-font",
		DisableSystemFontFallback: true,
	})
	if err := app.Run(func() dxui.View {
		return dxui.Text(dxui.TextProps{
			Value: "An application-owned font",
			Style: dxui.Style{
				Padding: dxui.Padding(24),
				Text: dxui.TextStyle{
					Families: []dxui.FontFamily{"application-font"},
					Size:     24,
				},
			},
		})
	}); err != nil {
		log.Fatal(err)
	}
}
