import {
  cloneTag,
  type CompoundTag,
  isCompound,
  isList,
  isNumeric,
  type ListTag,
  type NbtTag,
  Tag,
  compoundMatches,
  stringifySnbt,
} from "./nbt";

/**
 * SNBT parser. Covers the syntax a datapack author actually types:
 * compounds, lists, typed arrays, all numeric suffixes, quoted and unquoted
 * strings, escapes, and booleans.
 */

export class SnbtError extends Error {
  index: number;
  constructor(message: string, index: number) {
    super(message);
    this.name = "SnbtError";
    this.index = index;
  }
}

const UNQUOTED = /[A-Za-z0-9._+\-]/;

class Parser {
  text: string;
  pos = 0;
  constructor(text: string) {
    this.text = text;
  }

  error(message: string, at = this.pos): never {
    throw new SnbtError(message, at);
  }

  skip(): void {
    while (this.pos < this.text.length && /\s/.test(this.text[this.pos])) {
      this.pos++;
    }
  }

  parseValue(): NbtTag {
    this.skip();
    const ch = this.text[this.pos];
    if (ch === undefined) this.error("Expected a tag, reached the end");
    if (ch === "{") return this.parseCompound();
    if (ch === "[") return this.parseListOrArray();
    if (ch === '"' || ch === "'") return Tag.string(this.parseQuoted());
    return this.parseNumberOrString();
  }

  parseCompound(): CompoundTag {
    const start = this.pos;
    this.pos++; // {
    const value: Record<string, NbtTag> = {};
    this.skip();
    if (this.text[this.pos] === "}") {
      this.pos++;
      return Tag.compound(value);
    }
    while (true) {
      this.skip();
      let key: string;
      const ch = this.text[this.pos];
      if (ch === '"' || ch === "'") key = this.parseQuoted();
      else {
        const begin = this.pos;
        while (this.pos < this.text.length && UNQUOTED.test(this.text[this.pos])) {
          this.pos++;
        }
        key = this.text.slice(begin, this.pos);
      }
      if (key.length === 0) this.error("Expected a compound key");
      this.skip();
      if (this.text[this.pos] !== ":") this.error("Expected ':' after compound key");
      this.pos++;
      value[key] = this.parseValue();
      this.skip();
      const next = this.text[this.pos];
      if (next === ",") {
        this.pos++;
        this.skip();
        if (this.text[this.pos] === "}") {
          this.pos++;
          return Tag.compound(value);
        }
        continue;
      }
      if (next === "}") {
        this.pos++;
        return Tag.compound(value);
      }
      this.error("Expected ',' or '}' in compound");
    }
    void start;
  }

  parseListOrArray(): NbtTag {
    this.pos++; // [
    this.skip();
    const marker = this.text[this.pos];
    const isArray =
      (marker === "B" || marker === "I" || marker === "L") &&
      this.text[this.pos + 1] === ";";
    if (isArray) {
      this.pos += 2;
      const numbers: number[] = [];
      this.skip();
      if (this.text[this.pos] === "]") {
        this.pos++;
        return marker === "B"
          ? Tag.byteArray(numbers)
          : marker === "I"
            ? Tag.intArray(numbers)
            : Tag.longArray(numbers);
      }
      while (true) {
        this.skip();
        const numStart = this.pos;
        const value = this.parseNumberOrString();
        if (value.type === "string") {
          this.error("Expected a number in an NBT array", numStart);
        }
        if (!isNumeric(value) || value.type === "float" || value.type === "double") {
          this.error("Arrays only accept integer numbers", numStart);
        }
        numbers.push(value.value);
        this.skip();
        const next = this.text[this.pos];
        if (next === ",") {
          this.pos++;
          this.skip();
          if (this.text[this.pos] === "]") {
            this.pos++;
            break;
          }
          continue;
        }
        if (next === "]") {
          this.pos++;
          break;
        }
        this.error("Expected ',' or ']' in array");
      }
      return marker === "B"
        ? Tag.byteArray(numbers)
        : marker === "I"
          ? Tag.intArray(numbers)
          : Tag.longArray(numbers);
    }

    const items: NbtTag[] = [];
    if (this.text[this.pos] === "]") {
      this.pos++;
      return Tag.list([], "end");
    }
    while (true) {
      items.push(this.parseValue());
      this.skip();
      const next = this.text[this.pos];
      if (next === ",") {
        this.pos++;
        this.skip();
        if (this.text[this.pos] === "]") {
          this.pos++;
          break;
        }
        continue;
      }
      if (next === "]") {
        this.pos++;
        break;
      }
      this.error("Expected ',' or ']' in list");
    }
    // [int[]] style typed lists from 1.21.5 components: not needed for this
    // build, but the parser reports a clear message rather than a crash.
    return Tag.list(items);
  }

  parseQuoted(): string {
    const quote = this.text[this.pos];
    const start = this.pos;
    this.pos++;
    let out = "";
    while (this.pos < this.text.length) {
      const ch = this.text[this.pos];
      if (ch === quote) {
        this.pos++;
        return out;
      }
      if (ch === "\\") {
        this.pos++;
        const esc = this.text[this.pos];
        switch (esc) {
          case "n":
            out += "\n";
            break;
          case "t":
            out += "\t";
            break;
          case "r":
            out += "\r";
            break;
          case "b":
            out += "\b";
            break;
          case "f":
            out += "\f";
            break;
          case "\\":
            out += "\\";
            break;
          case '"':
            out += '"';
            break;
          case "'":
            out += "'";
            break;
          case "u": {
            const hex = this.text.slice(this.pos + 1, this.pos + 5);
            if (!/^[0-9a-fA-F]{4}$/.test(hex)) this.error("Invalid unicode escape");
            out += String.fromCharCode(parseInt(hex, 16));
            this.pos += 4;
            break;
          }
          default:
            // SNBT is lenient: unknown escapes keep the escaped character.
            out += esc ?? "";
        }
        this.pos++;
        continue;
      }
      out += ch;
      this.pos++;
    }
    this.error("Unterminated string", start);
  }

  parseNumberOrString(): NbtTag {
    const start = this.pos;
    while (this.pos < this.text.length && UNQUOTED.test(this.text[this.pos])) {
      this.pos++;
    }
    const raw = this.text.slice(start, this.pos);
    if (raw.length === 0) this.error("Expected a value");
    if (raw === "true") return Tag.byte(1);
    if (raw === "false") return Tag.byte(0);
    const numeric = /^([+-]?(?:\d+\.?\d*|\.\d+)(?:[eE][+-]?\d+)?)([bBsSlLfFdD])?$/.exec(raw);
    if (numeric) {
      const n = Number(numeric[1]);
      if (Number.isNaN(n)) return Tag.string(raw);
      const suffix = numeric[2]?.toLowerCase();
      switch (suffix) {
        case "b":
          return Tag.byte(n);
        case "s":
          return Tag.short(n);
        case "l":
          return Tag.long(n);
        case "f":
          return Tag.float(n);
        case "d":
          return Tag.double(n);
        default:
          return Number.isInteger(n) ? Tag.int(n) : Tag.double(n);
      }
    }
    return Tag.string(raw);
  }
}

export function parseSnbt(text: string): NbtTag {
  const parser = new Parser(text);
  const value = parser.parseValue();
  parser.skip();
  if (parser.pos < text.length) {
    throw new SnbtError(
      `Unexpected trailing characters: ${text.slice(parser.pos, parser.pos + 16)}`,
      parser.pos,
    );
  }
  return value;
}

/** Parse a value that is expected to be a compound (macro args, storage). */
export function parseSnbtCompound(text: string): CompoundTag {
  const tag = parseSnbt(text);
  if (!isCompound(tag)) {
    throw new SnbtError("Expected a compound tag", 0);
  }
  return tag;
}

export function parseJsonToTag(text: string): NbtTag {
  return jsonToTag(JSON.parse(text));
}

export function jsonToTag(value: unknown): NbtTag {
  if (value === null || value === undefined) return Tag.string("");
  if (typeof value === "boolean") return Tag.byte(value ? 1 : 0);
  if (typeof value === "number") {
    return Number.isInteger(value) ? Tag.int(value) : Tag.double(value);
  }
  if (typeof value === "string") return Tag.string(value);
  if (Array.isArray(value)) return Tag.list(value.map(jsonToTag));
  if (typeof value === "object") {
    const out: Record<string, NbtTag> = {};
    for (const [k, v] of Object.entries(value as Record<string, unknown>)) {
      out[k] = jsonToTag(v);
    }
    return Tag.compound(out);
  }
  return Tag.string(String(value));
}

export function tryParseSnbt(text: string): { ok: true; tag: NbtTag } | { ok: false; message: string; index: number } {
  try {
    return { ok: true, tag: parseSnbt(text) };
  } catch (error) {
    if (error instanceof SnbtError) {
      return { ok: false, message: error.message, index: error.index };
    }
    return { ok: false, message: String(error), index: 0 };
  }
}

/**
 * Text components arrived at "SNBT, not JSON" in 1.21.5. Older packs use JSON.
 * This accepts both, and a bare string as SNBT would.
 */
export function parseTextComponent(raw: string): NbtTag {
  const trimmed = raw.trim();
  if (trimmed.startsWith('"') || trimmed.startsWith("'")) {
    try {
      return parseSnbt(trimmed);
    } catch {
      return Tag.string(trimmed.replace(/^["']|["']$/g, ""));
    }
  }
  if (trimmed.startsWith("{") || trimmed.startsWith("[")) {
    try {
      return parseSnbt(trimmed);
    } catch (snbtError) {
      try {
        return jsonToTag(JSON.parse(trimmed));
      } catch {
        throw snbtError;
      }
    }
  }
  // 26.3 still accepts a bare word in many component positions. Datapack
  // authors lean on it in `/title`; treat it as literal text.
  return Tag.string(trimmed);
}

export function serializeComponent(tag: NbtTag): string {
  return typeof tag.value === "string" ? tag.value : stringifySnbt(tag, { typed: false });
}

export { cloneTag, compoundMatches };
export type { CompoundTag, ListTag };
export { isCompound, isList };
