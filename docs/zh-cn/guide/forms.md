# 构建一个表单

这一节做一个设置表单：输入名称、选择语言、勾选条款，最后点击保存。

## 1. 运行示例

<<< ../../examples/form/main.go

## 2. 按顺序试一遍

1. 不填名称，观察 **Save** 按钮不可用。
2. 输入名称，选择语言，再勾选 **Accept terms**。
3. 点击 **Save**，下方显示 `Saved: ...`；在名称输入框中按 Enter 也能提交。
4. 清空名称或取消勾选，保存按钮再次变为不可用。

这里的“保存”只更新界面文字，没有写入文件。

## 3. 看懂表单逻辑

| 代码 | 作用 |
| --- | --- |
| `name`、`language`、`accepted` | 保存三个控件的值。 |
| `OnChange: dxui.Assign(...)` | 把操作结果写回对应变量。 |
| `valid()` | 检查名称不为空且已勾选条款。 |
| `Disabled: !valid()` | 未满足条件时禁用保存按钮。 |
| `submit()` | 再检查一次条件，然后更新提示文字。 |

按钮和输入框的 Enter 都调用 `submit`，所以校验要写在这个函数里，不能只依赖按钮禁用。

试着给 `valid()` 增加“名称至少两个字符”的条件，再分别用按钮和 Enter 提交。

需要更多控件配置时，查看 [Input](/zh-cn/components/input)、[Select](/zh-cn/components/select) 和 [Checkbox](/zh-cn/components/checkbox)。

下一步：[执行后台任务](./async)。
