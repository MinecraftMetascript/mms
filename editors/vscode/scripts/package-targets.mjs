// Packages one platform-specific .vsix per VS Code target, each with the mms
// binary for that platform in bin/, so users do not install mms themselves.
// The Marketplace (and `code --install-extension`) picks the right one.
//
//   node scripts/package-targets.mjs                 every target below
//   node scripts/package-targets.mjs linux-x64 ...   just these
//
// Run `npm run build` first (`npm run package:targets` does). Needs Go: the
// binaries are cross-compiled from ../../engine with CGO off. MMS_VERSION
// (such as v1.0.0) is stamped into the binary for `mms version`.
import { execFileSync } from 'node:child_process';
import fs from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const engine = path.resolve(root, '../../engine');
const bin = path.join(root, 'bin');

// VS Code target -> Go GOOS/GOARCH. Alpine gets the same static binary.
const targets = {
	'linux-x64': ['linux', 'amd64'],
	'linux-arm64': ['linux', 'arm64'],
	'alpine-x64': ['linux', 'amd64'],
	'alpine-arm64': ['linux', 'arm64'],
	'darwin-x64': ['darwin', 'amd64'],
	'darwin-arm64': ['darwin', 'arm64'],
	'win32-x64': ['windows', 'amd64'],
	'win32-arm64': ['windows', 'arm64']
};

const wanted = process.argv.slice(2);
for (const t of wanted) {
	if (!(t in targets)) {
		console.error(`unknown target ${t}; known: ${Object.keys(targets).join(', ')}`);
		process.exit(2);
	}
}

// vsce's own entry point, run with this node: no shell, so paths with spaces
// work on Windows. The workspace may install it in the repository's
// node_modules, so resolve it the way Node does.
const vsce = path.join(path.dirname(createRequire(import.meta.url).resolve('@vscode/vsce/package.json')), 'vsce');
const ldflags = ['-s', '-w'];
if (process.env.MMS_VERSION) ldflags.push(`-X main.version=${process.env.MMS_VERSION}`);

try {
	for (const target of wanted.length ? wanted : Object.keys(targets)) {
		const [goos, goarch] = targets[target];
		const exe = goos === 'windows' ? 'mms.exe' : 'mms';
		fs.rmSync(bin, { recursive: true, force: true });
		fs.mkdirSync(bin);
		console.log(`${target}: go build (${goos}/${goarch})`);
		execFileSync('go', ['build', '-trimpath', '-ldflags', ldflags.join(' '), '-o', path.join(bin, exe), './cmd/mms'], {
			cwd: engine,
			stdio: 'inherit',
			env: { ...process.env, GOOS: goos, GOARCH: goarch, CGO_ENABLED: '0' }
		});
		fs.copyFileSync(path.join(engine, 'THIRD_PARTY_NOTICES.txt'), path.join(bin, 'THIRD_PARTY_NOTICES.txt'));
		if (goos !== 'windows') fs.chmodSync(path.join(bin, exe), 0o755);
		execFileSync(process.execPath, [vsce, 'package', '--skip-license', '--no-dependencies', '--target', target], { cwd: root, stdio: 'inherit' });
	}
} finally {
	// The universal .vsix (`npm run package`) must not pick up a stray binary.
	fs.rmSync(bin, { recursive: true, force: true });
}
