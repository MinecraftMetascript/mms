// Written books, and the same markup in chat and titles. A Book's text is
// book markup: **bold**, {gold}colors{/}, {key:jump} for the reader's own
// key, [links](page:2) and ${scores}. A line of --- starts a new page. The
// text can sit in the declaration, one string per line, or in a .mmsbook
// file read with TextFile. Any other text, such as a tell or a title, is
// the same markup without pages; a .mmstext file holds it.

/** How many times the reader has died, for the guide to show. */
Objective deaths: deathCount

/** Given to every player the first time they join. */
Book guide {
  title = "Library Guide"
  author = "The Librarian"
  text = [
    "# Welcome",
    "This world hides {gold}three diaries{/} in old chests.",
    "",
    "Press {key:inventory} to see what you carry, and {key:jump} to climb.",
    "",
    "[Contents](page:2 \"Turn to the contents\")",
    "---",
    "# Contents",
    "[Find a chest](run:/function library:chest \"Places a chest here\")",
    "[About the diaries](page:3)",
    "",
    "You have died ${deaths[@s]} times. {dark_gray}Careful.{/}",
    "---",
    "# The diaries",
    "A miner kept them while digging too deep. Each chest holds one; collect the set to learn what ||happened||.",
  ]
}

/** One of the miner's diaries, worn by time. */
Book diary_12 { title = "Diary, Day 12", author = "A Miner", generation = tattered, text = TextFile("books/day_12.mmsbook") }
Book diary_13 { title = "Diary, Day 13", author = "A Miner", generation = tattered, text = TextFile("books/day_13.mmsbook") }
Book diary_14 { title = "Diary, Day 14", author = "A Miner", generation = tattered, text = TextFile("books/day_14.mmsbook") }

/** A chest of old things: one diary, and some clutter. */
LootTable old_chest {
  pools = [
    Pool(1, [diary_12, diary_13, diary_14]),
    Pool(1..3, [torch, coal, Item(name = paper, modifier = [SetCount(1..4)])]),
  ]
}

/** The welcome a new reader sees in chat. */
const welcome: Text = "Welcome to the {gold}Library{/}, ${@s}! Press {key:inventory} to find your guide."

OnFirstJoin {
  /give @s ${guide}
  title(@s, "{gold}The Library", subtitle = "{gray}*mind the shelves*")
  tell(@s, welcome)
  tell(@s, TextFile("text/rules.mmstext"))
}

/** /function library:guide gives another copy of the guide. */
Function guide {
  /give @s ${guide}
  actionbar(@s, "{green}You take a fresh copy of the guide")
}

/** /function library:chest places a chest of old things where you stand. */
Function chest {
  /setblock ~ ~ ~ chest{LootTable:"library:old_chest"}
}
