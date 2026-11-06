// Expands puts an "Expand" CodeLens above each Blueprint instance (such as
// a CustomItem) and each Commands Template call, and adds "MMS: Expand at
// Cursor". Both show the code it writes, from the language server's
// `mms/expand`, in a read-only document beside the source.
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';
import type { Range } from './messages';

interface Expansion {
	title: string;
	text: string;
	range: Range;
}

const scheme = 'mms-expand';

export class Expands implements vscode.Disposable, vscode.CodeLensProvider, vscode.TextDocumentContentProvider {
	private readonly lenses = new vscode.EventEmitter<void>();
	readonly onDidChangeCodeLenses = this.lenses.event;
	private readonly docs = new vscode.EventEmitter<vscode.Uri>();
	readonly onDidChange = this.docs.event;
	private readonly texts = new Map<string, string>();
	private readonly subs: vscode.Disposable[] = [this.lenses, this.docs];

	constructor(
		private readonly client: () => LanguageClient | undefined,
		changed: vscode.Event<void>
	) {
		this.subs.push(
			vscode.languages.registerCodeLensProvider({ language: 'mms' }, this),
			vscode.workspace.registerTextDocumentContentProvider(scheme, this),
			changed(() => this.lenses.fire()),
			vscode.commands.registerCommand('mms.showExpansion', (x: Expansion) => this.show(x)),
			vscode.commands.registerCommand('mms.expandAtCursor', () => this.atCursor())
		);
	}

	async provideCodeLenses(doc: vscode.TextDocument): Promise<vscode.CodeLens[]> {
		const c = this.client();
		if (!c?.isRunning()) return [];
		try {
			const xs = await c.sendRequest<Expansion[]>('mms/expand', { textDocument: { uri: doc.uri.toString() } });
			return xs.map(
				(x) =>
					new vscode.CodeLens(new vscode.Range(x.range.start.line, 0, x.range.start.line, 0), {
						title: 'Expand',
						tooltip: `Show the code ${x.title} writes`,
						command: 'mms.showExpansion',
						arguments: [x]
					})
			);
		} catch {
			return [];
		}
	}

	provideTextDocumentContent(uri: vscode.Uri): string {
		return this.texts.get(uri.toString()) ?? '';
	}

	private async atCursor() {
		const ed = vscode.window.activeTextEditor;
		const c = this.client();
		if (!ed || !c?.isRunning()) return;
		const pos = ed.selection.active;
		const x = await c.sendRequest<Expansion | null>('mms/expand', {
			textDocument: { uri: ed.document.uri.toString() },
			position: { line: pos.line, character: pos.character }
		});
		if (!x) {
			void vscode.window.showInformationMessage('Put the cursor in a Template call or a Blueprint declaration, such as a CustomItem, to expand it.');
			return;
		}
		await this.show(x);
	}

	private async show(x: Expansion) {
		// One document per expansion title, refreshed in place.
		const uri = vscode.Uri.from({ scheme, path: `/${x.title.replace(/[^\w.-]+/g, '_')}.mms` });
		this.texts.set(uri.toString(), `// What ${x.title} writes\n${x.text}\n`);
		this.docs.fire(uri);
		const doc = await vscode.workspace.openTextDocument(uri);
		await vscode.languages.setTextDocumentLanguage(doc, 'mms');
		await vscode.window.showTextDocument(doc, { viewColumn: vscode.ViewColumn.Beside, preview: true, preserveFocus: true });
	}

	dispose() {
		for (const s of this.subs) s.dispose();
	}
}
