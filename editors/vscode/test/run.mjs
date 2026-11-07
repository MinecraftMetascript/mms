// Runs the extension host tests: starts a real VS Code with the extension,
// opens a small MMS project, and runs test/suite.ts in it against the
// language server. MMS_PATH names the mms binary to use (go build
// ./cmd/mms); on Linux, run it under xvfb-run.
import { runTests } from '@vscode/test-electron';
import * as fs from 'node:fs';
import * as os from 'node:os';
import * as path from 'node:path';
import { fileURLToPath } from 'node:url';

const here = path.dirname(fileURLToPath(import.meta.url));
const ext = path.resolve(here, '..');
const mms = process.env.MMS_PATH;
if (!mms || !fs.existsSync(mms)) {
	console.error('Set MMS_PATH to an mms binary, such as the output of `go build -o /tmp/mms ./cmd/mms`.');
	process.exit(2);
}

// The project the tests open.
const dir = fs.mkdtempSync(path.join(os.tmpdir(), 'mms-ext-test-'));
const write = (rel, text) => {
	fs.mkdirSync(path.dirname(path.join(dir, rel)), { recursive: true });
	fs.writeFileSync(path.join(dir, rel), text);
};
write('mms.json', JSON.stringify({ namespace: 'demo' }, null, 2) + '\n');
write('.vscode/settings.json', JSON.stringify({ 'mms.path': path.resolve(mms) }, null, 2) + '\n');
write(
	'main.mms',
	[
		'const unused_value = 1',
		'const wave = 3',
		'/** Doubles a number. */',
		'Template twice(x: int): int = x * 2',
		'Template ring(p: Id<particle>): Commands {',
		'  /particle ${p} ~ ~1 ~',
		'}',
		'CustomItem ruby { base = diamond, name = "Ruby" }',
		'Function start {',
		'  ring(flame)',
		'  /say ${wave}',
		'}',
		'DensityFunction d = twice(wave)',
		''
	].join('\n')
);
write('text/welcome.mmstext', 'Welcome, **traveller**!\n');

try {
	await runTests({
		extensionDevelopmentPath: ext,
		extensionTestsPath: path.join(ext, 'out-test', 'suite.js'),
		launchArgs: [dir, '--disable-extensions', '--disable-workspace-trust', '--skip-welcome', '--skip-release-notes']
	});
} catch (e) {
	console.error('Extension host tests failed:', e instanceof Error ? e.message : e);
	process.exit(1);
} finally {
	fs.rmSync(dir, { recursive: true, force: true });
}
