# Slider：滑块

用 `Slider` 构建滑块。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/slider/main.go

## 参数与 API

```go
func Slider(props SliderProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Value` | `float32` | 应用拥有的当前值；行为、范围和零值见本页说明。 |
| `Min` | `float32` | 下界；与 Max 同为 0 时使用默认 0..100，否则按字面值；须有限且小于 Max。 |
| `Max` | `float32` | 上界；与 Min 同为 0 时默认 100，否则按字面值；须有限且大于 Min。 |
| `Step` | `float32` | 以 Min 为起点的步长；非正或非有限时默认 1，端点仍可到达。 |
| `Disabled` | `bool` | 默认 false；交互组件为 true 时移除焦点并取消交互。 |
| `OnChange` | `func(float32)` | 完整新值提案；写回业务变量后下一次构建生效。nil 行为见本页。 |

## 行为与边界

Min 和 Max 同为零表示 0..100；Step 非正或非有限时回退 1。其它边界按字面使用，非有限边界或 Min>=Max 使控件不可交互、不可聚焦。Value 超范围被钳制；NaN/-Inf 显示 Min，+Inf 显示 Max。提案以 Min 为起点对齐步长，端点仍可到达。支持轨道点击、捕获拖动、方向键、Home/End。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
