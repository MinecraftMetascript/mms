import type { NbtTag } from "./nbt";
import { Tag } from "./nbt";
import { jsonToTag } from "./snbt";
import type { InventoryItem } from "./player";

export interface ItemModifierDefinition {
  id: string;
  path: string;
  function: string;
  data: Record<string, unknown>;
}

export interface ItemModifierError {
  path: string;
  message: string;
}

const MODIFIER_PATH = /^data\/([^/]+)\/item_modifiers?\/(.+)\.json$/;

function asComponent(value: unknown): NbtTag {
  return jsonToTag(value);
}

export function parseItemModifierFiles(files: Record<string, string>): {
  modifiers: Map<string, ItemModifierDefinition>;
  errors: ItemModifierError[];
} {
  const modifiers = new Map<string, ItemModifierDefinition>();
  const errors: ItemModifierError[] = [];
  for (const [path, content] of Object.entries(files)) {
    const match = MODIFIER_PATH.exec(path);
    if (!match) continue;
    const [, namespace, name] = match;
    try {
      const data = JSON.parse(content) as Record<string, unknown>;
      if (!data || typeof data !== "object" || Array.isArray(data)) {
        errors.push({ path, message: "Item modifier must be an object" });
        continue;
      }
      const fn = typeof data.function === "string" ? data.function.replace(/^minecraft:/, "") : "";
      if (!fn) {
        errors.push({ path, message: "Item modifier is missing a function" });
        continue;
      }
      modifiers.set(`${namespace}:${name}`, {
        id: `${namespace}:${name}`,
        path,
        function: fn,
        data,
      });
    } catch (error) {
      errors.push({ path, message: `Invalid item modifier JSON: ${(error as Error).message}` });
    }
  }
  return { modifiers, errors };
}

/** Apply a supported item modifier function to one stack. */
export function applyItemModifier(
  modifier: ItemModifierDefinition,
  item: InventoryItem,
): InventoryItem {
  const next: InventoryItem = {
    ...item,
    components: { ...(item.components ?? {}) },
  };
  const data = modifier.data;
  switch (modifier.function) {
    case "set_count": {
      const count = typeof data.count === "number" ? data.count : Number(data.count);
      if (Number.isFinite(count)) next.count = Math.max(1, Math.floor(count));
      break;
    }
    case "set_damage": {
      const damage = typeof data.damage === "number" ? data.damage : Number(data.damage);
      if (Number.isFinite(damage)) next.components!["minecraft:damage"] = Tag.int(damage);
      break;
    }
    case "set_name": {
      if (data.name !== undefined) {
        next.components!["minecraft:item_name"] = asComponent(data.name);
      }
      break;
    }
    case "set_lore": {
      if (Array.isArray(data.lore)) {
        next.components!["minecraft:lore"] = Tag.list(
          data.lore.map((entry) => asComponent(entry)),
          "end",
        );
      }
      break;
    }
    case "set_custom_data": {
      if (data.tag !== undefined) {
        next.components!["minecraft:custom_data"] = asComponent(data.tag);
      }
      break;
    }
    case "set_components": {
      if (data.components && typeof data.components === "object" && !Array.isArray(data.components)) {
        for (const [key, value] of Object.entries(data.components as Record<string, unknown>)) {
          next.components![key.includes(":") ? key : `minecraft:${key}`] = asComponent(value);
        }
      }
      break;
    }
    case "enchant_randomly": {
      next.components!["minecraft:enchantments"] = Tag.compound();
      break;
    }
    default:
      break;
  }
  return next;
}
