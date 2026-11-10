# Building and playing

Writing a pack is a loop: edit, check, build, look at it in the game. This guide is about making that loop short, and about what to do with the pack once it works.

## The loop

| Step | Command line | Web editor | VS Code |
| --- | --- | --- | --- |
| Check | `mms check` | as you type | as you type |
| Test | `mms test` | Tests tab, on every edit | Testing view |
| Build into a world | `mms build --world <world>` | World tab | **Build into world** in the MMS side bar |
| Share | `mms build --zip pack.zip` | **Download pack** | |

Most mistakes never need the game. The checker knows the game's schema and command tree, [tests](testing.md) run functions in an interpreter, and the previews draw terrain, loot, recipes, books and structures. Open the game for what only the game can show: how it feels to play.

## Check, fix and format

```sh
mms check          # errors and warnings
mms check -hints   # style suggestions too
mms fix            # apply every safe fix
mms fmt            # format every file
```

`mms check` exits with 1 when there is an error, and `mms build` refuses to write a pack with errors in it, so a broken pack never reaches a world. `mms fix` applies only fixes that leave the built pack unchanged. `mms fmt -check` lists unformatted files without changing them, for CI. [Lint rules](lint.md) covers the rules and how to set their severity.

## Build into a world

```sh
mms build --world ~/.minecraft/saves/MyWorld
```

This writes the data pack to `<world>/datapacks/<namespace>`. If the project has textures, sounds or other resource pack files, it merges them into `<world>/resourcepacks/resources.zip`, the pack the game applies by itself when the world opens. Several projects can build into one world: each keeps its own files.

The world's folder is under `saves` in the Minecraft folder: `%APPDATA%\.minecraft\saves` on Windows, `~/Library/Application Support/minecraft/saves` on macOS and `~/.minecraft/saves` on Linux.

- **Web editor.** In Chrome and Edge, the World tab's **Choose world…** picks the folder, and **Build into world** does the same build. Other browsers cannot write to a folder and offer the downloads.
- **VS Code.** The MMS side bar lists the worlds of the Minecraft launcher, Prism Launcher and MultiMC. Pick one and click **Build into world**.

## Reload or rejoin

The last line of a build says what the game needs:

| The build says | When | Do |
| --- | --- | --- |
| `next step: run /reload in the world` | functions, loot tables, recipes, advancements, predicates, item modifiers, structure templates or tags of blocks, items and entities changed | run `/reload` |
| `next step: exit and rejoin the world` | worldgen, dimensions, enchantments and other registry entries, or the resource pack changed | save and quit to title, then open the world |
| `next step: none` | the build wrote the same files as before | nothing |

Function work is therefore the fastest loop: keep the game open, build, `/reload`, try.

Worldgen is slower, because the game never generates a chunk twice. Add `--reset-dimensions` to delete the saved chunks of the project's own dimensions whose worldgen changed, so the same place generates again. The world must be closed. See [Seeing worldgen changes](../start/tutorial-dimension.md#seeing-worldgen-changes).

## A world for testing

A few habits save time:

- Make a creative world just for the pack, with cheats on, and keep it out of worlds you care about until the pack works.
- Give every pack a `give` function that hands over its items, since custom items are not in the creative menu.
- `/function <namespace>:` and Tab lists the pack's functions. Names under `mms_gen/` are written by MMS, and [How functions compile](../concepts/how-functions-compile.md) explains them.
- To look at state, `/scoreboard objectives setdisplay sidebar <objective>` shows an objective and `/data get storage <namespace>:mms` shows every `Storage` record and string.
- When something does nothing, check the game's log. The game reports files it could not load there, and `mms check` should have caught it first: if it did not, that is a bug worth [reporting](https://github.com/MinecraftMetascript/mms/issues).

## What gets built

```sh
mms build              # writes build/ and, if needed, build-resources/
mms build --out dist   # another folder
```

The `build` folder is a complete data pack: `pack.mcmeta`, and under `data/<namespace>/` one JSON or `.mcfunction` file per declaration. It is plain output, so read it when you want to know exactly what a declaration became. The web editor's JSON tab shows the file of the declaration at the cursor, and its Pack tab (the **Pack** view in VS Code) lists every file.

`mms.json` at the project's root holds the settings:

```json
{
  "$schema": "https://minecraftmetascript.github.io/mms/mms.schema.json",
  "namespace": "gems",
  "description": "Gems that drop from the gem chest",
  "icon": "art/icon.png",
  "output": "build",
  "sources": ["src/**/*.mms"],
  "lint": { "unused": "warning" }
}
```

Only `namespace` is needed. `description` and `icon` are what the game's pack list shows, `sources` limits the project to some files, and `resourcePack.output` moves the resource pack. Editors complete and check the file from its schema.

## Share it

```sh
mms build --zip gems.zip
```

This writes `gems.zip`, and `gems-resources.zip` when there is a resource pack. In the web editor, **Download pack** and **Download resources** give the same two files.

- **A data pack** goes in a world's `datapacks` folder, as the zip. On a server, that is the `datapacks` folder of the server's world.
- **A resource pack** goes in each player's `resourcepacks` folder, or in `<world>/resourcepacks/resources.zip` for a single-player world. A server sends one to its players with `resource-pack` in `server.properties`.

A pack with custom item textures needs both. Without the resource pack the items still work and keep their names, and look like their base items.

## In CI

`mms` is one binary with no dependencies, so checking a pack on every push takes three lines:

```sh
mms fmt -check
mms check
mms test
```

Each exits with 1 on failure. `mms build --zip` after them gives a release file.

## Start from an existing pack

`mms import` turns a data pack into a project:

```sh
mms import old-pack.zip my-project
```

It writes `mms.json` and one `.mms` file per resource folder, such as `loot_table.mms`, using the shorthand wherever it applies. In the web editor, **Import zip…** does the same, and **Paste JSON…** converts one JSON file into a declaration at the cursor. Importing the vanilla data pack is a good way to see how the game's own loot tables or worldgen read in MMS.

Functions come across as they are, command for command. They are valid MMS, and you can rewrite them with variables and control flow one at a time: see [Functions](functions.md).
