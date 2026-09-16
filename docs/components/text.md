# Text: text and Label

Build text with `Text` and `Label`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/text/main.go

## Parameters and API

```go
func Text(props TextProps) View
```

```go
func Label(value string) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Value` | `string` | Current application-owned value; see this page for behavior and zero-value semantics. |
| `Wrap` | `TextWrap` | Defaults to TextNoWrap; TextWrapWords uses simple word wrapping. |
| `MaxLines` | `int` | Positive values limit visible lines; 0 is unlimited, negative values fail the build. |

## Behavior and limitations

Label(value) is default Text; TextNoWrap is the zero value. TextWrapWords uses simple word splitting and collapses whitespace, not full Unicode line breaking. Positive MaxLines limits lines. Invalid UTF-8 becomes U+FFFD. Text is not a selectable editor; CJK display depends on font coverage.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
