# 窗口管理 API

一个 App 管理主窗口及零个或多个独立顶层子窗口。组件 API 仍统一来自 dxui 根包；Window 是生命周期句柄，不能作为 View 使用，也不要自行构造零值 Window。先读[多窗口教程](/zh-cn/guide/multi-window)，精确声明见[运行时参考](/zh-cn/api/runtime)。

## 创建与调度

| API | 参数与结果 | 范围、限制与错误 |
| --- | --- | --- |
| App.CreateWindow(options WindowOptions, root func() View) (*Window, error) | 配置和非 nil 根构建函数；成功返回子窗口句柄。 | UI 回调或 App.Update 内调用。准备首帧后显示；失败清理已创建资源，不影响已有窗口。 |
| App.Update(update func()) error | FIFO 排队执行非 nil 状态修改。 | 仅主窗口；可跨 goroutine。未运行/关闭返回 ErrAppNotRunning/ErrAppClosed。 |
| App.Invalidate() error | 请求主窗口重新构建，不附带状态修改。 | 可跨 goroutine；生命周期错误同 App.Update。不能替代后台状态同步。 |
| Window.Update(update func()) error | FIFO 排队执行非 nil 状态修改。 | 仅目标子窗口；可跨 goroutine。已关闭或 nil 句柄返回 ErrWindowClosed，App 正在退出可返回 ErrAppClosed。 |
| Window.Invalidate() error | 请求目标子窗口重建。 | 相当于对该窗口提交空更新；生命周期错误同 Window.Update。 |
| Window.Close() | 请求关闭目标子窗口，无返回值。 | 幂等、可跨 goroutine；不关闭主窗口或兄弟窗口，不调用 OnCloseRequest。nil 句柄为空操作。 |
| Window.Closed() bool | 返回已请求或完成关闭。 | 并发安全，nil 句柄返回 true；不能作为下一次操作必定成功的保证。 |
| Window.Diagnostics() RuntimeDiagnostics | 返回该窗口的计数快照。 | 并发安全；nil 句柄返回零快照。[字段解释](/zh-cn/api/application#cachebudgets-与-runtimediagnostics)。 |

App.Close 关闭全部窗口并结束唯一事件循环；App.Diagnostics 仍以主窗口为对象，不是全部窗口统计之和。关闭后不要继续提交依赖该窗口存活的工作。

## WindowOptions 全部字段

| 字段 | 类型 | 默认值与行为 |
| --- | --- | --- |
| Title | string | 空字符串默认 dxui。 |
| Width / Height | float32 | 默认 640 / 480 逻辑单位；不同于主窗口默认的 800 / 600。有限非负，不超过 MaxInt32；显式正值四舍五入后至少为 1。 |
| MinWidth / MinHeight | float32 | 0 不设置最小尺寸；有限非负，不超过 MaxInt32。 |
| Background | RGBAColor | 全零使用 App 配置的背景；若仍全零，回退 RGBA(28,30,36,255)。 |
| Shortcuts | []Shortcut | 独立、复制的窗口快捷键；默认空，不自动继承 AppOptions.Shortcuts。 |
| OnCloseRequest | func(*Window) | nil 接受系统关闭请求；非 nil 时调用 Window.Close 才接受，否则拒绝。 |
| OnShown | func(*Window) | 完整首帧呈现并显示后仅调用一次；在 CreateWindow 返回前执行，使用回调传入的句柄。 |

Theme、Fonts、DefaultFont、DisableSystemFontFallback、Renderer、Caches、Diagnostics 等使用 App 配置。窗口不暴露独立 SetTheme：通过 App.SetTheme 更新所有存活窗口。预算按窗口分别使用，创建更多窗口会增加总资源容量。

源码 WindowOptions 注释目前笼统写了继承 App 快捷键，但创建实现明确用 options.Shortcuts 替换。因此本页与展示参考按实现说明为窗口独立配置，完整原始注释保留在维护快照中。

## 主窗口与子窗口通用操作

下面每个方法同时存在于 `*App`（主窗口）和 `*Window`（指定子窗口）。带原生操作的 SetTitle、SetSize、Maximize、Unmaximize、Minimize、Unminimize 必须在 UI 回调或 Update 闭包执行；Title、Size、IsMaximized、IsMinimized 是并发安全快照。

| 方法 | 返回值 | 用途与限制 |
| --- | --- | --- |
| Title() | string | 最近成功配置的标题；App 未运行时返回 AppOptions 的原始标题。nil Window 返回空字符串。 |
| SetTitle(title string) | error | 成功更新标题快照，不重建根；空字符串按请求设置，不套用创建时默认标题。 |
| Size() | Size | 逻辑客户区尺寸；成功 SetSize 与后续 resize/scale 事件均更新快照。App 未运行时返回配置尺寸；nil Window 返回零 Size。 |
| SetSize(width, height float32) | error | 两轴必须有限正数，四舍五入后至少 1，不超过 MaxInt32。请求成功更新快照，实际布局跟随后续视口事件，平台仍可能调整最终尺寸。 |
| Maximize() | error | 请求平台最大化；返回 nil 不代表状态事件已确认。 |
| Unmaximize() | error | 请求恢复正常状态。 |
| Minimize() | error | 请求平台最小化；返回 nil 不代表状态事件已确认。 |
| Unminimize() | error | 请求恢复正常状态；与 Unmaximize 共用平台恢复操作。 |
| IsMaximized() | bool | 最近窗口事件确认的最大化状态，不根据请求乐观设置；未创建主窗口或 nil Window 返回 false。 |
| IsMinimized() | bool | 最近窗口事件确认的最小化状态；未创建主窗口或 nil Window 返回 false。 |

主窗口命令在运行前返回 ErrWindowNotRunning，在 App 关闭后返回 ErrAppClosed。子窗口关闭后的命令返回 ErrWindowClosed；标题/尺寸/状态读取仍是快照，不用它们判断窗口是否存活。

WindowOptions 不提供状态事件回调。若把 IsMaximized/IsMinimized 显示为文字，不应在发出请求后立即假定它已变化；下面的 Read confirmed state 按钮在点击时重新读取已确认状态。

## 完整操作示例

这个例子给主窗口和子窗口提供同样的标题、尺寸、最大化与最小化按钮。子窗口首次构建早于 CreateWindow 返回，因此通过函数取得句柄，只在之后的按钮事件中执行操作。

<<< ../../examples/window-controls/main.go

最小化后可从系统任务栏或窗口管理器恢复窗口，或通过主窗口的 Restore child 按钮请求恢复。Refresh root 演示目标 Invalidate，Snapshot 演示目标 Diagnostics；这些接口不会创建轮询循环。

## 生命周期错误

| 错误 | 对应场景 |
| --- | --- |
| ErrWindowNotRunning | CreateWindow 在 App.Run 之外执行，或主窗口命令在运行前执行。 |
| ErrWindowClosed | nil/已关闭子窗口的 Update、Invalidate 或窗口命令。 |
| ErrAppClosed | App 正在退出时提交工作或执行主窗口命令。 |
| 其它 error | nil builder/更新闭包、非法尺寸、快捷键校验、首帧构建或平台操作失败。 |

使用 errors.Is 区分生命周期错误，处理返回值，不匹配错误文本。错误校验有先后顺序；例如 CreateWindow 在运行时已销毁后可能首先返回 ErrWindowNotRunning，不应把所有退出阶段一律看成同一种错误。
