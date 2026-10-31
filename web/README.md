# MMS web editor

The browser editor for MMS: SvelteKit 3, Tailwind CSS 4 and CodeMirror 6, built as a static site for GitHub Pages.

- The engine runs as WASM in a Web Worker (`static/mms-worker.js`) and speaks LSP to the editor through `@codemirror/lsp-client`. A supervisor restarts it if it crashes or hangs.
- Previews render with Deepslate in a second worker. Vanilla files that a pack references are fetched from mcmeta at the commit the spec pack was built from.
- Projects are kept in IndexedDB and can be imported and exported as zips. Importing a data pack zip, or pasting one JSON file, converts it to MMS through the engine (`mms/importPack`).

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
