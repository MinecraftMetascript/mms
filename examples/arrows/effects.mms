// What the custom arrows, the Thunder Bow and the Blast Potion do where
// they hit.
//
// on_projectile_hit_block runs as the arrow, at the point it hit, so
// `execute on origin` is whoever shot it. on_projectile_hit_entity runs
// as and at the entity it hurt, so `execute on attacker` is the shooter.
// on_projectile_land runs at the spot a thrown potion landed.

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

/**
 * Blast Potions fly further than other potions: the first tick one is
 * seen, its speed is multiplied by 1.75. The item hooks mark the potion
 * with mms_item in its custom_data, and a thrown potion keeps its item.
 * Players see the faster flight straight away, because they run a thrown
 * potion's physics themselves and are sent the new velocity.
 */
OnTick {
  /execute as @e[type=splash_potion,tag=!arrows.thrown] run function arrows:throw_further
}

/** Runs as a splash potion the first tick it is seen. */
Function throw_further {
  /tag @s add arrows.thrown
  /execute unless data entity @s Item.components."minecraft:custom_data"{mms_item:"arrows:blast_potion"} run return 0
  /execute store result entity @s Motion[0] double 0.00175 run data get entity @s Motion[0] 1000
  /execute store result entity @s Motion[1] double 0.00175 run data get entity @s Motion[1] 1000
  /execute store result entity @s Motion[2] double 0.00175 run data get entity @s Motion[2] 1000
}
