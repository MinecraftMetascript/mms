import { Cursor } from "../cursor";
import { type NbtTag } from "../nbt";
import { resolveEntityTargets, SelectorError } from "../selectors";
import { type CommandContext, type CommandResult, feedback } from "./context";

type Handler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

export const titleHandler: Handler = (cursor, ctx) => {
  const targetsStart = cursor.mark();
  const raw = cursor.readTargetsRaw().raw;
  try {
    const entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
    if (entities.length === 0 && !ctx.lint) {
      cursor.error("No player was found for /title", targetsStart, "selector");
    }
  } catch (error) {
    if (error instanceof SelectorError) {
      cursor.error(error.message, targetsStart, "selector");
    }
    throw error;
  }

  const actionStart = cursor.mark();
  const action = cursor.readWord();
  if (action === "times") {
    const fadeIn = cursor.readInt();
    const stay = cursor.readInt();
    const fadeOut = cursor.readInt();
    for (const value of [fadeIn, stay, fadeOut]) {
      if (value.value < 0) {
        cursor.error("Title timings must be zero or greater", value.start, "title-times");
      }
    }
    if (!ctx.lint) {
      ctx.world.setTitleTimes(fadeIn.value, stay.value, fadeOut.value);
    }
    feedback(
      ctx,
      "success",
      `Set title times to ${fadeIn.value} / ${stay.value} / ${fadeOut.value} ticks`,
    );
    return { result: 1, success: true };
  }
  if (action === "reset" || action === "clear") {
    if (!ctx.lint) {
      ctx.world.clearTitle(action === "reset" ? "title" : "all");
      if (action === "reset") ctx.world.clearTitle("subtitle");
    }
    feedback(ctx, "success", action === "reset" ? "Reset titles" : "Cleared titles");
    return { result: 1, success: true };
  }
  if (!["title", "subtitle", "actionbar"].includes(action)) {
    cursor.error(
      `Expected title, subtitle, actionbar, times, reset or clear, got '${action}'`,
      actionStart,
      "title-action",
    );
  }
  if (cursor.done()) {
    cursor.error(`Expected a text component after '${action}'`, cursor.mark(), "title-component");
  }
  const component: NbtTag = cursor.readComponent().tag;
  if (!ctx.lint) {
    ctx.world.setTitle(action as "title" | "subtitle" | "actionbar", component);
  }
  const preview = ctx.world.plain(component, ctx.source);
  feedback(
    ctx,
    "success",
    `Displayed ${action}${preview ? `: ${preview}` : ""}`,
  );
  return { result: 1, success: true };
};
