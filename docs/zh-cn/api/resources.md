# 字体、图片与图标资源：完整声明

[用途与示例](/zh-cn/api/resources-guide) · [全部 API 索引](/zh-cn/api/)

以下声明由 Go AST 提取，保留源码英文契约和字段注释，隐藏私有字段；`struct {}` 表示外部不可直接配置的内部表示，不表示可以用零值替代构造函数。常量块保留 iota 上下文。

## FontFamily {#api-fontfamily}

```text
FontFamily is an application-visible family name used for deterministic
ordered fallback. It is not an operating-system font handle.
```


```go
type FontFamily string
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:17`

## FontFamilyDefault {#api-fontfamilydefault}

```text
FontFamilyDefault is dxui's embedded, BSD-licensed Go Regular font. It
covers the MVP Latin samples but not CJK.
```


```go
const (
	// FontFamilyDefault is dxui's embedded, BSD-licensed Go Regular font. It
	// covers the MVP Latin samples but not CJK.
	FontFamilyDefault FontFamily = internaltext.BuiltinFamily
	// FontFamilySystemSans resolves one fixed candidate list per operating
	// system; it is never searched by fuzzy display name.
	FontFamilySystemSans FontFamily = "system-ui"
	// FontFamilySystemMono is the deterministic system monospace family.
	FontFamilySystemMono FontFamily = "system-monospace"
	// FontFamilySystemCJK is the deterministic system CJK fallback family.
	FontFamilySystemCJK FontFamily = "system-cjk"
)
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:22`

## FontFamilySystemSans {#api-fontfamilysystemsans}

```text
FontFamilySystemSans resolves one fixed candidate list per operating
system; it is never searched by fuzzy display name.
```

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:25`

## FontFamilySystemMono {#api-fontfamilysystemmono}

```text
FontFamilySystemMono is the deterministic system monospace family.
```

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:27`

## FontFamilySystemCJK {#api-fontfamilysystemcjk}

```text
FontFamilySystemCJK is the deterministic system CJK fallback family.
```

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:29`

## Font {#api-font}

```text
Font registers one TTF/OTF/TTC face. A FontBytes value owns its copied bytes
for as long as that value or an App configured with it remains reachable.
FontFile handles and all parsed faces live only from App.Run text startup until
renderer teardown completes. The file is not copied into the Go heap.
FaceIndex selects a collection face and is zero
for ordinary fonts.
```


```go
type Font struct {
	Family    FontFamily
	Weight    FontWeight
	Slant     FontSlant
	FaceIndex int
}
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:52`

## FontBytes {#api-fontbytes}

```text
FontBytes creates an application-owned font from a defensive copy of data.
```


```go
func FontBytes(family FontFamily, data []byte) Font
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:61`

## FontFile {#api-fontfile}

```text
FontFile creates a font loaded from the exact path during App.Run startup.
dxui does not search or copy the file into an application package.
```


```go
func FontFile(family FontFamily, path string) Font
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:67`

## SystemFont {#api-systemfont}

```text
SystemFont requests one of the documented generic system families. The
resolver tests a fixed path list for the current OS and returns a startup
error when none exists; it never depends on fontconfig or fuzzy name search.
```


```go
func SystemFont(family FontFamily) Font
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `font.go:74`

## Point {#api-point}

```text
Point is a position in icon view-box coordinates.
```


```go
type Point struct{ X, Y float32 }
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:12`

## Rect {#api-rect}

```text
Rect is a rectangle in icon view-box coordinates.
```


```go
type Rect struct{ X, Y, Width, Height float32 }
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:15`

## PathCommand {#api-pathcommand}

```text
PathCommand stores up to three points.
```


```go
type PathCommand struct {
	Verb   PathVerb
	Points [3]Point
}
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:29`

## PaintMode {#api-paintmode}

```text
PaintMode identifies whether a packed path is filled or stroked.
```


```go
type PaintMode uint8
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:35`

## PaintStroke {#api-paintstroke}


```go
const (
	PaintStroke PaintMode = iota
	PaintFill
)
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:38`

## PaintFill {#api-paintfill}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:39`

## Path {#api-path}

```text
Path is one decoded path with a single paint mode.
```


```go
type Path struct {
	Mode     PaintMode
	Commands []PathCommand
}
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:43`

## Data {#api-data}

```text
Data is immutable when returned by Packed. Commands remains exported for
compatibility with application-authored filled icons; callers own that
slice and dxui copies it when constructing a View.
```


```go
type Data struct {
	ViewBox  Rect
	Commands []PathCommand
}
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:51`

## IconData.IsPacked {#api-icondata-ispacked}

```text
IsPacked reports whether data uses the immutable generated representation.
```


```go
func (data Data) IsPacked() bool
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:70`

## IconData.Identity {#api-icondata-identity}

```text
Identity returns a stable content identity without exposing packed bytes.
```


```go
func (data Data) Identity() uint64
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:73`

## IconData.Paths {#api-icondata-paths}

```text
Paths validates and decodes data. Legacy Commands form one filled path.
```


```go
func (data Data) Paths(maxCommands int) ([]Path, error)
```

[对应示例](/zh-cn/api/resources-guide) · 源文件 `internal/icondata/icondata.go:99`

