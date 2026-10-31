Namespace example

// Pools take rolls and entries positionally; bare ids are item entries
// and #tags expand to one entry per item.
LootTable chests/gems {
  pools = [
    Pool(1..3, diamond, emerald, #minecraft:logs),
    Pool(1, [Item(name = gold_ingot).Weight(3)]),
  ]
}

// Conditions combine with && (all_of), || (any_of) and ! (inverted).
Predicate stormy_night = WeatherCheck(thundering = true) && !(TimeCheck(clock = overworld, value = 0..12000) || RandomChance(0.5))

// Score and count bounds are ranges; either end may be left open.
Predicate high_score = EntityScores(entity = this, scores = { kills = 10.., deaths = ..2 })

// Shaped(pattern, key, result) and Shapeless(ingredients, result); an
// item times a number is a stack.
Recipe ruby_block = Shaped(["RRR", "RRR", "RRR"], { R = emerald }, emerald_block)
Recipe rubies = Shapeless([emerald_block], emerald * 9)

// Criteria are trigger helpers whose arguments are the conditions.
Advancement first_gem = {
  criteria = { has_gem = InventoryChanged(items = [{ items = [diamond, emerald] }]) }
}

Predicate bad = WeatherCheck(raining = true) && 1
Recipe bad_count = Shapeless([emerald], emerald * gold)
