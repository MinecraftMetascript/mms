// Prints the extension version a CI build publishes, from the highest
// vMAJOR.MINOR.PATCH tag merged into HEAD (the repository must be fetched with
// its tags). Tags with a suffix, such as v1.0.0-rc.1, are left out.
//
//   node scripts/version.mjs              the tag itself: v0.5.0 -> 0.5.0
//   node scripts/version.mjs --pre-release N
//                                         one minor ahead, patch N:
//                                         v0.5.0 -> 0.6.N
//
// A pre-release is always newer than the release it follows, so VS Code
// moves people on the pre-release channel to it. N is the workflow's run
// number, which only goes up.
import { execFileSync } from 'node:child_process';

const stable = /^v(\d+)\.(\d+)\.(\d+)$/;
const versions = execFileSync('git', ['tag', '--merged', 'HEAD', '--list', 'v*'], { encoding: 'utf8' })
	.split('\n')
	.map((t) => stable.exec(t.trim()))
	.filter(Boolean)
	.map((m) => m.slice(1).map(Number))
	.sort((a, b) => a[0] - b[0] || a[1] - b[1] || a[2] - b[2]);
if (versions.length === 0) {
	console.error('no vMAJOR.MINOR.PATCH tag is merged into HEAD');
	process.exit(1);
}
const [major, minor, patch] = versions[versions.length - 1];
const args = process.argv.slice(2);
if (args[0] === '--pre-release') {
	const n = Number(args[1]);
	if (!Number.isInteger(n) || n < 1) {
		console.error('usage: version.mjs --pre-release <positive build number>');
		process.exit(2);
	}
	console.log(`${major}.${minor + 1}.${n}`);
} else {
	console.log(`${major}.${minor}.${patch}`);
}
