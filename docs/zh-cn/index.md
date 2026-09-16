---
layout: home
hero:
  text: 声明式构建，按需渲染。
  tagline: 面向跨平台桌面应用的 Go GUI 框架，同时支持 Windows、MacOS、Linux。融合声明式视图与保留式更新，以事件驱动运行；支持无 CGO 构建、多窗口管理与类型安全的组件、布局和主题。
  image:
    src: /brand/logo.png
    alt: DXUI 标志
    width: 1983
    height: 793
  actions:
    - theme: brand
      text: 运行第一个窗口
      link: /zh-cn/guide/getting-started
    - theme: alt
      text: 查看组件示例
      link: /zh-cn/components/
features:
  - title: 声明式接口，保留式内核
    details: 以不可变 View 描述界面，由框架协调新旧组件树。稳定的 Key 保留焦点、编辑与交互状态，业务数据始终由应用掌握。
    link: /zh-cn/guide/state
    linkText: 理解视图与状态
  - title: 事件驱动的按需更新
    details: 多个窗口共用事件循环，分别维护视图与交互状态，按需更新各自界面。空闲时不持续重绘。区分布局与绘制变化，颜色等视觉更新无需重新测量布局。
    link: /zh-cn/guide/troubleshooting
    linkText: 了解运行与性能行为
  - title: 无需 CGO 的构建方式
    details: 使用 CGO_ENABLED=0 构建框架与应用，沿用 Go 的模块管理和编译流程，为不同桌面平台构建应用产物。
    link: /zh-cn/guide/getting-started
    linkText: 查看构建与平台要求
  - title: 类型安全的布局与主题
    details: 用具体 Go 类型配置布局、样式和交互状态。Primitive、Semantic、Component 三层 Token 支持统一定制，主题校验后原子替换。
    link: /zh-cn/guide/style
    linkText: 探索样式与主题
  - title: 面向长列表的虚拟化
    details: 固定行高 VirtualList 只挂载可见区域与预留行，支持稳定业务键和滚动锚定，让大数据列表的视图规模保持可控。
    link: /zh-cn/components/virtual-list
    linkText: 使用虚拟列表
  - title: 按需链接的矢量图标
    details: 提供 2,066 个官方图标名称与别名，支持主题着色和描边宽度配置。直接引用所需构造器，未使用的图标可由 Go 链接器移除。
    link: /zh-cn/components/icon
    linkText: 使用图标资源
---

<div class="dxui-home">

<section class="home-start" aria-labelledby="first-window">
<div class="home-start-copy">
<p class="home-eyebrow">从这里开始</p>
<h2 id="first-window">一个 main.go，<br>一个桌面窗口。</h2>
<p>组件组合、应用生命周期与事件回调，统一在类型明确的 Go API 中。下面的完整程序展示了一个窗口从构建到关闭的基本流程。</p>
<ol class="home-steps">
<li><strong>准备环境</strong><span>安装 Go 1.25+，创建模块并添加 dxui 依赖。</span></li>
<li><strong>描述界面</strong><span>在 Run 的构建函数里组合文字、容器和按钮。</span></li>
<li><strong>运行与交互</strong><span>执行 go run .，点击 Close 结束应用。</span></li>
</ol>
<a class="home-text-link" href="/zh-cn/guide/getting-started">查看安装步骤与运行说明 <span aria-hidden="true">→</span></a>
</div>
<div class="home-code">
<div class="home-code-title"><span>main.go</span><span>完整示例</span></div>

<<< ../examples/hello/main.go

</div>
</section>

<section class="home-explore" aria-labelledby="explore-docs">
<p class="home-eyebrow">按你的进度阅读</p>
<h2 id="explore-docs">从学会使用，到随时查阅。</h2>
<div class="home-paths">
<a href="/zh-cn/guide/" class="home-path">
<span class="home-path-number">01 / 学习</span>
<h3>跟着教程构建应用</h3>
<p>窗口 → 状态与事件 → 布局与样式 → 表单 → 后台任务。每一步都解释代码为什么这样写。</p>
<span class="home-path-action">开始指南 <span aria-hidden="true">↗</span></span>
</a>
<a href="/zh-cn/components/" class="home-path">
<span class="home-path-number">02 / 实践</span>
<h3>找到需要的组件</h3>
<p>23 种组件，附可独立运行的程序、完整属性与交互规则。先复制运行，再按需求调整。</p>
<span class="home-path-action">浏览完整组件示例 <span aria-hidden="true">↗</span></span>
</a>
<a href="/zh-cn/api/" class="home-path">
<span class="home-path-number">03 / 查阅</span>
<h3>查清 API 的实际行为</h3>
<p>查找用途、参数、默认值与限制。从应用生命周期到主题、字体、图片和全部官方图标。</p>
<span class="home-path-action">打开公开 API 索引 <span aria-hidden="true">↗</span></span>
</a>
</div>
</section>

<section class="home-topics" aria-labelledby="common-tasks">
<div><p class="home-eyebrow">常用入口</p><h2 id="common-tasks">你想实现什么？</h2></div>
<div class="home-topic-links">
<a href="/zh-cn/guide/forms">构建输入表单 <span aria-hidden="true">→</span></a>
<a href="/zh-cn/components/virtual-list">展示大量列表数据 <span aria-hidden="true">→</span></a>
<a href="/zh-cn/guide/style">定制样式与明暗主题 <span aria-hidden="true">→</span></a>
<a href="/zh-cn/api/icons">查找官方图标 <span aria-hidden="true">→</span></a>
<a href="/zh-cn/guide/async">接入后台任务 <span aria-hidden="true">→</span></a>
<a href="/zh-cn/guide/multi-window">管理多个窗口 <span aria-hidden="true">→</span></a>
<a href="/zh-cn/guide/troubleshooting">排查运行与交互问题 <span aria-hidden="true">→</span></a>
</div>
</section>

<p class="home-version">© 2026 DXUI · MIT</p>
</div>
