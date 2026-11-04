// The jigsaw connection graph as text, to paste into an issue, a chat or
// a Markdown file: a report listing every pool, piece and jigsaw block
// with what it joins and why a dead end is dead, and a Mermaid flowchart
// of which pieces join which.
//
// Free of DOM types, so it can be tested and reused anywhere.

import { deadEndText, type JigsawGraph, type Link, type Piece, type Pool } from './graph';
import type { OutputJigsaw as Jigsaw } from '../types';

const EMPTY = 'minecraft:empty';
const pct = (n: number) => (n >= 0.995 ? '100%' : n < 0.01 ? '<1%' : `${Math.round(n * 100)}%`);
const code = (s: string) => `\`${s}\``;
const sizeText = (s: [number, number, number]) => `${s[0]}×${s[1]}×${s[2]}`;
/** A piece as `id`, with its label when it has one. */
const pieceName = (p: Piece) => `${p.kind === 'single' ? code(p.id) : `*${p.id}*`}${p.label ? ` "${p.label}"` : ''}`;

/** Its facing, the way its top points, and joint, from its orientation. */
function facing(j: Jigsaw): { front: string; top: string } | null {
	const m = /^([a-z]+)_([a-z]+)$/.exec(j.orientation ?? '');
	return m ? { front: m[1], top: m[2] } : null;
}

/**
 * The sides of the piece a jigsaw block sits on, and a note when it faces
 * out of a side it is not on: the game then treats it as internal and only
 * places a piece there that fits wholly inside this one.
 */
function sides(j: Jigsaw, size: [number, number, number] | undefined): { on: string[]; warn?: string } {
	if (!size) return { on: [] };
	const [x, y, z] = j.pos;
	const on: string[] = [];
	if (x === 0) on.push('west');
	if (x === size[0] - 1) on.push('east');
	if (z === 0) on.push('north');
	if (z === size[2] - 1) on.push('south');
	if (y === 0) on.push('down');
	if (y === size[1] - 1) on.push('up');
	const f = facing(j);
	const warn = f && !on.includes(f.front) ? `is internal: it faces ${f.front} but is not on the ${f.front} side, so a piece joining it must fit inside this one` : undefined;
	return { on, warn };
}

/** The pieces a jigsaw block can join, in words, with full ids. */
export function joinsText(graph: JigsawGraph, piece: Piece, l: Link, id: (s: string) => string = (s) => s): string {
	if (l.dead) return deadEndText(l, id(piece.jigsaws![l.jigsaw].target), piece.jigsaws![l.jigsaw]);
	if (l.truncated) return `${id(l.pool)} (beyond the depths shown)`;
	const pool = l.to !== undefined ? graph.pools[l.to] : undefined;
	if (!pool?.declared) return `${id(l.pool)} (not read: vanilla or another pack)`;
	const named = (q: Piece) => (q.label ? `${id(q.id)} (${q.label})` : id(q.id));
	const names = l.pieces.map((i) => named(pool.pieces[i]));
	if (l.fallback) {
		const fb = graph.pools[l.fallback.to];
		names.push(`only the fallback ${id(fb.id)}: ${l.fallback.pieces.map((i) => named(fb.pieces[i])).join(', ')}`);
	}
	if (l.maybe) names.push('pieces whose jigsaws are unknown');
	return names.join(', ') || 'nothing (target is empty)';
}

/** The jigsaw names that jigsaw blocks drawing from each pool target. */
function targetsByPool(graph: JigsawGraph): Map<string, Set<string>> {
	const out = new Map<string, Set<string>>();
	for (const pool of graph.pools)
		for (const p of pool.pieces)
			for (const j of p.jigsaws ?? []) {
				if (j.pool === EMPTY || j.target === EMPTY) continue;
				if (!out.has(j.pool)) out.set(j.pool, new Set());
				out.get(j.pool)!.add(j.target);
			}
	// A pool's fallback is drawn from by the same jigsaw blocks.
	for (const pool of graph.pools) {
		const names = out.get(pool.id);
		if (!pool.fallback || !names) continue;
		if (!out.has(pool.fallback)) out.set(pool.fallback, new Set());
		for (const n of names) out.get(pool.fallback)!.add(n);
	}
	return out;
}

function poolNote(pool: Pool): string {
	if (pool.declared) return pool.pieces.length ? '' : 'It has no pieces.';
	return pool.id.split(':')[0] === 'minecraft' ? 'Vanilla pool, not read.' : 'No TemplatePool declares this pool.';
}

/**
 * The graph as a Markdown report: a summary, every problem with its
 * piece and jigsaw block, every pool's pieces and their jigsaw blocks
 * (position, facing, joint, name, pool, target, what joins), and a
 * Mermaid flowchart.
 */
export function jigsawMarkdown(graph: JigsawGraph): string {
	const out: string[] = [];
	const root = graph.root;
	const start = graph.pools[0];
	out.push(`# Jigsaw connections: ${root.id}`, '');
	const facts = [`${root.kind === 'structure' ? 'Jigsaw structure' : 'Template pool'} ${code(root.id)}`];
	if (root.kind === 'structure' && start) facts.push(`start pool ${code(start.id)}`);
	if (graph.size !== undefined) facts.push(`size ${graph.size} (how many pieces deep it grows from the start piece)`);
	if (graph.startName) facts.push(`start jigsaw name ${code(graph.startName)}`);
	out.push(`${facts.join(', ')}.`);
	if (graph.repeatsFrom !== undefined) {
		const last = graph.pools.reduce((n, p) => Math.max(n, p.depth), 0);
		out.push('', `Growth repeats: after depth ${last}, pieces join the same pools as at depth ${graph.repeatsFrom} again, so later depths are not listed.`);
	}
	out.push(
		'',
		'Positions are x y z inside the piece, from its north-west bottom corner; sizes are x×y×z. Facing is the way a jigsaw block points (where the next piece goes), top is the way its top points. ' +
			'A jigsaw block joins one whose name is its target, facing the opposite way; a joint only matters for up or down facing blocks (rollable: the joined piece may turn; aligned: tops must match).'
	);

	// Problems first, so they are easy to find.
	const problems: string[] = [];
	const seen = new Set<string>();
	const targets = targetsByPool(graph);
	for (const pool of graph.pools) {
		if (pool.tooDeep && !seen.has(`deep|${pool.id}|${pool.depth}`)) {
			seen.add(`deep|${pool.id}|${pool.depth}`);
			problems.push(`- Past size: ${code(pool.id)} at depth ${pool.depth} is past the structure's size ${graph.size}, so nothing from it is placed.`);
		}
		pool.pieces.forEach((p, pi) => {
			p.links.forEach((l) => {
				if (!l.dead) return;
				const key = `dead|${pool.id}|${pi}|${l.jigsaw}`;
				if (seen.has(key)) return;
				seen.add(key);
				const j = p.jigsaws![l.jigsaw];
				const f = facing(j);
				problems.push(
					`- Dead end: ${pieceName(p)} in ${code(pool.id)}, jigsaw block at ${j.pos.join(' ')}${f ? ` facing ${f.front}` : ''} (name ${code(j.name)}, pool ${code(j.pool)}, target ${code(j.target)}): ${deadEndText(l, j.target, j)}.`
				);
			});
			if (p.unreachable && !seen.has(`never|${pool.id}|${pi}`)) {
				seen.add(`never|${pool.id}|${pi}`);
				const has = [...new Set((p.jigsaws ?? []).map((j) => j.name))].map(code).join(', ') || 'none';
				const wanted =
					pool === start && graph.startName ? code(graph.startName) : [...(targets.get(pool.id) ?? [])].map(code).join(', ') || 'none';
				problems.push(
					`- Never joined: ${pieceName(p)} in ${code(pool.id)}. Its jigsaw names are ${has}, but jigsaw blocks drawing from this pool only target ${wanted}${pool === start ? ' (the start jigsaw name)' : ''}.`
				);
			}
			if (p.kind === 'single' && p.jigsaws === undefined && !seen.has(`unknown|${p.id}`)) {
				seen.add(`unknown|${p.id}`);
				problems.push(`- Not read: ${pieceName(p)} in ${code(pool.id)} is not a structure file of this project, so its jigsaw blocks are unknown.`);
			}
			for (const j of p.jigsaws ?? []) {
				const w = sides(j, p.size).warn;
				if (w && !seen.has(`side|${p.id}|${j.pos.join(',')}`)) {
					seen.add(`side|${p.id}|${j.pos.join(',')}`);
					problems.push(`- Check: ${pieceName(p)} jigsaw block at ${j.pos.join(' ')} (${code(j.name)}) ${w} (size ${sizeText(p.size!)}).`);
				}
			}
		});
	}
	out.push('', '## Problems', '');
	out.push(problems.length ? problems.join('\n') : 'None found.');

	// Pools by depth; a pool shown at several depths lists its pieces once.
	out.push('', '## Pools by depth');
	const listed = new Map<string, number>();
	for (const pool of graph.pools) {
		const tags = [pool === start ? 'start' : `depth ${pool.depth}`];
		if (pool.tooDeep) tags.push('past size');
		out.push('', `### ${code(pool.id)} · ${tags.join(', ')}`, '');
		const notes: string[] = [];
		const note = poolNote(pool);
		if (note) notes.push(note);
		if (pool.fallback) notes.push(`Fallback ${code(pool.fallback)} (tried when no piece of this pool fits).`);
		const wanted = targets.get(pool.id);
		if (pool !== start && wanted?.size) notes.push(`Jigsaw blocks drawing from it target ${[...wanted].map(code).join(', ')}.`);
		const before = listed.get(pool.id);
		if (before !== undefined) {
			notes.push(`Same pieces as at depth ${before}, listed there.`);
			out.push(notes.join(' '));
			continue;
		}
		listed.set(pool.id, pool.depth);
		if (notes.length) out.push(notes.join(' '), '');
		for (const p of pool.pieces) {
			const head = [`weight ${p.weight} (${pct(p.share)})`];
			if (p.size) head.push(`size ${sizeText(p.size)}`);
			if (p.unreachable) head.push('**never joined**');
			out.push(`- ${pieceName(p)} · ${head.join(' · ')}`);
			if (p.kind !== 'single') continue;
			if (!p.jigsaws) {
				out.push('  - Jigsaw blocks unknown: not a structure file of this project.');
				continue;
			}
			if (!p.jigsaws.length) {
				out.push('  - No jigsaw blocks, so the structure ends here.');
				continue;
			}
			out.push('', '  | At | Side | Facing | Top | Joint | Name | Pool | Target | Joins |', '  |---|---|---|---|---|---|---|---|---|');
			p.jigsaws.forEach((j, ji) => {
				const f = facing(j);
				const s = sides(j, p.size);
				const l = p.links.find((x) => x.jigsaw === ji);
				const joins = j.pool === EMPTY ? 'nothing: it only receives' : l ? (l.dead ? `**dead end:** ${joinsText(graph, p, l)}` : joinsText(graph, p, l)) : '';
				const side = [s.on.join('+') || 'inside', s.warn ? `⚠ ${s.warn}` : ''].filter(Boolean).join(', ');
				out.push(
					`  | ${j.pos.join(' ')} | ${p.size ? side : '?'} | ${f?.front ?? '?'} | ${f?.top ?? '?'} | ${f && (f.front === 'up' || f.front === 'down') ? (j.joint ?? '?') : '–'} | ${code(j.name)} | ${code(j.pool)} | ${code(j.target)} | ${joins.replace(/\|/g, '\\|')} |`
				);
			});
			out.push('');
		}
		if (out[out.length - 1] === '') out.pop();
	}

	out.push('', '## Graph', '', '```mermaid', jigsawMermaid(graph), '```', '');
	return out.join('\n');
}

/**
 * The graph as a Mermaid flowchart: a box per pool and depth holding its
 * pieces, a line from each piece to the pieces that can join it labelled
 * with the jigsaw target, dotted where it loops back or only maybe joins,
 * and red dead ends.
 */
export function jigsawMermaid(graph: JigsawGraph): string {
	const ns = graph.root.id.split(':')[0];
	const strip = (s: string) => (s.startsWith(`${ns}:`) ? s.slice(ns.length + 1) : s.replace(/^minecraft:/, ''));
	const label = (s: string) => `"${s.replace(/"/g, '#quot;')}"`;
	const lines: string[] = ['flowchart LR'];
	const dead: string[] = [];
	const never: string[] = [];
	graph.pools.forEach((pool, pi) => {
		const title = `${strip(pool.id)} · ${pi === 0 ? 'start' : `depth ${pool.depth}`}${pool.tooDeep ? ' · past size' : ''}`;
		lines.push(`  subgraph P${pi}[${label(title)}]`);
		if (!pool.pieces.length) lines.push(`    P${pi}_none[${label(poolNote(pool) || 'no pieces')}]`);
		pool.pieces.forEach((p, qi) => {
			const bits = [pct(p.share)];
			if (p.size) bits.push(sizeText(p.size));
			if (p.jigsaws === undefined && p.kind === 'single') bits.push('jigsaws unknown');
			const name = p.kind === 'single' ? strip(p.id) : p.id;
			lines.push(`    P${pi}_${qi}[${label(`${p.label ? `${p.label}<br/>${name}` : name}<br/>${bits.join(' · ')}`)}]`);
			if (p.unreachable) never.push(`P${pi}_${qi}`);
		});
		lines.push('  end');
	});
	// Piece to piece, one line per pair, labelled with the targets.
	const edges = new Map<string, { from: string; to: string; names: Set<string>; dotted: boolean }>();
	const edge = (from: string, to: string, name: string, dotted: boolean) => {
		const key = `${from}>${to}`;
		const e = edges.get(key);
		if (e) {
			e.names.add(name);
			e.dotted &&= dotted;
		} else edges.set(key, { from, to, names: new Set([name]), dotted });
	};
	graph.pools.forEach((pool, pi) => {
		pool.pieces.forEach((p, qi) => {
			for (const l of p.links) {
				const j = p.jigsaws![l.jigsaw];
				const from = `P${pi}_${qi}`;
				const name = strip(j.target);
				if (l.dead) {
					const id = `D${pi}_${qi}_${l.jigsaw}`;
					lines.push(`  ${id}([${label(`dead end: ${deadEndText(l, name, j)}`)}])`);
					edge(from, id, name, false);
					dead.push(id);
					continue;
				}
				if (l.to === undefined) continue;
				const back = (to: number) => graph.pools[to].depth <= pool.depth;
				for (const t of l.pieces) edge(from, `P${l.to}_${t}`, name, back(l.to) || l.maybe);
				if (l.fallback) for (const t of l.fallback.pieces) edge(from, `P${l.fallback.to}_${t}`, `${name} (fallback)`, true);
				// Pieces whose jigsaws are unknown may join too: a dotted line to the pool.
				if ((!l.pieces.length && !l.fallback) || (l.maybe && l.pieces.length)) edge(from, `P${l.to}`, name, true);
			}
		});
		if (pool.fallbackTo !== undefined) edge(`P${pi}`, `P${pool.fallbackTo}`, 'fallback', true);
	});
	for (const e of edges.values()) lines.push(`  ${e.from} ${e.dotted ? '-.' : '--'} ${label([...e.names].join(', '))} ${e.dotted ? '.->' : '-->'} ${e.to}`);
	if (dead.length) lines.push('  classDef dead fill:#fee2e2,stroke:#dc2626,color:#7f1d1d', `  class ${dead.join(',')} dead`);
	if (never.length) lines.push('  classDef never fill:#fef3c7,stroke:#d97706,color:#78350f', `  class ${never.join(',')} never`);
	return lines.join('\n');
}
