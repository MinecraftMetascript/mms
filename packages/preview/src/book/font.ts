// Minecraft's text fonts, drawn the way the game draws them: bitmap
// glyph sheets, fixed-width spaces and the unifont fallback, with the
// game's glyph widths, bold, italic and the scrambled (obfuscated) style.

import { assetUrl, loadImage, loadJson, loadUnifont } from './assets';

/** One character of a font: its advance and its lit pixels. */
export interface Glyph {
	/** How far the pen moves, in GUI pixels, before bold. */
	advance: number;
	/** Lit pixels as x, y pairs in GUI pixels from the top left of the line. */
	pixels: number[];
	/** The size of one glyph pixel in GUI pixels (1 for ascii, 0.5 for unifont). */
	size: number;
}

/** The game's box for a character no provider has. */
const missing: Glyph = (() => {
	const px: number[] = [];
	for (let y = 0; y < 8; y++)
		for (let x = 0; x < 5; x++) if (y === 0 || y === 7 || x === 0 || x === 4) px.push(x, y);
	return { advance: 6, pixels: px, size: 1 };
})();

type Provider =
	| { type: 'glyphs'; glyphs: Map<number, Glyph> }
	| { type: 'unihex'; file: string; overrides: { from: number; to: number; left: number; right: number }[] };

interface ProviderJson {
	type: string;
	id?: string;
	file?: string;
	hex_file?: string;
	height?: number;
	ascent?: number;
	chars?: string[];
	advances?: Record<string, number>;
	filter?: Record<string, boolean>;
	size_overrides?: { from: string; to: string; left: number; right: number }[];
}

/** Reads a font's JSON from the project's resource pack, if it has it. */
export type PackJson = (path: string) => unknown;

/** A loaded font: its providers in order, the first with a glyph wins. */
export class Font {
	private providers: Provider[] = [];
	private hex = new Map<string, Map<number, string>>();
	private unihex = new Map<number, Glyph>();
	private byWidth = new Map<number, Glyph[]>();
	/** Problems loading parts of the font, to show under the preview. */
	notes: string[] = [];

	static async load(id: string, pack: PackJson): Promise<Font> {
		const f = new Font();
		await f.add(id, pack, new Set());
		for (const p of f.providers) {
			if (p.type !== 'glyphs') continue;
			for (const g of p.glyphs.values()) {
				if (!g.pixels.length) continue;
				const list = f.byWidth.get(g.advance) ?? [];
				list.push(g);
				f.byWidth.set(g.advance, list);
			}
		}
		return f;
	}

	private async add(id: string, pack: PackJson, seen: Set<string>) {
		if (seen.has(id)) return;
		seen.add(id);
		const [ns, path] = id.includes(':') ? id.split(':', 2) : ['minecraft', id];
		const own = pack(`assets/${ns}/font/${path}.json`) as { providers?: ProviderJson[] } | undefined;
		const def = own ?? (await loadJson<{ providers: ProviderJson[] }>(assetUrl(id, 'font', '.json')));
		if (!def) {
			this.notes.push(`There is no font ${id}.`);
			return;
		}
		for (const p of def.providers ?? []) {
			// The default options: not forced uniform, not Japanese variants.
			if (p.filter && Object.values(p.filter).some((v) => v)) continue;
			switch (p.type) {
				case 'reference':
					if (p.id) await this.add(p.id, pack, seen);
					break;
				case 'space': {
					const glyphs = new Map<number, Glyph>();
					for (const [ch, adv] of Object.entries(p.advances ?? {})) glyphs.set(ch.codePointAt(0)!, { advance: adv, pixels: [], size: 1 });
					this.providers.push({ type: 'glyphs', glyphs });
					break;
				}
				case 'bitmap': {
					const glyphs = await bitmapGlyphs(p);
					if (glyphs) this.providers.push({ type: 'glyphs', glyphs });
					else this.notes.push(`The preview can't draw ${p.file}; its glyphs show as boxes.`);
					break;
				}
				case 'unihex':
					if (p.hex_file)
						this.providers.push({
							type: 'unihex',
							file: p.hex_file,
							overrides: (p.size_overrides ?? []).map((o) => ({ from: o.from.codePointAt(0)!, to: o.to.codePointAt(0)!, left: o.left, right: o.right }))
						});
					break;
				default:
					this.notes.push(`The preview can't draw ${p.type} fonts.`);
			}
		}
	}

	/** Loads what drawing these characters needs, such as unifont. */
	async prepare(chars: Iterable<number>) {
		for (const cp of chars) {
			if (cp === 10 || this.bitmap(cp)) continue;
			for (const p of this.providers) {
				if (p.type !== 'unihex' || this.hex.has(p.file)) continue;
				try {
					this.hex.set(p.file, await loadUnifont(p.file));
				} catch {
					this.hex.set(p.file, new Map());
					this.notes.push("Couldn't load unifont; some characters show as boxes.");
				}
			}
			return;
		}
	}

	private bitmap(cp: number): Glyph | undefined {
		for (const p of this.providers) {
			if (p.type === 'glyphs') {
				const g = p.glyphs.get(cp);
				if (g) return g;
			}
		}
	}

	glyph(cp: number): Glyph {
		for (const p of this.providers) {
			if (p.type === 'glyphs') {
				const g = p.glyphs.get(cp);
				if (g) return g;
				continue;
			}
			const cached = this.unihex.get(cp);
			if (cached) return cached;
			const bits = this.hex.get(p.file)?.get(cp);
			if (!bits) continue;
			const g = hexGlyph(bits, p.overrides.find((o) => o.from <= cp && cp <= o.to));
			this.unihex.set(cp, g);
			return g;
		}
		return missing;
	}

	/** A random glyph as wide as g, for scrambled text. */
	scramble(g: Glyph): Glyph {
		const list = this.byWidth.get(g.advance);
		return list ? list[Math.floor(Math.random() * list.length)] : g;
	}
}

async function bitmapGlyphs(p: ProviderJson): Promise<Map<number, Glyph> | null> {
	if (!p.file || !p.chars?.length) return null;
	let img: ImageData | null = null;
	try {
		// Only vanilla textures load: a project font's texture is a file
		// the preview can't read.
		img = await loadImage(assetUrl(p.file, 'textures', ''));
	} catch {
		return null;
	}
	if (!img) return null;
	const rows = p.chars.map((r) => Array.from(r));
	const cols = Math.max(...rows.map((r) => r.length));
	const cw = Math.floor(img.width / cols);
	const ch = Math.floor(img.height / rows.length);
	const height = p.height ?? 8;
	const ascent = p.ascent ?? 7;
	const scale = height / ch;
	const top = 7 - ascent;
	const glyphs = new Map<number, Glyph>();
	rows.forEach((row, ry) =>
		row.forEach((c, rx) => {
			const cp = c.codePointAt(0)!;
			if (cp === 0 || glyphs.has(cp)) return;
			const px: number[] = [];
			let width = 0;
			for (let y = 0; y < ch; y++)
				for (let x = 0; x < cw; x++) {
					if (img!.data[((ry * ch + y) * img!.width + rx * cw + x) * 4 + 3] === 0) continue;
					px.push(x * scale, top + y * scale);
					width = Math.max(width, x + 1);
				}
			glyphs.set(cp, { advance: Math.floor(0.5 + width * scale) + 1, pixels: px, size: scale });
		})
	);
	return glyphs;
}

function hexGlyph(bits: string, o?: { left: number; right: number }): Glyph {
	// 32, 48 or 64 hex digits: 16 rows of 8, 12 or 16 pixels.
	const w = (bits.length / 16) * 4;
	const lit: [number, number][] = [];
	for (let y = 0; y < 16; y++) {
		const row = parseInt(bits.slice((y * w) / 4, ((y + 1) * w) / 4), 16);
		for (let x = 0; x < w; x++) if (row & (1 << (w - 1 - x))) lit.push([x, y]);
	}
	let left = o?.left ?? Math.min(...lit.map(([x]) => x));
	let right = o?.right ?? Math.max(...lit.map(([x]) => x));
	if (!lit.length && !o) left = right = 0;
	const px: number[] = [];
	for (const [x, y] of lit) px.push((x - left) / 2, y / 2);
	return { advance: Math.floor((right - left + 1) / 2) + 1, pixels: px, size: 0.5 };
}
