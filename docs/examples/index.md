# Running complete examples

Component pages include the same `docs/examples/<component>/main.go` files used by compilation checks. UI code uses only the root package and official icons. Image/Avatar examples generate pixels themselves and need no external files.

In an empty directory, follow the quick start's `go mod init` and `go get`, copy an example to `main.go`, and run `go run .`. These are native window programs, not Go running in the browser.

- [Multiple windows and shared state updates](/guide/multi-window#complete-example)
- [Main and child window controls](/api/windows#complete-window-controls-example)

- [Box complete example](/components/box#complete-example)
- [Text complete example](/components/text#complete-example)
- [Button complete example](/components/button#complete-example)
- [ButtonGroup complete example](/components/button-group#complete-example)
- [InputGroup complete example](/components/input-group#complete-example)
- [Input complete example](/components/input#complete-example)
- [Textarea complete example](/components/textarea#complete-example)
- [Checkbox complete example](/components/checkbox#complete-example)
- [Radio complete example](/components/radio#complete-example)
- [ToggleSwitch complete example](/components/toggle-switch#complete-example)
- [Slider complete example](/components/slider#complete-example)
- [Select complete example](/components/select#complete-example)
- [Tabs complete example](/components/tabs#complete-example)
- [Menu complete example](/components/menu#complete-example)
- [Scroll complete example](/components/scroll#complete-example)
- [VirtualList complete example](/components/virtual-list#complete-example)
- [Popover complete example](/components/popover#complete-example)
- [Tooltip complete example](/components/tooltip#complete-example)
- [Icon complete example](/components/icon#complete-example)
- [Image complete example](/components/image#complete-example)
- [Avatar complete example](/components/avatar#complete-example)
- [Badge complete example](/components/badge#complete-example)
- [ProgressBar complete example](/components/progress-bar#complete-example)

Maintainers can run `python scripts/check-docs.py <dxui-source>` in the website root to compile every complete example and compare the public API snapshot. Successful compilation does not replace native interaction testing.
