package main

import (
	"github.com/dxui-org/dxui"
	"log"
	"time"
)

func main() {

	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Tooltip",
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
			dxui.Tooltip(
				dxui.TooltipProps{
					Placement: dxui.OverlayTop,
					Delay:     700 * time.Millisecond,
				},
				dxui.TextButton(dxui.ButtonProps{}, "Hover or Tab here"),
				dxui.Label("A helpful hint"),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
