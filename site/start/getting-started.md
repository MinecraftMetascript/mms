# Getting started

MMS is a small language that compiles to Minecraft Java Edition data packs. You declare resources and functions in `.mms` files, and `mms build` writes the JSON and `.mcfunction` files. The target is Minecraft 26.3.

## Install

- **In the browser:** the [editor](https://minecraftmetascript.github.io/mms/) needs no install. Projects are saved in your browser (click the project name to switch between them), and **Download pack** gives you the data pack as a zip.
- **Command line:** download `mms` for your system from the [releases page](https://github.com/MinecraftMetascript/mms/releases). Or, with Go installed, run `go install github.com/minecraftmetascript/mms/engine/cmd/mms@latest`.
- **VS Code:** install the `.vsix` for your system from the same release, such as `mms-win32-x64-1.0.0.vsix` or `mms-darwin-arm64-1.0.0.vsix`, through **Extensions → … → Install from VSIX**. It includes `mms`, so you do not need the command line to use the editor. (The universal `mms-1.0.0.vsix` does not, and runs `mms` from your `PATH` or `mms.path`.)

## A first project

```sh
mkdir gems && cd gems
mms init -namespace gems
```

`mms init` writes `mms.json` (the namespace, the pack description, the output folder) and a starter file. Its `$schema` line points at the [mms.json schema](https://minecraftmetascript.github.io/mms/mms.schema.json), so editors with JSON Schema support (VS Code, JetBrains) complete, describe and check every field. The VS Code extension knows the schema even without that line. Every `.mms` file in the folder belongs to the project. Put this in `gems.mms`:

```
/** Gems that drop from the gem chest. */
ItemTag gems = [diamond, emerald, amethyst_shard]

LootTable chests/gems {
  pools = [Pool(1..3, #gems)]
}

Recipe gem_block = Shaped(["GG", "GG"], { G = #gems }, emerald_block)

Objective found = dummy

Function find {
  found[@s] += 1
  /loot give @s loot gems:chests/gems
  tell(@s, "You have found ${found[@s]} gem chests")
}
```

Then build:

```sh
mms build                                     # writes build/
mms build -world ~/.minecraft/saves/MyWorld   # installs into a world
```

## Declarations

Each declaration is `Kind name = value`, or `Kind name { fields }` for an object:

- **Kinds** are PascalCase: `LootTable`, `DensityFunction`, `ItemTag`. Every resource folder in a data pack has one. The [reference](https://minecraftmetascript.github.io/mms/docs/reference) lists them all.
- **Names** become ids in the file's namespace: `chests/gems` in namespace `gems` is `gems:chests/gems`, written to `data/gems/loot_table/chests/gems.json`.
- **Doc comments** (`/** ... */`) describe the declaration right below them. See [Doc comments](#doc-comments).

A `Namespace` line switches the namespace for the rest of the file:

```
Namespace minecraft

ItemTag swords = [diamond_sword]
```

Replacing a vanilla resource like this is allowed. The editor notes that it happens, since vanilla's version is lost.

## Doc comments

A comment that starts with `/**` documents the declaration, `let` or parameter list that follows it. Only blank space may come between the two. The text is Markdown, and it appears on hover over the name, both where it is declared and wherever it is used: in other files, in calls, in `kills[@s]` scores, inside commands such as `/function gems:find`, and in `${...}`. Completion and signature help show it too.

```
/**
 * Gives a player their reward and counts it.
 *
 * @param count how many diamonds to give
 * @returns the player's new total
 * @deprecated use `reward_all`
 */
Function reward(count: int): int {
  /** Rewards before this one. */
  let before = found[@s]
  /give @s diamond ${count}
  found[@s] = before + 1
  return found[@s]
}
```

These JSDoc-style tags are understood:

- `@param name text` documents a parameter. Hovering the parameter shows its text, and so does signature help while typing that argument.
- `@returns text` (or `@return`) describes the result.
- `@deprecated [text]` marks the declaration as deprecated. Every use gets a `deprecated` lint warning and is struck through in the editor.
- `@example` starts a code example, and `@see text` adds a pointer elsewhere.
- `@label text` names an entry of a template pool's `elements` list. A doc comment works above an entry too, and the jigsaw preview and its export show its label (or else the first line of its text) next to the piece. That tells apart one structure listed several times with different processors:

  ```
  elements = [
    WeightedElement(50, SinglePoolElement(rigid, Structure("structures/road_cap.nbt"), [])),
    /** @label Worn road cap */
    WeightedElement(10, SinglePoolElement(rigid, Structure("structures/road_cap.nbt"), road_wear_medium)),
  ]
  ```

Other tags are shown as written. The formatter lines doc comments up with a ` * ` on every line and keeps the indentation inside them, so indented code and nested lists survive.

## Projects with several files

A project is every `.mms` file under the folder that holds `mms.json`, including subfolders. The build output folder, `node_modules` and any folder whose name starts with `.` are skipped, at any depth. To limit the project to some files, list glob patterns in `sources` in `mms.json`, such as `["src/**/*.mms"]`.

- **One namespace of names.** A declaration in any file can be used from any other, and the order of files and declarations does not matter. Declaring the same id twice for one kind is an error that points to the other declaration. Tags are the exception: declarations of the same tag merge, so a tag can be built up across files. Different kinds can share an id, such as a `Recipe` and a `LootTable` both named `gem_block`.
- **Namespaces per file.** Each file starts in the namespace from `mms.json`. A `Namespace` line changes it for the rest of that file only.
- **Folders are for you.** A file's folder does not change the ids it declares: `LootTable chests/buried` in `loot/chests.mms` is still `<namespace>:chests/buried`.
- **Editors see the whole project.** The language server loads every project file, not only the open ones, so go-to-definition, references and rename work across files. The web editor's sidebar adds, renames and deletes files, and file names can include folders, such as `recipes/crafting.mms`.

The [`treasure`](../../examples/treasure/) and [`functions-advanced`](../../examples/functions-advanced/) examples are split this way.

## Values

- **Object literals** mirror the JSON: `{ key = value }`. Fields are separated by newlines or commas.
- **Lists** are `[a, b, c]`.
- **Enum values** are bare words: `action = show_text`, `color = red`. They are written exactly as the game spells them; only ids get `minecraft:`.
- **Colors** (fields such as `shadow_color` or a biome's `water_color`) take `"#RRGGBB"`, or `"#AARRGGBB"` where the color has an alpha (alpha first, so `"#80FF0000"` is half-transparent red). The game's own forms work too: an integer, or `[r, g, b]` / `[r, g, b, a]` with each part from 0 to 1. MMS writes a `"#hex"` string as the integer where the game wants one.
- **Ids** can drop `minecraft:`. A bare name such as `diamond` or `#gems` means your own declaration if one exists, otherwise the vanilla one.
- **Helpers** build typed objects: `Pool(...)` is a loot pool, and `Noise(...)` is `{ type = noise, ... }`. Every dispatched type has a helper named after its type, and arguments can be positional or named. See [Writing resources](../guides/resources.md) and the [reference](https://minecraftmetascript.github.io/mms/docs/reference).

Every value is checked against the 26.3 schema as you type. Completion offers the fields, helpers and ids that fit where the cursor is, and hover shows the field's documentation.

## Formatting and lint

`mms fmt` (and format-on-save in the editors) prints every file in one style, so there is nothing to configure. `mms check` reports errors and [lint](../guides/lint.md) findings, and `mms fix` applies the fixes that are always safe.

## Next

- [Writing resources](../guides/resources.md): the shorthand for loot tables, recipes, density functions and more.
- [Functions](../guides/functions.md): variables, control flow, calls and strings in function code.
- [Tutorial: a custom dimension](tutorial-dimension.md): a complete worldgen project.

## Importing

`mms import` reads a data pack folder or zip and writes one `.mms` file per resource folder, such as `loot_table.mms`, using the shorthand wherever it applies. In the browser, use **Project ▾ → Import zip…**, or paste one JSON file with **Paste JSON as MMS…**.
