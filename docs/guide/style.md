# Styles and themes

Customize control appearance and switch between light and dark themes with a button.

## 1. Run the example

<<< ../examples/theme/main.go

Click **Switch theme** to change themes. Hover over the button to make it slightly translucent; press Tab to see keyboard focus.

## 2. Choose the right setting

| What to change | Where |
| --- | --- |
| Control size, spacing, and background | The control's Style. |
| Hover appearance | States.Hover; the example changes opacity. |
| Application colors | Call app.SetTheme. |

Start with Style for normal appearance. For example:

```go
dxui.Style{
    Width: dxui.Px(240),
    Padding: dxui.Padding(16),
    Background: dxui.ColorRGBA(240, 245, 255, 255),
}
```

## 3. Switch themes

The example starts from LightTheme() or DarkTheme(), changes the accent color, then calls app.SetTheme(next). Named theme values are tokens that give controls consistent colors.

Try changing Blue600 to Red600. Call SetTheme after modifying the theme; editing a Go map alone does not update the UI. Theme changes propagate to all live windows.

::: details Advanced: precedence, explicit clearing, and tokens

### Style precedence

```text
Primitive → Semantic → Component Base/Default
→ local Style → local States.Default
→ Hover → Focus → Checked → Pressed → Disabled
→ Style.Force
```

Within an active state, component theme patches apply before local patches. Disabled suppresses Hover/Focus/Pressed; a controlled Checked layer remains below it. State patches only change painting, not dimensions on hover.

### Explicit zero values

A Go zero value often means unset, not explicitly cleared.

| Goal | Setting |
| --- | --- |
| Clear default padding | `Padding: dxui.Padding(0)` |
| Explicit zero opacity | `Opacity: dxui.Some(float32(0))` |
| Prevent shrinking | `Shrink: dxui.NoShrink()` |
| Clear a normal border | `Border: dxui.NoBorder()` |
| Clear theme shadows | `Shadow: []dxui.Shadow{}`; nil does not clear |
| Also clear state shadows | `Force: dxui.StylePatch{Shadow: dxui.Some([]dxui.Shadow{})}` |

Removing focus rings makes keyboard position harder to see. Usually retain default Focus styling and change only normal appearance.

### Literals, tokens, and colors

RGBA returns raw RGBAColor for AppOptions.Background and primitive color maps. ColorRGBA returns explicit ColorValue for Style.Background. Wrap existing RGBAColor with LiteralColor; TokenColor references the theme. Metric supplies an explicit logical number; TokenMetric reads a theme metric.

Color.Primitive.Blue600 is a palette token; Color.Semantic.Accent is a semantic role. Themes may map semantic colors to other palette colors. Legacy flat constants are in the [token reference](/api/tokens); old White/Black/Blue/Gray are not identical to the new palette entries.

Start from a complete LightTheme() or DarkTheme() before editing maps. A Theme containing only Accent does not auto-fill missing entries. SetTheme rejects missing references, cycles, and invalid values. Accepted maps/slices are copied; editing the original map later does not update the UI without another SetTheme call. Equal themes or unused-token changes need not produce a frame.


:::

Next: [Build a form](./forms). See [theme APIs](/api/theme) for full configuration.
