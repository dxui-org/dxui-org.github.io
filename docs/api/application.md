# Application, events, and diagnostics API

Application APIs are in `github.com/dxui-org/dxui`. This page explains their use and constraints; see [runtime declarations](/api/runtime) for exact signatures, fields, and source comments.

| API | Parameters and results | Requirements and examples |
| --- | --- | --- |
| NewApp(AppOptions) *App | Copies configuration such as fonts and shortcuts; returns an App managing the main and child windows. | Does not create a window; see [getting started](/guide/getting-started). |
| App.Run(func() View) error | A parameterless root returns a valid View; blocks until shutdown or failure. | Call directly on the main thread; each App runs once. Nil roots are rejected. |
| App.RunResponsive(func(LayoutContext) View) error | Context supplies logical Width/Height. | Rebuilds after coalesced resize/scale events; [responsive example](/guide/layout). |
| App.Update(func()) error | Queues main-window state changes in FIFO order on the UI thread; rebuilds after a batch. | Requires a running App and a non-nil closure; [background tasks](/guide/async). |
| App.Invalidate() error | Requests a main-window rebuild. | Safe across goroutines; use Window.Invalidate for children. [Multi-window updates](/guide/multi-window). |
| App.CreateWindow(WindowOptions, func() View) (*Window, error) | Creates an independent top-level child window, shown after its first frame succeeds. | Call in a UI callback or App.Update; [multiple windows](/guide/multi-window). |
| App.Close() | Requests shutdown of the main window and all children; no result. | Concurrent-safe and idempotent; a no-op before Run. |
| App.SetTheme(Theme) error | Validates, copies, and propagates the theme to live windows. | During execution, call in a UI callback or Update; [theme example](/guide/style). |
| App.SetClipboardText(string) error | Writes UTF-8 text; invalid UTF-8 becomes U+FFFD. | UI callback/Update only; reports unavailable lifecycle or native clipboard errors. |
| App.Diagnostics() RuntimeDiagnostics | Concurrent-safe main-window snapshot; use Window.Diagnostics for children. | Available during or after execution; see the metrics below. |
| `Assign[T](*T) func(T)` | Creates an assignment callback; a nil pointer panics immediately. | For UI callbacks; does not replace Update. |
| `Some[T](T) Option[T]` | Explicitly supplies a value, including zero. | Zero Option means unset; there is no public getter. |
| ErrAppNotRunning / ErrAppClosed | Lifecycle errors matched with errors.Is. | [Background task shutdown](/guide/async). |

App and Window both expose title, size, maximize/minimize, and state queries. See [window management](/api/windows) for all parameters, thread requirements, and a complete program.

## AppOptions fields

| Field | Purpose, defaults, and constraints |
| --- | --- |
| Title | Window title; an empty string defaults to dxui. |
| Width / Height | Initial logical size; zero defaults to 800 / 600. Finite, nonnegative, at most MaxInt32; explicit positive values must round to at least 1. |
| MinWidth / MinHeight | Minimum logical size; 0 leaves it unspecified. Run reports invalid dimensions. |
| Renderer | Defaults to RendererAuto; RendererSoftware requests the named software renderer. |
| Background | Raw RGBAColor clear color, constructed with RGBA. All-zero defaults to RGBA(28,30,36,255). Tutorials explicitly choose a light background. |
| Caches | CacheBudgets; 0 selects defaults, negative values disable the corresponding cache. Budgets apply separately to each window. |
| Fonts | Font descriptions from FontBytes/FontFile/SystemFont; loaded at startup, with errors returned by Run. |
| DefaultFont | Default FontFamily; use with registered fonts. |
| DisableSystemFontFallback | Defaults to false; true disables automatic system CJK fallback without changing registered fonts. |
| Theme | Zero Theme selects the built-in light theme; a partially populated theme is not automatically completed. |
| Shortcuts | Copied main-window shortcuts; children configure WindowOptions.Shortcuts separately. |
| OnCloseRequest | Main-window close request; nil closes all windows. A callback decides whether to call App.Close. |
| OnError | Observes recoverable build/callback failures on the UI thread; nil terminates Run with the error. |
| OnShown | Called once after the main window's complete first frame is presented and shown; children use WindowOptions.OnShown. |
| Diagnostics | Defaults to false; true enables bounded event, timing, and resource metrics. |

## Complete shortcuts, clipboard, shutdown, and diagnostics example

<<< ../examples/lifecycle/main.go

Shortcut.Key accepts only the public ShortcutKey values: Enter, Backspace, 0..9, Plus, Minus, Multiply, Divide, Decimal, and Equals. Arbitrary key strings are unsupported. Shift/Control/Alt/Super/Primary default to false and match exactly. OnPress takes no arguments; Repeat defaults to false. Editors and built-in control keys take priority, so the example uses Primary+1 rather than Enter.

## CacheBudgets and RuntimeDiagnostics

Budgets and CPU caches are independent per window, so total capacity grows with the number of live windows. Window.Diagnostics describes a child; App.Diagnostics is not an aggregate.

| Cache field | Default budget | Meaning |
| --- | --- | --- |
| FontBytes | 32 MiB | Application and automatic CJK font sources; negative values disallow application fonts. |
| TextSourceBytes | 2 MiB | Retained text/icon mask sources; negative values disallow these retained sources. |
| GlyphBytes | 2 MiB | Glyph cache. |
| TextMeasureBytes | 1 MiB | Text measurement cache. |
| ImageBytes | 2 MiB | Inactive reusable CPU images/native textures; unique visible images form a working set not strictly capped by this budget. |
| ShadowBytes | 2 MiB | Bounded shadow geometry cache. |

RendererName/SoftwareFallback identify the rendering mode. LogicalSize/PixelSize/PixelDensity/DisplayScale distinguish logical dimensions and pixel scaling. FrameCount, BuildCount, LayoutCount, PaintCount, ReconcileCount, and PaintNodeCount distinguish presentation, building, layout, painting, and tree work.

WindowCreates, RendererCreateAttempts/Creates, ExposeEvents, ResizeEvents, ScaleEvents, NoopViewportEvents, and RendererResetEvents track lifecycle and viewport events; EventCount counts events. TextureCreates/Destroys, CacheBytes/BudgetBytes/Entries, and FontResources/ImageResources/RendererResources describe resources. Goroutines, GoHeapBytes/Objects, GoTotalAllocBytes, and GoMallocs are process-wide Go metrics, not memory owned by one component.

EventToPresent and FrameTime use TimingSummary: Count is the total count, Samples is a bounded sample count, and P50NS/P95NS/P99NS are nanoseconds. Check CountersEnabled first; disabled or unavailable metrics and absent samples are not zero-cost measurements. See [RuntimeDiagnostics](/api/runtime#api-runtimediagnostics) for field types.
