// The language's own declaration kinds. The compiler builds them in Go;
// these declarations give them their names and docs.

/** Sets the namespace for the declarations after it in the file. */
extern Blueprint Namespace
/** A function: commands, variables and control flow, compiled to .mcfunction files. */
extern Blueprint Function
/** A scoreboard objective, created by the generated load function. */
extern Blueprint Objective
/** A value for the whole pack, set when the pack first loads. */
extern Blueprint Global
/** A record kept in command storage: `Storage game: GameState` keeps a `Type GameState { wave: int = 0 }` at `<namespace>:mms game`. Functions read and write its fields as `game.wave`; the pack sets each missing field to its default when it loads. */
extern Blueprint Storage
/** A test of the pack's functions, run by `mms test` and the editor: a `setup { }` block, then `run f()` and `expect condition` lines. Tests run in a fresh world after the pack loads and are never written to the pack. */
extern Blueprint Test
/** Code that runs when the pack loads, from #minecraft:load. */
extern Blueprint OnLoad
/** Code that runs every tick, from #minecraft:tick. */
extern Blueprint OnTick
/** Code that runs once for each player, the first time they are in the world with the pack, as and at that player. */
extern Blueprint OnFirstJoin
/** Code that runs as and at a player right after they respawn. */
extern Blueprint OnRespawn
/** Code that runs as and at a player each time they meet a trigger: `OnAdvancement name = Trigger(...) { ... }` writes the advancement, its reward function and the revoke. */
extern Blueprint OnAdvancement
/** A name for jigsaw blocks, used only by the editor: `JigsawHole road_end` lists `<namespace>:road_end` in the Jigsaw view's Name and Target lists, so pieces built in the game join by names spelled the same way. It writes no file. */
extern Blueprint JigsawHole

/** Text in English, or one text per language: `{ en_us = "Ruby", de_de = "Rubin" }`. */
Type LangText = string | {string: string}

/** The cosine of an angle in degrees. */
extern Template cos(degrees: double): double
/** The sine of an angle in degrees. */
extern Template sin(degrees: double): double
/** The number of items in a list, fields in an object, or characters in a string. */
extern Template len(value: any): int
/** A LangText as an object of language code to text; plain text is English. */
extern Template by_language(text: LangText?): {string: string}
/** The English text of a LangText, or its first text when it has no English. */
extern Template english(text: LangText): string
/** Whether a list holds a value. */
extern Template contains(list: [any], value: any): bool
/** Whether a string matches a regular expression, as a whole. */
extern Template matches(text: string, pattern: string): bool
/** A string with every `old` replaced by `new`. */
extern Template replace(text: string, old: string, new: string): string
/** An asset call such as Texture("art/ruby.png") with an id name, unless it has one: named(texture, "item/ruby"). */
extern Template named(asset: any, name: string): any
/** The kind of a value: "none", "bool", "number", "string", "name", "list", "object", "range", "commands", or "runtime" for a value only known when the pack runs. */
extern Template kind_of(value: any): string
/** A whole number from 0 to 2147483647 made from a string, the same in every build: `hash("dungeon:dungeon")`. */
extern Template hash(text: string): int

/** Sends `message` to the chat of each player `target` selects: `tell(@a, "Wave ${wave} begins")`. The message is text markup, and `${...}` in it shows a score, a Global, a selector's names or a constant. */
extern Template tell(target: Selector, message: Text): Commands
/** Shows `message` in large text in the middle of the screen of each player `target` selects, with `subtitle` under it. */
extern Template title(target: Selector, message: Text, subtitle: Text?): Commands
/** Shows `message` just above the hotbar of each player `target` selects. */
/** Whether `target` selects at least one entity, in a function: `if (exists(@e[tag=boss])) { ... }`. */
extern Template exists(target: Selector): bool
/** How many entities `target` selects, in a function: `let n = count(@e[type=zombie, distance=..16])`. */
extern Template count(target: Selector): int
extern Template actionbar(target: Selector, message: Text): Commands
