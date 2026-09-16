package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	value := "First line\nSecond line"
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Textarea",
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
			dxui.Textarea(dxui.TextareaProps{
				Key:         "notes",
				Value:       value,
				OnChange:    dxui.Assign(&value),
				Wrap:        dxui.TextWrapWords,
				Placeholder: "Notes",
				Style: dxui.Style{
					Width:  dxui.Px(400),
					Height: dxui.Px(180),
				},
			}),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
