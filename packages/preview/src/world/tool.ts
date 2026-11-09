// The World and Structures pages' state (WorldTool.svelte), and the rules
// both editors share for structures saved in a world. The host (the VS
// Code extension's src/world.ts, or the web editor's World tab) owns the
// state and sends all of it after each change; the page only draws it and
// sends back what was clicked.

/** What the player does in the game after a build, from `mms build --world`. */
export type NextAction = 'reload' | 'rejoin' | 'none';

export type WorldResult =
	| {
			ok: true;
			files: number;
			at: string;
			next?: { action: NextAction; line: string; changed?: string };
			resources: boolean;
			/** Dimensions whose saved chunks the build took out, so they generate again. */
			reset?: string[];
			/** Dimensions it was asked to reset that had no chunks yet. */
			unvisited?: string[];
	  }
	| { ok: false; message: string; detail: string; at: string; problems: boolean; worldOpen?: boolean };

/** A dimension the project's pack declares. */
export interface WorldDimensionRow {
	/** Such as mypack:caves. */
	id: string;
	/** The overworld, nether or end, which the build never resets by itself. */
	vanilla: boolean;
	/** Whether the build world has saved chunks for it. */
	generated: boolean;
}

/** A structure a structure block saved in the world. */
export interface WorldStructureRow {
	/** Its .nbt file in the world, which names it in messages. */
	file: string;
	/** Its id in the game, such as tower:base. */
	id: string;
	status: StructureStatus;
	/** Where it goes (or is) in the project, relative to the mms.json folder. */
	rel: string;
	call: string;
	saved: string;
	/** When it was saved, in ms since the epoch, for "saved 5 min ago". */
	savedMs: number;
	/** Imported, but no StructureFile(...) call names the project copy (lint/unused-asset). */
	unused: boolean;
	/** Its id in the pack, when it was saved under another namespace. */
	packId?: string;
}

export interface WorldState {
	/** No project means no mms.json in the workspace. */
	project?: { name: string; ns?: string; path: string; canChoose: boolean };
	/** The world the pack is built into. None means none chosen yet for the project. */
	world?: { name: string; folder?: string; path: string };
	/**
	 * The world structures are imported from: a separate one when chosen
	 * (separate is true), else the build world.
	 */
	structureWorld?: { name: string; folder?: string; path: string; separate: boolean };
	building: boolean;
	buildOnSave: boolean;
	/** Reset the project's dimensions whose worldgen changed on each build (mms build --reset-dimensions). */
	resetDimensions: boolean;
	/** The pack's dimensions, shown with the build world. */
	dimensions: WorldDimensionRow[];
	result?: WorldResult;
	structures: WorldStructureRow[];
	/** Project .nbt files no StructureFile(...) call names that are not saves of the structure world. */
	unusedFiles: string[];
	/** The unused check waits for a project without errors. */
	unusedPaused: boolean;
}

/** What the page sends back when something is clicked. */
export type WorldAction =
	| {
			type: 'run';
			action: 'build' | 'chooseWorld' | 'chooseProject' | 'revealWorld' | 'forgetWorld' | 'showProblems' | 'showOutput' | 'copyReload' | 'dismiss' | 'importWaiting'
				| 'chooseStructureWorld' | 'sameStructureWorld' | 'revealStructureWorld';
	  }
	| { type: 'buildOnSave'; on: boolean }
	| { type: 'resetDimensions'; on: boolean }
	/** Build, and reset this one dimension (a WorldDimensionRow's id) whatever changed. */
	| { type: 'resetDimension'; id: string }
	/** file is a WorldStructureRow's file. */
	| { type: 'import'; file: string }
	| { type: 'copyCall'; file: string }
	| { type: 'deleteFromWorld'; file: string }
	/** files are WorldStructureRow files, imported together (a folder's Import button). */
	| { type: 'importFiles'; files: string[] }
	/** rel is relative to the project folder. */
	| { type: 'revealFile'; rel: string };

/**
 * What the host can do beyond building and importing. A missing label
 * leaves its button or menu item out.
 */
export interface WorldFeatures {
	/** The menu item that shows a project's structure file, such as "Show file in Explorer". */
	revealFile?: string;
	/** The build world's menu item that opens its datapacks folder. */
	revealWorld?: string;
	/** The structure world's menu item that opens its generated folder. */
	revealStructureWorld?: string;
	/** The label of the checkbox that builds after each save. */
	buildOnSave?: string;
	/** A failed build offers to show the build's output. */
	showOutput?: boolean;
	/** A structure's menu offers to delete it from the world. */
	deleteFromWorld?: boolean;
	/** A line under Choose world… on where worlds are found. */
	chooseWorldHint?: string;
}

/** How a world's structure compares with the project's copy. */
export type StructureStatus = 'new' | 'imported' | 'changed';

/** A structure saved in a world: <world>/generated/<ns>/structure/<path>.nbt. */
export interface SavedStructure {
	/** The namespace it was saved under in the game. */
	ns: string;
	/** Its id path, such as houses/small. */
	path: string;
}

/**
 * Where a structure goes in the project, relative to its folder:
 * <folder>/<path>.nbt for the project's namespace, and
 * <folder>-<namespace>/<path>.nbt for another. Each namespace gets its own
 * folder, so a:b/foo and b:foo never share a file.
 */
export function importPath(s: SavedStructure, projectNs: string | undefined, folder: string): string {
	const base = folder.replace(/\\/g, '/').replace(/^\.?\/+|\/+$/g, '') || 'structures';
	return normalize(`${s.ns === projectNs ? base : `${base}-${s.ns}`}/${s.path}.nbt`);
}

/** A slash path with "." parts and doubled slashes taken out, and ".." applied where it can be. */
function normalize(p: string): string {
	const out: string[] = [];
	for (const part of p.split('/')) {
		if (part === '' || part === '.') continue;
		if (part === '..' && out.length && out[out.length - 1] !== '..') out.pop();
		else out.push(part);
	}
	return out.join('/');
}

/**
 * The call that puts an imported structure in the pack. The id name
 * keeps the folders of the path; without it the id would be the bare
 * file name.
 */
export function structureCall(rel: string, s: SavedStructure): string {
	// Piece("tower/base") reads structures/tower/base.nbt as a pool entry,
	// with the id tower/base.
	if (rel === `structures/${s.path}.nbt`) return `Piece("${s.path}")`;
	const base = s.path.split('/').pop()!;
	const defaultName = base.toLowerCase().replace(/[^a-z0-9_.-]/g, '_');
	return s.path === defaultName ? `StructureFile("${rel}")` : `StructureFile("${rel}", "${s.path}")`;
}

/**
 * The structure a world file holds, from its path relative to the world:
 * generated/<ns>/structure/<path>.nbt, or structures/ before 1.21. A
 * world played in both versions may have both; the newer save wins.
 */
export function savedStructure(rel: string): (SavedStructure & { folder: string }) | undefined {
	const m = /^generated\/([^/]+)\/(structures?)\/(.+)\.nbt$/.exec(rel);
	return m ? { ns: m[1], folder: m[2], path: m[3] } : undefined;
}

/** The dimensions a pack declares: data/<ns>/dimension/<path>.json. */
export function packDimensions(paths: string[]): string[] {
	const ids = new Set<string>();
	for (const p of paths) {
		const m = /^data\/([^/]+)\/dimension\/(.+)\.json$/.exec(p);
		if (m) ids.add(`${m[1]}:${m[2]}`);
	}
	return [...ids].sort();
}

export const vanillaDimensions = new Set(['minecraft:overworld', 'minecraft:the_nether', 'minecraft:the_end']);

/**
 * The folders, relative to the world, where a dimension's chunks may be
 * saved: dimensions/<ns>/<path> since 26.1, and for a vanilla dimension
 * the older place (the world itself, DIM-1 or DIM1).
 */
export function dimensionFolders(id: string): string[][] {
	const [ns, name] = id.split(':');
	const dirs = [['dimensions', ns, ...name.split('/')]];
	if (id === 'minecraft:overworld') dirs.push([]);
	if (id === 'minecraft:the_nether') dirs.push(['DIM-1']);
	if (id === 'minecraft:the_end') dirs.push(['DIM1']);
	return dirs;
}

/** The part of a "next step" line after it, as the World page shows it. */
export function nextView(action: NextAction, line: string): { action: NextAction; line: string; changed?: string } {
	return { action, line, changed: /; (.*) changed\b/.exec(line)?.[1] };
}
