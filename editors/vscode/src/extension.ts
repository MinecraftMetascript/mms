// The MMS extension starts `mms lsp` and connects VS Code to it. All
// language features (diagnostics, completion, hover, rename, formatting,
// quick fixes, semantic highlighting) come from the server.
import * as vscode from 'vscode';
import { LanguageClient, type LanguageClientOptions, type ServerOptions } from 'vscode-languageclient/node';

let client: LanguageClient | undefined;

function serverOptions(): ServerOptions {
	const configured = vscode.workspace.getConfiguration('mms').get<string>('path')?.trim();
	const command = configured || 'mms';
	return { command, args: ['lsp'] };
}

async function start(context: vscode.ExtensionContext) {
	const clientOptions: LanguageClientOptions = {
		documentSelector: [{ scheme: 'file', language: 'mms' }],
		synchronize: {
			// mms.json changes the namespace and lint settings; asset files
			// used by Texture, Sound and Structure are checked on change.
			fileEvents: [
				vscode.workspace.createFileSystemWatcher('**/mms.json'),
				vscode.workspace.createFileSystemWatcher('**/*.{png,ogg,nbt}')
			]
		}
	};
	client = new LanguageClient('mms', 'MMS Language Server', serverOptions(), clientOptions);
	try {
		await client.start();
	} catch (e) {
		const msg = e instanceof Error ? e.message : String(e);
		const choice = await vscode.window.showErrorMessage(
			`Could not start the MMS language server (${msg}). Install mms from the releases page, or set "mms.path".`,
			'Open releases',
			'Open settings'
		);
		if (choice === 'Open releases') {
			vscode.env.openExternal(vscode.Uri.parse('https://github.com/MinecraftMetascript/mms/releases'));
		} else if (choice === 'Open settings') {
			vscode.commands.executeCommand('workbench.action.openSettings', 'mms.path');
		}
	}
	context.subscriptions.push({ dispose: () => void client?.stop() });
}

export async function activate(context: vscode.ExtensionContext) {
	context.subscriptions.push(
		vscode.commands.registerCommand('mms.restartServer', async () => {
			await client?.stop();
			await start(context);
		}),
		vscode.workspace.onDidChangeConfiguration(async (e) => {
			if (e.affectsConfiguration('mms.path')) {
				await client?.stop();
				await start(context);
			}
		})
	);
	await start(context);
}

export async function deactivate() {
	await client?.stop();
}
