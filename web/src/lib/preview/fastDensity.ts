import { DensityFunction, NoiseRouter, type RandomState } from 'deepslate';

// Faster stand-ins for two of Deepslate's density function nodes, used by
// the terrain previews. Profiling a chunk fill showed most of its time in:
//
// - `interpolated`: each block looks up 8 cell corners in a Map keyed by a
//   string built per lookup, through closures (and the Map never shrinks).
// - `cache`: Deepslate doesn't cache at all, so terrain shaping that the
//   game computes once per column (continentalness, erosion, the offset,
//   factor and jaggedness splines) is recomputed at every cell corner.
//
// The replacements give the same values: corners are cached under numeric
// keys, and a `cache` whose input doesn't depend on y is computed once per
// column. Other caches remember their last position, which covers the same
// function being reached twice for one block.

type Context = { x: number; y: number; z: number };

/** Entries a cache keeps before starting over, to bound its memory. */
const MAX_ENTRIES = 1 << 18;

/** Packs block or cell coordinates (|x|, |z| < 2^20, |y| < 2^10) into one exact number. */
function pack(x: number, y: number, z: number): number {
	return ((x + 1048576) * 2097152 + (z + 1048576)) * 2048 + (y + 1024);
}

/** Deepslate's lazyLerp on values: exactly a when t is 0. */
function lerp(t: number, a: number, b: number): number {
	return t === 0 ? a : a + t * (b - a);
}

class FastInterpolated extends DensityFunction {
	private readonly corners = new Map<number, number>();
	constructor(
		readonly input: DensityFunction,
		readonly cellSizeXz: number,
		readonly cellSizeY: number
	) {
		super();
	}
	private corner(cx: number, cy: number, cz: number): number {
		const key = pack(cx, cy, cz);
		let v = this.corners.get(key);
		if (v === undefined) {
			if (this.corners.size >= MAX_ENTRIES) this.corners.clear();
			const w = this.cellSizeXz;
			v = this.input.compute(DensityFunction.context(cx * w, cy * this.cellSizeY, cz * w));
			this.corners.set(key, v);
		}
		return v;
	}
	// The corners of the last cell, since blocks are mostly asked for cell by cell.
	private cell = NaN;
	private c000 = 0; private c100 = 0; private c010 = 0; private c110 = 0;
	private c001 = 0; private c101 = 0; private c011 = 0; private c111 = 0;
	compute({ x, y, z }: Context): number {
		const w = this.cellSizeXz;
		const h = this.cellSizeY;
		const cx = Math.floor(x / w);
		const cy = Math.floor(y / h);
		const cz = Math.floor(z / w);
		const cell = pack(cx, cy, cz);
		if (cell !== this.cell) {
			this.cell = cell;
			this.c000 = this.corner(cx, cy, cz);
			this.c100 = this.corner(cx + 1, cy, cz);
			this.c010 = this.corner(cx, cy + 1, cz);
			this.c110 = this.corner(cx + 1, cy + 1, cz);
			this.c001 = this.corner(cx, cy, cz + 1);
			this.c101 = this.corner(cx + 1, cy, cz + 1);
			this.c011 = this.corner(cx, cy + 1, cz + 1);
			this.c111 = this.corner(cx + 1, cy + 1, cz + 1);
		}
		const tx = (x - cx * w) / w;
		const ty = (y - cy * h) / h;
		const tz = (z - cz * w) / w;
		// The same lerps, in the same order, as Deepslate's lazyLerp3.
		const near = lerp(ty, lerp(tx, this.c000, this.c100), lerp(tx, this.c010, this.c110));
		if (tz === 0) return near;
		return lerp(tz, near, lerp(ty, lerp(tx, this.c001, this.c101), lerp(tx, this.c011, this.c111)));
	}
	mapChildren(visitor: DensityFunction.Visitor): DensityFunction {
		return new FastInterpolated(visitor.apply(this.input), this.cellSizeXz, this.cellSizeY);
	}
	range() {
		return this.input.range();
	}
}

class ColumnCache extends DensityFunction {
	private readonly values = new Map<number, number>();
	constructor(readonly input: DensityFunction) {
		super();
	}
	compute(context: Context): number {
		if (!Number.isInteger(context.x) || !Number.isInteger(context.z)) return this.input.compute(context);
		const key = pack(context.x, 0, context.z);
		let v = this.values.get(key);
		if (v === undefined) {
			if (this.values.size >= MAX_ENTRIES) this.values.clear();
			v = this.input.compute(context);
			this.values.set(key, v);
		}
		return v;
	}
	mapChildren(visitor: DensityFunction.Visitor): DensityFunction {
		return new ColumnCache(visitor.apply(this.input));
	}
	range() {
		return this.input.range();
	}
}

class LastCache extends DensityFunction {
	private x = NaN;
	private y = NaN;
	private z = NaN;
	private value = 0;
	constructor(readonly input: DensityFunction) {
		super();
	}
	compute(context: Context): number {
		if (context.x !== this.x || context.y !== this.y || context.z !== this.z) {
			this.value = this.input.compute(context);
			this.x = context.x;
			this.y = context.y;
			this.z = context.z;
		}
		return this.value;
	}
	mapChildren(visitor: DensityFunction.Visitor): DensityFunction {
		return new LastCache(visitor.apply(this.input));
	}
	range() {
		return this.input.range();
	}
}

/** Whether a function's value can change with y. Unknown kinds are assumed to. */
function dependsOnY(fn: DensityFunction, memo: WeakMap<DensityFunction, boolean>): boolean {
	const known = memo.get(fn);
	if (known !== undefined) return known;
	let result: boolean;
	if (fn instanceof DensityFunction.Constant) {
		result = false;
	} else if (fn instanceof DensityFunction.Gradient) {
		result = fn.axis === 'y';
	} else if (fn instanceof DensityFunction.ShiftA || fn instanceof DensityFunction.ShiftB) {
		result = false; // sampled at y = 0
	} else if (fn instanceof DensityFunction.NoiseFunction) {
		result = fn.yScale !== 0 || [fn.shiftX, fn.shiftY, fn.shiftZ].some((c) => dependsOnY(c, memo));
	} else if (fn instanceof DensityFunction.EndOuterIslands) {
		result = false;
	} else {
		// Combinators depend on y when a child does; a leaf of another kind is assumed to.
		const children: DensityFunction[] = [];
		fn.mapChildren({ apply: (c: DensityFunction) => (children.push(c), c) });
		const leafKinds = [DensityFunction.HolderHolder].some((k) => fn instanceof k);
		result = children.length === 0 || leafKinds ? true : children.some((c) => dependsOnY(c, memo));
	}
	memo.set(fn, result);
	return result;
}

/** Swaps a random state's router for one using the faster nodes above. */
export function speedUp(rs: RandomState): RandomState {
	const memo = new WeakMap<DensityFunction, boolean>();
	const visitor: DensityFunction.Visitor = {
		apply(fn: DensityFunction): DensityFunction {
			if (fn instanceof DensityFunction.Interpolated) {
				return new FastInterpolated(fn.input, fn.cellSizeXz, fn.cellSizeY);
			}
			if (fn instanceof DensityFunction.Cache) {
				return dependsOnY(fn.input, memo) ? new LastCache(fn.input) : new ColumnCache(fn.input);
			}
			return fn;
		}
	};
	(rs as { router: NoiseRouter }).router = NoiseRouter.mapAll(rs.router, visitor);
	return rs;
}
