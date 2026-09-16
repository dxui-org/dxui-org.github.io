package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	app := dxui.NewApp(dxui.AppOptions{
		Background: dxui.RGBA(248, 250, 252, 255),
		Title:      "Hello dxui",
		Width:      640,
		Height:     400,
	})
	if err := app.Run(func() dxui.View {
		return dxui.Box(
			dxui.BoxProps{
				Gap: 16,
				Style: dxui.Style{
					Padding: dxui.Padding(24),
				},
			},
			dxui.Label("Hello, dxui!"),
			dxui.TextButton(dxui.ButtonProps{
				OnPress: app.Close,
			}, "Close"),
		)
	}); err != nil {
		log.Fatal(err)
	}
}
