# Changelog

DXUI features, public APIs, and behavior changes.

## 2026-10-15 · Typed pointer cursors

DXUI source commit: `ea1f823`.

- Added `Cursor`, `CursorAuto`, `CursorDefault`, `CursorText`, and `CursorPointer`.
- Added `Style.Cursor` for typed hover-pointer overrides without changing focus or activation behavior.
- `CursorAuto` uses the text I-beam for Input and Textarea, the pointing hand for Button, and the default arrow elsewhere.
- Decorative children follow the enclosing control's effective cursor unless they explicitly override it.

## 2026-09-14 · Multiple windows

DXUI source commit: `4e5ee7d`.

- Added `App.CreateWindow`, `Window`, and `WindowOptions` for independent child windows.
- Added title, size, maximize, and minimize operations for main and child windows.
- Added `App.Invalidate`, `Window.Update`, and `Window.Invalidate` for updates targeting individual windows.
- Added `Window.Close`, `Window.Closed`, and `Window.Diagnostics` for child shutdown, close status, and diagnostic snapshots.
- Windows share an event loop but maintain independent UI, focus, and input state. Shortcuts are configured per window.
- `App.SetTheme` updates all live windows; closing the main window ends the application and closes all children.
