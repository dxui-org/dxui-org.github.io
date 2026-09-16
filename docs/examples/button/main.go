package main

import (
	"fmt"
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	count := 0
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Button",
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
				dxui.Label(fmt.Sprint(count)),
				dxui.TextButton(
					dxui.ButtonProps{
						Variant: dxui.ButtonOutline,
						Tone:    dxui.ButtonSuccess,
						Size:    dxui.ButtonNormal,
						OnPress: func() {

							count++
						},
					},
					"Add",
				),
				dxui.Button(dxui.ButtonProps{
					Disabled: true,
				}, dxui.Label("Disabled")),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
