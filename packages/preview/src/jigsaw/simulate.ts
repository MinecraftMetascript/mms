// A best-effort simulation of how the game grows a jigsaw structure, for
// the jigsaw simulation preview: which pieces it places, where, and turned
// which way, for a given layout seed.
//
// It follows the game's jigsaw placement step by step: a random rotation
// and start piece, then a queue of placed pieces (by placement priority,
// then first in, first out). For each piece, its jigsaw blocks are
// shuffled and sorted by selection priority; for each jigsaw block, the
// pool's pieces (one entry per weight) and then its fallback's are
// shuffled, and each candidate is tried in the four rotations, shuffled,
// joined by each of its jigsaw blocks that faces this one with the target
// name. A candidate is placed when its box fits the free space: inside the
// max distance box (or inside its parent's box, when the jigsaw block
// points into the parent), and clear of every piece placed in that space.
// Random numbers come from Java's LCG, as the game's do.
//
// Left out on purpose: terrain (terrain_matching pieces are placed like
// rigid ones, and the start piece is not projected to a heightmap), the
// expansion hack, pool aliases, processors and world height limits. So a
// layout seed is not a world seed: the same layout seed gives the same
// layout here, but not the one a world shows in some chunk.
//
// Free of DOM types, so it can be tested and reused anywhere.

import type { OutputAsset, OutputFile, OutputJigsaw } from '../types';
import { fullId } from './graph';

/** A box of blocks, min and max corners, both inclusive. */
export type Box = [[number, number, number], [number, number, number]];

/** 0: none, 1: clockwise 90°, 2: 180°, 3: counterclockwise 90°, as the game numbers them. */
export type Rotation = 0 | 1 | 2 | 3;

export interface PlacedJigsaw {
	/** Where it is in the world. */
	pos: [number, number, number];
	/** Its orientation once its piece is turned, such as east_up. */
	orientation: string;
	pool: string;
	name: string;
	target: string;
	/** The block it turns into once the structure is placed. */
	finalState: string;
	/**
	 * What happened when the game grew from it: the piece joined here, or
	 * why nothing was. A jigsaw block a child joined by has `joined`.
	 */
	result?: { piece: number } | { none: NoneReason };
	/** The parent's jigsaw block this piece joined by, for the piece's own joining block. */
	joined?: boolean;
}

/** Why a jigsaw block placed nothing. */
export type NoneReason =
	/** Its pool is minecraft:empty: an end on purpose. */
	| 'empty'
	/** No pool of that name is in the pack (a vanilla pool, or undeclared). */
	| 'unknown-pool'
	/** The piece is at the structure's size, and the fallback had nothing that fit. */
	| 'depth'
	/** No piece had a jigsaw block that could join, or none fit the free space. */
	| 'no-fit'
	/** The pool drew the empty piece. */
	| 'empty-piece';

export interface PlacedPiece {
	/** The structure id of the piece, or a description for a feature. */
	id: string;
	kind: 'single' | 'feature' | 'list';
	/** For a list element, the structures it places together. */
	parts?: string[];
	/** The pool it came from. */
	pool: string;
	depth: number;
	rotation: Rotation;
	/** Where its structure's (0, 0, 0) block lands. */
	origin: [number, number, number];
	box: Box;
	jigsaws: PlacedJigsaw[];
	/** The piece it joined onto, and that piece's jigsaw block, by index. */
	parent?: { piece: number; jigsaw: number };
}

export interface Simulation {
	seed: number;
	pieces: PlacedPiece[];
	/** The box no piece may leave, from max_distance_from_center. */
	bounds: Box;
	/** The structure's size: how many pieces deep it grows. */
	depth: number;
	/** The deepest piece placed. */
	reached: number;
	/** Why there is no structure at all, when there is none. */
	failed?: string;
	/** Things the simulation could not follow, in plain words, once each. */
	notes: string[];
}

type Json = null | boolean | number | string | Json[] | { [k: string]: Json };
type Obj = { [k: string]: Json };
const isObj = (v: Json | undefined): v is Obj => typeof v === 'object' && v !== null && !Array.isArray(v);
const EMPTY = 'minecraft:empty';

/** Java's java.util.Random, which the game's legacy random source is. */
export class JavaRandom {
	private hi = 0; // the seed's high 24 bits
	private lo = 0; // the seed's low 24 bits

	constructor(seed: number | bigint) {
		const s = (BigInt.asUintN(64, BigInt(seed)) ^ 0x5deece66dn) & ((1n << 48n) - 1n);
		this.hi = Number(s >> 24n);
		this.lo = Number(s & 0xffffffn);
	}

	next(bits: number): number {
		// seed = seed * 0x5DEECE66D + 0xB, mod 2^48, in two 24-bit halves.
		const p = this.lo * 0xece66d + 0xb;
		const lo = p % 0x1000000;
		const carry = Math.floor(p / 0x1000000);
		const hi = (this.hi * 0xece66d + this.lo * 0x5de + carry) % 0x1000000;
		this.hi = hi;
		this.lo = lo;
		const r = Math.floor((hi * 0x1000000 + lo) / 2 ** (48 - bits));
		return r >= 2 ** (bits - 1) && bits === 32 ? r - 2 ** 32 : r;
	}

	nextInt(bound: number): number {
		if ((bound & -bound) === bound) return Number((BigInt(bound) * BigInt(this.next(31))) >> 31n);
		for (;;) {
			const bits = this.next(31);
			const val = bits % bound;
			if (bits - val + (bound - 1) < 2 ** 31) return val;
		}
	}
}

/** The game's Util.shuffle: Fisher–Yates from the end, in place. */
function shuffle<T>(list: T[], random: JavaRandom): T[] {
	for (let i = list.length; i > 1; i--) {
		const j = random.nextInt(i);
		[list[i - 1], list[j]] = [list[j], list[i - 1]];
	}
	return list;
}

const HORIZONTAL = ['north', 'east', 'south', 'west'];
const OPPOSITE: Record<string, string> = { north: 'south', south: 'north', east: 'west', west: 'east', up: 'down', down: 'up' };
const STEP: Record<string, [number, number, number]> = { north: [0, 0, -1], south: [0, 0, 1], east: [1, 0, 0], west: [-1, 0, 0], up: [0, 1, 0], down: [0, -1, 0] };

function turnDir(d: string, r: Rotation): string {
	const i = HORIZONTAL.indexOf(d);
	if (i < 0) return d;
	const by = r === 3 ? 3 : r; // counterclockwise 90° = three quarter turns clockwise
	return HORIZONTAL[(i + by) % 4];
}

/** An orientation such as north_up, turned about the y axis. */
export function turnOrientation(o: string, r: Rotation): string {
	const [front, top] = o.split('_');
	return `${turnDir(front, r)}_${turnDir(top ?? 'up', r)}`;
}

/** A position in a structure, turned about its (0, 0, 0) as the game turns it. */
export function turnPos([x, y, z]: readonly number[], r: Rotation): [number, number, number] {
	switch (r) {
		case 1:
			return [-z, y, x];
		case 2:
			return [-x, y, -z];
		case 3:
			return [z, y, -x];
		default:
			return [x, y, z];
	}
}

function sizeBox(size: readonly number[], r: Rotation, at: readonly number[]): Box {
	const a = turnPos([0, 0, 0], r);
	const b = turnPos([size[0] - 1, size[1] - 1, size[2] - 1], r);
	return [
		[Math.min(a[0], b[0]) + at[0], Math.min(a[1], b[1]) + at[1], Math.min(a[2], b[2]) + at[2]],
		[Math.max(a[0], b[0]) + at[0], Math.max(a[1], b[1]) + at[1], Math.max(a[2], b[2]) + at[2]]
	];
}

const inside = (b: Box, p: readonly number[]) => p.every((v, i) => v >= b[0][i] && v <= b[1][i]);
const within = (inner: Box, outer: Box) => inner[0].every((v, i) => v >= outer[0][i]) && inner[1].every((v, i) => v <= outer[1][i]);
const overlaps = (a: Box, b: Box) => a[0].every((v, i) => v <= b[1][i]) && a[1].every((v, i) => v >= b[0][i]);
const moved = (b: Box, d: readonly number[]): Box => [b[0].map((v, i) => v + d[i]) as Box[0], b[1].map((v, i) => v + d[i]) as Box[1]];

/** The space pieces may still go in: a box less the boxes placed in it. Shared, as the game shares it. */
interface Free {
	box: Box;
	taken: Box[];
}
const fits = (free: Free, b: Box) => within(b, free.box) && !free.taken.some((t) => overlaps(t, b));

/** A jigsaw block of an element, relative to where the element is placed. */
interface Jig {
	pos: [number, number, number];
	orientation: string;
	pool: string;
	name: string;
	target: string;
	joint: string;
	finalState: string;
	selection: number;
	placement: number;
}

/** A pool element the simulation can place. */
interface Element {
	kind: 'single' | 'feature' | 'list' | 'empty' | 'unknown';
	id: string;
	parts?: string[];
	/** The element's box placed at at, turned by r. */
	box(at: readonly number[], r: Rotation): Box;
	/** Its jigsaw blocks turned by r, relative to its (0, 0, 0). */
	jigsaws(r: Rotation): Jig[];
}

const front = (o: string) => o.split('_')[0];
const top = (o: string) => o.split('_')[1] ?? 'up';

function jigOf(j: OutputJigsaw): Jig {
	const orientation = j.orientation ?? 'north_up';
	const f = front(orientation);
	return {
		pos: [...j.pos],
		orientation,
		pool: fullId(j.pool),
		name: fullId(j.name),
		target: fullId(j.target),
		joint: j.joint ?? (f === 'up' || f === 'down' ? 'rollable' : 'aligned'),
		finalState: j.final_state ?? 'minecraft:air',
		selection: j.selection_priority ?? 0,
		placement: j.placement_priority ?? 0
	};
}

const turnJig = (j: Jig, r: Rotation): Jig => ({ ...j, pos: turnPos(j.pos, r), orientation: turnOrientation(j.orientation, r) });

/** The game's JigsawBlock.canAttach: facing each other, the same way up unless rollable, and the target name. */
const canAttach = (a: Jig, b: Jig) => front(a.orientation) === OPPOSITE[front(b.orientation)] && (a.joint === 'rollable' || top(a.orientation) === top(b.orientation)) && a.target === b.name;

/** The jigsaw block the game gives a feature element: named minecraft:bottom, facing down. */
const featureJig: Jig = { pos: [0, 0, 0], orientation: 'down_south', pool: EMPTY, name: 'minecraft:bottom', target: EMPTY, joint: 'rollable', finalState: 'minecraft:air', selection: 0, placement: 0 };

const nbtPath = /^data\/([^/]+)\/structures?\/(.+)\.nbt$/;

/** The structures by id, from the pack's copied structure files. */
export function structureFiles(assets: readonly OutputAsset[]): Map<string, OutputAsset> {
	const out = new Map<string, OutputAsset>();
	for (const a of assets) {
		const m = nbtPath.exec(a.path);
		if (m) out.set(`${m[1]}:${m[2]}`, a);
	}
	return out;
}

function element(el: Json | undefined, structures: Map<string, OutputAsset>): Element {
	const unknown = (id: string): Element => ({ kind: 'unknown', id, box: (at) => [[...at] as Box[0], [...at] as Box[1]], jigsaws: () => [] });
	if (!isObj(el)) return unknown('?');
	const type = String(el.element_type ?? '').replace(/^minecraft:/, '');
	switch (type) {
		case 'single_pool_element':
		case 'legacy_single_pool_element': {
			const id = fullId(String(el.location ?? '?'));
			const a = structures.get(id);
			if (!a?.size || !a.jigsaws) return unknown(id);
			const size = a.size;
			const jigs = a.jigsaws.map(jigOf);
			return { kind: 'single', id, box: (at, r) => sizeBox(size, r, at), jigsaws: (r) => jigs.map((j) => turnJig(j, r)) };
		}
		case 'list_pool_element': {
			// Placed together as one piece, joined by the first element's jigsaw blocks.
			const parts = (Array.isArray(el.elements) ? el.elements : []).map((e) => element(e, structures));
			if (!parts.length) return { kind: 'empty', id: 'empty list', box: (at) => [[...at] as Box[0], [...at] as Box[1]], jigsaws: () => [] };
			const bad = parts.find((p) => p.kind === 'unknown');
			if (bad) return unknown(bad.id);
			return {
				kind: 'list',
				id: parts.map((p) => p.id).join(' + '),
				parts: parts.map((p) => p.id),
				box: (at, r) => {
					const bs = parts.map((p) => p.box(at, r));
					return [[0, 1, 2].map((i) => Math.min(...bs.map((b) => b[0][i]))) as Box[0], [0, 1, 2].map((i) => Math.max(...bs.map((b) => b[1][i]))) as Box[1]];
				},
				jigsaws: (r) => parts[0].jigsaws(r)
			};
		}
		case 'feature_pool_element': {
			const id = `feature ${String(el.feature ?? '?').replace(/^minecraft:/, '')}`;
			return { kind: 'feature', id, box: (at) => [[...at] as Box[0], [...at] as Box[1]], jigsaws: () => [featureJig] };
		}
		case 'empty_pool_element':
			return { kind: 'empty', id: 'empty', box: (at) => [[...at] as Box[0], [...at] as Box[1]], jigsaws: () => [] };
		default:
			return unknown(type || '?');
	}
}

interface Pool {
	id: string;
	/** Each element once per weight, as the game keeps them. */
	templates: Element[];
	fallback: string;
}

function parse(text: string | undefined): Json | undefined {
	if (text === undefined) return undefined;
	try {
		return JSON.parse(text);
	} catch {
		return undefined;
	}
}

/** Options for a simulation; the structure's own values are used for the rest. */
export interface SimulateOptions {
	seed?: number;
	/** At most this many pieces, so a structure that grows without end still returns. */
	limit?: number;
}

/**
 * Simulates a jigsaw structure at an output path, or returns null when the
 * path is not a jigsaw structure in the output.
 */
export function simulate(path: string, output: readonly OutputFile[], assets: readonly OutputAsset[] = [], opts: SimulateOptions = {}): Simulation | null {
	const files = new Map(output.map((f) => [f.path, f]));
	const json = parse(files.get(path)?.text);
	if (!isObj(json) || String(json.type ?? '').replace(/^minecraft:/, '') !== 'jigsaw' || typeof json.start_pool !== 'string') return null;
	const seed = opts.seed ?? 0;
	const limit = opts.limit ?? 4096;
	const depth = typeof json.size === 'number' ? json.size : 7;
	const maxDistance = typeof json.max_distance_from_center === 'number' ? json.max_distance_from_center : 80;
	const startName = typeof json.start_jigsaw_name === 'string' ? fullId(json.start_jigsaw_name) : undefined;
	const structures = structureFiles(assets);
	const noted = new Set<string>();
	const sim: Simulation = { seed, pieces: [], bounds: [[0, 0, 0], [0, 0, 0]], depth, reached: 0, notes: [] };
	const note = (s: string) => {
		if (!noted.has(s)) {
			noted.add(s);
			sim.notes.push(s);
		}
	};
	if (json.use_expansion_hack === true) note('The expansion hack is not simulated.');
	if (Array.isArray(json.pool_aliases) && json.pool_aliases.length) note('Pool aliases are not simulated; jigsaw blocks draw from the pools they name.');

	const pools = new Map<string, Pool | null>();
	const pool = (id: string): Pool | null => {
		if (pools.has(id)) return pools.get(id)!;
		let p: Pool | null = null;
		if (id === EMPTY) p = { id, templates: [], fallback: EMPTY };
		else {
			const [ns, name] = id.split(':', 2);
			const j = parse(files.get(`data/${ns}/worldgen/template_pool/${name}.json`)?.text);
			if (isObj(j)) {
				const templates: Element[] = [];
				for (const e of Array.isArray(j.elements) ? j.elements : []) {
					if (!isObj(e)) continue;
					const el = element(e.element, structures);
					if (el.kind === 'unknown') note(`${el.id} is not one of this project's structure files, so it is never placed here.`);
					const w = typeof e.weight === 'number' ? e.weight : 1;
					for (let i = 0; i < w; i++) templates.push(el);
				}
				p = { id, templates, fallback: typeof j.fallback === 'string' ? fullId(j.fallback) : EMPTY };
			}
		}
		pools.set(id, p);
		return p;
	};

	const random = new JavaRandom(seed);
	// The start: a rotation, then a piece from the start pool.
	const rotation = random.nextInt(4) as Rotation;
	const startPool = pool(fullId(json.start_pool));
	if (!startPool || !startPool.templates.length) {
		sim.failed = startPool ? `The start pool ${json.start_pool} has no pieces.` : `The start pool ${json.start_pool} is not in the pack.`;
		return sim;
	}
	const start = startPool.templates[random.nextInt(startPool.templates.length)];
	if (start.kind === 'empty') {
		sim.failed = 'The start pool drew the empty piece, so nothing is placed for this seed.';
		return sim;
	}
	if (start.kind === 'unknown') {
		sim.failed = `The start piece ${start.id} is not one of this project's structure files.`;
		return sim;
	}
	let offset: [number, number, number] = [0, 0, 0];
	if (startName) {
		const named = shuffled(start, rotation, random).find((j) => j.name === startName);
		if (!named) {
			sim.failed = `The start piece ${start.id} has no jigsaw block named ${startName}.`;
			return sim;
		}
		offset = named.pos;
	}
	const startAt: [number, number, number] = [-offset[0], -offset[1], -offset[2]];
	let box = start.box(startAt, rotation);
	const cx = Math.trunc((box[1][0] + box[0][0]) / 2);
	const cz = Math.trunc((box[1][2] + box[0][2]) / 2);
	// The game moves the start piece so its ground level (a block above its
	// bottom) is at the start height.
	const dy = startAt[1] - (box[0][1] + 1);
	box = moved(box, [0, dy, 0]);
	const origin: [number, number, number] = [startAt[0], startAt[1] + dy, startAt[2]];
	const cy = startAt[1] + offset[1];
	sim.bounds = [
		[cx - maxDistance, cy - maxDistance, cz - maxDistance],
		[cx + maxDistance, cy + maxDistance, cz + maxDistance]
	];

	type State = { piece: number; free: Free; depth: number; priority: number; order: number };
	const queue: State[] = [];
	let order = 0;
	const add = (piece: PlacedPiece, jigs: Jig[]) => {
		piece.jigsaws = jigs.map((j) => ({ pos: j.pos, orientation: j.orientation, pool: j.pool, name: j.name, target: j.target, finalState: j.finalState }));
		sim.pieces.push(piece);
		sim.reached = Math.max(sim.reached, piece.depth);
		return sim.pieces.length - 1;
	};
	const startJigs = start.jigsaws(rotation).map((j) => ({ ...j, pos: j.pos.map((v, i) => v + origin[i]) as Jig['pos'] }));
	add({ id: start.id, kind: start.kind, ...(start.parts ? { parts: start.parts } : {}), pool: startPool.id, depth: 0, rotation, origin, box, jigsaws: [] }, startJigs);
	const elements: Element[] = [start];
	if (depth > 0) queue.push({ piece: 0, free: { box: sim.bounds, taken: [box] }, depth: 0, priority: 0, order: order++ });

	while (queue.length) {
		// Highest placement priority first, then first in, first out.
		let best = 0;
		for (let i = 1; i < queue.length; i++) if (queue[i].priority > queue[best].priority) best = i;
		const state = queue.splice(best, 1)[0];
		const piece = sim.pieces[state.piece];
		const el = elements[state.piece];
		const pbox = piece.box;
		let interior: Free | undefined;
		// The piece's jigsaw blocks, shuffled the way the game shuffles them,
		// in the world; the indices tie them back to piece.jigsaws.
		const own = el.jigsaws(piece.rotation).map((j, i) => ({ j: { ...j, pos: j.pos.map((v, k) => v + piece.origin[k]) as Jig['pos'] }, i }));
		shuffle(own, random);
		own.sort((a, b) => b.j.selection - a.j.selection);
		outer: for (const { j, i: ji } of own) {
			// The block this piece joined by is tried too, as the game does;
			// nothing can fit back into the parent, but it uses up random numbers.
			const out = piece.jigsaws[ji];
			const set = (r: NonNullable<PlacedJigsaw['result']>) => {
				if (!out.joined) out.result = r;
			};
			const f = front(j.orientation);
			const step = STEP[f] ?? [0, 0, 0];
			const targetPos: [number, number, number] = [j.pos[0] + step[0], j.pos[1] + step[1], j.pos[2] + step[2]];
			const p = pool(j.pool);
			if (!p || (!p.templates.length && p.id !== EMPTY)) {
				set({ none: j.pool === EMPTY ? 'empty' : 'unknown-pool' });
				if (!p && !j.pool.startsWith('minecraft:')) note(`No TemplatePool declares ${j.pool}.`);
				else if (!p) note(`${j.pool} is a vanilla pool, which the simulation does not have.`);
				continue;
			}
			const fb = pool(p.fallback);
			if (!fb || (!fb.templates.length && fb.id !== EMPTY)) {
				set({ none: 'unknown-pool' });
				continue;
			}
			if (j.pool === EMPTY) {
				set({ none: 'empty' });
				continue;
			}
			let free: Free;
			if (inside(pbox, targetPos)) free = interior ??= { box: pbox, taken: [] };
			else free = state.free;
			const candidates: { c: Element; from: string }[] = [];
			if (state.depth !== depth) for (const c of shuffle([...p.templates], random)) candidates.push({ c, from: p.id });
			for (const c of shuffle([...fb.templates], random)) candidates.push({ c, from: fb.id });
			let why: NoneReason = state.depth === depth ? 'depth' : 'no-fit';
			for (const { c, from } of candidates) {
				if (c.kind === 'empty') {
					why = 'empty-piece';
					break;
				}
				for (const r of shuffle([0, 1, 2, 3] as Rotation[], random)) {
					const cj = shuffled(c, r, random);
					for (const k of cj) {
						if (!canAttach(j, k)) continue;
						const at: [number, number, number] = [targetPos[0] - k.pos[0], targetPos[1] - k.pos[1], targetPos[2] - k.pos[2]];
						const cbox = c.box(at, r);
						if (!fits(free, cbox)) continue;
						free.taken.push(cbox);
						if (sim.pieces.length >= limit) {
							note(`The simulation stops at ${limit} pieces.`);
							break outer;
						}
						const jigs = c.jigsaws(r).map((q) => ({ ...q, pos: q.pos.map((v, n) => v + at[n]) as Jig['pos'] }));
						const idx = add({ id: c.id, kind: c.kind as PlacedPiece['kind'], ...(c.parts ? { parts: c.parts } : {}), pool: from, depth: state.depth + 1, rotation: r, origin: at, box: cbox, jigsaws: [], parent: { piece: state.piece, jigsaw: ji } }, jigs);
						elements.push(c);
						const back = sim.pieces[idx].jigsaws.findIndex((q) => q.pos.every((v, n) => v === k.pos[n] + at[n]));
						if (back >= 0) sim.pieces[idx].jigsaws[back].joined = true;
						set({ piece: idx });
						if (state.depth + 1 <= depth) queue.push({ piece: idx, free, depth: state.depth + 1, priority: j.placement, order: order++ });
						continue outer;
					}
				}
			}
			set({ none: why });
		}
	}
	return sim;
}

/** An element's jigsaw blocks turned by r and shuffled, then sorted by selection priority, as the game lists them. */
function shuffled(el: Element, r: Rotation, random: JavaRandom): Jig[] {
	const list = shuffle(el.jigsaws(r), random);
	return list.sort((a, b) => b.selection - a.selection);
}
