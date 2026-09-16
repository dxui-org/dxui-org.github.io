# Textarea: multiline input

Build multiline inputs with `Textarea`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/textarea/main.go

## Parameters and API

```go
func Textarea(props TextareaProps) View
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
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `ReadOnly` | `bool` | Defaults to false; true prevents editing/preedit while preserving navigation and selection. |
| `Wrap` | `TextWrap` | Defaults to TextNoWrap; TextWrapWords uses simple word wrapping. |

## Behavior and limitations

Uses Input's controlled-value and rune-selection rules. Enter inserts a newline when editable; no OnSubmit, Password, or password toggle. Wrap defaults to TextNoWrap, with TextWrapWords available. Retains internal caret scrolling and undo history. ReadOnly still allows selection, copy, and scrolling. Not a rich-text editor; no grapheme editing or RTL layout.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
