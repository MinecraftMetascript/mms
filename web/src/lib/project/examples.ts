import type { AssetFile, ProjectData } from './project.svelte';

// The repository's examples, bundled at build time, with their subfolders.
const raw = import.meta.glob('../../../../examples/**/*.{mms,json}', {
	query: '?raw',
	import: 'default',
	eager: true
}) as Record<string, string>;

// Their textures and sounds, as data: URLs.
const binary = import.meta.glob('../../../../examples/**/*.{png,ogg}', {
	query: '?inline',
	import: 'default',
	eager: true
}) as Record<string, string>;

export interface Example {
	id: string;
	description: string;
	data: ProjectData;
}

export const examples: Example[] = (() => {
	type Dir = { config?: string; files: { path: string; text: string }[]; assets: AssetFile[] };
	const byDir = new Map<string, Dir>();
	// ../../../../examples/<dir>/<path inside the example>
	const split = (path: string) => {
		const [dir, ...rest] = path.slice('../../../../examples/'.length).split('/');
		const e = byDir.get(dir) ?? { files: [], assets: [] };
		byDir.set(dir, e);
		return [e, rest.join('/')] as const;
	};
	for (const [path, text] of Object.entries(raw)) {
		const [e, name] = split(path);
		if (name === 'mms.json') e.config = text;
		else if (name.endsWith('.mms')) e.files.push({ path: name, text });
	}
	for (const [path, url] of Object.entries(binary)) {
		const [e, name] = split(path);
		const data = url.slice(url.indexOf(',') + 1);
		if (url.startsWith('data:') && url.includes(';base64,')) e.assets.push({ path: name, data });
	}
	return [...byDir.entries()]
		.sort(([a], [b]) => a.localeCompare(b))
		.filter(([, e]) => e.files.length)
		.map(([id, e]) => {
			// Top-level files first, then each folder's, alphabetically.
			e.files.sort((a, b) => a.path.split('/').length - b.path.split('/').length || a.path.localeCompare(b.path));
			const cfg = e.config ? JSON.parse(e.config) : {};
			return {
				id,
				description: cfg.description ?? id,
				data: {
					name: id,
					config: { namespace: cfg.namespace ?? 'mms', description: cfg.description ?? '' },
					files: e.files,
					assets: e.assets.sort((a, b) => a.path.localeCompare(b.path)),
					active: e.files[0]?.path ?? ''
				}
			};
		});
})();
