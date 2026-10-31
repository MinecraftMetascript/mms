export interface Node {
	name: string;
	path: string;
	children: Node[];
	leaf: boolean;
}

/** Builds a folder tree from slash-separated paths, folders first. */
export function buildTree(paths: string[]): Node[] {
	const root: Node = { name: '', path: '', children: [], leaf: false };
	for (const p of paths) {
		const parts = p.split('/');
		let cur = root;
		parts.forEach((part, i) => {
			const path = parts.slice(0, i + 1).join('/');
			const leaf = i === parts.length - 1;
			let next = cur.children.find((c) => c.name === part && c.leaf === leaf);
			if (!next) {
				next = { name: part, path, children: [], leaf };
				cur.children.push(next);
			}
			cur = next;
		});
	}
	const sort = (n: Node) => {
		n.children.sort((a, b) => (a.leaf === b.leaf ? a.name.localeCompare(b.name) : a.leaf ? 1 : -1));
		n.children.forEach(sort);
	};
	sort(root);
	// Collapse single-child folder chains (data/ns/worldgen → one row).
	const squash = (n: Node): Node => {
		n.children = n.children.map(squash);
		if (!n.leaf && n.children.length === 1 && !n.children[0].leaf && n.name) {
			const c = n.children[0];
			return { ...c, name: `${n.name}/${c.name}` };
		}
		return n;
	};
	return squash(root).children;
}
