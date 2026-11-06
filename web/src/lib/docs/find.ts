// Site search over the index from search.ts, run in the browser.

/** One search result: a page, a section of one, or a reference entry. */
export interface Entry {
	title: string;
	/** Where it is, such as "Guides" or "Helpers · Loot function". */
	where: string;
	/** The path under the site's base, such as "/docs/guides/functions#calls". */
	path: string;
	text: string;
}

/**
 * The entries that match q, best first: a title that is the query, then
 * one that starts with it, contains it, and last text that contains it.
 * Every word of q must match somewhere.
 */
export function search(entries: Entry[], q: string, limit = 30): Entry[] {
	const words = q.toLowerCase().split(/\s+/).filter(Boolean);
	if (!words.length) return [];
	const scored: { e: Entry; s: number }[] = [];
	for (const e of entries) {
		const title = e.title.toLowerCase();
		const hay = `${title} ${e.where.toLowerCase()} ${e.text.toLowerCase()}`;
		if (!words.every((w) => hay.includes(w))) continue;
		const w = words.join(' ');
		const s = title === w ? 0 : title.startsWith(w) ? 1 : title.includes(w) ? 2 : words.every((x) => title.includes(x)) ? 3 : 4;
		scored.push({ e, s: s * 10 + (e.path.includes('#') ? 1 : 0) });
	}
	return scored
		.sort((a, b) => a.s - b.s || a.e.title.length - b.e.title.length)
		.slice(0, limit)
		.map((x) => x.e);
}
