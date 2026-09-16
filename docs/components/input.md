# Input: single-line input

Build single-line inputs with `Input`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/input/main.go

## Parameters and API

```go
func Input(props InputProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Value` | `string` | Current application-owned value; see this page for behavior and zero-value semantics. |
| `OnChange` | `func(string)` | Proposes the complete new value; write it to state for the next build. See behavior for nil callbacks. |
| `Selection` | `Option[TextRange]` | Unset uses internal selection; Some(TextRange) controls rune-based selection. |
| `OnSelectionChange` | `func(TextRange)` | Proposes a complete rune selection. Nil retains internal selection; controlled Selection takes precedence on rebuild. |
| `Placeholder` | `string` | Displayed for an empty Value; defaults to empty. |
| `Password` | `bool` | Defaults to false; true masks text and disables copy/cut. |
| `ShowPasswordToggle` | `bool` | Defaults to false; effective only with Password=true. Visibility persists while identity is retained. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `ReadOnly` | `bool` | Defaults to false; true prevents editing/preedit while preserving navigation and selection. |
| `OnSubmit` | `func()` | Input Enter callback; still available when ReadOnly, unavailable when Disabled. |

## Behavior and limitations

Value is a fully controlled string; rejecting OnChange proposals rejects edits. Selection uses rune indices; controlled selection must accept OnSelectionChange to move. ReadOnly preserves navigation, selection, non-password copy, and OnSubmit; Disabled removes focus and stops input. Passwords cannot be copied/cut even when visible. ShowPasswordToggle only works with Password=true and changes internal visibility without OnChange. IME composition does not call OnChange until text is committed.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
