# Application lifecycle: Full declarations

[Usage and examples](/api/application) · [All APIs](/api/)

These declarations are extracted from the Go AST, retaining English source contracts and field comments while hiding private fields. `struct {}` represents an opaque implementation, not permission to use a zero value instead of its constructor. Constant blocks retain iota context.

## ErrAppNotRunning {#api-errappnotrunning}

```text
ErrAppNotRunning reports an Update attempted outside App.Run.
```


```go
var // ErrAppNotRunning reports an Update attempted outside App.Run.
ErrAppNotRunning = errors.New("dxui: app is not running")
```

[Related example](/api/application) · Source `app.go:27`

## ErrAppClosed {#api-errappclosed}

```text
ErrAppClosed reports an operation submitted after shutdown was requested.
```


```go
var // ErrAppClosed reports an operation submitted after shutdown was requested.
ErrAppClosed = errors.New("dxui: app is closing")
```

[Related example](/api/application) · Source `app.go:29`

## RuntimeDiagnostics {#api-runtimediagnostics}

```text
RuntimeDiagnostics is a backend-neutral snapshot of the running or most
recently stopped native runtime.
```


```go
type RuntimeDiagnostics struct {
	RendererName           string
	LogicalSize            Size
	PixelSize              Size
	PixelDensity           float32
	DisplayScale           float32
	FrameCount             uint64
	WindowCreates          uint64
	RendererCreateAttempts uint64
	RendererCreates        uint64
	ExposeEvents           uint64
	ResizeEvents           uint64
	ScaleEvents            uint64
	NoopViewportEvents     uint64
	RendererResetEvents    uint64
	BuildCount             uint64
	LayoutCount            uint64
	PaintCount             uint64
	SoftwareFallback       bool
	CountersEnabled        bool
	EventCount             uint64
	ReconcileCount         uint64
	PaintNodeCount         uint64
	TextureCreates         uint64
	TextureDestroys        uint64
	CacheBytes             uint64
	CacheBudgetBytes       uint64
	CacheEntries           uint64
	FontResources          uint64
	ImageResources         uint64
	RendererResources      uint64
	Goroutines             int
	GoHeapBytes            uint64
	GoHeapObjects          uint64
	GoTotalAllocBytes      uint64
	GoMallocs              uint64
	EventToPresent         TimingSummary
	FrameTime              TimingSummary
}
```

[Related example](/api/application) · Source `app.go:41`

## TimingSummary {#api-timingsummary}

```text
TimingSummary reports a bounded percentile distribution in nanoseconds.
```


```go
type TimingSummary struct {
	Count, Samples      uint64
	P50NS, P95NS, P99NS int64
}
```

[Related example](/api/application) · Source `app.go:83`

## App {#api-app}

```text
App owns one application runtime, one main window and any child windows. An
App is single-use: Run may be called exactly once.
```


```go
type App struct {
}
```

[Related example](/api/application) · Source `app.go:90`

## LayoutContext {#api-layoutcontext}

```text
LayoutContext is the logical space available to the root builder. It
contains no backend values. A constraint-aware build runs once initially
and once for the final resize/scale event in each drained event batch.
```


```go
type LayoutContext struct{ Width, Height float32 }
```

[Related example](/api/application) · Source `app.go:152`

## NewApp {#api-newapp}

```text
NewApp creates an application configuration without starting the native runtime.
```


```go
func NewApp(options AppOptions) *App
```

[Related example](/api/application) · Source `app.go:155`

## App.RunResponsive {#api-app-runresponsive}

```text
RunResponsive is Run with a root builder that may choose a different view
structure from the current logical window size. Measuring a result never
invokes the builder again; only a later coalesced viewport event can do so.
```


```go
func (a *App) RunResponsive(root func(LayoutContext) View) error
```

[Related example](/api/application) · Source `app.go:181`

## App.Run {#api-app-run}

```text
Run creates the native window, builds root, and owns the process main thread
until the app closes. Call it directly from main, before moving UI work to
other goroutines. Run is blocking and may be called only once, including
after a startup or runtime error. A nil root is rejected before the App is
consumed. Startup, build, renderer, callback, and event-loop failures are
returned; OnError also observes runtime failures when configured.
```


```go
func (a *App) Run(root func() View) (runErr error)
```

[Related example](/api/application) · Source `app.go:206`

## App.SetTheme {#api-app-settheme}

```text
SetTheme validates and copies a complete Primitive -> Semantic -> Component
theme atomically. Equal themes are a no-op. A relevant metric-token change
schedules layout; visual-only resolved changes schedule display/paint only.
While Run is active, call SetTheme from a UI callback or inside Update.
```


```go
func (a *App) SetTheme(source Theme) error
```

[Related example](/api/application) · Source `app.go:411`

## App.SetClipboardText {#api-app-setclipboardtext}

```text
SetClipboardText writes UTF-8 text to the system clipboard. While Run is
active, call it from a UI callback or inside Update so the native operation
remains on the UI thread. Invalid UTF-8 is normalized to replacement runes.
```


```go
func (a *App) SetClipboardText(value string) error
```

[Related example](/api/application) · Source `app.go:499`

## App.Close {#api-app-close}

```text
Close requests application shutdown and safely wakes a blocked event wait.
It may be called from callbacks or any goroutine. Repeated calls and calls
made before Run are no-ops.
```


```go
func (a *App) Close()
```

[Related example](/api/application) · Source `app.go:529`

## App.Update {#api-app-update}

```text
Update queues a state update for FIFO execution on the UI thread, then
rebuilds the root once after the batch. It never executes update on the
caller's goroutine and is safe to call from any goroutine.
```


```go
func (a *App) Update(update func()) error
```

[Related example](/api/application) · Source `app.go:548`

## App.Invalidate {#api-app-invalidate}

```text
Invalidate requests a rebuild of the main window only. It is safe from any
goroutine; child windows use Window.Invalidate.
```


```go
func (a *App) Invalidate() error
```

[Related example](/api/windows) · Source `app.go:574`

## App.Diagnostics {#api-app-diagnostics}

```text
Diagnostics returns a race-safe, backend-independent runtime snapshot.
```


```go
func (a *App) Diagnostics() RuntimeDiagnostics
```

[Related example](/api/application) · Source `app.go:597`

## Assign {#api-assign}

```text
Assign returns a callback that stores its argument in target. It is intended
for simple controlled UI callbacks. Assign panics when target is nil; it does
not schedule work or replace App.Update for cross-goroutine mutation.
```


```go
func Assign[T any](target *T) func(T)
```

[Related example](/api/application) · Source `assign.go:6`

## Option {#api-option}

```text
Option distinguishes an explicitly supplied zero value from an unset value.
```


```go
type Option[T any] struct {
}
```

[Related example](/api/application) · Source `dxui.go:16`

## Some {#api-some}

```text
Some creates a set option.
```


```go
func Some[T any](value T) Option[T]
```

[Related example](/api/application) · Source `dxui.go:22`

## Point {#api-point}

```text
Point is a position or offset in logical units.
```


```go
type Point = icondata.Point
```

[Related example](/api/application) · Source `dxui.go:27`

## Size {#api-size}

```text
Size is a width and height in logical units.
```


```go
type Size struct{ Width, Height float32 }
```

[Related example](/api/application) · Source `dxui.go:30`

## RGBAColor {#api-rgbacolor}

```text
RGBAColor is an 8-bit non-premultiplied RGBA color.

The former Color type name is now the discoverable color-token namespace.
```


```go
type RGBAColor struct{ R, G, B, A uint8 }
```

[Related example](/api/application) · Source `dxui.go:35`

## RGBA {#api-rgba}

```text
RGBA creates an 8-bit non-premultiplied color.
```


```go
func RGBA(r, g, b, a uint8) RGBAColor
```

[Related example](/api/application) · Source `dxui.go:38`

## RendererPreference {#api-rendererpreference}

```text
RendererPreference selects the preferred renderer creation policy.
```


```go
type RendererPreference uint8
```

[Related example](/api/application) · Source `dxui.go:41`

## RendererAuto {#api-rendererauto}


```go
const (
	RendererAuto RendererPreference = iota
	RendererSoftware
)
```

[Related example](/api/application) · Source `dxui.go:44`

## RendererSoftware {#api-renderersoftware}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:45`

## CacheBudgets {#api-cachebudgets}

```text
CacheBudgets bounds CPU and renderer-owned text, icon, and image resources. Zero
selects defaults: FontBytes 32 MiB, TextSourceBytes 2 MiB, GlyphBytes 2 MiB,
TextMeasureBytes 1 MiB, ImageBytes 2 MiB, and ShadowBytes 2 MiB. ImageBytes
bounds inactive reusable CPU pixels and renderer textures. Unique images in
the committed display are working-set resources charged at four bytes per
source pixel until that display releases them. A negative cache budget
disables that cache; negative FontBytes or TextSourceBytes permits no
application fonts or retained text/icon masks respectively.
```


```go
type CacheBudgets struct {
	FontBytes        int
	TextSourceBytes  int
	GlyphBytes       int
	TextMeasureBytes int
	ImageBytes       int
	ShadowBytes      int
}
```

[Related example](/api/application) · Source `dxui.go:56`

## AppOptions {#api-appoptions}

```text
AppOptions configures an App's main window and shared runtime. Its
zero value selects documented window, renderer, cache, font, and theme
defaults; invalid dimensions, renderer values, fonts, or themes are reported
by App.Run before native event processing begins.
```


```go
type AppOptions struct {
	Title               string
	Width, Height       float32
	MinWidth, MinHeight float32
	Renderer            RendererPreference
	Background          RGBAColor
	Caches              CacheBudgets
	Fonts               []Font
	DefaultFont         FontFamily
	// DisableSystemFontFallback prevents lazy deterministic system-CJK font
	// loading. The zero value enables fallback after all application fonts and
	// dxui's built-in Latin font.
	DisableSystemFontFallback bool
	Theme                     Theme
	Shortcuts                 []Shortcut
	// OnCloseRequest handles a native window-close request on the UI thread.
	// A nil callback closes the App. A non-nil callback must call Close when it
	// accepts the request.
	OnCloseRequest func(*App)
	// OnError observes recoverable build and callback failures on the UI
	// thread. When nil, the failure terminates Run and is returned.
	OnError func(error)
	// OnShown runs once on the UI thread after the complete first frame was
	// presented and the native window was shown successfully. It is not called
	// after startup failure or on later builds/presents.
	OnShown func(*App)
	// Diagnostics enables bounded event/timing/resource counters.
	// It is false by default so release event and render paths avoid the work.
	Diagnostics bool
}
```

[Related example](/api/application) · Source `dxui.go:69`

## WindowOptions {#api-windowoptions}

```text
WindowOptions configures an independent child window. Fonts, theme, rendering and cache settings come from the owner App. Shortcuts are configured separately for each window.
```


```go
type WindowOptions struct {
	Title               string
	Width, Height       float32
	MinWidth, MinHeight float32
	Background          RGBAColor
	Shortcuts           []Shortcut
	// OnCloseRequest may reject a native close request by returning without
	// calling Window.Close. A nil callback accepts the request.
	OnCloseRequest func(*Window)
	// OnShown runs once after the complete first frame is presented and the
	// hidden native window has been shown successfully.
	OnShown func(*Window)
}
```

[Related example](/api/windows) · Source `dxui.go:103`

## ShortcutKey {#api-shortcutkey}

```text
ShortcutKey is a backend-neutral semantic key used by an App shortcut.
```


```go
type ShortcutKey uint8
```

[Related example](/api/application) · Source `dxui.go:118`

## KeyEnter {#api-keyenter}


```go
const (
	KeyEnter ShortcutKey = iota + 1
	KeyBackspace
	Key0
	Key1
	Key2
	Key3
	Key4
	Key5
	Key6
	Key7
	Key8
	Key9
	KeyPlus
	KeyMinus
	KeyMultiply
	KeyDivide
	KeyDecimal
	KeyEquals
)
```

[Related example](/api/application) · Source `dxui.go:121`

## KeyBackspace {#api-keybackspace}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:122`

## Key0 {#api-key0}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:123`

## Key1 {#api-key1}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:124`

## Key2 {#api-key2}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:125`

## Key3 {#api-key3}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:126`

## Key4 {#api-key4}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:127`

## Key5 {#api-key5}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:128`

## Key6 {#api-key6}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:129`

## Key7 {#api-key7}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:130`

## Key8 {#api-key8}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:131`

## Key9 {#api-key9}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:132`

## KeyPlus {#api-keyplus}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:133`

## KeyMinus {#api-keyminus}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:134`

## KeyMultiply {#api-keymultiply}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:135`

## KeyDivide {#api-keydivide}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:136`

## KeyDecimal {#api-keydecimal}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:137`

## KeyEquals {#api-keyequals}

See the corresponding constant block on this page for types and values.

[Related example](/api/application) · Source `dxui.go:138`

## ShortcutModifiers {#api-shortcutmodifiers}

```text
ShortcutModifiers are matched exactly. Primary substitutes for Command on
macOS and Control elsewhere; callers do not also set that physical field.
```


```go
type ShortcutModifiers struct{ Shift, Control, Alt, Super, Primary bool }
```

[Related example](/api/application) · Source `dxui.go:143`

## Shortcut {#api-shortcut}

```text
Shortcut binds one application-window key chord to a semantic action.
Focused editors and built-in control keys have priority. Repeat enables
repeated key-down activation; otherwise native repeat is consumed silently.
```


```go
type Shortcut struct {
	Key       ShortcutKey
	Modifiers ShortcutModifiers
	Repeat    bool
	OnPress   func()
}
```

[Related example](/api/application) · Source `dxui.go:148`

## ErrWindowClosed {#api-errwindowclosed}

```text
ErrWindowClosed reports an operation on a child window after close.
```


```go
var // ErrWindowClosed reports an operation on a child window after close.
ErrWindowClosed = errors.New("dxui: window is closed")
```

[Related example](/api/windows) · Source `window.go:15`

## ErrWindowNotRunning {#api-errwindownotrunning}

```text
ErrWindowNotRunning reports child-window creation outside App.Run.
```


```go
var // ErrWindowNotRunning reports child-window creation outside App.Run.
ErrWindowNotRunning = errors.New("dxui: window runtime is not running")
```

[Related example](/api/windows) · Source `window.go:17`

## Window {#api-window}

```text
Window is an opaque, concurrency-safe handle to a child native window.
Component construction remains on the root dxui API; a Window only owns
window-level lifecycle and scheduling operations.
```


```go
type Window struct {
}
```

[Related example](/api/windows) · Source `window.go:23`

## App.CreateWindow {#api-app-createwindow}

```text
CreateWindow creates a hidden native child window and prepares its complete
first frame. It must be called from a UI callback or App.Update closure.
Failure releases every partially created child resource and leaves existing
windows unchanged.
```


```go
func (a *App) CreateWindow(options WindowOptions, root func() View) (*Window, error)
```

[Related example](/api/windows) · Source `window.go:49`

## Window.Update {#api-window-update}

```text
Update queues a child-window state mutation and rebuilds only that window.
```


```go
func (w *Window) Update(update func()) error
```

[Related example](/api/windows) · Source `window.go:167`

## Window.Invalidate {#api-window-invalidate}

```text
Invalidate requests a rebuild of only this child window.
```


```go
func (w *Window) Invalidate() error
```

[Related example](/api/windows) · Source `window.go:187`

## Window.SetTitle {#api-window-settitle}

```text
SetTitle changes the native title on the UI thread. It may be called from a
UI callback or App.Update closure and does not rebuild the root.
```


```go
func (w *Window) SetTitle(title string) error
```

[Related example](/api/windows) · Source `window.go:191`

## Window.SetSize {#api-window-setsize}

```text
SetSize requests a new positive logical client size. The resulting native
viewport event drives layout; it is not presented speculatively.
```


```go
func (w *Window) SetSize(width, height float32) error
```

[Related example](/api/windows) · Source `window.go:207`

## Window.Title {#api-window-title}

```text
Title returns the last successfully configured title.
```


```go
func (w *Window) Title() string
```

[Related example](/api/windows) · Source `window.go:234`

## Window.Size {#api-window-size}

```text
Size returns the latest known logical client size. Native resize events and
successful SetSize calls update this snapshot.
```


```go
func (w *Window) Size() Size
```

[Related example](/api/windows) · Source `window.go:245`

## Window.Maximize {#api-window-maximize}

```text
Maximize requests the platform's maximized window state.
```


```go
func (w *Window) Maximize() error
```

[Related example](/api/windows) · Source `window.go:255`

## Window.Unmaximize {#api-window-unmaximize}

```text
Unmaximize restores a maximized window to its normal state.
```


```go
func (w *Window) Unmaximize() error
```

[Related example](/api/windows) · Source `window.go:267`

## Window.Minimize {#api-window-minimize}

```text
Minimize requests the platform's minimized window state.
```


```go
func (w *Window) Minimize() error
```

[Related example](/api/windows) · Source `window.go:270`

## Window.Unminimize {#api-window-unminimize}

```text
Unminimize restores a minimized window to its normal state.
```


```go
func (w *Window) Unminimize() error
```

[Related example](/api/windows) · Source `window.go:282`

## Window.IsMaximized {#api-window-ismaximized}

```text
IsMaximized reports the latest state confirmed by native window events.
```


```go
func (w *Window) IsMaximized() bool
```

[Related example](/api/windows) · Source `window.go:295`

## Window.IsMinimized {#api-window-isminimized}

```text
IsMinimized reports the latest state confirmed by native window events.
```


```go
func (w *Window) IsMinimized() bool
```

[Related example](/api/windows) · Source `window.go:305`

## App.Title {#api-app-title}

```text
Main-window counterparts preserve App as the main lifecycle handle.
```


```go
func (a *App) Title() string
```

[Related example](/api/windows) · Source `window.go:345`

## App.Size {#api-app-size}


```go
func (a *App) Size() Size
```

[Related example](/api/windows) · Source `window.go:355`

## App.SetTitle {#api-app-settitle}


```go
func (a *App) SetTitle(title string) error
```

[Related example](/api/windows) · Source `window.go:365`

## App.SetSize {#api-app-setsize}


```go
func (a *App) SetSize(width, height float32) error
```

[Related example](/api/windows) · Source `window.go:372`

## App.Maximize {#api-app-maximize}


```go
func (a *App) Maximize() error
```

[Related example](/api/windows) · Source `window.go:379`

## App.Unmaximize {#api-app-unmaximize}


```go
func (a *App) Unmaximize() error
```

[Related example](/api/windows) · Source `window.go:386`

## App.Minimize {#api-app-minimize}


```go
func (a *App) Minimize() error
```

[Related example](/api/windows) · Source `window.go:393`

## App.Unminimize {#api-app-unminimize}


```go
func (a *App) Unminimize() error
```

[Related example](/api/windows) · Source `window.go:400`

## App.IsMaximized {#api-app-ismaximized}


```go
func (a *App) IsMaximized() bool
```

[Related example](/api/windows) · Source `window.go:407`

## App.IsMinimized {#api-app-isminimized}


```go
func (a *App) IsMinimized() bool
```

[Related example](/api/windows) · Source `window.go:413`

## Window.Close {#api-window-close}

```text
Close closes this child window without affecting its owner or siblings.
Repeated calls are no-ops.
```


```go
func (w *Window) Close()
```

[Related example](/api/windows) · Source `window.go:434`

## Window.Diagnostics {#api-window-diagnostics}

```text
Diagnostics returns this window's backend-neutral counters.
```


```go
func (w *Window) Diagnostics() RuntimeDiagnostics
```

[Related example](/api/windows) · Source `window.go:455`

## Window.Closed {#api-window-closed}

```text
Closed reports whether close has been requested or completed.
```


```go
func (w *Window) Closed() bool
```

[Related example](/api/windows) · Source `window.go:463`

