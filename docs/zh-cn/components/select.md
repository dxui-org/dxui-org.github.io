# Select：下拉选择

用 `Select` 构建下拉选择。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/select/main.go

## 参数与 API

```go
func Select(props SelectProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Value` | `string` | 应用拥有的当前值；行为、范围和零值见本页说明。 |
| `Options` | `[]SelectOption` | 复制的 SelectOption 列表；值非空唯一，上限 4096。 |
| `Placeholder` | `string` | 空 Value 时显示；默认空字符串。 |
| `Disabled` | `bool` | 默认 false；交互组件为 true 时移除焦点并取消交互。 |
| `OnChange` | `func(string)` | 完整新值提案；写回业务变量后下一次构建生效。nil 行为见本页。 |

## 行为与边界

Options 的 Value 必须非空且唯一，作为重排身份。支持空选项列表和未匹配 Value。nil OnChange 仍可打开、浏览、关闭，只不接纳选择。支持 Enter/Space、方向键、Home/End、Escape、Tab；禁用选项跳过。弹层逃逸祖先 Scroll 裁剪，并在窗口边缘翻转/钳制。不是原生系统选择框，不虚拟化，最多 4096 项。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
