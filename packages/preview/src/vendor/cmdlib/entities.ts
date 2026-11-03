import { type CompoundTag, isCompound, isList, type NbtTag, stringifySnbt, Tag } from "./nbt";
import { parseTextComponent } from "./snbt";
import { plainText } from "./text";
import { entityInfo } from "./data";

export interface Vec3 {
  x: number;
  y: number;
  z: number;
}

export interface Vec2 {
  yaw: number;
  pitch: number;
}

export interface EntityData {
  uuid: string;
  type: string;
  nbt: CompoundTag;
  isPlayer: boolean;
  username?: string;
}

let uuidCounter = 0;

export function nextUuid(): string {
  if (typeof crypto !== "undefined" && "randomUUID" in crypto) {
    return crypto.randomUUID();
  }
  uuidCounter++;
  return `00000000-0000-4000-8000-${String(uuidCounter).padStart(12, "0")}`;
}

export function normalizeType(type: string): string {
  return type.includes(":") ? type : `minecraft:${type}`;
}

export function createEntity(
  type: string,
  pos: Vec3,
  overrides: CompoundTag | null = null,
  opts: { isPlayer?: boolean; username?: string } = {},
): EntityData {
  const fullType = normalizeType(type);
  const info = entityInfo(fullType);
  const base: Record<string, NbtTag> = {};
  const defaults = info.defaults ?? {
    Health: Tag.float(20),
    Tags: Tag.list([], "string"),
  };
  for (const [key, value] of Object.entries(defaults)) {
    base[key] = clone(value);
  }
  base.Pos = Tag.list([Tag.double(pos.x), Tag.double(pos.y), Tag.double(pos.z)], "double");
  if (opts.isPlayer) {
    base.Rotation = Tag.list([Tag.float(0), Tag.float(0)], "float");
    if (!base.Health) base.Health = Tag.float(20);
    if (!base.foodLevel) base.foodLevel = Tag.int(20);
  }
  // NBT supplied by a command (including Pos) is applied after the defaults.
  if (overrides) {
    for (const [key, value] of Object.entries(overrides.value)) {
      base[key] = clone(value);
    }
  }
  const nbtUuid = base.UUID;
  const uuid =
    nbtUuid && nbtUuid.type === "intArray" && nbtUuid.value.length >= 4
      ? uuidFromInts(nbtUuid.value)
      : nextUuid();
  if (!base.UUID) base.UUID = Tag.intArray(uuidToInts(uuid));
  const entity: EntityData = {
    uuid,
    type: fullType,
    nbt: Tag.compound(base),
    isPlayer: opts.isPlayer ?? false,
    username: opts.username,
  };
  if (entity.isPlayer && entity.username) {
    entity.nbt.value.CustomName = Tag.string(
      `{"text":${JSON.stringify(entity.username)}}`,
    );
  }
  return entity;
}

function clone(tag: NbtTag): NbtTag {
  return JSON.parse(JSON.stringify(tag)) as NbtTag;
}

export function uuidFromInts(ints: number[]): string {
  const hex = ints
    .slice(0, 4)
    .map((n) => (n >>> 0).toString(16).padStart(8, "0"))
    .join("");
  return `${hex.slice(0, 8)}-${hex.slice(8, 12)}-${hex.slice(12, 16)}-${hex.slice(16, 20)}-${hex.slice(20, 32)}`;
}

/**
 * Expand a canonical or abbreviated UUID (`91bb5-0-0-0-ffff`) to the
 * 8-4-4-4-12 form used by entity NBT.
 */
export function canonicalUuid(raw: string): string | null {
  const value = raw.trim().toLowerCase();
  if (!/^[0-9a-f-]+$/.test(value)) return null;
  let groups: string[];
  if (value.includes("-")) {
    groups = value.split("-");
    if (groups.length !== 5) return null;
  } else {
    if (value.length > 32) return null;
    const padded = value.padStart(32, "0");
    groups = [padded.slice(0, 8), padded.slice(8, 12), padded.slice(12, 16), padded.slice(16, 20), padded.slice(20, 32)];
  }
  const widths = [8, 4, 4, 4, 12];
  if (groups.length !== widths.length) return null;
  const expanded = groups.map((group, index) => {
    if (group.length === 0 || group.length > widths[index]) return null;
    return group.padStart(widths[index], "0");
  });
  if (expanded.some((group) => group === null)) return null;
  return expanded.join("-");
}

function uuidToInts(uuid: string): number[] {
  const hex = uuid.replace(/-/g, "").padEnd(32, "0");
  const ints: number[] = [];
  for (let i = 0; i < 32; i += 8) {
    ints.push(parseInt(hex.slice(i, i + 8), 16) | 0);
  }
  return ints;
}

export function readVec3(tag: NbtTag | undefined, fallback: Vec3): Vec3 {
  if (!tag || !isList(tag) || tag.value.length < 3) return { ...fallback };
  const nums = tag.value.slice(0, 3).map((t) => (t.type === "string" ? 0 : t.value));
  return {
    x: Number(nums[0]) || 0,
    y: Number(nums[1]) || 0,
    z: Number(nums[2]) || 0,
  };
}

export function entityPos(entity: EntityData): Vec3 {
  return readVec3(entity.nbt.value.Pos, { x: 0, y: 64, z: 0 });
}

export function setEntityPos(entity: EntityData, pos: Vec3): void {
  entity.nbt.value.Pos = Tag.list(
    [Tag.double(pos.x), Tag.double(pos.y), Tag.double(pos.z)],
    "double",
  );
}

export function entityRotation(entity: EntityData): Vec2 {
  const rotation = entity.nbt.value.Rotation;
  if (rotation && isList(rotation) && rotation.value.length >= 2) {
    const yaw = rotation.value[0].value as number;
    const pitch = rotation.value[1].value as number;
    return { yaw: Number(yaw) || 0, pitch: Number(pitch) || 0 };
  }
  return { yaw: 0, pitch: 0 };
}

export function setEntityRotation(entity: EntityData, rotation: Vec2): void {
  entity.nbt.value.Rotation = Tag.list(
    [Tag.float(rotation.yaw), Tag.float(rotation.pitch)],
    "float",
  );
}

export function entityTags(entity: EntityData): string[] {
  const tags = entity.nbt.value.Tags;
  if (!isList(tags)) return [];
  return tags.value
    .filter((t) => t.type === "string")
    .map((t) => (t as { value: string }).value);
}

export function entityCustomName(entity: EntityData): string {
  if (entity.isPlayer && entity.username) return entity.username;
  const custom = entity.nbt.value.CustomName;
  if (!custom) return entityInfo(entity.type).name;
  try {
    if (custom.type === "string") {
      const raw = custom.value.trim();
      if (raw.length === 0) return entityInfo(entity.type).name;
      if (raw.startsWith("{") || raw.startsWith('"') || raw.startsWith("[")) {
        const rendered = plainText(parseTextComponent(raw));
        return rendered.length > 0 ? rendered : entityInfo(entity.type).name;
      }
      return raw;
    }
    const rendered = plainText(custom);
    return rendered.length > 0 ? rendered : entityInfo(entity.type).name;
  } catch {
    return entityInfo(entity.type).name;
  }
}

export function entityHealth(entity: EntityData): number {
  const health = entity.nbt.value.Health;
  return health && typeof health.value === "number" ? health.value : 20;
}

export function describeEntity(entity: EntityData): string {
  const pos = entityPos(entity);
  return `${entityCustomName(entity)} (${entity.type.replace("minecraft:", "")}) @ ${pos.x.toFixed(1)}, ${pos.y.toFixed(1)}, ${pos.z.toFixed(1)}`;
}

export function entityNbtPreview(entity: EntityData, maxLength = 160): string {
  const text = stringifySnbt(entity.nbt, { typed: false, quoteKeys: true });
  return text.length > maxLength ? text.slice(0, maxLength - 1) + "…" : text;
}

export function isMarker(entity: EntityData): boolean {
  return entity.type === "minecraft:marker";
}

export function findNbtKey(compound: NbtTag, key: string): NbtTag | undefined {
  return isCompound(compound) ? compound.value[key] : undefined;
}
