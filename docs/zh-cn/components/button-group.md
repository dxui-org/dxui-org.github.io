# ButtonGroup：连接按钮组

用 `ButtonGroup` 构建连接按钮组。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/button-group/main.go

## 参数与 API

```go
func ButtonGroup(props ButtonGroupProps, buttons ...View) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Orientation` | `ButtonGroupOrientation` | 默认水平（ButtonGroup）或垂直（Menu）；详见本页。 |
| `Dividers` | `bool` | 默认 false；true 绘制组内主题分隔边。 |

## 行为与边界

仅接收 Button（含 TextButton），允许零个。默认水平且没有分隔线；Orientation=ButtonGroupVertical 改为纵向。Dividers 启用主题边框重叠，使共享边只绘制一次。组无独立焦点，每个按钮保留 Disabled、回调和源顺序 Tab；没有方向键导航。纵向自动宽度对齐最宽项，显式子 Width 优先。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
