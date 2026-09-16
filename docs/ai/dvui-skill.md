# dxui-skill

A dedicated AI coding skill for DXUI, with complete public API references, component guides, and runnable Go examples for building and modifying applications.

Repository: [dxui/dxui-skill](https://github.com/dxui-org/dxui-skill)

## What is included

- **Coding guidance**: State, events, layout, styles, background tasks, and multiple windows.
- **Complete API references**: The root package, official icons and catalog, and types and methods exposed through public aliases.
- **Components and examples**: Behavior guides and complete window, form, theme, and multi-window programs.
- **Synchronization checks**: Compare the bundled API references with a local DXUI checkout.

The assistant reads `SKILL.md`, then loads only the references and examples needed for the task.

## How to use it

Clone the repository:

```sh
git clone https://github.com/dxui-org/dxui-skill.git
```

Follow your AI tool's skill installation instructions to add the entire `dxui-skill` directory. Keep `SKILL.md`, `references`, `assets`, and `scripts` together so the tool can access the supporting files.

After installation, name the skill in your request. For example:

```text
Use dxui-skill to add a settings window to this Go project.
Include a name input, theme selector, and Save button.
Check the DXUI version in go.mod before choosing the APIs.
```

If your tool cannot install skills but can read local files, ask it to read the repository's `SKILL.md` and follow its reference links.

## Using it with llms.txt

[llms.txt](./llms-txt) provides an online documentation index. dxui-skill keeps coding guidance, API references, and examples locally for project development.

The bundled API references target a particular source version. Use the version in your project's `go.mod`, then build locally and check the generated UI.
