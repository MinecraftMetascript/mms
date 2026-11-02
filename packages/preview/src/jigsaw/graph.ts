// Which jigsaw pieces can join which, for the jigsaw connection preview.
//
// A template pool lists pieces (structure files). Each piece holds jigsaw
// blocks: a jigsaw block names the pool the next piece is drawn from, and
// its target, the name of a jigsaw block in that piece to join onto. The
// graph starts at a pool, or at a jigsaw Structure's start pool, and
// follows every jigsaw block outwards, one depth per step.
//
// Free of DOM types, so it can be tested and reused anywhere.

import { jigsawable } from '../paths';
import type { OutputAsset, OutputFile, OutputJigsaw as Jigsaw } from '../types';

export type { Jigsaw };

/** Why a jigsaw block joins nothing. */
export type DeadEnd =
	/** Its pool is in a project namespace but nothing declares it. */
	| 'undeclared'
	/** Its pool has no pieces. */
	| 'empty-pool'
	/** No piece of its pool has a jigsaw block named its target. */
	| 'no-target';

export interface Link {
	/** The jigsaw block, by its index in the piece's jigsaws. */
	jigsaw: number;
	pool: string;
	/** The pool's index in the graph's pools, when the graph has it. */
	to?: number;
	/** The pieces of that pool that can join here, by index. */
	pieces: number[];
	/** Some pieces of the pool have unknown jigsaws, so they may join too. */
	maybe: boolean;
	/**
	 * When no piece of the pool can join, the game tries the pool's
	 * fallback: the fallback's index and its pieces that can join.
	 */
	fallback?: { to: number; pieces: number[] };
	dead?: DeadEnd;
}

export interface Piece {
	/** The structure id, or a description for a feature or empty element. */
	id: string;
	kind: 'single' | 'feature' | 'empty' | 'other';
	weight: number;
	/** Its share of the pool's total weight, 0..1. */
	share: number;
	/** Its jigsaw blocks; undefined when unknown (not a project structure file). */
	jigsaws?: Jigsaw[];
	links: Link[];
	/**
	 * No jigsaw block that draws from this pool targets a name this piece
	 * has, so it is never placed (except as the start piece).
	 */
	unreachable: boolean;
}

export interface Pool {
	id: string;
	/** The output path of its declaration, when the project declares it. */
	path?: string;
	/** Whether the project declares it; false for a vanilla pool. */
	declared: boolean;
	fallback?: string;
	/** The fallback's index in the graph's pools. */
	fallbackTo?: number;
	pieces: Piece[];
	/**
	 * Steps from the start pool: the depth of the pieces it places. A pool
	 * reached at several depths is shown once at each, so a pool that
	 * draws from itself grows out to the right like the structure does.
	 */
	depth: number;
	/** Past the structure's size, so nothing from it is placed. */
	tooDeep: boolean;
}

export interface JigsawGraph {
	/** What the graph starts from: a pool, or a jigsaw structure. */
	root: { kind: 'pool' | 'structure'; id: string };
	/** The structure's size (max depth), for a structure. */
	size?: number;
	/** The structure's start_jigsaw_name, if any. */
	startName?: string;
	/** In order of depth; the first is the start pool. */
	pools: Pool[];
	/**
	 * Growth repeats from here: the last depth's pools are all already
	 * shown at this earlier depth, and its jigsaw blocks link back to it.
	 */
	repeatsFrom?: number;
	problems: { deadEnds: number; unreachable: number; tooDeep: number; unknown: number };
}

type Json = null | boolean | number | string | Json[] | { [k: string]: Json };
type Obj = { [k: string]: Json };
const isObj = (v: Json | undefined): v is Obj => typeof v === 'object' && v !== null && !Array.isArray(v);
const strip = (s: string) => s.replace(/^minecraft:/, '');
const EMPTY = 'minecraft:empty';

/** A resource id with the minecraft namespace added. */
export const fullId = (s: string) => (s.includes(':') ? s : `minecraft:${s}`);

const nbtPath = /^data\/([^/]+)\/structures?\/(.+)\.nbt$/;

function parse(text: string | undefined): Json | undefined {
	if (text === undefined) return undefined;
	try {
		return JSON.parse(text);
	} catch {
		return undefined;
	}
}

/** A pool element's piece; a list element is one piece. */
function elementPieces(el: Json | undefined, weight: number, structures: Map<string, Jigsaw[]>, out: Omit<Piece, 'share' | 'links' | 'unreachable'>[]) {
	if (!isObj(el)) {
		out.push({ id: '?', kind: 'other', weight });
		return;
	}
	const type = strip(String(el.element_type ?? ''));
	switch (type) {
		case 'single_pool_element':
		case 'legacy_single_pool_element': {
			const id = fullId(String(el.location ?? '?'));
			out.push({ id, kind: 'single', weight, jigsaws: structures.get(id) });
			break;
		}
		case 'list_pool_element': {
			// The game places a list's elements together, as one piece
			// with one weight, and joins it by its first element's jigsaw
			// blocks only.
			const parts: Omit<Piece, 'share' | 'links' | 'unreachable'>[] = [];
			for (const e of Array.isArray(el.elements) ? el.elements : []) elementPieces(e, 1, structures, parts);
			const first = parts[0];
			if (!first) out.push({ id: 'empty list', kind: 'empty', weight, jigsaws: [] });
			else out.push({ id: parts.map((q) => q.id).join(' + '), kind: first.kind, weight, jigsaws: first.jigsaws });
			break;
		}
		case 'feature_pool_element':
			out.push({ id: `feature ${strip(String(el.feature ?? '?'))}`, kind: 'feature', weight, jigsaws: [] });
			break;
		case 'empty_pool_element':
			out.push({ id: 'empty', kind: 'empty', weight, jigsaws: [] });
			break;
		default:
			out.push({ id: type || '?', kind: 'other', weight });
	}
}

/**
 * The jigsaw graph from a pool or a jigsaw structure, or null when the
 * path is neither or the structure is not a jigsaw one.
 */
export function jigsawGraph(path: string | null, output: readonly OutputFile[], assets: readonly OutputAsset[] = []): JigsawGraph | null {
	const root = jigsawable(path);
	if (!root) return null;
	const files = new Map(output.map((f) => [f.path, f]));
	const structures = new Map<string, Jigsaw[]>();
	for (const a of assets) {
		const m = nbtPath.exec(a.path);
		if (m && a.jigsaws) structures.set(`${m[1]}:${m[2]}`, a.jigsaws);
	}
	// Namespaces the project writes to: a pool there must be declared.
	const namespaces = new Set<string>();
	for (const f of output) {
		const m = /^data\/([^/]+)\//.exec(f.path);
		if (m) namespaces.add(m[1]);
	}
	const poolFile = (id: string) => {
		const [ns, name] = id.split(':', 2);
		return `data/${ns}/worldgen/template_pool/${name}.json`;
	};

	let start: string;
	const graph: JigsawGraph = { root, pools: [], problems: { deadEnds: 0, unreachable: 0, tooDeep: 0, unknown: 0 } };
	if (root.kind === 'structure') {
		const json = parse(files.get(path!)?.text);
		if (!isObj(json) || strip(String(json.type ?? '')) !== 'jigsaw' || typeof json.start_pool !== 'string') return null;
		start = fullId(json.start_pool);
		if (typeof json.size === 'number') graph.size = json.size;
		if (typeof json.start_jigsaw_name === 'string') graph.startName = fullId(json.start_jigsaw_name);
	} else {
		start = root.id;
	}

	// Pool definitions, read once.
	type Def = Omit<Pool, 'depth' | 'tooDeep' | 'fallbackTo' | 'pieces'> & { pieces: Omit<Piece, 'links' | 'unreachable'>[] };
	const defs = new Map<string, Def>();
	const def = (id: string): Def => {
		let d = defs.get(id);
		if (d) return d;
		const file = files.get(poolFile(id));
		const json = parse(file?.text);
		d = { id, declared: !!file, pieces: [] };
		if (file) d.path = file.path;
		if (isObj(json)) {
			if (typeof json.fallback === 'string' && fullId(json.fallback) !== EMPTY) d.fallback = fullId(json.fallback);
			const raw: Omit<Piece, 'share' | 'links' | 'unreachable'>[] = [];
			for (const e of Array.isArray(json.elements) ? json.elements : []) {
				if (!isObj(e)) continue;
				elementPieces(e.element, typeof e.weight === 'number' ? e.weight : 1, structures, raw);
			}
			const total = raw.reduce((n, p) => n + p.weight, 0);
			d.pieces = raw.map((p) => ({ ...p, share: total ? p.weight / total : 0 }));
		}
		defs.set(id, d);
		return d;
	};

	// One column per growth step: the pools the previous step's jigsaw
	// blocks draw from, with their fallbacks, which are tried in their
	// place. It stops when a step draws from exactly the pools of an
	// earlier step (other than the start, whose pieces are placed whole),
	// since from there growth repeats, and at the structure's size.
	const withFallbacks = (ids: Iterable<string>) => {
		const out: string[] = [];
		const add = (id: string) => {
			if (id === EMPTY || out.includes(id)) return;
			out.push(id);
			const fb = def(id).fallback;
			if (fb) add(fb);
		};
		for (const id of ids) add(id);
		return out;
	};
	const sameSet = (a: string[], b: string[]) => a.length === b.length && a.every((id) => b.includes(id));
	const maxColumns = graph.size !== undefined ? Math.min(graph.size + 2, 16) : 8;
	const columns: string[][] = [withFallbacks([start])];
	for (;;) {
		const next: string[] = [];
		for (const id of columns[columns.length - 1])
			for (const p of def(id).pieces) for (const j of p.jigsaws ?? []) if (j.pool !== EMPTY && !next.includes(j.pool)) next.push(j.pool);
		if (!next.length) break;
		const col = withFallbacks(next);
		const again = columns.findIndex((c, i) => i > 0 && sameSet(c, col));
		if (again >= 0) {
			graph.repeatsFrom = again;
			break;
		}
		if (columns.length >= maxColumns) break;
		columns.push(col);
	}

	const at = new Map<string, number>(); // id@depth -> index in graph.pools
	columns.forEach((col, depth) => {
		for (const id of col) {
			const d = def(id);
			at.set(`${id}@${depth}`, graph.pools.length);
			graph.pools.push({ ...d, depth, tooDeep: graph.size !== undefined && depth > graph.size, pieces: d.pieces.map((p) => ({ ...p, links: [], unreachable: false })) });
		}
	});
	// Where a jigsaw block at a depth draws from: its pool one step on,
	// or where growth repeats, the pool at the earlier depth.
	const find = (id: string, depth: number): number | undefined => {
		const i = at.get(`${id}@${depth}`);
		if (i !== undefined || graph.repeatsFrom === undefined) return i;
		for (let d = depth; d >= graph.repeatsFrom; d--) {
			const j = at.get(`${id}@${d}`);
			if (j !== undefined) return j;
		}
		return undefined;
	};

	// Link each jigsaw block to the pieces of its pool that have a jigsaw
	// block named its target. Problems count once per piece, however many
	// depths show it.
	const targeted = new Map<string, Set<string>>();
	const counted = new Set<string>();
	const once = (key: string) => !counted.has(key) && !!counted.add(key);
	const matches = (pool: Pool, target: string, link: Link) => {
		const pieces: number[] = [];
		pool.pieces.forEach((q, qi) => {
			if (q.jigsaws === undefined) link.maybe = true;
			else if (q.jigsaws.some((o) => o.name === target)) pieces.push(qi);
		});
		return pieces;
	};
	for (const pool of graph.pools) {
		if (pool.tooDeep) graph.problems.tooDeep++;
		pool.fallbackTo = pool.fallback === undefined ? undefined : at.get(`${pool.fallback}@${pool.depth}`);
		pool.pieces.forEach((p, pi) => {
			if (p.jigsaws === undefined && p.kind === 'single' && once(`u|${pool.id}|${pi}`)) graph.problems.unknown++;
			(p.jigsaws ?? []).forEach((j, ji) => {
				if (j.pool === EMPTY) return;
				const to = find(j.pool, pool.depth + 1);
				const target = to === undefined ? undefined : graph.pools[to];
				const link: Link = { jigsaw: ji, pool: j.pool, to, pieces: [], maybe: false };
				if (!def(j.pool).declared) {
					const ns = j.pool.split(':')[0];
					if (ns !== 'minecraft' && namespaces.has(ns)) link.dead = 'undeclared';
					else link.maybe = true;
				} else if (j.target !== EMPTY) {
					if (!targeted.has(j.pool)) targeted.set(j.pool, new Set());
					targeted.get(j.pool)!.add(j.target);
					if (target) link.pieces = matches(target, j.target, link);
					else link.maybe = true; // past the last depth shown
					// The game tries the pool's fallback when none of its pieces fit.
					const fb = target?.fallbackTo ?? (target?.fallback === undefined ? undefined : at.get(`${target.fallback}@${target.depth}`));
					if (target && !link.pieces.length && !link.maybe && fb !== undefined) {
						const f = graph.pools[fb];
						if (!f.declared && f.id.split(':')[0] === 'minecraft') link.maybe = true;
						const pieces = matches(f, j.target, link);
						if (pieces.length) link.fallback = { to: fb, pieces };
					}
					if (target && !link.pieces.length && !link.maybe && !link.fallback) link.dead = target.pieces.length ? 'no-target' : 'empty-pool';
				}
				if (link.dead && once(`d|${pool.id}|${pi}|${ji}`)) graph.problems.deadEnds++;
				p.links.push(link);
			});
		});
	}

	// A piece outside the start is only placed when some jigsaw block
	// drawing from its pool targets a name it has. A pool's fallback is
	// drawn from by the same jigsaw blocks as the pool.
	for (const [id, names] of [...targeted]) {
		const fb = def(id).fallback;
		if (!fb) continue;
		const into = targeted.get(fb) ?? new Set();
		for (const n of names) into.add(n);
		targeted.set(fb, into);
	}
	for (const [i, pool] of graph.pools.entries()) {
		if (i === 0) {
			// The start piece is placed whole; with a start_jigsaw_name only
			// pieces with that jigsaw block can start.
			if (root.kind === 'structure' && graph.startName) pool.pieces.forEach((p, pi) => markUnreachable(p, new Set([graph.startName!]), graph, once(`r|start|${pi}`)));
			continue;
		}
		const names = targeted.get(pool.id);
		if (names) pool.pieces.forEach((p, pi) => markUnreachable(p, names, graph, once(`r|${pool.id}|${pi}`)));
	}
	return graph;
}

function markUnreachable(p: Piece, names: Set<string>, graph: JigsawGraph, count: boolean) {
	if (p.kind !== 'single' || p.jigsaws === undefined) return;
	if (p.jigsaws.some((j) => names.has(j.name))) return;
	p.unreachable = true;
	if (count) graph.problems.unreachable++;
}

/** Plain words for a dead end. */
export function deadEndText(l: Link, target: string): string {
	switch (l.dead) {
		case 'undeclared':
			return `no TemplatePool declares ${l.pool}`;
		case 'empty-pool':
			return `${l.pool} has no pieces`;
		case 'no-target':
			return `no piece of ${l.pool} has a jigsaw block named ${target}`;
		default:
			return '';
	}
}
