// The biomes tool: which biomes a dimension's biome source places in an
// area, the climate at a point, and why a biome does or doesn't win. The
// same sampling and the same explanations as the editors' biome map.

import { climateParams, gapLayer, missing, nearest, nearestOf, type ClimateEntries, type NearEntry } from '../preview/climate';
import type { PreviewFile } from '../preview/protocol';
import { ArgError, fmt, num, pct, pick, qualified, str } from './common';
import { render, type Rendered } from './shim';

const options = { biome: 'minecraft:plains', terrain: 'minecraft:overworld' };

/** Samples a square of the dimension's biome map, n by n samples, centred on x and z. */
function sample(id: string, seed: string, x: number, y: number, z: number, n: number, scale: number): Promise<Rendered> {
	return render({ kind: 'dimension', mode: 'biomes', id, width: n, height: n, view: { x, y, z, facing: 'north', scale }, seed, options });
}

/** One entry against a climate point: its distance and the parameters it misses. */
function entryLine(e: NearEntry, winner: boolean): string {
	const misses = e.params.filter((p) => p.off > 0).map((p) => `${p.param} ${fmt(p.value)} is outside ${fmt(p.min)}..${fmt(p.max)} by ${fmt(p.off)}`);
	if (e.offset) misses.push(`offset ${fmt(e.offset)}`);
	const why = misses.length ? `: ${misses.join('; ')}` : ': every parameter in range';
	return `${e.biome} (entry ${e.index}), distance ${fmt(e.distance)}${winner ? ', wins' : ''}${why}`;
}

/** The climate of the sample at index i of a render with n samples. */
function climateAt(r: Rendered, i: number): number[] {
	const n = r.width * r.height;
	return climateParams.map((_, c) => r.climate![c * n + i]);
}

export async function biomes(args: Record<string, unknown>, files: PreviewFile[]): Promise<string> {
	const dims = files.filter((f) => f.kind === 'dimension').map((f) => f.id);
	const given = str(args, 'dimension');
	if (!given && dims.length !== 1) {
		throw new ArgError(dims.length ? `Give dimension: one of ${dims.join(', ')}.` : 'The project has no Dimension with a noise generator to map.');
	}
	const id = given ? pick('dimension', given, dims) : dims[0];
	const seed = String(args.seed ?? '0');
	const x = Math.floor(num(args, 'x', 0));
	const z = Math.floor(num(args, 'z', 0));
	const y = Math.floor(num(args, 'y', 64, -2048, 2048));
	const radius = Math.floor(num(args, 'radius', 2048, 16, 32768));
	const n = 128;
	const scale = (radius * 2) / n;
	const wanted = str(args, 'biome');

	const map = await sample(id, seed, x, y, z, n, scale);
	const entries = map.entries;
	const labels = map.labels ?? [];
	const total = n * n;
	const counts = new Map<string, number>();
	for (let i = 0; i < total; i++) {
		const b = labels[map.values[i]];
		if (b) counts.set(b, (counts.get(b) ?? 0) + 1);
	}

	const out: string[] = [];
	out.push(
		`Dimension ${id}, seed ${seed}, at y ${y}: ${n}×${n} samples ${fmt(scale)} blocks apart, ` +
			`x ${x - radius}..${x + radius}, z ${z - radius}..${z + radius}.`
	);
	if (!entries) out.push("Its biome source isn't an inline multi-noise list, so there are no entries to explain; the climate is still sampled.");
	out.push('', 'Biomes in the area, by share of the samples:');
	for (const [b, c] of [...counts].sort((a, b) => b[1] - a[1])) out.push(`- ${b} ${pct(c, total)}`);
	const absent = missing(entries, counts.keys());
	if (absent.length) out.push(`Listed by the source but not in the area: ${absent.join(', ')}.`);

	// The centre: its climate and the entries nearest it.
	const point = await sample(id, seed, x, y, z, 1, 1);
	const at = climateAt(point, 0);
	out.push('', `At x ${x}, z ${z} (y ${y}): ${point.labels?.[point.values[0]] ?? 'no biome'}.`);
	out.push(`Climate: ${climateParams.map((p, c) => `${p} ${fmt(at[c])}`).join(', ')}.`);
	if (entries) {
		out.push('Nearest entries (the lowest distance wins; at most one per biome):');
		nearest(entries, at, 5).forEach((e, k) => out.push(`${k + 1}. ${entryLine(e, k === 0)}`));
	}

	if (wanted) out.push('', ...explain(biomeId(wanted, [...counts.keys(), ...(entries?.biomes ?? [])]), map, entries, counts, total, x, z, scale));
	return out.join('\n');
}

/**
 * The biome a name means: a bare name matches the biomes the source and the
 * map know in any namespace (the pack's own included), else minecraft:'s.
 */
function biomeId(name: string, known: string[]): string {
	if (name.includes(':')) return name;
	const hits = [...new Set(known)].filter((b) => b.slice(b.indexOf(':') + 1) === name);
	if (hits.length > 1) throw new ArgError(`${name} could be ${hits.join(' or ')}; give the namespace.`);
	return hits[0] ?? qualified(name);
}

/** Why one biome does or doesn't win in the area. */
function explain(biome: string, map: Rendered, entries: ClimateEntries | undefined, counts: Map<string, number>, total: number, x: number, z: number, scale: number): string[] {
	const n = map.width;
	const labels = map.labels ?? [];
	// World coordinates of sample i (its centre), as the worker samples them.
	const where = (i: number) => {
		const px = i % n;
		const py = Math.floor(i / n);
		return { x: Math.floor(x + (px + 0.5 - n / 2) * scale), z: Math.floor(z + (py + 0.5 - n / 2) * scale) };
	};
	const wins = counts.get(biome) ?? 0;
	if (wins > 0) {
		// The winning sample nearest the centre.
		let best = -1;
		let bestD = Infinity;
		for (let i = 0; i < total; i++) {
			if (labels[map.values[i]] !== biome) continue;
			const w = where(i);
			const d = (w.x - x) ** 2 + (w.z - z) ** 2;
			if (d < bestD) {
				bestD = d;
				best = i;
			}
		}
		const w = where(best);
		return [`${biome} wins at ${pct(wins, total)} of the area; the nearest sample to the centre is x ${w.x}, z ${w.z}.`];
	}
	if (!entries) return [`${biome} wins nowhere in the area.`];
	if (!entries.biomes.includes(biome)) {
		return [`${biome} wins nowhere: the biome source has no entry for it.`];
	}
	const out = [`${biome} wins nowhere in the area.`];
	// Where it comes closest: the smallest gap between its nearest entry and the winner.
	const gaps = gapLayer(entries, biome, map.climate!, map.fit!);
	let best = -1;
	for (let i = 0; i < total; i++) if (!Number.isNaN(gaps[i]) && (best < 0 || gaps[i] < gaps[best])) best = i;
	if (best >= 0) {
		const at = climateAt(map, best);
		const own = nearestOf(entries, biome, at)!;
		const winner = nearest(entries, at, 1)[0];
		const w = where(best);
		out.push(`It comes closest at x ${w.x}, z ${w.z}, losing to ${winner.biome} by ${fmt(gaps[best])}.`);
		out.push(`There: ${entryLine(own, false)}.`);
		out.push(`The winner: ${entryLine(winner, true)}.`);
	}
	// Each parameter's range in the area against the ranges its entries ask for.
	const own: number[] = [];
	entries.biomes.forEach((b, i) => b === biome && own.push(i));
	out.push('', `Climate in the area against what ${biome}'s entries ask for:`);
	let depthMiss = false;
	climateParams.forEach((p, c) => {
		let lo = Infinity;
		let hi = -Infinity;
		for (let i = 0; i < total; i++) {
			const v = map.climate![c * total + i];
			if (v < lo) lo = v;
			if (v > hi) hi = v;
		}
		const asks = own.map((e) => `${fmt(entries.params[e * 13 + c * 2])}..${fmt(entries.params[e * 13 + c * 2 + 1])}`);
		const reach = own.some((e) => entries.params[e * 13 + c * 2] <= hi && entries.params[e * 13 + c * 2 + 1] >= lo);
		out.push(`- ${p}: area ${fmt(lo)}..${fmt(hi)}; asks ${[...new Set(asks)].join(' or ')}${reach ? '' : ' (never in range here)'}`);
		if (!reach && p === 'depth') depthMiss = true;
	});
	// Depth follows the height sampled, unlike the other five.
	if (depthMiss) out.push('Depth changes with y (it is 0 at the surface and grows downwards), so try another y.');
	return out;
}
