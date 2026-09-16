# 后台任务与生命周期

耗时计算放进 goroutine，完成后通过 `App.Update` 更新界面，避免窗口卡住。

## 1. 运行示例

<<< ../../examples/async/main.go

点击 **Compute**，完成后会显示计算结果。计算期间按钮禁用；任务较快时，可能看不到中间的 Computing 提示。

## 2. 按三个步骤处理任务

1. **点击时准备数据**：设置 `busy`，把计算需要的输入放进局部变量 `n`。
2. **后台计算**：goroutine 只使用准备好的输入，算出 `result`。
3. **回到界面更新**：在 `app.Update` 的闭包中修改 `message` 和 `busy`。

不要在 goroutine 中直接修改界面正在读取的变量。`Assign` 只是赋值函数，不能代替 `Update`。

`Update` 返回 nil 说明工作已排队，不代表闭包已经执行。应用关闭后提交会失败，因此示例检查并处理返回的错误。

## 3. 关闭时停止任务

示例使用 `context.WithCancel`。系统请求关闭窗口时，`OnCloseRequest` 先调用 `cancel()` 停止后台工作，再调用 `app.Close()` 退出。

设置了 `OnCloseRequest` 后，需要显式调用 `Close` 才会接受关闭。`Run` 返回后也调用一次 `cancel()`，覆盖其它退出路径。

试着增加计算量，并在任务执行中关闭窗口。后台任务应响应取消，不再继续提交结果。

::: details 应用启动、关闭和错误回调

NewApp 不加载原生窗口；配置/字体/主题/原生启动错误由 Run 返回。OnShown 仅在完整第一帧呈现并成功显示后调用一次，适合启动依赖窗口已就绪的工作。不要把 root builder 当作首次显示通知。

主窗口默认系统关闭请求会结束应用，同时关闭全部子窗口。如果设置 `OnCloseRequest func(*App)`，你必须在接受请求时调用 Close；不调用就是保持窗口打开，可用于未保存提示。App.Close 幂等且允许跨 goroutine，Run 前调用为空操作。Window.Close 只关闭目标子窗口，见[多窗口教程](./multi-window)。

OnError 在 UI 线程观察可恢复的构建/回调失败；为 nil 时错误终止 Run 并返回。错误处理器应记录或调整导致问题的业务状态，不能依赖一个持续出错的 builder 自动恢复。零 View、非法条目键等使树更新事务失败，保留上一个有效帧；首次构建没有旧帧可保留。不要用 OnError 忽略启动失败。


:::

快捷键、剪贴板和生命周期参数见[应用 API](/zh-cn/api/application)。

下一步：[打开多个窗口](./multi-window)。
