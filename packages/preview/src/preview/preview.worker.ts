/// <reference lib="webworker" />
import {
	BiomeSource,
	BlockPos,
	Chunk,
	ChunkPos,
	DensityFunction,
	BlockState,
	DensityVolume,
	FixedBiomeSource,
	MaterialRuleContext,
	Holder,
	HolderSet,
	Identifier,
	NoiseChunk,
	NoiseChunkGenerator,
	NoiseGeneratorSettings,
	NoiseRouter,
	RandomState,
	WorldgenRegistries,
	type MaterialRule
} from 'deepslate';
import { biomeColor, blockColor, CAVE, SKY, VOID, type Rgb } from './colors';
import { speedUp } from './fastDensity';
import { screenAxes, vanillaSummary, type PreviewFile, type PreviewRequest, type PreviewResponse, type PreviewView } from './protocol';

// Renders previews with Deepslate. User files are registered under their
// ids; minecraft: ids the pack references but does not define come from
// mcmeta's per-registry bundles, fetched on demand and kept for the session.

type Folder = Exclude<PreviewFile['kind'], 'dimension'>;
type RenderRequest = Extract<PreviewRequest, { type: 'render' }>;
type Rendered = Extract<PreviewResponse, { type: 'rendered' }>;

const registries = {
	noise: WorldgenRegistries.NOISE,
	density_function: WorldgenRegistries.DENSITY_FUNCTION,
	noise_settings: WorldgenRegistries.NOISE_SETTINGS,
	material_rule: WorldgenRegistries.MATERIAL_RULE,
	material_condition: WorldgenRegistries.MATERIAL_CONDITION
} as const;

// Surface rules only ask whether a biome is in a set, so every biome id
// resolves; what a biome contains does not matter here.
const biomeRegistry = WorldgenRegistries.BIOME as unknown as { get(id: Identifier): unknown; getOrThrow(id: Identifier): unknown };
biomeRegistry.get = biomeRegistry.getOrThrow = () => ({});
const biomeTags = WorldgenRegistries.BIOME.getTagRegistry();

let userFiles: PreviewFile[] = [];
let dimensions = new Map<string, unknown>();
let version = 0; // bumped when the files change, to drop cached terrain
/**
 * Bumped only when files that shape the terrain change (noises, density
 * functions, noise settings and dimensions, ignoring surface rules), so
 * editing a surface rule reuses the filled chunks.
 */
let shapeVersion = 0;
let shapeSignature = '';
function shapeOf(files: PreviewFile[]): string {
	const shaping = files.filter((f) => ['noise', 'density_function', 'noise_settings', 'dimension'].includes(f.kind));
	return JSON.stringify(shaping, (k, v) => (k === 'surface_rule' ? undefined : v));
}
const bundles = new Map<Folder, Record<string, unknown> | null>();

async function fetchBundle(folder: Folder): Promise<boolean> {
	if (!bundles.has(folder)) {
		let json: Record<string, unknown> | null = null;
		try {
			const path = folder.startsWith('tag/') ? `tag/${folder.slice(4)}` : `worldgen/${folder}`;
			const res = await fetch(`${vanillaSummary}/${path}/data.min.json`);
			if (res.ok) json = await res.json();
		} catch {
			json = null;
		}
		bundles.set(folder, json);
	}
	return bundles.get(folder) != null;
}

function register(kind: Folder, id: string, json: unknown) {
	if (kind === 'tag/worldgen/biome') {
		biomeTags.register(Identifier.parse(id), () => HolderSet.fromJson(WorldgenRegistries.BIOME, json));
		return;
	}
	const r = registries[kind] as unknown as { register(id: Identifier, v: () => unknown): void; parse(o: unknown): unknown };
	r.register(Identifier.parse(id), () => r.parse(json));
}

function registerAll() {
	biomeTags.clear();
	for (const r of Object.values(registries)) {
		// Builtin entries (the surface noises) survive; everything else
		// is re-registered from the current files.
		for (const key of r.keys()) {
			if (!(key.equals(Identifier.create('surface')) || key.equals(Identifier.create('surface_secondary')))) r.delete(key);
		}
	}
	for (const [folder, bundle] of bundles) {
		if (!bundle) continue;
		for (const [path, json] of Object.entries(bundle)) register(folder, `minecraft:${path}`, json);
	}
	for (const f of userFiles) {
		if (f.kind !== 'dimension') register(f.kind, f.id, f.json);
	}
}

/** Runs build, fetching the vanilla bundle of each missing key it reports, until it succeeds. */
async function withVanilla<T>(build: () => T): Promise<T> {
	for (;;) {
		registerAll();
		try {
			return build();
		} catch (e) {
			const m = /Missing key in (?:minecraft:)?(worldgen\/[a-z_]+|tags\/worldgen\/biome): minecraft:/.exec(String(e));
			const folder = (m?.[1] === 'tags/worldgen/biome' ? 'tag/worldgen/biome' : m?.[1].replace('worldgen/', '')) as Folder | undefined;
			if (!folder || !(folder in registries || folder === 'tag/worldgen/biome') || bundles.has(folder) || !(await fetchBundle(folder))) throw e;
		}
	}
}

function seedOf(s: string): bigint {
	try {
		return BigInt(s);
	} catch {
		let h = 0n;
		for (const c of s) h = (h * 31n + BigInt(c.codePointAt(0)!)) & 0xffffffffffffffffn;
		return BigInt.asIntN(64, h);
	}
}

// Diverging palette for density values: blue below zero, orange above.
function diverging(v: number, out: Uint8ClampedArray, i: number) {
	const t = Math.max(-1, Math.min(1, v));
	if (t < 0) {
		const a = -t;
		out[i] = 245 - 200 * a;
		out[i + 1] = 245 - 140 * a;
		out[i + 2] = 245 - 30 * a;
	} else {
		out[i] = 245 - 10 * t;
		out[i + 1] = 245 - 120 * t;
		out[i + 2] = 245 - 200 * t;
	}
	out[i + 3] = 255;
}

function put(out: Uint8ClampedArray, i: number, c: Rgb, shade = 1) {
	out[i] = c[0] * shade;
	out[i + 1] = c[1] * shade;
	out[i + 2] = c[2] * shade;
	out[i + 3] = 255;
}

/** Yields to the message loop, so a newer request can supersede this one. */
const breathe = () => new Promise((r) => setTimeout(r, 0));
class Superseded extends Error {}

// ---- Dimensions -------------------------------------------------------

interface Dimension {
	settings: NoiseGeneratorSettings;
	biomes: BiomeSource;
}

function readDimension(id: string): Dimension {
	const json = dimensions.get(id) as { generator?: { type?: string; settings?: unknown; biome_source?: unknown } } | undefined;
	if (!json) throw new Error(`Missing dimension ${id}`);
	const gen = json.generator ?? {};
	const type = String(gen.type ?? '').replace(/^minecraft:/, '');
	if (type !== 'noise') throw new Error(`Only noise generators have a preview; this one is ${type || 'missing'}.`);
	const settings =
		typeof gen.settings === 'string'
			? WorldgenRegistries.NOISE_SETTINGS.getOrThrow(Identifier.parse(gen.settings))
			: NoiseGeneratorSettings.fromJson(gen.settings);
	const source = (gen.biome_source ?? {}) as { type?: string; preset?: unknown; biomes?: unknown };
	const sourceType = String(source.type ?? '').replace(/^minecraft:/, '');
	if (sourceType === 'multi_noise' && (source.preset !== undefined || typeof source.biomes === 'string')) {
		throw new Error('Multi-noise presets have no preview: their parameters are built into the game. List the biomes inline to map them.');
	}
	return { settings, biomes: BiomeSource.fromJson(source) };
}

// ---- Terrain ----------------------------------------------------------

/** A palette of block or biome ids shared by every render, so cached columns stay valid. */
const palette: string[] = [];
const paletteIndex = new Map<string, number>();
function indexOf(id: string): number {
	let i = paletteIndex.get(id);
	if (i === undefined) {
		i = palette.length;
		palette.push(id);
		paletteIndex.set(id, i);
	}
	return i;
}

/**
 * Built strips of the slice for the last terrain rendered: 16 blocks along
 * the slice (one chunk's width) by the world height, keyed by stripKey.
 * Indexed (along * height + y), along counting up the world axis.
 */
let terrainKey = '';
const terrainStrips = new Map<string, Uint16Array>();
function stripKey(alongX: boolean, depth: number, chunk: number): string {
	return `${alongX ? 'x' : 'z'}|${depth}|${chunk}`;
}

interface Terrain {
	settings: NoiseGeneratorSettings;
	rule: MaterialRule;
	rs: RandomState;
	biomeAt: (x: number, y: number, z: number) => Identifier;
	/** What the terrain's shape depends on, so filled chunks can be shared across surface rule edits. */
	shape: string;
}

type Generator = { globalFluidPicker: unknown; getWorldgenContext(): unknown };

/**
 * Builds 16 blocks of a vertical slice: along x at z = depth, or along z
 * at x = depth, for the chunk `chunk` along that axis. Only the slice is
 * filled and painted, so it costs about what one row of a chunk does.
 */
function buildStrip(t: Terrain, alongX: boolean, depth: number, chunk: number): Uint16Array {
	const { settings } = t;
	const minY = settings.noise.minY;
	const height = settings.noise.height;
	const at = chunk * 16;
	const pos = alongX ? ChunkPos.create(chunk, Math.floor(depth / 16)) : ChunkPos.create(Math.floor(depth / 16), chunk);
	const blocks = new Chunk(minY, height, pos);
	const gen = new NoiseChunkGenerator(new FixedBiomeSource(Identifier.create('plains')), settings) as unknown as Generator;
	const volume = alongX ? new DensityVolume(16, height, 1, at, minY, depth) : new DensityVolume(1, height, 16, depth, minY, at);
	const noiseChunk = new NoiseChunk(t.rs, settings, gen.globalFluidPicker as never, volume);
	const density = t.rs.router.finalDensity.computeVolume(volume);
	for (let i = 0; i < 16; i++) {
		const [ix, iz] = alongX ? [i, 0] : [0, i];
		const bx = volume.blockX(ix);
		const bz = volume.blockZ(iz);
		for (let y = 0; y < height; y++) {
			const by = minY + y;
			const state = noiseChunk.aquifer.computeSubstance(bx, by, bz, density[volume.indexUnchecked(ix, y, iz)]) ?? settings.defaultBlock;
			blocks.setBlockState(BlockPos.create(bx, by, bz), state);
		}
	}
	const columns: [number, number][] = [];
	for (let i = 0; i < 16; i++) columns.push(alongX ? [at + i, depth] : [depth, at + i]);
	paintSurface(t, blocks, noiseChunk, gen.getWorldgenContext(), columns);
	const out = new Uint16Array(16 * height);
	for (let i = 0; i < 16; i++) {
		const [bx, bz] = columns[i];
		for (let y = 0; y < height; y++) out[i * height + y] = stateIndex(blocks.getBlockState(BlockPos.create(bx, minY + y, bz)));
	}
	return out;
}

/**
 * Runs the surface rule down each of the given block columns of a filled
 * chunk. Deepslate's MaterialSystem.buildSurface only paints the first z
 * row of a chunk (its previews are slices along x), so this is its loop
 * for any columns. Blocks it replaces are pushed to `replaced`.
 */
function paintSurface(t: Terrain, chunk: Chunk, noiseChunk: unknown, worldgen: unknown, columns: Iterable<[number, number]>, replaced?: [BlockPos, BlockState][]) {
	const minY = t.settings.noise.minY;
	const maxY = minY + t.settings.noise.height - 1;
	const context = new MaterialRuleContext(
		t.rs.surfaceSystem,
		t.rs,
		chunk,
		noiseChunk as never,
		worldgen as never,
		(pos: BlockPos) => t.biomeAt(pos[0], pos[1], pos[2])
	);
	const rule = t.rule.compile(context);
	for (const [bx, bz] of columns) {
		context.updateXZ(bx, bz);
		let stoneAbove = 0;
		let waterHeight = Number.MIN_SAFE_INTEGER;
		let nextCeiling = Number.MAX_SAFE_INTEGER;
		for (let y = maxY; y >= minY; y--) {
			const pos = BlockPos.create(bx, y, bz);
			const old = chunk.getBlockState(pos);
			if (old.equals(BlockState.AIR)) {
				stoneAbove = 0;
				waterHeight = Number.MIN_SAFE_INTEGER;
				continue;
			}
			if (old.isFluid()) {
				if (waterHeight === Number.MIN_SAFE_INTEGER) waterHeight = y + 1;
				continue;
			}
			if (nextCeiling >= y) {
				nextCeiling = Number.MIN_SAFE_INTEGER;
				for (let ly = y - 1; ly >= minY; ly--) {
					const next = chunk.getBlockState(BlockPos.create(bx, ly, bz));
					if (next.equals(BlockState.AIR) || next.isFluid()) {
						nextCeiling = ly + 1;
						break;
					}
				}
			}
			stoneAbove += 1;
			context.updateY(stoneAbove, y - nextCeiling + 1, waterHeight, y);
			const state = rule(bx, y, bz);
			if (state) {
				replaced?.push([pos, old]);
				chunk.setBlockState(pos, state);
			}
		}
	}
}

// ---- Isometric terrain -------------------------------------------------

/** Palette indices by block state object, so reading a chunk back doesn't build a name per block. */
const stateIndices = new Map<BlockState, number>();
function stateIndex(state: BlockState): number {
	let i = stateIndices.get(state);
	if (i === undefined) {
		if (stateIndices.size > 4096) stateIndices.clear();
		i = indexOf(state.getName().toString());
		stateIndices.set(state, i);
	}
	return i;
}

/** Filled chunks (terrain shape before surface rules), most recently used last. */
const fills = new Map<string, { chunk: Chunk; noiseChunk: unknown }>();
const MAX_FILLS = 64;

/** A full 16×16 chunk with surface rules, indexed ((x * 16 + z) * height + y). */
function buildChunk3D(t: Terrain, cx: number, cz: number): Uint16Array<ArrayBuffer> {
	const { settings } = t;
	const minY = settings.noise.minY;
	const height = settings.noise.height;
	const gen = new NoiseChunkGenerator(new FixedBiomeSource(Identifier.create('plains')), settings) as unknown as {
		createNoiseChunk(rs: RandomState, chunk: Chunk, onlyFirstZ: boolean): unknown;
		doFill(chunk: Chunk, noiseChunk: unknown): void;
		getWorldgenContext(): unknown;
	};
	// Filling is most of the cost, and it doesn't depend on the surface rule.
	const fillKey = `${t.shape}|${cx},${cz}`;
	let fill = fills.get(fillKey);
	if (fill) {
		fills.delete(fillKey); // re-inserted below as the most recent
	} else {
		const chunk = new Chunk(minY, height, ChunkPos.create(cx, cz));
		const noiseChunk = gen.createNoiseChunk(t.rs, chunk, false);
		gen.doFill(chunk, noiseChunk);
		fill = { chunk, noiseChunk };
	}
	fills.set(fillKey, fill);
	if (fills.size > MAX_FILLS) fills.delete(fills.keys().next().value!);
	const { chunk, noiseChunk } = fill;
	// Blocks the surface rule replaces, put back afterwards so the fill can be reused.
	const replaced: [BlockPos, BlockState][] = [];
	try {
		const x0 = cx * 16;
		const z0 = cz * 16;
		const columns: [number, number][] = [];
		for (let x = 0; x < 16; x++) for (let z = 0; z < 16; z++) columns.push([x0 + x, z0 + z]);
		paintSurface(t, chunk, noiseChunk, gen.getWorldgenContext(), columns, replaced);
		const out = new Uint16Array(16 * 16 * height);
		for (let x = 0; x < 16; x++) {
			for (let z = 0; z < 16; z++) {
				const base = (x * 16 + z) * height;
				for (let y = 0; y < height; y++) {
					out[base + y] = stateIndex(chunk.getBlockState(BlockPos.create(x0 + x, minY + y, z0 + z)));
				}
			}
		}
		return out;
	} finally {
		for (const [pos, old] of replaced) chunk.setBlockState(pos, old);
	}
}

let isoKey = '';

/** A built chunk for the isometric view, with what its faces need. */
interface IsoChunk {
	/** Block palette indices, ((x * 16 + z) * height + y). */
	data: Uint16Array;
	/** Lowest and highest top block of its columns (h and -1 when every column is open to the void). */
	lowest: number;
	highest: number;
	/** Its visible faces, by the area bottom and which neighbours are present. */
	faces: Map<string, ChunkFaces>;
}
interface ChunkFaces {
	/** Per face: chunk-local block x, y, z and side (0 top, 1 +x, 2 +z). */
	faces: Int16Array;
	fills: string[];
}
/** Built chunks for the current terrain, least recently used first. */
const isoChunks = new Map<string, IsoChunk>();
const MAX_ISO_CHUNKS = 400;

/**
 * The visible faces of an area of chunks, in back-to-front order. Hidden
 * blocks are pruned per chunk, so a pan or zoom only projects and fills
 * what can be seen, and a chunk arriving only re-prunes its neighbours.
 */
interface IsoArea {
	key: string;
	n: number;
	/** Lowest and highest drawn y (area-local, 0 = world bottom), or bottom > top when empty. */
	bottom: number;
	top: number;
	/** Per face: block x, y, z and side (0 top, 1 +x, 2 +z). */
	faces: Int16Array;
	/** Per face: its fill colour. */
	fills: string[];
}
let isoArea: IsoArea | null = null;
/** The built chunks drawn at one block size around the canvas, kept so panning only moves it. */
let isoImage: { key: string; s: number; x0: number; y0: number; x1: number; y1: number; image: OffscreenCanvas; empty: boolean } | null = null;

const shades = [1, 0.62, 0.8];
const fillCache = new Map<string, string[]>();
function fillsOf(id: string): string[] {
	let f = fillCache.get(id);
	if (!f) {
		const c = blockColor(id);
		f = shades.map((k) => `rgb(${Math.round(c[0] * k)} ${Math.round(c[1] * k)} ${Math.round(c[2] * k)})`);
		fillCache.set(id, f);
	}
	return f;
}

let airIndex = -1;
let caveAirIndex = -1;
function isEmpty(b: number): boolean {
	if (airIndex < 0) {
		airIndex = indexOf('minecraft:air');
		caveAirIndex = indexOf('minecraft:cave_air');
	}
	return b === airIndex || b === caveAirIndex;
}

function isoChunkOf(data: Uint16Array, h: number): IsoChunk {
	let lowest = h;
	let highest = -1;
	for (let col = 0; col < 256; col++) {
		const base = col * h;
		let y = h - 1;
		while (y >= 0 && isEmpty(data[base + y])) y--;
		if (y < 0) continue; // open to the void, as under sky islands
		if (y < lowest) lowest = y;
		if (y > highest) highest = y;
	}
	return { data, lowest, highest, faces: new Map() };
}

/**
 * A chunk's exposed top, +x and +z faces from bottom up, in back-to-front
 * order (x, then z, then y). Faces against a missing neighbour or the area's
 * edge are drawn, so the area's sides show its layers.
 */
function chunkFaces(c: IsoChunk, east: IsoChunk | undefined, south: IsoChunk | undefined, bottom: number, h: number): ChunkFaces {
	const key = `${bottom}|${east ? 1 : 0}|${south ? 1 : 0}`;
	const hit = c.faces.get(key);
	if (hit) return hit;
	const { data } = c;
	const faces: number[] = [];
	const fills: string[] = [];
	for (let x = 0; x < 16; x++) {
		for (let z = 0; z < 16; z++) {
			const base = (x * 16 + z) * h;
			const eastBase = x < 15 ? base + 16 * h : z * h;
			const eastData = x < 15 ? data : east?.data;
			const southBase = z < 15 ? base + h : x * 16 * h;
			const southData = z < 15 ? data : south?.data;
			for (let y = bottom; y <= c.highest; y++) {
				const b = data[base + y];
				if (isEmpty(b)) continue;
				const f = fillsOf(palette[b]);
				if (y === h - 1 || isEmpty(data[base + y + 1])) faces.push(x, y, z, 0), fills.push(f[0]);
				if (!eastData || isEmpty(eastData[eastBase + y])) faces.push(x, y, z, 1), fills.push(f[1]);
				if (!southData || isEmpty(southData[southBase + y])) faces.push(x, y, z, 2), fills.push(f[2]);
			}
		}
	}
	const out = { faces: Int16Array.from(faces), fills };
	// Keep a few variants (bottoms differ between areas); drop the oldest.
	if (c.faces.size >= 4) c.faces.delete(c.faces.keys().next().value!);
	c.faces.set(key, out);
	return out;
}

/** The area's faces, chunk by chunk. Chunk order (x-major, then z) is a valid back-to-front order too. */
function isoAreaOf(key: string, chunks: (IsoChunk | undefined)[][], n: number, h: number): IsoArea {
	let lowest = h;
	let highest = -1;
	for (const row of chunks) {
		for (const c of row) {
			if (!c || c.highest < 0) continue;
			lowest = Math.min(lowest, c.lowest);
			highest = Math.max(highest, c.highest);
		}
	}
	// From a little below the lowest surface, so the cut sides show the layers the surface rules paint.
	const bottom = Math.max(0, lowest - 16);
	const parts: ChunkFaces[] = [];
	const offsets: [number, number][] = [];
	let total = 0;
	for (let i = 0; i < chunks.length; i++) {
		for (let j = 0; j < chunks[i].length; j++) {
			const c = chunks[i][j];
			if (!c || c.highest < 0) continue;
			const part = chunkFaces(c, chunks[i + 1]?.[j], chunks[i][j + 1], bottom, h);
			parts.push(part);
			offsets.push([i * 16, j * 16]);
			total += part.faces.length;
		}
	}
	const faces = new Int16Array(total);
	const fills: string[] = [];
	let o = 0;
	parts.forEach((part, k) => {
		const [dx, dz] = offsets[k];
		const f = part.faces;
		for (let i = 0; i < f.length; i += 4) {
			faces[o++] = f[i] + dx;
			faces[o++] = f[i + 1];
			faces[o++] = f[i + 2] + dz;
			faces[o++] = f[i + 3];
		}
		for (const s of part.fills) fills.push(s);
	});
	return { key, n, bottom, top: highest, faces, fills };
}

/** Fills an area's faces, with block (x, y, z) at screen (ox + (x - z)·s, oy + (x + z)·s/2 - y·s), skipping those outside w × h. */
function drawFaces(ctx: OffscreenCanvasRenderingContext2D, area: IsoArea, s: number, ox: number, oy: number, w: number, h: number) {
	const { faces, fills } = area;
	ctx.lineWidth = 0.5;
	// Runs of faces with the same colour are filled as one path; the order
	// between runs is kept, so the back-to-front drawing still holds.
	let current = '';
	const flush = () => {
		if (!current) return;
		ctx.fill();
		ctx.stroke(); // hides hairline seams between faces
	};
	for (let i = 0, f = 0; i < faces.length; i += 4, f++) {
		const x = faces[i], y = faces[i + 1], z = faces[i + 2], side = faces[i + 3];
		// The block's top-back corner; a block spans 2s across and 2s down from it.
		const bx = ox + (x - z) * s;
		const by = oy + ((x + z) * s) / 2 - (y + 1) * s;
		if (bx + s < 0 || bx - s > w || by + 2 * s < 0 || by > h) continue;
		if (fills[f] !== current) {
			flush();
			current = fills[f];
			ctx.fillStyle = ctx.strokeStyle = current;
			ctx.beginPath();
		}
		if (side === 0) {
			ctx.moveTo(bx, by);
			ctx.lineTo(bx + s, by + s / 2);
			ctx.lineTo(bx, by + s);
			ctx.lineTo(bx - s, by + s / 2);
		} else if (side === 1) {
			ctx.moveTo(bx + s, by + s / 2);
			ctx.lineTo(bx, by + s);
			ctx.lineTo(bx, by + 2 * s);
			ctx.lineTo(bx + s, by + 1.5 * s);
		} else {
			ctx.moveTo(bx - s, by + s / 2);
			ctx.lineTo(bx, by + s);
			ctx.lineTo(bx, by + 2 * s);
			ctx.lineTo(bx - s, by + 1.5 * s);
		}
		ctx.closePath();
	}
	flush();
}

/** The largest image kept for panning, in pixels. */
const MAX_IMAGE = 4096 * 4096;

// ---- Chunk building helpers -------------------------------------------
//
// Building a chunk is most of the isometric view's cost (Deepslate's noise
// fill, about 0.4 s each), so the page starts a few helper copies of this
// worker and hands their ports over. Visible chunks are built first, in
// parallel; the ring just outside the view is built next, so a pan rarely
// waits. Without helpers, chunks are built here one at a time.

type HelperMessage =
	| { type: 'setup'; files: PreviewFile[] }
	| { type: 'chunk'; job: number; req: RenderRequest; cx: number; cz: number };
type HelperReply =
	| { type: 'chunk'; job: number; data: Uint16Array<ArrayBuffer>; ids: string[] }
	| { type: 'chunk'; job: number; error: string };

interface Job {
	id: number;
	key: string;
	cx: number;
	cz: number;
	req: RenderRequest;
	prefetch: boolean;
	promise: Promise<IsoChunk>;
	resolve: (c: IsoChunk) => void;
	reject: (e: unknown) => void;
}
const helpers: { port: MessagePort; job: Job | null }[] = [];
let lastSetup: HelperMessage | null = null;
let queue: Job[] = [];
/** Queued and running jobs, by chunk. */
const jobs = new Map<string, Job>();
let nextJob = 0;

function addHelpers(ports: MessagePort[]) {
	for (const port of ports) {
		const helper: (typeof helpers)[number] = { port, job: null };
		port.onmessage = (e: MessageEvent<HelperReply>) => finish(helper, e.data);
		if (lastSetup) port.postMessage(lastSetup);
		helpers.push(helper);
	}
	dispatch();
}

function finish(helper: (typeof helpers)[number], reply: HelperReply) {
	const job = helper.job;
	helper.job = null;
	if (job && job.id === reply.job) {
		if (jobs.get(`${job.cx},${job.cz}`) === job) jobs.delete(`${job.cx},${job.cz}`);
		if ('error' in reply) {
			job.reject(new Error(reply.error));
		} else if (job.key !== isoKey) {
			job.reject(new Superseded());
		} else {
			// The helper's palette differs from this one's.
			const map = Uint16Array.from(reply.ids, indexOf);
			const data = reply.data;
			for (let i = 0; i < data.length; i++) data[i] = map[data[i]];
			const c = keep(`${job.cx},${job.cz}`, isoChunkOf(data, data.length / 256));
			job.resolve(c);
			if (job.prefetch) refreshSoon();
		}
	}
	dispatch();
}

function keep(at: string, c: IsoChunk): IsoChunk {
	isoChunks.set(at, c);
	noteHeights(c);
	if (isoChunks.size > MAX_ISO_CHUNKS) isoChunks.delete(isoChunks.keys().next().value!);
	return c;
}

/**
 * Each chunk goes to the same helper every time, so that helper's filled
 * chunks are reused when a surface rule changes.
 */
function helperOf(cx: number, cz: number): number {
	const n = helpers.length;
	return (((cx * 31 + cz) % n) + n) % n;
}

function dispatch() {
	helpers.forEach((helper, h) => {
		if (helper.job) return;
		// The queue holds visible chunks first, in the order asked.
		const k = queue.findIndex((j) => helperOf(j.cx, j.cz) === h);
		if (k < 0) return;
		const job = queue.splice(k, 1)[0];
		helper.job = job;
		helper.port.postMessage({ type: 'chunk', job: job.id, req: job.req, cx: job.cx, cz: job.cz } satisfies HelperMessage);
	});
}

/**
 * Asks for the chunks an area needs, visible ones first, then the ring
 * around them. Queued chunks no longer wanted are dropped.
 */
function plan(req: RenderRequest, visible: [number, number][], ring: [number, number][]): Map<string, Promise<IsoChunk>> {
	const wanted = new Map<string, boolean>();
	for (const [cx, cz] of visible) wanted.set(`${cx},${cz}`, false);
	for (const [cx, cz] of ring) if (!wanted.has(`${cx},${cz}`)) wanted.set(`${cx},${cz}`, true);
	queue = queue.filter((j) => {
		const at = `${j.cx},${j.cz}`;
		if (j.key === isoKey && wanted.has(at)) return true;
		if (jobs.get(at) === j) jobs.delete(at);
		j.reject(new Superseded());
		return false;
	});
	const out = new Map<string, Promise<IsoChunk>>();
	for (const [at, prefetch] of wanted) {
		if (isoChunks.has(at)) continue;
		let job = jobs.get(at);
		if (job && job.key !== isoKey) job = undefined;
		if (!job) {
			const [cx, cz] = at.split(',').map(Number);
			let resolve!: (c: IsoChunk) => void;
			let reject!: (e: unknown) => void;
			const promise = new Promise<IsoChunk>((res, rej) => ((resolve = res), (reject = rej)));
			promise.catch(() => {}); // a dropped prefetch is nobody's error
			job = { id: nextJob++, key: isoKey, cx, cz, req, prefetch, promise, resolve, reject };
			jobs.set(at, job);
			queue.push(job);
		}
		job.prefetch = prefetch;
		if (!prefetch) out.set(at, job.promise);
	}
	// Keep visible jobs in the order asked, ahead of the ring.
	queue.sort((a, b) => Number(a.prefetch) - Number(b.prefetch));
	dispatch();
	return out;
}

/**
 * Builds a chunk, fetching vanilla entries its surface rule references the
 * first time they are missing (re-registering everything only then).
 */
async function buildIsoChunk(t: Terrain, cx: number, cz: number): Promise<Uint16Array<ArrayBuffer>> {
	try {
		return buildChunk3D(t, cx, cz);
	} catch (e) {
		if (!String(e).includes('Missing key')) throw e;
		return withVanilla(() => buildChunk3D(t, cx, cz));
	}
}

/** In a helper: builds one chunk for the main worker. */
async function serve(port: MessagePort, msg: HelperMessage) {
	if (msg.type === 'setup') {
		setup(msg.files);
		return;
	}
	try {
		const { terrain } = await terrainOf(msg.req);
		const data = await buildIsoChunk(terrain, msg.cx, msg.cz);
		port.postMessage({ type: 'chunk', job: msg.job, data, ids: palette.slice() } satisfies HelperReply, [data.buffer]);
	} catch (err) {
		port.postMessage({ type: 'chunk', job: msg.job, error: String(err instanceof Error ? err.message : err) } satisfies HelperReply);
	}
}

const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

/**
 * Draws the terrain as an isometric block diagram filling the canvas. In
 * this mode the view's x and y are the world x and z at the centre, and its
 * scale is the size of a block in pixels (0 for the default).
 *
 * Screen positions are "global": block (x, y, z) is at
 * ((x - z)·s, (x + z)·s/2 - (y - yc)·s), where yc is a height fixed per
 * terrain, so the picture doesn't jump as chunks arrive. The canvas shows
 * the global rectangle around the view's centre.
 *
 * Every chunk the canvas shows is built, nearest the centre first, then the
 * ones in a margin around it. The picture is drawn once into an image a
 * margin larger than the canvas, so a pan only moves it until it leaves the
 * margin or more chunks arrive.
 */
async function renderIso(req: RenderRequest, t: Terrain, key: string): Promise<Rendered> {
	const { width, height } = req;
	const h = t.settings.noise.height;
	if (key !== isoKey) {
		isoKey = key;
		isoChunks.clear();
		isoArea = null;
		isoImage = null;
	}
	if (isoCentre?.shape !== t.shape) {
		// Until a chunk says otherwise, centre on sea level.
		isoCentre = { shape: t.shape, y: t.settings.seaLevel - t.settings.noise.minY, fixed: false };
	}
	const sMin = isoMinScale(width, height);
	const s = Math.max(sMin, req.view.scale > 0 ? req.view.scale : width / 72);
	const vx = req.view.x;
	const vz = req.view.y;
	// The canvas, in global screen coordinates.
	const view = { x0: (vx - vz) * s - width / 2, y0: ((vx + vz) * s) / 2 - height / 2, x1: 0, y1: 0 };
	view.x1 = view.x0 + width;
	view.y1 = view.y0 + height;
	const mx = Math.min(width / 4, 32 * s);
	const my = Math.min(height / 4, 32 * s);
	const around = { x0: view.x0 - mx, y0: view.y0 - my, x1: view.x1 + mx, y1: view.y1 + my };

	const byDistance = (a: [number, number], b: [number, number]) =>
		Math.hypot(a[0] * 16 + 8 - vx, a[1] * 16 + 8 - vz) - Math.hypot(b[0] * 16 + 8 - vx, b[1] * 16 + 8 - vz);
	/** The chunks the canvas shows, and those not built yet; estimates improve as chunks arrive. */
	const needed = () => {
		const shown = chunksIn(view, s, h).sort(byDistance);
		const missing = shown.filter(([cx, cz]) => !touch(`${cx},${cz}`));
		return { shown, missing };
	};
	let { shown, missing } = needed();

	/** Draws (if needed) the image covering the canvas and returns the frame. */
	const frame = (): Rendered => {
		const canvas = new OffscreenCanvas(width, height);
		const ctx = canvas.getContext('2d')!;
		ctx.fillStyle = `rgb(${SKY.join(' ')})`;
		ctx.fillRect(0, 0, width, height);
		const image = isoImageFor(view, around, s, h);
		if (image) ctx.drawImage(image.image, Math.round(image.x0 - view.x0), Math.round(image.y0 - view.y0));
		const empty = !image || image.empty;
		const notice = empty && missing.length === 0 ? 'No blocks here. Drag to pan to some terrain.' : undefined;
		const bitmap = canvas.transferToImageBitmap();
		return {
			type: 'rendered', seq: req.seq, width, height,
			pixels: new Uint8ClampedArray(0), values: new Float32Array(0),
			min: Infinity, max: -Infinity, scale: s, minScale: sMin, bitmap, notice
		};
	};
	// Once the ring is built, redraw the image in the background so the next pan doesn't wait.
	refreshIso = () => isoImageFor(view, around, s, h);

	let posted = -1;
	let built = shown.length - missing.length;
	// Built chunks may show that more are needed (a tall one just off the
	// bottom edge, say), so this repeats until nothing shown is missing.
	while (missing.length > 0) {
		if (helpers.length === 0) {
			// No helpers: build here, showing chunks as they finish.
			for (const [cx, cz] of missing) {
				if (built > 0) post({ ...frame(), partial: true });
				await breathe();
				if (req.seq !== latest) throw new Superseded();
				keep(`${cx},${cz}`, isoChunkOf(await buildIsoChunk(t, cx, cz), h));
				built++;
			}
		} else {
			const ring = chunksIn(around, s, h).filter(([cx, cz]) => !shown.some((v) => v[0] === cx && v[1] === cz)).sort(byDistance);
			const wanted = plan(req, missing, ring);
			let waiting = missing.length;
			const all = Promise.all(missing.map(([cx, cz]) => wanted.get(`${cx},${cz}`)!.then(() => void (waiting--, built++))));
			all.catch(() => {});
			while (waiting > 0) {
				// Show what is built so far, a few times a second at most.
				if (built !== posted && built > 0) post({ ...frame(), partial: true });
				posted = built;
				await Promise.race([all, sleep(150)]);
				if (req.seq !== latest) throw new Superseded();
			}
			await all; // rethrows a helper's error
		}
		({ shown, missing } = needed());
	}
	return frame();
}

/** Marks a built chunk as recently used; whether it is built. */
function touch(at: string): boolean {
	const c = isoChunks.get(at);
	if (!c) return false;
	isoChunks.delete(at);
	isoChunks.set(at, c);
	return true;
}

/** The most chunks the canvas may show; zooming out stops there. */
const MAX_SHOWN = 100;
/**
 * The smallest block size at which a canvas shows about MAX_SHOWN chunks.
 * It spans w/s blocks of x - z and 2h/s of x + z, and the chunks along
 * its edges count too.
 */
function isoMinScale(width: number, height: number): number {
	const chunks = (s: number) => ((width / s + 32) * ((2 * height) / s + 64)) / 2 / 256;
	let s = 1;
	while (chunks(s) > MAX_SHOWN && s < 64) s *= 1.05;
	return s;
}

/** The height the view is centred on, fixed per terrain shape. */
let isoCentre: { shape: string; y: number; fixed: boolean } | null = null;

interface Rect { x0: number; y0: number; x1: number; y1: number }

/** The global screen rectangle of a chunk's blocks between heights lo and hi. */
function chunkRect(cx: number, cz: number, lo: number, hi: number, s: number): Rect {
	const yc = isoCentre!.y;
	const x = cx * 16;
	const z = cz * 16;
	return {
		x0: (x - z - 16) * s - s,
		x1: (x + 16 - z) * s + s,
		y0: ((x + z) * s) / 2 - (hi + 1 - yc) * s,
		y1: ((x + z + 32) * s) / 2 - (lo - yc) * s + s
	};
}
const overlaps = (a: Rect, b: Rect) => a.x0 < b.x1 && b.x0 < a.x1 && a.y0 < b.y1 && b.y0 < a.y1;

/** The heights a chunk's faces span: its own once built, else what built chunks suggest. */
function heightsOf(cx: number, cz: number, h: number): [number, number] {
	const c = isoChunks.get(`${cx},${cz}`);
	// Only the surface counts: a chunk's cut sides only show at the edge of the drawn area.
	if (c) return c.highest < 0 ? [0, -1] : [c.lowest, c.highest];
	// Not built: guess from built neighbours, else from the centre height.
	let lo = h;
	let hi = -1;
	for (let i = -1; i <= 1; i++) {
		for (let j = -1; j <= 1; j++) {
			const n = isoChunks.get(`${cx + i},${cz + j}`);
			if (!n || n.highest < 0) continue;
			lo = Math.min(lo, n.lowest);
			hi = Math.max(hi, n.highest);
		}
	}
	if (hi >= 0) return [lo, hi];
	const yc = isoCentre!.y;
	return [Math.max(0, yc - 16), Math.min(h - 1, yc + 16)];
}

/** The chunks whose blocks may show in a global screen rectangle. */
function chunksIn(r: Rect, s: number, h: number): [number, number][] {
	const yc = isoCentre!.y;
	// u = x - z and v = x + z bound the rectangle, allowing for every height.
	const u0 = r.x0 / s - 2, u1 = r.x1 / s + 2;
	const v0 = 2 * (r.y0 / s - yc) - 2, v1 = 2 * (r.y1 / s + h - yc) + 2;
	const cx0 = Math.floor((u0 + v0) / 32) - 1, cx1 = Math.floor((u1 + v1) / 32) + 1;
	const cz0 = Math.floor((v0 - u1) / 32) - 1, cz1 = Math.floor((v1 - u0) / 32) + 1;
	const out: [number, number][] = [];
	for (let cx = cx0; cx <= cx1; cx++) {
		// Only chunks whose u range meets the rectangle's.
		for (let cz = Math.max(cz0, Math.floor((cx * 16 - u1) / 16) - 1); cz <= Math.min(cz1, Math.floor((cx * 16 + 16 - u0) / 16) + 1); cz++) {
			const [lo, hi] = heightsOf(cx, cz, h);
			if (hi >= lo && overlaps(chunkRect(cx, cz, lo, hi, s), r)) out.push([cx, cz]);
		}
	}
	return out;
}

/** Called when a built chunk arrives: the first one sets the height the view is centred on. */
function noteHeights(c: IsoChunk) {
	if (c.highest < 0) return;
	if (isoCentre && !isoCentre.fixed) {
		// The first chunk (nearest the centre) sets the height the view is centred on.
		isoCentre.y = Math.round((c.lowest + c.highest) / 2);
		isoCentre.fixed = true;
	}
}

/** Redraws the image after background chunks arrive; set by the last render. */
let refreshIso: (() => void) | null = null;
let refreshTimer: ReturnType<typeof setTimeout> | undefined;
function refreshSoon() {
	clearTimeout(refreshTimer);
	refreshTimer = setTimeout(() => refreshIso?.(), 200);
}

/**
 * The image of the built chunks in `around`, reused while the canvas stays
 * inside it at the same scale and no chunk inside it has changed.
 */
function isoImageFor(view: Rect, around: Rect, s: number, h: number): (Rect & { image: OffscreenCanvas; empty: boolean }) | null {
	const reuse = isoImage && isoImage.s === s && isoImage.key.startsWith(isoKey + '#') &&
		view.x0 >= isoImage.x0 && view.y0 >= isoImage.y0 && view.x1 <= isoImage.x1 && view.y1 <= isoImage.y1;
	const r: Rect = reuse ? isoImage! : around;
	const built = chunksIn(r, s, h).filter(([cx, cz]) => isoChunks.has(`${cx},${cz}`));
	const key = `${isoKey}#${s}#${built.map((c) => c.join(',')).sort().join(';')}`;
	if (reuse && isoImage!.key === key) return isoImage;
	const iw = Math.ceil(r.x1 - r.x0);
	const ih = Math.ceil(r.y1 - r.y0);
	if (iw * ih > MAX_IMAGE) return r === view ? null : isoImageFor(view, view, s, h);
	// The built chunks as a grid, in back-to-front order.
	let gx0 = Infinity, gz0 = Infinity, gx1 = -Infinity, gz1 = -Infinity;
	for (const [cx, cz] of built) {
		gx0 = Math.min(gx0, cx); gz0 = Math.min(gz0, cz);
		gx1 = Math.max(gx1, cx); gz1 = Math.max(gz1, cz);
	}
	const grid: (IsoChunk | undefined)[][] = [];
	for (let cx = gx0; cx <= gx1; cx++) {
		const row: (IsoChunk | undefined)[] = [];
		for (let cz = gz0; cz <= gz1; cz++) row.push(isoChunks.get(`${cx},${cz}`));
		grid.push(row);
	}
	if (isoArea?.key !== key) isoArea = isoAreaOf(key, grid, 0, h);
	const area = isoArea;
	const image = new OffscreenCanvas(iw, ih);
	const empty = built.length === 0 || area.bottom > area.top;
	if (!empty) {
		const yc = isoCentre!.y;
		const ox = (gx0 * 16 - gz0 * 16) * s - r.x0;
		const oy = ((gx0 * 16 + gz0 * 16) * s) / 2 + yc * s - r.y0;
		drawFaces(image.getContext('2d')!, area, s, ox, oy, iw, ih);
	}
	isoImage = { key, s, x0: r.x0, y0: r.y0, x1: r.x1, y1: r.y1, image, empty };
	return isoImage;
}

/**
 * A 2D view's pixel geometry, from the camera at the image centre and the
 * screen axes of its facing. Pixels are sampled at their centres, so a
 * view facing south (x decreasing to the right) lines up like north.
 */
function planeOf(width: number, height: number, view: PreviewView) {
	const axes = screenAxes(view.facing);
	const alongX = axes.rx !== 0;
	const sign = alongX ? axes.rx : axes.rz;
	const across = (px: number) => (px + 0.5 - width / 2) * view.scale;
	const down = (py: number) => (py + 0.5 - height / 2) * view.scale;
	return {
		...axes,
		/** For side views: whether the slice runs along x, and the other coordinate it is at. */
		alongX,
		depth: Math.floor(alongX ? view.z : view.x),
		/** For side views: the world coordinate along the slice at a pixel column. */
		along: (px: number) => (alongX ? view.x : view.z) + sign * across(px),
		/** For side views: the world y at a pixel row. */
		y: (py: number) => view.y - down(py),
		/** For top-down views: the world x and z at a pixel. */
		x: (px: number, py: number) => view.x + axes.rx * across(px) + axes.dx * down(py),
		z: (px: number, py: number) => view.z + axes.rz * across(px) + axes.dz * down(py),
		/** For side views: the world x and z of a point along the slice. */
		at: (u: number): [number, number] => (alongX ? [u, view.z] : [view.x, u])
	};
}

/**
 * Builds the strip at the view's centre inside withVanilla, so entries the
 * surface rules reference are fetched before the full pass.
 */
async function warm(req: RenderRequest, t: Terrain, key: string): Promise<Terrain> {
	if (key !== terrainKey) {
		terrainKey = key;
		terrainStrips.clear();
	}
	const plane = planeOf(req.width, req.height, req.view);
	const chunk = Math.floor(Math.floor(plane.along(req.width / 2)) / 16);
	const k = stripKey(plane.alongX, plane.depth, chunk);
	if (!terrainStrips.has(k)) {
		await withVanilla(() => terrainStrips.set(k, buildStrip(t, plane.alongX, plane.depth, chunk)));
	}
	return t;
}

async function renderTerrain(req: RenderRequest, t: Terrain, key: string): Promise<Rendered> {
	const { width, height, view } = req;
	const minY = t.settings.noise.minY;
	const maxY = minY + t.settings.noise.height;
	const sea = t.settings.seaLevel;
	if (key !== terrainKey) {
		terrainKey = key;
		terrainStrips.clear();
	}
	const plane = planeOf(width, height, view);
	const values = new Float32Array(width * height).fill(NaN);
	const pixels = new Uint8ClampedArray(width * height * 4);
	const h = t.settings.noise.height;
	const inView = new Set<string>();
	let started = performance.now();
	for (let px = 0; px < width; px++) {
		const u = Math.floor(plane.along(px));
		const chunk = Math.floor(u / 16);
		const k = stripKey(plane.alongX, plane.depth, chunk);
		inView.add(k);
		let strip = terrainStrips.get(k);
		if (!strip) {
			if (performance.now() - started > 40) {
				await breathe();
				if (req.seq !== latest) throw new Superseded();
				started = performance.now();
			}
			strip = buildStrip(t, plane.alongX, plane.depth, chunk);
			terrainStrips.set(k, strip);
		}
		const local = u - chunk * 16;
		const [x, z] = plane.at(u);
		// Air below the router's estimate of the surface is cave, drawn darker than sky.
		const surface = t.rs.router.chunkSurfaceLevel.compute(DensityFunction.context(x, 0, z));
		for (let py = 0; py < height; py++) {
			const y = Math.floor(plane.y(py));
			const i = py * width + px;
			if (y < minY || y >= maxY) {
				put(pixels, i * 4, VOID);
				continue;
			}
			const b = strip[local * h + (y - minY)];
			const id = palette[b];
			values[i] = b;
			if (id === 'minecraft:air') {
				put(pixels, i * 4, y < surface ? CAVE : SKY);
			} else {
				put(pixels, i * 4, blockColor(id));
			}
		}
	}
	// Keep the cache from growing without bound while panning and turning.
	if (terrainStrips.size > 512) {
		for (const k of terrainStrips.keys()) if (!inView.has(k)) terrainStrips.delete(k);
	}
	return { type: 'rendered', seq: req.seq, width, height, pixels, values, min: Infinity, max: -Infinity, labels: palette.slice(), seaLevel: sea, minY, maxY };
}

// ---- Biome maps -------------------------------------------------------

async function renderBiomes(req: RenderRequest, dim: Dimension, rs: RandomState): Promise<Rendered> {
	const { width, height, view } = req;
	const values = new Float32Array(width * height);
	const pixels = new Uint8ClampedArray(width * height * 4);
	const plane = planeOf(width, height, view);
	const qy = Math.floor(view.y) >> 2;
	const counts = new Map<number, number>();
	// Biomes are stored per 4×4×4 cell, so neighbouring pixels in one cell share a lookup.
	const cache = new Map<string, number>();
	let started = performance.now();
	for (let py = 0; py < height; py++) {
		if (performance.now() - started > 40) {
			await breathe();
			if (req.seq !== latest) throw new Superseded();
			started = performance.now();
		}
		for (let px = 0; px < width; px++) {
			const qx = Math.floor(plane.x(px, py)) >> 2;
			const qz = Math.floor(plane.z(px, py)) >> 2;
			const k = `${qx},${qz}`;
			let b = cache.get(k);
			if (b === undefined) {
				b = indexOf(dim.biomes.getBiome(qx, qy, qz, rs.sampler).toString());
				cache.set(k, b);
			}
			const i = py * width + px;
			values[i] = b;
			counts.set(b, (counts.get(b) ?? 0) + 1);
			put(pixels, i * 4, biomeColor(palette[b]));
		}
	}
	// Darken the edges between biomes so neighbours with close colours stay apart.
	for (let py = 0; py < height; py++) {
		for (let px = 0; px < width; px++) {
			const i = py * width + px;
			if ((px + 1 < width && values[i + 1] !== values[i]) || (py + 1 < height && values[i + width] !== values[i])) {
				pixels[i * 4] *= 0.75;
				pixels[i * 4 + 1] *= 0.75;
				pixels[i * 4 + 2] *= 0.75;
			}
		}
	}
	const legend = [...counts.entries()]
		.sort((a, b) => b[1] - a[1])
		.map(([b]) => ({ id: palette[b], color: `rgb(${biomeColor(palette[b]).join(' ')})` }));
	return { type: 'rendered', seq: req.seq, width, height, pixels, values, min: Infinity, max: -Infinity, labels: palette.slice(), legend };
}

// ---- Render -----------------------------------------------------------

/** The last terrain set up, so panning doesn't rebuild its random state each frame. */
let cached: { key: string; t: Terrain } | null = null;

/** The terrain a block view draws: a dimension's, or noise settings with a surface rule. */
async function terrainOf(req: RenderRequest): Promise<{ terrain: Terrain; key: string }> {
	const seed = seedOf(req.seed);
	const id = Identifier.parse(req.id);
	const fixedBiome = Identifier.parse(req.options.biome || 'minecraft:plains');
	const what = req.kind === 'dimension' ? '' : `${req.options.terrain}|${fixedBiome}`;
	const key = [version, req.seed, req.kind, req.id, req.mode, what].join('|');
	if (cached?.key === key) return { terrain: cached.t, key };
	const terrain =
		req.kind === 'dimension'
			? await withVanilla((): Terrain => {
					const dim = readDimension(req.id);
					const rs = speedUp(new RandomState(dim.settings, seed));
					return {
						settings: dim.settings,
						rule: dim.settings.materialRule,
						rs,
						biomeAt: (x, y, z) => dim.biomes.getBiome(x >> 2, y >> 2, z >> 2, rs.sampler),
						shape: [shapeVersion, req.seed, 'dimension', req.id].join('|')
					};
				})
			: await withVanilla((): Terrain => {
					const settingsId = req.kind === 'material_rule' ? Identifier.parse(req.options.terrain || 'minecraft:overworld') : id;
					const settings = WorldgenRegistries.NOISE_SETTINGS.getOrThrow(settingsId);
					const rule = req.kind === 'material_rule' ? WorldgenRegistries.MATERIAL_RULE.getOrThrow(id) : settings.materialRule;
					const rs = speedUp(new RandomState(settings, seed));
					const shape = [shapeVersion, req.seed, 'noise_settings', settingsId].join('|');
					return { settings, rule, rs, biomeAt: () => fixedBiome, shape };
				});
	cached = { key, t: terrain };
	return { terrain, key };
}

async function render(req: RenderRequest): Promise<PreviewResponse> {
	const { width, height, view } = req;
	const seed = seedOf(req.seed);
	const id = Identifier.parse(req.id);
	const values = new Float32Array(width * height);
	const pixels = new Uint8ClampedArray(width * height * 4);
	const plane = planeOf(width, height, view);
	let min = Infinity;
	let max = -Infinity;
	const track = (v: number) => {
		if (v < min) min = v;
		if (v > max) max = v;
	};
	if (req.kind === 'dimension' && req.mode === 'biomes') {
		const { dim, rs } = await withVanilla(() => {
			const dim = readDimension(req.id);
			const rs = new RandomState(dim.settings, seed);
			// Touch the climate router, so missing entries are fetched now.
			rs.sampler.sample(0, 0, 0);
			return { dim, rs };
		});
		return renderBiomes(req, dim, rs);
	}

	if (req.kind === 'dimension' || req.kind === 'material_rule' || (req.kind === 'noise_settings' && req.mode !== 'density')) {
		const { terrain, key } = await terrainOf(req);
		if (req.mode === 'iso') return renderIso(req, terrain, key);
		await warm(req, terrain, key);
		return renderTerrain(req, terrain, key);
	}

	if (req.kind === 'noise') {
		// Top-down, at the camera's y.
		const noise = await withVanilla(() => {
			const rs = new RandomState(NoiseGeneratorSettings.create({}), seed);
			return rs.getOrCreateNoise(id);
		});
		for (let py = 0; py < height; py++) {
			for (let px = 0; px < width; px++) {
				const v = noise.get3D(plane.x(px, py), view.y, plane.z(px, py));
				values[py * width + px] = v;
				track(v);
			}
		}
		for (let i = 0; i < values.length; i++) {
			const g = Math.max(0, Math.min(255, (values[i] + 1) * 127.5));
			pixels[i * 4] = pixels[i * 4 + 1] = pixels[i * 4 + 2] = g;
			pixels[i * 4 + 3] = 255;
		}
		return { type: 'rendered', seq: req.seq, width, height, pixels, values, min, max };
	}

	if (req.kind === 'density_function') {
		// Side view: the slice through the camera, y up.
		const fn = await withVanilla(() => {
			const rs = new RandomState(NoiseGeneratorSettings.create({}), seed);
			// A reference holder, so the visitor binds noises and caches by id.
			const df = new DensityFunction.HolderHolder(Holder.reference(WorldgenRegistries.DENSITY_FUNCTION, id));
			return df.mapAll(rs.createVisitor(false));
		});
		for (let py = 0; py < height; py++) {
			const y = plane.y(py);
			for (let px = 0; px < width; px++) {
				const [x, z] = plane.at(plane.along(px));
				const v = fn.compute(DensityFunction.context(x, y, z));
				values[py * width + px] = v;
				track(v);
			}
		}
		const range = Math.max(Math.abs(min), Math.abs(max), 1e-9);
		for (let i = 0; i < values.length; i++) diverging(values[i] / range, pixels, i * 4);
		return { type: 'rendered', seq: req.seq, width, height, pixels, values, min, max };
	}

	// Noise settings, density mode: the terrain's final density, a side view.
	const { settings, router } = await withVanilla(() => {
		const settings = WorldgenRegistries.NOISE_SETTINGS.getOrThrow(id);
		const rs = speedUp(new RandomState(settings, seed));
		return { settings, router: rs.router as NoiseRouter };
	});
	const minY = settings.noise.minY;
	const maxY = minY + settings.noise.height;
	const sea = settings.seaLevel;
	for (let py = 0; py < height; py++) {
		const y = plane.y(py);
		for (let px = 0; px < width; px++) {
			const i = py * width + px;
			const o = i * 4;
			let v = NaN;
			if (y >= minY && y < maxY) {
				const [x, z] = plane.at(plane.along(px));
				v = router.finalDensity.compute(DensityFunction.context(x, y, z));
				track(v);
			}
			values[i] = v;
			if (Number.isNaN(v)) {
				put(pixels, o, VOID);
			} else if (v > 0) {
				const shade = 120 + Math.min(1, v) * 40;
				pixels[o] = shade; pixels[o + 1] = shade; pixels[o + 2] = shade + 6;
			} else if (y < sea) {
				pixels[o] = 52; pixels[o + 1] = 104; pixels[o + 2] = 196;
			} else {
				put(pixels, o, SKY);
			}
			pixels[o + 3] = 255;
		}
	}
	return { type: 'rendered', seq: req.seq, width, height, pixels, values, min, max, seaLevel: sea, minY, maxY };
}

function post(res: Rendered) {
	const transfer: Transferable[] = [res.pixels.buffer, res.values.buffer];
	if (res.bitmap) transfer.push(res.bitmap);
	(self as unknown as Worker).postMessage(res, transfer);
}

function setup(files: PreviewFile[]) {
	userFiles = files;
	dimensions = new Map(files.filter((f) => f.kind === 'dimension').map((f) => [f.id, f.json]));
	version++;
	const shape = shapeOf(files);
	if (shape !== shapeSignature) {
		shapeSignature = shape;
		shapeVersion++;
	}
	lastSetup = { type: 'setup', files };
	for (const helper of helpers) helper.port.postMessage(lastSetup);
}

let latest = 0;
self.onmessage = async (e: MessageEvent<PreviewRequest>) => {
	const req = e.data;
	if (req.type === 'setup') {
		setup(req.files);
		return;
	}
	if (req.type === 'helpers') {
		addHelpers(req.ports);
		return;
	}
	if (req.type === 'serve') {
		const { port } = req;
		port.onmessage = (m: MessageEvent<HelperMessage>) => serve(port, m.data);
		return;
	}
	latest = req.seq;
	try {
		const res = await render(req);
		if (res.seq !== latest) return; // superseded while fetching
		if (res.type === 'rendered') post(res);
	} catch (err) {
		if (err instanceof Superseded || req.seq !== latest) return;
		(self as unknown as Worker).postMessage({ type: 'error', seq: req.seq, message: String(err instanceof Error ? err.message : err) } satisfies PreviewResponse);
	}
};
