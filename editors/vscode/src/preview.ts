// The preview panel: the web editor's worldgen previews and inspectors
// (packages/preview, built into out/webview) in a webview beside the
// editor. A "Preview" CodeLens above each declaration that has one opens
// the panel on it, and the panel stays on that declaration while you edit
// elsewhere, redrawing as the pack changes. The pack comes from the
// server's `mms/output` (output.ts). A Book previews as its pages: its declaration
// has a lens, and so does each page of its .mmsbook file.
import { bookPageAt, itemPath, shows } from '@mms/preview/paths';
import * as vscode from 'vscode';
import type { OutputItem, ToExtension, ToPanel } from './messages';
import { isSource, key, type PackOutput } from './output';

export const viewType = 'mms.preview';

export class Previews implements vscode.Disposable, vscode.CodeLensProvider {
	private panel: vscode.WebviewPanel | undefined;
	/** The output path the panel shows. */
	private target: string | null = null;
	/** The book page the panel opens on, from 0. */
	private page = 0;
	private readonly lenses = new vscode.EventEmitter<void>();
	readonly onDidChangeCodeLenses = this.lenses.event;
	private readonly subscriptions: vscode.Disposable[] = [];

	constructor(
		private readonly context: vscode.ExtensionContext,
		private readonly output: PackOutput
	) {
		this.subscriptions.push(
			this.lenses,
			vscode.languages.registerCodeLensProvider([{ language: 'mms' }, { language: 'mmsbook' }], this),
			output.onDidChange(() => {
				this.lenses.fire();
				this.post({ type: 'output', files: this.files, items: this.items, assets: this.output.assets });
				if (this.panel) this.panel.title = title(this.target, this.items);
			}),
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

	private get files() {
		return this.output.files;
	}
	private get items() {
		return this.output.items;
	}

	provideCodeLenses(document: vscode.TextDocument): vscode.CodeLens[] {
		const here = key(document.uri);
		const lenses: vscode.CodeLens[] = [];
		const lines = new Set<number>();
		for (const f of this.files) {
			if (!f.uri || !f.range || f.implicit || key(vscode.Uri.parse(f.uri)) !== here) continue;
			const what = shows(f.path);
			if (!what || lines.has(f.range.start.line)) continue;
			// Only a jigsaw structure has pieces to connect.
			if (what.kind === 'structure' && !/"type"\s*:\s*"(minecraft:)?jigsaw"/.test(f.text)) continue;
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
			this.post({ type: 'output', files: this.files, items: this.items, assets: this.output.assets });
			this.setTarget(this.target, this.page);
			await this.output.refresh();
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
		this.panel?.dispose();
		for (const s of this.subscriptions) s.dispose();
	}
}

/** The panel's tab title: the declaration it shows. */
function title(path: string | null, items: OutputItem[]): string {
	const id = path ? shows(path, items)?.id : undefined;
	return id ? `Preview ${id.replace(/^[^:]+:/, '')}` : 'MMS Preview';
}
