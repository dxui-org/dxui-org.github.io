# DXUI Documentation

Tutorials, component examples, and API references for [DXUI](https://github.com/dxui-org/dxui), built with VitePress.

[English](https://dxui-org.github.io/) · [中文](https://dxui-org.github.io/zh-cn/)

## Development

```sh
pnpm install
pnpm docs:dev
```

Build and preview:

```sh
pnpm docs:build
pnpm docs:preview
```

## Structure

- `docs/` — English pages (default).
- `docs/zh-cn/` — Chinese pages.
- `docs/examples/` — Shared, runnable Go examples.
- `docs/public/` — Brand assets and `llms.txt`.
- `scripts/` — API generation, translations, and validation.

Keep both languages in sync. Edit generated content in `scripts/build-reference.py` and `scripts/i18n-en.json`; edit tutorials and other handwritten pages directly. Changelogs record DXUI changes only.

## Validation

After building, check language coverage and internal links (requires Python 3):

```sh
pnpm docs:check-i18n
```

Deployment uses the GitHub Pages workflow.
