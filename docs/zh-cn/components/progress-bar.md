# ProgressBar：进度条

用 `ProgressBar` 构建进度条。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/progress-bar/main.go

## 参数与 API

```go
func ProgressBar(props ProgressBarProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Value` | `float32` | 应用拥有的当前值；行为、范围和零值见本页说明。 |

## 行为与边界

Value 为 0..1 比例，不是百分数；越界钳制，NaN/-Inf 为空，+Inf 为满。无事件和内部进度计时。Background 改轨道色，Text.Color 改填充色；默认 160×12，轨道高 8。没有不确定进度、分段、竖向或动画模式。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
