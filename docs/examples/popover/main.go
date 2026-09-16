package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	open := false
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Popover",
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
			dxui.Popover(
				dxui.PopoverProps{
					Open:         open,
					OnOpenChange: dxui.Assign(&open),
					Placement:    dxui.OverlayBottomStart,
					Offset:       dxui.Metric(6),
				},
				dxui.Label("Open details"),
				dxui.Box(
					dxui.BoxProps{
						Gap: 12,
					},
					dxui.Label("Details"),
					dxui.TextButton(
						dxui.ButtonProps{
							OnPress: func() {
								open = false
							},
						},
						"Close",
					),
				),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
