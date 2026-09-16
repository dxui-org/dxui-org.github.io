package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	value := "home"
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Tabs",
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
				dxui.Tabs(dxui.TabsProps{
					Value:    value,
					OnChange: dxui.Assign(&value),
					Items: []dxui.TabItem{
						{
							Value: "home",
							Label: "Home",
						},
						{
							Value: "settings",
							Label: "Settings",
						},
					},
				}),
				dxui.Label("Panel: "+value),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
