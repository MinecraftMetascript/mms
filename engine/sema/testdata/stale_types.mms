// mcdoc still lists feature and structure types the game has dropped or
// renamed. The game's registry decides, so these are rejected, with the
// current name when mcdoc maps both to the same config.
Feature lichen = GlowLichen(block = glow_lichen, can_be_placed_on = [stone])
Feature patch = RandomPatch(tries = 4)
Structure village = Village()
Feature raw = { type = glow_lichen }

// The current spellings work.
Feature growth = MultifaceGrowth(block = glow_lichen, can_be_placed_on = [stone])
