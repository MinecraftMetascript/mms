// The spline editor's model: a density function's cubic spline as the
// language server's `mms/spline` gives it, the game's way of working out
// its value, and the text edits that move, add and remove points.
import type { SourceRange } from '../types';

/** A number of a point. range is set only for a literal that can be rewritten. */
export interface SplineNum {
	/** null when it could not be worked out. */
	value: number | null;
	range?: SourceRange;
}

export interface SplinePointData {
	range: SourceRange;
	location: SplineNum;
	derivative: SplineNum;
	/** The value, when it is a number. */
	value?: SplineNum;
	/** The value, when it is a spline. */
	spline?: SplineData;
	/** What deleting the point removes; set when the points are a list in place. */
	remove?: SourceRange;
}

export interface SplineData {
	range: SourceRange;
	/** The coordinate as written, such as `overworld/continents`. */
	coordinate: string;
	/** The points are a list in place, so points can be added and removed. */
	editable: boolean;
	/** A new point's source, with $location, $derivative and $value for its numbers. */
	template?: string;
	/** Where a new point goes: inserts[i] puts it before point i, the last after the last point. */
	inserts?: { range: SourceRange; before: string; after: string }[];
	points: SplinePointData[];
}

/** The spline at the cursor: the outermost of its nesting, and the point indices down to the one at the cursor. */
export interface SplineAt {
	root: SplineData;
	focus: number[];
}

export interface TextEdit {
	range: SourceRange;
	newText: string;
}

/** Values to use for coordinates, by their source text, when working out nested splines. */
export type Coordinates = Record<string, number>;

/** A point's value: its number, or its nested spline at the given coordinates. */
export function pointValue(p: SplinePointData, at: Coordinates): number {
	if (p.spline) return evaluate(p.spline, at[p.spline.coordinate] ?? 0, at);
	return p.value?.value ?? 0;
}

/**
 * A spline's value at x, the way the game works it out: a cubic Hermite
 * curve between the two points around x, and a straight line along the
 * end point's derivative past either end.
 */
export function evaluate(s: SplineData, x: number, at: Coordinates = {}): number {
	const pts = s.points;
	const n = pts.length;
	if (n === 0) return 0;
	const loc = (i: number) => pts[i].location.value ?? 0;
	const der = (i: number) => pts[i].derivative.value ?? 0;
	// The last point at or before x.
	let i = -1;
	for (let j = 0; j < n; j++) if (loc(j) <= x) i = j;
	if (i < 0) return pointValue(pts[0], at) + der(0) * (x - loc(0));
	if (i === n - 1) return pointValue(pts[i], at) + der(i) * (x - loc(i));
	const x0 = loc(i);
	const x1 = loc(i + 1);
	const v0 = pointValue(pts[i], at);
	const v1 = pointValue(pts[i + 1], at);
	const w = x1 - x0;
	const t = (x - x0) / w;
	const p = der(i) * w - (v1 - v0);
	const q = -der(i + 1) * w + (v1 - v0);
	return lerp(t, v0, v1) + t * (1 - t) * lerp(t, p, q);
}

function lerp(t: number, a: number, b: number) {
	return a + t * (b - a);
}

/** The spline at a path of point indices from root, or the deepest one the path still reaches. */
export function splineAt(root: SplineData, path: number[]): { spline: SplineData; path: number[] } {
	let s = root;
	const kept: number[] = [];
	for (const i of path) {
		const next = s.points[i]?.spline;
		if (!next) break;
		s = next;
		kept.push(i);
	}
	return { spline: s, path: kept };
}

/** The coordinates of the splines nested under s, each with the range of locations its splines use. */
export function nestedCoordinates(s: SplineData): { name: string; min: number; max: number }[] {
	const out = new Map<string, { name: string; min: number; max: number }>();
	const visit = (sp: SplineData) => {
		for (const p of sp.points) {
			if (!p.spline) continue;
			const c = p.spline.coordinate;
			const e = out.get(c) ?? { name: c, min: Infinity, max: -Infinity };
			for (const q of p.spline.points) {
				if (q.location.value === null) continue;
				e.min = Math.min(e.min, q.location.value);
				e.max = Math.max(e.max, q.location.value);
			}
			out.set(c, e);
			visit(p.spline);
		}
	};
	visit(s);
	return [...out.values()].map((e) => (e.min <= e.max ? e : { ...e, min: -1, max: 1 }));
}

/** A number as the editor writes it: at most `digits` decimals, no trailing zeros. */
export function formatNumber(v: number, digits = 3): string {
	const s = Number(v.toFixed(digits)).toString();
	return s === '-0' ? '0' : s;
}

/** The edit that writes v over a literal, or null when the number is not one. */
export function setNumber(n: SplineNum, v: number, digits = 3): TextEdit | null {
	return n.range ? { range: n.range, newText: formatNumber(v, digits) } : null;
}

/** The edit that adds a point, in location order, or null when the points are not a list in place. */
export function addPoint(s: SplineData, location: number, value: number, derivative = 0): TextEdit | null {
	if (!s.editable || !s.template || !s.inserts?.length) return null;
	let gap = 0;
	while (gap < s.points.length && (s.points[gap].location.value ?? -Infinity) <= location) gap++;
	const ins = s.inserts[Math.min(gap, s.inserts.length - 1)];
	const text = s.template
		.replace('$location', formatNumber(location))
		.replace('$derivative', formatNumber(derivative))
		.replace('$value', formatNumber(value));
	return { range: ins.range, newText: ins.before + text + ins.after };
}

/** The edit that deletes a point, or null when the points are not a list in place. */
export function removePoint(p: SplinePointData): TextEdit | null {
	return p.remove ? { range: p.remove, newText: '' } : null;
}

/** Whether two splines are the same one, by where they start. */
export function sameSpline(a: SplineData | undefined, b: SplineData | undefined): boolean {
	return !!a && !!b && a.range.start.line === b.range.start.line && a.range.start.character === b.range.start.character;
}
