# 布局与响应式

这一节把导航和内容排在一起，并让它们随窗口宽度切换排列方式。

## 1. 运行示例

<<< ../../examples/responsive/main.go

拖窄窗口：宽度小于 600 时，左右排列变成上下排列。

## 2. 用 Box 排列控件

| 配置 | 效果 |
| --- | --- |
| `Direction: dxui.Horizontal` | 从左到右排列。 |
| `Direction: dxui.Vertical` | 从上到下排列，也是默认值。 |
| `Gap: 16` | 控件之间留 16 个逻辑单位。 |
| `Padding: dxui.Padding(24)` | 容器四周留 24 个逻辑单位。 |

`Gap` 放在 `BoxProps` 中，`Padding` 放在它的 `Style` 中。逻辑单位会随屏幕缩放，不等同于物理像素。

## 3. 分配宽度和高度

示例中的导航使用 `Width: dxui.Px(160)` 固定宽度，配合 `Shrink: dxui.NoShrink()` 防止被挤小；内容区使用 `Grow: 1` 占用排列方向上的剩余空间。

`RunResponsive` 提供当前窗口的 `ctx.Width` 和 `ctx.Height`。把它们设为根容器尺寸，再根据 `ctx.Width` 选择横排或竖排。

试着把切换条件中的 `600` 改成 `800`，观察变化。

## 内容太多时怎么办

内容超出容器不会自动滚动。使用 [Scroll](/zh-cn/components/scroll)，并为它设置明确的高度；大量固定高度的行使用 [VirtualList](/zh-cn/components/virtual-list)。

::: details 更多尺寸、定位和裁剪规则

| 写法 | 含义与约束 |
| --- | --- |
| 未设置 Width/Height | auto，由内容和父约束决定。 |
| Px(n) | 非负固定逻辑尺寸。 |
| Percent(50) | 父级该轴尺寸确定时取一半；无法确定时按 auto。 |
| Fill() | Percent(100)，也依赖父级确定尺寸。 |
| Grow: 1 | 按比例分配主轴剩余空间。 |
| Shrink 未设置 | 默认 1，空间不足时可缩小。 |
| Shrink: NoShrink() | 显式 0，不参与缩小；可能导致溢出。 |
| Basis | flex 基准尺寸；支持 auto、Px、Percent。 |
| MinWidth/MaxWidth 等 | 限制最终尺寸；最大值小于最小值时提升到最小值。 |

不要在父级高度由子内容决定时，期望子项 Fill 自动撑满窗口。需要填满窗口时用响应式上下文给根明确宽高。

### 滚动不是自动产生的

溢出内容默认不滚动。使用 Scroll 并给视口确定的高度；它在允许滚动的轴上对内容无界测量。完整例子见 [Scroll](/zh-cn/components/scroll)。需要几千或更多行时使用 [VirtualList](/zh-cn/components/virtual-list)，并理解它的固定高度与卸载规则。

### 绝对定位、裁剪与层叠

`Style.Position=PositionAbsolute` 配合 `Insets{Top, Right, Bottom, Left}`。绝对定位子项不参加 flex 尺寸分配和 Gap 计数。`ZIndex` 只改变同级绘制/命中顺序，不改变布局和 Tab 源顺序。

`OverflowClip` 是矩形裁剪。设置 Radius 不会自动获得圆角子树裁剪。边框在矩形内部绘制，但**不占据布局空间**，需要内容留白时设置 Padding。Shadow 也不改变布局或点击区域。

`Hidden` 保留占位但不绘制/命中；Opacity 为 0 的子树也不会命中。透明背景颜色本身不会让几何点击区域消失。

只支持单行 flex；不支持 wrap、order、baseline、负间距、自动 margin、百分比间距、space-around/evenly、sticky、完整 CSS 布局。全部字段见[布局与样式值](/zh-cn/api/values)。


:::

下一步：[调整样式与主题](./style)。
