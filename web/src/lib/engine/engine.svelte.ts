import { LSPClient, languageServerExtensions } from '@codemirror/lsp-client';
import type { EditorView } from '@codemirror/view';
import { Supervisor, type EngineStatus } from './supervisor';
import { ProjectWorkspace } from './workspace';
import { sanitizeHTML } from './sanitize';
import { showDiagnostics } from '../editor/diagnostics';
import type { InlayHint } from '../editor/inlay';
import type { OutputJigsaw } from '@mms/preview';

export interface Diagnostic {
	range: { start: { line: number; character: number }; end: { line: number; character: number } };
	severity?: number;
	message: string;
	code?: string;
	/** 1 marks unnecessary code. */
	tags?: number[];
	data?: { fixes?: Fix[] };
	/** Other places that explain the problem, such as the Template line that wrote it. */
	relatedInformation?: { location: { uri: string; range: Diagnostic['range'] }; message: string }[];
}

export interface TextEdit {
	range: Diagnostic['range'];
	newText: string;
}

/**
 * A refactoring such as "Extract { ... } to a Noise declaration": its
 * edits, and the default name it picked (renamed after applying).
 */
export interface Refactor {
	title: string;
	edits: TextEdit[];
	name?: string;
}

/** A fix the engine attached to a diagnostic; safe fixes never change the pack. */
export interface Fix {
	title: string;
	edits: TextEdit[];
	safe: boolean;
}

export interface OutputFile {
	path: string;
	text: string;
	/** "resources" for a resource pack file. */
	pack?: string;
	uri?: string;
	range?: Diagnostic['range'];
	/** Written for an inline value where an id is required; range is that value's. */
	implicit?: boolean;
	/** Labels for parts of the file by JSON pointer, from doc comments, such as "/elements/2" of a template pool. */
	labels?: Record<string, string>;
	/** Set when this is a custom item's stack shown as JSON, not a pack file. */
	item?: OutputItem;
}

/**
 * A CustomItem: the item stack its name stands for. It is not a file of
 * its own, so it is listed apart from the pack's files.
 */
export interface OutputItem {
	id: string;
	/** The stack as JSON: {id, components}. */
	text: string;
	/** The stack as a command writes it, such as stick[max_stack_size=1]. */
	snbt: string;
	uri: string;
	range: Diagnostic['range'];
	/** For a Book: where each page of its text starts, maybe in a .mmsbook file. */
	pages?: { uri: string; range: Diagnostic['range'] }[];
}

/** A message tell, title or actionbar shows, at the command or over the .mmstext file it reads. */
/** What a Template call or Blueprint instance writes, as `mms/expand` returns it. */
export interface Expansion {
	title: string;
	text: string;
	range: Diagnostic['range'];
}

/** A world build as `mms/worldInstall` works it out; bytes are base64. */
export interface WorldInstall {
	namespace: string;
	/** The data pack, for datapacks/<namespace>. */
	data: { path: string; data: string }[];
	resources: number;
	dimensions: string[];
	/** Dimensions whose saved chunks to take out. */
	reset: string[];
	next: 'reload' | 'rejoin' | 'none';
	/** The "next step: ..." line. */
	message: string;
	/** The world's new resourcepacks/resources.zip; absent when the build leaves it alone. */
	zip?: string;
	resourcesChanged?: boolean;
	others?: string[];
	replaced?: string[];
	adopted?: boolean;
}

/** One Test's outcome, as `mms/runTests` returns it. */
export interface TestOutcome {
	name: string;
	uri: string;
	range: Diagnostic['range'];
	pass: boolean;
	message?: string;
	/** The failing expect, from 1, and where it is. */
	expect?: number;
	expectRange?: Diagnostic['range'];
	said: string[];
}

const testLine = /^[ \t]*Test[ \t]+[A-Za-z_]/m;

export interface OutputText {
	/** text:<uri>:<offset>, the key the editor selects it by. */
	id: string;
	kind: 'chat' | 'title' | 'actionbar';
	/** The message as a text component, in JSON. */
	text: string;
	subtitle?: string;
	uri: string;
	range: Diagnostic['range'];
}

/**
 * The key the editor selects an output by. A data pack file keeps its
 * path; a resource pack file is under resources/, since both packs have a
 * pack.mcmeta; a custom item is item:<id>.
 */
export const outputKey = (f: OutputFile) => (f.pack === 'resources' ? `resources/${f.path}` : f.path);
export const itemKey = (it: OutputItem) => `item:${it.id}`;

/** A project file the pack copies, such as a texture: its output path and source. */
/** A `JigsawHole name` declaration. */
export interface OutputHole {
	id: string;
	doc?: string;
	uri: string;
	range: Diagnostic['range'];
}

export interface OutputAsset {
	path: string;
	src: string;
	pack?: string;
	/** For a structure file: its jigsaw blocks. */
	jigsaws?: OutputJigsaw[];
	/** For a structure file: its size in blocks, x y z. */
	size?: [number, number, number];
}

export interface DocSymbol {
	name: string;
	detail: string;
	kind: number;
	range: Diagnostic['range'];
	selectionRange: Diagnostic['range'];
}

export interface PackImport {
	files: { path: string; text: string }[];
	namespace: string;
	description: string;
	skipped: string[];
	count: number;
}

export interface ResourceKind {
	kind: string;
	folder: string;
}

export interface EngineHooks {
	onText(uri: string, text: string): void;
	display(uri: string): Promise<EditorView | null>;
}

/**
 * Engine is the page's view of the language server: the LSP client and
 * workspace, the worker supervisor, and reactive copies of what the
 * sidebar and previews show (diagnostics, the generated pack, symbols).
 */
export class Engine {
	status = $state<EngineStatus>('starting');
	statusDetail = $state('');
	diagnostics = $state<Record<string, Diagnostic[]>>({});
	output = $state<OutputFile[]>([]);
	outputAssets = $state<OutputAsset[]>([]);
	/** The project's JigsawHole declarations: jigsaw block names, which write no file. */
	outputHoles = $state<OutputHole[]>([]);
	outputItems = $state<OutputItem[]>([]);
	/** The messages tell, title and actionbar show, for the chat preview. */
	outputTexts = $state<OutputText[]>([]);
	/** Project files nothing uses, so the packs leave them out. */
	unusedAssets = $state<string[]>([]);
	hasErrors = $state(false);
	/** Whether the project declares a Test; the results of running them after each analysis; and why none ran. */
	hasTests = $state(false);
	tests = $state<TestOutcome[]>([]);
	testsError = $state('');
	/** Why the server refused the project settings, such as a bad namespace. */
	configError = $state('');
	/** Settings the server took but doubts, such as a lint rule nothing has. */
	configWarnings = $state<string[]>([]);
	symbols = $state<Record<string, DocSymbol[]>>({});

	readonly client: LSPClient;
	readonly workspace!: ProjectWorkspace;
	private supervisor: Supervisor;
	private config: Record<string, unknown> = { namespace: 'mms', description: '' };
	private refreshTimer: ReturnType<typeof setTimeout> | null = null;
	/** The binary files the server holds, path to base64. */
	private assets = new Map<string, string>();

	constructor(workerURL: string, hooks: EngineHooks) {
		let ws: ProjectWorkspace | null = null;
		this.client = new LSPClient({
			timeout: 10000,
			// Docs come from project files, which may be imported from anyone.
			sanitizeHTML,
			workspace: (client) => (ws = new ProjectWorkspace(client, hooks)),
			extensions: languageServerExtensions(),
			notificationHandlers: {
				'textDocument/publishDiagnostics': (client, params: { uri: string; diagnostics: Diagnostic[] }) => {
					this.diagnostics[params.uri] = params.diagnostics;
					this.scheduleRefresh();
					// Shown here rather than by the client so fixes become actions.
					const view = client.workspace.getFile(params.uri)?.getView();
					if (view) showDiagnostics(view, params.diagnostics);
					return true;
				}
			}
		});
		(this as { workspace: ProjectWorkspace }).workspace = ws!;
		this.supervisor = new Supervisor({
			workerURL,
			onStatus: (s, detail) => {
				this.status = s;
				this.statusDetail = detail ?? '';
			},
			onRestart: () => this.reconnect()
		});
		this.connect();
	}

	private connect() {
		this.client.connect(this.supervisor.transport);
		this.client.initializing.then(() => {
			this.sendConfig();
			if (this.assets.size) {
				const files = [...this.assets].map(([path, data]) => ({ path, data }));
				this.client.notification('mms/setAssets', { files });
			}
			this.scheduleRefresh();
		});
	}

	private reconnect() {
		this.client.disconnect();
		this.diagnostics = {};
		this.connect();
	}

	restart() {
		this.supervisor.restart();
	}

	/** Sends the settings; the server checks them and says what is wrong. */
	private sendConfig() {
		this.client.request<unknown, { warnings?: string[] } | null>('mms/setConfig', this.config).then(
			(r) => {
				this.configError = '';
				this.configWarnings = r?.warnings ?? [];
			},
			(e: unknown) => (this.configError = e instanceof Error ? e.message : String(e))
		);
	}

	/** Sends the whole mms.json, lint settings and all, as VS Code's server reads it from disk. */
	setConfig(config: { namespace: string; description: string; rest?: Record<string, unknown> }) {
		const { rest, ...known } = config;
		this.config = { ...rest, ...known };
		if (this.client.connected) this.sendConfig();
	}

	/** Replaces the binary files the server reads, sending only changes. */
	setAssets(files: { path: string; data: string }[]) {
		const want = new Map(files.map((f) => [f.path, f.data]));
		const changed = files.filter((f) => this.assets.get(f.path) !== f.data);
		const remove = [...this.assets.keys()].filter((p) => !want.has(p));
		this.assets = want;
		if ((changed.length || remove.length) && this.client.connected) {
			this.client.notification('mms/setAssets', { files: changed, remove });
		}
	}

	/** Refreshes the generated pack and symbols shortly after analysis. */
	scheduleRefresh() {
		if (this.refreshTimer) clearTimeout(this.refreshTimer);
		this.refreshTimer = setTimeout(() => this.refresh(), 120);
	}

	private async refresh() {
		if (!this.client.connected) return;
		try {
			await this.client.initializing;
			this.client.sync();
			const out = await this.client.request<object, { files: OutputFile[]; assets?: OutputAsset[]; jigsawHoles?: OutputHole[]; items?: OutputItem[]; texts?: OutputText[]; unusedAssets?: string[]; hasErrors: boolean }>('mms/output', {});
			this.output = out.files;
			this.outputAssets = out.assets ?? [];
			this.outputHoles = out.jigsawHoles ?? [];
			this.outputItems = out.items ?? [];
			this.outputTexts = out.texts ?? [];
			this.unusedAssets = out.unusedAssets ?? [];
			this.hasErrors = out.hasErrors;
			// Tests run in the same interpreter as `mms test`, after every
			// analysis, so their results stay current as the code changes.
			this.hasTests = this.workspace.files.some((f) => f.languageId === 'mms' && testLine.test(f.doc.toString()));
			if (this.hasTests) {
				const res = await this.client.request<object, { tests: TestOutcome[]; error?: string }>('mms/runTests', {});
				this.tests = res.tests;
				this.testsError = res.error ?? '';
			} else {
				this.tests = [];
				this.testsError = '';
			}
			const syms: Record<string, DocSymbol[]> = {};
			await Promise.all(
				this.workspace.files.map(async (f) => {
					syms[f.uri] =
						(await this.client.request<object, DocSymbol[] | null>('textDocument/documentSymbol', {
							textDocument: { uri: f.uri }
						})) ?? [];
				})
			);
			this.symbols = syms;
		} catch {
			// A restart or timeout; the next analysis refreshes again.
		}
	}

	/** Converts data pack files into formatted MMS files. */
	async importPack(files: { path: string; text: string }[]): Promise<PackImport> {
		await this.client.initializing;
		return this.client.request<object, PackImport>('mms/importPack', { files });
	}

	/**
	 * What building the pack into a world writes there, worked out from the
	 * world's data pack folder (before, by path, base64) and its
	 * resources.zip (base64, null when there is none). Rejects when the
	 * project has errors.
	 */
	async worldInstall(p: { before: Record<string, string>; oldZip: string | null; reset: boolean; resetIds?: string[] }): Promise<WorldInstall> {
		await this.client.initializing;
		return this.client.request<object, WorldInstall>('mms/worldInstall', p);
	}

	/** Converts one JSON resource into a declaration for the file uri. */
	async convertJSON(uri: string, kind: string, name: string, text: string): Promise<string> {
		await this.client.initializing;
		const res = await this.client.request<object, { text: string }>('mms/convertJSON', { uri, kind, name, text });
		return res.text;
	}

	/** The JSON resource kinds and their data pack folders. */
	async resourceKinds(): Promise<ResourceKind[]> {
		await this.client.initializing;
		return this.client.request<object, ResourceKind[]>('mms/resourceKinds', {});
	}

	/** The text of a file of the language's prelude, by its mms-prelude: URI. */
	async preludeFile(uri: string): Promise<string | null> {
		await this.client.initializing;
		try {
			return (await this.client.request<object, { text: string }>('mms/preludeFile', { uri })).text;
		} catch {
			return null;
		}
	}

	/** What the Template call or Blueprint instance at a position writes, or null. */
	async expand(uri: string, position: { line: number; character: number }): Promise<Expansion | null> {
		if (!this.client.connected) return null;
		this.client.sync();
		try {
			return await this.client.request<object, Expansion | null>('mms/expand', { textDocument: { uri }, position });
		} catch {
			return null;
		}
	}

	/** Refactorings for a range of a file, such as extracting an inline value to a declaration. */
	async refactors(uri: string, range: Diagnostic['range']): Promise<Refactor[]> {
		this.client.sync();
		const actions = await this.client.request<
			object,
			{ title: string; edit?: { changes?: Record<string, TextEdit[]> }; data?: { rename?: { name: string } } }[]
		>('textDocument/codeAction', { textDocument: { uri }, range, context: { diagnostics: [], only: ['refactor'] } });
		return (actions ?? []).map((a) => ({ title: a.title, edits: a.edit?.changes?.[uri] ?? [], name: a.data?.rename?.name }));
	}

	/** The edits that fix every safely fixable problem in a file. */
	async fixAll(uri: string): Promise<TextEdit[]> {
		const actions = await this.client.request<object, { edit?: { changes?: Record<string, TextEdit[]> } }[]>('textDocument/codeAction', {
			textDocument: { uri },
			range: { start: { line: 0, character: 0 }, end: { line: 0, character: 0 } },
			context: { diagnostics: [], only: ['source.fixAll'] }
		});
		return actions[0]?.edit?.changes?.[uri] ?? [];
	}

	async format(uri: string): Promise<string | null> {
		this.client.sync();
		const edits = await this.client.request<object, { newText: string }[] | null>('textDocument/formatting', {
			textDocument: { uri },
			options: { tabSize: 2, insertSpaces: true }
		});
		return edits && edits.length ? edits[0].newText : null;
	}

	/** Semantic tokens for a document, in LSP's relative encoding. */
	async semanticTokens(uri: string): Promise<number[]> {
		this.client.sync();
		const res = await this.client.request<object, { data: number[] } | null>('textDocument/semanticTokens/full', {
			textDocument: { uri }
		});
		return res?.data ?? [];
	}

	/** Inlay hints for a document, such as a custom item's base item. */
	async inlayHints(uri: string): Promise<InlayHint[]> {
		this.client.sync();
		const res = await this.client.request<object, InlayHint[] | null>('textDocument/inlayHint', {
			textDocument: { uri },
			range: { start: { line: 0, character: 0 }, end: { line: 1 << 30, character: 0 } }
		});
		return res ?? [];
	}

	dispose() {
		this.client.disconnect();
		this.supervisor.dispose();
	}
}

export const tokenTypes = [
	'keyword',
	'comment',
	'string',
	'number',
	'operator',
	'type',
	'variable',
	'function',
	'method',
	'property',
	'enumMember',
	'namespace',
	'macro',
	'parameter'
];
