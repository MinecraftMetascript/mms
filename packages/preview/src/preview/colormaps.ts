// Colour maps for density previews (noises, density functions, a noise
// settings' final density). Values are mapped onto a palette over a range:
// by default the smallest and largest value in view, so a field that only
// varies between 0.2 and 0.3 still uses the whole palette.

import type { Rgb } from './colors';

export type PaletteId = 'grayscale' | 'viridis' | 'magma' | 'terrain' | 'diverging' | 'solid';

export interface Palette {
	id: PaletteId;
	name: string;
	/** Evenly spaced colour stops, low to high. */
	stops: Rgb[];
	/** Centred on zero: the range is made symmetric, so zero is always the middle colour. */
	centred?: boolean;
}

const hex = (s: string): Rgb => [parseInt(s.slice(1, 3), 16), parseInt(s.slice(3, 5), 16), parseInt(s.slice(5, 7), 16)];

export const palettes: Palette[] = [
	{ id: 'grayscale', name: 'Grayscale', stops: [hex('#000000'), hex('#ffffff')] },
	{
		id: 'viridis',
		name: 'Viridis',
		stops: ['#440154', '#482878', '#3e4989', '#31688e', '#26828e', '#1f9e89', '#35b779', '#6ece58', '#b5de2b', '#fde725'].map(hex)
	},
	{
		id: 'magma',
		name: 'Magma',
		stops: ['#000004', '#180f3d', '#440f76', '#721f81', '#9e2f7f', '#cd4071', '#f1605d', '#fd9668', '#feca8d', '#fcfdbf'].map(hex)
	},
	{
		id: 'terrain',
		name: 'Terrain',
		// Deep water, shallows, lowland, grass, hills, rock, snow.
		stops: ['#1a2a6c', '#2f6fbf', '#e8d9a0', '#5aa04a', '#2f7a3a', '#8a6f52', '#9a948c', '#ffffff'].map(hex)
	},
	{
		id: 'diverging',
		name: 'Blue / red (around 0)',
		stops: ['#2166ac', '#4393c3', '#92c5de', '#d1e5f0', '#f7f7f7', '#fddbc7', '#f4a582', '#d6604d', '#b2182b'].map(hex),
		centred: true
	}
];

/** For a noise settings' final density: solid stone where it is positive, water or sky elsewhere (the worker draws it). */
export const solidPalette: Palette = { id: 'solid', name: 'Solid / air', stops: [hex('#6e6e74'), hex('#b4b4ba')] };

export function paletteOf(id: PaletteId): Palette {
	return id === 'solid' ? solidPalette : (palettes.find((p) => p.id === id) ?? palettes[0]);
}

const tables = new Map<PaletteId, Uint8ClampedArray>();
/** 256 RGB entries, interpolated between the stops. */
function table(p: Palette): Uint8ClampedArray {
	let t = tables.get(p.id);
	if (t) return t;
	t = new Uint8ClampedArray(256 * 3);
	const n = p.stops.length - 1;
	for (let i = 0; i < 256; i++) {
		const f = (i / 255) * n;
		const k = Math.min(n - 1, Math.floor(f));
		const w = f - k;
		const a = p.stops[k];
		const b = p.stops[k + 1];
		for (let c = 0; c < 3; c++) t[i * 3 + c] = a[c] + (b[c] - a[c]) * w;
	}
	tables.set(p.id, t);
	return t;
}

/** The value range a palette is drawn over. */
export interface Range {
	lo: number;
	hi: number;
}

/**
 * The range to draw values in [min, max] over. Centred palettes are made
 * symmetric about zero. A flat field (min = max) gets a range around it, so
 * it is drawn in the palette's middle colour rather than an end.
 */
export function rangeFor(p: Palette, min: number, max: number, fixed?: Range): Range {
	let { lo, hi } = fixed ?? { lo: min, hi: max };
	if (!Number.isFinite(lo) || !Number.isFinite(hi)) return { lo: -1, hi: 1 };
	if (p.centred && !fixed) {
		const m = Math.max(Math.abs(lo), Math.abs(hi));
		lo = -m;
		hi = m;
	}
	if (hi - lo < 1e-9) {
		const pad = Math.max(Math.abs(lo) * 0.5, 0.5);
		return { lo: lo - pad, hi: hi + pad };
	}
	return { lo, hi };
}

/** Where a value falls along the palette, 0 to 1 (clamped). */
export function position(v: number, r: Range): number {
	return Math.max(0, Math.min(1, (v - r.lo) / (r.hi - r.lo)));
}

/** Fills RGBA pixels from values; NaN (outside the world) gets `none`. */
export function colorize(values: Float32Array, p: Palette, r: Range, none: Rgb, out = new Uint8ClampedArray(values.length * 4)) {
	const t = table(p);
	const scale = 255 / (r.hi - r.lo);
	for (let i = 0; i < values.length; i++) {
		const v = values[i];
		const o = i * 4;
		if (Number.isNaN(v)) {
			out[o] = none[0];
			out[o + 1] = none[1];
			out[o + 2] = none[2];
		} else {
			const k = Math.max(0, Math.min(255, Math.round((v - r.lo) * scale))) * 3;
			out[o] = t[k];
			out[o + 1] = t[k + 1];
			out[o + 2] = t[k + 2];
		}
		out[o + 3] = 255;
	}
	return out;
}

/** A CSS linear-gradient of the palette, left (low) to right (high). */
export function cssGradient(p: Palette): string {
	const n = p.stops.length - 1;
	return `linear-gradient(to right, ${p.stops.map((c, i) => `rgb(${c.join(' ')}) ${((i / n) * 100).toFixed(1)}%`).join(', ')})`;
}

/** Tick labels for a range: a few round numbers, trimmed to what the range needs. */
export function formatValue(v: number, r: Range): string {
	const span = Math.abs(r.hi - r.lo);
	const digits = span >= 100 ? 0 : span >= 10 ? 1 : span >= 1 ? 2 : span >= 0.1 ? 3 : 4;
	const s = v.toFixed(digits);
	return s === `-${(0).toFixed(digits)}` ? s.slice(1) : s;
}
