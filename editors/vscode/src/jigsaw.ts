// The Jigsaw view in the MMS side bar: a form for a jigsaw block (target
// pool, target name, name, what it turns into, joint) that copies a /give
// command for a block set up that way, so pieces built in the game get
// their jigsaw blocks right without retyping ids in the jigsaw screen.
//
// The lists come from the pack (output.ts): the project's TemplatePools,
// its JigsawHole declarations (jigsaw names, which write no file), and the
// jigsaw blocks already in the pack's structure files, which load into the
// form to copy one. Copied blocks are remembered per project as Recent.
// The page (webview/Jigsaw.svelte) owns the form and builds the command.
import * as fs from 'fs';
import * as path from 'path';
import * as vscode from 'vscode';
import { sameBlock } from '@mms/preview/jigsaw/give';
import type { FromJigsaw, JigsawConfig, JigsawExisting, JigsawRecent, JigsawState, ToJigsaw } from './messages';
import { key, type PackOutput } from './output';

/** How many copied blocks Recent keeps per project. */
const RECENT_MAX = 20;

export class JigsawView implements vscode.Disposable, vscode.WebviewViewProvider {
	private view: vscode.WebviewView | undefined;
	private readonly subscriptions: vscode.Disposable[] = [];

	constructor(
		private readonly context: vscode.ExtensionContext,
		private readonly output: PackOutput,
		/** The project folder (holding mms.json) the side bar is about. */
		private readonly project: () => string | undefined
	) {
		this.subscriptions.push(
			vscode.window.registerWebviewViewProvider('mms.jigsaw', this),
			output.onDidChange(() => this.redraw()),
			output.onDidChangeDocument(() => this.redraw())
		);
	}

	resolveWebviewView(view: vscode.WebviewView) {
		this.view = view;
		const root = vscode.Uri.joinPath(this.context.extensionUri, 'out', 'jigsaw');
		view.webview.options = { enableScripts: true, enableCommandUris: false, localResourceRoots: [root] };
		view.webview.html = html(view.webview, root);
		view.webview.onDidReceiveMessage((m: FromJigsaw) => void this.receive(m), undefined, this.subscriptions);
		view.onDidChangeVisibility(() => view.visible && this.redraw(), undefined, this.subscriptions);
		view.onDidDispose(() => {
			if (this.view === view) this.view = undefined;
		});
	}

	private async receive(m: FromJigsaw) {
		switch (m.type) {
			case 'ready':
				return this.redraw();
			case 'copy':
				await vscode.env.clipboard.writeText(m.command);
				void vscode.window.setStatusBarMessage('$(check) Copied the give command; paste it in the game (creative, as an operator)', 4000);
				return this.remember(m.config);
			case 'forget':
				return this.setRecent(this.recent().filter((r) => !sameBlock(r.config, m.config, this.ns()) || r.config.label !== m.config.label));
			case 'forgetAll':
				return this.setRecent([]);
			case 'declare':
				return this.declare(m.id);
			case 'revealHole': {
				const h = this.output.jigsawHoles.find((h) => h.id === m.id);
				if (h) await reveal(h.uri, h.range);
				return;
			}
			case 'revealFile': {
				const project = this.project();
				if (project) await vscode.commands.executeCommand('revealInExplorer', vscode.Uri.file(path.join(project, ...m.src.split('/'))));
				return;
			}
		}
	}

	private ns(): string | undefined {
		const project = this.project();
		return project ? namespace(project) : undefined;
	}

	/** The project's recent blocks, newest first. */
	private recent(): JigsawRecent[] {
		const project = this.project();
		if (!project) return [];
		return this.context.workspaceState.get<Record<string, JigsawRecent[]>>('mms.jigsawRecent')?.[key(vscode.Uri.file(project))] ?? [];
	}

	private async setRecent(list: JigsawRecent[]) {
		const project = this.project();
		if (!project) return;
		const all = { ...this.context.workspaceState.get<Record<string, JigsawRecent[]>>('mms.jigsawRecent') };
		all[key(vscode.Uri.file(project))] = list.slice(0, RECENT_MAX);
		await this.context.workspaceState.update('mms.jigsawRecent', all);
		this.redraw();
	}

	/** Puts a copied block at the top of Recent; the same block copied again moves up, with its new label. */
	private remember(config: JigsawConfig) {
		const ns = this.ns();
		const rest = this.recent().filter((r) => !sameBlock(r.config, config, ns));
		return this.setRecent([{ config, at: Date.now() }, ...rest]);
	}

	private state(): JigsawState {
		const project = this.project();
		const ns = project ? namespace(project) : undefined;
		const state: JigsawState = { holes: [], pools: [], recent: [], existing: [], hasErrors: this.output.hasErrors };
		if (!project) return state;
		state.project = { name: path.basename(project), ns };
		state.holes = this.output.jigsawHoles.map((h) => ({ id: h.id, doc: h.doc }));
		state.recent = this.recent();
		// Structure files by id, with their jigsaw blocks.
		const structures = new Map<string, NonNullable<(typeof this.output.assets)[number]['jigsaws']>>();
		for (const a of this.output.assets) {
			const m = /^data\/([^/]+)\/structures?\/(.+)\.nbt$/.exec(a.path);
			if (!m || !a.jigsaws) continue;
			const id = `${m[1]}:${m[2]}`;
			structures.set(id, a.jigsaws);
			for (const j of a.jigsaws) {
				const existing: JigsawExisting = {
					src: a.src,
					id,
					pos: j.pos,
					orientation: j.orientation,
					config: { label: j.label ?? '', pool: field(j.pool), target: field(j.target), name: field(j.name), finalState: j.final_state ?? 'minecraft:air', joint: j.joint === 'aligned' ? 'aligned' : 'rollable' }
				};
				state.existing.push(existing);
			}
		}
		for (const f of this.output.files) {
			const m = /^data\/([^/]+)\/worldgen\/template_pool\/(.+)\.json$/.exec(f.path);
			if (!m) continue;
			const names = new Set<string>();
			for (const loc of poolLocations(f.text)) for (const j of structures.get(loc) ?? []) if (j.name !== 'minecraft:empty') names.add(j.name);
			state.pools.push({ id: `${m[1]}:${m[2]}`, names: [...names].sort() });
		}
		state.pools.sort((a, b) => a.id.localeCompare(b.id));
		return state;
	}

	private redraw() {
		if (!this.view) return;
		void this.view.webview.postMessage({ type: 'state', state: this.state() } satisfies ToJigsaw);
	}

	/**
	 * Adds `JigsawHole <name>` after the project's last JigsawHole, or at
	 * the end of the file of its first TemplatePool, and shows it. The
	 * file is left unsaved, to undo or keep.
	 */
	private async declare(id: string) {
		const [ns, name] = id.includes(':') ? id.split(':', 2) : [this.ns() ?? 'minecraft', id];
		const holes = this.output.jigsawHoles;
		// The last hole of the file the holes start in.
		const last = holes.filter((h) => h.uri === holes[0]?.uri).reduce<(typeof holes)[number] | undefined>((a, b) => (!a || b.range.end.line > a.range.end.line ? b : a), undefined);
		let uri: vscode.Uri | undefined;
		let line: number | undefined;
		let fileNs: string | undefined;
		if (last) {
			uri = vscode.Uri.parse(last.uri);
			line = last.range.end.line + 1;
			fileNs = last.id.split(':')[0];
		} else {
			const pool = this.output.files.find((f) => /^data\/[^/]+\/worldgen\/template_pool\//.test(f.path) && f.uri);
			if (pool?.uri) {
				uri = vscode.Uri.parse(pool.uri);
				fileNs = pool.path.split('/')[1];
			} else if (vscode.window.activeTextEditor?.document.languageId === 'mms') {
				uri = vscode.window.activeTextEditor.document.uri;
				fileNs = this.ns();
			}
		}
		if (!uri) {
			void vscode.window.showWarningMessage(`Add \`JigsawHole ${name}\` to one of the project's .mms files to declare it.`);
			return;
		}
		const doc = await vscode.workspace.openTextDocument(uri);
		const text = `JigsawHole ${ns === fileNs ? name : `${ns}:${name}`}`;
		const at = line !== undefined && line <= doc.lineCount ? new vscode.Position(line, 0) : doc.positionAt(doc.getText().length);
		const atEnd = line === undefined || line >= doc.lineCount;
		const before = atEnd && doc.getText().length > 0 && !doc.getText().endsWith('\n') ? '\n' : '';
		const edit = new vscode.WorkspaceEdit();
		edit.insert(uri, at, last ? `${before}${text}\n` : `${before}\n${text}\n`);
		await vscode.workspace.applyEdit(edit);
		const editor = await vscode.window.showTextDocument(doc, { preview: false, preserveFocus: true });
		const added = doc.getText().indexOf(text, doc.offsetAt(at));
		if (added >= 0) {
			const r = new vscode.Range(doc.positionAt(added), doc.positionAt(added + text.length));
			editor.selection = new vscode.Selection(r.start, r.end);
			editor.revealRange(r, vscode.TextEditorRevealType.InCenterIfOutsideViewport);
		}
	}

	dispose() {
		for (const s of this.subscriptions) s.dispose();
	}
}

/** An id as the form holds it: minecraft:empty, the game's default, as an empty field. */
function field(id: string): string {
	return id === 'minecraft:empty' ? '' : id;
}

/** The structure ids a template pool's JSON places, list elements included. */
function poolLocations(text: string): string[] {
	const out: string[] = [];
	const walk = (el: unknown) => {
		if (!el || typeof el !== 'object') return;
		const o = el as { location?: unknown; elements?: unknown[] };
		if (typeof o.location === 'string') out.push(o.location.includes(':') ? o.location : `minecraft:${o.location}`);
		if (Array.isArray(o.elements)) o.elements.forEach(walk);
	};
	try {
		const pool = JSON.parse(text) as { elements?: { element?: unknown }[] };
		for (const e of pool.elements ?? []) walk(e.element);
	} catch {
		// Not JSON yet; no pieces.
	}
	return out;
}

async function reveal(uri: string, range: { start: { line: number; character: number } }) {
	const doc = await vscode.workspace.openTextDocument(vscode.Uri.parse(uri));
	const at = new vscode.Position(range.start.line, range.start.character);
	await vscode.window.showTextDocument(doc, { selection: new vscode.Range(at, at), preview: true });
}

/** The namespace in a project's mms.json, or mms, the engine's default. */
function namespace(project: string): string {
	try {
		const ns = JSON.parse(fs.readFileSync(path.join(project, 'mms.json'), 'utf8')).namespace;
		return typeof ns === 'string' && ns ? ns : 'mms';
	} catch {
		return 'mms';
	}
}

function html(webview: vscode.Webview, root: vscode.Uri): string {
	const nonce = Array.from({ length: 32 }, () => Math.floor(Math.random() * 36).toString(36)).join('');
	const script = webview.asWebviewUri(vscode.Uri.joinPath(root, 'jigsaw.js'));
	const style = webview.asWebviewUri(vscode.Uri.joinPath(root, 'jigsaw.css'));
	const csp = ["default-src 'none'", `style-src ${webview.cspSource} 'unsafe-inline'`, `script-src 'nonce-${nonce}'`, `font-src ${webview.cspSource}`].join('; ');
	return `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta http-equiv="Content-Security-Policy" content="${csp}">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="${style}">
<title>Jigsaw</title>
</head>
<body>
<div id="app"></div>
<script type="module" nonce="${nonce}" src="${script}"></script>
</body>
</html>`;
}
