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
	/** Steps from the start pool: the depth of the pieces it places. */
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

/** A pool element's pieces; a list element adds each of its elements. */
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
		case 'list_pool_element':
			// The game places a list's elements together, as one piece;
			// each is listed so its jigsaws show.
			for (const e of Array.isArray(el.elements) ? el.elements : []) elementPieces(e, weight, structures, out);
			break;
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

	// Breadth first from the start pool, so each pool sits at the
	// smallest depth it is reached at. A fallback is tried in place of
	// its pool, at the same depth.
	const index = new Map<string, number>();
	const queue: { id: string; depth: number }[] = [{ id: start, depth: 0 }];
	const add = (id: string, depth: number) => {
		if (id === EMPTY || index.has(id) || queue.some((q) => q.id === id)) return;
		queue.push({ id, depth });
	};
	while (queue.length) {
		const { id, depth } = queue.shift()!;
		if (index.has(id)) continue;
		const file = files.get(poolFile(id));
		const json = parse(file?.text);
		const pool: Pool = { id, declared: !!file, depth, tooDeep: graph.size !== undefined && depth > graph.size, pieces: [] };
		if (file) pool.path = file.path;
		if (isObj(json)) {
			if (typeof json.fallback === 'string' && fullId(json.fallback) !== EMPTY) pool.fallback = fullId(json.fallback);
			const raw: Omit<Piece, 'share' | 'links' | 'unreachable'>[] = [];
			for (const e of Array.isArray(json.elements) ? json.elements : []) {
				if (!isObj(e)) continue;
				elementPieces(e.element, typeof e.weight === 'number' ? e.weight : 1, structures, raw);
			}
			const total = raw.reduce((n, p) => n + p.weight, 0);
			pool.pieces = raw.map((p) => ({ ...p, share: total ? p.weight / total : 0, links: [], unreachable: false }));
		}
		index.set(id, graph.pools.length);
		graph.pools.push(pool);
		if (pool.fallback) add(pool.fallback, depth);
		for (const p of pool.pieces) for (const j of p.jigsaws ?? []) if (j.pool !== EMPTY) add(j.pool, depth + 1);
	}

	// Link each jigsaw block to the pieces of its pool that have a jigsaw
	// block named its target.
	const targeted = new Map<number, Set<string>>();
	for (const pool of graph.pools) {
		if (pool.tooDeep) graph.problems.tooDeep++;
		for (const p of pool.pieces) {
			if (p.jigsaws === undefined && p.kind === 'single') graph.problems.unknown++;
			(p.jigsaws ?? []).forEach((j, ji) => {
				if (j.pool === EMPTY) return;
				const to = index.get(j.pool);
				const target = to === undefined ? undefined : graph.pools[to];
				const link: Link = { jigsaw: ji, pool: j.pool, to, pieces: [], maybe: false };
				if (!target?.declared) {
					const ns = j.pool.split(':')[0];
					if (ns !== 'minecraft' && namespaces.has(ns)) link.dead = 'undeclared';
					else link.maybe = true;
				} else if (j.target !== EMPTY) {
					if (!targeted.has(to!)) targeted.set(to!, new Set());
					targeted.get(to!)!.add(j.target);
					target.pieces.forEach((q, qi) => {
						if (q.jigsaws === undefined) link.maybe = true;
						else if (q.jigsaws.some((o) => o.name === j.target)) link.pieces.push(qi);
					});
					if (!link.pieces.length && !link.maybe) link.dead = target.pieces.length ? 'no-target' : 'empty-pool';
				}
				if (link.dead) graph.problems.deadEnds++;
				p.links.push(link);
			});
		}
	}

	// A piece outside the start pool is only placed when some jigsaw
	// block drawing from its pool targets a name it has. A pool's
	// fallback is drawn from by the same jigsaw blocks as the pool.
	for (const [i, pool] of graph.pools.entries()) {
		if (pool.fallback && index.has(pool.fallback)) {
			pool.fallbackTo = index.get(pool.fallback);
			const names = targeted.get(i);
			if (names) {
				const into = targeted.get(pool.fallbackTo!) ?? new Set();
				for (const n of names) into.add(n);
				targeted.set(pool.fallbackTo!, into);
			}
		}
	}
	for (const [i, pool] of graph.pools.entries()) {
		if (i === 0 && root.kind === 'structure') {
			// The start piece is placed whole; with a start_jigsaw_name only
			// pieces with that jigsaw block can start.
			if (graph.startName) for (const p of pool.pieces) markUnreachable(p, new Set([graph.startName]), graph);
			continue;
		}
		if (i === 0) continue;
		const names = targeted.get(i);
		if (!names) continue;
		for (const p of pool.pieces) markUnreachable(p, names, graph);
	}
	return graph;
}

function markUnreachable(p: Piece, names: Set<string>, graph: JigsawGraph) {
	if (p.kind !== 'single' || p.jigsaws === undefined) return;
	if (p.jigsaws.some((j) => names.has(j.name))) return;
	p.unreachable = true;
	graph.problems.unreachable++;
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
