import { isCompound, isList, type NbtTag, numericValue, stringifySnbt, Tag } from "./nbt";
import { parseTextComponent } from "./snbt";

/**
 * Text component rendering. Minecraft's component tree is deep; cmdlib renders
 * the subset that matters for titles, chat and scoreboards: text, translate,
 * score, selector, nbt, keybind, extra, and full style inheritance.
 */

export const MINECRAFT_COLORS: Record<string, string> = {
  black: "#000000",
  dark_blue: "#0000AA",
  dark_green: "#00AA00",
  dark_aqua: "#00AAAA",
  dark_red: "#AA0000",
  dark_purple: "#AA00AA",
  gold: "#FFAA00",
  gray: "#AAAAAA",
  dark_gray: "#555555",
  blue: "#5555FF",
  green: "#55FF55",
  aqua: "#55FFFF",
  red: "#FF5555",
  light_purple: "#FF55FF",
  yellow: "#FFFF55",
  white: "#FFFFFF",
};

export interface TextStyle {
  color?: string;
  bold?: boolean;
  italic?: boolean;
  underlined?: boolean;
  strikethrough?: boolean;
  obfuscated?: boolean;
}

export interface TextRun extends TextStyle {
  text: string;
}

export interface TextRenderContext {
  resolveScoreHolder?: (name: string) => string;
  resolveScore?: (holder: string, objective: string) => number | null;
  resolveSelector?: (selector: string) => string[];
  resolveSelectorComponents?: (selector: string) => NbtTag[];
  resolveEntityNbt?: (selector: string, path: string) => string | null;
  resolveNbtTag?: (
    source: "storage" | "entity" | "block",
    id: string,
    path: string,
  ) => NbtTag | null;
  resolveNbtTags?: (
    source: "storage" | "entity" | "block",
    id: string,
    path: string,
  ) => NbtTag[];
  resolveKeybind?: (key: string) => string;
  resolveTranslation?: (key: string, fallback: string | null) => string | null;
}

const DEFAULT_STYLE: TextStyle = { color: "#FFFFFF" };

function readBoolean(compound: NbtTag, key: string): boolean | undefined {
  if (!isCompound(compound)) return undefined;
  const tag = compound.value[key];
  if (!tag) return undefined;
  if (tag.type === "byte" || tag.type === "short" || tag.type === "int" || tag.type === "long") {
    return tag.value !== 0;
  }
  return undefined;
}

function readString(compound: NbtTag, key: string): string | undefined {
  if (!isCompound(compound)) return undefined;
  const tag = compound.value[key];
  return tag && tag.type === "string" ? tag.value : undefined;
}

function styleFrom(component: NbtTag, parent: TextStyle): TextStyle {
  if (!isCompound(component)) return parent;
  const style: TextStyle = { ...parent };
  const color = readString(component, "color");
  if (color) {
    if (color.startsWith("#")) style.color = color;
    else style.color = MINECRAFT_COLORS[color] ?? parent.color;
  }
  for (const key of [
    "bold",
    "italic",
    "underlined",
    "strikethrough",
    "obfuscated",
  ] as const) {
    const value = readBoolean(component, key);
    if (value !== undefined) style[key] = value;
  }
  return style;
}

function componentText(component: NbtTag, ctx: TextRenderContext): string {
  if (component.type === "string") return component.value;
  if (component.type === "list") {
    return component.value
      .map((c) => plainText(c, ctx))
      .join("");
  }
  if (!isCompound(component)) return "";
  const text = readString(component, "text");
  if (text !== undefined) return text;

  const translate = readString(component, "translate");
  if (translate) {
    const fallback = readString(component, "fallback");
    const resolved = ctx.resolveTranslation?.(translate, fallback ?? null);
    if (resolved !== null && resolved !== undefined) return resolved;
    const withTags = component.value["with"];
    if (isList(withTags) && withTags.value.length > 0) {
      const args = withTags.value.map((t) => plainText(t, ctx));
      return translate.replace(/%(\d+\$)?s/g, () => args.shift() ?? "");
    }
    return fallback ?? translate;
  }

  const score = component.value["score"];
  if (isCompound(score)) {
    const holder = readString(score, "name") ?? "";
    const objective = readString(score, "objective") ?? "";
    const resolvedHolder = ctx.resolveScoreHolder?.(holder) ?? holder;
    const value = ctx.resolveScore?.(resolvedHolder, objective);
    return value === null || value === undefined ? "" : String(value);
  }

  const selector = readString(component, "selector");
  if (selector) {
    const separators = readString(component, "separator") ?? ", ";
    return (ctx.resolveSelector?.(selector) ?? []).join(separators);
  }

  const nbtValue = component.value["nbt"];
  if (typeof nbtValue?.value === "string") {
    const storage = readString(component, "storage");
    const entity = readString(component, "entity");
    const block = readString(component, "block");
    const source = storage ? "storage" : entity ? "entity" : block ? "block" : "storage";
    const id = storage ?? entity ?? block ?? "";
    const tag = ctx.resolveNbtTag?.(source, id, nbtValue.value);
    if (tag) {
      if (readBoolean(component, "interpret")) return plainText(tag, ctx);
      return formatTagForComponent(tag);
    }
    return "";
  }
  if (isCompound(nbtValue)) {
    const source = readString(nbtValue, "source") ?? "storage";
    const path = readString(nbtValue, "path") ?? "";
    return ctx.resolveEntityNbt?.(source, path) ?? "";
  }

  const keybind = readString(component, "keybind");
  if (keybind) return ctx.resolveKeybind?.(keybind) ?? keybind;

  return "";
}

function collect(
  component: NbtTag,
  ctx: TextRenderContext,
  parentStyle: TextStyle,
  out: TextRun[],
): void {
  if (isList(component)) {
    for (const child of component.value) collect(child, ctx, parentStyle, out);
    return;
  }
  if (component.type === "string") {
    if (component.value.length > 0) out.push({ text: component.value, ...parentStyle });
    return;
  }
  if (!isCompound(component)) return;
  const style = styleFrom(component, parentStyle);
  const nbtPath = readString(component, "nbt");
  if (nbtPath !== undefined && ctx.resolveNbtTags) {
    const storage = readString(component, "storage");
    const entity = readString(component, "entity");
    const block = readString(component, "block");
    const source = storage ? "storage" : entity ? "entity" : block ? "block" : "storage";
    const id = storage ?? entity ?? block ?? "";
    const tags = ctx.resolveNbtTags(source, id, nbtPath);
    if (tags.length > 0) {
      const interpret = readBoolean(component, "interpret") ?? false;
      const separator = component.value["separator"] ?? Tag.compound({
        text: Tag.string(", "),
        color: Tag.string("gray"),
      });
      tags.forEach((tag, index) => {
        if (index > 0) collect(separator, ctx, style, out);
        if (interpret) {
          collect(
            tag.type === "string" ? parseTextComponent(tag.value) : tag,
            ctx,
            style,
            out,
          );
        } else {
          out.push({ text: formatTagForComponent(tag), ...style });
        }
      });
      const nbtExtra = component.value["extra"];
      if (isList(nbtExtra)) {
        for (const child of nbtExtra.value) collect(child, ctx, style, out);
      }
      return;
    }
  }
  const selectorText = readString(component, "selector");
  if (selectorText && ctx.resolveSelectorComponents) {
    const selected = ctx.resolveSelectorComponents(selectorText);
    const separator = component.value["separator"] ?? Tag.compound({
      text: Tag.string(", "),
      color: Tag.string("gray"),
    });
    selected.forEach((child, index) => {
      if (index > 0) collect(separator, ctx, style, out);
      collect(child, ctx, style, out);
    });
    const selectorExtra = component.value["extra"];
    if (isList(selectorExtra)) {
      for (const child of selectorExtra.value) collect(child, ctx, style, out);
    }
    return;
  }
  const main = componentText(component, ctx);
  if (main.length > 0) out.push({ text: main, ...style });
  const extra = component.value["extra"];
  if (isList(extra)) {
    for (const child of extra.value) collect(child, ctx, style, out);
  }
}

export function renderTextComponent(
  component: NbtTag,
  ctx: TextRenderContext = {},
): TextRun[] {
  const out: TextRun[] = [];
  collect(component, ctx, DEFAULT_STYLE, out);
  return mergeRuns(out);
}

export function plainText(component: NbtTag, ctx: TextRenderContext = {}): string {
  return renderTextComponent(component, ctx)
    .map((r) => r.text)
    .join("");
}

export function mergeRuns(runs: TextRun[]): TextRun[] {
  const out: TextRun[] = [];
  for (const run of runs) {
    if (run.text.length === 0) continue;
    const previous = out[out.length - 1];
    if (
      previous &&
      previous.color === run.color &&
      previous.bold === run.bold &&
      previous.italic === run.italic &&
      previous.underlined === run.underlined &&
      previous.strikethrough === run.strikethrough &&
      previous.obfuscated === run.obfuscated
    ) {
      previous.text += run.text;
    } else {
      out.push({ ...run });
    }
  }
  return out;
}

/** Convert a component to the legacy section-sign string, for logs/tests. */
export function componentToLegacy(component: NbtTag, ctx: TextRenderContext = {}): string {
  const runs = renderTextComponent(component, ctx);
  const codes: Record<string, string> = {
    "#000000": "0",
    "#0000AA": "1",
    "#00AA00": "2",
    "#00AAAA": "3",
    "#AA0000": "4",
    "#AA00AA": "5",
    "#FFAA00": "6",
    "#AAAAAA": "7",
    "#555555": "8",
    "#5555FF": "9",
    "#55FF55": "a",
    "#55FFFF": "b",
    "#FF5555": "c",
    "#FF55FF": "d",
    "#FFFF55": "e",
    "#FFFFFF": "f",
  };
  return runs
    .map((run) => {
      let prefix = "";
      if (run.color && codes[run.color]) prefix += `§${codes[run.color]}`;
      if (run.bold) prefix += "§l";
      if (run.italic) prefix += "§o";
      if (run.underlined) prefix += "§n";
      if (run.strikethrough) prefix += "§m";
      if (run.obfuscated) prefix += "§k";
      return prefix + run.text;
    })
    .join("");
}

/** NBT preview helper used by the `nbt` component source. */
export function formatTagForComponent(tag: NbtTag): string {
  const value = numericValue(tag);
  if (value !== null) return String(value);
  return stringifySnbt(tag, { typed: false, quoteKeys: true });
}
