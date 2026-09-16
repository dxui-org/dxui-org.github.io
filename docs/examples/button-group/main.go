package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	choice := "A"
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "ButtonGroup",
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
				dxui.Label(choice),
				dxui.ButtonGroup(
					dxui.ButtonGroupProps{
						Dividers: true,
					},
					dxui.TextButton(
						dxui.ButtonProps{
							OnPress: func() {
								choice = "A"
							},
						},
						"A",
					),
					dxui.TextButton(
						dxui.ButtonProps{
							OnPress: func() {
								choice = "B"
							},
						},
						"B",
					),
				),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
