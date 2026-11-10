// The preview panel: the web editor's worldgen previews and inspectors
// (packages/preview, built into out/webview) in a webview beside the
// editor. A "Preview" CodeLens above each declaration that has one opens
// the panel on it, and the panel stays on that declaration while you edit
// elsewhere, redrawing as the pack changes. The pack comes from the
// server's `mms/output` (output.ts). A Book previews as its pages: its declaration
// has a lens, and so does each page of its .mmsbook file. A tell, title or
// actionbar message previews as the game shows it, from a lens at the
// command (or over the .mmstext file it reads).
import { bookPageAt, itemPath, shows } from '@mms/preview/paths';
import * as fs from 'node:fs';
import * as path from 'node:path';
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';
import type { OutputItem, OutputText, SplineAt, ToExtension, ToPanel } from './messages';
import { isSource, key, type PackOutput } from './output';
import { configDir } from './world';

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
	/**
	 * The last .mms editor and the spline at its cursor, for the spline
	 * editor. The editor is kept while the panel has focus, so a drag there
	 * still edits it.
	 */
	private editor: vscode.TextEditor | undefined;
	private spline: SplineAt | null = null;
	private splineTimer: ReturnType<typeof setTimeout> | undefined;

	constructor(
		private readonly context: vscode.ExtensionContext,
		private readonly output: PackOutput,
		private readonly client: () => LanguageClient | undefined = () => undefined
	) {
		this.editor = isMms(vscode.window.activeTextEditor) ? vscode.window.activeTextEditor : undefined;
		this.subscriptions.push(
			vscode.window.onDidChangeTextEditorSelection((e) => {
				if (!isMms(e.textEditor)) return;
				this.editor = e.textEditor;
				this.scheduleSpline();
			}),
			vscode.window.onDidChangeActiveTextEditor((e) => {
				if (e && isSource(e.document)) {
					this.editor = isMms(e) ? e : undefined;
					this.scheduleSpline();
				}
			}),
			output.onDidChange(() => this.scheduleSpline()),
			vscode.commands.registerCommand('mms.editSpline', (uri?: string, at?: { line: number; character: number }) => this.editSpline(uri, at)),
			this.lenses,
			vscode.languages.registerCodeLensProvider([{ language: 'mms' }, { language: 'mmsbook' }, { language: 'mmstext' }], this),
			output.onDidChange(() => {
				this.lenses.fire();
				this.post({ type: 'output', files: this.files, items: this.items, texts: this.output.texts, assets: this.output.assets, textures: this.textures(), structures: this.structures() });
				if (this.panel) this.panel.title = title(this.target, this.items, this.output.texts);
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

	private scheduleSpline() {
		if (!this.panel) return;
		clearTimeout(this.splineTimer);
		this.splineTimer = setTimeout(() => void this.refreshSpline(), 60);
	}

	/** Asks the language server for the spline at the cursor and sends it to the panel. */
	private async refreshSpline() {
		const ed = this.editor;
		const c = this.client();
		let spline: SplineAt | null = null;
		const version = ed?.document.version ?? -1;
		if (ed && !ed.document.isClosed && c?.isRunning()) {
			const pos = ed.selection.active;
			try {
				spline = await c.sendRequest<SplineAt | null>('mms/spline', {
					textDocument: { uri: ed.document.uri.toString() },
					position: { line: pos.line, character: pos.character }
				});
			} catch {
				spline = null;
			}
		}
		if (ed !== this.editor) return; // a newer request is on its way
		// The document changed while the server answered: its ranges are
		// out of date, and the change schedules a new request.
		if (ed && ed.document.version !== version) return void this.scheduleSpline();
		this.spline = spline;
		this.post({ type: 'spline', uri: ed?.document.uri.toString() ?? null, version, spline });
	}

	/** Opens the panel's spline editor on the spline at a position, or at the cursor. */
	private async editSpline(uri?: string, at?: { line: number; character: number }) {
		if (uri && at) {
			const doc = await vscode.workspace.openTextDocument(vscode.Uri.parse(uri));
			const column = vscode.window.visibleTextEditors.find((e) => e.document === doc)?.viewColumn ?? vscode.ViewColumn.One;
			const p = new vscode.Position(at.line, at.character);
			this.editor = await vscode.window.showTextDocument(doc, { viewColumn: column, selection: new vscode.Range(p, p) });
		}
		this.show(this.target ?? undefined);
		this.scheduleSpline();
	}

	private get files() {
		return this.output.files;
	}
	private get items() {
		return this.output.items;
	}

	/** Data URLs read for copied files, by file path, with the mtime they were read at. */
	private readonly dataCache = new Map<string, { mtime: number; url: string }>();

	/**
	 * The pack's copied files whose pack path matches as data URLs, by pack
	 * path, read from the project folder.
	 */
	private copied(match: RegExp, type: string): Record<string, string> {
		const out: Record<string, string> = {};
		const source = this.output.source;
		const project = source && configDir(path.dirname(source.fsPath));
		if (!project) return out;
		for (const a of this.output.assets) {
			if (!match.test(a.path)) continue;
			const file = path.join(project, a.src);
			try {
				const mtime = fs.statSync(file).mtimeMs;
				let hit = this.dataCache.get(file);
				if (!hit || hit.mtime !== mtime) {
					hit = { mtime, url: `data:${type};base64,` + fs.readFileSync(file).toString('base64') };
					this.dataCache.set(file, hit);
				}
				out[a.path] = hit.url;
			} catch {
				// A missing file is left out; the preview shows the name instead.
			}
		}
		return out;
	}

	/** The pack's copied .png textures, so the recipe inspector can draw custom item icons. */
	private textures(): Record<string, string> {
		return this.copied(/^assets\/[^/]+\/textures\/.+\.png$/, 'image/png');
	}

	/** The pack's structure files, so the jigsaw simulation can draw their blocks. */
	private structures(): Record<string, string> {
		return this.copied(/^data\/[^/]+\/structures?\/.+\.nbt$/, 'application/octet-stream');
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
		// A spline opens the spline editor.
		if (document.languageId === 'mms') {
			const text = document.getText();
			// Its own set: a Preview lens on the same line is another command.
			const splineLines = new Set<number>();
			for (const m of text.matchAll(/\bSpline\s*\(/g)) {
				const pos = document.positionAt(m.index! + m[0].length);
				if (splineLines.has(pos.line)) continue;
				splineLines.add(pos.line);
				lenses.push(
					new vscode.CodeLens(new vscode.Range(pos.line, 0, pos.line, 0), {
						title: '$(graph-line) Edit spline',
						tooltip: 'Show this spline as a graph in the MMS preview, with points to drag',
						command: 'mms.editSpline',
						arguments: [document.uri.toString(), { line: pos.line, character: pos.character }]
					})
				);
			}
		}
		for (const t of this.output.texts) {
			if (key(vscode.Uri.parse(t.uri)) !== here || lines.has(t.range.start.line)) continue;
			lines.add(t.range.start.line);
			lenses.push(
				new vscode.CodeLens(new vscode.Range(t.range.start.line, 0, t.range.start.line, 0), {
					title: '$(open-preview) Preview',
					tooltip: `Show this ${t.kind === 'chat' ? 'chat message' : t.kind} in the MMS preview`,
					command: 'mms.previewDeclaration',
					arguments: [t.id]
				})
			);
		}
		return lenses;
	}

	private setTarget(path: string | null, page = 0) {
		this.target = path;
		this.page = page;
		this.post({ type: 'target', path, page });
		if (this.panel) this.panel.title = title(path, this.items, this.output.texts);
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
		const t = this.output.texts.find((t) => key(vscode.Uri.parse(t.uri)) === here && t.range.start.line <= line && line <= t.range.end.line);
		if (t) return { path: t.id };
	}

	private adopt(panel: vscode.WebviewPanel) {
		this.panel = panel;
		const root = vscode.Uri.joinPath(this.context.extensionUri, 'out', 'webview');
		panel.title = title(this.target, this.items, this.output.texts);
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
			this.post({ type: 'output', files: this.files, items: this.items, texts: this.output.texts, assets: this.output.assets, textures: this.textures(), structures: this.structures() });
			this.setTarget(this.target, this.page);
			this.scheduleSpline();
			await this.output.refresh();
		} else if (m.type === 'splineEdit') {
			const uri = vscode.Uri.parse(m.uri);
			// The edit's ranges are in the version of the spline the page
			// showed; against any other version they would land on the
			// wrong text, so drop it and show the spline as it is now.
			const doc = vscode.workspace.textDocuments.find((d) => d.uri.toString() === uri.toString());
			if (!doc || doc.version !== m.version) return void this.scheduleSpline();
			const edit = new vscode.WorkspaceEdit();
			for (const e of m.edits) {
				edit.replace(uri, new vscode.Range(e.range.start.line, e.range.start.character, e.range.end.line, e.range.end.character), e.newText);
			}
			await vscode.workspace.applyEdit(edit);
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
		clearTimeout(this.splineTimer);
		this.panel?.dispose();
		for (const s of this.subscriptions) s.dispose();
	}
}

function isMms(e: vscode.TextEditor | undefined): e is vscode.TextEditor {
	return !!e && e.document.languageId === 'mms' && e.document.uri.scheme === 'file';
}

/** The panel's tab title: the declaration it shows. */
function title(path: string | null, items: OutputItem[], texts: OutputText[]): string {
	if (path?.startsWith('text:')) return 'Preview message';
	const id = path ? shows(path, items, texts)?.id : undefined;
	return id ? `Preview ${id.replace(/^[^:]+:/, '')}` : 'MMS Preview';
}
