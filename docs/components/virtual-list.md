# VirtualList: fixed-height virtual lists

Build fixed-height virtual lists with `VirtualList`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/virtual-list/main.go

## Parameters and API

```go
func VirtualList(props VirtualListProps) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Count` | `int` | Snapshot item count; 0 is empty, maximum MaxVirtualListItems. |
| `Version` | `uint64` | Increment when content/keys change; do not modify inside Build. |
| `RowHeight` | `float32` | Complete finite positive logical row height, including spacing. |
| `Overscan` | `int` | Retained rows outside the viewport, 0..MaxVirtualListOverscan. |
| `InitialOffset` | `Option[Point]` | Optional initial mount offset; later changes do not reset internal position. |
| `Offset` | `Option[Point]` | Optional controlled scroll position. |
| `Scrollbar` | `ScrollbarPolicy` | Defaults to ScrollbarAuto; Always always shows, Hidden hides the track but permits scrolling. |
| `ItemKey` | `func(index int) string` | Returns stable nonempty keys unique across the full list; a pure UI-thread function. |
| `Build` | `func(index int) View` | Builds a row by index only in the mounted range; must have no side effects. |
| `OnScroll` | `func(Point)` | Proposes a complete offset; uncontrolled internal scrolling does not depend on this callback. |

## Behavior and limitations

Vertical fixed-height rows only; supply a finite pixel Height. RowHeight includes spacing; each row is constrained and clipped to it. MaxVirtualListItems=10,000,000; MaxVirtualListOverscan=256. ItemKey must be nonempty and globally unique within the list. Count/Version describe an immutable snapshot; increment Version for content/key changes. Callbacks run on the UI thread and must be pure. Leaving overscan unmounts rows and discards focus, IME, editor, and overlay state; keep business values outside the list. Uncontrolled scrolling anchors to the first surviving visible key; controlled Offset takes precedence.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
