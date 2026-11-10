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

import type { SplineAt, TextEdit as SplineEdit } from '@mms/preview/spline';
export type { SplineAt, SplineEdit };

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
	| { type: 'target'; path: string | null; page?: number }
	/** The density function spline at the cursor (`mms/spline`), or null, for the spline editor. */
	| { type: 'spline'; uri: string | null; spline: SplineAt | null };

export type ToExtension =
	/** The page has loaded and wants the current state. */
	| { type: 'ready' }
	/** Show nothing (the shown declaration left the pack and the user cleared it). */
	| { type: 'clear' }
	/** Go to a declaration's source. */
	| { type: 'reveal'; uri: string; range: Range }
	/** Apply the spline editor's edits to a source file. */
	| { type: 'splineEdit'; uri: string; edits: SplineEdit[] };

// The World and Structures views' page (webview/World.svelte, drawing the
// shared WorldTool). The extension owns the state and sends all of it after
// each change; the page only draws it and sends back what was clicked.

import type { NextAction, WorldAction, WorldDimensionRow, WorldResult, WorldState, WorldStructureRow } from '@mms/preview/world/tool';
export type { NextAction, WorldDimensionRow, WorldResult, WorldState, WorldStructureRow };

export type ToWorld = { type: 'state'; state: WorldState };

export type FromWorld = { type: 'ready' } | WorldAction;

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
