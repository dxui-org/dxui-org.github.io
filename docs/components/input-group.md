# InputGroup: input prefixes and suffixes

Build input groups with `InputGroup`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/input-group/main.go

## Parameters and API

```go
func InputGroup(props InputGroupProps, content InputGroupContent) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |

## Behavior and limitations

InputGroupContent.Input must be exactly one Input. Supply Prefix/Suffix explicitly with Some; they may contain noninteractive content or Button, but no other focusable controls or nested editors. The group paints the shared background, border, corners, and focus-within appearance; Input takes the remaining width. Editor state and Key are retained. A suffix button is a separate Tab stop.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
