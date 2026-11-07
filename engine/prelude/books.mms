// Book: a written book whose pages are text markup, written in MMS on
// top of the written_book_content component.

/** How many times a written book has been copied. */
Type BookGeneration = "original" | "copy" | "copy_of_copy" | "tattered"

/** Lays out book markup as a book's pages: a `---` line starts the next page. The text is a list of lines, or `TextFile("books/guide.mmsbook")`. */
extern Template book_pages(text: BookText): [Text]

/**
 * A written book with pages in book markup:
 * `Book guide { title = "Guide", author = "The Garrison", text = ["# Welcome", "Plant the **banner** on gold."] }`.
 * Its name stands for the book in item stacks, loot entries and `${...}`
 * in commands.
 */
Blueprint Book {
  /** The title on the cover and the item's name, at most 32 characters. */
  title: string
  /** The author shown under the title: "by <author>". */
  author: string
  /** How many times the book has been copied: `original` (the default), `copy`, `copy_of_copy` or `tattered`. The last two cannot be copied. */
  generation: BookGeneration?
  /**
   * The book's text in book markup, one string per line, or `TextFile("books/guide.mmsbook")` to read it from a file. A line of `---` starts the next page.
   *
   * Markup: `# heading`, `**bold**`, `*italic*`, `__underlined__`, `~~struck~~`, `||obfuscated||`, `{gold}...{/}` (colors, `#ff8800`, `bold`, `shadow:#000000`, `font:uniform`, combined as `{bold red}`), `{key:jump}`, `{sprite:item/diamond}`, `{head:Notch}`, `{t:translation.key}`, `[text](page:2)` (also `run:/command`, `suggest:/command`, `copy:text`, `hover:text`, `https://...`; add `"hover text"` after the target), `${kills[@s]}` for a score or Global, `\` before a character to write it as is.
   */
  text: BookText?
  /** The pages as text components, for full control. Use `text` instead for book markup. */
  pages: [Text]?
  /** More item components, such as `lore` or `custom_data`. Every use of the book gets them. */
  components: DataComponentPatch = {}

  let content = {
    pages = if (text != none) book_pages(text) else pages
    title = title
    author = author
    generation = if (generation == "copy") 1 else if (generation == "copy_of_copy") 2 else if (generation == "tattered") 3
  }
  let book_components = { written_book_content = content } + components

  check {
    if (len(title) > 32) {
      error(title, "a book title is at most 32 characters")
    }
    if (text != none && pages != none) {
      error(pages, "a book takes text or pages, not both")
    }
    if (components.written_book_content != none) {
      error(components, "the book writes written_book_content itself; set title, author and text instead")
    }
  }

  /** In a recipe result, a loot entry, or `${...}` in a command: a written book with these pages. */
  as ItemStackTemplate = { id = written_book, components = book_components }
  as ItemStack = { id = written_book, components = book_components }
  as LootPoolEntry = Item(name = written_book, modifier = [{ type = set_components, components = book_components }])
  /** A Book used where the game takes a plain item id, such as a recipe ingredient or an item tag. It stands for written_book there, so any written book matches. */
  as Id<item> = written_book hint("custom-item-as-base", "'${self.path}' stands for written_book here: this field takes a plain item id, so any written book matches, not only ${self.path}")
}
