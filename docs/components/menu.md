# Menu: action lists

Build action lists with `Menu`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/menu/main.go

## Parameters and API

```go
func Menu(props MenuProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Orientation` | `MenuOrientation` | Defaults to vertical; see behavior below. |
| `Value` | `string` | Current application-owned value; see this page for behavior and zero-value semantics. |
| `Items` | `[]MenuItem` | Copied item list with nonempty unique values; may be empty. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `OnAction` | `func(string)` | Receives Value whenever an enabled item executes; the same item may execute repeatedly. |

## Behavior and limitations

An action list within normal layout. Defaults to MenuVertical; MenuHorizontal selects horizontal. Value optionally controls selected appearance; OnAction still fires for an already-selected item. Item.Value must be nonempty and unique. One menu contributes one Tab stop. Orientation-specific arrows/Home/End move the internal current item without wrapping; Enter/Space executes. No built-in popup, submenus, separators, routing, or shortcut fields. Compose long menus with Scroll.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
