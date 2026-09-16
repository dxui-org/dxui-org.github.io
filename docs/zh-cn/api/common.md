# 组件共同契约

23 种组件的 Props 直接声明以下五个字段，不存在需要嵌套的公共 Props 包装器。

| 字段 | 用途、默认与限制 | 示例 |
| --- | --- | --- |
| Key string | 默认空，以同父级的类型/位置匹配；非空时用稳定且同级唯一业务键。跨父级不保留身份。 | [状态与身份](/zh-cn/guide/state) |
| Style Style | 零值选默认；局部尺寸、间距和绘制属性。 | [样式](/zh-cn/guide/style) |
| Token ComponentToken | 默认组件主题项；显式值选另一个已定义主题项。 | [主题](/zh-cn/api/theme) |
| States StateStyles | 给真实状态提供绘制补丁；不产生状态，不修改布局。 | [状态覆盖](/zh-cn/guide/style#实际覆盖顺序) |
| Pointer PointerBehavior | PointerAuto 默认；PointerNone 排除整个子树的指针参与。不是通用 Disabled。 | [布局命中](/zh-cn/guide/layout) |

所有构造函数返回不可变 View，不返回可变控件对象。`View.WithKey(string)` 与 `View.WithStyle(Style)` 返回独立描述；后者替换完整 Style。零 View 作为根或子项是错误，空 Box 则有效。

事件在 UI 线程执行；nil 回调允许存在，其交互含义由组件决定，不能直接推导为 Disabled。悬停、焦点、按压、编辑组合、弹层活动和非受控滚动由运行时拥有。业务值如 Value/Checked/Open 由应用拥有。[具体矩阵与示例](/zh-cn/guide/state)。

结构体、枚举和函数的完整声明见[组件 API](/zh-cn/api/components)；共有 Style 和主题结构见[样式声明](/zh-cn/api/style)。
