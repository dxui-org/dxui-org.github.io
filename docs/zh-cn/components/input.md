# Input：单行输入

用 `Input` 构建单行输入。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/input/main.go

## 参数与 API

```go
func Input(props InputProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Value` | `string` | 应用拥有的当前值；行为、范围和零值见本页说明。 |
| `OnChange` | `func(string)` | 完整新值提案；写回业务变量后下一次构建生效。nil 行为见本页。 |
| `Selection` | `Option[TextRange]` | 默认未设置，使用内部选区；Some(TextRange) 后为受控 rune 选区。 |
| `OnSelectionChange` | `func(TextRange)` | 完整 rune 选区提案；nil 保留内部选区，受控 Selection 在重建时优先。 |
| `Placeholder` | `string` | 空 Value 时显示；默认空字符串。 |
| `Password` | `bool` | 默认 false；true 遮蔽显示并禁止复制/剪切。 |
| `ShowPasswordToggle` | `bool` | 默认 false；仅 Password=true 生效，保留身份期间记住显示状态。 |
| `Disabled` | `bool` | 默认 false；交互组件为 true 时移除焦点并取消交互。 |
| `ReadOnly` | `bool` | 默认 false；true 禁止编辑和预编辑，保留导航与选择。 |
| `OnSubmit` | `func()` | Input 的 Enter 回调；ReadOnly 时仍可调用，Disabled 时不可。 |

## 行为与边界

Value 是完整受控字符串；不写回 OnChange 提案就拒绝编辑。Selection 用 rune 下标，Some 后必须接纳 OnSelectionChange 才能移动选区。ReadOnly 保留导航、选区、非密码复制与 OnSubmit；Disabled 移除焦点并停止输入。密码即使显示明文也禁止复制/剪切。ShowPasswordToggle 仅在 Password=true 时有效，只改变内部显示，不触发 OnChange。IME 组合阶段不触发 OnChange；提交后才产生值。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
