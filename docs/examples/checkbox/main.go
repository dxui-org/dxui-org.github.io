package main

import (
	"fmt"
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	checked := false
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Checkbox",
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
				dxui.Checkbox(
					dxui.CheckboxProps{
						Checked:  checked,
						OnChange: dxui.Assign(&checked),
					},
					dxui.Label("Accept terms"),
				),
				dxui.Label(fmt.Sprint(checked)),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
