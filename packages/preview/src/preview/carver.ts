// A port of Minecraft 26.3's carvers (worldgen/carver), which Deepslate
// does not have. Given a carver's JSON and a seed, it finds the blocks the
// carver cuts in a chunk: the game's carving mask, before the mask is
// applied to the terrain. The random calls follow the game's order exactly
// (Java's 48-bit LCG, float maths rounded to float, Mth's sine table), so a
// seed gives the same caves as in game for a carver in the same slot of its
// biome's carver list.

const f = Math.fround;

// ---- Java's random ----------------------------------------------------

const TWO24 = 16777216;
const ML = 0xece66d; // 0x5DEECE66D split into 24-bit halves
const MH = 0x5de;
const M_BIG = 0x5deece66dn;
const MASK48 = (1n << 48n) - 1n;

/** Minecraft's LegacyRandomSource (java.util.Random's LCG), kept as two 24-bit halves so it needs no BigInt per call. */
export class JavaRandom {
	private hi = 0;
	private lo = 0;
	private gaussian: number | null = null;

	constructor(seed: bigint) {
		this.setSeed(seed);
	}

	setSeed(seed: bigint) {
		const s = (BigInt.asIntN(64, seed) ^ M_BIG) & MASK48;
		this.hi = Number(s >> 24n);
		this.lo = Number(s & 0xffffffn);
		this.gaussian = null;
	}

	next(bits: number): number {
		const p = this.lo * ML + 11;
		const lo = p % TWO24;
		const q = this.hi * ML + this.lo * MH + Math.floor(p / TWO24);
		this.hi = q % TWO24;
		this.lo = lo;
		const v = Math.floor((this.hi * TWO24 + this.lo) / 2 ** (48 - bits));
		return bits === 32 && v >= 2 ** 31 ? v - 2 ** 32 : v;
	}

	nextInt(bound: number): number {
		if (!(bound > 0)) throw new Error(`Bound must be positive (got ${bound})`);
		if ((bound & (bound - 1)) === 0) return Math.floor((bound * this.next(31)) / 2 ** 31);
		for (;;) {
			const r = this.next(31);
			const j = r % bound;
			if (r - j + (bound - 1) <= 0x7fffffff) return j;
		}
	}

	nextLong(): bigint {
		const a = BigInt(this.next(32));
		const b = BigInt(this.next(32));
		return BigInt.asIntN(64, (a << 32n) + b);
	}

	nextFloat(): number {
		return f(this.next(24) / TWO24);
	}

	nextDouble(): number {
		const a = this.next(26);
		const b = this.next(27);
		return (a * 2 ** 27 + b) / 2 ** 53;
	}

	/** java.util.Random's polar method, as MarsagliaPolarGaussian. */
	nextGaussian(): number {
		if (this.gaussian !== null) {
			const g = this.gaussian;
			this.gaussian = null;
			return g;
		}
		for (;;) {
			const a = 2 * this.nextDouble() - 1;
			const b = 2 * this.nextDouble() - 1;
			const s = a * a + b * b;
			if (s < 1 && s !== 0) {
				const m = Math.sqrt((-2 * Math.log(s)) / s);
				this.gaussian = b * m;
				return a * m;
			}
		}
	}

	/** WorldgenRandom.setLargeFeatureSeed: the random a carver starting in chunk (x, z) uses. */
	setLargeFeatureSeed(seed: bigint, x: number, z: number) {
		this.setSeed(seed);
		const a = this.nextLong();
		const b = this.nextLong();
		this.setSeed(BigInt.asIntN(64, (BigInt(x) * a) ^ (BigInt(z) * b) ^ seed));
	}
}

// ---- Mth ----------------------------------------------------------------

const SIN = new Float32Array(65536);
for (let i = 0; i < 65536; i++) SIN[i] = Math.sin((i * Math.PI * 2) / 65536);
const index = (v: number) => ((Math.trunc(v) % 65536) + 65536) % 65536;
const sin = (d: number) => SIN[index(d * 10430.378350470453)];
const cos = (d: number) => SIN[index(d * 10430.378350470453 + 16384)];
const PI_F = f(Math.PI);
const TAU_F = f(Math.PI * 2);
const HALF_PI_F = f(Math.PI / 2);

// ---- Providers --------------------------------------------------------

/** What height anchors resolve against: the generator's lowest y, height and sea level. */
export interface CarverContext {
	minY: number;
	height: number;
	seaLevel: number;
}

type Json = Record<string, unknown>;
const obj = (v: unknown): Json => (v && typeof v === 'object' && !Array.isArray(v) ? (v as Json) : {});
const num = (v: unknown, d = 0): number => (typeof v === 'number' ? v : d);
const typeOf = (v: Json) => String(v.type ?? '').replace(/^minecraft:/, '');

type FloatProvider = (r: JavaRandom) => number;
type IntProvider = (r: JavaRandom) => number;
type HeightProvider = (r: JavaRandom, c: CarverContext) => number;

const randomBetween = (r: JavaRandom, min: number, max: number) => f(f(r.nextFloat() * f(max - min)) + min);
const randomBetweenInclusive = (r: JavaRandom, min: number, max: number) => r.nextInt(max - min + 1) + min;
const nextIntBetween = (r: JavaRandom, min: number, max: number) => (min >= max ? min : r.nextInt(max - min + 1) + min);
const clamp = (v: number, lo: number, hi: number) => (v < lo ? lo : v > hi ? hi : v);

export function floatProvider(json: unknown, d = 0): FloatProvider {
	if (typeof json === 'number') {
		const v = f(json);
		return () => v;
	}
	const o = obj(json);
	switch (typeOf(o)) {
		case 'constant': {
			const v = f(num(o.value, d));
			return () => v;
		}
		case 'uniform': {
			const min = f(num(o.min_inclusive));
			const max = f(num(o.max_exclusive));
			return (r) => randomBetween(r, min, max);
		}
		case 'clamped_normal': {
			const mean = f(num(o.mean));
			const dev = f(num(o.deviation));
			const min = f(num(o.min));
			const max = f(num(o.max));
			return (r) => clamp(f(mean + f(f(r.nextGaussian()) * dev)), min, max);
		}
		case 'trapezoid': {
			const min = f(num(o.min));
			const max = f(num(o.max));
			const plateau = f(num(o.plateau));
			const range = f(max - min);
			const slope = f(f(range - plateau) / 2);
			const top = f(range - slope);
			return (r) => f(f(min + f(r.nextFloat() * top)) + f(r.nextFloat() * slope));
		}
	}
	const v = f(d);
	return () => v;
}

export function intProvider(json: unknown, d = 0): IntProvider {
	if (typeof json === 'number') return () => Math.trunc(json);
	const o = obj(json);
	const min = Math.trunc(num(o.min_inclusive));
	const max = Math.trunc(num(o.max_inclusive));
	switch (typeOf(o)) {
		case 'constant': {
			const v = Math.trunc(num(o.value, d));
			return () => v;
		}
		case 'uniform':
			return (r) => randomBetweenInclusive(r, min, max);
		case 'biased_to_bottom':
			return (r) => min + r.nextInt(r.nextInt(max - min + 1) + 1);
		case 'very_biased_to_bottom':
			return (r) => min + r.nextInt(r.nextInt(r.nextInt(max - min + 1) + 1) + 1);
		case 'clamped': {
			const source = intProvider(o.source);
			return (r) => clamp(source(r), min, max);
		}
		case 'clamped_normal': {
			const mean = f(num(o.mean));
			const dev = f(num(o.deviation));
			return (r) => Math.trunc(clamp(f(mean + f(f(r.nextGaussian()) * dev)), min, max));
		}
		case 'weighted_list': {
			const entries = (Array.isArray(o.distribution) ? o.distribution : []).map((e) => ({ weight: Math.trunc(num(obj(e).weight, 1)), data: intProvider(obj(e).data) }));
			const total = entries.reduce((s, e) => s + e.weight, 0);
			return (r) => {
				if (total <= 0) return 0;
				let k = r.nextInt(total);
				for (const e of entries) {
					k -= e.weight;
					if (k < 0) return e.data(r);
				}
				return 0;
			};
		}
	}
	return () => d;
}

function anchor(json: unknown): (c: CarverContext) => number {
	const o = obj(json);
	if (o.absolute !== undefined) return () => num(o.absolute);
	if (o.above_bottom !== undefined) return (c) => c.minY + num(o.above_bottom);
	if (o.below_top !== undefined) return (c) => c.minY + c.height - 1 - num(o.below_top);
	if (o.relative_to_sea_level !== undefined) return (c) => c.seaLevel + num(o.relative_to_sea_level);
	return () => 0;
}

export function heightProvider(json: unknown): HeightProvider {
	const o = obj(json);
	const type = typeOf(o);
	if (!type) {
		const a = anchor(json);
		return (_, c) => a(c);
	}
	const lo = anchor(o.min_inclusive);
	const hi = anchor(o.max_inclusive);
	const inner = Math.trunc(num(o.inner, 1));
	switch (type) {
		case 'constant': {
			const a = anchor(o.value);
			return (_, c) => a(c);
		}
		case 'uniform':
			return (r, c) => {
				const min = lo(c);
				const max = hi(c);
				return min > max ? min : randomBetweenInclusive(r, min, max);
			};
		case 'biased_to_bottom':
			return (r, c) => {
				const min = lo(c);
				const max = hi(c);
				if (max - min - inner + 1 <= 0) return min;
				return r.nextInt(r.nextInt(max - min - inner + 1) + inner) + min;
			};
		case 'very_biased_to_bottom':
			return (r, c) => {
				const min = lo(c);
				const max = hi(c);
				if (max - min - inner + 1 <= 0) return min;
				const a = nextIntBetween(r, min + inner, max);
				const b = nextIntBetween(r, min, a - 1);
				return nextIntBetween(r, min, b - 1 + inner);
			};
		case 'trapezoid': {
			const plateau = Math.trunc(num(o.plateau));
			return (r, c) => {
				const min = lo(c);
				const max = hi(c);
				if (min > max) return min;
				const range = max - min;
				if (plateau >= range) return randomBetweenInclusive(r, min, max);
				const slope = Math.trunc((range - plateau) / 2);
				return min + randomBetweenInclusive(r, 0, range - slope) + randomBetweenInclusive(r, 0, slope);
			};
		}
		case 'weighted_list': {
			const entries = (Array.isArray(o.distribution) ? o.distribution : []).map((e) => ({ weight: Math.trunc(num(obj(e).weight, 1)), data: heightProvider(obj(e).data) }));
			const total = entries.reduce((s, e) => s + e.weight, 0);
			return (r, c) => {
				if (total <= 0) return 0;
				let k = r.nextInt(total);
				for (const e of entries) {
					k -= e.weight;
					if (k < 0) return e.data(r, c);
				}
				return 0;
			};
		}
	}
	return () => 0;
}

// ---- Carvers ------------------------------------------------------------

/**
 * One step of a tunnel that carves an ellipsoid: its centre and radii, and
 * the step index the game's canReach check uses.
 */
interface Step {
	x: number;
	y: number;
	z: number;
	hr: number;
	vr: number;
	i: number;
}

/**
 * A tunnel as the game walks it, independent of the chunk being carved:
 * the steps that carve, then the branches it splits into (if it splits).
 * The game stops a tunnel early for a chunk it can no longer reach, so
 * which steps carve depends on the chunk; that is decided when carving.
 */
interface Tunnel {
	steps: Step[];
	/** Steps' count and thickness, for canReach. Rooms have no reach check. */
	count: number;
	reach: number;
	room: boolean;
	branches: Tunnel[];
	skip: Skip;
}

/** The game's CarveSkipChecker: relative x, y, z inside the ellipsoid and the block y. */
type Skip = (rx: number, ry: number, rz: number, y: number) => boolean;

export interface Carver {
	type: 'cave' | 'canyon';
	probability: number;
	/** The tunnels a carver starting in a chunk makes, or none. */
	plan(r: JavaRandom, c: CarverContext, chunkX: number, chunkZ: number): Tunnel[];
}

/** Reads a carver's JSON (the 26.3 format), or throws if it isn't one this can run. */
export function readCarver(json: unknown): Carver {
	const o = obj(json);
	const type = typeOf(o);
	const probability = f(num(o.probability));
	if (type === 'cave') return caveCarver(o, probability);
	if (type === 'canyon') return canyonCarver(o, probability);
	throw new Error(`Carvers of type ${type || '(missing)'} have no preview.`);
}

const RANGE_BLOCKS = (4 * 2 - 1) * 16; // getRange() is 4 chunks

function caveCarver(o: Json, probability: number): Carver {
	const y = heightProvider(o.y);
	const count = intProvider(o.count);
	const thickness = floatProvider(o.thickness);
	const weird = o.weird_thickness_bias === true;
	const roomRatio = floatProvider(o.room_vertical_radius_multiplier);
	const hMul = floatProvider(o.horizontal_radius_multiplier);
	const vMul = floatProvider(o.vertical_radius_multiplier);
	const startRatio = floatProvider(o.start_vertical_radius_multiplier, 1);
	const floorLevel = floatProvider(o.floor_level);

	const thick = (r: JavaRandom) => {
		let t = thickness(r);
		if (weird && r.nextInt(10) === 0) t = f(t * f(f(f(r.nextFloat() * r.nextFloat()) * 3) + 1));
		return t;
	};

	function tunnel(seed: bigint, x: number, yy: number, z: number, hm: number, vm: number, thickness: number, yaw: number, pitch: number, from: number, n: number, ratio: number, skip: Skip): Tunnel {
		const r = new JavaRandom(seed);
		const t: Tunnel = { steps: [], count: n, reach: f(f(thickness + 2) + 16), room: false, branches: [], skip };
		const branchAt = r.nextInt(Math.trunc(n / 2)) + Math.trunc(n / 4);
		const steep = r.nextInt(6) === 0;
		let dYaw = 0;
		let dPitch = 0;
		for (let i = from; i < n; i++) {
			const hr = 1.5 + f(sin(f(f(PI_F * i) / n)) * thickness);
			const vr = hr * ratio;
			const h = cos(pitch);
			x += f(cos(yaw) * h);
			yy += sin(pitch);
			z += f(sin(yaw) * h);
			pitch = f(pitch * (steep ? f(0.92) : f(0.7)));
			pitch = f(pitch + f(dPitch * f(0.1)));
			yaw = f(yaw + f(dYaw * f(0.1)));
			dPitch = f(dPitch * f(0.9));
			dYaw = f(dYaw * f(0.75));
			dPitch = f(dPitch + f(f(f(r.nextFloat() - r.nextFloat()) * r.nextFloat()) * 2));
			dYaw = f(dYaw + f(f(f(r.nextFloat() - r.nextFloat()) * r.nextFloat()) * 4));
			if (i === branchAt && thickness > 1) {
				const s1 = r.nextLong();
				const w1 = f(f(r.nextFloat() * 0.5) + 0.5);
				t.branches.push(tunnel(s1, x, yy, z, hm, vm, w1, f(yaw - HALF_PI_F), f(pitch / 3), i, n, 1, skip));
				const s2 = r.nextLong();
				const w2 = f(f(r.nextFloat() * 0.5) + 0.5);
				t.branches.push(tunnel(s2, x, yy, z, hm, vm, w2, f(yaw + HALF_PI_F), f(pitch / 3), i, n, 1, skip));
				return t;
			}
			if (r.nextInt(4) !== 0) t.steps.push({ x, y: yy, z, hr: hr * hm, vr: vr * vm, i });
		}
		return t;
	}

	return {
		type: 'cave',
		probability,
		plan(r, c, cx, cz) {
			const out: Tunnel[] = [];
			const n = count(r);
			for (let k = 0; k < n; k++) {
				const x = cx * 16 + r.nextInt(16);
				const yy = y(r, c);
				const z = cz * 16 + r.nextInt(16);
				const hm = hMul(r);
				const vm = vMul(r);
				const start = startRatio(r);
				const floor = floorLevel(r);
				const skip: Skip = (rx, ry, rz) => ry <= floor || rx * rx + ry * ry + rz * rz >= 1;
				let tunnels = 1;
				if (r.nextInt(4) === 0) {
					const ratio = roomRatio(r);
					const radius = f(1 + f(r.nextFloat() * 6));
					const hr = 1.5 + f(sin(HALF_PI_F) * radius);
					out.push({ steps: [{ x: x + 1, y: yy, z, hr, vr: hr * ratio, i: 0 }], count: 0, reach: 0, room: true, branches: [], skip });
					tunnels += r.nextInt(4);
				}
				for (let j = 0; j < tunnels; j++) {
					const yaw = f(r.nextFloat() * TAU_F);
					const pitch = f(f(r.nextFloat() - 0.5) / 4);
					const thickness = thick(r);
					const steps = RANGE_BLOCKS - r.nextInt(Math.trunc(RANGE_BLOCKS / 4));
					out.push(tunnel(r.nextLong(), x, yy, z, hm, vm, thickness, yaw, pitch, 0, steps, start, skip));
				}
			}
			return out;
		}
	};
}

function canyonCarver(o: Json, probability: number): Carver {
	const y = heightProvider(o.y);
	const verticalRotation = floatProvider(o.vertical_rotation);
	const shape = obj(o.shape);
	const yScale = floatProvider(shape.y_scale, 3);
	const thickness = floatProvider(shape.thickness);
	const distance = floatProvider(shape.distance_factor);
	const hFactor = floatProvider(shape.horizontal_radius_factor);
	const vDefault = f(num(shape.vertical_radius_default_factor, 1));
	const vCentre = f(num(shape.vertical_radius_center_factor));
	const smoothness = Math.trunc(num(shape.width_smoothness, 3));

	return {
		type: 'canyon',
		probability,
		plan(r, c, cx, cz) {
			let x = cx * 16 + r.nextInt(16);
			let yy = y(r, c);
			let z = cz * 16 + r.nextInt(16);
			let yaw = f(r.nextFloat() * TAU_F);
			let pitch = verticalRotation(r);
			const ratio = yScale(r);
			const w = thickness(r);
			const n = Math.trunc(f(RANGE_BLOCKS * distance(r)));
			const tr = new JavaRandom(r.nextLong());
			// Width factors per block y, from the bottom.
			const widths = new Float32Array(c.height);
			let wf = 1;
			for (let j = 0; j < c.height; j++) {
				if (j === 0 || (smoothness > 0 && tr.nextInt(smoothness) === 0)) wf = f(1 + f(tr.nextFloat() * tr.nextFloat()));
				widths[j] = f(wf * wf);
			}
			const skip: Skip = (rx, ry, rz, by) => (rx * rx + rz * rz) * widths[by - c.minY - 1] + (ry * ry) / 6 >= 1;
			const t: Tunnel = { steps: [], count: n, reach: f(f(w + 2) + 16), room: false, branches: [], skip };
			let dYaw = 0;
			let dPitch = 0;
			for (let i = 0; i < n; i++) {
				let hr = 1.5 + f(sin(f(f(i * PI_F) / n)) * w);
				let vr = hr * ratio;
				hr *= hFactor(tr);
				const k = f(1 - f(Math.abs(f(0.5 - f(i / n))) * 2));
				vr = f(vDefault + f(vCentre * k)) * vr * randomBetween(tr, f(0.75), 1);
				const h = cos(pitch);
				const v = sin(pitch);
				x += f(cos(yaw) * h);
				yy += v;
				z += f(sin(yaw) * h);
				pitch = f(pitch * f(0.7));
				pitch = f(pitch + f(dPitch * f(0.05)));
				yaw = f(yaw + f(dYaw * f(0.05)));
				dPitch = f(dPitch * f(0.8));
				dYaw = f(dYaw * f(0.5));
				dPitch = f(dPitch + f(f(f(tr.nextFloat() - tr.nextFloat()) * tr.nextFloat()) * 2));
				dYaw = f(dYaw + f(f(f(tr.nextFloat() - tr.nextFloat()) * tr.nextFloat()) * 4));
				if (tr.nextInt(4) !== 0) t.steps.push({ x, y: yy, z, hr, vr, i });
			}
			return [t];
		}
	};
}

// ---- Carving a chunk ----------------------------------------------------

/**
 * The blocks carved in one chunk, as the game's CarvingMask: indexed
 * ((x * 16 + z) * height + y - minY) with chunk-local x and z.
 */
export type Mask = Uint8Array;

/** The lowest and highest y a carver may cut (the game keeps the bottom block and the top 8). */
export function carvedRange(c: CarverContext): [number, number] {
	return [c.minY + 1, c.minY + c.height - 1 - 7];
}

/** The tunnels starting in each chunk, cached, since every chunk within 8 asks for them. */
export class CarverRun {
	private plans = new Map<string, Tunnel[] | null>();
	private random = new JavaRandom(0n);

	/**
	 * @param slot the carver's index in its biome's carver list, which the
	 * game adds to the seed.
	 */
	readonly carver: Carver;
	readonly seed: bigint;
	readonly slot: number;
	readonly context: CarverContext;
	constructor(carver: Carver, seed: bigint, slot: number, context: CarverContext) {
		this.carver = carver;
		this.seed = seed;
		this.slot = slot;
		this.context = context;
	}

	tunnelsAt(cx: number, cz: number): Tunnel[] | null {
		const key = `${cx},${cz}`;
		let plan = this.plans.get(key);
		if (plan === undefined) {
			const r = this.random;
			r.setLargeFeatureSeed(BigInt.asIntN(64, this.seed + BigInt(this.slot)), cx, cz);
			plan = r.nextFloat() <= this.carver.probability ? this.carver.plan(r, this.context, cx, cz) : null;
			if (this.plans.size > 20000) this.plans.clear();
			this.plans.set(key, plan);
		}
		return plan;
	}

	/** The chunks within 8 of (cx, cz) that start the carver. */
	startsNear(cx: number, cz: number): [number, number][] {
		const out: [number, number][] = [];
		for (let dx = -8; dx <= 8; dx++) for (let dz = -8; dz <= 8; dz++) if (this.tunnelsAt(cx + dx, cz + dz)) out.push([cx + dx, cz + dz]);
		return out;
	}

	/** Carves chunk (cx, cz) into mask (cleared first), returning how many blocks it cut. */
	carve(cx: number, cz: number, mask: Mask): number {
		mask.fill(0);
		const { height, minY } = this.context;
		const [lo, hi] = carvedRange(this.context);
		const midX = cx * 16 + 8;
		const midZ = cz * 16 + 8;
		let carved = 0;
		const ellipsoid = (s: Step, skip: Skip) => {
			const limit = 16 + s.hr * 2;
			if (Math.abs(s.x - midX) > limit || Math.abs(s.z - midZ) > limit) return;
			const x0 = Math.max(Math.floor(s.x - s.hr) - cx * 16 - 1, 0);
			const x1 = Math.min(Math.floor(s.x + s.hr) - cx * 16, 15);
			const y0 = Math.max(Math.floor(s.y - s.vr) - 1, lo);
			const y1 = Math.min(Math.floor(s.y + s.vr) + 1, hi);
			const z0 = Math.max(Math.floor(s.z - s.hr) - cz * 16 - 1, 0);
			const z1 = Math.min(Math.floor(s.z + s.hr) - cz * 16, 15);
			for (let lx = x0; lx <= x1; lx++) {
				const rx = (cx * 16 + lx + 0.5 - s.x) / s.hr;
				for (let lz = z0; lz <= z1; lz++) {
					const rz = (cz * 16 + lz + 0.5 - s.z) / s.hr;
					if (rx * rx + rz * rz >= 1) continue;
					const base = (lx * 16 + lz) * height - minY;
					for (let y = y1; y > y0; y--) {
						const ry = (y - 0.5 - s.y) / s.vr;
						if (skip(rx, ry, rz, y)) continue;
						if (!mask[base + y]) {
							mask[base + y] = 1;
							carved++;
						}
					}
				}
			}
		};
		const walk = (t: Tunnel) => {
			for (const s of t.steps) {
				// canReach: the game gives up on a tunnel that can no longer reach this chunk.
				if (!t.room) {
					const dx = s.x - midX;
					const dz = s.z - midZ;
					const left = t.count - s.i;
					if (dx * dx + dz * dz - left * left > t.reach * t.reach) return;
				}
				ellipsoid(s, t.skip);
			}
			for (const b of t.branches) walk(b);
		};
		for (const [sx, sz] of this.startsNear(cx, cz)) {
			for (const t of this.tunnelsAt(sx, sz)!) walk(t);
		}
		return carved;
	}
}
