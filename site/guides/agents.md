# AI agents

`mms mcp` is a [Model Context Protocol](https://modelcontextprotocol.io) server. It gives an AI agent that writes MMS the same engine the editors use, so the agent can check its work, see what it compiles to, run its tests and look names up without guessing.

## Setting it up

The server is part of the `mms` command line and speaks over stdin and stdout. It needs `mms` 0.5.1 or newer: run `mms mcp -check`, which tests the server and prints the lines below. Its `claude mcp add` line has the path of your own `mms` filled in; the `.mcp.json` entry keeps plain `mms`, which everyone who shares the file needs on their PATH. If it says `unknown command "mcp"`, the `mms` it found is too old.

In Claude Code, for yourself in one project folder:

```
claude mcp add mms -- mms mcp
```

To share it with everyone who works on a project, put a `.mcp.json` file at the project's root instead. `mms init` writes this one:

```
{
  "mcpServers": {
    "mms": { "command": "mms", "args": ["mcp"] },
    "nbt": { "command": "nbt-mcp", "args": ["-dir", "structures"] }
  }
}
```

The `nbt` entry is [nbt-mcp](https://github.com/MinecraftMetascript/nbt-go/tree/main/cmd/nbt-mcp), a separate server for building structure files. With `-dir structures` it saves `halls/hall.nbt` where `Piece("halls/hall")` looks for it. Install it with `go install github.com/MinecraftMetascript/nbt-go/cmd/nbt-mcp@latest`, or from the nbt-go releases page; remove the entry if you don't build structures. If the folder already has a `.mcp.json`, `mms init` adds whichever of the two servers it does not name and leaves the rest as it was.

Claude Code does not start a server from `.mcp.json` until you approve it. It asks the next time you start `claude` in that folder, and until you answer the server is not loaded.

Claude Desktop, Cursor and other clients take the same `mcpServers` entry in their own settings file.

The server has no settings and does not use the network. Everything it knows is inside the `mms` binary, including these docs and the vanilla data of the Minecraft version MMS targets.

The three [preview tools](#preview-tools) also need a JavaScript runtime: Node.js 20 or later, Bun or Deno, found on `PATH` (or set `MMS_JS` to its path). `mms mcp -check` says which one it found. The other tools work without one.

## Checking that it is connected

`claude mcp list` starts each server and says how it went. Inside Claude Code, `/mcp` shows the same and lists the tools.

| `claude mcp list` says | What it means |
| --- | --- |
| `Connected` | It works. The agent has the tools below. |
| `Pending approval` | The entry is in `.mcp.json` and nobody has approved it yet. Start `claude` in the folder and accept the server. |
| `Failed to connect`, `Connection closed` | `mms` started and quit, which is what a version without `mcp` does. Run `mms mcp -check`. |
| `Failed to connect`, `ENOENT` | The client did not find `mms`. Give the full path as the command; `mms mcp -check` prints it. |
| no `mms` line | The client did not read the entry. Check the file's name and place, and that it is valid JSON. |

`mms mcp` on its own waits for a client. It writes one line to stderr and nothing to stdout, so use `mms mcp -check` to test it by hand.

## The tools

The tools that read a project take either `dir`, the project folder on disk, or `source`, the text of one `.mms` file checked on its own. An agent without file access can use `source` for everything.

| Tool | What it does |
| --- | --- |
| `check` | Lists every problem with its place, the source line and the fixes MMS offers. `hints` adds style suggestions. |
| `explain` | Lists the files the project writes. With `name`, prints the JSON or `.mcfunction` files of one declaration, and what its Blueprint instance or Template calls expand to. |
| `test` | Runs the project's [`Test` declarations](functions.md#tests) in the interpreter and reports each failing `expect`. `run` keeps the tests whose names match a regular expression. |
| `fix` | Applies every safe [lint](lint.md) fix. With `dir` it writes the files; with `source` it returns the text. |
| `fmt` | Formats, in place for `dir` or returning the text for `source`. |
| `convert_json` | Turns the JSON of one pack file into the MMS declaration that writes it. |
| `lookup` | Searches the reference: kinds, helpers, the prelude's Templates, Blueprints and Types, and lint rules. An exact name gives the whole entry. |
| `vanilla` | Lists the game's registries and ids, a block's states, and the files of the vanilla data pack. |
| `docs` | Lists these pages, returns one, or searches them. |
| `biomes`, `density`, `jigsaw` | Answer what the editors' previews show, as text. See [Preview tools](#preview-tools). |

`fix` and `fmt` with `dir` are the only calls that write files, and they write only the project's own source files.

### Preview tools

An agent can't see the editors' preview pane, so these tools describe it in words and numbers. They run the same code the editors' previews do, so their numbers match the previews.

| Tool | What it does |
| --- | --- |
| `biomes` | The text form of the [biome map](editors.md#why-a-biome-doesnt-appear). Samples a Dimension's biome source over a square (`x`, `z`, `radius`, at height `y`) and lists the biomes in it by share. It gives the climate at the centre and the nearest multi-noise entries, with the parameters each one misses. With `biome`, it says where that biome wins. If it wins nowhere, it says where it comes closest, what beats it there, and which climate parameters never reach its ranges. |
| `density` | Samples a noise, a density function, the final density of noise settings, or a terrain. `mode` `area` gives the values over a square at one height. `column` gives the values down one column and where they cross 0 (for a final density, the ground and cave edges). `surface` gives the top block of each column, with surface rules and aquifers applied. Vanilla ids such as `minecraft:overworld` work too. |
| `jigsaw` | Grows a JigsawStructure with the [simulator](resources.md#simulating-a-jigsaw-structure) the editors use. One run lists every piece and each jigsaw block that placed nothing, with the reason. With `runs` it covers many seeds and says how often each piece appears, and which pieces the pools list but no run placed. |

All three take a `seed`. The rendered images are not part of these tools; an agent gets the numbers behind them.

The pages are also MCP resources, as `mms://docs/<page>`, and so is the prelude's source, as `mms://prelude/<file>`.

## Without MCP

An agent that can run commands can use `mms check`, `mms test` and `mms fix` directly. For reading, [llms.txt](https://minecraftmetascript.github.io/mms/llms.txt) links every page and [llms-full.txt](https://minecraftmetascript.github.io/mms/llms-full.txt) is the docs and the whole reference in one file.
