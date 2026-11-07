// Site search over the index from search.ts, run in the browser with
// FlexSearch: partial words, several words in any order, and a close
// match for a typo.
import { Charset, Index } from 'flexsearch';

/** One search result: a page, a section of one, or a reference entry. */
export interface Entry {
	title: string;
	/** Where it is, such as "Guides" or "Helpers · Loot function". */
	where: string;
	/** The path under the site's base, such as "/docs/guides/functions#calls". */
	path: string;
	text: string;
}

/** Splits a name into words: "SetCount" also reads as "set count". */
const words = (s: string) => `${s} ${s.replace(/([a-z0-9])([A-Z])/g, '$1 $2').replace(/[_:/.-]+/g, ' ')}`;

/** A search over a list of entries, built once when search first opens. */
export class Finder {
	private titles = new Index({ tokenize: 'full', encoder: Charset.LatinBalance });
	private texts = new Index({ tokenize: 'forward', encoder: Charset.LatinBalance });

	constructor(private entries: Entry[]) {
		entries.forEach((e, i) => {
			this.titles.add(i, words(e.title));
			this.texts.add(i, `${words(e.where)} ${e.text}`);
		});
	}

	/**
	 * The entries that match q, best first: titles that are the query,
	 * start with it or contain it, then titles that match its words, then
	 * entries whose text does. Pages come before their sections. When
	 * nothing matches every word, close spellings are tried.
	 */
	search(q: string, limit = 30): Entry[] {
		const query = q.trim();
		if (!query) return [];
		const lower = query.toLowerCase();
		const rank = (i: number, base: number) => {
			const t = this.entries[i].title.toLowerCase();
			const exact = t === lower ? 0 : t.startsWith(lower) ? 1 : t.includes(lower) ? 2 : 3;
			return base + exact * 10 + (this.entries[i].path.includes('#') ? 1 : 0);
		};
		const run = (suggest: boolean) => {
			const seen = new Map<number, number>();
			const add = (ids: unknown[], base: number) => {
				for (const id of ids as number[]) if (!seen.has(id)) seen.set(id, rank(id, base));
			};
			add(this.titles.search(query, { limit: 200, suggest }), 0);
			add(this.texts.search(query, { limit: 200, suggest }), 100);
			return [...seen].sort((a, b) => a[1] - b[1] || this.entries[a[0]].title.length - this.entries[b[0]].title.length);
		};
		let hits = run(false);
		if (!hits.length) hits = run(true);
		return hits.slice(0, limit).map(([i]) => this.entries[i]);
	}
}
