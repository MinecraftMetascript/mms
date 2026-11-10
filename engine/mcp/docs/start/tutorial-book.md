# Tutorial: a guide book

This tutorial writes the book a map hands to new players: a few pages with headings, colours, links that turn pages and run commands, and a live death count. Then it puts longer books in files of their own and hides them in a chest.

In the game's own format, each page of a written book is a tree of text components. In MMS a page is a few lines of markup, and the editor shows the open book as you type. The finished project is close to the [`library`](../../examples/library/) example.

## 1. Set up the project

```sh
mkdir library && cd library
mms init -namespace library
```

## 2. A first book

```mms
Book guide {
  title = "Library Guide"
  author = "The Librarian"
  text = [
    "# Welcome",
    "This world hides {gold}three diaries{/} in old chests.",
    "",
    "Press {key:inventory} to see what you carry.",
  ]
}

Function guide {
  /give @s ${guide}
}
```

`text` holds one string per line of the book. `# ` makes a bold heading, `{gold}...{/}` colours what is between, and `{key:inventory}` shows whichever key the reader has bound, so it is right for everyone. `${guide}` in a command is the finished book item.

Put the cursor in the book and open the **Preview** tab of the web editor (in VS Code, click **Preview** above the `Book`). It draws the open book with the game's own font and texture.

## 3. More pages, and links

A line of `---` starts a new page. A link looks like Markdown's, and its target says what a click does:

```mms
Book guide {
  title = "Library Guide"
  author = "The Librarian"
  text = [
    "# Welcome",
    "This world hides {gold}three diaries{/} in old chests.",
    "",
    "[Contents](page:2 \"Turn to the contents\")",
    "---",
    "# Contents",
    "[Find a chest](run:/function library:chest \"Places a chest here\")",
    "[About the diaries](page:3)",
    "---",
    "# The diaries",
    "A miner kept them while digging too deep. Collect the set to learn what ||happened||.",
  ]
}

Function chest {
  /setblock ~ ~ ~ chest
}
```

- `page:2` turns to a page, and the quoted text after a target is what shows on hover.
- `run:/function library:chest` runs a command as the reader. MMS checks that command like any other, so a link to a function that does not exist is an error here and not a dead link in the game.
- `||happened||` is obfuscated: the scrambled text the game uses for secrets.

In the preview, click the links and the page arrows to walk through the book. The full list of markup is in [Text markup](../guides/resources.md#text-markup).

## 4. Values that change

`${...}` puts a value in the text. The game fills it in when the reader first opens the book:

```mms
/** How many times each player has died. */
Objective deaths: deathCount

Book guide {
  title = "Library Guide"
  author = "The Librarian"
  text = [
    "# Your record",
    "Reader: ${@s}",
    "You have died ${deaths[@s]} times. {dark_gray}Careful.{/}",
  ]
}
```

`Objective deaths: deathCount` is a score the game keeps by itself. `${deaths[@s]}` is the reader's score and `${@s}` their name. The preview lists the values a book uses, with boxes to try numbers and names in.

## 5. Keep long text in a file

Strings in a list are fine for a short guide. A story reads better in a file of its own. Make `books/day_12.mmsbook`:

```
# Day 12
The lamps went out again at the bottom of the shaft. We relit them with the last of the *dry* matches.

Tomas says the stone down here is **warm**.
---
Something knocks behind the east wall. Three knocks, a pause, three knocks.

{dark_red}We do not knock back.{/}
```

and read it with `TextFile`:

```mms in=examples/library
Book diary_12 {
  title = "Diary, Day 12"
  author = "A Miner"
  generation = tattered
  text = TextFile("books/day_12.mmsbook")
}
```

A `.mmsbook` file is the same markup without the quotes. The editor highlights and completes it, previews each page, and checks it in place.

A book page shows 14 lines of about 19 characters, and the game cuts off the rest without a word. MMS measures each page with the vanilla font and warns where one runs over ([`book-page-overflow`](../guides/lint.md)). Add a `---` there.

## 6. Many books from one declaration

Three diaries would be three declarations that differ by a number. A `for` at the top of a file repeats a declaration, and `${day}` in the name gives each copy its own id:

```mms in=examples/library
for day in [12, 13, 14] {
  /** The miner's diary for day ${day}. */
  Book diary_${day} {
    title = "Diary, Day ${day}"
    author = "A Miner"
    generation = tattered
    text = TextFile("books/day_${day}.mmsbook")
  }
}

/** A chest of old things: one diary, and some clutter. */
LootTable old_chest {
  pools = [
    Pool(1, [diary_12, diary_13, diary_14]),
    Pool(1..3, [torch, coal, paper]),
  ]
}

/** Places a chest of old things where you stand. */
Function chest {
  /setblock ~ ~ ~ chest{LootTable:"library:old_chest"}
}
```

A book's name works as a loot entry, as a recipe result and in `/give`, like a [custom item](tutorial-custom-item.md)'s. The first pool drops one diary, and the second one to three bits of clutter.

## 7. Greet new players

The same markup works in chat, titles and the action bar. `OnFirstJoin` runs once for each player, the first time they join:

```mms
Book guide {
  title = "Library Guide"
  author = "The Librarian"
  text = ["# Welcome"]
}

/** The welcome a new reader sees in chat. */
const welcome: Text = "Welcome to the {gold}Library{/}, ${@s}! Press {key:inventory} to find your guide."

OnFirstJoin {
  /give @s ${guide}
  title(@s, "{gold}The Library", subtitle = "{gray}*mind the shelves*")
  tell(@s, welcome)
}
```

A `const` of type `Text` names a piece of text. Hover it to see it drawn, and put the cursor on a `tell` or `title` to see it in the Preview tab as chat or as a title. Text for chat can live in a file too: a `.mmstext` file is markup without pages, read with `TextFile` in the same way.

## 8. Build and read

```sh
mms build --world ~/.minecraft/saves/MyWorld
```

Run `/reload` in the world. You have already joined it, so the first-join greeting will not fire for you: run `/function library:guide` for a copy, or `/advancement revoke @s only library:mms_gen/first_join` to be new again.

## Next steps

- [Text](../reference/language/text.md) in the language reference lists every piece of markup and where values can go.
- Use the same markup in an advancement's title, an item's `lore` or a sign: anywhere the game takes a text component, a string is markup.
- The [`library`](../../examples/library/) example adds a rules message in a `.mmstext` file.
