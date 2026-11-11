// The jigsaw tool: grows a jigsaw structure with the simulator the
// editors' Simulation tab uses, once or over a run of seeds, and reports
// what was placed, what never was, and why jigsaw blocks stayed empty.

import { MAX_DEPTH, simulate, type NoneReason, type Simulation } from '../jigsaw/simulate';
import type { OutputAsset, OutputFile } from '../types';
import { ArgError, fmt, num, pick, str, table } from './common';

const reasons: Record<NoneReason, string> = {
	empty: 'its pool is minecraft:empty',
	'unknown-pool': 'its pool is not in the pack',
	depth: 'the structure was at its last depth and the fallback had nothing that fit',
	'no-fit': 'no piece could join or fit the free space',
	'empty-piece': 'the pool drew its empty piece'
};

const structurePath = /^data\/([^/]+)\/worldgen\/structure\/(.+)\.json$/;

export function jigsaw(args: Record<string, unknown>, files: OutputFile[], assets: OutputAsset[]): string {
	const ids: string[] = [];
	for (const f of files) {
		const m = structurePath.exec(f.path);
		if (m && /"type"\s*:\s*"(minecraft:)?jigsaw"/.test(f.text)) ids.push(`${m[1]}:${m[2]}`);
	}
	const given = str(args, 'structure');
	if (!given && ids.length !== 1) throw new ArgError(ids.length ? `Give structure: one of ${ids.join(', ')}.` : 'The project has no jigsaw structure.');
	const id = given ? pick('jigsaw structure', given, ids) : ids[0];
	const [ns, name] = id.split(':', 2);
	const path = `data/${ns}/worldgen/structure/${name}.json`;
	const seed = Math.floor(num(args, 'seed', 0));
	const runs = Math.floor(num(args, 'runs', 1, 1, 200));
	const depth = args.depth === undefined ? undefined : Math.floor(num(args, 'depth', 0, 0, MAX_DEPTH));

	const sims: Simulation[] = [];
	for (let s = seed; s < seed + runs; s++) sims.push(simulate(path, files, assets, { seed: s, depth })!);
	return runs === 1 ? one(id, sims[0], files) : many(id, sims, files);
}

/** The pieces each reached pool lists, by pool, for the "never placed" report. */
function poolPieces(files: OutputFile[], pools: Iterable<string>): Map<string, string[]> {
	const out = new Map<string, string[]>();
	for (const pool of pools) {
		const [ns, name] = pool.split(':', 2);
		const f = files.find((f) => f.path === `data/${ns}/worldgen/template_pool/${name}.json`);
		if (!f) continue;
		try {
			const j = JSON.parse(f.text) as { elements?: { element?: { location?: string; elements?: { location?: string }[] } }[] };
			const ids = new Set<string>();
			for (const e of j.elements ?? []) {
				const loc = e.element?.location;
				if (typeof loc === 'string') ids.add(loc.includes(':') ? loc : `minecraft:${loc}`);
			}
			out.set(pool, [...ids]);
		} catch {
			// Has a diagnostic of its own.
		}
	}
	return out;
}

const boxText = (b: Simulation['bounds']) => `${b[0].join(' ')} to ${b[1].join(' ')}`;

function header(id: string, s: Simulation): string[] {
	const out = [`Jigsaw structure ${id}: depth ${s.depth}${s.depth !== s.size ? ` (its size is ${s.size})` : ''}.`];
	if (s.notes.length) out.push(...s.notes.map((n) => `Note: ${n}`));
	return out;
}

/** One seed: every piece, and every jigsaw block that stayed empty. */
function one(id: string, s: Simulation, files: OutputFile[]): string {
	const out = header(id, s);
	out.push(`Seed ${s.seed}.`);
	if (s.failed) return [...out, `Nothing is placed: ${s.failed}`].join('\n');
	out.push(`${s.pieces.length} pieces, ${s.reached} deep. Bounds (max_distance_from_center): ${boxText(s.bounds)}.`);
	out.push('', 'Pieces (index, depth, structure, pool, rotation, origin, joined onto piece):');
	const rows = [['#', 'depth', 'piece', 'pool', 'rot', 'origin', 'parent']];
	s.pieces.forEach((p, i) => {
		rows.push([String(i), String(p.depth), p.id, p.pool, `${p.rotation * 90}°`, p.origin.join(' '), p.parent ? String(p.parent.piece) : '-']);
	});
	out.push(table(rows));
	const empty: string[] = [];
	s.pieces.forEach((p, i) =>
		p.jigsaws.forEach((j) => {
			if (j.result && 'none' in j.result && j.result.none !== 'empty') {
				empty.push(`- piece ${i} (${p.id}) at ${j.pos.join(' ')}, pool ${j.pool}, target ${j.target}: ${reasons[j.result.none]}`);
			}
		})
	);
	if (empty.length) out.push('', 'Jigsaw blocks that placed nothing (minecraft:empty pools left out):', ...empty);
	out.push(...never(files, [s]));
	return out.join('\n');
}

/** Many seeds: how often each piece appears, and how the structure's size varies. */
function many(id: string, sims: Simulation[], files: OutputFile[]): string {
	const out = header(id, sims[0]);
	out.push(`Seeds ${sims[0].seed}..${sims[sims.length - 1].seed} (${sims.length} runs).`);
	const failed = sims.filter((s) => s.failed);
	if (failed.length) {
		const why = new Map<string, number>();
		for (const s of failed) why.set(s.failed!, (why.get(s.failed!) ?? 0) + 1);
		out.push(`${failed.length} runs placed nothing:`, ...[...why].map(([w, c]) => `- ${c}×: ${w}`));
	}
	const ok = sims.filter((s) => !s.failed);
	if (ok.length) {
		const sizes = ok.map((s) => s.pieces.length);
		const deep = ok.map((s) => s.reached);
		out.push(
			`Pieces per run: ${Math.min(...sizes)}..${Math.max(...sizes)}, mean ${fmt(sizes.reduce((a, b) => a + b, 0) / ok.length, 1)}. ` +
				`Depth reached: ${Math.min(...deep)}..${Math.max(...deep)}.`
		);
		// How many runs place each piece, and how many of it a run places.
		const runsWith = new Map<string, number>();
		const total = new Map<string, number>();
		for (const s of ok) {
			const seen = new Set<string>();
			for (const p of s.pieces) {
				total.set(p.id, (total.get(p.id) ?? 0) + 1);
				seen.add(p.id);
			}
			for (const p of seen) runsWith.set(p, (runsWith.get(p) ?? 0) + 1);
		}
		out.push('', 'Pieces (runs that place it, mean count per run):');
		const rows = [['piece', 'runs', 'mean']];
		for (const [p, c] of [...runsWith].sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]))) rows.push([p, `${c}/${ok.length}`, fmt(total.get(p)! / ok.length, 2)]);
		out.push(table(rows));
		const empties = new Map<string, number>();
		for (const s of ok)
			for (const p of s.pieces)
				for (const j of p.jigsaws)
					if (j.result && 'none' in j.result && j.result.none !== 'empty') {
						const k = `${p.id}: pool ${j.pool}, target ${j.target}: ${reasons[j.result.none]}`;
						empties.set(k, (empties.get(k) ?? 0) + 1);
					}
		if (empties.size) {
			out.push('', 'Jigsaw blocks that placed nothing, over all runs (minecraft:empty pools left out):');
			for (const [k, c] of [...empties].sort((a, b) => b[1] - a[1])) out.push(`- ${c}×  ${k}`);
		}
	}
	out.push(...never(files, ok));
	return out.join('\n');
}

/** The pieces the pools they drew from list that no run placed. */
function never(files: OutputFile[], sims: Simulation[]): string[] {
	const placed = new Set<string>();
	const pools = new Set<string>();
	for (const s of sims)
		for (const p of s.pieces) {
			placed.add(p.id);
			pools.add(p.pool);
			for (const j of p.jigsaws) pools.add(j.pool.includes(':') ? j.pool : `minecraft:${j.pool}`);
		}
	const out: string[] = [];
	for (const [pool, pieces] of poolPieces(files, pools)) {
		const missed = pieces.filter((p) => !placed.has(p));
		if (missed.length) out.push(`- ${pool}: ${missed.join(', ')}`);
	}
	return out.length ? ['', `Pieces their pools list that ${sims.length > 1 ? 'no run' : 'this run'} placed:`, ...out] : [];
}
