# ProgressBar: progress indicators

Build progress indicators with `ProgressBar`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/progress-bar/main.go

## Parameters and API

```go
func ProgressBar(props ProgressBarProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Value` | `float32` | Current application-owned value; see this page for its range and zero-value behavior. |

## Behavior and limitations

Value is a 0..1 ratio, not a percentage. Out-of-range values clamp; NaN/-Inf are empty and +Inf is full. No events or internal progress timer. Background controls the track; Text.Color controls the fill. Default size is 160×12 with an 8-unit track. No indeterminate, segmented, vertical, or animated modes.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
