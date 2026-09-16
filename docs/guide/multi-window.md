# Multiple windows and independent updates

Open children from the main window and change each independently. Every child has its own input and counter.

## Complete example

<<< ../examples/multi-window/main.go

## 1. Open two child windows

Click **Create child** twice, enter different text in each, and click **Increment this child**. Each window keeps independent text and numbers.

Use app.CreateWindow: the first argument sets title and size, and the second describes the content. Call it inside a button callback, not a builder or background goroutine.

## 2. Update a child from the main window

Click **Update newest child** in the main window to increment the most recently opened child that is still alive.

| Update target | How |
| --- | --- |
| Current window component callback | Modify that window's state directly. |
| Main window | Wrap state changes in app.Update. |
| A child window | Wrap state changes in that window.Update. |

Sharing a Go variable does not automatically refresh every window. The example uses app.Update from a child to update the main-window message.

## 3. Close windows

Uncheck **Allow close** in a child, then attempt to close it: the window remains open. Check it again to allow closure.

Window.Close affects one child. App.Close or closing the main window ends the whole app without asking each child. Handle update errors because the target may have closed.

::: details Two details about window creation

The first build occurs before CreateWindow returns. Do not immediately read an outer handle that has not been assigned yet; the example uses it only in later button callbacks.

OnShown also runs before CreateWindow returns. Use the *Window argument passed to the callback.

:::

App.SetTheme changes themes across windows; child shortcuts need separate configuration. See [window management](/api/windows) for options, refresh methods, size, and maximization.

Next: [Troubleshooting](./troubleshooting).
