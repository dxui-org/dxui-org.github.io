# 字体、图片与图标资源

精确签名见[资源声明](/zh-cn/api/resources)；ImageSource/ImageFit 等也在[组件类型](/zh-cn/api/components)。资源描述应在 root builder 外创建复用。

## 字体

| API | 参数/结果 | 生命周期与限制 |
| --- | --- | --- |
| FontBytes(family, data) Font | FontFamily 名称和 TTF/OTF/TTC 字节 | 立即复制字节；只要 Font/App 可达即保有数据。 |
| FontFile(family, path) Font | 精确文件路径 | Run 启动时打开，文件不自动打包；保持到文本资源销毁，不热加载。 |
| SystemFont(family) Font | 固定通用系统字体族 | 尝试当前 OS 固定路径表，找不到时报启动错误，不模糊搜索字体显示名。 |
| Font | Family、Weight、Slant、FaceIndex | 用构造函数提供隐藏源；FaceIndex 为 TTC 字面索引，普通字体用 0；注册到 AppOptions.Fonts。 |
| FontFamilyDefault | 内嵌 Go Regular | 覆盖基础拉丁字符，不包含 CJK。 |
| FontFamilySystemSans / Mono / CJK | system-ui / system-monospace / system-cjk | 固定平台候选，并非所有系统必有。 |

```go
// 配置片段：将字体随应用提供，路径以运行目录为基础。
font := dxui.FontFile("app-cjk", "assets/NotoSansCJK-Regular.ttc")
font.FaceIndex = 2 // 仅在这份字体集合的该索引确为所需字面时使用
app := dxui.NewApp(dxui.AppOptions{

	Fonts: []dxui.Font{font},

	DefaultFont: "app-cjk",

	DisableSystemFontFallback: true,
})
_ = app
```

回退考虑 TextStyle.Families、应用注册/default 字体、内嵌拉丁字体、懒加载系统 CJK 候选，最后替代符/缺字框。自动 CJK 回退默认启用，候选缺失不会让整个布局失败；显式 SystemFont 失败则是启动错误。要可重复发布，使用自己有权分发的字体文件或嵌入字节，并保留字体许可证；不能依赖开发机字体恰好存在。

## 完整字体示例

下面是无需外部字体文件的完整 FontBytes 程序，使用 dxui 当前依赖中的 Go Regular 字体字节。它演示字节注册和默认字体选择，仅显示拉丁字符；不把该字体当作 CJK 字体。

<<< ../../examples/fonts/main.go

这份示例需保留 `golang.org/x/image` 模块依赖（`go mod tidy` 会按 dxui 的依赖图解析），所用字体为 Go 字体项目的 BSD 许可资源。

## 图片

`ImageBytes([]byte)` 立即复制编码数据；`ImageFile(string)` 记录按需读取的路径；`ImageFromGo(image.Image)` 保留按契约不可修改的 Go 图像。ImageSource 是可比较句柄，不暴露像素读取方法。每次新构造可能产生新资源身份，应在 builder 外缓存并按业务需要替换。

只解码 PNG、JPEG、GIF 首帧；不读取在线 URL，不动画，不异步解码。Image.MaxPixels 默认 16×1024×1024，限制解码尺寸；Avatar 共享资源管线，但不暴露 MaxPixels 字段。OnLoad(Size) 和 OnError(error) 是通知，nil 不阻止加载。[无外部素材的 Image 完整示例](/zh-cn/components/image)与[Avatar 完整示例](/zh-cn/components/avatar)可直接运行。

## 官方与自定义图标

常见界面使用 `icon.Search()` 一类无参数构造器；全部 2,066 个名称/别名见[图标目录](/zh-cn/api/icons)。构造器返回 IconData，用 IconProps.Size/Color/StrokeWidth 配置绘制。普通应用无需导入 catalog；catalog.Icons 全量目录会让整个图标库进入可达集合。

自定义 IconData 用 ViewBox 和 []PathCommand 描述填充路径。PathMove/PathLine 使用 Points[0]，PathQuad 使用前两个点，PathCubic 使用三个点，PathClose 关闭路径。不是 SVG 语法，也没有运行时 SVG 解析。

<<< ../../examples/custom-icon/main.go

IconData 是公开别名，因此 `IsPacked() bool`、`Identity() uint64`、`Paths(maxCommands int)` 也可调用。IsPacked 区分官方不可变编码与自定义 Commands；Identity 是内容哈希，不保证无碰撞，不应用作业务主键；Paths 校验并解码，maxCommands 必须为正且足够，失败返回 error。

Paths 的返回元素来自内部定义（含 Mode 与 Commands），外部可以通过类型推断访问，但不应导入 internal 来命名它。Mode 的内部值为描边 0、填充 1；这是别名带出的低层表示，组件应用应直接使用 IconData。不要修改官方图标内部描述或解码结果来期望改变资源。自定义 Commands 在构建 View 时复制。
