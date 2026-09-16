package main

import (
	"github.com/dxui-org/dxui"
	"log"
	"strings"
)

func main() {
	name, language, message := "", "go", "Fill in the form"
	accepted := false
	valid := func() bool {
		return strings.TrimSpace(name) != "" && accepted
	}
	submit := func() {
		if valid() {
			message = "Saved: " + name + " / " + language
		}
	}
	app := dxui.NewApp(dxui.AppOptions{
		Background: dxui.RGBA(248, 250, 252, 255),
		Title:      "Preferences",
		Width:      640,
		Height:     480,
	})
	if err := app.Run(func() dxui.View {
		return dxui.Box(
			dxui.BoxProps{
				Gap: 16,
				Style: dxui.Style{
					Padding: dxui.Padding(24),
					Width:   dxui.Px(400),
				},
			},
			dxui.Input(dxui.InputProps{
				Key:         "name",
				Value:       name,
				Placeholder: "Name",
				OnChange:    dxui.Assign(&name),
				OnSubmit:    submit,
			}),
			dxui.Select(dxui.SelectProps{
				Value:    language,
				OnChange: dxui.Assign(&language),
				Options: []dxui.SelectOption{
					{
						Value: "go",
						Label: "Go",
					},
					{
						Value: "rust",
						Label: "Rust",
					},
				},
			}),
			dxui.Checkbox(
				dxui.CheckboxProps{
					Checked:  accepted,
					OnChange: dxui.Assign(&accepted),
				},
				dxui.Label("Accept terms"),
			),
			dxui.TextButton(
				dxui.ButtonProps{
					Disabled: !valid(),
					OnPress:  submit,
				},
				"Save",
			),
			dxui.Label(message),
		)
	}); err != nil {
		log.Fatal(err)
	}
}
