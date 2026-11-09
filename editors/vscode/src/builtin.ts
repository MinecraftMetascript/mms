// BuiltinFiles shows files that come with the language server, where go to
// definition leads: the language's prelude (a built-in Template, Blueprint
// or Type) under the mms-prelude: scheme, and vanilla's own data pack JSON
// (a vanilla noise, loot table or tag) under the mms-vanilla: scheme. Both
// are read-only documents whose text comes from a server request.
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';

export class BuiltinFiles implements vscode.Disposable, vscode.TextDocumentContentProvider {
	private readonly subs: vscode.Disposable[];

	/** Serves scheme's documents with the server's request method, shown as languageId. */
	constructor(
		private readonly client: () => LanguageClient | undefined,
		scheme: string,
		private readonly method: string,
		languageId: string
	) {
		this.subs = [
			vscode.workspace.registerTextDocumentContentProvider(scheme, this),
			// Highlight a file as its language when it opens.
			vscode.workspace.onDidOpenTextDocument((doc) => {
				if (doc.uri.scheme === scheme && doc.languageId !== languageId) void vscode.languages.setTextDocumentLanguage(doc, languageId);
			})
		];
	}

	async provideTextDocumentContent(uri: vscode.Uri): Promise<string> {
		const c = this.client();
		if (!c?.isRunning()) return '// The MMS language server is not running.';
		const res = await c.sendRequest<{ text: string }>(this.method, { uri: uri.toString() });
		return res.text;
	}

	dispose() {
		for (const s of this.subs) s.dispose();
	}
}
