# @mms/preview

The worldgen previews (Deepslate renders of noise, density functions, noise settings, surface rules and dimensions) and the loot table, recipe and biome inspectors. The web editor (`web/`) and the VS Code extension's preview panel (`editors/vscode/webview/`) both import these components; there is no other copy.

It is a source-only package: hosts compile the Svelte and TypeScript themselves. A host:

- links it with `"@mms/preview": "file:…/packages/preview"`,
- installs `svelte` and `deepslate` and lists both in Vite's `resolve.dedupe`,
- imports `@mms/preview/theme.css` after `tailwindcss`, defines the `--c-*` colours and a `dark` variant,
- passes `createWorker`, starting `@mms/preview/worker` (`?worker` on the web, `?worker&inline` in a VS Code webview, which only allows blob workers).

`PreviewPane` is the whole pane (inspector or preview, plus the "pinned output is missing" note); `PinButton` is the pin toggle. `npm run check` type-checks the package on its own.
