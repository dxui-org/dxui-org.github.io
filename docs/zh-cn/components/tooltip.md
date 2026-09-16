# Tooltip：提示浮层

用 `Tooltip` 构建提示浮层。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/tooltip/main.go

## 参数与 API

```go
func Tooltip(props TooltipProps, anchor, content View) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Placement` | `OverlayPlacement` | 默认 OverlayBottomStart；共八个方向/对齐枚举，见参考。 |
| `Offset` | `MetricValue` | Scroll/VirtualList 为可选受控 Point；浮层为可选语义的 MetricValue 间距。 |
| `Delay` | `time.Duration` | time.Duration；0 使用 500ms 默认延迟，负数无效。 |
| `Disabled` | `bool` | 默认 false；交互组件为 true 时移除焦点并取消交互。 |

## 行为与边界

接受 anchor、content；鼠标悬停或键盘焦点启动 Delay，离开且失焦后关闭。Delay=0 表示默认 500ms，不是立即出现。Disabled 抑制提示。内容不可交互，不抢焦点、不阻挡指针；需要按钮等交互内容时使用 Popover。由事件/截止时间驱动，无需应用计时循环。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
