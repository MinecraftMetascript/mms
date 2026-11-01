// Which generated files the previews and inspectors show, by output path.
// Free of DOM types, so the VS Code extension host can use it too.

export type PreviewKind = 'noise' | 'density_function' | 'noise_settings' | 'material_rule' | 'dimension';
/** The registry folders the preview worker reads. */
export type PreviewFileKind = PreviewKind | 'material_condition' | 'tag/worldgen/biome';

const previewPattern = /^data\/([^/]+)\/(worldgen\/(?:noise|density_function|noise_settings|material_rule|material_condition)|dimension|tags\/worldgen\/biome)\/(.+)\.json$/;

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
	return f && f.kind !== 'material_condition' && f.kind !== 'tag/worldgen/biome' ? { kind: f.kind, id: f.id } : null;
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

/** What the preview pane shows for an output path: a preview, an inspector, or nothing. */
export const shows = (path: string | null) => previewable(path) ?? inspectable(path);
