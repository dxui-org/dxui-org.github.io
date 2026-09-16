# Component types: Full declarations

[Usage and examples](/components/) · [All APIs](/api/)

These declarations are extracted from the Go AST, retaining English source contracts and field comments while hiding private fields. `struct {}` represents an opaque implementation, not permission to use a zero value instead of its constructor. Constant blocks retain iota context.

## ButtonVariant {#api-buttonvariant}

```text
ButtonVariant selects a Button's surface and border recipe.
```


```go
type ButtonVariant uint8
```

[Related example](/components/) · Source `button.go:10`

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

[Related example](/components/) · Source `button.go:13`

## ButtonSoft {#api-buttonsoft}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:14`

## ButtonOutline {#api-buttonoutline}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:15`

## ButtonDashed {#api-buttondashed}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:16`

## ButtonGhost {#api-buttonghost}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:17`

## ButtonLink {#api-buttonlink}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:18`

## ButtonTone {#api-buttontone}

```text
ButtonTone selects semantic colors independently of the surface recipe.
```


```go
type ButtonTone uint8
```

[Related example](/components/) · Source `button.go:22`

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

[Related example](/components/) · Source `button.go:25`

## ButtonSecondary {#api-buttonsecondary}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:26`

## ButtonSuccess {#api-buttonsuccess}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:27`

## ButtonInfo {#api-buttoninfo}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:28`

## ButtonWarn {#api-buttonwarn}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:29`

## ButtonDanger {#api-buttondanger}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:30`

## ButtonDefault {#api-buttondefault}

```text
ButtonDefault is an alias for the zero-value Primary tone.
```


```go
// ButtonDefault is an alias for the zero-value Primary tone.
const ButtonDefault = ButtonPrimary
```

[Related example](/components/) · Source `button.go:34`

## ButtonSize {#api-buttonsize}

```text
ButtonSize selects overridable padding, minimum height, and inherited text metrics.
```


```go
type ButtonSize uint8
```

[Related example](/components/) · Source `button.go:37`

## ButtonNormal {#api-buttonnormal}


```go
const (
	ButtonNormal ButtonSize = iota
	ButtonSmall
	ButtonLarge
)
```

[Related example](/components/) · Source `button.go:40`

## ButtonSmall {#api-buttonsmall}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:41`

## ButtonLarge {#api-buttonlarge}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `button.go:42`

## ScrollAxis {#api-scrollaxis}

```text
ScrollAxis selects the axes whose content is measured without a maximum
and whose retained offset may change.
```


```go
type ScrollAxis uint8
```

[Related example](/components/) · Source `components.go:12`

## ScrollVertical {#api-scrollvertical}


```go
const (
	ScrollVertical ScrollAxis = iota
	ScrollHorizontal
	ScrollBoth
)
```

[Related example](/components/) · Source `components.go:15`

## ScrollHorizontal {#api-scrollhorizontal}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:16`

## ScrollBoth {#api-scrollboth}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:17`

## ScrollbarPolicy {#api-scrollbarpolicy}

```text
ScrollbarPolicy controls the overlay scrollbar. Hidden suppresses scrollbar
paint and pointer interaction without disabling wheel/trackpad scrolling.
```


```go
type ScrollbarPolicy uint8
```

[Related example](/components/) · Source `components.go:22`

## ScrollbarAuto {#api-scrollbarauto}


```go
const (
	ScrollbarAuto ScrollbarPolicy = iota
	ScrollbarAlways
	ScrollbarHidden
)
```

[Related example](/components/) · Source `components.go:25`

## ScrollbarAlways {#api-scrollbaralways}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:26`

## ScrollbarHidden {#api-scrollbarhidden}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:27`

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

[Related example](/components/) · Source `components.go:36`

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

[Related example](/components/) · Source `components.go:54`

## TextWrap {#api-textwrap}

```text
TextWrap selects the MVP simple wrapping policy.
```


```go
type TextWrap uint8
```

[Related example](/components/) · Source `components.go:73`

## TextNoWrap {#api-textnowrap}


```go
const (
	TextNoWrap TextWrap = iota
	TextWrapWords
)
```

[Related example](/components/) · Source `components.go:76`

## TextWrapWords {#api-textwrapwords}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:77`

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

[Related example](/components/) · Source `components.go:83`

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

[Related example](/components/) · Source `components.go:95`

## ButtonGroupOrientation {#api-buttongrouporientation}

```text
ButtonGroupOrientation selects the main axis used to arrange buttons.
```


```go
type ButtonGroupOrientation uint8
```

[Related example](/components/) · Source `components.go:111`

## ButtonGroupHorizontal {#api-buttongrouphorizontal}


```go
const (
	ButtonGroupHorizontal ButtonGroupOrientation = iota
	ButtonGroupVertical
)
```

[Related example](/components/) · Source `components.go:114`

## ButtonGroupVertical {#api-buttongroupvertical}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:115`

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

[Related example](/components/) · Source `components.go:120`

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

[Related example](/components/) · Source `components.go:133`

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

[Related example](/components/) · Source `components.go:144`

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

[Related example](/components/) · Source `components.go:152`

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

[Related example](/components/) · Source `components.go:163`

## PathVerb {#api-pathverb}

```text
PathVerb identifies one command in dxui's deliberately small vector icon
format. It is not an SVG parser or a general scene graph.
```


```go
type PathVerb = icondata.PathVerb
```

[Related example](/components/) · Source `components.go:174`

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

[Related example](/components/) · Source `components.go:177`

## PathLine {#api-pathline}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:178`

## PathQuad {#api-pathquad}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:179`

## PathCubic {#api-pathcubic}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:180`

## PathClose {#api-pathclose}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:181`

## PathCommand {#api-pathcommand}

```text
PathCommand stores up to three points. Move/Line use Points[0], Quad uses
Points[0:2], and Cubic uses all three points.
```


```go
type PathCommand = icondata.PathCommand
```

[Related example](/components/) · Source `components.go:186`

## Rect {#api-rect}

```text
Rect is a rectangle in logical units.
```


```go
type Rect = icondata.Rect
```

[Related example](/components/) · Source `components.go:189`

## IconData {#api-icondata}

```text
IconData is immutable dxui path data in ViewBox coordinates.
```


```go
type IconData = icondata.Data
```

[Related example](/components/) · Source `components.go:192`

## IconStrokeWidth {#api-iconstrokewidth}

```text
IconStrokeWidth is a stroke width in icon view-box units. Zero selects the
Lucide default of 2. It affects immutable stroked resources and is ignored
by legacy filled IconData.
```


```go
type IconStrokeWidth float32
```

[Related example](/components/) · Source `components.go:197`

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

[Related example](/components/) · Source `components.go:201`

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

[Related example](/components/) · Source `components.go:216`

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

[Related example](/components/) · Source `components.go:232`

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

[Related example](/components/) · Source `components.go:249`

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

[Related example](/components/) · Source `components.go:263`

## ImageSource {#api-imagesource}

```text
ImageSource is an immutable, comparable handle to application image data.
```


```go
type ImageSource struct{}
```

[Related example](/components/) · Source `components.go:284`

## ImageBytes {#api-imagebytes}

```text
ImageBytes copies encoded PNG, JPEG, or GIF bytes immediately.
```


```go
func ImageBytes(data []byte) ImageSource
```

[Related example](/components/) · Source `components.go:287`

## ImageFile {#api-imagefile}

```text
ImageFile records a path that is read by the image engine on demand.
```


```go
func ImageFile(path string) ImageSource
```

[Related example](/components/) · Source `components.go:292`

## ImageFromGo {#api-imagefromgo}

```text
ImageFromGo records an immutable-by-contract Go image source.
```


```go
func ImageFromGo(value image.Image) ImageSource
```

[Related example](/components/) · Source `components.go:295`

## ImageFit {#api-imagefit}

```text
ImageFit selects the supported destination fitting policy.
```


```go
type ImageFit uint8
```

[Related example](/components/) · Source `components.go:298`

## ImageContain {#api-imagecontain}


```go
const (
	ImageContain ImageFit = iota
	ImageCover
	ImageFill
	ImageNone
)
```

[Related example](/components/) · Source `components.go:301`

## ImageCover {#api-imagecover}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:302`

## ImageFill {#api-imagefill}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:303`

## ImageNone {#api-imagenone}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:304`

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

[Related example](/components/) · Source `components.go:309`

## AvatarShape {#api-avatarshape}

```text
AvatarShape selects the fixed Avatar clipping shape.
```


```go
type AvatarShape uint8
```

[Related example](/components/) · Source `components.go:325`

## AvatarCircle {#api-avatarcircle}


```go
const (
	AvatarCircle AvatarShape = iota
	AvatarSquare
)
```

[Related example](/components/) · Source `components.go:328`

## AvatarSquare {#api-avatarsquare}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:329`

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

[Related example](/components/) · Source `components.go:335`

## TextRange {#api-textrange}

```text
TextRange uses Unicode code-point (rune) offsets, never UTF-8 byte offsets.
```


```go
type TextRange struct{ Start, End int }
```

[Related example](/components/) · Source `components.go:350`

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

[Related example](/components/) · Source `components.go:359`

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

[Related example](/components/) · Source `components.go:389`

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

[Related example](/components/) · Source `components.go:412`

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

[Related example](/components/) · Source `components.go:422`

## OverlayPlacement {#api-overlayplacement}

```text
OverlayPlacement selects the preferred side and alignment of a window-level
overlay. Placement automatically flips to the opposite side when it has
more usable room and the preferred side cannot fit the content.
```


```go
type OverlayPlacement uint8
```

[Related example](/components/) · Source `components.go:440`

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

[Related example](/components/) · Source `components.go:443`

## OverlayBottom {#api-overlaybottom}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:444`

## OverlayBottomEnd {#api-overlaybottomend}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:445`

## OverlayTopStart {#api-overlaytopstart}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:446`

## OverlayTop {#api-overlaytop}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:447`

## OverlayTopEnd {#api-overlaytopend}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:448`

## OverlayLeft {#api-overlayleft}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:449`

## OverlayRight {#api-overlayright}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:450`

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

[Related example](/components/) · Source `components.go:458`

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

[Related example](/components/) · Source `components.go:474`

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

[Related example](/components/) · Source `components.go:488`

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

[Related example](/components/) · Source `components.go:497`

## MenuOrientation {#api-menuorientation}

```text
MenuOrientation selects the axis used to arrange Menu items.
```


```go
type MenuOrientation uint8
```

[Related example](/components/) · Source `components.go:512`

## MenuVertical {#api-menuvertical}


```go
const (
	MenuVertical MenuOrientation = iota
	MenuHorizontal
)
```

[Related example](/components/) · Source `components.go:515`

## MenuHorizontal {#api-menuhorizontal}

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `components.go:516`

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

[Related example](/components/) · Source `components.go:521`

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

[Related example](/components/) · Source `components.go:530`

## View {#api-view}

```text
View is an immutable description. Its representation is deliberately
private and never contains backend handles. The zero value is invalid
as a root or child; a failed tree update leaves the last valid view visible.
```


```go
type View struct{}
```

[Related example](/components/) · Source `view.go:42`

## View.WithKey {#api-view-withkey}

```text
WithKey returns an independent description whose sibling-local key is key.
The receiver and any descriptions that share its node are unchanged. This
does not add a container or retained identity level.
```


```go
func (view View) WithKey(key string) View
```

[Related example](/components/) · Source `view.go:47`

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

[Related example](/components/) · Source `view.go:55`

## Box {#api-box}

```text
Box creates a single-line flex description. Direction defaults to Vertical.
Zero children are valid; a zero View among the supplied children is rejected
transactionally.
```


```go
func Box(props BoxProps, children ...View) View
```

[Related example](/components/) · Source `view.go:231`

## Scroll {#api-scroll}

```text
Scroll creates a clipped, one-child scrolling viewport.
```


```go
func Scroll(props ScrollProps, child View) View
```

[Related example](/components/) · Source `view.go:241`

## VirtualList {#api-virtuallist}

```text
VirtualList creates a vertical fixed-row virtualized viewport. The runtime
builds only the visible rows plus the configured bounded overscan.
```


```go
func VirtualList(props VirtualListProps) View
```

[Related example](/components/) · Source `view.go:249`

## Select {#api-select}

```text
Select creates a controlled custom popup Select. Options are copied.
```


```go
func Select(props SelectProps) View
```

[Related example](/components/) · Source `view.go:256`

## Popover {#api-popover}

```text
Popover creates a controlled interactive overlay around one anchor and one
content view. The content remains retained but is painted only in the
window-level overlay layer while Open is true.
```


```go
func Popover(props PopoverProps, anchor, content View) View
```

[Related example](/components/) · Source `view.go:273`

## Tooltip {#api-tooltip}

```text
Tooltip creates a delayed, non-interactive overlay around one anchor and
one content view.
```


```go
func Tooltip(props TooltipProps, anchor, content View) View
```

[Related example](/components/) · Source `view.go:282`

## Tabs {#api-tabs}

```text
Tabs creates a controlled horizontal tab selector. Items are copied.
```


```go
func Tabs(props TabsProps) View
```

[Related example](/components/) · Source `view.go:300`

## Menu {#api-menu}

```text
Menu creates an inline vertical or horizontal action list with an optional
controlled selected Value. Items are copied. Vertical is the zero value.
```


```go
func Menu(props MenuProps) View
```

[Related example](/components/) · Source `view.go:314`

## Text {#api-text}

```text
Text creates a text description.
```


```go
func Text(props TextProps) View
```

[Related example](/components/) · Source `view.go:327`

## Label {#api-label}

```text
Label creates text with the default text style.
```


```go
func Label(value string) View
```

[Related example](/components/) · Source `view.go:335`

## Button {#api-button}

```text
Button creates a semantic button description.
```


```go
func Button(props ButtonProps, child View) View
```

[Related example](/components/) · Source `view.go:338`

## TextButton {#api-textbutton}

```text
TextButton creates a Button whose text is centered in its content box.
Use Button directly when the content is not a simple label.
```


```go
func TextButton(props ButtonProps, label string) View
```

[Related example](/components/) · Source `view.go:348`

## ButtonGroup {#api-buttongroup}

```text
ButtonGroup creates a non-focusable horizontal or vertical group containing
only Button views. Horizontal is the zero-value orientation.
```


```go
func ButtonGroup(props ButtonGroupProps, buttons ...View) View
```

[Related example](/components/) · Source `view.go:359`

## InputGroup {#api-inputgroup}

```text
InputGroup creates one horizontal visual control from a required Input and
optional prefix/suffix views. The supplied Input keeps its own identity and
editing state; the group suppresses only the Input's internal surface paint.
```


```go
func InputGroup(props InputGroupProps, content InputGroupContent) View
```

[Related example](/components/) · Source `view.go:385`

## Badge {#api-badge}

```text
Badge creates a compact, non-interactive one-child label.
```


```go
func Badge(props BadgeProps, child View) View
```

[Related example](/components/) · Source `view.go:422`

## ProgressBar {#api-progressbar}

```text
ProgressBar creates a deterministic, non-interactive progress indicator.
```


```go
func ProgressBar(props ProgressBarProps) View
```

[Related example](/components/) · Source `view.go:432`

## ToggleSwitch {#api-toggleswitch}

```text
ToggleSwitch creates a controlled semantic switch description.
```


```go
func ToggleSwitch(props ToggleSwitchProps) View
```

[Related example](/components/) · Source `view.go:439`

## Slider {#api-slider}

```text
Slider creates a controlled single-value horizontal slider description.
```


```go
func Slider(props SliderProps) View
```

[Related example](/components/) · Source `view.go:446`

## Checkbox {#api-checkbox}

```text
Checkbox creates a controlled semantic checkbox with one label child.
```


```go
func Checkbox(props CheckboxProps, label View) View
```

[Related example](/components/) · Source `view.go:453`

## Radio {#api-radio}

```text
Radio creates a controlled semantic radio button with one label child.
```


```go
func Radio(props RadioProps, label View) View
```

[Related example](/components/) · Source `view.go:461`

## Icon {#api-icon}

```text
Icon creates a lightweight vector icon description.
```


```go
func Icon(props IconProps) View
```

[Related example](/components/) · Source `view.go:485`

## Image {#api-image}

```text
Image creates a guarded raster image description.
```


```go
func Image(props ImageProps) View
```

[Related example](/components/) · Source `view.go:496`

## Avatar {#api-avatar}

```text
Avatar creates a centered cover image with circular or square clipping.
```


```go
func Avatar(props AvatarProps) View
```

[Related example](/components/) · Source `view.go:503`

## Input {#api-input}

```text
Input creates a controlled, single-line input description.
```


```go
func Input(props InputProps) View
```

[Related example](/components/) · Source `view.go:534`

## Textarea {#api-textarea}

```text
Textarea creates a controlled multiline text editor description.
```


```go
func Textarea(props TextareaProps) View
```

[Related example](/components/) · Source `view.go:544`

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

[Related example](/components/) · Source `virtual_list.go:12`

## MaxVirtualListOverscan {#api-maxvirtuallistoverscan}

```text
MaxVirtualListOverscan bounds work retained outside the viewport.
```

See the corresponding constant block on this page for types and values.

[Related example](/components/) · Source `virtual_list.go:14`

