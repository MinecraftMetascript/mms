// The preview panel: the web editor's worldgen previews and inspectors
// (packages/preview, built into out/webview) in a webview beside the
// editor. A "Preview" CodeLens above each declaration that has one opens
// the panel on it, and the panel stays on that declaration while you edit
// elsewhere, redrawing as the pack changes. The pack comes from the
// server's `mms/output`.
import { shows } from '@mms/preview/paths';
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';
import type { OutputFile, ToExtension, ToPanel } from './messages';

export const viewType = 'mms.preview';

export class Previews implements vscode.Disposable, vscode.CodeLensProvider {
	private panel: vscode.WebviewPanel | undefined;
	private files: OutputFile[] = [];
	/** The output path the panel shows. */
	private target: string | null = null;
	/** The .mms document last active, whose project's pack is fetched. */
	private document: vscode.Uri | undefined;
	private timer: ReturnType<typeof setTimeout> | undefined;
	private readonly lenses = new vscode.EventEmitter<void>();
	readonly onDidChangeCodeLenses = this.lenses.event;
	private readonly subscriptions: vscode.Disposable[] = [];

	constructor(
		private readonly context: vscode.ExtensionContext,
		private readonly client: () => LanguageClient | undefined
	) {
		this.subscriptions.push(
			this.lenses,
			vscode.languages.registerCodeLensProvider({ language: 'mms' }, this),
			vscode.window.onDidChangeActiveTextEditor((e) => {
				if (e?.document.languageId !== 'mms' || e.document.uri.toString() === this.document?.toString()) return;
				// Another file may belong to another project: fetch its pack.
				this.document = e.document.uri;
				this.scheduleRefresh();
			}),
			vscode.window.registerWebviewPanelSerializer(viewType, {
				deserializeWebviewPanel: async (panel, state?: { path?: string | null }) => {
					this.target = state?.path ?? null;
					this.adopt(panel);
				}
			})
		);
	}

	/** Opens the panel on an output path, or on the declaration at the cursor. */
	show(path?: string) {
		if (path === undefined) {
			const editor = vscode.window.activeTextEditor;
			path = (editor && this.at(editor.document.uri, editor.selection.active.line)?.path) ?? this.target ?? undefined;
		}
		if (path !== undefined) this.setTarget(path);
		if (this.panel) {
			this.panel.reveal(undefined, true);
			return;
		}
		const panel = vscode.window.createWebviewPanel(viewType, 'MMS Preview', { viewColumn: vscode.ViewColumn.Beside, preserveFocus: true }, {});
		this.adopt(panel);
	}

	/** Refreshes the generated pack shortly after the server publishes an analysis. */
	scheduleRefresh() {
		if (this.timer) clearTimeout(this.timer);
		this.timer = setTimeout(() => void this.refresh(), 150);
	}

	provideCodeLenses(document: vscode.TextDocument): vscode.CodeLens[] {
		const here = key(document.uri);
		const lenses: vscode.CodeLens[] = [];
		const lines = new Set<number>();
		for (const f of this.files) {
			if (!f.uri || !f.range || f.implicit || key(vscode.Uri.parse(f.uri)) !== here) continue;
			const what = shows(f.path);
			if (!what || lines.has(f.range.start.line)) continue;
			lines.add(f.range.start.line);
			const at = new vscode.Range(f.range.start.line, 0, f.range.start.line, 0);
			lenses.push(
				new vscode.CodeLens(at, {
					title: '$(open-preview) Preview',
					tooltip: `Show ${what.id} in the MMS preview`,
					command: 'mms.previewDeclaration',
					arguments: [f.path]
				})
			);
		}
		return lenses;
	}

	private setTarget(path: string | null) {
		this.target = path;
		this.post({ type: 'target', path });
		if (this.panel) this.panel.title = title(path);
	}

	/** The output declared at a line of a document, if it has a preview. */
	private at(uri: vscode.Uri, line: number): OutputFile | undefined {
		const here = key(uri);
		return this.files.find(
			(f) => f.uri && f.range && !f.implicit && f.range.start.line <= line && line <= f.range.end.line && shows(f.path) && key(vscode.Uri.parse(f.uri)) === here
		);
	}

	private adopt(panel: vscode.WebviewPanel) {
		this.panel = panel;
		const root = vscode.Uri.joinPath(this.context.extensionUri, 'out', 'webview');
		panel.title = title(this.target);
		panel.iconPath = vscode.Uri.joinPath(this.context.extensionUri, 'icons', 'preview.svg');
		panel.webview.options = { enableScripts: true, localResourceRoots: [root] };
		panel.webview.html = this.html(panel.webview, root);
		panel.webview.onDidReceiveMessage((m: ToExtension) => this.receive(m), undefined, this.subscriptions);
		panel.onDidDispose(() => {
			if (this.panel === panel) this.panel = undefined;
		});
	}

	private async receive(m: ToExtension) {
		if (m.type === 'ready') {
			this.post({ type: 'output', files: this.files });
			this.setTarget(this.target);
			await this.refresh();
		} else if (m.type === 'clear') {
			this.setTarget(null);
		} else if (m.type === 'reveal') {
			const doc = await vscode.workspace.openTextDocument(vscode.Uri.parse(m.uri));
			const range = new vscode.Range(m.range.start.line, m.range.start.character, m.range.end.line, m.range.end.character);
			// Open it where the source is being edited, not over the panel.
			const column = vscode.window.visibleTextEditors.find((e) => e.document.languageId === 'mms')?.viewColumn ?? vscode.ViewColumn.One;
			const editor = await vscode.window.showTextDocument(doc, { viewColumn: column, selection: new vscode.Range(range.start, range.start) });
			editor.revealRange(range, vscode.TextEditorRevealType.InCenterIfOutsideViewport);
		}
	}

	private async refresh() {
		const client = this.client();
		if (!client?.isRunning()) return;
		try {
			// The uri picks the project (the nearest mms.json) in a workspace
			// with several; servers that predate that ignore it.
			const active = vscode.window.activeTextEditor?.document;
			const uri = (this.document ?? (active?.languageId === 'mms' ? active.uri : undefined))?.toString();
			const out = await client.sendRequest<{ files: OutputFile[]; hasErrors: boolean }>('mms/output', uri ? { uri } : {});
			this.files = out.files;
			this.lenses.fire();
			this.post({ type: 'output', files: out.files });
		} catch {
			// A restart or shutdown; the next analysis refreshes again.
		}
	}

	private post(m: ToPanel) {
		void this.panel?.webview.postMessage(m);
	}

	private html(webview: vscode.Webview, root: vscode.Uri): string {
		const nonce = Array.from({ length: 32 }, () => Math.floor(Math.random() * 36).toString(36)).join('');
		const script = webview.asWebviewUri(vscode.Uri.joinPath(root, 'preview.js'));
		const style = webview.asWebviewUri(vscode.Uri.joinPath(root, 'preview.css'));
		// The preview worker starts from a blob URL and fetches vanilla
		// worldgen from mcmeta on GitHub.
		const csp = [
			"default-src 'none'",
			`style-src ${webview.cspSource} 'unsafe-inline'`,
			`script-src 'nonce-${nonce}'`,
			'worker-src blob:',
			'connect-src https://raw.githubusercontent.com',
			`img-src ${webview.cspSource} data: blob:`,
			`font-src ${webview.cspSource}`
		].join('; ');
		return `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta http-equiv="Content-Security-Policy" content="${csp}">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="${style}">
<title>MMS Preview</title>
</head>
<body>
<div id="app"></div>
<script type="module" nonce="${nonce}" src="${script}"></script>
</body>
</html>`;
	}

	dispose() {
		if (this.timer) clearTimeout(this.timer);
		this.panel?.dispose();
		for (const s of this.subscriptions) s.dispose();
	}
}

/** The panel's tab title: the declaration it shows. */
function title(path: string | null): string {
	const id = path ? shows(path)?.id : undefined;
	return id ? `Preview ${id.replace(/^[^:]+:/, '')}` : 'MMS Preview';
}

/** A file's identity: its path, ignoring case on Windows. */
function key(uri: vscode.Uri): string {
	return process.platform === 'win32' ? uri.fsPath.toLowerCase() : uri.fsPath;
}
