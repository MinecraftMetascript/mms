import { LSPClient, languageServerExtensions } from '@codemirror/lsp-client';
import type { EditorView } from '@codemirror/view';
import { Supervisor, type EngineStatus } from './supervisor';
import { ProjectWorkspace } from './workspace';

export interface Diagnostic {
	range: { start: { line: number; character: number }; end: { line: number; character: number } };
	severity?: number;
	message: string;
	code?: string;
}

export interface OutputFile {
	path: string;
	text: string;
	uri?: string;
	range?: Diagnostic['range'];
}

export interface DocSymbol {
	name: string;
	detail: string;
	kind: number;
	range: Diagnostic['range'];
	selectionRange: Diagnostic['range'];
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
	hasErrors = $state(false);
	symbols = $state<Record<string, DocSymbol[]>>({});

	readonly client: LSPClient;
	readonly workspace!: ProjectWorkspace;
	private supervisor: Supervisor;
	private config: { namespace: string; description: string } = { namespace: 'mms', description: '' };
	private refreshTimer: ReturnType<typeof setTimeout> | null = null;

	constructor(workerURL: string, hooks: EngineHooks) {
		let ws: ProjectWorkspace | null = null;
		this.client = new LSPClient({
			timeout: 10000,
			workspace: (client) => (ws = new ProjectWorkspace(client, hooks)),
			extensions: languageServerExtensions(),
			notificationHandlers: {
				'textDocument/publishDiagnostics': (_client, params: { uri: string; diagnostics: Diagnostic[] }) => {
					this.diagnostics[params.uri] = params.diagnostics;
					this.scheduleRefresh();
					return false; // let the editor extension show them too
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
			const out = await this.client.request<object, { files: OutputFile[]; hasErrors: boolean }>('mms/output', {});
			this.output = out.files;
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
