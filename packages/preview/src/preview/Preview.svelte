<script lang="ts">
	import { onMount, untrack } from 'svelte';
	import type { OutputFile } from '../types';
	import {
		canLookDown,
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
	import { climateParams, gapLayer, missing, nearest, nearestOf, type BiomeLayer, type NearEntry } from './climate';
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
	/** A render has been running long enough to say so, even over a density view. */
	let slow = $state(false);
	let slowTimer: ReturnType<typeof setTimeout> | undefined;
	function setBusy(on: boolean) {
		if (on && !busy) {
			// Every worker has until the next stall check from now on.
			const now = performance.now();
			for (const w of alive.keys()) alive.set(w, now);
			slowTimer = setTimeout(() => (slow = busy), 500);
		} else if (!on) {
			clearTimeout(slowTimer);
			slow = false;
		}
		busy = on;
	}
	// Raw: renders are replaced whole, and a biome map's entries are large.
	let last = $state.raw<Extract<PreviewResponse, { type: 'rendered' }> | null>(null);
	let hover = $state<{ px: number; py: number } | null>(null);
	type Point = { x: number; y: number; z: number };
	let lock = $state<Point | null>(null);
	let seq = 0;
	let lastTarget: string | null = null;
	let mode = $state<PreviewMode>('density');
	let biome = $state('minecraft:plains');
	let terrain = $state('minecraft:overworld');
	// The top-down switch, for views that otherwise show a side slice. It
	// stays on across targets and modes, and applies where it can.
	let lookDown = $state(false);
	const canDown = $derived(info ? canLookDown(info.kind, mode) : false);
	const flat = $derived(info ? topDown(info.kind, mode, lookDown) : false);
	/** Top-down terrain: the first block at or below the camera's y. */
	const topBlocks = $derived(flat && canDown && mode === 'blocks');
	/** The side view's y while top-down terrain looks down from the top of the world. */
	let sideY: number | null = null;

	// Density colouring: a palette per kind, remembered in this browser, and
	// the range it spans (the values in view, -1 to 1, or the user's own).
	// A biome map's climate layers and its distance-to-win layer have their own.
	type DensityKind = 'noise' | 'density_function' | 'noise_settings' | 'climate' | 'gap';
	const paletteDefaults: Record<DensityKind, PaletteId> = {
		noise: 'grayscale',
		density_function: 'viridis',
		noise_settings: 'solid',
		climate: 'diverging',
		gap: 'magma'
	};
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
		noise_settings: storedPalette('noise_settings'),
		climate: storedPalette('climate'),
		gap: storedPalette('gap')
	});

	// What a biome map shows: the biomes, one climate parameter, or how far
	// a chosen biome is from winning (multi-noise sources only).
	let layer = $state<BiomeLayer>('biomes');
	/** The biome the map explains: its nearest entry under the pointer, and the distance layer. */
	let focus = $state('');
	const layerNames: Record<BiomeLayer, string> = {
		biomes: 'Biomes',
		temperature: 'Temperature',
		humidity: 'Humidity',
		continentalness: 'Continentalness',
		erosion: 'Erosion',
		depth: 'Depth',
		weirdness: 'Weirdness',
		gap: 'Distance to win'
	};
	const biomeMap = $derived(mode === 'biomes' && info?.kind === 'dimension');
	const entries = $derived(biomeMap ? last?.entries : undefined);
	const sourceBiomes = $derived(entries ? [...new Set(entries.biomes)].sort() : []);
	const layers = $derived<BiomeLayer[]>(['biomes', ...climateParams, ...(entries ? (['gap'] as const) : [])]);
	/** The layer shown: a layer the render can't draw (no entries for the distance) falls back to the biomes. */
	const shownLayer = $derived<BiomeLayer>(!biomeMap || !last?.climate || !layers.includes(layer) ? 'biomes' : layer);
	$effect(() => {
		// Keep the focus on a biome the source has: the first one not in view, else the first.
		if (entries && !sourceBiomes.includes(focus)) focus = notInView[0] ?? sourceBiomes[0] ?? '';
	});
	/** The values a climate or distance layer colours, and their range in view. */
	const layerValues = $derived.by(() => {
		if (shownLayer === 'biomes' || !last?.climate) return null;
		const n = last.width * last.height;
		const values =
			shownLayer === 'gap'
				? entries && last.fit
					? gapLayer(entries, focus, last.climate, last.fit)
					: null
				: last.climate.subarray(climateParams.indexOf(shownLayer) * n, (climateParams.indexOf(shownLayer) + 1) * n);
		if (!values) return null;
		let min = Infinity;
		let max = -Infinity;
		for (const v of values) {
			if (v < min) min = v;
			if (v > max) max = v;
		}
		return { values, min, max };
	});
	/** Biomes the source lists that win nowhere in view. */
	const notInView = $derived(entries && last?.legend ? missing(entries, last.legend.map((b) => b.id)) : []);
	function setLayer(l: BiomeLayer) {
		layer = l;
	}
	// A new layer, biome to explain or range redraws what is there.
	$effect(() => {
		void [layerValues, palette, valueRange];
		untrack(draw);
	});

	const densityKind = $derived<DensityKind | null>(
		layerValues
			? shownLayer === 'gap'
				? 'gap'
				: 'climate'
			: info && mode === 'density' && info.kind !== 'material_rule' && info.kind !== 'dimension'
				? (info.kind as DensityKind)
				: null
	);
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
		if (!palette || !last || (last.labels && !layerValues)) return null;
		const fixed = rangeMode === 'unit' ? { lo: -1, hi: 1 } : rangeMode === 'custom' ? customRange : undefined;
		const { min, max } = layerValues ?? last;
		return rangeFor(palette, min, max, fixed);
	});
	function setRangeMode(m: RangeMode) {
		if (m === 'custom' && rangeMode !== 'custom' && last && Number.isFinite((layerValues ?? last).min)) {
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
	/**
	 * The furthest a 2D view zooms out, in blocks per pixel. Terrain is built
	 * a chunk at a time, so zooming out far gets slow, and top-down terrain
	 * needs every chunk in view: it stops at about 150 of them.
	 */
	function maxScale(): number {
		if (mode === 'density' || mode === 'biomes') return 64;
		if (!topBlocks || !box) return 4;
		const w = Math.max(16, box.clientWidth / resolution);
		const h = Math.max(16, box.clientHeight / resolution);
		// (w·s/16 + 1)(h·s/16 + 1) chunks at s blocks per pixel.
		const a = (w * h) / 256;
		const b = (w + h) / 16;
		const c = 1 - 150;
		return Math.min(4, (-b + Math.sqrt(b * b - 4 * a * c)) / (2 * a));
	}
	/**
	 * Fits the camera to whether it now shows top-down terrain: that looks
	 * down from the top of the world (lower Y to cut into it), and going
	 * back to a side view puts the camera's height back.
	 */
	function fitTop(wasTopBlocks: boolean) {
		if (topBlocks && !wasTopBlocks) {
			sideY = view.y;
			view.y = (last?.maxY ?? 320) - 1;
		} else if (!topBlocks && wasTopBlocks && sideY !== null) {
			view.y = sideY;
			sideY = null;
		}
		view.scale = Math.min(view.scale, maxScale());
	}
	function setDown(on: boolean) {
		if (!info || on === lookDown) return;
		const was = topBlocks;
		lookDown = on;
		fitTop(was);
		lock = null;
		last = null;
		request();
	}

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

	/**
	 * Workers can get stuck: a density function that repeats work can take
	 * hours to sample, and a worker busy sampling can't take the next
	 * request, so the preview would show the old picture forever. While a
	 * render runs, every worker is pinged each second; one that hasn't
	 * answered for stallMs is stuck, so all of them are replaced. A newer
	 * request the stuck worker never got to is sent again; otherwise the
	 * next change renders again. A crashed worker is replaced the same way.
	 */
	const stallMs = 6000;
	/** Each worker and when it last answered. */
	const alive = new Map<Worker, number>();
	let files: PreviewFile[] = [];
	/** The newest render the main worker has started. */
	let started = 0;

	function startWorker() {
		const w = createWorker();
		worker = w;
		alive.set(w, performance.now());
		w.onmessage = (e: MessageEvent<PreviewResponse>) => {
			alive.set(w, performance.now());
			const r = e.data;
			if (r.type === 'pong') return;
			if (r.type === 'started') {
				started = r.seq;
				return;
			}
			if (r.seq !== seq) {
				if (r.type === 'rendered') r.bitmap?.close();
				return;
			}
			setBusy(r.type === 'rendered' && !!r.partial);
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
		w.onerror = (e) => {
			e.preventDefault();
			restart(`the preview's worker crashed${e.message ? ` (${e.message})` : ''}`);
		};
		w.postMessage({ type: 'setup', files } satisfies PreviewRequest);
	}

	function stopWorkers() {
		worker?.terminate();
		worker = null;
		for (const h of helpers) h.terminate();
		helpers = [];
		alive.clear();
	}

	/**
	 * Replaces every worker. The render that broke them isn't asked for
	 * again (it would only break them again), and the preview says why; but
	 * if a newer one is waiting behind it, that one is.
	 */
	function restart(why: string) {
		const waiting = seq > started;
		stopWorkers();
		startWorker();
		started = seq;
		if (waiting) {
			request();
			return;
		}
		seq++; // drops anything still on its way from the old workers
		setBusy(false);
		error = `${why[0].toUpperCase()}${why.slice(1)}. It starts again when you change the code or the view.`;
	}

	function watch() {
		if (!busy) return;
		const now = performance.now();
		for (const [w, t] of alive) {
			if (now - t > stallMs) {
				restart(`it stopped responding after ${Math.round(stallMs / 1000)} seconds, so this view may take too long to compute`);
				return;
			}
			w.postMessage({ type: 'ping' } satisfies PreviewRequest);
		}
	}

	onMount(() => {
		startWorker();
		const watcher = setInterval(watch, 1000);
		const ro = new ResizeObserver(() => request());
		ro.observe(box);
		return () => {
			ro.disconnect();
			clearInterval(watcher);
			clearTimeout(slowTimer);
			stopWorkers();
		};
	});

	// Re-register files whenever the pack changes, then re-render.
	$effect(() => {
		files = parsed.map(({ kind, id, json }) => ({ kind, id, json }));
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
				sideY = null;
				fitTop(false);
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
		const wasTopBlocks = topBlocks;
		mode = m;
		if (topDown(info.kind, m, lookDown) !== wasFlat) {
			// Keep the camera where it is, at the new view's zoom.
			view = { ...view, scale: defaultView(info.kind, m).scale };
			lock = null;
		}
		fitTop(wasTopBlocks);
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
			alive.set(helper, performance.now());
			helper.onmessage = () => alive.set(helper, performance.now());
			helper.onerror = (e) => {
				e.preventDefault();
				restart(`a worker building chunks crashed${e.message ? ` (${e.message})` : ''}`);
			};
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
			setBusy(true);
			if (mode === 'iso' || topBlocks) startHelpers();
			worker.postMessage({
				type: 'render',
				seq: ++seq,
				kind: info.kind,
				mode,
				id: info.id,
				width,
				height,
				view: mode === 'iso' ? $state.snapshot(isoView) : { ...$state.snapshot(view), down: lookDown && canDown },
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
		} else if (palette && valueRange && layerValues) {
			const px = colorize(layerValues.values, palette, valueRange, VOID);
			// Biome edges stay drawn, so the climate lines up with the map.
			const { width: w, height: h, values: b } = last;
			for (let i = 0; i < w * h; i++) {
				const x = i % w;
				if ((x + 1 < w && b[i + 1] !== b[i]) || (i + w < w * h && b[i + w] !== b[i])) {
					px[i * 4] *= 0.6;
					px[i * 4 + 1] *= 0.6;
					px[i * 4 + 2] *= 0.6;
				}
			}
			ctx.putImageData(new ImageData(px, w, h), 0, 0);
		} else {
			ctx.putImageData(new ImageData(last.pixels, last.width, last.height), 0, 0);
		}
		if (last.seaLevel !== undefined && mode === 'density' && !flat) {
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
		if (flat) {
			// Top-down terrain shows a block at its own height; maps are at the camera's.
			const i = Math.floor(py) * (last?.width ?? 0) + Math.floor(px);
			const y = last?.heights?.[i];
			return { x: Math.floor(u * a.rx + v * a.dx), y: y === undefined || Number.isNaN(y) ? Math.floor(view.y) : y, z: Math.floor(u * a.rz + v * a.dz) };
		}
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
		return flat && !topBlocks ? `x ${p.x}, z ${p.z}` : `x ${p.x}, y ${p.y}, z ${p.z}`;
	}
	const axes = $derived(screenAxes(view.facing));
	function coordTitle(axis: 'x' | 'y' | 'z'): string {
		if (axis === 'y') {
			if (topBlocks) return 'The height the camera looks down from: blocks above it are cut away';
			return flat ? 'The height the map is sampled at' : 'The height at the middle of the view';
		}
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
			const scale = Math.min(maxScale(), Math.max(1 / 16, start.scale / factor));
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
		if (!at || !last || (mode !== 'density' && !layerValues)) return NaN;
		const ix = Math.floor(at.px);
		const iy = Math.floor(at.py);
		if (ix < 0 || iy < 0 || ix >= last.width || iy >= last.height) return NaN;
		return (layerValues?.values ?? last.values)[iy * last.width + ix];
	});
	/**
	 * A biome map's climate at the pinned point (or, with none, under the
	 * pointer), and which entries came closest there: why that biome won and
	 * the others didn't.
	 */
	const climateAt = $derived.by(() => {
		if (!biomeMap || !last?.climate) return null;
		let ix: number;
		let iy: number;
		if (lock) [ix, iy] = toPixel(lock).map(Math.floor);
		else if (hover) [ix, iy] = [Math.floor(hover.px), Math.floor(hover.py)];
		else return null;
		if (ix < 0 || iy < 0 || ix >= last.width || iy >= last.height) return null;
		const n = last.width * last.height;
		const i = iy * last.width + ix;
		const at = climateParams.map((_, c) => last!.climate![c * n + i]);
		const near = entries ? nearest(entries, at, 4) : [];
		const focused = entries && focus && !near.some((e) => e.biome === focus) ? nearestOf(entries, focus, at) : null;
		return { at, near, focused };
	});
	/**
	 * A pinned point on a biome map holds the panel on it: hovering the map
	 * doesn't change it, and the pointer can go over the panel to read it.
	 */
	const pinned = $derived(biomeMap && !!lock);
	const short = (id: string) => id.replace(/^minecraft:/, '');
	const num = (v: number) => (Math.abs(v) < 5e-4 ? '0' : v.toFixed(3));
	/** A range, short: 2 decimals, trailing zeros dropped. */
	const tight = (v: number) => (Math.abs(v) < 5e-3 ? '0' : String(+v.toFixed(2)));
	const range = (min: number, max: number) => (min === max ? tight(min) : `${tight(min)}…${tight(max)}`);
	/** The table's entry rows: the nearest ones, then the biome being explained if it isn't among them. */
	const nearRows = $derived.by((): { e: NearEntry; focused: boolean }[] =>
		climateAt ? [...climateAt.near.map((e) => ({ e, focused: false })), ...(climateAt.focused ? [{ e: climateAt.focused, focused: true }] : [])] : []
	);
	/** Short column heads, so the table fits beside the map (the full name is the title). */
	const paramHeads: Record<(typeof climateParams)[number], string> = {
		temperature: 'temp',
		humidity: 'humid',
		continentalness: 'contin',
		erosion: 'erosion',
		depth: 'depth',
		weirdness: 'weird'
	};
	const showOffset = $derived(nearRows.some((r) => r.e.offset !== 0));
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
			{#if biomeMap && last?.climate}
				<label class="flex items-center gap-1" title="Show the biomes, or the climate parameter the biome source samples">
					Show
					<select aria-label="Layer" class="rounded border border-line bg-panel px-1 py-0.5 text-ink" value={shownLayer} onchange={(e) => setLayer(e.currentTarget.value as BiomeLayer)}>
						{#each layers as l (l)}<option value={l}>{layerNames[l]}</option>{/each}
					</select>
				</label>
				{#if entries}
					<label class="flex items-center gap-1" title="The biome to explain: how far it is from winning, and what keeps it from winning under the pointer">
						Explain
						<select aria-label="Explain biome" class="rounded border border-line bg-panel px-1 py-0.5 text-ink" bind:value={focus}>
							{#each sourceBiomes as b (b)}<option value={b}>{short(b)}{notInView.includes(b) ? ' (not in view)' : ''}</option>{/each}
						</select>
					</label>
				{/if}
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
					{#if canDown}
						<button
							class="ml-1 rounded border border-line px-1.5 py-0.5 {lookDown ? 'bg-accent text-white' : 'hover:bg-hover'}"
							title={lookDown ? 'Back to a side view' : 'Look straight down from the camera'}
							aria-pressed={lookDown}
							onclick={() => setDown(!lookDown)}>Top</button
						>
					{/if}
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
						sideY = null;
						fitTop(false);
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
				{#if readout && !pinned}<span>{readout}</span>{/if}
				{#if lockValue}<span class="text-red-300">● {lockValue}</span>{/if}
				{#if busy && (!last || mode !== 'density' || slow)}<span>Rendering…</span>{/if}
			</div>
		{/if}
		{#if info && last && mode !== 'iso'}
			<div class="pointer-events-none absolute top-2 left-2 rounded bg-black/60 px-2 py-1 font-mono text-[11px] text-white" aria-label="Camera facing">
				{#if flat && canDown}Looking down, top is {compass(-axes.dx, -axes.dz)} ·{:else}Facing {view.facing} ·{/if}
				← {compass(-axes.rx, -axes.rz)} {compass(axes.rx, axes.rz)} →
			</div>
		{/if}
		{#if info && last?.notice}
			<div class="pointer-events-none absolute inset-x-3 rounded bg-black/60 px-3 py-2 text-xs text-white {mode === 'iso' ? 'top-3' : 'top-10'}">{last.notice}</div>
		{/if}
		{#if info && last?.legend}
			<div class="pointer-events-none absolute top-2 right-2 flex max-h-[60%] flex-col gap-0.5 overflow-hidden rounded bg-black/60 px-2 py-1 text-[11px] text-white" aria-label="Biomes in view">
				{#each last.legend.slice(0, 12) as b (b.id)}
					<span class="flex items-center gap-1.5"><span class="inline-block h-2.5 w-2.5 rounded-sm" style:background={b.color}></span>{b.id.replace(/^minecraft:/, '')}</span>
				{/each}
				{#if last.legend.length > 12}<span class="text-white/70">+{last.legend.length - 12} more</span>{/if}
				{#if notInView.length}
					<span class="mt-0.5 max-w-48 border-t border-white/20 pt-0.5 text-white/70" aria-label="Biomes not in view">
						Not in view: {notInView.slice(0, 8).map(short).join(', ')}{notInView.length > 8 ? ` +${notInView.length - 8} more` : ''}
					</span>
				{/if}
			</div>
		{/if}
		{#if info && last}
			<div class="pointer-events-none absolute right-2 bottom-2 flex flex-col items-end gap-1">
				{#if palette && valueRange && palette.id !== 'solid'}
					{@const r = valueRange}
					{@const stat = layerValues ?? last}
					{@const flatField = stat.min === stat.max}
					<div class="flex w-56 flex-col gap-1 rounded bg-black/60 px-2 py-1.5 font-mono text-[11px] text-white" aria-label="Value scale">
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
							{#if stat.min === Infinity}
								no values in view
							{:else if flatField}
								flat: {stat.min.toFixed(3)} everywhere in view
							{:else}
								in view {stat.min.toFixed(3)} … {stat.max.toFixed(3)}
							{/if}
						</div>
					</div>
				{:else if last.min !== Infinity}
					<div class="rounded bg-black/60 px-2 py-1 font-mono text-[11px] text-white">
						{last.min.toFixed(3)} … {last.max.toFixed(3)}
					</div>
				{/if}
				<!-- Every preview here is rendered with Misode's Deepslate. -->
				<p class="pointer-events-auto rounded bg-black/40 px-1.5 py-0.5 text-[10px] text-white/80 hover:text-white" data-credit="deepslate">
					Rendered with <a class="underline-offset-2 hover:underline" href="https://github.com/misode/deepslate" target="_blank" rel="noopener">Deepslate</a>
					by <a class="underline-offset-2 hover:underline" href="https://github.com/misode" target="_blank" rel="noopener">Misode</a>
				</p>
			</div>
		{/if}
	</div>
	{#if info && biomeMap}
		<!--
			The climate where the pointer is (or at the pinned point), and the
			entries nearest it. Below the map, at a fixed height, so it never
			covers the map or resizes it; with a point pinned it holds still,
			so the pointer can go over it.
		-->
		<div class="h-40 shrink-0 overflow-auto border-t border-line bg-panel px-3 py-1.5 font-mono text-[11px] text-ink" aria-label="Climate details">
			{#if climateAt}
				<div class="mb-1 flex items-center gap-2 font-sans text-xs text-muted">
					<span>
						{#if pinned && lockValue}Pinned: {lockValue}{:else}{readout}{/if}
						· sampled at y {Math.floor(view.y)}
					</span>
					{#if pinned}
						<button
							class="rounded border border-line px-1.5 py-0.5 hover:bg-hover"
							title="Unpin, to follow the pointer again"
							onclick={() => {
								lock = null;
								draw();
							}}>Unpin</button
						>
					{:else}
						<span class="text-muted/80">Click the map to pin a point.</span>
					{/if}
				</div>
				<table class="[&_td]:px-1 [&_td]:py-px [&_th]:px-1 [&_th]:py-px [&_th]:font-normal" aria-label="Nearest entries">
					<thead>
						<tr class="text-muted">
							<th class="text-left">{climateAt.near.length ? 'Entry' : ''}</th>
							{#if climateAt.near.length}<th class="text-right" title="The game's distance from the entry to the climate here: lowest wins">distance</th>{/if}
							{#each climateParams as p (p)}<th class="text-right" title={p}>{paramHeads[p]}</th>{/each}
							{#if showOffset}<th class="text-right" title="The entry's offset, added to its distance">offset</th>{/if}
						</tr>
					</thead>
					<tbody>
						<tr class="border-t border-line font-semibold" aria-label="Climate at the pointer">
							<td>{pinned ? 'pinned' : 'here'}</td>
							{#if climateAt.near.length}<td></td>{/if}
							{#each climateParams as p, c (p)}<td class="text-right">{num(climateAt.at[c])}</td>{/each}
							{#if showOffset}<td></td>{/if}
						</tr>
						{#each nearRows as { e, focused }, k (`${e.index}|${focused}`)}
							<tr class="border-t border-line {focused ? 'text-amber-700 dark:text-amber-300' : ''}" title={focused ? 'The biome being explained' : undefined}>
								<td class="max-w-44 truncate" title={e.biome}>{focused ? '→' : k === 0 ? '✓' : `${k + 1}.`} {short(e.biome)}</td>
								<td class="text-right">{num(e.distance)}</td>
								{#each e.params as p (p.param)}
									<td
										class="text-right whitespace-nowrap {p.off > 0 ? 'font-semibold text-red-600 dark:text-red-400' : 'text-muted'}"
										title="{p.param}: {range(p.min, p.max)}{p.off > 0 ? `; ${num(p.value)} here is ${num(p.off)} outside` : ''}"
										>{range(p.min, p.max)}</td
									>
								{/each}
								{#if showOffset}<td class="text-right {e.offset ? 'font-semibold text-red-600 dark:text-red-400' : 'text-muted'}">{tight(e.offset)}</td>{/if}
							</tr>
						{/each}
					</tbody>
				</table>
				{#if climateAt.near.length}
					<p class="mt-1 font-sans text-[11px] text-muted">Ranges in red are the ones the climate here is outside of; hover a cell for how far.</p>
				{/if}
			{:else}
				<p class="font-sans text-xs text-muted">Hover the map to see the climate there and the biome entries nearest it, or click to pin a point.</p>
			{/if}
		</div>
	{/if}
</div>
