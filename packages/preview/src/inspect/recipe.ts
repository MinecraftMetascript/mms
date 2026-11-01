import { isObj, strip, type Json } from './common';

/** One ingredient slot: the items or tags that fit it. */
export type Slot = string[] | null;

export interface RecipeInfo {
	type: string;
	/** The layout the view draws. */
	layout: 'grid' | 'cooking' | 'stonecutting' | 'smithing' | 'brewing' | 'special';
	/** Grid recipes: rows of slots, at most 3×3. */
	grid?: Slot[][];
	shapeless?: boolean;
	/** Single-input recipes (cooking, stonecutting). */
	input?: Slot;
	/** Smithing: template, base, addition. */
	smithing?: { template: Slot; base: Slot; addition: Slot };
	/** Brewing: the potion in, the reagent on top. */
	brewing?: { input: Slot; reagent: Slot };
	/** Potion names or a trim pattern, shown under the slots. */
	notes?: { input?: string; result?: string };
	result: { id: string; count: number } | null;
	cookingTime?: number;
	experience?: number;
	group?: string;
	category?: string;
}

/** An ingredient: an id, a #tag, or a list of either. */
function slot(v: Json | undefined): Slot {
	if (v === undefined || v === null) return null;
	if (typeof v === 'string') return [v];
	if (Array.isArray(v)) return v.flatMap((x) => slot(x) ?? []);
	if (isObj(v)) {
		if (typeof v.item === 'string') return [v.item];
		if (typeof v.tag === 'string') return ['#' + v.tag];
		if (typeof v.id === 'string') return [v.id];
	}
	return null;
}

function result(v: Json | undefined): RecipeInfo['result'] {
	if (typeof v === 'string') return { id: v, count: 1 };
	if (!isObj(v)) return null;
	const id = v.id ?? v.item;
	if (typeof id !== 'string') return null;
	return { id, count: typeof v.count === 'number' ? v.count : 1 };
}

const cooking: Record<string, number> = { smelting: 200, blasting: 100, smoking: 100, campfire_cooking: 600 };

/** Lays out a recipe for drawing. */
export function analyseRecipe(r: Json): RecipeInfo | null {
	if (!isObj(r)) return null;
	const type = strip(String(r.type ?? ''));
	const base = {
		type,
		result: result(r.result),
		group: typeof r.group === 'string' ? r.group : undefined,
		category: typeof r.category === 'string' ? r.category : undefined
	};
	if (type === 'crafting_shaped' && Array.isArray(r.pattern) && isObj(r.key)) {
		const key = r.key;
		const grid = r.pattern.slice(0, 3).map((row) =>
			[...String(row)].slice(0, 3).map((ch) => (ch === ' ' ? null : slot(key[ch])))
		);
		return { ...base, layout: 'grid', grid };
	}
	if (type === 'crafting_shapeless' && Array.isArray(r.ingredients)) {
		const items = r.ingredients.slice(0, 9).map((i) => slot(i));
		const grid: Slot[][] = [];
		for (let i = 0; i < items.length; i += 3) grid.push(items.slice(i, i + 3));
		return { ...base, layout: 'grid', grid, shapeless: true };
	}
	if (type === 'crafting_transmute') {
		// An empty result keeps the input item.
		const input = slot(r.input);
		const res = base.result ?? (input ? { id: input[0], count: 1 } : null);
		return { ...base, result: res, layout: 'grid', grid: [[input, slot(r.material)]], shapeless: true };
	}
	if (type === 'crafting_dye') {
		return { ...base, layout: 'grid', grid: [[slot(r.target), slot(r.dye)]], shapeless: true };
	}
	if (type === 'brewing') {
		const potion = (v: Json | undefined) => {
			const pc = isObj(v) ? (isObj(v.potion_contents) ? v.potion_contents : isObj(v.components) ? v.components['minecraft:potion_contents'] : undefined) : undefined;
			const p = isObj(pc) ? (pc.potions ?? pc.potion) : typeof pc === 'string' ? pc : undefined;
			return typeof p === 'string' ? strip(p).replace(/_/g, ' ') : Array.isArray(p) ? p.map((x) => strip(String(x)).replace(/_/g, ' ')).join(', ') : undefined;
		};
		return {
			...base,
			result: result(r.output),
			layout: 'brewing',
			brewing: { input: slot(r.input), reagent: slot(r.reagent) },
			notes: { input: potion(r.input), result: potion(r.output) }
		};
	}
	if (type === 'crafting_imbue') {
		const m = slot(r.material);
		return { ...base, layout: 'grid', grid: [[m, m, m], [m, slot(r.source), m], [m, m, m]] };
	}
	if (type in cooking) {
		return {
			...base,
			layout: 'cooking',
			input: slot(r.ingredient),
			cookingTime: typeof r.cookingtime === 'number' ? r.cookingtime : cooking[type],
			experience: typeof r.experience === 'number' ? r.experience : 0
		};
	}
	if (type === 'stonecutting') return { ...base, layout: 'stonecutting', input: slot(r.ingredient) };
	if (type === 'smithing_transform' || type === 'smithing_trim') {
		const smithing = { template: slot(r.template), base: slot(r.base), addition: slot(r.addition) };
		if (type === 'smithing_trim') {
			// The result is the base item with the trim applied.
			const pattern = typeof r.pattern === 'string' ? strip(r.pattern).replace(/_/g, ' ') + ' trim' : 'trimmed';
			return { ...base, result: smithing.base ? { id: smithing.base[0], count: 1 } : null, layout: 'smithing', smithing, notes: { result: pattern } };
		}
		return { ...base, layout: 'smithing', smithing };
	}
	return { ...base, layout: 'special' };
}
