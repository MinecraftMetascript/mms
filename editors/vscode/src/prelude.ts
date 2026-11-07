// PreludeFiles shows the language's prelude, where go to definition on a
// built-in Template, Blueprint or Type leads: read-only documents under the
// mms-prelude: scheme, whose text comes from the server's mms/preludeFile.
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';

export const preludeScheme = 'mms-prelude';

export class PreludeFiles implements vscode.Disposable, vscode.TextDocumentContentProvider {
	private readonly subs: vscode.Disposable[];

	constructor(private readonly client: () => LanguageClient | undefined) {
		this.subs = [
			vscode.workspace.registerTextDocumentContentProvider(preludeScheme, this),
			// Highlight a prelude file as MMS when it opens.
			vscode.workspace.onDidOpenTextDocument((doc) => {
				if (doc.uri.scheme === preludeScheme && doc.languageId !== 'mms') void vscode.languages.setTextDocumentLanguage(doc, 'mms');
			})
		];
	}

	async provideTextDocumentContent(uri: vscode.Uri): Promise<string> {
		const c = this.client();
		if (!c?.isRunning()) return '// The MMS language server is not running.';
		const res = await c.sendRequest<{ text: string }>('mms/preludeFile', { uri: uri.toString() });
		return res.text;
	}

	dispose() {
		for (const s of this.subs) s.dispose();
	}
}
