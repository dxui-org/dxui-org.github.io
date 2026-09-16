# Scroll: scrolling containers

Build scrolling containers with `Scroll`. This page includes a runnable example, all props, and behavior constraints.

## Complete example

Complete [environment setup](/guide/getting-started) first. Copy this complete `main.go` and run `go run .`.

<<< ../examples/scroll/main.go

## Parameters and API

```go
func Scroll(props ScrollProps, child View) View
```

| Field | Type | Purpose, defaults, and constraints |
| --- | --- | --- |
| `Key` | `string` | Stable unique identity within one parent; empty by default. Use business IDs in dynamic lists, not changing indices. |
| `Style` | `Style` | Local layout, paint, and text styling; zero uses intrinsic sizes and theme defaults. |
| `Token` | `ComponentToken` | Component theme entry; empty selects its default. |
| `States` | `StateStyles` | Paint patches for actual interaction states; cannot fabricate state or change layout. |
| `Pointer` | `PointerBehavior` | Defaults to PointerAuto; PointerNone excludes the entire subtree from pointer participation, unlike Disabled. |
| `Axis` | `ScrollAxis` | Defaults to ScrollVertical; children are measured without bounds on enabled axes. |
| `InitialOffset` | `Option[Point]` | Optional initial mount offset; later changes do not reset internal position. |
| `Offset` | `Option[Point]` | Optional controlled scroll position. |
| `Scrollbar` | `ScrollbarPolicy` | Defaults to ScrollbarAuto; Always always shows, Hidden hides the track but permits scrolling. |
| `OnScroll` | `func(Point)` | Proposes a complete offset; uncontrolled internal scrolling does not depend on this callback. |

## Behavior and limitations

Exactly one child. Axis defaults to ScrollVertical; ScrollHorizontal/ScrollBoth are available. Children are measured unbounded on enabled axes, so the viewport needs definite dimensions. Unset Offset retains internal state; InitialOffset applies only at first mount. With Some Offset, OnScroll only proposes changes. ScrollbarHidden still permits wheel scrolling. Nested scrollers pass remaining delta at boundaries. No inertia or smooth scrolling; ordinary Scroll mounts all content.

All components share [identity, style, and pointer rules](/api/common). See [full declarations](/api/components) for referenced enums, structs, and comments. [Running examples](/examples/) explains build verification.
