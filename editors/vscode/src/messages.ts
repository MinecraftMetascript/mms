// Messages between the extension and the preview panel's page (webview/).
// Type-only, so both builds can import it. OutputFile matches the one in
// packages/preview, which the Node build cannot import (it has Svelte files).

export type Range = { start: { line: number; character: number }; end: { line: number; character: number } };

/** A generated pack file, as the language server's `mms/output` returns it. */
export interface OutputFile {
	path: string;
	text: string;
	uri?: string;
	range?: Range;
	/** Written for an inline value where an id is required; range is that value's. */
	implicit?: boolean;
	/** Labels for parts of the file by JSON pointer, from doc comments, such as "/elements/2" of a template pool. */
	labels?: Record<string, string>;
	/** "resources" for a resource pack file. */
	pack?: string;
}

/** A project file the pack copies, such as a texture or a structure. */
export interface OutputAsset {
	/** Its path in the pack. */
	path: string;
	/** Its path in the project, relative to the mms.json folder. */
	src: string;
	/** "resources" for a resource pack file. */
	pack?: string;
	/** For a structure file: its jigsaw blocks. */
	jigsaws?: { pos: [number, number, number]; pool: string; name: string; target: string; orientation?: string; joint?: string }[];
	/** For a structure file: its size in blocks, x y z. */
	size?: [number, number, number];
}

/** A CustomItem or Book: the item stack its name stands for, which is not a file. */
export interface OutputItem {
	id: string;
	text: string;
	uri: string;
	range: Range;
	/** For a Book: where each page of its text starts. */
	pages?: { uri: string; range: Range }[];
}

export type ToPanel =
	/** The generated pack and custom items, refreshed after each analysis. */
	| {
			type: 'output';
			files: OutputFile[];
			items: OutputItem[];
			assets: OutputAsset[];
			/** The pack's copied .png textures as data URLs, by pack path: custom item icons. */
			textures?: Record<string, string>;
	  }
	/**
	 * The output path to show, chosen with a Preview CodeLens or the
	 * open-preview command; for a book, item:<id> and the page to open.
	 */
	| { type: 'target'; path: string | null; page?: number };

export type ToExtension =
	/** The page has loaded and wants the current state. */
	| { type: 'ready' }
	/** Show nothing (the shown declaration left the pack and the user cleared it). */
	| { type: 'clear' }
	/** Go to a declaration's source. */
	| { type: 'reveal'; uri: string; range: Range };

// The World view's page (webview/World.svelte). The extension owns the
// state and sends all of it after each change; the page only draws it and
// sends back what was clicked.

/** What the player does in the game after a build, from `mms build --world`. */
export type NextAction = 'reload' | 'rejoin' | 'none';

export type WorldResult =
	| { ok: true; files: number; at: string; next?: { action: NextAction; line: string; changed?: string }; resources: boolean }
	| { ok: false; message: string; detail: string; at: string; problems: boolean };

/** A structure a structure block saved in the world. */
export interface WorldStructureRow {
	/** Its .nbt file in the world, which names it in messages. */
	file: string;
	/** Its id in the game, such as tower:base. */
	id: string;
	status: 'new' | 'imported' | 'changed';
	/** Where it goes (or is) in the project, relative to the mms.json folder. */
	rel: string;
	call: string;
	saved: string;
	/** When it was saved, in ms since the epoch, for "saved 5 min ago". */
	savedMs: number;
	/** Imported, but no Structure(...) call names the project copy (lint/unused-asset). */
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
	result?: WorldResult;
	structures: WorldStructureRow[];
	/** Project .nbt files no Structure(...) call names that are not saves of the structure world. */
	unusedFiles: string[];
	/** The unused check waits for a project without errors. */
	unusedPaused: boolean;
}

export type ToWorld = { type: 'state'; state: WorldState };

export type FromWorld =
	| { type: 'ready' }
	| {
			type: 'run';
			action: 'build' | 'chooseWorld' | 'chooseProject' | 'revealWorld' | 'forgetWorld' | 'showProblems' | 'showOutput' | 'copyReload' | 'dismiss' | 'refresh' | 'importWaiting'
				| 'chooseStructureWorld' | 'sameStructureWorld' | 'revealStructureWorld';
	  }
	| { type: 'buildOnSave'; on: boolean }
	/** file is a WorldStructureRow's file. */
	| { type: 'import'; file: string }
	| { type: 'copyCall'; file: string }
	| { type: 'deleteFromWorld'; file: string }
	/** files are WorldStructureRow files, imported together (a folder's Import button). */
	| { type: 'importFiles'; files: string[] }
	/** rel is relative to the project folder. */
	| { type: 'revealFile'; rel: string };
