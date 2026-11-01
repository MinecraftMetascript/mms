import { strFromU8, strToU8, unzipSync, zipSync, type Zippable } from 'fflate';
import type { OutputFile } from '../engine/engine.svelte';
import type { ProjectData } from './project.svelte';

const fixed = new Date(1980, 0, 1);

/** The built data pack as a zip. */
export function packZip(files: OutputFile[]): Uint8Array {
	const z: Zippable = {};
	for (const f of files) z[f.path] = [strToU8(f.text), { mtime: fixed }];
	return zipSync(z, { level: 6 });
}

/** The project's sources and mms.json as a zip, the layout `mms build` reads. */
export function projectZip(p: ProjectData): Uint8Array {
	const z: Zippable = {};
	z['mms.json'] = [strToU8(JSON.stringify({ $schema: 'https://minecraftmetascript.github.io/mms/mms.schema.json', namespace: p.config.namespace, description: p.config.description }, null, 2) + '\n'), { mtime: fixed }];
	for (const f of p.files) z[f.path] = [strToU8(f.text), { mtime: fixed }];
	return zipSync(z, { level: 6 });
}

/** A zip's text files, with a single top-level folder stripped. */
export function readZip(data: Uint8Array): Map<string, string> {
	const entries = unzipSync(data);
	const paths = Object.keys(entries).filter((p) => !p.endsWith('/') && !p.startsWith('__MACOSX/'));
	const tops = new Set(paths.map((p) => p.split('/')[0]));
	let prefix = '';
	if (tops.size === 1 && paths.every((p) => p.includes('/'))) prefix = [...tops][0] + '/';
	const out = new Map<string, string>();
	for (const p of paths) {
		const rel = p.slice(prefix.length);
		// Structures (.nbt) and other binary files have no MMS form.
		if (/\.(mms|json|mcmeta|mcfunction)$/.test(rel) || rel === 'mms.json') out.set(rel, strFromU8(entries[p]));
	}
	return out;
}

/** Whether a zip holds a data pack rather than an MMS project. */
export function isDataPack(files: Map<string, string>): boolean {
	const paths = [...files.keys()];
	return !paths.some((p) => p.endsWith('.mms')) && paths.some((p) => p === 'pack.mcmeta' || p.startsWith('data/'));
}

/** Reads an MMS project zip. */
export function readProjectZip(name: string, entries: Map<string, string>): ProjectData {
	const config = { namespace: 'mms', description: '' };
	const files = [];
	for (const [rel, text] of entries) {
		if (rel === 'mms.json') {
			try {
				Object.assign(config, JSON.parse(text));
			} catch {
				/* keep defaults */
			}
		} else if (rel.endsWith('.mms')) {
			files.push({ path: rel, text });
		}
	}
	if (!files.length) throw new Error('the zip has no .mms files');
	files.sort((a, b) => a.path.localeCompare(b.path));
	return { name: name.replace(/\.zip$/i, ''), config, files, active: files[0].path };
}

export function download(name: string, data: Uint8Array | string, type = 'application/zip') {
	const blob = new Blob([data as BlobPart], { type });
	const url = URL.createObjectURL(blob);
	const a = document.createElement('a');
	a.href = url;
	a.download = name;
	a.click();
	setTimeout(() => URL.revokeObjectURL(url), 1000);
}
