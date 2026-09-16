# Radio: radio buttons

Build radio buttons with `Radio`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/radio/main.go

## Parameters and API

```go
func Radio(props RadioProps, label View) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Selected` | `bool` | Selected state, default false; derive it from a shared application value. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `OnSelect` | `func()` | Called without arguments when an enabled unselected Radio activates; update the shared value in the app. |

## Behavior and limitations

Selected is authoritative. Pointer or Space calls OnSelect only when enabled and unselected; a selected item neither deselects nor calls back. Use shared application state for exclusivity; there is no RadioGroup. Every Radio keeps its own focus stop.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
