// The volume view: density in 3D, ray-marched on the GPU.
//
// The worker samples tiles of VOLUME_TILE × VOLUME_TILE columns on a grid
// (noise settings at their cell corners, which the GPU's trilinear
// filtering blends as the game's interpolated nodes do). The tiles around
// the camera live in a window of W × W slots in one 3D texture that wraps
// around: tile (tx, tz) always sits in slot (tx mod W, tz mod W), so moving
// only fills the slots of the tiles that come into view, and nothing is
// copied. Beside the density, each tile carries:
//
// - how fast density can change near each cell, so rays sphere-trace:
//   they jump through open air or solid stone and slow down near a surface;
// - the highest solid block of each column, so the caves look can tell air
//   under the ground (a cave) from the open air above it.

import * as THREE from 'three';
import { OrbitControls } from 'three/addons/controls/OrbitControls.js';
import { VOLUME_TILE, type VolumeTile } from './protocol';

/** How each look draws density. */
export type VolumeLook = 'caves' | 'shell' | 'solid';
const lookIndex: Record<VolumeLook, number> = { caves: 0, shell: 1, solid: 2 };

/** Tiles from the camera's tile to each side, so the window is 2R + 1 tiles across. */
const R = 2;
/** Slots per side: one spare, so the window's far edge never wraps onto its near one. */
const W = 2 * R + 2;
/** Tiles kept on the CPU, so moving back doesn't sample them again. */
const MAX_CACHED = 96;

/** A tile as the window stores it, on the function's finest grid. */
interface Tile {
	tx: number;
	tz: number;
	/** How much coarser than the grid it was sampled (1 for the grid itself). */
	lod: number;
	values: Float32Array;
	/** Per cell: the steepest change per block within two cells (see lipOf). */
	lip: Float32Array;
	/** Per block column: the highest y whose density is over the threshold (-Infinity for none). */
	tops: Float32Array | null;
	topsFor: number;
}

/** The grid the current function is sampled on, from its first tile. */
interface Grid {
	sx: number;
	sy: number;
	nx: number;
	ny: number;
	minY: number;
	/** Corners per tile along x and z, the last one left to the next tile. */
	c: number;
}

export interface ValueRange {
	min: number;
	max: number;
	median: number;
}

export interface VolumeSettings {
	look: VolumeLook;
	threshold: number;
	/** The highest y drawn. */
	cutY: number;
	/** How much of the window to keep along x and z, 0 to 1, cut from the east and south. */
	cutX: number;
	cutZ: number;
	haze: number;
	caveOpacity: number;
	/** Draws whole blocks, as the game places them, rather than the smooth surface between samples. */
	blocks: boolean;
}

export const defaultSettings: VolumeSettings = { look: 'caves', threshold: 0, cutY: 320, cutX: 1, cutZ: 1, haze: 0.004, caveOpacity: 0.55, blocks: true };

const mod = (a: number, n: number) => ((a % n) + n) % n;

/** A tile the window wants, at a level of detail. */
export interface TileWant {
	tx: number;
	tz: number;
	lod: number;
}

/** Screen pixels between samples that still look smooth: the GPU blends between them. */
const PIXELS_PER_SAMPLE = 4;

/** Puts a coarse tile on the finest grid, blending its samples trilinearly. */
function upsample(g: Grid, tile: Extract<VolumeTile, { values: Float32Array }>): Float32Array {
	const { lod } = tile;
	if (lod === 1) return tile.values;
	const n = VOLUME_TILE / (g.sx * lod) + 1;
	const layers = Math.floor((g.ny - 1) / lod) + 1;
	const src = tile.values;
	const at = (i: number, j: number, k: number) => src[(j * n + k) * n + i];
	const out = new Float32Array(g.nx * g.nx * g.ny);
	for (let j = 0; j < g.ny; j++) {
		const gj = Math.min(j / lod, layers - 1);
		const j0 = Math.min(Math.floor(gj), layers - 2 < 0 ? 0 : layers - 2);
		const ty = layers > 1 ? gj - j0 : 0;
		const j1 = Math.min(j0 + 1, layers - 1);
		for (let k = 0; k < g.nx; k++) {
			const gk = k / lod;
			const k0 = Math.min(Math.floor(gk), n - 2);
			const tz = gk - k0;
			for (let i = 0; i < g.nx; i++) {
				const gi = i / lod;
				const i0 = Math.min(Math.floor(gi), n - 2);
				const tx = gi - i0;
				const lo =
					(1 - tz) * ((1 - tx) * at(i0, j0, k0) + tx * at(i0 + 1, j0, k0)) + tz * ((1 - tx) * at(i0, j0, k0 + 1) + tx * at(i0 + 1, j0, k0 + 1));
				const hi =
					(1 - tz) * ((1 - tx) * at(i0, j1, k0) + tx * at(i0 + 1, j1, k0)) + tz * ((1 - tx) * at(i0, j1, k0 + 1) + tx * at(i0 + 1, j1, k0 + 1));
				out[(j * g.nx + k) * g.nx + i] = lo + ty * (hi - lo);
			}
		}
	}
	return out;
}

/**
 * Per cell, the steepest slope of its edges in density per block; then
 * the largest within two cells, which is how far one step can go. Cells
 * within two of the tile's side can't see the neighbour tile, so they take
 * the tile's largest instead.
 */
function lipOf(g: Grid, values: Float32Array): Float32Array {
	const { sx, sy, nx, ny, c } = g;
	const cy = ny - 1;
	const at = (i: number, j: number, k: number) => values[(j * nx + k) * nx + i];
	const cell = new Float32Array(c * cy * c);
	let most = 0;
	for (let j = 0; j < cy; j++) {
		for (let k = 0; k < c; k++) {
			for (let i = 0; i < c; i++) {
				let gx = 0;
				let gy = 0;
				let gz = 0;
				for (let a = 0; a < 2; a++) {
					for (let b = 0; b < 2; b++) {
						gx = Math.max(gx, Math.abs(at(i + 1, j + a, k + b) - at(i, j + a, k + b)) / sx);
						gy = Math.max(gy, Math.abs(at(i + a, j + 1, k + b) - at(i + a, j, k + b)) / sy);
						gz = Math.max(gz, Math.abs(at(i + a, j + b, k + 1) - at(i + a, j + b, k)) / sx);
					}
				}
				const l = Math.hypot(gx, gy, gz);
				cell[(j * c + k) * c + i] = l;
				if (l > most) most = l;
			}
		}
	}
	// The largest within two cells along each axis in turn (a separable max filter).
	const a = cell;
	const b = new Float32Array(c * cy * c);
	const idx = (i: number, j: number, k: number) => (j * c + k) * c + i;
	for (let j = 0; j < cy; j++)
		for (let k = 0; k < c; k++)
			for (let i = 0; i < c; i++) {
				let m = 0;
				for (let d = Math.max(0, i - 2); d <= Math.min(c - 1, i + 2); d++) m = Math.max(m, a[idx(d, j, k)]);
				b[idx(i, j, k)] = m;
			}
	for (let j = 0; j < cy; j++)
		for (let k = 0; k < c; k++)
			for (let i = 0; i < c; i++) {
				let m = 0;
				for (let d = Math.max(0, k - 2); d <= Math.min(c - 1, k + 2); d++) m = Math.max(m, b[idx(i, j, d)]);
				a[idx(i, j, k)] = m;
			}
	const out = new Float32Array(c * cy * c);
	for (let j = 0; j < cy; j++)
		for (let k = 0; k < c; k++)
			for (let i = 0; i < c; i++) {
				let m = 0;
				if (i < 2 || k < 2 || i >= c - 2 || k >= c - 2) m = most;
				else for (let d = Math.max(0, j - 2); d <= Math.min(cy - 1, j + 2); d++) m = Math.max(m, a[idx(i, d, k)]);
				out[idx(i, j, k)] = m * 1.02 + 1e-5;
			}
	return out;
}

/** The highest y in each of a tile's block columns where trilinear density is over the threshold. */
function topsOf(g: Grid, values: Float32Array, threshold: number): Float32Array {
	const { sx, sy, nx, ny, minY } = g;
	const n = VOLUME_TILE;
	const tops = new Float32Array(n * n);
	const at = (i: number, j: number, k: number) => values[(j * nx + k) * nx + i];
	for (let bz = 0; bz < n; bz++) {
		const gz = bz / sx;
		const k = Math.floor(gz);
		const tz = gz - k;
		for (let bx = 0; bx < n; bx++) {
			const gx = bx / sx;
			const i = Math.floor(gx);
			const tx = gx - i;
			let top = -Infinity;
			// The column's corner values, bilinear in x and z, then linear up each cell.
			let above = NaN;
			for (let j = ny - 1; j >= 0 && top === -Infinity; j--) {
				const v =
					(1 - tz) * ((1 - tx) * at(i, j, k) + tx * at(i + 1, j, k)) + tz * ((1 - tx) * at(i, j, k + 1) + tx * at(i + 1, j, k + 1));
				if (v > threshold) {
					// The crossing between this corner and the one above, to the block.
					const t = Number.isNaN(above) ? 0 : (v - threshold) / (v - above);
					top = Math.floor(minY + (j + Math.min(1, Math.max(0, t))) * sy);
				}
				above = v;
			}
			tops[bz * n + bx] = top;
		}
	}
	return tops;
}

const vertexShader = /* glsl */ `
	out vec3 vPos;
	void main() {
		vPos = (modelMatrix * vec4(position, 1.0)).xyz;
		gl_Position = projectionMatrix * viewMatrix * vec4(vPos, 1.0);
	}`;

const fragmentShader = /* glsl */ `
	precision highp float;
	precision highp sampler3D;
	uniform sampler3D vol, lipTex;
	uniform sampler2D topTex, readyTex;
	uniform vec3 camPos, boxMin, boxMax, bg;
	// The scene's origin modulo the window, in blocks, so texture lookups wrap the same way tiles do.
	uniform vec2 wrap;
	uniform float sx, sy, minY, ny, cornersAcross, thr, haze, caveA;
	uniform int look;
	uniform bool cubes;
	in vec3 vPos;
	out vec4 outColor;

	const float TILE = ${VOLUME_TILE.toFixed(1)};
	const float SLOTS = ${W.toFixed(1)};

	vec3 world(vec3 p) { return vec3(p.x + wrap.x, p.y, p.z + wrap.y); }
	float dens(vec3 p) {
		vec3 w = world(p);
		vec3 t = vec3(w.x / sx, w.z / sx, (w.y - minY) / sy);
		return texture(vol, (t + 0.5) / vec3(cornersAcross, cornersAcross, ny)).r - thr;
	}
	float lipAt(vec3 p) {
		vec3 w = world(p);
		vec3 c = floor(vec3(w.x / sx, w.z / sx, clamp((w.y - minY) / sy, 0.0, ny - 2.0)));
		return texture(lipTex, (c + 0.5) / vec3(cornersAcross, cornersAcross, ny - 1.0)).r;
	}
	float topAt(vec3 p) {
		vec3 w = floor(world(p));
		return texture(topTex, (w.xz + 0.5) / (SLOTS * TILE)).r;
	}
	bool readyTile(vec3 p) {
		vec3 w = world(p);
		return texture(readyTex, (floor(w.xz / TILE) + 0.5) / SLOTS).r > 0.5;
	}
	// A point's value blends the corners of its cell, which may reach into the next tile.
	bool ready(vec3 p) {
		return readyTile(p) && readyTile(p + vec3(sx, 0.0, 0.0)) && readyTile(p + vec3(0.0, 0.0, sx)) && readyTile(p + vec3(sx, 0.0, sx));
	}
	vec3 normalAt(vec3 p) {
		vec3 g = vec3(dens(p + vec3(1, 0, 0)) - dens(p - vec3(1, 0, 0)),
		              dens(p + vec3(0, 1, 0)) - dens(p - vec3(0, 1, 0)),
		              dens(p + vec3(0, 0, 1)) - dens(p - vec3(0, 0, 1)));
		return -normalize(g + 1e-7);
	}
	// Height: deep magenta, red, amber, green, sky blue, from the bottom of the world to y 200.
	vec3 ramp(float y) {
		float t = clamp((y - minY) / (200.0 - minY), 0.0, 1.0) * 4.0;
		vec3 c0 = vec3(0.45, 0.12, 0.55), c1 = vec3(0.86, 0.27, 0.22), c2 = vec3(0.93, 0.68, 0.20), c3 = vec3(0.25, 0.66, 0.45), c4 = vec3(0.35, 0.58, 0.92);
		if (t < 1.0) return mix(c0, c1, t);
		if (t < 2.0) return mix(c1, c2, t - 1.0);
		if (t < 3.0) return mix(c2, c3, t - 2.0);
		return mix(c3, c4, t - 3.0);
	}
	float light(vec3 n, vec3 rd) {
		n = faceforward(n, rd, n);
		vec3 L = normalize(vec3(0.5, 0.9, 0.3));
		return 0.38 + 0.52 * max(dot(n, L), 0.0) + 0.18 * max(dot(n, -rd), 0.0);
	}
	vec2 hitBox(vec3 ro, vec3 rd) {
		vec3 inv = 1.0 / rd;
		vec3 a = (boxMin - ro) * inv, b = (boxMax - ro) * inv;
		vec3 mn = min(a, b), mx = max(a, b);
		return vec2(max(max(mn.x, mn.y), max(mn.z, 0.0)), min(min(mx.x, mx.y), mx.z));
	}
	void add(inout vec4 acc, vec3 col, float a) {
		acc.rgb += (1.0 - acc.a) * a * col;
		acc.a += (1.0 - acc.a) * a;
	}
	void hazeOver(inout vec4 acc, float blocks) {
		add(acc, vec3(0.62, 0.64, 0.66), 1.0 - pow(1.0 - haze, blocks * 2.0)); // haze per half block
	}

	// Blocks: a block is solid where density at its position (its lowest
	// corner, where the game samples it) is over the threshold. The ray
	// steps block by block (Amanatides and Woo) near a surface, and jumps
	// ahead where the step bound says no block's solidity can change.
	vec4 marchBlocks(vec3 ro, vec3 rd, vec2 tt) {
		vec4 acc = vec4(0.0);
		float t = tt.x + 1e-4;
		vec3 cell = floor(ro + rd * t);
		vec3 dir = sign(rd);
		vec3 tDelta = abs(1.0 / rd);
		vec3 tMax = (cell + max(dir, 0.0) - ro) / rd;
		bool ok = ready(cell);
		bool solid = ok && dens(cell) > 0.0;
		if (look == 2 && solid) {
			// The solid look starts on a cut face: shaded flat by height.
			return vec4(ramp(cell.y + 0.5) * 0.78, 1.0);
		}
		float maxJump = 2.0 * min(sx, sy) - 1.0;
		for (int i = 0; i < 4000; i++) {
			if (acc.a > 0.97) break;
			if (ok) {
				// Within s blocks along the ray, every block corner is within s + 1.74 blocks of this block's centre.
				vec3 c = cell + 0.5;
				float L = lipAt(c);
				float s = min(0.9 * abs(dens(c)) / L - 1.74, maxJump);
				if (s > 2.0) {
					if (look == 0 && solid && haze > 0.0) hazeOver(acc, s);
					t += s;
					if (t >= tt.y) break;
					cell = floor(ro + rd * t);
					tMax = (cell + max(dir, 0.0) - ro) / rd;
					continue;
				}
			}
			float tPrev = t;
			int axis;
			if (tMax.x < tMax.y && tMax.x < tMax.z) { axis = 0; t = tMax.x; tMax.x += tDelta.x; cell.x += dir.x; }
			else if (tMax.y < tMax.z) { axis = 1; t = tMax.y; tMax.y += tDelta.y; cell.y += dir.y; }
			else { axis = 2; t = tMax.z; tMax.z += tDelta.z; cell.z += dir.z; }
			if (t >= tt.y) break;
			if (look == 0 && solid && ok && haze > 0.0) hazeOver(acc, t - tPrev);
			if (!ready(cell)) { ok = false; continue; }
			bool s1 = dens(cell) > 0.0;
			if (!ok) { ok = true; solid = s1; continue; }
			if (s1 != solid) {
				vec3 moved = vec3(axis == 0 ? dir.x : 0.0, axis == 1 ? dir.y : 0.0, axis == 2 ? dir.z : 0.0);
				vec3 solidCell = s1 ? cell : cell - moved;
				vec3 airCell = s1 ? cell - moved : cell;
				// Flat light per face, as blocks are lit in the game: tops brightest.
				float shade = axis == 1 ? (moved.y < 0.0 ? 1.0 : 0.5) : axis == 0 ? 0.8 : 0.65;
				// A darker line along each block's edges.
				vec3 h = fract(ro + rd * t);
				vec2 uv = axis == 0 ? h.yz : axis == 1 ? h.xz : h.xy;
				float edge = min(min(uv.x, uv.y), min(1.0 - uv.x, 1.0 - uv.y));
				if (edge < 0.06) shade *= 0.82;
				float y = solidCell.y + 0.5;
				vec3 col;
				float a;
				if (look == 2) {
					col = mix(vec3(0.72, 0.74, 0.70), ramp(y), 0.35) * shade;
					a = 1.0;
				} else if (look == 1) {
					col = ramp(y) * shade;
					a = 0.32;
				} else if (airCell.y < topAt(airCell + 0.5) - 3.0) {
					col = ramp(y) * shade;
					a = caveA;
				} else {
					col = vec3(0.80, 0.83, 0.80) * shade;
					a = 0.07;
				}
				add(acc, col, a);
			}
			solid = s1;
		}
		return acc;
	}

	void main() {
		vec3 ro = camPos;
		vec3 rd = normalize(vPos - camPos);
		vec2 tt = hitBox(ro, rd);
		if (tt.x >= tt.y) { outColor = vec4(bg, 1.0); return; }
		if (cubes) {
			vec4 acc = marchBlocks(ro, rd, tt);
			outColor = vec4(acc.rgb + (1.0 - acc.a) * bg, 1.0);
			return;
		}
		float minStep = min(1.0, sx / 4.0), maxStep = 2.0 * min(sx, sy);
		vec4 acc = vec4(0.0);
		float t = tt.x + 1e-3;
		vec3 p = ro + rd * t;
		bool ok = ready(p);
		float d = ok ? dens(p) : 0.0;
		if (look == 2 && ok && d > 0.0) {
			// The solid look starts on a cut face: shaded flat by height.
			outColor = vec4(ramp(p.y) * 0.78, 1.0);
			return;
		}
		for (int i = 0; i < 3000; i++) {
			float dt = ok ? clamp(0.9 * abs(d) / lipAt(p), minStep, maxStep) : sx;
			float t1 = t + dt;
			if (t1 > tt.y || acc.a > 0.97) break;
			vec3 p1 = ro + rd * t1;
			bool ok1 = ready(p1);
			if (!ok1) { t = t1; p = p1; ok = false; continue; }
			float d1 = dens(p1);
			if (!ok) { t = t1; p = p1; d = d1; ok = true; continue; }
			if (look == 0 && d1 > 0.0 && haze > 0.0) {
				float a = 1.0 - pow(1.0 - haze, dt * 2.0); // haze per half block
				acc.rgb += (1.0 - acc.a) * a * vec3(0.62, 0.64, 0.66);
				acc.a += (1.0 - acc.a) * a;
			}
			if ((d > 0.0) != (d1 > 0.0)) {
				float a0 = t, a1 = t1;
				for (int k = 0; k < 6; k++) {
					float m = 0.5 * (a0 + a1);
					if ((dens(ro + rd * m) > 0.0) == (d > 0.0)) a0 = m; else a1 = m;
				}
				vec3 pc = ro + rd * (0.5 * (a0 + a1));
				float l = light(normalAt(pc), rd);
				vec3 col;
				float a;
				if (look == 2) {
					col = mix(vec3(0.72, 0.74, 0.70), ramp(pc.y), 0.35) * l;
					a = 1.0;
				} else if (look == 1) {
					col = ramp(pc.y) * l;
					a = 0.32;
				} else if (pc.y < topAt(pc) - 3.0) {
					col = ramp(pc.y) * l;
					a = caveA;
				} else {
					col = vec3(0.80, 0.83, 0.80) * l;
					a = 0.07;
				}
				acc.rgb += (1.0 - acc.a) * a * col;
				acc.a += (1.0 - acc.a) * a;
			}
			t = t1; p = p1; d = d1;
		}
		outColor = vec4(acc.rgb + (1.0 - acc.a) * bg, 1.0);
	}`;

/**
 * Draws the volume into a canvas and keeps the tiles around the camera.
 * It asks for tiles through `want`, nearest first; the host samples them
 * and hands them back with `add`.
 */
export class VolumeRenderer {
	readonly renderer: THREE.WebGLRenderer;
	readonly camera: THREE.PerspectiveCamera;
	readonly controls: OrbitControls;
	private scene = new THREE.Scene();
	private material: THREE.ShaderMaterial;
	private box: THREE.Mesh;
	private edges: THREE.LineSegments;
	private uniforms: Record<string, THREE.IUniform>;
	private grid: Grid | null = null;
	private tiles = new Map<string, Tile>();
	/** Which tile each slot holds, by slot index (tz mod W)·W + (tx mod W). */
	private slots: (Tile | null)[] = new Array(W * W).fill(null);
	private volData: Uint16Array<ArrayBuffer> | null = null;
	private lipData: Uint16Array<ArrayBuffer> | null = null;
	private topData = new Float32Array(W * VOLUME_TILE * W * VOLUME_TILE).fill(-1e9);
	private readyData = new Uint8Array(W * W);
	private volTex: THREE.Data3DTexture | null = null;
	private lipTex: THREE.Data3DTexture | null = null;
	private topTex: THREE.DataTexture;
	private readyTex: THREE.DataTexture;
	/** World x and z of the scene's origin (whole blocks), so scene coordinates stay small. */
	private origin = { x: 0, z: 0 };
	/** The tile the camera looks at. */
	private centre = { tx: Number.NaN, tz: Number.NaN };
	private settings: VolumeSettings = { ...defaultSettings };
	/** Whether the camera has been moved by hand; until then a resize refits it. */
	private touched = false;
	private full: number;
	private queued = false;
	private sharpen: ReturnType<typeof setTimeout> | undefined;
	private dirty = { vol: false, top: false };
	/**
	 * Called when the camera moves: the tiles the window needs, missing
	 * ones first (coarse, so the view fills fast), then finer ones where
	 * the camera is close, nearest first.
	 */
	want: (tiles: TileWant[]) => void = () => {};
	private lastWant = '';
	/** Called when the camera's target moves, with its world x and z. */
	moved: (x: number, z: number) => void = () => {};
	/** Called when a tile arrives, with the range of density in the window. */
	ranged: (r: ValueRange) => void = () => {};

	constructor(readonly canvas: HTMLCanvasElement) {
		this.renderer = new THREE.WebGLRenderer({ canvas, antialias: false });
		if (!this.renderer.capabilities.isWebGL2) throw new Error('the 3D density view needs WebGL2');
		this.full = Math.min(globalThis.devicePixelRatio || 1, 1.5);
		this.renderer.setPixelRatio(this.full);
		this.camera = new THREE.PerspectiveCamera(45, 1, 1, 8000);
		this.controls = new OrbitControls(this.camera, canvas);
		// Panning slides along the ground, so dragging with the right button travels.
		this.controls.screenSpacePanning = false;
		this.controls.maxPolarAngle = Math.PI * 0.95;
		this.controls.addEventListener('change', () => this.moving());
		this.controls.addEventListener('start', () => (this.touched = true));

		this.topTex = new THREE.DataTexture(this.topData, W * VOLUME_TILE, W * VOLUME_TILE, THREE.RedFormat, THREE.FloatType);
		this.readyTex = new THREE.DataTexture(this.readyData, W, W, THREE.RedFormat, THREE.UnsignedByteType);
		for (const t of [this.topTex, this.readyTex]) {
			t.minFilter = t.magFilter = THREE.NearestFilter;
			t.wrapS = t.wrapT = THREE.RepeatWrapping;
			t.unpackAlignment = 1;
			t.needsUpdate = true;
		}
		this.uniforms = {
			vol: { value: null },
			lipTex: { value: null },
			topTex: { value: this.topTex },
			readyTex: { value: this.readyTex },
			camPos: { value: new THREE.Vector3() },
			boxMin: { value: new THREE.Vector3() },
			boxMax: { value: new THREE.Vector3() },
			bg: { value: new THREE.Color() },
			wrap: { value: new THREE.Vector2() },
			sx: { value: 4 },
			sy: { value: 8 },
			minY: { value: -64 },
			ny: { value: 2 },
			cornersAcross: { value: 1 },
			thr: { value: 0 },
			haze: { value: defaultSettings.haze },
			caveA: { value: defaultSettings.caveOpacity },
			look: { value: 0 },
			cubes: { value: true }
		};
		this.material = new THREE.ShaderMaterial({
			glslVersion: THREE.GLSL3,
			uniforms: this.uniforms,
			side: THREE.BackSide,
			vertexShader,
			fragmentShader
		});
		this.box = new THREE.Mesh(new THREE.BoxGeometry(1, 1, 1), this.material);
		this.box.visible = false;
		this.scene.add(this.box);
		this.edges = new THREE.LineSegments(new THREE.EdgesGeometry(new THREE.BoxGeometry(1, 1, 1)), new THREE.LineBasicMaterial({ transparent: true, opacity: 0.35 }));
		this.scene.add(this.edges);
	}

	/** Sets the background and edge colours from CSS colours. */
	setColors(background: string, edge: string) {
		THREE.ColorManagement.enabled = false; // the shader writes display colours directly
		const bg = new THREE.Color(background);
		(this.uniforms.bg.value as THREE.Color).copy(bg);
		this.renderer.setClearColor(bg);
		(this.edges.material as THREE.LineBasicMaterial).color.set(edge);
		this.redraw();
	}

	resize(w: number, h: number) {
		this.renderer.setSize(w, h, false);
		this.camera.aspect = w / Math.max(1, h);
		this.camera.updateProjectionMatrix();
		if (!this.touched && w > 0 && h > 0) this.fit();
		this.redraw();
	}

	/** Puts the camera over world (x, z), looking down at an angle from the south-east, far enough back to see the window. */
	lookAt(x: number, z: number) {
		this.origin = { x: Math.round(x), z: Math.round(z) };
		const g = this.grid;
		const minY = g?.minY ?? -64;
		const maxY = g ? g.minY + (g.ny - 1) * g.sy : 320;
		this.controls.target.set(x - this.origin.x, (minY + maxY) / 2, z - this.origin.z);
		this.fit();
		this.followTarget();
	}

	/** Backs the camera off along its current direction until the window fits. */
	private fit() {
		const g = this.grid;
		const height = g ? (g.ny - 1) * g.sy : 384;
		const width = (2 * R + 1) * VOLUME_TILE * Math.SQRT2; // seen corner on
		const tan = Math.tan(THREE.MathUtils.degToRad(this.camera.fov / 2));
		const dist = 1.1 * Math.max(height / (2 * tan), width / (2 * tan * this.camera.aspect));
		const dir = this.camera.position.clone().sub(this.controls.target);
		if (dir.lengthSq() < 1e-6) dir.set(1.35, 0.95, 1.6);
		this.camera.position.copy(this.controls.target).addScaledVector(dir.normalize(), dist);
		this.controls.update();
	}

	/** Moves the camera and its target along the ground, by blocks to the screen's right and forward. */
	travel(right: number, forward: number) {
		const f = new THREE.Vector3();
		this.camera.getWorldDirection(f);
		f.y = 0;
		if (f.lengthSq() < 1e-6) f.set(0, 0, -1);
		f.normalize();
		const r = new THREE.Vector3(-f.z, 0, f.x);
		const step = r.multiplyScalar(right).add(f.multiplyScalar(forward));
		this.touched = true;
		this.camera.position.add(step);
		this.controls.target.add(step);
		this.controls.update();
	}

	/** Forgets every tile: the function, its seed or the pack changed. */
	clear() {
		this.tiles.clear();
		this.slots.fill(null);
		this.grid = null;
		this.volData = null;
		this.lipData = null;
		this.volTex?.dispose();
		this.lipTex?.dispose();
		this.volTex = this.lipTex = null;
		this.box.visible = false;
		this.topData.fill(-1e9);
		this.readyData.fill(0);
		this.readyTex.needsUpdate = this.topTex.needsUpdate = true;
		this.centre = { tx: Number.NaN, tz: Number.NaN };
		this.lastWant = '';
		this.followTarget();
	}

	apply(s: VolumeSettings) {
		const thresholdChanged = s.threshold !== this.settings.threshold;
		this.settings = { ...s };
		const u = this.uniforms;
		u.look.value = lookIndex[s.look];
		u.thr.value = s.threshold;
		u.haze.value = s.haze;
		u.caveA.value = s.caveOpacity;
		u.cubes.value = s.blocks;
		if (thresholdChanged) this.refreshTops();
		this.placeBox();
		this.moving();
	}

	/** Takes a sampled tile. */
	add(tile: Extract<VolumeTile, { values: Float32Array }>) {
		const g = this.grid;
		if (!g || g.sx !== tile.sx || g.sy !== tile.sy || g.ny !== tile.ny || g.minY !== tile.minY) this.setGrid(tile);
		const grid = this.grid!;
		const key = `${tile.tx},${tile.tz}`;
		const had = this.tiles.get(key);
		if (had && had.lod <= tile.lod) return; // a finer one came first
		const values = upsample(grid, tile);
		const t: Tile = { tx: tile.tx, tz: tile.tz, lod: tile.lod, values, lip: lipOf(grid, values), tops: null, topsFor: NaN };
		this.tiles.delete(key);
		this.tiles.set(key, t);
		if (this.tiles.size > MAX_CACHED) {
			for (const [k, old] of this.tiles) {
				if (this.inWindow(old.tx, old.tz)) continue;
				this.tiles.delete(k);
				if (this.tiles.size <= MAX_CACHED) break;
			}
		}
		if (this.inWindow(t.tx, t.tz)) this.place(t);
		this.replan();
		const r = this.valueRange();
		if (r) this.ranged(r);
		this.redraw();
	}

	/** The smallest, largest and median density in the window's tiles (the median from a sample of each). */
	valueRange(): ValueRange | null {
		let min = Infinity;
		let max = -Infinity;
		const some: number[] = [];
		for (let dz = -R; dz <= R; dz++) {
			for (let dx = -R; dx <= R; dx++) {
				const t = this.tiles.get(`${this.centre.tx + dx},${this.centre.tz + dz}`);
				if (!t) continue;
				const v = t.values;
				for (let i = 0; i < v.length; i++) {
					if (v[i] < min) min = v[i];
					if (v[i] > max) max = v[i];
				}
				const stride = Math.max(1, Math.floor(v.length / 997));
				for (let i = 0; i < v.length; i += stride) some.push(v[i]);
			}
		}
		if (!some.length) return null;
		some.sort((a, b) => a - b);
		return { min, max, median: some[some.length >> 1] };
	}

	/** How many of the window's tiles are missing, and how many are coarser than the zoom calls for. */
	missing(): { missing: number; coarse: number } {
		let missing = 0;
		let coarse = 0;
		for (let dz = -R; dz <= R; dz++) {
			for (let dx = -R; dx <= R; dx++) {
				const tx = this.centre.tx + dx;
				const tz = this.centre.tz + dz;
				const t = this.tiles.get(`${tx},${tz}`);
				if (!t) missing++;
				else if (t.lod > this.lodFor(tx, tz)) coarse++;
			}
		}
		return { missing, coarse };
	}

	/** The coarsest level of detail: 16 blocks apart. */
	private maxLod(): number {
		return Math.max(1, Math.floor(16 / (this.grid?.sx ?? 4)));
	}

	/**
	 * The level of detail a tile needs: samples about PIXELS_PER_SAMPLE
	 * screen pixels apart at its distance, in whole powers of two of the
	 * function's spacing (so always whole blocks).
	 */
	private lodFor(tx: number, tz: number): number {
		const g = this.grid;
		if (!g) return this.maxLod();
		const h = this.canvas.clientHeight || 1;
		const centre = new THREE.Vector3(
			(tx + 0.5) * VOLUME_TILE - this.origin.x,
			this.controls.target.y,
			(tz + 0.5) * VOLUME_TILE - this.origin.z
		);
		// The tile's nearest point is about half a tile closer than its centre.
		const dist = Math.max(1, this.camera.position.distanceTo(centre) - VOLUME_TILE / 2);
		const blocksPerPixel = (2 * dist * Math.tan(THREE.MathUtils.degToRad(this.camera.fov / 2))) / h;
		const spacing = blocksPerPixel * PIXELS_PER_SAMPLE;
		let lod = 1;
		while (lod < this.maxLod() && g.sx * lod * 2 <= spacing) lod *= 2;
		return lod;
	}

	/** Asks for what the window needs now: missing tiles coarse, then refinements. */
	private replan() {
		if (!Number.isFinite(this.centre.tx)) return;
		const { tx, tz } = this.centre;
		const near = (a: TileWant, b: TileWant) => Math.hypot(a.tx - tx, a.tz - tz) - Math.hypot(b.tx - tx, b.tz - tz);
		const missing: TileWant[] = [];
		const finer: TileWant[] = [];
		const coarsest = this.maxLod();
		for (let dz = -R; dz <= R; dz++) {
			for (let dx = -R; dx <= R; dx++) {
				const at = { tx: tx + dx, tz: tz + dz };
				const t = this.tiles.get(`${at.tx},${at.tz}`);
				const lod = this.lodFor(at.tx, at.tz);
				if (!t) {
					missing.push({ ...at, lod: coarsest });
					if (lod < coarsest) finer.push({ ...at, lod });
				} else if (t.lod > lod) {
					finer.push({ ...at, lod });
				}
			}
		}
		missing.sort(near);
		finer.sort(near);
		const wants = [...missing, ...finer];
		const key = wants.map((w) => `${w.tx},${w.tz},${w.lod}`).join(' ');
		if (key === this.lastWant) return;
		this.lastWant = key;
		this.want(wants);
	}

	/** The world x and z the camera looks at. */
	target(): { x: number; z: number } {
		return { x: this.controls.target.x + this.origin.x, z: this.controls.target.z + this.origin.z };
	}

	dispose() {
		clearTimeout(this.sharpen);
		this.controls.dispose();
		this.volTex?.dispose();
		this.lipTex?.dispose();
		this.topTex.dispose();
		this.readyTex.dispose();
		this.material.dispose();
		this.renderer.dispose();
	}

	// ---- internals ----

	private setGrid(tile: Extract<VolumeTile, { values: Float32Array }>) {
		const c = VOLUME_TILE / tile.sx;
		this.grid = { sx: tile.sx, sy: tile.sy, nx: c + 1, ny: tile.ny, minY: tile.minY, c };
		this.tiles.clear();
		this.slots.fill(null);
		const across = W * c;
		this.volData = new Uint16Array(across * across * tile.ny);
		this.lipData = new Uint16Array(across * across * (tile.ny - 1));
		this.volTex?.dispose();
		this.lipTex?.dispose();
		// The texture's u, v, w are world x, z, y.
		this.volTex = new THREE.Data3DTexture(this.volData, across, across, tile.ny);
		this.volTex.minFilter = this.volTex.magFilter = THREE.LinearFilter;
		this.lipTex = new THREE.Data3DTexture(this.lipData, across, across, tile.ny - 1);
		this.lipTex.minFilter = this.lipTex.magFilter = THREE.NearestFilter;
		for (const t of [this.volTex, this.lipTex]) {
			t.format = THREE.RedFormat;
			t.type = THREE.HalfFloatType;
			t.wrapS = t.wrapT = THREE.RepeatWrapping;
			t.wrapR = THREE.ClampToEdgeWrapping;
			t.unpackAlignment = 1;
			t.needsUpdate = true;
		}
		const u = this.uniforms;
		u.vol.value = this.volTex;
		u.lipTex.value = this.lipTex;
		u.sx.value = tile.sx;
		u.sy.value = tile.sy;
		u.minY.value = tile.minY;
		u.ny.value = tile.ny;
		u.cornersAcross.value = across;
		this.box.visible = true;
		if (!this.touched) {
			// Centre on the world's height now that it's known.
			this.controls.target.y = tile.minY + ((tile.ny - 1) * tile.sy) / 2;
			this.fit();
		}
		this.placeBox();
	}

	private inWindow(tx: number, tz: number): boolean {
		return Math.abs(tx - this.centre.tx) <= R && Math.abs(tz - this.centre.tz) <= R;
	}

	/** Copies a tile into its slot. */
	private place(t: Tile) {
		const g = this.grid!;
		const { c, nx, ny } = g;
		const across = W * c;
		const sxSlot = mod(t.tx, W);
		const szSlot = mod(t.tz, W);
		const slot = szSlot * W + sxSlot;
		this.slots[slot] = t;
		const vol = this.volData!;
		const lip = this.lipData!;
		const half = THREE.DataUtils.toHalfFloat;
		for (let j = 0; j < ny; j++) {
			for (let k = 0; k < c; k++) {
				const row = (j * across + szSlot * c + k) * across + sxSlot * c;
				const src = (j * nx + k) * nx;
				for (let i = 0; i < c; i++) vol[row + i] = half(t.values[src + i]);
				if (j < ny - 1) {
					const lsrc = (j * c + k) * c;
					for (let i = 0; i < c; i++) lip[row + i] = half(t.lip[lsrc + i]);
				}
			}
		}
		this.placeTops(t);
		this.updateReady();
		this.dirty.vol = true;
	}

	private placeTops(t: Tile) {
		const g = this.grid!;
		if (t.topsFor !== this.settings.threshold) {
			t.tops = topsOf(g, t.values, this.settings.threshold);
			t.topsFor = this.settings.threshold;
		}
		const n = VOLUME_TILE;
		const across = W * n;
		const x0 = mod(t.tx, W) * n;
		const z0 = mod(t.tz, W) * n;
		for (let bz = 0; bz < n; bz++) this.topData.set(t.tops!.subarray(bz * n, bz * n + n), (z0 + bz) * across + x0);
		this.dirty.top = true;
	}

	private refreshTops() {
		if (!this.grid) return;
		for (const t of this.slots) if (t && this.inWindow(t.tx, t.tz)) this.placeTops(t);
	}

	private updateReady() {
		for (let s = 0; s < W * W; s++) {
			const t = this.slots[s];
			this.readyData[s] = t && this.inWindow(t.tx, t.tz) ? 255 : 0;
		}
		this.readyTex.needsUpdate = true;
	}

	/** After the target moves: rebases the scene if it went far, and moves the window. */
	private followTarget() {
		const target = this.controls.target;
		if (Math.abs(target.x) > 4096 || Math.abs(target.z) > 4096) {
			const dx = Math.round(target.x);
			const dz = Math.round(target.z);
			this.origin.x += dx;
			this.origin.z += dz;
			target.x -= dx;
			target.z -= dz;
			this.camera.position.x -= dx;
			this.camera.position.z -= dz;
			this.controls.update();
		}
		const { x, z } = this.target();
		this.moved(x, z);
		const tx = Math.floor(x / VOLUME_TILE);
		const tz = Math.floor(z / VOLUME_TILE);
		if (tx !== this.centre.tx || tz !== this.centre.tz) {
			this.centre = { tx, tz };
			for (let dz = -R; dz <= R; dz++) {
				for (let dx = -R; dx <= R; dx++) {
					const t = this.tiles.get(`${tx + dx},${tz + dz}`);
					if (t && this.grid && this.slots[mod(t.tz, W) * W + mod(t.tx, W)] !== t) this.place(t);
				}
			}
			this.updateReady();
			this.placeBox();
		}
		// Zooming changes the detail tiles need, so this runs on every move.
		this.replan();
	}

	/** Fits the box to the window, less the cuts, in scene coordinates. */
	private placeBox() {
		const g = this.grid;
		if (!Number.isFinite(this.centre.tx)) return;
		const s = this.settings;
		const x0 = (this.centre.tx - R) * VOLUME_TILE - this.origin.x;
		const z0 = (this.centre.tz - R) * VOLUME_TILE - this.origin.z;
		// The last cell along x and z blends into the spare slot, so it's left out.
		const span = (2 * R + 1) * VOLUME_TILE - (g?.sx ?? 4);
		const minY = g?.minY ?? -64;
		const maxY = g ? g.minY + (g.ny - 1) * g.sy : 320;
		const top = Math.max(minY + 1, Math.min(maxY, s.cutY));
		const min = new THREE.Vector3(x0, minY, z0);
		const max = new THREE.Vector3(x0 + Math.max(1, span * s.cutX), top, z0 + Math.max(1, span * s.cutZ));
		(this.uniforms.boxMin.value as THREE.Vector3).copy(min);
		(this.uniforms.boxMax.value as THREE.Vector3).copy(max);
		const size = max.clone().sub(min);
		const mid = min.clone().add(max).multiplyScalar(0.5);
		for (const m of [this.box, this.edges]) {
			m.scale.copy(size);
			m.position.copy(mid);
		}
		(this.uniforms.wrap.value as THREE.Vector2).set(mod(this.origin.x, W * VOLUME_TILE), mod(this.origin.z, W * VOLUME_TILE));
	}

	/** While the view moves, draw at half resolution, then sharpen once it has been still a moment. */
	private moving() {
		if (this.renderer.getPixelRatio() !== this.full / 2) this.renderer.setPixelRatio(this.full / 2);
		clearTimeout(this.sharpen);
		this.sharpen = setTimeout(() => {
			this.renderer.setPixelRatio(this.full);
			this.redraw();
		}, 160);
		this.followTarget();
		this.redraw();
	}

	redraw() {
		if (this.queued) return;
		this.queued = true;
		requestAnimationFrame(() => this.draw());
	}

	private draw() {
		this.queued = false;
		if (this.dirty.vol) {
			if (this.volTex) this.volTex.needsUpdate = true;
			if (this.lipTex) this.lipTex.needsUpdate = true;
			this.dirty.vol = false;
		}
		if (this.dirty.top) {
			this.topTex.needsUpdate = true;
			this.dirty.top = false;
		}
		(this.uniforms.camPos.value as THREE.Vector3).copy(this.camera.position);
		this.renderer.render(this.scene, this.camera);
	}
}
