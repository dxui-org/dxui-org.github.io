# Icon: vector icons

Build vector icons with `Icon`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/icon/main.go

## Parameters and API

```go
func Icon(props IconProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Data` | `IconData` | Valid icon data from an official constructor or custom filled paths. |
| `Size` | `float32` | Logical size; 0 uses the theme default. |
| `Color` | `ColorValue` | Explicit color or TokenColor; inherits theme color when unset. |
| `StrokeWidth` | `IconStrokeWidth` | Finite nonnegative view-box stroke width; 0 means 2. Ignored for filled paths. |

## Behavior and limitations

Size uses logical units; 0 selects the theme default of 20. StrokeWidth uses view-box units, must be finite/nonnegative, and defaults to 2 when zero; custom filled icons ignore it. Icons have no action semantics; compose with Button. Individual icon.Name() calls link on demand, while icon/catalog retains the full catalog. IconData does not parse SVG.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
