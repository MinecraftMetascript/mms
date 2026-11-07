import { strFromU8, strToU8, unzipSync, zipSync, type Zippable } from 'fflate';
import type { OutputAsset, OutputFile } from '../engine/engine.svelte';
import { fromBase64, toBase64 } from './base64';
import { isAssetPath, isSourcePath, safeProjectPath, type AssetFile, type ProjectData } from './project.svelte';

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

/** Entries larger than this are left out: no project file comes near it. */
const maxEntry = 32 << 20;

/** A zip's text files and assets, read in one pass. */
export interface ZipContents {
	/** Text files (sources, JSON, mcfunction), by path. */
	texts: Map<string, string>;
	/** Textures, sounds and structures. */
	assets: AssetFile[];
}

const isTextPath = (p: string) => isSourcePath(p) || /\.(json|mcmeta|mcfunction)$/.test(p);

/**
 * Reads a zip once: its files by path, with a single top-level folder
 * stripped and every path made safe (no '..', no leading '/', no '\\').
 * Folders, macOS metadata and oversized entries are skipped.
 */
export function readZip(data: Uint8Array): ZipContents {
	const entries = unzipSync(data, {
		filter: (f) => !f.name.endsWith('/') && !f.name.startsWith('__MACOSX/') && f.originalSize <= maxEntry
	});
	const paths = Object.keys(entries);
	const tops = new Set(paths.map((p) => p.split('/')[0]));
	let prefix = '';
	if (tops.size === 1 && paths.every((p) => p.includes('/'))) prefix = [...tops][0] + '/';
	const texts = new Map<string, string>();
	const assets: AssetFile[] = [];
	for (const p of paths) {
		const rel = safeProjectPath(p.slice(prefix.length));
		if (!rel) continue;
		if (isTextPath(rel)) texts.set(rel, strFromU8(entries[p]));
		else if (isAssetPath(rel)) assets.push({ path: rel, data: toBase64(entries[p]) });
	}
	assets.sort((a, b) => a.path.localeCompare(b.path));
	return { texts, assets };
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
				const j = JSON.parse(text);
				if (typeof j?.namespace === 'string') config.namespace = j.namespace;
				if (typeof j?.description === 'string') config.description = j.description;
			} catch {
				/* keep defaults */
			}
		} else if (isSourcePath(rel)) {
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
