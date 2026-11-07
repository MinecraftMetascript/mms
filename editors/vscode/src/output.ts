// The generated pack of the project being edited, from the language
// server's `mms/output`. The preview panel and the Pack view both read it.
// It is fetched again shortly after each analysis (publishDiagnostics) and
// edit, and when another project's file becomes active: the request's uri
// picks the project (the nearest mms.json) in a workspace with several.
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';
import type { OutputAsset, OutputFile, OutputHole, OutputItem } from './messages';

export class PackOutput implements vscode.Disposable {
	files: OutputFile[] = [];
	items: OutputItem[] = [];
	assets: OutputAsset[] = [];
	/** The project's JigsawHole declarations. */
	jigsawHoles: OutputHole[] = [];
	hasErrors = false;
	/** Project files nothing uses (lint/unused-asset), relative to the mms.json folder. */
	unusedAssets: string[] = [];
	/** The source document whose project's pack is fetched. */
	private document: vscode.Uri | undefined;
	private timer: ReturnType<typeof setTimeout> | undefined;
	private readonly changed = new vscode.EventEmitter<void>();
	/** Fires after each fetch. */
	readonly onDidChange = this.changed.event;
	private readonly projectChanged = new vscode.EventEmitter<void>();
	/** Fires when the pack starts following another source document. */
	readonly onDidChangeDocument = this.projectChanged.event;
	private readonly subscriptions: vscode.Disposable[] = [];

	constructor(private readonly client: () => LanguageClient | undefined) {
		const active = vscode.window.activeTextEditor?.document;
		if (isSource(active)) this.document = active!.uri;
		this.subscriptions.push(
			this.changed,
			this.projectChanged,
			vscode.window.onDidChangeActiveTextEditor((e) => {
				if (e && isSource(e.document)) this.follow(e.document.uri);
			}),
			// A project with no problems publishes no diagnostics, so edits
			// refresh the pack too.
			vscode.workspace.onDidOpenTextDocument((d) => isSource(d) && this.scheduleRefresh()),
			vscode.workspace.onDidChangeTextDocument((e) => isSource(e.document) && this.scheduleRefresh())
		);
	}

	/** The source document the pack follows, if any. */
	get source(): vscode.Uri | undefined {
		return this.document;
	}

	/** Follows another source document, which may be in another project. */
	follow(uri: vscode.Uri) {
		if (uri.toString() === this.document?.toString()) return;
		this.document = uri;
		// The old project's unused files are not this one's; the next fetch fills them.
		this.unusedAssets = [];
		this.projectChanged.fire();
		this.scheduleRefresh();
	}

	/** Fetches the pack shortly, so a burst of analyses fetches once. */
	scheduleRefresh() {
		if (this.timer) clearTimeout(this.timer);
		this.timer = setTimeout(() => void this.refresh(), 150);
	}

	async refresh() {
		const client = this.client();
		if (!client?.isRunning()) return;
		try {
			const asked = this.document;
			const uri = asked?.toString();
			const out = await client.sendRequest<{ files: OutputFile[]; items?: OutputItem[]; assets?: OutputAsset[]; jigsawHoles?: OutputHole[]; unusedAssets?: string[]; hasErrors: boolean }>(
				'mms/output',
				uri ? { uri } : {}
			);
			// A reply for a document the pack no longer follows is stale; the
			// fetch scheduled by follow() brings the current one.
			if (asked !== this.document) return;
			this.files = out.files;
			this.items = out.items ?? [];
			this.assets = out.assets ?? [];
			this.jigsawHoles = out.jigsawHoles ?? [];
			this.hasErrors = out.hasErrors;
			this.unusedAssets = out.unusedAssets ?? [];
			this.changed.fire();
		} catch {
			// A restart or shutdown; the next analysis refreshes again.
		}
	}

	dispose() {
		if (this.timer) clearTimeout(this.timer);
		for (const s of this.subscriptions) s.dispose();
	}
}

/** Whether a document is MMS source: a .mms file, a .mmsbook book text or a .mmstext text. */
export function isSource(doc: vscode.TextDocument | undefined): boolean {
	return doc?.languageId === 'mms' || doc?.languageId === 'mmsbook' || doc?.languageId === 'mmstext';
}

/** A file's identity: its path, ignoring case on Windows. */
export function key(uri: vscode.Uri): string {
	return process.platform === 'win32' ? uri.fsPath.toLowerCase() : uri.fsPath;
}
