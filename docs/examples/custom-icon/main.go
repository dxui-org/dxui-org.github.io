package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	data := dxui.IconData{
		ViewBox: dxui.Rect{
			Width:  24,
			Height: 24,
		},
		Commands: []dxui.PathCommand{
			{
				Verb: dxui.PathMove,
				Points: [3]dxui.Point{{
					X: 12,
					Y: 2,
				}},
			},
			{
				Verb: dxui.PathLine,
				Points: [3]dxui.Point{{
					X: 22,
					Y: 22,
				}},
			},
			{
				Verb: dxui.PathLine,
				Points: [3]dxui.Point{{
					X: 2,
					Y: 22,
				}},
			},
			{
				Verb: dxui.PathClose,
			},
		},
	}
	if _, err := data.Paths(100); err != nil {
		log.Fatal(err)
	}
	log.Printf("packed=%t identity=%d", data.IsPacked(), data.Identity())
	app := dxui.NewApp(dxui.AppOptions{
		Background: dxui.RGBA(248, 250, 252, 255),
		Title:      "Custom icon",
		Width:      400,
		Height:     300,
	})
	if err := app.Run(func() dxui.View {
		return dxui.Icon(dxui.IconProps{
			Data: data,
			Size: 80,
		})
	}); err != nil {
		log.Fatal(err)
	}
}
