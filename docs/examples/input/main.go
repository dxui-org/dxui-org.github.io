package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	value := ""
	selection := dxui.TextRange{}
	message := "Press Enter"
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Input",
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
				dxui.Input(dxui.InputProps{
					Key:                "password",
					Value:              value,
					OnChange:           dxui.Assign(&value),
					Selection:          dxui.Some(selection),
					OnSelectionChange:  dxui.Assign(&selection),
					Password:           true,
					ShowPasswordToggle: true,
					Placeholder:        "Password",
					OnSubmit: func() {
						message = "Submitted"
					},
				}),
				dxui.Label(message),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
