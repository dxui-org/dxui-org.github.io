package main

import (
	"github.com/dxui-org/dxui"
	"image"
	"image/color"
	"log"
)

func main() {
	pixels := image.NewRGBA(image.Rect(0, 0, 48, 48))
	for y := 0; y < 48; y++ {

		for x := 0; x < 48; x++ {
			pixels.SetRGBA(x, y, color.RGBA{
				R: 40,
				G: 160,
				B: 120,
				A: 255,
			})
		}
	}
	source := dxui.ImageFromGo(pixels)
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Avatar",
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
			dxui.Avatar(dxui.AvatarProps{
				Source: source,
				Shape:  dxui.AvatarCircle,
				Size:   64,
			}),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
