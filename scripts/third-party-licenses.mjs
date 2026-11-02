#!/usr/bin/env node
// Writes the license notices of the npm packages a build bundles, so the
// web editor and the VS Code extension ship them as MIT and friends ask.
//
//   node third-party-licenses.mjs <package dir> <out file> <root>... [--notice <file>]...
//
// Each root (such as deepslate) is resolved from the package dir the way
// Node does, then its dependencies are followed. --notice appends a file
// as is, for data that is not an npm package (vanilla-mcdoc in the engine).
import { existsSync, readdirSync, readFileSync, mkdirSync, writeFileSync } from 'node:fs';
import { dirname, join, resolve } from 'node:path';

const args = process.argv.slice(2);
const notices = [];
for (let i = args.indexOf('--notice'); i !== -1; i = args.indexOf('--notice')) notices.push(...args.splice(i, 2).slice(1));
const [pkgDir, outFile, ...roots] = args;
if (!pkgDir || !outFile || !roots.length) {
	console.error('usage: third-party-licenses.mjs <package dir> <out file> <root>... [--notice <file>]...');
	process.exit(2);
}

/** The directory of package `name` as Node would load it from `from`. */
function find(name, from) {
	for (let dir = resolve(from); ; dir = dirname(dir)) {
		const p = join(dir, 'node_modules', name);
		if (existsSync(join(p, 'package.json'))) return p;
		if (dirname(dir) === dir) return null;
	}
}

const seen = new Map(); // name@version -> { name, version, license, text, homepage }
const missing = [];
function visit(name, from) {
	const dir = find(name, from);
	if (!dir) return missing.push(name);
	const pkg = JSON.parse(readFileSync(join(dir, 'package.json'), 'utf8'));
	const key = `${pkg.name}@${pkg.version}`;
	if (seen.has(key)) return;
	const file = readdirSync(dir).find((f) => /^(licen[sc]e|copying)(\.(md|txt))?$/i.test(f));
	const repo = typeof pkg.repository === 'string' ? pkg.repository : pkg.repository?.url;
	seen.set(key, {
		name: pkg.name,
		version: pkg.version,
		license: typeof pkg.license === 'string' ? pkg.license : (pkg.license?.type ?? 'UNKNOWN'),
		text: file ? readFileSync(join(dir, file), 'utf8').trim() : null,
		url: pkg.homepage ?? repo?.replace(/^git\+/, '').replace(/\.git$/, '')
	});
	for (const dep of Object.keys(pkg.dependencies ?? {})) visit(dep, dir);
}
for (const r of roots) visit(r, pkgDir);
if (missing.length) {
	console.error(`third-party-licenses: not installed: ${[...new Set(missing)].join(', ')} (run npm ci)`);
	process.exit(1);
}

const rule = '-'.repeat(72);
const parts = [
	'Third-party software in this build, and the notices its licenses ask to keep with it.',
	'Previews are rendered with Deepslate by Misode: https://github.com/misode/deepslate'
];
for (const n of notices) parts.push(rule, readFileSync(n, 'utf8').trim());
for (const p of [...seen.values()].sort((a, b) => a.name.localeCompare(b.name))) {
	const head = `${p.name} ${p.version} (${p.license})${p.url ? `\n${p.url}` : ''}`;
	parts.push(rule, p.text ? `${head}\n\n${p.text}` : `${head}\n\n(No license file in the package; see its license field above.)`);
}
mkdirSync(dirname(outFile), { recursive: true });
writeFileSync(outFile, parts.join('\n\n') + '\n');
console.log(`third-party-licenses: ${seen.size} packages -> ${outFile}`);
