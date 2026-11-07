// CustomItem and the item hooks: a vanilla item with its own look, name
// and behaviour, written in MMS on top of models, item models, lang
// files, advancements and functions.

/**
 * A vanilla item with its own look and name. It writes the model drawn
 * from its texture, the item model that its item_model component names
 * and its name in the lang files. Its name then stands for the base item
 * with those components wherever an item stack goes: recipe results,
 * loot entries and `${...}` in commands.
 */
Blueprint CustomItem {
  /** The vanilla item this one is made from. It keeps the base item's behaviour; recipes and item tests that ask for an item id match the base. */
  base: Id<item>
  /** The item's name: text for `en_us`, or one text per language such as `{ en_us = "Ember Shard", de_de = "Glutsplitter" }`. It becomes an `item.<namespace>.<name>` entry in the lang files and the item's `item_name` component. */
  name: LangText?
  /** The item's picture, usually `Texture("art/name.png")`, which is copied to textures/item/<name>.png. It makes a flat item model, like most vanilla items. */
  texture: Texture?
  /** The parent of the model drawn from `texture`: `item/generated` (the default) for flat items, `item/handheld` for tools held like a sword. */
  parent: Id<model>?
  /** An existing model to show instead of one drawn from `texture`, such as a `Model` declared in the project. */
  model: Id<model>?
  /** An item model to use instead of the generated one: the id of an `ItemModel`, or its body, such as `{ model = Select(...) }` for an item that changes its look. */
  item_model: ItemModel?
  /** More item components, such as `max_stack_size` or `rarity`. Every use of the item gets them. */
  components: DataComponentPatch = {}
  /** How many of the item each right-click use takes from the stack, paid once per press by on_right_click, on_use_on_block or on_right_click_entity. */
  consume_on_use: int?
  /** How much durability each right-click use costs, paid like consume_on_use. The item needs a `max_damage` component, and gets `max_stack_size = 1` unless it sets one. When the durability runs out the item breaks, and on_break runs. */
  damage_on_use: int?
  /** A `Function` without parameters to run, as and at the player, when they right-click with the item: once per press, not again while the button is held. When the item also has on_use_on_block or on_right_click_entity, a press that runs one of those does not run this one too, so this is the click at anything else, such as the air. */
  on_right_click: Function?
  /** A `Function` without parameters to run, as and at the player, on every tick they hold right click with the item. */
  on_hold_right_click: Function?
  /** A `Function` without parameters to run when a player right-clicks a block with the item, once per press: as the player, positioned at the centre of the block. */
  on_use_on_block: Function?
  /** A `Function` without parameters to run when a player right-clicks an entity with the item, once per press: as the player, at the entity, which has the tag `mms_target` while it runs. */
  on_right_click_entity: Function?
  /** A `Function` without parameters to run, as and at the player, when they hurt an entity while holding the item in their main hand. */
  on_hit: Function?
  /** A `Function` without parameters to run, as and at the player, when they kill an entity while holding the item in their main hand. */
  on_kill: Function?
  /** A `Function` without parameters to run, as and at the player, when they finish eating or drinking the item. The item needs a `consumable` component, or a base item that is food or a potion. It cannot be combined with the right-click hooks. */
  on_consume: Function?
  /** A `Function` without parameters to run, as and at the player, when they pick the item up from the ground. */
  on_pickup: Function?
  /** A `Function` without parameters to run, as and at the player, when the item runs out of durability and breaks. The item needs durability: a tool, weapon or armor base, or a `max_damage` component. */
  on_break: Function?
  /** A `Function` without parameters to run when an arrow hits a block: an arrow shot from the item, for a bow or crossbow base, or the arrow that is the item, for an arrow, tipped_arrow or spectral_arrow base. It runs as the arrow, at the point it hit; `execute on origin` is the shooter. */
  on_projectile_hit_block: Function?
  /** A `Function` without parameters to run when an arrow hurts an entity: an arrow shot from the item, for a bow or crossbow base, or the arrow that is the item, for an arrow, tipped_arrow or spectral_arrow base. It runs as and at the entity hit; `execute on attacker` is the shooter. */
  on_projectile_hit_entity: Function?
  /** A `Function` without parameters to run where a thrown splash or lingering potion that is the item lands, whatever it hits. It runs at the spot, as a marker that is removed afterwards. Needs a splash_potion or lingering_potion base. */
  on_projectile_land: Function?

  let dotted = replace(self.path, "/", ".")
  let lang_key = "item.${self.namespace}.${dotted}"
  let draws_model = model == none && item_model == none && texture != none
  let item_model_id = if (item_model != none || model != none || texture != none) self.id
  let presses = on_right_click != none || on_use_on_block != none || on_right_click_entity != none
  let holds = presses || on_hold_right_click != none
  let projectile = on_projectile_hit_block != none || on_projectile_hit_entity != none
  let hooked = holds || projectile || on_projectile_land != none || on_hit != none || on_kill != none || on_consume != none || on_pickup != none || on_break != none
  let costs = consume_on_use != none || damage_on_use != none
  let weapons = [bow, crossbow]
  let ammo = [arrow, tipped_arrow, spectral_arrow]
  let thrown = [splash_potion, lingering_potion]

  let item_components = {
    item_model = item_model_id
    item_name = if (name != none) { translate = lang_key, fallback = english(name) }
  } + components + {
    max_stack_size = if (damage_on_use != none && components.max_stack_size == none) 1
    // About eleven days: right click "uses" the item without ever
    // finishing, and without the eating sounds and particles.
    consumable = if (holds && components.consumable == none) {
      consume_seconds = 1000000
      animation = none
      has_consume_particles = false
    }
    use_effects = if (holds && components.use_effects == none) {
      can_sprint = true
      speed_multiplier = 1
      interact_vibrations = false
    }
    custom_data = if (hooked) (components.custom_data ?? {}) + { mms_item = "${self.id}" }
  }

  check {
    for lang, text in by_language(name) {
      if (!matches(lang, "[a-z]{2,3}_[a-z0-9]{2,3}")) {
        error(name, "'${lang}' is not a language code such as en_us or de_de")
      }
    }
    if (item_model != none && (texture != none || model != none || parent != none)) {
      error(item_model, "item_model replaces the generated models, so texture, model and parent have no effect; leave out one or the other")
    } else if (model != none && (texture != none || parent != none)) {
      error(model, "model names the model to show, so texture and parent have no effect; leave out one or the other")
    } else if (parent != none && texture == none) {
      error(parent, "parent is the parent of the model drawn from texture; add a texture, or leave it out")
    }
    if (consume_on_use != none && consume_on_use < 1) {
      error(consume_on_use, "consume_on_use is at least 1")
    }
    if (damage_on_use != none && damage_on_use < 1) {
      error(damage_on_use, "damage_on_use is at least 1")
    }
    if (costs && !presses) {
      error(if (consume_on_use != none) consume_on_use else damage_on_use, "the use cost is paid on each right-click use; add on_right_click, on_use_on_block or on_right_click_entity")
    }
    if (damage_on_use != none && components.max_damage == none) {
      error(damage_on_use, "damage_on_use needs the item's durability: add max_damage to its components, such as components = { max_damage = 32 }")
    }
    if (damage_on_use != none && components.max_stack_size != none && components.max_stack_size != 1) {
      error(components, "an item with durability cannot stack: the game requires max_stack_size = 1 with max_damage")
    }
    if (projectile && !contains(weapons, base) && !contains(ammo, base)) {
      error(base, "on_projectile_hit_block and on_projectile_hit_entity follow arrows, so the base must be a bow or crossbow (for the arrows it shoots) or an arrow, tipped_arrow or spectral_arrow (for itself)")
    }
    if (on_projectile_land != none && !contains(thrown, base)) {
      error(on_projectile_land, "on_projectile_land follows a thrown potion, so the base must be splash_potion or lingering_potion")
    }
    if (hooked && components.custom_data != none && kind_of(components.custom_data) != "object") {
      error(components, "the on_* fields mark the item in its custom_data, which must be written as an object such as { kind = \"wand\" }")
    }
    if (on_consume != none && holds) {
      error(on_consume, "on_consume needs the item to be eaten, but the right-click hooks make right click use it without ever finishing; leave out one or the other")
    }
  }

  /** In a recipe result, a loot entry, or `${...}` in a command: the base item with the custom item's components. */
  as ItemStackTemplate = { id = base, components = item_components }
  as ItemStack = { id = base, components = item_components }
  as LootPoolEntry = Item(name = base, modifier = [{ type = set_components, components = item_components }])
  /** A CustomItem used where the game takes a plain item id, such as a recipe ingredient or an item tag. It stands for its base item there, so it also matches the base item without the custom components. */
  as Id<item> = base hint("custom-item-as-base", "'${self.path}' stands for its base item ${base} here: this field takes a plain item id, so any ${base} matches, not only ${self.path}")

  emit {
    if (draws_model) {
      Model item/${self.path} {
        parent = parent ?? item/generated
        textures = { layer0 = named(texture, "item/${self.path}") }
      }
    }
    if (item_model == none && (model != none || texture != none)) {
      ItemModel ${self.path} { model = { type = "minecraft:model", model = model ?? item/${self.path} } }
    }
    for lang, text in by_language(name) {
      Lang ${lang} { "${lang_key}" = text }
    }
    if (holds) {
      PressHook mms_gen/item_hook/${self.path}/press {
        item = self.id
        base = base
        hold = on_hold_right_click
        click = on_right_click
        block = on_use_on_block
        entity = on_right_click_entity
        consume = consume_on_use ?? 0
        damage = damage_on_use ?? 0
        max_damage = components.max_damage ?? 0
        on_break = on_break
      }
    }
    if (on_hit != none) {
      HeldHook mms_gen/item_hook/${self.path}/hit { item = self.id, base = base, trigger = player_hurt_entity, action = on_hit }
    }
    if (on_kill != none) {
      HeldHook mms_gen/item_hook/${self.path}/kill { item = self.id, base = base, trigger = player_killed_entity, action = on_kill }
    }
    if (on_consume != none) {
      ItemHook mms_gen/item_hook/${self.path}/consume { item = self.id, base = base, trigger = consume_item, action = on_consume }
    }
    if (on_pickup != none) {
      ItemHook mms_gen/item_hook/${self.path}/pickup { item = self.id, base = base, trigger = thrown_item_picked_up_by_player, action = on_pickup }
    }
    if (on_break != none) {
      ItemHook mms_gen/item_hook/${self.path}/break { item = self.id, base = base, trigger = item_durability_changed, broken = true, action = on_break }
    }
    if (projectile) {
      ArrowHook mms_gen/item_hook/${self.path}/projectile {
        item = self.id
        slot = if (contains(ammo, base)) "item" else "weapon"
        hit_block = on_projectile_hit_block
        hit_entity = on_projectile_hit_entity
      }
    }
    if (on_projectile_land != none) {
      LandHook mms_gen/item_hook/${self.path}/land { item = self.id, action = on_projectile_land }
    }
  }
}

/**
 * An advancement that runs a function, as and at the player, each time
 * the player meets its trigger with a custom item, and is revoked so it
 * fires again.
 *
 * @internal
 */
Blueprint ItemHook {
  /** The custom item's id. */
  item: string
  /** Its base item. */
  base: Id<item>
  /** The advancement trigger, which tests the item in its `item` condition. */
  trigger: Id<trigger_type>
  /** Whether the hook runs when the item breaks. */
  broken: bool = false
  action: Function

  emit {
    primary Advancement ${self.path} {
      criteria = {
        "${trigger}" = {
          trigger = trigger
          conditions = {
            item = { items = base, predicates = { "minecraft:custom_data" = { mms_item = item } } }
            durability = if (broken) { max = 0 }
          }
        }
      }
      rewards = { function = self.id }
    }
    Function ${self.path} {
      /advancement revoke @s only ${self.id}
      action()
    }
  }
}

/**
 * An ItemHook whose trigger tests the player's main hand.
 *
 * @internal
 */
Blueprint HeldHook {
  item: string
  base: Id<item>
  trigger: Id<trigger_type>
  action: Function

  emit {
    primary Advancement ${self.path} {
      criteria = {
        "${trigger}" = {
          trigger = trigger
          conditions = {
            player = {
              type = entity_properties
              entity = this
              predicate = { equipment = { mainhand = { items = base, predicates = { "minecraft:custom_data" = { mms_item = item } } } } }
            }
          }
        }
      }
      rewards = { function = self.id }
    }
    Function ${self.path} {
      /advancement revoke @s only ${self.id}
      action()
    }
  }
}

/**
 * The right-click hooks, after datapack.wiki/guide/right-click: the
 * item's consumable component never finishes, so holding right click
 * "uses" it every tick, which a using_item advancement sees. A per-player
 * score, set to 2 on each tick of use and counted down each tick, makes
 * the press hooks run once per press; a second score names the item, so
 * the countdown of one item does not swallow the press of another. The
 * block and entity hooks cast a ray from the player's eyes, as far as
 * their interaction ranges. A press runs one hook: the block or entity
 * hook when the ray reaches what it hooks, otherwise the click hook.
 *
 * @internal
 */
Blueprint PressHook {
  item: string
  base: Id<item>
  hold: Function?
  click: Function?
  block: Function?
  entity: Function?
  consume: int = 0
  damage: int = 0
  max_damage: int = 0
  on_break: Function?

  let obj = "mms_click.${self.namespace}"
  // Which item the player last pressed with, so holding the button while
  // switching to another custom item counts as a new press of that one.
  let held = "mms_click_item.${self.namespace}"
  let item_hash = hash(item)
  let marker = "${base}[custom_data~{mms_item:\"${item}\"}]"
  let pays = consume > 0 || damage > 0
  // set_damage works in fractions of max_damage and rounds the new
  // damage down, so half a point more keeps float error from costing
  // one point less.
  let fraction = if (max_damage > 0) (damage + 0.5) / max_damage else 0
  let ray = block != none || entity != none

  emit {
    primary Advancement ${self.path} {
      criteria = {
        press = {
          trigger = using_item
          conditions = { item = { items = base, predicates = { "minecraft:custom_data" = { mms_item = item } } } }
        }
      }
      rewards = { function = self.id }
    }
    Function ${self.path} {
      /advancement revoke @s only ${self.id}
      if (hold != none) {
        hold()
      }
      if (click != none || ray) {
        // Still held since the last press with this item: keep the score
        // up and stop.
        /execute if score @s ${obj} matches 1.. if score @s ${held} matches ${item_hash} run return run scoreboard players set @s ${obj} 2
        /scoreboard players set @s ${obj} 2
        /scoreboard players set @s ${held} ${item_hash}
      }
      if (ray) {
        /tag @s add mms_clicker
        /execute store result score #block ${obj} run attribute @s block_interaction_range get 10
        /execute store result score #entity ${obj} run attribute @s entity_interaction_range get 10
        /scoreboard players set #ray ${obj} 0
        if (click != none) {
          // Set to 1 by the block or entity hook when it runs.
          /scoreboard players set #hit ${obj} 0
        }
        // Back to feet anchoring, or each step would add the eye height.
        /execute anchored eyes positioned ^ ^ ^ anchored feet run function ${self.id}/ray
        /tag @s remove mms_clicker
      }
      if (click != none) {
        if (ray) {
          /execute if score #hit ${obj} matches 0 run function ${self.id}/click
        } else {
          click()
          if (pays) {
            /function ${self.id}/cost
          }
        }
      }
    }
    if (click != none && ray) {
      Function ${self.path}/click {
        click()
        if (pays) {
          /function ${self.id}/cost
        }
      }
    }
    if (click != none || ray) {
      Objective mms_click.${self.namespace}
      Objective mms_click_item.${self.namespace}
      OnTick {
        /scoreboard players remove @a[scores={${obj}=1..}] ${obj} 1
      }
    }
    if (pays) {
      Function ${self.path}/cost {
        /execute if items entity @s weapon.mainhand ${marker} run return run function ${self.id}/cost/mainhand
        /execute if items entity @s weapon.offhand ${marker} run function ${self.id}/cost/offhand
      }
      for hand in ["mainhand", "offhand"] {
        Function ${self.path}/cost/${hand} {
          if (consume > 0) {
            /item modify entity @s weapon.${hand} {type:"minecraft:set_count",count:${0 - consume},add:true}
          }
          if (damage > 0) {
            /item modify entity @s weapon.${hand} {type:"minecraft:set_damage",damage:${0 - (if (fraction > 1) 1 else fraction)},add:true}
            /execute if items entity @s weapon.${hand} *[damage~{durability:{max:0}}] run function ${self.id}/cost/${hand}/break
          }
        }
        if (damage > 0) {
          Function ${self.path}/cost/${hand}/break {
            /item replace entity @s weapon.${hand} with air
            /playsound entity.item.break player @a
            if (on_break != none) {
              on_break()
            }
          }
        }
      }
    }
    if (ray) {
      Function ${self.path}/ray {
        /scoreboard players add #ray ${obj} 1
        // An entity whose hitbox holds this point: its box meets both
        // the unit box below the point and the one above it.
        if (entity != none) {
          /execute if score #ray ${obj} <= #entity ${obj} positioned ~-0.99 ~-0.99 ~-0.99 as @e[dx=0,dy=0,dz=0,tag=!mms_clicker,type=!item,type=!experience_orb,type=!marker,type=!area_effect_cloud,type=!arrow,type=!spectral_arrow,type=!trident,type=!block_display,type=!item_display,type=!text_display,type=!lightning_bolt] positioned ~0.99 ~0.99 ~0.99 if entity @s[dx=0,dy=0,dz=0] run return run function ${self.id}/ray/entity
        } else {
          /execute if score #ray ${obj} <= #entity ${obj} positioned ~-0.99 ~-0.99 ~-0.99 as @e[dx=0,dy=0,dz=0,tag=!mms_clicker,type=!item,type=!experience_orb,type=!marker,type=!area_effect_cloud,type=!arrow,type=!spectral_arrow,type=!trident,type=!block_display,type=!item_display,type=!text_display,type=!lightning_bolt] positioned ~0.99 ~0.99 ~0.99 if entity @s[dx=0,dy=0,dz=0] run return 1
        }
        if (block != none) {
          /execute unless block ~ ~ ~ #minecraft:air unless block ~ ~ ~ water unless block ~ ~ ~ lava unless block ~ ~ ~ light unless block ~ ~ ~ structure_void unless block ~ ~ ~ bubble_column run return run execute align xyz positioned ~0.5 ~0.5 ~0.5 run function ${self.id}/ray/block
        } else {
          /execute unless block ~ ~ ~ #minecraft:air unless block ~ ~ ~ water unless block ~ ~ ~ lava unless block ~ ~ ~ light unless block ~ ~ ~ structure_void unless block ~ ~ ~ bubble_column run return 1
        }
        /execute if score #ray ${obj} < #block ${obj} positioned ^ ^ ^0.1 run function ${self.id}/ray
      }
      if (block != none) {
        Function ${self.path}/ray/block {
          if (click != none) {
            /scoreboard players set #hit ${obj} 1
          }
          block()
          if (pays) {
            /function ${self.id}/cost
          }
        }
      }
      if (entity != none) {
        Function ${self.path}/ray/entity {
          if (click != none) {
            /scoreboard players set #hit ${obj} 1
          }
          /tag @s add mms_target
          /execute at @s as @p[tag=mms_clicker] run function ${self.id}/ray/entity/run
          if (pays) {
            /execute as @p[tag=mms_clicker] run function ${self.id}/cost
          }
          /tag @s remove mms_target
        }
        Function ${self.path}/ray/entity/run {
          entity()
        }
      }
    }
  }
}

/**
 * The projectile hooks. An arrow keeps the bow or crossbow it was shot
 * from as its weapon, and when it hits, the game runs the weapon's
 * enchantment effects at the exact spot: hit_block as the arrow,
 * post_attack as the entity it hurt. So the item gets an enchantment
 * that runs its hooks, and the tick function adds it to the weapon of
 * each new arrow that is the item or was shot from it. It is never on
 * the item itself, so it shows in no tooltip and adds no glint.
 *
 * @internal
 */
Blueprint ArrowHook {
  item: string
  /** Where the arrow keeps the item: `weapon`, or `item` for an arrow that is the item. */
  slot: string
  hit_block: Function?
  hit_entity: Function?

  emit {
    primary Enchantment ${self.path} {
      description = { text = item }
      supported_items = [bow, crossbow]
      weight = 1
      max_level = 1
      min_cost = { base = 1, per_level_above_first = 0 }
      max_cost = { base = 1, per_level_above_first = 0 }
      anvil_cost = 0
      slots = [any]
      effects = {
        "minecraft:hit_block" = if (hit_block != none) [{ effect = { type = run_function, function = "${self.id}/hit_block" } }]
        "minecraft:post_attack" = if (hit_entity != none) [{ enchanted = attacker, affected = victim, effect = { type = run_function, function = "${self.id}/hit_entity" } }]
      }
    }
    if (hit_block != none) {
      Function ${self.path}/hit_block {
        hit_block()
      }
    }
    if (hit_entity != none) {
      Function ${self.path}/hit_entity {
        hit_entity()
      }
    }
    // An arrow from a dispenser has no weapon, so it gets a plain bow.
    Function ${self.path} {
      /execute unless data entity @s weapon run data modify entity @s weapon set value {id:"minecraft:bow",count:1}
      /data modify entity @s weapon.components."minecraft:enchantments"."${self.id}" set value 1
    }
    Function ${self.path}/spot {
      /execute if data entity @s ${slot}.components."minecraft:custom_data"{mms_item:"${item}"} run function ${self.id}
    }
    FunctionTag mms_gen/projectile/spotted = [${self.path}/spot]
    ProjectileSetup mms_gen/projectile { }
  }
}

/**
 * on_projectile_land. A thrown potion keeps no weapon, so no enchantment
 * runs when it lands. Instead the tick function puts a marker on each new
 * one, riding it. The game removes the potion when it lands and drops its
 * passengers where they are, so a marker that is no longer riding
 * anything marks the spot. The marker holds the function to run in its
 * data.
 *
 * @internal
 */
Blueprint LandHook {
  item: string
  action: Function

  let dotted = replace(replace(self.id, ":", "."), "/", ".")
  let marks = "mms_land.${dotted}"

  emit {
    primary Function ${self.path} {
      /tag @s add mms_new
      // The tick function runs this as the potion but not at it.
      /execute at @s summon marker run function ${self.id}/mount
      /tag @s remove mms_new
    }
    Function ${self.path}/mount {
      /tag @s add mms_land.${self.namespace}
      /tag @s add ${marks}
      /ride @s mount @n[tag=mms_new]
      // A marker that failed to mount would set the hook off right
      // where the potion was thrown.
      /execute unless predicate ${self.namespace}:mms_gen/projectile/riding run kill @s
    }
    Function ${self.path}/landed {
      /execute if entity @s[tag=${marks}] run function ${self.id}/run
    }
    Function ${self.path}/run {
      action()
    }
    Function ${self.path}/spot {
      /execute if data entity @s Item.components."minecraft:custom_data"{mms_item:"${item}"} run function ${self.id}
    }
    FunctionTag mms_gen/projectile/spotted = [${self.path}/spot]
    FunctionTag mms_gen/projectile/landed = [${self.path}/landed]
    ProjectileSetup mms_gen/projectile { thrown = true }
  }
}

/**
 * What a namespace with projectile hooks needs: a tick function that
 * finds each new arrow or thrown potion once.
 *
 * @internal
 */
Blueprint ProjectileSetup {
  thrown: bool = false

  let tag = "mms_shot.${self.namespace}"

  emit {
    Function ${self.path}/spotted_new {
      /tag @s add ${tag}
      /function #${self.namespace}:mms_gen/projectile/spotted
    }
    OnTick {
      /execute as @e[type=#minecraft:arrows,tag=!${tag}] run function ${self.id}/spotted_new
    }
    if (thrown) {
      OnTick {
        /execute as @e[type=splash_potion,tag=!${tag}] run function ${self.id}/spotted_new
        /execute as @e[type=lingering_potion,tag=!${tag}] run function ${self.id}/spotted_new
        /execute as @e[type=marker,tag=mms_land.${self.namespace},predicate=!${self.id}/riding] at @s run function ${self.id}/land
      }
      Predicate ${self.path}/riding { type = entity_properties, entity = this, predicate = { vehicle = {} } }
      Function ${self.path}/land {
        /function #${self.namespace}:mms_gen/projectile/landed
        /kill @s
      }
    }
  }
}
