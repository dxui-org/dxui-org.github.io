# Select: dropdown selection

Build dropdowns with `Select`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/select/main.go

## Parameters and API

```go
func Select(props SelectProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Value` | `string` | Current application-owned value; see this page for behavior and zero-value semantics. |
| `Options` | `[]SelectOption` | Copied option list; nonempty unique values, at most 4096 items. |
| `Placeholder` | `string` | Displayed for an empty Value; defaults to empty. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `OnChange` | `func(string)` | Proposes the complete new value; write it to state for the next build. See behavior for nil callbacks. |

## Behavior and limitations

Option Value must be nonempty and unique for reorder identity. Empty lists and unmatched Value are supported. Nil OnChange still permits opening, browsing, and closing without accepting a selection. Enter/Space, arrows, Home/End, Escape, and Tab are supported; disabled options are skipped. The overlay escapes ancestor Scroll clipping and flips/clamps at window edges. Not a native system selector or virtualized list; at most 4096 items.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
