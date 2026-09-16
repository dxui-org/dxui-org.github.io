package main

import (
	"github.com/dxui-org/dxui"
	"github.com/dxui-org/dxui/icon"
	"log"
)

func main() {

	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Icon",
		Width:      640,
		Height:     480,
		Background: dxui.RGBA(248, 250, 252, 255),
	})
	if err := app.Run(func() dxui.View {
		return dxui.Box(
			dxui.BoxProps{
				Gap: 16,
				Style: dxui.Style{
					Padding: dxui.Padding(24),
				},
			},
			dxui.Icon(dxui.IconProps{
				Data:        icon.Search(),
				Size:        32,
				StrokeWidth: 2,
				Color:       dxui.TokenColor(dxui.Color.Semantic.Accent),
			}),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
