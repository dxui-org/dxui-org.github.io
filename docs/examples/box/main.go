package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {

	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Box",
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
			dxui.Box(
				dxui.BoxProps{
					Direction: dxui.Horizontal,
					Gap:       12,
					Align:     dxui.AlignCenter,
				},
				dxui.Label("Left"),
				dxui.TextButton(dxui.ButtonProps{}, "Right"),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
