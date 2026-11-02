// Placing and breaking. A Storage Barrel places an invisible item frame;
// the next tick turns it into a barrel block with three entities on its
// front. The tick function also notices barrels whose block is gone.

/** The direction the new barrel's front faces, as an item frame Facing: 2 north, 3 south, 4 west, 5 east. */
Global dir: int = 0

OnTick {
  /execute as @e[type=item_frame,tag=barrels.place] at @s run function barrels:place
  // Each interaction entity sits half a block in front of its barrel's
  // centre, facing out, so ^ ^0.375 ^-0.5 is the centre again.
  /execute as @e[type=interaction,tag=barrels.hit] at @s positioned ^ ^0.375 ^-0.5 unless block ~ ~ ~ minecraft:barrel run function barrels:broken
}

/** Runs as and at a Storage Barrel's item frame, the tick after it was placed. */
Function place {
  // A frame on a wall faces away from it, toward the player.
  /execute store result score $dir mms.barrels run data get entity @s Facing
  if (dir < 2) {
    // On a floor or ceiling the frame faces up or down, so face the
    // player who placed it instead.
    dir = 3
    /execute as @p if entity @s[y_rotation=-45..45] run scoreboard players set $dir mms.barrels 2
    /execute as @p if entity @s[y_rotation=45..135] run scoreboard players set $dir mms.barrels 5
    /execute as @p if entity @s[y_rotation=-135..-45] run scoreboard players set $dir mms.barrels 4
  }
  Align(xyz).Positioned(~0.5 ~0.5 ~0.5) {
    /execute if block ~ ~ ~ #minecraft:replaceable run function barrels:build
    /execute unless block ~ ~ ~ #minecraft:replaceable run loot spawn ~ ~ ~ loot barrels:storage_barrel
  }
  /kill @s
}

/** Builds a barrel at the centre of this block, its front facing dir. */
Function build {
  // The lock keeps the barrel's own 27 slots shut: right-clicking its
  // sides says "Barrel is locked" instead of opening a second, hidden
  // inventory. No real item matches the key.
  if (dir == 2) {
    /setblock ~ ~ ~ barrel[facing=north]{lock:{items:"minecraft:structure_void",components:{"minecraft:custom_data":{barrels:"key"}}}}
    parts(180)
  } else if (dir == 3) {
    /setblock ~ ~ ~ barrel[facing=south]{lock:{items:"minecraft:structure_void",components:{"minecraft:custom_data":{barrels:"key"}}}}
    parts(0)
  } else if (dir == 4) {
    /setblock ~ ~ ~ barrel[facing=west]{lock:{items:"minecraft:structure_void",components:{"minecraft:custom_data":{barrels:"key"}}}}
    parts(90)
  } else {
    /setblock ~ ~ ~ barrel[facing=east]{lock:{items:"minecraft:structure_void",components:{"minecraft:custom_data":{barrels:"key"}}}}
    parts(-90)
  }
  /playsound minecraft:block.barrel.close block @a
}

/**
 * Summons the front of a barrel at the centre of this block. yaw is the
 * way the front faces (0 south, 90 west, 180 north, -90 east), so ^ ^ ^0.5
 * is the middle of the front face.
 */
Function parts(yaw: int) {
  next_id += 1
  // Clicks land on this 0.75 block box, which sticks out of the front
  // face. The sides, top and edges of the block are still the block, so
  // it can be mined as usual.
  /execute rotated ${yaw} 0 run summon interaction ^ ^-0.375 ^0.5 {Tags:["barrels.part","barrels.hit","barrels.new"],width:0.75f,height:0.75f,response:true,Rotation:[${yaw}f,0f]}
  // The stored item, drawn like an item in an item frame and flattened
  // so that blocks do not sink into the barrel.
  /execute rotated ${yaw} 0 run summon item_display ^ ^0.08 ^0.52 {Tags:["barrels.part","barrels.face","barrels.new"],item_display:"fixed",Rotation:[${yaw}f,0f],transformation:{left_rotation:[0f,0f,0f,1f],right_rotation:[0f,0f,0f,1f],translation:[0f,0f,0f],scale:[0.45f,0.45f,0.02f]}}
  // The count, under the item.
  /execute rotated ${yaw} 0 run summon text_display ^ ^-0.42 ^0.52 {Tags:["barrels.part","barrels.label","barrels.new"],text:{text:"Empty",color:"gray"},billboard:"fixed",background:0,Rotation:[${yaw}f,0f],transformation:{left_rotation:[0f,0f,0f,1f],right_rotation:[0f,0f,0f,1f],translation:[0f,0f,0f],scale:[0.5f,0.5f,0.5f]}}
  barrel_id[@e[tag=barrels.new]] = next_id
  barrel_count[@e[tag=barrels.new]] = 0
  barrel_stack[@e[tag=barrels.new]] = 0
  /tag @e[tag=barrels.new] remove barrels.new
}

/**
 * Runs as a barrel's interaction entity, at the centre of the block, when
 * the barrel block is gone: mined, blown up or replaced.
 */
Function broken {
  /tag @s add barrels.this
  id = barrel_id[@s]
  count = barrel_count[@s]
  stack = barrel_stack[@s]
  As(@e[tag=barrels.part, distance=..2]) {
    if (barrel_id[@s] == id) {
      /tag @s add barrels.this
    }
  }
  // Spill the contents, a stack at a time.
  while (count > 0) {
    amount = stack
    if (amount > count) {
      amount = count
    }
    count -= amount
    /summon item ~ ~ ~ {Item:{id:"minecraft:stone"},Tags:["barrels.drop"]}
    give_drop()
  }
  // A mined barrel drops a plain barrel; swap it for a Storage Barrel. In
  // creative nothing drops, so nothing is swapped.
  /execute as @e[type=item,distance=..1,limit=1,sort=nearest] if items entity @s contents minecraft:barrel at @s run function barrels:swap_drop
  /kill @e[tag=barrels.this]
}

/** Runs as the plain barrel item a broken Storage Barrel dropped. */
Function swap_drop {
  /loot spawn ~ ~ ~ loot barrels:storage_barrel
  /kill @s
}

/** Fills the new item entity tagged barrels.drop with amount of the stored item. */
Function give_drop {
  /item replace entity @e[type=item,tag=barrels.drop,limit=1] contents from entity @e[type=item_display,tag=barrels.this,limit=1] contents
  /item modify entity @e[type=item,tag=barrels.drop,limit=1] contents barrels:take_amount
  /tag @e[type=item,tag=barrels.drop] remove barrels.drop
}
