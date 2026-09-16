# State and events

Build a counter that increments each time its button is clicked.

## 1. Run the counter

Replace the previous main.go with this program and run `go run .`. Use the same procedure for later complete examples.

<<< ../examples/button/main.go

Click **Add** and watch the number change.

## 2. Understand the update

1. `count` stores the number outside `app.Run`.
2. `OnPress` executes `count++` on activation.
3. DXUI runs the builder again; Label reads the new number and updates the UI.

**Keep state outside builders and change it in callbacks.** Declaring count := 0 inside the builder resets it on every rebuild. Builders only describe the UI; do not start tasks or change state there.

Try replacing count++ with count += 2 and run again.

## 3. Retain text input

Inputs follow the same pattern: Value reads the variable, and OnChange writes the new text back.

```go
// Declare outside app.Run.
name := ""
```

```go
// Add this control to the Box returned by the builder.
dxui.Input(dxui.InputProps{
    Value: name,
    OnChange: func(next string) {
        name = next
    },
})
```

next is the complete new input value, so do not concatenate it. The callback can be shortened to `OnChange: dxui.Assign(&name)`. Without writing the new value back, the input cannot retain the edit.

Regular callbacks that change the current window need no manual refresh. Background work uses [App.Update](./async); for other windows see [multiple windows](./multi-window).

::: details Why list controls need Key

When inserting, deleting, or sorting controls, set Key to a stable business ID so selection and focus follow the right control. Do not use changing positions as IDs.

Keys are local to one parent; runtime state need not survive unmounting. Keep business data outside the control. See [common contracts](/api/common).

:::

See [Input](/components/input) for length limits, read-only behavior, and disabled state.

Next: [Arrange the UI](./layout).
