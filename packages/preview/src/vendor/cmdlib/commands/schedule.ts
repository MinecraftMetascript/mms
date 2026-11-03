import { Cursor } from "../cursor";
import { type CommandContext, type CommandResult, feedback } from "./context";

type Handler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

export const TICKS_PER_SECOND = 20;
export const TICKS_PER_DAY = 24000;

export const SCHEDULE_USAGE =
  "/schedule function <function> <time> [append|replace] | /schedule clear <function>";

const TIME_UNITS: Record<string, number> = {
  t: 1,
  s: TICKS_PER_SECOND,
  d: TICKS_PER_DAY,
};

/**
 * Parse a `/schedule` time token.
 *
 * The grammar mirrors vanilla's `time` argument: an unsigned integer followed
 * by `t` (ticks), `s` (seconds) or `d` (days). No decimals, signs or
 * whitespace are accepted. Returns `null` when the token is invalid.
 */
export function parseScheduleTime(raw: string): number | null {
  const match = /^(\d+)([tsd])$/.exec(raw);
  if (!match) return null;
  const amount = Number(match[1]);
  const unit = TIME_UNITS[match[2]];
  if (!Number.isSafeInteger(amount) || unit === undefined) return null;
  const ticks = amount * unit;
  return Number.isSafeInteger(ticks) ? ticks : null;
}

export const scheduleHandler: Handler = (cursor, ctx) => {
  const actionStart = cursor.mark();
  const action = cursor.peekWord();

  if (action === "clear") {
    cursor.readWord();
    const idStart = cursor.mark();
    const id = cursor.readResourceLocation();
    if (!ctx.world.getFunction(id)) {
      cursor.error(`Unknown function '${id}'`, idStart, "unknown-function");
    }
    if (!cursor.done()) cursor.error("Unexpected arguments", cursor.mark());
    if (ctx.lint) return { result: 1, success: true };
    const removed = ctx.world.clearSchedules(id);
    if (removed > 0) {
      feedback(
        ctx,
        "success",
        `Cleared ${removed} scheduled ${removed === 1 ? "call" : "calls"} for ${id}`,
      );
    } else {
      feedback(ctx, "info", `No scheduled functions for ${id}`);
    }
    return { result: removed, success: true };
  }

  if (action !== "function") {
    cursor.error(
      `Expected 'function' or 'clear', got '${action || "end of line"}'. Usage: ${SCHEDULE_USAGE}`,
      actionStart,
      "schedule-action",
    );
  }

  cursor.readWord();
  const idStart = cursor.mark();
  const id = cursor.readResourceLocation();
  if (!ctx.world.getFunction(id)) {
    cursor.error(`Unknown function '${id}'`, idStart, "unknown-function");
  }

  const timeStart = cursor.mark();
  const rawTime = cursor.readWord();
  const delay = parseScheduleTime(rawTime);
  if (delay === null) {
    return cursor.error(
      `Invalid time '${rawTime}'. Expected an integer followed by t, s or d (for example 20t, 3s, 1d).`,
      timeStart,
      "schedule-time",
    );
  }

  let mode: "append" | "replace" = "replace";
  if (!cursor.done()) {
    const modeStart = cursor.mark();
    const rawMode = cursor.readWord();
    if (rawMode !== "append" && rawMode !== "replace") {
      cursor.error(
        `Expected 'append' or 'replace', got '${rawMode}'`,
        modeStart,
        "schedule-mode",
      );
    }
    mode = rawMode === "append" ? "append" : "replace";
  }

  if (!cursor.done()) cursor.error("Unexpected arguments", cursor.mark());
  if (ctx.lint) return { result: 1, success: true };

  const entry = ctx.world.scheduleFunction(id, delay, mode);
  feedback(
    ctx,
    "success",
    `Scheduled function ${id} in ${delay} tick${delay === 1 ? "" : "s"} (tick ${entry.dueTick})${mode === "append" ? " [append]" : ""}`,
  );
  return { result: 1, success: true };
};
