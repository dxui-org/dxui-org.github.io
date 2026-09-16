# 更新日志

记录 DXUI 的功能、公开 API 与行为变化。

## 2026-09-14 · 多窗口支持

对应 DXUI 源码提交：`4e5ee7d`。

- 新增 `App.CreateWindow`、`Window` 和 `WindowOptions`，支持创建独立子窗口。
- 新增主窗口与子窗口的标题、尺寸、最大化和最小化操作。
- 新增 `App.Invalidate`、`Window.Update` 和 `Window.Invalidate`，支持按窗口更新界面。
- 新增 `Window.Close`、`Window.Closed` 和 `Window.Diagnostics`，用于关闭子窗口、查询关闭状态和读取诊断信息。
- 多个窗口共享事件循环，各自维护界面、焦点和输入状态；快捷键按窗口独立配置。
- `App.SetTheme` 更新所有存活窗口的主题；关闭主窗口会结束应用并关闭全部子窗口。
