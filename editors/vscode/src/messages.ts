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
}

export type ToPanel =
	/** The generated pack, refreshed after each analysis. */
	| { type: 'output'; files: OutputFile[] }
	/** The output path to show, chosen with a Preview CodeLens or the open-preview command. */
	| { type: 'target'; path: string | null };

export type ToExtension =
	/** The page has loaded and wants the current state. */
	| { type: 'ready' }
	/** Show nothing (the shown declaration left the pack and the user cleared it). */
	| { type: 'clear' }
	/** Go to a declaration's source. */
	| { type: 'reveal'; uri: string; range: Range };
