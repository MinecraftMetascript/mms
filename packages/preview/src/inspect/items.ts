// How the inspectors show an item stack: a CustomItem's result carries
// its base id plus components, so its name comes from the item_name or
// custom_name component (looked up in the pack's lang files) and its
// icon from the item_model it names, followed through the pack's item
// model definition and model to a texture the pack copies.
import type { OutputFile, OutputItem } from '../types';
import { isObj, label, type Json, type Obj } from './common';

/** What the inspectors know about the pack's items. */
export interface ItemContext {
	/** The pack's en_us lines. */
	lang: Record<string, string>;
	/** Parsed output files, by path. */
	json: (path: string) => Json | undefined;
	/** Image URLs of the pack's copied textures, by their path in the pack. */
	textures: Record<string, string>;
	/** The pack's custom items, with their stacks parsed. */
	items: { id: string; stack: Obj }[];
}

/** The Svelte context key the inspector shares its ItemContext under. */
export const itemsKey = Symbol('mms-items');

export const emptyItems: ItemContext = { lang: {}, json: () => undefined, textures: {}, items: [] };

export function itemContext(output: OutputFile[], items: OutputItem[], textures: Record<string, string>): ItemContext {
	const byPath = new Map(output.map((f) => [f.path, f.text]));
	const parsed = new Map<string, Json | undefined>();
	const json = (path: string) => {
		if (!parsed.has(path)) {
			const text = byPath.get(path);
			let v: Json | undefined;
			try {
				v = text === undefined ? undefined : JSON.parse(text);
			} catch {
				v = undefined;
			}
			parsed.set(path, v);
		}
		return parsed.get(path);
	};
	const lang: Record<string, string> = {};
	for (const f of output) {
		if (!/^assets\/[^/]+\/lang\/en_us\.json$/.test(f.path)) continue;
		const l = json(f.path);
		if (isObj(l)) for (const [k, v] of Object.entries(l)) if (typeof v === 'string') lang[k] = v;
	}
	const stacks: ItemContext['items'] = [];
	for (const it of items) {
		try {
			const s = JSON.parse(it.text);
			// Books show as their pages; they are not crafted items here.
			if (isObj(s) && typeof s.id === 'string') stacks.push({ id: it.id, stack: s });
		} catch {
			// A broken stack shows by its base item.
		}
	}
	return { lang, json, textures, items: stacks };
}

/** How one item (or #tag) shows in a slot. */
export interface ItemLook {
	/** The id or #tag as written. */
	id: string;
	/** The name shown on the tile. */
	name: string;
	/** An image URL for the tile, when the pack has the texture. */
	icon?: string;
	/** The CustomItem this stack is, by its declaration id. */
	custom?: string;
	/** True when the name comes from the stack's components. */
	named?: boolean;
	/** The tile's tooltip. */
	title: string;
}

const comp = (c: Obj | undefined, k: string): Json | undefined => c?.['minecraft:' + k] ?? c?.[k];

/** A text component as plain text, with translations looked up. */
export function plainText(v: Json | undefined, lang: Record<string, string>): string {
	if (typeof v === 'string') return v;
	if (typeof v === 'number' || typeof v === 'boolean') return String(v);
	if (Array.isArray(v)) return v.map((x) => plainText(x, lang)).join('');
	if (!isObj(v)) return '';
	let s = '';
	if (typeof v.text === 'string') s = v.text;
	else if (typeof v.translate === 'string') {
		const fmt = lang[v.translate] ?? (typeof v.fallback === 'string' ? v.fallback : v.translate);
		const args = Array.isArray(v.with) ? v.with.map((x) => plainText(x, lang)) : [];
		let i = 0;
		s = fmt.replace(/%(?:(\d+)\$)?s/g, (_, n) => args[n ? Number(n) - 1 : i++] ?? '');
	}
	if (Array.isArray(v.extra)) s += v.extra.map((x) => plainText(x, lang)).join('');
	return s;
}

/** Deep equality of two JSON values, ignoring key order and minecraft: prefixes on keys. */
function same(a: Json | undefined, b: Json | undefined): boolean {
	if (Array.isArray(a)) return Array.isArray(b) && a.length === b.length && a.every((x, i) => same(x, b[i]));
	if (isObj(a)) {
		if (!isObj(b)) return false;
		const norm = (o: Obj) => new Map(Object.entries(o).map(([k, v]) => [k.replace(/^minecraft:/, ''), v]));
		const ma = norm(a);
		const mb = norm(b);
		return ma.size === mb.size && [...ma].every(([k, v]) => mb.has(k) && same(v, mb.get(k)));
	}
	return a === b;
}

const full = (id: string) => (id.includes(':') ? id : 'minecraft:' + id);

/** The CustomItem a stack is: by its mms_item marker, else by equal components. */
function customOf(id: string, components: Obj | undefined, ctx: ItemContext): string | undefined {
	if (!components) return undefined;
	const data = comp(components, 'custom_data');
	if (isObj(data) && typeof data.mms_item === 'string' && ctx.items.some((it) => it.id === data.mms_item)) return data.mms_item;
	return ctx.items.find((it) => full(String(it.stack.id)) === full(id) && same(it.stack.components ?? {}, components))?.id;
}

/** A resource location's file in the pack: assets/ns/folder/path.ext. */
function assetPath(id: string, folder: string, ext: string): string {
	const [ns, path] = id.includes(':') ? id.split(':', 2) : ['minecraft', id];
	return `assets/${ns}/${folder}/${path}${ext}`;
}

/** The model an item model definition draws first: a select's fallback or first case, and so on. */
function firstModel(m: Json | undefined, depth = 0): string | undefined {
	if (!isObj(m) || depth > 8) return undefined;
	if (typeof m.model === 'string') return m.model;
	for (const k of ['model', 'fallback', 'on_true', 'on_false']) {
		const r = firstModel(m[k], depth + 1);
		if (r) return r;
	}
	for (const k of ['cases', 'entries', 'models']) {
		const list = m[k];
		if (!Array.isArray(list)) continue;
		for (const c of list) {
			const r = firstModel(isObj(c) && 'model' in c && isObj(c.model) ? c.model : c, depth + 1);
			if (r) return r;
		}
	}
	return undefined;
}

/** The icon for an item_model: its definition's model, that model's layer0 texture (following parents), as a pack texture. */
function iconOf(itemModel: string, ctx: ItemContext): string | undefined {
	const def = ctx.json(assetPath(itemModel, 'items', '.json'));
	let model = isObj(def) ? firstModel(def.model) : undefined;
	for (let depth = 0; model && depth < 8; depth++) {
		const m = ctx.json(assetPath(model, 'models', '.json'));
		if (!isObj(m)) return undefined;
		const tex = isObj(m.textures) ? m.textures : undefined;
		const t = tex?.layer0 ?? tex?.particle ?? (tex ? Object.values(tex)[0] : undefined);
		if (typeof t === 'string' && !t.startsWith('#')) return ctx.textures[assetPath(t, 'textures', '.png')];
		model = typeof m.parent === 'string' ? m.parent : undefined;
	}
	return undefined;
}

/** The custom items whose base is this id, by name: an ingredient of that id matches them too. */
export function customsOfBase(id: string, ctx: ItemContext): string[] {
	if (id.startsWith('#')) return [];
	return ctx.items
		.filter((it) => full(String(it.stack.id)) === full(id))
		.map((it) => look(String(it.stack.id), isObj(it.stack.components) ? it.stack.components : undefined, ctx).name);
}

/** How an item id (or #tag), with the stack's components if any, shows. */
export function look(id: string, components: Obj | undefined, ctx: ItemContext): ItemLook {
	const base = label(id);
	if (id.startsWith('#')) return { id, name: base, title: `${base} (${id})` };
	const custom = customOf(id, components, ctx);
	const nameComp = comp(components, 'custom_name') ?? comp(components, 'item_name');
	const name = nameComp !== undefined ? plainText(nameComp, ctx.lang) : '';
	const model = comp(components, 'item_model');
	const icon = typeof model === 'string' ? iconOf(model, ctx) : undefined;
	const shown = name || base;
	const lines = [shown];
	if (custom) lines.push(`Custom item ${custom}`);
	if (name || custom) lines.push(`Base: ${base} (${full(id)})`);
	else lines[0] = `${base} (${full(id)})`;
	return { id, name: shown, icon, custom, named: !!name, title: lines.join('\n') };
}
