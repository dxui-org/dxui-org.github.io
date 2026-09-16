package main

import (
	"fmt"
	"github.com/dxui-org/dxui"
	"image"
	"image/color"
	"log"
)

func main() {
	pixels := image.NewRGBA(image.Rect(0, 0, 64, 64))
	for y := 0; y < 64; y++ {

		for x := 0; x < 64; x++ {
			pixels.SetRGBA(
				x,
				y,
				color.RGBA{
					R: uint8(x * 4),
					G: uint8(y * 4),
					B: 160,
					A: 255,
				},
			)
		}
	}
	source := dxui.ImageFromGo(pixels)
	status := "Loading"
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Image",
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
				dxui.Image(dxui.ImageProps{
					Source: source,
					Fit:    dxui.ImageContain,
					Alignment: dxui.Point{
						X: .5,
						Y: .5,
					},
					MaxPixels: 1024 * 1024,
					Style: dxui.Style{
						Width:  dxui.Px(180),
						Height: dxui.Px(120),
					},
					OnLoad: func(s dxui.Size) {
						status = fmt.Sprint(s)
					},
					OnError: func(err error) {
						status = err.Error()
					},
				}),
				dxui.Label(status),
			),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
