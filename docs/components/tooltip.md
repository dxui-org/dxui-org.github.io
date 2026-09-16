# Tooltip: hints

Build hints with `Tooltip`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/tooltip/main.go

## Parameters and API

```go
func Tooltip(props TooltipProps, anchor, content View) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Placement` | `OverlayPlacement` | Defaults to OverlayBottomStart; eight placement/alignment values are available. |
| `Offset` | `MetricValue` | Optional literal/token spacing for the overlay. |
| `Delay` | `time.Duration` | 0 selects the default 500ms delay; negative values are invalid. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |

## Behavior and limitations

Accepts anchor and content. Hover or keyboard focus starts Delay; leaving and losing focus closes it. Delay=0 means 500ms, not immediate display. Disabled suppresses the hint. Content is noninteractive, takes no focus, and does not block pointers; use Popover for interactive content. Driven by events/deadlines, with no application timer loop.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
