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
/** A vanilla item with its own look and name: `CustomItem ember_shard { base = amethyst_shard, name = "Ember Shard", texture = Texture("art/ember_shard.png") }` writes the model, the item model and the lang entry, and its name stands for the item in item stacks, loot entries and `${...}` in commands. */
extern Blueprint CustomItem
/** A written book with pages in book markup: `Book guide = { title = "Guide", author = "The Garrison", text = ["# Welcome", "Plant the **banner** on gold."] }`. Its name stands for the book in item stacks, loot entries and `${...}` in commands. */
extern Blueprint Book
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
