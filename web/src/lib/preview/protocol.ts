export type PreviewKind = 'noise' | 'density_function' | 'noise_settings' | 'material_rule' | 'dimension';

/**
 * What a render shows:
 * - density: sampled values (noise, density functions, noise settings' final density)
 * - blocks: a terrain slice from the chunk generator, with aquifers and surface rules
 * - biomes: a top-down map of a dimension's biome source
 */
export type PreviewMode = 'density' | 'blocks' | 'biomes';

/** The modes each kind offers, the first being the default. */
export const modesOf: Record<PreviewKind, PreviewMode[]> = {
	noise: ['density'],
	density_function: ['density'],
	noise_settings: ['blocks', 'density'],
	material_rule: ['blocks'],
	dimension: ['biomes', 'blocks']
};

/** Whether a view looks down on x/z (otherwise it is a side view of x/y at z = 0). */
export function topDown(kind: PreviewKind, mode: PreviewMode): boolean {
	return kind === 'noise' || mode === 'biomes';
}

export interface PreviewView {
	/** World coordinate at the centre of the image. */
	x: number;
	y: number;
	/** Blocks per pixel. */
	scale: number;
}

export interface PreviewOptions {
	/** The biome surface rules see, for kinds without a biome source. */
	biome: string;
	/** The noise settings a surface rule is drawn on. */
	terrain: string;
	/** The height a biome map is sampled at. */
	y: number;
}

export interface PreviewFile {
	kind: PreviewKind | 'material_condition' | 'tag/worldgen/biome';
	id: string;
	json: unknown;
}

export type PreviewRequest =
	| { type: 'setup'; files: PreviewFile[] }
	| {
			type: 'render';
			seq: number;
			kind: PreviewKind;
			mode: PreviewMode;
			id: string;
			width: number;
			height: number;
			view: PreviewView;
			seed: string;
			options: PreviewOptions;
	  };

export type PreviewResponse =
	| {
			type: 'rendered';
			seq: number;
			width: number;
			height: number;
			pixels: Uint8ClampedArray<ArrayBuffer>;
			/** Sampled values, or for block and biome modes an index into labels (NaN for none). */
			values: Float32Array<ArrayBuffer>;
			min: number;
			max: number;
			/** Block or biome ids that values index, for block and biome modes. */
			labels?: string[];
			/** For biome maps: the biomes in view with their colours, most common first. */
			legend?: { id: string; color: string }[];
			/** For terrain: the sea level and world height. */
			seaLevel?: number;
			minY?: number;
			maxY?: number;
	  }
	| { type: 'error'; seq: number; message: string };

/**
 * Vanilla worldgen that MMS output references comes from mcmeta's summary
 * of 26.3 (one bundle per registry), the same version the spec pack uses.
 */
export const vanillaSummary = 'https://raw.githubusercontent.com/misode/mcmeta/d96c75fec200c4580dd75e033a76341521165461/data';

const pattern = /^data\/([^/]+)\/(worldgen\/(?:noise|density_function|noise_settings|material_rule|material_condition)|dimension|tags\/worldgen\/biome)\/(.+)\.json$/;

/** The registry folder and id of an output path the preview worker reads, or null. */
export function previewFile(path: string): { kind: PreviewFile['kind']; id: string } | null {
	const m = pattern.exec(path);
	if (!m) return null;
	const kind = m[2] === 'tags/worldgen/biome' ? 'tag/worldgen/biome' : m[2].replace('worldgen/', '');
	return { kind: kind as PreviewFile['kind'], id: `${m[1]}:${m[3]}` };
}

/** The preview kind and id of an output path, or null if it has no preview. */
export function previewable(path: string | null): { kind: PreviewKind; id: string } | null {
	const f = path ? previewFile(path) : null;
	return f && f.kind !== 'material_condition' && f.kind !== 'tag/worldgen/biome' ? { kind: f.kind, id: f.id } : null;
}
