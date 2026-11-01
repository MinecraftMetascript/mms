import { strFromU8, strToU8, unzipSync, zipSync, type Zippable } from 'fflate';
import type { OutputAsset, OutputFile } from '../engine/engine.svelte';
import { fromBase64, toBase64 } from './base64';
import { isAssetPath, type AssetFile, type ProjectData } from './project.svelte';

const fixed = new Date(1980, 0, 1);

/**
 * One built pack as a zip: the data pack, or with resources the resource
 * pack. Copied files (textures, sounds) come from the project's assets.
 */
export function packZip(files: OutputFile[], copies: OutputAsset[], assets: AssetFile[], resources = false): Uint8Array {
	const z: Zippable = {};
	const mine = (pack?: string) => (pack === 'resources') === resources;
	for (const f of files) if (mine(f.pack)) z[f.path] = [strToU8(f.text), { mtime: fixed }];
	for (const c of copies) {
		const a = assets.find((a) => a.path === c.src);
		if (a && mine(c.pack)) z[c.path] = [fromBase64(a.data), { mtime: fixed, level: 0 }];
	}
	return zipSync(z, { level: 6 });
}

/** Whether the build has a resource pack as well as a data pack. */
export const hasResourcePack = (files: OutputFile[], copies: OutputAsset[]) =>
	files.some((f) => f.pack === 'resources') || copies.some((c) => c.pack === 'resources');

/** The project's sources and mms.json as a zip, the layout `mms build` reads. */
export function projectZip(p: ProjectData): Uint8Array {
	const z: Zippable = {};
	z['mms.json'] = [strToU8(JSON.stringify({ $schema: 'https://minecraftmetascript.github.io/mms/mms.schema.json', namespace: p.config.namespace, description: p.config.description }, null, 2) + '\n'), { mtime: fixed }];
	for (const f of p.files) z[f.path] = [strToU8(f.text), { mtime: fixed }];
	for (const a of p.assets ?? []) z[a.path] = [fromBase64(a.data), { mtime: fixed, level: 0 }];
	return zipSync(z, { level: 6 });
}

/** A zip's files, by path with a single top-level folder stripped. */
function unzipEntries(data: Uint8Array): Map<string, Uint8Array> {
	const entries = unzipSync(data);
	const paths = Object.keys(entries).filter((p) => !p.endsWith('/') && !p.startsWith('__MACOSX/'));
	const tops = new Set(paths.map((p) => p.split('/')[0]));
	let prefix = '';
	if (tops.size === 1 && paths.every((p) => p.includes('/'))) prefix = [...tops][0] + '/';
	return new Map(paths.map((p) => [p.slice(prefix.length), entries[p]]));
}

/** A zip's text files, with a single top-level folder stripped. */
export function readZip(data: Uint8Array): Map<string, string> {
	const out = new Map<string, string>();
	for (const [rel, bytes] of unzipEntries(data)) {
		if (/\.(mms|json|mcmeta|mcfunction)$/.test(rel) || rel === 'mms.json') out.set(rel, strFromU8(bytes));
	}
	return out;
}

/** A project zip's textures, sounds and structures. */
export function readZipAssets(data: Uint8Array): AssetFile[] {
	const out: AssetFile[] = [];
	for (const [rel, bytes] of unzipEntries(data)) if (isAssetPath(rel)) out.push({ path: rel, data: toBase64(bytes) });
	return out.sort((a, b) => a.path.localeCompare(b.path));
}

/** Whether a zip holds a data pack rather than an MMS project. */
export function isDataPack(files: Map<string, string>): boolean {
	const paths = [...files.keys()];
	return !paths.some((p) => p.endsWith('.mms')) && paths.some((p) => p === 'pack.mcmeta' || p.startsWith('data/'));
}

/** Reads an MMS project zip. */
export function readProjectZip(name: string, entries: Map<string, string>, assets: AssetFile[] = []): ProjectData {
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
	return { name: name.replace(/\.zip$/i, ''), config, files, assets, active: files[0].path };
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
