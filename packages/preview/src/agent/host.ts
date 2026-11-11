// Runs the preview worker outside a browser, for `mms mcp`'s preview tools.
// The worker is imported as it is, so an agent sees the same numbers the
// editors draw: `self` becomes a stand-in that collects what the worker
// posts, and the vanilla bundles it would fetch from mcmeta come from the
// engine's embedded vanilla files instead.
//
// Import this module before anything that imports the worker (see shim.ts).

import type { PreviewFile, PreviewRequest, PreviewResponse } from '../preview/protocol';
import { vanillaSummary } from '../preview/protocol';

export type RenderRequest = Extract<PreviewRequest, { type: 'render' }>;
export type Rendered = Extract<PreviewResponse, { type: 'rendered' }>;

/** Vanilla files by data pack folder ("worldgen/noise", "tags/worldgen/biome"), then path. */
export type VanillaFiles = Record<string, Record<string, unknown>>;

interface Pending {
	resolve: (r: Rendered) => void;
	reject: (e: Error) => void;
}

const pending = new Map<number, Pending>();
let nextSeq = 1;

/** What the worker's onmessage handler is, once it is imported. */
type Handler = (e: { data: PreviewRequest }) => unknown;

const fakeSelf = {
	onmessage: null as Handler | null,
	postMessage(res: PreviewResponse) {
		if (res.type !== 'rendered' && res.type !== 'error') return;
		if (res.type === 'rendered' && res.partial) return;
		const p = pending.get(res.seq);
		if (!p) return;
		pending.delete(res.seq);
		if (res.type === 'error') p.reject(new Error(res.message));
		else p.resolve(res);
	}
};

let vanilla: VanillaFiles = {};

/**
 * Serves the worker's requests for mcmeta bundles from the vanilla files;
 * anything else fails as if offline.
 */
async function fakeFetch(input: unknown): Promise<{ ok: boolean; json(): Promise<unknown> }> {
	const url = String(input);
	const prefix = `${vanillaSummary}/`;
	const suffix = '/data.min.json';
	if (url.startsWith(prefix) && url.endsWith(suffix)) {
		// The worker asks for worldgen/<folder> and tag/worldgen/biome.
		const path = url.slice(prefix.length, -suffix.length);
		const folder = path.startsWith('tag/') ? `tags/${path.slice(4)}` : path;
		const files = vanilla[folder];
		if (files) return { ok: true, json: async () => files };
	}
	return { ok: false, json: async () => null };
}

const g = globalThis as unknown as { self: unknown; fetch: unknown };
g.self = fakeSelf;
g.fetch = fakeFetch;

/** Hands the worker the pack's files and the vanilla files it may need. */
export function setup(files: PreviewFile[], vanillaFiles: VanillaFiles) {
	vanilla = vanillaFiles;
	send({ type: 'setup', files });
}

function send(req: PreviewRequest) {
	if (!fakeSelf.onmessage) throw new Error('the preview worker is not loaded');
	void fakeSelf.onmessage({ data: req });
}

/** Renders one view, as the editors' preview pane asks for it. */
export function render(req: Omit<RenderRequest, 'type' | 'seq'>): Promise<Rendered> {
	const seq = nextSeq++;
	return new Promise((resolve, reject) => {
		pending.set(seq, { resolve, reject });
		send({ ...req, type: 'render', seq });
	});
}
