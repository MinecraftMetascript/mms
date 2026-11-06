# How functions compile

What a [function](../guides/functions.md) becomes in the data pack: where its values live, and how branches, loops and early exits turn into commands. You don't need this to write functions, but it helps when reading the generated `.mcfunction` files.

## Values

Values live on one scoreboard objective per namespace, `mms.<namespace>`:

| Holder | Value |
|---|---|
| `$<name>` | a Global |
| `#<function>.<name>` | a parameter or local |
| `#<function>._t<n>`, `#<function>._f<n>` | temporaries and condition flags |
| `#c<n>` | the constant `n`, for multiplication, division and comparison |
| `#cf`, `#rv` | control flow, see below |

String values and macro arguments live in storage `<namespace>:mms`: `g.<name>` for globals, `l."<function>".<name>` for locals, and `m` for the arguments of the macro being called.

**Branches and loops.** `if` with a one-command body becomes `execute if ... run <command>`. Longer bodies become sub-functions next to the function, such as `demo:check/_if0`.

A loop is a function that checks its condition, runs its body and calls itself. `break` returns from it, and `continue` calls it again.

**Leaving early.** When `break`, `continue` or `return` happens inside a sub-function, it sets `#cf`, and the calling function passes it on. The return value travels in `#rv`.

**Results.** Functions return values with the `return` command, and callers read them with `execute store result score`.

**Recursion.** A call that can lead back to the calling function, directly or through other functions, pushes the caller's scores onto a stack in storage, `<namespace>:mms stack`, and pops them afterwards. Only the scores read after the call are saved. Every other call costs nothing extra.

**Optimisation.** Temporaries that only carry a value from one place to the next are removed, so `x = x + 1` becomes one `scoreboard players add`.

**Load.** The generated load function creates the objective, sets the constants, and gives each global its first value.

## Generated functions

`OnLoad { ... }` and `OnTick { ... }` bodies, and `Objective` declarations, are collected into `<namespace>:mms_gen/load` and `<namespace>:mms_gen/tick`. These are registered in `#minecraft:load` and `#minecraft:tick`. `OnFirstJoin` and `OnRespawn` bodies are collected the same way, into `<namespace>:mms_gen/first_join` and `<namespace>:mms_gen/respawn`.

A hook body that holds only commands is copied into these functions as it is. A body with variables or control flow compiles to its own function, such as `<namespace>:mms_gen/load/<hash>`, which the generated function calls.
