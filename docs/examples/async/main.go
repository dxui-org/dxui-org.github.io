package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	message, busy := "Ready", false
	app := dxui.NewApp(dxui.AppOptions{
		Background: dxui.RGBA(248, 250, 252, 255),
		Title:      "Background work",
		Width:      640,
		Height:     400,
		OnCloseRequest: func(a *dxui.App) {
			cancel()
			a.Close()
		},
	})
	err := app.Run(func() dxui.View {
		return dxui.Box(
			dxui.BoxProps{
				Gap: 16,
				Style: dxui.Style{
					Padding: dxui.Padding(24),
				},
			},
			dxui.Label(message),
			dxui.TextButton(
				dxui.ButtonProps{
					Disabled: busy,
					OnPress: func() {
						busy = true
						message = "Computing"
						n := int64(5_000_000)
						go func() {
							var sum int64
							for i := int64(0); i < n; i++ {
								if i%4096 == 0 {
									select {
									case <-ctx.Done():
										return
									default:
									}
								}
								sum += i
							}
							result := fmt.Sprintf("Sum: %d", sum)
							if err := app.Update(func() {
								message = result
								busy = false
							}); err != nil &&
								!errors.Is(err, dxui.ErrAppClosed) &&
								!errors.Is(err, dxui.ErrAppNotRunning) {
								log.Print(err)
							}
						}()
					},
				},
				"Compute",
			),
		)
	})
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}
