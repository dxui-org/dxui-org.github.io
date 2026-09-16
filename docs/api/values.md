# Geometry, layout, and paint values

This page groups shared APIs by purpose. See [style declarations](/api/style) for exact fields, enums, and contracts, and [runtime declarations](/api/runtime) for coordinates. The [layout](/guide/layout) and [style](/guide/style) examples demonstrate composition.

## Coordinates and constructors

| API | Parameters and purpose | Zero values, order, and constraints |
| --- | --- | --- |
| Point{X,Y} / Size{Width,Height} / Rect{X,Y,Width,Height} | float32 logical positions, offsets, sizes, and rectangles; view-box coordinates for icons. | No backend objects. |
| Px(float32) Length | Explicit logical length. | Layout requires finite nonnegative values; zero Length means auto, unlike Px(0). |
| Percent(float32) / Fill() | Percentage 0..100; Fill equals Percent(100). | Acts as auto when the parent axis is indefinite. |
| Metric(float32) MetricValue | Explicit logical metric, including 0. | Distinct from Length; used for padding, borders, etc. Constraints depend on the property. |
| TokenMetric(MetricToken) | Reads a theme metric. | Missing or cyclic references are rejected. |
| NoShrink() Option[float32] | Some(float32(0)). | Unlike the default shrink factor of 1. |
| Padding(n) / Margin(n) | Uniform logical spacing on four sides. | Returns EdgeValues with explicit zero support. |
| PaddingXY(x,y) / MarginXY(x,y) | Horizontal and vertical spacing. | x is left/right; y is top/bottom. |
| Edges(t,r,b,l) | Explicit values for four sides. | Top, right, bottom, left; negative spacing is unsupported. |
| UniformEdges(MetricValue) | One literal/token metric for all sides. | Zero EdgeValues fields are unset; each can be configured separately. |
| Round(n) / Corners(tl,tr,br,bl) | Uniform or per-corner logical radii. | Top-left, top-right, bottom-right, bottom-left; clamped to geometry when painted. |
| UniformCorners(MetricValue) | Uniform literal/token radius. | Zero CornerValues fields are unset. |
| RGBA(r,g,b,a) | Non-premultiplied uint8 RGBAColor. | 0..255; a=0 is transparent. |
| ColorRGBA(r,g,b,a) / LiteralColor(RGBAColor) | Returns a ColorValue for styles. | Supports explicit transparency; RGBAColor and ColorValue are not interchangeable. |
| TokenColor(ColorToken) | Returns a token-reference ColorValue. | Resolved by the current theme. |
| Stroke(width,color) / NoBorder() | Solid Border / explicit zero border width. | Later state patches can still override a normal Style. |

## Every Style field

| Field | Purpose and default |
| --- | --- |
| Width/Height, MinWidth/MinHeight, MaxWidth/MaxHeight | Length dimensions and limits, default auto; minima win over conflicting maxima. |
| Margin/Padding | Outer/inner spacing through EdgeValues; components may provide defaults when unset. |
| Position | Defaults to PositionFlow; PositionAbsolute leaves flex flow. |
| Insets | Top/Right/Bottom/Left Length constraints for absolute positioning. |
| Grow | Nonnegative growth factor, default 0. |
| Shrink | Option[float32] nonnegative shrink factor, default 1 when unset. |
| Basis | Flex base Length, default auto; shrink weights include the original basis. |
| AlignSelf | Option[Align]; inherits parent Align when unset. |
| ZIndex | Sibling paint-order integer, default 0; ties use source order. |
| Overflow | Defaults to OverflowVisible; OverflowClip clips to a rectangle. |
| Background | ColorValue background; unset retains the component appearance. |
| Border | Painted inside without taking layout space; see below. |
| Radius | Four CornerValues; does not imply rounded descendant clipping. |
| Shadow | Up to 4 outer shadows; nil inherits, non-nil empty slice clears. |
| Opacity | Option[float32], 0..1, multiplied into the subtree; 0 excludes hit testing. |
| Visibility | Defaults to Visible; Hidden retains layout but does not paint or hit-test. |
| Text | Text styling and inherited color; see below. |
| Force | Final StylePatch; painting only, cannot change layout. |

Align has Start/Center/End/Stretch; Justify has Start/Center/End/SpaceBetween; Direction has Vertical/Horizontal. Use the exported enum constants, not arbitrary integers.

## Border, Shadow, and TextStyle

Border.Width is MetricValue; Color is ColorValue; Pattern is BorderSolid (default) or BorderDashed. Sides combines BorderTop/Right/Bottom/Left bits; Sides=0 means all sides, **not no border**. Setting Width/Color in a normal Style also resets Sides. StylePatch.Border replaces the entire Border.

Shadow OffsetX/OffsetY and Blur/Spread are MetricValue; Color is ColorValue. Offsets may be signed; Blur/Spread must be finite and nonnegative. Shadows use bounded geometry approximations, not CSS Gaussian blur, and do not affect layout or hit testing.

TextStyle.Families is an ordered []FontFamily. Size/LineHeight are logical metrics (0 uses theme defaults). Weight selects the nearest registered weight; constants include WeightRegular=400, WeightMedium=500, and WeightBold=700. SlantNormal/SlantItalic select font faces without promising synthesized variants. Color controls text; Align is TextStart/Center/End. Supports simple LTR/CJK text, not Arabic/Indic shaping, bidi/RTL, color emoji, or vertical text.

## Override patches

StylePatch Background, Border, Radius, Shadow, Opacity, Visibility, and TextColor are Options. Unset inherits; Some explicitly overrides, including zero. StateStyles has Default, Hover, Focus, Checked, Pressed, and Disabled StylePatch fields. See [style precedence](/guide/style).
