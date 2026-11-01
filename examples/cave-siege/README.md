# Cave Siege prototype

Phase 1 of the Cave Siege game mode: can mobs be steered to a banner, damage it, and dig through walls? There is no cave world yet.

Build it into a world with `mms build --world <world folder> examples/cave-siege`. Use a superflat creative world. The pack adds a `siege:scratch` dimension that holds the path field, so the world must be opened (or restarted) after the pack is installed: `/reload` does not add dimensions.

How enemies find the banner: when the banner is planted, a flood fill marks every cell an enemy can stand in within 32 blocks across and 16 up or down, with the way back towards the banner. It runs over a few ticks and again every 10 seconds, or within a second of a breaker opening a wall. Each enemy chases a waypoint, an invisible, tiny iron golem that walks that path a few blocks ahead of it. Zombies, skeletons and spiders go for iron golems on their own, so they are never hurt to steer them. Creepers only attack players, so each creeper gets one silent 0.01 damage ping from its first waypoint. Waypoints that reach the banner stay there, and the damage they take is the banner's damage.

1. Stand on open ground and run `/function siege:arena_maze`. It builds a stone maze towards +x and +z and sets the time to midnight.
2. Run `/function siege:give_banner` and place the banner on the gold block. The boss bar shows its health.
3. Stand on the emerald block, run `/function siege:spawn_zombies`, then fly up out of reach (enemies within 8 blocks of a survival player fight that player instead).
4. Watch the zombies walk the maze to the banner. `/function siege:status` shows the banner's health.
5. `/function siege:spawn_mixed` tries skeletons, spiders and creepers. `/function siege:spawn_sappers` spawns zombies that always dig.
6. `/function siege:seal_banner` walls the banner in with cobblestone. Enemies should stop, get flagged as stuck after about 5 seconds, then dig in (3 seconds per stone block). `seal_banner_obsidian` takes 15 seconds per block.
7. `/function siege:reset` clears everything.

What to look for:

- Do all four mob types walk to the banner, or do some ignore their waypoints? No enemy should flash red or make a hurt sound unless something really hit it (a creeper flashes once when it spawns).
- Do enemies keep the target across the whole maze, or wander? Skeletons fight hand to hand until they reach the banner, where they get a bow.
- Break the banner yourself: it should come back within a second with its design and facing, and drop nothing.
- Does the banner lose health at a fair rate? Creeper explosions may hit very hard.
- Do breakers dig the right blocks, and only towards the banner?
