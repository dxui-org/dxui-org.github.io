# Background tasks and lifecycle

Run expensive computations in a goroutine and update the UI through `App.Update` when finished, keeping the window responsive.

## 1. Run the example

<<< ../examples/async/main.go

Click **Compute** to display the result. The button is disabled during calculation; fast tasks may finish before the Computing message is visible.

## 2. Handle tasks in three steps

1. **Prepare input on click**: Set `busy` and copy the required input into local variable `n`.
2. **Compute in the background**: The goroutine uses only the prepared input to produce `result`.
3. **Update the UI**: Change `message` and `busy` inside the `app.Update` closure.

Do not directly change variables read by the UI from a goroutine. `Assign` is only an assignment function, not a replacement for `Update`.

A nil result from `Update` means the work was queued, not executed. Submission fails after shutdown, so the example handles the returned error.

## 3. Stop work during shutdown

The example uses `context.WithCancel`. On a system close request, `OnCloseRequest` calls `cancel()` to stop background work, then `app.Close()` to exit.

When OnCloseRequest is set, explicitly call Close to accept closure. Calling cancel() again after Run returns covers other exit paths.

Try increasing the workload and closing the window during computation. The task should honor cancellation instead of submitting further results.

::: details Startup, shutdown, and error callbacks

NewApp does not create native windows. Run returns configuration, font, theme, and native startup errors. OnShown runs once after the complete first frame is presented and shown; use it for work requiring a ready window, not the root builder.

By default, a main-window system close request ends the application and closes all children. With `OnCloseRequest func(*App)`, call Close to accept; returning without it keeps the window open, useful for unsaved changes. App.Close is idempotent, concurrent-safe, and a no-op before Run. Window.Close affects only its child; see [multiple windows](./multi-window).

OnError observes recoverable build/callback failures on the UI thread; nil ends Run with the error. Log or correct the problematic state rather than relying on a persistently failing builder to recover. Zero Views and invalid item keys fail tree-update transactions, retaining the previous valid frame; the first build has no previous frame. Do not use OnError to ignore startup failures.


:::

See [application APIs](/api/application) for shortcuts, clipboard, and lifecycle parameters.

Next: [Open multiple windows](./multi-window).
