# Cave Siege prototype

Phase 1 of the Cave Siege game mode: can mobs be steered to a banner, damage it, and dig through walls? There is no cave world yet.

Build it into a world with `mms build --world <world folder> examples/cave-siege`. Use a superflat creative world.

1. Stand on open ground and run `/function siege:arena_maze`. It builds a stone maze towards +x and +z and sets the time to midnight.
2. Run `/function siege:give_banner` and place the banner on the gold block. The boss bar shows its health.
3. Stand on the emerald block, run `/function siege:spawn_zombies`, then fly up out of reach (enemies within 8 blocks of a survival player fight that player instead).
4. Watch the zombies walk the maze to the banner. `/function siege:status` shows the banner's health.
5. `/function siege:spawn_mixed` tries skeletons, spiders and creepers. `/function siege:spawn_sappers` spawns zombies that always dig.
6. `/function siege:seal_banner` walls the banner in with cobblestone. Enemies should stop, get flagged as stuck after about 5 seconds, then dig in (3 seconds per stone block). `seal_banner_obsidian` takes 15 seconds per block.
7. `/function siege:reset` clears everything.

What to look for:

- Do all four mob types walk to the banner, or do some ignore the lure?
- Do enemies keep the target across the whole maze, or wander?
- Does the banner lose health at a fair rate? Creeper explosions may hit very hard.
- Do breakers dig the right blocks, and only towards the banner?
