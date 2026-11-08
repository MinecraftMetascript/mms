// Groups the World view's structures into folders, so a world with
// hundreds of saved pieces stays readable. A structure id's path is split at
// `/`, like the folders it is saved in (village/houses/small → village ›
// houses › small), and a folder is only made where two or more structures
// share it: a lone structure is shown directly, and a folder holding one
// subfolder is merged into it (village/houses). Only the list is grouped;
// files and ids keep their names.

export interface Leaf<T> {
	kind: 'leaf';
	item: T;
}

export interface Group<T> {
	kind: 'group';
	/** Unique among all groups, for remembering which are open. */
	key: string;
	/** The folder's name, without a trailing separator, such as `houses` or `village/houses` when merged. */
	label: string;
	children: Node<T>[];
	/** Every item under this group, at any depth. */
	items: T[];
}

export type Node<T> = Leaf<T> | Group<T>;

/**
 * Splits a name after each `/` (and the `:` of a namespaced id), keeping
 * the separators: ns:a/b_c → ["ns:", "a/", "b_c"].
 */
export function segments(name: string): string[] {
	return name.match(/[^:/]*(?:[:/]|$)/g)!.filter(Boolean);
}

/** A folder's shown name: its path without the trailing separator. */
const folderName = (raw: string) => raw.replace(/[:/]$/, '');

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
		let g: Group<T> = { kind: 'group', key, label: folderName(seg), children: build(list, depth + 1, key), items: list.map((e) => e.item) };
		// A folder whose only content is one folder reads as one: village › houses → village/houses.
		while (g.children.length === 1 && g.children[0].kind === 'group') {
			const inner: Group<T> = g.children[0];
			g = { ...inner, label: folderName(inner.key.slice(prefix.length)) };
		}
		out.push(g);
	}
	leaves.sort((a, b) => a.at - b.at);
	return [...out, ...leaves.map((e): Node<T> => ({ kind: 'leaf', item: e.item }))];
}
