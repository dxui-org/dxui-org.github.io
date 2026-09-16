# Tabs: tab selection

Build tab selectors with `Tabs`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/tabs/main.go

## Parameters and API

```go
func Tabs(props TabsProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Value` | `string` | Current application-owned value; see this page for behavior and zero-value semantics. |
| `Items` | `[]TabItem` | Copied item list with nonempty unique values; may be empty. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `OnChange` | `func(string)` | Proposes the complete new value; write it to state for the next build. See behavior for nil callbacks. |

## Behavior and limitations

Renders horizontal tabs only; the app builds content based on Value. Items.Value must be nonempty and unique. Empty/unmatched Value shows no selected indicator. Left/right/Home/End move the internal current item; Enter/Space proposes selection. Selected or disabled items do not call back. No TabPanel, vertical tabs, scrolling, closing, or drag-reordering API.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
