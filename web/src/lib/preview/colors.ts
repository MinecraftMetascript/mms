// Map colours for blocks and biomes. Vanilla names get a hand-picked
// colour; anything else gets a stable colour from a hash of its id.

export type Rgb = readonly [number, number, number];

export const SKY: Rgb = [186, 218, 247];
export const VOID: Rgb = [30, 30, 36];
export const CAVE: Rgb = [62, 58, 66];

const blocks: Record<string, Rgb> = {
	air: SKY,
	cave_air: [120, 120, 130],
	void_air: VOID,
	stone: [125, 125, 125],
	deepslate: [80, 80, 86],
	cobblestone: [116, 116, 116],
	cobbled_deepslate: [76, 76, 80],
	bedrock: [55, 55, 55],
	tuff: [108, 109, 102],
	calcite: [223, 224, 220],
	granite: [149, 103, 85],
	diorite: [188, 188, 188],
	andesite: [136, 136, 136],
	dirt: [134, 96, 67],
	coarse_dirt: [119, 85, 59],
	rooted_dirt: [144, 103, 76],
	grass_block: [109, 153, 48],
	podzol: [91, 63, 24],
	mycelium: [111, 99, 105],
	mud: [60, 57, 60],
	moss_block: [89, 109, 45],
	sand: [219, 207, 163],
	sandstone: [216, 203, 155],
	red_sand: [190, 102, 33],
	red_sandstone: [186, 99, 29],
	gravel: [131, 127, 126],
	clay: [160, 166, 179],
	snow: [249, 254, 254],
	snow_block: [249, 254, 254],
	powder_snow: [248, 253, 253],
	ice: [145, 183, 253],
	packed_ice: [141, 180, 250],
	blue_ice: [116, 167, 253],
	water: [52, 104, 196],
	lava: [207, 92, 15],
	obsidian: [15, 10, 24],
	netherrack: [97, 38, 38],
	soul_sand: [81, 62, 50],
	soul_soil: [75, 57, 46],
	basalt: [80, 81, 86],
	blackstone: [42, 36, 41],
	crimson_nylium: [130, 31, 31],
	warped_nylium: [43, 114, 101],
	nether_wart_block: [114, 2, 2],
	warped_wart_block: [22, 119, 121],
	end_stone: [219, 222, 158],
	terracotta: [152, 94, 67],
	white_terracotta: [209, 178, 161],
	orange_terracotta: [161, 83, 37],
	yellow_terracotta: [186, 133, 35],
	brown_terracotta: [77, 51, 35],
	red_terracotta: [143, 61, 46],
	light_gray_terracotta: [135, 106, 97],
	honey_block: [251, 185, 52],
	slime_block: [111, 192, 91],
	magma_block: [142, 63, 31],
	glowstone: [171, 131, 84],
	prismarine: [99, 156, 151],
	dripstone_block: [134, 107, 92],
	sculk: [12, 29, 36],
	amethyst_block: [133, 97, 191]
};

const biomes: Record<string, Rgb> = {
	the_void: [0, 0, 0],
	plains: [141, 179, 96],
	sunflower_plains: [181, 219, 136],
	snowy_plains: [255, 255, 255],
	ice_spikes: [180, 220, 220],
	desert: [250, 148, 24],
	swamp: [7, 249, 178],
	mangrove_swamp: [44, 204, 142],
	forest: [5, 102, 33],
	flower_forest: [45, 142, 73],
	birch_forest: [48, 116, 68],
	dark_forest: [64, 81, 26],
	pale_garden: [178, 184, 178],
	old_growth_birch_forest: [88, 156, 108],
	old_growth_pine_taiga: [89, 102, 81],
	old_growth_spruce_taiga: [129, 142, 121],
	taiga: [11, 102, 89],
	snowy_taiga: [49, 85, 74],
	savanna: [189, 178, 95],
	savanna_plateau: [167, 157, 100],
	windswept_hills: [96, 96, 96],
	windswept_gravelly_hills: [136, 136, 136],
	windswept_forest: [80, 112, 80],
	windswept_savanna: [229, 218, 135],
	jungle: [83, 123, 9],
	sparse_jungle: [98, 139, 23],
	bamboo_jungle: [118, 142, 20],
	badlands: [217, 69, 21],
	eroded_badlands: [255, 109, 61],
	wooded_badlands: [176, 151, 101],
	meadow: [96, 164, 69],
	cherry_grove: [255, 145, 200],
	grove: [71, 114, 108],
	snowy_slopes: [196, 196, 196],
	frozen_peaks: [160, 160, 160],
	jagged_peaks: [220, 220, 200],
	stony_peaks: [123, 143, 116],
	river: [0, 0, 255],
	frozen_river: [160, 160, 255],
	beach: [250, 222, 85],
	snowy_beach: [250, 240, 192],
	stony_shore: [162, 162, 132],
	warm_ocean: [0, 0, 172],
	lukewarm_ocean: [0, 0, 144],
	deep_lukewarm_ocean: [0, 0, 64],
	ocean: [0, 0, 112],
	deep_ocean: [0, 0, 48],
	cold_ocean: [32, 32, 112],
	deep_cold_ocean: [32, 32, 56],
	frozen_ocean: [112, 112, 214],
	deep_frozen_ocean: [64, 64, 144],
	mushroom_fields: [255, 0, 255],
	dripstone_caves: [134, 96, 67],
	lush_caves: [0, 255, 0],
	deep_dark: [3, 31, 41],
	nether_wastes: [191, 59, 59],
	warped_forest: [73, 144, 123],
	crimson_forest: [221, 8, 8],
	soul_sand_valley: [94, 56, 48],
	basalt_deltas: [64, 54, 54],
	the_end: [128, 128, 255],
	end_highlands: [181, 181, 54],
	end_midlands: [255, 255, 153],
	small_end_islands: [144, 144, 223],
	end_barrens: [112, 112, 204]
};

function hashed(id: string, light: number): Rgb {
	let h = 2166136261;
	for (let i = 0; i < id.length; i++) h = Math.imul(h ^ id.charCodeAt(i), 16777619);
	const hue = ((h >>> 0) % 360) / 360;
	// HSL → RGB with fixed saturation and the given lightness.
	const s = 0.55;
	const q = light < 0.5 ? light * (1 + s) : light + s - light * s;
	const p = 2 * light - q;
	const f = (t: number) => {
		t = (t + 1) % 1;
		const v = t < 1 / 6 ? p + (q - p) * 6 * t : t < 1 / 2 ? q : t < 2 / 3 ? p + (q - p) * (2 / 3 - t) * 6 : p;
		return Math.round(v * 255);
	};
	return [f(hue + 1 / 3), f(hue), f(hue - 1 / 3)];
}

function vanillaName(id: string): string | null {
	const i = id.indexOf(':');
	if (i < 0) return id;
	return id.slice(0, i) === 'minecraft' ? id.slice(i + 1) : null;
}

/** A colour for a block id such as `minecraft:stone`. */
export function blockColor(id: string): Rgb {
	const name = vanillaName(id);
	if (name !== null) {
		const c = blocks[name];
		if (c) return c;
		// Families: anything terracotta, wool, concrete, ore, log…
		if (name.endsWith('_ore')) return name.startsWith('deepslate') ? blocks.deepslate : blocks.stone;
		if (name.includes('terracotta')) return blocks.terracotta;
	}
	return hashed(id, 0.55);
}

/** A colour for a biome id such as `minecraft:plains`. */
export function biomeColor(id: string): Rgb {
	const name = vanillaName(id);
	return (name !== null && biomes[name]) || hashed(id, 0.5);
}
