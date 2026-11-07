# Tech debt audit

Audit of `main` at b469b74 (the overhaul, 16da2fe, plus #213 and #215 to #220), run on 2026-11-07. Every finding cites the line it is about. Findings marked **(repro)** were reproduced with a CLI built from this commit or a scratch test, not only read.

## Executive summary

1. **The function compiler miscompiles ordinary code (repro).** User locals, compiler temps and shadowed locals share one score-holder namespace, so `let _t0 = 5` is overwritten by a temp and `let x_1` collides with a shadowed `x`. `mms check` reports nothing. (F01)
2. **`mms build` can write and delete outside the build folder (repro).** A declaration named `a/../../evil`, an `"output": "../x"` or `"."` in mms.json, or a namespace of `..` with `-world` all escape. One case deleted a world's `data/raids.dat`. A cloned project is enough to trigger it. (F02 to F04)
3. **The web editor renders doc comments as unsanitised HTML.** `LSPClient` has no `sanitizeHTML`, so `/// <img onerror=...>` in an imported project runs script in the editor origin, which holds every saved project. (F05)
4. **The editor and the CLI checker disagree with the formatter on where declarations start (repro).** The incremental chunker splits at any `word word` line, so `else 2` on its own line passes `mms fmt` and fails in the editor and `mms check`. That breaks the whitespace-agnostic goal. (F09)
5. **Each keystroke redoes most of the work.** Blueprint instances (every CustomItem, Book and JigsawStructure) are expanded and checked again with no reuse (400 CustomItems cost 158 ms natively for a comment edit). Typed consts and instances are type-checked two or three times. VS Code runs Fix All on every lightbulb request and re-reads the Minecraft world from disk on every pack refresh. (F20 to F27)
6. **The `mms/*` LSP payloads are `map[string]any` in Go and hand-copied three times in TypeScript, and the copies have drifted.** VS Code drops `texts`, so the chat/title preview from W5 never appears there. (F30, F31)
7. **Two of the "one way" rules don't hold yet.** `MaterialRule` and `SurfaceRule` both resolve (and the importer and error messages print the undocumented one), `Shaped` and `CraftingShaped` are both valid, list commas are optional, and two mcfunction simulators (Go `mcfn` and TS cmdlib) can disagree about the same pack. (F40 to F44)
8. **Quick fixes and lints can break files.** The "quotes are not needed" fix strips quotes from `setup`, `run`, `check` and other keywords that can't be written bare, and the missing-field fix parses the field names out of an English error message. (F07, F33)
9. **Web project data loss.** Export then import drops every `.mmsbook` and `.mmstext` file, and IndexedDB save failures (including quota) are swallowed while the UI says "saved". (F11, F12)
10. **CI runs only `gofmt`, `go vet` and the tests.** No staticcheck, unused or ineffassign, which is how an ignored error in `mms import` (F08), seven unused helpers and the dead `fn` coalesce call (F55) reached `main`.

## Architectural mental model

MMS is a Go engine with three hosts. The CLI (`engine/cmd/mms`), the LSP over stdio, and the same LSP compiled to WASM (`engine/cmd/mms-wasm`) all run one pipeline:

`syntax` (ANTLR grammar, a token-rewriting recovery pass in `recover.go`, a hand-written AST builder, and incremental chunk reparsing) → `eval` (the prelude first, then the user files; consts, `for`, Templates and Blueprints expanded into plain `syntax.Decl`s placed at the user's call span) → `types` (turns MMS type expressions into mcdoc `spec.Type`s and calls `sema` to check each Blueprint field, Template argument and typed const) → `sema` (the real checker: walks syntax against mcdoc views, lowers to JSON, records editor `Info`; it also holds the function compiler front end, the text markup parser, and the kind and symbol tables) → `fn` (function IR, optimiser, lowering to mcfunction) → `project` (memoised `Analyze`, per-declaration reuse, lints, emit, tests through `internal/mcfn`, world packs) → `ide` and `lsp` (completion, hover and the rest, built fresh per request from `Result`).

`web/` is a SvelteKit app. Its `Engine` drives the WASM LSP through `@codemirror/lsp-client` and, after each diagnostics publish, pulls the whole generated pack (`mms/output`), test results, and document symbols for every file. `packages/preview` is a source-only Svelte library of previews (Deepslate worldgen worker, jigsaw, books, chat, function runs on vendored cmdlib), shared by the web editor and the VS Code webviews. `editors/vscode` is a thin client around `mms lsp`, plus World, Structures and Jigsaw views that only it has.

Two places where this differs from the docs:

- `docs/tasks.md` says W3 is done when "`sema/check.go` and its helpers are deleted", W4 when "the hardcoded kind map ... is gone", and W5 adds "a new `text` package". In the code, `types` imports `sema` and delegates all checking to it (`engine/types/check.go:101`), `sema/check.go` is 1,144 lines and the core of checking, kind behaviour is still decided by about 47 string comparisons such as `k.MMS == "Function"` across sema, project and ide, and markup lives in `sema/bookmarkup.go`, `book.go`, `text.go` and `tell.go`. "types resolves, sema checks" may well be the right final design, but the task list says otherwise. (F60)
- `engine/eval/eval.go` says provenance serves diagnostics, go to definition and previews. Only eval's own tests read `Prov.Def`, `Frames` or the values of `Result.Nodes`. (F51)

## Tooling run

| Tool | Result |
| --- | --- |
| `go vet` (excluding `syntax/parser`) | clean |
| `golangci-lint` 2.5 (standard + staticcheck, unused, errcheck, ineffassign) | 7 unused, 2 ineffassign (one is a real bug, F08), 16 staticcheck style notes, 184 errcheck (about 150 are `fmt.Fprint*` in the CLI and are noise; the real ones are listed below) |
| `deadcode` (golang.org/x/tools) | 8 unreachable functions; 27 reachable only from tests |
| `go test -coverpkg=./...` | 83.7% total. Lowest: `cmd/mms` 54%, `jsonv` 51%, `syntax` 71% (excluding generated code). `fn` reads 80% but has no tests of its own (F70) |
| `govulncheck` | **not run**: vuln.go.dev is blocked by this environment's network policy |
| `npm audit` | web: fflate 0.8.2 moderate (GHSA-px8p-9vwx-vf98, used on user zips), 4 low via deepmerge in the Svelte toolchain. VS Code: glob and minimatch moderate via `@vscode/vsce` (dev only). preview: 0 |
| `svelte-check` / `tsc --noEmit` | preview and VS Code clean. web: 3 errors, all from generated files missing in a fresh checkout (`reference.json`, `third-party-licenses.txt`), not real |
| `madge --circular` | web: 2 cycles, both type-only (`import type`). preview: cycles only inside vendored cmdlib |
| `knip`, `depcheck` | not run; unused dependencies were checked by hand (none found apart from F86) |

## Findings

Severity: **Critical** is a bug that loses data, miscompiles or is exploitable now. **High** is a real bug or a large cost on the editor hot path. **Medium** raises the cost of change or causes visible DX problems. **Low** is cleanup.

### Correctness and security

| ID | Category | File:Line | Severity | Effort | Description | Recommendation |
| --- | --- | --- | --- | --- | --- | --- |
| F01 | Correctness | engine/sema/fn.go:369-373, engine/fn/gen.go:28, engine/fn/opt.go:19 | Critical | S | **(repro)** A user local is `#f.<name>`, a temp is `#f._t<n>`, a flag `#f._f<n>`, and a shadowed local `<name>_<k>`. `let _t0 = 5; let y = a*2 + a*3; return _t0 + y` emits `set #f._t0 5`, then overwrites `#f._t0` with `a*2`. `let x_1 = 7` plus a shadowed `x` makes the inner `x` write `#g.x_1`, so `g()` returns 4, not 9. `isTemp` (`Contains("._t")`) also treats user names containing `._t` as temps. String temps have the same problem: runtime concatenation writes `_s<n>` in local storage (engine/sema/fn.go:1343), which a string local named `_s0` shares. | Put generated and shadow names, both score holders and storage paths, behind a character no identifier can contain (Top 5, #1). |
| F02 | Security / data loss | engine/project/emit.go:310-322 | Critical | S | **(repro)** `WriteDir` joins `f.Path` without a containment check. `Function a/../../../../../evil { /say hi }` passes `mms check`, and `mms build` writes `<project>/evil.mcfunction`. The same paths go into `Zip`, a zip-slip for whoever unpacks the download. | Reject `..` segments in resource names in sema (as `CheckDimensionID` does, engine/project/dimreset.go:368). In `WriteDir` and `Zip`, convert each path with `filepath.Localize` (which, unlike `fs.ValidPath`, rejects `\` on Windows) and refuse what fails. |
| F03 | Security / data loss | engine/project/config.go:66-84 | Critical | S | **(repro)** `Validate` never checks `output` or `resourcePack.output`. `"output": "."` deleted the project's own `data/`; `"../victim"` wipes `../victim/data` and `assets`. | Require `filepath.IsLocal` and not `.` for both outputs. Before `WriteDir` removes anything, resolve the target with `filepath.EvalSymlinks` and check it is still inside the project, since a cloned project can ship `build` as a symlink to `../victim`. Keep the CLI `-out` flag unrestricted, since that is an explicit choice. |
| F04 | Security / data loss | engine/project/config.go:63, engine/cmd/mms/build.go:189 | High | S | **(repro)** `nsPattern` accepts `.` and `..`. With namespace `..`, `mms build -world w` resolves the datapack folder to `w` itself, and `WriteDir` removed `w/data/raids.dat` (world saved data). | `^[a-z0-9_][a-z0-9_.\-]*$`, plus `filepath.IsLocal(ns)` before the join. |
| F05 | Security | web/src/lib/engine/engine.svelte.ts:184 | High | S | `LSPClient` is built without `sanitizeHTML`; lsp-client assigns rendered Markdown to `innerHTML` (`@codemirror/lsp-client/dist/index.js:105`), and `engine/syntax/doc.go:161` passes doc text through as is. A doc comment with `<img onerror>` in an imported or example project runs script with access to IndexedDB. `javascript:` links in docs are not intercepted either (web/src/lib/editor/Editor.svelte:223). | Pass a `sanitizeHTML` (DOMPurify, links limited to http(s), `#` and relative) (Top 5, #3). |
| F06 | Correctness | engine/fn/gen.go:440-442 | High | S | **(repro)** `fn.QuoteSNBT` escapes only `\` and `"`. `let s = "a\nb"` emits a command split over two mcfunction lines, which breaks the function. `sema.quoteSNBT` (engine/sema/customitem.go:119) already escapes control characters, and `quoteText` (engine/sema/tell.go:257) escapes a third set. | One quoter: move `sema.quoteSNBT`'s body into `fn.QuoteSNBT`, use it from sema and tell, and add a newline test. |
| F07 | Correctness (quick fix) | engine/sema/lint.go:17 | High | S | **(repro)** `reserved` lists 11 keywords; the lexer has more (`const`, `extern`, `primary`, `as`, `check`, `emit`, `setup`, `expect`, `run`, `Type`, MMSLexer.g4:16-41) that can't be written bare. The safe fix for `quoted-name` turns `FunctionTag t = ["setup"]` into a parse error. `engine/importer/importer.go:17` keeps a second copy with the same gap. | Add `syntax.IsReserved`, derived from the lexer's literal names, used by lint and importer, with a test that every reserved word fails to parse bare. |
| F08 | Error handling | engine/cmd/mms/import.go:70-76 | High | S | **(repro)** The `err` from `MkdirAll`/`WriteFile` is shadowed inside the `if`, line 73 checks the outer `err`, and `mms import` prints "imported" after writing nothing (ineffassign flags :71). | `if err := ...; err != nil { fail }` for each call. |
| F09 | Correctness / whitespace | engine/syntax/chunks.go:162 | High | S | **(repro)** `startsDecl` treats any `word word` at column 0 as a new declaration. `const x: int = if (flag) 1` followed by `else 2` on the next line is valid for `syntax.Parse` and `mms fmt` but an error under `Reparse`, which `project.SetFile` and so the editor and `mms check` use. The fuzz test excuses exactly this case (engine/syntax/fuzz_test.go:46). | Split only where `startsTopItem` (engine/syntax/recover.go:141) would; fall back to the whole-file parse when a chunk has errors and the file doesn't (Top 5, #4). |
| F10 | Correctness | engine/syntax/parse.go:1003-1011 | Medium | S | `unquote` uses Go's `strconv.Unquote`; any escape Go doesn't know (`\{`, `\$`) makes it fail, and the fallback `strings.Trim(s, "\"'")` leaves `\t` literal and strips quote characters inside the string. `"'quoted' \{x}"` becomes `quoted' \{x}`. | An MMS unescaper: resolve `\" \' \\ \n \t`, keep `\{` and `\$` for the markup layer, with table tests. |
| F11 | Data loss | web/src/lib/project/zip.ts:50 | High | S | `readZip` keeps only `.mms`, `.json`, `.mcmeta`, `.mcfunction`, so export then import of a project silently drops every `.mmsbook` and `.mmstext` file (the `library` example breaks). `readProjectZip` (:79) was updated to `isSourcePath` but this filter wasn't. | Filter with `isSourcePath(rel) \|\| /\.(json\|mcmeta\|mcfunction)$/` and add a round-trip test. |
| F12 | Error handling | web/src/lib/project/project.svelte.ts:264-270 | High | S | `flush()` swallows every save error, `QuotaExceededError` included (likely, since assets are stored as base64), and then records `updated: Date.now()`, so the switcher says the project was just saved. | `saveError` state shown in the header; update the meta only after a successful write. |
| F13 | Error handling | web/src/lib/components/ProjectsDialog.svelte:27-35 | Medium | S | `run()` has `try/finally` and no catch; open, rename, duplicate and delete don't catch IndexedDB errors, so private windows get unhandled rejections and no feedback. | Catch in `run()` and show the error. |
| F14 | Security | web/src/lib/project/zip.ts:37-43 | Medium | S | Zip entry paths aren't normalised: `../x.mms`, `/x.mms` and `a\b.mms` are stored and exported again. Three path normalisers disagree (`project.svelte.ts:236`, `:274`, and none here). | One `safeProjectPath()` used for import, add, rename and `setAsset`. |
| F15 | Dependency | web/package.json (fflate ^0.8.2), web/src/lib/project/zip.ts:38 | Medium | S | fflate 0.8.2 has GHSA-px8p-9vwx-vf98 (infinite loop on malformed ZIP64) and runs on user zips on the main thread, twice per import (`web/src/routes/+page.svelte:250`, `:266`), with no size limit. | Upgrade fflate; unzip once with a `filter` that skips oversized entries. |
| F16 | Correctness | engine/lsp/workspace.go:31-46 | Medium | S | `contains` ignores `Config.Sources`, so a `.mms` file that `sources` excludes joins the project when it changes on disk, and editor diagnostics stop matching `mms build`. | Match against the source globs (reuse `globRegexp`). |
| F17 | Correctness | engine/lsp/workspace.go:67-68 | Low | S | After an mms.json reload, `discover` builds new workspaces but never clears `s.assetDiags` (keyed by the old pointers), so stale unused-asset diagnostics stay. | Clear published asset diagnostics before `discover`. |
| F18 | Correctness (import) | engine/importer/pack.go:211 | Medium | S | **(repro)** The `data/<ns>` folder name is used as the namespace without validation: `data/My Pack/` imports "successfully" and writes an mms.json that `mms check` then rejects. | Skip invalid namespaces into `Skipped` with a reason; call `cfg.Validate()` before writing. |
| F90 | Error handling | engine/cmd/mms/fmt.go:44 | Medium | S | `runFmt` ignores the error from `io.WriteString(stdout, out)` and exits 0, so a script redirecting `mms fmt` output to a full disk accepts a truncated file. This is the formatted output itself, unlike the diagnostic prints listed as noise. | Return 1 when the write fails. |
| F19 | Correctness | engine/cmd/mms/build.go:308-311 | Medium | S | `runInit` ignores the `WriteFile(main.mms)` error and says "created mms.json and main.mms" even when main.mms already existed. 0% coverage. | Check the error, say "kept existing main.mms", add a test. |

### Editor hot path and performance

| ID | Category | File:Line | Severity | Effort | Description | Recommendation |
| --- | --- | --- | --- | --- | --- | --- |
| F20 | Performance | engine/project/expand.go:85-116 | High | M | **(repro)** Every Blueprint instance is expanded and checked again (`CheckType`, `ConvertedStack`) on each `Analyze`, with no reuse. 400 small CustomItems cost 158 ms natively for a comment-only edit (100 cost 38 ms; the 2,000-line DensityFunction bench is 10 ms). WASM is several times slower. | Reuse instance Decls the way `checkDecl` does (`inst.D`, `moved`, `DepsHold`); add a Blueprint workload to `engine/ide/bench_test.go`. |
| F21 | Performance | engine/eval/stmt.go:307, engine/eval/eval.go:279 | Medium | M | `eval.Run` runs on the whole project each time and rebuilds every declaration that touches a const, Template or Blueprint as new pointers, so `project.reuse` (keyed by pointer) misses them all. `emittedBefore` also runs `format.File` per emitted declaration per run. | Memoise expansion output per source Decl in `eval.Cache`, keyed by a fingerprint of what it read. |
| F22 | Performance / consistency | engine/types/check.go:84, engine/project/project.go:474-483, engine/project/expand.go:101, :141 | High | M | Typed consts, instances and value Templates are checked two or three times per analysis: once by `types.Env.Check` on the lowered value for diagnostics, then again by project on the written value for `Info`, whose diagnostics are dropped. The trees differ, so hover and errors can disagree, and Book markup diagnostics come only from the second run (`c.BookDiags`). | Return the per-site `Checker` from `Env.Check` and reuse its `Info` and diagnostics in project. |
| F23 | Performance / correctness | engine/lsp/actions.go:90-95 | High | S | `source.fixAll.mms` is computed whenever `context.only` is empty, which is VS Code's automatic lightbulb request on each cursor move: a full `Analyze`, up to 16 rounds, run on the live project and then restored with `SetFile(orig)`, which leaves `p.last` and `p.moved` describing the fixed text. | Offer it only when `only` names `source.fixAll` (the web already asks that way, engine.svelte.ts:349), and run it on `w.proj.Fork()`. |
| F24 | Performance | editors/vscode/src/world.ts:128, editors/vscode/src/structures.ts:25-80 | High | S | `output.onDidChange` (about 150 ms after each keystroke) runs `redraw`, which walks `<world>/generated`, reads and hashes every `.nbt` twice, gunzips `level.dat` for both worlds and re-reads mms.json per structure, on the extension host thread, even when the views are hidden. | Recompute the world part only from file watchers; cache hashes by mtime; skip when no view is visible. |
| F25 | Performance | engine/lsp/server.go:998, engine/project/testrun.go:46, web/src/lib/engine/engine.svelte.ts:255-291 | Medium | M | After each analysis the web pulls `mms/output` (a full uncached `Emit`, every file's text in one JSON string across the worker boundary), then `mms/runTests` (`HasErrors` and `emit(true)` again), then `documentSymbol` for every file. | Memoise `Emit` on `Result`; return a generation number so clients skip unchanged output; fetch symbols only for changed files. |
| F26 | Performance | packages/preview/src/preview/Preview.svelte:277-281, packages/preview/src/preview/preview.worker.ts:1429 | Medium | S | Any output change, including an edit to an unrelated function, re-parses every previewable file and posts it to the worker, which bumps `version` unconditionally and throws away the iso and terrain caches. | Post `setup` only when a signature of the files changes; bump `version` only on a real change, as `shapeVersion` does. |
| F27 | Performance | editors/vscode/src/preview.ts:35-38, :84-115, :208 | Medium | S | Each pack refresh re-posts the whole pack plus every texture and structure as base64 data URLs to the panel. | Send assets only when the `(path, mtime)` set changes, in their own message. |
| F28 | Performance | engine/project/project.go:255-329, engine/lsp/server.go:552-558 | Medium | S | `Diagnostics(path)` walks all of `r.Order` and `r.Instances` per file, and `analyze` calls it for every unit and republishes every file each keystroke (O(files × decls)). `HasErrors` repeats the walk. | Build per-path diagnostics once, lazily; publish only files whose diagnostics changed. |
| F29 | Performance | web/src/lib/project/project.svelte.ts:110-118, :252 | Medium | M | Each debounced save rewrites the whole project, base64 assets included, to IndexedDB. | Store assets under their own keys and write them only from `setAsset` and `removeAsset`. |

### Contracts and error shapes

| ID | Category | File:Line | Severity | Effort | Description | Recommendation |
| --- | --- | --- | --- | --- | --- | --- |
| F30 | Type & contract | engine/lsp/server.go:985-1178, editors/vscode/src/messages.ts:1-51, web/src/lib/engine/engine.svelte.ts:41-125, packages/preview/src/types.ts:1-82 | Medium | M | `mms/output`, `mms/runTests`, `mms/expand` and `mms/importPack` are built as `map[string]any` and typed by hand three times in TS. Drift so far: preview's `OutputFile` lacks `pack`, VS Code's `OutputItem` lacks `snbt`, web lacks `jigsawHoles`, three names for one range type, `TestOutcome` defined twice. messages.ts:1-3 says the Node build "cannot import" preview, but :179 does. | Go structs in `engine/lsp/protocol.go`; one type-only `packages/preview/src/protocol.ts` imported by both hosts (Top 5, #5). |
| F31 | Editor parity | editors/vscode/src/output.ts:73-85, editors/vscode/webview/App.svelte:63-76 | Medium | S | The server sends `texts`, `PackOutput` drops it, so the tell/title/actionbar preview (W5) exists only in the web editor. | Carry `texts` to `PreviewPane`; add lenses. Follows from F30. |
| F32 | Error handling | engine/lsp/server.go:473-480, engine/lsp/workspace.go:117-122 | Medium | S | `mms/setConfig` is a notification, so a `Validate` failure goes nowhere and the old config stays. An invalid mms.json on disk is only logged, then `DefaultConfig()` is used silently. | Make `setConfig` a request and show the error; publish config errors as a diagnostic on mms.json. |
| F33 | Fragile contract | engine/project/expand.go:163-167 | Medium | S | The missing-fields quick fix recovers field names by cutting the message from engine/eval/blueprint.go:99 at `" needs "`. A rewording silently drops the fix. `mms/missing-field` is also reused for unrelated problems (engine/sema/shorthand.go:220, engine/sema/other.go:132). | Carry the field names in a structured field on the diagnostic; give the other two cases their own codes. |
| F34 | Error handling | engine/lsp/server.go:134-155, engine/cmd/mms-wasm/main.go:43 | Medium | S | Panics in notifications are logged only through `s.Log`, which is nil in WASM, so a `didChange` panic in the browser vanishes and leaves stale diagnostics. | Log to the console in mms-wasm; send `window/logMessage` for notification panics. |
| F35 | Error shape | engine/eval/blueprint.go:222-232, engine/project/hints.go:125 | Medium | S | Blueprint diagnostic codes come in three shapes: `mms/<kebab>`, the bare rule name from `warn(f, "rule", msg)`, and `lint/<rule>` from `hint(...)`. Bare codes can't be configured in mms.json; the `"check"` fallback at :222 is unreachable. | One prefix; decide whether Blueprint rules are configurable (see Q3). |
| F36 | Error shape | engine/sema/tell.go:108, engine/sema/storage.go:439, engine/sema/exec.go:369, engine/sema/check.go:470 | Low | S | Argument-count problems use `missing-arg`/`too-many-args` in some places and `mms/arity` in others; unknown helpers are `mms/unknown-callable` or `mms/unknown-helper`. | Pick one set and rename. |
| F37 | Error shape | engine/sema/fn.go:313, :616-619 | Low | S | `Function f(): string` is accepted at the signature, then every `return` errors "functions return int or bool". | Reject `string` once, on the return type. |
| F38 | Validation | engine/project/config.go:95, web/src/lib/project/zip.ts:30, :75 | Low | S | mms.json is decoded without `DisallowUnknownFields`, so `"namepsace"` is silent. The web `Object.assign`s any imported mms.json into a two-field `Config`, then exports only namespace and description, losing `lint`, `icon` and `sources`. | Reject unknown keys (allowing `$schema`, `libraries`); decide whether the web keeps the full config (Q5). |
| F39 | Contract | editors/vscode/src/world.ts:47, :62, :624-637 | Medium | M | The World view's build result is scraped from human-readable CLI text: `files` sums every `(\d+) files?` on stdout, and failures match exact phrasings from engine/cmd/mms/build.go:150 and :202. A comment keeps an "older mms builds" fallback. | `mms build --world --json`; drop the fallback (the extension bundles its `mms`). |

### One way to do things

| ID | Category | File:Line | Severity | Effort | Description | Recommendation |
| --- | --- | --- | --- | --- | --- | --- |
| F40 | Consistency | engine/sema/kinds.go:63, :73, :181-184 | High | S | `MaterialRule`/`SurfaceRule` and `MaterialCondition`/`SurfaceCondition` both resolve. `ForFolder` returns the folder name, so the importer (engine/importer/pack.go:99), hoisting and the wrong-kind message say `MaterialRule` while the docs teach `SurfaceRule` (site/start/tutorial-dimension.md:12). The `NoiseSettings` alias is redundant. | Add both folders to `renamed` so only the alias names them; drop `NoiseSettings`. |
| F41 | Consistency | engine/sema/structs.go:374, :386 | Medium | S | `Shaped(...)` and `CraftingShaped(...)`, `AboveWater` and `Water`, `Pool` and `LootPool` are all valid. Completion shows only the alias. | Reject the unaliased spelling where an alias exists. |
| F42 | Consistency | engine/syntax/grammar/MMSParser.g4:196 | Medium | S | List elements take optional commas, `(COMMA? value)*`, so `[a (b)]` is a call, `[1 -2]` is `[-1]`, and `[x` / `-y]` over two lines is a subtraction. The formatter always writes commas (engine/format/format.go:737). | Require commas between list elements (Q7). |
| F43 | Dead syntax | engine/syntax/grammar/MMSParser.g4:168 | Medium | S | `->` is lexed, parsed (engine/syntax/parse.go:816) and formatted (engine/format/format.go:621), but nothing gives it a meaning; sema rejects it. | Remove it, or give it a form in prelude/forms.mms. |
| F44 | Duplication | engine/internal/mcfn/mcfn.go:1-8, packages/preview/src/run/run.ts:1-4 | Medium | L | Two mcfunction simulators run the same output: Go `mcfn` for `mms test` and LSP tests, vendored TS cmdlib for the function-run preview. They already diverge (cmdlib needed a local patch for `stack append`; mcfn rejects commands outside a short list, mcfn.go:178), so a pack can pass one and fail the other. | Decide which is authoritative (Q1). Cheapest step: a shared golden corpus run by both CI suites. |
| F45 | Duplication | engine/syntax/parse.go:337, engine/eval/eval.go:550, engine/sema/fn.go:137, engine/ide/evalnames.go:208, engine/sema/bookmarkup.go:338, engine/project/lint.go:617 | Medium | M | `${...}` is split six ways with different rules (nested or not, byte or rune offsets, regex). Rename and references disagree with eval on `${f({a = 1})}`; the lint version counts a field `.x` as a use of const `x`. | One `syntax.Pieces(s)`, parsed once in the builder as names already are. |
| F46 | Duplicated semantics | engine/sema/forms.go:174, engine/eval/blueprint.go:115 | Medium | L | `as T = value` is evaluated by eval for Blueprints, but `extern Type` forms are rewritten textually by sema's `substitute` with its own constant folding. A form conversion can't use `if`, Templates or arithmetic the way a Blueprint `as` can. | Have sema call eval with `self` bound; delete `substitute` and `fold`. |
| F47 | Duplication | engine/eval/eval.go:402, engine/ide/extracteval.go:11, engine/sema/hooks.go:26-36 | Low | S | The list of kinds whose body is function code is kept three times in Go; the prelude declares those kinds but doesn't mark them. | Mark them in core.mms (for example an `@code` tag) and export `prelude.CodeKinds()`. |
| F48 | Leftover kind map | engine/sema/hooks.go:35, engine/project/project.go:583-623 | Medium | M | Kind behaviour is decided by about 47 string literals (`k.MMS == "Function"`, 17 in sema, 15 in project, 15 in ide). An `extern Blueprint` added to core.mms that no switch handles does nothing, silently. | Typed constants or predicates, plus a test that every `prelude.Kinds()` name is handled by `checkDecl`. |
| F49 | Duplication | engine/sema/check.go:1139-1170, engine/types/types.go:265-283, engine/project/jigsaw.go:936-966, engine/sema/storage.go:394 | Low | S | Three edit-distance implementations and four "did you mean" thresholds, so suggestions differ by layer. | One `sema.Closest`. |
| F50 | Duplication | engine/project/project.go:367-370, :412-421, :432-433, :455, :477-478, :562-564, engine/sema/check.go:566, engine/ide/ide.go:93-95 | Low | S | `NewChecker` plus copying `Files`, `Texts`, `Src` is pasted 8 times; `trialErrors` and ide's `at()` forget `Texts`, so checks there behave differently from analysis. | `p.NewChecker(ns, src)` and `(*Checker).sub(...)`. |

### Architectural decay and dead code

| ID | Category | File:Line | Severity | Effort | Description | Recommendation |
| --- | --- | --- | --- | --- | --- | --- |
| F51 | Unused abstraction | engine/eval/eval.go:102, :120, :590 | Medium | M | `Prov.Def`, `Prov.Frames` and the values of `Result.Nodes` are read only by eval's tests; ide/tokens.go:98 checks membership only. Template-body diagnostics land on the call span with no "in Template X". | Either wire `Frames` into `relatedInformation` and go to definition (a DX win), or cut it to a set and delete `Def`/`Frames` (Q2). |
| F52 | Correctness / perf | engine/eval/expr.go:402, :421 | Medium | S | `object()` evaluates each spread twice (three times when it isn't a record), so Template calls in spreads are recorded twice in `Result.Calls` and nested spreads cost 2^depth. | Reduce each spread once. |
| F53 | Dead code / parity | engine/sema/book.go:101 | Medium | S | `bookMarkupDoc` is unused. At a3bea4e it was appended to the Book `text` hover; after the overhaul the hover comes from prelude/books.mms:22 without it. A hover parity regression. | Show it on `BookText` fields, or move it into books.mms and delete the constant. |
| F54 | Dead code | engine/sema/customitem.go:20-50 | Medium | S | `CustomItemOut` keeps pre-overhaul fields (`Model`, `ModelPath`, `ItemModel`, `LangKey`, `Names`) nothing sets; `Languages()` has no callers; `idType` is unused. `check.go:116` `textureName` is never assigned, so assets.go:176 never runs. | Reduce to `{Stack, Book}`; delete the rest. |
| F55 | Dead code / bug | engine/fn/opt.go:138 | Low | S | `coalesce(&Block{Ops: o.Cond.Pre}, uses)` discards its result, so `If` condition pre-ops are never coalesced; `Loop` pre isn't attempted. | Write the result back, or delete the line. |
| F56 | Dead code | engine/sema/ide.go:48, :56, :90, :399, :547, engine/sema/symbols.go:98 | Low | S | `Info.TypeName`, `Info.FieldName`, `Checker.Fits`, `Checker.TypeInfo`, `Info.TypeDoc`, `Table.InNamespace`: no callers, tests included. | Delete. |
| F57 | Dead code | engine/syntax/clone.go:29-32, engine/syntax/parse.go:996, engine/format/doc.go:43, engine/sema/shorthand.go:34, :209, engine/sema/lint.go:200, engine/spec/pack.go:94-96 | Low | S | `CloneBlock`, `CloneDecl`, `builder.sourceText`, `format.join`, `tLootCondition`, `word`, `btoi`, `spec.Embedded` are unreachable (deadcode, golangci). `Embedded`'s comment says the web serves the pack; nothing in web/ does. | Delete. |
| F58 | Dead code | engine/command/args.go:85-86, :97-98, engine/sema/kinds.go:50, engine/sema/blockstate.go:30, :218 | Low | S | Unreachable `greedy`/`message` cases (line 68 returns first), a `case "CustomItem"` in `Folder()` (CustomItem isn't an extern Blueprint), and a `"minecraft:"+path` block lookup that never matches (`Pack.Blocks` keys are bare). | Delete. |
| F59 | Naming leftover | engine/sema/shorthand.go:87, :128, engine/sema/check.go:300, :581, engine/importer/importer.go:2-3, engine/ide/bench_test.go:13, engine/ide/ide.go:937, site/guides/resources.md:11 | Low | S | `sugar.go` became `shorthand.go`, but `sugarRule`, `Checker.sugar`, `sugarCallables`, several comments and a docs heading ("**Sugar.**") still say sugar. | Rename. |
| F60 | Plan drift | docs/tasks.md:11-13 | Medium | S | W3, W4 and W5 done conditions describe deletions and a package that didn't happen (see the mental model). | Rewrite tasks.md to the real design, or file follow-ups (Q4). |
| F61 | God file | engine/sema/fn.go (1,698 lines) | Medium | M | 43 `fc` methods: signatures, statements, expressions, strings and commands. With exec.go, storage.go, tell.go and test.go the function compiler is about 3,000 lines inside the schema-checker package, sharing `*Checker` state. | Pure move into `fnsig.go`, `fnstmt.go`, `fnexpr.go`, `fnstr.go`. |
| F62 | God file | packages/preview/src/preview/preview.worker.ts (1,467 lines) | Medium | M | Vanilla registry fetch, strip terrain, 3D chunk building, the iso painter, the helper-worker pool, view maths and render dispatch, coupled through module-level caches. | Split along those seams; the entry only routes messages. |
| F63 | Consistency | packages/preview/src/preview/preview.worker.ts:1111-1133, packages/preview/src/preview/Preview.svelte:412-456 | Medium | M | Two `planeOf` camera mappings: the worker samples pixel centres (`px + 0.5`), the component doesn't, so hover readouts can be half a pixel off what was drawn. | One pure `plane()` in `protocol.ts`, used by both, with a round-trip test. |
| F64 | God file | web/src/routes/+page.svelte (528 lines) | Medium | M | Engine lifecycle, theme, output-under-cursor rules, pinning, expansion, zip import/export, asset data URLs, status and an inline tab. The view-waiter logic is duplicated (:62, :203) and never rejects. | Move selection, import/export and the Expanded tab out; one `awaitView`. |
| F65 | God function | engine/lsp/server.go:985-1178, engine/project/project.go:337-503 | Low | M | `output` is 194 lines with seven dead `l == nil` checks; `Analyze` is 167 lines mixing pipeline, reuse policy and four editor passes, which hides F20. | Per-section builders on typed structs (F30); move editor passes next to `instanceDecls`. |
| F66 | Leftover | web/src/lib/project/store.ts:25, :50-57, web/src/lib/docs/site.ts:74-80, web/src/routes/docs/reference/+page.server.ts:11-16, web/src/lib/index.ts:1 | Low | S | A migration from the pre-index `mms-project` layout, two redirect maps for old docs URLs, and the SvelteKit scaffold `index.ts` with an unused `#lib` alias. No users yet. | Delete (but see Q6 for the docs URLs). |
| F67 | Leftover | engine/lsp/server.go:492-493 | Low | S | `mms/setFiles` is described as how the browser sends the project; the web uses `didOpen` per file and only tests call it. | Delete it, or switch the web to it (cheaper than N `didOpen`s). |
| F68 | Dead code | editors/vscode/src/jigsawGive.ts:1-3 | Low | S | A re-export shim of `@mms/preview/jigsaw/give`, which messages.ts:179 already imports directly. | Delete the shim. |
| F69 | Leftover fixtures | engine/syntax/testdata/sketch.mms:14, engine/format/testdata/sketch.in.mms | Low | S | Pre-overhaul syntax (`.Entry().When()`, `Shaped(...) -> x`, `Score(...)`); `mms check` gives 10 errors, yet it pins the parser outline, formatter goldens and fuzz seeds. | Replace with an overhaul-era sample. |

### Test debt

| ID | Category | File:Line | Severity | Effort | Description | Recommendation |
| --- | --- | --- | --- | --- | --- | --- |
| F70 | Test debt | engine/fn/ (no `_test.go`) | Medium | M | The optimiser and codegen have no unit tests; their 80% coverage comes from end-to-end runs, which missed F01, F06 and F55. | `fn/opt_test.go` on hand-built IR, plus a VM property test comparing optimised and unoptimised output. |
| F71 | Test debt | engine/sema/sema_test.go:24 | Medium | M | `TestGolden` runs `sema.Collect` and `Checker.Resource` straight on parsed syntax, skipping eval, prelude conversions and `types`, so it covers a path production doesn't use. | Drive the goldens through `project.Analyze`. |
| F72 | Test debt | engine/syntax/fuzz_test.go:46, engine/syntax/whitespace_test.go:18 | Medium | S | The fuzz test exempts the Reparse/Parse mismatch that is F09; the whitespace test only checks `Parse`, and only adds whitespace, so it misses `total/2` lexing as a resource location and `a:int` as `NS:PATH`. | Drop the exemption after F09; add Reparse and whitespace-removal variants, with a "write `a / b`" diagnostic. |
| F73 | Test debt | engine/lsp/server.go:473, :722, :820, :902 | Medium | S | 0% coverage at the LSP layer for `setConfig`, `signatureHelp`, `prepareRename` and `inlayHint`, all parity-table features. | Round-trip tests in `server_test.go`. |
| F74 | Test debt | web/playwright.config.ts:5-20, web/tests/smoke.spec.ts | Medium | M | 36 tests in one file (the most-churned file in the repo), run serially; each cold-boots WASM; the goto-and-wait-for-"No problems" preamble is repeated about 34 times; fixed sleeps at :39, :68 (800 ms past a 300 ms debounce), :128, :686. No unit tests for `zip.ts`, `store.ts`, `signature.ts`, `site.ts`. | `fullyParallel`, an `editor` fixture, split by area, poll instead of sleeping; unit tests for the pure modules. |
| F75 | Test debt | packages/preview/package.json:21-23, editors/vscode/src/world.ts:51-70 | Medium | M | The shared package has no tests at all (jigsaw simulation, graph, give command, book flattening, path classification), and VS Code's exported pure helpers in world.ts and structures.ts have none either. | vitest in preview; a Node-only unit file in VS Code. |
| F76 | Test debt | engine/cmd/mms/build.go:91, :278 | Low | S | `runCheck`, `runInit`, `sanitizeNS`, `runReference`, `runLSP` at 0%. | Table tests; the `runX(args, out, err)` signatures already allow it. |

### Dependencies, config and docs

| ID | Category | File:Line | Severity | Effort | Description | Recommendation |
| --- | --- | --- | --- | --- | --- | --- |
| F80 | CI | .github/workflows/engine.yml:23-26 | Medium | S | Only gofmt, vet and tests. | Add staticcheck (or golangci-lint with unused and ineffassign, and errcheck excluding `fmt.Fprint*`), excluding `syntax/parser`. |
| F81 | Dependency drift | editors/vscode/package.json:423-435, web/package.json:21-28, packages/preview/package.json:25-32 | Medium | S | TypeScript 5.9.3 in VS Code vs ^6.0.3 elsewhere, though VS Code's svelte-check type-checks the preview sources; `@types/node` 20 vs 22; exact pins in VS Code vs carets elsewhere; three lockfiles, no workspace. | Align TS now; pick one pinning policy or a root npm workspace (Q8). |
| F82 | Repo hygiene | engine/syntax/.cache/antlr-4.13.1-complete.jar, engine/go.mod:5-8 | Low | S | A 2 MB jar is committed although `gen.sh` downloads it and `.gitignore` covers only the specgen cache. `go.mod` marks the direct `antlr4-go/antlr/v4` dependency `// indirect` (`go mod tidy` changes it). | Untrack the jar and ignore the folder; commit `go mod tidy`. |
| F83 | Docs drift | site/guides/lint.md:5-19 | Medium | S | The table lists 13 rules; `project.Rules` has 17 (missing `jigsaw-facing`, `jigsaw-fit`, `jigsaw-overlap`, `jigsaw-unknown-pool`, `jigsaw-unreachable-target`). It lists `custom-item-as-base`, a prelude `hint(...)` and not a rule, so the documented `"lint": {"custom-item-as-base": "off"}` fails with "unknown rule". `book-page-overflow` is described differently than in engine/project/lint.go:28. | Link to the generated reference instead of a hand table, and decide on Q3. |
| F84 | Docs drift | web/src/lib/components/Sidebar.svelte:122, site/guides/lint.md:15-16, web/src/lib/project/project.svelte.ts:14 | Medium | S | The asset tooltip and lint docs say `Structure("x.nbt")`; the helper is `StructureFile` and `Structure` is the worldgen kind. | Rename in all three. |
| F85 | Docs drift | editors/vscode/README.md:3, :15 | Medium | S | Omits the Testing view, Expand lenses and command, prelude go-to-definition, and the `.mmsbook`/`.mmstext` languages; Previews omits books, function runs and template pools. | Update. |
| F86 | Docs drift | engine/syntax/grammar/MMSLexer.g4:3, :8, engine/format/verify.go:10, site/reference/language/values.md:76 | Low | S | Comments cite plan.md sections that don't exist and a function (`separateLines`) that is `separateWords`; values.md lists 7 language types while `types.Names` offers 19, including undocumented `BookText`. | Fix the comments; generate the type list. |
| F87 | Config | Taskfile.yml:5-6 | Low | S | `GREETING: Hello, world!` is template residue; the Taskfile covers only CLI and VS Code builds, not wasm, web or tests. | Delete the var, and either complete the Taskfile or remove it in favour of the npm scripts CI uses. |
| F88 | Consistency | web/src/lib/engine/workspace.ts:15, web/src/lib/editor/Editor.svelte:155 | Low | S | `.mmstext` opens as `mmsbook` and is highlighted with page breaks, though `---` is an error there; VS Code has a separate language. | An `mmstext` language id without page breaks. |
| F89 | Consistency | editors/vscode/src/preview.ts:227-257, src/world.ts:349-367, src/jigsaw.ts:239-258, src/world.ts:703, src/jigsaw.ts:230 | Low | S | The webview HTML shell and CSP are copied three times (nonce from `Math.random`), and two `namespace(project)` readers disagree on the default (`undefined` vs `'mms'`). | One `webviewHtml()` with `crypto.randomBytes`; one `projectNamespace()` matching the engine default. |

## Top 5: if you fix nothing else, fix these

### 1. Separate compiler-made score holders from user locals (F01)

```go
// engine/fn/gen.go
 func Holder(path, name string) string { return "#" + path + "." + name } // user names only
+// GenHolder names compiler-made scores; '%' can't appear in an MMS identifier.
+func GenHolder(path, name string) string { return "#" + path + ".%" + name }
 // flag()
-	s := g.score(Holder(g.path, "_f"+strconv.Itoa(g.n)))
+	s := g.score(GenHolder(g.path, "f"+strconv.Itoa(g.n)))
// engine/fn/opt.go
-func isTemp(s Score) bool { return strings.Contains(s.Holder, "._t") }
+func isTemp(s Score) bool { return strings.Contains(s.Holder, ".%t") }
// engine/sema/fn.go temp() and declare()
-	s := f.score(fn.Holder(f.path, "_t"+strconv.Itoa(f.temps)))
+	s := f.score(fn.GenHolder(f.path, "t"+strconv.Itoa(f.temps)))
-		holder = n.Text + "_" + strconv.Itoa(k)
+		holder = n.Text + "%" + strconv.Itoa(k)
// sema/fn.go string temps (line 1343): the same prefix for local storage paths
-	fn.LocalPath(f.path, "_s"+strconv.Itoa(f.temps))
+	fn.LocalPath(f.path, "%s"+strconv.Itoa(f.temps))
```

Refresh the pinned strings in engine/project/functions_test.go:375-377 and the example goldens, and add VM regression tests for `let _t0`, `let x_1` with a shadowed `x`, and a string local `_s0` next to a concatenation. Do F06 (one SNBT quoter) in the same pass, since it is the other codegen bug that end-to-end tests missed, and add `fn/opt_test.go` (F70).

### 2. Keep every write inside the build folder (F02, F03, F04)

```go
// engine/project/config.go
-var nsPattern = regexp.MustCompile(`^[a-z0-9_.\-]+$`)
+var nsPattern = regexp.MustCompile(`^[a-z0-9_][a-z0-9_.\-]*$`) // never "." or ".."
 func (c Config) Validate() error {
+	for _, o := range []struct{ key, dir string }{{"output", c.Output}, {"resourcePack.output", c.ResourcePack.Output}} {
+		if o.dir != "" && (!filepath.IsLocal(filepath.FromSlash(o.dir)) || filepath.Clean(o.dir) == ".") {
+			return fmt.Errorf("mms.json: %s %q must be a folder inside the project", o.key, o.dir)
+		}
+	}
// engine/project/emit.go, in WriteDir and Zip before each file
+	rel, err := filepath.Localize(f.Path) // rejects "..", absolute paths, and "\" on Windows
+	if err != nil {
+		return fmt.Errorf("refusing to write %q outside the pack", f.Path)
+	}
+	p := filepath.Join(dir, rel)
```

`WriteDir` should also resolve `dir` with `filepath.EvalSymlinks` and confirm it is inside the project before its `RemoveAll` calls, so an in-project `build` symlink can't point the cleanup elsewhere.

Then add a sema error for `..` or empty path segments in declaration names (the same rule as `CheckDimensionID`), so the problem shows at check time in the editor and not only at build time.

### 3. Sanitise LSP Markdown in the web editor (F05)

```ts
// web/src/lib/engine/engine.svelte.ts
+import DOMPurify from 'dompurify'; // pin exactly; add to the `licenses` list in package.json
+const safeHTML = (html: string) => DOMPurify.sanitize(html, { ALLOWED_URI_REGEXP: /^(?:https?:|#|\/|\.)/i });
 this.client = new LSPClient({
   timeout: 10000,
+  sanitizeHTML: safeHTML,
```

Add a Playwright case: a declaration with `/// <img src=x onerror="window.__x=1">`, hover a use, and assert `window.__x` is undefined.

### 4. Make the chunker agree with the parser (F09, F72)

```go
// engine/syntax/chunks.go
-// startsDecl reports whether the line at i starts `Ident name` ...
+// startsDecl reports whether the line at i starts a top-level item: a
+// PascalCase kind then a name, or const/for/extern then a word. Same rule as
+// startsTopItem in recover.go; keep them in one helper.
 func startsDecl(src []rune, i int) bool {
```

In `Reparse`, when a new chunk has errors, parse the whole file once; if that parse is clean, return it unchunked. Then delete the `straddles` exemption in fuzz_test.go:46 and run `TestWhitespaceAgnostic` through `Reparse` with the `else`-on-its-own-line case.

### 5. One typed protocol for the `mms/*` methods (F30, F31, F25)

- `engine/lsp/protocol.go`: `OutputResult{Files, Assets, Items, Texts, JigsawHoles, UnusedAssets, HasErrors}`, `TestsResult`, `Expansion`, `ImportResult`, with JSON tags. Rewrite `output`, `runTests`, `expand` and `importPack` to fill them, which also splits the 194-line `output` (F65).
- `packages/preview/src/protocol.ts` (type-only, no Svelte), exported as `@mms/preview/protocol`, with the method-name constants and a `Methods` map from method to params and result. Delete the copies in editors/vscode/src/messages.ts, web/src/lib/engine/engine.svelte.ts and packages/preview/src/types.ts.
- A Go test marshals a filled `OutputResult` to `testdata/protocol.json`, and a TS check reads the same fixture with `satisfies OutputResult`, so drift fails CI.
- While the types are open: add a `generation` field so clients can skip unchanged output, and pass `texts` through in VS Code.

## Quick wins (Low effort, Medium or higher severity)

- [ ] F01 Separate generated score holders from user locals
- [ ] F02 / F03 / F04 Validate output folders, namespaces and declaration path segments
- [ ] F05 `sanitizeHTML` on the web `LSPClient`
- [ ] F06 One SNBT quoter that escapes control characters
- [ ] F07 Derive the reserved-word list from the lexer
- [ ] F08 / F90 Fix the shadowed `err` in `mms import` and the ignored write in `mms fmt`
- [ ] F09 Chunk only where a top-level item can start
- [ ] F11 Keep `.mmsbook`/`.mmstext` in zip import
- [ ] F12 Show save failures in the web editor
- [ ] F23 Offer Fix All only when asked for it, on a fork
- [ ] F24 Stop re-reading the world on every keystroke in VS Code
- [ ] F40 One name for surface rules and conditions
- [ ] F15 Upgrade fflate and unzip once
- [ ] F33 Carry missing field names in the diagnostic, not the message
- [ ] F53 Restore the book markup reference in hover
- [ ] F80 Add staticcheck to engine CI
- [ ] F81 Align TypeScript to 6.0 in the VS Code extension
- [ ] F83 / F84 / F85 Fix the lint table, the `StructureFile` name and the VS Code README

## Things that look bad but are actually fine

- **About 150 unchecked `fmt.Fprint*` errors in `engine/cmd/mms`.** A CLI printing to stdout/stderr has nothing useful to do on failure. The real unchecked errors are listed separately (F08, F19, F90).
- **Unchecked `binary.Write` in engine/nbt/encode.go.** It writes to a `bytes.Buffer`, which can't fail, and `Encode` is used only by tests.
- **`recover()` in `lsp.dispatch` (engine/lsp/server.go:135) and `project.checkDecl` (engine/project/project.go:570).** Boundary guards that turn a checker panic into an `mms/internal` diagnostic and re-panic under `debugPanics`, not control flow.
- **Zip reading in `mms import`, `importPack` and `MergeWorldPack`.** All go through `fs.FS` (`zip.Reader` and `fstest.MapFS` reject `..`), and the world pack is only re-zipped in memory and written to one fixed path. The import problem is namespace validation (F18), not traversal.
- **`go vet` "unreachable code" in `syntax/parser`.** Generated ANTLR output; CI already excludes it.
- **The `${` wildcard matcher in engine/command/command.go:147.** It looked exponential; timed flat at 1 to 8 interpolations, dominated by spec decode.
- **No origin check on webview `message` listeners, and unvalidated casts of webview messages in the extension.** Only the host and the page's own frames can post, every CSP is `default-src 'none'`, and the extension re-resolves every path the webview sends before acting on it.
- **`{@html}` in web/src/lib/components/JsonView.svelte:38 and Prose.svelte:7.** The first escapes before wrapping spans; the second renders the repo's own `site/` Markdown at build time.
- **The madge cycles in web.** Both are `import type`, so there is no runtime cycle; F30 removes them anyway.
- **`SignatureShim` and the in-band `mms/ping` heartbeat in the web supervisor.** Both carry real weight (lsp-client shows only signature docs; the server keeps parameter docs separate for VS Code), and string ping ids can't collide with numeric request ids.
- **Command lines depend on line breaks.** A raw `/command` runs to the end of its line by design; nothing else reads line breaks on valid input.
- **Two copies of `mms.schema.json`.** engine/project/schema_test.go keeps both in sync with the loader and the lint ids.
- **`ctx.site` collapsing spans inside Templates to the call site.** Diagnostics have to land in the user's file, and ide/tokens.go:92 skips those nodes deliberately.
- **Union narrowing by fewest trial errors (engine/sema/check.go:477).** It stops at the first clean arm, so first-match holds; fewest errors only picks which arm's diagnostics to show.
- **`deepslate` in web dependencies with no imports.** It is the peer for `@mms/preview`, deduped in Vite.
- **No TODO, FIXME or HACK comments anywhere outside vendored and generated code.** Debt here is in the code, not in markers.

## Open questions for the maintainer

1. **Which simulator is authoritative?** Go `mcfn` (`mms test`) and TS cmdlib (function-run preview) run the same output and can disagree (F44). Should one go, or do they share a golden corpus?
2. **Provenance:** wire `Frames` into related information and Template-body go to definition, or delete it (F51)?
3. **Are Blueprint `check` and `hint(...)` rule names configurable in mms.json** like the Go rules? This decides F35 and the `custom-item-as-base` docs bug (F83).
4. **Is "types resolves, sema checks" the final design?** If yes, tasks.md should say so (F60) and the double check (F22) belongs in project; if not, it belongs in types.
5. **Should the web editor keep the full mms.json** (lint, icon, sources), or only namespace and description (F38)?
6. **Do old docs URLs need to keep working?** No users, but the site is public on GitHub Pages, so external links may exist (F66).
7. **Were optional list commas deliberate** (for a recovery case), or a side effect of the overhaul grammar (F42)?
8. **A root npm workspace?** It would remove three lockfiles and the version drift (F81), at the cost of changing CI caching.
9. **Are `SurfaceRule`, `Shaped`, `AboveWater` and friends the official names**, or compatibility aliases that "no aliases" says to remove (F40, F41)?
10. **Is the jigsaw hole picker, and the World/Structures view, intentionally VS Code only?** `mms/output` returns `jigsawHoles`, and nothing in web reads them.
11. **Function signatures use string type names** (`tInt`, `tBool`, engine/sema/fn.go:17-21), so `Type N = int` is "unknown type" in a function. Intended?
12. **Should namespaces `.` and `..` be banned?** Minecraft allows them; F04 recommends banning them.
13. **Is cmdlib's licence (the author's verbal permission, packages/preview/src/vendor/cmdlib/NOTICE.md:7-10) enough for a release that bundles it?**
