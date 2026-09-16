# Scroll：滚动容器

用 `Scroll` 构建滚动容器。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/scroll/main.go

## 参数与 API

```go
func Scroll(props ScrollProps, child View) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Axis` | `ScrollAxis` | 默认 ScrollVertical；启用轴无界测量子项。 |
| `InitialOffset` | `Option[Point]` | 可选首次挂载偏移；后续修改不会重置内部位置。 |
| `Offset` | `Option[Point]` | Scroll/VirtualList 为可选受控 Point；浮层为可选语义的 MetricValue 间距。 |
| `Scrollbar` | `ScrollbarPolicy` | 默认 ScrollbarAuto；Always 始终显示，Hidden 隐藏轨道但仍可滚动。 |
| `OnScroll` | `func(Point)` | 完整 Point 提案；未设置 Offset 时内部滚动不依赖此回调。 |

## 行为与边界

恰好一个 child。Axis 默认 ScrollVertical，可选 ScrollHorizontal/ScrollBoth；启用轴对子项无上限测量，因此视口需要确定大小。Offset 未设置时保留内部偏移，InitialOffset 只首次挂载使用；Some Offset 后 OnScroll 仅提案。ScrollbarHidden 仍允许滚轮。嵌套滚动在边界传递剩余量。没有惯性/平滑滚动；普通 Scroll 挂载全部内容。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
