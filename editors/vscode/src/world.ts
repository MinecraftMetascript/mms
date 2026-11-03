// The World view in the MMS side bar: the project being edited, the
// Minecraft world chosen for it, and a button that builds the pack into
// that world. Building runs `mms build --world <world> <project>`, the same
// install as the command line: the data pack goes to
// <world>/datapacks/<namespace> and a resource pack is merged into
// <world>/resourcepacks/resources.zip, next to other projects' files. It builds the files on disk, so the
// project's unsaved files are saved first.
//
// Under the world, Structures lists what structure blocks saved in it
// (structures.ts), so pieces built in the game can be pulled into the
// project for jigsaw template pools, and pulled again after each re-save.
import { execFile } from 'child_process';
import * as fs from 'fs';
import * as os from 'os';
import * as path from 'path';
import * as vscode from 'vscode';
import { key, type PackOutput } from './output';
import { findStructures, importPath, importStructure, status, structureCall, type Status, type WorldStructure } from './structures';
import { findWorlds, readWorld, savesFolders, type World } from './worlds';

type Row = { id: string; item: vscode.TreeItem; children?: () => Row[]; structure?: WorldStructure };

/** The outcome of the last build, shown under the button. */
type Result =
	| { ok: true; world: string; files: number; resources: boolean; next: NextStep | undefined; at: Date }
	| { ok: false; message: string; detail: string; at: Date };

/**
 * What the player does in the game for the build to take effect, from the
 * "next step:" line `mms build --world` prints last: run /reload, or exit
 * and rejoin the world when worldgen, other registries or the resource
 * pack changed. Older mms builds print no such line.
 */
type NextStep = { action: 'reload' | 'rejoin' | 'none'; line: string };

export function nextStep(stdout: string): NextStep | undefined {
	const m = /^next step: (run \/reload|exit and rejoin|none)\b.*$/m.exec(stdout);
	if (!m) return undefined;
	const action = m[1] === 'none' ? 'none' : m[1] === 'exit and rejoin' ? 'rejoin' : 'reload';
	return { action, line: m[0].replace(/^next step: /, '') };
}

export class WorldView implements vscode.Disposable, vscode.TreeDataProvider<Row> {
	private readonly changed = new vscode.EventEmitter<void>();
	readonly onDidChangeTreeData = this.changed.event;
	private readonly view: vscode.TreeView<Row>;
	/** Every folder with an mms.json in the workspace. */
	private projects: string[] = [];
	private building = false;
	private result: Result | undefined;
	private saveTimer: ReturnType<typeof setTimeout> | undefined;
	/** Watches the world's generated structures while one is chosen. */
	private worldWatcher: { world: string; watcher: vscode.FileSystemWatcher } | undefined;
	private refreshTimer: ReturnType<typeof setTimeout> | undefined;
	private readonly subscriptions: vscode.Disposable[] = [];

	constructor(
		private readonly context: vscode.ExtensionContext,
		private readonly output: PackOutput,
		/** The mms executable, or why there is none. */
		private readonly mms: () => { command: string } | { error: string },
		private readonly log: (line: string) => void
	) {
		this.view = vscode.window.createTreeView('mms.world', { treeDataProvider: this });
		const watcher = vscode.workspace.createFileSystemWatcher('**/mms.json', false, true, false);
		// An imported structure changing in the project changes its status.
		const nbt = vscode.workspace.createFileSystemWatcher('**/*.nbt');
		this.subscriptions.push(
			this.changed,
			this.view,
			watcher,
			watcher.onDidCreate(() => void this.findProjects()),
			watcher.onDidDelete(() => void this.findProjects()),
			nbt,
			nbt.onDidCreate(() => this.refreshSoon()),
			nbt.onDidChange(() => this.refreshSoon()),
			nbt.onDidDelete(() => this.refreshSoon()),
			output.onDidChangeDocument(() => this.update()),
			vscode.workspace.onDidChangeWorkspaceFolders(() => void this.findProjects()),
			vscode.workspace.onDidSaveTextDocument((d) => this.saved(d)),
			vscode.commands.registerCommand('mms.chooseWorld', () => this.chooseWorld()),
			vscode.commands.registerCommand('mms.chooseProject', () => this.chooseProject()),
			vscode.commands.registerCommand('mms.buildToWorld', () => this.build()),
			vscode.commands.registerCommand('mms.revealWorld', () => {
				const w = this.world();
				if (w) void vscode.commands.executeCommand('revealFileInOS', vscode.Uri.file(path.join(w, 'datapacks')));
			}),
			vscode.commands.registerCommand('mms.forgetWorld', () => this.setWorld(undefined)),
			vscode.commands.registerCommand('mms.importStructure', (row?: Row) => row?.structure && this.importStructures([row.structure])),
			vscode.commands.registerCommand('mms.importStructures', () => this.chooseStructures()),
			vscode.commands.registerCommand('mms.copyStructureCall', (row?: Row) => {
				const st = row?.structure && this.structureInfo(row.structure);
				if (st) void vscode.env.clipboard.writeText(st.call).then(() => vscode.window.setStatusBarMessage(`$(check) Copied ${st.call}`, 4000));
			}),
			vscode.commands.registerCommand('mms.refreshStructures', () => this.changed.fire())
		);
		void this.findProjects();
	}

	/** The project folder (holding mms.json) the panels are about. */
	project(): string | undefined {
		const source = this.output.source;
		if (source?.scheme === 'file') {
			const dir = configDir(path.dirname(source.fsPath));
			if (dir) return dir;
		}
		const chosen = this.context.workspaceState.get<string>('mms.project');
		if (chosen && this.projects.some((p) => key(vscode.Uri.file(p)) === key(vscode.Uri.file(chosen)))) return chosen;
		return this.projects[0];
	}

	/** The world chosen for the project, if it still exists. */
	world(project = this.project()): string | undefined {
		if (!project) return undefined;
		const w = this.context.workspaceState.get<Record<string, string>>('mms.worlds')?.[key(vscode.Uri.file(project))];
		return w && fs.existsSync(path.join(w, 'level.dat')) ? w : undefined;
	}

	getTreeItem(row: Row): vscode.TreeItem {
		return row.item;
	}

	getChildren(row?: Row): Row[] {
		if (row) return row.children?.() ?? [];
		const project = this.project();
		const world = project && this.world(project);
		// No rows shows the welcome text with its buttons (package.json).
		if (!project || !world) return [];
		const rows: Row[] = [];
		const p = new vscode.TreeItem(path.basename(project));
		p.iconPath = new vscode.ThemeIcon('package');
		p.description = namespace(project);
		p.tooltip = `Project: ${path.join(project, 'mms.json')}` + (this.projects.length > 1 ? '\n\nClick to choose another project.' : '');
		if (this.projects.length > 1) p.command = { command: 'mms.chooseProject', title: 'Choose Project' };
		rows.push({ id: 'project', item: p });

		const info = readWorld(world);
		const w = new vscode.TreeItem(info?.name ?? path.basename(world));
		w.iconPath = new vscode.ThemeIcon('globe');
		w.description = path.basename(world) === info?.name ? undefined : path.basename(world);
		w.tooltip = new vscode.MarkdownString(`World: \`${world}\`\n\nClick to choose another world.`);
		w.command = { command: 'mms.chooseWorld', title: 'Choose World' };
		w.contextValue = 'world';
		rows.push({ id: 'world', item: w });

		const b = new vscode.TreeItem(this.building ? 'Building…' : 'Build into world');
		b.iconPath = new vscode.ThemeIcon(this.building ? 'sync~spin' : 'play');
		b.tooltip = `Run mms build --world, writing ${path.join('datapacks', '<namespace>')} and ${path.join('resourcepacks', 'resources.zip')} in the world`;
		if (!this.building) b.command = { command: 'mms.buildToWorld', title: 'Build into World' };
		rows.push({ id: 'build', item: b });

		const r = this.result;
		if (r && !this.building) {
			const time = r.at.toLocaleTimeString(undefined, { hour: '2-digit', minute: '2-digit' });
			const s = new vscode.TreeItem(r.ok ? `Built ${r.files} ${r.files === 1 ? 'file' : 'files'}` : r.message);
			s.description = time;
			if (r.ok) {
				s.iconPath = new vscode.ThemeIcon('check', new vscode.ThemeColor('testing.iconPassed'));
				const step = r.next ? r.next.line.charAt(0).toUpperCase() + r.next.line.slice(1) + '.' : '';
				if (r.next?.action === 'reload') s.description = `${time} · run /reload`;
				else if (r.next?.action === 'rejoin') s.description = `${time} · rejoin the world`;
				s.tooltip = new vscode.MarkdownString(
					`Built into \`${r.world}\`.\n\n` +
						(r.next
							? step
							: 'Run `/reload` in the game to load the data pack.' + (r.resources ? ' Rejoin the world to load the resource pack.' : ''))
				);
			} else {
				s.iconPath = new vscode.ThemeIcon('error', new vscode.ThemeColor('errorForeground'));
				s.tooltip = r.detail;
				s.command = { command: 'workbench.actions.view.problems', title: 'Show Problems' };
			}
			rows.push({ id: 'result', item: s });
		}

		const found = findStructures(world);
		const st = new vscode.TreeItem('Structures', found.length ? vscode.TreeItemCollapsibleState.Collapsed : vscode.TreeItemCollapsibleState.None);
		st.id = `structures:${world}`;
		st.iconPath = new vscode.ThemeIcon('symbol-structure');
		const waiting = found.filter((f) => this.structureInfo(f)?.status !== 'imported').length;
		st.description = found.length ? (waiting ? `${waiting} to import` : 'all imported') : 'none saved yet';
		st.tooltip = new vscode.MarkdownString(
			'What structure blocks saved in this world (`generated/<namespace>/structure`). Import one to use it in a template pool with `Structure(...)`; after you save it again in the game, it shows as changed.'
		);
		st.contextValue = found.length ? 'structures' : 'structures-empty';
		rows.push({ id: 'structures', item: st, children: () => found.map((f) => this.structureRow(f)) });
		return rows;
	}

	/** Where a world structure goes in the project, its status and the call that uses it. */
	private structureInfo(s: WorldStructure): { rel: string; dest: string; status: Status; call: string } | undefined {
		const project = this.project();
		if (!project) return undefined;
		const folder = vscode.workspace.getConfiguration('mms').get<string>('structuresFolder') || 'structures';
		const rel = importPath(s, namespace(project), folder);
		const dest = path.join(project, ...rel.split('/'));
		return { rel, dest, status: status(s, dest), call: structureCall(rel, s) };
	}

	private structureRow(s: WorldStructure): Row {
		const info = this.structureInfo(s)!;
		const t = new vscode.TreeItem(`${s.ns}:${s.path}`);
		t.id = `structure:${s.file}`;
		const label: Record<Status, string> = { new: 'new', imported: 'imported', changed: 'changed in world' };
		t.description = label[info.status];
		t.iconPath = new vscode.ThemeIcon(
			info.status === 'imported' ? 'pass' : info.status === 'changed' ? 'diff-modified' : 'circle-outline',
			info.status === 'changed' ? new vscode.ThemeColor('gitDecoration.modifiedResourceForeground') : undefined
		);
		const ns = namespace(this.project()!);
		t.tooltip = new vscode.MarkdownString(
			`Saved ${s.saved.toLocaleString()} in \`${s.file}\`\n\n` +
				(info.status === 'new' ? `Import copies it to \`${info.rel}\`.` : `In the project at \`${info.rel}\`.`) +
				`\n\nUse it with \`${info.call}\`` +
				(s.ns !== ns && ns ? `; its id in the pack is \`${ns}:${s.path}\`, in the project's namespace.` : '.')
		);
		t.contextValue = `structure-${info.status}`;
		// An imported structure shows its project copy; importing is the row's button.
		if (info.status !== 'new') t.command = { command: 'revealInExplorer', title: 'Reveal in Explorer', arguments: [vscode.Uri.file(info.dest)] };
		return { id: t.id, item: t, structure: s };
	}

	/** Lets the user pick structures to import, with the new and changed ones ticked. */
	private async chooseStructures() {
		const world = this.world();
		if (!world) return;
		const found = findStructures(world);
		if (!found.length) {
			void vscode.window.showInformationMessage('No structures saved in this world yet. Save one with a structure block (SAVE mode), then import it here.');
			return;
		}
		const picks = found.map((s) => {
			const info = this.structureInfo(s)!;
			return { label: `${s.ns}:${s.path}`, description: info.status === 'imported' ? 'imported' : info.status === 'changed' ? 'changed in world' : 'new', detail: `→ ${info.rel}`, picked: info.status !== 'imported', structure: s };
		});
		const chosen = await vscode.window.showQuickPick(picks, { canPickMany: true, title: 'Import structures into the project', placeHolder: 'New and changed structures are ticked' });
		if (chosen?.length) await this.importStructures(chosen.map((c) => c.structure));
	}

	/** Copies structures into the project and offers the Structure(...) call. */
	private async importStructures(list: WorldStructure[]) {
		const done: { rel: string; call: string }[] = [];
		for (const s of list) {
			const info = this.structureInfo(s);
			if (!info) return;
			try {
				// The folder setting must not lead out of the project.
				const inside = path.relative(this.project()!, info.dest);
				if (inside.startsWith('..') || path.isAbsolute(inside)) throw new Error(`"mms.structuresFolder" points outside the project (${info.rel})`);
				importStructure(s, info.dest);
				done.push(info);
				this.log(`imported ${s.file} to ${info.dest}`);
			} catch (e) {
				void vscode.window.showErrorMessage(`Could not import ${s.ns}:${s.path}: ${e instanceof Error ? e.message : String(e)}`);
			}
		}
		this.changed.fire();
		if (done.length === 1) {
			const [d] = done;
			const choice = await vscode.window.showInformationMessage(`Imported ${d.rel}. Use it in a template pool with ${d.call}.`, 'Copy call');
			if (choice) await vscode.env.clipboard.writeText(d.call);
		} else if (done.length > 1) {
			void vscode.window.setStatusBarMessage(`$(check) Imported ${done.length} structures`, 5000);
		}
	}

	/** Redraws soon, so a burst of file events redraws once. */
	private refreshSoon() {
		if (this.refreshTimer) clearTimeout(this.refreshTimer);
		this.refreshTimer = setTimeout(() => this.changed.fire(), 300);
	}

	/** Watches the chosen world's generated folder, where structure blocks save. */
	private watchWorld(world: string | undefined) {
		if (this.worldWatcher?.world === world) return;
		this.worldWatcher?.watcher.dispose();
		this.worldWatcher = undefined;
		if (!world) return;
		const watcher = vscode.workspace.createFileSystemWatcher(new vscode.RelativePattern(vscode.Uri.file(world), 'generated/**/*.nbt'));
		watcher.onDidCreate(() => this.refreshSoon());
		watcher.onDidChange(() => this.refreshSoon());
		watcher.onDidDelete(() => this.refreshSoon());
		this.worldWatcher = { world, watcher };
	}

	/** Lets the user pick a world from the saves folders, or browse to one. */
	async chooseWorld() {
		const project = this.project();
		if (!project) {
			void vscode.window.showWarningMessage('Open a folder with an mms.json to build its pack into a world.');
			return;
		}
		const extra = (vscode.workspace.getConfiguration('mms').get<string[]>('savesFolders') ?? []).map(expandHome);
		const worlds = findWorlds(savesFolders(extra));
		const current = this.world(project);
		type Pick = vscode.QuickPickItem & { world?: World; browse?: true };
		const picks: Pick[] = [{ label: '$(folder-opened) Browse for a world folder…', browse: true }];
		for (const w of worlds) {
			picks.push({
				label: w.name,
				description: [w.source, w.name === path.basename(w.dir) ? '' : path.basename(w.dir)].filter(Boolean).join(' · '),
				detail: `Played ${w.saved.toLocaleString()}`,
				picked: current !== undefined && key(vscode.Uri.file(current)) === key(vscode.Uri.file(w.dir)),
				world: w
			});
		}
		const pick = await vscode.window.showQuickPick(picks, {
			title: `Build ${path.basename(project)} into a world`,
			placeHolder: worlds.length ? 'Choose a Minecraft world (most recently played first)' : 'No worlds found in the usual saves folders; browse for one',
			matchOnDescription: true
		});
		if (!pick) return;
		let dir = pick.world?.dir;
		if (pick.browse) {
			const chosen = await vscode.window.showOpenDialog({
				canSelectFolders: true,
				canSelectFiles: false,
				openLabel: 'Use this world',
				title: 'Choose a world folder (it has a level.dat)',
				defaultUri: savesFolders(extra)[0] ? vscode.Uri.file(savesFolders(extra)[0].dir) : undefined
			});
			if (!chosen?.[0]) return;
			dir = chosen[0].fsPath;
			if (!fs.existsSync(path.join(dir, 'level.dat'))) {
				void vscode.window.showErrorMessage(`${dir} is not a Minecraft world: it has no level.dat.`);
				return;
			}
		}
		if (dir) await this.setWorld(dir, project);
	}

	private async setWorld(dir: string | undefined, project = this.project()) {
		if (!project) return;
		const worlds = { ...this.context.workspaceState.get<Record<string, string>>('mms.worlds') };
		if (dir) worlds[key(vscode.Uri.file(project))] = dir;
		else delete worlds[key(vscode.Uri.file(project))];
		await this.context.workspaceState.update('mms.worlds', worlds);
		this.result = undefined;
		this.update();
	}

	/** Lets the user pick which project the side bar is about. */
	async chooseProject() {
		await this.findProjects();
		const root = (p: string) => vscode.workspace.asRelativePath(p, true) || path.basename(p);
		const pick = await vscode.window.showQuickPick(
			this.projects.map((p) => ({ label: path.basename(p), description: root(p), dir: p })),
			{ title: 'Choose an MMS project', placeHolder: 'The folder with the mms.json to build and explore' }
		);
		if (!pick) return;
		await this.context.workspaceState.update('mms.project', pick.dir);
		// The Pack view shows the pack of the project one of these files is in.
		const [file] = await vscode.workspace.findFiles(new vscode.RelativePattern(pick.dir, '**/*.mms'), '**/node_modules/**', 1);
		if (file && configDir(path.dirname(file.fsPath)) === pick.dir) this.output.follow(file);
		this.update();
	}

	/** Builds the project into its world. */
	async build(): Promise<void> {
		const project = this.project();
		if (!project) {
			void vscode.window.showWarningMessage('Open a folder with an mms.json to build its pack into a world.');
			return;
		}
		let world = this.world(project);
		if (!world) {
			await this.chooseWorld();
			world = this.world(project);
			if (!world) return;
		}
		if (this.building) return;
		const mms = this.mms();
		if ('error' in mms) {
			void vscode.window.showErrorMessage(`Cannot build: ${mms.error} Set "mms.path" to the full path of mms.`, 'Open settings').then((c) => {
				if (c) void vscode.commands.executeCommand('workbench.action.openSettings', 'mms.path');
			});
			return;
		}
		// The build reads the files on disk.
		const here = key(vscode.Uri.file(project)) + path.sep;
		await Promise.all(vscode.workspace.textDocuments.filter((d) => d.isDirty && key(d.uri).startsWith(here)).map((d) => d.save()));
		this.building = true;
		this.update();
		const args = ['build', '--world', world, project];
		this.log(`running ${mms.command} ${args.join(' ')}`);
		const run = await new Promise<{ code: number; stdout: string; stderr: string }>((resolve) => {
			execFile(mms.command, args, { cwd: project, windowsHide: true }, (err, stdout, stderr) => {
				const code = err ? (typeof err.code === 'number' ? err.code : 1) : 0;
				if (err && typeof err.code !== 'number') stderr = `${stderr}${err.message}`;
				resolve({ code, stdout, stderr });
			});
		});
		this.building = false;
		for (const line of `${run.stdout}${run.stderr}`.split('\n').filter(Boolean)) this.log(`  ${line}`);
		const name = readWorld(world)?.name ?? path.basename(world);
		if (run.code === 0) {
			const files = [...run.stdout.matchAll(/(\d+) files?/g)].reduce((n, m) => n + Number(m[1]), 0);
			const resources = /resource pack/.test(run.stdout);
			const next = nextStep(run.stdout);
			this.result = { ok: true, world, files, resources, next, at: new Date() };
			if (next?.action === 'reload') {
				void vscode.window.showInformationMessage(`MMS: built into ${name}. Run /reload in the game to apply it.`, 'Copy /reload').then((c) => {
					if (c) void vscode.env.clipboard.writeText('/reload');
				});
			} else if (next?.action === 'rejoin') {
				const changed = /; (.*) changed\b/.exec(next.line)?.[1];
				void vscode.window.showInformationMessage(
					`MMS: built into ${name}. Exit and rejoin the world to apply it` + (changed ? ` (${changed} changed; /reload is not enough).` : '.')
				);
			} else {
				void vscode.window.setStatusBarMessage(next ? `$(check) MMS: built into ${name}; nothing changed` : `$(check) MMS: built into ${name}`, 5000);
			}
		} else {
			const errors = /(\d+) errors?; nothing written/.exec(run.stderr);
			const message = errors ? `${errors[1]} ${errors[1] === '1' ? 'error' : 'errors'}; nothing written` : 'Build failed';
			this.result = { ok: false, message, detail: run.stderr.trim(), at: new Date() };
			const last = run.stderr.trim().split('\n').pop()?.replace(/^mms: /, '') ?? '';
			void vscode.window
				.showErrorMessage(errors ? `MMS: ${message} to ${name}. Fix the problems and build again.` : `MMS: building into ${name} failed: ${last}`, errors ? 'Show problems' : 'Show output')
				.then((c) => {
					if (c === 'Show problems') void vscode.commands.executeCommand('workbench.actions.view.problems');
					else if (c === 'Show output') void vscode.commands.executeCommand('mms.showOutput');
				});
		}
		this.update();
	}

	/** With mms.buildOnSave, saving a project file rebuilds its world. */
	private saved(doc: vscode.TextDocument) {
		if (!vscode.workspace.getConfiguration('mms').get<boolean>('buildOnSave')) return;
		const project = this.project();
		if (!project || !this.world(project) || doc.uri.scheme !== 'file') return;
		if (configDir(path.dirname(doc.uri.fsPath)) !== project) return;
		// Save All saves many files at once: build once after them.
		if (this.saveTimer) clearTimeout(this.saveTimer);
		this.saveTimer = setTimeout(() => void this.build(), 300);
	}

	private async findProjects() {
		const found = await vscode.workspace.findFiles('**/mms.json', '**/node_modules/**');
		this.projects = found.map((u) => path.dirname(u.fsPath)).sort();
		this.update();
	}

	private update() {
		const project = this.project();
		void vscode.commands.executeCommand('setContext', 'mms.hasProject', project !== undefined);
		void vscode.commands.executeCommand('setContext', 'mms.hasWorld', this.world(project) !== undefined);
		this.watchWorld(this.world(project));
		this.changed.fire();
	}

	dispose() {
		if (this.saveTimer) clearTimeout(this.saveTimer);
		if (this.refreshTimer) clearTimeout(this.refreshTimer);
		this.worldWatcher?.watcher.dispose();
		for (const s of this.subscriptions) s.dispose();
	}
}

/** The nearest folder at or above dir that holds an mms.json. */
export function configDir(dir: string): string | undefined {
	for (;;) {
		if (fs.existsSync(path.join(dir, 'mms.json'))) return dir;
		const up = path.dirname(dir);
		if (up === dir) return undefined;
		dir = up;
	}
}

/** The namespace in a project's mms.json, if it can be read. */
function namespace(project: string): string | undefined {
	try {
		const ns = JSON.parse(fs.readFileSync(path.join(project, 'mms.json'), 'utf8')).namespace;
		return typeof ns === 'string' ? ns : undefined;
	} catch {
		return undefined;
	}
}

function expandHome(p: string): string {
	return p === '~' || p.startsWith('~/') || p.startsWith('~\\') ? path.join(os.homedir(), p.slice(1)) : p;
}
