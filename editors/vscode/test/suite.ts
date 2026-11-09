// The extension host tests: the editor features of the language overhaul,
// checked through VS Code's own commands against the real language server.
// test/run.mjs starts VS Code on a small project and calls run().
import * as assert from 'node:assert/strict';
import * as vscode from 'vscode';

type Test = { name: string; fn: () => Promise<void> };
const tests: Test[] = [];
const test = (name: string, fn: () => Promise<void>) => tests.push({ name, fn });

let doc: vscode.TextDocument;

/** The position of the first occurrence of text in the document, plus delta characters. */
function at(text: string, delta = 0): vscode.Position {
	const off = doc.getText().indexOf(text);
	assert.ok(off >= 0, `"${text}" is not in the document`);
	return doc.positionAt(off + delta);
}

/** Polls f until it returns a value that passes ok, or fails after a while. */
async function until<T>(what: string, f: () => Thenable<T> | T, ok: (v: T) => boolean, ms = 60_000): Promise<T> {
	const end = Date.now() + ms;
	let last: T | undefined;
	while (Date.now() < end) {
		try {
			last = await f();
			if (ok(last)) return last;
		} catch {
			// The server may not be ready yet.
		}
		await new Promise((r) => setTimeout(r, 250));
	}
	throw new Error(`timed out waiting for ${what}; last: ${JSON.stringify(last)}`);
}

test('the outline lists consts, Templates and declarations', async () => {
	const syms = await until(
		'document symbols',
		() => vscode.commands.executeCommand<vscode.DocumentSymbol[]>('vscode.executeDocumentSymbolProvider', doc.uri),
		(s) => (s?.length ?? 0) >= 5
	);
	const names = syms.map((s) => s.name);
	for (const n of ['wave', 'twice', 'ring', 'start']) assert.ok(names.includes(n), `outline lacks ${n}: ${names}`);
});

test('unused consts are reported', async () => {
	const diags = await until(
		'diagnostics',
		() => vscode.languages.getDiagnostics(doc.uri),
		(d) => d.some((x) => String(x.message).includes('unused_value'))
	);
	const d = diags.find((x) => x.message.includes('unused_value'))!;
	assert.ok(d.tags?.includes(vscode.DiagnosticTag.Unnecessary));
});

test('hover on a Template call shows its signature and doc', async () => {
	const hovers = await vscode.commands.executeCommand<vscode.Hover[]>('vscode.executeHoverProvider', doc.uri, at('twice(wave)', 1));
	const text = hovers.flatMap((h) => h.contents.map((c) => (typeof c === 'string' ? c : c.value))).join('\n');
	assert.match(text, /Template twice\(x: int\): int/);
	assert.match(text, /Doubles a number/);
});

test('signature help covers Template parameters', async () => {
	const help = await vscode.commands.executeCommand<vscode.SignatureHelp>('vscode.executeSignatureHelpProvider', doc.uri, at('twice(wave)', 6));
	assert.equal(help?.signatures[0]?.label, 'twice(x: int): int');
});

test('go to definition on a built-in opens the prelude read-only', async () => {
	const defs = await vscode.commands.executeCommand<(vscode.Location | vscode.LocationLink)[]>('vscode.executeDefinitionProvider', doc.uri, at('CustomItem', 2));
	assert.equal(defs.length, 1);
	const uri = 'targetUri' in defs[0] ? defs[0].targetUri : defs[0].uri;
	assert.equal(uri.scheme, 'mms-prelude');
	const prelude = await vscode.workspace.openTextDocument(uri);
	assert.match(prelude.getText(), /Blueprint CustomItem/);
});

test('go to definition on a vanilla id opens its JSON read-only', async () => {
	const defs = await vscode.commands.executeCommand<(vscode.Location | vscode.LocationLink)[]>('vscode.executeDefinitionProvider', doc.uri, at('Noise(ridge', 7));
	assert.equal(defs.length, 1);
	const uri = 'targetUri' in defs[0] ? defs[0].targetUri : defs[0].uri;
	assert.equal(uri.toString(), 'mms-vanilla:/data/minecraft/worldgen/noise/ridge.json');
	const file = await vscode.workspace.openTextDocument(uri);
	assert.match(file.getText(), /"base_octave"/);
});

test('rename of a const edits its declaration and every use', async () => {
	const edit = await vscode.commands.executeCommand<vscode.WorkspaceEdit>('vscode.executeDocumentRenameProvider', doc.uri, at('wave', 1), 'level');
	const edits = edit.get(doc.uri);
	// const wave, ${wave} in the command, and twice(wave).
	assert.equal(edits.length, 3, JSON.stringify(edits));
});

test('completion offers the new top-level keywords', async () => {
	const list = await vscode.commands.executeCommand<vscode.CompletionList>('vscode.executeCompletionItemProvider', doc.uri, new vscode.Position(doc.lineCount - 1, 0));
	const labels = list.items.map((i) => (typeof i.label === 'string' ? i.label : i.label.label));
	for (const k of ['const', 'Template', 'Blueprint', 'CustomItem']) assert.ok(labels.includes(k), `completion lacks ${k}`);
});

test('a Commands Template call and a Blueprint instance have Expand lenses', async () => {
	const lenses = await until(
		'Expand lenses',
		() => vscode.commands.executeCommand<vscode.CodeLens[]>('vscode.executeCodeLensProvider', doc.uri, 20),
		(ls) => ls.filter((l) => l.command?.title === 'Expand').length >= 2
	);
	const lines = lenses.filter((l) => l.command?.title === 'Expand').map((l) => l.range.start.line);
	assert.ok(lines.includes(at('ring(flame)').line), `no lens on ring(flame): ${lines}`);
	assert.ok(lines.includes(at('CustomItem').line), `no lens on CustomItem: ${lines}`);
});

test('semantic tokens know parameters', async () => {
	const legend = await vscode.commands.executeCommand<vscode.SemanticTokensLegend>('vscode.provideDocumentSemanticTokensLegend', doc.uri);
	assert.ok(legend.tokenTypes.includes('parameter'));
});

test('.mmstext files open as MMS Text', async () => {
	const [file] = await vscode.workspace.findFiles('text/welcome.mmstext');
	const text = await vscode.workspace.openTextDocument(file);
	assert.equal(text.languageId, 'mmstext');
});

export async function run(): Promise<void> {
	const [main] = await vscode.workspace.findFiles('main.mms');
	doc = await vscode.workspace.openTextDocument(main);
	await vscode.window.showTextDocument(doc);
	const failed: string[] = [];
	for (const t of tests) {
		try {
			await t.fn();
			console.log(`  ok   ${t.name}`);
		} catch (e) {
			failed.push(t.name);
			console.log(`  FAIL ${t.name}\n       ${e instanceof Error ? e.message : String(e)}`);
		}
	}
	console.log(`${tests.length - failed.length} passed, ${failed.length} failed`);
	if (failed.length) throw new Error(`${failed.length} extension host test(s) failed: ${failed.join('; ')}`);
}
