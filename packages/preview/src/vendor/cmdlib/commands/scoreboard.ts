import { Cursor } from "../cursor";
import { type CompoundTag, type NbtTag, Tag } from "../nbt";
import { DISPLAY_SLOTS, SCOREBOARD_CRITERIA } from "../data";
import { SCOREBOARD_MAX, SCOREBOARD_MIN } from "../world";
import {
  resolveHolderTargets,
  resolveSingleHolder,
  SelectorError,
} from "../selectors";
import type { HolderTarget } from "../selectors";
import { type CommandContext, type CommandResult, feedback, print } from "./context";
import { CommandError } from "../cursor";

type Handler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

function readScoreValue(cursor: Cursor): number {
  const value = cursor.readInt();
  if (value.value < SCOREBOARD_MIN || value.value > SCOREBOARD_MAX) {
    cursor.error(
      `Score must be between ${SCOREBOARD_MIN} and ${SCOREBOARD_MAX}`,
      value.start,
      "score-range",
    );
  }
  return value.value;
}

export const scoreboardHandler: Handler = (cursor, ctx) => {
  const group = cursor.readWord();
  if (group === "objectives") return objectives(cursor, ctx);
  if (group === "players") return players(cursor, ctx);
  cursor.error(`Expected 'objectives' or 'players', got '${group}'`, cursor.mark() - group.length);
  return { result: 0, success: false };
};

function objectives(cursor: Cursor, ctx: CommandContext): CommandResult {
  const start = cursor.mark();
  const action = cursor.readWord();
  switch (action) {
    case "add": {
      const name = readObjective(cursor);
      const criterion = cursor.readWord();
      const known = SCOREBOARD_CRITERIA.some((c) => c.name === criterion) || criterion.includes(":");
      if (!known) {
        // Minecraft rejects unknown criteria; being strict here catches typos
        // that would otherwise silently create a dead objective.
        cursor.error(`Unknown criterion '${criterion}'`, cursor.mark() - criterion.length, "criterion");
      }
      let displayName: NbtTag = Tag.compound({
        text: Tag.string(name),
      });
      if (!cursor.done()) {
        displayName = cursor.readComponent().tag;
      }
      if (ctx.lint) {
        // The lint world preloads every declared objective so cross-file
        // references resolve. Re-declaration is therefore expected.
        return { result: 1, success: true };
      }
      try {
        ctx.world.addObjective(name, criterion, displayName);
      } catch (error) {
        throw new CommandError(String((error as Error).message), start, "objective");
      }
      feedback(ctx, "success", `Added new objective '${name}' using '${criterion}' criterion`);
      return { result: 1, success: true };
    }
    case "remove": {
      const name = readObjective(cursor);
      if (!ctx.world.objectives.has(name)) {
        if (!ctx.lint) {
          print(
            ctx,
            "warning",
            `Objective '${name}' does not exist; nothing to remove`,
          );
        }
        return { result: 0, success: false };
      }
      if (!ctx.lint) ctx.world.removeObjective(name);
      feedback(ctx, "success", `Removed objective '${name}'`);
      return { result: 1, success: true };
    }
    case "list": {
      const names = [...ctx.world.objectives.keys()];
      if (names.length === 0) print(ctx, "info", "There are no objectives.");
      else {
        for (const name of names) {
          const objective = ctx.world.objectives.get(name)!;
          print(
            ctx,
            "info",
            `${name}: ${ctx.world.plain(objective.displayName) || name} (${objective.criterion})`,
          );
        }
      }
      return { result: names.length, success: true };
    }
    case "modify": {
      const name = readObjective(cursor);
      const objective = ctx.world.objectives.get(name);
      if (!objective) {
        cursor.error(`Unknown objective '${name}'`, cursor.mark() - name.length, "unknown-objective");
      }
      const property = cursor.readWord();
      if (property === "displayname") {
        const component = cursor.readComponent();
        if (objective && !ctx.lint) objective.displayName = component.tag;
        feedback(ctx, "success", `Set the display name of objective '${name}'`);
        return { result: 1, success: true };
      }
      if (property === "rendertype" || property === "numberformat") {
        if (property === "rendertype") {
          const type = cursor.readWord();
          if (type !== "integer" && type !== "hearts") {
            cursor.error(`Expected 'integer' or 'hearts', got '${type}'`, cursor.mark() - type.length);
          }
          if (objective && !ctx.lint) objective.renderType = type;
          feedback(ctx, "success", `Set the render type of objective '${name}' to '${type}'`);
        } else {
          // Could be a literal, a styled component, or `fixed`.
          if (!cursor.done()) cursor.readComponent();
          feedback(ctx, "success", `Updated the number format of objective '${name}'`);
        }
        return { result: 1, success: true };
      }
      cursor.error(`Unknown objective property '${property}'`, cursor.mark() - property.length);
      break;
    }
    case "setdisplay": {
      const slot = cursor.readWord();
      const valid = DISPLAY_SLOTS.some((s) => s.name === slot);
      if (!valid) {
        cursor.error(
          `Unknown display slot '${slot}'. Try sidebar, list, below_name, or sidebar.team.<colour>.`,
          cursor.mark() - slot.length,
          "display-slot",
        );
      }
      let objectiveName: string | null = null;
      if (!cursor.done()) {
        objectiveName = readObjective(cursor);
        if (!ctx.world.objectives.has(objectiveName)) {
          cursor.error(
            `Unknown objective '${objectiveName}'`,
            cursor.mark() - objectiveName.length,
            "unknown-objective",
          );
        }
      }
      if (!ctx.lint) ctx.world.setDisplaySlot(slot, objectiveName);
      feedback(
        ctx,
        "success",
        objectiveName
          ? `Displayed objective '${objectiveName}' in slot '${slot}'`
          : `Cleared display slot '${slot}'`,
      );
      return { result: 1, success: true };
    }
    default:
      cursor.error(`Unknown objectives action '${action}'`, cursor.mark() - action.length);
  }
  void start;
  return { result: 0, success: false };
}

function players(cursor: Cursor, ctx: CommandContext): CommandResult {
  const action = cursor.readWord();
  if (action === "display") return display(cursor, ctx);
  if (action === "enable") return enable(cursor, ctx);
  if (!["set", "add", "remove", "reset", "get", "operation"].includes(action)) {
    cursor.error(`Unknown players action '${action}'`, cursor.mark() - action.length);
  }

  const targetsStart = cursor.mark();
  const targetsRaw = cursor.readTargetsRaw().raw;
  const isStar = targetsRaw === "*";
  if (isStar && action !== "reset") {
    cursor.error("'*' is only valid for scoreboard players reset", targetsStart);
  }

  let objectiveName: string | null = null;
  if (!cursor.done()) {
    const objectiveStart = cursor.mark();
    objectiveName = readObjective(cursor);
    if (!ctx.world.objectives.has(objectiveName)) {
      cursor.error(
        `Unknown objective '${objectiveName}'`,
        objectiveStart,
        "unknown-objective",
      );
    }
  } else if (action !== "reset") {
    cursor.error(
      `Expected an objective after the target list`,
      cursor.mark(),
      "objective-required",
    );
  }

  let holders: HolderTarget[];
  try {
    if (isStar) {
      const tracked = objectiveName
        ? [...(ctx.world.scores.get(objectiveName)?.keys() ?? [])]
        : [...new Set([...ctx.world.scores.values()].flatMap((m) => [...m.keys()]))];
      holders = tracked.map((holder) => ({ holder }));
    } else {
      holders = resolveHolderTargets(ctx.world, targetsRaw, {
        source: ctx.source,
        origin: ctx.pos,
        lenient: ctx.lint,
      });
    }
  } catch (error) {
    if (error instanceof SelectorError) {
      cursor.error(error.message, targetsStart, "selector");
    }
    throw error;
  }

  switch (action) {
    case "set": {
      const value = { value: readScoreValue(cursor) };
      for (const target of holders) {
        if (!ctx.lint) ctx.world.setScore(target.holder, objectiveName!, value.value);
      }
      feedback(
        ctx,
        "success",
        holders.length === 1
          ? `Set score of ${ctx.world.holderDisplayName(holders[0].holder)} to ${value.value} in objective '${objectiveName}'`
          : `Set ${holders.length} scores to ${value.value} in objective '${objectiveName}'`,
      );
      return { result: holders.length, success: true };
    }
    case "add":
    case "remove": {
      const value = { value: readScoreValue(cursor) };
      const delta = action === "add" ? value.value : -value.value;
      for (const target of holders) {
        if (!ctx.lint) ctx.world.addScore(target.holder, objectiveName!, delta);
      }
      feedback(
        ctx,
        "success",
        holders.length === 1
          ? `${action === "add" ? "Added" : "Removed"} ${Math.abs(value.value)} ${action === "add" ? "to" : "from"} ${ctx.world.holderDisplayName(holders[0].holder)}'s score`
          : `${action === "add" ? "Added" : "Removed"} ${Math.abs(value.value)} across ${holders.length} holders`,
      );
      return { result: holders.length, success: true };
    }
    case "reset": {
      for (const target of holders) {
        if (ctx.lint) continue;
        if (cursor.done() || !objectiveName) ctx.world.resetAllScores(target.holder);
        else ctx.world.resetScore(target.holder, objectiveName!);
      }
      feedback(ctx, "success", `Reset ${holders.length} score(s)`);
      return { result: holders.length, success: true };
    }
    case "get": {
      let last = 0;
      for (const target of holders) {
        const value = ctx.world.getScore(target.holder, objectiveName!);
        if (value === undefined) {
          if (!ctx.lint) {
            print(
              ctx,
              "error",
              `Could not get value of ${target.holder} for objective '${objectiveName}' (score not set)`,
            );
          }
          continue;
        }
        last = value;
        print(
          ctx,
          "data",
          `${ctx.world.holderDisplayName(target.holder)} has ${value} [${objectiveName}]`,
        );
      }
      return { result: last, success: holders.length > 0 };
    }
    case "operation": {
      const opStart = cursor.mark();
      const operation = cursor.readWord();
      if (!["=", "+=", "-=", "*=", "/=", "%=", "<", ">", "><"].includes(operation)) {
        cursor.error(`Unknown operation '${operation}'`, opStart, "operation");
      }
      const sourceStart = cursor.mark();
      const sourceRaw = cursor.readWord();
      let sourceHolder: string;
      try {
        sourceHolder = resolveSingleHolder(
          ctx.world,
          sourceRaw,
          ctx.source,
          ctx.pos,
        );
      } catch (error) {
        if (error instanceof SelectorError) {
          cursor.error(error.message, sourceStart, "selector");
        }
        throw error;
      }
      const sourceObjectiveStart = cursor.mark();
      const sourceObjective = readObjective(cursor);
      if (!ctx.world.objectives.has(sourceObjective)) {
        cursor.error(
          `Unknown objective '${sourceObjective}'`,
          sourceObjectiveStart,
          "unknown-objective",
        );
      }
      let sourceValue = ctx.world.getScore(sourceHolder, sourceObjective) ?? 0;
      for (const target of holders) {
        if (!ctx.lint) {
          const current = ctx.world.getScore(target.holder, objectiveName!) ?? 0;
          let next = current;
          switch (operation) {
            case "=":
              next = sourceValue;
              break;
            case "+=":
              next = current + sourceValue;
              break;
            case "-=":
              next = current - sourceValue;
              break;
            case "*=":
              next = current * sourceValue;
              break;
            case "/=":
              if (sourceValue === 0) {
                throw new CommandError("Cannot divide by zero", opStart, "division-by-zero");
              }
              next = Math.trunc(current / sourceValue);
              break;
            case "%=":
              if (sourceValue === 0) {
                throw new CommandError("Cannot divide by zero", opStart, "division-by-zero");
              }
              next = current % sourceValue;
              break;
            case "<":
              next = Math.min(current, sourceValue);
              break;
            case ">":
              next = Math.max(current, sourceValue);
              break;
            case "><":
              next = sourceValue;
              sourceValue = current;
              break;
          }
          ctx.world.setScore(target.holder, objectiveName!, next);
        }
      }
      feedback(
        ctx,
        "success",
        `Applied operation '${operation}' to ${holders.length} score(s) in objective '${objectiveName}'`,
      );
      return { result: holders.length, success: true };
    }
  }
  return { result: 0, success: false };
}

function enable(cursor: Cursor, ctx: CommandContext): CommandResult {
  const targetsStart = cursor.mark();
  const targetsRaw = cursor.readTargetsRaw().raw;
  const objectiveStart = cursor.mark();
  const objectiveName = readObjective(cursor);
  const objective = ctx.world.objectives.get(objectiveName);
  if (!objective) {
    cursor.error(
      `Unknown objective '${objectiveName}'`,
      objectiveStart,
      "unknown-objective",
    );
  }
  if (objective.criterion !== "trigger") {
    cursor.error(
      `Objective '${objectiveName}' uses the '${objective.criterion}' criterion, not 'trigger'`,
      objectiveStart,
      "trigger-objective",
    );
  }

  let holders: HolderTarget[];
  try {
    holders = resolveHolderTargets(ctx.world, targetsRaw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
  } catch (error) {
    if (error instanceof SelectorError) {
      cursor.error(error.message, targetsStart, "selector");
    }
    throw error;
  }
  const enabledHolders: string[] = [];
  for (const target of holders) {
    if (ctx.world.isTriggerEnabled(target.holder, objectiveName)) continue;
    enabledHolders.push(target.holder);
    if (!ctx.lint) {
      ctx.world.setTriggerEnabled(target.holder, objectiveName, true);
    }
  }
  if (enabledHolders.length === 0) {
    return { result: 0, success: false };
  }
  feedback(
    ctx,
    "success",
    enabledHolders.length === 1
      ? `Enabled trigger '${objectiveName}' for ${enabledHolders[0]}`
      : `Enabled trigger '${objectiveName}' for ${enabledHolders.length} holders`,
  );
  return { result: enabledHolders.length, success: true };
}

function display(cursor: Cursor, ctx: CommandContext): CommandResult {
  const kind = cursor.readWord();
  if (kind === "name") {
    const targetsRaw = cursor.readTargetsRaw().raw;
    const objectiveName = readObjective(cursor);
    const component = cursor.readComponent();
    const holders = resolveHolderTargets(ctx.world, targetsRaw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
    for (const target of holders) {
      if (!ctx.lint) {
        ctx.world.setHolderDisplayName(
          target.holder,
          ctx.world.plain(component.tag) || target.holder,
        );
      }
    }
    feedback(ctx, "success", `Updated the display name in objective '${objectiveName}'`);
    return { result: holders.length, success: true };
  }
  if (kind === "numberformat") {
    // Consume the rest of the format so the line still parses.
    if (!cursor.done()) cursor.readComponent();
    feedback(ctx, "success", "Updated the number format");
    return { result: 1, success: true };
  }
  cursor.error(`Unknown display action '${kind}'`, cursor.mark() - kind.length);
}

function readObjective(cursor: Cursor): string {
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
