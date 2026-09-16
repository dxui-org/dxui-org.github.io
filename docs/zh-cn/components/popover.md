# Popover：交互弹层

用 `Popover` 构建交互弹层。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/popover/main.go

## 参数与 API

```go
func Popover(props PopoverProps, anchor, content View) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Open` | `bool` | 受控打开状态，默认 false。 |
| `Placement` | `OverlayPlacement` | 默认 OverlayBottomStart；共八个方向/对齐枚举，见参考。 |
| `Offset` | `MetricValue` | Scroll/VirtualList 为可选受控 Point；浮层为可选语义的 MetricValue 间距。 |
| `OnOpenChange` | `func(bool)` | 打开/关闭布尔提案；nil 不改变 Open。 |

## 行为与边界

接受 anchor、content 两个 View；Open 必须由应用接纳 OnOpenChange。触发器指针/Enter/Space、顶层 Escape 或外部主键点击会提案。弹层位于窗口级，自动翻转/钳制；Tab 可转入内容，关闭时恢复仍存活触发器焦点。没有 Disabled 字段；禁用 anchor 子组件不等于禁用 Popover 宿主。无模态、焦点陷阱、箭头或动画。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
