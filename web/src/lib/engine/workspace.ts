import { LSPPlugin, Workspace, type LSPClient, type WorkspaceFile } from '@codemirror/lsp-client';
import { ChangeSet, Text, type TransactionSpec } from '@codemirror/state';
import type { EditorView } from '@codemirror/view';

class ProjectFile implements WorkspaceFile {
	view: EditorView | null = null;
	/** Changes applied while no editor showed the file, not yet synced. */
	pending: { changes: ChangeSet; prevDoc: Text } | null = null;
	constructor(
		readonly uri: string,
		public version: number,
		public doc: Text
	) {}
	get languageId() {
		return 'mms';
	}
	getView() {
		return this.view;
	}
}

export interface ProjectWorkspaceHooks {
	/** A file's text changed outside its editor (a rename edit). */
	onText(uri: string, text: string): void;
	/** Show a file in the editor and resolve to its view once it is there. */
	display(uri: string): Promise<EditorView | null>;
}

/**
 * ProjectWorkspace keeps every project file open on the server, not only
 * the one in the editor, so references, rename and diagnostics see the
 * whole project. Edits to files without an editor (from rename) are
 * applied to the stored text and synced like editor changes.
 */
export class ProjectWorkspace extends Workspace {
	files: ProjectFile[] = [];
	private versions: Record<string, number> = {};

	constructor(
		client: LSPClient,
		private hooks: ProjectWorkspaceHooks
	) {
		super(client);
	}

	private nextVersion(uri: string) {
		return (this.versions[uri] = (this.versions[uri] ?? -1) + 1);
	}

	/** Replaces the set of files, opening and closing them on the server. */
	setFiles(files: { uri: string; text: string }[]) {
		const want = new Map(files.map((f) => [f.uri, f.text]));
		for (const f of [...this.files]) {
			if (!want.has(f.uri)) this.removeFile(f.uri);
		}
		for (const [uri, text] of want) {
			if (!this.getFile(uri)) this.addFile(uri, text);
		}
	}

	addFile(uri: string, text: string) {
		if (this.getFile(uri)) return;
		const f = new ProjectFile(uri, this.nextVersion(uri), Text.of(text.split('\n')));
		this.files.push(f);
		if (this.client.connected) this.client.didOpen(f);
	}

	removeFile(uri: string) {
		const f = this.getFile(uri) as ProjectFile | null;
		if (!f) return;
		this.files = this.files.filter((x) => x !== f);
		if (this.client.connected) this.client.didClose(uri);
	}

	syncFiles() {
		const out = [];
		for (const f of this.files) {
			if (f.pending) {
				out.push({ file: f, prevDoc: f.pending.prevDoc, changes: f.pending.changes });
				f.pending = null;
				f.version = this.nextVersion(f.uri);
			}
			const plugin = f.view && LSPPlugin.get(f.view);
			if (!plugin) continue;
			const changes = plugin.unsyncedChanges;
			if (!changes.empty) {
				out.push({ file: f, prevDoc: f.doc, changes });
				f.doc = f.view!.state.doc;
				f.version = this.nextVersion(f.uri);
				plugin.clear();
			}
		}
		return out;
	}

	openFile(uri: string, _languageId: string, view: EditorView) {
		let f = this.getFile(uri) as ProjectFile | null;
		if (!f) {
			f = new ProjectFile(uri, this.nextVersion(uri), view.state.doc);
			this.files.push(f);
			this.client.didOpen(f);
		}
		f.view = view;
		if (f.doc.toString() !== view.state.doc.toString()) {
			// The editor was created from newer text than the server has.
			const prevDoc = f.doc;
			f.pending = {
				prevDoc,
				changes: ChangeSet.of([{ from: 0, to: prevDoc.length, insert: view.state.doc }], prevDoc.length)
			};
			f.doc = view.state.doc;
		}
	}

	closeFile(uri: string, view: EditorView) {
		const f = this.getFile(uri) as ProjectFile | null;
		if (f && f.view === view) {
			// Flush the editor's last edits before it goes away.
			const plugin = LSPPlugin.get(view);
			if (plugin && !plugin.unsyncedChanges.empty) this.client.sync();
			f.view = null;
		}
	}

	updateFile(uri: string, update: TransactionSpec) {
		const f = this.getFile(uri) as ProjectFile | null;
		if (!f) return;
		if (f.view) {
			f.view.dispatch(update);
			return;
		}
		if (!update.changes) return;
		const changes = ChangeSet.of(update.changes, f.doc.length);
		const prevDoc = f.pending ? f.pending.prevDoc : f.doc;
		const all = f.pending ? f.pending.changes.compose(changes) : changes;
		f.doc = changes.apply(f.doc);
		f.pending = { prevDoc, changes: all };
		this.hooks.onText(uri, f.doc.toString());
		this.client.sync();
	}

	displayFile(uri: string): Promise<EditorView | null> {
		return this.hooks.display(uri);
	}
}
