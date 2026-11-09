// The Jigsaw tool both editors show (JigsawTool.svelte): a form for one
// jigsaw block that builds its /give command, the blocks copied before
// (Recent), and the jigsaw blocks of the pack's structure files. A host
// works out its state from the pack (jigsawToolState) and carries out the
// actions the tool asks for: copy, forget, declare a hole, show a file.
import type { OutputAsset, OutputFile } from '../types';
import type { JigsawConfig } from './give';

export interface JigsawPoolRow {
	id: string;
	/** The jigsaw names its pieces have, when its pieces are project structure files. */
	names: string[];
}

export interface JigsawHoleRow {
	id: string;
	doc?: string;
}

/** A jigsaw block in one of the pack's structure files. */
export interface JigsawExisting {
	/** The structure's project file, relative to the project. */
	src: string;
	/** Its id in the pack, such as dungeon:halls/hall. */
	id: string;
	pos: [number, number, number];
	orientation?: string;
	config: JigsawConfig;
}

export interface JigsawRecent {
	config: JigsawConfig;
	/** When it was last copied, in ms since the epoch. */
	at: number;
}

export interface JigsawState {
	project?: { name: string; ns?: string };
	holes: JigsawHoleRow[];
	pools: JigsawPoolRow[];
	recent: JigsawRecent[];
	existing: JigsawExisting[];
	/** The holes can only be read from a project without errors in their files. */
	hasErrors: boolean;
}

/** What the tool asks its host to do. */
export type JigsawAction =
	/** Copy the command, and remember the block as recent. */
	| { type: 'copy'; config: JigsawConfig; command: string }
	| { type: 'forget'; config: JigsawConfig }
	| { type: 'forgetAll' }
	/** Add `JigsawHole <name>` to the project, next to its other holes. */
	| { type: 'declare'; id: string }
	| { type: 'revealHole'; id: string }
	| { type: 'revealFile'; src: string };

/** The tool's state for a project, from the pack the engine built. */
export function jigsawToolState(p: {
	project?: { name: string; ns?: string };
	holes: readonly { id: string; doc?: string }[];
	files: readonly OutputFile[];
	assets: readonly OutputAsset[];
	recent: JigsawRecent[];
	hasErrors: boolean;
}): JigsawState {
	const state: JigsawState = { holes: [], pools: [], recent: [], existing: [], hasErrors: p.hasErrors };
	if (!p.project) return state;
	state.project = p.project;
	state.holes = p.holes.map((h) => ({ id: h.id, doc: h.doc }));
	state.recent = p.recent;
	// Structure files by id, with their jigsaw blocks.
	const structures = new Map<string, NonNullable<OutputAsset['jigsaws']>>();
	for (const a of p.assets) {
		const m = /^data\/([^/]+)\/structures?\/(.+)\.nbt$/.exec(a.path);
		if (!m || !a.jigsaws) continue;
		const id = `${m[1]}:${m[2]}`;
		structures.set(id, a.jigsaws);
		for (const j of a.jigsaws) {
			state.existing.push({
				src: a.src,
				id,
				pos: j.pos,
				orientation: j.orientation,
				config: { label: j.label ?? '', pool: field(j.pool), target: field(j.target), name: field(j.name), finalState: j.final_state ?? 'minecraft:air', joint: j.joint === 'aligned' ? 'aligned' : 'rollable' }
			});
		}
	}
	for (const f of p.files) {
		const m = /^data\/([^/]+)\/worldgen\/template_pool\/(.+)\.json$/.exec(f.path);
		if (!m) continue;
		const names = new Set<string>();
		for (const loc of poolLocations(f.text)) for (const j of structures.get(loc) ?? []) if (j.name !== 'minecraft:empty') names.add(j.name);
		state.pools.push({ id: `${m[1]}:${m[2]}`, names: [...names].sort() });
	}
	state.pools.sort((a, b) => a.id.localeCompare(b.id));
	return state;
}

/** An id as the form holds it: minecraft:empty, the game's default, as an empty field. */
function field(id: string): string {
	return id === 'minecraft:empty' ? '' : id;
}

/** The structure ids a template pool's JSON places, list elements included. */
function poolLocations(text: string): string[] {
	const out: string[] = [];
	const walk = (el: unknown) => {
		if (!el || typeof el !== 'object') return;
		const o = el as { location?: unknown; elements?: unknown[] };
		if (typeof o.location === 'string') out.push(o.location.includes(':') ? o.location : `minecraft:${o.location}`);
		if (Array.isArray(o.elements)) o.elements.forEach(walk);
	};
	try {
		const pool = JSON.parse(text) as { elements?: { element?: unknown }[] };
		for (const e of pool.elements ?? []) walk(e.element);
	} catch {
		// Not JSON yet; no pieces.
	}
	return out;
}

/** How many copied blocks Recent keeps per project. */
export const RECENT_MAX = 20;

/**
 * Where `JigsawHole <name>` goes: after the project's last hole, in the
 * file the holes start in, or else at the end of the file of its first
 * TemplatePool. `line` is the 0-based line to insert before, or undefined
 * for the end of the file; `text` is what to insert there.
 */
export function holePlacement(
	id: string,
	ns: string | undefined,
	holes: readonly { id: string; uri: string; range: { end: { line: number } } }[],
	files: readonly OutputFile[]
): { uri: string; line?: number; text: string; name: string } | null {
	const [hns, name] = id.includes(':') ? (id.split(':', 2) as [string, string]) : [ns ?? 'minecraft', id];
	const first = holes[0];
	const last = holes.filter((h) => h.uri === first?.uri).reduce<(typeof holes)[number] | undefined>((a, b) => (!a || b.range.end.line > a.range.end.line ? b : a), undefined);
	let uri: string | undefined;
	let line: number | undefined;
	let fileNs: string | undefined;
	if (last) {
		uri = last.uri;
		line = last.range.end.line + 1;
		fileNs = last.id.split(':')[0];
	} else {
		const pool = files.find((f) => /^data\/[^/]+\/worldgen\/template_pool\//.test(f.path) && f.uri);
		if (!pool?.uri) return null;
		uri = pool.uri;
		fileNs = pool.path.split('/')[1];
	}
	const decl = `JigsawHole ${hns === fileNs ? name : `${hns}:${name}`}`;
	return { uri, line, text: last ? `${decl}\n` : `\n${decl}\n`, name: decl };
}
