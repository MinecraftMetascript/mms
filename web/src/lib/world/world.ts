// Building a project's pack into a world folder and finding the structures
// a world saved, for the World tab. The engine works out what a build
// writes (mms/worldInstall, the same plan `mms build --world` follows);
// this writes it, and takes the chunks of reset dimensions out the way
// project.ResetDimensions does.
import { dimensionFolders, savedStructure, vanillaDimensions, type WorldDimensionRow } from '@mms/preview/world/tool';
import { fromBase64, toBase64 } from '#lib/project/base64.ts';
import type { WorldInstall } from '#lib/engine/engine.svelte.ts';
import type { WorldFolder } from './folder';

/** Where a reset moves a dimension's old chunks, as in project.ResetFolder. */
export const resetFolder = 'mms-reset';
const chunkDirs = ['region', 'entities', 'poi'];

/** The data pack already in the world, as project.ReadPackDir reads it, base64 by path. */
export async function readPackDir(world: WorldFolder, dir: string): Promise<Record<string, string>> {
	const out: Record<string, string> = {};
	for (const f of ['pack.mcmeta', 'pack.png']) {
		const file = await world.file(`${dir}/${f}`);
		if (file) out[f] = toBase64(new Uint8Array(await file.arrayBuffer()));
	}
	for (const { path, file } of await world.walk(`${dir}/data`)) out[path.slice(dir.length + 1)] = toBase64(new Uint8Array(await file.arrayBuffer()));
	return out;
}

/** Writes the data pack the way project.WriteDir does: data/, assets/ and pack.png are replaced. */
export async function writePackDir(world: WorldFolder, dir: string, files: WorldInstall['data']): Promise<void> {
	for (const old of ['data', 'assets', 'pack.png']) await world.remove(`${dir}/${old}`);
	for (const f of files) {
		if (f.path.split('/').some((p) => p === '' || p === '.' || p === '..')) throw new Error(`${f.path} is not a path inside the pack`);
		await world.write(`${dir}/${f.path}`, fromBase64(f.data));
	}
}

/** Where a world saves a dimension, as project.DimensionFolder says. */
async function dimensionFolder(world: WorldFolder, id: string): Promise<string> {
	const [now, legacy] = dimensionFolders(id).map((p) => p.join('/'));
	if (legacy !== undefined && !(await world.isDir(now)) && (await world.isDir(legacy ? `${legacy}/region` : 'region'))) return legacy;
	return now;
}

/**
 * Takes the saved chunks of each dimension out of the world into
 * mms-reset/, after clearing an earlier reset out of it. Moved lists the
 * folders taken out; none means the dimension had not been visited.
 */
export async function resetDimensions(world: WorldFolder, ids: string[]): Promise<{ id: string; moved: string[] }[]> {
	let cleared = false;
	const out: { id: string; moved: string[] }[] = [];
	for (const id of ids) {
		const dir = await dimensionFolder(world, id);
		const r = { id, moved: [] as string[] };
		for (const sub of chunkDirs) {
			const from = dir ? `${dir}/${sub}` : sub;
			if (!(await world.isDir(from))) continue;
			if (!cleared) {
				await world.remove(resetFolder);
				cleared = true;
			}
			// No folder move in the browser: copy, then delete.
			for (const { path, file } of await world.walk(from)) await world.write(`${resetFolder}/${path}`, new Uint8Array(await file.arrayBuffer()));
			await world.remove(from);
			r.moved.push(from);
		}
		out.push(r);
	}
	return out;
}

/** The pack's dimensions, and whether the world has saved chunks for each. */
export async function dimensionRows(world: WorldFolder, ids: string[]): Promise<WorldDimensionRow[]> {
	const rows: WorldDimensionRow[] = [];
	for (const id of ids) {
		let generated = false;
		for (const d of dimensionFolders(id)) if (await world.isDir([...d, 'region'].join('/'))) generated = true;
		rows.push({ id, vanilla: vanillaDimensions.has(id), generated });
	}
	return rows;
}

/** A structure saved in the world, by its file's path in the world. */
export interface FoundStructure {
	ns: string;
	path: string;
	file: string;
	saved: Date;
	data: () => Promise<Uint8Array>;
}

/** The structures a world saved with structure blocks, by id; the newer save wins when both folders have one. */
export async function findStructures(world: WorldFolder): Promise<FoundStructure[]> {
	const out: FoundStructure[] = [];
	for (const { path, file } of await world.walk('generated')) {
		const s = savedStructure(path);
		if (!s) continue;
		const found: FoundStructure = { ns: s.ns, path: s.path, file: path, saved: new Date(file.lastModified), data: async () => new Uint8Array(await file.arrayBuffer()) };
		const i = out.findIndex((o) => o.ns === s.ns && o.path === s.path);
		if (i < 0) out.push(found);
		else if (out[i].saved < found.saved) out[i] = found;
	}
	return out.sort((a, b) => `${a.ns}:${a.path}`.localeCompare(`${b.ns}:${b.path}`));
}
