# Resources, performance, and troubleshooting

Check terminal errors first, then find the matching symptom below.

## The window does not open

Check that a desktop session is available and inspect the error returned by app.Run. If renderer creation failed, add this configuration to AppOptions and retry:

```go
dxui.AppOptions{
    Renderer: dxui.RendererSoftware,
}
```

Software rendering still needs a window system; it cannot open a window on a server without a display environment.

## The UI behaves unexpectedly

| Symptom | Check first |
| --- | --- |
| Typed text does not stay | OnChange writes the value back and the variable is outside Run. |
| Background changes do not appear | Submit assignments through App.Update or the target Window.Update. |
| Another window stays stale | Call Update or Invalidate for that window. |
| Controls do not fill the window | Use RunResponsive to set root width/height. |
| Overflow does not scroll | Use Scroll with a definite viewport height. |
| Focus follows the wrong reordered item | Use stable business Keys. |
| CJK characters appear as boxes | Check font glyph coverage; see [fonts](/api/resources-guide). |
| Hover changes a color | States override normal Style. |
| Child shortcuts do not work | Configure that child's WindowOptions.Shortcuts. |

## Slow rendering or increasing memory

Start with these three checks:

1. **Expensive work in callbacks**: Move it to a [background task](./async).
2. **Repeated image creation**: Reuse ImageSource outside builders; see [Image](/components/image).
3. **Too many mounted list items**: Use [VirtualList](/components/virtual-list) for fixed-height rows.

DXUI updates on demand; periodic Update calls are not needed to keep it refreshing.

::: details Performance metrics and capability limits

The event loop wakes for events and deadlines, not fixed-rate idle frames. Submit updates for real result/progress changes, not to keep refreshing. ProgressBar has no internal animation timer. Each window has independent resource budgets, so total capacity increases with more windows. Cache budgets are not hard working-set caps; visible images still consume resources.

Enable `AppOptions.Diagnostics: true` for comparable snapshots. Distinguish Build/Layout/Paint from FrameCount; frame counts need not grow while idle. Diagnostics() is concurrent-safe, but do not use polling to drive production UI refresh. See [diagnostic fields](/api/application#cachebudgets-and-runtimediagnostics).

VirtualList builds only visible and overscan rows, but a new Count/Version snapshot still needs O(N) key validation. It does not support variable row heights; unmounting offscreen rows discards runtime editor/focus state.

### Current capability limits

No full CSS Flexbox, rich text, grapheme-cluster editing, bidi/RTL/complex-script shaping, platform accessibility bridge, animated images, SVG parsing, RadioGroup, tri-state checkbox, TabPanel, or modal focus trap. Compose existing public types where appropriate and assess framework extensions when requirements exceed them.

Automated checks compare source APIs, compile complete programs, and build VitePress. Native input, real IME, DPI changes, visuals, and platform distribution still require manual testing. Compiling is not the same as running on every platform.

:::

Continue with [component examples](/components/) or the [public API](/api/).
