# Radio：单选按钮

用 `Radio` 构建单选按钮。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/radio/main.go

## 参数与 API

```go
func Radio(props RadioProps, label View) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Selected` | `bool` | 当前是否选中，默认 false；与共享业务值比较得到。 |
| `Disabled` | `bool` | 默认 false；交互组件为 true 时移除焦点并取消交互。 |
| `OnSelect` | `func()` | 未选中 Radio 激活时调用，无参数；由应用更新共享选择值。 |

## 行为与边界

Selected 是权威值。启用且未选中时，指针或 Space 调用 OnSelect；已选中不会反选，也不回调。互斥由应用共享业务变量实现，没有 RadioGroup。每个 Radio 保留自己的焦点停靠点。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
