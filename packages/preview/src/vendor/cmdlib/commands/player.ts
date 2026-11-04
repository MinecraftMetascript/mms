import { Cursor } from "../cursor";
import { type CompoundTag, isCompound, type NbtTag, Tag } from "../nbt";
import { STATUS_EFFECT_IDS } from "../data";
import { isGameMode, type GameMode } from "../player";
import { normalizeItemId, parseItemComponents } from "../items";
import { applyItemModifier } from "../item_modifiers";
import { SCOREBOARD_MAX, SCOREBOARD_MIN } from "../world";
import type { EntityData } from "../entities";
import {
  holderName,
  resolveEntityTargets,
  resolveHolderTargets,
  SelectorError,
} from "../selectors";
import {
  type CommandContext,
  type CommandResult,
  feedback,
  readObjectiveName,
} from "./context";

type Handler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

const TICKS_PER_SECOND = 20;
const DEFAULT_EFFECT_SECONDS = 30;
const MAX_EFFECT_SECONDS = 1_000_000;
const MAX_EFFECT_AMPLIFIER = 255;

/* ------------------------------- targets -------------------------------- */

function resolvePlayers(
  cursor: Cursor,
  ctx: CommandContext,
  raw: string,
  start: number,
): EntityData[] {
  let entities: EntityData[];
  try {
    entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
  } catch (error) {
    if (error instanceof SelectorError) {
      cursor.error(error.message, start, "selector");
    }
    throw error;
  }
  const players = entities.filter((entity) => entity.isPlayer);
  if (players.length === 0) {
    cursor.error(
      `No player was found matching '${raw}'. This command only works on players.`,
      start,
      "player-target",
    );
  }
  return players;
}

function resolveEffectTargets(
  cursor: Cursor,
  ctx: CommandContext,
  raw: string,
  start: number,
): EntityData[] {
  let entities: EntityData[];
  try {
    entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
  } catch (error) {
    if (error instanceof SelectorError) {
      cursor.error(error.message, start, "selector");
    }
    throw error;
  }
  if (entities.length === 0 && !ctx.lint) {
    cursor.error(`No entity was found matching '${raw}'`, start, "entity-target");
  }
  return entities;
}

function entityLabel(ctx: CommandContext, entities: EntityData[]): string {
  if (entities.length === 1) {
    return ctx.world.holderDisplayName(holderName(entities[0]));
  }
  return `${entities.length} entities`;
}

function playerLabel(ctx: CommandContext, players: EntityData[]): string {
  if (players.length === 1) {
    return ctx.world.holderDisplayName(holderName(players[0]));
  }
  return `${players.length} players`;
}

function isKnownPlayerName(ctx: CommandContext, token: string): boolean {
  return ctx.world.entities.some(
    (entity) => entity.isPlayer && entity.username === token,
  );
}

/* -------------------------------- items --------------------------------- */

function readOptionalItemNbt(cursor: Cursor): CompoundTag | undefined {
  cursor.skip();
  if (cursor.done() || cursor.peek() !== "{") return undefined;
  const parsed = cursor.readSnbt();
  if (!isCompound(parsed.tag)) {
    cursor.error("Item NBT must be a compound", parsed.start, "nbt-type");
  }
  return parsed.tag;
}

function readGiveItem(
  cursor: Cursor,
): { id: string; count: number; nbt?: CompoundTag; components?: Record<string, NbtTag> } {
  const id = normalizeItemId(cursor.readResourceLocation());
  let count = 1;
  let nbt: CompoundTag | undefined;
  let components: Record<string, NbtTag> | undefined;

  cursor.skip();
  if (cursor.peek() === "[") {
    components = parseItemComponents(cursor.readBalanced());
  }
  if (cursor.peek() === "{") {
    nbt = readOptionalItemNbt(cursor);
  } else {
    const lookahead = cursor.peekWord();
    if (/^[+-]?\d+$/.test(lookahead)) {
      const parsed = cursor.readInt();
      if (parsed.value <= 0) {
        cursor.error(
          "Item count must be a positive integer",
          parsed.start,
          "item-count",
        );
      }
      count = parsed.value;
      nbt = readOptionalItemNbt(cursor);
    }
  }
  return { id, count, nbt, components };
}

/* ------------------------------ /gamemode ------------------------------- */

export const gamemodeHandler: Handler = (cursor, ctx) => {
  const modeStart = cursor.mark();
  const rawMode = cursor.readWord().toLowerCase();
  if (!isGameMode(rawMode)) {
    cursor.error(
      `Expected survival, creative, adventure or spectator, got '${rawMode}'`,
      modeStart,
      "gamemode",
    );
    return { result: 0, success: false };
  }
  const mode: GameMode = rawMode;
  const targetsStart = cursor.mark();
  const raw = cursor.done() ? "@s" : cursor.readTargetsRaw().raw;
  const players = resolvePlayers(cursor, ctx, raw, targetsStart);

  if (!ctx.lint) {
    for (const player of players) {
      ctx.world.setPlayerGameMode(player, mode);
    }
  }
  feedback(
    ctx,
    "success",
    `Set ${playerLabel(ctx, players)} to ${mode} mode`,
  );
  return { result: players.length, success: true };
};

/* -------------------------------- /give --------------------------------- */

export const giveHandler: Handler = (cursor, ctx) => {
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  const item = readGiveItem(cursor);
  const players = resolvePlayers(cursor, ctx, raw, targetsStart);

  if (!ctx.lint) {
    for (const player of players) {
      ctx.world.givePlayerItem(player, item.id, item.count, item.nbt, item.components);
    }
  }
  feedback(
    ctx,
    "success",
    `Gave ${item.count} ${item.id} to ${playerLabel(ctx, players)}`,
  );
  return { result: item.count * players.length, success: true };
};

/* -------------------------------- /clear -------------------------------- */

export const clearHandler: Handler = (cursor, ctx) => {
  let raw = "@s";
  let itemId: string | undefined;
  let itemNbt: CompoundTag | undefined;
  let itemComponents: Record<string, NbtTag> | undefined;
  let maxCount: number | undefined;

  if (!cursor.done()) {
    const token = cursor.peekWord();
    if (token.startsWith("@") || isKnownPlayerName(ctx, token)) {
      raw = cursor.readTargetsRaw().raw;
    }
  }
  if (!cursor.done()) {
    itemId = normalizeItemId(cursor.readResourceLocation());
    cursor.skip();
    if (cursor.peek() === "[") {
      itemComponents = parseItemComponents(cursor.readBalanced());
    }
    itemNbt = readOptionalItemNbt(cursor);
  }
  if (!cursor.done()) {
    const parsed = cursor.readInt();
    if (parsed.value < 0) {
      cursor.error(
        "maxCount must be zero or a positive integer",
        parsed.start,
        "clear-count",
      );
    }
    maxCount = parsed.value;
  }

  const targetsStart = cursor.mark();
  const players = resolvePlayers(cursor, ctx, raw, targetsStart);
  let removed = 0;
  if (!ctx.lint) {
    for (const player of players) {
      removed += ctx.world.clearPlayerItems(
        player,
        itemId,
        itemNbt,
        maxCount,
        itemComponents,
      );
    }
  }
  feedback(
    ctx,
    "success",
    itemId
      ? `Removed ${removed} ${itemId} from ${playerLabel(ctx, players)}`
      : `Cleared ${removed} item(s) from ${playerLabel(ctx, players)}`,
  );
  return { result: removed, success: true };
};

/* -------------------------------- /effect ------------------------------- */

function readEffectId(cursor: Cursor): string {
  const start = cursor.mark();
  const id = cursor.readResourceLocation();
  if (!STATUS_EFFECT_IDS.has(id)) {
    const raw = cursor.text.slice(start, cursor.mark());
    cursor.error(
      `Unknown status effect '${raw}'. Try speed, regeneration or resistance.`,
      start,
      "effect-id",
    );
  }
  return id;
}

export const effectHandler: Handler = (cursor, ctx) => {
  const actionStart = cursor.mark();
  const action = cursor.readWord().toLowerCase();
  if (action === "give") return giveEffect(cursor, ctx);
  if (action === "clear") return clearEffect(cursor, ctx);
  cursor.error(
    `Expected 'give' or 'clear', got '${action || "end of line"}'`,
    actionStart,
    "effect-action",
  );
  return { result: 0, success: false };
};

function giveEffect(cursor: Cursor, ctx: CommandContext): CommandResult {
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  const effectId = readEffectId(cursor);

  let seconds = DEFAULT_EFFECT_SECONDS;
  if (!cursor.done()) {
    const parsed = cursor.readInt();
    if (parsed.value <= 0 || parsed.value > MAX_EFFECT_SECONDS) {
      cursor.error(
        `Duration must be between 1 and ${MAX_EFFECT_SECONDS} seconds`,
        parsed.start,
        "effect-duration",
      );
    }
    seconds = parsed.value;
  }

  let amplifier = 0;
  if (!cursor.done()) {
    const parsed = cursor.readInt();
    if (parsed.value < 0 || parsed.value > MAX_EFFECT_AMPLIFIER) {
      cursor.error(
        `Amplifier must be between 0 and ${MAX_EFFECT_AMPLIFIER}`,
        parsed.start,
        "effect-amplifier",
      );
    }
    amplifier = parsed.value;
  }

  let hidden = false;
  if (!cursor.done()) {
    const booleanStart = cursor.mark();
    const value = cursor.readWord().toLowerCase();
    if (value === "true") hidden = true;
    else if (value === "false") hidden = false;
    else {
      cursor.error(
        `Expected 'true' or 'false', got '${value}'`,
        booleanStart,
        "effect-hide-particles",
      );
    }
  }

  const targets = resolveEffectTargets(cursor, ctx, raw, targetsStart);
  if (!ctx.lint) {
    for (const target of targets) {
      ctx.world.applyStatusEffect(target, {
        id: effectId,
        amplifier,
        durationTicks: seconds * TICKS_PER_SECOND,
        hidden,
      });
    }
  }
  feedback(
    ctx,
    "success",
    `Applied ${effectId} (${seconds}s, amplifier ${amplifier}) to ${entityLabel(ctx, targets)}`,
  );
  return { result: targets.length, success: true };
}

function clearEffect(cursor: Cursor, ctx: CommandContext): CommandResult {
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  let effectId: string | undefined;
  if (!cursor.done()) effectId = readEffectId(cursor);
  const targets = resolveEffectTargets(cursor, ctx, raw, targetsStart);

  let removed = 0;
  if (!ctx.lint) {
    for (const target of targets) {
      removed += ctx.world.clearStatusEffects(target, effectId);
    }
  }
  feedback(
    ctx,
    "success",
    effectId
      ? `Removed ${removed} '${effectId}' effect(s) from ${entityLabel(ctx, targets)}`
      : `Removed ${removed} effect(s) from ${entityLabel(ctx, targets)}`,
  );
  return { result: removed, success: true };
}

/* ------------------------------- /spectate ------------------------------ */

export const spectateHandler: Handler = (cursor, ctx) => {
  const targetStart = cursor.mark();
  const targetRaw = cursor.readTargetsRaw().raw;
  const targets = resolveEffectTargets(cursor, ctx, targetRaw, targetStart);
  if (targets.length !== 1) {
    cursor.error(
      `Expected exactly one entity to spectate, got ${targets.length}`,
      targetStart,
      "spectate-target",
    );
    return { result: 0, success: false };
  }
  const viewerStart = cursor.mark();
  const viewerRaw = cursor.done() ? "@s" : cursor.readTargetsRaw().raw;
  const viewers = resolvePlayers(cursor, ctx, viewerRaw, viewerStart);
  if (!ctx.lint) {
    for (const viewer of viewers) {
      ctx.world.setSpectatorTarget(viewer, targets[0].uuid);
    }
  }
  feedback(
    ctx,
    "success",
    `Made ${playerLabel(ctx, viewers)} spectate ${entityLabel(ctx, targets)}`,
  );
  return { result: viewers.length, success: true };
};

/* -------------------------------- /trigger ------------------------------ */

export const triggerHandler: Handler = (cursor, ctx) => {
  const objectiveStart = cursor.mark();
  const objective = readObjectiveName(cursor);
  const definition = ctx.world.getObjective(objective);
  if (!definition) {
    cursor.error(
      `Unknown objective '${objective}'`,
      objectiveStart,
      "unknown-objective",
    );
    return { result: 0, success: false };
  }
  if (definition.criterion !== "trigger") {
    cursor.error(
      `Objective '${objective}' uses the '${definition.criterion}' criterion, not 'trigger'`,
      objectiveStart,
      "trigger-objective",
    );
    return { result: 0, success: false };
  }

  const source = ctx.source;
  if (!source || !source.isPlayer) {
    cursor.error(
      "Only players can use /trigger",
      cursor.mark(),
      "trigger-player",
    );
    return { result: 0, success: false };
  }
  const holder = holderName(source);
  // Lint runs files in filesystem order, which is not necessarily load-tag
  // order. Runtime stays strict; lint only validates the objective/player.
  if (!ctx.lint && !ctx.world.isTriggerEnabled(holder, objective)) {
    cursor.error(
      `Trigger '${objective}' is not enabled for ${ctx.world.holderDisplayName(holder)}. Run /scoreboard players enable first.`,
      cursor.mark(),
      "trigger-disabled",
    );
  }

  let action: "add" | "set" = "add";
  let amount = 1;
  let explicitAction = false;
  if (!cursor.done()) {
    explicitAction = true;
    const actionStart = cursor.mark();
    const rawAction = cursor.readWord().toLowerCase();
    if (rawAction !== "add" && rawAction !== "set") {
      cursor.error(
        `Expected 'add <value>' or 'set <value>', got '${rawAction}'`,
        actionStart,
        "trigger-action",
      );
    }
    action = rawAction === "set" ? "set" : "add";
    if (cursor.done()) {
      cursor.error(
        `Expected an integer after '${action}'`,
        cursor.mark(),
        "trigger-value",
      );
    }
    const parsed = cursor.readInt();
    if (parsed.value < SCOREBOARD_MIN || parsed.value > SCOREBOARD_MAX) {
      cursor.error(
        `Trigger value must be between ${SCOREBOARD_MIN} and ${SCOREBOARD_MAX}`,
        parsed.start,
        "score-range",
      );
    }
    amount = parsed.value;
  }

  let result = 1;
  if (!ctx.lint) {
    const current = ctx.world.getScore(holder, objective);
    // `/trigger <objective>` adds one when a score exists, but initialises an
    // unset score to 0 (the vanilla quirk for first use).
    const next = explicitAction
      ? action === "set"
        ? amount
        : (current ?? 0) + amount
      : current === undefined
        ? 0
        : current + 1;
    ctx.world.setScore(holder, objective, next);
    result = ctx.world.getScore(holder, objective) ?? 0;
    ctx.world.setTriggerEnabled(holder, objective, false);
  }
  feedback(
    ctx,
    "success",
    `Triggered '${objective}' for ${ctx.world.holderDisplayName(holder)} (${action} ${amount}, now ${result})`,
  );
  return { result, success: true };
};

export const damageHandler: Handler = (cursor, ctx) => {
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  const amount = cursor.readNumber().value;
  if (!Number.isFinite(amount) || amount < 0) {
    cursor.error("Damage amount must be a non-negative number", cursor.mark(), "damage");
  }
  // Optional damage type / at / by clauses are accepted but do not alter the
  // simulated health reduction.
  while (!cursor.done()) cursor.readWord();
  let entities: EntityData[] = [];
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
  if (!ctx.lint) {
    for (const entity of entities) {
      if (entity.isPlayer) {
        ctx.world.damagePlayer(entity, amount);
      } else {
        const current = entity.nbt.value.Health;
        const value = typeof current?.value === "number" ? current.value : 20;
        entity.nbt.value.Health = Tag.float(Math.max(0, value - amount));
      }
    }
    ctx.world.touch();
  }
  feedback(
    ctx,
    "success",
    `Applied ${amount} damage to ${entities.length} ${entities.length === 1 ? "entity" : "entities"}`,
  );
  return { result: entities.length, success: true };
};

export const itemHandler: Handler = (cursor, ctx) => {
  const actionStart = cursor.mark();
  const action = cursor.readWord();
  if (action !== "replace" && action !== "modify") {
    cursor.error("Expected replace or modify", actionStart, "item-action");
    return { result: 0, success: false };
  }
  const source = cursor.readWord();
  if (source !== "entity") {
    cursor.error("Only entity item slots are simulated", cursor.mark(), "item-source");
  }
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  const slotStart = cursor.mark();
  const slot = cursor.readWord();
  const players = resolvePlayers(cursor, ctx, raw, targetsStart);
  if (action === "replace") {
    cursor.expectWord("with");
    const item = readGiveItem(cursor);
    if (!ctx.lint) {
      for (const player of players) {
        if (!ctx.world.setItemAt(player, slot, {
          id: item.id,
          count: item.count,
          nbt: item.nbt,
          components: item.components,
        })) {
          cursor.error(`Unknown item slot '${slot}'`, slotStart, "item-slot");
        }
      }
    }
    feedback(ctx, "success", `Replaced ${slot} with ${item.count} ${item.id}`);
    return { result: players.length, success: true };
  }
  const modifierId = cursor.readResourceLocation();
  const modifier = ctx.world.getItemModifier(modifierId);
  if (!modifier) {
    cursor.error(`Unknown item modifier '${modifierId}'`, cursor.mark(), "unknown-item-modifier");
  }
  if (!ctx.lint) {
    for (const player of players) {
      const current = ctx.world.getItemAt(player, slot);
      if (!current) continue;
      if (modifier) {
        const next = applyItemModifier(modifier, current);
        ctx.world.setItemAt(player, slot, next);
      }
    }
  }
  feedback(ctx, "success", `Modified ${slot} with ${modifierId}`);
  return { result: players.length, success: true };
};
