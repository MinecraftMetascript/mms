# tree-sitter-mms

A [tree-sitter](https://tree-sitter.github.io) grammar for Minecraft Metascript, used by the [Zed extension](../zed) for highlighting, brackets, indentation and the outline. Other tree-sitter editors (Neovim, Helix) can use it too.

The engine does not use this grammar. Its parser is ANTLR (`engine/syntax/grammar/*.g4`), which stays the definition of the language. This grammar follows it closely, and is a little more lenient where that helps an editor recover from half-typed code; for example, list elements need no separator.

`src/scanner.c` lexes raw commands: a `/` at the start of a line followed by a lowercase letter, then the command's text around `${...}` interpolations.

```sh
npm ci
npx tree-sitter generate   # after editing grammar.js; commit src/ too
npm test                   # corpus tests, plus scripts/check.sh
```

`scripts/check.sh` checks that `src/` is up to date, that every example and engine test file parses without errors (the deliberately broken `recovery/` inputs aside), and that the Zed queries compile.
