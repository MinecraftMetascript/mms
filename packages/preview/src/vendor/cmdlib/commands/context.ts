import { World } from "../world";
import { entityPos, entityRotation } from "../entities";
import type { EntityData, Vec2, Vec3 } from "../entities";
import type { NbtTag } from "../nbt";
import { Cursor, CommandError, UnsupportedCommandError } from "../cursor";
import { COMMANDS } from "../data";
import { HANDLERS } from "./index";

export interface CommandResult {
  result: number;
  success: boolean;
}

export interface Runtime {
  world: World;
  runFunction(
    id: string,
    args: Record<string, NbtTag> | null,
    ctx: CommandContext,
  ): CommandResult;
  /** Advance the simulation by one tick (used by /tick step and sprint). */
  tick(): void;
  grantAdvancement(holder: string, id: string): boolean;
  revokeAdvancement(holder: string, id: string): boolean;
}

export interface CommandContext {
  world: World;
  runtime: Runtime;
  source: EntityData | null;
  pos: Vec3;
  rotation: Vec2;
  /** Position anchor set by `execute anchored`; affects later position clauses. */
  anchor: "feet" | "eyes";
  dimension: string;
  /** Inside a function: suppress routine feedback, like vanilla does. */
  quiet: boolean;
  lint: boolean;
  depth: number;
  macroArgs?: Record<string, NbtTag>;
  functionId?: string;
  lineNumber?: number;
}

export class ReturnSignal extends Error {
  value: number;
  constructor(value: number) {
    super("return");
    this.name = "ReturnSignal";
    this.value = value;
  }
}

export interface ParseContext {
  ctx: CommandContext;
  cursor: Cursor;
}

export function makeContext(
  world: World,
  runtime: Runtime,
  source: EntityData | null,
  options: { quiet?: boolean; lint?: boolean; depth?: number } = {},
): CommandContext {
  const pos: Vec3 = source ? entityPos(source) : { x: 0, y: 64, z: 0 };
  const rotation: Vec2 = source ? entityRotation(source) : { yaw: 0, pitch: 0 };
  return {
    world,
    runtime,
    source,
    pos,
    rotation,
    anchor: "feet",
    dimension: "minecraft:overworld",
    quiet: options.quiet ?? false,
    lint: options.lint ?? false,
    depth: options.depth ?? 0,
  };
}

export function forkContext(
  ctx: CommandContext,
  patch: Partial<CommandContext>,
): CommandContext {
  return { ...ctx, ...patch };
}

export function feedback(
  ctx: CommandContext,
  kind: Parameters<World["log"]>[0],
  text: string,
  component?: NbtTag,
): void {
  if (ctx.lint || ctx.quiet) return;
  ctx.world.log(kind, text, component);
}

export function print(ctx: CommandContext, kind: Parameters<World["log"]>[0], text: string): void {
  if (ctx.lint) return;
  ctx.world.log(kind, text);
}

export function executeCommandLine(
  line: string,
  ctx: CommandContext,
): CommandResult {
  const text = line.replace(/^\s*\//, "").trim();
  if (text.length === 0 || text.startsWith("#")) {
    return { result: 0, success: true };
  }
  ctx.world.consumeCommand();
  if (ctx.world.revision < 0) {
    // never true; keeps the method honest about world ownership
  }
  const cursor = new Cursor(text);
  const start = cursor.mark();
  const name = cursor.readWord().toLowerCase();
  const handler = HANDLERS[name];
  if (!handler) {
    const spec = COMMANDS[name];
    if (spec && !spec.supported) {
      throw new UnsupportedCommandError(name, start);
    }
    throw new CommandError(
      `Unknown command '${name}'. Try /scoreboard, /data, /function, /title, /execute, /summon or /say.`,
      start,
      "unknown-command",
    );
  }
  const result = handler(cursor, ctx);
  cursor.skip();
  if (!cursor.done()) {
    cursor.error(
      `Unexpected trailing input '${cursor.readRest()}'`,
      cursor.mark(),
      "trailing",
    );
  }
  return result;
}

/* ----------------------------- shared readers ----------------------------- */

export function readTargets(
  cursor: Cursor,
  ctx: CommandContext,
): { raw: string; start: number; end: number } {
  return cursor.readTargetsRaw();
}

export function readObjectiveName(cursor: Cursor): string {
  const start = cursor.mark();
  const name = cursor.readWord();
  if (!/^[A-Za-z0-9_.+-]+$/.test(name)) {
    cursor.error(
      `Invalid objective name '${name}'. Use letters, digits, '_', '.', '+' or '-'.`,
      start,
      "objective-name",
    );
  }
  return name;
}

export function readScoreHolderName(cursor: Cursor): string {
  return cursor.readWord();
}

export { Cursor, CommandError, UnsupportedCommandError };
