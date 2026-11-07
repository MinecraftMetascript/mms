/** A generated pack file, as the language server's `mms/output` returns it. */
export interface OutputFile {
	path: string;
	text: string;
	/** The source document and range the file came from. */
	uri?: string;
	range?: { start: { line: number; character: number }; end: { line: number; character: number } };
	/** Written for an inline value where an id is required; range is that value's. */
	implicit?: boolean;
	/** Labels for parts of the file by JSON pointer, from doc comments, such as "/elements/2" of a template pool. */
	labels?: Record<string, string>;
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

/**
 * A message a `tell`, `title` or `actionbar` command shows, from
 * `mms/output`'s texts: at the command, or over the .mmstext file it reads.
 */
export interface OutputText {
	/** text:<uri>:<offset>, the key an editor selects it by. */
	id: string;
	kind: 'chat' | 'title' | 'actionbar';
	/** The message as a text component, in JSON. */
	text: string;
	/** A title's subtitle, in JSON. */
	subtitle?: string;
	uri: string;
	range: SourceRange;
}

/** A jigsaw block of a structure file: where it is, and what it joins. */
export interface OutputJigsaw {
	pos: [number, number, number];
	pool: string;
	name: string;
	target: string;
	/** Its orientation state, such as north_up: the way it faces, then the way its top points. */
	orientation?: string;
	/** rollable or aligned: whether a piece joined to an up or down facing block may turn. */
	joint?: string;
	/** The block it turns into once the structure is placed, such as minecraft:air. */
	final_state?: string;
	/** The name of the item it was placed from, if it had one. */
	label?: string;
	/** Higher goes first among its piece's jigsaw blocks when the structure grows. */
	selection_priority?: number;
	/** Higher grows first among the pieces joined so far. */
	placement_priority?: number;
}

/** A project file the pack copies, such as a texture or a structure. */
export interface OutputAsset {
	/** Its path in the pack. */
	path: string;
	/** Its path in the project. */
	src: string;
	/** "resources" for a resource pack file. */
	pack?: string;
	/** For a structure file: its jigsaw blocks. */
	jigsaws?: OutputJigsaw[];
	/** For a structure file: its size in blocks, x y z. */
	size?: [number, number, number];
}
