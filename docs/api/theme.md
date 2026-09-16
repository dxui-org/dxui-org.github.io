# Theme and token API

`LightTheme() Theme` and `DarkTheme() Theme` return independent, editable complete themes. `App.SetTheme(Theme) error` validates, copies, and propagates a theme to all live windows; call from a UI callback or Update while running. New children use the current App theme; Window has no SetTheme. [Complete theme switcher](/guide/style).

| Type/field | Purpose | Constraints and examples |
| --- | --- | --- |
| Theme.Primitive | PrimitiveTokens | Colors: map[ColorToken]RGBAColor; Metrics: map[MetricToken]float32. Base literals. |
| Theme.Semantic | SemanticTokens | Colors: map[ColorToken]ColorValue; Metrics: map[MetricToken]MetricValue. Map semantic roles to literals or other tokens. |
| Theme.Components | map[ComponentToken]ComponentTheme | Default component appearances; token names must be nonempty. |
| ComponentTheme.Base | StylePatch | Base component paint values. |
| ComponentTheme.States | StateStyles | Default/Hover/Focus/Checked/Pressed/Disabled patches. |
| ColorToken / MetricToken / ComponentToken | Distinct string-based types | Keys for their respective maps; do not interchange them or reference undefined tokens. |
| Color.Primitive.* | Discoverable palette keys | 26 families with 50, 100..900, and 950 shades, plus White/Black; see token declarations for all fields. |
| Color.Semantic.* | Surface, SurfaceHigh, Text, Accent, AccentHover, OnAccent, FocusRing, Danger, Success, Info, Warn, Border, Shadow | Customize by role rather than hardcoding a palette in each button. |

```go
// Fragment: app already exists; execute in a UI callback while running.
theme := dxui.LightTheme()
theme.Semantic.Colors[dxui.Color.Semantic.Accent] =
	dxui.TokenColor(dxui.Color.Primitive.Violet600)
theme.Semantic.Metrics[dxui.MetricSemanticControlHeight] = dxui.Metric(40)
if err := app.SetTheme(theme); err != nil {
	log.Print(err)
}
```

Component size tokens live in Semantic.Metrics. Check the built-in map before editing so a same-named key in another layer does not hide the change. Explicit local sizes, padding, and font metrics still override theme defaults.

## Finding and using tokens

See [all token declarations](/api/tokens), including legacy flat names. `ColorPrimitive*` names palette colors; `ColorSemantic*` names roles. `MetricPrimitive*` provides base scales; `MetricSemantic*` provides spacing, radii, font sizes, and similar values. `MetricComponent*` defines component dimensions; `Component*` selects appearance entries.

For example, ProgressBar uses `ColorSemanticProgressTrack/Fill`; Tabs uses `ColorSemanticTabsIndicator`; Menu distinguishes Hover/Active/Selected/Pressed/Disabled colors. Reference pages retain actual string values, searchable by name. Tokens are keys, not RGBA or float32 values: wrap them in TokenColor/TokenMetric.

Button Variant/Tone select common appearances. An explicit Token replaces that recipe while Size remains independent. Use `ComponentTheme` for shared changes and Style/States for one button. [Button example](/components/button).

Themes reject missing references, cycles, and invalid numbers. A completely zero Theme selects the built-in light theme; a partially populated Theme is not a patch over the default. Equal themes or changes to unused values may not repaint; SetTheme is not a forced refresh API.
