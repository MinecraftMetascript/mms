// The path field: a flood fill from the banner over every cell an enemy can stand in.
//
// Each reachable cell gets a wool block at the same position in siege:scratch, an empty
// dimension. Its colour says which of the 12 neighbouring cells (4 sides, each one block up,
// level or down) is one step nearer the banner, so a waypoint walks the shortest path by
// reading the block under itself. The banner's own cell is a gold block.
//
// The fill runs over several ticks, 40 cells a tick, within 32 blocks across and
// 16 blocks up or down of the banner. It runs in whatever dimension the banner is in.

/** An empty void dimension that only holds the path field. It loads when the world starts, not on /reload. */
DimensionType scratch = {
  height = 384
  logical_height = 384
  min_y = -64
  coordinate_scale = 1
  ambient_light = 0
  has_skylight = false
  has_ceiling = false
  has_ender_dragon_fight = false
  infiniburn = "#minecraft:infiniburn_overworld"
  default_clock = overworld
  monster_spawn_block_light_limit = 0
  monster_spawn_light_level = 0
}

Dimension scratch = {
  type = scratch
  generator = {
    type = flat
    settings = {
      biome = "minecraft:the_void"
      layers = []
      features = false
      lakes = false
    }
  }
}

/** Blocks an enemy must not stand in or on. */
BlockTag hazard = [
  minecraft:lava,
  minecraft:water,
  minecraft:bubble_column,
  minecraft:fire,
  minecraft:soul_fire,
  minecraft:magma_block,
  minecraft:powder_snow,
  minecraft:cactus,
  minecraft:sweet_berry_bush,
]

/** 0 when idle; 1 to 5 while clearing the old field, one slab a tick; 6 while flooding. */
Global path_phase: int = 0

/** Set when a block between the enemies and the banner was broken: the field is rebuilt within a second. */
Global path_dirty: bool = false

/** Starts rebuilding the field around the banner. Runs anywhere. */
Function repath {
  let x = cs_tmp["#cx"]
  let z = cs_tmp["#cz"]
  // Keep the scratch area loaded: the fill only works in loaded chunks.
  let from = "" + (x - 32) + " " + (z - 32)
  let to = "" + (x + 32) + " " + (z + 32)
  /execute in siege:scratch run forceload add ${from} ${to}
  path_dirty = false
  path_phase = 1
}

/** Runs every tick while the banner is planted and the field is being built. */
Function path_tick {
  if (path_phase <= 5) {
    // Forceloaded chunks take a few ticks to load, and a fill in an unloaded chunk does nothing.
    if (!scratch_loaded()) {
      return
    }
    clear_slab(path_phase)
    path_phase += 1
    if (path_phase == 6) {
      seed()
    }
    return
  }
  /execute at @e[tag=cs_core,limit=1] run function siege:path_work
}

/** Whether the 4 corners of the scratch area are loaded. */
Function scratch_loaded: bool {
  let x = cs_tmp["#cx"]
  let y = cs_tmp["#cy"]
  let z = cs_tmp["#cz"]
  let a = "" + (x - 32) + " " + y + " " + (z - 32)
  let b = "" + (x + 32) + " " + y + " " + (z - 32)
  let c = "" + (x - 32) + " " + y + " " + (z + 32)
  let d = "" + (x + 32) + " " + y + " " + (z + 32)
  /execute store success score #ld cs_tmp in siege:scratch if loaded ${a} if loaded ${b} if loaded ${c} if loaded ${d}
  return cs_tmp["#ld"] == 1
}

/** Clears one of 5 slabs of the scratch area, 7 blocks high, so no tick fills more than 30,000 blocks. */
Function clear_slab(n: int) {
  let x = cs_tmp["#cx"]
  let y = cs_tmp["#cy"] - 16 + (n - 1) * 7
  let z = cs_tmp["#cz"]
  let top = y + 6
  if (top > cs_tmp["#cy"] + 16) {
    top = cs_tmp["#cy"] + 16
  }
  let from = "" + (x - 32) + " " + y + " " + (z - 32)
  let to = "" + (x + 32) + " " + top + " " + (z + 32)
  /execute in siege:scratch run fill ${from} ${to} minecraft:air
  if (n == 1) {
    save_banner()
  }
}

/** Starts the fill at the banner's own cell, at distance 0. */
Function seed {
  let x = cs_tmp["#cx"]
  let y = cs_tmp["#cy"]
  let z = cs_tmp["#cz"]
  let at = "" + x + " " + y + " " + z
  /execute in siege:scratch run setblock ${at} minecraft:gold_block
  /data modify storage siege:path queue set value []
  /data modify storage siege:path queue append value {x:0,y:0,z:0}
  /execute store result storage siege:path queue[0].x int 1 run scoreboard players get #cx cs_tmp
  /execute store result storage siege:path queue[0].y int 1 run scoreboard players get #cy cs_tmp
  /execute store result storage siege:path queue[0].z int 1 run scoreboard players get #cz cs_tmp
}

/** Runs at the core: takes up to 40 cells off the queue. */
Function path_work {
  let n = 0
  while (n < 40) {
    /execute store result score #more cs_tmp if data storage siege:path queue[0]
    if (cs_tmp["#more"] == 0) {
      path_phase = 0
      return
    }
    path_step()
    n += 1
  }
}

/** Takes the first cell off the queue and visits its 4 neighbours, one block up, level or down. */
Function path_step {
  /execute store result score #px cs_tmp run data get storage siege:path queue[0].x
  /execute store result score #py cs_tmp run data get storage siege:path queue[0].y
  /execute store result score #pz cs_tmp run data get storage siege:path queue[0].z
  /data remove storage siege:path queue[0]
  let x = cs_tmp["#px"]
  let y = cs_tmp["#py"]
  let z = cs_tmp["#pz"]
  // #h is the side the new cell is on: 0 +x, 1 -x, 2 +z, 3 -z.
  cs_tmp["#h"] = 0
  visit(x + 1, y, z)
  cs_tmp["#h"] = 1
  visit(x - 1, y, z)
  cs_tmp["#h"] = 2
  visit(x, y, z + 1)
  cs_tmp["#h"] = 3
  visit(x, y, z - 1)
}

/** Tries the column at x, z, at the current cell's height and one block up or down. */
Function visit(x: int, y: int, z: int) {
  let dx = x - cs_tmp["#cx"]
  let dz = z - cs_tmp["#cz"]
  if (dx > 32 || dx < -32 || dz > 32 || dz < -32) {
    return
  }
  cs_tmp["#nx"] = x
  cs_tmp["#nz"] = z
  /execute positioned ${x} ${y} ${z} run function siege:visit_column
}

/** Runs at the neighbour cell, level with the current cell. At most one of the 3 heights can be standable. */
Function visit_column {
  // #v is the new cell's height: 0 level, 1 one block up, 2 one block down.
  cs_tmp["#v"] = 0
  cs_tmp["#ny"] = cs_tmp["#py"]
  /execute if function siege:standable run return run function siege:mark
  cs_tmp["#v"] = 1
  cs_tmp["#ny"] = cs_tmp["#py"] + 1
  /execute positioned ~ ~1 ~ if function siege:standable run return run function siege:mark
  cs_tmp["#v"] = 2
  cs_tmp["#ny"] = cs_tmp["#py"] - 1
  /execute positioned ~ ~-1 ~ if function siege:standable run return run function siege:mark
}

/**
 * Whether an enemy can stand here and the cell is not in the field yet: room for feet and head,
 * something solid underneath, no hazards, and within 16 blocks up or down of the banner.
 */
Function standable: bool {
  let dy = cs_tmp["#ny"] - cs_tmp["#cy"]
  if (dy > 16 || dy < -16) {
    return false
  }
  /execute store success score #ok cs_tmp if block ~ ~ ~ #siege:walkable if block ~ ~1 ~ #siege:walkable unless block ~ ~-1 ~ #siege:walkable unless block ~ ~ ~ #siege:hazard unless block ~ ~-1 ~ #siege:hazard in siege:scratch if block ~ ~ ~ minecraft:air
  return cs_tmp["#ok"] == 1
}

/** Runs at a new cell: colours it with the way back to the cell it was reached from, and queues it. */
Function mark {
  let d = cs_tmp["#h"] * 3 + cs_tmp["#v"]
  if (d == 0) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:white_wool
  } else if (d == 1) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:orange_wool
  } else if (d == 2) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:magenta_wool
  } else if (d == 3) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:light_blue_wool
  } else if (d == 4) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:yellow_wool
  } else if (d == 5) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:lime_wool
  } else if (d == 6) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:pink_wool
  } else if (d == 7) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:gray_wool
  } else if (d == 8) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:light_gray_wool
  } else if (d == 9) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:cyan_wool
  } else if (d == 10) {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:purple_wool
  } else {
    /execute in siege:scratch run setblock ~ ~ ~ minecraft:blue_wool
  }
  /data modify storage siege:path queue append value {x:0,y:0,z:0}
  /execute store result storage siege:path queue[-1].x int 1 run scoreboard players get #nx cs_tmp
  /execute store result storage siege:path queue[-1].y int 1 run scoreboard players get #ny cs_tmp
  /execute store result storage siege:path queue[-1].z int 1 run scoreboard players get #nz cs_tmp
}

/**
 * Runs as and at a waypoint: moves it one cell nearer the banner. "at @s" after the block test
 * goes back to the waypoint's own dimension, so the teleport does not take it into scratch.
 */
Function carrot_step {
  /execute if entity @s[tag=cs_parked] run return 0
  /execute if entity @e[tag=cs_core,distance=..0.6] run return run tag @s add cs_parked
  /execute in siege:scratch if block ~ ~ ~ minecraft:white_wool at @s run return run tp @s ~-1 ~ ~
  /execute in siege:scratch if block ~ ~ ~ minecraft:orange_wool at @s run return run tp @s ~-1 ~-1 ~
  /execute in siege:scratch if block ~ ~ ~ minecraft:magenta_wool at @s run return run tp @s ~-1 ~1 ~
  /execute in siege:scratch if block ~ ~ ~ minecraft:light_blue_wool at @s run return run tp @s ~1 ~ ~
  /execute in siege:scratch if block ~ ~ ~ minecraft:yellow_wool at @s run return run tp @s ~1 ~-1 ~
  /execute in siege:scratch if block ~ ~ ~ minecraft:lime_wool at @s run return run tp @s ~1 ~1 ~
  /execute in siege:scratch if block ~ ~ ~ minecraft:pink_wool at @s run return run tp @s ~ ~ ~-1
  /execute in siege:scratch if block ~ ~ ~ minecraft:gray_wool at @s run return run tp @s ~ ~-1 ~-1
  /execute in siege:scratch if block ~ ~ ~ minecraft:light_gray_wool at @s run return run tp @s ~ ~1 ~-1
  /execute in siege:scratch if block ~ ~ ~ minecraft:cyan_wool at @s run return run tp @s ~ ~ ~1
  /execute in siege:scratch if block ~ ~ ~ minecraft:purple_wool at @s run return run tp @s ~ ~-1 ~1
  /execute in siege:scratch if block ~ ~ ~ minecraft:blue_wool at @s run return run tp @s ~ ~1 ~1
  // Not in the field (walled off, or the field is being rebuilt): one block straight at the banner, through open space only.
  /execute facing entity @e[tag=cs_core,limit=1] feet rotated ~ 0 positioned ^ ^ ^1 align xyz positioned ~0.5 ~ ~0.5 if block ~ ~ ~ #siege:walkable if block ~ ~1 ~ #siege:walkable run tp @s ~ ~ ~
}

/** Runs at the core: copies the banner, with its design and facing, to the top of the scratch dimension. */
Function save_banner {
  /execute at @e[tag=cs_core,limit=1] if block ~ ~ ~ #minecraft:banners run clone ~ ~ ~ ~ ~ ~ to siege:scratch ~ 319 ~
}

/** Runs at the core: puts the saved banner back. Before the first save, a new Siege Banner is placed. */
Function restore_banner {
  /execute store success score #saved cs_tmp in siege:scratch if block ~ 319 ~ #minecraft:banners
  // The clone runs in the banner's own dimension, so it copies from scratch into that one.
  if (cs_tmp["#saved"] == 1) {
    /clone from siege:scratch ~ 319 ~ ~ 319 ~ ~ ~ ~
    return
  }
  /setblock ~ ~ ~ minecraft:white_banner{patterns:[{pattern:"minecraft:flower",color:"red"},{pattern:"minecraft:border",color:"black"}]}
}
