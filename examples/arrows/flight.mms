// How the custom arrows find out where they land.
//
// An arrow that sticks in a block stays there, with inGround set, so the
// tick function can find it. An arrow that hits a mob is gone by the next
// tick, and with it the place it hit. So each custom arrow gets a marker
// riding on it as a stand-in. When the game removes an arrow it drops its
// passengers where they are, and a stand-in that is no longer riding
// anything marks the spot.
//
// The stand-in is placed at the arrow's position from the end of the
// tick before, so on a mob hit it is up to one tick of flight short of
// the mob: about three blocks at full draw. That is close enough for an
// explosion or a teleport. The Thunder Bow in bow.mms shows the other
// way, which is exact but works per bow rather than per arrow.

/** A number for each player or mob that has shot a custom arrow. */
Objective shooter_id

/** On a stand-in: the shooter_id of whoever shot its arrow, or 0. */
Objective shot_by

Global shooters: int = 0
Global owner: int = 0

/** Whether the entity is riding something. */
Predicate riding = EntityProperties(entity = this, predicate = { vehicle = {} })

OnTick {
  // Look at each new arrow once. Arrows from any bow, crossbow or
  // dispenser count: the arrow entity keeps the item it was shot as.
  As(@e[type=arrow, tag=!arrows.seen]).At(@s) { spotted() }
  // An arrow stuck in a block: its stand-in is still riding it.
  /execute as @e[type=arrow,tag=arrows.flying,nbt={inGround:1b}] at @s run function arrows:stuck
  // A stand-in without its arrow: the arrow hit something and is gone.
  /execute as @e[type=marker,tag=arrows.stand_in,predicate=!arrows:riding] at @s run function arrows:impact
}

/** Runs as and at an arrow the first tick it is seen. */
Function spotted {
  /tag @s add arrows.seen
  /execute if data entity @s item.components."minecraft:custom_data".arrow run function arrows:arm
}

/** Runs as and at a newly shot custom arrow. */
Function arm {
  /tag @s add arrows.flying
  // Remember who shot it, as a number, for arrows that act on the
  // shooter. A mob shooter gets a number too; a dispenser has none.
  owner = 0
  /execute on origin unless score @s shooter_id matches 1.. run function arrows:number_shooter
  /execute on origin run scoreboard players operation $owner mms.arrows = @s shooter_id
  /tag @s add arrows.new
  /execute summon marker run function arrows:mount
  /tag @s remove arrows.new
}

/** Runs as a shooter that has no shooter_id yet. */
Function number_shooter {
  shooters += 1
  shooter_id[@s] = shooters
}

/**
 * Runs as a new marker, at the arrow tagged arrows.new: copies the
 * arrow's kind onto it and puts it on the arrow.
 */
Function mount {
  /tag @s add arrows.stand_in
  /execute if data entity @n[type=arrow,tag=arrows.new] item.components."minecraft:custom_data"{arrow:"explosive"} run tag @s add arrows.explosive
  /execute if data entity @n[type=arrow,tag=arrows.new] item.components."minecraft:custom_data"{arrow:"ender"} run tag @s add arrows.ender
  shot_by[@s] = owner
  /ride @s mount @n[type=arrow,tag=arrows.new]
}

/** Runs as and at an arrow that has just stuck in a block. */
Function stuck {
  /execute on passengers if entity @s[tag=arrows.stand_in] run function arrows:impact
  /kill @s
}

/**
 * Runs as a stand-in, at the spot its arrow hit: does what the arrow's
 * kind does there, then removes the stand-in.
 */
Function impact {
  /execute if entity @s[tag=arrows.explosive] run function arrows:explode
  /execute if entity @s[tag=arrows.ender] run function arrows:ender_jump
  /kill @s
}
