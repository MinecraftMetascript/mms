import {
  cloneTag,
  type CompoundTag,
  isByteArray,
  isCompound,
  isIntArray,
  isList,
  isLongArray,
  isNumericArray,
  type ListTag,
  type NbtTag,
  numericValue,
  Tag,
} from "./nbt";
import { parseSnbt, SnbtError } from "./snbt";

/**
 * NBT path support for `/data`.
 *
 * Implemented: dotted keys, quoted keys, list indexes, `[]` wildcards and
 * `[{...}]` compound filters. That is the vocabulary datapack authors use for
 * config storage and score snapshots.
 */

export type NbtPathSegment =
  | { kind: "key"; key: string }
  | { kind: "index"; index: number }
  | { kind: "all" }
  | { kind: "match"; pattern: CompoundTag }
  | { kind: "filter"; pattern: CompoundTag }
  | { kind: "filterIndex"; pattern: CompoundTag; index: number };

export class NbtPathError extends Error {
  index: number;
  constructor(message: string, index: number) {
    super(message);
    this.name = "NbtPathError";
    this.index = index;
  }
}

const UNQUOTED_KEY = /[^.[\]{}]/;

export function parseNbtPath(text: string): NbtPathSegment[] {
  const segments: NbtPathSegment[] = [];
  let i = 0;
  const error = (msg: string, at = i): never => {
    throw new NbtPathError(msg, at);
  };

  while (i < text.length) {
    if (text[i] === ".") {
      i++;
      continue;
    }
    if (text[i] === "{") {
      const start = i;
      let depth = 0;
      let inString: string | null = null;
      for (; i < text.length; i++) {
        const ch = text[i];
        if (inString) {
          if (ch === "\\") i++;
          else if (ch === inString) inString = null;
          continue;
        }
        if (ch === '"' || ch === "'") inString = ch;
        else if (ch === "{") depth++;
        else if (ch === "}") {
          depth--;
          if (depth === 0) {
            i++;
            break;
          }
        }
      }
      if (depth !== 0) error("Unclosed compound filter", start);
      const raw = text.slice(start, i);
      let pattern: CompoundTag = Tag.compound();
      try {
        pattern = parseSnbt(raw) as CompoundTag;
      } catch (e) {
        const snbt = e as SnbtError;
        error(snbt.message, start + snbt.index);
      }
      segments.push({ kind: "match", pattern });
      continue;
    }
    if (text[i] === "[") {
      const open = i;
      i++;
      while (text[i] === " ") i++;
      if (text[i] === "]") {
        segments.push({ kind: "all" });
        i++;
        continue;
      }
      if (text[i] === "{") {
        // Find the balanced compound, then an optional index before ']'.
        const start = i;
        let depth = 0;
        let inString: string | null = null;
        for (; i < text.length; i++) {
          const ch = text[i];
          if (inString) {
            if (ch === "\\") i++;
            else if (ch === inString) inString = null;
            continue;
          }
          if (ch === '"' || ch === "'") inString = ch;
          else if (ch === "{") depth++;
          else if (ch === "}") {
            depth--;
            if (depth === 0) {
              i++;
              break;
            }
          }
        }
        if (depth !== 0) error("Unclosed compound filter", start);
        const raw = text.slice(start, i);
        let pattern: CompoundTag = Tag.compound();
        try {
          pattern = parseSnbt(raw) as CompoundTag;
        } catch (e) {
          const snbt = e as SnbtError;
          error(snbt.message, start + snbt.index);
        }
        while (text[i] === " ") i++;
        if (text[i] === "]") {
          segments.push({ kind: "filter", pattern });
          i++;
          continue;
        }
        if (text[i] === ",") {
          // `[{...}, 0]` is not valid NBT path syntax; keep the message clear.
          error("Compound filters only accept one match", i);
        }
        error("Expected ']' after compound filter", i);
        void open;
      }
      const start = i;
      while (i < text.length && /[0-9-]/.test(text[i])) i++;
      if (start === i) error("Expected a list index");
      const index = Number(text.slice(start, i));
      while (text[i] === " ") i++;
      if (text[i] !== "]") error("Expected ']' after list index");
      i++;
      if (!Number.isInteger(index)) {
        error("List indexes must be integers", start);
      }
      segments.push({ kind: "index", index });
      continue;
    }

    // Key: quoted or a run of characters up to the next separator.
    let key: string;
    if (text[i] === '"' || text[i] === "'") {
      const quote = text[i];
      i++;
      let out = "";
      while (i < text.length && text[i] !== quote) {
        if (text[i] === "\\") {
          i++;
          const esc = text[i];
          out += esc === "n" ? "\n" : esc === "t" ? "\t" : esc ?? "";
          i++;
          continue;
        }
        out += text[i];
        i++;
      }
      if (text[i] !== quote) error("Unclosed quoted path key");
      i++;
      key = out;
    } else {
      const start = i;
      while (i < text.length && UNQUOTED_KEY.test(text[i])) i++;
      key = text.slice(start, i);
      if (key.length === 0) error(`Unexpected '${text[i]}' in path`, i);
    }
    segments.push({ kind: "key", key });
    if (i < text.length && text[i] !== "." && text[i] !== "[" && text[i] !== "{") {
      error(`Unexpected '${text[i]}' in path`, i);
    }
  }
  return segments;
}

export interface PathNode {
  parent: NbtTag | null;
  /** Parent slot: compound key or list index. */
  key: string | number | null;
  value: NbtTag;
}

const TRANSFORM_FIELDS = new Set([
  "translation",
  "scale",
  "left_rotation",
  "right_rotation",
]);

function isTransformationMatrix(tag: NbtTag | undefined): tag is ListTag {
  return (
    isList(tag) &&
    tag.value.length === 16 &&
    tag.value.every((item) => numericValue(item) !== null)
  );
}

function matrixField(tag: ListTag, field: string): NbtTag | null {
  const value = tag.value.map((item) => numericValue(item) ?? 0);
  // The final cell is the homogeneous w divisor. Projective matrices use it
  // to scale translation/scale, which is how gm encodes division and products.
  const w = value[15] === 0 ? 1 : value[15];
  if (field === "translation") {
    return Tag.list(
      [
        Tag.float(value[3] / w),
        Tag.float(value[7] / w),
        Tag.float(value[11] / w),
      ],
      "float",
    );
  }
  if (field === "scale") {
    const lengths = [0, 1, 2].map(
      (column) =>
        Math.hypot(
          value[column],
          value[4 + column],
          value[8 + column],
        ) / w,
    );
    return Tag.list(lengths.map((n) => Tag.float(n)), "float");
  }
  // Matrix decomposition rotations are only exposed as the identity unless
  // the matrix is used for rendering. Translation and scale cover the common
  // datapack use cases.
  return Tag.list(
    [Tag.float(0), Tag.float(0), Tag.float(0), Tag.float(1)],
    "float",
  );
}

export function resolvePath(root: NbtTag, segments: NbtPathSegment[]): PathNode[] {
  let nodes: PathNode[] = [{ parent: null, key: null, value: root }];

  for (const segment of segments) {
    const next: PathNode[] = [];
    for (const node of nodes) {
      const current = node.value;
      switch (segment.kind) {
        case "key": {
          if (!isCompound(current)) {
            if (isTransformationMatrix(current) && TRANSFORM_FIELDS.has(segment.key)) {
              const value = matrixField(current, segment.key);
              if (value) {
                next.push({ parent: current, key: null, value });
              }
            }
            continue;
          }
          const value = current.value[segment.key];
          if (value === undefined) continue;
          next.push({ parent: current, key: segment.key, value });
          break;
        }
        case "index": {
          const length = isList(current)
            ? current.value.length
            : isNumericArray(current)
              ? current.value.length
              : 0;
          const index = segment.index < 0 ? length + segment.index : segment.index;
          if (isList(current)) {
            const value = current.value[index];
            if (value === undefined) continue;
            next.push({ parent: current, key: index, value });
          } else if (isNumericArray(current)) {
            const value = current.value[index];
            if (value === undefined) continue;
            next.push({
              parent: current,
              key: index,
              value: arrayElementTag(current, value),
            });
          }
          break;
        }
        case "match": {
          if (!isCompound(current)) continue;
          const matches = Object.entries(segment.pattern.value).every(
            ([key, value]) =>
              current.value[key] !== undefined &&
              tagMatches(current.value[key], value),
          );
          if (matches) next.push(node);
          break;
        }
        case "all": {
          if (isList(current)) {
            current.value.forEach((value, index) =>
              next.push({ parent: current, key: index, value }),
            );
          } else if (isNumericArray(current)) {
            current.value.forEach((value, index) =>
              next.push({
                parent: current,
                key: index,
                value: arrayElementTag(current, value),
              }),
            );
          } else if (isCompound(current)) {
            for (const [key, value] of Object.entries(current.value)) {
              next.push({ parent: current, key, value });
            }
          }
          break;
        }
        case "filter": {
          const source = isList(current)
            ? current.value
            : isCompound(current)
              ? Object.values(current.value)
              : [];
          source.forEach((candidate, index) => {
            if (!isCompound(candidate)) return;
            const matches = Object.entries(segment.pattern.value).every(
              ([k, v]) => candidate.value[k] && tagMatches(candidate.value[k], v),
            );
            if (!matches) return;
            const parent = isList(current) ? current : current;
            next.push({
              parent,
              key: isList(current) ? index : Object.keys(current.value)[index] ?? null,
              value: candidate,
            });
          });
          break;
        }
        case "filterIndex": {
          const matches = next.filter((n) => n.key === segment.index);
          next.push(...matches);
          break;
        }
      }
    }
    nodes = next;
  }
  return nodes;
}

function arrayElementTag(
  array: ReturnType<typeof Tag.byteArray> | ReturnType<typeof Tag.intArray> | ReturnType<typeof Tag.longArray>,
  value: number,
): NbtTag {
  if (isByteArray(array)) return Tag.byte(value);
  if (isIntArray(array)) return Tag.int(value);
  return Tag.long(value);
}

function tagMatches(a: NbtTag, b: NbtTag): boolean {
  if (isCompound(a) && isCompound(b)) {
    return Object.entries(b.value).every(
      ([k, v]) => a.value[k] !== undefined && tagMatches(a.value[k], v),
    );
  }
  if (isList(a) && isList(b)) {
    return (
      a.value.length === b.value.length &&
      a.value.every((v, i) => tagMatches(v, b.value[i]))
    );
  }
  if (a.type !== b.type) return false;
  if (a.type === "list" || a.type === "compound") return false;
  return JSON.stringify(a.value) === JSON.stringify(b.value);
}

export function getPathValues(root: NbtTag, segments: NbtPathSegment[]): NbtTag[] {
  return resolvePath(root, segments).map((n) => n.value);
}

/** Resolve exactly one node, or throw a human-readable error. */
export function getSinglePath(
  root: NbtTag,
  segments: NbtPathSegment[],
): PathNode | null {
  const matches = resolvePath(root, segments);
  if (matches.length === 0) return null;
  return matches[0];
}

export type ModifyMode =
  | "set"
  | "merge"
  | "append"
  | "prepend"
  | "insert"
  | "remove";

function listItems(value: NbtTag): NbtTag[] {
  return isList(value)
    ? value.value.map((item) => cloneTag(item))
    : [cloneTag(value)];
}

export function applyPathModify(
  root: CompoundTag,
  segments: NbtPathSegment[],
  mode: ModifyMode,
  value: NbtTag | null,
  insertIndex = 0,
): void {
  if (segments.length === 0) {
    throw new Error("A target path is required");
  }
  const last = segments[segments.length - 1];
  if (last.kind === "all" || last.kind === "filter" || last.kind === "filterIndex") {
    throw new Error("Wildcard paths cannot be modified");
  }

  const parentNodes = resolvePath(
    root,
    segments.slice(0, -1),
  ).filter((n) => n.key !== null || segments.length === 1);

  let container: NbtTag | null;
  if (segments.length === 1) {
    container = root;
  } else {
    const found = parentNodes.length > 0 ? parentNodes[0].value : null;
    container = found ?? createPath(root, segments.slice(0, -1));
  }

  if (!container) throw new Error("Could not resolve the target path");

  if (last.kind === "key") {
    if (!isCompound(container)) {
      throw new Error("Cannot modify a key on a non-compound tag");
    }
    const current = container.value[last.key];
    switch (mode) {
      case "set":
        container.value[last.key] = cloneTag(value!);
        break;
      case "merge": {
        if (!isCompound(value)) throw new Error("merge requires a compound value");
        if (current && isCompound(current)) {
          mergeCompound(current, value);
        } else {
          container.value[last.key] = cloneTag(value);
        }
        break;
      }
      case "append":
      case "prepend":
      case "insert": {
        // MMS: like the game, a missing tag becomes an empty list first.
        if (!current) container.value[last.key] = Tag.list([], "end");
        const target = container.value[last.key];
        if (!isList(target)) {
          throw new Error(`Cannot ${mode} to a tag that is not a list`);
        }
        if (!value) throw new Error(`${mode} requires a value`);
        const items = listItems(value);
        if (mode === "append") target.value.push(...items);
        else if (mode === "prepend") target.value.unshift(...items);
        else {
          if (insertIndex < 0 || insertIndex > target.value.length) {
            throw new Error(`Index ${insertIndex} is outside the list`);
          }
          target.value.splice(insertIndex, 0, ...items);
        }
        break;
      }
      case "remove":
        delete container.value[last.key];
        break;
    }
    return;
  }

  if (last.kind === "index") {
    if (isNumericArray(container)) {
      if (mode !== "set") {
        throw new Error(`Cannot ${mode} an element of a ${container.type}`);
      }
      const index = last.index < 0 ? container.value.length + last.index : last.index;
      if (index < 0 || index >= container.value.length) {
        throw new Error(`Index ${last.index} is outside the array`);
      }
      const number = numericValue(value);
      if (number === null) {
        throw new Error("Array elements must be numbers");
      }
      container.value[index] =
        container.type === "byteArray"
          ? Tag.byte(number).value
          : container.type === "intArray"
            ? Tag.int(number).value
            : Tag.long(number).value;
      return;
    }
    if (!isList(container)) {
      throw new Error("Cannot index a non-list tag");
    }
    const index = last.index < 0 ? container.value.length + last.index : last.index;
    switch (mode) {
      case "set": {
        if (index >= container.value.length) {
          if (index === container.value.length) {
            container.value.push(cloneTag(value!));
          } else {
            throw new Error(`Index ${last.index} is outside the list`);
          }
        } else {
          container.value[index] = cloneTag(value!);
        }
        break;
      }
      case "merge": {
        const current = container.value[index];
        if (!isCompound(current) || !isCompound(value)) {
          throw new Error("merge requires two compounds");
        }
        mergeCompound(current, value);
        break;
      }
      case "append":
      case "prepend":
      case "insert": {
        const current = container.value[index];
        if (!current || !isList(current)) {
          throw new Error(`Cannot ${mode} to a tag that is not a list`);
        }
        const items = listItems(value!);
        if (mode === "append") current.value.push(...items);
        else if (mode === "prepend") current.value.unshift(...items);
        else {
          if (insertIndex < 0 || insertIndex > current.value.length) {
            throw new Error(`Index ${insertIndex} is outside the list`);
          }
          current.value.splice(insertIndex, 0, ...items);
        }
        break;
      }
      case "remove":
        if (index < 0 || index >= container.value.length) {
          throw new Error(`Index ${last.index} is outside the list`);
        }
        container.value.splice(index, 1);
        break;
    }
    return;
  }
}

/** Walk/create intermediate compounds for `data modify ... set`. */
function createPath(root: CompoundTag, segments: NbtPathSegment[]): NbtTag | null {
  let current: NbtTag = root;
  for (const segment of segments) {
    if (segment.kind === "key") {
      if (!isCompound(current)) return null;
      const compound = current as CompoundTag;
      let next: NbtTag | undefined = compound.value[segment.key];
      if (!next) {
        next = Tag.compound();
        compound.value[segment.key] = next;
      }
      current = next;
    } else if (segment.kind === "index") {
      if (!isList(current)) return null;
      const list = current as ListTag;
      let next: NbtTag | undefined = list.value[segment.index];
      if (!next) {
        if (segment.index !== list.value.length) return null;
        next = Tag.compound();
        list.value.push(next);
      }
      current = next;
    } else {
      return null;
    }
  }
  return current;
}

export function mergeCompound(target: CompoundTag, source: CompoundTag): void {
  for (const [key, value] of Object.entries(source.value)) {
    const existing = target.value[key];
    if (isCompound(existing) && isCompound(value)) {
      mergeCompound(existing, value);
    } else {
      target.value[key] = cloneTag(value);
    }
  }
}

/** Format a path back to text, for diagnostics and quick-fix previews. */
export function pathToString(segments: NbtPathSegment[]): string {
  return segments
    .map((s) => {
      switch (s.kind) {
        case "key":
          return s.key;
        case "index":
          return `[${s.index}]`;
        case "all":
          return "[]";
        case "filter":
          return `[{${
            Object.entries(s.pattern.value)
              .map(([k, v]) => `${k}:${JSON.stringify(v.value)}`)
              .join(",")
          }}]`;
        case "filterIndex":
          return `[${s.index}]`;
      }
    })
    .join(".");
}
