import { defineConfig, type DefaultTheme } from 'vitepress'

const components = [
  ['Box', 'box'], ['Text / Label', 'text'], ['Button / TextButton', 'button'],
  ['ButtonGroup', 'button-group'], ['InputGroup', 'input-group'],
  ['Input', 'input'], ['Textarea', 'textarea'], ['Checkbox', 'checkbox'],
  ['Radio', 'radio'], ['ToggleSwitch', 'toggle-switch'], ['Slider', 'slider'],
  ['Select', 'select'], ['Tabs', 'tabs'], ['Menu', 'menu'], ['Scroll', 'scroll'],
  ['VirtualList', 'virtual-list'], ['Popover', 'popover'], ['Tooltip', 'tooltip'],
  ['Icon', 'icon'], ['Image', 'image'], ['Avatar', 'avatar'], ['Badge', 'badge'],
  ['ProgressBar', 'progress-bar'],
]

function localeTheme(zh: boolean): DefaultTheme.Config {
  const prefix = zh ? '/zh-cn' : ''
  const label = (en: string, cn: string) => zh ? cn : en
  const item = (en: string, cn: string, path: string) => ({ text: label(en, cn), link: prefix + path })
  return {
    nav: [
      item('Guide', '指南', '/guide/'), item('Components', '组件', '/components/'),
      item('Other APIs', '其它 API', '/api/application'), item('API Index', '完整索引', '/api/'),
    ],
    sidebar: [
      { text: label('Guide', '指南'), collapsed: false, items: [
        item('Learning path', '学习路线', '/guide/'),
        item('1. First window', '1. 第一个窗口', '/guide/getting-started'),
        item('2. State and events', '2. 状态与事件', '/guide/state'),
        item('3. Layout and responsive UI', '3. 布局与响应式', '/guide/layout'),
        item('4. Styles and themes', '4. 样式与主题', '/guide/style'),
        item('5. Building a form', '5. 构建表单', '/guide/forms'),
        item('6. Background tasks', '6. 后台任务与生命周期', '/guide/async'),
        item('7. Multiple windows', '7. 多窗口与独立更新', '/guide/multi-window'),
        item('8. Troubleshooting', '8. 资源、性能与排错', '/guide/troubleshooting'),
      ] },
      { text: label('AI Integration', 'AI 集成'), collapsed: false, items: [
        item('llms.txt', 'llms.txt', '/ai/llms-txt'),
        item('dxui-skill', 'dxui-skill', '/ai/dxui-skill'),
      ] },
      { text: label('Components', '组件'), collapsed: false, items: [
        item('Overview', '组件总览', '/components/'),
        ...components.map(([text, path]) => ({ text, link: `${prefix}/components/${path}` })),
      ] },
      { text: label('Other APIs', '其它 API'), collapsed: false, items: [
        item('Common component contracts', '组件共同契约', '/api/common'),
        item('Application and diagnostics', '应用、快捷键与诊断', '/api/application'),
        item('Window management', '窗口管理', '/api/windows'),
        item('Geometry, layout, and painting', '几何、布局与绘制值', '/api/values'),
        item('Themes and tokens', '主题与 Token', '/api/theme'),
        item('Fonts, images, and icons', '字体、图片与图标资源', '/api/resources-guide'),
        item('Complete icon catalog', '完整图标目录', '/api/icons'),
      ] },
      { text: label('Full Reference', '完整声明与维护'), collapsed: true, items: [
        item('Public API index', '全部公开 API 索引', '/api/'),
        item('Runtime', '运行时', '/api/runtime'),
        item('Component types', '组件关联类型', '/api/components'),
        item('Style types', '样式类型', '/api/style'),
        item('All tokens', '全部 Token', '/api/tokens'),
        item('Resource types', '资源类型', '/api/resources'),
        item('Running examples', '运行完整示例', '/examples/'),
        item('Changelog', '更新日志', '/api/coverage'),
      ] },
    ],
    outline: { label: label('On this page', '本页目录'), level: 2 },
    docFooter: { prev: label('Previous page', '上一篇'), next: label('Next page', '下一篇') },
    returnToTopLabel: label('Back to top', '回到顶部'),
    sidebarMenuLabel: label('Menu', '目录'),
    darkModeSwitchLabel: label('Appearance', '外观'),
    darkModeSwitchTitle: label('Switch to dark theme', '切换至深色主题'),
    lightModeSwitchTitle: label('Switch to light theme', '切换至浅色主题'),
    langMenuLabel: label('Change language', '切换语言'),
    skipToContentLabel: label('Skip to content', '跳转到内容'),
  }
}

export default defineConfig({
  lang: 'en', title: 'DXUI', cleanUrls: true,
  sitemap: { hostname: 'https://dxui-org.github.io' },
  description: 'DXUI tutorials, complete component examples, and public API references for Go developers.',
  locales: {
    root: { label: 'English', lang: 'en', themeConfig: localeTheme(false) },
    'zh-cn': {
      label: '简体中文', lang: 'zh-CN',
      description: '面向 Go 开发者的 DXUI 教程、完整组件示例与公开 API 参考。',
      themeConfig: localeTheme(true),
    },
  },
  head: [
    ['link', { rel: 'describedby', type: 'text/plain', href: '/llms.txt' }],
    ['link', { rel: 'icon', type: 'image/png', href: '/brand/icon.png' }],
    ['link', { rel: 'apple-touch-icon', href: '/brand/icon.png' }],
    ['meta', { name: 'theme-color', content: '#087f8c' }],
  ],
  themeConfig: {
    logo: { src: '/brand/icon.png', alt: '' },
    i18nRouting: (_data, route, targetLocale) => {
      const page = route.data.relativePath.replace(/^zh-cn\//, '')
        .replace(/(^|\/)index\.md$/, '$1').replace(/\.md$/, '')
      // API IDs are shared; translated heading IDs may differ between languages.
      const hash = /^#(?:api-|first-window$|explore-docs$|common-tasks$)/.test(route.hash)
        ? route.hash : ''
      return (targetLocale === 'zh-cn' ? '/zh-cn/' : '/') + page + route.query + hash
    },
    socialLinks: [{ icon: 'github', link: 'https://github.com/dxui-org/dxui' }],
    search: {
      provider: 'local',
      options: {
        locales: {
          'zh-cn': {
            translations: {
              button: { buttonText: '搜索', buttonAriaLabel: '搜索文档' },
              modal: {
                displayDetails: '显示详细列表', resetButtonTitle: '清除搜索',
                backButtonTitle: '返回', noResultsText: '没有找到相关结果',
                footer: { selectText: '选择', navigateText: '切换', closeText: '关闭' },
              },
            },
          },
        },
      },
    },
  },
})
