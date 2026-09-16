# Image: raster images

Build raster images with `Image`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/image/main.go

## Parameters and API

```go
func Image(props ImageProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Source` | `ImageSource` | Stable image handle; reuse it rather than allocating during every build. |
| `Fit` | `ImageFit` | Defaults to ImageContain; also supports Cover, Fill, and None. |
| `Alignment` | `Point` | X/Y each range from 0 to 1; default (0,0) is top-left, (.5,.5) is centered. |
| `MaxPixels` | `int64` | Decoded pixel budget; 0 selects 16,777,216, and negative values are invalid. |
| `OnLoad` | `func(Size)` | Success notification with decoded size; nil still loads. |
| `OnError` | `func(error)` | Resource error notification; the app may render a fallback. Nil still attempts loading. |

## Behavior and limitations

Construct Source with ImageBytes, ImageFile, or ImageFromGo and reuse it outside builders. Contain preserves the entire image and aspect ratio; Cover fills and crops; Fill stretches; None keeps original size. Alignment is 0..1 per axis, default top-left. MaxPixels limits decoded pixels. Only PNG/JPEG/first-frame GIF; no network URLs or asynchronous decoding. Nil load callbacks do not prevent decoding. Do not modify a Go image after constructing its source.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
