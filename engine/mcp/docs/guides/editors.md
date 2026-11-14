# A tour of the editors

MMS has two editors: the [web editor](https://minecraftmetascript.github.io/mms/), which needs no install, and the VS Code extension. Both run the same engine as the `mms` command line, so they report the same problems and build the same pack. This page shows what they can do and where to find it.

## Choosing one

- **The web editor** is the quickest start. Projects are kept in your browser, examples are one click away, and **Download pack** gives you a zip. Use it to try the language, to share a small pack, or on a computer where you cannot install anything.
- **VS Code** suits a project you keep: the files are on your disk, so they work with git, with other tools and with `mms` in a terminal. Install **Minecraft Metascript** from the Extensions view. It also works in Cursor, VSCodium and other editors that use Open VSX.

A project moves between them as a zip: **Export project zip** and **Import zip…** in the web editor's project menu.

## While you type

These work the same in both editors.

- **Errors as you type.** Every value is checked against the game's schema and every command against its command tree. Unknown ids, missing fields, wrong types and misspelled commands are underlined with the reason.
- **Completion.** It offers what fits at the cursor: the fields of an object, the helpers and ids a field takes, the next word of a command, selector options, your own functions and objectives, and the markup inside text.
- **Hover.** A field shows its documentation, a helper its signature, a name its doc comment, a `const` its value, a command the forms it can take, and a piece of text how the game will draw it.
- **Go to definition, references and rename** work across files, on declarations and on ids inside commands. Go to definition on a vanilla id, such as a loot table or a noise, opens the game's own JSON for it. On `CustomItem` or another built-in Blueprint it opens the prelude source.
- **Quick fixes.** Ctrl/⌘+. lists the fixes at the cursor, such as adding missing fields, correcting a misspelled name or declaring something that does not exist yet. **Fix all** applies every safe one.
- **Format.** One style, nothing to configure. Ctrl/⌘+S formats in the web editor, and VS Code formats on save.
- **Inlay hints** mark template calls that are written out at compile time (`inlined`) and show loop values beside generated names.

## Previews

Put the cursor on a declaration and open the **Preview** tab in the web editor. In VS Code, click the **Preview** lens above the declaration, and the panel stays on it while you edit.

| Declaration | The preview shows |
| --- | --- |
| `Noise`, `DensityFunction`, `NoiseSettings`, `Dimension` | the terrain: density, blocks or 3D. See [Shaping terrain](terrain.md) |
| `Dimension` | also a map of its biomes and the climate behind them. See [Why a biome doesn't appear](#why-a-biome-doesnt-appear) |
| `SurfaceRule` | the blocks it places on the terrain |
| a `Spline` | its curve, which you can drag. See [Splines](resources.md#splines) |
| `LootTable`, `Recipe`, `Biome` | an inspector: a table's pools and entries, a recipe's grid, a biome's settings |
| `Book`, a `.mmsbook` file | the open book, page by page. See [Previewing a book](resources.md#previewing-a-book) |
| `tell`, `title`, `actionbar`, a `Text` | the message in chat, as a title or above the hotbar |
| `JigsawStructure`, `TemplatePool` | which pieces join which, and a 3D simulation of one layout. See [Simulating a jigsaw structure](resources.md#simulating-a-jigsaw-structure) |

### Why a biome doesn't appear

A `MultiNoise` biome source picks, at each spot, the entry whose `parameters` are nearest the climate sampled there. The dimension's **Biomes** map helps when a biome never wins:

- Hover the map to fill the table below it: the six climate values at that spot (at the camera's Y, which is what `depth` depends on), then the nearest entries, best first, with each one's distance and parameter ranges. The ranges the climate is outside of are in red. Click the map to pin a point: the table then stays on it while you move the pointer, until you click **Unpin**.
- **Show** switches the map to one climate parameter, coloured over its range in view, with the biome outlines kept on top.
- **Explain** picks a biome to follow. Biomes the source lists that win nowhere in view are marked *not in view*, and also listed under the legend. If the chosen biome isn't among the nearest entries, its nearest entry is added to the table in amber. **Show: Distance to win** maps how far that biome is from winning: 0 where it wins, and the darker the closer.

Distances are the game's: the square root of the sum of squares of how far each value is outside its range, with the entry's `offset` added as one more term.

Previews redraw as you type. The pin button in the web editor keeps the Preview tab on one declaration while you edit another, which is how you watch terrain change while editing the noise behind it.

Each declaration's preview remembers how you left it: the mode, camera and zoom, the 3D view's look, cuts and threshold, a jigsaw simulation's seed and slice, and a spline's coordinate sliders. Go back to a declaration and it shows the same view. The web editor keeps this in your browser for each project, and VS Code keeps it with the preview panel. **Reset view** puts the camera back.

## Seeing the output

- **JSON** (web) shows the file the declaration at the cursor becomes.
- **Expanded** (web), and the **Expand** lens in VS Code, show what a Blueprint instance such as a `CustomItem`, or a command template call, writes.
- **Pack**, in the sidebar of both, lists every file of the data pack and the resource pack.

Reading the output is the fastest way to learn what a piece of shorthand means, and to trust it.

## Tests

Once a project has a [`Test`](testing.md), the web editor shows a **Tests** tab that runs them all after every edit. VS Code lists them in its Testing view. A failing `expect` is marked at its line.

## Working with a world

| | Web editor | VS Code |
| --- | --- | --- |
| Build the pack into a world | **World** tab (Chrome and Edge) | **World** view in the MMS side bar |
| Import pieces saved with structure blocks | **Structures**, in the World tab | **Structures** view |
| Set up a jigsaw block and copy its `/give` | **Jigsaw** tab | **Jigsaw** view |

After each build the editor says whether the game needs `/reload` or a rejoin. See [Building and playing](building.md).

## The web editor's own corners

- **Projects.** Click the project's name to switch projects or start a new one. They are saved in the browser, so keep a copy of anything you would be sad to lose: **Export project zip**, in the project menu, saves the sources. **Download pack** is the built pack, which cannot be turned back into a project.
- **Examples** opens any of the example projects as a new project.
- **Paste JSON…** turns a JSON file from a data pack into the MMS declaration that writes it.
- **+ Add texture or sound** adds `.png`, `.ogg` and `.nbt` files to the project.
- **Symbols**, in the sidebar, lists the project's declarations.
- **Docs search.** On these pages, Ctrl/⌘+K or `/` searches every page and the reference.

## VS Code's own corners

- **`mms.json`** gets completion, hover and validation, including every lint rule's name.
- **A separate world for structures.** Build pieces in one world and test the pack in another: **Use a separate world…** in the Structures view.
- **Pre-release builds.** **Switch to Pre-Release Version** on the extension's page follows every change merged to `main`.

## For AI agents

An agent that writes MMS can use the same checks through [`mms mcp`](agents.md), along with text versions of the biome map, the terrain previews and the jigsaw simulation ([Preview tools](agents.md#preview-tools)). It can also read these pages as Markdown: every page has a `.md` twin, and [llms.txt](https://minecraftmetascript.github.io/mms/llms.txt) links them all.
