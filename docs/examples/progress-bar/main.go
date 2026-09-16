package main

import (
	"fmt"
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	value := float32(.25)
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "ProgressBar",
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
				dxui.ProgressBar(dxui.ProgressBarProps{
					Value: value,
				}),
				dxui.TextButton(
					dxui.ButtonProps{
						OnPress: func() {
							value += .1
							if value > 1 {
								value = 0
							}
						},
					},
					"Advance",
				),
				dxui.Label(fmt.Sprintf("%.0f%%", value*100)),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
