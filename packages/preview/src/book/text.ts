// Text components as a book page shows them: flattened into styled
// pieces, then wrapped into lines the way the book screen wraps them.

import type { Font, Glyph } from './font';

export type Json = null | boolean | number | string | Json[] | { [k: string]: Json };
type Obj = { [k: string]: Json };

export interface Style {
	/** #rrggbb */
	color: string;
	bold: boolean;
	italic: boolean;
	underlined: boolean;
	strikethrough: boolean;
	obfuscated: boolean;
	font: string;
	click?: Obj;
	hover?: Json;
}

export const black: Style = {
	color: '#000000',
	bold: false,
	italic: false,
	underlined: false,
	strikethrough: false,
	obfuscated: false,
	font: 'minecraft:default'
};

/** A piece of a page: text in one style, or an inline sprite or head. */
export type Piece = { style: Style; text: string } | { style: Style; object: 'atlas' | 'player'; id: string; atlas?: string };

/** A value the book fills in when first opened, which the preview asks for. */
export interface Sample {
	key: string;
	label: string;
	kind: 'score' | 'nbt' | 'selector';
}

export interface Resolver {
	lang: Record<string, string>;
	/** The sample value for a score, nbt or selector. */
	sample: (s: Sample) => string;
	/** Every sample a page used, collected while flattening. */
	used: Map<string, Sample>;
}

export const textColors: Record<string, string> = {
	black: '#000000',
	dark_blue: '#0000AA',
	dark_green: '#00AA00',
	dark_aqua: '#00AAAA',
	dark_red: '#AA0000',
	dark_purple: '#AA00AA',
	gold: '#FFAA00',
	gray: '#AAAAAA',
	dark_gray: '#555555',
	blue: '#5555FF',
	green: '#55FF55',
	aqua: '#55FFFF',
	red: '#FF5555',
	light_purple: '#FF55FF',
	yellow: '#FFFF55',
	white: '#FFFFFF'
};

/** The keys each keybind is on by default, as the key's translation. */
const defaultKeys: Record<string, string> = {
	'key.attack': 'key.mouse.left',
	'key.use': 'key.mouse.right',
	'key.pickItem': 'key.mouse.middle',
	'key.forward': 'W',
	'key.left': 'A',
	'key.back': 'S',
	'key.right': 'D',
	'key.jump': 'key.keyboard.space',
	'key.sneak': 'key.keyboard.left.shift',
	'key.sprint': 'key.keyboard.left.control',
	'key.drop': 'Q',
	'key.inventory': 'E',
	'key.chat': 'T',
	'key.command': '/',
	'key.playerlist': 'key.keyboard.tab',
	'key.swapOffhand': 'F',
	'key.socialInteractions': 'P',
	'key.advancements': 'L',
	'key.screenshot': 'key.keyboard.f2',
	'key.togglePerspective': 'key.keyboard.f5',
	'key.fullscreen': 'key.keyboard.f11',
	'key.smoothCamera': 'key.keyboard.unknown',
	'key.spectatorOutlines': 'key.keyboard.unknown',
	'key.saveToolbarActivator': 'C',
	'key.loadToolbarActivator': 'X'
};

export function keyLabel(keybind: string, lang: Record<string, string>): string {
	const hotbar = /^key\.hotbar\.(\d)$/.exec(keybind);
	const key = hotbar ? hotbar[1] : defaultKeys[keybind];
	if (!key) return keybind;
	return lang[key] ?? key;
}

function styleOf(parent: Style, o: Obj): Style {
	const s = { ...parent };
	if (typeof o.color === 'string') s.color = o.color.startsWith('#') ? o.color : (textColors[o.color] ?? s.color);
	for (const k of ['bold', 'italic', 'underlined', 'strikethrough', 'obfuscated'] as const) if (typeof o[k] === 'boolean') s[k] = o[k] as boolean;
	if (typeof o.font === 'string') s.font = o.font.includes(':') ? o.font : 'minecraft:' + o.font;
	if (o.click_event && typeof o.click_event === 'object' && !Array.isArray(o.click_event)) s.click = o.click_event as Obj;
	if (o.hover_event && typeof o.hover_event === 'object' && !Array.isArray(o.hover_event)) {
		const h = o.hover_event as Obj;
		if (h.action === 'show_text') s.hover = h.value ?? '';
	}
	return s;
}

/** Flattens a text component into pieces, filling in translations and samples. */
export function flatten(c: Json, parent: Style, r: Resolver, out: Piece[] = [], depth = 0): Piece[] {
	if (depth > 64 || c === null) return out;
	if (typeof c !== 'object') {
		out.push({ style: parent, text: String(c) });
		return out;
	}
	if (Array.isArray(c)) {
		if (!c.length) return out;
		// [a, b, c] is a with b and c appended: they take a's style.
		const head = c[0];
		const s = head && typeof head === 'object' && !Array.isArray(head) ? styleOf(parent, head) : parent;
		flatten(head, parent, r, out, depth + 1);
		for (const e of c.slice(1)) flatten(e, s, r, out, depth + 1);
		return out;
	}
	const s = styleOf(parent, c);
	const type = typeof c.type === 'string' ? c.type : undefined;
	if (type === 'text' || (!type && 'text' in c)) {
		out.push({ style: s, text: String(c.text ?? '') });
	} else if (type === 'translatable' || (!type && 'translate' in c)) {
		translate(c, s, r, out, depth);
	} else if (type === 'score' || (!type && 'score' in c)) {
		const sc = (c.score ?? {}) as Obj;
		const name = String(sc.name ?? ''),
			obj = String(sc.objective ?? '');
		out.push({ style: s, text: r.sample(use(r, { key: `score ${name} ${obj}`, label: `${obj}[${name}]`, kind: 'score' })) });
	} else if (type === 'selector' || (!type && 'selector' in c)) {
		const sel = String(c.selector ?? '');
		out.push({ style: s, text: r.sample(use(r, { key: `selector ${sel}`, label: sel, kind: 'selector' })) });
	} else if (type === 'keybind' || (!type && 'keybind' in c)) {
		out.push({ style: s, text: keyLabel(String(c.keybind ?? ''), r.lang) });
	} else if (type === 'nbt' || (!type && 'nbt' in c)) {
		const where = String(c.storage ?? c.entity ?? c.block ?? '');
		out.push({ style: s, text: r.sample(use(r, { key: `nbt ${where} ${c.nbt}`, label: `${where} ${c.nbt}`, kind: 'nbt' })) });
	} else if (type === 'object' || (!type && ('sprite' in c || 'player' in c))) {
		if (c.object === 'player' || 'player' in c) {
			const p = c.player;
			const name = typeof p === 'string' ? p : p && typeof p === 'object' && !Array.isArray(p) ? String(p.name ?? '') : '';
			out.push({ style: s, object: 'player', id: name });
		} else {
			out.push({ style: s, object: 'atlas', id: String(c.sprite ?? ''), atlas: typeof c.atlas === 'string' ? c.atlas : undefined });
		}
	}
	if (Array.isArray(c.extra)) for (const e of c.extra) flatten(e, s, r, out, depth + 1);
	return out;
}

function use(r: Resolver, s: Sample): Sample {
	r.used.set(s.key, s);
	return s;
}

function translate(c: Obj, s: Style, r: Resolver, out: Piece[], depth: number) {
	const key = String(c.translate ?? '');
	const fmt = r.lang[key] ?? (typeof c.fallback === 'string' ? c.fallback : key);
	const args = Array.isArray(c.with) ? c.with : [];
	let next = 0;
	const re = /%(?:(\d+)\$)?([s%])/g;
	let at = 0;
	for (let m = re.exec(fmt); m; m = re.exec(fmt)) {
		if (m.index > at) out.push({ style: s, text: fmt.slice(at, m.index) });
		at = m.index + m[0].length;
		if (m[2] === '%') {
			out.push({ style: s, text: '%' });
			continue;
		}
		const i = m[1] ? parseInt(m[1]) - 1 : next++;
		if (i < args.length) flatten(args[i], s, r, out, depth + 1);
	}
	if (at < fmt.length) out.push({ style: s, text: fmt.slice(at) });
}

/** A laid-out glyph: where it is and which piece it came from. */
export interface Placed {
	x: number;
	line: number;
	/** The advance, bold included. */
	w: number;
	piece: Piece;
	/** The character, or -1 for a sprite or head. */
	cp: number;
	glyph?: Glyph;
}

/** How far bold draws the glyph again to the right, and widens it. */
export const boldOffset = (g: Glyph) => (g.size < 1 ? 0.5 : 1);

export const pageWidth = 114;
export const pageLines = 14;

/**
 * Wraps a page's pieces into lines pageWidth wide the way the game does:
 * at the last space that fits, else mid-word; a space that ends a line
 * is not drawn. Returns every glyph with its line, including those past
 * the bottom of the page, which the game cuts off.
 */
export function layout(pieces: Piece[], font: (id: string) => Font): { glyphs: Placed[]; lines: number } {
	const items: Placed[] = [];
	for (const p of pieces) {
		if ('object' in p) {
			items.push({ x: 0, line: 0, w: 9, piece: p, cp: -1 });
			continue;
		}
		const f = font(p.style.font);
		for (const ch of p.text) {
			const cp = ch.codePointAt(0)!;
			if (cp === 10) {
				items.push({ x: 0, line: 0, w: 0, piece: p, cp });
				continue;
			}
			const g = f.glyph(cp);
			items.push({ x: 0, line: 0, w: g.advance + (p.style.bold ? boldOffset(g) : 0), piece: p, cp, glyph: g });
		}
	}
	const out: Placed[] = [];
	let line = 0,
		x = 0,
		start = 0,
		lastSpace = -1;
	for (let i = 0; i < items.length; i++) {
		const g = items[i];
		if (g.cp === 10) {
			line++;
			x = 0;
			start = out.length;
			lastSpace = -1;
			continue;
		}
		if (x + g.w <= pageWidth) {
			g.x = x;
			g.line = line;
			x += g.w;
			if (g.cp === 32) lastSpace = out.length;
			out.push(g);
			continue;
		}
		if (g.cp === 32) {
			// The space ends the line and is not drawn.
			line++;
			x = 0;
			start = out.length;
			lastSpace = -1;
			continue;
		}
		line++;
		x = 0;
		if (lastSpace >= start) {
			// Move the last word down, dropping the space before it.
			const word = out.splice(lastSpace + 1);
			out.pop();
			start = out.length;
			for (const w of word) {
				w.x = x;
				w.line = line;
				x += w.w;
				out.push(w);
			}
		} else {
			start = out.length;
		}
		lastSpace = -1;
		g.x = x;
		g.line = line;
		x += g.w;
		out.push(g);
	}
	return { glyphs: out, lines: line + 1 };
}
