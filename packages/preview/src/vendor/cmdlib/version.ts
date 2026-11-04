/**
 * Version profiles.
 *
 * cmdlib is version-aware on purpose. 26.3 is the headline target, but every
 * path, folder name and pack metadata rule is driven by this table, so a new
 * game drop is a new row rather than a rewrite. Versions older than 1.21 use
 * singular/plural folder changes, and versions before 1.21.9 use pack_format
 * instead of min_format/max_format — those differences are exposed as
 * capability flags, not scattered `if` statements.
 */

import type { TagKind } from "./tag-kinds";
import { TAG_KINDS, tagFolderFor } from "./tag-kinds";

/** Folder of every supported tag kind for a version, e.g. `tags/item`. */
export interface TagFolderMap extends Record<TagKind, string> {
  function: "tags/functions" | "tags/function";
}

export interface VersionProfile {
  id: string;
  label: string;
  release: string;
  packFormat: [number, number];
  dataVersion?: number;
  protocol?: number;
  functionFolder: "functions" | "function";
  /** Function tag folder. Kept for older call sites; mirrors `tagFolders.function`. */
  tagFolder: "tags/functions" | "tags/function";
  /** Tag folder for every supported kind (function, item, block, …). */
  tagFolders: TagFolderMap;
  packMeta: "legacy" | "minmax";
  macros: boolean;
  textComponents: "json" | "snbt";
  dataStorage: boolean;
  returnCommand: boolean;
  displayNumberFormat: boolean;
  /** Docs link shown in the version picker. */
  notes?: string;
}

function tagFoldersFor(packFormat: number): TagFolderMap {
  const singular = packFormat >= 48;
  if (singular) {
    return {
      function: "tags/function",
      item: "tags/item",
      block: "tags/block",
      entity_type: "tags/entity_type",
      fluid: "tags/fluid",
      game_event: "tags/game_event",
    };
  }
  return {
    function: "tags/functions",
    item: "tags/items",
    block: "tags/blocks",
    entity_type: "tags/entity_types",
    fluid: "tags/fluids",
    game_event: "tags/game_events",
  };
}

const profile = (
  id: string,
  label: string,
  release: string,
  packFormat: [number, number],
  opts: Partial<VersionProfile> = {},
): VersionProfile => {
  const tagFolders = opts.tagFolders ?? tagFoldersFor(packFormat[0]);
  return {
    id,
    label,
    release,
    packFormat,
    functionFolder: packFormat[0] >= 48 ? "function" : "functions",
    tagFolder: tagFolders.function,
    tagFolders,
    packMeta: packFormat[0] >= 82 ? "minmax" : "legacy",
    macros: packFormat[0] >= 18,
    textComponents: packFormat[0] >= 71 ? "snbt" : "json",
    dataStorage: true,
    returnCommand: packFormat[0] >= 18,
    displayNumberFormat: packFormat[0] >= 26,
    ...opts,
  };
};

export const VERSIONS: VersionProfile[] = [
  profile("26.3", "26.3 · Wilderness Bound", "2026-09-15", [121, 0], {
    dataVersion: 5023,
    protocol: 777,
    notes:
      "Adds /compute, /posteffect, data modify … compute, and execute if slots. cmdlib targets this release.",
  }),
  profile("1.21.10", "1.21.9 – 1.21.10", "2025-10-07", [88, 0], {}),
  profile("1.21.8", "1.21.8", "2025-07-17", [81, 0], {}),
  profile("1.21.7", "1.21.7", "2025-06-30", [81, 0], {}),
  profile("1.21.6", "1.21.6", "2025-06-17", [80, 0], {}),
  profile("1.21.5", "1.21.5", "2025-03-25", [71, 0], {}),
  profile("1.21.4", "1.21.4", "2024-12-03", [61, 0], { dataVersion: 4189 }),
  profile("1.21.2", "1.21.2 – 1.21.3", "2024-10-22", [57, 0], { dataVersion: 4082 }),
  profile("1.21", "1.21 – 1.21.1", "2024-06-13", [48, 0], { dataVersion: 3953 }),
  profile("1.20.5", "1.20.5 – 1.20.6", "2024-04-23", [41, 0], {}),
  profile("1.20.4", "1.20.3 – 1.20.4", "2023-12-05", [26, 0], { dataVersion: 3700 }),
  profile("1.20.2", "1.20.2", "2023-09-20", [18, 0], { dataVersion: 3578 }),
  profile("1.20", "1.20 – 1.20.1", "2023-06-07", [15, 0], { dataVersion: 3463 }),
  profile("1.19.4", "1.19.4", "2023-03-14", [12, 0], { dataVersion: 3337 }),
  profile("1.19", "1.19 – 1.19.3", "2022-06-07", [10, 0], {}),
  profile("1.18.2", "1.18.2", "2022-02-28", [9, 0], {}),
  profile("1.18", "1.18 – 1.18.1", "2021-11-30", [8, 0], {}),
  profile("1.17", "1.17 – 1.17.1", "2021-06-08", [7, 0], {}),
  profile("1.16.5", "1.16.2 – 1.16.5", "2020-08-11", [6, 0], {}),
  profile("1.16", "1.16 – 1.16.1", "2020-06-23", [5, 0], {}),
  profile("1.15", "1.15 – 1.15.2", "2019-12-10", [5, 0], {}),
  profile("1.14", "1.14 – 1.14.4", "2019-04-23", [4, 0], {}),
  profile("1.13", "1.13 – 1.13.2", "2018-07-18", [4, 0], {}),
];

export const DEFAULT_VERSION = "26.3";

export function getVersion(id: string): VersionProfile {
  return VERSIONS.find((v) => v.id === id) ?? VERSIONS.find((v) => v.id === DEFAULT_VERSION)!;
}

export function latestVersion(): VersionProfile {
  return getVersion(DEFAULT_VERSION);
}

export function formatPackFormat(format: [number, number]): string {
  return format[1] === 0 ? `${format[0]}.0` : `${format[0]}.${format[1]}`;
}

export function functionDir(version: VersionProfile): string {
  return version.functionFolder;
}

export function functionPath(id: string, version: VersionProfile): string {
  const [namespace, path] = splitId(id);
  return `data/${namespace}/${version.functionFolder}/${path}.mcfunction`;
}

export function tagPath(id: string, version: VersionProfile): string {
  const [namespace, path] = splitId(id);
  return `data/${namespace}/${version.tagFolder}/${path}.json`;
}

export function splitId(id: string): [string, string] {
  const index = id.indexOf(":");
  if (index === -1) return ["minecraft", id];
  return [id.slice(0, index), id.slice(index + 1)];
}

export function buildPackMeta(version: VersionProfile, description: string): string {
  if (version.packMeta === "legacy") {
    return JSON.stringify(
      {
        pack: {
          pack_format: version.packFormat[0],
          description,
        },
      },
      null,
      2,
    );
  }
  return JSON.stringify(
    {
      pack: {
        description,
        min_format: [version.packFormat[0], version.packFormat[1]],
        max_format: [version.packFormat[0], version.packFormat[1]],
      },
    },
    null,
    2,
  );
}

/** Migrate every folder-level break between two versions. */
export function migrateFiles(
  files: Record<string, string>,
  from: VersionProfile,
  to: VersionProfile,
  packDescription: string,
): Record<string, string> {
  const out: Record<string, string> = {};
  for (const [rawPath, content] of Object.entries(files)) {
    // Old projects and hand-edited imports occasionally carry a leading file
    // separator. Normalise before matching so migration is idempotent.
    let next = rawPath.replace(/^\/+/, "");
    if (from.functionFolder !== to.functionFolder) {
      next = next.replace(
        new RegExp(`^data/([^/]+)/${from.functionFolder}/`),
        `data/$1/${to.functionFolder}/`,
      );
    }
    // Every tag kind moved from a plural folder to a singular one in 1.21.
    // Run all kinds (including functions) so item/block/entity_type/fluid/
    // game_event tags migrate alongside function tags.
    for (const kind of TAG_KINDS) {
      const fromFolder = tagFolderFor(kind, from);
      const toFolder = tagFolderFor(kind, to);
      if (fromFolder === toFolder) continue;
      next = next.replace(
        new RegExp(`^data/([^/]+)/${fromFolder}/`),
        `data/$1/${toFolder}/`,
      );
    }
    out[next] = content;
  }
  if ("pack.mcmeta" in out) {
    out["pack.mcmeta"] = buildPackMeta(to, packDescription);
  }
  return out;
}
