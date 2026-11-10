# AI agents

`mms mcp` is a [Model Context Protocol](https://modelcontextprotocol.io) server. It gives an AI agent that writes MMS the same engine the editors use, so the agent can check its work, see what it compiles to, run its tests and look names up without guessing.

## Setting it up

The server is part of the `mms` command line and speaks over stdin and stdout. It needs a version of `mms` newer than 0.5.0: run `mms mcp -check`, which tests the server and prints the lines below with the path of your own `mms` filled in. If it says `unknown command "mcp"`, the `mms` it found is too old.

In Claude Code, for yourself in one project folder:

```
claude mcp add mms -- mms mcp
```

To share it with everyone who works on a project, put a `.mcp.json` file at the project's root instead:

```
{
  "mcpServers": {
    "mms": { "command": "mms", "args": ["mcp"] }
  }
}
```

Claude Code does not start a server from `.mcp.json` until you approve it. It asks the next time you start `claude` in that folder, and until you answer the server is not loaded.

Claude Desktop, Cursor and other clients take the same `mcpServers` entry in their own settings file.

The server has no settings and does not use the network. Everything it knows is inside the `mms` binary, including these docs and the vanilla data of the Minecraft version MMS targets.

## Checking that it is connected

`claude mcp list` starts each server and says how it went. Inside Claude Code, `/mcp` shows the same and lists the tools.

| `claude mcp list` says | What it means |
| --- | --- |
| `Connected` | It works. The agent has the tools below. |
| `Pending approval` | The entry is in `.mcp.json` and nobody has approved it yet. Start `claude` in the folder and accept the server. |
| `Failed to connect`, `Connection closed` | `mms` started and quit, which is what a version without `mcp` does. Run `mms mcp -check`. |
| `Failed to connect`, `ENOENT` | The client did not find `mms`. Give the full path as the command; `mms mcp -check` prints it. |
| no `mms` line | The client did not read the entry. Check the file's name and place, and that it is valid JSON. |

`mms mcp` on its own waits for a client and prints nothing else, so use `mms mcp -check` to test it by hand.

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

`fix` and `fmt` with `dir` are the only calls that write files, and they write only the project's own source files.

The pages are also MCP resources, as `mms://docs/<page>`, and so is the prelude's source, as `mms://prelude/<file>`.

## Without MCP

An agent that can run commands can use `mms check`, `mms test` and `mms fix` directly. For reading, [llms.txt](https://minecraftmetascript.github.io/mms/llms.txt) links every page and [llms-full.txt](https://minecraftmetascript.github.io/mms/llms-full.txt) is the docs and the whole reference in one file.
