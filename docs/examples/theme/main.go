package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	dark := false
	app := dxui.NewApp(dxui.AppOptions{
		Background: dxui.RGBA(248, 250, 252, 255),
		Title:      "Theme",
		Width:      640,
		Height:     400,
	})
	if err := app.Run(func() dxui.View {
		return dxui.Box(
			dxui.BoxProps{
				Gap: 16,
				Style: dxui.Style{
					Padding: dxui.Padding(24),
				},
			},
			dxui.TextButton(
				dxui.ButtonProps{
					States: dxui.StateStyles{
						Hover: dxui.StylePatch{
							Opacity: dxui.Some(float32(.8)),
						},
					},
					OnPress: func() {
						next := dxui.DarkTheme()
						if dark {
							next = dxui.LightTheme()
						}
						next.Semantic.Colors[dxui.Color.Semantic.Accent] =
							dxui.TokenColor(dxui.Color.Primitive.Blue600)
						if err := app.SetTheme(next); err != nil {
							log.Print(err)
							return
						}
						dark = !dark
					},
				},
				"Switch theme",
			),
		)
	}); err != nil {
		log.Fatal(err)
	}
}
