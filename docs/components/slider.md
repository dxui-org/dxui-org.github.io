# Slider: numeric sliders

Build numeric sliders with `Slider`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/slider/main.go

## Parameters and API

```go
func Slider(props SliderProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Value` | `float32` | Current application-owned value; see this page for its range and zero-value behavior. |
| `Min` | `float32` | Lower bound; if both bounds are 0, uses 0..100. Otherwise literal, finite, and less than Max. |
| `Max` | `float32` | Upper bound; if both bounds are 0, defaults to 100. Otherwise literal, finite, and greater than Min. |
| `Step` | `float32` | Steps from Min; nonpositive or nonfinite values default to 1. Endpoints remain reachable. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `OnChange` | `func(float32)` | Proposes the complete new value; write it to state for the next build. See behavior for nil callbacks. |

## Behavior and limitations

Both bounds at zero mean 0..100; nonpositive/nonfinite Step falls back to 1. Other bounds are literal. Nonfinite bounds or Min>=Max disable interaction and focus. Value clamps; NaN/-Inf display Min and +Inf displays Max. Proposals align steps from Min while preserving endpoint access. Supports track clicks, captured dragging, arrows, Home, and End.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
