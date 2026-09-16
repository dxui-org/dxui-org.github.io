# ToggleSwitch：开关

用 `ToggleSwitch` 构建开关。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/toggle-switch/main.go

## 参数与 API

```go
func ToggleSwitch(props ToggleSwitchProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Checked` | `bool` | 当前受控布尔值，默认 false；在 OnChange 中接纳提案。 |
| `Disabled` | `bool` | 默认 false；交互组件为 true 时移除焦点并取消交互。 |
| `OnChange` | `func(bool)` | 完整新值提案；写回业务变量后下一次构建生效。nil 行为见本页。 |

## 行为与边界

Checked / OnChange 与 Checkbox 相同，通过指针或 Space 提案。默认尺寸由主题指标决定（40×24 逻辑单位）。没有内置文字标签，可与 Text/Label 组合；nil 回调仍有交互外观。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
