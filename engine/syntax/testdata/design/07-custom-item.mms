/** A vanilla item with its own look and name. */
Blueprint CustomItem {
  /** The vanilla item it is made from, which decides how it behaves. */
  base: Id<item>
  /** Its name in English, or one per language. */
  name: LangText?
  /** The picture for a flat item model. */
  texture: Texture?
  parent: Model = item/generated
  model: Model?
  components: ItemComponents = {}
  on_right_click: Function?
  on_hit: Function?

  let lang_key = "item.${self.namespace}.${self.path}"
  let clicks = on_right_click != none
  let item_components = components + {
    item_model = self.id
    item_name = if (name != none) Translate(lang_key, fallback = english(name))
    consumable = if (clicks && components.consumable == none) never_finished_consumable
    custom_data = if (clicks || on_hit != none) { mms_item = "${self.id}" }
  }

  check {
    if (texture != none && model != none) {
      error(model, "Set texture or model, not both")
    }
    if (clicks && components.consumable != none) {
      warn(on_right_click, "right-click", "A consumable of your own changes how the click is detected")
    }
  }

  /** In a recipe result, a loot entry, or ${...} in a command: the item with its components. */
  as ItemStack = { id = base, components = item_components }
  as LootPoolEntry = Item(name = base, functions = [SetComponents(item_components)])
  /** Ingredients and tags match by id only. */
  as Id<item> = base hint("custom-item-base", "Matches any ${base}, not only ${self.path}")

  emit {
    if (model == none && texture != none) {
      Model item/${self.path} { parent = parent, textures = { layer0 = texture } }
    }
    ItemModel ${self.path} { model = Plain(model ?? item/${self.path}) }
    for lang, text in by_language(name) {
      Lang ${lang} { "${lang_key}" = text }
    }
    if (clicks) {
      RightClickHook mms_gen/item_hook/${self.path}/right_click { item = self, run = on_right_click }
    }
    if (on_hit != none) {
      HitHook mms_gen/item_hook/${self.path}/hit { item = self, run = on_hit }
    }
  }
}
