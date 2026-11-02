/** A generated pack file, as the language server's `mms/output` returns it. */
export interface OutputFile {
	path: string;
	text: string;
	/** The source document and range the file came from. */
	uri?: string;
	range?: { start: { line: number; character: number }; end: { line: number; character: number } };
	/** Written for an inline value where an id is required; range is that value's. */
	implicit?: boolean;
}

/** A source position, as the language server sends it. */
export interface SourceRange {
	start: { line: number; character: number };
	end: { line: number; character: number };
}

/**
 * A CustomItem or Book from `mms/output`'s items: the item stack its name
 * stands for, which is not a file of its own.
 */
export interface OutputItem {
	id: string;
	/** The stack as JSON: {id, components}. */
	text: string;
	uri: string;
	range: SourceRange;
	/** For a Book: where each page of its text starts. */
	pages?: { uri: string; range: SourceRange }[];
}
