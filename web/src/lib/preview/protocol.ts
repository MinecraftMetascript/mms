export type PreviewKind = 'noise' | 'density_function' | 'noise_settings';

export interface PreviewView {
	/** World coordinate at the centre of the image. */
	x: number;
	y: number;
	/** Blocks per pixel. */
	scale: number;
}

export type PreviewRequest =
	| { type: 'setup'; files: { kind: PreviewKind; id: string; json: unknown }[] }
	| {
			type: 'render';
			seq: number;
			kind: PreviewKind;
			id: string;
			width: number;
			height: number;
			view: PreviewView;
			seed: string;
	  };

export type PreviewResponse =
	| {
			type: 'rendered';
			seq: number;
			width: number;
			height: number;
			pixels: Uint8ClampedArray<ArrayBuffer>;
			values: Float32Array<ArrayBuffer>;
			min: number;
			max: number;
			/** For noise settings: the sea level and world height. */
			seaLevel?: number;
			minY?: number;
			maxY?: number;
	  }
	| { type: 'error'; seq: number; message: string };

/** Vanilla files that MMS output references are fetched from mcmeta at the same commit the spec pack uses. */
export const vanillaBase = 'https://raw.githubusercontent.com/misode/mcmeta/538b2b167248c648b2198f2c0d56eced10dfc0cf/data';

const pattern = /^data\/([^/]+)\/worldgen\/(noise|density_function|noise_settings)\/(.+)\.json$/;

/** The preview kind and id of an output path, or null if it has no preview. */
export function previewable(path: string | null): { kind: PreviewKind; id: string } | null {
	const m = path ? pattern.exec(path) : null;
	return m ? { kind: m[2] as PreviewKind, id: `${m[1]}:${m[3]}` } : null;
}
