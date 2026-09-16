package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	app := dxui.NewApp(dxui.AppOptions{
		Background: dxui.RGBA(248, 250, 252, 255),
		Title:      "Responsive",
		Width:      800,
		Height:     480,
	})
	if err := app.RunResponsive(func(ctx dxui.LayoutContext) dxui.View {
		direction := dxui.Horizontal
		if ctx.Width < 600 {
			direction = dxui.Vertical
		}
		return dxui.Box(
			dxui.BoxProps{
				Direction: direction,
				Gap:       16,
				Style: dxui.Style{
					Width:   dxui.Px(ctx.Width),
					Height:  dxui.Px(ctx.Height),
					Padding: dxui.Padding(24),
				},
			},
			dxui.Box(
				dxui.BoxProps{
					Style: dxui.Style{
						Width:  dxui.Px(160),
						Shrink: dxui.NoShrink(),
					},
				},
				dxui.Label("Navigation"),
			),
			dxui.Box(
				dxui.BoxProps{
					Style: dxui.Style{
						Grow: 1,
					},
				},
				dxui.Label("Content"),
			),
		)
	}); err != nil {
		log.Fatal(err)
	}
}
