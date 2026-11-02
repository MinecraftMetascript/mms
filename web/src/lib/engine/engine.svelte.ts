import { LSPClient, languageServerExtensions } from '@codemirror/lsp-client';
import type { EditorView } from '@codemirror/view';
import { Supervisor, type EngineStatus } from './supervisor';
import { ProjectWorkspace } from './workspace';
import { showDiagnostics } from '../editor/diagnostics';

export interface Diagnostic {
	range: { start: { line: number; character: number }; end: { line: number; character: number } };
	severity?: number;
	message: string;
	code?: string;
	/** 1 marks unnecessary code. */
	tags?: number[];
	data?: { fixes?: Fix[] };
}

export interface TextEdit {
	range: Diagnostic['range'];
	newText: string;
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
}

/** A project file the pack copies, such as a texture: its output path and source. */
export interface OutputAsset {
	path: string;
	src: string;
	pack?: string;
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
	hasErrors = $state(false);
	symbols = $state<Record<string, DocSymbol[]>>({});

	readonly client: LSPClient;
	readonly workspace!: ProjectWorkspace;
	private supervisor: Supervisor;
	private config: { namespace: string; description: string } = { namespace: 'mms', description: '' };
	private refreshTimer: ReturnType<typeof setTimeout> | null = null;
	/** The binary files the server holds, path to base64. */
	private assets = new Map<string, string>();

	constructor(workerURL: string, hooks: EngineHooks) {
		let ws: ProjectWorkspace | null = null;
		this.client = new LSPClient({
			timeout: 10000,
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
			this.client.notification('mms/setConfig', this.config);
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

	setConfig(config: { namespace: string; description: string }) {
		this.config = { ...config };
		if (this.client.connected) this.client.notification('mms/setConfig', this.config);
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
			const out = await this.client.request<object, { files: OutputFile[]; assets?: OutputAsset[]; hasErrors: boolean }>('mms/output', {});
			this.output = out.files;
			this.outputAssets = out.assets ?? [];
			this.hasErrors = out.hasErrors;
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

	/** Formats a whole document, returning the new text or null. */
	/** Converts data pack files into formatted MMS files. */
	async importPack(files: { path: string; text: string }[]): Promise<PackImport> {
		await this.client.initializing;
		return this.client.request<object, PackImport>('mms/importPack', { files });
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
	'macro'
];
