# 几何、布局与绘制值

本页按用途解释共有 API；精确字段类型、枚举常量与源码契约见[样式完整声明](/zh-cn/api/style)，坐标基础类型见[运行时声明](/zh-cn/api/runtime)。[布局完整示例](/zh-cn/guide/layout)和[样式完整示例](/zh-cn/guide/style)展示组合方式。

## 坐标与构造函数

| API | 参数与用途 | 零值、顺序与限制 |
| --- | --- | --- |
| Point{X,Y} / Size{Width,Height} / Rect{X,Y,Width,Height} | float32 逻辑位置、偏移、尺寸和矩形；图标中为 view-box 坐标。 | 不含后端对象。 |
| Px(float32) Length | 显式逻辑长度。 | 布局要求有限非负；零 Length 为 auto，不同于 Px(0)。 |
| Percent(float32) / Fill() | 百分比 0..100；Fill 等价 Percent(100)。 | 父轴不确定时按 auto。 |
| Metric(float32) MetricValue | 显式逻辑指标，可表达设置为 0。 | 与 Length 类型不同，用于 padding、边框等；约束依属性而定。 |
| TokenMetric(MetricToken) | 读取主题指标。 | 缺失/循环引用被拒绝。 |
| NoShrink() Option[float32] | Some(float32(0))。 | 与默认 shrink=1 不同。 |
| Padding(n) / Margin(n) | 四边统一逻辑数值。 | 返回 EdgeValues，包含显式零。 |
| PaddingXY(x,y) / MarginXY(x,y) | 水平与竖直边距。 | x 对应左右，y 对应上下。 |
| Edges(t,r,b,l) | 四边显式数值。 | 上、右、下、左；不支持负间距。 |
| UniformEdges(MetricValue) | 四边同一字面/Token 指标。 | EdgeValues 各字段零值为未设置，可单独配置。 |
| Round(n) / Corners(tl,tr,br,bl) | 统一/逐角逻辑圆角。 | 左上、右上、右下、左下；绘制时按几何钳制。 |
| UniformCorners(MetricValue) | 统一字面/Token 圆角。 | CornerValues 各字段零值为未设置。 |
| RGBA(r,g,b,a) | uint8 非预乘 RGBAColor。 | 0..255，a=0 全透明。 |
| ColorRGBA(r,g,b,a) / LiteralColor(RGBAColor) | 返回可设置样式的 ColorValue。 | 包括显式透明色；RGBAColor 与 ColorValue 不可混用。 |
| TokenColor(ColorToken) | 返回 Token 引用 ColorValue。 | 当前主题负责解析。 |
| Stroke(width,color) / NoBorder() | 实线 Border / 显式边宽 0。 | 普通 Style 后面的状态补丁仍可覆盖。 |

## Style 每个字段

| 字段 | 作用与默认 |
| --- | --- |
| Width/Height、MinWidth/MinHeight、MaxWidth/MaxHeight | Length 尺寸与限制，默认 auto；最小值优先于冲突最大值。 |
| Margin/Padding | 外/内边距，EdgeValues 四边指标；未设置时组件可给默认。 |
| Position | PositionFlow 默认；PositionAbsolute 退出 flex 流。 |
| Insets | Top/Right/Bottom/Left 的 Length 绝对位置约束。 |
| Grow | 非负增长系数，默认 0。 |
| Shrink | Option[float32] 非负收缩系数，未设置为 1。 |
| Basis | flex 基准 Length，默认 auto；收缩权重结合原始 basis。 |
| AlignSelf | Option[Align]，未设置则使用父级 Align。 |
| ZIndex | 同级绘制顺序整数，默认 0；同值按源码顺序。 |
| Overflow | OverflowVisible 默认；OverflowClip 矩形裁剪。 |
| Background | ColorValue 背景；未设置沿用组件样式。 |
| Border | 内侧绘制且不占布局空间；见下表。 |
| Radius | CornerValues 四角；不等同圆角后代裁剪。 |
| Shadow | 最多 4 个外阴影；nil 继承，非 nil 空切片清除。 |
| Opacity | Option[float32]，0..1，乘到子树；0 不命中。 |
| Visibility | Visible 默认；Hidden 保留布局，不绘制/命中。 |
| Text | 文本样式与继承颜色，见下文。 |
| Force | 最后应用的 StylePatch，仅绘制值，不能改变布局。 |

Align 为 Start/Center/End/Stretch；Justify 为 Start/Center/End/SpaceBetween；Direction 为 Vertical/Horizontal。这些为枚举类型，使用公开常量，不自行构造越界整数。

## Border、Shadow 与 TextStyle

Border.Width 为 MetricValue，Color 为 ColorValue，Pattern 为 BorderSolid（默认）或 BorderDashed，Sides 用 BorderTop/Right/Bottom/Left 位或组合；Sides=0 表示所有边，**不是无边框**。普通 Style 设置 Width/Color 时也会重置 Sides；StylePatch.Border 替换整个 Border。

Shadow 的 OffsetX/OffsetY、Blur/Spread 都是 MetricValue，Color 为 ColorValue；偏移可带方向，Blur/Spread 必须有限非负。阴影是有界几何近似，非 CSS 高斯模糊；不影响布局和命中。

TextStyle.Families 是有序 []FontFamily，Size/LineHeight 是逻辑字体大小/行高（0 取主题默认）；Weight 选最近注册字重，公开常量 WeightRegular=400、WeightMedium=500、WeightBold=700；SlantNormal/SlantItalic 选择字体斜体面，不承诺合成所有字体变体。Color 设置文字颜色，Align 为 TextStart/Center/End。只提供简单 LTR/CJK 文本，不支持 Arabic/Indic shaping、bidi/RTL、彩色 emoji、竖排。

## 覆盖补丁

StylePatch 的 Background、Border、Radius、Shadow、Opacity、Visibility、TextColor 全部是 Option，未设置沿用前一层，Some 显式覆盖（包括零值）。StateStyles 含 Default、Hover、Focus、Checked、Pressed、Disabled，每项都是 StylePatch；状态优先级见[样式教程](/zh-cn/guide/style)。
