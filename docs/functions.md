# Functions

`Function`, `OnLoad` and `OnTick` bodies compile to `.mcfunction` files.

## Raw commands

A line that starts with `/` is a command. It is written to the output as it is, without the slash:

```
Function greet {
  /say hello
  /execute as @a at @s run particle minecraft:heart ~ ~2 ~
}
```

Every command is checked against the 26.3 command tree, the same one the game uses. The checker reports:

- unknown commands and options, with the alternatives
- missing and extra arguments
- argument syntax: numbers and their bounds, selectors and their options, coordinates, ids, ranges, and balanced NBT and text
- unknown vanilla ids, such as `minecraft:diamnd`
- ids in your own namespaces that nothing declares, such as `function demo:missing`

Completion inside a command offers the next literal and values for the argument being typed. These include selectors, game modes, registry ids, and your project's functions, tags and objectives.

The checker checks syntax, not meaning. Whether an entity exists, or whether an NBT path fits the entity, is left to the game.

## Generated functions

`OnLoad { ... }` and `OnTick { ... }` bodies, and `Objective` declarations, are collected into `<namespace>:mms_gen/load` and `<namespace>:mms_gen/tick`. These are registered in `#minecraft:load` and `#minecraft:tick`.
