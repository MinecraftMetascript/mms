// Runs generated mcfunction files in cmdlib's command simulator, for the
// function preview. The simulator is vendored in ../vendor/cmdlib; it runs
// commands against a small in-memory world (scores, storage, entities as
// NBT, chat) without a game: no AI, physics or block ticks.
import type { OutputFile } from '../types';
export { runnable } from '../paths';
import type { Project } from '../vendor/cmdlib/datapack';
import { Simulator } from '../vendor/cmdlib/simulator';
import { renderTextComponent, type TextRun } from '../vendor/cmdlib/text';
import { stringifySnbt } from '../vendor/cmdlib/nbt';
import type { LogKind } from '../vendor/cmdlib/world';
import { entityCustomName, entityPos } from '../vendor/cmdlib/entities';

const maxLog = 5000;


/** The pack as a cmdlib project: every data file, keyed by its pack path. */
export function projectOf(output: readonly OutputFile[]): Project {
	const files: Record<string, string> = {};
	for (const f of output) if (f.path.startsWith('data/') || f.path === 'pack.mcmeta') files[f.path] = f.text;
	return { id: 'mms', name: 'MMS', versionId: '26.3', files, activePath: '', updatedAt: 0 };
}

export interface LogLine {
	kind: LogKind;
	/** Chat and action bar lines keep their styled runs. */
	runs: TextRun[];
}

export interface ScoreRow {
	/** The score holder as the world keys it. */
	key: string;
	/** Its display name. */
	holder: string;
	values: Record<string, number>;
}

/** What the preview shows of the world after a run or a tick. */
export interface RunSnapshot {
	tick: number;
	log: LogLine[];
	objectives: string[];
	scores: ScoreRow[];
	storage: { id: string; snbt: string }[];
	entities: { type: string; name: string; tags: string[]; pos: string }[];
	scheduled: { id: string; due: number }[];
}

/**
 * One simulated world for a pack: start() loads it (the load tags run) and
 * calls the function, tick() steps the tick tag and schedules.
 */
export class Run {
	sim: Simulator;
	constructor(output: readonly OutputFile[]) {
		this.sim = new Simulator({ project: projectOf(output) });
	}

	/** A fresh world: run the load tags, then call the function. */
	start(id: string): void {
		this.sim.reset({ runLoad: false });
		// cmdlib keeps the last 600 lines; its per-command echo ("data") can
		// push the chat out of that within a few ticks, so keep more.
		const w = this.sim.world;
		let n = 0;
		w.log = (kind, text, component) => {
			w.logs.push({ id: ++n, kind, text, component, time: 0 });
			if (w.logs.length > maxLog) w.logs.splice(0, w.logs.length - maxLog);
		};
		this.sim.runLoadTags();
		this.call(id);
	}


	/** Call a function as the player (Dev, at 0 64 0), as a player's /function would. */
	call(id: string): void {
		this.sim.world.log('command', `/function ${id}`);
		this.sim.runCommand(`function ${id}`);
	}

	tick(n = 1): void {
		for (let i = 0; i < n; i++) this.sim.tick();
	}

	snapshot(): RunSnapshot {
		const w = this.sim.world;
		const ctx = w.textContext(w.player() ?? null);
		const log = w.logs.map((l) => ({
			kind: l.kind,
			runs: l.component ? renderTextComponent(l.component, ctx) : [{ text: l.text }]
		}));
		const objectives = [...w.objectives.keys()].sort();
		// The world keeps scores by objective, then holder; the table is by holder.
		const holders = new Map<string, Record<string, number>>();
		for (const [obj, byHolder] of w.scores) {
			for (const [holder, v] of byHolder) {
				let row = holders.get(holder);
				if (!row) holders.set(holder, (row = {}));
				row[obj] = v;
			}
		}
		const scores = [...holders].map(([key, values]) => ({ key, holder: w.holderDisplayName(key), values }));
		scores.sort((a, b) => a.holder.localeCompare(b.holder));
		const storage = [...w.storages].map(([id, tag]) => ({ id, snbt: stringifySnbt(tag, { pretty: true }) }));
		const entities = w.entities.map((e) => {
			const p = entityPos(e);
			const tags = e.nbt.value.Tags;
			return {
				type: e.type.replace(/^minecraft:/, ''),
				name: entityCustomName(e),
				tags: tags?.type === 'list' ? tags.value.map((t) => String(t.value)) : [],
				pos: [p.x, p.y, p.z].map((v) => +v.toFixed(1)).join(' ')
			};
		});
		const scheduled = this.sim.scheduled.map((s) => ({ id: s.functionId, due: s.dueTick }));
		return { tick: w.tick, log, objectives, scores, storage, entities, scheduled };
	}
}
