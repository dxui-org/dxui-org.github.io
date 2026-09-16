package main

import (
	"github.com/dxui-org/dxui"
	"log"
)

func main() {
	message := "Primary+1 copies text"
	var app *dxui.App
	copyText := func() {
		if err := app.SetClipboardText("Hello dxui"); err != nil {
			message = err.Error()
		} else {
			message = "Copied"
		}
	}
	app = dxui.NewApp(dxui.AppOptions{
		Background:  dxui.RGBA(248, 250, 252, 255),
		Title:       "Lifecycle",
		Width:       640,
		Height:      400,
		Diagnostics: true,
		Shortcuts: []dxui.Shortcut{
			{
				Key: dxui.Key1,
				Modifiers: dxui.ShortcutModifiers{
					Primary: true,
				},
				OnPress: copyText,
			},
		},
		OnShown: func(a *dxui.App) {
			log.Printf("shown: %+v", a.Diagnostics().LogicalSize)
		},
		OnCloseRequest: func(a *dxui.App) {
			a.Close()
		},
	})
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
				OnPress: copyText,
			}, "Copy"),
		)
	}); err != nil {
		log.Fatal(err)
	}
	log.Printf("final diagnostics: %+v", app.Diagnostics())
}
