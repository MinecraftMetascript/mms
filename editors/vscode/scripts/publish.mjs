// Publishes every .vsix in a folder to the VS Code Marketplace and Open VSX.
//
//   node scripts/publish.mjs <folder> [--pre-release]
//
// VSCE_PAT (an Azure DevOps token with Marketplace Manage) publishes to the
// Marketplace and OVSX_PAT (an Open VSX access token) to Open VSX. A store
// whose token is not set is skipped with a note, so the workflows run before
// the secrets exist. Versions already published are skipped, so a failed
// run can be re-run.
import { execFileSync } from 'node:child_process';
import fs from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';

const args = process.argv.slice(2);
const preRelease = args.includes('--pre-release');
const dir = args.find((a) => !a.startsWith('--'));
if (!dir) {
	console.error('usage: publish.mjs <folder> [--pre-release]');
	process.exit(2);
}
const vsixes = fs
	.readdirSync(dir)
	.filter((f) => f.endsWith('.vsix'))
	.map((f) => path.resolve(dir, f));
if (vsixes.length === 0) {
	console.error(`no .vsix in ${dir}`);
	process.exit(1);
}

const run = (cmd, argv) => execFileSync(cmd, argv, { stdio: 'inherit' });

if (process.env.VSCE_PAT) {
	const vsce = path.join(path.dirname(createRequire(import.meta.url).resolve('@vscode/vsce/package.json')), 'vsce');
	// The token is read from VSCE_PAT, so it is not on the command line.
	run(process.execPath, [vsce, 'publish', '--skip-duplicate', ...(preRelease ? ['--pre-release'] : []), '--packagePath', ...vsixes]);
} else {
	console.log('VSCE_PAT is not set: skipping the VS Code Marketplace.');
}

if (process.env.OVSX_PAT) {
	// ovsx reads OVSX_PAT, and the pre-release flag from each package.
	for (const vsix of vsixes) run('npx', ['--yes', 'ovsx@1.3.0', 'publish', vsix, '--skip-duplicate']);
} else {
	console.log('OVSX_PAT is not set: skipping Open VSX.');
}
