package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/dxui-org/dxui"
)

type childState struct {
	window     *dxui.Window
	count      int
	text       string
	allowClose bool
	message    string
}

func report(err error) {
	if err != nil && !errors.Is(err, dxui.ErrWindowClosed) && !errors.Is(err, dxui.ErrAppClosed) {
		log.Print(err)
	}
}

func main() {
	app := dxui.NewApp(dxui.AppOptions{
		Title:      "Multi-window",
		Width:      640,
		Height:     420,
		Background: dxui.RGBA(248, 250, 252, 255),
	})
	message := "Create independent child windows"
	children := []*childState{}
	create := func() {
		number := len(children) + 1
		state := &childState{
			allowClose: true,
		}
		requestClose := func(w *dxui.Window) {
			if !state.allowClose {
				state.message = "Enable Allow close first"
				return
			}
			w.Close()
			report(app.Update(func() {
				message = fmt.Sprintf("Child %d closed", number)
			}))
		}
		window, err := app.CreateWindow(
			dxui.WindowOptions{
				Title:          fmt.Sprintf("Child %d", number),
				Width:          420,
				Height:         400,
				OnCloseRequest: requestClose,
				OnShown: func(w *dxui.Window) {
					log.Printf("Shown: %s", w.Title())
				},
			},
			func() dxui.View {
				return dxui.Box(
					dxui.BoxProps{
						Gap: 12,
						Style: dxui.Style{
							Padding: dxui.Padding(24),
						},
					},
					dxui.Label(fmt.Sprintf("Count: %d", state.count)),
					dxui.Input(dxui.InputProps{
						Key:         "notes",
						Value:       state.text,
						Placeholder: "Independent input",
						OnChange:    dxui.Assign(&state.text),
					}),
					dxui.TextButton(
						dxui.ButtonProps{
							OnPress: func() {
								state.count++
								report(app.Update(func() {
									message = fmt.Sprintf("Child %d changed", number)
								}))
							},
						},
						"Increment this child",
					),
					dxui.Checkbox(
						dxui.CheckboxProps{
							Checked:  state.allowClose,
							OnChange: dxui.Assign(&state.allowClose),
						},
						dxui.Label("Allow close"),
					),
					dxui.Label(state.message),
					dxui.TextButton(
						dxui.ButtonProps{
							OnPress: func() {
								requestClose(state.window)
							},
						},
						"Close this child",
					),
				)
			},
		)
		if err != nil {
			message = err.Error()
			return
		}
		state.window = window
		children = append(children, state)
		message = fmt.Sprintf("Created child %d", number)
	}
	updateNewest := func() {
		for i := len(children) - 1; i >= 0; i-- {
			state := children[i]
			if !state.window.Closed() {
				report(state.window.Update(func() {
					state.count++
				}))
				return
			}
		}
	}
	if err := app.Run(func() dxui.View {
		return dxui.Box(
			dxui.BoxProps{
				Gap: 16,
				Style: dxui.Style{
					Padding: dxui.Padding(24),
				},
			},
			dxui.Label(message),
			dxui.TextButton(dxui.ButtonProps{
				OnPress: create,
			}, "Create child"),
			dxui.TextButton(
				dxui.ButtonProps{
					OnPress: updateNewest,
				},
				"Update newest child",
			),
			dxui.TextButton(dxui.ButtonProps{
				OnPress: app.Close,
			}, "Close all windows"),
		)
	}); err != nil {
		log.Fatal(err)
	}
}
