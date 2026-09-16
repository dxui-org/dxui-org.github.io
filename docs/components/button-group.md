# ButtonGroup: connected buttons

Build connected buttons with `ButtonGroup`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/button-group/main.go

## Parameters and API

```go
func ButtonGroup(props ButtonGroupProps, buttons ...View) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Orientation` | `ButtonGroupOrientation` | Defaults to horizontal; see behavior below. |
| `Dividers` | `bool` | Defaults to false; true draws themed separators. |

## Behavior and limitations

Accepts only Button (including TextButton), with zero children allowed. Horizontal without dividers by default; ButtonGroupVertical selects vertical. Dividers overlaps theme borders so each shared edge is painted once. No group focus; buttons retain Disabled, callbacks, and source-order Tab behavior. No arrow navigation. Vertical auto-width aligns to the widest button; explicit child Width wins.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
