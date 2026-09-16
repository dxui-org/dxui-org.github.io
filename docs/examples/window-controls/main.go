package main

import (
	"fmt"
	"log"

	"github.com/dxui-org/dxui"
)

// App and Window share these window-level operations.
type windowControl interface {
	Title() string
	SetTitle(string) error
	Size() dxui.Size
	SetSize(float32, float32) error
	Maximize() error
	Unmaximize() error
	Minimize() error
	Unminimize() error
	IsMaximized() bool
	IsMinimized() bool
	Invalidate() error
	Diagnostics() dxui.RuntimeDiagnostics
}

func controls(target func() windowControl, message *string) dxui.View {
	action := func(label string, run func() error) dxui.View {
		return dxui.TextButton(
			dxui.ButtonProps{
				OnPress: func() {
					if err := run(); err != nil {
						*message = err.Error()
						return
					}
					*message = "Requested: " + label
				},
			},
			label,
		)
	}
	row := func(children ...dxui.View) dxui.View {
		return dxui.Box(dxui.BoxProps{
			Direction: dxui.Horizontal,
			Gap:       8,
		}, children...)
	}
	return dxui.Box(
		dxui.BoxProps{
			Gap: 12,
		},
		row(
			action("Set title", func() error {
				return target().SetTitle("Updated title")
			}),
			action("Set size", func() error {
				return target().SetSize(720, 600)
			}),
		),
		row(
			action("Maximize", func() error {
				return target().Maximize()
			}),
			action("Unmaximize", func() error {
				return target().Unmaximize()
			}),
		),
		row(
			action("Minimize", func() error {
				return target().Minimize()
			}),
			action("Unminimize", func() error {
				return target().Unminimize()
			}),
		),
		row(
			action("Refresh root", func() error {
				return target().Invalidate()
			}),
			dxui.TextButton(
				dxui.ButtonProps{
					OnPress: func() {
						w := target()
						*message = fmt.Sprintf(
							"%s | size=%v | maximized=%t | minimized=%t",
							w.Title(),
							w.Size(),
							w.IsMaximized(),
							w.IsMinimized(),
						)
					},
				},
				"Read confirmed state",
			),
		),
		dxui.TextButton(
			dxui.ButtonProps{
				OnPress: func() {
					d := target().Diagnostics()
					*message = fmt.Sprintf("Builds: %d, paints: %d", d.BuildCount, d.PaintCount)
				},
			},
			"Snapshot",
		),
		dxui.Text(dxui.TextProps{
			Value: *message,
			Wrap:  dxui.TextWrapWords,
		}),
	)
}

func main() {
	mainMessage, childMessage := "Ready", "Ready"
	var child *dxui.Window
	app := dxui.NewApp(dxui.AppOptions{
		Title:       "Window controls",
		Width:       720,
		Height:      600,
		Diagnostics: true,
		Background:  dxui.RGBA(248, 250, 252, 255),
	})
	open := func() {
		if child != nil && !child.Closed() {
			mainMessage = "Child is already open"
			return
		}
		childMessage = "Ready"
		window, err := app.CreateWindow(
			dxui.WindowOptions{
				Title:  "Child controls",
				Width:  720,
				Height: 600,
			},
			func() dxui.View {
				return dxui.Box(
					dxui.BoxProps{
						Style: dxui.Style{
							Padding: dxui.Padding(24),
						},
					},
					controls(func() windowControl {
						return child
					}, &childMessage),
				)
			},
		)
		if err != nil {
			mainMessage = err.Error()
			return
		}
		child = window
		mainMessage = "Child opened"
	}
	if err := app.Run(func() dxui.View {
		return dxui.Box(
			dxui.BoxProps{
				Gap: 16,
				Style: dxui.Style{
					Padding: dxui.Padding(24),
				},
			},
			dxui.TextButton(dxui.ButtonProps{
				OnPress: open,
			}, "Open child"),
			dxui.TextButton(
				dxui.ButtonProps{
					OnPress: func() {
						if child == nil || child.Closed() {
							mainMessage = "Open a child first"
							return
						}
						if err := child.Unminimize(); err != nil {
							mainMessage = err.Error()
						}
					},
				},
				"Restore child",
			),
			controls(func() windowControl {
				return app
			}, &mainMessage),
		)
	}); err != nil {
		log.Fatal(err)
	}
}
