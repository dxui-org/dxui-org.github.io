# Box: containers and layout

Build containers and layouts with `Box`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/box/main.go

## Parameters and API

```go
func Box(props BoxProps, children ...View) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Direction` | `Direction` | Defaults to Vertical; Horizontal arranges left to right. |
| `Gap` | `float32` | Nonnegative fixed logical spacing between children, default 0. |
| `Justify` | `Justify` | Main-axis Start/Center/End/SpaceBetween, default Start. |
| `Align` | `Align` | Cross-axis Start/Center/End/Stretch, default Start. |

## Behavior and limitations

Defaults to Vertical and a single row/column; Gap uses logical units. Justify controls the main axis, Align the cross axis. Children may be empty but cannot contain zero Views. No wrapping, order, reverse direction, or full CSS Flexbox.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
