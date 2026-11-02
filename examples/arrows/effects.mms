// What each scripted arrow does where it lands. Each runs as the arrow's
// stand-in marker (see flight.mms), at the spot the arrow hit.

/**
 * A small explosion, half the strength of TNT. It is a TNT entity with
 * no fuse, so it goes off on its first tick; block_state air keeps it
 * from showing for that tick. The tnt_explodes game rule turns it off,
 * as it does TNT.
 */
Function explode {
  /summon tnt ~ ~ ~ {fuse:0s,explosion_power:2f,block_state:{Name:"minecraft:air"}}
}

/**
 * Moves the shooter to the spot, half a block back towards where they
 * stood, so they land beside a wall the arrow stuck in rather than
 * inside it.
 */
Function ender_jump {
  owner = shot_by[@s]
  if (owner == 0) {
    return
  }
  /execute as @e[scores={shooter_id=1..}] if score @s shooter_id = $owner mms.arrows run tag @s add arrows.jumper
  /execute facing entity @e[tag=arrows.jumper,limit=1] feet positioned ^ ^ ^0.5 as @e[tag=arrows.jumper,limit=1] run function arrows:land_jump
  /tag @e[tag=arrows.jumper] remove arrows.jumper
}

/** Runs as the shooter, at where they should land. */
Function land_jump {
  /particle portal ~ ~1 ~ 0.3 0.6 0.3 0.5 40
  /execute at @s run particle portal ~ ~1 ~ 0.3 0.6 0.3 0.5 40
  /tp @s ~ ~ ~
  /playsound entity.player.teleport player @a
}
