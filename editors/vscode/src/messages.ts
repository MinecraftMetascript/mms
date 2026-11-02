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
	jigsaws?: { pos: [number, number, number]; pool: string; name: string; target: string }[];
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
	| { type: 'output'; files: OutputFile[]; items: OutputItem[]; assets: OutputAsset[] }
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
