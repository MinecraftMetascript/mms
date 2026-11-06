// The MMS extension starts `mms lsp` and connects VS Code to it. All
// language features (diagnostics, completion, hover, rename, formatting,
// quick fixes, semantic highlighting) come from the server.
//
// Activation never waits on the server: the commands are registered first,
// and a server that cannot start is reported with a notification that names
// the fix, while "MMS: Restart Language Server" keeps working.
//
// The preview panel (preview.ts) shows the web editor's worldgen previews.
// The MMS side bar has the World view (world.ts), which builds the pack
// into a Minecraft world, the Pack view (packView.ts), which lists the
// generated files, the Structures view (also world.ts), which imports
// structures saved in the game, and the Jigsaw view (jigsaw.ts), which
// copies give commands for set-up jigsaw blocks. The panels read the pack from output.ts.
// Test declarations show in VS Code's Testing view (tests.ts), and an
// Expand lens shows what a template call or blueprint writes (expand.ts).
import * as fs from 'fs';
import * as os from 'os';
import * as path from 'path';
import * as vscode from 'vscode';
import { LanguageClient, type LanguageClientOptions, type ServerOptions } from 'vscode-languageclient/node';
import { Expands } from './expand';
import { JigsawView } from './jigsaw';
import { PackOutput } from './output';
import { PackView } from './packView';
import { Previews } from './preview';
import { Tests } from './tests';
import { WorldView } from './world';

const releasesUrl = 'https://github.com/MinecraftMetascript/mms/releases';
const exe = process.platform === 'win32' ? 'mms.exe' : 'mms';

// bundled is the mms that platform-specific builds of the extension carry in
// bin/ (see scripts/package-targets.mjs). Set on activation.
let bundled: string | undefined;

let client: LanguageClient | undefined;
let output: vscode.OutputChannel | undefined;
let pack: PackOutput | undefined;
// Starts and restarts run one at a time, so a restart pressed while the
// server is still starting does not race it.
let queue: Promise<void> = Promise.resolve();

function log(line: string) {
	output?.appendLine(`[mms] ${line}`);
}

function isFile(p: string): boolean {
	try {
		return fs.statSync(p).isFile();
	} catch {
		return false;
	}
}

// expand resolves ~ and ${workspaceFolder} in mms.path, and makes a relative
// path relative to the first workspace folder.
function expand(p: string): string {
	const folder = vscode.workspace.workspaceFolders?.[0]?.uri.fsPath;
	p = p.replace(/\$\{workspaceFolder\}/g, folder ?? '');
	if (p === '~' || p.startsWith('~/') || p.startsWith('~\\')) {
		p = path.join(os.homedir(), p.slice(1));
	}
	if (!path.isAbsolute(p) && folder && (p.includes('/') || p.includes('\\'))) {
		p = path.join(folder, p);
	}
	return p;
}

// findMms looks for the mms executable on PATH, then where `go install`
// puts it. VS Code started from the macOS Dock or Finder may not see the
// PATH your shell sets up, so ~/go/bin is checked by hand.
function findMms(): string | undefined {
	const dirs = (process.env.PATH ?? '').split(path.delimiter).filter(Boolean);
	if (process.env.GOBIN) dirs.push(process.env.GOBIN);
	for (const gopath of (process.env.GOPATH ?? '').split(path.delimiter).filter(Boolean)) {
		dirs.push(path.join(gopath, 'bin'));
	}
	dirs.push(path.join(os.homedir(), 'go', 'bin'));
	if (process.platform === 'darwin') dirs.push('/opt/homebrew/bin', '/usr/local/bin');
	for (const dir of dirs) {
		const candidate = path.join(dir, exe);
		if (isFile(candidate)) return candidate;
	}
	return undefined;
}

type Resolved = { command: string } | { error: string };

// findBundled returns the mms shipped inside the extension, if this build of
// the extension has one. The universal build (and a build from source) has
// none and falls back to PATH.
function findBundled(): string | undefined {
	if (!bundled || !isFile(bundled)) return undefined;
	if (process.platform !== 'win32') {
		// Unpacking a .vsix can drop the executable bit.
		try {
			fs.accessSync(bundled, fs.constants.X_OK);
		} catch {
			try {
				fs.chmodSync(bundled, 0o755);
			} catch (e) {
				log(`cannot make ${bundled} executable: ${e instanceof Error ? e.message : String(e)}`);
				return undefined;
			}
		}
	}
	return bundled;
}

// resolveServer picks the mms to run: mms.path if set, then the one bundled
// with the extension, then mms on PATH (or where `go install` puts it).
function resolveServer(): Resolved {
	const configured = vscode.workspace.getConfiguration('mms').get<string>('path')?.trim();
	if (configured) {
		const p = expand(configured);
		if (path.isAbsolute(p) && !isFile(p)) {
			return { error: `"mms.path" is set to ${p}, but there is no file there.` };
		}
		return { command: p };
	}
	const own = findBundled();
	if (own) return { command: own };
	const found = findMms();
	if (found) return { command: found };
	return { error: 'This build of the extension does not include mms for your platform, and the mms executable was not found on your PATH or in ~/go/bin.' };
}

function reportStartFailure(reason: string) {
	log(reason);
	// Not awaited: a notification the user ignores must not hold up anything.
	void vscode.window
		.showErrorMessage(
			`MMS language server not started. ${reason} Set "mms.path" to the full path of mms (for example the binary from \`go build ./cmd/mms\`), or install it from the releases page, then run "MMS: Restart Language Server".`,
			'Open settings',
			'Open releases',
			'Show output'
		)
		.then((choice) => {
			if (choice === 'Open settings') {
				void vscode.commands.executeCommand('workbench.action.openSettings', 'mms.path');
			} else if (choice === 'Open releases') {
				void vscode.env.openExternal(vscode.Uri.parse(releasesUrl));
			} else if (choice === 'Show output') {
				output?.show(true);
			}
		});
}

async function stop() {
	const c = client;
	client = undefined;
	if (!c) return;
	try {
		// A client whose start failed cannot be stopped, only dropped.
		if (c.needsStop()) await c.stop();
	} catch (e) {
		log(`stopping the server: ${e instanceof Error ? e.message : String(e)}`);
	}
}

/**
 * Puts the cursor on a name and opens the rename box, as after an
 * "Extract to ... declaration" action: the server picks a default name and
 * this lets the user type a better one straight away.
 */
async function renameAt(uri: string, position: { line: number; character: number }) {
	const doc = await vscode.workspace.openTextDocument(vscode.Uri.parse(uri));
	const shown = vscode.window.visibleTextEditors.find((e) => e.document === doc);
	const editor = await vscode.window.showTextDocument(doc, { viewColumn: shown?.viewColumn });
	const at = new vscode.Position(position.line, position.character);
	editor.selection = new vscode.Selection(at, at);
	editor.revealRange(new vscode.Range(at, at));
	await vscode.commands.executeCommand('editor.action.rename');
}

async function start() {
	await stop();
	const resolved = resolveServer();
	if ('error' in resolved) {
		reportStartFailure(resolved.error);
		return;
	}
	log(`starting ${resolved.command} lsp`);
	const serverOptions: ServerOptions = { command: resolved.command, args: ['lsp'] };
	const clientOptions: LanguageClientOptions = {
		documentSelector: [
			{ scheme: 'file', language: 'mms' },
			{ scheme: 'file', language: 'mmsbook' }
		],
		outputChannel: output,
		// Extract-to-declaration actions then call mms.renameAt on the new name.
		initializationOptions: { renameAfterExtract: true },
		// The server registers its own file watchers (mms.json, sources
		// and asset files), so no synchronize.fileEvents here.
		middleware: {
			// Each analysis ends with diagnostics; the preview then fetches the new pack.
			handleDiagnostics(uri, diagnostics, next) {
				next(uri, diagnostics);
				pack?.scheduleRefresh();
			}
		}
	};
	const c = new LanguageClient('mms', 'MMS Language Server', serverOptions, clientOptions);
	client = c;
	try {
		await c.start();
		log('server started');
		pack?.scheduleRefresh();
	} catch (e) {
		const msg = e instanceof Error ? e.message : String(e);
		reportStartFailure(`Running ${resolved.command} failed (${msg}).`);
	}
}

function restart(): Promise<void> {
	queue = queue.then(start, start);
	return queue;
}

export function activate(context: vscode.ExtensionContext) {
	bundled = path.join(context.extensionPath, 'bin', exe);
	output = vscode.window.createOutputChannel('MMS Language Server');
	pack = new PackOutput(() => client);
	const previews = new Previews(context, pack);
	const world = new WorldView(context, pack, resolveServer, (line) => log(line));
	context.subscriptions.push(
		output,
		pack,
		previews,
		world,
		new PackView(pack, () => world.project()),
		new JigsawView(context, pack, () => world.project()),
		new Tests(() => client),
		new Expands(() => client, pack.onDidChange),
		vscode.commands.registerCommand('mms.showPreview', () => previews.show()),
		vscode.commands.registerCommand('mms.previewDeclaration', (path: string, page?: number) => previews.show(path, page)),
		vscode.commands.registerCommand('mms.renameAt', renameAt),
		vscode.commands.registerCommand('mms.restartServer', () => restart()),
		vscode.commands.registerCommand('mms.showOutput', () => output?.show(true)),
		vscode.workspace.onDidChangeConfiguration((e) => {
			if (e.affectsConfiguration('mms.path')) void restart();
		}),
		{ dispose: () => void stop() }
	);
	void restart();
}

export async function deactivate() {
	await queue.catch(() => undefined);
	await stop();
}
