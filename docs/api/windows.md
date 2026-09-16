# Window management API

An App manages one main window and zero or more independent top-level child windows. Components remain in the root package; Window is a lifecycle handle, not a View, and should not be constructed as a zero value. Start with [multiple windows](/guide/multi-window); exact signatures are in [runtime references](/api/runtime).

## Creation and scheduling

| API | Parameters and results | Scope, constraints, and errors |
| --- | --- | --- |
| App.CreateWindow(options WindowOptions, root func() View) (*Window, error) | Configuration and non-nil root builder; returns a child handle on success. | Call in a UI callback or App.Update. Shows after preparing the first frame; failures clean up partial resources without changing existing windows. |
| App.Update(update func()) error | Queues non-nil state changes in FIFO order. | Main window only; safe across goroutines. Returns ErrAppNotRunning/ErrAppClosed outside its lifecycle. |
| App.Invalidate() error | Requests a main-window rebuild without changing state. | Concurrent-safe; lifecycle errors match App.Update. Does not synchronize background state writes. |
| Window.Update(update func()) error | Queues non-nil state changes in FIFO order. | Target child only; concurrent-safe. Closed/nil handles return ErrWindowClosed; shutdown may return ErrAppClosed. |
| Window.Invalidate() error | Requests a target child rebuild. | Equivalent to an empty update; lifecycle errors match Window.Update. |
| Window.Close() | Requests child shutdown; no result. | Idempotent and concurrent-safe. Leaves other windows open and bypasses OnCloseRequest. Nil is a no-op. |
| Window.Closed() bool | Reports requested or completed closure. | Concurrent-safe; nil returns true. Does not guarantee a subsequent operation succeeds. |
| Window.Diagnostics() RuntimeDiagnostics | Returns this window's metrics snapshot. | Concurrent-safe; nil returns a zero snapshot. [Fields](/api/application#cachebudgets-and-runtimediagnostics). |

App.Close closes every window and ends the shared event loop. App.Diagnostics still describes the main window rather than an aggregate. Do not submit work requiring a window after it closes.

## All WindowOptions fields

| Field | Type | Default and behavior |
| --- | --- | --- |
| Title | string | Empty defaults to dxui. |
| Width / Height | float32 | Defaults to 640 / 480 logical units, unlike the main window's 800 / 600. Finite, nonnegative, at most MaxInt32; explicit positive values must round to at least 1. |
| MinWidth / MinHeight | float32 | 0 leaves minima unspecified; finite, nonnegative, at most MaxInt32. |
| Background | RGBAColor | All-zero uses the App background, then RGBA(28,30,36,255) if still zero. |
| Shortcuts | []Shortcut | Independent copied shortcuts; empty by default, with no automatic inheritance from AppOptions.Shortcuts. |
| OnCloseRequest | func(*Window) | Nil accepts system close requests. A callback must call Window.Close to accept; otherwise it rejects. |
| OnShown | func(*Window) | Called once after the complete first frame is presented and shown, before CreateWindow returns. Use its handle argument. |

Theme, Fonts, DefaultFont, DisableSystemFontFallback, Renderer, Caches, and Diagnostics use App configuration. Windows have no individual SetTheme; App.SetTheme updates all live windows. Budgets apply per window, so total resource capacity grows with more windows.

The baseline WindowOptions source comment broadly claims shortcut inheritance, but creation explicitly replaces them with options.Shortcuts. This page and the displayed reference follow the implementation; the raw maintenance snapshot retains the original comment.

## Operations shared by main and child windows

Each method below exists on `*App` (main window) and `*Window` (target child). Native SetTitle, SetSize, Maximize, Unmaximize, Minimize, and Unminimize operations require a UI callback or Update closure. Title, Size, IsMaximized, and IsMinimized are concurrent-safe snapshots.

| Method | Result | Purpose and constraints |
| --- | --- | --- |
| Title() | string | Last successfully configured title; before Run, App returns the original option. Nil Window returns an empty string. |
| SetTitle(title string) | error | Updates the title snapshot on success without rebuilding the root. Empty strings are used literally, not replaced by the creation default. |
| Size() | Size | Logical client size, updated by successful SetSize and later resize/scale events. Before Run, App returns configured dimensions; nil Window returns zero Size. |
| SetSize(width, height float32) | error | Both dimensions must be finite, positive, round to at least 1, and not exceed MaxInt32. Success updates the snapshot; layout follows viewport events, and the platform may adjust the final size. |
| Maximize() | error | Requests maximization; nil does not mean a state event has confirmed it. |
| Unmaximize() | error | Requests restoration to normal state. |
| Minimize() | error | Requests minimization; nil does not mean a state event has confirmed it. |
| Unminimize() | error | Requests normal state; shares the platform restore operation with Unmaximize. |
| IsMaximized() | bool | State confirmed by the latest window event, not optimistically set by requests. False for an uncreated main window or nil Window. |
| IsMinimized() | bool | Minimized state confirmed by the latest window event. False for an uncreated main window or nil Window. |

Main-window commands return ErrWindowNotRunning before Run and ErrAppClosed after closure. Closed-child commands return ErrWindowClosed. Title/size/state getters remain snapshots; do not use them as liveness checks.

WindowOptions has no state-event callback. If displaying IsMaximized/IsMinimized, do not assume a request immediately changes them. The Read confirmed state button below reads confirmed state when clicked.

## Complete window controls example

This example provides the same title, size, maximize, and minimize buttons for main and child windows. The child first builds before CreateWindow returns, so it retrieves the handle through a function only during later button events.

<<< ../examples/window-controls/main.go

After minimizing, restore through the taskbar/window manager or the main window's Restore child button. Refresh root demonstrates targeted Invalidate; Snapshot demonstrates Diagnostics. These APIs do not create polling loops.

## Lifecycle errors

| Error | Scenario |
| --- | --- |
| ErrWindowNotRunning | CreateWindow outside App.Run, or main-window commands before Run. |
| ErrWindowClosed | Update, Invalidate, or commands on nil/closed children. |
| ErrAppClosed | Submitting work or executing main-window commands during shutdown. |
| Other errors | Nil builders/closures, invalid dimensions, shortcut validation, first-frame build failures, or platform failures. |

Use errors.Is for lifecycle errors and handle return values instead of matching text. Checks have an order: CreateWindow may return ErrWindowNotRunning first after runtime destruction. Not every shutdown-stage failure has the same error.
