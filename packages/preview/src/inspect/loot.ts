import { isObj, numberRange, strip, type Json, type Obj } from './common';

/** One leaf entry of a pool, with its chance per roll. */
export interface LootRow {
	/** An item id, `#tag`, loot table id, or a description such as "nothing". */
	name: string;
	kind: 'item' | 'tag' | 'loot_table' | 'empty' | 'dynamic' | 'other';
	weight: number;
	/** Chance that one roll picks this entry, if every condition passes. */
	chance: number;
	count: { min: number; max: number; mean: number } | null;
	/** Expected number of items per table roll. */
	expected: number | null;
	conditional: boolean;
	functions: string[];
	/** Set when the entry sits in an alternatives, group or sequence. */
	via?: string;
}

export interface LootPoolInfo {
	rolls: { min: number; max: number; mean: number } | null;
	bonus: boolean;
	conditional: boolean;
	functions: string[];
	rows: LootRow[];
}

const conds = (o: Obj) => {
	const c = o.conditions ?? o.condition;
	return Array.isArray(c) ? c.length > 0 : c !== undefined;
};

const fnNames = (o: Obj): string[] => {
	const fs = o.functions ?? o.modifier ?? o.modifiers;
	const list = Array.isArray(fs) ? fs : fs !== undefined ? [fs] : [];
	return list.map((f) => (isObj(f) ? strip(String(f.function ?? f.type ?? '?')) : strip(String(f))));
};

/** The item count an entry's set_count function gives, if any. */
function countOf(o: Obj): LootRow['count'] {
	const fs = o.functions ?? o.modifier ?? o.modifiers;
	const list = Array.isArray(fs) ? fs : fs !== undefined ? [fs] : [];
	let count: LootRow['count'] = { min: 1, max: 1, mean: 1 };
	for (const f of list) {
		if (!isObj(f) || strip(String(f.function ?? f.type)) !== 'set_count') continue;
		const r = numberRange(f.count);
		count = r && count && f.add === true ? { min: count.min + r.min, max: count.max + r.max, mean: count.mean + r.mean } : r;
	}
	return count;
}

function leaves(entries: Json[], via?: string): { o: Obj; via?: string }[] {
	const out: { o: Obj; via?: string }[] = [];
	for (const e of entries) {
		if (!isObj(e)) continue;
		const type = strip(String(e.type ?? ''));
		if ((type === 'alternatives' || type === 'group' || type === 'sequence') && Array.isArray(e.children)) {
			out.push(...leaves(e.children, via ?? type));
		} else {
			out.push({ o: e, via });
		}
	}
	return out;
}

/** Summarises a loot table: per pool, each entry's chance and count. */
export function analyseLoot(table: Json): LootPoolInfo[] {
	if (!isObj(table) || !Array.isArray(table.pools)) return [];
	return table.pools.filter(isObj).map((pool) => {
		const rolls = numberRange(pool.rolls);
		const entries = Array.isArray(pool.entries) ? leaves(pool.entries) : [];
		const total = entries.reduce((s, { o }) => s + (typeof o.weight === 'number' ? o.weight : 1), 0) || 1;
		const rows: LootRow[] = entries.map(({ o, via }) => {
			const type = strip(String(o.type ?? ''));
			const weight = typeof o.weight === 'number' ? o.weight : 1;
			const chance = weight / total;
			let name = String(o.name ?? o.value ?? '');
			let kind: LootRow['kind'] = 'other';
			if (type === 'item') kind = 'item';
			else if (type === 'tag') {
				kind = 'tag';
				name = String(o.items ?? o.name ?? '');
				if (!name.startsWith('#')) name = '#' + name;
			} else if (type === 'loot_table') {
				kind = 'loot_table';
				name = typeof o.value === 'string' ? o.value : 'inline table';
			} else if (type === 'empty') {
				kind = 'empty';
				name = 'nothing';
			} else if (type === 'dynamic') {
				kind = 'dynamic';
				name = `dynamic: ${strip(name)}`;
			} else name = type || 'entry';
			const count = kind === 'empty' ? { min: 0, max: 0, mean: 0 } : countOf(o);
			const expected = rolls && count ? rolls.mean * chance * count.mean : null;
			return { name, kind, weight, chance, count, expected, conditional: conds(o), functions: fnNames(o), via };
		});
		return { rolls, bonus: pool.bonus_rolls !== undefined && pool.bonus_rolls !== 0, conditional: conds(pool), functions: fnNames(pool), rows };
	});
}

/** One simulated roll of the table, ignoring conditions. */
export function rollLoot(pools: LootPoolInfo[], random: () => number): { name: string; count: number }[] {
	const out = new Map<string, number>();
	const int = (r: { min: number; max: number }) => r.min + Math.floor(random() * (r.max - r.min + 1));
	for (const pool of pools) {
		if (!pool.rows.length || !pool.rolls) continue;
		const n = int(pool.rolls);
		const total = pool.rows.reduce((s, r) => s + r.weight, 0);
		for (let i = 0; i < n; i++) {
			let pick = random() * total;
			const row = pool.rows.find((r) => (pick -= r.weight) < 0) ?? pool.rows[pool.rows.length - 1];
			if (row.kind === 'empty') continue;
			const c = row.count ? int(row.count) : 1;
			if (c > 0) out.set(row.name, (out.get(row.name) ?? 0) + c);
		}
	}
	return [...out].map(([name, count]) => ({ name, count }));
}
