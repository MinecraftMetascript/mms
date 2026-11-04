import { jsonToTag } from "./snbt";
import type { NbtTag } from "./nbt";

export interface AdvancementCriterion {
  name: string;
  trigger: string;
  conditions: Record<string, unknown> | null;
}

export interface AdvancementDisplay {
  title?: NbtTag;
  description?: NbtTag;
  icon?: string;
  frame?: string;
  background?: string;
  hidden?: boolean;
}

export interface AdvancementRewards {
  function?: string;
  experience?: number;
  loot?: string[];
  recipes?: string[];
}

export interface AdvancementDefinition {
  id: string;
  path: string;
  parent?: string;
  display?: AdvancementDisplay;
  criteria: AdvancementCriterion[];
  requirements?: string[][];
  rewards: AdvancementRewards;
  data: Record<string, unknown>;
}

export interface AdvancementError {
  path: string;
  message: string;
}

const ADVANCEMENT_PATH = /^data\/([^/]+)\/advancements?\/(.+)\.json$/;

function stringArray(value: unknown): string[] {
  if (typeof value === "string") return [value];
  return Array.isArray(value)
    ? value.filter((entry): entry is string => typeof entry === "string")
    : [];
}

export function parseAdvancementFiles(files: Record<string, string>): {
  advancements: Map<string, AdvancementDefinition>;
  errors: AdvancementError[];
} {
  const advancements = new Map<string, AdvancementDefinition>();
  const errors: AdvancementError[] = [];
  for (const [path, content] of Object.entries(files)) {
    const match = ADVANCEMENT_PATH.exec(path);
    if (!match) continue;
    const [, namespace, name] = match;
    let raw: unknown;
    try {
      raw = JSON.parse(content);
    } catch (error) {
      errors.push({ path, message: `Invalid advancement JSON: ${(error as Error).message}` });
      continue;
    }
    if (!raw || typeof raw !== "object" || Array.isArray(raw)) {
      errors.push({ path, message: "Advancement must be a JSON object" });
      continue;
    }
    const data = raw as Record<string, unknown>;
    const rawCriteria = data.criteria;
    if (!rawCriteria || typeof rawCriteria !== "object" || Array.isArray(rawCriteria)) {
      errors.push({ path, message: "Advancement is missing a criteria object" });
      continue;
    }
    const criteria: AdvancementCriterion[] = [];
    for (const [nameKey, value] of Object.entries(rawCriteria as Record<string, unknown>)) {
      if (!value || typeof value !== "object" || Array.isArray(value)) {
        errors.push({ path, message: `Criterion '${nameKey}' must be an object` });
        continue;
      }
      const entry = value as Record<string, unknown>;
      if (typeof entry.trigger !== "string" || !entry.trigger) {
        errors.push({ path, message: `Criterion '${nameKey}' is missing a trigger` });
        continue;
      }
      criteria.push({
        name: nameKey,
        trigger: entry.trigger.replace(/^minecraft:/, ""),
        conditions:
          entry.conditions && typeof entry.conditions === "object" && !Array.isArray(entry.conditions)
            ? (entry.conditions as Record<string, unknown>)
            : null,
      });
    }
    if (criteria.length === 0) {
      errors.push({ path, message: "Advancement has no valid criteria" });
      continue;
    }

    const displayRaw =
      data.display && typeof data.display === "object" && !Array.isArray(data.display)
        ? (data.display as Record<string, unknown>)
        : null;
    let display: AdvancementDisplay | undefined;
    if (displayRaw) {
      const iconRaw =
        displayRaw.icon && typeof displayRaw.icon === "object" && !Array.isArray(displayRaw.icon)
          ? (displayRaw.icon as Record<string, unknown>)
          : null;
      display = {
        title: displayRaw.title !== undefined ? jsonToTag(displayRaw.title) : undefined,
        description:
          displayRaw.description !== undefined ? jsonToTag(displayRaw.description) : undefined,
        icon: typeof iconRaw?.id === "string" ? iconRaw.id : undefined,
        frame: typeof displayRaw.frame === "string" ? displayRaw.frame : undefined,
        background: typeof displayRaw.background === "string" ? displayRaw.background : undefined,
        hidden: displayRaw.hidden === true,
      };
    }

    const rewardsRaw =
      data.rewards && typeof data.rewards === "object" && !Array.isArray(data.rewards)
        ? (data.rewards as Record<string, unknown>)
        : {};
    const requirements = Array.isArray(data.requirements)
      ? (data.requirements as unknown[])
          .filter((group): group is unknown[] => Array.isArray(group))
          .map((group) => group.filter((entry): entry is string => typeof entry === "string"))
      : undefined;

    advancements.set(`${namespace}:${name}`, {
      id: `${namespace}:${name}`,
      path,
      parent: typeof data.parent === "string" ? data.parent : undefined,
      display,
      criteria,
      requirements,
      rewards: {
        function: typeof rewardsRaw.function === "string" ? rewardsRaw.function : undefined,
        experience:
          typeof rewardsRaw.experience === "number" ? rewardsRaw.experience : undefined,
        loot: stringArray(rewardsRaw.loot),
        recipes: stringArray(rewardsRaw.recipes),
      },
      data,
    });
  }
  return { advancements, errors };
}

/** True when every requirement group has at least one completed criterion. */
export function requirementsMet(
  definition: AdvancementDefinition,
  completed: Set<string>,
): boolean {
  if (!definition.requirements || definition.requirements.length === 0) {
    return definition.criteria.every((criterion) => completed.has(criterion.name));
  }
  return definition.requirements.every(
    (group) => group.length === 0 || group.some((name) => completed.has(name)),
  );
}
