# Text

Wherever the game takes a text component, a string is markup: a small Markdown-like format for Minecraft text. A string without markup is the same plain string in the output.

```mms
CustomItem ember_shard {
  base = amethyst_shard
  name = "{#ff8800}Ember Shard"
  components = { lore = ["{gray}*Still warm.*"] }
}
```

## Markup

| Markup | Shows |
| --- | --- |
| `**bold**`, `*italic*`, `__underlined__`, `~~struck~~`, `\|\|obfuscated\|\|` | styled text |
| `{gold}...{/}`, `{#ff8800}...{/}` | a color |
| `{bold red}...{/}` | styles combined, also `shadow:#3f1f00` and `font:uniform` |
| `{key:jump}` | the key the player has bound |
| `{sprite:item/diamond}`, `{head:Notch}` | an inline icon, a player's head |
| `{t:block.minecraft.stone}` | a translated text |
| `[text](run:/function demo:start)` | runs a command when clicked; also `suggest:`, `copy:` and `https://` |
| `[text](hover:A red gem)`, `[text](run:/say hi "Says hi")` | hover text, alone or after a target |
| `# Heading` | a bold line |
| `${kills[@s]}`, `${@s}` | a value, filled in by the game |
| `\*` | the character itself |

[Writing resources](../../guides/resources.md#text-markup) has the full table, with what each piece becomes.

## Values in text

`${...}` in text becomes part of the component, so the game fills it in when it shows the text, with no macro: a score (`${kills[@s]}`), a Global or a storage field (`${game.wave}`), or the names a selector finds (`${@s}`). A constant is written in.

## Chat, titles and the action bar

`tell(target, text)` writes `/tellraw`, `title(target, text, subtitle = text)` a title, and `actionbar(target, text)` a line above the hotbar. In a raw command, `${text}` writes the text's JSON.

```mms
Objective deaths: deathCount

OnRespawn {
  tell(@s, "You have died **${deaths[@s]}** times, ${@s}.")
  title(@s, "{red}Back again", subtitle = "{gray}*try the stairs*")
  actionbar(@s, "Press {key:inventory} to check your gear")
}
```

## Named text

There is no text declaration. A fixed piece of text is a `const` of type `Text`, and text with parameters is a `Template` that returns `Text`. Hover shows either one rendered, and the web editor's Preview tab shows it as chat, a title or a book page.

```mms
const welcome: Text = "Welcome to **the Library**, ${@s}! Press {key:inventory} to open your bag."

/** Shown when a wave starts. */
Template wave_start(wave: int): Text = "{red}Wave ${wave}{/} begins. [Rules](run:/say rules \"Show the rules\")"

OnFirstJoin {
  tell(@s, welcome)
  tell(@s, wave_start(1))
}
```

## Text files

`TextFile("text/rules.mmstext")` reads markup from a file in the project. The editor checks and completes markup in the file itself.

- A `.mmstext` file is text for anywhere: chat, a title, an item's lore.
- A `.mmsbook` file is a book's text. A `---` line starts a new page, `[..](page:2)` turns to a page, and the `book-page-overflow` lint marks text past the end of a page.

Page breaks and page links work only in book text, in a `.mmsbook` file or a `Book`'s `text`. Anywhere else they are an error with a fix that removes them.

```mms
Objective deaths: deathCount

Book guide {
  title = "Library Guide"
  author = "The Librarian"
  text = [
    "# Welcome",
    "This world hides {gold}three diaries{/} in old chests.",
    "---",
    "You have died ${deaths[@s]} times.",
  ]
}

OnFirstJoin {
  /give @s ${guide}
}
```

See [Books](../../guides/resources.md#books) for the `Book` fields and the book preview.
