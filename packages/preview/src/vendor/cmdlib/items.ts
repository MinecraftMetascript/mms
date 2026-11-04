import { isCompound, isList, isNumeric, isString, type NbtTag, stringifySnbt } from "./nbt";
import { parseSnbt } from "./snbt";
import type { InventoryItem } from "./player";

export const HOTBAR_SIZE = 9;
export const MAIN_INVENTORY_SIZE = 27;
export const ENDERCHEST_SIZE = 27;
export const ARMOR_SLOTS = ["armor.head", "armor.chest", "armor.legs", "armor.feet"] as const;

export function normalizeItemId(id: string): string {
  return id.includes(":") ? id : `minecraft:${id}`;
}

export function normalizeComponentKey(key: string): string {
  return key.includes(":") ? key : `minecraft:${key}`;
}

function splitComponents(raw: string): string[] {
  const body = raw.trim().replace(/^\[/, "").replace(/\]$/, "");
  const parts: string[] = [];
  let depth = 0;
  let quote: string | null = null;
  let start = 0;
  for (let i = 0; i < body.length; i++) {
    const ch = body[i];
    if (quote) {
      if (ch === "\\") i++;
      else if (ch === quote) quote = null;
      continue;
    }
    if (ch === '"' || ch === "'") quote = ch;
    else if (ch === "{" || ch === "[") depth++;
    else if (ch === "}" || ch === "]") depth--;
    else if (ch === "," && depth === 0) {
      parts.push(body.slice(start, i).trim());
      start = i + 1;
    }
  }
  const tail = body.slice(start).trim();
  if (tail) parts.push(tail);
  return parts.filter(Boolean);
}

/** Parse `[custom_data={...},item_name="..." ]` into component tags. */
export function parseItemComponents(raw: string): Record<string, NbtTag> {
  const components: Record<string, NbtTag> = {};
  for (const part of splitComponents(raw)) {
    const equals = part.indexOf("=");
    if (equals <= 0) continue;
    const key = normalizeComponentKey(part.slice(0, equals).trim());
    const value = part.slice(equals + 1).trim();
    try {
      components[key] = parseSnbt(value);
    } catch {
      // Invalid components are caught by the command parser's caller.
    }
  }
  return components;
}

export function itemStackKey(item: InventoryItem): string {
  return [
    item.id,
    item.nbt ? stringifySnbt(item.nbt, { typed: true, quoteKeys: true }) : "",
    item.components ? stringifySnbt({ type: "compound", value: item.components }, { typed: true, quoteKeys: true }) : "",
  ].join("|");
}

export function itemComponentsMatch(
  actual: Record<string, NbtTag> | undefined,
  expected: Record<string, NbtTag>,
): boolean {
  if (!actual) return Object.keys(expected).length === 0;
  const matches = (left: NbtTag, right: NbtTag): boolean => {
    if (isCompound(left) && isCompound(right)) {
      return Object.entries(right.value).every(
        ([key, value]) =>
          left.value[key] !== undefined && matches(left.value[key], value),
      );
    }
    if (isList(left) && isList(right)) {
      return (
        left.value.length === right.value.length &&
        left.value.every((entry, index) => matches(entry, right.value[index]))
      );
    }
    if (isString(left) && isString(right)) return left.value === right.value;
    if (isNumeric(left) && isNumeric(right)) {
      return Number(left.value) === Number(right.value);
    }
    return false;
  };
  return Object.entries(expected).every(
    ([key, value]) => actual[key] !== undefined && matches(actual[key], value),
  );
}

export function itemComponentSummary(item: InventoryItem): string {
  const keys = Object.keys(item.components ?? {});
  if (item.nbt) keys.push("nbt");
  return keys.join(", ");
}

export function formatSlot(slot: string): string {
  return slot
    .replace(/^weapon\./, "")
    .replace(/^armor\./, "armor ")
    .replace(/^inventory\./, "inv ")
    .replace(/^hotbar\./, "hotbar ")
    .replace(/^enderchest\./, "ender ");
}

export function slotList(): string[] {
  return [
    "weapon.mainhand",
    "weapon.offhand",
    "armor.head",
    "armor.chest",
    "armor.legs",
    "armor.feet",
    ...Array.from({ length: HOTBAR_SIZE }, (_, i) => `hotbar.${i}`),
    ...Array.from({ length: MAIN_INVENTORY_SIZE }, (_, i) => `inventory.${i}`),
    ...Array.from({ length: ENDERCHEST_SIZE }, (_, i) => `enderchest.${i}`),
  ];
}

export interface ResolvedSlot {
  storage: "inventory" | "enderchest";
  key: string;
}

export function resolveSlot(slot: string, selectedSlot: number): ResolvedSlot | null {
  const value = slot.replace(/^container\./, "inventory.");
  if (value === "weapon.mainhand") {
    return { storage: "inventory", key: `hotbar.${selectedSlot}` };
  }
  if (value === "weapon.offhand") {
    return { storage: "inventory", key: "weapon.offhand" };
  }
  if (ARMOR_SLOTS.includes(value as (typeof ARMOR_SLOTS)[number])) {
    return { storage: "inventory", key: value };
  }
  const hotbar = /^hotbar\.(\d+)$/.exec(value);
  if (hotbar) {
    const index = Number(hotbar[1]);
    return index < HOTBAR_SIZE
      ? { storage: "inventory", key: `hotbar.${index}` }
      : null;
  }
  const inventory = /^inventory\.(\d+)$/.exec(value);
  if (inventory) {
    const index = Number(inventory[1]);
    return index < MAIN_INVENTORY_SIZE
      ? { storage: "inventory", key: `inventory.${index}` }
      : null;
  }
  const ender = /^enderchest\.(\d+)$/.exec(value);
  if (ender) {
    const index = Number(ender[1]);
    return index < ENDERCHEST_SIZE
      ? { storage: "enderchest", key: `enderchest.${index}` }
      : null;
  }
  return null;
}
