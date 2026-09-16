# dxui-skill

DXUI 专属 AI 编程技能，提供完整公开 API 参考、组件用法和可运行的 Go 示例，帮助 AI 编写和修改 DXUI 应用。

仓库：[dxui/dxui-skill](https://github.com/dxui-org/dxui-skill)

## 包含什么

- **编程指导**：状态与事件、布局、样式、后台任务和多窗口的使用规则。
- **完整 API 参考**：DXUI 根包、官方图标及目录包，以及公开别名涉及的类型和方法。
- **组件与示例**：组件行为说明，以及窗口、表单、主题和多窗口等完整程序。
- **同步检查**：核对技能内的 API 参考与本地 DXUI 源码是否一致。

AI 从 `SKILL.md` 读取使用指导，再按任务查阅相关参考和示例。

## 怎么使用

先获取仓库：

```sh
git clone https://github.com/dxui-org/dxui-skill.git
```

按照所用 AI 工具的技能安装方式，将整个 `dxui-skill` 目录加入技能目录。保留 `SKILL.md`、`references`、`assets` 和 `scripts`，以便工具读取参考和示例。

安装后，在任务中明确使用这个技能，例如：

```text
请使用 dxui-skill，为当前 Go 项目添加一个设置窗口。
窗口包含名称输入框、主题选择和保存按钮。
先检查 go.mod 中的 DXUI 版本，再根据对应 API 实现。
```

如果工具不支持技能安装，但能读取本地文件，可以让它先读取仓库中的 `SKILL.md`，再根据其中的链接查阅所需内容。

## 与 llms.txt 配合使用

[llms.txt](./llms-txt) 提供在线文档入口；dxui-skill 将编程指导、API 参考和示例放在本地，适合在开发项目时使用。

技能中的 API 参考对应特定源码版本。已有项目应以 `go.mod` 中实际使用的版本为准，生成代码后在本地构建并检查交互。
