# Button: buttons and TextButton

Build buttons with `Button`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/button/main.go

## Parameters and API

```go
func Button(props ButtonProps, child View) View
```

```go
func TextButton(props ButtonProps, label string) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Variant` | `ButtonVariant` | Defaults to ButtonFilled; see appearance recipes below. |
| `Tone` | `ButtonTone` | Defaults to ButtonPrimary; ButtonDefault is an alias. |
| `Size` | `ButtonSize` | Size enum; defaults to ButtonNormal. |
| `Disabled` | `bool` | Defaults to false; true removes focus and cancels interaction on interactive components. |
| `OnPress` | `func()` | Runs on the UI thread after activation; nil performs no action. |

## Behavior and limitations

Button accepts one arbitrary child View; TextButton centers text. Defaults: Filled / Primary / Normal. Variants: Filled, Soft, Outline, Dashed, Ghost, Link. Tones: Primary, Secondary, Success, Info, Warn, Danger. Sizes: Small, Normal, Large. Explicit Token replaces the appearance recipe; local Style still overrides it. A matching pointer or Enter/Space release activates once. Nil OnPress retains focus and interaction appearance. Link does not navigate.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
