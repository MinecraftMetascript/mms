import { type NbtTag, Tag } from "./nbt";

/* ------------------------------------------------------------------ */
/* Vanilla registries (curated, 26.3)                                  */
/* ------------------------------------------------------------------ */

export interface EntityTypeInfo {
  id: string;
  name: string;
  category: "player" | "hostile" | "passive" | "neutral" | "object" | "other";
  color: string;
  shape: "humanoid" | "quadruped" | "bird" | "item" | "marker" | "orb" | "object";
  height: number;
  defaults?: Record<string, NbtTag>;
}

const commonNbt = (health = 20): Record<string, NbtTag> => ({
  Pos: Tag.list([Tag.double(0), Tag.double(64), Tag.double(0)], "double"),
  Motion: Tag.list([Tag.double(0), Tag.double(0), Tag.double(0)], "double"),
  Rotation: Tag.list([Tag.float(0), Tag.float(0)], "float"),
  Health: Tag.float(health),
  CustomName: Tag.string('{"text":""}'),
  CustomNameVisible: Tag.byte(0),
  Tags: Tag.list([], "string"),
  Invulnerable: Tag.byte(0),
  NoGravity: Tag.byte(0),
  Glowing: Tag.byte(0),
  Silent: Tag.byte(0),
  NoAI: Tag.byte(1),
});

const entity = (
  id: string,
  name: string,
  category: EntityTypeInfo["category"],
  color: string,
  shape: EntityTypeInfo["shape"],
  height = 1.8,
  defaults?: Record<string, NbtTag>,
): EntityTypeInfo => ({ id: `minecraft:${id}`, name, category, color, shape, height, defaults });

export const ENTITY_TYPES: EntityTypeInfo[] = [
  entity("player", "Player", "player", "#FFD166", "humanoid", 1.8, {
    ...commonNbt(20),
    foodLevel: Tag.int(20),
    foodSaturation: Tag.float(5),
    XpLevel: Tag.int(0),
    XpP: Tag.float(0),
    XpTotal: Tag.int(0),
    SelectedItemSlot: Tag.int(0),
  }),
  entity("armor_stand", "Armor Stand", "object", "#C9CED6", "marker", 1.9, {
    ...commonNbt(20),
    Invisible: Tag.byte(0),
    Marker: Tag.byte(0),
    Small: Tag.byte(0),
    ShowArms: Tag.byte(0),
    NoBasePlate: Tag.byte(0),
  }),
  entity("zombie", "Zombie", "hostile", "#5E9C54", "humanoid", 1.95, {
    ...commonNbt(20),
    IsBaby: Tag.byte(0),
    CanBreakDoors: Tag.byte(0),
  }),
  entity("husk", "Husk", "hostile", "#B9A57B", "humanoid", 1.95),
  entity("drowned", "Drowned", "hostile", "#5C9389", "humanoid", 1.95),
  entity("skeleton", "Skeleton", "hostile", "#D6D6D6", "humanoid", 1.99),
  entity("stray", "Stray", "hostile", "#9FB8C8", "humanoid", 1.99),
  entity("creeper", "Creeper", "hostile", "#63B25A", "object", 1.7),
  entity("spider", "Spider", "hostile", "#4A3A36", "quadruped", 0.9),
  entity("witch", "Witch", "hostile", "#4B3F5D", "humanoid", 1.95),
  entity("villager", "Villager", "passive", "#B78962", "humanoid", 1.95),
  entity("wandering_trader", "Wandering Trader", "passive", "#5B6EA8", "humanoid", 1.95),
  entity("iron_golem", "Iron Golem", "neutral", "#C6C0B2", "humanoid", 2.7),
  entity("snow_golem", "Snow Golem", "passive", "#E9F2F4", "humanoid", 1.9),
  entity("cow", "Cow", "passive", "#8A6248", "quadruped", 1.4),
  entity("pig", "Pig", "passive", "#E39AA4", "quadruped", 0.9),
  entity("sheep", "Sheep", "passive", "#E9E5DA", "quadruped", 1.3),
  entity("chicken", "Chicken", "passive", "#EDE7D6", "bird", 0.7),
  entity("wolf", "Wolf", "neutral", "#B8B4AA", "quadruped", 0.85),
  entity("cat", "Cat", "passive", "#C8A87C", "quadruped", 0.7),
  entity("horse", "Horse", "passive", "#9A6B45", "quadruped", 1.6),
  entity("bee", "Bee", "neutral", "#E8C34A", "orb", 0.6),
  entity("allay", "Allay", "passive", "#6FD3E0", "orb", 0.6),
  entity("item", "Item", "object", "#D9B45B", "item", 0.25, {
    ...commonNbt(5),
    Item: Tag.compound({ id: Tag.string("minecraft:stone"), count: Tag.int(1) }),
    PickupDelay: Tag.short(0),
    Age: Tag.short(0),
  }),
  entity("item_frame", "Item Frame", "object", "#B08A52", "object", 0.5),
  entity("item_display", "Item Display", "object", "#C5A86A", "item", 0.5, {
    ...commonNbt(20),
    item: Tag.compound({ id: Tag.string("minecraft:stone"), count: Tag.int(1) }),
    transformation: Tag.compound({
      translation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0)], "float"),
      left_rotation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0), Tag.float(1)], "float"),
      scale: Tag.list([Tag.float(1), Tag.float(1), Tag.float(1)], "float"),
      right_rotation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0), Tag.float(1)], "float"),
    }),
    billboard: Tag.string("fixed"),
  }),
  entity("block_display", "Block Display", "object", "#A88D62", "object", 1, {
    ...commonNbt(20),
    block_state: Tag.compound({ Name: Tag.string("minecraft:stone") }),
    transformation: Tag.compound({
      translation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0)], "float"),
      left_rotation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0), Tag.float(1)], "float"),
      scale: Tag.list([Tag.float(1), Tag.float(1), Tag.float(1)], "float"),
      right_rotation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0), Tag.float(1)], "float"),
    }),
  }),
  entity("text_display", "Text Display", "object", "#CBB8E8", "marker", 0.5, {
    ...commonNbt(20),
    text: Tag.string('{"text":""}'),
    transformation: Tag.compound({
      translation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0)], "float"),
      left_rotation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0), Tag.float(1)], "float"),
      scale: Tag.list([Tag.float(1), Tag.float(1), Tag.float(1)], "float"),
      right_rotation: Tag.list([Tag.float(0), Tag.float(0), Tag.float(0), Tag.float(1)], "float"),
    }),
  }),
  entity("interaction", "Interaction", "object", "#9BA7B4", "marker", 1),
  entity("marker", "Marker", "object", "#7CE3D2", "marker", 0.1, {
    ...commonNbt(20),
    data: Tag.compound(),
  }),
  entity("area_effect_cloud", "Area Effect Cloud", "other", "#B87CD8", "orb", 0.5),
  entity("experience_orb", "Experience Orb", "other", "#A6F06B", "orb", 0.3),
  entity("arrow", "Arrow", "other", "#C9B38A", "object", 0.5),
  entity("fireball", "Fireball", "other", "#F08B3E", "orb", 1),
  entity("tnt", "Primed TNT", "other", "#E05A4E", "object", 1),
  entity("boat", "Boat", "object", "#A87645", "object", 0.6),
  entity("minecart", "Minecart", "object", "#8B8F94", "object", 0.7),
  entity("falling_block", "Falling Block", "other", "#8C8C8C", "object", 1),
  entity("lightning_bolt", "Lightning Bolt", "other", "#C8E6FF", "marker", 0.1),
];

export const ENTITY_BY_ID = new Map(
  ENTITY_TYPES.map((e) => [e.id.replace("minecraft:", ""), e]),
);

export function entityInfo(type: string): EntityTypeInfo {
  return (
    ENTITY_BY_ID.get(type.replace("minecraft:", "")) ??
    ENTITY_BY_ID.get("marker")!
  );
}

/** NBT keys worth completing, with Minecraft's own scalar types. */
export interface NbtKeyInfo {
  key: string;
  type: string;
  description: string;
  entities?: string[];
}

export const COMMON_NBT_KEYS: NbtKeyInfo[] = [
  { key: "Pos", type: "list<double>", description: "World position [x, y, z]." },
  { key: "Motion", type: "list<double>", description: "Velocity in blocks per tick." },
  { key: "Rotation", type: "list<float>", description: "Yaw and pitch, in degrees." },
  { key: "Health", type: "float", description: "Current health. 0 kills the entity." },
  { key: "CustomName", type: "string", description: "JSON/SNBT text component shown above the entity." },
  { key: "CustomNameVisible", type: "byte", description: "Render the custom name even without line of sight." },
  { key: "Tags", type: "list<string>", description: "Scoreboard-style tags used by selectors." },
  { key: "Invulnerable", type: "byte", description: "Ignore most damage sources." },
  { key: "NoGravity", type: "byte", description: "Disable gravity." },
  { key: "Glowing", type: "byte", description: "Apply the glowing outline." },
  { key: "Silent", type: "byte", description: "Suppress idle sounds." },
  { key: "NoAI", type: "byte", description: "Freeze mob AI. cmdlib never runs AI anyway." },
  { key: "Air", type: "short", description: "Remaining breath ticks." },
  { key: "Fire", type: "short", description: "Remaining fire ticks." },
  { key: "OnGround", type: "byte", description: "Whether the entity is on the ground." },
  { key: "FallDistance", type: "float", description: "Distance fallen, used for fall damage." },
  { key: "PortalCooldown", type: "int", description: "Ticks before the entity can portal again." },
  { key: "HurtTime", type: "short", description: "Red flash timer." },
  { key: "DeathTime", type: "short", description: "Death animation timer." },
];

export const ENTITY_SPECIFIC_NBT: Record<string, NbtKeyInfo[]> = {
  "minecraft:player": [
    { key: "foodLevel", type: "int", description: "Hunger bar value (0–20)." },
    { key: "XpLevel", type: "int", description: "Experience level." },
    { key: "XpP", type: "float", description: "Progress toward the next level." },
    { key: "SelectedItemSlot", type: "int", description: "Hotbar slot (0–8)." },
    { key: "Inventory", type: "list<compound>", description: "Main inventory." },
    { key: "EnderItems", type: "list<compound>", description: "Ender chest contents." },
    { key: "abilities", type: "compound", description: "Flight and invulnerability flags." },
    { key: "Score", type: "int", description: "Displayed score." },
  ],
  "minecraft:zombie": [
    { key: "IsBaby", type: "byte", description: "Spawn the baby variant." },
    { key: "CanBreakDoors", type: "byte", description: "Allow door breaking." },
  ],
  "minecraft:armor_stand": [
    { key: "Invisible", type: "byte", description: "Hide the stand." },
    { key: "Marker", type: "byte", description: "Remove the hitbox and render nothing but the name." },
    { key: "Small", type: "byte", description: "Half-size armor stand." },
    { key: "ShowArms", type: "byte", description: "Render the arms." },
    { key: "NoBasePlate", type: "byte", description: "Remove the stone base." },
    { key: "Pose", type: "compound", description: "Per-part rotations." },
  ],
  "minecraft:item": [
    { key: "Item", type: "compound", description: "Item stack: {id, count, components}." },
    { key: "PickupDelay", type: "short", description: "Ticks before a player can pick it up." },
    { key: "Age", type: "short", description: "Ticks the item has existed." },
    { key: "Owner", type: "uuid", description: "Player allowed to pick it up." },
  ],
  "minecraft:villager": [
    { key: "VillagerData", type: "compound", description: "Profession, type and level." },
    { key: "Offers", type: "compound", description: "Trade offers." },
  ],
  "minecraft:marker": [
    { key: "data", type: "compound", description: "Free-form command storage attached to the marker." },
  ],
  "minecraft:tnt": [
    { key: "fuse", type: "short", description: "Ticks until detonation." },
  ],
};

export function nbtKeysFor(type: string): NbtKeyInfo[] {
  return [
    ...(ENTITY_SPECIFIC_NBT[type] ?? []),
    ...COMMON_NBT_KEYS,
  ];
}

export const SCOREBOARD_CRITERIA = [
  { name: "dummy", description: "Free-form, changed only by commands." },
  { name: "trigger", description: "Changed by players with /trigger." },
  { name: "deathCount", description: "Times the player has died." },
  { name: "playerKillCount", description: "Players killed." },
  { name: "totalKillCount", description: "All mobs and players killed." },
  { name: "health", description: "Current health." },
  { name: "food", description: "Current hunger." },
  { name: "air", description: "Current breath." },
  { name: "armor", description: "Armor points." },
  { name: "xp", description: "Experience points." },
  { name: "level", description: "Experience level." },
  ...["black", "dark_blue", "dark_green", "dark_aqua", "dark_red", "dark_purple", "gold", "gray", "dark_gray", "blue", "green", "aqua", "red", "light_purple", "yellow", "white"].flatMap(
    (color) => [
      { name: `teamkill.${color}`, description: `Kills of ${color} team members.` },
      { name: `killedByTeam.${color}`, description: `Deaths caused by ${color} team.` },
    ],
  ),
];

export const SCOREBOARD_OPERATIONS = [
  { name: "=", description: "Set target to source." },
  { name: "+=", description: "Add source to target." },
  { name: "-=", description: "Subtract source from target." },
  { name: "*=", description: "Multiply target by source." },
  { name: "/=", description: "Integer-divide target by source." },
  { name: "%=", description: "Modulo of target by source." },
  { name: "<", description: "Keep the smaller value." },
  { name: ">", description: "Keep the larger value." },
  { name: "><", description: "Swap target and source." },
];

export const DISPLAY_SLOTS = [
  { name: "sidebar", description: "The right-hand sidebar." },
  { name: "list", description: "The tab player list." },
  { name: "below_name", description: "Below the player's name tag." },
  ...["black", "dark_blue", "dark_green", "dark_aqua", "dark_red", "dark_purple", "gold", "gray", "dark_gray", "blue", "green", "aqua", "red", "light_purple", "yellow", "white"].map(
    (color) => ({
      name: `sidebar.team.${color}`,
      description: `Sidebar for the ${color} team.`,
    }),
  ),
];

export const TEXT_COMPONENT_KEYS = [
  { key: "text", description: "Literal text." },
  { key: "translate", description: "Translation key." },
  { key: "fallback", description: "Fallback used when the key is missing." },
  { key: "with", description: "Arguments for the translation." },
  { key: "score", description: "{name, objective} — read a scoreboard value." },
  { key: "selector", description: "Resolve an entity selector to names." },
  { key: "nbt", description: "Read a tag from storage or an entity." },
  { key: "keybind", description: "Render a key binding name." },
  { key: "color", description: "Named colour or #rrggbb." },
  { key: "bold", description: "Bold style." },
  { key: "italic", description: "Italic style." },
  { key: "underlined", description: "Underline style." },
  { key: "strikethrough", description: "Strikethrough style." },
  { key: "obfuscated", description: "Scrambled glyphs." },
  { key: "extra", description: "Child components." },
];

export const TEXT_COLORS = [
  "black", "dark_blue", "dark_green", "dark_aqua", "dark_red", "dark_purple",
  "gold", "gray", "dark_gray", "blue", "green", "aqua", "red", "light_purple",
  "yellow", "white",
];

export const GAMEMODES = [
  "survival",
  "creative",
  "adventure",
  "spectator",
] as const;

/** Curated status-effect registry used by `/effect` validation. */
export const STATUS_EFFECTS = [
  { name: "speed", description: "Increases movement speed." },
  { name: "slowness", description: "Decreases movement speed." },
  { name: "haste", description: "Increases mining speed." },
  { name: "mining_fatigue", description: "Decreases mining speed." },
  { name: "strength", description: "Increases attack damage." },
  { name: "instant_health", description: "Heals instantly." },
  { name: "instant_damage", description: "Deals damage instantly." },
  { name: "jump_boost", description: "Increases jump height." },
  { name: "nausea", description: "Warps the screen." },
  { name: "regeneration", description: "Regenerates health over time." },
  { name: "resistance", description: "Reduces incoming damage." },
  { name: "fire_resistance", description: "Prevents fire damage." },
  { name: "water_breathing", description: "Prevents drowning." },
  { name: "invisibility", description: "Hides the entity." },
  { name: "blindness", description: "Blinds the player." },
  { name: "night_vision", description: "Brightens the view." },
  { name: "hunger", description: "Increases hunger." },
  { name: "weakness", description: "Reduces attack damage." },
  { name: "poison", description: "Deals damage over time." },
  { name: "wither", description: "Deals wither damage over time." },
  { name: "health_boost", description: "Raises maximum health." },
  { name: "absorption", description: "Adds absorption hearts." },
  { name: "saturation", description: "Restores hunger and saturation." },
  { name: "glowing", description: "Outlines the entity." },
  { name: "levitation", description: "Makes the entity float." },
  { name: "luck", description: "Improves loot luck." },
  { name: "unluck", description: "Reduces loot luck." },
  { name: "slow_falling", description: "Reduces fall speed." },
  { name: "conduit_power", description: "Improves underwater activity." },
  { name: "dolphins_grace", description: "Increases swim speed." },
  { name: "bad_omen", description: "Triggers raids." },
  { name: "hero_of_the_village", description: "Improves villager trades." },
  { name: "darkness", description: "Darkens the screen." },
  { name: "trial_omen", description: "Triggers trial spawners." },
  { name: "raid_omen", description: "Triggers a raid." },
  { name: "wind_charged", description: "Allows wind charging." },
  { name: "weaving", description: "Spreads cobwebs on death." },
  { name: "oozing", description: "Spawns slimes on death." },
  { name: "infested", description: "Spawns silverfish on death." },
];

export const STATUS_EFFECT_IDS = new Set(
  STATUS_EFFECTS.map((effect) => `minecraft:${effect.name}`),
);

/* ------------------------------------------------------------------ */
/* Command metadata                                                    */
/* ------------------------------------------------------------------ */

export type ArgType =
  | "word"
  | "integer"
  | "number"
  | "objective"
  | "newObjective"
  | "targets"
  | "selector"
  | "scoreHolder"
  | "score"
  | "function"
  | "storage"
  | "path"
  | "nbt"
  | "component"
  | "coordinates"
  | "rotation"
  | "range"
  | "enum"
  | "greedy"
  | "literal";

export interface CommandArg {
  name: string;
  type: ArgType;
  description?: string;
  optional?: boolean;
  enumValues?: string[];
  rest?: boolean;
}

export interface CommandSpec {
  name: string;
  summary: string;
  category: "player" | "world" | "data" | "server" | "integration";
  supported: boolean;
  /** Shown in the reference panel. */
  usage?: string;
  args?: CommandArg[];
  children?: Record<string, CommandSpec>;
  since?: string;
}

const arg = (name: string, type: ArgType, description?: string, optional = false): CommandArg => ({
  name,
  type,
  description,
  optional,
});

const group = (
  name: string,
  summary: string,
  children: Record<string, CommandSpec>,
  category: CommandSpec["category"] = "data",
  supported = true,
): CommandSpec => ({ name, summary, category, supported, children });

const subject = (targets: string, objective: string, score?: string): CommandArg[] => [
  arg(targets, "targets", "A player name, selector, or * for every tracked holder."),
  arg(objective, "objective", "Objective name."),
  ...(score ? [arg(score, "integer", "Score value.")] : []),
];

export const COMMANDS: Record<string, CommandSpec> = {
  execute: {
    name: "execute",
    summary: "Change the command context, test conditions, store a result, or run a command.",
    category: "server",
    supported: true,
    usage: "/execute <as|at|anchored|positioned|align|in|rotated|facing|if|unless|store> … run <command>",
    children: {
      as: { name: "as", summary: "Run once for each selected entity, as that entity.", category: "server", supported: true, args: [arg("targets", "targets")] },
      at: { name: "at", summary: "Run at an entity's position and rotation.", category: "server", supported: true, args: [arg("targets", "targets")] },
      anchored: { name: "anchored", summary: "Use the eyes or feet as the command position.", category: "server", supported: true, args: [arg("anchor", "enum", "eyes or feet.")] },
      positioned: { name: "positioned", summary: "Set the command position.", category: "server", supported: true, args: [arg("pos", "coordinates")] },
      align: { name: "align", summary: "Round the command position to the block grid.", category: "server", supported: true, args: [arg("axes", "enum", "Any combination of x, y, z.", false)] },
      rotated: { name: "rotated", summary: "Set the command rotation.", category: "server", supported: true, args: [arg("rotation", "rotation")] },
      facing: { name: "facing", summary: "Face a position or entity.", category: "server", supported: true, args: [arg("target", "coordinates")] },
      in: { name: "in", summary: "Switch dimension.", category: "server", supported: true, args: [arg("dimension", "word")] },
      if: {
        name: "if",
        summary: "Run only if a condition is true.",
        category: "server",
        supported: true,
        children: {
          predicate: { name: "predicate", summary: "Test a datapack predicate by resource location.", category: "server", supported: true, args: [arg("predicate", "word", "Predicate id, e.g. cmdlib:is_hungry.")] },
        },
      },
      unless: {
        name: "unless",
        summary: "Run only if a condition is false.",
        category: "server",
        supported: true,
        children: {
          predicate: { name: "predicate", summary: "Test a datapack predicate by resource location.", category: "server", supported: true, args: [arg("predicate", "word", "Predicate id, e.g. cmdlib:is_hungry.")] },
        },
      },
      store: { name: "store", summary: "Store the result or success value.", category: "server", supported: true, children: {} },
      run: { name: "run", summary: "The command to run.", category: "server", supported: true, args: [arg("command", "greedy")] },
    },
  },
  scoreboard: group("scoreboard", "Manage objectives, players and display slots.", {
    objectives: group("objectives", "Create, remove and display objectives.", {
      add: { name: "add", summary: "Create a new objective.", category: "data", supported: true, args: [arg("objective", "newObjective"), arg("criterion", "enum"), arg("displayName", "component", undefined, true)] },
      remove: { name: "remove", summary: "Delete an objective and all its scores.", category: "data", supported: true, args: [arg("objective", "objective")] },
      list: { name: "list", summary: "List every objective.", category: "data", supported: true },
      modify: { name: "modify", summary: "Change a display name or number format.", category: "data", supported: true, args: [arg("objective", "objective")] },
      setdisplay: { name: "setdisplay", summary: "Put an objective in a display slot.", category: "data", supported: true, args: [arg("slot", "enum"), arg("objective", "objective", undefined, true)] },
    }),
    players: group("players", "Read and write score holders.", {
      set: { name: "set", summary: "Set a score.", category: "data", supported: true, args: subject("targets", "objective", "score") },
      add: { name: "add", summary: "Add to a score.", category: "data", supported: true, args: subject("targets", "objective", "score") },
      remove: { name: "remove", summary: "Subtract from a score.", category: "data", supported: true, args: subject("targets", "objective", "score") },
      reset: { name: "reset", summary: "Remove a score (or all of a holder's scores).", category: "data", supported: true, args: [arg("targets", "targets"), arg("objective", "objective", undefined, true)] },
      get: { name: "get", summary: "Print a score.", category: "data", supported: true, args: subject("targets", "objective") },
      operation: { name: "operation", summary: "Combine two scores.", category: "data", supported: true, args: [arg("targets", "targets"), arg("objective", "objective"), arg("operation", "enum"), arg("source", "scoreHolder"), arg("sourceObjective", "objective")] },
      display: { name: "display", summary: "Set the name shown for a holder.", category: "data", supported: true, args: [arg("name", "scoreHolder")] },
      enable: { name: "enable", summary: "Enable a trigger objective.", category: "data", supported: true, usage: "/scoreboard players enable <targets> <objective>", args: subject("targets", "objective") },
    }),
  }, "data"),

  title: {
    name: "title",
    summary: "Show a title, subtitle, actionbar, or configure timings.",
    category: "server",
    supported: true,
    usage: "/title <targets> <title|subtitle|actionbar> <component>",
    args: [arg("targets", "targets"), arg("action", "enum"), arg("component", "component", undefined, true)],
  },
  tick: {
    name: "tick",
    summary: "Freeze, unfreeze, step or change the simulation tick rate.",
    category: "server",
    supported: true,
    usage: "/tick <freeze|unfreeze|rate <tps>|step [ticks]|sprint [ticks]|query>",
    args: [arg("action", "enum"), arg("value", "number", undefined, true)],
  },
  forceload: {
    name: "forceload",
    summary: "Keep chunks loaded in the simulation.",
    category: "server",
    supported: true,
    usage: "/forceload <add|remove|query> <x> <z> [<x2> <z2>]",
    args: [arg("action", "enum"), arg("x", "integer"), arg("z", "integer")],
  },
  function: {
    name: "function",
    summary: "Run a function, optionally with macro arguments.",
    category: "data",
    supported: true,
    usage: "/function <id> [<arguments>] | /function <id> with <storage|entity|block> …",
    args: [arg("function", "function"), arg("arguments", "nbt", undefined, true)],
  },
  data: group("data", "Read and modify NBT on storage, entities and block entities.", {
    get: { name: "get", summary: "Print NBT.", category: "data", supported: true, args: [arg("source", "enum"), arg("target", "word"), arg("path", "path", undefined, true), arg("scale", "number", undefined, true)] },
    merge: { name: "merge", summary: "Merge a compound into a target.", category: "data", supported: true, args: [arg("source", "enum"), arg("target", "word"), arg("nbt", "nbt")] },
    modify: { name: "modify", summary: "Change a tag in place.", category: "data", supported: true, args: [arg("source", "enum"), arg("target", "word"), arg("path", "path")] },
    remove: { name: "remove", summary: "Remove a tag.", category: "data", supported: true, args: [arg("source", "enum"), arg("target", "word"), arg("path", "path")] },
  }, "data"),
  summon: {
    name: "summon",
    summary: "Create an entity at a position with optional NBT.",
    category: "world",
    supported: true,
    args: [arg("entity", "word"), arg("pos", "coordinates", undefined, true), arg("nbt", "nbt", undefined, true)],
  },
  kill: {
    name: "kill",
    summary: "Remove entities.",
    category: "player",
    supported: true,
    args: [arg("targets", "targets", undefined, true)],
  },
  teleport: {
    name: "teleport",
    summary: "Move entities to a position or entity.",
    category: "player",
    supported: true,
    args: [arg("targets", "targets"), arg("destination", "coordinates")],
  },
  tag: {
    name: "tag",
    summary: "Add, remove or list entity tags.",
    category: "data",
    supported: true,
    args: [arg("targets", "targets"), arg("action", "enum"), arg("name", "word", undefined, true)],
  },
  say: {
    name: "say",
    summary: "Broadcast a chat message.",
    category: "server",
    supported: true,
    args: [arg("message", "greedy")],
  },
  tellraw: {
    name: "tellraw",
    summary: "Send a text component to players.",
    category: "server",
    supported: true,
    args: [arg("targets", "targets"), arg("component", "component")],
  },
  me: {
    name: "me",
    summary: "Broadcast an action message.",
    category: "server",
    supported: true,
    args: [arg("action", "greedy")],
  },
  return: {
    name: "return",
    summary: "Return a value from a function.",
    category: "server",
    supported: true,
    args: [arg("value", "integer", "The value to return.", true)],
  },
  reload: { name: "reload", summary: "Reload functions and tags from the datapack.", category: "server", supported: true },
  datapack: { name: "datapack", summary: "Enable, disable or list data packs.", category: "server", supported: false },
  help: { name: "help", summary: "List commands.", category: "server", supported: false },
  time: { name: "time", summary: "Read or set world time.", category: "world", supported: false, args: [arg("action", "enum")] },
  weather: { name: "weather", summary: "Set the weather.", category: "world", supported: false },
  gamerule: { name: "gamerule", summary: "Read or set a game rule.", category: "world", supported: false },
  setblock: { name: "setblock", summary: "Place a block.", category: "world", supported: true },
  fill: { name: "fill", summary: "Fill a region with blocks.", category: "world", supported: false },
  clone: { name: "clone", summary: "Copy blocks.", category: "world", supported: false },
  particle: { name: "particle", summary: "Spawn particles.", category: "world", supported: false },
  playsound: { name: "playsound", summary: "Play a sound.", category: "world", supported: false },
  bossbar: { name: "bossbar", summary: "Manage boss bars.", category: "server", supported: false },
  give: {
    name: "give",
    summary: "Give items to players.",
    category: "player",
    supported: true,
    usage: "/give <targets> <item> [<count>] [<nbt>]",
    args: [
      arg("targets", "targets", "Players to receive the item."),
      arg("item", "word", "Item id, with or without the minecraft: namespace."),
      arg("count", "integer", "Positive stack count. Defaults to 1.", true),
      arg("nbt", "nbt", "Optional SNBT payload used to distinguish stacks.", true),
    ],
  },
  clear: {
    name: "clear",
    summary: "Clear items from player inventories.",
    category: "player",
    supported: true,
    usage: "/clear [<targets>] [<item>] [<maxCount>]",
    args: [
      arg("targets", "targets", "Players to clear. Defaults to @s.", true),
      arg("item", "word", "Only clear this item id.", true),
      arg("maxCount", "integer", "Maximum number of items to remove.", true),
    ],
  },
  damage: {
    name: "damage",
    summary: "Damage entities and players.",
    category: "player",
    supported: true,
    usage: "/damage <targets> <amount> [damageType]",
    args: [arg("targets", "targets"), arg("amount", "number"), arg("damageType", "word", undefined, true)],
  },
  effect: {
    name: "effect",
    summary: "Apply or remove status effects.",
    category: "player",
    supported: true,
    usage:
      "/effect give <targets> <effect> [<seconds>] [<amplifier>] [<hideParticles>] | /effect clear <targets> [<effect>]",
    args: [
      arg("action", "enum", "give or clear."),
      arg("targets", "targets", "Players affected by the effect."),
      arg("effect", "word", "Status effect id, e.g. speed or minecraft:speed."),
      arg("seconds", "integer", "Duration in seconds. Defaults to 30.", true),
      arg("amplifier", "integer", "Effect amplifier. Defaults to 0.", true),
      arg("hideParticles", "enum", "true or false. Defaults to false.", true),
    ],
  },
  gamemode: {
    name: "gamemode",
    summary: "Change a player's game mode.",
    category: "player",
    supported: true,
    usage: "/gamemode <survival|creative|adventure|spectator> [<targets>]",
    args: [
      arg("mode", "enum", "survival, creative, adventure or spectator."),
      arg("targets", "targets", "Players to update. Defaults to @s.", true),
    ],
  },
  spectate: {
    name: "spectate",
    summary: "Make a player spectate an entity.",
    category: "player",
    supported: true,
    usage: "/spectate <target> [<player>]",
    args: [
      arg("target", "targets", "Entity to spectate. Must match exactly one entity."),
      arg("player", "targets", "Player whose camera changes. Defaults to @s.", true),
    ],
  },
  trigger: {
    name: "trigger",
    summary: "Add to or set a trigger objective for the executing player.",
    category: "player",
    supported: true,
    usage: "/trigger <objective> | /trigger <objective> add <value> | /trigger <objective> set <value>",
    args: [
      arg("objective", "objective", "A trigger-criterion objective enabled for this player."),
      arg("action", "enum", "add or set. Omitting both adds 1.", true),
      arg("value", "integer", "32-bit value to add or set.", true),
    ],
  },
  xp: { name: "xp", summary: "Add or query experience.", category: "player", supported: false },
  advancement: {
    name: "advancement",
    summary: "Grant or revoke advancements for players.",
    category: "player",
    supported: true,
    usage: "/advancement <grant|revoke> <targets> <only|everything|...> [<advancement>]",
    args: [arg("action", "enum"), arg("targets", "targets"), arg("mode", "enum"), arg("advancement", "word", undefined, true)],
  },
  attribute: { name: "attribute", summary: "Read or set attributes.", category: "player", supported: false },
  item: {
    name: "item",
    summary: "Replace or modify item slots.",
    category: "player",
    supported: true,
    usage: "/item replace entity <targets> <slot> with <item> [count] | /item modify entity <targets> <slot> <modifier>",
    args: [arg("action", "enum"), arg("source", "enum"), arg("targets", "targets"), arg("slot", "word")],
  },
  loot: { name: "loot", summary: "Insert loot.", category: "player", supported: false },
  schedule: {
    name: "schedule",
    summary: "Run a function later, or clear its queued runs.",
    category: "server",
    supported: true,
    usage:
      "/schedule function <function> <time> [append|replace] | /schedule clear <function>",
    args: [
      arg("action", "enum", "function or clear."),
      arg("function", "function", "Function to schedule or clear."),
      arg("time", "word", "Delay as an integer plus t, s or d (20 t/s, 24000 t/d).", true),
      arg("mode", "enum", "append keeps existing runs; replace is the default.", true),
    ],
  },
  random: {
    name: "random",
    summary: "Draw a random value from a range.",
    category: "server",
    supported: true,
    usage: "/random <value|roll> <min>..<max> [sequence] | /random reset <sequence>",
    args: [arg("action", "enum"), arg("range", "word", "Integer range, e.g. 1..6."), arg("sequence", "word", undefined, true)],
  },
  compute: { name: "compute", summary: "Evaluate a number provider (26.3).", category: "server", supported: false, since: "26.3" },
  posteffect: { name: "posteffect", summary: "Apply a post-processing effect (26.3).", category: "world", supported: false, since: "26.3" },
  dialog: { name: "dialog", summary: "Open a dialog screen.", category: "server", supported: false },
  version: { name: "version", summary: "Print the game version.", category: "server", supported: false },
  list: { name: "list", summary: "List players.", category: "server", supported: false },
  seed: { name: "seed", summary: "Print the world seed.", category: "server", supported: false },
};

export const COMMAND_NAMES = Object.keys(COMMANDS).sort();

export function commandSpec(name: string): CommandSpec | undefined {
  return COMMANDS[name];
}

export const SINCE_26_3 = ["compute", "posteffect"];

export const VANILLA_COMMAND_COUNT = 68;
