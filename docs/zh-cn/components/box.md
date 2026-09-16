# Box：容器与布局

用 `Box` 构建容器与布局。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/box/main.go

## 参数与 API

```go
func Box(props BoxProps, children ...View) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Direction` | `Direction` | 默认 Vertical；Horizontal 从左到右排列。 |
| `Gap` | `float32` | 相邻子项的非负固定逻辑间距，默认 0。 |
| `Justify` | `Justify` | 主轴 Start/Center/End/SpaceBetween，默认 Start。 |
| `Align` | `Align` | 交叉轴 Start/Center/End/Stretch，默认 Start。 |

## 行为与边界

默认 Vertical，单行排列；Gap 为逻辑单位。Justify 控制主轴，Align 控制交叉轴。children 可为空，但不能包含零 View。不支持换行、order、反向排列或完整 CSS Flexbox。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
