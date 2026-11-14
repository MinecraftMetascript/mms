// Rulers and a scale bar for the previews: steps that read well at any
// zoom, and where their ticks land.

/** The smallest of 1, 2 and 5 times a power of ten (whole blocks, at least 1) that is at least `min`. */
export function niceStep(min: number): number {
	if (!(min > 1)) return 1;
	const p = 10 ** Math.floor(Math.log10(min));
	for (const m of [1, 2, 5, 10]) if (m * p >= min) return m * p;
	return 10 * p;
}

/**
 * The multiples of step between lo and hi (either way round). Counted, not
 * stepped, and capped, so a huge coordinate can't loop forever.
 */
export function multiples(lo: number, hi: number, step: number): number[] {
	const a = Math.min(lo, hi);
	const b = Math.max(lo, hi);
	const first = Math.ceil(a / step);
	const n = Math.max(0, Math.min(200, Math.floor(b / step) - first + 1));
	return Array.from({ length: n }, (_, i) => (first + i) * step);
}

/** A tick on a ruler: its world coordinate and its offset along the edge in CSS pixels. */
export interface Tick {
	value: number;
	at: number;
}

/**
 * The ticks along an edge `length` CSS pixels long, where the coordinate at
 * offset p is origin + per·p (per is negative when it counts down the edge).
 * Ticks are at least `gap` pixels apart.
 */
export function ticks(origin: number, per: number, length: number, gap: number): Tick[] {
	if (!per || !Number.isFinite(per) || !Number.isFinite(origin) || length <= 0) return [];
	const step = niceStep(Math.abs(per) * gap);
	return multiples(origin, origin + per * length, step).map((value) => ({ value, at: (value - origin) / per }));
}

/**
 * A scale bar for a view at `blocksPerPixel`: a power of two blocks (so
 * it lands on chunks and regions) whose bar is about 60 to 120 pixels long.
 */
export function scaleBar(blocksPerPixel: number, target = 80): { blocks: number; px: number; label: string } | null {
	if (!(blocksPerPixel > 0) || !Number.isFinite(blocksPerPixel)) return null;
	const blocks = 2 ** Math.max(0, Math.round(Math.log2(target * blocksPerPixel)));
	const px = blocks / blocksPerPixel;
	let label = `${blocks} block${blocks === 1 ? '' : 's'}`;
	if (blocks >= 512) label += ` · ${blocks / 512} region${blocks === 512 ? '' : 's'}`;
	else if (blocks >= 16) label += ` · ${blocks / 16} chunk${blocks === 16 ? '' : 's'}`;
	return { blocks, px, label };
}
