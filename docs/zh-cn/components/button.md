# Button：按钮与 TextButton

用 `Button` 构建按钮与 TextButton。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/button/main.go

## 参数与 API

```go
func Button(props ButtonProps, child View) View
```

```go
func TextButton(props ButtonProps, label string) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Variant` | `ButtonVariant` | 默认 ButtonFilled；外观配方，详见本页。 |
| `Tone` | `ButtonTone` | 默认 ButtonPrimary；ButtonDefault 是其别名。 |
| `Size` | `ButtonSize` | 尺寸；Button 是枚举，其它组件为逻辑单位，0 取主题默认。 |
| `Disabled` | `bool` | 默认 false；交互组件为 true 时移除焦点并取消交互。 |
| `OnPress` | `func()` | 激活后在 UI 线程执行；nil 不执行动作。 |

## 行为与边界

Button 接受一个任意子 View；TextButton 自动居中显示文字。默认 Filled / Primary / Normal。Variant 支持 Filled、Soft、Outline、Dashed、Ghost、Link；Tone 支持 Primary、Secondary、Success、Info、Warn、Danger；Size 支持 Small、Normal、Large。显式 Token 替换外观配方，局部 Style 保持覆盖权。指针或 Enter/Space 匹配释放时激活一次；nil OnPress 仍保留焦点与交互外观。Link 不执行导航。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
