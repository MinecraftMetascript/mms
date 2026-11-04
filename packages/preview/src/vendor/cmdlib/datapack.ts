import type { FunctionDef, FunctionTag } from "./world";
import { type VersionProfile, getVersion, splitId } from "./version";
import {
  TAG_KINDS,
  TAG_KIND_META,
  isValidNamespace,
  isValidResourcePath,
  parseTagEntry,
  parseTagPath,
  tagFolderFor,
  tagKindForFolderName,
} from "./tag-kinds";
import type { TagKind, TagValueInfo } from "./tag-kinds";
import { newId } from "./ids";

export interface Project {
  id: string;
  name: string;
  versionId: string;
  files: Record<string, string>;
  activePath: string;
  updatedAt: number;
}

/** A parsed tag of any supported kind. Function tags are also exposed in `tags`. */
export interface ParsedTag {
  kind: TagKind;
  id: string;
  namespace: string;
  /** Tag name without the `.json` suffix, e.g. `ores/iron`. */
  name: string;
  path: string;
  /** Normalised ids; references keep their leading `#`. */
  values: string[];
  entries: TagValueInfo[];
}

export type DatapackErrorCode =
  | "invalid-folder"
  | "invalid-tag-id"
  | "invalid-tag-json"
  | "missing-tag-values"
  | "invalid-tag-value";

export interface DatapackError {
  path: string;
  message: string;
  code: DatapackErrorCode;
  /** For value errors, the index in the tag's `values` array. */
  valueIndex?: number;
}

export interface ParsedDatapack {
  functions: FunctionDef[];
  /** Function tags only, in the shape the runtime/command layer expects. */
  tags: FunctionTag[];
  /** Every supported tag kind, including function tags. */
  allTags: ParsedTag[];
  errors: DatapackError[];
}

const FUNCTION_RE = /^data\/([^/]+)\/([^/]+)\/(.+)\.mcfunction$/;
const TAG_FILE_RE = /^data\/([^/]+)\/tags\/([^/]+)\/(.+)\.json$/;

export function parseDatapack(
  files: Record<string, string>,
  version: VersionProfile,
): ParsedDatapack {
  const functions: FunctionDef[] = [];
  const tags: FunctionTag[] = [];
  const allTags: ParsedTag[] = [];
  const errors: ParsedDatapack["errors"] = [];

  for (const [path, content] of Object.entries(files)) {
    const fn = FUNCTION_RE.exec(path);
    if (fn) {
      const [, namespace, folder, rest] = fn;
      if (folder !== version.functionFolder) {
        errors.push({
          path,
          code: "invalid-folder",
          message: `Function folder '${folder}' is not valid for ${version.id}; expected '${version.functionFolder}'.`,
        });
        continue;
      }
      const lines = content.split(/\r?\n/);
      const macroKeys = new Set<string>();
      let hasMacroLines = false;
      for (const line of lines) {
        if (!line.startsWith("$")) continue;
        hasMacroLines = true;
        const regex = /\$\(([^)]*)\)/g;
        let match: RegExpExecArray | null;
        while ((match = regex.exec(line))) {
          if (match[1]) macroKeys.add(match[1]);
        }
      }
      functions.push({
        id: `${namespace}:${rest}`,
        path,
        lines,
        macroKeys: [...macroKeys],
        hasMacroLines,
      });
      continue;
    }
    if (!path.startsWith("data/")) continue;
    const tagFile = TAG_FILE_RE.exec(path);
    if (!tagFile) continue;
    const [, namespace, folderName, name] = tagFile;
    const folder = `tags/${folderName}`;
    const kind = TAG_KINDS.find((candidate) => tagFolderFor(candidate, version) === folder);
    if (!kind) {
      // Only report folders that belong to a tag kind cmdlib knows about.
      // Datapacks also contain vanilla tag registries we do not model yet
      // (painting_variant, damage_type, …); those must not become errors.
      const knownKind = tagKindForFolderName(folderName);
      if (knownKind) {
        errors.push({
          path,
          code: "invalid-folder",
          message: `Tag folder '${folder}' is not valid for ${version.id}; expected '${tagFolderFor(knownKind, version)}'.`,
        });
      }
      continue;
    }
    if (!isValidNamespace(namespace) || !isValidResourcePath(name)) {
      errors.push({
        path,
        code: "invalid-tag-id",
        message: `Invalid ${TAG_KIND_META[kind].label.toLowerCase()} tag id '${namespace}:${name}'. Use a lower-case namespace and path.`,
      });
      continue;
    }
    let parsed: unknown;
    try {
      parsed = JSON.parse(content);
    } catch (error) {
      errors.push({
        path,
        code: "invalid-tag-json",
        message: `Invalid tag JSON: ${(error as Error).message}`,
      });
      continue;
    }
    if (!parsed || typeof parsed !== "object" || Array.isArray(parsed)) {
      errors.push({
        path,
        code: "invalid-tag-json",
        message: "A tag file must contain a JSON object with a 'values' array.",
      });
      continue;
    }
    const rawValues = (parsed as { values?: unknown }).values;
    if (!Array.isArray(rawValues)) {
      errors.push({
        path,
        code: "missing-tag-values",
        message: `${TAG_KIND_META[kind].label} tag is missing a 'values' array.`,
      });
      continue;
    }

    const entries: TagValueInfo[] = [];
    rawValues.forEach((rawEntry, index) => {
      const entry = parseTagEntry(kind, rawEntry);
      if (!entry.ok) {
        errors.push({
          path,
          code: "invalid-tag-value",
          message: entry.message ?? `Invalid ${TAG_KIND_META[kind].valueLabel} value.`,
          valueIndex: index,
        });
        return;
      }
      entries.push(entry);
    });

    const values = entries.map((entry) =>
      entry.reference ? `#${entry.id}` : entry.id,
    );
    const id = `${namespace}:${name}`;
    allTags.push({
      kind,
      id,
      namespace,
      name,
      path,
      values,
      entries,
    });
    if (kind === "function") {
      tags.push({ id, path, values });
    }
  }

  return { functions, tags, allTags, errors };
}

export function functionFiles(files: Record<string, string>, version: VersionProfile): string[] {
  return Object.keys(files).filter(
    (path) =>
      path.startsWith("data/") &&
      path.includes(`/${version.functionFolder}/`) &&
      path.endsWith(".mcfunction"),
  );
}

export function listFunctionIds(
  files: Record<string, string>,
  version: VersionProfile,
): Array<{ id: string; path: string; hasMacros: boolean }> {
  return functionFiles(files, version).map((path) => {
    const fn = FUNCTION_RE.exec(path)!;
    return {
      id: `${fn[1]}:${fn[3]}`,
      path,
      hasMacros: files[path].split(/\r?\n/).some((line) => line.startsWith("$")),
    };
  });
}

export function listFunctionTags(
  files: Record<string, string>,
  version: VersionProfile,
): Array<{ id: string; path: string; values: string[] }> {
  const out: Array<{ id: string; path: string; values: string[] }> = [];
  for (const [path, content] of Object.entries(files)) {
    if (!path.startsWith("data/") || !path.endsWith(".json")) continue;
    const info = parseTagPath(path, version);
    if (!info || info.kind !== "function") continue;
    try {
      const parsed = JSON.parse(content) as { values?: unknown };
      const values = Array.isArray(parsed.values)
        ? parsed.values
            .map((entry) => parseTagEntry("function", entry))
            .filter((entry) => entry.ok)
            .map((entry) => (entry.reference ? `#${entry.id}` : entry.id))
        : [];
      out.push({ id: `${info.namespace}:${info.name}`, path, values });
    } catch {
      // The linter surfaces malformed JSON separately.
    }
  }
  return out;
}

/** List every parsed tag of a kind (or all kinds) for Explorer/completion use. */
export function listTags(
  files: Record<string, string>,
  version: VersionProfile,
  kind?: TagKind,
): Array<{ kind: TagKind; id: string; path: string; values: string[] }> {
  return parseDatapack(files, version)
    .allTags.filter((tag) => !kind || tag.kind === kind)
    .map((tag) => ({ kind: tag.kind, id: tag.id, path: tag.path, values: tag.values }));
}

export function collectStorageIds(files: Record<string, string>): string[] {
  const ids = new Set<string>();
  const regex = /\b(?:data\s+(?:get|modify|merge|remove)\s+storage|if\s+data\s+storage|with\s+storage)\s+([a-z0-9_.-]+(?::[a-z0-9_./-]+)?)/g;
  for (const content of Object.values(files)) {
    let match: RegExpExecArray | null;
    while ((match = regex.exec(content))) {
      const id = match[1].includes(":") ? match[1] : `minecraft:${match[1]}`;
      ids.add(id);
    }
  }
  return [...ids].sort();
}

export function collectObjectiveIds(files: Record<string, string>): string[] {
  const ids = new Set<string>();
  const regex = /scoreboard\s+objectives\s+add\s+([A-Za-z0-9_.+-]+)/g;
  for (const content of Object.values(files)) {
    let match: RegExpExecArray | null;
    while ((match = regex.exec(content))) ids.add(match[1]);
  }
  return [...ids].sort();
}

export function collectScoreHolders(files: Record<string, string>): string[] {
  const ids = new Set<string>();
  const regex = /scoreboard\s+players\s+(?:set|add|remove|get|reset)\s+([A-Za-z0-9_.+-]+)/g;
  for (const content of Object.values(files)) {
    let match: RegExpExecArray | null;
    while ((match = regex.exec(content))) {
      if (!match[1].startsWith("@")) ids.add(match[1]);
    }
  }
  return [...ids].sort();
}

export function createStarterFiles(version: VersionProfile): Record<string, string> {
  const fn = version.functionFolder;
  const tag = version.tagFolder;
  const packMeta =
    version.packMeta === "legacy"
      ? JSON.stringify(
          { pack: { pack_format: version.packFormat[0], description: "cmdlib datapack" } },
          null,
          2,
        )
      : JSON.stringify(
          {
            pack: {
              description: "cmdlib datapack",
              min_format: [version.packFormat[0], version.packFormat[1]],
              max_format: [version.packFormat[0], version.packFormat[1]],
            },
          },
          null,
          2,
        );

  return {
    "pack.mcmeta": packMeta + "\n",
    [`data/cmdlib/${fn}/load.mcfunction`]: `# Runs once when the pack is loaded.
`,
    [`data/cmdlib/${fn}/tick.mcfunction`]: `# Runs once every game tick.
`,
    [`data/minecraft/${tag}/load.json`]: JSON.stringify(
      { values: ["cmdlib:load"] },
      null,
      2,
    ) + "\n",
    [`data/minecraft/${tag}/tick.json`]: JSON.stringify(
      { values: ["cmdlib:tick"] },
      null,
      2,
    ) + "\n",
  };
}

export function createProject(versionId = "26.3", name = "Untitled Pack"): Project {
  const version = getVersion(versionId);
  const files = createStarterFiles(version);
  return {
    id: newId("pack"),
    name,
    versionId,
    files,
    activePath: `data/cmdlib/${version.functionFolder}/load.mcfunction`,
    updatedAt: Date.now(),
  };
}

/** A small project used by the landing page live simulation. */
export function createLandingProject(versionId = "26.3"): Project {
  const version = getVersion(versionId);
  const fn = version.functionFolder;
  const tag = version.tagFolder;
  const files: Record<string, string> = {
    "pack.mcmeta": JSON.stringify(
      {
        pack: {
          description: "cmdlib landing demo",
          min_format: [version.packFormat[0], version.packFormat[1]],
          max_format: [version.packFormat[0], version.packFormat[1]],
        },
      },
      null,
      2,
    ) + "\n",
    [`data/landing/${fn}/load.mcfunction`]: `# The landing page runs this real load tag on boot.
scoreboard objectives add progress dummy {"text":"PROGRESS","color":"gold","bold":true}
scoreboard objectives setdisplay sidebar progress
scoreboard players set @s progress 0

data modify storage landing:state stage set value "boot"
data modify storage landing:state ticks set value 0

summon armor_stand 3 64 1 {CustomName:'{"text":"Beacon","color":"gold"}',Tags:["beacon"],Glowing:1b,NoGravity:1b}
summon zombie -3 64 2 {CustomName:'{"text":"Target","color":"red"}',Tags:["target"],NoAI:1b}
summon marker 0 64 -3 {CustomName:'{"text":"Anchor","color":"aqua"}',Tags:["anchor"]}

title @s times 5 55 10
tellraw @s [{"text":"world booted","color":"gray"}]
`,
    [`data/landing/${fn}/stamp.mcfunction`]: `# Macro function used by the scroll demo.
$data modify storage landing:state stage set value "$(stage)"
$tellraw @s [{"text":"macro → ","color":"gray"},{"text":"$(stage)","color":"aqua"}]
$title @s subtitle {"text":"stage: $(stage)","color":"gray"}
`,
    [`data/minecraft/${tag}/load.json`]: JSON.stringify({ values: ["landing:load"] }, null, 2) + "\n",
  };
  return {
    id: newId("pack"),
    name: "Landing",
    versionId,
    files,
    activePath: `data/landing/${fn}/load.mcfunction`,
    updatedAt: Date.now(),
  };
}

export function projectFunctionIds(project: Project): string[] {
  return listFunctionIds(project.files, getVersion(project.versionId)).map((f) => f.id);
}

export function displayNameForFile(path: string): string {
  const parts = path.split("/");
  return parts[parts.length - 1];
}

export function fileKind(path: string): "mcfunction" | "json" | "mcmeta" | "other" {
  if (path.endsWith(".mcfunction")) return "mcfunction";
  if (path.endsWith(".json")) return "json";
  if (path.endsWith("pack.mcmeta")) return "mcmeta";
  return "other";
}

export function splitPath(path: string): string[] {
  return path.split("/").filter(Boolean);
}

export function functionIdFromPath(
  path: string,
  version: VersionProfile,
): string | null {
  const match = FUNCTION_RE.exec(path);
  if (!match) return null;
  if (match[2] !== version.functionFolder) return null;
  return `${match[1]}:${match[3]}`;
}

export function namespaced(id: string): string {
  const [namespace, path] = splitId(id);
  return `${namespace}:${path}`;
}
