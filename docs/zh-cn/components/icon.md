# Icon：矢量图标

用 `Icon` 构建矢量图标。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/icon/main.go

## 参数与 API

```go
func Icon(props IconProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Data` | `IconData` | 有效 IconData；可用官方 icon 构造器或自定义填充路径。 |
| `Size` | `float32` | 尺寸；Button 是枚举，其它组件为逻辑单位，0 取主题默认。 |
| `Color` | `ColorValue` | 图标显式颜色或 TokenColor；未设置时继承主题色。 |
| `StrokeWidth` | `IconStrokeWidth` | 有限非负 view-box 描边宽，0 为 2；填充路径忽略。 |

## 行为与边界

Size 是逻辑单位，0 使用主题默认 20。StrokeWidth 为 view-box 单位，0 表示默认 2，必须有限且非负；对自定义填充图标忽略。图标无语义动作，需要组合 Button。icon.Name() 按需链接；icon/catalog 会引入完整目录。IconData 不是 SVG 解析器。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
