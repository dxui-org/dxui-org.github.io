package main

import (
	"fmt"
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	items := make([]string, 10000)
	for i := range items {

		items[i] = fmt.Sprintf("item-%d", i)
	}
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "VirtualList",
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
			dxui.VirtualList(dxui.VirtualListProps{
				Key: "items",
				Style: dxui.Style{
					Height: dxui.Px(320),
					Width:  dxui.Px(400),
				},
				Count:     len(items),
				Version:   1,
				RowHeight: 36,
				Overscan:  3,
				ItemKey: func(i int) string {
					return items[i]
				},
				Build: func(i int) dxui.View {
					return dxui.Label(items[i])
				},
			}),
		)
	}); err != nil {

		log.Fatal(err)
	}
}
