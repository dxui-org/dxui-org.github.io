# Checkbox: checkboxes

Build checkboxes with `Checkbox`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/checkbox/main.go

## Parameters and API

```go
func Checkbox(props CheckboxProps, label View) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Checked` | `bool` | Controlled boolean, default false; accept changes in OnChange. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `OnChange` | `func(bool)` | Proposes a complete new value; write it to state for the next build. See behavior for nil callbacks. |

## Behavior and limitations

Accepts one label View. Checked is controlled; pointer or Space proposes toggling it. Nil OnChange does not change the business value and is not Disabled. Disabled removes focus and cancels pressing. No tri-state support.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
