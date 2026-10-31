<script lang="ts">
	import { onMount } from 'svelte';
	import type { OutputFile } from '../engine/engine.svelte';
	import {
		modesOf,
		previewable,
		previewFile,
		topDown,
		type PreviewFile,
		type PreviewKind,
		type PreviewMode,
		type PreviewRequest,
		type PreviewResponse,
		type PreviewView
	} from './protocol';

	interface Props {
		output: OutputFile[];
		/** The output path to preview. */
		target: string | null;
	}
	let { output, target }: Props = $props();

	// A target the pack no longer has (renamed or deleted) shows nothing.
	const info = $derived(target && output.some((f) => f.path === target) ? previewable(target) : null);
	const defaults: Record<PreviewKind | 'biomes', PreviewView> = {
		noise: { x: 0, y: 0, scale: 4 },
		density_function: { x: 0, y: 64, scale: 2 },
		noise_settings: { x: 0, y: 64, scale: 2 },
		material_rule: { x: 0, y: 64, scale: 0.5 },
		dimension: { x: 0, y: 64, scale: 2 },
		biomes: { x: 0, y: 0, scale: 8 }
	};
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
	let isoView = $state<PreviewView>({ x: 0, y: 0, scale: 0 });
	let seed = $state('0');
	let error = $state('');
	let busy = $state(false);
	let last: Extract<PreviewResponse, { type: 'rendered' }> | null = $state(null);
	let hover = $state<{ px: number; py: number } | null>(null);
	let lock = $state<{ x: number; y: number } | null>(null);
	let seq = 0;
	let lastTarget: string | null = null;
	let mode = $state<PreviewMode>('density');
	let biome = $state('minecraft:plains');
	let terrain = $state('minecraft:overworld');
	let biomeY = $state(64);
	const flat = $derived(info ? topDown(info.kind, mode) : false);
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
		worker = new Worker(new URL('./preview.worker.ts', import.meta.url), { type: 'module' });
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
			view = defaultView(info.kind, m);
			lock = null;
		}
		view.scale = Math.min(view.scale, maxScale);
		if (m === 'iso') {
			// Start the 3D view over the slice's x, with the slice (z = 0) through the middle.
			isoView = { x: Math.round(view.x), y: 0, scale: 0 };
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
			const helper = new Worker(new URL('./preview.worker.ts', import.meta.url), { type: 'module' });
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
				options: { biome, terrain, y: biomeY }
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
		if (lock) {
			const [px, py] = toPixel(lock.x, lock.y);
			ctx.strokeStyle = '#ef4444';
			ctx.lineWidth = 1.5;
			ctx.beginPath();
			ctx.arc(px, py, 4, 0, Math.PI * 2);
			ctx.stroke();
		}
	}

	function toWorld(px: number, py: number): [number, number] {
		const w = last?.width ?? 0;
		const h = last?.height ?? 0;
		const x = view.x + (px - w / 2) * view.scale;
		// Side views have y up; the noise view has z down.
		const y = flat ? view.y + (py - h / 2) * view.scale : view.y - (py - h / 2) * view.scale;
		return [Math.floor(x), Math.floor(y)];
	}
	function toPixel(x: number, y: number): [number, number] {
		const w = last?.width ?? 0;
		const h = last?.height ?? 0;
		const px = (x - view.x) / view.scale + w / 2;
		const py = flat ? (y - view.y) / view.scale + h / 2 : (view.y - y) / view.scale + h / 2;
		return [px, py];
	}

	let drag: { x: number; y: number; vx: number; vy: number; moved: boolean } | null = null;
	function pointer(e: PointerEvent) {
		const r = canvas.getBoundingClientRect();
		return { px: ((e.clientX - r.left) / r.width) * (last?.width ?? 1), py: ((e.clientY - r.top) / r.height) * (last?.height ?? 1) };
	}
	function down(e: PointerEvent) {
		canvas.setPointerCapture(e.pointerId);
		const v = mode === 'iso' ? isoView : view;
		drag = { x: e.clientX, y: e.clientY, vx: v.x, vy: v.y, moved: false };
	}
	function move(e: PointerEvent) {
		hover = pointer(e);
		if (!drag) return;
		const dx = (e.clientX - drag.x) / resolution;
		const dy = (e.clientY - drag.y) / resolution;
		if (Math.abs(dx) + Math.abs(dy) > 2) drag.moved = true;
		if (!drag.moved) return;
		if (mode === 'iso') {
			// Undo the projection (screen x = (x - z)·s, screen y = (x + z)·s/2),
			// so the ground follows the pointer.
			const s = isoView.scale || 8;
			isoView.x = drag.vx - (dx / s + (2 * dy) / s) / 2;
			isoView.y = drag.vy - ((2 * dy) / s - dx / s) / 2;
			request();
			return;
		}
		view.x = drag.vx - dx * view.scale;
		view.y = flat ? drag.vy - dy * view.scale : drag.vy + dy * view.scale;
		request();
	}
	function up(e: PointerEvent) {
		if (drag && !drag.moved && mode !== 'iso') {
			const p = pointer(e);
			const [x, y] = toWorld(p.px, p.py);
			lock = lock && lock.x === x && lock.y === y ? null : { x, y };
			draw();
		}
		drag = null;
	}
	function wheel(e: WheelEvent) {
		e.preventDefault();
		if (mode === 'iso') {
			// Zooming out past a 2×2-chunk area builds more chunks, up to 4×4.
			const s = isoView.scale || 8;
			isoView.scale = Math.min(48, Math.max(last?.minScale ?? 2, s * Math.exp(-e.deltaY * 0.0015)));
			request();
			return;
		}
		const p = pointer(e as unknown as PointerEvent);
		const [wx, wy] = toWorld(p.px, p.py);
		const f = Math.exp(e.deltaY * 0.0015);
		const scale = Math.min(maxScale, Math.max(1 / 16, view.scale * f));
		// Keep the point under the cursor fixed.
		const w = last?.width ?? 0;
		const h = last?.height ?? 0;
		view.scale = scale;
		view.x = wx - (p.px - w / 2) * scale;
		view.y = flat ? wy - (p.py - h / 2) * scale : wy + (p.py - h / 2) * scale;
		request();
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
		const [x, y] = toWorld(at.px, at.py);
		const coord = flat ? `x ${x}, z ${y}` : `x ${x}, y ${y}`;
		return Number.isNaN(v) ? coord : `${coord} · ${show(v)}`;
	});
	const lockValue = $derived.by(() => {
		if (!lock || !last) return '';
		const [px, py] = toPixel(lock.x, lock.y);
		const ix = Math.floor(px);
		const iy = Math.floor(py);
		if (ix < 0 || iy < 0 || ix >= last.width || iy >= last.height) return `${lock.x}, ${lock.y}`;
		const v = last.values[iy * last.width + ix];
		return `${lock.x}, ${lock.y} · ${Number.isNaN(v) ? '–' : show(v)}`;
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
			{#if mode === 'biomes'}
				<label class="flex items-center gap-1">
					Y
					<input
						type="number"
						class="w-16 rounded border border-line bg-panel px-1.5 py-0.5 text-ink"
						bind:value={biomeY}
						oninput={() => request()}
					/>
				</label>
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
					if (mode === 'iso') isoView = { x: Math.round(view.x), y: 0, scale: 0 };
					else view = defaultView(info.kind, mode);
					request();
				}}>Reset view</button
			>
		{:else}
			<span>Put the cursor in a noise, density function, noise settings, surface rule or dimension declaration to preview it.</span>
		{/if}
	</div>
	<div class="relative min-h-0 flex-1 bg-canvas" bind:this={box}>
		<canvas
			bind:this={canvas}
			class="absolute inset-0 h-full w-full touch-none [image-rendering:pixelated]"
			class:invisible={!info || !last}
			onpointerdown={down}
			onpointermove={move}
			onpointerup={up}
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
		{#if info && last && last.min !== Infinity}
			<div class="pointer-events-none absolute right-2 bottom-2 rounded bg-black/60 px-2 py-1 font-mono text-[11px] text-white">
				{last.min.toFixed(3)} … {last.max.toFixed(3)}
			</div>
		{/if}
	</div>
</div>
