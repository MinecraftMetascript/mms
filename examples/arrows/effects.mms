// What the custom arrows and the Thunder Bow do where an arrow hits.
//
// on_projectile_hit_block runs as the arrow, at the point it hit, so
// `execute on origin` is whoever shot it. on_projectile_hit_entity runs
// as and at the entity it hurt, so `execute on attacker` is the shooter.

/**
 * A small explosion, half the strength of TNT. It is a TNT entity with
 * no fuse, so it goes off on its first tick; block_state air keeps it
 * from showing for that tick. The tnt_explodes game rule turns it off,
 * as it does TNT.
 */
Function explode {
  /summon tnt ~ ~ ~ {fuse:0s,explosion_power:2f,block_state:{Name:"minecraft:air"}}
}

/** An Explosive Arrow in a block: explode, and take the arrow away so it can't be picked up. */
Function explode_here {
  explode()
  /kill @s
}

/**
 * An Ender Arrow in a block: the shooter lands half a block back towards
 * where they stood, beside the wall rather than inside it.
 */
Function ender_block {
  /execute on origin facing entity @s feet positioned ^ ^ ^0.5 run function arrows:land_jump
  /kill @s
}

/** An Ender Arrow in a mob: the shooter lands a block in front of it. */
Function ender_entity {
  /execute on attacker facing entity @s feet positioned ^ ^ ^1 run function arrows:land_jump
}

/** Runs as the shooter, at where they should land. */
Function land_jump {
  /execute at @s run particle portal ~ ~1 ~ 0.3 0.6 0.3 0.5 40
  /tp @s ~ ~ ~
  /particle portal ~ ~1 ~ 0.3 0.6 0.3 0.5 40
  /playsound entity.player.teleport player @a
}

/** The Thunder Bow: lightning where the arrow hit. */
Function strike {
  /summon lightning_bolt
}
