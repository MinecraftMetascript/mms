// Which generated files the previews and inspectors show, by output path.
// Free of DOM types, so the VS Code extension host can use it too.

export type PreviewKind = 'noise' | 'density_function' | 'noise_settings' | 'material_rule' | 'dimension' | 'carver';
/** The registry folders the preview worker reads. Biomes are read for the carvers they list. */
export type PreviewFileKind = PreviewKind | 'material_condition' | 'tag/worldgen/biome' | 'biome';

const previewPattern = /^data\/([^/]+)\/(worldgen\/(?:noise|density_function|noise_settings|material_rule|material_condition|carver|biome)|dimension|tags\/worldgen\/biome)\/(.+)\.json$/;

/** The registry folder and id of an output path the preview worker reads, or null. */
export function previewFile(path: string): { kind: PreviewFileKind; id: string } | null {
	const m = previewPattern.exec(path);
	if (!m) return null;
	const kind = m[2] === 'tags/worldgen/biome' ? 'tag/worldgen/biome' : m[2].replace('worldgen/', '');
	return { kind: kind as PreviewFileKind, id: `${m[1]}:${m[3]}` };
}

/** The preview kind and id of an output path, or null if it has no preview. */
export function previewable(path: string | null): { kind: PreviewKind; id: string } | null {
	const f = path ? previewFile(path) : null;
	return f && f.kind !== 'material_condition' && f.kind !== 'tag/worldgen/biome' && f.kind !== 'biome' ? { kind: f.kind, id: f.id } : null;
}

/** Which inspector, if any, shows a generated file. */
export type InspectKind = 'loot_table' | 'recipe' | 'biome';

const inspectPattern = /^data\/([^/]+)\/(loot_table|recipe|worldgen\/biome)\/(.+)\.json$/;

export function inspectable(path: string | null): { kind: InspectKind; id: string } | null {
	const m = path ? inspectPattern.exec(path) : null;
	if (!m) return null;
	const kind = (m[2] === 'worldgen/biome' ? 'biome' : m[2]) as InspectKind;
	return { kind, id: `${m[1]}:${m[3]}` };
}

const jigsawPattern = /^data\/([^/]+)\/worldgen\/(template_pool|structure)\/(.+)\.json$/;

/**
 * The template pool or structure a jigsaw connection graph starts from.
 * Any structure qualifies by path; only a jigsaw one has a graph.
 */
export function jigsawable(path: string | null): { kind: 'pool' | 'structure'; id: string } | null {
	const m = path ? jigsawPattern.exec(path) : null;
	return m ? { kind: m[2] === 'template_pool' ? 'pool' : 'structure', id: `${m[1]}:${m[3]}` } : null;
}

const functionPattern = /^data\/([^/]+)\/function\/(.+)\.mcfunction$/;

/** The function an output path names, which the function preview runs. */
export function runnable(path: string | null): { kind: 'function'; id: string } | null {
	const m = path ? functionPattern.exec(path) : null;
	return m ? { kind: 'function', id: `${m[1]}:${m[2]}` } : null;
}

/** The key an editor selects a custom item or book by. */
export const itemPath = (id: string) => `item:${id}`;

/** Whether an item stack (as JSON) is a written book with content. */
export const isBook = (stack: string) => /"(minecraft:)?written_book_content"\s*:/.test(stack);

/** The book an output path names, if it is item:<id> of a book in items. */
export function bookAt(path: string | null, items: readonly { id: string; text: string }[] = []): { kind: 'book'; id: string } | null {
	if (!path?.startsWith('item:')) return null;
	const it = items.find((it) => itemPath(it.id) === path);
	return it && isBook(it.text) ? { kind: 'book', id: it.id } : null;
}

/**
 * The page of a book at a source line, from 0: the last page starting at
 * or before it in the same file, or null when the line is in none.
 */
export function bookPageAt(pages: readonly { uri: string; range: { start: { line: number } } }[] | undefined, uri: string, line: number): number | null {
	let at: number | null = null;
	pages?.forEach((p, i) => {
		if (p.uri === uri && p.range.start.line <= line) at = i;
	});
	return at;
}

/**
 * What the preview pane shows for an output path: a preview, an
 * inspector, a jigsaw graph, a book, a function run, or nothing. Books
 * need the items to tell.
 */
export const shows = (path: string | null, items?: readonly { id: string; text: string }[]) =>
	previewable(path) ?? inspectable(path) ?? jigsawable(path) ?? bookAt(path, items) ?? runnable(path);
