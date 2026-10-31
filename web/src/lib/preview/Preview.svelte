<script lang="ts">
	import { onMount } from 'svelte';
	import type { OutputFile } from '../engine/engine.svelte';
	import { previewable, type PreviewKind, type PreviewRequest, type PreviewResponse, type PreviewView } from './protocol';

	interface Props {
		output: OutputFile[];
		/** The output path to preview. */
		target: string | null;
	}
	let { output, target }: Props = $props();

	// A target the pack no longer has (renamed or deleted) shows nothing.
	const info = $derived(target && output.some((f) => f.path === target) ? previewable(target) : null);
	const defaults: Record<PreviewKind, PreviewView> = {
		noise: { x: 0, y: 0, scale: 4 },
		density_function: { x: 0, y: 64, scale: 2 },
		noise_settings: { x: 0, y: 64, scale: 2 }
	};

	let canvas: HTMLCanvasElement;
	let box: HTMLDivElement;
	let worker: Worker | null = null;
	let view = $state<PreviewView>({ ...defaults.density_function });
	let seed = $state('0');
	let error = $state('');
	let busy = $state(false);
	let last: Extract<PreviewResponse, { type: 'rendered' }> | null = $state(null);
	let hover = $state<{ px: number; py: number } | null>(null);
	let lock = $state<{ x: number; y: number } | null>(null);
	let seq = 0;
	let lastKind: PreviewKind | null = null;
	// CSS pixels per rendered pixel; terrain is costly to sample, so coarser.
	const resolution = $derived(info?.kind === 'noise_settings' ? 4 : 2);

	onMount(() => {
		worker = new Worker(new URL('./preview.worker.ts', import.meta.url), { type: 'module' });
		worker.onmessage = (e: MessageEvent<PreviewResponse>) => {
			const r = e.data;
			if (r.seq !== seq) return;
			busy = false;
			if (r.type === 'error') {
				error = r.message;
				return;
			}
			error = '';
			last = r;
			draw();
		};
		const ro = new ResizeObserver(() => request());
		ro.observe(box);
		return () => {
			ro.disconnect();
			worker?.terminate();
		};
	});

	// Re-register files whenever the pack changes, then re-render.
	$effect(() => {
		const files: { kind: PreviewKind; id: string; json: unknown }[] = [];
		for (const f of output) {
			const p = previewable(f.path);
			if (!p) continue;
			try {
				files.push({ ...p, json: JSON.parse(f.text) });
			} catch {
				/* skip */
			}
		}
		worker?.postMessage({ type: 'setup', files } satisfies PreviewRequest);
		request();
	});

	$effect(() => {
		if (info && info.kind !== lastKind) {
			lastKind = info.kind;
			view = { ...defaults[info.kind] };
			lock = null;
		}
		request();
	});

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
			worker.postMessage({
				type: 'render',
				seq: ++seq,
				kind: info.kind,
				id: info.id,
				width,
				height,
				view: $state.snapshot(view),
				seed
			} satisfies PreviewRequest);
		});
	}

	function draw() {
		if (!last || !canvas) return;
		canvas.width = last.width;
		canvas.height = last.height;
		const ctx = canvas.getContext('2d')!;
		ctx.putImageData(new ImageData(last.pixels, last.width, last.height), 0, 0);
		if (last.seaLevel !== undefined && info?.kind !== 'noise') {
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
		const y = info?.kind === 'noise' ? view.y + (py - h / 2) * view.scale : view.y - (py - h / 2) * view.scale;
		return [Math.floor(x), Math.floor(y)];
	}
	function toPixel(x: number, y: number): [number, number] {
		const w = last?.width ?? 0;
		const h = last?.height ?? 0;
		const px = (x - view.x) / view.scale + w / 2;
		const py = info?.kind === 'noise' ? (y - view.y) / view.scale + h / 2 : (view.y - y) / view.scale + h / 2;
		return [px, py];
	}

	let drag: { x: number; y: number; vx: number; vy: number; moved: boolean } | null = null;
	function pointer(e: PointerEvent) {
		const r = canvas.getBoundingClientRect();
		return { px: ((e.clientX - r.left) / r.width) * (last?.width ?? 1), py: ((e.clientY - r.top) / r.height) * (last?.height ?? 1) };
	}
	function down(e: PointerEvent) {
		canvas.setPointerCapture(e.pointerId);
		drag = { x: e.clientX, y: e.clientY, vx: view.x, vy: view.y, moved: false };
	}
	function move(e: PointerEvent) {
		hover = pointer(e);
		if (!drag) return;
		const dx = (e.clientX - drag.x) / resolution;
		const dy = (e.clientY - drag.y) / resolution;
		if (Math.abs(dx) + Math.abs(dy) > 2) drag.moved = true;
		if (!drag.moved) return;
		view.x = drag.vx - dx * view.scale;
		view.y = info?.kind === 'noise' ? drag.vy - dy * view.scale : drag.vy + dy * view.scale;
		request();
	}
	function up(e: PointerEvent) {
		if (drag && !drag.moved) {
			const p = pointer(e);
			const [x, y] = toWorld(p.px, p.py);
			lock = lock && lock.x === x && lock.y === y ? null : { x, y };
			draw();
		}
		drag = null;
	}
	function wheel(e: WheelEvent) {
		e.preventDefault();
		const p = pointer(e as unknown as PointerEvent);
		const [wx, wy] = toWorld(p.px, p.py);
		const f = Math.exp(e.deltaY * 0.0015);
		const scale = Math.min(64, Math.max(1 / 16, view.scale * f));
		// Keep the point under the cursor fixed.
		const w = last?.width ?? 0;
		const h = last?.height ?? 0;
		view.scale = scale;
		view.x = wx - (p.px - w / 2) * scale;
		view.y = info?.kind === 'noise' ? wy - (p.py - h / 2) * scale : wy + (p.py - h / 2) * scale;
		request();
	}

	const readout = $derived.by(() => {
		const at = hover;
		if (!at || !last) return '';
		const ix = Math.floor(at.px);
		const iy = Math.floor(at.py);
		if (ix < 0 || iy < 0 || ix >= last.width || iy >= last.height) return '';
		const v = last.values[iy * last.width + ix];
		const [x, y] = toWorld(at.px, at.py);
		const coord = info?.kind === 'noise' ? `x ${x}, z ${y}` : `x ${x}, y ${y}`;
		return Number.isNaN(v) ? coord : `${coord} · ${v.toFixed(4)}`;
	});
	const lockValue = $derived.by(() => {
		if (!lock || !last) return '';
		const [px, py] = toPixel(lock.x, lock.y);
		const ix = Math.floor(px);
		const iy = Math.floor(py);
		if (ix < 0 || iy < 0 || ix >= last.width || iy >= last.height) return `${lock.x}, ${lock.y}`;
		const v = last.values[iy * last.width + ix];
		return `${lock.x}, ${lock.y} · ${Number.isNaN(v) ? '–' : v.toFixed(4)}`;
	});
</script>

<div class="flex h-full min-h-0 flex-col">
	<div class="flex flex-wrap items-center gap-2 border-b border-line px-3 py-1.5 text-xs text-muted">
		{#if info}
			<span class="font-medium text-ink">{info.id}</span>
			<span>{info.kind.replace('_', ' ')}</span>
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
					view = { ...defaults[info.kind] };
					request();
				}}>Reset view</button
			>
		{:else}
			<span>Put the cursor in a noise, density function or noise settings declaration to preview it.</span>
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
				{#if busy && !last}<span>Rendering…</span>{/if}
			</div>
		{/if}
		{#if info && last && last.min !== Infinity}
			<div class="pointer-events-none absolute right-2 bottom-2 rounded bg-black/60 px-2 py-1 font-mono text-[11px] text-white">
				{last.min.toFixed(3)} … {last.max.toFixed(3)}
			</div>
		{/if}
	</div>
</div>
