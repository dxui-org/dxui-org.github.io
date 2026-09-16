# 状态与事件

这一节做一个计数器：每点一次按钮，数字加一。

## 1. 运行计数器

用下面的代码替换上一节的 `main.go`，执行 `go run .`。后面的完整示例也用这个方式运行。

<<< ../../examples/button/main.go

点击 **Add**，观察数字变化。

## 2. 看懂更新过程

1. `count` 保存当前数字，声明在 `app.Run` 外面。
2. `OnPress` 在点击时执行 `count++`。
3. DXUI 再次运行构建函数，`Label` 读到新数字，界面随之更新。

**状态放在构建函数外，修改放在事件回调里。** 如果把 `count := 0` 放进构建函数，每次更新都会归零。构建函数只负责描述界面，不要在里面启动任务或修改状态。

试着把 `count++` 改成 `count += 2`，再运行一次。

## 3. 保存输入的文字

输入框也使用同样的方式：`Value` 读取变量，`OnChange` 把新文字写回变量。

```go
// 放在 app.Run 外。
name := ""
```

```go
// 在构建函数返回的 Box 中添加这个控件。
dxui.Input(dxui.InputProps{
    Value: name,
    OnChange: func(next string) {
        name = next
    },
})
```

`next` 是输入框的完整新内容，不需要再拼接。上面的回调也可以简写为 `OnChange: dxui.Assign(&name)`。不写回新值，输入框就不会保留这次修改。

普通组件回调修改本窗口状态后，无需手动刷新。后台任务使用 [App.Update](./async)；更新其它窗口见[多窗口](./multi-window)。

::: details 列表中的控件为什么需要 Key

动态添加、删除或排序控件时，用稳定的业务 ID 设置 `Key`，让输入选区和焦点跟随正确的控件。不要用会变化的位置作为 ID。

Key 只在同一父级内有效；卸载后不保证保留运行时状态。业务值仍需放在控件外。完整规则见[组件共同契约](/zh-cn/api/common)。

:::

输入长度限制、只读和禁用的区别见 [Input](/zh-cn/components/input)。

下一步：[排列界面](./layout)。
