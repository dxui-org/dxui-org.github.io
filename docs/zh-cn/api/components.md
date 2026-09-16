# 组件关联类型：完整声明

[用途与示例](/zh-cn/components/) · [全部 API 索引](/zh-cn/api/)

以下声明由 Go AST 提取，保留源码英文契约和字段注释，隐藏私有字段；`struct {}` 表示外部不可直接配置的内部表示，不表示可以用零值替代构造函数。常量块保留 iota 上下文。

## ButtonVariant {#api-buttonvariant}

```text
ButtonVariant selects a Button's surface and border recipe.
```


```go
type ButtonVariant uint8
```

[对应示例](/zh-cn/components/) · 源文件 `button.go:10`

## ButtonFilled {#api-buttonfilled}


```go
const (
	ButtonFilled ButtonVariant = iota // Filled is the compatible default.
	ButtonSoft
	ButtonOutline
	ButtonDashed
	ButtonGhost
	ButtonLink // Link is an action, with compact spacing; it never navigates.
)
```

[对应示例](/zh-cn/components/) · 源文件 `button.go:13`

## ButtonSoft {#api-buttonsoft}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:14`

## ButtonOutline {#api-buttonoutline}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:15`

## ButtonDashed {#api-buttondashed}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:16`

## ButtonGhost {#api-buttonghost}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:17`

## ButtonLink {#api-buttonlink}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:18`

## ButtonTone {#api-buttontone}

```text
ButtonTone selects semantic colors independently of the surface recipe.
```


```go
type ButtonTone uint8
```

[对应示例](/zh-cn/components/) · 源文件 `button.go:22`

## ButtonPrimary {#api-buttonprimary}


```go
const (
	ButtonPrimary ButtonTone = iota
	ButtonSecondary
	ButtonSuccess
	ButtonInfo
	ButtonWarn
	ButtonDanger
)
```

[对应示例](/zh-cn/components/) · 源文件 `button.go:25`

## ButtonSecondary {#api-buttonsecondary}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:26`

## ButtonSuccess {#api-buttonsuccess}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:27`

## ButtonInfo {#api-buttoninfo}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:28`

## ButtonWarn {#api-buttonwarn}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:29`

## ButtonDanger {#api-buttondanger}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:30`

## ButtonDefault {#api-buttondefault}

```text
ButtonDefault is an alias for the zero-value Primary tone.
```


```go
// ButtonDefault is an alias for the zero-value Primary tone.
const ButtonDefault = ButtonPrimary
```

[对应示例](/zh-cn/components/) · 源文件 `button.go:34`

## ButtonSize {#api-buttonsize}

```text
ButtonSize selects overridable padding, minimum height, and inherited text metrics.
```


```go
type ButtonSize uint8
```

[对应示例](/zh-cn/components/) · 源文件 `button.go:37`

## ButtonNormal {#api-buttonnormal}


```go
const (
	ButtonNormal ButtonSize = iota
	ButtonSmall
	ButtonLarge
)
```

[对应示例](/zh-cn/components/) · 源文件 `button.go:40`

## ButtonSmall {#api-buttonsmall}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:41`

## ButtonLarge {#api-buttonlarge}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `button.go:42`

## ScrollAxis {#api-scrollaxis}

```text
ScrollAxis selects the axes whose content is measured without a maximum
and whose retained offset may change.
```


```go
type ScrollAxis uint8
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:12`

## ScrollVertical {#api-scrollvertical}


```go
const (
	ScrollVertical ScrollAxis = iota
	ScrollHorizontal
	ScrollBoth
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:15`

## ScrollHorizontal {#api-scrollhorizontal}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:16`

## ScrollBoth {#api-scrollboth}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:17`

## ScrollbarPolicy {#api-scrollbarpolicy}

```text
ScrollbarPolicy controls the overlay scrollbar. Hidden suppresses scrollbar
paint and pointer interaction without disabling wheel/trackpad scrolling.
```


```go
type ScrollbarPolicy uint8
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:22`

## ScrollbarAuto {#api-scrollbarauto}


```go
const (
	ScrollbarAuto ScrollbarPolicy = iota
	ScrollbarAlways
	ScrollbarHidden
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:25`

## ScrollbarAlways {#api-scrollbaralways}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:26`

## ScrollbarHidden {#api-scrollbarhidden}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:27`

## ScrollProps {#api-scrollprops}

```text
ScrollProps configures a one-child clipped viewport. Offset is authoritative
when set. Otherwise InitialOffset is used only when the view is first
mounted; later positions are preserved while the view keeps its identity.
OnScroll reports the complete offset. Without Offset, movement updates
retained state even with a nil callback; with Offset it only proposes a
change. Scroll has no Disabled or ReadOnly property.
```


```go
type ScrollProps struct {
	Key           string
	Style         Style
	Token         ComponentToken
	States        StateStyles
	Pointer       PointerBehavior
	Axis          ScrollAxis
	InitialOffset Option[Point]
	Offset        Option[Point]
	Scrollbar     ScrollbarPolicy
	// OnScroll reports movement. Nil still scrolls internally when Offset is unset; a set Offset stays authoritative.
	OnScroll func(Point)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:36`

## VirtualListProps {#api-virtuallistprops}

```text
VirtualListProps configures a vertical, fixed-row-height virtual list.
Count and Version identify one immutable data snapshot; increment Version
whenever row keys or content may have changed. ItemKey and Build run on the
UI thread and must be deterministic and side-effect free for that snapshot.
```


```go
type VirtualListProps struct {
	Key           string
	Style         Style
	Token         ComponentToken
	States        StateStyles
	Pointer       PointerBehavior
	Count         int
	Version       uint64
	RowHeight     float32
	Overscan      int
	InitialOffset Option[Point]
	Offset        Option[Point]
	Scrollbar     ScrollbarPolicy
	ItemKey       func(index int) string
	Build         func(index int) View
	OnScroll      func(Point)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:54`

## TextWrap {#api-textwrap}

```text
TextWrap selects the MVP simple wrapping policy.
```


```go
type TextWrap uint8
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:73`

## TextNoWrap {#api-textnowrap}


```go
const (
	TextNoWrap TextWrap = iota
	TextWrapWords
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:76`

## TextWrapWords {#api-textwrapwords}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:77`

## TextProps {#api-textprops}

```text
TextProps configures simple left-to-right text. Invalid UTF-8 is normalized
to U+FFFD when Text is constructed. TextWrapWords collapses whitespace and
wraps only at word boundaries; it is not Unicode line-break conformance.
```


```go
type TextProps struct {
	Key      string
	Style    Style
	Token    ComponentToken
	States   StateStyles
	Pointer  PointerBehavior
	Value    string
	Wrap     TextWrap
	MaxLines int
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:83`

## ButtonProps {#api-buttonprops}

```text
ButtonProps configures a semantic button.
```


```go
type ButtonProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	Variant ButtonVariant
	Tone    ButtonTone
	Size    ButtonSize
	// Disabled removes focus and activation and suppresses Hover, Pressed, and Focus styles.
	Disabled bool
	// OnPress runs on activation. Nil preserves focus and interaction visuals but emits no action.
	OnPress func()
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:95`

## ButtonGroupOrientation {#api-buttongrouporientation}

```text
ButtonGroupOrientation selects the main axis used to arrange buttons.
```


```go
type ButtonGroupOrientation uint8
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:111`

## ButtonGroupHorizontal {#api-buttongrouphorizontal}


```go
const (
	ButtonGroupHorizontal ButtonGroupOrientation = iota
	ButtonGroupVertical
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:114`

## ButtonGroupVertical {#api-buttongroupvertical}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:115`

## ButtonGroupProps {#api-buttongroupprops}

```text
ButtonGroupProps configures a non-focusable connected container of Button
views.
```


```go
type ButtonGroupProps struct {
	Key         string
	Style       Style
	Token       ComponentToken
	States      StateStyles
	Pointer     PointerBehavior
	Orientation ButtonGroupOrientation
	// Dividers enables one themed border between adjacent Buttons.
	Dividers bool
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:120`

## InputGroupProps {#api-inputgroupprops}

```text
InputGroupProps configures the non-focusable visual container around one
Input and its optional leading and trailing content.
```


```go
type InputGroupProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:133`

## InputGroupContent {#api-inputgroupcontent}

```text
InputGroupContent identifies the required Input and optional adornments.
Prefix and Suffix may contain passive views or Buttons; other focusable
controls and nested text editors are rejected.
```


```go
type InputGroupContent struct {
	Input  View
	Prefix Option[View]
	Suffix Option[View]
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:144`

## BadgeProps {#api-badgeprops}

```text
BadgeProps configures a compact, non-interactive label around one arbitrary
child. The child inherits the Badge text/icon tint unless locally overridden.
```


```go
type BadgeProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:152`

## ProgressBarProps {#api-progressbarprops}

```text
ProgressBarProps configures a deterministic, non-interactive progress
indicator. Value is clamped to [0,1] for painting. NaN and negative
infinity paint empty; positive infinity paints complete.
```


```go
type ProgressBarProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	Value   float32
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:163`

## PathVerb {#api-pathverb}

```text
PathVerb identifies one command in dxui's deliberately small vector icon
format. It is not an SVG parser or a general scene graph.
```


```go
type PathVerb = icondata.PathVerb
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:174`

## PathMove {#api-pathmove}


```go
const (
	PathMove  = icondata.PathMove
	PathLine  = icondata.PathLine
	PathQuad  = icondata.PathQuad
	PathCubic = icondata.PathCubic
	PathClose = icondata.PathClose
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:177`

## PathLine {#api-pathline}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:178`

## PathQuad {#api-pathquad}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:179`

## PathCubic {#api-pathcubic}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:180`

## PathClose {#api-pathclose}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:181`

## PathCommand {#api-pathcommand}

```text
PathCommand stores up to three points. Move/Line use Points[0], Quad uses
Points[0:2], and Cubic uses all three points.
```


```go
type PathCommand = icondata.PathCommand
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:186`

## Rect {#api-rect}

```text
Rect is a rectangle in logical units.
```


```go
type Rect = icondata.Rect
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:189`

## IconData {#api-icondata}

```text
IconData is immutable dxui path data in ViewBox coordinates.
```


```go
type IconData = icondata.Data
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:192`

## IconStrokeWidth {#api-iconstrokewidth}

```text
IconStrokeWidth is a stroke width in icon view-box units. Zero selects the
Lucide default of 2. It affects immutable stroked resources and is ignored
by legacy filled IconData.
```


```go
type IconStrokeWidth float32
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:197`

## IconProps {#api-iconprops}

```text
IconProps configures a vector icon. Size is measured in logical units; zero
uses the current component theme default.
```


```go
type IconProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	Data    IconData
	Size    float32
	Color   ColorValue
	// StrokeWidth uses icon view-box units and defaults to 2. Values must be
	// finite and non-negative; zero means the default rather than no stroke.
	StrokeWidth IconStrokeWidth
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:201`

## ToggleSwitchProps {#api-toggleswitchprops}

```text
ToggleSwitchProps configures a controlled switch.
```


```go
type ToggleSwitchProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	// Checked is authoritative. Disabled cancels interaction and removes the focus stop.
	Checked, Disabled bool
	// OnChange proposes Checked. Nil preserves focus/press visuals without changing Checked.
	OnChange func(bool)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:216`

## SliderProps {#api-sliderprops}

```text
SliderProps configures a controlled single-value horizontal slider. Value
is authoritative; OnChange receives a clamped, step-aligned proposal. Zero
Min and Max select the default 0..100 range. Step defaults to 1 when it is
non-positive or non-finite. Other invalid ranges are inert.
```


```go
type SliderProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	Value   float32
	Min     float32
	Max     float32
	Step    float32
	// Disabled cancels dragging and removes focus and pointer interaction.
	Disabled bool
	// OnChange proposes Value. Nil preserves focus and drag visuals without changing Value.
	OnChange func(float32)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:232`

## CheckboxProps {#api-checkboxprops}

```text
CheckboxProps configures a controlled two-state checkbox.
```


```go
type CheckboxProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	// Checked is authoritative. Disabled cancels interaction and removes the focus stop.
	Checked, Disabled bool
	// OnChange proposes Checked. Nil preserves focus/press visuals without changing Checked.
	OnChange func(bool)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:249`

## RadioProps {#api-radioprops}

```text
RadioProps configures a controlled radio button. Selected is authoritative;
OnSelect is called only when an enabled, unselected Radio is activated.
```


```go
type RadioProps struct {
	Key      string
	Style    Style
	Token    ComponentToken
	States   StateStyles
	Pointer  PointerBehavior
	Selected bool
	// Disabled cancels interaction and removes the focus stop.
	Disabled bool
	// OnSelect proposes selection of an unselected Radio. Nil preserves focus and press visuals.
	OnSelect func()
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:263`

## ImageSource {#api-imagesource}

```text
ImageSource is an immutable, comparable handle to application image data.
```


```go
type ImageSource struct{}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:284`

## ImageBytes {#api-imagebytes}

```text
ImageBytes copies encoded PNG, JPEG, or GIF bytes immediately.
```


```go
func ImageBytes(data []byte) ImageSource
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:287`

## ImageFile {#api-imagefile}

```text
ImageFile records a path that is read by the image engine on demand.
```


```go
func ImageFile(path string) ImageSource
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:292`

## ImageFromGo {#api-imagefromgo}

```text
ImageFromGo records an immutable-by-contract Go image source.
```


```go
func ImageFromGo(value image.Image) ImageSource
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:295`

## ImageFit {#api-imagefit}

```text
ImageFit selects the supported destination fitting policy.
```


```go
type ImageFit uint8
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:298`

## ImageContain {#api-imagecontain}


```go
const (
	ImageContain ImageFit = iota
	ImageCover
	ImageFill
	ImageNone
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:301`

## ImageCover {#api-imagecover}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:302`

## ImageFill {#api-imagefill}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:303`

## ImageNone {#api-imagenone}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:304`

## ImageProps {#api-imageprops}

```text
ImageProps configures a decoded raster image. Alignment components are in
[0,1].
```


```go
type ImageProps struct {
	Key       string
	Style     Style
	Token     ComponentToken
	States    StateStyles
	Pointer   PointerBehavior
	Source    ImageSource
	Fit       ImageFit
	Alignment Point
	MaxPixels int64
	// OnLoad and OnError are optional notifications; nil does not stop decoding.
	OnLoad  func(Size)
	OnError func(error)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:309`

## AvatarShape {#api-avatarshape}

```text
AvatarShape selects the fixed Avatar clipping shape.
```


```go
type AvatarShape uint8
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:325`

## AvatarCircle {#api-avatarcircle}


```go
const (
	AvatarCircle AvatarShape = iota
	AvatarSquare
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:328`

## AvatarSquare {#api-avatarsquare}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:329`

## AvatarProps {#api-avatarprops}

```text
AvatarProps configures a square, centered cover image. Size is measured in
logical units; zero uses the current component theme default. Explicit Style
width and height take precedence.
```


```go
type AvatarProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	Source  ImageSource
	Shape   AvatarShape
	Size    float32
	// OnLoad and OnError are optional notifications; nil does not stop decoding.
	OnLoad  func(Size)
	OnError func(error)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:335`

## TextRange {#api-textrange}

```text
TextRange uses Unicode code-point (rune) offsets, never UTF-8 byte offsets.
```


```go
type TextRange struct{ Start, End int }
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:350`

## InputProps {#api-inputprops}

```text
InputProps configures a controlled single-line input. OnChange receives the
complete proposed value after committed text, paste, cut, deletion, undo, or
redo. The application accepts the edit by returning that value from the next
build; leaving Value unchanged rejects it. IME composition does not call
OnChange. Selection and TextRange use rune indices. A nil OnChange makes the
input read-only by behavior, without Disabled styling. ReadOnly also blocks
pre-edit; focus, selection, non-password copy, and OnSubmit remain available.
```


```go
type InputProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	Value   string
	// OnChange proposes Value. Nil blocks editing and pre-edit, but preserves selection and copy.
	OnChange  func(string)
	Selection Option[TextRange]
	// OnSelectionChange reports rune selection. Nil retains internal selection; Selection, when set, wins on rebuild.
	OnSelectionChange func(TextRange)
	Placeholder       string
	Password          bool
	// ShowPasswordToggle adds an internal trailing visibility button only when
	// Password is also true. Visibility is temporary state owned by the Input.
	ShowPasswordToggle bool
	// Disabled removes focus, stops native text input, and cancels composition, drag, and press.
	Disabled bool
	// ReadOnly blocks edits and pre-edit while allowing focus, navigation, selection, and non-password copy.
	ReadOnly bool
	// OnSubmit handles Enter independently of OnChange and ReadOnly. Nil emits no submit; Disabled blocks it.
	OnSubmit func()
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:359`

## TextareaProps {#api-textareaprops}

```text
TextareaProps configures a controlled multiline editor and follows the same
controlled Value and rune-indexed selection contract as Input. TextWrapWords
is a simple LTR/CJK word-wrap policy, not full Unicode line breaking. A nil
OnChange preserves navigation, selection, copy, and scrolling without edits
or pre-edit; it does not imply Disabled styling.
```


```go
type TextareaProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	Value   string
	// OnChange proposes Value. Nil blocks editing and pre-edit, but preserves selection and copy.
	OnChange  func(string)
	Selection Option[TextRange]
	// OnSelectionChange reports rune selection. Nil retains internal selection; Selection, when set, wins on rebuild.
	OnSelectionChange func(TextRange)
	Placeholder       string
	// Disabled removes focus, stops native text input, and cancels composition, drag, and press.
	Disabled bool
	// ReadOnly blocks edits and pre-edit while allowing focus, navigation, selection, and non-password copy.
	ReadOnly bool
	Wrap     TextWrap
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:389`

## SelectOption {#api-selectoption}

```text
SelectOption is one immutable entry in a Select. Value must be non-empty and
unique within the Select; it is both the controlled application value and
the stable identity used when options reorder.
```


```go
type SelectOption struct {
	Value    string
	Label    string
	Disabled bool
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:412`

## SelectProps {#api-selectprops}

```text
SelectProps configures a controlled custom popup Select. Value is
authoritative; an empty or unmatched Value displays Placeholder.
```


```go
type SelectProps struct {
	Key         string
	Style       Style
	Token       ComponentToken
	States      StateStyles
	Pointer     PointerBehavior
	Value       string
	Options     []SelectOption
	Placeholder string
	// Disabled closes the popup, cancels interaction, and removes the focus stop.
	Disabled bool
	// OnChange proposes Value. Nil still allows popup browsing, navigation, and dismissal.
	OnChange func(string)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:422`

## OverlayPlacement {#api-overlayplacement}

```text
OverlayPlacement selects the preferred side and alignment of a window-level
overlay. Placement automatically flips to the opposite side when it has
more usable room and the preferred side cannot fit the content.
```


```go
type OverlayPlacement uint8
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:440`

## OverlayBottomStart {#api-overlaybottomstart}


```go
const (
	OverlayBottomStart OverlayPlacement = iota
	OverlayBottom
	OverlayBottomEnd
	OverlayTopStart
	OverlayTop
	OverlayTopEnd
	OverlayLeft
	OverlayRight
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:443`

## OverlayBottom {#api-overlaybottom}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:444`

## OverlayBottomEnd {#api-overlaybottomend}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:445`

## OverlayTopStart {#api-overlaytopstart}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:446`

## OverlayTop {#api-overlaytop}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:447`

## OverlayTopEnd {#api-overlaytopend}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:448`

## OverlayLeft {#api-overlayleft}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:449`

## OverlayRight {#api-overlayright}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:450`

## PopoverProps {#api-popoverprops}

```text
PopoverProps configures a controlled, interactive window-level overlay.
Open is authoritative. Anchor activation, Escape, and an outside primary
click submit the proposed state through OnOpenChange. Nil leaves Open
unchanged and preserves open-content interaction. There is no Disabled or
ReadOnly property; an anchor child does not disable the Popover host.
```


```go
type PopoverProps struct {
	Key       string
	Style     Style
	Token     ComponentToken
	States    StateStyles
	Pointer   PointerBehavior
	Open      bool
	Placement OverlayPlacement
	Offset    MetricValue
	// OnOpenChange proposes Open. Nil leaves Open unchanged; open content remains interactive.
	OnOpenChange func(bool)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:458`

## TooltipProps {#api-tooltipprops}

```text
TooltipProps configures a non-interactive window-level hint. Hovering or
keyboard-focusing the anchor starts Delay; leaving both closes it. A zero
Delay selects the built-in 500 ms default.
```


```go
type TooltipProps struct {
	Key       string
	Style     Style
	Token     ComponentToken
	States    StateStyles
	Pointer   PointerBehavior
	Placement OverlayPlacement
	Offset    MetricValue
	Delay     time.Duration
	Disabled  bool
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:474`

## TabItem {#api-tabitem}

```text
TabItem is one immutable label in Tabs. Value must be non-empty and unique
within the component.
```


```go
type TabItem struct {
	Value    string
	Label    string
	Disabled bool
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:488`

## TabsProps {#api-tabsprops}

```text
TabsProps configures a controlled horizontal tab selector. Value is
authoritative; Tabs renders only the selector and applications render the
corresponding content separately.
```


```go
type TabsProps struct {
	Key     string
	Style   Style
	Token   ComponentToken
	States  StateStyles
	Pointer PointerBehavior
	Value   string
	Items   []TabItem
	// Disabled cancels interaction and removes the focus stop.
	Disabled bool
	// OnChange proposes a different Value. Nil retains active-item navigation and press visuals.
	OnChange func(string)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:497`

## MenuOrientation {#api-menuorientation}

```text
MenuOrientation selects the axis used to arrange Menu items.
```


```go
type MenuOrientation uint8
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:512`

## MenuVertical {#api-menuvertical}


```go
const (
	MenuVertical MenuOrientation = iota
	MenuHorizontal
)
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:515`

## MenuHorizontal {#api-menuhorizontal}

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `components.go:516`

## MenuItem {#api-menuitem}

```text
MenuItem is one immutable action in a Menu. Value must be non-empty and
unique within the component.
```


```go
type MenuItem struct {
	Value    string
	Label    string
	Disabled bool
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:521`

## MenuProps {#api-menuprops}

```text
MenuProps configures an inline action list. Value optionally identifies the
application-controlled selected item. Activating an enabled item calls
OnAction with that item's Value, including when it is already selected.
```


```go
type MenuProps struct {
	Key         string
	Style       Style
	Token       ComponentToken
	States      StateStyles
	Pointer     PointerBehavior
	Orientation MenuOrientation
	Value       string
	Items       []MenuItem
	// Disabled cancels interaction and removes the focus stop.
	Disabled bool
	// OnAction invokes an enabled item, including the selected one. Nil retains navigation and visuals.
	OnAction func(string)
}
```

[对应示例](/zh-cn/components/) · 源文件 `components.go:530`

## View {#api-view}

```text
View is an immutable description. Its representation is deliberately
private and never contains backend handles. The zero value is invalid
as a root or child; a failed tree update leaves the last valid view visible.
```


```go
type View struct{}
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:42`

## View.WithKey {#api-view-withkey}

```text
WithKey returns an independent description whose sibling-local key is key.
The receiver and any descriptions that share its node are unchanged. This
does not add a container or retained identity level.
```


```go
func (view View) WithKey(key string) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:47`

## View.WithStyle {#api-view-withstyle}

```text
WithStyle returns an independent description whose complete local Style is
style. Replacement is intentional: zero values, nil slices, and explicit
empty slices keep their normal Style meanings. This does not add a container
or retained identity level.
```


```go
func (view View) WithStyle(style Style) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:55`

## Box {#api-box}

```text
Box creates a single-line flex description. Direction defaults to Vertical.
Zero children are valid; a zero View among the supplied children is rejected
transactionally.
```


```go
func Box(props BoxProps, children ...View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:231`

## Scroll {#api-scroll}

```text
Scroll creates a clipped, one-child scrolling viewport.
```


```go
func Scroll(props ScrollProps, child View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:241`

## VirtualList {#api-virtuallist}

```text
VirtualList creates a vertical fixed-row virtualized viewport. The runtime
builds only the visible rows plus the configured bounded overscan.
```


```go
func VirtualList(props VirtualListProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:249`

## Select {#api-select}

```text
Select creates a controlled custom popup Select. Options are copied.
```


```go
func Select(props SelectProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:256`

## Popover {#api-popover}

```text
Popover creates a controlled interactive overlay around one anchor and one
content view. The content remains retained but is painted only in the
window-level overlay layer while Open is true.
```


```go
func Popover(props PopoverProps, anchor, content View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:273`

## Tooltip {#api-tooltip}

```text
Tooltip creates a delayed, non-interactive overlay around one anchor and
one content view.
```


```go
func Tooltip(props TooltipProps, anchor, content View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:282`

## Tabs {#api-tabs}

```text
Tabs creates a controlled horizontal tab selector. Items are copied.
```


```go
func Tabs(props TabsProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:300`

## Menu {#api-menu}

```text
Menu creates an inline vertical or horizontal action list with an optional
controlled selected Value. Items are copied. Vertical is the zero value.
```


```go
func Menu(props MenuProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:314`

## Text {#api-text}

```text
Text creates a text description.
```


```go
func Text(props TextProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:327`

## Label {#api-label}

```text
Label creates text with the default text style.
```


```go
func Label(value string) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:335`

## Button {#api-button}

```text
Button creates a semantic button description.
```


```go
func Button(props ButtonProps, child View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:338`

## TextButton {#api-textbutton}

```text
TextButton creates a Button whose text is centered in its content box.
Use Button directly when the content is not a simple label.
```


```go
func TextButton(props ButtonProps, label string) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:348`

## ButtonGroup {#api-buttongroup}

```text
ButtonGroup creates a non-focusable horizontal or vertical group containing
only Button views. Horizontal is the zero-value orientation.
```


```go
func ButtonGroup(props ButtonGroupProps, buttons ...View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:359`

## InputGroup {#api-inputgroup}

```text
InputGroup creates one horizontal visual control from a required Input and
optional prefix/suffix views. The supplied Input keeps its own identity and
editing state; the group suppresses only the Input's internal surface paint.
```


```go
func InputGroup(props InputGroupProps, content InputGroupContent) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:385`

## Badge {#api-badge}

```text
Badge creates a compact, non-interactive one-child label.
```


```go
func Badge(props BadgeProps, child View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:422`

## ProgressBar {#api-progressbar}

```text
ProgressBar creates a deterministic, non-interactive progress indicator.
```


```go
func ProgressBar(props ProgressBarProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:432`

## ToggleSwitch {#api-toggleswitch}

```text
ToggleSwitch creates a controlled semantic switch description.
```


```go
func ToggleSwitch(props ToggleSwitchProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:439`

## Slider {#api-slider}

```text
Slider creates a controlled single-value horizontal slider description.
```


```go
func Slider(props SliderProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:446`

## Checkbox {#api-checkbox}

```text
Checkbox creates a controlled semantic checkbox with one label child.
```


```go
func Checkbox(props CheckboxProps, label View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:453`

## Radio {#api-radio}

```text
Radio creates a controlled semantic radio button with one label child.
```


```go
func Radio(props RadioProps, label View) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:461`

## Icon {#api-icon}

```text
Icon creates a lightweight vector icon description.
```


```go
func Icon(props IconProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:485`

## Image {#api-image}

```text
Image creates a guarded raster image description.
```


```go
func Image(props ImageProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:496`

## Avatar {#api-avatar}

```text
Avatar creates a centered cover image with circular or square clipping.
```


```go
func Avatar(props AvatarProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:503`

## Input {#api-input}

```text
Input creates a controlled, single-line input description.
```


```go
func Input(props InputProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:534`

## Textarea {#api-textarea}

```text
Textarea creates a controlled multiline text editor description.
```


```go
func Textarea(props TextareaProps) View
```

[对应示例](/zh-cn/components/) · 源文件 `view.go:544`

## MaxVirtualListItems {#api-maxvirtuallistitems}

```text
MaxVirtualListItems bounds key metadata and extent arithmetic.
```


```go
const (
	// MaxVirtualListItems bounds key metadata and extent arithmetic.
	MaxVirtualListItems = 10_000_000
	// MaxVirtualListOverscan bounds work retained outside the viewport.
	MaxVirtualListOverscan = 256
)
```

[对应示例](/zh-cn/components/) · 源文件 `virtual_list.go:12`

## MaxVirtualListOverscan {#api-maxvirtuallistoverscan}

```text
MaxVirtualListOverscan bounds work retained outside the viewport.
```

类型和值见本页同组常量块。

[对应示例](/zh-cn/components/) · 源文件 `virtual_list.go:14`

