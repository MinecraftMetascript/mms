import { Cursor } from "../cursor";
import { isCompound, isList, isNumeric, isString, type NbtTag, Tag } from "../nbt";
import { applyPathModify, getPathValues, parseNbtPath } from "../nbtPath";
import {
  resolveEntityTargets,
  resolveHolderTargets,
  resolveSingleHolder,
  SelectorError,
} from "../selectors";
import { entityPos, entityRotation, type EntityData } from "../entities";
import { parseItemComponents } from "../items";
import type { InventoryItem } from "../player";
import { evaluatePredicate } from "../predicates";
import { executeCommandLine } from "./context";
import {
  type CommandContext,
  type CommandResult,
  feedback,
} from "./context";
import { CommandError } from "../cursor";

type Handler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

interface StoreSpec {
  mode: "result" | "success";
  kind: "score" | "storage" | "entity" | "block";
  targetsRaw?: string;
  objective?: string;
  storageId?: string;
  pathText?: string;
  valueType?: string;
  scale: number;
  pos?: { x: number; y: number; z: number };
}

/** Player eye height above the feet, matching Minecraft's standing eye level. */
export const EYE_HEIGHT = 1.62;

export const executeHandler: Handler = (cursor, ctx) => runChain(cursor, ctx);

function runChain(cursor: Cursor, ctx: CommandContext): CommandResult {
  cursor.skip();
  const subStart = cursor.mark();
  const sub = cursor.readWord();
  switch (sub) {
    case "as": {
      const raw = cursor.readTargetsRaw().raw;
      const restStart = cursor.mark();
      let entities: EntityData[];
      try {
        entities = resolveEntityTargets(ctx.world, raw, {
          source: ctx.source,
          origin: ctx.pos,
          lenient: true,
        });
      } catch (error) {
        if (error instanceof SelectorError) {
          cursor.error(error.message, subStart, "selector");
        }
        throw error;
      }
      if (entities.length === 0) {
        if (ctx.lint) {
          const probe = new Cursor(cursor.text, restStart);
          runChain(probe, { ...ctx, lint: true, quiet: true });
        }
        cursor.pos = cursor.end;
        return { result: 0, success: false };
      }
      let result = 0;
      let success = true;
      for (const entity of entities) {
        const branch = new Cursor(cursor.text, restStart);
        const branchResult = runChain(branch, { ...ctx, source: entity });
        result += branchResult.result;
        success = success && branchResult.success;
      }
      cursor.pos = cursor.end;
      return { result, success };
    }
    case "anchored": {
      const anchorStart = cursor.mark();
      const anchor = cursor.readWord();
      if (anchor !== "eyes" && anchor !== "feet") {
        cursor.error(
          `Expected 'eyes' or 'feet', got '${anchor}'`,
          anchorStart,
          "execute-anchor",
        );
      }
      ctx.anchor = anchor;
      if (ctx.source) {
        const feet = entityPos(ctx.source);
        ctx.pos =
          anchor === "eyes" ? { ...feet, y: feet.y + EYE_HEIGHT } : feet;
      }
      return runChain(cursor, ctx);
    }
    case "at": {
      const raw = cursor.readTargetsRaw().raw;
      const entities = resolveEntityTargets(ctx.world, raw, {
        source: ctx.source,
        origin: ctx.pos,
        lenient: true,
      });
      for (const entity of entities.slice(0, 1)) {
        const feet = entityPos(entity);
        ctx.pos =
          ctx.anchor === "eyes"
            ? { ...feet, y: feet.y + EYE_HEIGHT }
            : feet;
        ctx.rotation = entityRotation(entity);
      }
      return runChain(cursor, ctx);
    }
    case "positioned": {
      const next = cursor.peekWord();
      if (next === "as" || next === "over") {
        cursor.readWord();
        const raw = cursor.readTargetsRaw().raw;
        const entities = resolveEntityTargets(ctx.world, raw, {
          source: ctx.source,
          origin: ctx.pos,
          lenient: true,
        });
        const target = entities[0];
        if (target) {
          const pos = entityPos(target);
          ctx.pos = next === "over" ? { ...pos, y: pos.y + 1 } : pos;
        }
      } else {
        ctx.pos = cursor.readVec3(ctx.pos, ctx.rotation);
      }
      return runChain(cursor, ctx);
    }
    case "align": {
      const axes = cursor.readWord();
      if (!/^[xyz]+$/.test(axes)) {
        cursor.error(`Expected x, y, z or a combination, got '${axes}'`, subStart, "align");
      }
      if (axes.includes("x")) ctx.pos.x = Math.floor(ctx.pos.x);
      if (axes.includes("y")) ctx.pos.y = Math.floor(ctx.pos.y);
      if (axes.includes("z")) ctx.pos.z = Math.floor(ctx.pos.z);
      return runChain(cursor, ctx);
    }
    case "rotated": {
      const next = cursor.peekWord();
      if (next === "as") {
        cursor.readWord();
        const raw = cursor.readTargetsRaw().raw;
        const entities = resolveEntityTargets(ctx.world, raw, {
          source: ctx.source,
          origin: ctx.pos,
          lenient: true,
        });
        if (entities[0]) ctx.rotation = entityRotation(entities[0]);
      } else {
        ctx.rotation = cursor.readVec2();
      }
      return runChain(cursor, ctx);
    }
    case "facing": {
      const next = cursor.peekWord();
      if (next === "entity") {
        cursor.readWord();
        const raw = cursor.readTargetsRaw().raw;
        const anchor = cursor.done() ? "feet" : cursor.peekWord();
        if (anchor === "feet" || anchor === "eyes") cursor.readWord();
        const entities = resolveEntityTargets(ctx.world, raw, {
          source: ctx.source,
          origin: ctx.pos,
          lenient: true,
        });
        const target = entities[0];
        if (target) {
          const pos = entityPos(target);
          ctx.rotation = lookAt(ctx.pos, anchor === "eyes" ? { ...pos, y: pos.y + 1.62 } : pos);
        }
      } else {
        const target = cursor.readVec3(ctx.pos, ctx.rotation);
        ctx.rotation = lookAt(ctx.pos, target);
      }
      return runChain(cursor, ctx);
    }
    case "in": {
      const dimension = cursor.readResourceLocation();
      ctx.dimension = dimension;
      return runChain(cursor, ctx);
    }
    case "if":
    case "unless": {
      const condition = parseCondition(cursor, ctx);
      const value = condition();
      const negate = sub === "unless";
      const passed = negate ? !value : value;
      // `/execute if …` and `/execute unless …` are valid terminal
      // subcommands: they return success/failure without requiring `run`.
      if (cursor.done()) {
        return { result: passed ? 1 : 0, success: passed };
      }
      if (!passed) {
        if (ctx.lint) {
          // Still parse the tail so the linter sees its syntax.
          const probe = cursor.mark();
          try {
            runChain(cursor, { ...ctx, lint: true, quiet: true });
          } catch {
            cursor.pos = probe;
            cursor.readRest();
          }
          return { result: 0, success: false };
        }
        cursor.readRest();
        return { result: 0, success: false };
      }
      return runChain(cursor, ctx);
    }
    case "store": {
      const spec = parseStore(cursor, ctx);
      const result = runChain(cursor, ctx);
      if (!ctx.lint) applyStore(ctx, spec, result);
      return result;
    }
    case "run": {
      const rest = cursor.readRest();
      if (rest.length === 0) {
        cursor.error("Expected a command after 'run'", cursor.mark(), "execute-run");
      }
      return executeCommandLine(rest, ctx);
    }
    default:
      cursor.error(
        `Unknown execute subcommand '${sub}'. Expected as, at, anchored, positioned, align, rotated, facing, in, if, unless, store or run.`,
        subStart,
        "execute-subcommand",
      );
  }
}

type Condition = () => boolean;

function parseCondition(cursor: Cursor, ctx: CommandContext): Condition {
  const start = cursor.mark();
  const kind = cursor.readWord();
  switch (kind) {
    case "entity": {
      const raw = cursor.readTargetsRaw().raw;
      return () => {
        try {
          return (
            resolveEntityTargets(ctx.world, raw, {
              source: ctx.source,
              origin: ctx.pos,
              lenient: true,
            }).length > 0
          );
        } catch {
          return false;
        }
      };
    }
    case "score": {
      const targetsStart = cursor.mark();
      const targetsRaw = cursor.readTargetsRaw().raw;
      const objectiveStart = cursor.mark();
      const objective = cursor.readWord();
      if (!ctx.world.objectives.has(objective)) {
        cursor.error(`Unknown objective '${objective}'`, objectiveStart, "unknown-objective");
      }
      const operatorStart = cursor.mark();
      const op = cursor.readWord();
      if (op === "matches") {
        const range = parseRange(cursor.readWord(), cursor, operatorStart);
        return () => {
          const holders = resolveHolderTargets(ctx.world, targetsRaw, {
            source: ctx.source,
            origin: ctx.pos,
            lenient: true,
          });
          return holders.some((holder) => {
            const value = ctx.world.getScore(holder.holder, objective);
            return (
              value !== undefined &&
              value >= range.min &&
              value <= range.max
            );
          });
        };
      }
      if (!["=", "<", "<=", ">", ">=", "=="].includes(op)) {
        cursor.error(
          `Expected a comparison (=, <, <=, >, >=) or 'matches', got '${op}'`,
          operatorStart,
          "comparison",
        );
      }
      const sourceStart = cursor.mark();
      const sourceRaw = cursor.readWord();
      const sourceObjectiveStart = cursor.mark();
      const sourceObjective = cursor.readWord();
      if (!ctx.world.objectives.has(sourceObjective)) {
        cursor.error(
          `Unknown objective '${sourceObjective}'`,
          sourceObjectiveStart,
          "unknown-objective",
        );
      }
      return () => {
        const targetHolders = resolveHolderTargets(ctx.world, targetsRaw, {
          source: ctx.source,
          origin: ctx.pos,
          lenient: true,
        });
        if (targetHolders.length === 0) return false;
        let sourceHolder = sourceRaw;
        try {
          sourceHolder = resolveSingleHolder(ctx.world, sourceRaw, ctx.source, ctx.pos);
        } catch {
          sourceHolder = sourceRaw;
        }
        const sourceValue = ctx.world.getScore(sourceHolder, sourceObjective);
        if (sourceValue === undefined) return false;
        return targetHolders.some((holder) => {
          const value = ctx.world.getScore(holder.holder, objective);
          if (value === undefined) return false;
          switch (op) {
            case "=":
            case "==":
              return value === sourceValue;
            case "<":
              return value < sourceValue;
            case "<=":
              return value <= sourceValue;
            case ">":
              return value > sourceValue;
            case ">=":
              return value >= sourceValue;
            default:
              return false;
          }
        });
      };
      void sourceStart;
    }
    case "block": {
      const pos = cursor.readVec3(ctx.pos, ctx.rotation);
      const blockStart = cursor.mark();
      const block = cursor.readWord();
      return () => {
        const actual = ctx.world.getBlock(pos);
        if (!actual) return false;
        return normalizeBlock(actual) === normalizeBlock(block);
      };
      void blockStart;
    }
    case "data": {
      const sourceKindStart = cursor.mark();
      const sourceKind = cursor.readWord();
      if (!["block", "entity", "storage"].includes(sourceKind)) {
        cursor.error(`Expected block, entity or storage, got '${sourceKind}'`, sourceKindStart);
      }
      let exists: () => boolean;
      if (sourceKind === "entity") {
        const targetsStart = cursor.mark();
        const raw = cursor.readTargetsRaw().raw;
        const pathStart = cursor.mark();
        const pathText = cursor.readNbtPathText();
        const path = parsePath(cursor, pathText, pathStart);
        exists = () => {
          const entities = resolveEntityTargets(ctx.world, raw, {
            source: ctx.source,
            origin: ctx.pos,
            lenient: true,
          });
          return entities.some((entity) => getPathValues(entity.nbt, path).length > 0);
        };
        void targetsStart;
      } else if (sourceKind === "storage") {
        const id = cursor.readResourceLocation();
        const pathStart = cursor.mark();
        const pathText = cursor.readNbtPathText();
        const path = parsePath(cursor, pathText, pathStart);
        exists = () => {
          const storage = ctx.world.storage(id);
          return !!storage && getPathValues(storage, path).length > 0;
        };
      } else {
        const pos = cursor.readVec3(ctx.pos, ctx.rotation);
        const pathStart = cursor.mark();
        const pathText = cursor.readNbtPathText();
        const path = parsePath(cursor, pathText, pathStart);
        exists = () => getPathValues(ctx.world.getBlockNbt(pos), path).length > 0;
      }
      return exists;
    }
    case "dimension": {
      const id = cursor.readResourceLocation();
      return () => ctx.dimension === id;
    }
    case "predicate": {
      const idStart = cursor.mark();
      const id = cursor.readResourceLocation();
      if (ctx.lint && !ctx.world.getPredicate(id)) {
        cursor.error(`Unknown predicate '${id}'`, idStart, "unknown-predicate");
      }
      return () => {
        if (!ctx.world.getPredicate(id)) return false;
        return evaluatePredicate(ctx.world, id, {
          source: ctx.source,
          pos: ctx.pos,
          dimension: ctx.dimension,
        });
      };
    }
    case "function": {
      const idStart = cursor.mark();
      const id = cursor.readResourceLocation();
      if (ctx.lint && !ctx.world.getFunction(id)) {
        cursor.error(`Unknown function '${id}'`, idStart, "unknown-function");
      }
      return () => {
        if (ctx.lint) return true;
        return ctx.runtime.runFunction(id, null, ctx).success;
      };
    }
    case "loaded": {
      cursor.readVec3(ctx.pos, ctx.rotation);
      return () => true;
    }
    case "items": {
      const sourceKind = cursor.readWord();
      if (sourceKind === "entity") {
        const raw = cursor.readTargetsRaw().raw;
        const slotsText = cursor.readWord();
        const slots = slotNames(slotsText);
        const predicate = readItemPredicate(cursor);
        return () => {
          const entities = resolveEntityTargets(ctx.world, raw, {
            source: ctx.source,
            origin: ctx.pos,
            lenient: true,
          });
          return entities.some((entity) =>
            slots.some((slot) => {
              const item = ctx.world.getItemAt(entity, slot);
              return !!item && itemMatchesPredicate(item, predicate);
            }),
          );
        };
      }
      cursor.error(
        "Only 'execute if items entity <targets> <slots> <item>' is simulated",
        start,
        "items-source",
      );
      break;
    }
    default:
      cursor.error(
        `Unknown condition '${kind}'. Expected entity, score, block, data, dimension, function, items, loaded or predicate.`,
        start,
        "condition",
      );
  }
  return () => false;
}

function parseStore(cursor: Cursor, ctx: CommandContext): StoreSpec {
  const modeStart = cursor.mark();
  const mode = cursor.readWord();
  if (mode !== "result" && mode !== "success") {
    cursor.error(`Expected 'result' or 'success', got '${mode}'`, modeStart, "store-mode");
  }
  const kindStart = cursor.mark();
  const kind = cursor.readWord();
  const spec: StoreSpec = { mode, kind: kind as StoreSpec["kind"], scale: 1 };
  if (kind === "score") {
    spec.targetsRaw = cursor.readTargetsRaw().raw;
    spec.objective = cursor.readWord();
    if (!ctx.world.objectives.has(spec.objective)) {
      cursor.error(`Unknown objective '${spec.objective}'`, cursor.mark(), "unknown-objective");
    }
    return spec;
  }
  if (kind === "storage") {
    spec.storageId = cursor.readResourceLocation();
    spec.pathText = cursor.readNbtPathText();
    spec.valueType = cursor.readWord();
    assertStoreType(cursor, spec.valueType);
    if (!cursor.done()) spec.scale = cursor.readNumber().value;
    return spec;
  }
  if (kind === "entity") {
    spec.targetsRaw = cursor.readTargetsRaw().raw;
    spec.pathText = cursor.readNbtPathText();
    spec.valueType = cursor.readWord();
    assertStoreType(cursor, spec.valueType);
    if (!cursor.done()) spec.scale = cursor.readNumber().value;
    return spec;
  }
  if (kind === "block") {
    spec.pos = cursor.readVec3(ctx.pos, ctx.rotation);
    spec.pathText = cursor.readNbtPathText();
    spec.valueType = cursor.readWord();
    assertStoreType(cursor, spec.valueType);
    if (!cursor.done()) spec.scale = cursor.readNumber().value;
    return spec;
  }
  cursor.error(`Expected score, storage, entity or block, got '${kind}'`, kindStart, "store-kind");
}

function assertStoreType(cursor: Cursor, type: string): void {
  if (!["byte", "short", "int", "long", "float", "double"].includes(type)) {
    cursor.error(
      `Expected byte, short, int, long, float or double, got '${type}'`,
      cursor.mark() - type.length,
      "store-type",
    );
  }
}

function applyStore(ctx: CommandContext, spec: StoreSpec, result: CommandResult): void {
  const raw = spec.mode === "result" ? result.result : result.success ? 1 : 0;
  const value = Math.floor(raw * spec.scale);
  if (spec.kind === "score") {
    for (const target of resolveHolderTargets(ctx.world, spec.targetsRaw!, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: true,
    })) {
      ctx.world.setScore(target.holder, spec.objective!, value);
    }
    return;
  }
  const tag = makeNumberTag(spec.valueType!, value);
  if (spec.kind === "storage") {
    const storage = ctx.world.storage(spec.storageId!, true)!;
    applyPathModify(storage, parseNbtPath(spec.pathText!), "set", tag);
    return;
  }
  if (spec.kind === "entity") {
    const entities = resolveEntityTargets(ctx.world, spec.targetsRaw!, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: true,
    });
    for (const entity of entities) {
      applyPathModify(entity.nbt, parseNbtPath(spec.pathText!), "set", tag);
    }
    return;
  }
  if (spec.kind === "block" && spec.pos) {
    applyPathModify(ctx.world.getBlockNbt(spec.pos), parseNbtPath(spec.pathText!), "set", tag);
  }
}

function makeNumberTag(type: string, value: number): NbtTag {
  switch (type) {
    case "byte":
      return Tag.byte(value);
    case "short":
      return Tag.short(value);
    case "long":
      return Tag.long(value);
    case "float":
      return Tag.float(value);
    case "double":
      return Tag.double(value);
    default:
      return Tag.int(value);
  }
}

function parseRange(raw: string, cursor: Cursor, start: number): { min: number; max: number } {
  if (raw.includes("..")) {
    const [a, b] = raw.split("..");
    const min = a === "" ? -Infinity : Number(a);
    const max = b === "" ? Infinity : Number(b);
    if (Number.isNaN(min) || Number.isNaN(max)) {
      cursor.error(`Invalid range '${raw}'`, start, "range");
    }
    return { min, max };
  }
  const value = Number(raw);
  if (Number.isNaN(value)) cursor.error(`Invalid range '${raw}'`, start, "range");
  return { min: value, max: value };
}

function parsePath(cursor: Cursor, text: string, start: number) {
  try {
    return parseNbtPath(text);
  } catch (error) {
    const pathError = error as { message: string; index: number };
    cursor.error(`Invalid path '${text}': ${pathError.message}`, start + (pathError.index ?? 0), "path");
  }
}

function lookAt(from: { x: number; y: number; z: number }, to: { x: number; y: number; z: number }) {
  const dx = to.x - from.x;
  const dy = to.y - from.y;
  const dz = to.z - from.z;
  const horizontal = Math.hypot(dx, dz);
  const yaw = (Math.atan2(-dx, dz) * 180) / Math.PI;
  const pitch = (-Math.atan2(dy, horizontal) * 180) / Math.PI;
  return { yaw, pitch };
}

function normalizeBlock(id: string): string {
  return id.includes(":") ? id : `minecraft:${id}`;
}

void isList;
void feedback;


interface ItemPredicate {
  id: string;
  components: Record<string, NbtTag>;
  tag?: boolean;
}

function readItemPredicate(cursor: Cursor): ItemPredicate {
  if (cursor.peek() === "#") {
    cursor.readWord();
    return { id: "", components: {}, tag: true };
  }
  const id = cursor.readResourceLocation();
  let components: Record<string, NbtTag> = {};
  cursor.skip();
  if (cursor.peek() === "[") {
    components = parseItemComponents(cursor.readBalanced());
  }
  return { id: id.includes(":") ? id : `minecraft:${id}`, components };
}

function slotNames(source: string): string[] {
  const expand = (prefix: string, count: number) =>
    Array.from({ length: count }, (_, i) => `${prefix}.${i}`);
  if (source === "*") {
    return [
      "weapon.mainhand", "weapon.offhand",
      "armor.head", "armor.chest", "armor.legs", "armor.feet",
      ...expand("hotbar", 9), ...expand("inventory", 27),
      ...expand("enderchest", 27),
    ];
  }
  if (source === "weapon.*") return ["weapon.mainhand", "weapon.offhand"];
  if (source === "armor.*") return ["armor.head", "armor.chest", "armor.legs", "armor.feet"];
  if (source === "hotbar.*") return expand("hotbar", 9);
  if (source === "inventory.*" || source === "container.*") return expand("inventory", 27);
  if (source === "enderchest.*") return expand("enderchest", 27);
  const range = /^(hotbar|inventory|container|enderchest)\.(\d+)-(\d+)$/.exec(source);
  if (range) {
    const prefix = range[1] === "container" ? "inventory" : range[1];
    const start = Number(range[2]);
    const end = Number(range[3]);
    const out: string[] = [];
    for (let i = start; i <= end; i++) out.push(`${prefix}.${i}`);
    return out;
  }
  return [source];
}

function predicateComponentMatches(actual: NbtTag, expected: NbtTag): boolean {
  if (isCompound(actual) && isCompound(expected)) {
    return Object.entries(expected.value).every(
      ([key, value]) =>
        actual.value[key] !== undefined && predicateComponentMatches(actual.value[key], value),
    );
  }
  if (isList(actual) && isList(expected)) {
    return (
      actual.value.length === expected.value.length &&
      actual.value.every((entry, index) => predicateComponentMatches(entry, expected.value[index]))
    );
  }
  if (isString(actual) && isString(expected)) return actual.value === expected.value;
  if (isNumeric(actual) && isNumeric(expected)) return Number(actual.value) === Number(expected.value);
  return false;
}

function itemMatchesPredicate(item: InventoryItem, predicate: ItemPredicate): boolean {
  if (predicate.tag) return false;
  if (item.id !== (predicate.id.includes(":") ? predicate.id : `minecraft:${predicate.id}`)) {
    return false;
  }
  return Object.entries(predicate.components).every(([key, expected]) => {
    const actual = item.components?.[key];
    return actual !== undefined && predicateComponentMatches(actual, expected);
  });
}
