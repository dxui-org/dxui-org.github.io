# Building a form

Build a settings form: enter a name, select a language, accept the terms, and save.

## 1. Run the example

<<< ../examples/form/main.go

## 2. Try it in order

1. Leave the name empty and observe that **Save** is disabled.
2. Enter a name, choose a language, and check **Accept terms**.
3. Click **Save** to display `Saved: ...`. Pressing Enter in the name input also submits.
4. Clear the name or uncheck the terms; Save becomes disabled again.

Saving in this example updates text only; it does not write a file.

## 3. Understand the form logic

| Code | Purpose |
| --- | --- |
| `name`, `language`, `accepted` | Store the three control values. |
| `OnChange: dxui.Assign(...)` | Writes changes into the corresponding variable. |
| `valid()` | Checks for a nonempty name and accepted terms. |
| `Disabled: !valid()` | Disables Save while the conditions are unmet. |
| `submit()` | Rechecks conditions, then updates the message. |

Both the button and input Enter action call submit, so validate inside that function instead of relying only on the disabled button.

Try requiring at least two characters in valid(), then submit through both the button and Enter.

For additional configuration, see [Input](/components/input), [Select](/components/select), and [Checkbox](/components/checkbox).

Next: [Run background tasks](./async).
