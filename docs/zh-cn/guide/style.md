# 样式与主题

这一节调整控件外观，并用按钮切换明暗主题。

## 1. 运行示例

<<< ../../examples/theme/main.go

点击 **Switch theme** 切换主题。把鼠标移到按钮上，按钮会稍微变淡；按 Tab 可以看到键盘焦点效果。

## 2. 分清三个入口

| 想改什么 | 在哪里设置 |
| --- | --- |
| 控件的尺寸、间距、背景 | 控件的 `Style`。 |
| 鼠标悬停时的外观 | `States.Hover`，示例中设置了透明度。 |
| 整个应用的配色 | 调用 `app.SetTheme`。 |

普通外观从 `Style` 开始设置。例如，下面是一个样式片段：

```go
dxui.Style{
    Width: dxui.Px(240),
    Padding: dxui.Padding(16),
    Background: dxui.ColorRGBA(240, 245, 255, 255),
}
```

## 3. 切换主题

示例从 `LightTheme()` 或 `DarkTheme()` 取得完整主题，修改强调色，再传给 `app.SetTheme(next)`。主题中的颜色名称称为 Token，控件通过它们使用统一配色。

试着把示例中的 `Blue600` 改成 `Red600`。修改主题后要再次调用 `SetTheme`，只修改 Go map 不会更新界面；有多个窗口时，主题会传播到所有存活窗口。

::: details 进阶：覆盖顺序、清除默认样式与 Token

### 实际覆盖顺序

```text
Primitive → Semantic → Component Base/Default
→ 局部 Style → 局部 States.Default
→ Hover → Focus → Checked → Pressed → Disabled
→ Style.Force
```

同一个活动状态内先应用组件主题补丁，再应用局部补丁。Disabled 抑制 Hover/Focus/Pressed，但受控 Checked 层仍保留在其下。状态补丁只支持绘制属性，不能让按钮 hover 时改变尺寸。

### 显式零值

Go 的零值经常表示“未设置”，并非“强制清空”。

| 目的 | 写法 |
| --- | --- |
| 清除默认内边距 | `Padding: dxui.Padding(0)` |
| 强制透明度为 0 | `Opacity: dxui.Some(float32(0))` |
| 禁止缩小 | `Shrink: dxui.NoShrink()` |
| 清除普通边框 | `Border: dxui.NoBorder()` |
| 清除主题阴影 | `Shadow: []dxui.Shadow{}`，nil 不清除 |
| 连状态阴影一起清除 | `Force: dxui.StylePatch{Shadow: dxui.Some([]dxui.Shadow{})}` |

清除焦点环会使键盘定位不明显；大多数场景保留默认 Focus 样式，只改普通外观即可。

### 字面值、Token 与颜色空间

`RGBA` 返回原始 RGBAColor，适合 AppOptions.Background 和 Primitive 颜色表。`ColorRGBA` 返回显式 ColorValue，适合 Style.Background。已有 RGBAColor 用 LiteralColor 包装；TokenColor 指向主题颜色。Metric 是显式逻辑数值，TokenMetric 从主题读取指标。

`Color.Primitive.Blue600` 是调色板 token，`Color.Semantic.Accent` 表示强调用途。主题允许把语义色重新映射到另一原始色。旧的平铺常量也在[Token 参考](/zh-cn/api/tokens)中，旧 White/Black/Blue/Gray 与新调色板条目并非同一组值。

主题先从 `LightTheme()` 或 `DarkTheme()` 取完整副本，再改映射；不要创建一个只含 Accent 的 Theme 期待自动补全其余项。SetTheme 校验缺失引用、循环和非法值，失败时整份拒绝。调用成功会复制 map/slice；之后修改应用原始 map 不会隐式更新界面，需要再次 SetTheme。相同或未使用 Token 的变化不必产生新帧。


:::

下一步：[把控件组合成表单](./forms)。完整配置见[主题 API](/zh-cn/api/theme)。
