/**
 * Minecraft tag kinds.
 *
 * A datapack can put JSON tag files under `data/<namespace>/tags/<folder>/`.
 * Until 1.21 the folder names were plural (`tags/items`, `tags/blocks`, …);
 * 1.21+ uses singular names (`tags/item`, `tags/block`, …). This module is the
 * single source of truth for those folders, the value syntax, and the metadata
 * the editor/linter needs to offer tag kinds that are not functions.
 *
 * Everything here is pure data + string helpers. `parseDatapack` and the
 * linter consume it; the Explorer can build creation menus from
 * `TAG_KIND_META` and `tagPathFor` without duplicating folder rules.
 */

import type { VersionProfile } from "./version";

/** Tag registries cmdlib understands today. */
export const TAG_KINDS = [
  "function",
  "item",
  "block",
  "entity_type",
  "fluid",
  "game_event",
] as const;

export type TagKind = (typeof TAG_KINDS)[number];

/** True when a string is one of the supported tag kinds. */
export function isTagKind(value: string): value is TagKind {
  return (TAG_KINDS as readonly string[]).includes(value);
}

export interface TagKindMeta {
  kind: TagKind;
  /** Human label for menus and diagnostics, e.g. "Item". */
  label: string;
  /** Plural label, e.g. "Items". */
  pluralLabel: string;
  /** Folder name under `tags/` on 1.21+ (singular). */
  singularFolder: string;
  /** Folder name under `tags/` on 1.20.4 and older (plural). */
  pluralFolder: string;
  /** One sentence shown in the editor when choosing a tag kind. */
  description: string;
  /** Registry kind the tag's values point at. */
  valueKind: TagKind;
  /** Lower-case value label used in validation messages. */
  valueLabel: string;
  /** A few realistic values for completions and empty-state hints. */
  examples: string[];
  /**
   * Function tags point at `.mcfunction` files. Other kinds point at registry
   * entries (items, blocks, entity types, fluids, game events), so a missing
   * function must never be reported for them.
   */
  requiresFunctionTarget: boolean;
}

export const TAG_KIND_META: Record<TagKind, TagKindMeta> = {
  function: {
    kind: "function",
    label: "Function",
    pluralLabel: "Functions",
    singularFolder: "function",
    pluralFolder: "functions",
    description: "Runs a list of functions, e.g. the load and tick entry points.",
    valueKind: "function",
    valueLabel: "function",
    examples: ["cmdlib:load", "#minecraft:tick"],
    requiresFunctionTarget: true,
  },
  item: {
    kind: "item",
    label: "Item",
    pluralLabel: "Items",
    singularFolder: "item",
    pluralFolder: "items",
    description: "Groups items so recipes, loot tables and commands can target them together.",
    valueKind: "item",
    valueLabel: "item",
    examples: ["minecraft:diamond", "#minecraft:planks"],
    requiresFunctionTarget: false,
  },
  block: {
    kind: "block",
    label: "Block",
    pluralLabel: "Blocks",
    singularFolder: "block",
    pluralFolder: "blocks",
    description: "Groups blocks for recipes, loot tables, worldgen and selectors.",
    valueKind: "block",
    valueLabel: "block",
    examples: ["minecraft:stone", "#minecraft:logs"],
    requiresFunctionTarget: false,
  },
  entity_type: {
    kind: "entity_type",
    label: "Entity type",
    pluralLabel: "Entity types",
    singularFolder: "entity_type",
    pluralFolder: "entity_types",
    description: "Groups entity types for predicates, advancements and selectors.",
    valueKind: "entity_type",
    valueLabel: "entity type",
    examples: ["minecraft:zombie", "#minecraft:skeletons"],
    requiresFunctionTarget: false,
  },
  fluid: {
    kind: "fluid",
    label: "Fluid",
    pluralLabel: "Fluids",
    singularFolder: "fluid",
    pluralFolder: "fluids",
    description: "Groups fluids for predicates, block interactions and worldgen.",
    valueKind: "fluid",
    valueLabel: "fluid",
    examples: ["minecraft:water", "#minecraft:lava"],
    requiresFunctionTarget: false,
  },
  game_event: {
    kind: "game_event",
    label: "Game event",
    pluralLabel: "Game events",
    singularFolder: "game_event",
    pluralFolder: "game_events",
    description: "Groups sculk/game events for vibration listeners and predicates.",
    valueKind: "game_event",
    valueLabel: "game event",
    examples: ["minecraft:block_break", "#minecraft:vibrations"],
    requiresFunctionTarget: false,
  },
};

export function tagKindMeta(kind: TagKind): TagKindMeta {
  return TAG_KIND_META[kind];
}

const NAMESPACE_RE = /^[a-z0-9_.-]+$/;
const PATH_RE = /^[a-z0-9_./-]+$/;

/** Minecraft namespaces are lower-case `[a-z0-9_.-]+`. */
export function isValidNamespace(namespace: string): boolean {
  return NAMESPACE_RE.test(namespace);
}

/** Minecraft resource paths are lower-case `[a-z0-9_./-]+` without empty segments. */
export function isValidResourcePath(path: string): boolean {
  return (
    PATH_RE.test(path) &&
    !path.startsWith("/") &&
    !path.endsWith("/") &&
    !path.includes("//")
  );
}

/**
 * Normalise a bare or namespaced resource location.
 * Bare values use `minecraft:`, matching Minecraft's JSON parsing.
 * Returns `null` when the value is not a valid resource location.
 */
export function normalizeResourceLocation(value: string): string | null {
  if (typeof value !== "string" || value.length === 0) return null;
  const parts = value.split(":");
  if (parts.length > 2) return null;
  if (parts.length === 1) {
    return isValidResourcePath(parts[0]) ? `minecraft:${parts[0]}` : null;
  }
  const [namespace, path] = parts;
  if (!isValidNamespace(namespace) || !isValidResourcePath(path)) return null;
  return `${namespace}:${path}`;
}

/**
 * Folder for a tag kind, e.g. `tags/item` on 26.3 and `tags/items` on 1.20.4.
 * `version.tagFolders` is authoritative; the singular/plural metadata is a
 * fallback for hand-built VersionProfile-like objects.
 */
export function tagFolderFor(kind: TagKind, version: VersionProfile): string {
  const mapped = version.tagFolders?.[kind];
  if (mapped) return mapped;
  const meta = TAG_KIND_META[kind];
  const singular = version.functionFolder === "function";
  return `tags/${singular ? meta.singularFolder : meta.pluralFolder}`;
}

/**
 * Full source path for a tag file.
 *
 * `tagPathFor("item", "mypack", "ores/iron", getVersion("26.3"))`
 * → `data/mypack/tags/item/ores/iron.json`.
 */
export function tagPathFor(
  kind: TagKind,
  namespace: string,
  name: string,
  version: VersionProfile,
): string {
  const cleanNamespace = namespace.trim().replace(/:$/, "") || "minecraft";
  const cleanName = name
    .trim()
    .replace(/^#/, "")
    .replace(/^\/+/, "")
    .replace(/\/+$/, "");
  return `data/${cleanNamespace}/${tagFolderFor(kind, version)}/${cleanName}.json`;
}

const TAG_PATH_RE = /^data\/([^/]+)\/tags\/([^/]+)\/(.+)\.json$/;

export interface TagPathInfo {
  kind: TagKind;
  namespace: string;
  /** Tag name/path without the `.json` suffix, e.g. `ores/iron`. */
  name: string;
  /** Folder including the `tags/` prefix, e.g. `tags/item`. */
  folder: string;
}

/**
 * Parse a datapack path into a tag kind, but only when the folder is the one
 * this version expects. Legacy tag folders return `null` on new versions so
 * `parseDatapack` can report a migration error instead of silently accepting
 * the wrong layout.
 */
export function parseTagPath(path: string, version: VersionProfile): TagPathInfo | null {
  const match = TAG_PATH_RE.exec(path);
  if (!match) return null;
  const [, namespace, folderName, name] = match;
  const folder = `tags/${folderName}`;
  const kind = TAG_KINDS.find((candidate) => tagFolderFor(candidate, version) === folder);
  if (!kind) return null;
  return { kind, namespace, name, folder };
}

/** Detect the tag kind for a path on a given version, or `null`. */
export function detectTagKind(path: string, version: VersionProfile): TagKind | null {
  return parseTagPath(path, version)?.kind ?? null;
}

/** Find the kind whose singular or plural folder name matches, if any. */
export function tagKindForFolderName(folderName: string): TagKind | null {
  return (
    TAG_KINDS.find((kind) => {
      const meta = TAG_KIND_META[kind];
      return folderName === meta.singularFolder || folderName === meta.pluralFolder;
    }) ?? null
  );
}

/**
 * Result of validating one `values` entry.
 *
 * `id` is the normalised resource location without a leading `#`; `reference`
 * records whether the source value started with `#`. On failure `message` is
 * set and `id`/`namespace`/`path` are empty strings.
 */
export interface TagValueInfo {
  ok: boolean;
  message?: string;
  /** Original textual id (without any `{ id, required }` wrapper). */
  raw: string;
  /** Normalised `namespace:path`, no leading `#`. Empty on failure. */
  id: string;
  namespace: string;
  path: string;
  /** True for `#namespace:tag` references. */
  reference: boolean;
  /** Defaults to `true`; `{ id, required: false }` sets this to false. */
  required: boolean;
}

function invalidValue(
  raw: string,
  message: string,
  reference: boolean,
  required = true,
): TagValueInfo {
  return {
    ok: false,
    message,
    raw,
    id: "",
    namespace: "",
    path: "",
    reference,
    required,
  };
}

/**
 * Validate one tag value (string form) for a kind. This checks resource
 * location syntax and namespace only — registry membership is not knowable
 * from a datapack, and function existence is the linter's job.
 */
export function validateTagValue(kind: TagKind, value: string): TagValueInfo {
  const meta = TAG_KIND_META[kind];
  if (typeof value !== "string") {
    return invalidValue(
      String(value),
      `A ${meta.label.toLowerCase()} tag value must be a string id or reference.`,
      false,
    );
  }
  const raw = value;
  const reference = raw.startsWith("#");
  const bare = reference ? raw.slice(1) : raw;
  const what = reference ? `tag reference '${raw}'` : `${meta.valueLabel} id '${raw}'`;
  if (bare.length === 0) {
    return invalidValue(raw, `Empty ${meta.valueLabel} id. Expected namespace:path${reference ? " after '#'" : ""}.`, reference);
  }
  const normalized = normalizeResourceLocation(bare);
  if (!normalized) {
    return invalidValue(
      raw,
      `Invalid ${what}. Use namespace:path with lower-case letters, digits, '.', '_', '-' and '/'.`,
      reference,
    );
  }
  const [namespace, path] = normalized.split(":") as [string, string];
  return {
    ok: true,
    raw,
    id: normalized,
    namespace,
    path,
    reference,
    required: true,
  };
}

/** Convenience wrapper: the validation error message, or `null` when valid. */
export function tagValueError(kind: TagKind, value: string): string | null {
  const info = validateTagValue(kind, value);
  return info.ok ? null : (info.message ?? "Invalid tag value.");
}

/**
 * Parse a `values` element: either a string id/reference or the optional
 * object form `{ "id": "namespace:path", "required": false }`.
 *
 * `required` defaults to true; a non-boolean `required` is an error rather
 * than being silently coerced.
 */
export function parseTagEntry(kind: TagKind, entry: unknown): TagValueInfo {
  if (typeof entry === "string") return validateTagValue(kind, entry);
  const meta = TAG_KIND_META[kind];
  if (!entry || typeof entry !== "object" || Array.isArray(entry)) {
    return invalidValue(
      String(entry),
      `Tag value must be a string or an object with an 'id' (${meta.label.toLowerCase()} tag).`,
      false,
    );
  }
  const record = entry as Record<string, unknown>;
  if (typeof record.id !== "string") {
    return invalidValue(
      "",
      `Tag entry must have a string 'id' (${meta.label.toLowerCase()} tag).`,
      false,
    );
  }
  const info = validateTagValue(kind, record.id);
  if (!info.ok) return info;
  if (record.required !== undefined && typeof record.required !== "boolean") {
    return invalidValue(
      record.id,
      `Tag entry '${record.id}' has a non-boolean 'required'; use true or false.`,
      info.reference,
    );
  }
  return {
    ...info,
    raw: record.id,
    required: record.required !== false,
  };
}
