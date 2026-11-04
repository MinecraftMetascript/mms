import type { PreviewKind, PreviewFileKind } from '../paths';
export type { PreviewKind } from '../paths';

/**
 * What a render shows:
 * - density: sampled values (noise, density functions, noise settings' final density)
 * - blocks: a terrain slice from the chunk generator, with aquifers and surface rules
 * - biomes: a top-down map of a dimension's biome source
 * - iso: an isometric block diagram of a couple of chunks of terrain
 */
export type PreviewMode = 'density' | 'blocks' | 'biomes' | 'iso';

/** The modes each kind offers, the first being the default. */
export const modesOf: Record<PreviewKind, PreviewMode[]> = {
	noise: ['density'],
	density_function: ['density'],
	noise_settings: ['blocks', 'iso', 'density'],
	material_rule: ['blocks', 'iso'],
	dimension: ['biomes', 'blocks', 'iso']
};

/** Whether a view looks down on x/z (otherwise it is a side view: a vertical slice through the camera). */
export function topDown(kind: PreviewKind, mode: PreviewMode): boolean {
	return kind === 'noise' || mode === 'biomes';
}

/** The way a 2D view's camera faces. Always along an axis, so a side view slices along x or z. */
export type Facing = 'north' | 'east' | 'south' | 'west';
export const facings: Facing[] = ['north', 'east', 'south', 'west'];

/**
 * The world directions of a 2D view's screen axes, as unit x/z steps:
 * right is the screen's right, and down is the screen's down in a
 * top-down view (the way behind the camera). A side view's slice runs
 * along right, through the camera, so north and south slice along x
 * (at the camera's z), east and west along z (at the camera's x).
 */
export function screenAxes(facing: Facing): { rx: number; rz: number; dx: number; dz: number } {
	switch (facing) {
		case 'north':
			return { rx: 1, rz: 0, dx: 0, dz: 1 };
		case 'east':
			return { rx: 0, rz: 1, dx: -1, dz: 0 };
		case 'south':
			return { rx: -1, rz: 0, dx: 0, dz: -1 };
		case 'west':
			return { rx: 0, rz: -1, dx: 1, dz: 0 };
	}
}

/**
 * Where a view looks. For the 2D views this is the camera: the world
 * point at the centre of the image and the way it faces. A side view
 * shows the slice through it; a top-down view samples at its y (a biome
 * map's height). The isometric view uses x and y as the world x and z
 * at the centre, and ignores z and facing.
 */
export interface PreviewView {
	x: number;
	y: number;
	z: number;
	facing: Facing;
	/** Blocks per pixel (for the isometric view, a block's size in pixels). */
	scale: number;
}

export interface PreviewOptions {
	/** The biome surface rules see, for kinds without a biome source. */
	biome: string;
	/** The noise settings a surface rule is drawn on. */
	terrain: string;
}

export interface PreviewFile {
	kind: PreviewFileKind;
	id: string;
	json: unknown;
}

export type PreviewRequest =
	| { type: 'setup'; files: PreviewFile[] }
	/** Ports to helper workers that build chunks for the isometric view. */
	| { type: 'helpers'; ports: MessagePort[] }
	/** Makes this worker a helper, taking chunk jobs on the port. */
	| { type: 'serve'; port: MessagePort }
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
			/** For the isometric view: the block size in pixels it was drawn at. */
			scale?: number;
			/** For the isometric view: the smallest block size it allows (it fills the canvas with chunks). */
			minScale?: number;
			/** For the isometric view: the image itself, drawn in the worker (pixels is then empty). */
			bitmap?: ImageBitmap;
			/** More of this render follows (chunks are still being built). */
			partial?: boolean;
			/** Something to tell the viewer about an otherwise empty render. */
			notice?: string;
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

export { previewFile, previewable } from '../paths';
