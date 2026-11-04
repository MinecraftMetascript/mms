// The Pack view in the MMS side bar: the files the project generates, as
// the web editor's Pack tab lists them. The data pack, the resource pack
// (when the project has one) and the custom items each get a tree, with
// single-child folder chains squashed into one row (@mms/preview/tree, as
// on the web). Opening a file shows its text read-only; it updates as you
// edit. A file with a preview or inspector can open in the preview panel,
// and Go to Source jumps to the declaration that wrote it.
import { itemPath, shows } from '@mms/preview/paths';
import { buildTree, type Node } from '@mms/preview/tree';
import * as path from 'path';
import * as vscode from 'vscode';
import type { OutputAsset, OutputFile, OutputItem, Range } from './messages';
import type { PackOutput } from './output';

export const scheme = 'mms-output';

type Pack = 'data' | 'resources';

type PackNode =
	| { kind: 'section'; pack: Pack | 'items'; label: string }
	| { kind: 'folder'; pack: Pack; node: Node }
	| { kind: 'file'; pack: Pack; node: Node; file?: OutputFile; asset?: OutputAsset }
	| { kind: 'item'; item: OutputItem };

export class PackView implements vscode.Disposable, vscode.TreeDataProvider<PackNode>, vscode.TextDocumentContentProvider {
	private readonly changed = new vscode.EventEmitter<void>();
	readonly onDidChangeTreeData = this.changed.event;
	private readonly docs = new vscode.EventEmitter<vscode.Uri>();
	readonly onDidChange = this.docs.event;
	private readonly view: vscode.TreeView<PackNode>;
	private trees: Record<Pack, Node[]> = { data: [], resources: [] };
	private readonly subscriptions: vscode.Disposable[] = [];

	constructor(
		private readonly output: PackOutput,
		/** The project folder the pack is from, for asset sources. */
		private readonly project: () => string | undefined
	) {
		this.view = vscode.window.createTreeView('mms.pack', { treeDataProvider: this, showCollapseAll: true });
		this.subscriptions.push(
			this.changed,
			this.docs,
			this.view,
			vscode.workspace.registerTextDocumentContentProvider(scheme, this),
			output.onDidChange(() => this.update()),
			vscode.commands.registerCommand('mms.openOutput', (n: PackNode) => this.open(n)),
			vscode.commands.registerCommand('mms.revealOutputSource', (n: PackNode) => this.reveal(n)),
			vscode.commands.registerCommand('mms.previewOutput', (n: PackNode) => {
				const p = previewPath(n);
				if (p) void vscode.commands.executeCommand('mms.previewDeclaration', p);
			})
		);
		this.update();
	}

	getTreeItem(n: PackNode): vscode.TreeItem {
		switch (n.kind) {
			case 'section': {
				const t = new vscode.TreeItem(n.label, vscode.TreeItemCollapsibleState.Expanded);
				t.id = `section:${n.pack}`;
				if (n.pack === 'items') t.tooltip = 'A custom item is not a file: its name stands for this item stack in recipes, loot tables and commands.';
				return t;
			}
			case 'folder': {
				const t = new vscode.TreeItem(n.node.name, vscode.TreeItemCollapsibleState.Collapsed);
				t.id = `${n.pack}:${n.node.path}/`;
				t.iconPath = vscode.ThemeIcon.Folder;
				return t;
			}
			case 'file': {
				const t = new vscode.TreeItem(n.node.name);
				t.id = `${n.pack}:${n.node.path}`;
				// A file uri gives the row the theme's icon for its extension.
				t.resourceUri = vscode.Uri.file(n.node.name);
				t.iconPath = vscode.ThemeIcon.File;
				t.command = { command: 'mms.openOutput', title: 'Open', arguments: [n] };
				const from = n.file?.uri && n.file.range ? `${vscode.workspace.asRelativePath(vscode.Uri.parse(n.file.uri))}:${n.file.range.start.line + 1}` : undefined;
				if (n.asset) {
					t.description = 'copied';
					t.tooltip = `${n.node.path}\n\nCopied from ${n.asset.src}`;
				} else {
					if (n.file?.implicit) t.description = 'inline value';
					t.tooltip = from ? `${n.node.path}\n\nWritten by ${from}` : `${n.node.path}\n\nGenerated`;
				}
				t.contextValue = [n.file?.uri ? 'source' : '', previewPath(n) ? 'preview' : ''].filter(Boolean).join(' ') || 'file';
				return t;
			}
			case 'item': {
				const t = new vscode.TreeItem(n.item.id);
				t.id = `item:${n.item.id}`;
				t.iconPath = new vscode.ThemeIcon(previewPath(n, this.output.items) ? 'book' : 'symbol-object');
				t.tooltip = 'A custom item: the item stack its name stands for.';
				t.command = { command: 'mms.openOutput', title: 'Open', arguments: [n] };
				t.contextValue = ['source', previewPath(n, this.output.items) ? 'preview' : ''].filter(Boolean).join(' ');
				return t;
			}
		}
	}

	getChildren(n?: PackNode): PackNode[] {
		if (!n) {
			const items = this.output.items;
			const sections = this.trees.resources.length > 0 || items.length > 0;
			if (!sections) return this.nodes('data', this.trees.data);
			const out: PackNode[] = [{ kind: 'section', pack: 'data', label: 'Data pack' }];
			if (this.trees.resources.length) out.push({ kind: 'section', pack: 'resources', label: 'Resource pack' });
			if (items.length) out.push({ kind: 'section', pack: 'items', label: 'Custom items' });
			return out;
		}
		if (n.kind === 'section') {
			if (n.pack === 'items') return this.output.items.map((item) => ({ kind: 'item', item }));
			return this.nodes(n.pack, this.trees[n.pack]);
		}
		if (n.kind === 'folder') return this.nodes(n.pack, n.node.children);
		return [];
	}

	/** The text of an opened output: a pack file, or an item stack as JSON. */
	provideTextDocumentContent(uri: vscode.Uri): string {
		const [, pack, ...rest] = uri.path.split('/');
		const p = rest.join('/');
		if (pack === 'item') {
			const it = this.output.items.find((it) => fileName(it.id) === p);
			return it ? pretty(it.text) : '// This item is no longer in the pack.\n';
		}
		const f = this.output.files.find((f) => packOf(f) === pack && f.path === p);
		return f ? f.text : '// This file is no longer in the pack.\n';
	}

	private nodes(pack: Pack, nodes: Node[]): PackNode[] {
		return nodes.map((node) => {
			if (!node.leaf) return { kind: 'folder', pack, node };
			const file = this.output.files.find((f) => packOf(f) === pack && f.path === node.path);
			const asset = file ? undefined : this.output.assets.find((a) => packOf(a) === pack && a.path === node.path);
			return { kind: 'file', pack, node, file, asset };
		});
	}

	private async open(n: PackNode) {
		if (n.kind === 'file' && n.asset) {
			// A copied texture or sound opens from the project.
			const project = this.project();
			if (project) await vscode.commands.executeCommand('vscode.open', vscode.Uri.file(path.join(project, n.asset.src)), { preview: true });
			return;
		}
		const uri = outputUri(n);
		if (!uri) return;
		const doc = await vscode.workspace.openTextDocument(uri);
		await vscode.window.showTextDocument(doc, { preview: true, preserveFocus: true });
	}

	private async reveal(n: PackNode) {
		const at = n.kind === 'file' ? n.file : n.kind === 'item' ? n.item : undefined;
		if (!at?.uri || !at.range) return;
		const r: Range = at.range;
		const range = new vscode.Range(r.start.line, r.start.character, r.end.line, r.end.character);
		await vscode.window.showTextDocument(vscode.Uri.parse(at.uri), { selection: new vscode.Range(range.start, range.start) });
	}

	private update() {
		const tree = (pack: Pack) => buildTree([...this.output.files.filter((f) => packOf(f) === pack), ...this.output.assets.filter((a) => packOf(a) === pack)].map((f) => f.path));
		this.trees = { data: tree('data'), resources: tree('resources') };
		this.view.description = this.project() ? path.basename(this.project()!) : undefined;
		this.view.message = this.output.hasErrors ? 'The project has errors, so this pack is incomplete.' : undefined;
		this.changed.fire();
		// Opened outputs show the new text.
		for (const d of vscode.workspace.textDocuments) if (d.uri.scheme === scheme) this.docs.fire(d.uri);
	}

	dispose() {
		for (const s of this.subscriptions) s.dispose();
	}
}

const packOf = (f: { pack?: string }): Pack => (f.pack === 'resources' ? 'resources' : 'data');

/** An item id as a file name, for its tab: ns:sword → ns_sword.json. */
const fileName = (id: string) => `${id.replace(/[:/]/g, '_')}.json`;

/** The uri an output opens at: mms-output:/<data|resources|item>/<path>. */
function outputUri(n: PackNode): vscode.Uri | undefined {
	if (n.kind === 'file') return vscode.Uri.from({ scheme, path: `/${n.pack}/${n.node.path}` });
	if (n.kind === 'item') return vscode.Uri.from({ scheme, path: `/item/${fileName(n.item.id)}` });
	return undefined;
}

/** The preview panel's path for an output that has a preview or inspector. */
function previewPath(n: PackNode, items?: OutputItem[]): string | undefined {
	if (n.kind === 'file' && n.pack === 'data' && n.file && shows(n.file.path)) return n.file.path;
	if (n.kind === 'item' && shows(itemPath(n.item.id), items ?? [n.item])) return itemPath(n.item.id);
	return undefined;
}

function pretty(json: string): string {
	try {
		return JSON.stringify(JSON.parse(json), null, 2) + '\n';
	} catch {
		return json;
	}
}
