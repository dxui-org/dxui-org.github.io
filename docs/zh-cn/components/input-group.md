# InputGroup：带前后缀的输入框

用 `InputGroup` 构建带前后缀的输入框。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/input-group/main.go

## 参数与 API

```go
func InputGroup(props InputGroupProps, content InputGroupContent) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |

## 行为与边界

InputGroupContent.Input 必须恰好是一个 Input；Prefix/Suffix 用 Some 显式提供，可包含非交互内容或 Button，不能包含其它可聚焦控件或嵌套编辑器。组统一绘制背景、边框、圆角和 focus-within，Input 使用剩余宽度；编辑器状态与 Key 保留。后缀按钮是独立 Tab 停靠点。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
