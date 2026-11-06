# Library

Written books, made with `Book` declarations in `books.mms`, and the same text markup in chat and titles.

- **`guide`** is given to each player the first time they join (`OnFirstJoin`), along with a title, a welcome in chat that names them with `${@s}`, and the rules from `text/rules.mmstext`. `/function library:guide` gives another copy, with a note above the hotbar. It shows the reader's own keys with `{key:...}`, links between its pages, a link that runs `/function library:chest`, and the reader's death count from a score.
- **`diary_12`, `diary_13`, `diary_14`** are tattered diaries whose text is in `books/*.mmsbook` files. The `old_chest` loot table holds one of them, with some clutter.

Try it in a world:

```
mms build --world ~/.minecraft/saves/MyWorld
```

Join the world for the guide, then use its "Find a chest" link, or run `/function library:chest`, to place a chest of old things. `/loot give @s loot library:old_chest` gives its contents straight away.

In the web editor, the cursor on a `tell`, `title` or `actionbar` line, or in `text/rules.mmstext`, previews the message as the game shows it. In VS Code or the web editor, `.mmsbook` and `.mmstext` files get completion after `{`, `](` and `${`, hover on tags and links, and errors in place. A page with more text than the book screen shows gets a `book-page-overflow` warning where the text is cut off.
