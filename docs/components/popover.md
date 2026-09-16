# Popover: interactive overlays

Build interactive overlays with `Popover`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/popover/main.go

## Parameters and API

```go
func Popover(props PopoverProps, anchor, content View) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Open` | `bool` | Controlled open state, default false. |
| `Placement` | `OverlayPlacement` | Defaults to OverlayBottomStart; eight placement/alignment values are available. |
| `Offset` | `MetricValue` | Optional literal/token spacing for the overlay. |
| `OnOpenChange` | `func(bool)` | Proposes opening/closing; nil leaves Open unchanged. |

## Behavior and limitations

Accepts anchor and content Views. The app must accept OnOpenChange proposals to update Open. Pointer/Enter/Space on the trigger, topmost Escape, or an outside primary click proposes a change. Window-level overlays flip/clamp automatically. Tab can enter content; closing restores a surviving trigger's focus. No Disabled field; disabling an anchor child does not disable the Popover host. No modal behavior, focus trap, arrow, or animation.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
