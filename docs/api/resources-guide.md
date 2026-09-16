# Fonts, images, and icon resources

See [resource declarations](/api/resources) for exact signatures; ImageSource/ImageFit and related types are in [component types](/api/components). Create and reuse resource descriptions outside the root builder.

## Fonts

| API | Parameters/results | Lifetime and constraints |
| --- | --- | --- |
| FontBytes(family, data) Font | FontFamily name and TTF/OTF/TTC bytes | Copies bytes immediately; data remains while the Font/App is reachable. |
| FontFile(family, path) Font | Exact file path | Opens at Run startup; files are not automatically bundled. Retained until text resources are destroyed; no hot reload. |
| SystemFont(family) Font | Fixed generic system family | Tries the current OS candidate paths; missing fonts cause startup errors. Does not search arbitrary display names. |
| Font | Family, Weight, Slant, FaceIndex | Use constructors for the hidden source; FaceIndex selects a TTC face, or 0 for a normal font. Register in AppOptions.Fonts. |
| FontFamilyDefault | Embedded Go Regular | Basic Latin coverage; no CJK. |
| FontFamilySystemSans / Mono / CJK | system-ui / system-monospace / system-cjk | Fixed platform candidates; availability varies. |

```go
// Configuration fragment: distribute the font with the app; paths use the working directory.
font := dxui.FontFile("app-cjk", "assets/NotoSansCJK-Regular.ttc")
font.FaceIndex = 2 // Use only if this collection has the intended face at index 2.
app := dxui.NewApp(dxui.AppOptions{

	Fonts: []dxui.Font{font},

	DefaultFont: "app-cjk",

	DisableSystemFontFallback: true,
})
_ = app
```

Fallback considers TextStyle.Families, registered/default application fonts, embedded Latin fonts, lazily loaded system CJK candidates, then replacement/missing-glyph boxes. Automatic CJK fallback is enabled by default; missing candidates do not fail the entire layout, but explicit SystemFont failures are startup errors. For reproducible distribution, bundle font files or bytes you may redistribute and retain their licenses; do not rely on fonts installed on a development machine.

## Complete font example

This FontBytes program needs no external font file: it uses Go Regular bytes from DXUI's current dependencies. It demonstrates byte registration and default family selection for Latin text; this is not a CJK font.

<<< ../examples/fonts/main.go

The example needs `golang.org/x/image` (`go mod tidy` resolves it through DXUI's dependency graph). The Go font project distributes this font under a BSD license.

## Images

`ImageBytes([]byte)` copies encoded data immediately; `ImageFile(string)` records a path read on demand; `ImageFromGo(image.Image)` retains a Go image that must not be mutated. ImageSource is comparable and exposes no pixel reader. New constructors may create new resource identities; cache handles outside builders and replace them when needed.

Decodes PNG, JPEG, and only the first GIF frame. No online URLs, animation, or asynchronous decoding. Image.MaxPixels defaults to 16×1024×1024 and limits decoded size; Avatar uses the same pipeline but exposes no MaxPixels. OnLoad(Size)/OnError(error) are notifications; nil does not stop loading. Run the self-contained [Image](/components/image) and [Avatar](/components/avatar) examples.

## Official and custom icons

Use parameterless constructors such as `icon.Search()`; the [catalog](/api/icons) lists all 2,066 names and aliases. Constructors return IconData; IconProps.Size/Color/StrokeWidth configure rendering. Ordinary applications do not need catalog: catalog.Icons retains the entire library.

Custom IconData describes filled paths with ViewBox and []PathCommand. PathMove/PathLine use Points[0], PathQuad uses two points, PathCubic uses three, and PathClose closes a path. This is not SVG syntax and there is no runtime SVG parser.

<<< ../examples/custom-icon/main.go

IconData is a public alias, so `IsPacked() bool`, `Identity() uint64`, and `Paths(maxCommands int)` are callable. IsPacked distinguishes official immutable encodings from custom Commands. Identity is a content hash, not a collision-free business key. Paths validates and decodes; maxCommands must be positive and sufficient, otherwise it returns an error.

Paths returns internally defined elements with Mode and Commands, accessible through type inference without importing internal packages. Mode uses 0 for stroke and 1 for fill. These are low-level alias details; component code should use IconData directly. Do not mutate official descriptions or decoded results to change resources. Custom Commands are copied when building a View.
