import { isObj, numberRange, strip, type Json } from './common';

export interface Swatch {
	name: string;
	color: string;
	/** True when the game derives it from temperature and downfall. */
	derived?: boolean;
}

export interface Spawn {
	mob: string;
	weight: number;
	chance: number;
	count: string;
}

export interface BiomeInfo {
	temperature: number | null;
	downfall: number | null;
	precipitation: 'rain' | 'snow' | 'none';
	frozen: boolean;
	colors: Swatch[];
	spawns: { category: string; mobs: Spawn[] }[];
	/** Placed features per generation step. */
	features: { step: string; count: number }[];
	carvers: number;
	music: string[];
}

const steps = [
	'raw_generation',
	'lakes',
	'local_modifications',
	'underground_structures',
	'surface_structures',
	'strongholds',
	'underground_ores',
	'underground_decoration',
	'fluid_springs',
	'vegetal_decoration',
	'top_layer_modification'
];

/** A colour as `#rrggbb`, from a hex string or a packed integer. */
export function hexColor(v: Json | undefined): string | null {
	if (typeof v === 'number') return '#' + (v & 0xffffff).toString(16).padStart(6, '0');
	if (typeof v === 'string' && /^#[0-9a-f]{6}([0-9a-f]{2})?$/i.test(v)) return v.slice(0, 7).toLowerCase();
	return null;
}

const rgb = (c: string) => [1, 3, 5].map((i) => parseInt(c.slice(i, i + 2), 16));

/** The colormap lookup, approximated by blending the triangle's corners. */
function colormap(corners: [string, string, string], temp: number, downfall: number): string {
	const t = Math.min(1, Math.max(0, temp));
	const d = Math.min(1, Math.max(0, downfall)) * t;
	const [cold, dry, wet] = corners.map(rgb);
	const mix = cold.map((_, i) => Math.round(cold[i] * (1 - t) + dry[i] * (t - d) + wet[i] * d));
	return '#' + mix.map((x) => x.toString(16).padStart(2, '0')).join('');
}

const grassMap: [string, string, string] = ['#80b497', '#bfb755', '#47cd33'];
const foliageMap: [string, string, string] = ['#60a17b', '#aea42a', '#1abf00'];

/** Summarises a biome: climate, colours, mob spawns, feature counts. */
export function analyseBiome(b: Json): BiomeInfo | null {
	if (!isObj(b)) return null;
	const temperature = typeof b.temperature === 'number' ? b.temperature : null;
	const downfall = typeof b.downfall === 'number' ? b.downfall : null;
	const frozen = b.temperature_modifier === 'frozen';
	const rains = b.has_precipitation !== false;
	const precipitation = !rains ? 'none' : temperature !== null && temperature < 0.15 ? 'snow' : 'rain';
	const effects = isObj(b.effects) ? b.effects : {};
	const attrs = isObj(b.attributes) ? b.attributes : {};
	const colors: Swatch[] = [];
	const add = (name: string, v: Json | undefined) => {
		const c = hexColor(isObj(v) && 'argument' in v ? v.argument : v);
		if (c) colors.push({ name, color: c });
	};
	for (const [k, v] of Object.entries(attrs)) {
		const m = /^(?:minecraft:)?visual\/(.+_color)$/.exec(k);
		if (m) add(m[1].replace(/_/g, ' '), v);
	}
	add('water color', effects.water_color);
	for (const [k, map] of [['grass_color', grassMap], ['foliage_color', foliageMap]] as const) {
		if (effects[k] !== undefined) add(k.replace('_', ' '), effects[k]);
		else if (temperature !== null && downfall !== null)
			colors.push({ name: k.replace('_', ' '), color: colormap(map, temperature, downfall), derived: true });
	}
	add('dry foliage color', effects.dry_foliage_color);

	const spawns: BiomeInfo['spawns'] = [];
	const ms = attrs['minecraft:gameplay/natural_mob_spawns'];
	const arg = isObj(ms) ? (isObj(ms.argument) ? ms.argument : ms) : null;
	const byCat = arg && isObj(arg.spawns_by_category) ? arg.spawns_by_category : isObj(b.spawners) ? b.spawners : null;
	if (byCat) {
		for (const [category, list] of Object.entries(byCat)) {
			if (!Array.isArray(list) || !list.length) continue;
			const mobs = list.filter(isObj);
			const total = mobs.reduce((s, m) => s + (typeof m.weight === 'number' ? m.weight : 1), 0) || 1;
			spawns.push({
				category: category.replace(/_/g, ' '),
				mobs: mobs.map((m) => {
					const weight = typeof m.weight === 'number' ? m.weight : 1;
					let count: string;
					if (m.minCount !== undefined) count = m.minCount === m.maxCount ? String(m.minCount) : `${m.minCount}–${m.maxCount}`;
					else {
						const r = numberRange(m.count);
						count = !r ? '?' : r.min === r.max ? String(r.min) : `${r.min}–${r.max}`;
					}
					return { mob: String(m.type ?? '?'), weight, chance: weight / total, count };
				})
			});
		}
	}
	const features = Array.isArray(b.features)
		? b.features.map((f, i) => ({ step: (steps[i] ?? `step ${i}`).replace(/_/g, ' '), count: typeof f === 'string' ? 1 : Array.isArray(f) ? f.length : 0 }))
		: [];
	const carvers = Array.isArray(b.carvers) ? b.carvers.length : typeof b.carvers === 'string' ? 1 : 0;
	const music: string[] = [];
	const bg = attrs['minecraft:audio/background_music'];
	const collect = (v: Json | undefined) => {
		if (isObj(v)) {
			if (typeof v.sound === 'string') music.push(strip(v.sound));
			for (const x of Object.values(v)) if (isObj(x)) collect(x);
		}
	};
	collect(isObj(bg) && 'argument' in bg ? bg.argument : bg);
	return { temperature, downfall, precipitation, frozen, colors, spawns, features, carvers, music: [...new Set(music)] };
}
