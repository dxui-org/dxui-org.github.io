# 主题与 Token API

`LightTheme() Theme`、`DarkTheme() Theme` 返回独立、可修改的完整主题。`App.SetTheme(Theme) error` 验证并复制主题，传播到所有存活窗口；运行中从 UI 回调或 Update 调用。新建子窗口采用 App 当前主题，Window 没有独立 SetTheme 方法。[完整明暗切换程序](/zh-cn/guide/style)。

| 类型/字段 | 用途 | 约束与示例 |
| --- | --- | --- |
| Theme.Primitive | PrimitiveTokens | Colors: map[ColorToken]RGBAColor；Metrics: map[MetricToken]float32。基础字面值。 |
| Theme.Semantic | SemanticTokens | Colors: map[ColorToken]ColorValue；Metrics: map[MetricToken]MetricValue。将用途映射为字面值或其它 Token。 |
| Theme.Components | map[ComponentToken]ComponentTheme | 每个组件默认外观；非空 token 名。 |
| ComponentTheme.Base | StylePatch | 组件基础绘制值。 |
| ComponentTheme.States | StateStyles | 组件 Default/Hover/Focus/Checked/Pressed/Disabled 补丁。 |
| ColorToken / MetricToken / ComponentToken | string 定义的不同类型 | 用作对应 map 的键，不混用，不传未定义引用。 |
| Color.Primitive.* | 可发现的调色板键 | 26 色族，每族 50、100..900、950，另有 White/Black；完整字段在 Token 声明页。 |
| Color.Semantic.* | Surface、SurfaceHigh、Text、Accent、AccentHover、OnAccent、FocusRing、Danger、Success、Info、Warn、Border、Shadow | 按语义定制，避免每个按钮硬编码调色板。 |

```go
// 片段：app 已创建；运行中在 UI 回调内执行。
theme := dxui.LightTheme()
theme.Semantic.Colors[dxui.Color.Semantic.Accent] =
	dxui.TokenColor(dxui.Color.Primitive.Violet600)
theme.Semantic.Metrics[dxui.MetricSemanticControlHeight] = dxui.Metric(40)
if err := app.SetTheme(theme); err != nil {
	log.Print(err)
}
```

组件尺寸 Token 在 Semantic.Metrics 表中；先核对内置主题对应 map，再修改，避免同名键在错误层被遮盖。局部显式尺寸、padding、字体指标仍优先于主题默认。

## Token 查找与使用

全部常量（包含旧平铺拼写）见[Token 完整声明](/zh-cn/api/tokens)。`ColorPrimitive*` 是原始颜色键，`ColorSemantic*` 是用途颜色键；`MetricPrimitive*` 是基础刻度，`MetricSemantic*` 是通用间距、圆角、字号等；`MetricComponent*` 是某组件的高度、间距、轨道、指示器尺寸；`Component*` 选择组件外观条目。

比如 ProgressBar 的 `ColorSemanticProgressTrack/Fill` 分别控制轨道/填充；Tabs 的 `ColorSemanticTabsIndicator` 控制选中指示；Menu 的 Hover/Active/Selected/Pressed/Disabled 颜色分开。每个常量页保留真实字符串值，索引可搜索名字。它们是键名，不能直接当 RGBA 或 float32 使用，要包在 TokenColor/TokenMetric 中。

Button 的 Variant/Tone 是常用外观入口，显式 Token 替换 Variant/Tone 配方，Size 仍独立生效。定制共享组件外观时用 `ComponentTheme`；仅一个按钮的例外用 Style/States。[按钮完整例子](/zh-cn/components/button)。

主题不接受缺失引用、循环依赖或非法数值。完全零 Theme 选内置亮色；创建一个部分非零 Theme 不等于“给默认主题打补丁”。SetTheme 相等或改变未使用值可不重绘，不能将其当作强制刷新接口。
