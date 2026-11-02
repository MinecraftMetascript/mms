// The preview panel: the web editor's worldgen previews and inspectors
// (packages/preview, built into out/webview) in a webview beside the
// editor. A "Preview" CodeLens above each declaration that has one opens
// the panel on it, and the panel stays on that declaration while you edit
// elsewhere, redrawing as the pack changes. The pack comes from the
// server's `mms/output`. A Book previews as its pages: its declaration
// has a lens, and so does each page of its .mmsbook file.
import { bookPageAt, itemPath, shows } from '@mms/preview/paths';
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';
import type { OutputFile, OutputItem, ToExtension, ToPanel } from './messages';

export const viewType = 'mms.preview';

export class Previews implements vscode.Disposable, vscode.CodeLensProvider {
	private panel: vscode.WebviewPanel | undefined;
	private files: OutputFile[] = [];
	private items: OutputItem[] = [];
	/** The output path the panel shows. */
	private target: string | null = null;
	/** The book page the panel opens on, from 0. */
	private page = 0;
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
			vscode.languages.registerCodeLensProvider([{ language: 'mms' }, { language: 'mmsbook' }], this),
			vscode.window.onDidChangeActiveTextEditor((e) => {
				if (!e || !isSource(e.document) || e.document.uri.toString() === this.document?.toString()) return;
				// Another file may belong to another project: fetch its pack.
				this.document = e.document.uri;
				this.scheduleRefresh();
			}),
			// A project with no problems publishes no diagnostics, so edits
			// refresh the pack (and the lenses) too.
			vscode.workspace.onDidOpenTextDocument((d) => isSource(d) && this.scheduleRefresh()),
			vscode.workspace.onDidChangeTextDocument((e) => isSource(e.document) && this.scheduleRefresh()),
			vscode.window.registerWebviewPanelSerializer(viewType, {
				deserializeWebviewPanel: async (panel, state?: { path?: string | null; page?: number }) => {
					this.target = state?.path ?? null;
					this.page = state?.page ?? 0;
					this.adopt(panel);
				}
			})
		);
	}

	/**
	 * Opens the panel on an output path (and a book on a page), or on the
	 * declaration at the cursor.
	 */
	show(path?: string, page?: number) {
		if (path === undefined) {
			const editor = vscode.window.activeTextEditor;
			const at = editor && this.at(editor.document.uri, editor.selection.active.line);
			if (at) [path, page] = [at.path, at.page];
			else path = this.target ?? undefined;
		}
		if (path !== undefined) this.setTarget(path, page);
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
		for (const it of this.items) {
			const path = itemPath(it.id);
			if (!shows(path, this.items)) continue;
			if (key(vscode.Uri.parse(it.uri)) === here && !lines.has(it.range.start.line)) {
				lines.add(it.range.start.line);
				lenses.push(
					new vscode.CodeLens(new vscode.Range(it.range.start.line, 0, it.range.start.line, 0), {
						title: '$(open-preview) Preview',
						tooltip: `Show the book ${it.id} in the MMS preview`,
						command: 'mms.previewDeclaration',
						arguments: [path, 0]
					})
				);
			}
			// Each page of a .mmsbook file opens the book on that page.
			it.pages?.forEach((p, i) => {
				if (p.uri === it.uri || key(vscode.Uri.parse(p.uri)) !== here || lines.has(p.range.start.line)) return;
				lines.add(p.range.start.line);
				lenses.push(
					new vscode.CodeLens(new vscode.Range(p.range.start.line, 0, p.range.start.line, 0), {
						title: `$(open-preview) Preview page ${i + 1}`,
						tooltip: `Show page ${i + 1} of ${it.id} in the MMS preview`,
						command: 'mms.previewDeclaration',
						arguments: [path, i]
					})
				);
			});
		}
		return lenses;
	}

	private setTarget(path: string | null, page = 0) {
		this.target = path;
		this.page = page;
		this.post({ type: 'target', path, page });
		if (this.panel) this.panel.title = title(path, this.items);
	}

	/** The output declared at a line of a document, if it has a preview. */
	private at(uri: vscode.Uri, line: number): { path: string; page?: number } | undefined {
		const here = key(uri);
		const f = this.files.find(
			(f) => f.uri && f.range && !f.implicit && f.range.start.line <= line && line <= f.range.end.line && shows(f.path) && key(vscode.Uri.parse(f.uri)) === here
		);
		if (f) return { path: f.path };
		for (const it of this.items) {
			const path = itemPath(it.id);
			if (!shows(path, this.items)) continue;
			const pages = it.pages?.filter((p) => key(vscode.Uri.parse(p.uri)) === here).map((p) => ({ ...p, uri: here }));
			const page = bookPageAt(pages, here, line);
			const inDecl = key(vscode.Uri.parse(it.uri)) === here && it.range.start.line <= line && line <= it.range.end.line;
			if (page !== null || inDecl) return { path, page: page ?? 0 };
		}
	}

	private adopt(panel: vscode.WebviewPanel) {
		this.panel = panel;
		const root = vscode.Uri.joinPath(this.context.extensionUri, 'out', 'webview');
		panel.title = title(this.target, this.items);
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
			this.post({ type: 'output', files: this.files, items: this.items });
			this.setTarget(this.target, this.page);
			await this.refresh();
		} else if (m.type === 'clear') {
			this.setTarget(null);
		} else if (m.type === 'reveal') {
			const doc = await vscode.workspace.openTextDocument(vscode.Uri.parse(m.uri));
			const range = new vscode.Range(m.range.start.line, m.range.start.character, m.range.end.line, m.range.end.character);
			// Open it where the source is being edited, not over the panel.
			const column = vscode.window.visibleTextEditors.find((e) => isSource(e.document))?.viewColumn ?? vscode.ViewColumn.One;
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
			const uri = (this.document ?? (isSource(active) ? active!.uri : undefined))?.toString();
			const out = await client.sendRequest<{ files: OutputFile[]; items?: OutputItem[]; hasErrors: boolean }>('mms/output', uri ? { uri } : {});
			this.files = out.files;
			this.items = out.items ?? [];
			this.lenses.fire();
			this.post({ type: 'output', files: out.files, items: this.items });
			if (this.panel) this.panel.title = title(this.target, this.items);
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
		// worldgen from mcmeta on GitHub; the book preview fetches the
		// game's font and textures from there too.
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
function title(path: string | null, items: OutputItem[]): string {
	const id = path ? shows(path, items)?.id : undefined;
	return id ? `Preview ${id.replace(/^[^:]+:/, '')}` : 'MMS Preview';
}

/** Whether a document is MMS source: a .mms file or a .mmsbook book text. */
function isSource(doc: vscode.TextDocument | undefined): boolean {
	return doc?.languageId === 'mms' || doc?.languageId === 'mmsbook';
}

/** A file's identity: its path, ignoring case on Windows. */
function key(uri: vscode.Uri): string {
	return process.platform === 'win32' ? uri.fsPath.toLowerCase() : uri.fsPath;
}
