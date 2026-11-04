# MMS branding

The MMS mark is a block held between code braces: a gold top over blue sides, inside `{ }`. The wordmark is "MMS" over a tracked "METASCRIPT", set in [Chakra Petch](https://fonts.google.com/specimen/Chakra+Petch) (SIL Open Font License) and converted to outlines, so no font is needed to show it.

| File | Use |
| --- | --- |
| `logo-light.svg`, `logo-dark.svg` | Full logo (mark and wordmark) for light and dark backgrounds |
| `mark.svg` | Mark alone; the braces switch colour with the viewer's colour scheme |
| `mark-light.svg`, `mark-dark.svg` | Mark alone with fixed colours, for light and dark backgrounds |
| `favicon.svg` | Mark on a blue tile, drawn to read at 16px. Used as the web favicon and the VS Code `.mms` file icon |
| `icon-128.png` | `favicon.svg` at 128px, the VS Code extension icon |

Copies live where each tool needs them:

- `web/src/lib/assets/favicon.svg` and `web/static/apple-touch-icon.png` (180px) for the web editor and docs
- `web/src/lib/components/Logo.svelte`, the header logo, with braces and text in `currentColor`
- `editors/vscode/icons/mms.png` and `editors/vscode/icons/mms-file.svg`

Zed extensions have no icon field, so the Zed extension carries no logo.

## Colours

| | Light | Dark |
| --- | --- | --- |
| Cube top | `#f2b01e` | `#fbbf24` |
| Cube left | `#2563eb` | `#3b82f6` |
| Cube right | `#1a46c4` | `#2459d6` |
| Braces and "MMS" | `#1b1f24` | `#e6e9ed` |
| "METASCRIPT" | `#65707d` | `#8d97a3` |
| Favicon tile | `#2563eb` | `#2563eb` |

## Rebuilding

`build.py` writes every SVG here plus `Logo.svelte`. It needs `fonttools` and the Chakra Petch Medium (500) and Bold (700) TTFs, named `cp500.ttf` and `cp700.ttf`:

```sh
pip install fonttools
python3 branding/build.py <font-dir> branding web/src/lib/components/Logo.svelte
```

Render the PNGs from `favicon.svg` at 128px and 180px with any SVG renderer (a headless browser screenshot works), then copy the files to the places listed above.
