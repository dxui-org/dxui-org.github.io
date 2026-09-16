# VirtualList：固定行高虚拟列表

用 `VirtualList` 构建固定行高虚拟列表。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/virtual-list/main.go

## 参数与 API

```go
func VirtualList(props VirtualListProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Count` | `int` | 快照条目数；0 为空，最大 MaxVirtualListItems。 |
| `Version` | `uint64` | 快照版本 uint64；内容/键改变时递增，不能在 Build 内修改。 |
| `RowHeight` | `float32` | 每行完整有限正逻辑高度，包括间距。 |
| `Overscan` | `int` | 视口外保留的行数，0..MaxVirtualListOverscan。 |
| `InitialOffset` | `Option[Point]` | 可选首次挂载偏移；后续修改不会重置内部位置。 |
| `Offset` | `Option[Point]` | Scroll/VirtualList 为可选受控 Point；浮层为可选语义的 MetricValue 间距。 |
| `Scrollbar` | `ScrollbarPolicy` | 默认 ScrollbarAuto；Always 始终显示，Hidden 隐藏轨道但仍可滚动。 |
| `ItemKey` | `func(index int) string` | 按索引返回全列表稳定、非空、唯一键；UI 线程纯函数。 |
| `Build` | `func(index int) View` | 按索引构建行 View；只为挂载窗口构建，不可有副作用。 |
| `OnScroll` | `func(Point)` | 完整 Point 提案；未设置 Offset 时内部滚动不依赖此回调。 |

## 行为与边界

仅垂直固定行高；必须给出有限的像素 Height。RowHeight 包括间距，行被限制为该高度并裁剪。Count 上限 MaxVirtualListItems=10,000,000，Overscan 上限 MaxVirtualListOverscan=256。ItemKey 全列表非空唯一；Count/Version 是不可变快照，内容或键变更必须递增 Version。回调在 UI 线程且应纯粹。离开 overscan 即卸载，焦点、IME、编辑和弹层状态不恢复；业务值放在列表外。非受控滚动按首个幸存可见键锚定；受控 Offset 优先。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
