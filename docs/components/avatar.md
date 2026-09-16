# Avatar: avatars

Build avatars with `Avatar`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/avatar/main.go

## Parameters and API

```go
func Avatar(props AvatarProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Source` | `ImageSource` | Stable image handle; reuse it rather than allocating during every build. |
| `Shape` | `AvatarShape` | Defaults to AvatarCircle; AvatarSquare selects a square. |
| `Size` | `float32` | Logical size; 0 uses the theme default. |
| `OnLoad` | `func(Size)` | Success notification with decoded size; nil still loads. |
| `OnError` | `func(error)` | Resource error notification; the app may render a fallback. Nil still attempts loading. |

## Behavior and limitations

Shares ImageSource and OnLoad/OnError, with centered Cover fitting. Shape defaults to AvatarCircle; AvatarSquare is optional. Size=0 uses the theme default of 40; positive values are logical units, overridden by explicit Style dimensions. No URLs, initials, automatic fallback, or status dots. Render fallback content in response to OnError.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
