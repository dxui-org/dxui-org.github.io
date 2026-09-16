---
layout: home
hero:
  text: Declarative views. Render on demand.
  tagline: A Go GUI framework for Windows, macOS, and Linux. Declarative views meet retained updates in an event-driven runtime, with CGO-free builds, multiple windows, and typed components, layouts, and themes.
  image:
    src: /brand/logo.png
    alt: DXUI logo
    width: 1983
    height: 793
  actions:
    - theme: brand
      text: Run your first window
      link: /guide/getting-started
    - theme: alt
      text: Explore components
      link: /components/
features:
  - title: Declarative API, retained core
    details: Describe the UI with immutable Views and let the framework reconcile the tree. Stable keys retain focus and editing state while your application owns its data.
    link: /guide/state
    linkText: Understand views and state
  - title: Event-driven updates
    details: Windows share an event loop while keeping independent UI and interaction state. Idle windows do not repaint continuously, and paint-only changes avoid unnecessary layout measurement.
    link: /guide/troubleshooting
    linkText: Explore runtime behavior
  - title: Build without CGO
    details: Build the framework and applications with CGO_ENABLED=0, using familiar Go modules and compilation workflows for each desktop platform.
    link: /guide/getting-started
    linkText: Run and build
  - title: Typed layouts and themes
    details: Configure layout, appearance, and interaction states with concrete Go types. Primitive, Semantic, and Component tokens provide consistent customization with validated theme updates.
    link: /guide/style
    linkText: Explore styles and themes
  - title: Virtualized long lists
    details: Fixed-height VirtualList mounts only visible and overscan rows. Stable keys and scroll anchoring keep large datasets manageable.
    link: /components/virtual-list
    linkText: Use virtual lists
  - title: Icons linked on demand
    details: Choose from 2,066 official icon names and aliases with theme colors and configurable stroke widths. Reference only the constructors you need so the Go linker can discard unused icons.
    link: /components/icon
    linkText: Use icon resources
---

<div class="dxui-home">

<section class="home-start" aria-labelledby="first-window">
<div class="home-start-copy">
<p class="home-eyebrow">Start here</p>
<h2 id="first-window">One main.go.<br>One desktop window.</h2>
<p>Compose components and handle lifecycle events through a typed Go API. This complete program takes a window from creation to shutdown.</p>
<ol class="home-steps">
<li><strong>Set up</strong><span>Install Go 1.25+, create a module, and add DXUI.</span></li>
<li><strong>Describe the UI</strong><span>Compose text, containers, and buttons in the Run builder.</span></li>
<li><strong>Run and interact</strong><span>Run go run . and click Close to exit.</span></li>
</ol>
<a class="home-text-link" href="/guide/getting-started">View setup and running instructions <span aria-hidden="true">→</span></a>
</div>
<div class="home-code">
<div class="home-code-title"><span>main.go</span><span>Complete example</span></div>

<<< ./examples/hello/main.go

</div>
</section>

<section class="home-explore" aria-labelledby="explore-docs">
<p class="home-eyebrow">Learn at your pace</p>
<h2 id="explore-docs">From getting started to looking it up.</h2>
<div class="home-paths">
<a href="/guide/" class="home-path">
<span class="home-path-number">01 / Learn</span>
<h3>Build along with the tutorials</h3>
<p>Window → state and events → layout and styles → forms → background tasks. Each step explains how the code works.</p>
<span class="home-path-action">Start the guide <span aria-hidden="true">↗</span></span>
</a>
<a href="/components/" class="home-path">
<span class="home-path-number">02 / Practice</span>
<h3>Find the right component</h3>
<p>23 components with standalone programs, complete props, and interaction rules. Run an example, then adapt it.</p>
<span class="home-path-action">Browse component examples <span aria-hidden="true">↗</span></span>
</a>
<a href="/api/" class="home-path">
<span class="home-path-number">03 / Reference</span>
<h3>Understand each API</h3>
<p>Look up parameters, defaults, and limitations for lifecycle, themes, fonts, images, and every official icon.</p>
<span class="home-path-action">Open the API index <span aria-hidden="true">↗</span></span>
</a>
</div>
</section>

<section class="home-topics" aria-labelledby="common-tasks">
<div><p class="home-eyebrow">Common tasks</p><h2 id="common-tasks">What are you building?</h2></div>
<div class="home-topic-links">
<a href="/guide/forms">Build an input form <span aria-hidden="true">→</span></a>
<a href="/components/virtual-list">Display large datasets <span aria-hidden="true">→</span></a>
<a href="/guide/style">Customize styles and themes <span aria-hidden="true">→</span></a>
<a href="/api/icons">Find official icons <span aria-hidden="true">→</span></a>
<a href="/guide/async">Run background tasks <span aria-hidden="true">→</span></a>
<a href="/guide/multi-window">Manage multiple windows <span aria-hidden="true">→</span></a>
<a href="/guide/troubleshooting">Troubleshoot runtime behavior <span aria-hidden="true">→</span></a>
</div>
</section>

<p class="home-version">© 2026 DXUI · MIT</p>
</div>
