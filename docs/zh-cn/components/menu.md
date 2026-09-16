# Menu：菜单动作列表

用 `Menu` 构建菜单动作列表。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/menu/main.go

## 参数与 API

```go
func Menu(props MenuProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Orientation` | `MenuOrientation` | 默认水平（ButtonGroup）或垂直（Menu）；详见本页。 |
| `Value` | `string` | 应用拥有的当前值；行为、范围和零值见本页说明。 |
| `Items` | `[]MenuItem` | 复制的 TabItem/MenuItem 列表；值非空唯一，允许空列表。 |
| `Disabled` | `bool` | 默认 false；交互组件为 true 时移除焦点并取消交互。 |
| `OnAction` | `func(string)` | 每次执行启用菜单项时接收 Value；允许重复执行同一项。 |

## 行为与边界

普通布局内的动作列表，默认 MenuVertical，MenuHorizontal 改为水平。Value 是可选受控选择外观；OnAction 对已选中项仍会再次触发。Item.Value 非空唯一。一个菜单贡献一个 Tab 停靠点；对应方向键/Home/End 移动内部当前项，不循环，Enter/Space 执行。无内置弹层、子菜单、分隔线、路由或快捷键字段。长菜单组合 Scroll。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
