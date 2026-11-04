import { World } from "./world";
import type { ScheduledFunction } from "./world";
import type { EntityData } from "./entities";
import { isCompound, isList, isNumeric, type NbtTag } from "./nbt";
import { CommandError, Cursor } from "./cursor";
import { executeCommandLine, makeContext, ReturnSignal } from "./commands/context";
import type { CommandContext, CommandResult, Runtime } from "./commands/context";
import { parseDatapack, type Project } from "./datapack";
import { parsePredicateFiles } from "./predicates";
import { parseAdvancementFiles } from "./advancements";
import { parseItemModifierFiles } from "./item_modifiers";
import { getVersion, type VersionProfile } from "./version";
import { macroValueToString } from "./commands/function";
import { HANDLERS } from "./commands";
import type { GameMode, InventoryItem, StatusEffect } from "./player";

export interface RunOutcome extends CommandResult {
  ok: boolean;
  error?: string;
}

/** Result of one `Simulator.tick()`, deliberately small and serialisable. */
export interface TickOutcome {
  /** World tick after the increment (starts at 1). */
  tick: number;
  /** True when a `minecraft:tick` function tag existed and was invoked. */
  tickTagRan: boolean;
  /** Number of scheduled functions that became due and were run. */
  scheduledRan: number;
  /** Errors raised by the tick tag or due scheduled functions. */
  errors: string[];
  ok: boolean;
}

export interface SimulatorOptions {
  project: Project;
  /** Seed a player at the origin. */
  player?: boolean;
}

/**
 * Owns one live world and the datapack functions loaded into it.
 *
 * The simulator is deliberately not a game loop. Entities do not have AI and
 * nothing ticks unless a command runs. That constraint is what makes the
 * preview deterministic, and it is why `cmdlib` can run entirely in a tab.
 */
export class Simulator implements Runtime {
  project: Project;
  version: VersionProfile;
  /** 0 when joined as a normal player; 2 for the editor/operator session. */
  permissionLevel = 2;
  world: World;
  revision = 0;
  private listeners = new Set<() => void>();
  private detachWorld: (() => void) | null = null;
  private loading = false;
  private tagStack = new Set<string>();

  constructor(options: SimulatorOptions) {
    this.project = options.project;
    this.version = getVersion(options.project.versionId);
    this.world = new World({
      version: this.version,
      seedPlayer: options.player !== false,
    });
    this.attachWorld();
    this.syncFunctions();
  }

  private attachWorld(): void {
    this.detachWorld?.();
    this.detachWorld = this.world.subscribe(() => this.touch());
  }

  subscribe = (listener: () => void): (() => void) => {
    this.listeners.add(listener);
    return () => this.listeners.delete(listener);
  };

  getSnapshot = (): number => this.revision;

  touch(): void {
    this.revision++;
    for (const listener of this.listeners) listener();
  }

  dispose(): void {
    this.detachWorld?.();
    this.listeners.clear();
  }

  setProject(project: Project): void {
    const versionChanged = project.versionId !== this.project.versionId;
    this.project = project;
    if (versionChanged) {
      this.version = getVersion(project.versionId);
      this.reset({ runLoad: false });
      return;
    }
    this.syncFunctions();
    this.touch();
  }

  /** Other packs in the same project whose functions, tags and predicates are visible. */
  contextProjects: Project[] = [];

  /** Re-read functions, tags and predicates without destroying world state. */
  syncFunctions(): void {
    const functions = new Map<string, ReturnType<typeof parseDatapack>["functions"][number]>();
    const tags = new Map<string, ReturnType<typeof parseDatapack>["tags"][number]>();
    const predicates = new Map<string, ReturnType<typeof parsePredicateFiles>["predicates"] extends Map<string, infer T> ? T : never>();
    const sources = [
      this.project,
      ...this.contextProjects.filter((pack) => pack.id !== this.project.id),
    ];
    const advancements = new Map<string, ReturnType<typeof parseAdvancementFiles>["advancements"] extends Map<string, infer T> ? T : never>();
    for (const source of sources) {
      const version = getVersion(source.versionId);
      const parsed = parseDatapack(source.files, version);
      for (const fn of parsed.functions) functions.set(fn.id, fn);
      for (const tag of parsed.tags) tags.set(tag.id, tag);
      for (const [id, predicate] of parsePredicateFiles(source.files).predicates) {
        predicates.set(id, predicate);
      }
      for (const [id, advancement] of parseAdvancementFiles(source.files).advancements) {
        advancements.set(id, advancement);
      }
    }
    const modifiers = new Map<string, ReturnType<typeof parseItemModifierFiles>["modifiers"] extends Map<string, infer T> ? T : never>();
    for (const source of sources) {
      for (const [id, modifier] of parseItemModifierFiles(source.files).modifiers) {
        modifiers.set(id, modifier);
      }
    }
    this.world.setFunctions([...functions.values()]);
    this.world.setFunctionTags([...tags.values()]);
    this.world.setPredicates(predicates);
    this.world.setAdvancements(advancements);
    this.world.setItemModifiers(modifiers);
  }

  /** Build a new world, load the pack, and run its load tags. */
  reset(options: { runLoad?: boolean } = {}): void {
    this.detachWorld?.();
    this.world = new World({
      version: this.version,
      seedPlayer: true,
    });
    this.attachWorld();
    this.syncFunctions();
    this.world.log(
      "system",
      `Simulation reset · Minecraft ${this.version.id} · pack format ${this.version.packFormat.join(".")}`,
    );
    if (options.runLoad !== false) {
      this.loading = true;
      try {
        this.runLoadTags();
      } finally {
        this.loading = false;
      }
    }
    this.world.touch();
  }

  runLoadTags(): number {
    let ran = 0;
    const tags = ["minecraft:load", "cmdlib:load"];
    for (const tag of tags) {
      if (!this.world.getFunctionTag(tag)) continue;
      try {
        this.runFunction(`#${tag}`, null, this.baseContext());
        ran++;
      } catch (error) {
        this.world.log("error", this.formatError(error));
      }
    }
    if (ran === 0) {
      this.world.log(
        "info",
        "No load tag found. Add data/minecraft/tags/function/load.json to run setup automatically.",
      );
    }
    return ran;
  }

  /* ------------------------------- ticking ---------------------------- */

  /** Read-only view of the world's `/schedule function` queue. */
  get scheduled(): readonly ScheduledFunction[] {
    return this.world.scheduled;
  }

  /* ------------------------- player state views ----------------------- */

  /** Game mode of the default player. `survival` until `/gamemode` runs. */
  get playerGameMode(): GameMode {
    return this.world.playerGameMode();
  }

  /** Read-only inventory of the default player. */
  get playerInventory(): readonly InventoryItem[] {
    return this.world.playerInventory();
  }

  /** Read-only active effects on the default player. */
  get playerEffects(): readonly StatusEffect[] {
    return this.world.effectsOf();
  }

  /** Remove queued schedules for one function, or all of them. */
  clearSchedules(functionId?: string): number {
    const removed = this.world.clearSchedules(functionId);
    if (removed > 0) this.world.touch();
    return removed;
  }

  /**
   * Advance the simulation by exactly one game tick.
   *
   * Order matches the runtime contract: bump `world.tick`, run every value of
   * the `minecraft:tick` function tag, then run the functions that were due at
   * the start of that pass. The due list is collected *after* the tag so a tick
   * function may safely call `/schedule` (including `0t`); entries added by a
   * due function itself wait for the next tick, which prevents a same-tick
   * reschedule loop from stalling the UI.
   *
   * Errors are logged and returned rather than thrown; the simulation always
   * ends with one `touch()` so React subscribers see the new tick.
   */
  tick(): TickOutcome {
    this.world.tick += 1;
    const tick = this.world.tick;
    const errors: string[] = [];
    let tickTagRan = false;
    this.world.resetCommandBudget();
    // Effects tick before the tick function so commands observe the state the
    // player will have during this game tick. Expiry is silent but the final
    // `touch()` below still notifies UI subscribers.
    this.world.tickStatusEffects();

    if (this.world.getFunctionTag("minecraft:tick")) {
      tickTagRan = true;
      try {
        // `runFunction` delegates tags to the existing cycle-safe runner.
        this.runFunction("#minecraft:tick", null, this.baseContext());
      } catch (error) {
        const message = this.formatError(error);
        errors.push(message);
        this.world.log("error", message);
      }
    }

    const due = this.world.collectDueSchedules(tick);
    for (const entry of due) {
      try {
        this.runFunction(entry.functionId, null, this.baseContext());
      } catch (error) {
        const message = this.formatError(error);
        errors.push(message);
        this.world.log(
          "error",
          `[scheduled ${entry.functionId}] ${message}`,
        );
      }
    }

    this.evaluateAdvancements("tick");
    this.evaluateAdvancements("inventory_changed");
    this.world.touch();
    return {
      tick,
      tickTagRan,
      scheduledRan: due.length,
      errors,
      ok: errors.length === 0,
    };
  }

  grantAdvancement(holder: string, id: string): boolean {
    const definition = this.world.getAdvancement(id);
    if (!definition) return false;
    if (!this.world.grantAdvancement(holder, id)) return false;
    if (definition.rewards.function) {
      const player =
        this.world.entities.find((entity) => entity.username === holder) ?? null;
      const ctx = {
        ...this.baseContext(true),
        source: player,
      };
      try {
        this.runFunction(definition.rewards.function, null, ctx);
      } catch (error) {
        this.world.log(
          "error",
          `Advancement reward '${definition.rewards.function}' failed: ${this.formatError(error)}`,
        );
      }
    }
    return true;
  }

  revokeAdvancement(holder: string, id: string): boolean {
    return this.world.revokeAdvancement(holder, id);
  }

  private evaluateAdvancements(trigger: string): void {
    for (const player of this.world.entities) {
      if (!player.isPlayer || !player.username) continue;
      const completed = this.world.advancementProgressFor(player.username);
      for (const definition of this.world.advancements.values()) {
        if (completed.has(definition.id)) continue;
        const relevant = definition.criteria.filter(
          (criterion) => criterion.trigger === trigger,
        );
        if (relevant.length === 0) continue;
        if (relevant.length !== definition.criteria.length) continue;
        if (!relevant.every((criterion) => this.advancementCriterionMet(criterion, player))) {
          continue;
        }
        this.grantAdvancement(player.username, definition.id);
      }
    }
  }

  private advancementCriterionMet(
    criterion: { trigger: string; conditions: Record<string, unknown> | null },
    player: EntityData,
  ): boolean {
    if (criterion.trigger === "tick") return true;
    if (criterion.trigger !== "inventory_changed") return false;
    const items = criterion.conditions?.items;
    if (!Array.isArray(items)) return false;
    const inventory = this.world.playerInventory(player);
    return items.some((entry) => {
      if (!entry || typeof entry !== "object") return false;
      const spec = entry as Record<string, unknown>;
      const ids = typeof spec.items === "string"
        ? [spec.items]
        : Array.isArray(spec.items)
          ? spec.items.filter((id): id is string => typeof id === "string")
          : [];
      return inventory.some((item) => {
        if (ids.length > 0) {
          const normalized = item.id.includes(":") ? item.id : `minecraft:${item.id}`;
          if (!ids.some((id) => id === normalized || id === item.id)) return false;
        }
        const components = spec.components;
        if (components && typeof components === "object" && !Array.isArray(components)) {
          for (const [rawKey, expected] of Object.entries(components as Record<string, unknown>)) {
            const key = rawKey.includes(":") ? rawKey : `minecraft:${rawKey}`;
            const actual = item.components?.[key];
            if (actual === undefined) return false;
            if (!this.componentMatches(actual, expected)) return false;
          }
        }
        return true;
      });
    });
  }

  private componentMatches(tag: NbtTag, expected: unknown): boolean {
    if (typeof expected === "boolean") {
      return isNumeric(tag) && Number(tag.value) === (expected ? 1 : 0);
    }
    if (typeof expected === "number") {
      return isNumeric(tag) && Number(tag.value) === expected;
    }
    if (typeof expected === "string") {
      return tag.type === "string" && tag.value === expected;
    }
    if (Array.isArray(expected)) {
      return (
        isList(tag) &&
        tag.value.length === expected.length &&
        tag.value.every((item, index) => this.componentMatches(item, expected[index]))
      );
    }
    if (expected && typeof expected === "object") {
      if (!isCompound(tag)) return false;
      return Object.entries(expected as Record<string, unknown>).every(
        ([key, value]) =>
          tag.value[key] !== undefined && this.componentMatches(tag.value[key], value),
      );
    }
    return false;
  }

  private baseContext(quiet = false): CommandContext {
    const player = this.world.player() ?? null;
    return makeContext(this.world, this, player, { quiet, lint: false });
  }

  private isNonFatal(error: unknown): error is CommandError {
    return (
      error instanceof CommandError &&
      (error.code === "unsupported" || error.code === "unsupported-command")
    );
  }

  /** Run one console line. Errors are logged, not thrown. */
  runCommand(line: string, options: { quiet?: boolean } = {}): RunOutcome {
    const commandName = line.trim().replace(/^\//, "").split(/\s+/, 1)[0]?.toLowerCase() ?? "";
    const playerCommands = new Set(["help", "list", "me", "trigger"]);
    if (this.permissionLevel === 0 && !playerCommands.has(commandName)) {
      const message = `You do not have permission to use /${commandName || "that command"} (permission level 0).`;
      this.world.log("error", message);
      this.world.touch();
      return { result: 0, success: false, ok: false, error: message };
    }
    this.world.resetCommandBudget();
    const ctx = this.baseContext(options.quiet ?? false);
    this.currentLine = line;
    try {
      const result = executeCommandLine(line, ctx);
      this.world.touch();
      return { ...result, ok: true };
    } catch (error) {
      if (error instanceof ReturnSignal) {
        this.world.touch();
        return { result: error.value, success: true, ok: true };
      }
      const message = this.formatError(error);
      this.world.log("error", message);
      this.world.touch();
      return { result: 0, success: false, ok: false, error: message };
    } finally {
      this.currentLine = undefined;
    }
  }

  /** Run a script of commands in one context (a function file, essentially). */
  runScript(source: string, options: { quiet?: boolean; macros?: Record<string, NbtTag> } = {}): RunOutcome {
    this.world.resetCommandBudget();
    const ctx = this.baseContext(options.quiet ?? false);
    ctx.macroArgs = options.macros;
    const lines = source.split(/\r?\n/);
    let last: { result: number; success: boolean } = { result: 0, success: true };
    try {
      for (let i = 0; i < lines.length; i++) {
        const prepared = this.prepareLine(lines[i], options.macros ?? null, "(console)", i + 1);
        if (prepared === null) continue;
        try {
          const result = executeCommandLine(prepared, ctx);
          last = result;
        } catch (error) {
          if (error instanceof ReturnSignal) throw error;
          if (this.isNonFatal(error)) {
            this.world.log(
              "warning",
              `Warning: ${error.message} (console line ${i + 1}); continuing`,
            );
            continue;
          }
          throw error;
        }
      }
      this.world.touch();
      return { ...last, ok: true };
    } catch (error) {
      if (error instanceof ReturnSignal) {
        this.world.touch();
        return { result: error.value, success: true, ok: true };
      }
      const message = this.formatError(error);
      this.world.log("error", message);
      this.world.touch();
      return { result: 0, success: false, ok: false, error: message };
    }
  }

  runFile(path: string): RunOutcome {
    const source = this.project.files[path];
    if (source === undefined) {
      const message = `File '${path}' does not exist in this project`;
      this.world.log("error", message);
      this.touch();
      return { result: 0, success: false, ok: false, error: message };
    }
    return this.runScript(source);
  }


  /** Runtime interface used by the /function command. */
  runFunction(
    id: string,
    args: Record<string, NbtTag> | null,
    parent: CommandContext,
  ): CommandResult {
    if (id.startsWith("#")) {
      return this.runFunctionTag(id.slice(1), args, parent);
    }
    const def = this.world.getFunction(id);
    if (!def) {
      throw new CommandError(`Unknown function '${id}'`, 0, "unknown-function");
    }
    if (def.hasMacroLines && (!args || Object.keys(args).length === 0)) {
      throw new CommandError(
        `Function '${id}' uses macros and must be called with arguments`,
        0,
        "macro-required",
      );
    }
    if (parent.depth > 64) {
      throw new CommandError(
        `Function recursion limit reached while calling '${id}'`,
        0,
        "recursion",
      );
    }
    const ctx: CommandContext = {
      ...parent,
      depth: parent.depth + 1,
      quiet: true,
      functionId: id,
      macroArgs: args ?? undefined,
    };
    let lastResult = 0;
    for (let i = 0; i < def.lines.length; i++) {
      const prepared = this.prepareLine(def.lines[i], args, id, i + 1);
      if (prepared === null) continue;
      try {
        const result = executeCommandLine(prepared, ctx);
        lastResult = result.result;
      } catch (error) {
        if (error instanceof ReturnSignal) {
          return { result: error.value, success: true };
        }
        if (this.isNonFatal(error)) {
          // Unsupported commands are lint warnings, not runtime failures. Keep
          // executing the rest of the function so a single unsupported command
          // cannot stop a setup/tick chain.
          this.world.log(
            "warning",
            `Warning: ${error.message} (function ${id}, line ${i + 1}); continuing`,
          );
          continue;
        }
        if (error instanceof CommandError) {
          // Cycles and runaway recursion are structural failures, not ordinary
          // command failures, so surface them to the tick/run result.
          if (error.code === "tag-cycle" || error.code === "recursion") {
            error.message = `${error.message} (function ${id}, line ${i + 1})`;
            throw error;
          }
          // A failed command in a function does not abort the remaining
          // lines in vanilla. Function command failures are also silent in
          // game chat; explicit diagnostics such as missing objectives are
          // emitted by the command itself. Keep going without noisy errors.
          continue;
        }
        throw error;
      }
    }
    return { result: lastResult, success: true };
  }

  private runFunctionTag(
    id: string,
    args: Record<string, NbtTag> | null,
    parent: CommandContext,
  ): CommandResult {
    const tag = this.world.getFunctionTag(id);
    if (!tag) {
      throw new CommandError(`Unknown function tag '${id}'`, 0, "unknown-function-tag");
    }
    if (this.tagStack.has(id)) {
      throw new CommandError(
        `Function tag cycle detected at '${id}'`,
        0,
        "tag-cycle",
      );
    }
    this.tagStack.add(id);
    try {
      let result = 0;
      let success = true;
      for (const value of tag.values) {
        try {
          const branch = this.runFunction(value, args, parent);
          result += branch.result;
          success = success && branch.success;
        } catch (error) {
          if (error instanceof ReturnSignal) continue;
          throw error;
        }
      }
      return { result, success };
    } finally {
      this.tagStack.delete(id);
    }
  }

  /**
   * Expand a function line. Lines that begin with `$` are macro lines; `$$(`
   * escapes a literal `$(`. Values are stringified the way vanilla does.
   */
  prepareLine(
    raw: string,
    args: Record<string, NbtTag> | null,
    sourceLabel: string,
    lineNumber: number,
  ): string | null {
    const line = raw.replace(/\s+$/, "");
    if (line.trim().length === 0) return null;
    if (line.trimStart().startsWith("#")) return null;
    if (!line.startsWith("$")) return line;
    const body = line.slice(1);
    if (!args) {
      throw new CommandError(
        `Macro function requires arguments (${sourceLabel}, line ${lineNumber})`,
        0,
        "macro-required",
      );
    }
    let out = "";
    let i = 0;
    while (i < body.length) {
      if (body.startsWith("$$(", i)) {
        out += "$(";
        i += 3;
        continue;
      }
      if (body.startsWith("$(", i)) {
        const close = body.indexOf(")", i + 2);
        if (close === -1) {
          throw new CommandError(
            `Unclosed macro interpolation at line ${lineNumber}`,
            i,
            "macro-syntax",
          );
        }
        const key = body.slice(i + 2, close);
        const value = args[key];
        if (!value) {
          throw new CommandError(
            `Missing macro argument '${key}' (${sourceLabel}, line ${lineNumber})`,
            i,
            "macro-argument",
          );
        }
        const text = macroValueToString(value);
        if (text.includes("\n") || text.includes("\r")) {
          throw new CommandError(
            `Macro argument '${key}' contains a newline, which Minecraft forbids`,
            i,
            "macro-newline",
          );
        }
        out += text;
        i = close + 1;
        continue;
      }
      out += body[i];
      i++;
    }
    return out;
  }

  formatError(error: unknown): string {
    if (error instanceof CommandError) {
      const line = this.currentLine ?? "";
      return line
        ? `${error.message}  ·  ${line.trim()}`
        : error.message;
    }
    if (error instanceof ReturnSignal) return `return ${error.value}`;
    return (error as Error)?.message ?? String(error);
  }

  private currentLine?: string;

  /** Used by the console to check a command without committing it. */
  previewCommand(line: string): { ok: boolean; offset?: number; message?: string } {
    const ctx = this.baseContext(true);
    ctx.lint = true;
    try {
      const cursor = new Cursor(line.replace(/^\s*\//, ""));
      const name = cursor.readWord();
      if (!name) return { ok: true };
      if (!HANDLERS[name]) {
        return { ok: false, message: `Unknown command '${name}'`, offset: 0 };
      }
      return { ok: true };
    } catch (error) {
      return {
        ok: false,
        message: (error as Error).message,
        offset: (error as { index?: number }).index,
      };
    }
  }
}
