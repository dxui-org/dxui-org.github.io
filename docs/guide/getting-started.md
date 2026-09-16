# Your first window

Run a desktop window with text and a button. You need Go 1.25 or newer and a Windows, macOS, or Linux desktop environment.

## 1. Create a project

```sh
mkdir hello-dxui
cd hello-dxui
go mod init example.com/hello-dxui
go get github.com/dxui-org/dxui
```

## 2. Copy the code

Create `main.go` in the project directory and copy this complete program:

<<< ../examples/hello/main.go

## 3. Run it

In the project directory, run:

```sh
go run .
```

If you see the **Hello dxui** window, it worked. Click **Close** to close it.

## What the code does

| Code | Purpose |
| --- | --- |
| `NewApp` | Configures the window title, width, and height. |
| `app.Run` | Opens the window and runs until the application closes. |
| `Box` | Arranges its children vertically by default. |
| `Label` / `TextButton` | Displays text / a button. |
| `OnPress` | Function called when the button is activated. |

Try changing Title and the Label text, then run again. Call Run from main, not from a goroutine.

::: details Optional: build an executable without CGO

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

Keep `go.mod` and `go.sum` to record dependency versions.

:::

If the window does not open, inspect the terminal error, then consult [troubleshooting](./troubleshooting).

Next: [Make a button update the UI](./state).
