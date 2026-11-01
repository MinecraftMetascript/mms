// The MMS extension starts `mms lsp` and connects VS Code to it. All
// language features (diagnostics, completion, hover, rename, formatting,
// quick fixes, semantic highlighting) come from the server.
//
// Activation never waits on the server: the commands are registered first,
// and a server that cannot start is reported with a notification that names
// the fix, while "MMS: Restart Language Server" keeps working.
import * as fs from 'fs';
import * as os from 'os';
import * as path from 'path';
import * as vscode from 'vscode';
import { LanguageClient, type LanguageClientOptions, type ServerOptions } from 'vscode-languageclient/node';

const releasesUrl = 'https://github.com/MinecraftMetascript/mms/releases';
const exe = process.platform === 'win32' ? 'mms.exe' : 'mms';

let client: LanguageClient | undefined;
let output: vscode.OutputChannel | undefined;
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

function resolveServer(): Resolved {
	const configured = vscode.workspace.getConfiguration('mms').get<string>('path')?.trim();
	if (configured) {
		const p = expand(configured);
		if (path.isAbsolute(p) && !isFile(p)) {
			return { error: `"mms.path" is set to ${p}, but there is no file there.` };
		}
		return { command: p };
	}
	const found = findMms();
	if (found) return { command: found };
	return { error: 'The mms executable was not found on your PATH or in ~/go/bin.' };
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
		documentSelector: [{ scheme: 'file', language: 'mms' }],
		outputChannel: output,
		synchronize: {
			// mms.json changes the namespace and lint settings; asset files
			// used by Texture, Sound and Structure are checked on change.
			fileEvents: [
				vscode.workspace.createFileSystemWatcher('**/mms.json'),
				vscode.workspace.createFileSystemWatcher('**/*.{png,ogg,nbt}')
			]
		}
	};
	const c = new LanguageClient('mms', 'MMS Language Server', serverOptions, clientOptions);
	client = c;
	try {
		await c.start();
		log('server started');
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
	output = vscode.window.createOutputChannel('MMS Language Server');
	context.subscriptions.push(
		output,
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
