# 第一个窗口

这一节先运行一个带文字和按钮的桌面窗口。需要 Go 1.25 或更新版本，以及 Windows、macOS 或 Linux 桌面环境。

## 1. 创建项目

```sh
mkdir hello-dxui
cd hello-dxui
go mod init example.com/hello-dxui
go get github.com/dxui-org/dxui
```

## 2. 复制代码

在项目目录新建 `main.go`，复制下面的完整代码：

<<< ../../examples/hello/main.go

## 3. 运行

在项目目录执行：

```sh
go run .
```

看到 **Hello dxui** 窗口就成功了。点击 **Close** 关闭窗口。

## 这段代码做了什么

| 代码 | 作用 |
| --- | --- |
| `NewApp` | 设置窗口标题、宽度和高度。 |
| `app.Run` | 打开窗口，运行到应用关闭。 |
| `Box` | 把里面的控件排列在一起，默认竖排。 |
| `Label` / `TextButton` | 显示文字 / 显示按钮。 |
| `OnPress` | 按钮点击时执行的函数。 |

试着修改 `Title` 和 `Label` 中的文字，再运行一次。`Run` 要在 `main` 中调用，不要放进 goroutine。

::: details 可选：构建可执行文件、关闭 CGO

Windows PowerShell：

```powershell
$env:CGO_ENABLED = "0"
go build -o hello-dxui.exe .
./hello-dxui.exe
```

macOS / Linux：

```sh
CGO_ENABLED=0 go build -o hello-dxui .
./hello-dxui
```

保留 `go.mod` 和 `go.sum`，记录项目使用的依赖版本。

:::

打不开窗口时，先查看终端错误，再看[排错指南](./troubleshooting)。

下一步：[让按钮改变界面](./state)。
