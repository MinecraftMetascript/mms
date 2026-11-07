// A small WebGL 2 view for the jigsaw simulation: flat-shaded triangles
// and lines in one colour per vertex, an orbiting camera, and rays for
// picking. The meshes are built on the CPU from blocks or boxes.

import type { Box } from './simulate';

export type Rgb = readonly [number, number, number];
type V3 = [number, number, number];

/** The slice: keep only what is at or below `at` along `axis` (x, y or z). */
export interface Slice {
	axis: 0 | 1 | 2;
	at: number;
}

// The six faces of a unit cube: the direction out, a shade (top lit
// brightest), and its corners as offsets.
const FACES: { n: V3; shade: number; c: V3[] }[] = [
	{ n: [0, 1, 0], shade: 1, c: [[0, 1, 0], [0, 1, 1], [1, 1, 1], [1, 1, 0]] },
	{ n: [0, -1, 0], shade: 0.5, c: [[0, 0, 0], [1, 0, 0], [1, 0, 1], [0, 0, 1]] },
	{ n: [1, 0, 0], shade: 0.8, c: [[1, 0, 0], [1, 1, 0], [1, 1, 1], [1, 0, 1]] },
	{ n: [-1, 0, 0], shade: 0.7, c: [[0, 0, 0], [0, 0, 1], [0, 1, 1], [0, 1, 0]] },
	{ n: [0, 0, 1], shade: 0.75, c: [[0, 0, 1], [1, 0, 1], [1, 1, 1], [0, 1, 1]] },
	{ n: [0, 0, -1], shade: 0.65, c: [[0, 0, 0], [0, 1, 0], [1, 1, 0], [1, 0, 0]] }
];

function pushQuad(out: number[], c: V3[], lo: V3, size: V3, rgb: Rgb, shade: number) {
	const r = (rgb[0] / 255) * shade, g = (rgb[1] / 255) * shade, b = (rgb[2] / 255) * shade;
	for (const i of [0, 1, 2, 0, 2, 3]) out.push(lo[0] + c[i][0] * size[0], lo[1] + c[i][1] * size[1], lo[2] + c[i][2] * size[2], r, g, b);
}

/**
 * Triangles for blocks given as x, y, z, colour index, piece (5 ints
 * each). A face is drawn only where no kept block is beside it. When
 * `focus` is a piece, the other pieces' blocks are drawn faded.
 */
export function blockMesh(blocks: Int32Array, colors: readonly Rgb[], slice: Slice | null, focus: number | null): Float32Array {
	const kept = new Set<string>();
	const keep = (i: number) => !slice || blocks[i + slice.axis] <= slice.at;
	for (let i = 0; i < blocks.length; i += 5) if (keep(i)) kept.add(`${blocks[i]},${blocks[i + 1]},${blocks[i + 2]}`);
	const out: number[] = [];
	for (let i = 0; i < blocks.length; i += 5) {
		if (!keep(i)) continue;
		const x = blocks[i], y = blocks[i + 1], z = blocks[i + 2];
		let rgb = colors[blocks[i + 3]];
		if (focus !== null && blocks[i + 4] !== focus) rgb = fade(rgb);
		for (const f of FACES) {
			if (kept.has(`${x + f.n[0]},${y + f.n[1]},${z + f.n[2]}`)) continue;
			// A face on the slice plane is the cut: drawn a little darker.
			const cut = slice && f.n[slice.axis] === 1 && [x, y, z][slice.axis] === slice.at;
			pushQuad(out, f.c, [x, y, z], [1, 1, 1], rgb, cut ? f.shade * 0.8 : f.shade);
		}
	}
	return Float32Array.from(out);
}

/** A box clipped to the slice, or null when the slice leaves none of it. */
export function clip(b: Box, slice: Slice | null): Box | null {
	if (!slice) return b;
	if (b[0][slice.axis] > slice.at) return null;
	const max = [...b[1]] as Box[1];
	max[slice.axis] = Math.min(max[slice.axis], slice.at);
	return [b[0], max];
}

/** Triangles for solid boxes, each a little inset so neighbours show apart. */
export function boxMesh(boxes: readonly { box: Box; rgb: Rgb }[], slice: Slice | null, inset = 0.06): Float32Array {
	const out: number[] = [];
	for (const { box, rgb } of boxes) {
		const b = clip(box, slice);
		if (!b) continue;
		const lo: V3 = [b[0][0] + inset, b[0][1] + inset, b[0][2] + inset];
		const size: V3 = [b[1][0] + 1 - b[0][0] - 2 * inset, b[1][1] + 1 - b[0][1] - 2 * inset, b[1][2] + 1 - b[0][2] - 2 * inset];
		for (const f of FACES) pushQuad(out, f.c, lo, size, rgb, f.shade);
	}
	return Float32Array.from(out);
}

/** Lines for the edges of boxes. */
export function edgeLines(boxes: readonly { box: Box; rgb: Rgb }[], slice: Slice | null, grow = 0): Float32Array {
	const out: number[] = [];
	for (const { box, rgb } of boxes) {
		const b = clip(box, slice);
		if (!b) continue;
		const lo = [b[0][0] - grow, b[0][1] - grow, b[0][2] - grow];
		const hi = [b[1][0] + 1 + grow, b[1][1] + 1 + grow, b[1][2] + 1 + grow];
		const c = [rgb[0] / 255, rgb[1] / 255, rgb[2] / 255];
		const p = (i: number): V3 => [i & 1 ? hi[0] : lo[0], i & 2 ? hi[1] : lo[1], i & 4 ? hi[2] : lo[2]];
		for (const [a, d] of [[0, 1], [2, 3], [4, 5], [6, 7], [0, 2], [1, 3], [4, 6], [5, 7], [0, 4], [1, 5], [2, 6], [3, 7]]) out.push(...p(a), ...c, ...p(d), ...c);
	}
	return Float32Array.from(out);
}

export const fade = (c: Rgb): Rgb => [Math.round(c[0] * 0.35 + 150 * 0.65), Math.round(c[1] * 0.35 + 150 * 0.65), Math.round(c[2] * 0.35 + 150 * 0.65)];

const VS = `#version 300 es
in vec3 pos;
in vec3 color;
uniform mat4 mvp;
out vec3 vColor;
void main() { vColor = color; gl_Position = mvp * vec4(pos, 1.0); }`;
const FS = `#version 300 es
precision mediump float;
in vec3 vColor;
out vec4 outColor;
void main() { outColor = vec4(vColor, 1.0); }`;

export interface Camera {
	target: V3;
	/** Turn about the y axis, in radians. */
	yaw: number;
	/** Tilt down from level, in radians. */
	pitch: number;
	distance: number;
}

const sub = (a: V3, b: V3): V3 => [a[0] - b[0], a[1] - b[1], a[2] - b[2]];
const cross = (a: V3, b: V3): V3 => [a[1] * b[2] - a[2] * b[1], a[2] * b[0] - a[0] * b[2], a[0] * b[1] - a[1] * b[0]];
const norm = (a: V3): V3 => {
	const l = Math.hypot(...a) || 1;
	return [a[0] / l, a[1] / l, a[2] / l];
};
const FOV = Math.PI / 4;

/** Where the camera is, and its forward, right and up directions. */
export function eye(c: Camera): { at: V3; f: V3; r: V3; u: V3 } {
	const at: V3 = [
		c.target[0] + c.distance * Math.cos(c.pitch) * Math.sin(c.yaw),
		c.target[1] + c.distance * Math.sin(c.pitch),
		c.target[2] + c.distance * Math.cos(c.pitch) * Math.cos(c.yaw)
	];
	const f = norm(sub(c.target, at));
	const r = norm(cross(f, [0, 1, 0]));
	const u = cross(r, f);
	return { at, f, r, u };
}

function matrix(c: Camera, aspect: number, far: number): Float32Array {
	const { at, f, r, u } = eye(c);
	const near = Math.max(0.05, c.distance / 200);
	const t = 1 / Math.tan(FOV / 2);
	// Projection times view, column-major.
	const p = [t / aspect, 0, 0, 0, 0, t, 0, 0, 0, 0, (far + near) / (near - far), -1, 0, 0, (2 * far * near) / (near - far), 0];
	const dot = (a: V3, b: V3) => a[0] * b[0] + a[1] * b[1] + a[2] * b[2];
	const v = [r[0], u[0], -f[0], 0, r[1], u[1], -f[1], 0, r[2], u[2], -f[2], 0, -dot(r, at), -dot(u, at), dot(f, at), 1];
	const out = new Float32Array(16);
	for (let col = 0; col < 4; col++) for (let row = 0; row < 4; row++) {
		let s = 0;
		for (let k = 0; k < 4; k++) s += p[k * 4 + row] * v[col * 4 + k];
		out[col * 4 + row] = s;
	}
	return out;
}

/** The ray from the camera through a point of the canvas, in CSS pixels. */
export function ray(c: Camera, x: number, y: number, width: number, height: number): { from: V3; dir: V3 } {
	const { at, f, r, u } = eye(c);
	const t = Math.tan(FOV / 2);
	const nx = ((2 * x) / width - 1) * t * (width / height);
	const ny = (1 - (2 * y) / height) * t;
	return { from: at, dir: norm([f[0] + r[0] * nx + u[0] * ny, f[1] + r[1] * nx + u[1] * ny, f[2] + r[2] * nx + u[2] * ny]) };
}

/** How far along a ray it meets a box (blocks min to max + 1), or null. */
export function hit(from: V3, dir: V3, b: Box): number | null {
	let lo = -Infinity, hi = Infinity;
	for (let i = 0; i < 3; i++) {
		const a = b[0][i], z = b[1][i] + 1;
		if (Math.abs(dir[i]) < 1e-9) {
			if (from[i] < a || from[i] > z) return null;
			continue;
		}
		let t0 = (a - from[i]) / dir[i], t1 = (z - from[i]) / dir[i];
		if (t0 > t1) [t0, t1] = [t1, t0];
		lo = Math.max(lo, t0);
		hi = Math.min(hi, t1);
		if (lo > hi) return null;
	}
	return hi < 0 ? null : Math.max(lo, 0);
}

/** Draws triangles and lines into a canvas. */
export class VoxelView {
	private readonly gl: WebGL2RenderingContext;
	private readonly program: WebGLProgram;
	private readonly mvp: WebGLUniformLocation;
	private readonly tris: WebGLBuffer;
	private readonly lines: WebGLBuffer;
	private readonly vaoTris: WebGLVertexArrayObject;
	private readonly vaoLines: WebGLVertexArrayObject;
	private triCount = 0;
	private lineCount = 0;
	far = 1000;

	constructor(canvas: HTMLCanvasElement) {
		const gl = canvas.getContext('webgl2', { antialias: true });
		if (!gl) throw new Error('WebGL 2 is not available');
		this.gl = gl;
		const shader = (type: number, src: string) => {
			const s = gl.createShader(type)!;
			gl.shaderSource(s, src);
			gl.compileShader(s);
			if (!gl.getShaderParameter(s, gl.COMPILE_STATUS)) throw new Error(gl.getShaderInfoLog(s) ?? 'shader');
			return s;
		};
		const p = gl.createProgram()!;
		gl.attachShader(p, shader(gl.VERTEX_SHADER, VS));
		gl.attachShader(p, shader(gl.FRAGMENT_SHADER, FS));
		gl.linkProgram(p);
		this.program = p;
		this.mvp = gl.getUniformLocation(p, 'mvp')!;
		const vao = (buf: WebGLBuffer) => {
			const v = gl.createVertexArray()!;
			gl.bindVertexArray(v);
			gl.bindBuffer(gl.ARRAY_BUFFER, buf);
			const pos = gl.getAttribLocation(p, 'pos'), color = gl.getAttribLocation(p, 'color');
			gl.enableVertexAttribArray(pos);
			gl.vertexAttribPointer(pos, 3, gl.FLOAT, false, 24, 0);
			gl.enableVertexAttribArray(color);
			gl.vertexAttribPointer(color, 3, gl.FLOAT, false, 24, 12);
			gl.bindVertexArray(null);
			return v;
		};
		this.tris = gl.createBuffer()!;
		this.lines = gl.createBuffer()!;
		this.vaoTris = vao(this.tris);
		this.vaoLines = vao(this.lines);
	}

	set(tris: Float32Array, lines: Float32Array) {
		const gl = this.gl;
		gl.bindBuffer(gl.ARRAY_BUFFER, this.tris);
		gl.bufferData(gl.ARRAY_BUFFER, tris, gl.STATIC_DRAW);
		gl.bindBuffer(gl.ARRAY_BUFFER, this.lines);
		gl.bufferData(gl.ARRAY_BUFFER, lines, gl.STATIC_DRAW);
		this.triCount = tris.length / 6;
		this.lineCount = lines.length / 6;
	}

	draw(camera: Camera, background: Rgb) {
		const gl = this.gl;
		const w = gl.canvas.width, h = gl.canvas.height;
		gl.viewport(0, 0, w, h);
		gl.clearColor(background[0] / 255, background[1] / 255, background[2] / 255, 1);
		gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
		gl.enable(gl.DEPTH_TEST);
		gl.useProgram(this.program);
		gl.uniformMatrix4fv(this.mvp, false, matrix(camera, w / Math.max(h, 1), this.far));
		gl.bindVertexArray(this.vaoTris);
		gl.drawArrays(gl.TRIANGLES, 0, this.triCount);
		gl.bindVertexArray(this.vaoLines);
		gl.drawArrays(gl.LINES, 0, this.lineCount);
		gl.bindVertexArray(null);
	}

	dispose() {
		const gl = this.gl;
		gl.deleteBuffer(this.tris);
		gl.deleteBuffer(this.lines);
		gl.deleteVertexArray(this.vaoTris);
		gl.deleteVertexArray(this.vaoLines);
		gl.deleteProgram(this.program);
	}
}
