# Avatar：头像

用 `Avatar` 构建头像。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/avatar/main.go

## 参数与 API

```go
func Avatar(props AvatarProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Source` | `ImageSource` | 稳定 ImageSource；复用资源句柄，避免每次构建重新分配。 |
| `Shape` | `AvatarShape` | 默认 AvatarCircle；AvatarSquare 为方形。 |
| `Size` | `float32` | 尺寸；Button 是枚举，其它组件为逻辑单位，0 取主题默认。 |
| `OnLoad` | `func(Size)` | 加载成功通知，参数为解码 Size；nil 仍加载。 |
| `OnError` | `func(error)` | 资源错误通知；应用可切换占位视图，nil 仍尝试加载。 |

## 行为与边界

共享 ImageSource 和 OnLoad/OnError，固定居中 Cover。Shape 默认 AvatarCircle，可选 AvatarSquare。Size=0 用主题默认 40，正数为逻辑尺寸；显式 Style 宽高优先。没有 URL、文字缩写、自动降级或状态徽点；需要降级时由应用根据 OnError 改渲染内容。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
