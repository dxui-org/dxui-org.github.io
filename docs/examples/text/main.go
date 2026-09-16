package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {

	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Text",
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
			dxui.Text(dxui.TextProps{
				Value:    "Hello dxui",
				Wrap:     dxui.TextWrapWords,
				MaxLines: 2,
				Style: dxui.Style{
					Width: dxui.Px(240),
					Text: dxui.TextStyle{
						Size: 20,
					},
				},
			}),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
