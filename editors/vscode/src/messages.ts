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
	jigsaws?: { pos: [number, number, number]; pool: string; name: string; target: string; orientation?: string; joint?: string; final_state?: string; label?: string; selection_priority?: number; placement_priority?: number }[];
	/** For a structure file: its size in blocks, x y z. */
	size?: [number, number, number];
}

/** A `JigsawHole name` declaration: a jigsaw block name, which writes no file. */
export interface OutputHole {
	id: string;
	doc?: string;
	uri: string;
	range: Range;
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

/** A message a tell, title or actionbar command shows; matches packages/preview's OutputText. */
export interface OutputText {
	/** text:<uri>:<offset>, the key the preview selects it by. */
	id: string;
	kind: 'chat' | 'title' | 'actionbar';
	text: string;
	subtitle?: string;
	uri: string;
	range: Range;
}

export type ToPanel =
	/** The generated pack and custom items, refreshed after each analysis. */
	| {
			type: 'output';
			files: OutputFile[];
			items: OutputItem[];
			texts: OutputText[];
			assets: OutputAsset[];
			/** The pack's copied .png textures as data URLs, by pack path: custom item icons. */
			textures?: Record<string, string>;
			/** The pack's structure files as data URLs, by pack path: blocks for the jigsaw simulation. */
			structures?: Record<string, string>;
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

// The World and Structures views' page (webview/World.svelte). The extension owns the
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
	status: 'new' | 'imported' | 'changed';
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

export type ToWorld = { type: 'state'; state: WorldState };

export type FromWorld =
	| { type: 'ready' }
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

// The Jigsaw view's page (webview/Jigsaw.svelte). The page owns the form;
// the extension sends the lists to choose from and the recent blocks, and
// copies the give command the page builds.

import type { JigsawConfig } from '@mms/preview/jigsaw/give';
import type { JigsawAction, JigsawExisting, JigsawHoleRow, JigsawPoolRow, JigsawRecent, JigsawState } from '@mms/preview/jigsaw/tool';
export type { JigsawConfig, JigsawExisting, JigsawHoleRow, JigsawPoolRow, JigsawRecent, JigsawState };

export type ToJigsaw = { type: 'state'; state: JigsawState };

export type FromJigsaw = { type: 'ready' } | JigsawAction;

/** One Test's outcome, as the language server's `mms/runTests` returns it. */
export interface TestOutcome {
	name: string;
	uri: string;
	/** The Test's name. */
	range: Range;
	pass: boolean;
	message?: string;
	/** The failing expect, from 1, and where it is. */
	expect?: number;
	expectRange?: Range;
	/** What the test's commands said, through say and tellraw. */
	said: string[];
}
