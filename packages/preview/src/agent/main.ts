// The entry `mms mcp` runs for its preview tools: one AgentRequest as JSON
// on stdin, one {text} or {error} as JSON on stdout. Built by
// scripts/build-agent.mjs into a single file any JS runtime can run.

/// <reference types="node" />
import process from 'node:process';
import { biomes } from './biomes';
import { ArgError, previewFiles, type AgentRequest } from './common';
import { density } from './density';
import { setup } from './shim';
import { jigsaw } from './jigsaw';

export async function run(req: AgentRequest): Promise<{ text: string } | { error: string; args?: boolean }> {
	try {
		const files = previewFiles(req.files);
		setup(files, req.vanilla ?? {});
		const args = req.args ?? {};
		switch (req.tool) {
			case 'biomes':
				return { text: await biomes(args, files) };
			case 'density':
				return { text: await density(args, files, req.vanilla ?? {}) };
			case 'jigsaw':
				return { text: jigsaw(args, req.files, req.assets ?? []) };
		}
		return { error: `unknown tool ${String(req.tool)}` };
	} catch (e) {
		// MMS_JS_DEBUG=1 prints the stack, for working on these tools.
		if (process.env.MMS_JS_DEBUG) console.error(e);
		return { error: e instanceof Error ? e.message : String(e), args: e instanceof ArgError };
	}
}

async function readStdin(): Promise<string> {
	const chunks: Buffer[] = [];
	for await (const c of process.stdin) chunks.push(c as Buffer);
	return Buffer.concat(chunks).toString('utf8');
}

const res = await run(JSON.parse(await readStdin()) as AgentRequest);
process.stdout.write(JSON.stringify(res));
// The worker may leave a timer behind; the answer is out.
process.exit(0);
