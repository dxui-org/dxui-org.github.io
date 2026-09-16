# llms.txt

Give a web-enabled AI assistant the following address so it can consult DXUI documentation before writing your UI:

[https://dxui-org.github.io/llms.txt](https://dxui-org.github.io/llms.txt)

## What it contains

`llms.txt` is a Markdown documentation index with a DXUI introduction and annotated links to tutorials, components, and public APIs. An assistant can follow the relevant links for parameters and examples.

This file is an index, not the full documentation. Its format follows the [llms.txt specification](https://llmstxt.org/).

## How to use it

Provide the documentation entry point together with your request. For example:

```text
First read https://dxui-org.github.io/llms.txt,
then consult the State and Events, Input, and Button pages.

Use Go and DXUI to build a window with a name input and a Submit button.
After submission, display the entered name in the window.
Provide a complete, runnable main.go and explain how to run it.
```

For an existing project, also provide its DXUI version from `go.mod` and relevant code so the assistant can check API compatibility.

## When the tool cannot read web pages

Open [llms.txt](/llms.txt) and paste its contents into the conversation. Add the relevant page text or complete examples as needed. Supplying a URL alone does not mean the tool has read it.

Tools access external documentation differently. Ask the assistant to list the pages it actually consulted, then run the generated code locally and check its behavior.
