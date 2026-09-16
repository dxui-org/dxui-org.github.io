package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	offset := dxui.Point{}
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Scroll",
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
			dxui.Scroll(
				dxui.ScrollProps{
					Style: dxui.Style{
						Height: dxui.Px(240),
						Width:  dxui.Px(400),
					},
					Offset:    dxui.Some(offset),
					OnScroll:  dxui.Assign(&offset),
					Scrollbar: dxui.ScrollbarAlways,
				},
				dxui.Box(
					dxui.BoxProps{
						Gap: 12,
					},
					dxui.Label("Top"),
					dxui.Box(dxui.BoxProps{
						Style: dxui.Style{
							Height: dxui.Px(600),
							Shrink: dxui.NoShrink(),
						},
					}),
					dxui.Label("Bottom"),
				),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
