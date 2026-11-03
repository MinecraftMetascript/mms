import type { World } from "./world";
import type { EntityData, Vec3 } from "./entities";
import {
  canonicalUuid,
  entityCustomName,
  entityPos,
  entityTags,
  normalizeType,
  uuidFromInts,
} from "./entities";
import {
  type CompoundTag,
  isCompound,
  isIntArray,
  isList,
  isString,
  type NbtTag,
  numericValue,
  Tag,
  compoundMatches,
} from "./nbt";
import { parseSnbt, SnbtError } from "./snbt";

export type SelectorVariable = "s" | "p" | "a" | "e" | "r" | "n";

export interface ParsedSelector {
  variable: SelectorVariable;
  args: Record<string, string>;
  raw: string;
}

export type TargetToken =
  | { kind: "selector"; selector: ParsedSelector; raw: string }
  | { kind: "name"; name: string; raw: string };

export class SelectorError extends Error {
  index: number;
  constructor(message: string, index = 0) {
    super(message);
    this.name = "SelectorError";
    this.index = index;
  }
}

const VARIABLES: SelectorVariable[] = ["s", "p", "a", "e", "r", "n"];

export function isSelector(text: string): boolean {
  return /^@[spaern]/i.test(text.trim());
}

export function parseSelector(raw: string): ParsedSelector {
  const text = raw.trim();
  if (text.length < 2 || text[0] !== "@") {
    throw new SelectorError("Expected an entity selector starting with @", 0);
  }
  const variable = text[1] as SelectorVariable;
  if (!VARIABLES.includes(variable)) {
    throw new SelectorError(`Unknown selector variable @${text[1]}`, 1);
  }
  const args: Record<string, string> = {};
  if (text.length > 2) {
    if (text[2] !== "[") {
      throw new SelectorError(`Unexpected '${text[2]}' after @${variable}`, 2);
    }
    if (!text.endsWith("]")) {
      throw new SelectorError("Selector is missing its closing ']'", text.length - 1);
    }
    const body = text.slice(3, -1);
    for (const [key, value] of splitSelectorArgs(body)) {
      if (!key) throw new SelectorError("Empty selector argument", 2);
      args[key] = value;
    }
  }
  return { variable, args, raw: text };
}

function splitSelectorArgs(body: string): Array<[string, string]> {
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
      parts.push(body.slice(start, i));
      start = i + 1;
    }
  }
  if (start <= body.length) parts.push(body.slice(start));
  return parts
    .map((part) => part.trim())
    .filter(Boolean)
    .map((part) => {
      const eq = part.indexOf("=");
      if (eq === -1) throw new SelectorError(`Selector argument '${part}' is missing '='`);
      return [part.slice(0, eq).trim(), part.slice(eq + 1).trim()] as [string, string];
    });
}

export function splitTargets(raw: string): string[] {
  const parts: string[] = [];
  let depth = 0;
  let quote: string | null = null;
  let start = 0;
  for (let i = 0; i < raw.length; i++) {
    const ch = raw[i];
    if (quote) {
      if (ch === "\\") i++;
      else if (ch === quote) quote = null;
      continue;
    }
    if (ch === '"' || ch === "'") quote = ch;
    else if (ch === "[" || ch === "{") depth++;
    else if (ch === "]" || ch === "}") depth--;
    else if (ch === "," && depth === 0) {
      parts.push(raw.slice(start, i).trim());
      start = i + 1;
    }
  }
  parts.push(raw.slice(start).trim());
  return parts.filter((p) => p.length > 0);
}

export function parseTargets(raw: string): TargetToken[] {
  return splitTargets(raw).map((token) => {
    if (token.startsWith("@")) {
      return { kind: "selector", selector: parseSelector(token), raw: token };
    }
    return { kind: "name", name: token, raw: token };
  });
}

interface Range {
  min: number;
  max: number;
  negate: boolean;
  exact: boolean;
}

function parseRange(raw: string): Range {
  let text = raw.trim();
  let negate = false;
  if (text.startsWith("!")) {
    negate = true;
    text = text.slice(1);
  }
  if (text.includes("..")) {
    const [a, b] = text.split("..");
    const min = a === "" ? -Infinity : Number(a);
    const max = b === "" ? Infinity : Number(b);
    return { min, max, negate, exact: false };
  }
  const value = Number(text);
  return { min: value, max: value, negate, exact: true };
}

function inRange(value: number, range: Range): boolean {
  const inside = value >= range.min && value <= range.max;
  return range.negate ? !inside : inside;
}

function parseScoreFilter(raw: string): Record<string, Range> {
  const body = raw.trim();
  if (!body.startsWith("{") || !body.endsWith("}")) {
    throw new SelectorError("scores must be a compound like {points=1..5}", 0);
  }
  const inner = body.slice(1, -1);
  const out: Record<string, Range> = {};
  for (const part of splitSelectorArgs(inner.replace(/:/g, "="))) {
    out[part[0]] = parseRange(part[1]);
  }
  return out;
}

export interface SelectOptions {
  source: EntityData | null;
  origin: Vec3;
  /** Do not throw when nothing matches. */
  lenient?: boolean;
  random?: () => number;
}

export function selectEntities(
  world: World,
  selector: ParsedSelector,
  options: SelectOptions,
): EntityData[] {
  const { source, origin, lenient = false } = options;
  const random = options.random ?? Math.random;
  const all = world.entities;
  let candidates: EntityData[];
  switch (selector.variable) {
    case "s":
      if (!source) {
        if (lenient) return [];
        throw new SelectorError("No entity was found (there is no @s in this context)");
      }
      candidates = [source];
      break;
    case "a":
      candidates = all.filter((e) => e.isPlayer);
      break;
    case "p":
      candidates = all.filter((e) => e.isPlayer);
      break;
    case "r":
      candidates = all.filter((e) => e.isPlayer);
      break;
    case "n":
      candidates = all.slice();
      break;
    case "e":
      candidates = all.slice();
      break;
  }

  candidates = candidates.filter((entity) => matchesArgs(world, entity, selector.args, origin));

  const sort = selector.args.sort ?? (selector.variable === "p" || selector.variable === "n" ? "nearest" : "arbitrary");
  const distanceOf = (entity: EntityData): number => {
    const p = entityPos(entity);
    return Math.hypot(p.x - origin.x, p.y - origin.y, p.z - origin.z);
  };

  switch (sort) {
    case "nearest":
      candidates.sort((a, b) => distanceOf(a) - distanceOf(b));
      break;
    case "furthest":
      candidates.sort((a, b) => distanceOf(b) - distanceOf(a));
      break;
    case "random":
      for (let i = candidates.length - 1; i > 0; i--) {
        const j = Math.floor(random() * (i + 1));
        [candidates[i], candidates[j]] = [candidates[j], candidates[i]];
      }
      break;
    case "arbitrary":
      break;
    default:
      throw new SelectorError(`Unknown sort mode '${sort}'`);
  }

  const defaultLimit = selector.variable === "p" || selector.variable === "n" || selector.variable === "r" ? 1 : Infinity;
  let limit = defaultLimit;
  if (selector.args.limit !== undefined) {
    limit = Number(selector.args.limit);
    if (!Number.isInteger(limit) || limit < 1) {
      throw new SelectorError(`limit must be a positive integer, got '${selector.args.limit}'`);
    }
  }
  candidates = candidates.slice(0, limit);

  if (candidates.length === 0 && !lenient && selector.variable !== "s") {
    throw new SelectorError(`No entity was found matching ${selector.raw}`);
  }
  return candidates;
}

function matchesArgs(
  world: World,
  entity: EntityData,
  args: Record<string, string>,
  origin: Vec3,
): boolean {
  for (const [key, raw] of Object.entries(args)) {
    switch (key) {
      case "type": {
        let expected = raw;
        let negate = false;
        if (expected.startsWith("!")) {
          negate = true;
          expected = expected.slice(1);
        }
        const match = entity.type === normalizeType(expected);
        if (negate ? match : !match) return false;
        break;
      }
      case "name": {
        const value = unquote(raw);
        if (entityCustomName(entity) !== value) return false;
        break;
      }
      case "tag": {
        let tag = unquote(raw);
        let negate = false;
        if (tag.startsWith("!")) {
          negate = true;
          tag = tag.slice(1);
        }
        const has = entityTags(entity).includes(tag);
        if (negate ? has : !has) return false;
        break;
      }
      case "nbt": {
        let parsed: NbtTag;
        try {
          parsed = parseSnbt(raw);
        } catch (e) {
          throw new SelectorError((e as SnbtError).message);
        }
        if (!isCompound(parsed) || !compoundMatches(entity.nbt, parsed)) return false;
        break;
      }
      case "scores": {
        const filters = parseScoreFilter(raw);
        for (const [objective, range] of Object.entries(filters)) {
          const score = world.getScore(holderName(entity), objective) ?? 0;
          if (!inRange(score, range)) return false;
        }
        break;
      }
      case "distance": {
        const p = entityPos(entity);
        const distance = Math.hypot(p.x - origin.x, p.y - origin.y, p.z - origin.z);
        if (!inRange(distance, parseRange(raw))) return false;
        break;
      }
      case "level": {
        if (!entity.isPlayer) return false;
        const level = numericValue(entity.nbt.value.XpLevel) ?? 0;
        if (!inRange(level, parseRange(raw))) return false;
        break;
      }
      case "x":
      case "y":
      case "z":
      case "dx":
      case "dy":
      case "dz": {
        // Box filters are evaluated together after this loop.
        break;
      }
      case "sort":
      case "limit":
        break;
      case "gamemode":
      case "advancements":
      case "predicate":
      case "team":
      case "scores.":
        // Accepted but not simulated in this build. Silently ignoring is better
        // than a selector that matches nothing.
        break;
      default:
        break;
    }
  }

  const hasBox = ["x", "y", "z"].some((k) => k in args);
  if (hasBox) {
    const p = entityPos(entity);
    const x = args.x !== undefined ? Number(args.x) : origin.x;
    const y = args.y !== undefined ? Number(args.y) : origin.y;
    const z = args.z !== undefined ? Number(args.z) : origin.z;
    const dx = args.dx !== undefined ? Number(args.dx) : 0;
    const dy = args.dy !== undefined ? Number(args.dy) : 0;
    const dz = args.dz !== undefined ? Number(args.dz) : 0;
    const minX = Math.min(x, x + dx);
    const maxX = Math.max(x, x + dx);
    const minY = Math.min(y, y + dy);
    const maxY = Math.max(y, y + dy);
    const minZ = Math.min(z, z + dz);
    const maxZ = Math.max(z, z + dz);
    if (p.x < minX || p.x > maxX || p.y < minY || p.y > maxY || p.z < minZ || p.z > maxZ) {
      return false;
    }
  }
  return true;
}

function unquote(value: string): string {
  const v = value.trim();
  if ((v.startsWith('"') && v.endsWith('"')) || (v.startsWith("'") && v.endsWith("'"))) {
    return v.slice(1, -1).replace(/\\(["'\\])/g, "$1");
  }
  return v;
}

export function holderName(entity: EntityData): string {
  return entity.isPlayer && entity.username ? entity.username : entity.uuid;
}

function entityUuid(entity: EntityData): string {
  const tag = entity.nbt.value.UUID;
  if (isIntArray(tag) && tag.value.length >= 4) return uuidFromInts(tag.value);
  if (isString(tag)) return canonicalUuid(tag.value) ?? entity.uuid;
  return entity.uuid;
}

function entityByUuid(world: World, raw: string): EntityData | undefined {
  const wanted = canonicalUuid(raw);
  if (!wanted) return undefined;
  return world.entities.find((entity) => entityUuid(entity) === wanted);
}

export function resolveEntityTargets(
  world: World,
  raw: string,
  options: SelectOptions,
): EntityData[] {
  const out: EntityData[] = [];
  for (const token of parseTargets(raw)) {
    if (token.kind === "selector") {
      out.push(...selectEntities(world, token.selector, options));
    } else {
      const player = world.entities.find(
        (e) => e.isPlayer && e.username === token.name,
      );
      if (player) {
        out.push(player);
        continue;
      }
      const byUuid = entityByUuid(world, token.name);
      if (byUuid) {
        out.push(byUuid);
        continue;
      }
      if (!options.lenient) {
        if (canonicalUuid(token.name)) {
          throw new SelectorError(`No entity was found with the UUID '${token.name}'`);
        }
        throw new SelectorError(`No player was found with the name '${token.name}'`);
      }
    }
  }
  return out;
}

export interface HolderTarget {
  holder: string;
  entity?: EntityData;
}

export function resolveHolderTargets(
  world: World,
  raw: string,
  options: SelectOptions,
): HolderTarget[] {
  const out: HolderTarget[] = [];
  for (const token of parseTargets(raw)) {
    if (token.kind === "selector") {
      for (const entity of selectEntities(world, token.selector, options)) {
        out.push({ holder: holderName(entity), entity });
      }
    } else {
      const entity = entityByUuid(world, token.name);
      out.push({
        holder: entity ? holderName(entity) : token.name,
        entity,
      });
    }
  }
  return out;
}

/** Used by text components and `/execute if score` for a single holder. */
export function resolveSingleHolder(
  world: World,
  raw: string,
  source: EntityData | null,
  origin: Vec3,
): string {
  const tokens = parseTargets(raw);
  if (tokens.length === 1 && tokens[0].kind === "name") {
    const entity = entityByUuid(world, tokens[0].name);
    return entity ? holderName(entity) : tokens[0].name;
  }
  const holders = resolveHolderTargets(world, raw, { source, origin, lenient: true });
  if (holders.length === 0) return raw;
  return holders[0].holder;
}

export function summaryOfTargets(entities: EntityData[]): string {
  if (entities.length === 0) return "no entities";
  if (entities.length === 1) return entityCustomName(entities[0]);
  return `${entities.length} entities`;
}

export function scoreOf(world: World, entity: EntityData, objective: string): number | undefined {
  return world.getScore(holderName(entity), objective);
}

export function listScoresFor(world: World, entity: EntityData): string[] {
  return world.scoreObjectives(holderName(entity));
}

export function selectorHasListArg(parsed: ParsedSelector, key: string): boolean {
  return key in parsed.args;
}

export function isTargetSelector(raw: string): boolean {
  return splitTargets(raw).some((t) => t.startsWith("@"));
}

export function entitiesInBox(
  world: World,
  min: Vec3,
  max: Vec3,
): EntityData[] {
  return world.entities.filter((e) => {
    const p = entityPos(e);
    return (
      p.x >= Math.min(min.x, max.x) &&
      p.x <= Math.max(min.x, max.x) &&
      p.y >= Math.min(min.y, max.y) &&
      p.y <= Math.max(min.y, max.y) &&
      p.z >= Math.min(min.z, max.z) &&
      p.z <= Math.max(min.z, max.z)
    );
  });
}

export function scoreboardEntries(
  world: World,
  objective: string,
): Array<{ holder: string; value: number; display?: string }> {
  const scores = world.scores.get(objective);
  if (!scores) return [];
  return [...scores.entries()]
    .map(([holder, value]) => ({
      holder,
      value,
      display: world.holderDisplayName(holder),
    }))
    .sort((a, b) => b.value - a.value || a.holder.localeCompare(b.holder));
}

export function tagListOf(entity: EntityData): string[] {
  const tags = entity.nbt.value.Tags;
  return isList(tags)
    ? tags.value.filter((t) => t.type === "string").map((t) => t.value as string)
    : [];
}

export function isCompoundTag(value: NbtTag | undefined): value is CompoundTag {
  return !!value && isCompound(value);
}

export function dummyTag(): NbtTag {
  return Tag.compound();
}
