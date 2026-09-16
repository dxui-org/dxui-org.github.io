package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	choice := "small"
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Radio",
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
					Gap: 12,
				},
				dxui.Radio(
					dxui.RadioProps{
						Selected: choice == "small",
						OnSelect: func() {
							choice = "small"
						},
					},
					dxui.Label("Small"),
				),
				dxui.Radio(
					dxui.RadioProps{
						Selected: choice == "large",
						OnSelect: func() {
							choice = "large"
						},
					},
					dxui.Label("Large"),
				),
				dxui.Label(choice),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
