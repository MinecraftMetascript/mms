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
// The vector then goes into the summon command itself, as the new
// entity's Motion. Setting Motion with /data after the entity exists
// works on the server, but the game told players about the entity when
// it was summoned, standing still, so it seems to hang in the air for a
// moment before it moves. An entity summoned with its Motion is seen
// moving from its first frame.

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
  // 0.7 below the eye line puts the middle of a cow (1.4 blocks tall) on
  // the line the player looks along.
  Anchored(eyes).Positioned(^ ^-0.7 ^1.5) {
    /summon ${entity} ~ ~ ~ {Motion:${launch_motion},Rotation:[${launch_yaw},0.0]}
  }
}

/** Runs as a scratch marker: notes its position as the launch vector, then removes it. */
Function read_vector {
  /data modify storage arrows:mms g.launch_motion set from entity @s Pos
  /kill @s
}
