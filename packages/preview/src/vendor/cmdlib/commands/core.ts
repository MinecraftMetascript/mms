import { Cursor } from "../cursor";
import {
  type CompoundTag,
  isCompound,
  isList,
  isString,
  type NbtTag,
  Tag,
} from "../nbt";
import { ENTITY_BY_ID, entityInfo } from "../data";
import {
  createEntity,
  entityCustomName,
  entityPos,
  entityRotation,
  normalizeType,
  setEntityPos,
  setEntityRotation,
  type EntityData,
  type Vec3,
} from "../entities";
import { resolveEntityTargets, SelectorError } from "../selectors";
import { type CommandContext, type CommandResult, executeCommandLine, feedback, print, ReturnSignal } from "./context";
import { CommandError } from "../cursor";

type Handler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

export const summonHandler: Handler = (cursor, ctx) => {
  const typeStart = cursor.mark();
  const rawType = cursor.readWord();
  const type = normalizeType(rawType);
  const info = ENTITY_BY_ID.get(type.replace("minecraft:", ""));
  if (!info && !ctx.lint) {
    cursor.error(
      `Unknown entity type '${rawType}'. Try zombie, armor_stand, item or marker.`,
      typeStart,
      "entity-type",
    );
  }
  let pos = ctx.pos;
  let nbt: CompoundTag | null = null;
  if (!cursor.done() && cursor.peek() !== "{") {
    pos = cursor.readVec3(ctx.pos, ctx.rotation);
  }
  if (!cursor.done()) {
    const parsed = cursor.readSnbt();
    if (!isCompound(parsed.tag)) {
      cursor.error("Entity NBT must be a compound", parsed.start, "nbt-type");
    }
    nbt = parsed.tag as CompoundTag;
  }
  if (ctx.lint) return { result: 1, success: true };
  const entity = createEntity(type, pos, nbt, { isPlayer: false });
  ctx.world.addEntity(entity);
  feedback(ctx, "success", `Summoned new ${info?.name ?? type}`);
  return { result: 1, success: true };
};

export const killHandler: Handler = (cursor, ctx) => {
  let raw = "@s";
  if (!cursor.done()) raw = cursor.readTargetsRaw().raw;
  let entities: EntityData[];
  try {
    entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
  } catch (error) {
    if (error instanceof SelectorError) cursor.error(error.message, cursor.mark(), "selector");
    throw error;
  }
  if (ctx.lint) return { result: entities.length, success: true };
  for (const entity of entities) {
    if (entity.isPlayer) {
      entity.nbt.value.Health = Tag.float(0);
    } else {
      ctx.world.removeEntity(entity);
    }
  }
  feedback(
    ctx,
    "success",
    entities.length === 1
      ? `Killed ${entityCustomName(entities[0])}`
      : `Killed ${entities.length} entities`,
  );
  return { result: entities.length, success: true };
};

function lookAt(from: Vec3, to: Vec3): { yaw: number; pitch: number } {
  const dx = to.x - from.x;
  const dy = to.y - from.y;
  const dz = to.z - from.z;
  const horizontal = Math.hypot(dx, dz);
  return {
    yaw: (Math.atan2(-dx, dz) * 180) / Math.PI,
    pitch: (-Math.atan2(dy, horizontal) * 180) / Math.PI,
  };
}

export const teleportHandler: Handler = (cursor, ctx) => {
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  let entities: EntityData[];
  try {
    entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
  } catch (error) {
    if (error instanceof SelectorError) cursor.error(error.message, targetsStart, "selector");
    throw error;
  }

  let pos = ctx.pos;
  let rotation = ctx.rotation;
  const next = cursor.peekWord();
  const looksLikeCoordinate = /^[~^+\-.\d]/.test(next);
  if (looksLikeCoordinate) {
    pos = cursor.readVec3(ctx.pos, ctx.rotation);
    if (!cursor.done()) {
      if (cursor.peekWord() === "facing") {
        cursor.readWord();
        if (cursor.peekWord() === "entity") {
          cursor.readWord();
          const targetRaw = cursor.readTargetsRaw().raw;
          const anchor = cursor.done() ? "feet" : cursor.readWord();
          const target = resolveEntityTargets(ctx.world, targetRaw, {
            source: ctx.source,
            origin: ctx.pos,
            lenient: ctx.lint,
          })[0];
          if (target) {
            const targetPos = entityPos(target);
            rotation = lookAt(pos, {
              ...targetPos,
              y: targetPos.y + (anchor === "eyes" ? 1.62 : 0),
            });
          }
        } else {
          rotation = lookAt(pos, cursor.readVec3(ctx.pos, ctx.rotation));
        }
      } else {
        rotation = cursor.readVec2();
      }
    }
  } else {
    const destinationStart = cursor.mark();
    const destinationRaw = cursor.readTargetsRaw().raw;
    let destination: EntityData[];
    try {
      destination = resolveEntityTargets(ctx.world, destinationRaw, {
        source: ctx.source,
        origin: ctx.pos,
        lenient: ctx.lint,
      });
    } catch (error) {
      if (error instanceof SelectorError) cursor.error(error.message, destinationStart, "selector");
      throw error;
    }
    if (destination[0]) {
      pos = entityPos(destination[0]);
      rotation = entityRotation(destination[0]);
    }
  }

  if (!ctx.lint) {
    for (const entity of entities) {
      setEntityPos(entity, pos);
      setEntityRotation(entity, rotation);
    }
  }
  feedback(
    ctx,
    "success",
    entities.length === 1
      ? `Teleported ${entityCustomName(entities[0])} to ${pos.x.toFixed(1)} ${pos.y.toFixed(1)} ${pos.z.toFixed(1)}`
      : `Teleported ${entities.length} entities`,
  );
  return { result: entities.length, success: true };
};

export const tagHandler: Handler = (cursor, ctx) => {
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  let entities: EntityData[];
  try {
    entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
  } catch (error) {
    if (error instanceof SelectorError) cursor.error(error.message, targetsStart, "selector");
    throw error;
  }
  const action = cursor.readWord();
  if (!["add", "remove", "list"].includes(action)) {
    cursor.error(`Expected add, remove or list, got '${action}'`, cursor.mark() - action.length);
  }
  if (action === "list") {
    for (const entity of entities) {
      const tags = tagsOf(entity);
      print(
        ctx,
        "info",
        tags.length === 0
          ? `${entityCustomName(entity)} has no tags`
          : `${entityCustomName(entity)} has ${tags.length} tag(s): ${tags.join(", ")}`,
      );
    }
    return { result: entities.length, success: true };
  }
  const nameStart = cursor.mark();
  const name = cursor.readWord();
  if (name.length === 0) cursor.error("Expected a tag name", nameStart);
  if (!ctx.lint) {
    for (const entity of entities) {
      const tags = tagsOf(entity);
      if (action === "add" && !tags.includes(name)) tags.push(name);
      if (action === "remove") {
        const index = tags.indexOf(name);
        if (index >= 0) tags.splice(index, 1);
      }
      entity.nbt.value.Tags = Tag.list(tags.map((t) => Tag.string(t)), "string");
    }
  }
  feedback(
    ctx,
    "success",
    action === "add"
      ? `Added tag '${name}' to ${entities.length} entities`
      : `Removed tag '${name}' from ${entities.length} entities`,
  );
  return { result: entities.length, success: true };
};

export const sayHandler: Handler = (cursor, ctx) => {
  const message = cursor.readRest();
  if (message.length === 0) cursor.error("Expected a message", cursor.mark());
  const name = ctx.source ? entityCustomName(ctx.source) : "Server";
  const component = Tag.list([
    Tag.compound({ text: Tag.string("["), color: Tag.string("gray") }),
    Tag.compound({ text: Tag.string(name), color: Tag.string("gray") }),
    Tag.compound({ text: Tag.string("] "), color: Tag.string("gray") }),
    Tag.compound({ text: Tag.string(message) }),
  ]);
  if (!ctx.lint) ctx.world.chat(component, `[${name}] ${message}`);
  return { result: 1, success: true };
};

export const helpHandler: Handler = (cursor, ctx) => {
  feedback(
    ctx,
    "info",
    "Level 0 commands: /help, /list, /me, /trigger",
  );
  return { result: 1, success: true };
};

export const listHandler: Handler = (_cursor, ctx) => {
  const players = ctx.world.entities
    .filter((entity) => entity.isPlayer && entity.username)
    .map((entity) => entity.username);
  feedback(ctx, "info", `Players online (${players.length}): ${players.join(", ")}`);
  return { result: players.length, success: true };
};

export const meHandler: Handler = (cursor, ctx) => {
  const message = cursor.readRest();
  const name = ctx.source ? entityCustomName(ctx.source) : "Server";
  const component = Tag.list([
    Tag.compound({ text: Tag.string("* "), color: Tag.string("gray") }),
    Tag.compound({ text: Tag.string(name), color: Tag.string("gray") }),
    Tag.compound({ text: Tag.string(` ${message}`), italic: Tag.byte(1) }),
  ]);
  if (!ctx.lint) ctx.world.chat(component, `* ${name} ${message}`);
  return { result: 1, success: true };
};

export const tellrawHandler: Handler = (cursor, ctx) => {
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  const component = cursor.readComponent();
  if (ctx.lint) return { result: 1, success: true };
  let entities: EntityData[] = [];
  try {
    entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: true,
    });
  } catch {
    entities = [];
  }
  if (entities.length > 0) {
    const text = ctx.world.plain(component.tag, ctx.source);
    ctx.world.chat(component.tag, text);
  }
  void targetsStart;
  return { result: entities.length, success: entities.length > 0 };
};

export const returnHandler: Handler = (cursor, ctx) => {
  if (cursor.done()) {
    cursor.error("Expected a value, 'run <command>', or 'fail'", cursor.mark(), "return");
  }
  const next = cursor.peekWord();
  if (next === "run") {
    cursor.readWord();
    const rest = cursor.readRest();
    const result = executeCommandLine(rest, ctx);
    throw new ReturnSignal(result.result);
  }
  if (next === "fail") {
    cursor.readWord();
    throw new ReturnSignal(0);
  }
  const value = cursor.readInt();
  throw new ReturnSignal(value.value);
};

export const setblockHandler: Handler = (cursor, ctx) => {
  const pos = cursor.readVec3(ctx.pos, ctx.rotation);
  const blockStart = cursor.mark();
  const block = cursor.readWord();
  if (!/^[a-z0-9_.-]+(:[a-z0-9_./-]+)?$/.test(block)) {
    cursor.error(`Invalid block id '${block}'`, blockStart, "block-id");
  }
  if (!cursor.done()) {
    const mode = cursor.readWord();
    if (!["replace", "destroy", "keep"].includes(mode)) {
      cursor.error(`Expected replace, destroy or keep, got '${mode}'`, cursor.mark() - mode.length);
    }
  }
  if (!ctx.lint) ctx.world.setBlock(pos, normalizeType(block));
  feedback(ctx, "success", `Set block at ${pos.x} ${pos.y} ${pos.z} to ${normalizeType(block)}`);
  return { result: 1, success: true };
};

export const reloadHandler: Handler = (cursor, ctx) => {
  if (!cursor.done()) cursor.error("Unexpected arguments", cursor.mark());
  feedback(ctx, "success", "Reloaded datapack functions. Run F5 to reset the world and re-run load tags.");
  return { result: 1, success: true };
};

function tagsOf(entity: EntityData): string[] {
  const tags = entity.nbt.value.Tags;
  if (!isList(tags)) return [];
  return tags.value
    .filter((tag) => isString(tag))
    .map((tag) => (tag as { value: string }).value);
}

void isCompound;
void entityInfo;
void CommandError;

export const forceloadHandler: Handler = (cursor, ctx) => {
  const action = cursor.readWord();
  if (!["add", "remove", "query"].includes(action)) {
    cursor.error(
      `Expected add, remove or query, got '${action}'`,
      cursor.mark() - action.length,
      "forceload-action",
    );
    return { result: 0, success: false };
  }
  const values: number[] = [];
  while (!cursor.done()) {
    values.push(cursor.readInt().value);
  }
  if (action === "add" && values.length !== 2 && values.length !== 4) {
    cursor.error(
      "Expected /forceload add <x> <z> [<x2> <z2>]",
      cursor.mark(),
      "forceload-coordinates",
    );
  }
  if (ctx.lint) return { result: 1, success: true };
  feedback(ctx, "success", `forceload ${action}: chunks marked ${action === "add" ? "loaded" : action === "remove" ? "unloaded" : "queried"}`);
  return { result: 1, success: true };
};


export const tickHandler: Handler = (cursor, ctx) => {
  const action = cursor.readWord();
  switch (action) {
    case "freeze": {
      ctx.world.tickFrozen = true;
      ctx.world.touch();
      feedback(ctx, "success", "Tick loop frozen");
      return { result: 1, success: true };
    }
    case "unfreeze": {
      ctx.world.tickFrozen = false;
      ctx.world.touch();
      feedback(ctx, "success", "Tick loop unfrozen");
      return { result: 1, success: true };
    }
    case "rate": {
      const start = cursor.mark();
      const rate = cursor.readNumber().value;
      if (!Number.isFinite(rate) || rate < 1 || rate > 10000) {
        cursor.error("Tick rate must be between 1 and 10000", start, "tick-rate");
      }
      ctx.world.tickRate = rate;
      ctx.world.touch();
      feedback(ctx, "success", `Tick rate set to ${rate} tps`);
      return { result: 1, success: true };
    }
    case "step":
    case "sprint": {
      let count = 1;
      if (!cursor.done()) count = cursor.readInt().value;
      if (count < 1 || count > 10000) {
        cursor.error("Tick count must be between 1 and 10000", cursor.mark(), "tick-count");
      }
      for (let i = 0; i < count; i++) ctx.runtime.tick();
      feedback(ctx, "success", `${action === "step" ? "Stepped" : "Sprinted"} ${count} tick${count === 1 ? "" : "s"}`);
      return { result: count, success: true };
    }
    case "query": {
      print(
        ctx,
        "info",
        `tick ${ctx.world.tick} · ${ctx.world.tickRate} tps · ${ctx.world.tickFrozen ? "frozen" : "running"}`,
      );
      return { result: 1, success: true };
    }
    default:
      cursor.error("Expected freeze, unfreeze, rate, step, sprint or query", cursor.mark(), "tick-action");
      return { result: 0, success: false };
  }
};


export const randomHandler: Handler = (cursor, ctx) => {
  const action = cursor.readWord();
  if (action === "reset") {
    const sequence = cursor.readWord();
    if (!sequence) cursor.error("Expected a sequence name", cursor.mark());
    return { result: 1, success: true };
  }
  if (action !== "value" && action !== "roll") {
    cursor.error("Expected value, roll or reset", cursor.mark() - action.length, "random-action");
    return { result: 0, success: false };
  }
  const rangeStart = cursor.mark();
  const rangeText = cursor.readWord();
  const parts = rangeText.split("..");
  const min = Number(parts[0]);
  const max = parts.length > 1 ? Number(parts[1]) : min;
  if (
    parts.length > 2 ||
    !Number.isFinite(min) ||
    !Number.isFinite(max) ||
    !Number.isInteger(min) ||
    !Number.isInteger(max) ||
    min > max
  ) {
    cursor.error(
      `Expected an integer range like 1..6 or 4, got '${rangeText}'`,
      rangeStart,
      "random-range",
    );
  }
  // An optional sequence name is accepted; the simulator uses the world RNG
  // rather than maintaining per-sequence deterministic state.
  if (!cursor.done()) cursor.readWord();
  const value = Math.floor(Math.random() * (max - min + 1)) + min;
  if (action === "roll") {
    feedback(ctx, "success", `Rolled ${value}`);
  } else {
    feedback(ctx, "success", `${value}`);
  }
  return { result: value, success: true };
};


export const advancementHandler: Handler = (cursor, ctx) => {
  const actionStart = cursor.mark();
  const action = cursor.readWord();
  if (action !== "grant" && action !== "revoke") {
    cursor.error("Expected grant or revoke", actionStart, "advancement-action");
    return { result: 0, success: false };
  }
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  let entities: EntityData[] = [];
  try {
    entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    }).filter((entity) => entity.isPlayer);
  } catch (error) {
    if (error instanceof SelectorError) cursor.error(error.message, targetsStart, "selector");
    throw error;
  }
  const modeStart = cursor.mark();
  const mode = cursor.readWord();
  if (!["only", "until", "from", "through", "everything"].includes(mode)) {
    cursor.error(
      "Expected only, until, from, through or everything",
      modeStart,
      "advancement-mode",
    );
  }
  const ids: string[] = [];
  if (mode === "everything") {
    ids.push(...ctx.world.advancements.keys());
  } else {
    const idStart = cursor.mark();
    const id = cursor.readResourceLocation();
    if (ctx.lint && !ctx.world.getAdvancement(id)) {
      cursor.error(`Unknown advancement '${id}'`, idStart, "unknown-advancement");
    }
    ids.push(id);
  }
  if (!ctx.lint) {
    for (const entity of entities) {
      const holder = entity.username ?? entity.uuid;
      for (const id of ids) {
        if (action === "grant") ctx.runtime.grantAdvancement(holder, id);
        else ctx.runtime.revokeAdvancement(holder, id);
      }
    }
  }
  feedback(
    ctx,
    "success",
    `${action === "grant" ? "Granted" : "Revoked"} ${ids.length} advancement${ids.length === 1 ? "" : "s"}`,
  );
  return { result: ids.length * entities.length, success: true };
};
