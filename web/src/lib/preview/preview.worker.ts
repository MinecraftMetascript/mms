/// <reference lib="webworker" />
import {
	BiomeSource,
	BlockPos,
	Chunk,
	ChunkPos,
	DensityFunction,
	BlockState,
	FixedBiomeSource,
	MaterialRuleContext,
	Holder,
	HolderSet,
	Identifier,
	NoiseChunkGenerator,
	NoiseGeneratorSettings,
	NoiseRouter,
	RandomState,
	WorldgenRegistries,
	type MaterialRule
} from 'deepslate';
import { biomeColor, blockColor, CAVE, SKY, VOID, type Rgb } from './colors';
import { vanillaSummary, type PreviewFile, type PreviewRequest, type PreviewResponse } from './protocol';

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

/** Built chunks (block palette indices, x-major then y) for the last terrain rendered. */
let terrainKey = '';
const terrainChunks = new Map<number, Uint16Array>();

interface Terrain {
	settings: NoiseGeneratorSettings;
	rule: MaterialRule;
	rs: RandomState;
	biomeAt: (x: number, y: number, z: number) => Identifier;
}

function buildChunk(t: Terrain, chunkX: number): Uint16Array {
	const { settings } = t;
	const minY = settings.noise.minY;
	const height = settings.noise.height;
	const chunk = new Chunk(minY, height, ChunkPos.create(chunkX, 0));
	// buildTerrain takes one biome per chunk; doing its steps here lets the
	// surface rules see the biome at each block.
	const gen = new NoiseChunkGenerator(new FixedBiomeSource(Identifier.create('plains')), settings) as unknown as {
		createNoiseChunk(rs: RandomState, chunk: Chunk, onlyFirstZ: boolean): unknown;
		doFill(chunk: Chunk, noiseChunk: unknown): void;
		getWorldgenContext(): unknown;
	};
	const noiseChunk = gen.createNoiseChunk(t.rs, chunk, true);
	gen.doFill(chunk, noiseChunk);
	(t.rs.surfaceSystem as unknown as { buildSurface(...a: unknown[]): void }).buildSurface(
		chunk,
		noiseChunk,
		t.rule,
		gen.getWorldgenContext(),
		(pos: BlockPos) => t.biomeAt(pos[0], pos[1], pos[2])
	);
	const out = new Uint16Array(16 * height);
	const x0 = chunkX * 16;
	for (let x = 0; x < 16; x++) {
		for (let y = 0; y < height; y++) {
			out[x * height + y] = indexOf(chunk.getBlockState(BlockPos.create(x0 + x, minY + y, 0)).getName().toString());
		}
	}
	return out;
}

// ---- Isometric terrain -------------------------------------------------

/** A full 16×16 chunk with surface rules, indexed ((x * 16 + z) * height + y). */
function buildChunk3D(t: Terrain, cx: number, cz: number): Uint16Array {
	const { settings } = t;
	const minY = settings.noise.minY;
	const height = settings.noise.height;
	const chunk = new Chunk(minY, height, ChunkPos.create(cx, cz));
	const gen = new NoiseChunkGenerator(new FixedBiomeSource(Identifier.create('plains')), settings) as unknown as {
		createNoiseChunk(rs: RandomState, chunk: Chunk, onlyFirstZ: boolean): unknown;
		doFill(chunk: Chunk, noiseChunk: unknown): void;
		getWorldgenContext(): unknown;
	};
	const noiseChunk = gen.createNoiseChunk(t.rs, chunk, false);
	gen.doFill(chunk, noiseChunk);
	// Deepslate's MaterialSystem.buildSurface only paints the first z row
	// (its previews are slices), so this is its loop over the whole chunk.
	const context = new MaterialRuleContext(
		t.rs.surfaceSystem,
		t.rs,
		chunk,
		noiseChunk as never,
		gen.getWorldgenContext() as never,
		(pos: BlockPos) => t.biomeAt(pos[0], pos[1], pos[2])
	);
	const rule = t.rule.compile(context);
	const x0 = cx * 16;
	const z0 = cz * 16;
	const maxY = minY + height - 1;
	for (let x = 0; x < 16; x++) {
		for (let z = 0; z < 16; z++) {
			const bx = x0 + x;
			const bz = z0 + z;
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
				if (state) chunk.setBlockState(pos, state);
			}
		}
	}
	const out = new Uint16Array(16 * 16 * height);
	for (let x = 0; x < 16; x++) {
		for (let z = 0; z < 16; z++) {
			const base = (x * 16 + z) * height;
			for (let y = 0; y < height; y++) {
				out[base + y] = indexOf(chunk.getBlockState(BlockPos.create(x0 + x, minY + y, z0 + z)).getName().toString());
			}
		}
	}
	return out;
}

let isoKey = '';
const isoChunks = new Map<string, Uint16Array>();
/** Chunks across the isometric view (it shows ISO_CHUNKS² chunks); even, so it centres on a chunk corner. */
const ISO_CHUNKS = 2;

/** Draws the terrain around the view's x (and z = 0) as an isometric block diagram. */
async function renderIso(req: RenderRequest, t: Terrain, key: string): Promise<Rendered> {
	const { width, height } = req;
	const h = t.settings.noise.height;
	if (key !== isoKey) {
		isoKey = key;
		isoChunks.clear();
	}
	const n = ISO_CHUNKS * 16;
	// Centred on the slice's x, with z = 0 (the slice) running through the middle.
	const cx0 = Math.round(req.view.x / 16) - ISO_CHUNKS / 2;
	const cz0 = -ISO_CHUNKS / 2;
	const chunks: Uint16Array[][] = [];
	for (let i = 0; i < ISO_CHUNKS; i++) {
		chunks.push([]);
		for (let j = 0; j < ISO_CHUNKS; j++) {
			const k = `${cx0 + i},${cz0 + j}`;
			let c = isoChunks.get(k);
			if (!c) {
				await breathe();
				if (req.seq !== latest) throw new Superseded();
				c = buildChunk3D(t, cx0 + i, cz0 + j);
				isoChunks.set(k, c);
			}
			chunks[i].push(c);
		}
	}
	const air = indexOf('minecraft:air');
	const caveAir = indexOf('minecraft:cave_air');
	const at = (x: number, y: number, z: number): number => {
		if (x < 0 || z < 0 || x >= n || z >= n || y < 0 || y >= h) return air;
		return chunks[x >> 4][z >> 4][((x & 15) * 16 + (z & 15)) * h + y];
	};
	const empty = (b: number) => b === air || b === caveAir;

	// Show from a little below the lowest surface to the highest one, so the
	// cut sides show the layers the surface rules paint.
	let lowest = h;
	let highest = 0;
	for (let x = 0; x < n; x++) {
		for (let z = 0; z < n; z++) {
			let y = h - 1;
			while (y > 0 && empty(at(x, y, z))) y--;
			if (empty(at(x, y, z))) continue; // open to the void, as under sky islands
			lowest = Math.min(lowest, y);
			highest = Math.max(highest, y);
		}
	}
	if (lowest > highest) {
		const pixels = new Uint8ClampedArray(width * height * 4);
		for (let i = 0; i < width * height; i++) put(pixels, i * 4, SKY);
		const notice = 'No blocks in these chunks. Pan the Blocks view to some terrain, then switch back.';
		return { type: 'rendered', seq: req.seq, width, height, pixels, values: new Float32Array(width * height).fill(NaN), min: Infinity, max: -Infinity, notice };
	}
	const bottom = Math.max(0, lowest - 16);
	const span = highest + 1 - bottom;

	const canvas = new OffscreenCanvas(width, height);
	const ctx = canvas.getContext('2d')!;
	ctx.fillStyle = `rgb(${SKY.join(' ')})`;
	ctx.fillRect(0, 0, width, height);
	const s = Math.max(1, Math.min(width / (2 * n), height / (n + span)) * 0.92);
	const ox = width / 2;
	const oy = (height - (n + span) * s) / 2 + span * s;
	const px = (x: number, y: number, z: number): [number, number] => [ox + (x - z) * s, oy + ((x + z) * s) / 2 - (y - bottom) * s];
	const face = (pts: [number, number][], c: Rgb, shade: number) => {
		const fill = `rgb(${Math.round(c[0] * shade)} ${Math.round(c[1] * shade)} ${Math.round(c[2] * shade)})`;
		ctx.fillStyle = fill;
		ctx.strokeStyle = fill;
		ctx.beginPath();
		ctx.moveTo(pts[0][0], pts[0][1]);
		for (let i = 1; i < pts.length; i++) ctx.lineTo(pts[i][0], pts[i][1]);
		ctx.closePath();
		ctx.fill();
		ctx.stroke(); // hides hairline seams between faces
	};
	ctx.lineWidth = 0.5;
	// Back to front: every block only hides blocks with smaller x, z or y.
	for (let x = 0; x < n; x++) {
		for (let z = 0; z < n; z++) {
			for (let y = bottom; y <= highest; y++) {
				const b = at(x, y, z);
				if (empty(b)) continue;
				const c = blockColor(palette[b]);
				const above = y === highest || empty(at(x, y + 1, z));
				const right = x === n - 1 || empty(at(x + 1, y, z));
				const front = z === n - 1 || empty(at(x, y, z + 1));
				if (above) face([px(x, y + 1, z), px(x + 1, y + 1, z), px(x + 1, y + 1, z + 1), px(x, y + 1, z + 1)], c, 1);
				if (right) face([px(x + 1, y + 1, z), px(x + 1, y + 1, z + 1), px(x + 1, y, z + 1), px(x + 1, y, z)], c, 0.62);
				if (front) face([px(x, y + 1, z + 1), px(x + 1, y + 1, z + 1), px(x + 1, y, z + 1), px(x, y, z + 1)], c, 0.8);
			}
		}
	}
	const pixels = new Uint8ClampedArray(ctx.getImageData(0, 0, width, height).data.buffer);
	const values = new Float32Array(width * height).fill(NaN);
	return { type: 'rendered', seq: req.seq, width, height, pixels, values, min: Infinity, max: -Infinity };
}

/**
 * Builds the chunk at the view's centre inside withVanilla, so entries the
 * surface rules reference are fetched before the full pass.
 */
async function warm(t: Terrain, key: string): Promise<Terrain> {
	if (key !== terrainKey) {
		terrainKey = key;
		terrainChunks.clear();
	}
	if (terrainChunks.size === 0) {
		await withVanilla(() => terrainChunks.set(0, buildChunk(t, 0)));
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
		terrainChunks.clear();
	}
	const x0 = view.x - (width / 2) * view.scale;
	const yTop = view.y + (height / 2) * view.scale;
	const values = new Float32Array(width * height).fill(NaN);
	const pixels = new Uint8ClampedArray(width * height * 4);
	let started = performance.now();
	for (let px = 0; px < width; px++) {
		const x = Math.floor(x0 + px * view.scale);
		const cx = Math.floor(x / 16);
		let col = terrainChunks.get(cx);
		if (!col) {
			if (performance.now() - started > 40) {
				await breathe();
				if (req.seq !== latest) throw new Superseded();
				started = performance.now();
			}
			col = buildChunk(t, cx);
			terrainChunks.set(cx, col);
		}
		const lx = x - cx * 16;
		const h = t.settings.noise.height;
		// Air below the router's estimate of the surface is cave, drawn darker than sky.
		const surface = t.rs.router.chunkSurfaceLevel.compute(DensityFunction.context(x, 0, 0));
		for (let py = 0; py < height; py++) {
			const y = Math.floor(yTop - py * view.scale);
			const i = py * width + px;
			if (y < minY || y >= maxY) {
				put(pixels, i * 4, VOID);
				continue;
			}
			const b = col[lx * h + (y - minY)];
			const id = palette[b];
			values[i] = b;
			if (id === 'minecraft:air') {
				put(pixels, i * 4, y < surface ? CAVE : SKY);
			} else {
				put(pixels, i * 4, blockColor(id));
			}
		}
	}
	// Keep the cache from growing without bound while panning.
	if (terrainChunks.size > 512) {
		const keep = new Set<number>();
		for (let px = 0; px < width; px++) keep.add(Math.floor(Math.floor(x0 + px * view.scale) / 16));
		for (const k of terrainChunks.keys()) if (!keep.has(k)) terrainChunks.delete(k);
	}
	return { type: 'rendered', seq: req.seq, width, height, pixels, values, min: Infinity, max: -Infinity, labels: palette.slice(), seaLevel: sea, minY, maxY };
}

// ---- Biome maps -------------------------------------------------------

async function renderBiomes(req: RenderRequest, dim: Dimension, rs: RandomState): Promise<Rendered> {
	const { width, height, view } = req;
	const values = new Float32Array(width * height);
	const pixels = new Uint8ClampedArray(width * height * 4);
	const x0 = view.x - (width / 2) * view.scale;
	const z0 = view.y - (height / 2) * view.scale;
	const qy = Math.floor(req.options.y) >> 2;
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
		const qz = Math.floor(z0 + py * view.scale) >> 2;
		for (let px = 0; px < width; px++) {
			const qx = Math.floor(x0 + px * view.scale) >> 2;
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

async function render(req: RenderRequest): Promise<PreviewResponse> {
	const { width, height, view } = req;
	const seed = seedOf(req.seed);
	const id = Identifier.parse(req.id);
	const values = new Float32Array(width * height);
	const pixels = new Uint8ClampedArray(width * height * 4);
	const x0 = view.x - (width / 2) * view.scale;
	const yTop = view.y + (height / 2) * view.scale;
	let min = Infinity;
	let max = -Infinity;
	const track = (v: number) => {
		if (v < min) min = v;
		if (v > max) max = v;
	};
	const fixedBiome = Identifier.parse(req.options.biome || 'minecraft:plains');
	const terrainKey = (what: string) => [version, req.seed, req.kind, req.id, req.mode, what].join('|');

	if (req.kind === 'dimension') {
		const { dim, rs } = await withVanilla(() => {
			const dim = readDimension(req.id);
			const rs = new RandomState(dim.settings, seed);
			// Touch the climate router, so missing entries are fetched now.
			rs.sampler.sample(0, 0, 0);
			return { dim, rs };
		});
		if (req.mode === 'biomes') return renderBiomes(req, dim, rs);
		const terrain: Terrain = {
			settings: dim.settings,
			rule: dim.settings.materialRule,
			rs,
			biomeAt: (x, y, z) => dim.biomes.getBiome(x >> 2, y >> 2, z >> 2, rs.sampler)
		};
		const key = terrainKey('');
		await warm(terrain, key);
		return req.mode === 'iso' ? renderIso(req, terrain, key) : renderTerrain(req, terrain, key);
	}

	if (req.kind === 'material_rule' || (req.kind === 'noise_settings' && req.mode !== 'density')) {
		const key = terrainKey(`${req.options.terrain}|${fixedBiome}`);
		const terrain = await withVanilla((): Terrain => {
			const settingsId = req.kind === 'material_rule' ? Identifier.parse(req.options.terrain || 'minecraft:overworld') : id;
			const settings = WorldgenRegistries.NOISE_SETTINGS.getOrThrow(settingsId);
			const rule = req.kind === 'material_rule' ? WorldgenRegistries.MATERIAL_RULE.getOrThrow(id) : settings.materialRule;
			const rs = new RandomState(settings, seed);
			return { settings, rule, rs, biomeAt: () => fixedBiome };
		});
		await warm(terrain, key);
		return req.mode === 'iso' ? renderIso(req, terrain, key) : renderTerrain(req, terrain, key);
	}

	if (req.kind === 'noise') {
		// Top-down: x across, z down, at y = 0.
		const noise = await withVanilla(() => {
			const rs = new RandomState(NoiseGeneratorSettings.create({}), seed);
			return rs.getOrCreateNoise(id);
		});
		const z0 = view.y - (height / 2) * view.scale;
		for (let py = 0; py < height; py++) {
			for (let px = 0; px < width; px++) {
				const v = noise.get3D(x0 + px * view.scale, 0, z0 + py * view.scale);
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
		// Side view: x across, y up, at z = 0.
		const fn = await withVanilla(() => {
			const rs = new RandomState(NoiseGeneratorSettings.create({}), seed);
			// A reference holder, so the visitor binds noises and caches by id.
			const df = new DensityFunction.HolderHolder(Holder.reference(WorldgenRegistries.DENSITY_FUNCTION, id));
			return df.mapAll(rs.createVisitor(false));
		});
		for (let py = 0; py < height; py++) {
			const y = yTop - py * view.scale;
			for (let px = 0; px < width; px++) {
				const v = fn.compute(DensityFunction.context(x0 + px * view.scale, y, 0));
				values[py * width + px] = v;
				track(v);
			}
		}
		const range = Math.max(Math.abs(min), Math.abs(max), 1e-9);
		for (let i = 0; i < values.length; i++) diverging(values[i] / range, pixels, i * 4);
		return { type: 'rendered', seq: req.seq, width, height, pixels, values, min, max };
	}

	// Noise settings, density mode: the terrain's final density, side view at z = 0.
	const { settings, router } = await withVanilla(() => {
		const settings = WorldgenRegistries.NOISE_SETTINGS.getOrThrow(id);
		const rs = new RandomState(settings, seed);
		return { settings, router: rs.router as NoiseRouter };
	});
	const minY = settings.noise.minY;
	const maxY = minY + settings.noise.height;
	const sea = settings.seaLevel;
	for (let py = 0; py < height; py++) {
		const y = yTop - py * view.scale;
		for (let px = 0; px < width; px++) {
			const i = py * width + px;
			const o = i * 4;
			let v = NaN;
			if (y >= minY && y < maxY) {
				v = router.finalDensity.compute(DensityFunction.context(x0 + px * view.scale, y, 0));
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

let latest = 0;
self.onmessage = async (e: MessageEvent<PreviewRequest>) => {
	const req = e.data;
	if (req.type === 'setup') {
		userFiles = req.files;
		dimensions = new Map(req.files.filter((f) => f.kind === 'dimension').map((f) => [f.id, f.json]));
		version++;
		return;
	}
	latest = req.seq;
	try {
		const res = await render(req);
		if (res.seq !== latest) return; // superseded while fetching
		if (res.type === 'rendered') {
			(self as unknown as Worker).postMessage(res, [res.pixels.buffer, res.values.buffer]);
		}
	} catch (err) {
		if (err instanceof Superseded || req.seq !== latest) return;
		(self as unknown as Worker).postMessage({ type: 'error', seq: req.seq, message: String(err instanceof Error ? err.message : err) } satisfies PreviewResponse);
	}
};
