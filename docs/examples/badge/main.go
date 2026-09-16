package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {

	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Badge",
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
			dxui.Badge(
				dxui.BadgeProps{
					Style: dxui.Style{
						Background: dxui.TokenColor(dxui.Color.Semantic.Success),
					},
				},
				dxui.Label("Ready"),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
