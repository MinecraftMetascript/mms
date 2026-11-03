import type { Vec2, Vec3 } from "./entities";
import type { NbtTag } from "./nbt";
import { parseSnbt, SnbtError } from "./snbt";

export class CommandError extends Error {
  index: number;
  code: string;
  constructor(message: string, index = 0, code = "command-error") {
    super(message);
    this.name = "CommandError";
    this.index = index;
    this.code = code;
  }
}

export class UnsupportedCommandError extends CommandError {
  constructor(command: string, index = 0) {
    super(
      `/${command} is a recognised Minecraft command, but cmdlib 0.1 does not simulate it yet. The editor still lints it against the 26.3 syntax.`,
      index,
      "unsupported-command",
    );
    this.name = "UnsupportedCommandError";
  }
}

export class Cursor {
  text: string;
  pos: number;
  end: number;

  constructor(text: string, start = 0, end = text.length) {
    this.text = text;
    this.pos = start;
    this.end = end;
  }

  skip(): this {
    while (this.pos < this.end && /\s/.test(this.text[this.pos])) this.pos++;
    return this;
  }

  done(): boolean {
    this.skip();
    return this.pos >= this.end;
  }

  peek(offset = 0): string {
    return this.text[this.pos + offset] ?? "";
  }

  mark(): number {
    return this.pos;
  }

  error(message: string, at = this.pos, code = "syntax"): never {
    throw new CommandError(message, at, code);
  }

  /** Read one whitespace-delimited word, honouring quotes. */
  /** Read an NBT path, keeping spaces inside `{...}` / `[...]` filters. */
  readNbtPathText(): string {
    this.skip();
    const start = this.pos;
    let depth = 0;
    let quote: string | null = null;
    while (this.pos < this.end) {
      const ch = this.text[this.pos];
      if (quote) {
        if (ch === "\\") this.pos++;
        else if (ch === quote) quote = null;
        this.pos++;
        continue;
      }
      if (ch === '"' || ch === "'") {
        quote = ch;
      } else if (ch === "{" || ch === "[") {
        depth++;
      } else if (ch === "}" || ch === "]") {
        depth = Math.max(0, depth - 1);
      } else if (/\s/.test(ch) && depth === 0) {
        break;
      }
      this.pos++;
    }
    return this.text.slice(start, this.pos);
  }

  readWord(): string {
    this.skip();
    const start = this.pos;
    if (this.peek() === '"' || this.peek() === "'") {
      return this.readQuoted();
    }
    while (this.pos < this.end && !/\s/.test(this.text[this.pos])) {
      this.pos++;
    }
    if (this.pos === start) this.error("Expected an argument");
    return this.text.slice(start, this.pos);
  }

  readQuoted(): string {
    const quote = this.peek();
    const start = this.pos;
    this.pos++;
    let out = "";
    while (this.pos < this.end) {
      const ch = this.text[this.pos];
      if (ch === "\\") {
        out += this.text[this.pos + 1] ?? "";
        this.pos += 2;
        continue;
      }
      if (ch === quote) {
        this.pos++;
        return out;
      }
      out += ch;
      this.pos++;
    }
    this.error("Unterminated quoted string", start);
  }

  /** Read raw text including the quotes, for text component arguments. */
  readQuotedRaw(): string {
    const start = this.pos;
    const quote = this.peek();
    this.pos++;
    while (this.pos < this.end) {
      if (this.text[this.pos] === "\\") {
        this.pos += 2;
        continue;
      }
      if (this.text[this.pos] === quote) {
        this.pos++;
        return this.text.slice(start, this.pos);
      }
      this.pos++;
    }
    this.error("Unterminated quoted string", start);
  }

  /** Read a balanced `{...}` or `[...]` block, including delimiters. */
  readBalanced(): string {
    this.skip();
    const start = this.pos;
    const open = this.peek();
    const close = open === "{" ? "}" : open === "[" ? "]" : "";
    if (!close) this.error("Expected a compound or list");
    let depth = 0;
    let quote: string | null = null;
    while (this.pos < this.end) {
      const ch = this.text[this.pos];
      if (quote) {
        if (ch === "\\") {
          this.pos += 2;
          continue;
        }
        if (ch === quote) quote = null;
        this.pos++;
        continue;
      }
      if (ch === '"' || ch === "'") quote = ch;
      else if (ch === open) depth++;
      else if (ch === close) {
        depth--;
        if (depth === 0) {
          this.pos++;
          return this.text.slice(start, this.pos);
        }
      }
      this.pos++;
    }
    this.error(`Unclosed '${open}' tag`, start);
  }

  readSnbt(): { tag: NbtTag; start: number; end: number; raw: string } {
    this.skip();
    const start = this.pos;
    let raw: string;
    if (this.peek() === "{" || this.peek() === "[") {
      raw = this.readBalanced();
    } else if (this.peek() === '"' || this.peek() === "'") {
      raw = this.readQuotedRaw();
    } else {
      raw = this.readWord();
    }
    try {
      return { tag: parseSnbt(raw), start, end: this.pos, raw };
    } catch (error) {
      const snbt = error as SnbtError;
      this.error(snbt.message, start + (snbt.index ?? 0), "nbt");
    }
  }

  /** Read a text component. Bare text runs to the end of the line. */
  readComponent(): { tag: NbtTag; start: number; end: number; raw: string } {
    this.skip();
    const start = this.pos;
    let raw: string;
    if (this.peek() === "{" || this.peek() === "[") {
      raw = this.readBalanced();
    } else if (this.peek() === '"' || this.peek() === "'") {
      raw = this.readQuotedRaw();
    } else {
      raw = this.readRest();
      if (raw.trim().length === 0) this.error("Expected a text component");
      return {
        tag: { type: "string", value: raw.trim() },
        start,
        end: this.pos,
        raw,
      };
    }
    try {
      return { tag: parseSnbt(raw), start, end: this.pos, raw };
    } catch (error) {
      const snbt = error as SnbtError;
      this.error(`Invalid text component: ${snbt.message}`, start, "component");
    }
  }

  /**
   * Read a target list (`@e[...]`, `Name`, `@a,@p`) up to the next space at
   * bracket depth zero. Comma-separated names stay together.
   */
  readTargetsRaw(): { raw: string; start: number; end: number } {
    this.skip();
    const start = this.pos;
    let depth = 0;
    let quote: string | null = null;
    while (this.pos < this.end) {
      const ch = this.text[this.pos];
      if (quote) {
        if (ch === "\\") {
          this.pos += 2;
          continue;
        }
        if (ch === quote) quote = null;
        this.pos++;
        continue;
      }
      if (ch === '"' || ch === "'") quote = ch;
      else if (ch === "[" || ch === "{") depth++;
      else if (ch === "]" || ch === "}") depth--;
      else if (/\s/.test(ch) && depth === 0) break;
      this.pos++;
    }
    const raw = this.text.slice(start, this.pos);
    if (raw.length === 0) this.error("Expected a target selector or name");
    return { raw, start, end: this.pos };
  }

  readResourceLocation(allowTag = false): string {
    this.skip();
    const start = this.pos;
    let tag = false;
    if (allowTag && this.peek() === "#") {
      tag = true;
      this.pos++;
    }
    while (
      this.pos < this.end &&
      /[a-z0-9_.\-:/]/.test(this.text[this.pos])
    ) {
      this.pos++;
    }
    const raw = this.text.slice(start, this.pos);
    const value = tag ? raw.slice(1) : raw;
    if (value.length === 0) this.error("Expected a resource location", start);
    if (!/^[a-z0-9_.-]+(?::[a-z0-9_./-]+)?$/.test(value)) {
      this.error(
        `Invalid resource location '${value}'. Use namespace:path with lowercase letters, digits, '.', '_', '-' and '/'.`,
        start,
        "resource-location",
      );
    }
    const normalized = value.includes(":") ? value : `minecraft:${value}`;
    return tag ? `#${normalized}` : normalized;
  }

  readInt(): { value: number; start: number } {
    this.skip();
    const start = this.pos;
    if (this.peek() === "+") this.pos++;
    while (/[0-9-]/.test(this.peek()) && this.peek() !== "") this.pos++;
    const raw = this.text.slice(start, this.pos);
    if (!/^[+-]?\d+$/.test(raw)) {
      this.error(`Expected an integer, got '${raw || this.peek()}'`, start, "integer");
    }
    return { value: Number(raw), start };
  }

  readNumber(): { value: number; start: number } {
    this.skip();
    const start = this.pos;
    const match = /^[+-]?(?:\d+\.?\d*|\.\d+)(?:[eE][+-]?\d+)?f?d?/.exec(
      this.text.slice(this.pos, this.end),
    );
    if (!match) this.error("Expected a number", start, "number");
    this.pos += match[0].length;
    return { value: Number(match[0].replace(/[fd]$/, "")), start };
  }

  readRest(): string {
    const raw = this.text.slice(this.pos, this.end).trim();
    this.pos = this.end;
    return raw;
  }

  /** Read a position from three coordinate tokens. */
  readVec3(origin: Vec3, rotation: Vec2 = { yaw: 0, pitch: 0 }): Vec3 {
    const tokens: string[] = [];
    const starts: number[] = [];
    for (let i = 0; i < 3; i++) {
      this.skip();
      starts.push(this.pos);
      tokens.push(this.readWord());
    }
    const local = tokens.some((t) => t.startsWith("^"));
    const values = tokens.map((token, axis) => {
      if (token.startsWith("^")) {
        const rest = token.slice(1);
        return { kind: "local" as const, value: rest === "" ? 0 : Number(rest) };
      }
      if (token.startsWith("~")) {
        const rest = token.slice(1);
        const base = axis === 0 ? origin.x : axis === 1 ? origin.y : origin.z;
        return { kind: "relative" as const, value: rest === "" ? 0 : Number(rest) };
      }
      return { kind: "absolute" as const, value: Number(token) };
    });
    for (const [i, value] of values.entries()) {
      if (!Number.isFinite(value.value)) {
        this.error(`Invalid coordinate '${tokens[i]}'`, starts[i], "coordinates");
      }
    }
    if (local) {
      const yaw = (rotation.yaw * Math.PI) / 180;
      const pitch = (rotation.pitch * Math.PI) / 180;
      const left = { x: Math.cos(yaw), y: 0, z: Math.sin(yaw) };
      const forward = {
        x: -Math.sin(yaw) * Math.cos(pitch),
        y: -Math.sin(pitch),
        z: Math.cos(yaw) * Math.cos(pitch),
      };
      const up = { x: 0, y: 1, z: 0 };
      const lx = values[0].kind === "local" ? values[0].value : 0;
      const ly = values[1].kind === "local" ? values[1].value : 0;
      const lz = values[2].kind === "local" ? values[2].value : 0;
      return {
        x: origin.x + lx * left.x + ly * up.x + lz * forward.x,
        y: origin.y + lx * left.y + ly * up.y + lz * forward.y,
        z: origin.z + lx * left.z + ly * up.z + lz * forward.z,
      };
    }
    return {
      x: values[0].kind === "absolute" ? values[0].value : origin.x + values[0].value,
      y: values[1].kind === "absolute" ? values[1].value : origin.y + values[1].value,
      z: values[2].kind === "absolute" ? values[2].value : origin.z + values[2].value,
    };
  }

  readVec2(): Vec2 {
    const yaw = this.readNumber();
    const pitch = this.readNumber();
    return { yaw: yaw.value, pitch: pitch.value };
  }

  expectWord(expected: string): void {
    const start = this.pos;
    const word = this.readWord();
    if (word !== expected) {
      this.error(`Expected '${expected}', got '${word}'`, start, "keyword");
    }
  }

  peekWord(): string {
    const save = this.pos;
    try {
      return this.readWord();
    } catch {
      return "";
    } finally {
      this.pos = save;
    }
  }
}
