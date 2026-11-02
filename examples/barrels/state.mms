// What every barrel keeps, and the working values of the one being used.
//
// A barrel is a barrel block plus three entities on its front, tied
// together by a shared barrel_id score:
//
// - an interaction entity (tag barrels.hit) that catches clicks on the
//   front and keeps the barrel's count and stack size as scores
// - an item display (tag barrels.face) that shows the stored item and is
//   the one place the item itself is kept, with a count of 1
// - a text display (tag barrels.label) that shows how many there are
//
// Every part has the tag barrels.part.

Objective barrel_id
Objective barrel_count
Objective barrel_stack

/** The id the next barrel gets. */
Global next_id: int = 0

// The barrel being used, loaded by open() and saved by close(). Some raw
// commands below write these with `scoreboard players set $<name>
// mms.barrels` or `execute store`, because a Global named x is the score
// $x on the objective mms.<namespace> (see docs/functions.md, "How it
// compiles"). The item modifiers at the end of this file read them the
// same way.
Global id: int = 0
Global count: int = 0
/** The stored item's stack size: 16 or 64, or 0 while the barrel is empty. */
Global stack: int = 0

// Scratch values for the raw commands.
/** 1 when the player clicked with an item in hand. */
Global holding: int = 0
/** 1 when the player was sneaking as they clicked. */
Global sneaking: int = 0
Global held: int = 0
Global diff: int = 0
Global amount: int = 0
/** Minus the number of items to take from a player's slot. */
Global minus: int = 0

/** How many items fit: 32 stacks of the stored item. */
Function capacity(): int {
  return stack * 32
}

/**
 * Runs as a player who clicked a barrel, once barrels.this is on the
 * barrel's interaction entity. Tags the barrel's other parts and loads
 * its scores into the Globals above.
 */
Function open {
  id = barrel_id[@e[type=interaction, tag=barrels.this, limit=1]]
  count = barrel_count[@e[type=interaction, tag=barrels.this, limit=1]]
  stack = barrel_stack[@e[type=interaction, tag=barrels.this, limit=1]]
  As(@e[tag=barrels.part, distance=..12]) {
    if (barrel_id[@s] == id) {
      /tag @s add barrels.this
    }
  }
}

/** Saves the Globals back to the barrel, updates its label and lets go of it. */
Function close {
  if (count == 0) {
    // An empty barrel forgets its item, so it can take any kind next.
    stack = 0
  }
  barrel_count[@e[type=interaction, tag=barrels.this, limit=1]] = count
  barrel_stack[@e[type=interaction, tag=barrels.this, limit=1]] = stack
  if (count == 0) {
    /item replace entity @e[type=item_display,tag=barrels.this,limit=1] contents with air
    /data modify entity @e[type=text_display,tag=barrels.this,limit=1] text set value {text:"Empty",color:"gray"}
  } else if (count >= capacity()) {
    /data modify entity @e[type=text_display,tag=barrels.this,limit=1] text set value {text:"${count}",color:"gold"}
  } else {
    /data modify entity @e[type=text_display,tag=barrels.this,limit=1] text set value {text:"${count}",color:"white"}
  }
  /tag @e[tag=barrels.this] remove barrels.this
}

/** Item modifiers that set a stack's count from the Globals. */
ItemModifier one = SetCount(1)
ItemModifier take_amount = SetCount(Score(Fixed("$amount"), "mms.barrels"))
ItemModifier remove_amount = SetCount(Score(Fixed("$minus"), "mms.barrels"), add = true)
