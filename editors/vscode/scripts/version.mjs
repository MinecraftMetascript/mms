// Prints the extension version a CI build publishes, from the latest v* tag
// reachable from HEAD (the repository must be fetched with its tags).
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

const tag = execFileSync('git', ['describe', '--tags', '--abbrev=0', '--match', 'v[0-9]*'], { encoding: 'utf8' }).trim();
const m = /^v(\d+)\.(\d+)\.(\d+)$/.exec(tag);
if (!m) {
	console.error(`latest tag ${tag} is not vMAJOR.MINOR.PATCH`);
	process.exit(1);
}
const [major, minor, patch] = m.slice(1).map(Number);
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
