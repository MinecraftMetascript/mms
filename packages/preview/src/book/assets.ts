// Vanilla client assets the book preview draws with: the book screen
// texture, the fonts, item and block textures, and the English language
// file. They come from mcmeta's copy of the 26.3 assets, fetched when
// first needed and kept for the session.

/** mcmeta's 26.3-assets tag, pinned so the preview never shifts under us. */
export const vanillaAssets = 'https://raw.githubusercontent.com/misode/mcmeta/79962f2b091fb6213b425b277a8725d737f3def8/assets';

/** A resource location's file: ns:path in a folder with an extension. */
export function assetUrl(id: string, folder: string, ext: string): string {
	const [ns, path] = id.includes(':') ? id.split(':', 2) : ['minecraft', id];
	return `${vanillaAssets}/${ns}/${folder}/${path}${ext}`;
}

const cache = new Map<string, Promise<unknown>>();

function once<T>(key: string, load: () => Promise<T>): Promise<T> {
	let p = cache.get(key) as Promise<T> | undefined;
	if (!p) {
		p = load();
		// A failed fetch can be tried again next time.
		p.catch(() => cache.delete(key));
		cache.set(key, p);
	}
	return p;
}

/** An image's pixels, or null when the game has no such file. */
export function loadImage(url: string): Promise<ImageData | null> {
	return once('img|' + url, async () => {
		const res = await fetch(url);
		if (res.status === 404) return null;
		if (!res.ok) throw new Error(`${url}: ${res.status}`);
		const bmp = await createImageBitmap(await res.blob());
		const c = new OffscreenCanvas(bmp.width, bmp.height);
		const ctx = c.getContext('2d')!;
		ctx.drawImage(bmp, 0, 0);
		return ctx.getImageData(0, 0, bmp.width, bmp.height);
	});
}

/** A JSON file, or null when there is none. */
export function loadJson<T>(url: string): Promise<T | null> {
	return once('json|' + url, async () => {
		const res = await fetch(url);
		if (res.status === 404) return null;
		if (!res.ok) throw new Error(`${url}: ${res.status}`);
		return (await res.json()) as T;
	});
}

/** The English translations. */
export const loadLang = () => loadJson<Record<string, string>>(`${vanillaAssets}/minecraft/lang/en_us.json`).then((l) => l ?? {});

/**
 * Unifont, the fallback for every character the bitmap fonts lack: a zip
 * holding one .hex file of code point:bitmap lines. It is 1.5 MB, so it
 * is only fetched when a page needs it.
 */
export function loadUnifont(id: string): Promise<Map<number, string>> {
	return once('hex|' + id, async () => {
		const res = await fetch(assetUrl(id, 'font', ''));
		if (!res.ok) throw new Error(`${id}: ${res.status}`);
		const text = await unzipFirst(new Uint8Array(await res.arrayBuffer()), '.hex');
		const glyphs = new Map<number, string>();
		for (const line of text.split('\n')) {
			const i = line.indexOf(':');
			if (i > 0) glyphs.set(parseInt(line.slice(0, i), 16), line.slice(i + 1).trim());
		}
		return glyphs;
	});
}

/** The text of the first file in a zip whose name ends with suffix. */
async function unzipFirst(zip: Uint8Array, suffix: string): Promise<string> {
	const dv = new DataView(zip.buffer, zip.byteOffset, zip.byteLength);
	// The end of central directory record, then each central entry.
	let eocd = zip.length - 22;
	while (eocd >= 0 && dv.getUint32(eocd, true) !== 0x06054b50) eocd--;
	if (eocd < 0) throw new Error('not a zip file');
	let at = dv.getUint32(eocd + 16, true);
	const count = dv.getUint16(eocd + 10, true);
	for (let n = 0; n < count; n++) {
		const method = dv.getUint16(at + 10, true);
		const size = dv.getUint32(at + 20, true);
		const nameLen = dv.getUint16(at + 28, true);
		const extraLen = dv.getUint16(at + 30, true);
		const commentLen = dv.getUint16(at + 32, true);
		const local = dv.getUint32(at + 42, true);
		const name = new TextDecoder().decode(zip.subarray(at + 46, at + 46 + nameLen));
		at += 46 + nameLen + extraLen + commentLen;
		if (!name.endsWith(suffix)) continue;
		const start = local + 30 + dv.getUint16(local + 26, true) + dv.getUint16(local + 28, true);
		const data = zip.subarray(start, start + size);
		if (method === 0) return new TextDecoder().decode(data);
		const stream = new Blob([data as Uint8Array<ArrayBuffer>]).stream().pipeThrough(new DecompressionStream('deflate-raw'));
		return await new Response(stream).text();
	}
	throw new Error(`no ${suffix} file in the zip`);
}
