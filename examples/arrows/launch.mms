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
// launch runs as the entity being launched, so it never has to find it
// again with a tag and @e.

OnLoad {
  /execute in minecraft:overworld run forceload add 0 0
}

/**
 * Runs as the player when they right-click with the Cow Wand: a cow
 * appears in front of them and flies the way they look.
 */
Function shoot_cow {
  // A cow is 1.4 blocks tall, so 0.7 below the eye line puts its middle
  // on the line the player looks along.
  Anchored(eyes).Positioned(^ ^-0.7 ^1.5).Summon(cow) { launch("1.5") }
  /playsound entity.cow.ambient neutral @a
}

/**
 * Runs as an entity, rotated the way it should go: turns it to face that
 * way and sets its Motion to speed blocks per tick that way. speed is a
 * number, such as "1.5".
 */
Function launch(speed: string) {
  /rotate @s ~ ~
  /execute anchored feet in minecraft:overworld positioned 0.0 0.0 0.0 positioned ^ ^ ^${speed} summon marker run function arrows:read_vector
  /data modify entity @s Motion set from storage arrows:launch motion
}

/** Runs as a scratch marker: notes its position as the launch vector, then removes it. */
Function read_vector {
  /data modify storage arrows:launch motion set from entity @s Pos
  /kill @s
}
