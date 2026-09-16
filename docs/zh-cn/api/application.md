# 应用、事件与诊断 API

组件之外的运行入口都在 `github.com/dxui-org/dxui`。本页解释用途与使用条件，精确签名、全部字段和源码注释见[运行时声明](/zh-cn/api/runtime)。

| API | 参数与返回值 | 使用条件与对应示例 |
| --- | --- | --- |
| NewApp(AppOptions) *App | 复制配置中的字体/快捷键等描述，返回管理主窗口和子窗口的 App。 | 不创建窗口；[快速开始](/zh-cn/guide/getting-started)。 |
| App.Run(func() View) error | 无参 root 返回有效 View；阻塞到关闭或失败。 | main 线程直接调用，一次性 App；nil root 拒绝。 |
| App.RunResponsive(func(LayoutContext) View) error | 上下文提供逻辑 Width/Height。 | 合并 resize/scale 后重建；[响应式完整程序](/zh-cn/guide/layout)。 |
| App.Update(func()) error | 排队修改主窗口状态，FIFO UI 执行，一批后重建主窗口。 | 只接受运行中的非 nil 提交；[后台任务](/zh-cn/guide/async)。 |
| App.Invalidate() error | 仅请求主窗口重建。 | 可跨 goroutine；子窗口用 Window.Invalidate。[多窗口更新](/zh-cn/guide/multi-window)。 |
| App.CreateWindow(WindowOptions, func() View) (*Window, error) | 创建独立顶层子窗口，首帧成功后显示。 | UI 回调或 App.Update 内调用；[多窗口教程](/zh-cn/guide/multi-window)。 |
| App.Close() | 请求关闭主窗口和全部子窗口，无返回值。 | 跨 goroutine 安全且幂等；Run 前为空操作。 |
| App.SetTheme(Theme) error | 校验、复制主题并传播到存活窗口。 | 运行中在 UI 回调或 Update 内调用；[主题示例](/zh-cn/guide/style)。 |
| App.SetClipboardText(string) error | 写入 UTF-8 文本，无效 UTF-8 替换为 U+FFFD。 | UI 回调/Update 中执行；关闭、未运行或原生剪贴板失败返回错误。 |
| App.Diagnostics() RuntimeDiagnostics | 主窗口的线程安全值快照；子窗口用 Window.Diagnostics。 | 可在运行中或结束后取样；统计数据含义见下文。 |
| `Assign[T](*T) func(T)` | 创建纯赋值回调；nil 指针立即 panic。 | 用于 UI 回调，不能替代 Update。 |
| `Some[T](T) Option[T]` | 显式提供一个值，包括零值。 | Option 零值为未设置，没有公开 get 方法。 |
| ErrAppNotRunning / ErrAppClosed | 可用 errors.Is 匹配的生命周期错误。 | [后台任务关闭处理](/zh-cn/guide/async)。 |

App 与 Window 都提供标题、尺寸、最大化、最小化和状态读取方法；全部参数、线程要求与完整程序见[窗口管理 API](/zh-cn/api/windows)。

## AppOptions 字段

| 字段 | 用途、默认与限制 |
| --- | --- |
| Title | 窗口标题；空字符串默认 dxui。 |
| Width / Height | 初始逻辑尺寸，零值分别默认 800 / 600；有限非负且不超过 MaxInt32，显式正值四舍五入后须至少为 1。 |
| MinWidth / MinHeight | 窗口最小逻辑尺寸，0 不设置限制；非法尺寸由 Run 报错。 |
| Renderer | RendererAuto 默认；RendererSoftware 请求命名的软件渲染器。 |
| Background | 原始 RGBAColor 窗口清屏颜色，用 RGBA 构造；全零默认 RGBA(28,30,36,255)。教程显式使用浅色背景。 |
| Caches | CacheBudgets，0 选默认，负数禁用相应缓存；预算按窗口分别应用。 |
| Fonts | 从 FontBytes/FontFile/SystemFont 得到的 Font 列表；启动加载，错误由 Run 返回。 |
| DefaultFont | 默认 FontFamily；配合注册字体使用。 |
| DisableSystemFontFallback | 默认 false；true 关闭自动系统 CJK 回退，不改变显式注册字体。 |
| Theme | 零 Theme 选内置亮色；部分非零主题不自动补齐所有缺失映射。 |
| Shortcuts | []Shortcut，复制主窗口快捷键；子窗口通过 WindowOptions.Shortcuts 独立配置。 |
| OnCloseRequest | 主窗口关闭请求；nil 默认关闭全部窗口，非 nil 时由回调决定是否调用 App.Close。 |
| OnError | 在 UI 线程观察可恢复构建/回调失败；nil 时终止 Run 并返回错误。 |
| OnShown | 主窗口完整首帧呈现并显示成功后仅调用一次；子窗口使用 WindowOptions.OnShown。 |
| Diagnostics | 默认 false；true 开启有界事件/时间/资源统计。 |

## 快捷键、剪贴板、关闭和诊断完整示例

<<< ../../examples/lifecycle/main.go

Shortcut.Key 只接受当前公开 ShortcutKey：Enter、Backspace、0..9、Plus、Minus、Multiply、Divide、Decimal、Equals；不支持任意字符串键名。Modifiers 的 Shift/Control/Alt/Super/Primary 默认全 false，精确匹配。OnPress 是无参动作，Repeat 默认 false。编辑器和控件语义按键优先，因此例子使用 Primary+1 而非争用输入框 Enter。

## CacheBudgets 与 RuntimeDiagnostics

以下预算按窗口分别使用，当前各窗口 CPU 缓存也独立；应用总容量随存活窗口数增长。Window.Diagnostics 查看子窗口资源，App.Diagnostics 不是全窗口汇总。

| 缓存字段 | 零值默认预算 | 限制含义 |
| --- | --- | --- |
| FontBytes | 32 MiB | 应用字体和自动 CJK 字体源预算；负值不允许应用字体。 |
| TextSourceBytes | 2 MiB | 保留文本/图标掩码源；负值不允许这些保留源。 |
| GlyphBytes | 2 MiB | 字形缓存。 |
| TextMeasureBytes | 1 MiB | 文本测量缓存。 |
| ImageBytes | 2 MiB | 非活动、可复用的 CPU 图片/原生纹理；当前显示的唯一图片为工作集，不由此硬性封顶。 |
| ShadowBytes | 2 MiB | 有界阴影几何缓存。 |

RuntimeDiagnostics 的 RendererName/SoftwareFallback 用来确认渲染方式；LogicalSize/PixelSize/PixelDensity/DisplayScale 区分逻辑尺寸与像素缩放。FrameCount、BuildCount、LayoutCount、PaintCount、ReconcileCount 和 PaintNodeCount 区分呈现、构建、布局、绘制与树工作。

WindowCreates、RendererCreateAttempts/Creates、ExposeEvents、ResizeEvents、ScaleEvents、NoopViewportEvents、RendererResetEvents 用于生命周期与视口诊断；EventCount 为事件数。TextureCreates/Destroys、CacheBytes/BudgetBytes/Entries、FontResources/ImageResources/RendererResources 用于资源观察。Goroutines 和 GoHeapBytes/Objects、GoTotalAllocBytes、GoMallocs 是进程 Go 统计快照，不可解读为一个组件独占内存。

EventToPresent 和 FrameTime 是 TimingSummary：Count 为总计数，Samples 为有界保留样本数，P50NS/P95NS/P99NS 单位为纳秒。先检查 CountersEnabled；未开启、无样本或平台未提供的数据不能解释为性能为零。字段类型逐项见[RuntimeDiagnostics](/zh-cn/api/runtime#api-runtimediagnostics)。
