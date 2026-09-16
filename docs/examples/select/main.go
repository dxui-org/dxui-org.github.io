package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	value := "go"
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Select",
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
			dxui.Select(dxui.SelectProps{
				Value:       value,
				OnChange:    dxui.Assign(&value),
				Placeholder: "Choose",
				Options: []dxui.SelectOption{
					{
						Value: "go",
						Label: "Go",
					},
					{
						Value: "rust",
						Label: "Rust",
					},
					{
						Value:    "later",
						Label:    "Later",
						Disabled: true,
					},
				},
			}),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
