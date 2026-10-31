import type { ProjectData } from './project.svelte';

// The repository's examples, bundled at build time, with their subfolders.
const raw = import.meta.glob('../../../../examples/**/*', {
	query: '?raw',
	import: 'default',
	eager: true
}) as Record<string, string>;

export interface Example {
	id: string;
	description: string;
	data: ProjectData;
}

export const examples: Example[] = (() => {
	const byDir = new Map<string, { config?: string; files: { path: string; text: string }[] }>();
	for (const [path, text] of Object.entries(raw)) {
		// ../../../../examples/<dir>/<path inside the example>
		const [dir, ...rest] = path.slice('../../../../examples/'.length).split('/');
		const name = rest.join('/');
		const e = byDir.get(dir) ?? { files: [] };
		if (name === 'mms.json') e.config = text;
		else if (name.endsWith('.mms')) e.files.push({ path: name, text });
		byDir.set(dir, e);
	}
	return [...byDir.entries()]
		.sort(([a], [b]) => a.localeCompare(b))
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
					active: e.files[0]?.path ?? ''
				}
			};
		});
})();
