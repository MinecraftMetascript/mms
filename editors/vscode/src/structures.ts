// Structures saved in a world with structure blocks, which a project can
// pull in for its template pools. The game writes a structure block's SAVE
// to <world>/generated/<namespace>/structure/<path>.nbt (structures/ before
// 1.21). Importing copies the file into the project, under the structures
// folder (mms.structuresFolder), where StructureFile("structures/<path>.nbt")
// puts it in the pack with the id <namespace>:<path>.
import * as crypto from 'crypto';
import * as fs from 'fs';
import * as path from 'path';

export interface WorldStructure {
	/** The namespace it was saved under in the game. */
	ns: string;
	/** Its id path, such as houses/small. */
	path: string;
	/** The .nbt file in the world. */
	file: string;
	saved: Date;
}

/** How a world's structure compares with the project's copy. */
export type Status = 'new' | 'imported' | 'changed';

/** The structures saved in a world, by id. */
export function findStructures(world: string): WorldStructure[] {
	const out: WorldStructure[] = [];
	const generated = path.join(world, 'generated');
	for (const ns of dirs(generated)) {
		// 1.21 renamed the folder; a world played in both may have both.
		for (const folder of ['structures', 'structure']) {
			const root = path.join(generated, ns, folder);
			for (const rel of walk(root)) {
				if (!rel.endsWith('.nbt')) continue;
				const p = rel.slice(0, -4).split(path.sep).join('/');
				const file = path.join(root, rel);
				const s: WorldStructure = { ns, path: p, file, saved: fs.statSync(file).mtime };
				// The newer save wins when both folders have it.
				const i = out.findIndex((o) => o.ns === ns && o.path === p);
				if (i < 0) out.push(s);
				else if (out[i].saved < s.saved) out[i] = s;
			}
		}
	}
	return out.sort((a, b) => `${a.ns}:${a.path}`.localeCompare(`${b.ns}:${b.path}`));
}

/**
 * Where a structure goes in the project, relative to its folder:
 * <folder>/<path>.nbt for the project's namespace, and
 * <folder>-<namespace>/<path>.nbt for another. Each namespace gets its own
 * folder, so a:b/foo and b:foo never share a file.
 */
export function importPath(s: WorldStructure, projectNs: string | undefined, folder: string): string {
	const base = folder.replace(/\\/g, '/').replace(/^\.?\/+|\/+$/g, '') || 'structures';
	return path.posix.normalize(`${s.ns === projectNs ? base : `${base}-${s.ns}`}/${s.path}.nbt`);
}

/**
 * The helper call that puts an imported structure in the pack. The id name
 * keeps the folders of the path; without it the id would be the bare
 * file name.
 */
export function structureCall(rel: string, s: WorldStructure): string {
	const base = s.path.split('/').pop()!;
	const defaultName = base.toLowerCase().replace(/[^a-z0-9_.-]/g, '_');
	return s.path === defaultName ? `StructureFile("${rel}")` : `StructureFile("${rel}", "${s.path}")`;
}

/** Whether the project has the structure, and the same bytes as the world. */
export function status(s: WorldStructure, dest: string): Status {
	let mine: Buffer;
	try {
		mine = fs.readFileSync(dest);
	} catch {
		return 'new';
	}
	return hash(mine) === hash(fs.readFileSync(s.file)) ? 'imported' : 'changed';
}

/** Copies a structure into the project, making folders as needed. */
export function importStructure(s: WorldStructure, dest: string) {
	fs.mkdirSync(path.dirname(dest), { recursive: true });
	fs.copyFileSync(s.file, dest);
}

function hash(b: Buffer): string {
	return crypto.createHash('sha1').update(b).digest('hex');
}

function dirs(dir: string): string[] {
	try {
		return fs
			.readdirSync(dir, { withFileTypes: true })
			.filter((e) => e.isDirectory())
			.map((e) => e.name);
	} catch {
		return [];
	}
}

/** Files under root, as paths relative to it. */
function walk(root: string, rel = ''): string[] {
	let entries: fs.Dirent[];
	try {
		entries = fs.readdirSync(path.join(root, rel), { withFileTypes: true });
	} catch {
		return [];
	}
	return entries.flatMap((e) => (e.isDirectory() ? walk(root, path.join(rel, e.name)) : e.isFile() ? [path.join(rel, e.name)] : []));
}
