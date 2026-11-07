// Using a barrel: right-click the front to put items in, punch it to take
// them out. Both are clicks on the barrel's interaction entity, which the
// game reports with two advancement triggers. Both triggers only match
// entities tagged barrels.hit, so other interaction entities are left
// alone.

/** Whether the player is sneaking. */
Predicate sneaking = EntityProperties(entity = this, predicate = { flags = { is_sneaking = true } })

/**
 * Right click puts the held stack in; an empty barrel takes the kind of
 * whatever goes in first. Sneaking and right-clicking puts in every
 * matching item from the hotbar and inventory instead.
 */
OnAdvancement use_barrel = PlayerInteractedWithEntity(
  entity = EntityProperties(
    entity = this,
    predicate = {
      entity_type = interaction
      entity_tags = { all_of = ["barrels.hit"] }
    },
  ),
) {
  /tag @s add barrels.user
  /execute as @e[type=interaction,tag=barrels.hit,distance=..12] if data entity @s interaction if function barrels:used_by_user run tag @s add barrels.this
  /execute as @e[type=interaction,tag=barrels.this] run data remove entity @s interaction
  Execute("if entity @e[type=interaction,tag=barrels.this]") {
    open()
    /execute at @s run summon item_display ~ ~ ~ {Tags:["barrels.tmp"]}
    /execute store success score $holding mms.barrels if items entity @s weapon.mainhand *
    /execute store success score $sneaking mms.barrels if predicate barrels:sneaking
    if (sneaking == 1) {
      // An empty barrel takes the kind of the held item first, if any.
      if (holding == 1 && stack == 0) {
        claim()
      }
      insert_all()
    } else if (holding == 1) {
      insert_held()
    } else {
      actionbar(@s, "{gray}Right-click with an item to put it in, or sneak to put in everything that matches")
    }
    /kill @e[type=item_display,tag=barrels.tmp]
    close()
  }
  /tag @s remove barrels.user
}

/**
 * Punching takes a stack out; sneaking and punching takes one item. They
 * drop at the player's feet and are picked up straight away, or stay on
 * the ground when the inventory is full.
 */
OnAdvancement punch_barrel = PlayerHurtEntity(
  entity = EntityProperties(
    entity = this,
    predicate = {
      entity_type = interaction
      entity_tags = { all_of = ["barrels.hit"] }
    },
  ),
) {
  /tag @s add barrels.user
  /execute as @e[type=interaction,tag=barrels.hit,distance=..12] if data entity @s attack if function barrels:punched_by_user run tag @s add barrels.this
  /execute as @e[type=interaction,tag=barrels.this] run data remove entity @s attack
  Execute("if entity @e[type=interaction,tag=barrels.this]") {
    open()
    take_out()
    close()
  }
  /tag @s remove barrels.user
}

/** Runs as an interaction entity: whether the player tagged barrels.user right-clicked it last. */
Function used_by_user {
  /return run execute on target if entity @s[tag=barrels.user]
}

/** Runs as an interaction entity: whether the player tagged barrels.user punched it last. */
Function punched_by_user {
  /return run execute on attacker if entity @s[tag=barrels.user]
}

/** Puts the player's held stack in, as much of it as fits. */
Function insert_held {
  if (stack == 0) {
    claim()
  }
  if (stack > 0) {
    absorb("weapon.mainhand")
    if (diff == 1) {
      actionbar(@s, "{red}This barrel holds something else")
    }
  }
}

/** Makes the empty barrel hold the kind of item in the player's hand. */
Function claim {
  /execute if items entity @s weapon.mainhand *[max_stack_size=64] run scoreboard players set $stack mms.barrels 64
  /execute if items entity @s weapon.mainhand *[max_stack_size=16] run scoreboard players set $stack mms.barrels 16
  if (stack == 0) {
    actionbar(@s, "{red}Only items that stack to 16 or 64 fit in a barrel")
    return
  }
  /item replace entity @e[type=item_display,tag=barrels.this,limit=1] contents from entity @s weapon.mainhand
  /item modify entity @e[type=item_display,tag=barrels.this,limit=1] contents barrels:one
}

/** Puts in every item in the player's hotbar and inventory that matches the barrel. */
Function insert_all {
  if (stack == 0) {
    actionbar(@s, "{gray}Right-click with an item to choose what this barrel holds")
    return
  }
  let before = count
  for i in 0..8 {
    absorb("hotbar." + i)
  }
  for i in 0..26 {
    absorb("inventory." + i)
  }
  if (count == before) {
    actionbar(@s, "{gray}Nothing in your inventory matches")
  }
}

/**
 * Moves the stack in one of the player's slots into the barrel, as much
 * as fits, if it is the same item: same id and same components. diff is 1
 * afterwards when the slot held something else.
 */
Function absorb(slot: string) {
  diff = 0
  held = 0
  // Copy the stack to the scratch item display, note its count, then make
  // it a single item so that it compares equal to the stored one. Setting
  // the copy from the stored item reports success only if they differed.
  /item replace entity @e[type=item_display,tag=barrels.tmp,limit=1] contents from entity @s ${slot}
  /execute store result score $held mms.barrels run data get entity @e[type=item_display,tag=barrels.tmp,limit=1] item.count
  if (held == 0) {
    return
  }
  /item modify entity @e[type=item_display,tag=barrels.tmp,limit=1] contents barrels:one
  /execute store success score $diff mms.barrels run data modify entity @e[type=item_display,tag=barrels.tmp,limit=1] item set from entity @e[type=item_display,tag=barrels.this,limit=1] item
  if (diff == 1) {
    return
  }
  let take = capacity() - count
  if (take <= 0) {
    actionbar(@s, "{gold}This barrel is full")
    return
  }
  if (take > held) {
    take = held
  }
  count += take
  minus = -take
  /item modify entity @s ${slot} barrels:remove_amount
  /playsound minecraft:item.bundle.insert player @a
}

/** Takes a stack, or one item when sneaking, and drops it at the player's feet. */
Function take_out {
  if (count == 0) {
    return
  }
  amount = stack
  /execute if predicate barrels:sneaking run scoreboard players set $amount mms.barrels 1
  if (amount > count) {
    amount = count
  }
  count -= amount
  if (count == 0) {
    stack = 0
  }
  /execute at @s run summon item ~ ~ ~ {Item:{id:"minecraft:stone"},PickupDelay:0s,Tags:["barrels.drop"]}
  give_drop()
  /playsound minecraft:item.bundle.remove_one player @a
}
