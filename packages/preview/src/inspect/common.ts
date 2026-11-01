/** Which inspector, if any, shows a generated file. */
export type InspectKind = 'loot_table' | 'recipe' | 'biome';

const pattern = /^data\/([^/]+)\/(loot_table|recipe|worldgen\/biome)\/(.+)\.json$/;

export function inspectable(path: string | null): { kind: InspectKind; id: string } | null {
	const m = path ? pattern.exec(path) : null;
	if (!m) return null;
	const kind = (m[2] === 'worldgen/biome' ? 'biome' : m[2]) as InspectKind;
	return { kind, id: `${m[1]}:${m[3]}` };
}

/** `minecraft:golden_apple` → `Golden Apple`; other namespaces are kept. */
export function label(id: string): string {
	const tag = id.startsWith('#');
	const raw = tag ? id.slice(1) : id;
	const [ns, path] = raw.includes(':') ? raw.split(':', 2) : ['minecraft', raw];
	const name = path
		.split('/')
		.pop()!
		.split('_')
		.map((w) => (w ? w[0].toUpperCase() + w.slice(1) : w))
		.join(' ');
	return (tag ? '#' : '') + (ns === 'minecraft' ? name : `${name} (${ns})`);
}

export const strip = (s: string) => s.replace(/^minecraft:/, '');

export type Json = null | boolean | number | string | Json[] | { [k: string]: Json };
export type Obj = { [k: string]: Json };

export const isObj = (v: Json | undefined): v is Obj => typeof v === 'object' && v !== null && !Array.isArray(v);

/** A number provider's range and mean, or null when it depends on context. */
export function numberRange(v: Json | undefined, fallback = 1): { min: number; max: number; mean: number } | null {
	if (v === undefined) return { min: fallback, max: fallback, mean: fallback };
	if (typeof v === 'number') return { min: v, max: v, mean: v };
	if (!isObj(v)) return null;
	const type = strip(String(v.type ?? (v.min !== undefined ? 'uniform' : 'constant')));
	const n = (x: Json | undefined) => (typeof x === 'number' ? x : null);
	switch (type) {
		case 'constant': {
			const c = n(v.value);
			return c === null ? null : { min: c, max: c, mean: c };
		}
		case 'uniform': {
			const lo = n(v.min ?? v.min_inclusive);
			const hi = n(v.max ?? v.max_inclusive);
			return lo === null || hi === null ? null : { min: lo, max: hi, mean: (lo + hi) / 2 };
		}
		case 'binomial': {
			const trials = n(v.n);
			const p = n(v.p);
			return trials === null || p === null ? null : { min: 0, max: trials, mean: trials * p };
		}
	}
	return null;
}

export function fmtRange(r: { min: number; max: number } | null): string {
	if (!r) return '?';
	const f = (x: number) => (Number.isInteger(x) ? String(x) : x.toFixed(2));
	return r.min === r.max ? f(r.min) : `${f(r.min)}–${f(r.max)}`;
}

export function pct(p: number): string {
	if (p >= 0.995) return '100%';
	if (p < 0.001 && p > 0) return '<0.1%';
	return `${(p * 100).toFixed(p < 0.1 ? 1 : 0)}%`;
}

/** A small seeded PRNG (mulberry32), so a roll can be repeated. */
export function rng(seed: number): () => number {
	let a = seed >>> 0;
	return () => {
		a = (a + 0x6d2b79f5) >>> 0;
		let t = a;
		t = Math.imul(t ^ (t >>> 15), t | 1);
		t ^= t + Math.imul(t ^ (t >>> 7), t | 61);
		return ((t ^ (t >>> 14)) >>> 0) / 4294967296;
	};
}
