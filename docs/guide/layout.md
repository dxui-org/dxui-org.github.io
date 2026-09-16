# Layout and responsive UI

Arrange navigation and content together, then switch their direction based on window width.

## 1. Run the example

<<< ../examples/responsive/main.go

Narrow the window: below 600 units, the horizontal arrangement becomes vertical.

## 2. Arrange controls with Box

| Configuration | Effect |
| --- | --- |
| `Direction: dxui.Horizontal` | Arranges left to right. |
| `Direction: dxui.Vertical` | Arranges top to bottom; the default. |
| `Gap: 16` | Leaves 16 logical units between children. |
| `Padding: dxui.Padding(24)` | Leaves 24 logical units around the container interior. |

Gap belongs in BoxProps; Padding belongs in its Style. Logical units scale with the screen and are not physical pixels.

## 3. Allocate width and height

The navigation uses `Width: dxui.Px(160)` and `Shrink: dxui.NoShrink()` to keep its width. The content uses `Grow: 1` to occupy remaining space along the layout direction.

RunResponsive supplies the current ctx.Width and ctx.Height. Use them for the root container size, then choose horizontal or vertical based on ctx.Width.

Try changing the breakpoint from 600 to 800 and observe the result.

## When content does not fit

Overflow does not scroll automatically. Use [Scroll](/components/scroll) with a definite height. For many fixed-height rows, use [VirtualList](/components/virtual-list).

::: details More sizing, positioning, and clipping rules

| Setting | Meaning and constraints |
| --- | --- |
| Unset Width/Height | Auto, determined by content and parent constraints. |
| Px(n) | Nonnegative fixed logical size. |
| Percent(50) | Half of a definite parent axis; otherwise auto. |
| Fill() | Percent(100), also requiring a definite parent size. |
| Grow: 1 | Proportional allocation of remaining main-axis space. |
| Unset Shrink | Defaults to 1; may shrink when space is insufficient. |
| Shrink: NoShrink() | Explicit 0; prevents shrinking and may overflow. |
| Basis | Flex base size; supports auto, Px, and Percent. |
| MinWidth/MaxWidth, etc. | Bound final dimensions; a maximum below the minimum is raised to the minimum. |

A child Fill cannot automatically fill the window when the parent height depends on its children. Give the root explicit responsive dimensions when it must fill the window.

### Scrolling is explicit

Overflow does not scroll by default. Give Scroll a definite viewport height; it measures content unbounded on enabled axes. See [Scroll](/components/scroll). For thousands of rows, use [VirtualList](/components/virtual-list) with its fixed-height and unmounting rules.

### Absolute positioning, clipping, and stacking

Use `Style.Position=PositionAbsolute` with `Insets{Top, Right, Bottom, Left}`. Absolute children do not participate in flex allocation or Gap counting. ZIndex changes only sibling paint/hit order, not layout or source-order Tab navigation.

OverflowClip uses rectangular clipping. Radius does not imply rounded subtree clipping. Borders paint inside the rectangle but **consume no layout space**; use Padding for content spacing. Shadows do not affect layout or hit areas.

Hidden preserves layout but does not paint or hit-test. Opacity=0 subtrees are also excluded from hits. A transparent background alone does not remove the geometric hit area.

Only single-line flex is supported. No wrapping, order, baseline, negative spacing, auto margins, percentage spacing, space-around/evenly, sticky positioning, or full CSS layout. See [layout and style values](/api/values) for all fields.


:::

Next: [Styles and themes](./style).
