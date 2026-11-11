// Climate parameters for biome maps. A multi-noise biome source picks, at
// each 4×4×4 cell, the entry whose parameter ranges are nearest the climate
// sampled there (the sum of squared distances outside each range, plus the
// entry's offset squared). These helpers explain that choice: what the
// climate is at a point, which entries came closest, and how far a biome is
// from winning.

/** The six climate parameters, in the order the game and the JSON list them. */
export const climateParams = ['temperature', 'humidity', 'continentalness', 'erosion', 'depth', 'weirdness'] as const;
export type ClimateParam = (typeof climateParams)[number];

/** What a biome map can show: the biomes, one climate parameter, or how far a biome is from winning. */
export type BiomeLayer = 'biomes' | ClimateParam | 'gap';

/**
 * A multi-noise source's entries, packed: entry i's biome is biomes[i], and
 * params[i * 13 ...] holds min and max for each climate parameter in order,
 * then the offset.
 */
export interface ClimateEntries {
	biomes: string[];
	params: Float32Array<ArrayBuffer>;
}
export const entryStride = 13;

/** A climate sample: the six values, in climateParams order. */
export type ClimatePoint = ArrayLike<number>;

/** How far value v is outside [min, max] (0 inside). */
function outside(v: number, min: number, max: number): number {
	return v < min ? min - v : v > max ? v - max : 0;
}

/** The fitness of entry i at a climate point: the game's squared distance (lower wins). */
export function fitness(entries: ClimateEntries, i: number, at: ClimatePoint): number {
	const p = entries.params;
	const o = i * entryStride;
	let d = 0;
	for (let c = 0; c < 6; c++) {
		const e = outside(at[c], p[o + c * 2], p[o + c * 2 + 1]);
		d += e * e;
	}
	const off = p[o + 12];
	return d + off * off;
}

/** One parameter of an entry against the climate at a point. */
export interface ParamMiss {
	param: ClimateParam;
	min: number;
	max: number;
	value: number;
	/** How far the value is outside the range (0 inside). */
	off: number;
}

export interface NearEntry {
	index: number;
	biome: string;
	/** The game's distance: the square root of the fitness, so it reads in parameter units. */
	distance: number;
	params: ParamMiss[];
	offset: number;
}

function describe(entries: ClimateEntries, i: number, at: ClimatePoint): NearEntry {
	const p = entries.params;
	const o = i * entryStride;
	const params = climateParams.map((param, c) => {
		const min = p[o + c * 2];
		const max = p[o + c * 2 + 1];
		return { param, min, max, value: at[c], off: outside(at[c], min, max) };
	});
	return { index: i, biome: entries.biomes[i], distance: Math.sqrt(fitness(entries, i, at)), params, offset: p[o + 12] };
}

/**
 * The entries nearest a climate point, best first, at most one per biome
 * (a biome's nearest entry stands for it). The first is the one the game
 * picks, up to ties.
 */
export function nearest(entries: ClimateEntries, at: ClimatePoint, count = 5): NearEntry[] {
	const best = new Map<string, { i: number; f: number }>();
	for (let i = 0; i < entries.biomes.length; i++) {
		const f = fitness(entries, i, at);
		const b = entries.biomes[i];
		const cur = best.get(b);
		if (!cur || f < cur.f) best.set(b, { i, f });
	}
	return [...best.values()]
		.sort((a, b) => a.f - b.f || a.i - b.i)
		.slice(0, count)
		.map(({ i }) => describe(entries, i, at));
}

/** A biome's nearest entry at a climate point, or null if the source has no entry for it. */
export function nearestOf(entries: ClimateEntries, biome: string, at: ClimatePoint): NearEntry | null {
	let bi = -1;
	let bf = Infinity;
	for (let i = 0; i < entries.biomes.length; i++) {
		if (entries.biomes[i] !== biome) continue;
		const f = fitness(entries, i, at);
		if (f < bf) {
			bf = f;
			bi = i;
		}
	}
	return bi < 0 ? null : describe(entries, bi, at);
}

/**
 * How far a biome is from winning at each pixel: its nearest entry's
 * distance minus the winner's (both as square roots of fitness), 0 where it
 * wins. `climate` is planar (parameter c of pixel i at c * n + i) and `fit`
 * is the winner's fitness per pixel; NaN stays NaN.
 */
export function gapLayer(entries: ClimateEntries, biome: string, climate: Float32Array, fit: Float32Array): Float32Array<ArrayBuffer> {
	const n = fit.length;
	const out = new Float32Array(n);
	const own: number[] = [];
	for (let i = 0; i < entries.biomes.length; i++) if (entries.biomes[i] === biome) own.push(i);
	const at = new Float64Array(6);
	// Neighbouring pixels in one cell share a climate, so reuse the last answer.
	let prevOut = NaN;
	let have = false;
	for (let i = 0; i < n; i++) {
		const w = fit[i];
		if (Number.isNaN(w) || own.length === 0) {
			out[i] = NaN;
			have = false;
			continue;
		}
		let same = have;
		for (let c = 0; c < 6; c++) {
			const v = climate[c * n + i];
			if (v !== at[c]) same = false;
			at[c] = v;
		}
		if (!same) {
			let f = Infinity;
			for (const e of own) f = Math.min(f, fitness(entries, e, at));
			prevOut = Math.max(0, Math.sqrt(f) - Math.sqrt(w));
			have = true;
		}
		out[i] = prevOut;
	}
	return out;
}

/** The biomes a source lists that don't win anywhere in view. */
export function missing(entries: ClimateEntries | undefined, inView: Iterable<string>): string[] {
	if (!entries) return [];
	const seen = new Set(inView);
	return [...new Set(entries.biomes)].filter((b) => !seen.has(b)).sort();
}
