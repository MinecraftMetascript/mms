/** A generated pack file, as the language server's `mms/output` returns it. */
export interface OutputFile {
	path: string;
	text: string;
	/** The source document and range the file came from. */
	uri?: string;
	range?: { start: { line: number; character: number }; end: { line: number; character: number } };
}
