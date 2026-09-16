package main

import (
	"fmt"
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	value := float32(25)
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Slider",
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
				dxui.Slider(dxui.SliderProps{
					Min:      0,
					Max:      100,
					Step:     5,
					Value:    value,
					OnChange: dxui.Assign(&value),
				}),
				dxui.Label(fmt.Sprint(value)),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
