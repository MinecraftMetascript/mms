// The preview panel: the web editor's worldgen previews and inspectors
// (packages/preview, built into out/webview) in a webview beside the
// editor. It follows the declaration at the cursor, like the web editor's
// Preview tab, using the generated pack from the server's `mms/output`.
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';
import type { OutputFile, ToExtension, ToPanel } from './messages';

export const viewType = 'mms.preview';

export class Previews implements vscode.Disposable {
	private panel: vscode.WebviewPanel | undefined;
	private files: OutputFile[] = [];
	/** The output at the cursor; kept while the cursor is outside declarations. */
	private cursor: string | null = null;
	private timer: ReturnType<typeof setTimeout> | undefined;
	private readonly subscriptions: vscode.Disposable[] = [];

	constructor(
		private readonly context: vscode.ExtensionContext,
		private readonly client: () => LanguageClient | undefined
	) {
		this.subscriptions.push(
			vscode.window.onDidChangeTextEditorSelection((e) => this.follow(e.textEditor)),
			vscode.window.onDidChangeActiveTextEditor((e) => e && this.follow(e)),
			vscode.window.registerWebviewPanelSerializer(viewType, {
				deserializeWebviewPanel: async (panel) => this.adopt(panel)
			})
		);
	}

	/** Opens the panel beside the editor, or brings it forward. */
	show() {
		if (this.panel) {
			this.panel.reveal(undefined, true);
			return;
		}
		const panel = vscode.window.createWebviewPanel(
			viewType,
			'MMS Preview',
			{ viewColumn: vscode.ViewColumn.Beside, preserveFocus: true },
			{}
		);
		this.adopt(panel);
	}

	/** Refreshes the generated pack shortly after the server publishes an analysis. */
	scheduleRefresh() {
		if (!this.panel) return;
		if (this.timer) clearTimeout(this.timer);
		this.timer = setTimeout(() => void this.refresh(), 150);
	}

	private adopt(panel: vscode.WebviewPanel) {
		this.panel = panel;
		const root = vscode.Uri.joinPath(this.context.extensionUri, 'out', 'webview');
		panel.iconPath = vscode.Uri.joinPath(this.context.extensionUri, 'icons', 'preview.svg');
		panel.webview.options = { enableScripts: true, localResourceRoots: [root] };
		panel.webview.html = this.html(panel.webview, root);
		panel.webview.onDidReceiveMessage((m: ToExtension) => this.receive(m), undefined, this.subscriptions);
		panel.onDidDispose(() => {
			if (this.panel === panel) this.panel = undefined;
		});
		const editor = vscode.window.activeTextEditor;
		if (editor) this.follow(editor);
	}

	private async receive(m: ToExtension) {
		if (m.type === 'ready') {
			this.post({ type: 'cursor', path: this.cursor });
			await this.refresh();
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
		if (!this.panel || !client?.isRunning()) return;
		try {
			const out = await client.sendRequest<{ files: OutputFile[]; hasErrors: boolean }>('mms/output', {});
			this.files = out.files;
			this.post({ type: 'output', files: out.files });
			const editor = vscode.window.activeTextEditor;
			if (editor) this.follow(editor);
		} catch {
			// A restart or shutdown; the next analysis refreshes again.
		}
	}

	/** Points the panel at the declaration under the editor's cursor. */
	private follow(editor: vscode.TextEditor) {
		if (!this.panel || editor.document.languageId !== 'mms') return;
		const line = editor.selection.active.line;
		const here = key(editor.document.uri);
		const f = this.files.find((f) => f.uri && f.range && f.range.start.line <= line && line <= f.range.end.line && key(vscode.Uri.parse(f.uri)) === here);
		if (!f || f.path === this.cursor) return;
		this.cursor = f.path;
		this.post({ type: 'cursor', path: f.path });
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

/** A file's identity: its path, ignoring case on Windows. */
function key(uri: vscode.Uri): string {
	return process.platform === 'win32' ? uri.fsPath.toLowerCase() : uri.fsPath;
}
