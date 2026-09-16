# 运行完整示例

每个组件页引用本站 `docs/examples/<组件>/main.go` 的同一份代码，避免展示代码与编译代码分叉。界面 API 仅使用根包与官方 icon 资源；Image/Avatar 示例自行创建像素，不需要外部文件。

在任意空目录执行快速开始的 `go mod init`、`go get`，把所选示例复制为 `main.go` 后 `go run .`。这些是原生窗口程序，不是在浏览器中运行的 Go。

- [多窗口与跨窗口状态更新](/zh-cn/guide/multi-window#完整示例)
- [主窗口与子窗口控制](/zh-cn/api/windows#完整操作示例)

- [Box 完整示例](/zh-cn/components/box#完整示例)
- [Text 完整示例](/zh-cn/components/text#完整示例)
- [Button 完整示例](/zh-cn/components/button#完整示例)
- [ButtonGroup 完整示例](/zh-cn/components/button-group#完整示例)
- [InputGroup 完整示例](/zh-cn/components/input-group#完整示例)
- [Input 完整示例](/zh-cn/components/input#完整示例)
- [Textarea 完整示例](/zh-cn/components/textarea#完整示例)
- [Checkbox 完整示例](/zh-cn/components/checkbox#完整示例)
- [Radio 完整示例](/zh-cn/components/radio#完整示例)
- [ToggleSwitch 完整示例](/zh-cn/components/toggle-switch#完整示例)
- [Slider 完整示例](/zh-cn/components/slider#完整示例)
- [Select 完整示例](/zh-cn/components/select#完整示例)
- [Tabs 完整示例](/zh-cn/components/tabs#完整示例)
- [Menu 完整示例](/zh-cn/components/menu#完整示例)
- [Scroll 完整示例](/zh-cn/components/scroll#完整示例)
- [VirtualList 完整示例](/zh-cn/components/virtual-list#完整示例)
- [Popover 完整示例](/zh-cn/components/popover#完整示例)
- [Tooltip 完整示例](/zh-cn/components/tooltip#完整示例)
- [Icon 完整示例](/zh-cn/components/icon#完整示例)
- [Image 完整示例](/zh-cn/components/image#完整示例)
- [Avatar 完整示例](/zh-cn/components/avatar#完整示例)
- [Badge 完整示例](/zh-cn/components/badge#完整示例)
- [ProgressBar 完整示例](/zh-cn/components/progress-bar#完整示例)

维护者在网站根目录执行 `python scripts/check-docs.py <dxui源码路径>`，会逐个编译完整示例并核对公开 API 快照。编译成功不能替代本机窗口交互验收。
