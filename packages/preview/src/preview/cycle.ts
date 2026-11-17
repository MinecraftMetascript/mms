// Density functions that refer back to themselves. Computing one never
// ends, so deepslate runs out of stack; the engine reports the cycle at
// the source (mms/density-cycle), and the preview says the same instead
// of "Maximum call stack size exceeded".

import type { PreviewFile } from './protocol';

/** Each density function's refs to the files' other density functions. */
function refsOf(files: readonly PreviewFile[]): Map<string, string[]> {
	const ids = new Set(files.filter((f) => f.kind === 'density_function').map((f) => f.id));
	const refs = new Map<string, string[]>();
	const walk = (v: unknown, key: string, out: string[]) => {
		if (typeof v === 'string') {
			// Noises are ids too, but only under "noise".
			if (key !== 'type' && key !== 'noise' && ids.has(v)) out.push(v);
		} else if (Array.isArray(v)) {
			for (const e of v) walk(e, key, out);
		} else if (v && typeof v === 'object') {
			for (const [k, e] of Object.entries(v)) walk(e, k, out);
		}
	};
	for (const f of files) {
		if (f.kind !== 'density_function') continue;
		const out: string[] = [];
		walk(f.json, '', out);
		refs.set(f.id, out);
	}
	return refs;
}

/** A cycle of density function ids, the first repeated last, or undefined. */
export function densityCycle(files: readonly PreviewFile[]): string[] | undefined {
	const refs = refsOf(files);
	const state = new Map<string, 'visiting' | 'done'>();
	const path: string[] = [];
	const visit = (id: string): string[] | undefined => {
		state.set(id, 'visiting');
		path.push(id);
		for (const to of refs.get(id) ?? []) {
			const s = state.get(to);
			if (s === 'visiting') return [...path.slice(path.indexOf(to)), to];
			if (!s) {
				const found = visit(to);
				if (found) return found;
			}
		}
		path.pop();
		state.set(id, 'done');
		return undefined;
	};
	for (const id of refs.keys()) {
		if (!state.has(id)) {
			const found = visit(id);
			if (found) return found;
		}
	}
	return undefined;
}

/** The message for a cycle: "density function demo:a refers to itself (demo:a → demo:b → demo:a)". */
export function cycleMessage(cyc: readonly string[]): string {
	return `density function ${cyc[0]} refers to itself (${cyc.join(' → ')}), so it can't be computed; fix the cycle in the code`;
}
