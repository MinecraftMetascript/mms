// Groups the World view's structures into folders, so a world with
// hundreds of saved pieces stays readable. A name is split at `/` and `_`
// (village/house_small_1 → village/ › house_ › small_ › 1), and a folder is
// only made where two or more structures share it: a lone structure is
// shown directly, and a folder holding one subfolder is merged into it
// (house_small_). Only the list is grouped; files and ids keep their names.

export interface Leaf<T> {
	kind: 'leaf';
	item: T;
}

export interface Group<T> {
	kind: 'group';
	/** Unique among all groups, for remembering which are open. */
	key: string;
	/** The shared prefix this group adds, with its separator, such as `house_`. */
	label: string;
	children: Node<T>[];
	/** Every item under this group, at any depth. */
	items: T[];
}

export type Node<T> = Leaf<T> | Group<T>;

/**
 * Splits a name after each `/` and `_` (and the `:` of a namespaced id),
 * keeping the separators: a_b/c → ["a_", "b/", "c"].
 */
export function segments(name: string): string[] {
	return name.match(/[^:/_]*(?:[:/_]|$)/g)!.filter(Boolean);
}

/**
 * The folder tree of items named by name(item). Items keep their order
 * within each folder; folders come first, then single items.
 */
export function group<T>(items: T[], name: (item: T) => string): Node<T>[] {
	const entries = items.map((item, at) => ({ item, at, segs: segments(name(item)) }));
	return build(entries, 0, '');
}

type Entry<T> = { item: T; at: number; segs: string[] };

function build<T>(entries: Entry<T>[], depth: number, prefix: string): Node<T>[] {
	const groups = new Map<string, Entry<T>[]>();
	const leaves: Entry<T>[] = [];
	for (const e of entries) {
		// The last segment is the item's own name, not a folder.
		if (e.segs.length <= depth + 1) {
			leaves.push(e);
			continue;
		}
		const seg = e.segs[depth];
		const list = groups.get(seg);
		if (list) list.push(e);
		else groups.set(seg, [e]);
	}
	const out: Node<T>[] = [];
	for (const [seg, list] of groups) {
		if (list.length === 1) {
			leaves.push(list[0]);
			continue;
		}
		const key = prefix + seg;
		let g: Group<T> = { kind: 'group', key, label: seg, children: build(list, depth + 1, key), items: list.map((e) => e.item) };
		// A folder whose only content is one folder reads as one: house_ › small_ → house_small_.
		while (g.children.length === 1 && g.children[0].kind === 'group') {
			const inner: Group<T> = g.children[0];
			g = { ...inner, label: g.label + inner.label };
		}
		out.push(g);
	}
	leaves.sort((a, b) => a.at - b.at);
	return [...out, ...leaves.map((e): Node<T> => ({ kind: 'leaf', item: e.item }))];
}
