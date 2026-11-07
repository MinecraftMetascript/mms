// Tests puts the project's Test declarations in VS Code's Testing view and
// runs them with the language server's `mms/runTests`, the interpreter
// `mms test` uses, so the results match the command line. Tests are found
// by their `Test name {` lines; a run reports a failing expect at its line.
import * as vscode from 'vscode';
import type { LanguageClient } from 'vscode-languageclient/node';
import type { Range, TestOutcome } from './messages';

const testLine = /^[ \t]*Test[ \t]+([A-Za-z_][A-Za-z0-9_]*)/gm;

function toRange(r: Range): vscode.Range {
	return new vscode.Range(r.start.line, r.start.character, r.end.line, r.end.character);
}

export class Tests implements vscode.Disposable {
	private readonly ctrl = vscode.tests.createTestController('mms', 'MMS');
	private readonly subs: vscode.Disposable[] = [this.ctrl];

	constructor(private readonly client: () => LanguageClient | undefined) {
		this.ctrl.resolveHandler = async () => {
			for (const uri of await vscode.workspace.findFiles('**/*.mms', '**/node_modules/**')) {
				this.scan(uri, Buffer.from(await vscode.workspace.fs.readFile(uri)).toString('utf8'));
			}
		};
		this.ctrl.refreshHandler = () => this.ctrl.resolveHandler?.(undefined);
		this.ctrl.createRunProfile('Run', vscode.TestRunProfileKind.Run, (req, token) => this.run(req, token), true);
		const watcher = vscode.workspace.createFileSystemWatcher('**/*.mms');
		this.subs.push(
			watcher,
			watcher.onDidDelete((uri) => this.ctrl.items.delete(uri.toString())),
			vscode.workspace.onDidChangeTextDocument((e) => {
				if (e.document.languageId === 'mms') this.scan(e.document.uri, e.document.getText());
			}),
			vscode.workspace.onDidOpenTextDocument((d) => {
				if (d.languageId === 'mms') this.scan(d.uri, d.getText());
			})
		);
	}

	/** scan lists one file's tests, under an item for the file. */
	private scan(uri: vscode.Uri, text: string) {
		const found: vscode.TestItem[] = [];
		let file = this.ctrl.items.get(uri.toString());
		testLine.lastIndex = 0;
		for (let m = testLine.exec(text); m; m = testLine.exec(text)) {
			if (!file) {
				file = this.ctrl.createTestItem(uri.toString(), vscode.workspace.asRelativePath(uri), uri);
			}
			const item = file.children.get(m[1]) ?? this.ctrl.createTestItem(m[1], m[1], uri);
			const before = text.slice(0, m.index + m[0].length - m[1].length);
			const line = before.split('\n').length - 1;
			const col = before.length - before.lastIndexOf('\n') - 1;
			item.range = new vscode.Range(line, col, line, col + m[1].length);
			found.push(item);
		}
		if (found.length === 0) {
			this.ctrl.items.delete(uri.toString());
			return;
		}
		file!.children.replace(found);
		this.ctrl.items.add(file!);
	}

	private async run(req: vscode.TestRunRequest, token: vscode.CancellationToken) {
		const run = this.ctrl.createTestRun(req);
		const items: vscode.TestItem[] = [];
		const add = (it: vscode.TestItem) => {
			if (it.children.size === 0) items.push(it);
			else it.children.forEach(add);
		};
		if (req.include) req.include.forEach(add);
		else this.ctrl.items.forEach(add);
		const skip = new Set(req.exclude ?? []);
		const client = this.client();
		for (const it of items) {
			if (token.isCancellationRequested || skip.has(it)) {
				run.skipped(it);
				continue;
			}
			if (!client?.isRunning() || !it.uri) {
				run.errored(it, new vscode.TestMessage('The MMS language server is not running.'));
				continue;
			}
			run.started(it);
			const started = Date.now();
			try {
				const out = await client.sendRequest<{ tests: TestOutcome[]; error?: string }>('mms/runTests', { uri: it.uri.toString(), name: it.id });
				const t = out.tests.find((x) => x.name === it.id);
				for (const line of t?.said ?? []) run.appendOutput(`${it.id}: ${line}\r\n`, undefined, it);
				if (out.error) {
					run.errored(it, new vscode.TestMessage(out.error));
				} else if (!t) {
					run.skipped(it);
				} else if (t.pass) {
					run.passed(it, Date.now() - started);
				} else {
					const msg = new vscode.TestMessage(t.message ?? 'failed');
					msg.location = new vscode.Location(vscode.Uri.parse(t.uri), toRange(t.expectRange ?? t.range));
					run.failed(it, msg, Date.now() - started);
				}
			} catch (e) {
				run.errored(it, new vscode.TestMessage(e instanceof Error ? e.message : String(e)));
			}
		}
		run.end();
	}

	dispose() {
		for (const s of this.subs) s.dispose();
	}
}
