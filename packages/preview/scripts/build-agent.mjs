// Builds the preview code `mms mcp` runs (src/agent/main.ts) into one
// JavaScript file, embedded in the engine at engine/mcp/preview/agent.mjs.
// Run it after changing the preview worker, the climate helpers, the jigsaw
// simulator or src/agent: `npm run build:agent -w packages/preview`.
// engine's TestAgentInSync fails while the committed file is stale.
import { build } from 'esbuild';
import { fileURLToPath } from 'node:url';

const root = fileURLToPath(new URL('..', import.meta.url));
const out = process.argv[2] ?? fileURLToPath(new URL('../../../engine/mcp/preview/agent.mjs', import.meta.url));

await build({
	absWorkingDir: root,
	entryPoints: ['src/agent/main.ts'],
	bundle: true,
	platform: 'node',
	format: 'esm',
	target: 'node20',
	minify: !process.env.DEV,
	legalComments: 'eof',
	outfile: out,
	logLevel: 'warning'
});
