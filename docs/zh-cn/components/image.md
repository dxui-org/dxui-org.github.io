# Image：栅格图片

用 `Image` 构建栅格图片。本页包含可运行示例、完整属性及行为限制。

## 完整示例

先完成[环境准备](/zh-cn/guide/getting-started)。下面是完整 `main.go`，执行 `go run .`。

<<< ../../examples/image/main.go

## 参数与 API

```go
func Image(props ImageProps) View
```

| 字段 | 类型 | 用途、默认与约束 |
| --- | --- | --- |
| `Key` | `string` | 同父级稳定唯一身份；默认空。动态列表使用业务 ID，勿用可变索引。 |
| `Style` | `Style` | 局部布局、绘制和文本样式；零值使用固有尺寸与主题默认。 |
| `Token` | `ComponentToken` | 组件主题条目；空值选择该组件默认。 |
| `States` | `StateStyles` | 真实交互状态的绘制补丁，不能伪造状态或改变布局。 |
| `Pointer` | `PointerBehavior` | 默认 PointerAuto；PointerNone 排除整个子树的指针参与，不等同 Disabled。 |
| `Source` | `ImageSource` | 稳定 ImageSource；复用资源句柄，避免每次构建重新分配。 |
| `Fit` | `ImageFit` | 默认 ImageContain；支持 Cover、Fill、None。 |
| `Alignment` | `Point` | X/Y 均为 0..1；默认 (0,0) 左上，(.5,.5) 居中。 |
| `MaxPixels` | `int64` | 解码像素预算；0 选引擎默认 16,777,216，负值无效。 |
| `OnLoad` | `func(Size)` | 加载成功通知，参数为解码 Size；nil 仍加载。 |
| `OnError` | `func(error)` | 资源错误通知；应用可切换占位视图，nil 仍尝试加载。 |

## 行为与边界

Source 由 ImageBytes、ImageFile 或 ImageFromGo 构造；放在 builder 外复用。Fit 默认 Contain（保比例完整显示），Cover 填满并裁剪，Fill 拉伸，None 原始尺寸。Alignment 每轴 0..1，默认左上。MaxPixels 限制解码像素数；只有 PNG/JPEG/单帧 GIF，无网络 URL 或异步解码。nil 加载回调不阻止解码。Go 图像构造后不得修改。

所有组件共享[身份、样式与指针规则](/zh-cn/api/common)。字段引用的枚举、结构体及源码注释见[完整声明参考](/zh-cn/api/components)。[示例运行方式](/zh-cn/examples/)包含构建验证方法。
