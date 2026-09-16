# llms.txt

让 AI 助手先读取 DXUI 文档，再帮助你编写界面。把下面的地址提供给支持网页读取的 AI 工具即可：

[https://dxui-org.github.io/llms.txt](https://dxui-org.github.io/llms.txt)

## 它包含什么

`llms.txt` 是用 Markdown 编写的文档导航，包含 DXUI 简介，以及教程、组件和公开 API 的链接与简短说明。AI 可以从中找到相关页面，再按需读取具体参数和示例。

本站的文件是文档索引，不包含全部文档正文。格式参考 [llms.txt 规范](https://llmstxt.org/)。

## 怎么使用

把文档入口和你的需求一起发给 AI。例如：

```text
请先读取 https://dxui-org.github.io/llms.txt，
再查看状态与事件、Input 和 Button 的文档。

用 Go 和 DXUI 编写一个包含名称输入框和提交按钮的窗口。
点击提交后，在窗口中显示输入的名称。
请提供可以直接运行的完整 main.go，并说明运行方法。
```

已有项目时，再提供 `go.mod` 中的 DXUI 版本和相关代码，方便 AI 核对 API。

## 工具不能读取网页时

打开 [llms.txt](/llms.txt)，将内容粘贴给 AI；然后按任务需要，补充对应文档页的正文或完整示例。只提供链接并不代表工具已经读取了链接中的内容。

不同工具读取外部文档的方式不同，提供入口后可要求它列出实际参考的页面。生成代码后，在本地运行并检查交互效果。
