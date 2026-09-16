# 布局、样式与主题类型：完整声明

[用途与示例](/zh-cn/api/values) · [全部 API 索引](/zh-cn/api/)

以下声明由 Go AST 提取，保留源码英文契约和字段注释，隐藏私有字段；`struct {}` 表示外部不可直接配置的内部表示，不表示可以用零值替代构造函数。常量块保留 iota 上下文。

## Length {#api-length}

```text
Length is an opaque automatic, logical-pixel, or percentage length. Its zero
value means automatic sizing; use Px or Percent for an explicit value.
```


```go
type Length struct {
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:13`

## Px {#api-px}

```text
Px creates a logical-pixel length.
```


```go
func Px(value float32) Length
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:19`

## Percent {#api-percent}

```text
Percent creates a percentage length in the range 0..100.
```


```go
func Percent(value float32) Length
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:22`

## Fill {#api-fill}

```text
Fill is the common full-available-axis length. It is equivalent to
Percent(100) and remains subject to the parent's definite-size rules.
```


```go
func Fill() Length
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:26`

## MetricToken {#api-metrictoken}

```text
MetricToken names a theme metric.
```


```go
type MetricToken string
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:29`

## ColorToken {#api-colortoken}

```text
ColorToken names a theme color.
```


```go
type ColorToken string
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:32`

## MetricValue {#api-metricvalue}

```text
MetricValue is either a literal logical-unit metric or a theme token.
```


```go
type MetricValue struct {
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:35`

## Metric {#api-metric}

```text
Metric creates a literal metric.
```


```go
func Metric(value float32) MetricValue
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:43`

## NoShrink {#api-noshrink}

```text
NoShrink explicitly disables flex shrinking. It is equivalent to Some(0)
and is distinct from an unset Shrink, whose default is one.
```


```go
func NoShrink() Option[float32]
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:47`

## Padding {#api-padding}

```text
Padding applies one literal logical-unit metric to every edge. The returned
metrics are explicitly set, including when value is zero.
```


```go
func Padding(value float32) EdgeValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:51`

## PaddingXY {#api-paddingxy}

```text
PaddingXY applies literal logical-unit metrics to the horizontal and vertical
edges. The returned metrics are explicitly set, including zero values.
```


```go
func PaddingXY(horizontal, vertical float32) EdgeValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:55`

## Margin {#api-margin}

```text
Margin applies one literal logical-unit metric to every edge, like Padding.
The returned metrics are explicitly set, including when value is zero.
```


```go
func Margin(value float32) EdgeValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:64`

## MarginXY {#api-marginxy}

```text
MarginXY applies literal logical-unit metrics to the horizontal and vertical
edges, like PaddingXY. The returned metrics are explicitly set, including zero values.
```


```go
func MarginXY(horizontal, vertical float32) EdgeValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:68`

## Round {#api-round}

```text
Round applies one literal logical-unit radius to every corner. The returned
metrics are explicitly set, including when value is zero.
```


```go
func Round(value float32) CornerValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:72`

## TokenMetric {#api-tokenmetric}

```text
TokenMetric creates a token-backed metric.
```


```go
func TokenMetric(token MetricToken) MetricValue
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:75`

## ColorValue {#api-colorvalue}

```text
ColorValue is either a literal color or a theme token.
```


```go
type ColorValue struct {
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:80`

## ColorRGBA {#api-colorrgba}

```text
ColorRGBA creates an explicitly set literal paint color from RGBA channels,
including transparent zero. Use RGBA for raw RGBAColor data instead.
```


```go
func ColorRGBA(r, g, b, a uint8) ColorValue
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:89`

## LiteralColor {#api-literalcolor}

```text
LiteralColor wraps existing raw RGBA data as a literal paint color.
Use ColorRGBA when supplying channels directly.
```


```go
func LiteralColor(value RGBAColor) ColorValue
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:93`

## TokenColor {#api-tokencolor}

```text
TokenColor creates a token-backed paint color.
```


```go
func TokenColor(token ColorToken) ColorValue
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:96`

## EdgeValues {#api-edgevalues}

```text
EdgeValues contains metrics in top, right, bottom, left order. Its zero-value
fields are unset.
```


```go
type EdgeValues struct{ Top, Right, Bottom, Left MetricValue }
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:102`

## Edges {#api-edges}

```text
Edges creates explicit logical-unit edges in top, right, bottom, left order.
```


```go
func Edges(top, right, bottom, left float32) EdgeValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:105`

## UniformEdges {#api-uniformedges}

```text
UniformEdges applies one metric to every edge.
```


```go
func UniformEdges(value MetricValue) EdgeValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:115`

## CornerValues {#api-cornervalues}

```text
CornerValues contains radii in top-left, top-right, bottom-right, bottom-left order.
Its zero-value fields are unset.
```


```go
type CornerValues struct{ TopLeft, TopRight, BottomRight, BottomLeft MetricValue }
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:121`

## Corners {#api-corners}

```text
Corners creates explicit logical-unit radii in top-left, top-right,
bottom-right, bottom-left order.
```


```go
func Corners(topLeft, topRight, bottomRight, bottomLeft float32) CornerValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:125`

## UniformCorners {#api-uniformcorners}

```text
UniformCorners applies one radius to every corner.
```


```go
func UniformCorners(value MetricValue) CornerValues
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:135`

## Position {#api-position}

```text
Position controls normal-flow versus absolute layout.
```


```go
type Position uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:140`

## PositionFlow {#api-positionflow}


```go
const (
	PositionFlow Position = iota
	PositionAbsolute
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:143`

## PositionAbsolute {#api-positionabsolute}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:144`

## Insets {#api-insets}

```text
Insets contains absolute-position insets.
```


```go
type Insets struct{ Top, Right, Bottom, Left Length }
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:148`

## Overflow {#api-overflow}

```text
Overflow controls container clipping.
```


```go
type Overflow uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:151`

## OverflowVisible {#api-overflowvisible}


```go
const (
	OverflowVisible Overflow = iota
	OverflowClip
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:154`

## OverflowClip {#api-overflowclip}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:155`

## Align {#api-align}

```text
Align controls cross-axis alignment.
```


```go
type Align uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:159`

## AlignStart {#api-alignstart}


```go
const (
	AlignStart Align = iota
	AlignCenter
	AlignEnd
	AlignStretch
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:162`

## AlignCenter {#api-aligncenter}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:163`

## AlignEnd {#api-alignend}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:164`

## AlignStretch {#api-alignstretch}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:165`

## Justify {#api-justify}

```text
Justify controls main-axis alignment.
```


```go
type Justify uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:169`

## JustifyStart {#api-justifystart}


```go
const (
	JustifyStart Justify = iota
	JustifyCenter
	JustifyEnd
	JustifySpaceBetween
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:172`

## JustifyCenter {#api-justifycenter}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:173`

## JustifyEnd {#api-justifyend}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:174`

## JustifySpaceBetween {#api-justifyspacebetween}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:175`

## Style {#api-style}

```text
Style contains the supported layout, paint, and typography controls. Box
is always single-line; there is intentionally no wrap or order
property. The zero value is safe and selects intrinsic sizing and theme
defaults.
```


```go
type Style struct {
	Width, Height       Length
	MinWidth, MinHeight Length
	MaxWidth, MaxHeight Length
	Margin, Padding     EdgeValues
	Position            Position
	Insets              Insets
	Grow                float32
	Shrink              Option[float32]
	Basis               Length
	AlignSelf           Option[Align]
	ZIndex              int
	Overflow            Overflow
	Background          ColorValue
	Border              Border
	Radius              CornerValues
	// Shadow is an explicitly requested list of outer shadows. A nil value
	// leaves a theme/state value unchanged; a non-nil empty slice removes it.
	Shadow     []Shadow
	Opacity    Option[float32]
	Visibility Visibility
	Text       TextStyle
	// Force is applied after component and interaction-state styles. It is for
	// deliberate paint overrides such as suppressing a focus ring; ordinary
	// base appearance belongs in the fields above so Hover/Pressed/Focus remain
	// visible. Force cannot affect layout.
	Force StylePatch
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:182`

## Visibility {#api-visibility}

```text
Visibility controls whether a node contributes display items.
```


```go
type Visibility uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:212`

## Visible {#api-visible}


```go
const (
	Visible Visibility = iota
	Hidden
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:215`

## Hidden {#api-hidden}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:216`

## PointerBehavior {#api-pointerbehavior}

```text
PointerBehavior controls pointer participation for a view subtree.
```


```go
type PointerBehavior uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:220`

## PointerAuto {#api-pointerauto}


```go
const (
	PointerAuto PointerBehavior = iota
	// PointerNone excludes the view and all descendants. It is intended for
	// decorative overlays that must not intercept content below them.
	PointerNone
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:223`

## PointerNone {#api-pointernone}

```text
PointerNone excludes the view and all descendants. It is intended for
decorative overlays that must not intercept content below them.
```

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:226`

## BorderPattern {#api-borderpattern}

```text
BorderPattern selects how a border is painted. The zero value is solid.
```


```go
type BorderPattern uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:230`

## BorderSolid {#api-bordersolid}


```go
const (
	BorderSolid BorderPattern = iota
	BorderDashed
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:233`

## BorderDashed {#api-borderdashed}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:234`

## BorderSides {#api-bordersides}

```text
BorderSides selects edges. Zero means all edges, not an absent border.
```


```go
type BorderSides uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:238`

## BorderTop {#api-bordertop}


```go
const (
	BorderTop BorderSides = 1 << iota
	BorderRight
	BorderBottom
	BorderLeft
	BorderAll = BorderTop | BorderRight | BorderBottom | BorderLeft
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:241`

## BorderRight {#api-borderright}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:242`

## BorderBottom {#api-borderbottom}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:243`

## BorderLeft {#api-borderleft}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:244`

## BorderAll {#api-borderall}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:245`

## Border {#api-border}

```text
Border is a paint-only border drawn inside a view's layout bounds.
Each side owns the nearest half of its two adjacent corner arcs.
```


```go
type Border struct {
	Width   MetricValue
	Color   ColorValue
	Pattern BorderPattern
	// Sides defaults to all edges. In Style, a Width/Color assignment resets
	// Sides too; otherwise nonzero Sides changes only edge selection.
	// StylePatch.Border replaces the complete Border.
	Sides BorderSides
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:250`

## NoBorder {#api-noborder}

```text
NoBorder explicitly clears the border width. In ordinary Style, later state
patches may restore a border; Style.Force is applied after those patches.
```


```go
func NoBorder() Border
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:262`

## Stroke {#api-stroke}

```text
Stroke creates a solid border with a literal logical-unit width.
```


```go
func Stroke(width float32, color ColorValue) Border
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:265`

## Shadow {#api-shadow}

```text
Shadow is an outer paint-only shadow. MVP shadows support finite,
non-negative blur/spread and are rendered by a bounded approximation.
```


```go
type Shadow struct {
	OffsetX, OffsetY MetricValue
	Blur, Spread     MetricValue
	Color            ColorValue
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:271`

## FontWeight {#api-fontweight}

```text
FontWeight selects the nearest registered face weight in a family.
```


```go
type FontWeight uint16
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:278`

## WeightRegular {#api-weightregular}


```go
const (
	WeightRegular FontWeight = 400
	WeightMedium  FontWeight = 500
	WeightBold    FontWeight = 700
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:281`

## WeightMedium {#api-weightmedium}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:282`

## WeightBold {#api-weightbold}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:283`

## FontSlant {#api-fontslant}

```text
FontSlant selects a registered normal or italic face.
```


```go
type FontSlant uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:287`

## SlantNormal {#api-slantnormal}


```go
const (
	SlantNormal FontSlant = iota
	SlantItalic
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:290`

## SlantItalic {#api-slantitalic}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:291`

## TextAlign {#api-textalign}

```text
TextAlign is horizontal alignment within a Text node's content box.
```


```go
type TextAlign uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:295`

## TextStart {#api-textstart}


```go
const (
	TextStart TextAlign = iota
	TextCenter
	TextEnd
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:298`

## TextCenter {#api-textcenter}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:299`

## TextEnd {#api-textend}

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:300`

## TextStyle {#api-textstyle}

```text
TextStyle contains the simple-LTR/CJK MVP text inputs. Families are tried in
order before the App default and built-in Latin fallback. Size is a font
size in logical units; zero uses the current theme's default. LineHeight is
a logical-unit line height; zero uses the current theme's default.
Arabic/Indic shaping, bidi/RTL, color emoji, and vertical text are not
supported by this API.
```


```go
type TextStyle struct {
	Families   []FontFamily
	Size       float32
	LineHeight float32
	Weight     FontWeight
	Slant      FontSlant
	Color      ColorValue
	Align      TextAlign
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:309`

## ComponentToken {#api-componenttoken}

```text
ComponentToken names a component-level theme entry.
```


```go
type ComponentToken string
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:320`

## PrimitiveTokens {#api-primitivetokens}

```text
PrimitiveTokens are the literal foundation of a theme.
```


```go
type PrimitiveTokens struct {
	Colors  map[ColorToken]RGBAColor
	Metrics map[MetricToken]float32
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:323`

## SemanticTokens {#api-semantictokens}

```text
SemanticTokens map product meaning onto primitive or earlier semantic
tokens. A literal ColorValue/MetricValue is also accepted.
```


```go
type SemanticTokens struct {
	Colors  map[ColorToken]ColorValue
	Metrics map[MetricToken]MetricValue
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:330`

## StylePatch {#api-stylepatch}

```text
StylePatch is an explicitly optional paint-only override. State styles are
intentionally paint-only in MVP, so interaction never moves layout.
```


```go
type StylePatch struct {
	Background Option[ColorValue]
	Border     Option[Border]
	Radius     Option[CornerValues]
	Shadow     Option[[]Shadow]
	Opacity    Option[float32]
	Visibility Option[Visibility]
	TextColor  Option[ColorValue]
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:337`

## StateStyles {#api-statestyles}

```text
StateStyles contains the deterministic MVP visual-state cascade.
```


```go
type StateStyles struct {
	Default  StylePatch
	Hover    StylePatch
	Focus    StylePatch
	Disabled StylePatch
	Pressed  StylePatch
	Checked  StylePatch
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:348`

## ComponentTheme {#api-componenttheme}

```text
ComponentTheme supplies semantic-token-backed defaults for one component.
```


```go
type ComponentTheme struct {
	Base   StylePatch
	States StateStyles
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:358`

## Theme {#api-theme}

```text
Theme is the public, type-safe Primitive -> Semantic -> Component token
structure. SetTheme validates and copies every map and slice atomically.
```


```go
type Theme struct {
	Primitive  PrimitiveTokens
	Semantic   SemanticTokens
	Components map[ComponentToken]ComponentTheme
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:365`

## Direction {#api-direction}

```text
Direction selects Box's main axis. Vertical is the zero value.
```


```go
type Direction uint8
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:384`

## Vertical {#api-vertical}

```text
Vertical lays out Box children from top to bottom and is the zero value.
```


```go
const (
	// Vertical lays out Box children from top to bottom and is the zero value.
	Vertical Direction = iota
	// Horizontal lays out Box children from left to right.
	Horizontal
)
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:388`

## Horizontal {#api-horizontal}

```text
Horizontal lays out Box children from left to right.
```

类型和值见本页同组常量块。

[对应示例](/zh-cn/api/values) · 源文件 `style.go:390`

## BoxProps {#api-boxprops}

```text
BoxProps configures a Box. Gap is a fixed logical-unit spacing; zero means
no spacing between children.
```


```go
type BoxProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior

	Direction Direction
	Gap       float32
	Justify   Justify
	Align     Align
}
```

[对应示例](/zh-cn/api/values) · 源文件 `style.go:395`

