# MMS web editor

The browser editor for MMS: SvelteKit 3, Tailwind CSS 4 and CodeMirror 6, built as a static site for GitHub Pages.

- The engine runs as WASM in a Web Worker (`static/mms-worker.js`) and speaks LSP to the editor through `@codemirror/lsp-client`. A supervisor restarts it if it crashes or hangs.
- Previews render with Deepslate in a second worker: noises and density functions as sampled values, noise settings and dimensions as block terrain slices (aquifers and surface rules included), surface rules drawn on a terrain, and dimensions as biome maps. Vanilla worldgen that a pack references is fetched from mcmeta's 26.3 summary, one bundle per registry.
- The preview follows the declaration under the cursor. Pin it (the button by the Preview tab) to keep one declaration in view while you edit others, such as a dimension while you tweak its surface rule; it still redraws on every change.
- Several projects are kept in IndexedDB (`src/lib/project/store.ts`), one key per project plus an index; the project name in the header opens the switcher. New projects, examples and imports open as new saved projects. Projects can be imported and exported as zips. Importing a data pack zip, or pasting one JSON file, converts it to MMS through the engine (`mms/importPack`).

## Develop

```sh
npm ci
npm run wasm     # builds static/mms.wasm and copies wasm_exec.js (needs Go)
npm run dev
```

## Check and test

```sh
npm run check
npm run build
npm test         # Playwright smoke tests against the build
```

Set `BASE_PATH=/mms` to build for the GitHub Pages path, and `PW_CHROMIUM` to use a preinstalled Chromium.
