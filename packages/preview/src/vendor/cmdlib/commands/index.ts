import { executeHandler } from "./execute";
import { functionHandler } from "./function";
import { dataHandler } from "./data";
import { scoreboardHandler } from "./scoreboard";
import { titleHandler } from "./title";
import { scheduleHandler } from "./schedule";
import {
  clearHandler,
  damageHandler,
  effectHandler,
  gamemodeHandler,
  giveHandler,
  itemHandler,
  spectateHandler,
  triggerHandler,
} from "./player";
import {
  advancementHandler,
  forceloadHandler,
  killHandler,
  helpHandler,
  listHandler,
  meHandler,
  reloadHandler,
  randomHandler,
  returnHandler,
  sayHandler,
  setblockHandler,
  summonHandler,
  tagHandler,
  teleportHandler,
  tellrawHandler,
  tickHandler,
} from "./core";
import type { CommandContext, CommandResult } from "./context";
import type { Cursor } from "../cursor";

export type CommandHandler = (cursor: Cursor, ctx: CommandContext) => CommandResult;

export const HANDLERS: Record<string, CommandHandler> = {
  execute: executeHandler,
  scoreboard: scoreboardHandler,
  title: titleHandler,
  function: functionHandler,
  advancement: advancementHandler,
  forceload: forceloadHandler,
  tick: tickHandler,
  data: dataHandler,
  summon: summonHandler,
  kill: killHandler,
  teleport: teleportHandler,
  tp: teleportHandler,
  tag: tagHandler,
  say: sayHandler,
  tellraw: tellrawHandler,
  help: helpHandler,
  list: listHandler,
  me: meHandler,
  return: returnHandler,
  random: randomHandler,
  setblock: setblockHandler,
  reload: reloadHandler,
  schedule: scheduleHandler,
  gamemode: gamemodeHandler,
  give: giveHandler,
  item: itemHandler,
  spectate: spectateHandler,
  clear: clearHandler,
  damage: damageHandler,
  effect: effectHandler,
  trigger: triggerHandler,
};

export { executeHandler, scoreboardHandler, titleHandler, functionHandler, dataHandler };
export { gamemodeHandler, giveHandler, clearHandler, effectHandler, triggerHandler };
export { scheduleHandler, parseScheduleTime, TICKS_PER_SECOND, TICKS_PER_DAY, SCHEDULE_USAGE } from "./schedule";
