// Draws the book screen: the open book texture, the page indicator, the
// page's lines, the page turn arrows and a hover tooltip, in GUI pixels
// scaled up without smoothing, as the game draws them.

import { assetUrl, loadImage, vanillaAssets } from './assets';
import { Font, type Glyph, type PackJson } from './font';
import { black, boldOffset, flatten, layout, pageLines, type Json, type Piece, type Placed, type Resolver, type Style } from './text';

/** The canvas in GUI pixels: the book with a little margin; tooltips stay inside it. */
export const screenW = 196;
export const screenH = 196;
/** The book's top left corner on the canvas. */
export const bookX = (screenW - 192) >> 1;
export const bookY = 2;

const textX = bookX + 36;
const textY = bookY + 30;
export const forwardButton = { x: bookX + 116, y: bookY + 157, w: 23, h: 13 };
export const backButton = { x: bookX + 43, y: bookY + 157, w: 23, h: 13 };

type Image = ImageBitmap | OffscreenCanvas;

export class BookRenderer {
	private fonts = new Map<string, Font>();
	private images = new Map<string, Image | null>();
	notes = new Set<string>();

	constructor(
		private pack: PackJson,
		public lang: Record<string, string>
	) {}

	font = (id: string): Font => this.fonts.get(id) ?? this.fonts.get('minecraft:default')!;

	/** Flattens and wraps a page, loading the fonts and images it needs. */
	async prepare(page: Json, r: Resolver, tooltip?: boolean): Promise<{ pieces: Piece[]; glyphs: Placed[] }> {
		await this.loadFont('minecraft:default');
		const pieces = flatten(page, tooltip ? { ...black, color: '#FFFFFF' } : black, r);
		await this.load(pieces);
		return { pieces, glyphs: layout(pieces, this.font).glyphs };
	}

	async load(pieces: Piece[]) {
		const chars = new Map<string, Set<number>>();
		for (const p of pieces) {
			if ('object' in p) {
				await this.loadImage(spriteKey(p));
				continue;
			}
			await this.loadFont(p.style.font);
			const set = chars.get(p.style.font) ?? new Set();
			for (const ch of p.text) set.add(ch.codePointAt(0)!);
			chars.set(p.style.font, set);
		}
		for (const [f, set] of chars) await this.font(f).prepare(set);
		for (const f of this.fonts.values()) for (const n of f.notes) this.notes.add(n);
	}

	private async loadFont(id: string) {
		if (this.fonts.has(id)) return;
		this.fonts.set(id, await Font.load(id, this.pack));
	}

	async loadChrome() {
		await this.loadFont('minecraft:default');
		for (const k of ['gui/book', 'widget/page_forward', 'widget/page_forward_highlighted', 'widget/page_backward', 'widget/page_backward_highlighted'])
			await this.loadImage(k.startsWith('widget/') ? `sprite|minecraft:gui|minecraft:${k}` : `texture|minecraft:${k}`);
	}

	/** An image by key: texture|id, sprite|atlas|id or head|name. */
	private async loadImage(key: string) {
		if (this.images.has(key)) return;
		const [kind, a, b] = key.split('|');
		let img: ImageData | null = null;
		try {
			if (kind === 'texture') img = await loadImage(assetUrl(a, 'textures', '.png'));
			else if (kind === 'sprite') img = await loadImage(assetUrl(b, a === 'minecraft:gui' ? 'textures/gui/sprites' : 'textures', '.png'));
			else img = await loadImage(`${vanillaAssets}/minecraft/textures/entity/player/wide/steve.png`);
		} catch {
			img = null;
		}
		if (!img) {
			if (kind === 'sprite') this.notes.add(`There is no vanilla texture ${b}; the game shows it if the resource pack has it.`);
			this.images.set(key, null);
			return;
		}
		if (kind === 'head') {
			// The face, then the hat layer over it.
			const c = new OffscreenCanvas(8, 8);
			const ctx = c.getContext('2d')!;
			const src = await createImageBitmap(img);
			ctx.drawImage(src, 8, 8, 8, 8, 0, 0, 8, 8);
			ctx.drawImage(src, 40, 8, 8, 8, 0, 0, 8, 8);
			this.images.set(key, c);
			return;
		}
		// An animated texture is a strip of square frames: show the first.
		const h = kind === 'sprite' && img.height > img.width ? img.width : img.height;
		this.images.set(key, await createImageBitmap(img, 0, 0, img.width, h));
	}

	/**
	 * Draws the screen. hoverAt is the pointer in GUI pixels; tooltip is
	 * the laid-out hover text to show there.
	 */
	draw(
		ctx: CanvasRenderingContext2D,
		scale: number,
		page: { glyphs: Placed[]; index: number; count: number },
		hoverAt: { x: number; y: number } | null,
		tooltip: Placed[] | null
	) {
		ctx.setTransform(1, 0, 0, 1, 0, 0);
		ctx.clearRect(0, 0, ctx.canvas.width, ctx.canvas.height);
		ctx.imageSmoothingEnabled = false;
		ctx.setTransform(scale, 0, 0, scale, 0, 0);
		const book = this.images.get('texture|minecraft:gui/book');
		if (book) ctx.drawImage(book, 0, 0, 192, 192, bookX, bookY, 192, 192);

		const indicator = flatten(fmt(this.lang['book.pageIndicator'] ?? 'Page %1$s of %2$s', page.index + 1, page.count), black, this.resolver());
		const iw = layout(indicator, this.font).glyphs;
		const width = iw.reduce((n, g) => n + g.w, 0);
		this.text(ctx, iw, bookX + 192 - 44 - width, bookY + 16, false);

		this.text(
			ctx,
			page.glyphs.filter((g) => g.line < pageLines),
			textX,
			textY,
			false
		);

		const over = (b: typeof forwardButton) => hoverAt && hoverAt.x >= b.x && hoverAt.x < b.x + b.w && hoverAt.y >= b.y && hoverAt.y < b.y + b.h;
		const button = (b: typeof forwardButton, name: string) => {
			const img = this.images.get(`sprite|minecraft:gui|minecraft:widget/${name}${over(b) ? '_highlighted' : ''}`);
			if (img) ctx.drawImage(img, b.x, b.y);
		};
		if (page.index < page.count - 1) button(forwardButton, 'page_forward');
		if (page.index > 0) button(backButton, 'page_backward');

		if (tooltip && hoverAt) this.tooltip(ctx, tooltip, hoverAt);
	}

	private resolver(): Resolver {
		return { lang: this.lang, sample: () => '', used: new Map() };
	}

	/** Draws laid-out glyphs with their line 0 at (x0, y0), 9 pixels a line. */
	private text(ctx: CanvasRenderingContext2D, glyphs: Placed[], x0: number, y0: number, shadow: boolean) {
		if (shadow) for (const g of glyphs) this.glyph(ctx, g, x0 + 1, y0 + 1, true);
		for (const g of glyphs) this.glyph(ctx, g, x0, y0, false);
	}

	private glyph(ctx: CanvasRenderingContext2D, g: Placed, x0: number, y0: number, shadow: boolean) {
		const s = g.piece.style;
		const x = x0 + g.x,
			y = y0 + g.line * 9;
		const color = shadow ? shadowOf(s.color) : s.color;
		if ('object' in g.piece) {
			if (shadow) return;
			const img = this.images.get(spriteKey(g.piece));
			if (img) ctx.drawImage(img, 0, 0, img.width, img.height, x, y, 8, 8);
			else missingTexture(ctx, x, y);
			return;
		}
		let gl = g.glyph;
		if (!gl) return;
		if (s.obfuscated && gl.pixels.length) gl = this.font(s.font).scramble(gl);
		ctx.fillStyle = color;
		pixels(ctx, gl, x, y, s.italic);
		if (s.bold) pixels(ctx, gl, x + boldOffset(gl), y, s.italic);
		if (s.underlined) ctx.fillRect(x - 1, y + 8, g.w + 1, 1);
		if (s.strikethrough) ctx.fillRect(x - 1, y + 3.5, g.w + 1, 1);
	}

	private tooltip(ctx: CanvasRenderingContext2D, glyphs: Placed[], at: { x: number; y: number }) {
		const lines = glyphs.reduce((n, g) => Math.max(n, g.line + 1), 1);
		const w = glyphs.reduce((n, g) => Math.max(n, g.x + g.w), 0);
		const h = lines * 9 - 1;
		let x = at.x + 12,
			y = at.y - 12;
		if (x + w + 4 > screenW) x = Math.max(4, at.x - 16 - w);
		y = Math.max(4, Math.min(y, screenH - h - 4));
		// The classic tooltip: a near-black box with a purple frame.
		ctx.fillStyle = 'rgba(16,0,16,0.94)';
		ctx.fillRect(x - 3, y - 4, w + 6, 1);
		ctx.fillRect(x - 3, y + h + 3, w + 6, 1);
		ctx.fillRect(x - 3, y - 3, w + 6, h + 6);
		ctx.fillRect(x - 4, y - 3, 1, h + 6);
		ctx.fillRect(x + w + 3, y - 3, 1, h + 6);
		const grad = ctx.createLinearGradient(0, y - 3, 0, y + h + 3);
		grad.addColorStop(0, 'rgba(80,0,255,0.31)');
		grad.addColorStop(1, 'rgba(40,0,127,0.31)');
		ctx.fillStyle = grad;
		ctx.fillRect(x - 3, y - 3, 1, h + 6);
		ctx.fillRect(x + w + 2, y - 3, 1, h + 6);
		ctx.fillRect(x - 3, y - 3, w + 6, 1);
		ctx.fillRect(x - 3, y + h + 2, w + 6, 1);
		this.text(ctx, glyphs, x, y, true);
	}
}

function pixels(ctx: CanvasRenderingContext2D, g: Glyph, x: number, y: number, italic: boolean) {
	const p = g.pixels,
		s = g.size;
	for (let i = 0; i < p.length; i += 2) {
		// Italic leans the glyph: its top one pixel right, its bottom one left.
		const dx = italic ? 1 - 0.25 * (p[i + 1] + s / 2) : 0;
		ctx.fillRect(x + p[i] + dx, y + p[i + 1], s, s);
	}
}

/** The game's drop shadow: the colour at a quarter brightness. */
function shadowOf(color: string): string {
	const n = parseInt(color.slice(1), 16) & 0xfcfcfc;
	return '#' + (n >> 2).toString(16).padStart(6, '0');
}

function missingTexture(ctx: CanvasRenderingContext2D, x: number, y: number) {
	ctx.fillStyle = '#000000';
	ctx.fillRect(x, y, 8, 8);
	ctx.fillStyle = '#f800f8';
	ctx.fillRect(x, y, 4, 4);
	ctx.fillRect(x + 4, y + 4, 4, 4);
}

export function spriteKey(p: Extract<Piece, { object: string }>): string {
	if (p.object === 'player') return `head|${p.id}`;
	const id = p.id.includes(':') ? p.id : 'minecraft:' + p.id;
	const atlas = p.atlas ? (p.atlas.includes(':') ? p.atlas : 'minecraft:' + p.atlas) : 'minecraft:blocks';
	return `sprite|${atlas}|${id}`;
}

function fmt(f: string, ...args: (string | number)[]): string {
	let next = 0;
	return f.replace(/%(?:(\d+)\$)?s/g, (_, n) => String(args[n ? parseInt(n) - 1 : next++] ?? ''));
}

/** The glyph under a point on the page, if any. */
export function glyphAt(glyphs: Placed[], at: { x: number; y: number }): Placed | null {
	const line = Math.floor((at.y - textY) / 9);
	if (at.y < textY || line >= pageLines) return null;
	const x = at.x - textX;
	return glyphs.find((g) => g.line === line && x >= g.x && x < g.x + g.w) ?? null;
}

export type { Style };
