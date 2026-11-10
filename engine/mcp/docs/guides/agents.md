# AI agents

`mms mcp` is a [Model Context Protocol](https://modelcontextprotocol.io) server. It gives an AI agent that writes MMS the same engine the editors use, so the agent can check its work, see what it compiles to, run its tests and look names up without guessing.

## Setting it up

The server is part of the `mms` command line and speaks over stdin and stdout. Any MCP client can start it. In Claude Code:

```
claude mcp add mms -- mms mcp
```

For a client that reads a JSON file of servers, such as Claude Desktop or Cursor:

```
{
  "mcpServers": {
    "mms": { "command": "mms", "args": ["mcp"] }
  }
}
```

The server has no settings and does not use the network. Everything it knows is inside the `mms` binary, including these docs and the vanilla data of the Minecraft version MMS targets.

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
