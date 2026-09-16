# Common component contracts

All 23 components declare these five fields directly in their Props; there is no nested common-props wrapper.

| Field | Purpose, defaults, and constraints | Example |
| --- | --- | --- |
| Key string | Empty matches by type/position within one parent; nonempty keys must be stable and unique among siblings. Moving to another parent does not retain identity. | [State and identity](/guide/state) |
| Style Style | Zero selects defaults; local dimensions, spacing, and painting. | [Styles](/guide/style) |
| Token ComponentToken | Defaults to the component theme entry; an explicit value selects another defined entry. | [Themes](/guide/style) |
| States StateStyles | Paint patches for actual interaction states; does not create states or change layout. | [Precedence](/guide/style#style-precedence) |
| Pointer PointerBehavior | Defaults to PointerAuto; PointerNone excludes the whole subtree from pointer participation. Not a generic Disabled flag. | [Layout and hit testing](/guide/layout) |

Constructors return immutable View descriptions, not mutable widgets. `View.WithKey(string)` and `View.WithStyle(Style)` return independent descriptions; WithStyle replaces the entire Style. A zero View is invalid as a root or child; an empty Box is valid.

Events run on the UI thread. Nil callbacks have component-specific meanings and do not imply Disabled. The runtime owns hover, focus, press, IME composition, overlays, and uncontrolled scrolling. The application owns values such as Value/Checked/Open. See [state and examples](/guide/state).

See [component declarations](/api/components) for structs, enums, and functions, and [style declarations](/api/style) for shared Style and theme structures.
