import { type CompoundTag, isCompound, isList, isString, type NbtTag, Tag, numericValue } from "./nbt";
import {
  createEntity,
  entityCustomName,
  entityPos,
  entityRotation,
  setEntityPos,
  type EntityData,
} from "./entities";
import type { Vec3 } from "./entities";
import { type VersionProfile, getVersion } from "./version";
import type { TextRenderContext } from "./text";
import { formatTagForComponent } from "./text";
import {
  holderName as entityHolderName,
  parseSelector,
  resolveSingleHolder,
  selectEntities,
  SelectorError,
} from "./selectors";
import { parseSnbt } from "./snbt";
import { plainText } from "./text";
import { getPathValues, parseNbtPath } from "./nbtPath";
import {
  cloneStatusEffect,
  createPlayerState,
  type GameMode,
  type InventoryItem,
  itemComponentKey,
  itemNbtKey,
  type PlayerState,
  type StatusEffect,
} from "./player";
import type { PredicateDefinition } from "./predicates";
import {
  HOTBAR_SIZE,
  MAIN_INVENTORY_SIZE,
  itemComponentsMatch,
  normalizeItemId,
  resolveSlot,
} from "./items";
import type { AdvancementDefinition } from "./advancements";
import type { ItemModifierDefinition } from "./item_modifiers";

export interface Objective {
  name: string;
  criterion: string;
  displayName: NbtTag;
  renderType: "integer" | "hearts";
}

export interface TitleTimings {
  fadeIn: number;
  stay: number;
  fadeOut: number;
}

export interface TitleState {
  title?: NbtTag;
  subtitle?: NbtTag;
  actionbar?: NbtTag;
  times: TitleTimings;
  revision: number;
  titleRevision: number;
  actionbarRevision: number;
  setAt: number;
  actionbarSetAt: number;
}

export type LogKind = "chat" | "actionbar" | "error" | "warning" | "info" | "success" | "command" | "data" | "system";

export interface WorldLogEntry {
  id: number;
  kind: LogKind;
  text: string;
  component?: NbtTag;
  time: number;
}

export interface FunctionDef {
  id: string;
  path: string;
  lines: string[];
  macroKeys: string[];
  hasMacroLines: boolean;
}

export interface FunctionTag {
  id: string;
  values: string[];
  /** Source file, when loaded from a datapack. */
  path?: string;
}

/**
 * One `/schedule function` entry. A function id may have more than one entry
 * when it was queued with `append`; `sequence` keeps that queue stable when
 * two entries share a due tick.
 */
/** Minecraft scoreboards store Java 32-bit signed integers. */
export const SCOREBOARD_MIN = -2147483648;
export const SCOREBOARD_MAX = 2147483647;

export function clampScore(value: number): number {
  const whole = Math.trunc(value);
  if (Number.isNaN(whole)) return 0;
  if (whole > SCOREBOARD_MAX) return SCOREBOARD_MAX;
  if (whole < SCOREBOARD_MIN) return SCOREBOARD_MIN;
  return whole;
}

export interface ScheduledFunction {
  /** Function resource location, e.g. `cmdlib:tick`. */
  functionId: string;
  /** World tick at which the command was run. */
  scheduledAtTick: number;
  /** Absolute world tick on which the function becomes due. */
  dueTick: number;
  /** Monotonic insertion order, used to break due-tick ties. */
  sequence: number;
}

export interface WorldOptions {
  version: VersionProfile;
  seedPlayer?: boolean;
  playerName?: string;
  playerPos?: Vec3;
}

let logId = 0;

export class World {
  readonly version: VersionProfile;
  tick = 0;
  /** Tick rate in ticks per second (Minecraft's /tick rate). */
  tickRate = 20;
  /** When true, the timeline's automatic tick loop is frozen. */
  tickFrozen = false;
  entities: EntityData[] = [];
  /**
   * Simulated player state, keyed by score-holder name (username for players).
   * The map is public for inspection; commands mutate through the helpers in
   * the "player runtime" section below.
   */
  playerStates = new Map<string, PlayerState>();
  /** Active status effects, keyed by the same holder key. */
  statusEffects = new Map<string, StatusEffect[]>();
  storages = new Map<string, CompoundTag>();
  objectives = new Map<string, Objective>();
  scores = new Map<string, Map<string, number>>();
  displaySlots = new Map<string, string | null>();
  holderNames = new Map<string, string>();
  functions = new Map<string, FunctionDef>();
  functionTags = new Map<string, FunctionTag>();
  /** Datapack predicates, keyed by resource location. */
  predicates = new Map<string, PredicateDefinition>();
  /** Datapack advancements, keyed by resource location. */
  advancements = new Map<string, AdvancementDefinition>();
  /** Datapack item modifiers, keyed by resource location. */
  itemModifiers = new Map<string, ItemModifierDefinition>();
  /** Completed advancements per score holder (player username). */
  advancementProgress = new Map<string, Set<string>>();
  blocks = new Map<string, string>();
  logs: WorldLogEntry[] = [];
  titles: TitleState = {
    times: { fadeIn: 10, stay: 70, fadeOut: 20 },
    revision: 0,
    titleRevision: 0,
    actionbarRevision: 0,
    setAt: 0,
    actionbarSetAt: 0,
  };
  revision = 0;
  selectedEntity: string | null = null;
  private listeners = new Set<() => void>();
  private commandBudget = 0;
  private scheduleQueue: ScheduledFunction[] = [];
  private scheduleSequence = 0;

  constructor(options: WorldOptions) {
    this.version = options.version ?? getVersion("26.3");
    if (options.seedPlayer !== false) {
      const player = createEntity(
        "minecraft:player",
        options.playerPos ?? { x: 0, y: 64, z: 0 },
        null,
        { isPlayer: true, username: options.playerName ?? "Dev" },
      );
      this.entities.push(player);
    }
  }

  /* ------------------------------ events ------------------------------ */

  subscribe = (listener: () => void): (() => void) => {
    this.listeners.add(listener);
    return () => this.listeners.delete(listener);
  };

  touch(): void {
    this.revision++;
    for (const listener of this.listeners) listener();
  }

  resetCommandBudget(): void {
    this.commandBudget = 0;
  }

  consumeCommand(count = 1): void {
    this.commandBudget += count;
    if (this.commandBudget > 65536) {
      throw new Error("Command chain exceeded 65536 commands (maxCommandChainLength)");
    }
  }

  /* ---------------------------- scoreboard ---------------------------- */

  addObjective(name: string, criterion: string, displayName?: NbtTag): Objective {
    if (this.objectives.has(name)) {
      throw new Error(`An objective named '${name}' already exists`);
    }
    const objective: Objective = {
      name,
      criterion,
      displayName: displayName ?? Tag.string('{"text":""}'),
      renderType: "integer",
    };
    this.objectives.set(name, objective);
    if (!this.scores.has(name)) this.scores.set(name, new Map());
    return objective;
  }

  removeObjective(name: string): boolean {
    const existed = this.objectives.delete(name);
    this.scores.delete(name);
    for (const [slot, objective] of this.displaySlots) {
      if (objective === name) this.displaySlots.set(slot, null);
    }
    return existed;
  }

  getObjective(name: string): Objective | undefined {
    return this.objectives.get(name);
  }

  requireObjective(name: string): Objective {
    const objective = this.objectives.get(name);
    if (!objective) throw new Error(`Unknown objective '${name}'`);
    return objective;
  }

  setScore(holder: string, objective: string, value: number): void {
    this.requireObjective(objective);
    let map = this.scores.get(objective);
    if (!map) {
      map = new Map();
      this.scores.set(objective, map);
    }
    map.set(holder, clampScore(value));
  }

  getScore(holder: string, objective: string): number | undefined {
    return this.scores.get(objective)?.get(holder);
  }

  addScore(holder: string, objective: string, delta: number): number {
    this.requireObjective(objective);
    const next = clampScore((this.getScore(holder, objective) ?? 0) + delta);
    this.setScore(holder, objective, next);
    return next;
  }

  resetScore(holder: string, objective: string): void {
    this.scores.get(objective)?.delete(holder);
  }

  resetAllScores(holder: string): void {
    for (const map of this.scores.values()) map.delete(holder);
  }

  scoreObjectives(holder: string): string[] {
    const out: string[] = [];
    for (const [objective, map] of this.scores) {
      if (map.has(holder)) out.push(objective);
    }
    return out;
  }

  setDisplaySlot(slot: string, objective: string | null): void {
    this.displaySlots.set(slot, objective);
  }

  setHolderDisplayName(holder: string, name: string): void {
    this.holderNames.set(holder, name);
  }

  holderDisplayName(holder: string): string {
    const explicit = this.holderNames.get(holder);
    if (explicit) return explicit;
    const entity = this.entities.find((e) => entityHolderName(e) === holder);
    if (entity) return entityCustomName(entity);
    return holder;
  }

  sidebar(): { objective: Objective; entries: Array<{ holder: string; display: string; value: number }> } | null {
    const objectiveName = this.displaySlots.get("sidebar");
    if (!objectiveName) return null;
    const objective = this.objectives.get(objectiveName);
    if (!objective) return null;
    const map = this.scores.get(objectiveName) ?? new Map();
    const entries = [...map.entries()]
      .map(([holder, value]) => ({
        holder,
        display: this.holderDisplayName(holder),
        value,
      }))
      .sort((a, b) => b.value - a.value || a.holder.localeCompare(b.holder))
      .slice(0, 15);
    return { objective, entries };
  }

  /* ------------------------------- titles ----------------------------- */

  setTitle(kind: "title" | "subtitle" | "actionbar", component: NbtTag): void {
    if (kind === "title") this.titles.title = component;
    else if (kind === "subtitle") this.titles.subtitle = component;
    else {
      this.titles.actionbar = component;
      this.titles.actionbarSetAt = Date.now();
      this.titles.actionbarRevision++;
      this.titles.revision++;
      return;
    }
    this.titles.titleRevision++;
    this.titles.setAt = Date.now();
    this.titles.revision++;
  }

  clearTitle(kind: "title" | "subtitle" | "actionbar" | "all"): void {
    if (kind === "title" || kind === "all") {
      this.titles.title = undefined;
      this.titles.titleRevision++;
    }
    if (kind === "subtitle" || kind === "all") {
      this.titles.subtitle = undefined;
      this.titles.titleRevision++;
    }
    if (kind === "actionbar" || kind === "all") {
      this.titles.actionbar = undefined;
      this.titles.actionbarRevision++;
    }
    this.titles.revision++;
  }

  setTitleTimes(fadeIn: number, stay: number, fadeOut: number): void {
    this.titles.times = { fadeIn, stay, fadeOut };
    this.titles.titleRevision++;
    this.titles.revision++;
  }

  /* ------------------------------ entities ---------------------------- */

  addEntity(entity: EntityData): void {
    this.entities.push(entity);
  }

  removeEntity(entity: EntityData): void {
    const index = this.entities.indexOf(entity);
    if (index >= 0) this.entities.splice(index, 1);
    if (this.selectedEntity === entity.uuid) this.selectedEntity = null;
  }

  player(): EntityData | undefined {
    return this.entities.find((e) => e.isPlayer);
  }

  entityByUuid(uuid: string): EntityData | undefined {
    return this.entities.find((e) => e.uuid === uuid);
  }

  /* --------------------------- player runtime ------------------------- */

  /**
   * Resolve a target to its runtime state key. Players use their username so
   * state survives an entity object being replaced (teleport, re-summon, …);
   * other entities fall back to their UUID.
   */
  runtimeKey(target?: EntityData | string | null): string {
    if (typeof target === "string") return target;
    if (target) return entityHolderName(target);
    const player = this.player();
    return player ? entityHolderName(player) : "Dev";
  }

  /** Read-only player state, or undefined when the holder has never acted. */
  playerState(target?: EntityData | string | null): PlayerState | undefined {
    return this.playerStates.get(this.runtimeKey(target));
  }

  playerGameMode(target?: EntityData | string | null): GameMode {
    return this.playerState(target)?.gameMode ?? "survival";
  }

  /** Read-only inventory view for the UI. Do not mutate the returned array. */
  playerInventory(target?: EntityData | string | null): readonly InventoryItem[] {
    return this.playerState(target)?.inventory ?? [];
  }

  /** Read-only status effects for a player or any entity holder. */
  effectsOf(target?: EntityData | string | null): readonly StatusEffect[] {
    return this.statusEffects.get(this.runtimeKey(target)) ?? [];
  }

  isTriggerEnabled(
    target: EntityData | string | null,
    objective: string,
  ): boolean {
    return this.playerState(target)?.enabledTriggers.has(objective) ?? false;
  }

  spectatorTarget(target?: EntityData | string | null): string | undefined {
    return this.playerState(target)?.spectating;
  }

  setSpectatorTarget(
    target: EntityData | string | null,
    entityUuid: string | null,
  ): void {
    const state = this.ensurePlayerState(target);
    if (entityUuid) state.spectating = entityUuid;
    else delete state.spectating;
  }

  private ensurePlayerState(target?: EntityData | string | null): PlayerState {
    const key = this.runtimeKey(target);
    let state = this.playerStates.get(key);
    if (!state) {
      state = createPlayerState();
      this.playerStates.set(key, state);
    }
    return state;
  }

  setPlayerGameMode(
    target: EntityData | string | null,
    mode: GameMode,
  ): void {
    this.ensurePlayerState(target).gameMode = mode;
  }

  /**
   * Add items to a holder's inventory, merging stacks with the same id and NBT
   * payload. Returns the affected stack.
   */
  givePlayerItem(
    target: EntityData | string | null,
    id: string,
    count: number,
    nbt?: CompoundTag,
    components?: Record<string, NbtTag>,
  ): InventoryItem {
    const state = this.ensurePlayerState(target);
    const normalized = normalizeItemId(id);
    const key = itemNbtKey(nbt);
    const componentKey = itemComponentKey(components);
    const existing = state.inventory.find(
      (item) =>
        item.id === normalized &&
        itemNbtKey(item.nbt) === key &&
        itemComponentKey(item.components) === componentKey,
    );
    if (existing) {
      existing.count += count;
      this.syncInventoryNbt(target);
      return existing;
    }
    const used = new Set(state.inventory.map((item) => item.slot).filter(Boolean));
    let slot: string | undefined;
    for (let i = 0; i < HOTBAR_SIZE; i++) {
      if (!used.has(`hotbar.${i}`)) {
        slot = `hotbar.${i}`;
        break;
      }
    }
    if (!slot) {
      for (let i = 0; i < MAIN_INVENTORY_SIZE; i++) {
        if (!used.has(`inventory.${i}`)) {
          slot = `inventory.${i}`;
          break;
        }
      }
    }
    const item: InventoryItem = {
      id: normalized,
      count,
      ...(nbt ? { nbt: JSON.parse(JSON.stringify(nbt)) as CompoundTag } : {}),
      ...(components ? { components: JSON.parse(JSON.stringify(components)) } : {}),
      ...(slot ? { slot } : {}),
    };
    state.inventory.push(item);
    this.syncInventoryNbt(target);
    return item;
  }

  getItemAt(target: EntityData | string | null, slot: string): InventoryItem | undefined {
    const state = this.playerState(target);
    if (!state) return undefined;
    const resolved = resolveSlot(slot, state.selectedSlot);
    if (!resolved) return undefined;
    const collection = resolved.storage === "enderchest" ? state.enderchest : state.inventory;
    const exact = collection.find((item) => item.slot === resolved.key);
    if (exact) return exact;
    if (resolved.key === `hotbar.${state.selectedSlot}`) {
      return collection.find((item) => !item.slot);
    }
    return undefined;
  }

  setItemAt(
    target: EntityData | string | null,
    slot: string,
    item: InventoryItem | null,
  ): boolean {
    if (!target) return false;
    const state = this.ensurePlayerState(target);
    const resolved = resolveSlot(slot, state.selectedSlot);
    if (!resolved) return false;
    const collection = resolved.storage === "enderchest" ? state.enderchest : state.inventory;
    const index = collection.findIndex((entry) => entry.slot === resolved.key);
    if (index >= 0) collection.splice(index, 1);
    if (item) {
      collection.push({ ...item, slot: resolved.key });
    }
    this.syncInventoryNbt(target);
    this.touch();
    return true;
  }

  renamePlayer(username: string): boolean {
    const player = this.player();
    if (!player || !player.isPlayer) return false;
    const existing = this.entities.find(
      (entity) => entity !== player && entity.isPlayer && entity.username === username,
    );
    if (existing) return false;
    const oldKey = entityHolderName(player);
    const state = this.playerStates.get(oldKey);
    player.username = username;
    player.nbt.value.CustomName = Tag.compound({ text: Tag.string(username) });
    if (state) {
      this.playerStates.delete(oldKey);
      this.playerStates.set(username, state);
    }
    this.touch();
    return true;
  }

  playerSelectedSlot(target: EntityData | string | null = null): number {
    return this.playerState(target)?.selectedSlot ?? 0;
  }

  playerHealth(target: EntityData | string | null = null): number {
    const entity = typeof target === "string"
      ? this.entities.find((entry) => entry.username === target || entry.uuid === target)
      : target ?? this.player();
    const health = entity?.nbt.value.Health;
    return typeof health?.value === "number" ? Math.max(0, health.value) : 20;
  }

  damagePlayer(target: EntityData | string | null, amount: number): number {
    const entity = typeof target === "string"
      ? this.entities.find((entry) => entry.username === target || entry.uuid === target)
      : target ?? this.player();
    if (!entity || amount <= 0) return this.playerHealth(entity ?? null);
    const next = Math.max(0, this.playerHealth(entity) - amount);
    entity.nbt.value.Health = Tag.float(next);
    this.touch();
    return next;
  }

  enderchestItems(target: EntityData | string | null): readonly InventoryItem[] {
    return this.playerState(target)?.enderchest ?? [];
  }

  /**
   * Remove matching items from a holder's inventory.
   *
   * With no `id`, the whole inventory is cleared. With `id` and no `maxCount`,
   * every matching stack is removed. `maxCount` limits how many items are
   * removed, stopping part-way through a stack. Returns the removed count.
   */
  clearPlayerItems(
    target: EntityData | string | null,
    id?: string,
    nbt?: CompoundTag,
    maxCount?: number,
    components?: Record<string, NbtTag>,
  ): number {
    const state = this.playerState(target);
    if (!state || state.inventory.length === 0) return 0;
    if (id === undefined) {
      const removed = state.inventory.reduce((sum, item) => sum + item.count, 0);
      state.inventory.length = 0;
      this.syncInventoryNbt(target);
      return removed;
    }
    const nbtKey = nbt === undefined ? undefined : itemNbtKey(nbt);
    const componentKey = components === undefined ? undefined : itemComponentKey(components);
    let budget = maxCount === undefined ? Infinity : Math.max(0, Math.trunc(maxCount));
    let removed = 0;
    const remaining: InventoryItem[] = [];
    for (const item of state.inventory) {
      const matches =
        item.id === id &&
        (nbtKey === undefined || itemNbtKey(item.nbt) === nbtKey) &&
        (componentKey === undefined ||
          (components !== undefined && itemComponentsMatch(item.components, components)));
      if (!matches || budget <= 0) {
        remaining.push(item);
        continue;
      }
      const take = Math.min(item.count, budget);
      removed += take;
      budget -= take;
      if (take < item.count) remaining.push({ ...item, count: item.count - take });
    }
    state.inventory.length = 0;
    state.inventory.push(...remaining);
    this.syncInventoryNbt(target);
    return removed;
  }

  /** Apply (or refresh) one effect on a holder, replacing the same effect id. */
  applyStatusEffect(
    target: EntityData | string | null,
    effect: StatusEffect,
  ): void {
    const key = this.runtimeKey(target);
    let list = this.statusEffects.get(key);
    if (!list) {
      list = [];
      this.statusEffects.set(key, list);
    }
    const index = list.findIndex((entry) => entry.id === effect.id);
    if (index >= 0) list[index] = cloneStatusEffect(effect);
    else list.push(cloneStatusEffect(effect));
  }

  /** Remove one effect id, or every effect when `id` is omitted. */
  clearStatusEffects(
    target: EntityData | string | null,
    id?: string,
  ): number {
    const key = this.runtimeKey(target);
    const list = this.statusEffects.get(key);
    if (!list) return 0;
    if (id === undefined) {
      const removed = list.length;
      this.statusEffects.delete(key);
      return removed;
    }
    const remaining = list.filter((effect) => effect.id !== id);
    const removed = list.length - remaining.length;
    if (remaining.length > 0) this.statusEffects.set(key, remaining);
    else this.statusEffects.delete(key);
    return removed;
  }

  /**
   * Advance every effect by one tick, dropping effects whose duration reaches
   * zero. Returns the number of effects that expired.
   */
  tickStatusEffects(): number {
    let expired = 0;
    for (const [key, list] of [...this.statusEffects]) {
      const remaining: StatusEffect[] = [];
      for (const effect of list) {
        effect.durationTicks -= 1;
        if (effect.durationTicks > 0) remaining.push(effect);
        else expired += 1;
      }
      if (remaining.length > 0) this.statusEffects.set(key, remaining);
      else this.statusEffects.delete(key);
    }
    return expired;
  }

  setTriggerEnabled(
    target: EntityData | string | null,
    objective: string,
    enabled = true,
  ): void {
    const state = this.ensurePlayerState(target);
    if (enabled) state.enabledTriggers.add(objective);
    else state.enabledTriggers.delete(objective);
  }

  /* ------------------------------- data ------------------------------- */

  storage(id: string, create = false): CompoundTag | undefined {
    const existing = this.storages.get(id);
    if (existing) return existing;
    if (!create) return undefined;
    const compound = Tag.compound();
    this.storages.set(id, compound);
    return compound;
  }

  getBlockNbt(pos: Vec3): CompoundTag {
    const key = `${pos.x},${pos.y},${pos.z}`;
    const id = this.blocks.get(key);
    if (!id) return Tag.compound();
    return Tag.compound({ id: Tag.string(id) });
  }

  setBlock(pos: Vec3, id: string): void {
    this.blocks.set(`${pos.x},${pos.y},${pos.z}`, id);
  }

  getBlock(pos: Vec3): string | undefined {
    return this.blocks.get(`${pos.x},${pos.y},${pos.z}`);
  }

  /* ------------------------------ functions --------------------------- */

  setFunctions(functions: FunctionDef[]): void {
    this.functions.clear();
    for (const fn of functions) this.functions.set(fn.id, fn);
  }

  setFunctionTags(tags: FunctionTag[]): void {
    this.functionTags.clear();
    for (const tag of tags) this.functionTags.set(tag.id, tag);
  }

  getFunction(id: string): FunctionDef | undefined {
    return this.functions.get(id);
  }

  getFunctionTag(id: string): FunctionTag | undefined {
    return this.functionTags.get(id);
  }

  /* ----------------------------- predicates --------------------------- */

  setPredicates(
    predicates:
      | Iterable<PredicateDefinition>
      | Map<string, PredicateDefinition>,
  ): void {
    this.predicates.clear();
    const values =
      predicates instanceof Map ? predicates.values() : predicates;
    for (const definition of values) {
      this.predicates.set(normalizePredicateId(definition.id), definition);
    }
  }

  getPredicate(id: string): PredicateDefinition | undefined {
    return this.predicates.get(normalizePredicateId(id));
  }

  predicateIds(): string[] {
    return [...this.predicates.keys()];
  }

  /* ----------------------------- advancements ------------------------- */

  setAdvancements(
    advancements:
      | Iterable<AdvancementDefinition>
      | Map<string, AdvancementDefinition>,
  ): void {
    this.advancements.clear();
    const values =
      advancements instanceof Map ? advancements.values() : advancements;
    for (const definition of values) {
      this.advancements.set(definition.id, definition);
    }
  }

  getAdvancement(id: string): AdvancementDefinition | undefined {
    return this.advancements.get(id);
  }

  isAdvancementGranted(holder: string, id: string): boolean {
    return this.advancementProgress.get(holder)?.has(id) ?? false;
  }

  grantAdvancement(holder: string, id: string): boolean {
    if (!this.advancements.has(id)) return false;
    let completed = this.advancementProgress.get(holder);
    if (!completed) {
      completed = new Set();
      this.advancementProgress.set(holder, completed);
    }
    if (completed.has(id)) return false;
    completed.add(id);
    this.touch();
    return true;
  }

  revokeAdvancement(holder: string, id: string): boolean {
    const completed = this.advancementProgress.get(holder);
    if (!completed?.delete(id)) return false;
    this.touch();
    return true;
  }

  advancementProgressFor(holder: string): ReadonlySet<string> {
    return this.advancementProgress.get(holder) ?? new Set();
  }

  setItemModifiers(modifiers: Iterable<ItemModifierDefinition> | Map<string, ItemModifierDefinition>): void {
    this.itemModifiers.clear();
    const values = modifiers instanceof Map ? modifiers.values() : modifiers;
    for (const modifier of values) this.itemModifiers.set(modifier.id, modifier);
  }

  getItemModifier(id: string): ItemModifierDefinition | undefined {
    return this.itemModifiers.get(id);
  }

  useItem(
    target: EntityData | string | null,
    slot: string,
  ): { kind: "book" | "used" | "info"; message: string } {
    const item = this.getItemAt(target, slot);
    const player = typeof target === "string"
      ? this.entities.find((entry) => entry.username === target || entry.uuid === target)
      : target ?? this.player();
    if (!item || !player) return { kind: "info", message: "That slot is empty" };

    if (item.id === "minecraft:writable_book" || item.id === "minecraft:written_book") {
      return { kind: "book", message: "Book opened" };
    }

    const consume = () => {
      const next = { ...item, count: item.count - 1 };
      this.setItemAt(player, slot, next.count > 0 ? next : null);
    };

    if (item.id === "minecraft:ender_pearl") {
      const pos = entityPos(player);
      const rotation = entityRotation(player);
      const yaw = (rotation.yaw * Math.PI) / 180;
      setEntityPos(player, {
        x: pos.x - Math.sin(yaw) * 8,
        y: pos.y,
        z: pos.z + Math.cos(yaw) * 8,
      });
      this.damagePlayer(player, 5);
      consume();
      return { kind: "used", message: "Teleported with an ender pearl" };
    }

    if (item.id === "minecraft:carrot_on_a_stick" || item.id === "minecraft:warped_fungus_on_a_stick") {
      return { kind: "info", message: "Requires a saddled pig or strider to use" };
    }

    const foodId = item.id.replace("minecraft:", "");
    const foodValues: Record<string, number> = {
      apple: 4, golden_apple: 4, enchanted_golden_apple: 4, bread: 5,
      cooked_beef: 8, cooked_porkchop: 8, cooked_mutton: 6, cooked_chicken: 6,
      carrot: 3, golden_carrot: 6, potato: 1, baked_potato: 5, beetroot: 1,
      cookie: 2, melon_slice: 2, sweet_berries: 2, glow_berries: 2,
    };
    const foodComponent = item.components?.["minecraft:food"];
    const componentNutrition = foodComponent && typeof foodComponent.value === "object" && !Array.isArray(foodComponent.value)
      ? numericValue(foodComponent.value.nutrition) ?? NaN
      : NaN;
    const nutrition = Number.isFinite(componentNutrition) ? componentNutrition : foodValues[foodId] ?? 0;
    if (nutrition > 0) {
      const state = this.ensurePlayerState(player);
      const food = player.nbt.value.foodLevel;
      const current = typeof food?.value === "number" ? food.value : 20;
      player.nbt.value.foodLevel = Tag.int(Math.min(20, current + nutrition));
      if (foodId.includes("golden_apple")) {
        player.nbt.value.Health = Tag.float(Math.min(20, this.playerHealth(player) + 4));
      }
      consume();
      return { kind: "used", message: `Ate ${foodId.replace(/_/g, " ")}` };
    }

    const potion = item.components?.["minecraft:potion_contents"];
    if (potion) {
      const potionTag = potion.type === "compound" ? potion.value.potion : undefined;
      const potionName = potionTag?.type === "string" ? potionTag.value : "";
      if (potionName.includes("healing") || potionName.includes("regeneration")) {
        const health = this.playerHealth(player);
        player.nbt.value.Health = Tag.float(Math.min(20, health + 4));
      } else {
        this.applyStatusEffect(player, {
          id: "minecraft:speed",
          amplifier: 0,
          durationTicks: 20 * 20,
          hidden: false,
        });
      }
      consume();
      return { kind: "used", message: `Used ${potionName || "potion"}` };
    }

    if (item.id === "minecraft:snowball" || item.id === "minecraft:egg") {
      consume();
      return { kind: "used", message: `Threw ${item.id.replace("minecraft:", "").replace(/_/g, " ")}` };
    }

    return { kind: "info", message: "No simulated right-click behavior for this item" };
  }

  dropItem(
    target: EntityData | string | null,
    slot: string,
    count = 1,
  ): { ok: boolean; message: string } {
    const item = this.getItemAt(target, slot);
    const player = typeof target === "string"
      ? this.entities.find((entry) => entry.username === target || entry.uuid === target)
      : target ?? this.player();
    if (!item || !player) return { ok: false, message: "That slot is empty" };
    const amount = Math.max(1, Math.min(item.count, count));
    if (item.count > amount) {
      this.setItemAt(player, slot, { ...item, count: item.count - amount });
    } else {
      this.setItemAt(player, slot, null);
    }
    const stack = Tag.compound({
      id: Tag.string(item.id),
      count: Tag.int(amount),
      ...(item.components && Object.keys(item.components).length > 0
        ? { components: Tag.compound({ ...item.components }) }
        : {}),
    });
    const entity = createEntity("minecraft:item", entityPos(player), Tag.compound({
      Item: stack,
    }));
    this.addEntity(entity);
    this.touch();
    return { ok: true, message: `Dropped ${amount} ${item.id.replace("minecraft:", "")}` };
  }

  deleteItem(target: EntityData | string | null, slot: string): boolean {
    const item = this.getItemAt(target, slot);
    if (!item) return false;
    this.setItemAt(target, slot, null);
    return true;
  }

  equipItem(target: EntityData | string | null, slot: string): { ok: boolean; message: string } {
    const item = this.getItemAt(target, slot);
    if (!item) return { ok: false, message: "That slot is empty" };
    const armorSlot = item.id.endsWith("helmet") ? "armor.head"
      : item.id.endsWith("chestplate") ? "armor.chest"
        : item.id.endsWith("leggings") ? "armor.legs"
          : item.id.endsWith("boots") ? "armor.feet"
            : null;
    if (!armorSlot) return { ok: false, message: "This item cannot be equipped" };
    this.setItemAt(target, slot, null);
    this.setItemAt(target, armorSlot, item);
    return { ok: true, message: `Equipped ${item.id.replace("minecraft:", "")}` };
  }

  /** Mirror the carried inventory, offhand, armor and ender chest into player NBT. */
  syncInventoryNbt(target: EntityData | string | null): void {
    const entity = typeof target === "string"
      ? this.entities.find((entry) => entry.username === target || entry.uuid === target)
      : target ?? this.player();
    if (!entity || !entity.isPlayer) return;
    const state = this.playerState(entity);
    if (!state) return;
    const slotNumber = (slot?: string): number | null => {
      if (!slot) return null;
      if (slot === "weapon.offhand") return -106;
      if (slot === "armor.head") return 103;
      if (slot === "armor.chest") return 102;
      if (slot === "armor.legs") return 101;
      if (slot === "armor.feet") return 100;
      const hotbar = /^hotbar\.(\d+)$/.exec(slot);
      if (hotbar) return Number(hotbar[1]);
      const inventory = /^inventory\.(\d+)$/.exec(slot);
      if (inventory) return 9 + Number(inventory[1]);
      return null;
    };
    const stackFields = (item: InventoryItem) => ({
      id: Tag.string(item.id),
      count: Tag.int(item.count),
      ...(item.components && Object.keys(item.components).length > 0
        ? { components: Tag.compound({ ...item.components }) }
        : {}),
    });
    const toStack = (item: InventoryItem, slot: number) =>
      Tag.compound({
        Slot: Tag.byte(slot),
        ...stackFields(item),
      });
    const inventory = state.inventory
      .map((item, index) => {
        const slot = slotNumber(item.slot) ?? index;
        return toStack(item, slot);
      });
    const ender = state.enderchest.map((item, index) =>
      toStack(item, Number(item.slot?.split(".")[1] ?? index)),
    );
    entity.nbt.value.Inventory = Tag.list(inventory, "compound");
    entity.nbt.value.EnderItems = Tag.list(ender, "compound");
    entity.nbt.value.SelectedItemSlot = Tag.int(state.selectedSlot);
    const selectedItem = state.inventory.find(
      (item) => item.slot === `hotbar.${state.selectedSlot}`,
    );
    if (selectedItem) {
      entity.nbt.value.SelectedItem = Tag.compound(stackFields(selectedItem));
    } else {
      delete entity.nbt.value.SelectedItem;
    }
  }

  /** Re-read inventory/ender chest after raw NBT edits such as /data modify. */
  importInventoryNbt(
    target: EntityData | string | null,
    prefer: "inventory" | "selected" = "selected",
  ): void {
    const entity = typeof target === "string"
      ? this.entities.find((entry) => entry.username === target || entry.uuid === target)
      : target ?? this.player();
    if (!entity?.isPlayer) return;
    const state = this.ensurePlayerState(entity);
    const slotName = (slot: number): string | undefined => {
      if (slot === -106) return "weapon.offhand";
      if (slot === 103) return "armor.head";
      if (slot === 102) return "armor.chest";
      if (slot === 101) return "armor.legs";
      if (slot === 100) return "armor.feet";
      if (slot >= 0 && slot <= 8) return `hotbar.${slot}`;
      if (slot >= 9 && slot <= 35) return `inventory.${slot - 9}`;
      return undefined;
    };
    const readList = (tag: NbtTag | undefined, ender: boolean): InventoryItem[] => {
      if (!isList(tag)) return [];
      const items: InventoryItem[] = [];
      for (const entry of tag.value) {
        if (!isCompound(entry)) continue;
        const idTag = entry.value.id;
        if (!isString(idTag)) continue;
        const countTag = entry.value.count ?? entry.value.Count;
        const count = numericValue(countTag) ?? 1;
        const slotTag = entry.value.Slot;
        const slotNumberValue = numericValue(slotTag) ?? 0;
        const components = isCompound(entry.value.components)
          ? { ...entry.value.components.value }
          : undefined;
        items.push({
          id: idTag.value,
          count,
          ...(components ? { components } : {}),
          slot: ender
            ? `enderchest.${slotNumberValue}`
            : slotName(slotNumberValue),
        });
      }
      return items;
    };
    const inventory = readList(entity.nbt.value.Inventory, false);
    if (inventory.length > 0 || isList(entity.nbt.value.Inventory)) {
      state.inventory = inventory;
    }
    const ender = readList(entity.nbt.value.EnderItems, true);
    if (ender.length > 0 || isList(entity.nbt.value.EnderItems)) {
      state.enderchest = ender;
    }
    const selected = entity.nbt.value.SelectedItemSlot;
    if (selected && typeof selected.value === "number") {
      state.selectedSlot = Math.max(0, Math.min(8, selected.value));
    }
    const selectedItemTag = entity.nbt.value.SelectedItem;
    if (prefer === "selected" && isCompound(selectedItemTag)) {
      const idTag = selectedItemTag.value.id;
      if (isString(idTag)) {
        const count = numericValue(selectedItemTag.value.count ?? selectedItemTag.value.Count) ?? 1;
        const components = isCompound(selectedItemTag.value.components)
          ? { ...selectedItemTag.value.components.value }
          : undefined;
        const slot = `hotbar.${state.selectedSlot}`;
        state.inventory = state.inventory.filter((item) => item.slot !== slot);
        state.inventory.push({
          id: idTag.value,
          count,
          ...(components ? { components } : {}),
          slot,
        });
      }
    }
  }

  /* ------------------------------ schedule ---------------------------- */

  /**
   * Current `/schedule function` queue, ordered by due tick then sequence.
   * Read-only: mutate through `scheduleFunction` / `clearSchedules` so the
   * sequence counter and ordering stay consistent.
   */
  get scheduled(): readonly ScheduledFunction[] {
    return this.scheduleQueue;
  }

  /**
   * Queue a function. `replace` (the default) drops any existing entries for
   * the same function first; `append` keeps them and adds another entry.
   */
  scheduleFunction(
    functionId: string,
    delayTicks: number,
    mode: "replace" | "append" = "replace",
  ): ScheduledFunction {
    if (mode === "replace") this.clearSchedules(functionId);
    // The command parser only accepts integers, but keep the state honest if
    // another caller passes a fractional or negative delay.
    const delay = Math.max(0, Math.trunc(delayTicks));
    const entry: ScheduledFunction = {
      functionId,
      scheduledAtTick: this.tick,
      dueTick: this.tick + delay,
      sequence: ++this.scheduleSequence,
    };
    this.scheduleQueue.push(entry);
    this.scheduleQueue.sort(
      (a, b) => a.dueTick - b.dueTick || a.sequence - b.sequence,
    );
    return entry;
  }

  /** Remove every entry for `functionId`, or the whole queue when omitted. */
  clearSchedules(functionId?: string): number {
    if (functionId === undefined) {
      const removed = this.scheduleQueue.length;
      this.scheduleQueue = [];
      return removed;
    }
    const before = this.scheduleQueue.length;
    this.scheduleQueue = this.scheduleQueue.filter(
      (entry) => entry.functionId !== functionId,
    );
    return before - this.scheduleQueue.length;
  }

  /**
   * Atomically take every entry due at or before `currentTick`. The tick
   * runtime calls this once per tick, so functions queued by the tick tag are
   * handled by the same pass while functions queued by a due function wait for
   * the next tick.
   */
  collectDueSchedules(currentTick: number = this.tick): ScheduledFunction[] {
    if (this.scheduleQueue.length === 0) return [];
    const due: ScheduledFunction[] = [];
    const remaining: ScheduledFunction[] = [];
    for (const entry of this.scheduleQueue) {
      if (entry.dueTick <= currentTick) due.push(entry);
      else remaining.push(entry);
    }
    if (due.length === 0) return [];
    this.scheduleQueue = remaining;
    return due;
  }

  /* -------------------------------- logs ------------------------------ */

  log(kind: LogKind, text: string, component?: NbtTag): void {
    this.logs.push({
      id: ++logId,
      kind,
      text,
      component,
      time: Date.now(),
    });
    if (this.logs.length > 600) this.logs.splice(0, this.logs.length - 600);
  }

  chat(component: NbtTag, fallback: string): void {
    const text = this.plain(component) || fallback;
    this.logs.push({
      id: ++logId,
      kind: "chat",
      text,
      component,
      time: Date.now(),
    });
    if (this.logs.length > 600) this.logs.splice(0, this.logs.length - 600);
  }

  /* ------------------------------- text ------------------------------- */

  plain(component: NbtTag, source: EntityData | null = null): string {
    return plainText(component, this.textContext(source));
  }

  textContext(source: EntityData | null, origin?: Vec3): TextRenderContext {
    const pos = origin ?? (source ? entityPos(source) : { x: 0, y: 64, z: 0 });
    const normalizeId = (id: string) => (id.includes(":") ? id : `minecraft:${id}`);
    return {
      resolveScoreHolder: (name) => {
        if (name === "@s" && source) return entityHolderName(source);
        try {
          return resolveSingleHolder(this, name, source, pos);
        } catch {
          return name;
        }
      },
      resolveScore: (holder, objective) => this.getScore(holder, objective) ?? null,
      resolveSelectorComponents: (selector) => {
        try {
          const parsed = parseSelector(selector);
          return selectEntities(this, parsed, {
            source,
            origin: pos,
            lenient: true,
          }).map((entity) => {
            if (entity.isPlayer && entity.username) {
              return Tag.compound({ text: Tag.string(entity.username) });
            }
            const custom = entity.nbt.value.CustomName;
            return custom ?? Tag.string(entityCustomName(entity));
          });
        } catch {
          return [];
        }
      },
      resolveSelector: (selector) => {
        try {
          const parsed = parseSelector(selector);
          return selectEntities(this, parsed, {
            source,
            origin: pos,
            lenient: true,
          }).map(entityCustomName);
        } catch {
          return [];
        }
      },
      resolveNbtTags: (sourceType, id, path) => {
        try {
          const segments = path ? parseNbtPath(path) : [];
          if (sourceType === "storage") {
            const storage = this.storages.get(normalizeId(id));
            return storage ? getPathValues(storage, segments) : [];
          }
          if (sourceType === "entity") {
            const target = this.entities.find(
              (entity) => entity.username === id || entity.uuid === id,
            );
            return target ? getPathValues(target.nbt, segments) : [];
          }
          return [];
        } catch {
          return [];
        }
      },
      resolveNbtTag: (sourceType, id, path) => {
        try {
          const segments = path ? parseNbtPath(path) : [];
          if (sourceType === "storage") {
            const storage = this.storages.get(normalizeId(id));
            if (!storage) return null;
            return getPathValues(storage, segments)[0] ?? null;
          }
          if (sourceType === "entity") {
            const target = this.entities.find(
              (entity) => entity.username === id || entity.uuid === id,
            );
            return target ? getPathValues(target.nbt, segments)[0] ?? null : null;
          }
          return null;
        } catch {
          return null;
        }
      },
      resolveEntityNbt: (sourceType, path) => {
        try {
          const parts = path.split(/\s+/);
          if (sourceType === "storage" && parts.length >= 1) {
            const storage = this.storages.get(parts[0]);
            if (!storage) return null;
            const path = parts.slice(1).join(" ");
            const values = getPathValues(storage, parseNbtPath(path));
            return values.map(formatTagForComponent).join(", ");
          }
          return null;
        } catch {
          return null;
        }
      },
    };
  }

  /* ------------------------------- misc ------------------------------- */

  cloneScoresFrom(other: World): void {
    for (const [name, objective] of other.objectives) {
      this.objectives.set(name, { ...objective, displayName: objective.displayName });
      this.scores.set(name, new Map(other.scores.get(name) ?? []));
    }
    for (const [slot, objective] of other.displaySlots) {
      this.displaySlots.set(slot, objective);
    }
    for (const [holder, name] of other.holderNames) {
      this.holderNames.set(holder, name);
    }
  }

  describeVersion(): string {
    return this.version.id;
  }
}

export function makeWorld(versionId: string, playerPos?: Vec3): World {
  return new World({ version: getVersion(versionId), playerPos });
}

export function readNumber(tag: NbtTag | undefined, fallback = 0): number {
  return numericValue(tag) ?? fallback;
}

function normalizePredicateId(id: string): string {
  return id.includes(":") ? id : `minecraft:${id}`;
}

export { entityPos, parseSnbt, SelectorError };
