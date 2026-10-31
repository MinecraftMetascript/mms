# Loot tables: entries, conditions and functions

## minecraft:loot_pool_entry/*
- weight: How likely this entry is to be picked against the other entries of the pool, relative to their weights. Defaults to 1.
- quality: Changes `weight` by the luck of the player: the weight becomes `weight + quality * luck`. Defaults to 0.
- condition: Conditions that must all pass for the entry to be used. Combine several with `&&` and `||`.
- modifier: Item modifiers (loot functions) applied to the items this entry gives.

## minecraft:loot_function/*
- condition: Conditions that must all pass for the function to apply. Combine several with `&&` and `||`.

## minecraft:loot_condition/*
- entity: Which entity to check: `this` (the entity that died, or the player that opened or broke something), `attacker`, `direct_attacker` or `attacking_player`.

## minecraft:loot_pool_entry/item
Gives one item. A plain item id where an entry is expected does the same thing.

- name: The item.

```mms
LootTable gems {
  pools = [Pool(1, [Item(name = diamond).Weight(1), Item(name = emerald).Weight(3)])]
}
```

## minecraft:loot_pool_entry/tag
Gives items from an item tag: all of them, or with `expand`, one picked at random. A `#tag` where an entry is expected does the same thing.

- items: The item tag.

## minecraft:loot_pool_entry/loot_table
Rolls another loot table and gives what it drops.

- value: The loot table, by id, or a table written inline.

## minecraft:loot_pool_entry/empty
Gives nothing. Give it a weight to make a pool sometimes drop nothing.

```mms
LootTable maybe {
  pools = [Pool(1, [Item(name = diamond), Empty().Weight(4)])]
}
```

## minecraft:loot_pool_entry/alternatives
Uses the first child whose conditions pass, and ignores the rest.

- children: The entries to try, in order.

## minecraft:loot_pool_entry/group
Uses every child whose conditions pass.

- children: The entries to use.

## minecraft:loot_pool_entry/sequence
Uses children in order until one's conditions fail.

- children: The entries to use, in order.

## minecraft:loot_pool_entry/dynamic
Drops a block's own contents: `contents` for a shulker box's items, `sherds` for a decorated pot's sherds.

- name: Which contents to drop.

## minecraft:loot_pool_entry/slots
Gives items taken from slots of an entity or container, picked by a slot source.

## minecraft:loot_condition/random_chance
Passes at random, with probability `chance`.

```mms
Predicate coin_flip = RandomChance(0.5)
```

## minecraft:loot_condition/random_chance_with_enchanted_bonus
Passes at random, with a higher chance when the attacker's weapon has an enchantment, as looting does for rare drops.

- unenchanted_chance: The chance without the enchantment, from 0 to 1.
- enchanted_chance: The chance by enchantment level.
- enchantment: The enchantment that raises the chance.

## minecraft:loot_condition/all_of
Passes when every condition passes. Usually written with `&&`.

## minecraft:loot_condition/any_of
Passes when any condition passes. Usually written with `||`.

## minecraft:loot_condition/inverted
Passes when `term` fails. Usually written with `!`.

- term: The condition to invert.

## minecraft:loot_condition/weather_check
Passes when the weather matches. Leave out what doesn't matter.

- raining: `true` to require rain, `false` to require no rain.
- thundering: `true` to require a thunderstorm, `false` to require none.

```mms
Predicate storm = WeatherCheck(thundering = true)
```

## minecraft:loot_condition/time_check
Passes when a world clock's time is in a range.

```mms
Predicate night = TimeCheck(clock = overworld, value = 13000..23000, period = 24000)
```

## minecraft:loot_condition/killed_by_player
Passes when a player killed the entity.

- inverse: `true` to pass when no player killed it instead.

## minecraft:loot_condition/survives_explosion
Passes with a chance of 1 divided by the explosion's radius, and always when there was no explosion. Vanilla block drops use it so explosions destroy some items.

## minecraft:loot_condition/match_tool
Passes when the tool used, such as the pickaxe that broke a block, matches an item predicate.

- predicate: What the tool must be.

## minecraft:loot_condition/entity_properties
Passes when an entity matches an entity predicate.

- predicate: What the entity must be.

## minecraft:loot_condition/entity_scores
Passes when an entity's scoreboard scores are in the given ranges.

- scores: A map from objective name to a range of scores.

## minecraft:loot_condition/location_check
Passes when the position, shifted by the offsets, matches a location predicate: biome, structure, light, block and more.

- predicate: What the location must be.
- offsetX: Blocks added to x before checking.
- offsetY: Blocks added to y before checking.
- offsetZ: Blocks added to z before checking.

## minecraft:loot_condition/match_block
Passes when the block, such as the one being broken, matches.

- blocks: The blocks or block tag that pass.
- state: Block states the block must have, such as `age = "7"`.

## minecraft:loot_condition/damage_source_properties
Passes when the damage that killed the entity matches a damage source predicate.

## minecraft:loot_condition/table_bonus
Passes at random, with the chance picked by the level of an enchantment on the tool. Vanilla uses it for saplings and flint.

- enchantment: The enchantment whose level picks the chance.

## minecraft:loot_condition/int_value_check
Passes when a whole number is in a range.

## minecraft:loot_condition/float_value_check
Passes when a number is in a range.

## minecraft:loot_condition/enchantment_active_check
Inside an enchantment's effects, passes when the enchantment is active (`true`) or inactive (`false`).

## minecraft:loot_function/set_count
Sets the stack size, or adds to it with `add`.

- count: The count: a number, or a range such as `1..3` for a random count.

```mms
ItemModifier few = SetCount(1..3)
```

## minecraft:loot_function/set_damage
Sets how damaged a tool or armour piece is.

## minecraft:loot_function/set_name
Sets the item's name.

- name: The name, as a text component.

## minecraft:loot_function/set_lore
Sets or changes the item's lore lines.

- lore: The lore lines, as text components.

## minecraft:loot_function/set_components
Sets data components on the item, such as `max_stack_size` or `rarity`.

- components: The components to set. A key starting with `!` removes that component.

## minecraft:loot_function/set_custom_data
Sets the item's `custom_data` component.

## minecraft:loot_function/set_enchantments
Sets enchantment levels on the item.

## minecraft:loot_function/set_item
Changes the item to another one, keeping count and components.

- item: The new item.

## minecraft:loot_function/set_potion
Sets the potion of a potion, splash potion, lingering potion or tipped arrow.

## minecraft:loot_function/set_loot_table
Sets the loot table a container item, such as a chest, rolls when placed and opened.

## minecraft:loot_function/set_attributes
Adds attribute modifiers, such as attack damage, to the item.

## minecraft:loot_function/apply_bonus
Raises the count by the level of an enchantment on the tool, as fortune does for ores.

- enchantment: The enchantment whose level raises the count.
- formula: How the level changes the count: `ore_drops`, `uniform_bonus_count` or `binomial_with_bonus_count`.

## minecraft:loot_function/enchanted_count_increase
Adds to the count by the level of an enchantment on the attacker's weapon, as looting does.

## minecraft:loot_function/enchant_randomly
Adds one random enchantment.

## minecraft:loot_function/enchant_with_levels
Enchants the item as an enchanting table would at this many levels.

## minecraft:loot_function/explosion_decay
Removes some of the stack when an explosion caused the drop, as vanilla blocks do.

## minecraft:loot_function/furnace_smelt
Replaces the item with what smelting it in a furnace gives, as vanilla does for mobs killed by fire.

## minecraft:loot_function/limit_count
Keeps the count within a range.

## minecraft:loot_function/copy_name
Copies the custom name of an entity or block entity onto the item.

## minecraft:loot_function/copy_components
Copies data components from a block entity, entity or item onto the item.

## minecraft:loot_function/copy_custom_data
Copies NBT from an entity, block entity or storage into the item's `custom_data`.

## minecraft:loot_function/copy_state
Copies block state properties of the broken block onto the item.

## minecraft:loot_function/fill_player_head
Makes a player head show an entity's skin.

## minecraft:loot_function/exploration_map
Turns an empty map into an explorer map that points to the nearest structure of a kind.

## minecraft:loot_function/filtered
Applies `on_pass` to items that match `item_filter` and `on_fail` to those that don't.

## minecraft:loot_function/sequence
Applies several item modifiers in order.

## minecraft:loot_function/discard
Removes the item.

## minecraft:loot_function/set_contents
Fills a container item, such as a shulker box or bundle, with entries.

## minecraft:loot_function/modify_contents
Applies an item modifier to every item inside a container item.

## minecraft:loot_function/toggle_tooltips
Shows or hides tooltip parts of the item, such as enchantments or attribute modifiers.
