package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	value := ""
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "InputGroup",
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
			dxui.InputGroup(
				dxui.InputGroupProps{
					Style: dxui.Style{
						Width: dxui.Px(360),
					},
				},
				dxui.InputGroupContent{
					Input: dxui.Input(dxui.InputProps{
						Key:      "query",
						Value:    value,
						OnChange: dxui.Assign(&value),
					}),
					Prefix: dxui.Some(dxui.Label("Search")),
					Suffix: dxui.Some(dxui.TextButton(
						dxui.ButtonProps{
							OnPress: func() {
								value = ""
							},
						},
						"Clear",
					)),
				},
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
