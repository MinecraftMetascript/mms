// Shared by the agent tools: the request the engine sends, finding a
// file by id, and number formatting.

import type { OutputAsset, OutputFile } from '../types';
import type { PreviewFile } from '../preview/protocol';
import { previewFile } from '../paths';
import type { VanillaFiles } from './host';

/** What `mms mcp` sends: one tool call with the project's output. */
export interface AgentRequest {
	tool: 'biomes' | 'density' | 'jigsaw';
	args: Record<string, unknown>;
	/** The data pack files the project writes. */
	files: OutputFile[];
	/** The files it copies, with structure files' jigsaw blocks. */
	assets?: OutputAsset[];
	/** The vanilla worldgen files previews fall back on. */
	vanilla?: VanillaFiles;
}

/** A tool's mistake the agent can fix by changing its arguments. */
export class ArgError extends Error {}

/** The output files the preview worker reads. */
export function previewFiles(files: readonly OutputFile[]): PreviewFile[] {
	const out: PreviewFile[] = [];
	for (const f of files) {
		const p = previewFile(f.path);
		if (!p) continue;
		try {
			out.push({ ...p, json: JSON.parse(f.text) });
		} catch {
			// A file that does not parse has a diagnostic of its own.
		}
	}
	return out;
}

/** An id with its namespace: minecraft: when it has none. */
export const qualified = (id: string) => (id.includes(':') ? id : `minecraft:${id}`);

/**
 * Finds the one id an agent meant among candidates: an exact match, or a
 * bare name that matches one candidate's path in any namespace.
 */
export function pick(what: string, id: string, candidates: readonly string[]): string {
	if (candidates.includes(id)) return id;
	if (!id.includes(':')) {
		const hits = candidates.filter((c) => c.slice(c.indexOf(':') + 1) === id);
		if (hits.length === 1) return hits[0];
		if (hits.length > 1) throw new ArgError(`${id} could be ${hits.join(' or ')}; give the namespace.`);
		if (candidates.includes(`minecraft:${id}`)) return `minecraft:${id}`;
	}
	const list = candidates.length ? `: ${candidates.slice(0, 30).join(', ')}${candidates.length > 30 ? ', …' : ''}` : '';
	throw new ArgError(`No ${what} ${id}. The ${what}s there are${list || ' none'}.`);
}

/** A number argument, or the default; non-numbers are an argument error. */
export function num(args: Record<string, unknown>, key: string, def: number, min = -Infinity, max = Infinity): number {
	const v = args[key];
	if (v === undefined || v === null) return def;
	if (typeof v !== 'number' || !Number.isFinite(v)) throw new ArgError(`${key} must be a number.`);
	if (v < min || v > max) throw new ArgError(`${key} must be between ${min} and ${max}.`);
	return v;
}

export function str(args: Record<string, unknown>, key: string): string | undefined {
	const v = args[key];
	if (v === undefined || v === null || v === '') return undefined;
	if (typeof v !== 'string') throw new ArgError(`${key} must be a string.`);
	return v;
}

/** A value to three decimals, without trailing zeros; NaN reads as "none". */
export function fmt(v: number, digits = 3): string {
	if (Number.isNaN(v)) return 'none';
	if (!Number.isFinite(v)) return v > 0 ? '∞' : '-∞';
	const s = v.toFixed(digits);
	return s.includes('.') ? s.replace(/\.?0+$/, '') : s;
}

/** A share as a percentage to one decimal. */
export const pct = (part: number, whole: number) => `${whole ? fmt((100 * part) / whole, 1) : '0'}%`;

/** A fixed-width table: rows of cells, padded per column; the first row is the header. Numbers line up on the right, text on the left. */
export function table(rows: string[][]): string {
	const widths: number[] = [];
	const numeric: boolean[] = [];
	for (const r of rows) r.forEach((c, i) => (widths[i] = Math.max(widths[i] ?? 0, c.length)));
	widths.forEach((_, i) => (numeric[i] = rows.slice(1).every((r) => /^(-?[\d.]|none$|-$)/.test(r[i] ?? ''))));
	return rows.map((r) => r.map((c, i) => (numeric[i] ? c.padStart(widths[i]) : c.padEnd(widths[i]))).join('  ').trimEnd()).join('\n');
}
