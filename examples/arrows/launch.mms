// Launching an entity the way the player is looking.
//
// Motion is a vector in blocks per tick. The trick for getting one that
// points where the player looks: a marker summoned at ^ ^ ^speed from
// the world origin has that vector as its Pos. Three things make the
// trick reliable:
//
// - The marker has to be summoned somewhere loaded, so OnLoad keeps the
//   origin chunk in the overworld loaded, and the marker is summoned in
//   the overworld whatever dimension the shooter is in. (Recent versions
//   no longer keep spawn chunks loaded.)
// - The marker has to be killed once it has been read, or every launch
//   leaves one behind.
// - ^ ^ ^ has to be read with anchored feet. Under anchored eyes, ^ ^ ^
//   starts at the eyes of whoever runs the command, which bends the
//   vector upwards.
//
// Then there is how players see it. A mob tells players where it is only
// every few ticks, and they slide it between those points, so a cow
// launched on its own seems to hang for a moment before it flies. A
// falling block is different: players run its physics themselves from
// the velocity it was summoned with, so it moves smoothly from its first
// frame. So the cow rides an invisible falling block (structure_void
// draws nothing) summoned with the Motion. When the block lands it is
// removed rather than placed (CancelDrop), and the cow drops off where it
// landed. A falling block has half a mob's gravity, so it also flies
// further for the same speed.

OnLoad {
  /execute in minecraft:overworld run forceload add 0 0
}

/**
 * The vector and the facing for the next launch. ${} writes whatever
 * the storage holds (Globals live in storage arrows:mms, at g.<name>),
 * so these hold a list of three doubles and a number, not text.
 */
Global launch_motion: string = ""
Global launch_yaw: string = ""

/**
 * Runs as the player when they right-click with the Cow Wand: a cow
 * appears in front of them and flies the way they look.
 */
Function shoot_cow {
  launch("cow", "1.5")
  /playsound entity.cow.ambient neutral @a
}

/**
 * Runs as and at a player: summons an entity in front of their eyes,
 * facing the way they look and already moving that way at speed blocks
 * per tick. entity is an entity id, such as "cow"; speed is a number,
 * such as "1.5".
 */
Function launch(entity: string, speed: string) {
  /execute anchored feet in minecraft:overworld positioned 0.0 0.0 0.0 positioned ^ ^ ^${speed} summon marker run function arrows:read_vector
  /data modify storage arrows:mms g.launch_yaw set from entity @s Rotation[0]
  // A rider sits on top of the block, so 1.7 below the eye line puts the
  // middle of a cow (1.4 blocks tall) on the line the player looks along.
  Anchored(eyes).Positioned(^ ^-1.7 ^1.5) {
    /summon falling_block ~ ~ ~ {BlockState:{Name:"minecraft:structure_void"},Time:1,DropItem:0b,CancelDrop:1b,HurtEntities:0b,Motion:${launch_motion},Passengers:[{id:"${entity}",Rotation:[${launch_yaw},0.0]}]}
  }
}

/** Runs as a scratch marker: notes its position as the launch vector, then removes it. */
Function read_vector {
  /data modify storage arrows:mms g.launch_motion set from entity @s Pos
  /kill @s
}
