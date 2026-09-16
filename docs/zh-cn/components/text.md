# Text：文本与 Label

用 `Text` 构建文本与 Label。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/text/main.go

## 参数与 API

```go
func Text(props TextProps) View
```

```go
func Label(value string) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Value` | `string` | 应用拥有的当前值；行为、范围和零值见本页说明。 |
| `Wrap` | `TextWrap` | 默认 TextNoWrap；TextWrapWords 简单按词换行。 |
| `MaxLines` | `int` | 正数限制可见行数；0 不限制，负数使构建失败。 |

## 行为与边界

Label(value) 等价于默认 Text；TextNoWrap 为零值。TextWrapWords 使用简单分词换行，合并空白，并非完整 Unicode 换行。MaxLines 为正时限制行数。无效 UTF-8 替换为 U+FFFD。文本本身不是可选择的编辑器；中文依赖字体覆盖。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
