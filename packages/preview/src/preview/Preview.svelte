<script lang="ts">
	import { onMount } from 'svelte';
	import type { OutputFile } from '../types';
	import {
		facings,
		modesOf,
		previewable,
		previewFile,
		screenAxes,
		topDown,
		type Facing,
		type PreviewFile,
		type PreviewKind,
		type PreviewMode,
		type PreviewRequest,
		type PreviewResponse,
		type PreviewView
	} from './protocol';
	import { VOID } from './colors';
	import { colorize, cssGradient, formatValue, paletteOf, palettes, position, rangeFor, solidPalette, type PaletteId, type Range } from './colormaps';

	interface Props {
		output: OutputFile[];
		/** The output path to preview. */
		target: string | null;
		/**
		 * Starts a preview worker (preview.worker.ts). The host bundles it:
		 * the web app as a worker file, VS Code's webview inline, since a
		 * webview can only start workers from blob URLs.
		 */
		createWorker: () => Worker;
		/** What to say when there is nothing to preview. */
		hint?: string;
	}
	let {
		output,
		target,
		createWorker,
		hint = 'Put the cursor in a noise, density function, noise settings, surface rule or dimension declaration to preview it.'
	}: Props = $props();

	// A target the pack no longer has (renamed or deleted) shows nothing.
	const info = $derived(target && output.some((f) => f.path === target) ? previewable(target) : null);
	// Each view starts with the camera at the origin facing north (x to the right).
	const defaults: Record<PreviewKind | 'biomes', PreviewView> = {
		noise: { x: 0, y: 0, z: 0, facing: 'north', scale: 4 },
		density_function: { x: 0, y: 64, z: 0, facing: 'north', scale: 2 },
		noise_settings: { x: 0, y: 64, z: 0, facing: 'north', scale: 2 },
		material_rule: { x: 0, y: 64, z: 0, facing: 'north', scale: 0.5 },
		dimension: { x: 0, y: 64, z: 0, facing: 'north', scale: 2 },
		biomes: { x: 0, y: 64, z: 0, facing: 'north', scale: 8 }
	};
	const facingLetters: Record<Facing, string> = { north: 'N', east: 'E', south: 'S', west: 'W' };
	/** The compass letter of the way a world x/z step points. */
	const compass = (dx: number, dz: number) => (dx > 0 ? 'E' : dx < 0 ? 'W' : dz > 0 ? 'S' : 'N');
	const modeNames: Record<PreviewMode, string> = { density: 'Density', blocks: 'Blocks', biomes: 'Biomes', iso: '3D' };
	const kindNames: Record<PreviewKind, string> = {
		noise: 'noise',
		density_function: 'density function',
		noise_settings: 'noise settings',
		material_rule: 'surface rule',
		dimension: 'dimension'
	};
	const defaultView = (kind: PreviewKind, mode: PreviewMode) => ({ ...defaults[mode === 'biomes' ? 'biomes' : kind] });

	let canvas: HTMLCanvasElement;
	let box: HTMLDivElement;
	let worker: Worker | null = null;
	let view = $state<PreviewView>({ ...defaults.density_function });
	// The 3D view's own camera: x and y are the world x and z at the centre,
	// scale is a block's size in pixels (0 until the first render fits it).
	let isoView = $state<PreviewView>({ x: 0, y: 0, z: 0, facing: 'north', scale: 0 });
	const isoAt = (v: PreviewView): PreviewView => ({ x: Math.round(v.x), y: Math.round(v.z), z: 0, facing: 'north', scale: 0 });
	let seed = $state('0');
	let error = $state('');
	let busy = $state(false);
	let last: Extract<PreviewResponse, { type: 'rendered' }> | null = $state(null);
	let hover = $state<{ px: number; py: number } | null>(null);
	type Point = { x: number; y: number; z: number };
	let lock = $state<Point | null>(null);
	let seq = 0;
	let lastTarget: string | null = null;
	let mode = $state<PreviewMode>('density');
	let biome = $state('minecraft:plains');
	let terrain = $state('minecraft:overworld');
	const flat = $derived(info ? topDown(info.kind, mode) : false);

	// Density colouring: a palette per kind, remembered in this browser, and
	// the range it spans (the values in view, -1 to 1, or the user's own).
	type DensityKind = 'noise' | 'density_function' | 'noise_settings';
	const paletteDefaults: Record<DensityKind, PaletteId> = { noise: 'grayscale', density_function: 'viridis', noise_settings: 'solid' };
	const paletteKey = (k: DensityKind) => `mms.preview.palette.${k}`;
	function storedPalette(k: DensityKind): PaletteId {
		try {
			const v = localStorage.getItem(paletteKey(k)) as PaletteId | null;
			if (v && (v === 'solid' ? k === 'noise_settings' : palettes.some((p) => p.id === v))) return v;
		} catch {
			// Storage can be unavailable (private windows, some webviews).
		}
		return paletteDefaults[k];
	}
	let paletteIds = $state<Record<DensityKind, PaletteId>>({
		noise: storedPalette('noise'),
		density_function: storedPalette('density_function'),
		noise_settings: storedPalette('noise_settings')
	});
	const densityKind = $derived(info && mode === 'density' && info.kind !== 'material_rule' && info.kind !== 'dimension' ? (info.kind as DensityKind) : null);
	const palette = $derived(densityKind ? paletteOf(paletteIds[densityKind]) : null);
	const paletteChoices = $derived(densityKind === 'noise_settings' ? [solidPalette, ...palettes] : palettes);
	function setPalette(id: PaletteId) {
		if (!densityKind) return;
		paletteIds[densityKind] = id;
		try {
			localStorage.setItem(paletteKey(densityKind), id);
		} catch {
			// Not remembered; still applies now.
		}
		draw();
	}
	type RangeMode = 'auto' | 'unit' | 'custom';
	let rangeMode = $state<RangeMode>('auto');
	let customRange = $state<Range>({ lo: -1, hi: 1 });
	/** The range the palette spans for the current render. */
	const valueRange = $derived.by((): Range | null => {
		if (!palette || !last || last.labels) return null;
		const fixed = rangeMode === 'unit' ? { lo: -1, hi: 1 } : rangeMode === 'custom' ? customRange : undefined;
		return rangeFor(palette, last.min, last.max, fixed);
	});
	function setRangeMode(m: RangeMode) {
		if (m === 'custom' && rangeMode !== 'custom' && last && Number.isFinite(last.min)) {
			// Start from what is shown, so switching doesn't jump.
			const r = valueRange ?? { lo: -1, hi: 1 };
			customRange = { lo: +r.lo.toPrecision(3), hi: +r.hi.toPrecision(3) };
		}
		rangeMode = m;
		draw();
	}
	function setCustom(end: 'lo' | 'hi', text: string) {
		const v = Number(text);
		if (text.trim() === '' || !Number.isFinite(v)) return;
		customRange = { ...customRange, [end]: v };
		draw();
	}
	// CSS pixels per rendered pixel; terrain is costly to sample, so coarser.
	// The isometric view is drawn with polygons, so it renders at device pixels.
	const resolution = $derived(
		mode === 'iso' ? 1 / Math.min(2, globalThis.devicePixelRatio || 1) : mode === 'density' ? (info?.kind === 'noise_settings' ? 4 : 2) : 3
	);
	// Terrain is built a chunk at a time, so zooming out far gets slow.
	const maxScale = $derived(mode === 'density' ? 64 : mode === 'biomes' ? 64 : 4);

	// Biomes a surface rule can be tested against: the pack's own, those its
	// rules name in Biome() conditions, and a few vanilla ones. With refs,
	// rules and conditions referenced by id are followed, so a rule that
	// says If(in_badlands) finds the biomes in in_badlands.
	function biomesIn(json: unknown, out: Set<string>, refs?: Map<string, unknown>, seen = new Set<string>()) {
		if (typeof json === 'string') {
			const ref = refs?.get(json);
			if (ref !== undefined && !seen.has(json)) {
				seen.add(json);
				biomesIn(ref, out, refs, seen);
			}
		} else if (Array.isArray(json)) for (const v of json) biomesIn(v, out, refs, seen);
		else if (json && typeof json === 'object') {
			for (const [k, v] of Object.entries(json)) {
				if (k === 'biome_is') for (const b of Array.isArray(v) ? v : [v]) if (typeof b === 'string' && !b.startsWith('#')) out.add(b.includes(':') ? b : `minecraft:${b}`);
				biomesIn(v, out, refs, seen);
			}
		}
	}
	const parsed = $derived.by(() => {
		const files: (PreviewFile & { text: string })[] = [];
		for (const f of output) {
			const p = previewFile(f.path);
			if (!p) continue;
			try {
				files.push({ ...p, json: JSON.parse(f.text), text: f.text });
			} catch {
				/* skip */
			}
		}
		return files;
	});
	const targetFile = $derived(info ? parsed.find((f) => f.kind === info.kind && f.id === info.id) : undefined);
	const biomeChoices = $derived.by(() => {
		const own = new Set<string>();
		const refs = new Map<string, unknown>();
		for (const f of parsed) if (f.kind === 'material_rule' || f.kind === 'material_condition') refs.set(f.id, f.json);
		if (targetFile) biomesIn(targetFile.json, own, refs);
		const all = new Set(own);
		for (const f of output) {
			const m = /^data\/([^/]+)\/worldgen\/biome\/(.+)\.json$/.exec(f.path);
			if (m) all.add(`${m[1]}:${m[2]}`);
		}
		for (const f of parsed) if (f.kind === 'material_rule' || f.kind === 'material_condition') biomesIn(f.json, all);
		for (const b of ['plains', 'desert', 'badlands', 'forest', 'snowy_plains', 'beach', 'ocean', 'swamp']) all.add(`minecraft:${b}`);
		return { first: [...own][0], all: [...all].sort() };
	});
	const terrainChoices = $derived(['minecraft:overworld', ...parsed.filter((f) => f.kind === 'noise_settings').map((f) => f.id)]);

	onMount(() => {
		worker = createWorker();
		worker.onmessage = (e: MessageEvent<PreviewResponse>) => {
			const r = e.data;
			if (r.seq !== seq) {
				if (r.type === 'rendered') r.bitmap?.close();
				return;
			}
			busy = r.type === 'rendered' && !!r.partial;
			if (r.type === 'error') {
				error = r.message;
				return;
			}
			error = '';
			last = r;
			// The worker picks the default size and keeps zooming out within what fills the canvas.
			if (r.scale && (!isoView.scale || r.scale > isoView.scale)) isoView.scale = r.scale;
			draw();
		};
		const ro = new ResizeObserver(() => request());
		ro.observe(box);
		return () => {
			ro.disconnect();
			worker?.terminate();
			for (const h of helpers) h.terminate();
		};
	});

	// Re-register files whenever the pack changes, then re-render.
	$effect(() => {
		const files = parsed.map(({ kind, id, json }) => ({ kind, id, json }));
		worker?.postMessage({ type: 'setup', files } satisfies PreviewRequest);
		request();
	});

	// A new target starts from its kind's default mode, view and options.
	$effect(() => {
		const key = info ? `${info.kind}|${info.id}` : null;
		if (info && key !== lastTarget) {
			const sameKind = lastTarget?.startsWith(`${info.kind}|`);
			lastTarget = key;
			if (!sameKind) {
				mode = modesOf[info.kind][0];
				view = defaultView(info.kind, mode);
				lock = null;
				last = null;
			}
			biome = biomeChoices.first ?? 'minecraft:plains';
			if (info.kind === 'material_rule') {
				const user = parsed.find((f) => f.kind === 'noise_settings' && f.text.includes(`"${info.id}"`));
				terrain = user?.id ?? 'minecraft:overworld';
			}
		}
		request();
	});

	function setMode(m: PreviewMode) {
		if (!info || m === mode) return;
		const wasFlat = flat;
		mode = m;
		if (topDown(info.kind, m) !== wasFlat) {
			// Keep the camera where it is, at the new view's zoom.
			view = { ...view, scale: defaultView(info.kind, m).scale };
			lock = null;
		}
		view.scale = Math.min(view.scale, maxScale);
		if (m === 'iso') {
			// Start the 3D view over the camera.
			isoView = isoAt(view);
			lock = null;
		}
		last = null;
		request();
	}

	/**
	 * Workers that build chunks for the 3D view in parallel, started the
	 * first time it is shown and talking to the main worker directly.
	 */
	let helpers: Worker[] = [];
	function startHelpers() {
		if (helpers.length > 0 || !worker) return;
		const count = Math.max(1, Math.min(6, (navigator.hardwareConcurrency || 4) - 1));
		const ports: MessagePort[] = [];
		for (let i = 0; i < count; i++) {
			const helper = createWorker();
			const channel = new MessageChannel();
			helper.postMessage({ type: 'serve', port: channel.port2 } satisfies PreviewRequest, [channel.port2]);
			ports.push(channel.port1);
			helpers.push(helper);
		}
		worker.postMessage({ type: 'helpers', ports } satisfies PreviewRequest, ports);
	}

	let pending = false;
	function request() {
		if (pending) return;
		pending = true;
		requestAnimationFrame(() => {
			pending = false;
			if (!worker || !info || !box) return;
			const width = Math.max(16, Math.floor(box.clientWidth / resolution));
			const height = Math.max(16, Math.floor(box.clientHeight / resolution));
			busy = true;
			if (mode === 'iso') startHelpers();
			worker.postMessage({
				type: 'render',
				seq: ++seq,
				kind: info.kind,
				mode,
				id: info.id,
				width,
				height,
				view: $state.snapshot(mode === 'iso' ? isoView : view),
				seed,
				options: { biome, terrain }
			} satisfies PreviewRequest);
		});
	}

	/** The worker-drawn image on screen, closed once it is replaced. */
	let shown: ImageBitmap | null = null;
	function draw() {
		if (!last || !canvas) return;
		canvas.width = last.width;
		canvas.height = last.height;
		const ctx = canvas.getContext('2d')!;
		if (last.bitmap) {
			ctx.drawImage(last.bitmap, 0, 0);
			if (shown && shown !== last.bitmap) shown.close();
			shown = last.bitmap;
		} else if (palette && palette.id !== 'solid' && valueRange && mode === 'density') {
			ctx.putImageData(new ImageData(colorize(last.values, palette, valueRange, VOID), last.width, last.height), 0, 0);
		} else {
			ctx.putImageData(new ImageData(last.pixels, last.width, last.height), 0, 0);
		}
		if (last.seaLevel !== undefined && mode === 'density') {
			const py = (view.y - last.seaLevel) / view.scale + last.height / 2;
			ctx.strokeStyle = 'rgba(30, 80, 160, 0.6)';
			ctx.setLineDash([4, 3]);
			ctx.beginPath();
			ctx.moveTo(0, py);
			ctx.lineTo(last.width, py);
			ctx.stroke();
			ctx.setLineDash([]);
		}
		if (lock && inView(lock)) {
			const [px, py] = toPixel(lock);
			ctx.strokeStyle = '#ef4444';
			ctx.lineWidth = 1.5;
			ctx.beginPath();
			ctx.arc(px, py, 4, 0, Math.PI * 2);
			ctx.stroke();
		}
	}

	/**
	 * The interaction code works in the screen's plane: u along the screen's
	 * right and v down it (top-down) or up it (side views, where v is y).
	 * These map the camera and world points into it and back.
	 */
	function planeOf(c: { x: number; y: number; z: number }, facing = view.facing): [number, number] {
		const a = screenAxes(facing);
		const u = c.x * a.rx + c.z * a.rz;
		return [u, flat ? c.x * a.dx + c.z * a.dz : c.y];
	}
	/** Moves the camera to plane point (u, v), keeping the slice's depth (or a top-down view's y). */
	function setPlane(u: number, v: number) {
		const a = screenAxes(view.facing);
		if (flat) {
			view.x = u * a.rx + v * a.dx;
			view.z = u * a.rz + v * a.dz;
		} else {
			if (a.rx) view.x = u * a.rx;
			else view.z = u * a.rz;
			view.y = v;
		}
	}
	/** The world block under a canvas point. */
	function toWorld(px: number, py: number): Point {
		const [u, v] = flatWorld(px, py, view.scale, ...planeOf(view));
		const a = screenAxes(view.facing);
		if (flat) return { x: Math.floor(u * a.rx + v * a.dx), y: Math.floor(view.y), z: Math.floor(u * a.rz + v * a.dz) };
		return a.rx
			? { x: Math.floor(u * a.rx), y: Math.floor(v), z: Math.floor(view.z) }
			: { x: Math.floor(view.x), y: Math.floor(v), z: Math.floor(u * a.rz) };
	}
	/** The canvas point at the middle of a world block. */
	function toPixel(p: Point): [number, number] {
		const w = last?.width ?? 0;
		const h = last?.height ?? 0;
		const [u, v] = planeOf({ x: p.x + 0.5, y: p.y + 0.5, z: p.z + 0.5 });
		const [cu, cv] = planeOf(view);
		return [(u - cu) / view.scale + w / 2, flat ? (v - cv) / view.scale + h / 2 : (cv - v) / view.scale + h / 2];
	}
	/** Whether a world block is in the slice or map shown (a side view shows one depth, a map one y). */
	function inView(p: Point): boolean {
		const a = screenAxes(view.facing);
		if (flat) return true;
		return a.rx ? p.z === Math.floor(view.z) : p.x === Math.floor(view.x);
	}
	function coords(p: Point): string {
		return flat ? `x ${p.x}, z ${p.z}` : `x ${p.x}, y ${p.y}, z ${p.z}`;
	}
	const axes = $derived(screenAxes(view.facing));
	function coordTitle(axis: 'x' | 'y' | 'z'): string {
		if (axis === 'y') return flat ? 'The height the map is sampled at' : 'The height at the middle of the view';
		if (flat) return `The ${axis} at the middle of the map`;
		const depth = axes.rx ? 'z' : 'x';
		return axis === depth ? `The slice is at this ${axis} (turn to slice along the other axis)` : `The ${axis} at the middle of the slice`;
	}
	/** Turns the camera a quarter turn (1 clockwise, -1 anticlockwise) or to a facing, about its own position. */
	function face(to: Facing | 1 | -1) {
		const next = typeof to === 'string' ? to : facings[(facings.indexOf(view.facing) + to + 4) % 4];
		if (next === view.facing) return;
		view.facing = next;
		request();
		draw();
	}
	/** Sets one camera coordinate from an input, ignoring what isn't a number yet. */
	function setCoord(axis: 'x' | 'y' | 'z', value: string) {
		const n = Number(value);
		if (value.trim() === '' || !Number.isFinite(n)) return;
		view[axis] = n;
		request();
	}

	let drag: { x: number; y: number; vx: number; vy: number; moved: boolean } | null = null;
	/** Pointers down on the canvas (fingers, or the mouse), in client coordinates. */
	const pointers = new Map<number, { x: number; y: number }>();
	/** A two-finger pinch: where it started, and the view then. */
	let pinch: { dist: number; mid: { px: number; py: number }; vx: number; vy: number; scale: number } | null = null;
	/** Whether this gesture pinched, so lifting a finger after it doesn't count as a tap. */
	let pinched = false;

	function pointer(e: { clientX: number; clientY: number }) {
		const r = canvas.getBoundingClientRect();
		return { px: ((e.clientX - r.left) / r.width) * (last?.width ?? 1), py: ((e.clientY - r.top) / r.height) * (last?.height ?? 1) };
	}
	/** The current block size of the 3D view (before its first render, what the worker picked). */
	function isoScale(): number {
		return isoView.scale || last?.scale || 8;
	}
	/** The world x/z under a canvas point of the 3D view, at block size s (screen x = (x - z)·s, y = (x + z)·s/2). */
	function isoWorld(px: number, py: number, s: number, vx: number, vz: number): [number, number] {
		const dx = px - (last?.width ?? 0) / 2;
		const dy = py - (last?.height ?? 0) / 2;
		return [vx + (dx / s + (2 * dy) / s) / 2, vz + ((2 * dy) / s - dx / s) / 2];
	}
	/** The world point under a canvas point of a 2D view. */
	function flatWorld(px: number, py: number, scale: number, vx: number, vy: number): [number, number] {
		const dx = px - (last?.width ?? 0) / 2;
		const dy = py - (last?.height ?? 0) / 2;
		return [vx + dx * scale, flat ? vy + dy * scale : vy - dy * scale];
	}
	/**
	 * Zooms to a new size, keeping the world point that was under `from`
	 * (at the starting view) under `to`. A pinch moves and zooms at once;
	 * the wheel zooms about the cursor.
	 */
	function zoomTo(from: { px: number; py: number }, to: { px: number; py: number }, start: { vx: number; vy: number; scale: number }, factor: number) {
		if (mode === 'iso') {
			const s = Math.min(48, Math.max(last?.minScale ?? 2, start.scale * factor));
			const [wx, wz] = isoWorld(from.px, from.py, start.scale, start.vx, start.vy);
			const [ox, oz] = isoWorld(to.px, to.py, s, 0, 0);
			isoView = { ...isoView, x: wx - ox, y: wz - oz, scale: s };
		} else {
			const scale = Math.min(maxScale, Math.max(1 / 16, start.scale / factor));
			const [wx, wy] = flatWorld(from.px, from.py, start.scale, start.vx, start.vy);
			const [ox, oy] = flatWorld(to.px, to.py, scale, 0, 0);
			view.scale = scale;
			setPlane(wx - ox, wy - oy);
		}
		request();
	}
	function startDrag(at: { x: number; y: number }) {
		const [vx, vy] = mode === 'iso' ? [isoView.x, isoView.y] : planeOf(view);
		drag = { x: at.x, y: at.y, vx, vy, moved: pinched };
	}
	function startPinch() {
		const [a, b] = [...pointers.values()];
		const pa = pointer({ clientX: a.x, clientY: a.y });
		const pb = pointer({ clientX: b.x, clientY: b.y });
		const [vx, vy] = mode === 'iso' ? [isoView.x, isoView.y] : planeOf(view);
		pinch = {
			dist: Math.hypot(pa.px - pb.px, pa.py - pb.py) || 1,
			mid: { px: (pa.px + pb.px) / 2, py: (pa.py + pb.py) / 2 },
			vx,
			vy,
			scale: mode === 'iso' ? isoScale() : view.scale
		};
		pinched = true;
		drag = null;
	}

	function down(e: PointerEvent) {
		canvas.setPointerCapture(e.pointerId);
		if (pointers.size === 0) pinched = false;
		pointers.set(e.pointerId, { x: e.clientX, y: e.clientY });
		if (pointers.size === 2) startPinch();
		else if (pointers.size === 1) startDrag({ x: e.clientX, y: e.clientY });
	}
	function move(e: PointerEvent) {
		hover = pointer(e);
		if (pointers.has(e.pointerId)) pointers.set(e.pointerId, { x: e.clientX, y: e.clientY });
		if (pinch && pointers.size >= 2) {
			const [a, b] = [...pointers.values()];
			const pa = pointer({ clientX: a.x, clientY: a.y });
			const pb = pointer({ clientX: b.x, clientY: b.y });
			const dist = Math.hypot(pa.px - pb.px, pa.py - pb.py) || 1;
			zoomTo(pinch.mid, { px: (pa.px + pb.px) / 2, py: (pa.py + pb.py) / 2 }, pinch, dist / pinch.dist);
			return;
		}
		if (!drag) return;
		const dx = (e.clientX - drag.x) / resolution;
		const dy = (e.clientY - drag.y) / resolution;
		if (Math.abs(dx) + Math.abs(dy) > 2) drag.moved = true;
		if (!drag.moved) return;
		if (mode === 'iso') {
			// Undo the projection (screen x = (x - z)·s, screen y = (x + z)·s/2),
			// so the ground follows the pointer.
			const s = isoScale();
			isoView.x = drag.vx - (dx / s + (2 * dy) / s) / 2;
			isoView.y = drag.vy - ((2 * dy) / s - dx / s) / 2;
			request();
			return;
		}
		setPlane(drag.vx - dx * view.scale, flat ? drag.vy - dy * view.scale : drag.vy + dy * view.scale);
		request();
	}
	function up(e: PointerEvent) {
		if (!pointers.has(e.pointerId)) return;
		pointers.delete(e.pointerId);
		if (pinch) {
			if (pointers.size >= 2) {
				// A third touch (a palm, say) lifted: keep pinching with the two left.
				startPinch();
				return;
			}
			// Back to one finger: carry on panning from where it is now.
			pinch = null;
			const rest = [...pointers.values()][0];
			if (rest) startDrag(rest);
			return;
		}
		if (e.type === 'pointerup' && drag && !drag.moved && !pinched && mode !== 'iso') {
			const p = pointer(e);
			const at = toWorld(p.px, p.py);
			lock = lock && lock.x === at.x && lock.y === at.y && lock.z === at.z ? null : at;
			draw();
		}
		drag = null;
	}
	function wheel(e: WheelEvent) {
		e.preventDefault();
		const p = pointer(e);
		// Trackpad pinches arrive as ctrl+wheel with small deltas; scale them up to match.
		const f = Math.exp(-e.deltaY * (e.ctrlKey ? 0.01 : 0.0015));
		const [vx, vy] = mode === 'iso' ? [isoView.x, isoView.y] : planeOf(view);
		zoomTo(p, p, { vx, vy, scale: mode === 'iso' ? isoScale() : view.scale }, f);
	}

	/** A sampled value, or the block or biome it indexes. */
	function show(v: number): string {
		const labels = last?.labels;
		return labels ? labels[v].replace(/^minecraft:/, '') : v.toFixed(4);
	}

	const readout = $derived.by(() => {
		const at = hover;
		if (!at || !last || mode === 'iso') return '';
		const ix = Math.floor(at.px);
		const iy = Math.floor(at.py);
		if (ix < 0 || iy < 0 || ix >= last.width || iy >= last.height) return '';
		const v = last.values[iy * last.width + ix];
		const coord = coords(toWorld(at.px, at.py));
		return Number.isNaN(v) ? coord : `${coord} · ${show(v)}`;
	});
	/** The value under the pointer, for the scale's marker. */
	const hoverValue = $derived.by(() => {
		const at = hover;
		if (!at || !last || mode !== 'density') return NaN;
		const ix = Math.floor(at.px);
		const iy = Math.floor(at.py);
		if (ix < 0 || iy < 0 || ix >= last.width || iy >= last.height) return NaN;
		return last.values[iy * last.width + ix];
	});
	const lockValue = $derived.by(() => {
		if (!lock || !last) return '';
		const [px, py] = toPixel(lock);
		const ix = Math.floor(px);
		const iy = Math.floor(py);
		if (!inView(lock) || ix < 0 || iy < 0 || ix >= last.width || iy >= last.height) return coords(lock);
		const v = last.values[iy * last.width + ix];
		return `${coords(lock)} · ${Number.isNaN(v) ? '–' : show(v)}`;
	});
</script>

<div class="flex h-full min-h-0 flex-col">
	<div class="flex flex-wrap items-center gap-2 border-b border-line px-3 py-1.5 text-xs text-muted">
		{#if info}
			<span class="font-medium text-ink">{info.id}</span>
			<span>{kindNames[info.kind]}</span>
			{#if modesOf[info.kind].length > 1}
				<div class="flex overflow-hidden rounded border border-line" role="group" aria-label="Preview mode">
					{#each modesOf[info.kind] as m (m)}
						<button
							class="px-1.5 py-0.5 {mode === m ? 'bg-accent text-white' : 'hover:bg-hover'}"
							aria-pressed={mode === m}
							onclick={() => setMode(m)}>{modeNames[m]}</button
						>
					{/each}
				</div>
			{/if}
			{#if info.kind === 'material_rule'}
				<label class="flex items-center gap-1">
					Terrain
					<select aria-label="Terrain" class="rounded border border-line bg-panel px-1 py-0.5 text-ink" bind:value={terrain} onchange={() => request()}>
						{#each terrainChoices as t (t)}<option value={t}>{t.replace(/^minecraft:/, '')}</option>{/each}
					</select>
				</label>
			{/if}
			{#if (mode === 'blocks' || mode === 'iso') && info.kind !== 'dimension'}
				<label class="flex items-center gap-1" title="The biome Biome() conditions in surface rules see">
					Biome
					<select aria-label="Biome" class="rounded border border-line bg-panel px-1 py-0.5 text-ink" bind:value={biome} onchange={() => request()}>
						{#each biomeChoices.all as b (b)}<option value={b}>{b.replace(/^minecraft:/, '')}</option>{/each}
					</select>
				</label>
			{/if}
			{#if palette}
				<label class="flex items-center gap-1">
					Palette
					<select
						aria-label="Palette"
						class="rounded border border-line bg-panel px-1 py-0.5 text-ink"
						value={palette.id}
						onchange={(e) => setPalette(e.currentTarget.value as PaletteId)}
					>
						{#each paletteChoices as p (p.id)}<option value={p.id}>{p.name}</option>{/each}
					</select>
				</label>
				{#if palette.id !== 'solid'}
					<label class="flex items-center gap-1" title="The values the palette spans: those in view, -1 to 1, or your own">
						Range
						<select
							aria-label="Range"
							class="rounded border border-line bg-panel px-1 py-0.5 text-ink"
							value={rangeMode}
							onchange={(e) => setRangeMode(e.currentTarget.value as RangeMode)}
						>
							<option value="auto">Auto (in view)</option>
							<option value="unit">−1 to 1</option>
							<option value="custom">Custom</option>
						</select>
					</label>
					{#if rangeMode === 'custom'}
						<input
							type="number"
							step="any"
							aria-label="Range from"
							class="w-16 rounded border border-line bg-panel px-1.5 py-0.5 text-ink"
							value={customRange.lo}
							oninput={(e) => setCustom('lo', e.currentTarget.value)}
						/>
						<input
							type="number"
							step="any"
							aria-label="Range to"
							class="w-16 rounded border border-line bg-panel px-1.5 py-0.5 text-ink"
							value={customRange.hi}
							oninput={(e) => setCustom('hi', e.currentTarget.value)}
						/>
					{/if}
				{/if}
			{/if}
			{#if mode !== 'iso'}
				<div class="flex items-center gap-1" role="group" aria-label="Camera">
					{#each ['x', 'y', 'z'] as const as axis (axis)}
						<label class="flex items-center gap-1" title={coordTitle(axis)}>
							{axis.toUpperCase()}
							<input
								type="number"
								aria-label="Camera {axis}"
								class="w-16 rounded border border-line bg-panel px-1.5 py-0.5 text-ink"
								value={Math.round(view[axis])}
								oninput={(e) => setCoord(axis, e.currentTarget.value)}
							/>
						</label>
					{/each}
				</div>
				<div class="flex items-center gap-0.5" role="group" aria-label="Facing">
					<button class="rounded px-1 py-0.5 hover:bg-hover" title="Turn left" aria-label="Turn left" onclick={() => face(-1)}>⟲</button>
					<div class="flex overflow-hidden rounded border border-line">
						{#each facings as f (f)}
							<button
								class="px-1.5 py-0.5 {view.facing === f ? 'bg-accent text-white' : 'hover:bg-hover'}"
								title="Face {f}"
								aria-label="Face {f}"
								aria-pressed={view.facing === f}
								onclick={() => face(f)}>{facingLetters[f]}</button
							>
						{/each}
					</div>
					<button class="rounded px-1 py-0.5 hover:bg-hover" title="Turn right" aria-label="Turn right" onclick={() => face(1)}>⟳</button>
				</div>
			{/if}
			<label class="ml-auto flex items-center gap-1">
				Seed
				<input
					class="w-24 rounded border border-line bg-panel px-1.5 py-0.5 text-ink"
					bind:value={seed}
					oninput={() => request()}
				/>
			</label>
			<button
				class="rounded px-1.5 py-0.5 hover:bg-hover"
				onclick={() => {
					if (mode === 'iso') isoView = isoAt(view);
					else {
						view = defaultView(info.kind, mode);
						lock = null;
					}
					request();
				}}>Reset view</button
			>
		{:else}
			<span>{hint}</span>
		{/if}
	</div>
	<div class="relative min-h-0 flex-1 bg-canvas" bind:this={box}>
		<canvas
			bind:this={canvas}
			class="absolute inset-0 h-full w-full touch-none [image-rendering:pixelated]"
			class:invisible={!info || !last}
			data-view={JSON.stringify(mode === 'iso' ? isoView : view)}
			onpointerdown={down}
			onpointermove={move}
			onpointerup={up}
			onpointercancel={up}
			onpointerleave={() => (hover = null)}
			onwheel={wheel}
		></canvas>
		{#if info && error}
			<div class="absolute inset-x-3 top-3 rounded border border-red-300 bg-red-50 px-3 py-2 text-xs text-red-800 dark:border-red-800 dark:bg-red-950 dark:text-red-200">
				Preview failed: {error}
			</div>
		{/if}
		{#if info && (readout || lockValue || busy)}
			<div class="pointer-events-none absolute bottom-2 left-2 flex flex-col gap-0.5 rounded bg-black/60 px-2 py-1 font-mono text-[11px] text-white">
				{#if readout}<span>{readout}</span>{/if}
				{#if lockValue}<span class="text-red-300">● {lockValue}</span>{/if}
				{#if busy && (!last || mode !== 'density')}<span>Rendering…</span>{/if}
			</div>
		{/if}
		{#if info && last && mode !== 'iso'}
			<div class="pointer-events-none absolute top-2 left-2 rounded bg-black/60 px-2 py-1 font-mono text-[11px] text-white" aria-label="Camera facing">
				Facing {view.facing} · ← {compass(-axes.rx, -axes.rz)} {compass(axes.rx, axes.rz)} →
			</div>
		{/if}
		{#if info && last?.notice}
			<div class="pointer-events-none absolute inset-x-3 top-3 rounded bg-black/60 px-3 py-2 text-xs text-white">{last.notice}</div>
		{/if}
		{#if info && last?.legend}
			<div class="pointer-events-none absolute top-2 right-2 flex max-h-[60%] flex-col gap-0.5 overflow-hidden rounded bg-black/60 px-2 py-1 text-[11px] text-white" aria-label="Biomes in view">
				{#each last.legend.slice(0, 12) as b (b.id)}
					<span class="flex items-center gap-1.5"><span class="inline-block h-2.5 w-2.5 rounded-sm" style:background={b.color}></span>{b.id.replace(/^minecraft:/, '')}</span>
				{/each}
				{#if last.legend.length > 12}<span class="text-white/70">+{last.legend.length - 12} more</span>{/if}
			</div>
		{/if}
		{#if info && last && palette && valueRange && palette.id !== 'solid'}
			{@const r = valueRange}
			{@const flatField = last.min === last.max}
			<div class="pointer-events-none absolute right-2 bottom-2 flex w-56 flex-col gap-1 rounded bg-black/60 px-2 py-1.5 font-mono text-[11px] text-white" aria-label="Value scale">
				<div class="relative h-3 rounded-sm" style:background={cssGradient(palette)}>
					{#if r.lo < 0 && r.hi > 0}
						<span class="absolute inset-y-0 w-px bg-white/80 mix-blend-difference" style:left="{position(0, r) * 100}%" title="0"></span>
					{/if}
					{#if !Number.isNaN(hoverValue)}
						<span class="absolute -top-1 -bottom-1 w-0.5 -translate-x-1/2 rounded bg-white shadow-[0_0_0_1px_black]" style:left="{position(hoverValue, r) * 100}%"></span>
					{/if}
				</div>
				<div class="relative flex justify-between">
					<span>{formatValue(r.lo, r)}</span>
					{#if r.lo < 0 && r.hi > 0 && position(0, r) > 0.2 && position(0, r) < 0.8}
						<span class="absolute -translate-x-1/2" style:left="{position(0, r) * 100}%">0</span>
					{/if}
					<span>{formatValue(r.hi, r)}</span>
				</div>
				<div class="text-white/70">
					{#if last.min === Infinity}
						no values in view
					{:else if flatField}
						flat: {last.min.toFixed(3)} everywhere in view
					{:else}
						in view {last.min.toFixed(3)} … {last.max.toFixed(3)}
					{/if}
				</div>
			</div>
		{:else if info && last && last.min !== Infinity}
			<div class="pointer-events-none absolute right-2 bottom-2 rounded bg-black/60 px-2 py-1 font-mono text-[11px] text-white">
				{last.min.toFixed(3)} … {last.max.toFixed(3)}
			</div>
		{/if}
	</div>
</div>
