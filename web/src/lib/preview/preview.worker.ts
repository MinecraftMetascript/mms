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

/** Built chunks (block palette indices, x-major then y) for the last terrain rendered. */
let terrainKey = '';
const terrainChunks = new Map<number, Uint16Array>();

interface Terrain {
	settings: NoiseGeneratorSettings;
	rule: MaterialRule;
	rs: RandomState;
	biomeAt: (x: number, y: number, z: number) => Identifier;
	/** What the terrain's shape depends on, so filled chunks can be shared across surface rule edits. */
	shape: string;
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

/** Filled chunks (terrain shape before surface rules), most recently used last. */
const fills = new Map<string, { chunk: Chunk; noiseChunk: unknown }>();
const MAX_FILLS = 32;

/** A full 16×16 chunk with surface rules, indexed ((x * 16 + z) * height + y). */
function buildChunk3D(t: Terrain, cx: number, cz: number): Uint16Array {
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
				if (state) {
					replaced.push([pos, old]);
					chunk.setBlockState(pos, state);
				}
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
	for (const [pos, old] of replaced) chunk.setBlockState(pos, old);
	return out;
}

let isoKey = '';
const isoChunks = new Map<string, Uint16Array>();

/**
 * The visible faces of an area of chunks, in back-to-front order. Hidden
 * blocks are pruned once per area, so a pan or zoom only projects and
 * fills what can be seen.
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
/** The area drawn at one block size, kept so panning only moves it. */
let isoImage: { key: string; s: number; image: OffscreenCanvas } | null = null;

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

function isoAreaOf(key: string, chunks: (Uint16Array | undefined)[][], n: number, h: number): IsoArea {
	const air = indexOf('minecraft:air');
	const caveAir = indexOf('minecraft:cave_air');
	const at = (x: number, y: number, z: number): number => {
		if (x < 0 || z < 0 || x >= n || z >= n || y < 0 || y >= h) return air;
		const c = chunks[x >> 4][z >> 4];
		return c ? c[((x & 15) * 16 + (z & 15)) * h + y] : air;
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
	const bottom = Math.max(0, lowest - 16);
	const faces: number[] = [];
	const fills: string[] = [];
	// Back to front: every block only hides blocks with smaller x, z or y.
	for (let x = 0; x < n; x++) {
		for (let z = 0; z < n; z++) {
			for (let y = bottom; y <= highest; y++) {
				const b = at(x, y, z);
				if (empty(b)) continue;
				const f = fillsOf(palette[b]);
				if (y === highest || empty(at(x, y + 1, z))) faces.push(x, y, z, 0), fills.push(f[0]);
				if (x === n - 1 || empty(at(x + 1, y, z))) faces.push(x, y, z, 1), fills.push(f[1]);
				if (z === n - 1 || empty(at(x, y, z + 1))) faces.push(x, y, z, 2), fills.push(f[2]);
			}
		}
	}
	return { key, n, bottom, top: highest, faces: Int16Array.from(faces), fills };
}

/** Fills an area's faces, with block (x, y, z) at screen (ox + (x - z)·s, oy + (x + z)·s/2 - y·s), skipping those outside w × h. */
function drawFaces(ctx: OffscreenCanvasRenderingContext2D, area: IsoArea, s: number, ox: number, oy: number, w: number, h: number) {
	const { faces, fills } = area;
	ctx.lineWidth = 0.5;
	for (let i = 0, f = 0; i < faces.length; i += 4, f++) {
		const x = faces[i], y = faces[i + 1], z = faces[i + 2], side = faces[i + 3];
		// The block's top-back corner; a block spans 2s across and 2s down from it.
		const bx = ox + (x - z) * s;
		const by = oy + ((x + z) * s) / 2 - (y + 1) * s;
		if (bx + s < 0 || bx - s > w || by + 2 * s < 0 || by > h) continue;
		ctx.fillStyle = ctx.strokeStyle = fills[f];
		ctx.beginPath();
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
		ctx.fill();
		ctx.stroke(); // hides hairline seams between faces
	}
}

/** The largest area image kept for panning, in pixels. */
const MAX_IMAGE = 4096 * 4096;

/**
 * Draws the terrain as an isometric block diagram. In this mode the view's
 * x and y are the world x and z at the centre, and its scale is the size of
 * a block in pixels (0 to fit a 2×2-chunk area to the canvas). Zooming out
 * builds more chunks, up to 4×4.
 *
 * Built chunks, the pruned face list and the drawn area are each cached, so
 * a pan only moves an image and a zoom only redraws the visible faces.
 * While chunks are being built, each one is shown as it finishes.
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
	const fit = !(req.view.scale > 0);
	// Blocks across the area: enough to fill the canvas at this zoom.
	const across = fit ? 32 : width / (2 * req.view.scale);
	const chunksAcross = across <= 40 ? 2 : across <= 56 ? 3 : 4;
	const n = chunksAcross * 16;
	const cx = req.view.x;
	const cz = req.view.y;
	const cx0 = Math.round(cx / 16 - chunksAcross / 2);
	const cz0 = Math.round(cz / 16 - chunksAcross / 2);
	const chunks: (Uint16Array | undefined)[][] = [];
	const missing: [number, number][] = [];
	for (let i = 0; i < chunksAcross; i++) {
		chunks.push([]);
		for (let j = 0; j < chunksAcross; j++) {
			const c = isoChunks.get(`${cx0 + i},${cz0 + j}`);
			chunks[i].push(c);
			if (!c) missing.push([i, j]);
		}
	}

	const frame = (): Rendered => {
		const areaKey = `${cx0},${cz0},${chunksAcross},${missing.length}`;
		if (isoArea?.key !== areaKey) {
			isoArea = isoAreaOf(areaKey, chunks, n, h);
			isoImage = null;
		}
		const area = isoArea;
		const canvas = new OffscreenCanvas(width, height);
		const ctx = canvas.getContext('2d')!;
		ctx.fillStyle = `rgb(${SKY.join(' ')})`;
		ctx.fillRect(0, 0, width, height);
		const empty = area.bottom > area.top;
		const span = area.top + 1 - area.bottom;
		const s = fit ? Math.max(1, Math.min(width / (2 * n), height / (n + Math.max(span, 1))) * 0.92) : req.view.scale;
		if (!empty) {
			// Screen position of block (x, y, z), area-local, with the view's
			// centre (at mid height) in the middle of the canvas.
			const lx = cx - cx0 * 16;
			const lz = cz - cz0 * 16;
			const yc = area.bottom + span / 2;
			const ox = width / 2 - (lx - lz) * s;
			const oy = height / 2 - ((lx + lz) * s) / 2 + yc * s;
			const iw = Math.ceil(2 * n * s) + 2;
			const ih = Math.ceil((n + span) * s) + 2;
			if (iw * ih <= MAX_IMAGE) {
				// Draw the whole area once at this size; panning then only moves it.
				if (!isoImage || isoImage.key !== area.key || isoImage.s !== s) {
					const image = new OffscreenCanvas(iw, ih);
					drawFaces(image.getContext('2d')!, area, s, n * s + 1, (area.top + 1) * s + 1, iw, ih);
					isoImage = { key: area.key, s, image };
				}
				ctx.drawImage(isoImage.image, Math.round(ox - n * s - 1), Math.round(oy - (area.top + 1) * s - 1));
			} else {
				drawFaces(ctx, area, s, ox, oy, width, height);
			}
		}
		const notice =
			empty && missing.length === 0 ? 'No blocks in these chunks. Drag to pan to some terrain.' : undefined;
		const bitmap = canvas.transferToImageBitmap();
		return {
			type: 'rendered', seq: req.seq, width, height,
			pixels: new Uint8ClampedArray(0), values: new Float32Array(0),
			min: Infinity, max: -Infinity, scale: s, bitmap, notice
		};
	};

	while (missing.length > 0) {
		// Show what is built so far, then build the next chunk.
		if (isoChunks.size > 0) post({ ...frame(), partial: true });
		await breathe();
		if (req.seq !== latest) throw new Superseded();
		const [i, j] = missing[0];
		const c = buildChunk3D(t, cx0 + i, cz0 + j);
		isoChunks.set(`${cx0 + i},${cz0 + j}`, c);
		chunks[i][j] = c;
		missing.shift();
	}
	return frame();
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

/** The last terrain set up, so panning doesn't rebuild its random state each frame. */
let cached: { key: string; t: Terrain } | null = null;

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

	if (req.kind === 'dimension' && req.mode !== 'biomes' && cached?.key === terrainKey('')) {
		const key = terrainKey('');
		await warm(cached.t, key);
		return req.mode === 'iso' ? renderIso(req, cached.t, key) : renderTerrain(req, cached.t, key);
	}
	if (req.kind === 'dimension') {
		const { dim, rs } = await withVanilla(() => {
			const dim = readDimension(req.id);
			const rs = new RandomState(dim.settings, seed);
			// Touch the climate router, so missing entries are fetched now.
			rs.sampler.sample(0, 0, 0);
			return { dim, rs };
		});
		if (req.mode === 'biomes') return renderBiomes(req, dim, rs);
		const key = terrainKey('');
		const terrain: Terrain = {
			settings: dim.settings,
			rule: dim.settings.materialRule,
			rs,
			biomeAt: (x, y, z) => dim.biomes.getBiome(x >> 2, y >> 2, z >> 2, rs.sampler),
			shape: [shapeVersion, req.seed, 'dimension', req.id].join('|')
		};
		cached = { key, t: terrain };
		await warm(terrain, key);
		return req.mode === 'iso' ? renderIso(req, terrain, key) : renderTerrain(req, terrain, key);
	}

	if (req.kind === 'material_rule' || (req.kind === 'noise_settings' && req.mode !== 'density')) {
		const key = terrainKey(`${req.options.terrain}|${fixedBiome}`);
		const terrain = cached?.key === key ? cached.t : await withVanilla((): Terrain => {
			const settingsId = req.kind === 'material_rule' ? Identifier.parse(req.options.terrain || 'minecraft:overworld') : id;
			const settings = WorldgenRegistries.NOISE_SETTINGS.getOrThrow(settingsId);
			const rule = req.kind === 'material_rule' ? WorldgenRegistries.MATERIAL_RULE.getOrThrow(id) : settings.materialRule;
			const rs = new RandomState(settings, seed);
			const shape = [shapeVersion, req.seed, 'noise_settings', settingsId].join('|');
			return { settings, rule, rs, biomeAt: () => fixedBiome, shape };
		});
		cached = { key, t: terrain };
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

function post(res: Rendered) {
	const transfer: Transferable[] = [res.pixels.buffer, res.values.buffer];
	if (res.bitmap) transfer.push(res.bitmap);
	(self as unknown as Worker).postMessage(res, transfer);
}

let latest = 0;
self.onmessage = async (e: MessageEvent<PreviewRequest>) => {
	const req = e.data;
	if (req.type === 'setup') {
		userFiles = req.files;
		dimensions = new Map(req.files.filter((f) => f.kind === 'dimension').map((f) => [f.id, f.json]));
		version++;
		const shape = shapeOf(req.files);
		if (shape !== shapeSignature) {
			shapeSignature = shape;
			shapeVersion++;
		}
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
