# 资源、性能与排错

遇到问题时，先查看终端中的错误，再按下面的现象定位。

## 窗口打不开

确认在有桌面环境的系统中运行，并检查 `app.Run` 返回的错误。若错误与渲染器创建有关，可在 `AppOptions` 中加入下面的配置，再试一次：

```go
dxui.AppOptions{
    Renderer: dxui.RendererSoftware,
}
```

软件渲染仍需要窗口系统，不能在没有显示环境的服务器上直接打开窗口。

## 界面不符合预期

| 现象 | 先这样检查 |
| --- | --- |
| 输入的文字留不住 | `OnChange` 是否把新值写回变量，变量是否放在 `Run` 外。 |
| 后台修改后界面不变 | 用 `App.Update` 或目标 `Window.Update` 提交赋值。 |
| 另一个窗口没有更新 | 对那个窗口单独调用 `Update` 或 `Invalidate`。 |
| 控件撑不满窗口 | 用 `RunResponsive` 给根容器设置窗口宽高。 |
| 内容超出后不能滚动 | 使用 `Scroll`，并设置明确的视口高度。 |
| 列表排序后焦点跟错行 | 给控件设置稳定的业务 `Key`。 |
| 中文显示为方块 | 检查字体是否包含中文字形，见[字体配置](/zh-cn/api/resources-guide)。 |
| 悬停后颜色变了 | 检查 `States`，交互状态会覆盖普通 `Style`。 |
| 子窗口快捷键不起作用 | 设置该窗口的 `WindowOptions.Shortcuts`。 |

## 运行变慢或内存增加

先检查三件事：

1. **事件回调是否执行耗时工作**：移到[后台任务](./async)。
2. **是否反复创建图片资源**：把 ImageSource 放到构建函数外复用，见 [Image](/zh-cn/components/image)。
3. **是否一次显示大量列表项**：固定行高列表使用 [VirtualList](/zh-cn/components/virtual-list)。

DXUI 在需要时更新界面，不需要定时调用 Update 来保持刷新。

::: details 进一步查看性能数据和能力限制

事件循环按事件和截止时间唤醒，不固定帧率空转。不要写定时 Update 来“保持刷新”。背景任务在结果或真实进度改变时提交；ProgressBar 本身不启动动画。每个窗口独立使用资源缓存预算，多个窗口的总容量会增加。资源缓存预算不等于当前工作集的硬上限，正在显示的图片仍占资源。

使用 `AppOptions.Diagnostics: true` 取得可比较的快照，观察 Build/Layout/Paint 与 FrameCount 的区别；空闲无变化时不应要求持续增加帧数。`Diagnostics()` 可跨 goroutine 读取，但不要把统计轮询作为生产 UI 更新机制。[诊断字段说明](/zh-cn/api/application#cachebudgets-与-runtimediagnostics)。

VirtualList 只构建可见区及 overscan；新 Count/Version 快照仍需 O(N) 验证所有键，不代表所有操作都是 O(可见行数)。它不支持可变行高，离屏卸载会丢失运行时编辑/焦点状态。

### 当前能力边界

没有完整 CSS Flexbox、富文本、grapheme-cluster 编辑、bidi/RTL/复杂脚本 shaping、平台无障碍桥接、动画图片、SVG 解析、RadioGroup、三态复选框、TabPanel、模态焦点陷阱。相关组合应使用现有公开类型实现，并在产品需要超出这些能力时评估框架扩展。

本站的自动验证包括源码 API 对照、完整程序编译和 VitePress 构建。原生窗口的输入、真实 IME、DPI 切换、视觉效果与平台发布仍需手动测试；“可编译”不等于“已在所有平台运行”。

:::

接下来按需要查[组件示例](/zh-cn/components/)或[公开 API](/zh-cn/api/)。
