/**
 * NBT (Named Binary Tag) model.
 *
 * cmdlib keeps the same tag system Minecraft uses, because a datapack's whole
 * job is reading and writing it. Values are tagged objects rather than plain
 * JS values so that `data get` can report `1b` vs `1s` vs `1L`, and so a
 * round-trip through the editor never silently changes a number's type.
 */

export type TagType =
  | "byte"
  | "short"
  | "int"
  | "long"
  | "float"
  | "double"
  | "string"
  | "list"
  | "compound"
  | "byteArray"
  | "intArray"
  | "longArray";

export type NumericTagType = Extract<
  TagType,
  "byte" | "short" | "int" | "long" | "float" | "double"
>;

export interface ByteTag {
  type: "byte";
  value: number;
}
export interface ShortTag {
  type: "short";
  value: number;
}
export interface IntTag {
  type: "int";
  value: number;
}
export interface LongTag {
  type: "long";
  value: number;
}
export interface FloatTag {
  type: "float";
  value: number;
}
export interface DoubleTag {
  type: "double";
  value: number;
}
export interface StringTag {
  type: "string";
  value: string;
}
export interface ListTag {
  type: "list";
  value: NbtTag[];
  elementType: TagType | "end";
}
export interface CompoundTag {
  type: "compound";
  value: Record<string, NbtTag>;
}
export interface ByteArrayTag {
  type: "byteArray";
  value: number[];
}
export interface IntArrayTag {
  type: "intArray";
  value: number[];
}
export interface LongArrayTag {
  type: "longArray";
  value: number[];
}

export type NbtTag =
  | ByteTag
  | ShortTag
  | IntTag
  | LongTag
  | FloatTag
  | DoubleTag
  | StringTag
  | ListTag
  | CompoundTag
  | ByteArrayTag
  | IntArrayTag
  | LongArrayTag;

export type NbtNumber =
  | ByteTag
  | ShortTag
  | IntTag
  | LongTag
  | FloatTag
  | DoubleTag;

export const Tag = {
  byte: (value = 0): ByteTag => ({ type: "byte", value: clampByte(value) }),
  short: (value = 0): ShortTag => ({ type: "short", value: clampShort(value) }),
  int: (value = 0): IntTag => ({ type: "int", value: clampInt(value) }),
  long: (value = 0): LongTag => ({ type: "long", value: Math.trunc(value) }),
  float: (value = 0): FloatTag => ({ type: "float", value: Math.fround(value) }),
  double: (value = 0): DoubleTag => ({ type: "double", value }),
  string: (value = ""): StringTag => ({ type: "string", value }),
  list: (value: NbtTag[] = [], elementType?: TagType | "end"): ListTag => ({
    type: "list",
    value,
    elementType: elementType ?? inferElementType(value),
  }),
  compound: (value: Record<string, NbtTag> = {}): CompoundTag => ({
    type: "compound",
    value,
  }),
  byteArray: (value: number[] = []): ByteArrayTag => ({
    type: "byteArray",
    value: value.map(clampByte),
  }),
  intArray: (value: number[] = []): IntArrayTag => ({
    type: "intArray",
    value: value.map(clampInt),
  }),
  longArray: (value: number[] = []): LongArrayTag => ({
    type: "longArray",
    value: value.map((n) => Math.trunc(n)),
  }),
};

export function clampByte(n: number): number {
  return ((Math.trunc(n) + 128) % 256 + 256) % 256 - 128;
}
export function clampShort(n: number): number {
  return ((Math.trunc(n) + 32768) % 65536 + 65536) % 65536 - 32768;
}
export function clampInt(n: number): number {
  const v = Math.trunc(n) % 4294967296;
  return v > 2147483647 ? v - 4294967296 : v < -2147483648 ? v + 4294967296 : v;
}

export function isNumeric(tag: NbtTag | undefined | null): tag is NbtNumber {
  if (!tag) return false;
  return (
    tag.type === "byte" ||
    tag.type === "short" ||
    tag.type === "int" ||
    tag.type === "long" ||
    tag.type === "float" ||
    tag.type === "double"
  );
}
export function isCompound(tag: NbtTag | undefined | null): tag is CompoundTag {
  return !!tag && tag.type === "compound";
}
export function isList(tag: NbtTag | undefined | null): tag is ListTag {
  return !!tag && tag.type === "list";
}
export function isString(tag: NbtTag | undefined | null): tag is StringTag {
  return !!tag && tag.type === "string";
}
export function isByteArray(tag: NbtTag | undefined | null): tag is ByteArrayTag {
  return !!tag && tag.type === "byteArray";
}
export function isIntArray(tag: NbtTag | undefined | null): tag is IntArrayTag {
  return !!tag && tag.type === "intArray";
}
export function isLongArray(tag: NbtTag | undefined | null): tag is LongArrayTag {
  return !!tag && tag.type === "longArray";
}
export function isNumericArray(
  tag: NbtTag | undefined | null,
): tag is ByteArrayTag | IntArrayTag | LongArrayTag {
  return isByteArray(tag) || isIntArray(tag) || isLongArray(tag);
}

export function numericValue(tag: NbtTag | undefined | null): number | null {
  return isNumeric(tag) ? tag.value : null;
}

export function inferElementType(items: NbtTag[]): TagType | "end" {
  if (items.length === 0) return "end";
  const first = items[0].type;
  // Minecraft lists are homogeneous. If a heterogeneous list is supplied we
  // keep the first type and let the stringifier print each element faithfully.
  return items.every((i) => i.type === first) ? first : "end";
}

export function cloneTag(tag: NbtTag): NbtTag {
  switch (tag.type) {
    case "compound": {
      const out: Record<string, NbtTag> = {};
      for (const [k, v] of Object.entries(tag.value)) out[k] = cloneTag(v);
      return { type: "compound", value: out };
    }
    case "list":
      return {
        type: "list",
        elementType: tag.elementType,
        value: tag.value.map(cloneTag),
      };
    case "byteArray":
      return { type: "byteArray", value: [...tag.value] };
    case "intArray":
      return { type: "intArray", value: [...tag.value] };
    case "longArray":
      return { type: "longArray", value: [...tag.value] };
    default:
      return { type: tag.type, value: tag.value } as NbtTag;
  }
}

export function tagEquals(a: NbtTag, b: NbtTag): boolean {
  if (a.type !== b.type) {
    // Numeric tags compare by value in Minecraft's NBT matcher.
    if (isNumeric(a) && isNumeric(b)) return a.value === b.value;
    return false;
  }
  if (isNumeric(a) && isNumeric(b)) return a.value === b.value;
  switch (a.type) {
    case "string":
      return a.value === (b as StringTag).value;
    case "list": {
      const other = b as ListTag;
      return (
        a.value.length === other.value.length &&
        a.value.every((item, i) => tagEquals(item, other.value[i]))
      );
    }
    case "compound": {
      const other = b as CompoundTag;
      const ak = Object.keys(a.value);
      if (ak.length !== Object.keys(other.value).length) return false;
      return ak.every(
        (k) => k in other.value && tagEquals(a.value[k], other.value[k]),
      );
    }
    default:
      return JSON.stringify(a.value) === JSON.stringify(b.value);
  }
}

/** Minecraft's selector NBT matcher is a subset match, not an equality match. */
export function compoundMatches(subject: NbtTag, pattern: NbtTag): boolean {
  if (!isCompound(subject) || !isCompound(pattern)) {
    return tagEquals(subject, pattern);
  }
  for (const [key, value] of Object.entries(pattern.value)) {
    const actual = subject.value[key];
    if (!actual) return false;
    if (isCompound(value) && isCompound(actual)) {
      if (!compoundMatches(actual, value)) return false;
    } else if (!tagEquals(actual, value)) {
      return false;
    }
  }
  return true;
}

export function getCompound(tag: CompoundTag, key: string): NbtTag | undefined {
  return tag.value[key];
}

export function setCompound(
  tag: CompoundTag,
  key: string,
  value: NbtTag,
): void {
  tag.value[key] = value;
}

export function tagTypeName(tag: NbtTag): string {
  switch (tag.type) {
    case "byte":
      return "byte";
    case "short":
      return "short";
    case "int":
      return "int";
    case "long":
      return "long";
    case "float":
      return "float";
    case "double":
      return "double";
    case "string":
      return "string";
    case "list":
      return "list";
    case "compound":
      return "compound";
    case "byteArray":
      return "byte array";
    case "intArray":
      return "int array";
    case "longArray":
      return "long array";
  }
}

export function tagToPlain(tag: NbtTag): unknown {
  switch (tag.type) {
    case "byte":
    case "short":
    case "int":
    case "long":
    case "float":
    case "double":
    case "string":
      return tag.value;
    case "list":
      return tag.value.map(tagToPlain);
    case "compound": {
      const out: Record<string, unknown> = {};
      for (const [k, v] of Object.entries(tag.value)) out[k] = tagToPlain(v);
      return out;
    }
    default:
      return [...tag.value];
  }
}

export function plainToTag(value: unknown): NbtTag {
  if (value === null || value === undefined) return Tag.string("");
  if (typeof value === "boolean") return Tag.byte(value ? 1 : 0);
  if (typeof value === "number") {
    return Number.isInteger(value) ? Tag.int(value) : Tag.double(value);
  }
  if (typeof value === "string") return Tag.string(value);
  if (Array.isArray(value)) return Tag.list(value.map(plainToTag));
  if (typeof value === "object") {
    const out: Record<string, NbtTag> = {};
    for (const [k, v] of Object.entries(value as Record<string, unknown>)) {
      out[k] = plainToTag(v);
    }
    return Tag.compound(out);
  }
  return Tag.string(String(value));
}

const NEEDS_QUOTE = /[^A-Za-z0-9._+\-]/;

export function quoteString(value: string): string {
  let out = '"';
  for (const ch of value) {
    switch (ch) {
      case "\\":
        out += "\\\\";
        break;
      case '"':
        out += '\\"';
        break;
      case "\n":
        out += "\\n";
        break;
      case "\t":
        out += "\\t";
        break;
      case "\r":
        out += "\\r";
        break;
      default:
        out += ch;
    }
  }
  return out + '"';
}

function numberSuffix(tag: NbtTag): string {
  switch (tag.type) {
    case "byte":
      return "b";
    case "short":
      return "s";
    case "long":
      return "L";
    case "float":
      return "f";
    case "double":
      return "d";
    default:
      return "";
  }
}

function formatNumber(tag: NbtTag & { value: number }): string {
  switch (tag.type) {
    case "float": {
      const v = tag.value;
      return `${Number.isFinite(v) ? (Number.isInteger(v) ? v.toFixed(1) : String(v)) : v > 0 ? "Infinity" : "-Infinity"}f`;
    }
    case "double": {
      const v = tag.value;
      return `${Number.isFinite(v) ? (Number.isInteger(v) ? v.toFixed(1) : String(v)) : v > 0 ? "Infinity" : "-Infinity"}d`;
    }
    default:
      return `${tag.value}${numberSuffix(tag)}`;
  }
}

export interface SnbtOptions {
  /** Print `1b` / `1.0d` type suffixes. Vanilla's data output does this. */
  typed?: boolean;
  /** Pretty print with newlines and indentation. */
  pretty?: boolean;
  /** Quote strings that would be valid unquoted SNBT. */
  quoteKeys?: boolean;
}

export function stringifySnbt(tag: NbtTag, options: SnbtOptions = {}): string {
  const typed = options.typed ?? true;
  const pretty = options.pretty ?? false;
  const quoteKeys = options.quoteKeys ?? false;

  const walk = (node: NbtTag, indent: number): string => {
    const pad = pretty ? "\n" + "  ".repeat(indent + 1) : "";
    const endPad = pretty ? "\n" + "  ".repeat(indent) : "";
    switch (node.type) {
      case "string":
        return quoteString(node.value);
      case "byte":
      case "short":
      case "int":
      case "long":
      case "float":
      case "double":
        return typed
          ? formatNumber(node)
          : String(node.value);
      case "list": {
        if (node.value.length === 0) return "[]";
        const items = node.value.map((v) => walk(v, indent + 1)).join(pretty ? "," + pad : ", ");
        return `[${pad}${items}${endPad}]`;
      }
      case "compound": {
        const keys = Object.keys(node.value);
        if (keys.length === 0) return "{}";
        const items = keys
          .map((key) => {
            const keyText =
              quoteKeys || NEEDS_QUOTE.test(key) || key.length === 0
                ? quoteString(key)
                : key;
            return `${keyText}: ${walk(node.value[key], indent + 1)}`;
          })
          .join(pretty ? "," + pad : ", ");
        return `{${pad}${items}${endPad}}`;
      }
      case "byteArray":
        return `[B;${node.value.join(", ")}]`;
      case "intArray":
        return `[I;${node.value.join(", ")}]`;
      case "longArray":
        return `[L;${node.value.map((v) => `${v}L`).join(", ")}]`;
    }
  };

  return walk(tag, 0);
}

/** Compact, colour-friendly representation for inspector trees. */
export function previewValue(tag: NbtTag, maxLength = 80): string {
  const text = stringifySnbt(tag, { typed: false, quoteKeys: true });
  return text.length > maxLength ? text.slice(0, maxLength - 1) + "…" : text;
}
