// The density tool: samples a noise, density function, noise settings'
// final density, or a terrain's surface, through the preview worker.

import type { PreviewFile, PreviewKind } from '../preview/protocol';
import { ArgError, fmt, num, pct, pick, str, table } from './common';
import type { VanillaFiles } from './host';
import { render, type Rendered } from './shim';

type Kind = Extract<PreviewKind, 'noise' | 'density_function' | 'noise_settings' | 'dimension'>;
const kinds: Kind[] = ['noise', 'density_function', 'noise_settings', 'dimension'];
const vanillaFolder: Partial<Record<Kind, string>> = {
	noise: 'worldgen/noise',
	density_function: 'worldgen/density_function',
	noise_settings: 'worldgen/noise_settings'
};
type Mode = 'area' | 'column' | 'surface';

const options = { biome: 'minecraft:plains', terrain: 'minecraft:overworld' };

/** The kind and full id an agent's id names: the pack's files first, then vanilla. */
function resolve(id: string, kind: Kind | undefined, files: PreviewFile[], vanilla: VanillaFiles): { kind: Kind; id: string } {
	// A bare name means the pack's own before vanilla's, as in MMS code.
	const ofKind = (k: Kind) => {
		const own = files.filter((f) => f.kind === k).map((f) => f.id);
		const folder = vanillaFolder[k];
		const theirs = folder ? Object.keys(vanilla[folder] ?? {}).map((p) => `minecraft:${p}`) : [];
		return { own, all: [...new Set([...own, ...theirs])] };
	};
	const find = (k: Kind) => {
		const { own, all } = ofKind(k);
		if (!id.includes(':')) {
			try {
				return pick(k.replace('_', ' '), id, own);
			} catch {
				// Not the pack's own: try vanilla's too.
			}
		}
		return pick(k.replace('_', ' '), id, all);
	};
	if (kind) return { kind, id: find(kind) };
	const found: { kind: Kind; id: string }[] = [];
	for (const k of kinds) {
		try {
			found.push({ kind: k, id: find(k) });
		} catch {
			// Not of this kind.
		}
	}
	if (found.length === 1) return found[0];
	if (found.length > 1) throw new ArgError(`${id} names more than one thing: give kind, one of ${found.map((f) => f.kind).join(', ')}.`);
	throw new ArgError(`No noise, density function, noise settings or dimension is called ${id}.`);
}

export async function density(args: Record<string, unknown>, files: PreviewFile[], vanilla: VanillaFiles): Promise<string> {
	const given = str(args, 'id');
	if (!given) throw new ArgError('Give id: the noise, density function, noise settings or dimension to sample.');
	const k = str(args, 'kind');
	if (k && !kinds.includes(k as Kind)) throw new ArgError(`kind must be one of ${kinds.join(', ')}.`);
	const { kind, id } = resolve(given, k as Kind | undefined, files, vanilla);
	const seed = String(args.seed ?? '0');
	const mode = (str(args, 'mode') ?? (kind === 'dimension' ? 'surface' : 'area')) as Mode;
	if (!['area', 'column', 'surface'].includes(mode)) throw new ArgError('mode must be area, column or surface.');
	const x = Math.floor(num(args, 'x', 0));
	const z = Math.floor(num(args, 'z', 0));
	const what = `${kind.replace('_', ' ')} ${id}, seed ${seed}`;

	if (mode === 'surface') {
		if (kind !== 'noise_settings' && kind !== 'dimension') throw new ArgError('mode surface needs noise settings or a dimension: it builds the terrain.');
		return surface(what, kind, id, seed, x, z, args);
	}
	if (kind === 'dimension') throw new ArgError("A dimension has no density of its own: use mode surface, or sample its noise settings' final density.");
	if (mode === 'column') {
		if (kind === 'noise') throw new ArgError('A noise is sampled top-down at one y; use mode area, with y.');
		return column(what, kind, id, seed, x, z, args);
	}
	return area(what, kind, id, seed, x, z, args);
}

/** Values over a top-down square at one y: their spread and a coarse table. */
async function area(what: string, kind: Kind, id: string, seed: string, x: number, z: number, args: Record<string, unknown>): Promise<string> {
	const y = num(args, 'y', 64);
	const radius = Math.floor(num(args, 'radius', 256, 1, 32768));
	const n = Math.floor(num(args, 'size', 8, 2, 32));
	// Statistics come from a finer grid than the table.
	const fine = 64;
	const stats = await render({ kind, mode: 'density', id, width: fine, height: fine, view: { x, y, z, facing: 'north', scale: (2 * radius) / fine, down: true }, seed, options });
	const grid = await render({ kind, mode: 'density', id, width: n, height: n, view: { x, y, z, facing: 'north', scale: (2 * radius) / n, down: true }, seed, options });
	const out = [`${what}, at y ${fmt(y)}, x ${x - radius}..${x + radius}, z ${z - radius}..${z + radius}.`];
	out.push(spread(stats));
	out.push('', `Values on a ${n}×${n} grid, ${fmt((2 * radius) / n)} blocks apart (rows run north to south, columns west to east):`);
	const step = (2 * radius) / n;
	// The coordinates the worker samples at, which fall between blocks when the step is odd.
	const xs = Array.from({ length: n }, (_, i) => x + (i + 0.5 - n / 2) * step);
	const rows = [['z \\ x', ...xs.map((v) => fmt(v, 1))]];
	for (let py = 0; py < n; py++) {
		const zz = z + (py + 0.5 - n / 2) * step;
		rows.push([fmt(zz, 1), ...xs.map((_, px) => fmt(grid.values[py * n + px], 2))]);
	}
	out.push(table(rows));
	return out.join('\n');
}

/** The spread of a render's values: range, mean, and how much is above zero. */
function spread(r: Rendered): string {
	let sum = 0;
	let count = 0;
	let positive = 0;
	for (const v of r.values) {
		if (Number.isNaN(v)) continue;
		sum += v;
		count++;
		if (v > 0) positive++;
	}
	if (!count) return 'No values: the area is outside the world height.';
	return `Min ${fmt(r.min)}, max ${fmt(r.max)}, mean ${fmt(sum / count)}; above 0 at ${pct(positive, count)} of ${count} samples.`;
}

/** Values down one column, with the heights where the sign changes. */
async function column(what: string, kind: Kind, id: string, seed: string, x: number, z: number, args: Record<string, unknown>): Promise<string> {
	const from = Math.floor(num(args, 'from', -64, -2048, 2048));
	const to = Math.floor(num(args, 'to', 320, -2048, 2048));
	if (to <= from) throw new ArgError('to must be above from.');
	// At most 512 samples: a longer range takes a larger step rather than stopping short.
	const step = Math.max(Math.floor(num(args, 'step', 8, 1, 512)), Math.ceil((to - from) / 511));
	const n = Math.floor((to - from) / step) + 1;
	// A side view one pixel wide: row r is y = from + (n - 1 - r) * step.
	const r = await render({ kind, mode: 'density', id, width: 1, height: n, view: { x, y: from + (n / 2 - 0.5) * step, z, facing: 'north', scale: step }, seed, options });
	const ys = Array.from({ length: n }, (_, i) => from + (n - 1 - i) * step);
	const out = [`${what}, the column at x ${x}, z ${z}, y ${from}..${ys[0]} every ${step}.`, spread(r)];
	// Where the sign flips, top down: for a final density, the surface and cave edges.
	const flips: string[] = [];
	for (let i = 1; i < n; i++) {
		const a = r.values[i - 1];
		const b = r.values[i];
		if (Number.isNaN(a) || Number.isNaN(b) || a > 0 === b > 0) continue;
		flips.push(`between y ${ys[i]} and ${ys[i - 1]} (${b > 0 ? 'positive below' : 'positive above'})`);
	}
	out.push(flips.length ? `Crosses 0 ${flips.join('; ')}.` : 'Never crosses 0 in this range.');
	out.push('', table([['y', 'value'], ...ys.map((y, i) => [String(y), fmt(r.values[i])])]));
	return out.join('\n');
}

/** The terrain's top blocks over a square: heights, and which blocks show. */
async function surface(what: string, kind: Kind, id: string, seed: string, x: number, z: number, args: Record<string, unknown>): Promise<string> {
	const radius = Math.floor(num(args, 'radius', 32, 1, 64));
	const n = Math.floor(num(args, 'size', 8, 2, 32));
	const top = num(args, 'y', 2032);
	const fine = Math.min(2 * radius, 64);
	const view = (size: number) => ({ x, y: top, z, facing: 'north' as const, scale: (2 * radius) / size, down: true });
	const stats = await render({ kind, mode: 'blocks', id, width: fine, height: fine, view: view(fine), seed, options });
	const grid = await render({ kind, mode: 'blocks', id, width: n, height: n, view: view(n), seed, options });
	const out = [`${what}, the top block of each column${top < 2032 ? ` at or below y ${fmt(top)}` : ''}, x ${x - radius}..${x + radius}, z ${z - radius}..${z + radius}. Sea level is ${stats.seaLevel}.`];
	const labels = stats.labels ?? [];
	const counts = new Map<string, number>();
	let lo = Infinity;
	let hi = -Infinity;
	let sum = 0;
	let count = 0;
	for (let i = 0; i < stats.values.length; i++) {
		const h = stats.heights?.[i] ?? NaN;
		if (Number.isNaN(h)) continue;
		lo = Math.min(lo, h);
		hi = Math.max(hi, h);
		sum += h;
		count++;
		const b = labels[stats.values[i]];
		if (b) counts.set(b, (counts.get(b) ?? 0) + 1);
	}
	if (!count) return [...out, 'No blocks: every column is empty.'].join('\n');
	out.push(`Heights ${lo}..${hi}, mean ${fmt(sum / count, 1)}.`);
	out.push(`Top blocks: ${[...counts].sort((a, b) => b[1] - a[1]).map(([b, c]) => `${b} ${pct(c, count)}`).join(', ')}.`);
	if (stats.notice) out.push(stats.notice);
	const step = (2 * radius) / n;
	const xs = Array.from({ length: n }, (_, i) => Math.floor(x + (i + 0.5 - n / 2) * step));
	const glabels = grid.labels ?? [];
	const short = (b: string | undefined) => (b ? b.replace(/^minecraft:/, '') : '-');
	out.push('', `Height and top block on a ${n}×${n} grid, ${fmt(step)} blocks apart (rows run north to south):`);
	const rows = [['z \\ x', ...xs.map(String)]];
	for (let py = 0; py < n; py++) {
		const zz = Math.floor(z + (py + 0.5 - n / 2) * step);
		rows.push([String(zz), ...xs.map((_, px) => {
			const i = py * n + px;
			const h = grid.heights?.[i] ?? NaN;
			return Number.isNaN(h) ? '-' : `${h} ${short(glabels[grid.values[i]])}`;
		})]);
	}
	out.push(table(rows));
	return out.join('\n');
}
