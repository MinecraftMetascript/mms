// The blocks of a simulated jigsaw structure, for the simulation's block
// view: each placed piece's structure file read from its .nbt, turned and
// moved as the piece was, with jigsaw blocks swapped for their final state.
// Blocks are drawn as plain coloured cubes, so only block names are kept.

import { NbtFile, NbtType, type NbtCompound } from 'deepslate';
import { turnPos, type Simulation } from './simulate';

/** A structure file's blocks: a palette of names and, per block, x, y, z and its palette index. */
export interface StructureBlocks {
	palette: string[];
	blocks: Int32Array;
}

/** Blocks that are not drawn: air, and the void a structure block leaves alone. */
const hidden = new Set(['minecraft:air', 'minecraft:cave_air', 'minecraft:void_air', 'minecraft:structure_void']);

const cache = new Map<string, StructureBlocks | null>();

function decode(url: string): Uint8Array {
	const comma = url.indexOf(',');
	const bin = atob(url.slice(comma + 1));
	const out = new Uint8Array(bin.length);
	for (let i = 0; i < bin.length; i++) out[i] = bin.charCodeAt(i);
	return out;
}

/** Reads a structure file from a data: URL; null when it can't be read. Cached by URL. */
export function readStructure(url: string): StructureBlocks | null {
	const hit = cache.get(url);
	if (hit !== undefined) return hit;
	let out: StructureBlocks | null = null;
	try {
		const root = NbtFile.read(decode(url)).root;
		// Palettes are {Name, Properties} up to 1.21 and {id, properties} in 26.x.
		let paletteTag = root.getList('palette', NbtType.Compound);
		if (!paletteTag.length && root.has('palettes')) paletteTag = root.getList('palettes', NbtType.List).getList(0, NbtType.Compound);
		const states: NbtCompound[] = [];
		paletteTag.forEach((t) => states.push(t));
		const palette = states.map((s) => s.getString('Name') || s.getString('id') || 'minecraft:air');
		const blocks: number[] = [];
		root.getList('blocks', NbtType.Compound).forEach((t) => {
			const state = t.getNumber('state');
			const pos = t.getList('pos');
			const name = palette[state];
			if (name === undefined || hidden.has(name)) return;
			let index = state;
			if (name === 'minecraft:jigsaw') {
				// A jigsaw block turns into its final state once placed.
				const final = t.getCompound('nbt').getString('final_state') || 'minecraft:air';
				const id = final.replace(/[[{].*$/, '');
				if (hidden.has(id)) return;
				index = palette.indexOf(id);
				if (index < 0) index = palette.push(id) - 1;
			}
			blocks.push(pos.getNumber(0), pos.getNumber(1), pos.getNumber(2), index);
		});
		out = { palette, blocks: Int32Array.from(blocks) };
	} catch {
		out = null;
	}
	cache.set(url, out);
	if (cache.size > 256) cache.delete(cache.keys().next().value!);
	return out;
}

/** A placed block: where it is, and its name. */
export interface PlacedBlocks {
	names: string[];
	/** Per block: x, y, z, index into names, and the piece it belongs to. */
	blocks: Int32Array;
	/** Pieces whose structure file could not be read, by id. */
	missing: string[];
}

/**
 * The blocks of every piece of a simulation. `structures` maps a structure
 * id (such as `dungeon:halls/hall`) to its .nbt as a data: URL. Later
 * pieces overwrite earlier ones where they share a block.
 */
export function placeBlocks(sim: Simulation, structures: ReadonlyMap<string, string>): PlacedBlocks {
	const names: string[] = [];
	const nameIndex = new Map<string, number>();
	const at = new Map<string, number>(); // "x,y,z" -> offset into out
	const out: number[] = [];
	const missing = new Set<string>();
	sim.pieces.forEach((piece, pi) => {
		if (piece.kind === 'feature') return;
		for (const id of piece.parts ?? [piece.id]) {
			const url = structures.get(id);
			const s = url ? readStructure(url) : null;
			if (!s) {
				missing.add(id);
				continue;
			}
			const map = s.palette.map((n) => {
				let i = nameIndex.get(n);
				if (i === undefined) {
					i = names.push(n) - 1;
					nameIndex.set(n, i);
				}
				return i;
			});
			for (let i = 0; i < s.blocks.length; i += 4) {
				const p = turnPos([s.blocks[i], s.blocks[i + 1], s.blocks[i + 2]], piece.rotation);
				const x = p[0] + piece.origin[0], y = p[1] + piece.origin[1], z = p[2] + piece.origin[2];
				const key = `${x},${y},${z}`;
				const o = at.get(key);
				if (o !== undefined) {
					out[o + 3] = map[s.blocks[i + 3]];
					out[o + 4] = pi;
				} else {
					at.set(key, out.length);
					out.push(x, y, z, map[s.blocks[i + 3]], pi);
				}
			}
		}
	});
	return { names, blocks: Int32Array.from(out), missing: [...missing] };
}
