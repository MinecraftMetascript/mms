import { Cursor } from "../cursor";
import { type CompoundTag, isCompound, type NbtTag, stringifySnbt, Tag } from "../nbt";
import { getPathValues, parseNbtPath } from "../nbtPath";
import { resolveEntityTargets } from "../selectors";
import { type CommandContext, type CommandResult, feedback } from "./context";
import { CommandError } from "../cursor";

type Handler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

export const functionHandler: Handler = (cursor, ctx) => {
  const idStart = cursor.mark();
  const id = cursor.readResourceLocation(true);

  let args: CompoundTag | null = null;
  let argsLabel = "";
  let fromWith = false;
  if (!cursor.done()) {
    if (cursor.peekWord() === "with") {
      cursor.readWord();
      args = readWithSource(cursor, ctx);
      fromWith = true;
      argsLabel = " from with";
    } else if (cursor.peek() === "{") {
      const parsed = cursor.readSnbt();
      if (!isCompound(parsed.tag)) {
        cursor.error("Macro arguments must be a compound", parsed.start, "macro-args");
      }
      args = parsed.tag as CompoundTag;
      argsLabel = " with inline arguments";
    } else {
      cursor.error(
        `Expected macro arguments as {...} or 'with', got '${cursor.peekWord()}'`,
        cursor.mark(),
        "macro-args",
      );
    }
  }

  const tag = id.startsWith("#");
  if (tag) {
    const name = id.slice(1);
    const functionTag = ctx.world.getFunctionTag(name);
    if (!functionTag) {
      cursor.error(`Unknown function tag '${name}'`, idStart, "unknown-function-tag");
    }
    if (ctx.lint) {
      return { result: 1, success: true };
    }
    return ctx.runtime.runFunction(id, args ? args.value : null, ctx);
  }

  const def = ctx.world.getFunction(id);
  if (!def) {
    cursor.error(`Unknown function '${id}'`, idStart, "unknown-function");
    return { result: 0, success: false };
  }

  if (ctx.lint) {
    if (def.hasMacroLines) {
      if (args && fromWith) {
        // A `with` source may be built at runtime (or be empty in the lint
        // world). Fill placeholders so the call site can still be validated.
        for (const key of def.macroKeys) {
          if (!(key in args.value)) args.value[key] = Tag.int(1);
        }
      }
      if (!args) {
        cursor.error(
          `Function '${id}' uses macros and must be called with arguments`,
          idStart,
          "macro-required",
        );
      } else {
        const missing = def.macroKeys.filter((key) => !(key in args!.value));
        if (missing.length > 0) {
          cursor.error(
            `Missing macro argument${missing.length === 1 ? "" : "s"} ${missing
              .map((m) => `'${m}'`)
              .join(", ")} for function '${id}'`,
            idStart,
            "macro-argument",
          );
        }
      }
    }
    return { result: 1, success: true };
  }

  if (ctx.depth > 64) {
    cursor.error(
      `Function recursion limit reached while calling '${id}'`,
      idStart,
      "recursion",
    );
  }

  feedback(ctx, "info", `Ran function ${id}${argsLabel}`);
  return ctx.runtime.runFunction(id, args ? args.value : null, ctx);
};

function readWithSource(cursor: Cursor, ctx: CommandContext): CompoundTag {
  const start = cursor.mark();
  const kind = cursor.readWord();
  if (kind === "storage") {
    const idStart = cursor.mark();
    const id = cursor.readResourceLocation();
    const storage = ctx.world.storage(id);
    if (!storage) {
      if (ctx.lint) return Tag.compound();
      throw new CommandError(`No storage named '${id}' exists`, idStart, "storage-not-found");
    }
    const pathStart = cursor.mark();
    const pathText = cursor.done() ? "" : cursor.readNbtPathText();
    const path = pathText ? parsePath(cursor, pathText, pathStart) : [];
    const values = getPathValues(storage, path);
    if (values.length !== 1 || !isCompound(values[0])) {
      if (ctx.lint) return Tag.compound();
      throw new CommandError(
        pathText
          ? `The path '${pathText}' must point at exactly one compound`
          : "The storage must contain a compound",
        pathStart,
        "macro-source",
      );
    }
    return values[0];
  }
  if (kind === "entity") {
    const raw = cursor.readTargetsRaw().raw;
    const pathStart = cursor.mark();
    const pathText = cursor.done() ? "" : cursor.readNbtPathText();
    const path = pathText ? parsePath(cursor, pathText, pathStart) : [];
    const entities = resolveEntityTargets(ctx.world, raw, {
      source: ctx.source,
      origin: ctx.pos,
      lenient: ctx.lint,
    });
    if (entities.length !== 1) {
      if (ctx.lint) return Tag.compound();
      throw new CommandError(
        `The entity selector must match exactly one entity`,
        cursor.mark(),
        "macro-source",
      );
    }
    const values = getPathValues(entities[0].nbt, path);
    if (values.length !== 1 || !isCompound(values[0])) {
      if (ctx.lint) return Tag.compound();
      throw new CommandError(
        pathText
          ? `The path '${pathText}' must point at exactly one compound`
          : "The entity NBT must contain a compound",
        pathStart,
        "macro-source",
      );
    }
    return values[0];
  }
  if (kind === "block") {
    const pos = cursor.readVec3(ctx.pos, ctx.rotation);
    const pathStart = cursor.mark();
    const pathText = cursor.done() ? "" : cursor.readNbtPathText();
    const path = pathText ? parsePath(cursor, pathText, pathStart) : [];
    const values = getPathValues(ctx.world.getBlockNbt(pos), path);
    if (values.length !== 1 || !isCompound(values[0])) {
      if (ctx.lint) return Tag.compound();
      throw new CommandError(
        pathText
          ? `The path '${pathText}' must point at exactly one compound`
          : "The block NBT must contain a compound",
        pathStart,
        "macro-source",
      );
    }
    return values[0];
  }
  cursor.error(
    `Expected 'storage', 'entity' or 'block', got '${kind}'`,
    start,
    "macro-source",
  );
}

function parsePath(cursor: Cursor, text: string, start: number) {
  try {
    return parseNbtPath(text);
  } catch (error) {
    const pathError = error as { message: string; index: number };
    cursor.error(`Invalid path '${text}': ${pathError.message}`, start + (pathError.index ?? 0), "path");
  }
}

export function macroValueToString(tag: NbtTag): string {
  switch (tag.type) {
    case "string":
      return tag.value;
    case "byte":
    case "short":
    case "int":
    case "long":
      return String(tag.value);
    case "float":
    case "double":
      return String(tag.value);
    case "compound":
    case "list":
    case "byteArray":
    case "intArray":
    case "longArray":
      return stringifySnbt(tag, { typed: false, quoteKeys: true });
  }
}
