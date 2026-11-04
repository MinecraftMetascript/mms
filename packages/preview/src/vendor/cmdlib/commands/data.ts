import { Cursor } from "../cursor";
import { type CompoundTag, isCompound, isList, isString, type NbtTag, stringifySnbt, Tag } from "../nbt";
import {
  applyPathModify,
  getPathValues,
  mergeCompound,
  NbtPathError,
  parseNbtPath,
  resolvePath,
} from "../nbtPath";
import { resolveEntityTargets, SelectorError } from "../selectors";
import type { EntityData, Vec3 } from "../entities";
import { entityCustomName } from "../entities";
import { type CommandContext, type CommandResult, feedback, print } from "./context";
import { CommandError } from "../cursor";

type Handler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

interface DataSource {
  kind: "block" | "entity" | "storage";
  label: string;
  root: CompoundTag;
  entity?: EntityData;
  id?: string;
  pos?: Vec3;
}

export const dataHandler: Handler = (cursor, ctx) => {
  const actionStart = cursor.mark();
  const action = cursor.readWord();
  switch (action) {
    case "get":
      return get(cursor, ctx);
    case "merge":
      return merge(cursor, ctx);
    case "modify":
      return modify(cursor, ctx);
    case "remove":
      return remove(cursor, ctx);
    default:
      cursor.error(`Unknown data action '${action}'`, actionStart);
      return { result: 0, success: false };
  }
};

function get(cursor: Cursor, ctx: CommandContext): CommandResult {
  const source = readSource(cursor, ctx, false);
  let pathText: string | null = null;
  let scale: number | null = null;

  if (!cursor.done()) {
    const firstStart = cursor.mark();
    const first =
      cursor.peek() === "{" ? cursor.readNbtPathText() : cursor.readWord();
    if (/^[+-]?(?:\d+\.?\d*|\.\d+)$/.test(first)) {
      scale = Number(first);
    } else {
      pathText = first;
      if (!cursor.done()) {
        const scaleStart = cursor.mark();
        const raw = cursor.readWord();
        if (!/^[+-]?(?:\d+\.?\d*|\.\d+)$/.test(raw)) {
          cursor.error(`Expected a numeric scale, got '${raw}'`, scaleStart);
        }
        scale = Number(raw);
      }
    }
    void firstStart;
  }

  let values: NbtTag[];
  if (pathText === null) {
    values = [source.root];
  } else {
    const path = parsePathSafe(cursor, pathText, 0);
    const matches = resolvePath(source.root, path);
    if (matches.length === 0) {
      if (!ctx.lint) {
        throw new CommandError(
          `No tag was found at path '${pathText}' in ${source.label}`,
          cursor.mark(),
          "path-not-found",
        );
      }
      values = [];
    } else {
      values = matches.map((m) => m.value);
    }
  }

  if (scale !== null) {
    values = values.map((value) => {
      if (isNumericTag(value)) {
        return Tag.int(Math.floor(value.value * scale!));
      }
      if (!ctx.lint) {
        throw new CommandError(
          `Cannot scale a ${value.type} tag; only numeric tags can be scaled`,
          cursor.mark(),
          "scale-type",
        );
      }
      return value;
    });
  }

  const tagResult = (tag: NbtTag): number => {
    if (isNumericTag(tag)) return tag.value;
    if (isString(tag)) return tag.value.length;
    if (isList(tag)) return tag.value.length;
    if (tag.type === "byteArray" || tag.type === "intArray" || tag.type === "longArray") {
      return tag.value.length;
    }
    if (isCompound(tag)) return Object.keys(tag.value).length;
    return 1;
  };
  const result =
    values.length === 1 ? tagResult(values[0]) : values.length;

  const label =
    source.kind === "storage"
      ? `Storage ${source.id}`
      : source.kind === "block"
        ? `Block at ${source.pos?.x} ${source.pos?.y} ${source.pos?.z}`
        : `Entity ${source.entity ? entityCustomName(source.entity) : "?"}`;

  if (ctx.lint) return { result, success: true };

  if (values.length === 0) {
    print(ctx, "info", `${label} has no data at that path`);
    return { result: 0, success: true };
  }
  const noun =
    source.kind === "storage" ? "contents" : source.kind === "block" ? "block data" : "entity data";
  const pathSuffix = pathText ? ` at path '${pathText}'` : "";
  if (values.length === 1) {
    print(
      ctx,
      "data",
      `${label} has the following ${noun}${pathSuffix}: ${stringifySnbt(values[0])}`,
    );
  } else {
    print(
      ctx,
      "data",
      `${label} has ${values.length} matching tags${pathSuffix}: ${values
        .map((v) => stringifySnbt(v))
        .join(", ")}`,
    );
  }
  return { result, success: true };
}

function merge(cursor: Cursor, ctx: CommandContext): CommandResult {
  const source = readSource(cursor, ctx, true);
  const nbt = cursor.readSnbt().tag;
  if (!cursor.done()) {
    cursor.error(`Unexpected trailing input '${cursor.readRest()}'`, cursor.mark(), "trailing");
  }
  if (!isCompound(nbt)) {
    cursor.error("data merge requires a compound tag", cursor.mark(), "nbt-type");
  }
  if (!isCompound(source.root)) {
    cursor.error("The target does not hold a compound", cursor.mark(), "nbt-type");
  }
  try {
    mergeCompound(source.root, nbt);
    if (source.kind === "entity" && source.entity) {
      ctx.world.importInventoryNbt(source.entity);
    }
  } catch (error) {
    throw new CommandError(String((error as Error).message), cursor.mark(), "merge");
  }
  feedback(ctx, "success", `Merged data into ${source.label}`);
  return { result: 1, success: true };
}

function modify(cursor: Cursor, ctx: CommandContext): CommandResult {
  const source = readSource(cursor, ctx, true);
  const pathStart = cursor.mark();
  const pathText = cursor.readNbtPathText();
  const path = parsePathSafe(cursor, pathText, pathStart);

  const modeStart = cursor.mark();
  const mode = cursor.readWord();
  let insertIndex = 0;
  let value: NbtTag | null = null;

  if (mode === "insert") {
    insertIndex = cursor.readInt().value;
  } else if (!["set", "merge", "append", "prepend", "remove"].includes(mode)) {
    cursor.error(
      `Unknown modify mode '${mode}'. Expected set, merge, append, prepend, insert or remove.`,
      modeStart,
      "modify-mode",
    );
  }

  if (mode !== "remove") {
    const valueKindStart = cursor.mark();
    const valueKind = cursor.readWord();
    if (valueKind === "value") {
      value = cursor.readSnbt().tag;
      if (!cursor.done()) {
        cursor.error(
          `Unexpected trailing input '${cursor.readRest()}'`,
          cursor.mark(),
          "trailing",
        );
      }
    } else if (valueKind === "from") {
      const fromSource = readSource(cursor, ctx, false);
      // The source path is optional in vanilla: with no path the entire
      // source payload is copied. This is what makes
      // `/data modify storage su:temp Entity set from entity @s` work.
      if (cursor.done()) {
        value = clone(fromSource.root);
      } else {
        const fromPathStart = cursor.mark();
        const fromPathText = cursor.readNbtPathText();
        const fromPath = parsePathSafe(cursor, fromPathText, fromPathStart);
        const matches = resolvePath(fromSource.root, fromPath);
        if (matches.length === 0) {
          if (!ctx.lint) {
            throw new CommandError(
              `No tag was found at path '${fromPathText}'`,
              fromPathStart,
              "path-count",
            );
          }
          value = Tag.compound();
        } else if (matches.length > 1) {
          if (ctx.lint || (mode !== "append" && mode !== "prepend" && mode !== "insert")) {
            if (!ctx.lint) {
              throw new CommandError(
                `The source path matched ${matches.length} tags; it must match exactly one`,
                fromPathStart,
                "path-count",
              );
            }
            value = Tag.compound();
          } else {
            value = Tag.list(
              matches.map((match) => clone(match.value)),
              "end",
            );
          }
        } else {
          value = clone(matches[0].value);
        }
      }
    } else if (valueKind === "string") {
      const fromSource = readSource(cursor, ctx, false);
      const fromPathStart = cursor.mark();
      const fromPathText = cursor.readNbtPathText();
      const fromPath = parsePathSafe(cursor, fromPathText, fromPathStart);
      const matches = resolvePath(fromSource.root, fromPath);
      const start = cursor.readInt().value;
      let end: number | null = null;
      if (!cursor.done()) {
        const save = cursor.mark();
        try {
          end = cursor.readInt().value;
        } catch {
          cursor.pos = save;
          end = null;
        }
      }
      const sourceTag = matches.length === 1 ? matches[0].value : null;
      if (
        !sourceTag ||
        (!isString(sourceTag) && !isNumericTag(sourceTag))
      ) {
        if (!ctx.lint) {
          throw new CommandError(
            `The source path must match exactly one string or numeric tag`,
            fromPathStart,
            "path-type",
          );
        }
        value = Tag.string("");
      } else {
        // Vanilla converts numeric tags to their SNBT form before slicing.
        const raw = isString(sourceTag)
          ? sourceTag.value
          : stringifySnbt(sourceTag);
        const at = (index: number) =>
          index < 0
            ? Math.max(0, raw.length + index)
            : Math.min(raw.length, index);
        value = Tag.string(raw.slice(at(start), end === null ? raw.length : at(end)));
      }
    } else if (valueKind === "compute") {
      cursor.error(
        "/data modify … compute arrived in 26.3 but cmdlib does not evaluate number providers yet.",
        valueKindStart,
        "unsupported",
      );
    } else {
      cursor.error(
        `Expected 'value', 'from' or 'string' as the modify source, got '${valueKind}'`,
        valueKindStart,
        "modify-source",
      );
    }
  }

  try {
    applyPathModify(source.root, path, mode as "set" | "merge" | "append" | "prepend" | "insert" | "remove", value, insertIndex);
    if (source.kind === "entity" && source.entity) {
      ctx.world.importInventoryNbt(
        source.entity,
        pathText.startsWith("Inventory") || pathText.startsWith("EnderItems")
          ? "inventory"
          : "selected",
      );
    }
  } catch (error) {
    // During lint the pack's setup functions have not run, so entity/storage
    // state may be absent. Do not turn that context-dependent type mismatch
    // into a false-positive error.
    if (ctx.lint) return { result: 0, success: true };
    throw new CommandError(String((error as Error).message), pathStart, "modify");
  }
  feedback(ctx, "success", `Modified ${source.label}: ${pathText} ${mode}`);
  return { result: 1, success: true };
}

function remove(cursor: Cursor, ctx: CommandContext): CommandResult {
  const source = readSource(cursor, ctx, true);
  const pathStart = cursor.mark();
  const pathText = cursor.readNbtPathText();
  const path = parsePathSafe(cursor, pathText, pathStart);
  try {
    applyPathModify(source.root, path, "remove", null);
    if (source.kind === "entity" && source.entity) {
      ctx.world.importInventoryNbt(
        source.entity,
        pathText.startsWith("Inventory") || pathText.startsWith("EnderItems")
          ? "inventory"
          : "selected",
      );
    }
  } catch (error) {
    if (ctx.lint) return { result: 0, success: true };
    throw new CommandError(String((error as Error).message), pathStart, "remove");
  }
  feedback(ctx, "success", `Removed tag ${pathText} from ${source.label}`);
  return { result: 1, success: true };
}

/* ------------------------------ helpers ------------------------------ */

function readSource(
  cursor: Cursor,
  ctx: CommandContext,
  create: boolean,
): DataSource {
  const start = cursor.mark();
  const kind = cursor.readWord();
  if (kind === "block") {
    const pos = cursor.readVec3(ctx.pos, ctx.rotation);
    const block = ctx.world.getBlock(pos);
    if (!block && !ctx.lint) {
      throw new CommandError(
        `No block entity data exists at ${pos.x} ${pos.y} ${pos.z}`,
        start,
        "block-not-found",
      );
    }
    return {
      kind: "block",
      label: `block at ${pos.x}, ${pos.y}, ${pos.z}`,
      root: ctx.world.getBlockNbt(pos),
      pos,
    };
  }
  if (kind === "entity") {
    const targetsStart = cursor.mark();
    const raw = cursor.readTargetsRaw().raw;
    let entities;
    try {
      entities = resolveEntityTargets(ctx.world, raw, {
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
    if (entities.length !== 1) {
      if (!ctx.lint) {
        throw new CommandError(
          entities.length === 0
            ? "No entity was found"
            : `The entity selector matched ${entities.length} entities; /data needs exactly one`,
          targetsStart,
          "entity-count",
        );
      }
      const dummy = entities[0] ?? ctx.source;
      return {
        kind: "entity",
        label: dummy ? `entity ${entityCustomName(dummy)}` : "entity",
        root: dummy?.nbt ?? Tag.compound(),
        entity: dummy ?? undefined,
      };
    }
    return {
      kind: "entity",
      label: `entity ${entityCustomName(entities[0])}`,
      root: entities[0].nbt,
      entity: entities[0],
    };
  }
  if (kind === "storage") {
    const idStart = cursor.mark();
    const id = cursor.readResourceLocation();
    const existing = ctx.world.storage(id, create && !ctx.lint);
    if (!existing && !ctx.lint) {
      throw new CommandError(
        `No storage named '${id}' exists yet`,
        idStart,
        "storage-not-found",
      );
    }
    const root = existing ?? Tag.compound();
    if (ctx.lint && create && !existing) {
      // Give lint a scratch compound so path validation continues.
    }
    return {
      kind: "storage",
      label: `storage ${id}`,
      root,
      id,
    };
  }
  cursor.error(
    `Expected 'block', 'entity' or 'storage', got '${kind}'`,
    start,
    "data-source",
  );
}

function parsePathSafe(
  cursor: Cursor,
  pathText: string,
  start: number,
): ReturnType<typeof parseNbtPath> {
  try {
    return parseNbtPath(pathText);
  } catch (error) {
    const pathError = error as NbtPathError;
    cursor.error(
      `Invalid NBT path '${pathText}': ${pathError.message}`,
      start + (pathError.index ?? 0),
      "path",
    );
  }
}

function isNumericTag(tag: NbtTag): tag is NbtTag & { value: number } {
  return (
    tag.type === "byte" ||
    tag.type === "short" ||
    tag.type === "int" ||
    tag.type === "long" ||
    tag.type === "float" ||
    tag.type === "double"
  );
}

function clone(tag: NbtTag): NbtTag {
  return JSON.parse(JSON.stringify(tag)) as NbtTag;
}
