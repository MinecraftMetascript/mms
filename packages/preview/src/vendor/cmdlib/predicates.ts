import type { World } from "./world";
import type { EntityData, Vec3 } from "./entities";
import { entityPos, normalizeType } from "./entities";
import { holderName, resolveEntityTargets, resolveHolderTargets } from "./selectors";
import {
  compoundMatches,
  isCompound,
  numericValue,
  type NbtTag,
} from "./nbt";
import { jsonToTag, parseSnbt } from "./snbt";
import { getPathValues, parseNbtPath } from "./nbtPath";
import { isValidNamespace, isValidResourcePath } from "./tag-kinds";

/**
 * Datapack predicates.
 *
 * A predicate is a data-driven condition tree. The runtime evaluates the common
 * vanilla condition types directly against the simulated `World`; unsupported
 * condition types are false at runtime and are reported by the linter rather
 * than throwing.
 */

export interface PredicateDefinition {
  /** Resource location, e.g. `cmdlib:is_hungry`. */
  id: string;
  /** Source file, e.g. `data/cmdlib/predicate/is_hungry.json`. */
  path: string;
  /** Parsed JSON payload (the object containing `condition`). */
  data: unknown;
}

export interface PredicateContext {
  source: EntityData | null;
  pos: Vec3;
  dimension: string;
}

export type PredicateErrorCode =
  | "invalid-predicate-id"
  | "invalid-predicate-json"
  | "missing-condition";

export interface PredicateParseError {
  path: string;
  message: string;
  code: PredicateErrorCode;
}

export interface ParsedPredicates {
  predicates: Map<string, PredicateDefinition>;
  errors: PredicateParseError[];
}

const PREDICATE_FILE_RE = /^data\/([^/]+)\/predicate\/(.+)\.json$/;

/** Condition types this runtime understands. `minecraft:` is optional. */
export const SUPPORTED_PREDICATE_CONDITIONS = new Set([
  "all_of",
  "any_of",
  "inverted",
  "reference",
  "entity_properties",
  "entity_present",
  "location_check",
  "score",
  "entity_scores",
  "random_chance",
  "random_chance_with_looting",
  "time_check",
  "weather_check",
  "value_check",
]);

export function parsePredicateFiles(files: Record<string, string>): ParsedPredicates {
  const predicates = new Map<string, PredicateDefinition>();
  const errors: PredicateParseError[] = [];

  for (const [path, content] of Object.entries(files)) {
    const match = PREDICATE_FILE_RE.exec(path);
    if (!match) continue;
    const [, namespace, name] = match;
    if (!isValidNamespace(namespace) || !isValidResourcePath(name)) {
      errors.push({
        path,
        code: "invalid-predicate-id",
        message: `Invalid predicate id '${namespace}:${name}'. Use a lower-case namespace and path.`,
      });
      continue;
    }
    let parsed: unknown;
    try {
      parsed = JSON.parse(content);
    } catch (error) {
      errors.push({
        path,
        code: "invalid-predicate-json",
        message: `Invalid predicate JSON: ${(error as Error).message}`,
      });
      continue;
    }
    if (!isRecord(parsed)) {
      errors.push({
        path,
        code: "invalid-predicate-json",
        message: "A predicate file must contain a JSON object with a 'condition' field.",
      });
      continue;
    }
    if (!("condition" in parsed)) {
      errors.push({
        path,
        code: "missing-condition",
        message: "Predicate is missing a 'condition' field.",
      });
      continue;
    }
    if (
      typeof parsed.condition !== "string" &&
      !isRecord(parsed.condition)
    ) {
      errors.push({
        path,
        code: "invalid-predicate-json",
        message: "The predicate 'condition' must be a condition id string.",
      });
      continue;
    }
    const id = `${namespace}:${name}`;
    predicates.set(id, { id, path, data: parsed });
  }

  return { predicates, errors };
}

/* ------------------------------ evaluation ------------------------------ */

/**
 * Evaluate a predicate by id or definition. Unknown ids, malformed conditions
 * and reference cycles all evaluate to `false` instead of throwing.
 */
export function evaluatePredicate(
  world: World,
  definitionOrId: PredicateDefinition | string,
  context: PredicateContext,
): boolean {
  const definition =
    typeof definitionOrId === "string"
      ? world.getPredicate(definitionOrId)
      : definitionOrId;
  if (!definition) return false;
  const seen = new Set<string>([canonicalId(definition.id)]);
  return evaluateCondition(world, predicateRoot(definition.data), context, seen);
}

function evaluateCondition(
  world: World,
  raw: unknown,
  context: PredicateContext,
  seen: Set<string>,
): boolean {
  if (!isRecord(raw)) return false;
  const typeRaw = raw.condition ?? raw.type;
  if (typeof typeRaw !== "string") return false;
  const type = normalizeConditionType(typeRaw);

  switch (type) {
    case "all_of": {
      const terms = Array.isArray(raw.terms) ? raw.terms : [];
      return terms.every((term) => evaluateCondition(world, term, context, seen));
    }
    case "any_of": {
      const terms = Array.isArray(raw.terms) ? raw.terms : [];
      return terms.some((term) => evaluateCondition(world, term, context, seen));
    }
    case "inverted":
      return !evaluateCondition(world, raw.term, context, seen);
    case "reference": {
      if (typeof raw.id !== "string") return false;
      const id = canonicalId(raw.id);
      if (seen.has(id)) return false;
      const definition = world.getPredicate(id);
      if (!definition) return false;
      const next = new Set(seen);
      next.add(id);
      return evaluateCondition(world, predicateRoot(definition.data), context, next);
    }
    case "entity_properties":
      return evalEntityProperties(world, raw, context);
    case "entity_present":
      return selectPredicateEntities(world, raw.entity, context).length > 0;
    case "location_check":
      return evalLocationCheck(world, raw, context);
    case "score":
      return evalScore(world, raw, context);
    case "entity_scores":
      return evalEntityScores(world, raw, context);
    case "random_chance":
    case "random_chance_with_looting":
      return evalRandomChance(raw);
    case "time_check": {
      const period =
        typeof raw.period === "number" && raw.period > 0 ? raw.period : 24000;
      const time = ((world.tick % period) + period) % period;
      return rangeContains(raw.value, time);
    }
    case "weather_check":
      // The simulator has no weather model, so no weather condition can hold.
      return false;
    case "value_check": {
      const value = resolveNumberProvider(world, raw.value, context);
      if (value === null) return false;
      const range = raw.range ?? { min: raw.min, max: raw.max };
      return rangeContains(range, value);
    }
    default:
      return false;
  }
}

function evalEntityProperties(
  world: World,
  raw: Record<string, unknown>,
  context: PredicateContext,
): boolean {
  const entities = selectPredicateEntities(world, raw.entity ?? "this", context);
  if (entities.length === 0) return false;
  const predicate = isRecord(raw.predicate) ? raw.predicate : {};
  return entities.some((entity) => matchesEntityPredicate(entity, predicate));
}

function matchesEntityPredicate(
  entity: EntityData,
  predicate: Record<string, unknown>,
): boolean {
  if (predicate.type !== undefined) {
    const types = Array.isArray(predicate.type) ? predicate.type : [predicate.type];
    const matches = types.some((raw) => {
      if (typeof raw !== "string" || raw.startsWith("#")) return false;
      return entity.type === normalizeType(raw);
    });
    if (!matches) return false;
  }

  if (isRecord(predicate.flags)) {
    for (const [flag, expected] of Object.entries(predicate.flags)) {
      if (typeof expected !== "boolean") return false;
      const actual = entityFlag(entity, flag);
      if (actual === null || actual !== expected) return false;
    }
  }

  if (predicate.nbt !== undefined) {
    const pattern = parsePredicateNbt(predicate.nbt);
    if (!pattern || !isCompound(pattern) || !compoundMatches(entity.nbt, pattern)) {
      return false;
    }
  }

  return true;
}

function entityFlag(entity: EntityData, flag: string): boolean | null {
  const nbt = entity.nbt.value;
  const number = (key: string): number => numericValue(nbt[key]) ?? 0;
  const truthy = (key: string): boolean => number(key) !== 0;
  switch (flag) {
    case "is_baby":
      return truthy("IsBaby") || (nbt.Age !== undefined && number("Age") < 0);
    case "is_on_fire":
      return number("Fire") > 0;
    case "is_sneaking":
      return truthy("Sneaking");
    case "is_sprinting":
      return truthy("Sprinting");
    case "is_swimming":
      return truthy("Swimming");
    case "is_on_ground":
      return truthy("OnGround");
    case "is_flying":
      return truthy("Flying");
    case "is_elytra_flying":
      return truthy("FallFlying");
    case "is_invisible":
      return truthy("Invisible");
    case "is_glowing":
      return truthy("Glowing");
    case "is_sleeping":
      return truthy("Sleeping");
    default:
      return null;
  }
}

function parsePredicateNbt(raw: unknown): NbtTag | null {
  if (typeof raw === "string") {
    try {
      return parseSnbt(raw);
    } catch {
      return null;
    }
  }
  return jsonToTag(raw);
}

function evalLocationCheck(
  world: World,
  raw: Record<string, unknown>,
  context: PredicateContext,
): boolean {
  const offset = (axis: "x" | "y" | "z"): number => {
    const value = raw[`offset${axis.toUpperCase()}`];
    return typeof value === "number" ? value : 0;
  };
  const pos: Vec3 = {
    x: context.pos.x + offset("x"),
    y: context.pos.y + offset("y"),
    z: context.pos.z + offset("z"),
  };

  if (raw.dimension !== undefined) {
    if (typeof raw.dimension !== "string") return false;
    if (canonicalId(raw.dimension) !== canonicalId(context.dimension)) return false;
  }

  if (isRecord(raw.position)) {
    for (const axis of ["x", "y", "z"] as const) {
      const bound = raw.position[axis];
      if (bound === undefined) continue;
      const range = parsePredicatePositionRange(bound, pos[axis]);
      if (!range || pos[axis] < range.min || pos[axis] > range.max) return false;
    }
  } else if (raw.position !== undefined) {
    return false;
  }

  if (raw.block !== undefined && !matchesBlock(world, pos, raw.block)) return false;
  if (raw.fluid !== undefined && !matchesFluid(raw.fluid)) return false;

  return true;
}

function matchesBlock(world: World, pos: Vec3, raw: unknown): boolean {
  const actualRaw = world.getBlock({
    x: Math.floor(pos.x),
    y: Math.floor(pos.y),
    z: Math.floor(pos.z),
  });
  if (!actualRaw) return false;
  const actual = splitBlockState(actualRaw);

  let expectedIds: string[] = [];
  let expectedStates: Record<string, string> | null = null;
  if (typeof raw === "string") {
    expectedIds = [raw];
  } else if (isRecord(raw)) {
    if (typeof raw.id === "string") expectedIds = [raw.id];
    else if (Array.isArray(raw.blocks)) {
      expectedIds = raw.blocks.filter((entry): entry is string => typeof entry === "string");
    }
    const states = raw.state ?? raw.states;
    if (isRecord(states)) {
      expectedStates = {};
      for (const [key, value] of Object.entries(states)) {
        expectedStates[key] = String(value);
      }
    }
  }
  if (expectedIds.length === 0) return false;
  const idMatches = expectedIds.some((id) => {
    if (id.startsWith("#")) return false;
    return splitBlockState(id).id === actual.id;
  });
  if (!idMatches) return false;
  if (!expectedStates) return true;
  return Object.entries(expectedStates).every(
    ([key, value]) => actual.states[key] === value,
  );
}

function splitBlockState(raw: string): { id: string; states: Record<string, string> } {
  const open = raw.indexOf("[");
  if (open === -1 || !raw.endsWith("]")) {
    return { id: normalizeType(raw), states: {} };
  }
  const id = normalizeType(raw.slice(0, open));
  const states: Record<string, string> = {};
  for (const entry of raw.slice(open + 1, -1).split(",")) {
    const [key, value] = entry.split("=");
    if (key && value !== undefined) states[key.trim()] = value.trim();
  }
  return { id, states };
}

function matchesFluid(raw: unknown): boolean {
  let id: string | null = null;
  if (typeof raw === "string") id = raw;
  else if (isRecord(raw)) {
    if (typeof raw.id === "string") id = raw.id;
    else if (typeof raw.fluid === "string") id = raw.fluid;
  }
  // No fluid simulation yet: only the explicit "empty" fluid matches.
  return !!id && !id.startsWith("#") && canonicalId(id) === "minecraft:empty";
}

function evalScore(
  world: World,
  raw: Record<string, unknown>,
  context: PredicateContext,
): boolean {
  const objective =
    typeof raw.objective === "string"
      ? raw.objective
      : typeof raw.score === "string"
        ? raw.score
        : null;
  if (!objective) return false;
  const target = raw.target ?? raw.name ?? "this";
  const holders = resolvePredicateHolders(world, target, context);
  return holders.some((holder) =>
    rangeContains(raw.range, world.getScore(holder, objective) ?? 0),
  );
}

function evalEntityScores(
  world: World,
  raw: Record<string, unknown>,
  context: PredicateContext,
): boolean {
  const entities = selectPredicateEntities(world, raw.entity, context);
  const scores = isRecord(raw.scores) ? raw.scores : null;
  if (!scores) return false;
  return entities.some((entity) => {
    const holder = holderName(entity);
    return Object.entries(scores).every(([objective, range]) =>
      rangeContains(range, world.getScore(holder, objective) ?? 0),
    );
  });
}

function evalRandomChance(raw: Record<string, unknown>): boolean {
  const chance = typeof raw.chance === "number" ? raw.chance : null;
  if (chance === null) return false;
  const looting =
    typeof raw.looting_multiplier === "number" ? raw.looting_multiplier : 0;
  // There is no enchantment model yet, so the looting bonus is only additive
  // with a looting level of zero. This still keeps chance 0/1 deterministic.
  const effective = chance + looting * 0;
  if (effective <= 0) return false;
  if (effective >= 1) return true;
  return Math.random() < effective;
}

function resolvePredicateHolders(
  world: World,
  raw: unknown,
  context: PredicateContext,
): string[] {
  if (typeof raw !== "string" || raw.length === 0 || raw === "this") {
    return context.source ? [holderName(context.source)] : [];
  }
  if (raw.startsWith("@") || raw.includes(",")) {
    try {
      return resolveHolderTargets(world, raw, {
        source: context.source,
        origin: context.pos,
        lenient: true,
      }).map((entry) => entry.holder);
    } catch {
      return [];
    }
  }
  return [raw];
}

function resolveNumberProvider(
  world: World,
  raw: unknown,
  context: PredicateContext,
): number | null {
  if (typeof raw === "number") return raw;
  if (typeof raw === "string") {
    const value = Number(raw);
    return Number.isNaN(value) ? null : value;
  }
  if (!isRecord(raw)) return null;
  const type =
    typeof raw.type === "string" ? normalizeConditionType(raw.type) : null;
  if (type === "score") {
    const objective =
      typeof raw.objective === "string"
        ? raw.objective
        : typeof raw.score === "string"
          ? raw.score
          : null;
    if (!objective) return null;
    const holders = resolvePredicateHolders(world, raw.target ?? "this", context);
    if (holders.length === 0) return null;
    const value = world.getScore(holders[0], objective) ?? 0;
    const scale = typeof raw.scale === "number" ? raw.scale : 1;
    return value * scale;
  }
  if (type === "storage") {
    if (typeof raw.storage !== "string" || typeof raw.path !== "string") return null;
    const storage = world.storage(canonicalId(raw.storage));
    if (!storage) return null;
    try {
      const values = getPathValues(storage, parseNbtPath(raw.path));
      const value = values.length > 0 ? numericValue(values[0]) : null;
      if (value === null) return null;
      const scale = typeof raw.scale === "number" ? raw.scale : 1;
      return value * scale;
    } catch {
      return null;
    }
  }
  return null;
}

/* --------------------------- condition traversal ------------------------- */

/** Every predicate id referenced by `reference` conditions (in tree order). */
export function collectPredicateReferences(condition: unknown): string[] {
  const out: string[] = [];
  walkCondition(predicateRoot(condition), (type, node) => {
    if (type === "reference" && typeof node.id === "string") out.push(node.id);
  });
  return out;
}

/** Every condition type used in the tree, normalised without `minecraft:`. */
export function collectPredicateConditionTypes(condition: unknown): string[] {
  const out: string[] = [];
  walkCondition(predicateRoot(condition), (type) => out.push(type));
  return out;
}

export function isSupportedPredicateCondition(type: string): boolean {
  return SUPPORTED_PREDICATE_CONDITIONS.has(normalizeConditionType(type));
}

function walkCondition(
  raw: unknown,
  visitor: (type: string, node: Record<string, unknown>) => void,
): void {
  if (!isRecord(raw)) return;
  const typeRaw = raw.condition ?? raw.type;
  if (typeof typeRaw !== "string") return;
  const type = normalizeConditionType(typeRaw);
  visitor(type, raw);
  if (type === "all_of" || type === "any_of") {
    if (Array.isArray(raw.terms)) {
      for (const term of raw.terms) walkCondition(term, visitor);
    }
  } else if (type === "inverted") {
    walkCondition(raw.term, visitor);
  }
}

/**
 * Detect reference cycles. Each returned array is a cycle path with the first
 * id repeated at the end, e.g. `["a:p", "a:q", "a:p"]`.
 */
export function findPredicateCycles(
  predicates: Map<string, PredicateDefinition>,
): string[][] {
  const cycles: string[][] = [];
  const state = new Map<string, "visiting" | "done">();
  const stack: string[] = [];

  const visit = (id: string): void => {
    if (state.get(id) === "visiting") {
      const start = stack.indexOf(id);
      if (start >= 0) cycles.push([...stack.slice(start), id]);
      return;
    }
    if (state.get(id) === "done") return;
    state.set(id, "visiting");
    stack.push(id);
    const definition = predicates.get(id);
    if (definition) {
      for (const reference of collectPredicateReferences(definition.data)) {
        const target = canonicalId(reference);
        if (predicates.has(target)) visit(target);
      }
    }
    stack.pop();
    state.set(id, "done");
  };

  for (const id of predicates.keys()) visit(id);
  return cycles;
}

/* -------------------------------- helpers ------------------------------- */

export function canonicalId(raw: string): string {
  const text = raw.trim();
  return text.includes(":") ? text : `minecraft:${text}`;
}

export function normalizeConditionType(raw: string): string {
  return raw.trim().replace(/^minecraft:/, "");
}

interface NumericRange {
  min: number;
  max: number;
}

/** Parse `1..5`, `..5`, `1..`, an exact number, or a `{min,max}` object. */
export function parsePredicateRange(raw: unknown, base?: number): NumericRange | null {
  if (typeof raw === "number") {
    return Number.isNaN(raw) ? null : { min: raw, max: raw };
  }
  if (typeof raw === "string") {
    const text = raw.trim();
    if (text.length === 0) return null;
    if (text.includes("..")) {
      const [a, b] = text.split("..");
      const min = a === "" ? -Infinity : parseRangeEndpoint(a, base);
      const max = b === "" ? Infinity : parseRangeEndpoint(b, base);
      if (Number.isNaN(min) || Number.isNaN(max)) return null;
      return { min, max };
    }
    const value = parseRangeEndpoint(text, base);
    return Number.isNaN(value) ? null : { min: value, max: value };
  }
  if (isRecord(raw)) {
    const min =
      raw.min === undefined
        ? -Infinity
        : typeof raw.min === "number"
          ? raw.min
          : typeof raw.min === "string"
            ? parseRangeEndpoint(raw.min, base)
            : NaN;
    const max =
      raw.max === undefined
        ? Infinity
        : typeof raw.max === "number"
          ? raw.max
          : typeof raw.max === "string"
            ? parseRangeEndpoint(raw.max, base)
            : NaN;
    if (Number.isNaN(min) || Number.isNaN(max)) return null;
    return { min, max };
  }
  return null;
}

function parsePredicatePositionRange(raw: unknown, base: number): NumericRange | null {
  return parsePredicateRange(raw, base);
}

function parseRangeEndpoint(text: string, base?: number): number {
  if (text.startsWith("~")) {
    const rest = text.slice(1);
    const delta = rest === "" ? 0 : Number(rest);
    if (Number.isNaN(delta)) return NaN;
    return (base ?? 0) + delta;
  }
  return Number(text);
}

function rangeContains(raw: unknown, value: number): boolean {
  const range = parsePredicateRange(raw);
  if (!range) return false;
  return value >= range.min && value <= range.max;
}

function selectPredicateEntities(
  world: World,
  raw: unknown,
  context: PredicateContext,
): EntityData[] {
  if (typeof raw !== "string") return [];
  const selector = raw.trim();
  if (selector.length === 0 || selector === "this") {
    return context.source ? [context.source] : [];
  }
  try {
    return resolveEntityTargets(world, selector, {
      source: context.source,
      origin: context.pos,
      lenient: true,
    });
  } catch {
    return [];
  }
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return !!value && typeof value === "object" && !Array.isArray(value);
}

/** Accept both the vanilla wrapper (`condition` is a type string) and a nested condition object. */
function predicateRoot(data: unknown): unknown {
  if (isRecord(data) && isRecord(data.condition)) return data.condition;
  return data;
}

// Keep entityPos in the public surface used by callers constructing a context
// from an entity without importing entities separately.
export function predicateContextFor(
  source: EntityData | null,
  pos?: Vec3,
  dimension = "minecraft:overworld",
): PredicateContext {
  return {
    source,
    pos: pos ?? (source ? entityPos(source) : { x: 0, y: 64, z: 0 }),
    dimension,
  };
}

