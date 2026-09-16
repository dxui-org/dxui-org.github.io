package main

import (
	"fmt"
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	value := ""
	actions := 0
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Menu",
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
				dxui.Menu(dxui.MenuProps{
					Value: value,
					Items: []dxui.MenuItem{
						{
							Value: "save",
							Label: "Save",
						},
						{
							Value: "export",
							Label: "Export",
						},
					},
					OnAction: func(v string) {
						value = v
						actions++
					},
				}),
				dxui.Label(fmt.Sprintf("%s: %d actions", value, actions)),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
