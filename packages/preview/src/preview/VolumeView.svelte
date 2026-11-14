<script lang="ts">
	// The volume view: a density function or noise settings drawn in 3D.
	// Tiles of samples come from the preview worker (and its helpers), and
	// volume.ts ray-marches them on the GPU. Drag to orbit; right-drag or
	// WASD to travel; the window of tiles follows the camera.

	import { onMount } from 'svelte';
	import type { PreviewKind, VolumeTile, VolumeTileRequest } from './protocol';
	import Slider from './Slider.svelte';
	import { defaultSettings, VolumeRenderer, type VolumeLook, type VolumeSettings } from './volume';

	interface Props {
		kind: PreviewKind;
		id: string;
		seed: string;
		/** Bumped when the pack's files change, so every tile is sampled again. */
		version: number;
		/** Where to start, in world x and z. */
		x: number;
		z: number;
		/** Tiles sampled at once (the worker and its helpers). */
		parallel: number;
		send: (req: VolumeTileRequest) => void;
		/** The camera's target moved, to world (x, z). */
		onmove?: (x: number, z: number) => void;
	}
	let { kind, id, seed, version, x, z, parallel, send, onmove }: Props = $props();

	let canvas: HTMLCanvasElement;
	let host: HTMLDivElement;
	let gl: VolumeRenderer | null = null;
	let failure = $state('');
	let tileError = $state('');
	let settings = $state<VolumeSettings>({ ...defaultSettings });
	let at = $state({ x: 0, z: 0 });
	let loading = $state(0);
	let goX = $state('');
	let goZ = $state('');

	// ---- tiles ----
	let session = 0;
	let nextJob = 0;
	/** Tiles asked for and not yet answered, by job. */
	const inFlight = new Map<number, string>();
	/** Tiles the window needs, nearest first. */
	let queue: [number, number][] = [];
	const failed = new Set<string>();

	function pump() {
		const busy = new Set(inFlight.values());
		while (inFlight.size < parallel && queue.length) {
			const [tx, tz] = queue.shift()!;
			const key = `${tx},${tz}`;
			if (busy.has(key) || failed.has(key)) continue;
			const job = nextJob++;
			inFlight.set(job, key);
			busy.add(key);
			send({ type: 'tile', session, job, kind, id, seed, tx, tz });
		}
		loading = gl ? gl.missing() : 0;
	}

	/** A tile from the worker. */
	export function tile(t: VolumeTile) {
		if (t.session !== session) return;
		inFlight.delete(t.job);
		if ('error' in t) {
			failed.add(`${t.tx},${t.tz}`);
			tileError = t.error;
		} else {
			tileError = '';
			gl?.add(t);
		}
		pump();
	}

	/** Drops tiles in flight (the workers were restarted) and asks again. */
	export function restart() {
		fresh();
	}

	function fresh() {
		session++;
		inFlight.clear();
		failed.clear();
		queue = [];
		tileError = '';
		gl?.clear();
	}

	// A new function, seed or pack: start over, keeping the camera.
	let key = '';
	$effect(() => {
		const k = `${kind}|${id}|${seed}|${version}`;
		if (k === key) return;
		key = k;
		if (gl) fresh();
	});

	$effect(() => {
		const s = $state.snapshot(settings); // read first, so the effect follows it
		gl?.apply(s);
	});

	function set<K extends keyof VolumeSettings>(k: K, v: VolumeSettings[K]) {
		settings[k] = v;
	}

	function colors() {
		if (!gl || !host) return;
		const style = getComputedStyle(host);
		gl.setColors(style.backgroundColor || '#eceff3', style.color || '#666');
	}

	// ---- travel with the keyboard ----
	const held = new Set<string>();
	let travelling = 0;
	let lastTick = 0;
	function tick(now: number) {
		const dt = Math.min(0.05, (now - lastTick) / 1000);
		lastTick = now;
		let r = 0;
		let f = 0;
		if (held.has('a') || held.has('arrowleft')) r -= 1;
		if (held.has('d') || held.has('arrowright')) r += 1;
		if (held.has('w') || held.has('arrowup')) f += 1;
		if (held.has('s') || held.has('arrowdown')) f -= 1;
		if ((r || f) && gl) {
			const speed = (held.has('shift') ? 480 : 160) * dt; // blocks per second
			gl.travel(r * speed, f * speed);
			travelling = requestAnimationFrame(tick);
		} else {
			travelling = 0;
		}
	}
	function keydown(e: KeyboardEvent) {
		const k = e.key.toLowerCase();
		if (!['w', 'a', 's', 'd', 'arrowleft', 'arrowright', 'arrowup', 'arrowdown', 'shift'].includes(k)) return;
		e.preventDefault();
		held.add(k);
		if (!travelling && k !== 'shift') {
			lastTick = performance.now();
			travelling = requestAnimationFrame(tick);
		}
	}
	function keyup(e: KeyboardEvent) {
		held.delete(e.key.toLowerCase());
	}

	function goTo() {
		const nx = parseFloat(goX);
		const nz = parseFloat(goZ);
		if (!gl || !Number.isFinite(nx) || !Number.isFinite(nz)) return;
		gl.lookAt(nx, nz);
	}

	onMount(() => {
		try {
			gl = new VolumeRenderer(canvas);
		} catch (e) {
			failure = e instanceof Error ? e.message : String(e);
			return;
		}
		gl.want = (tiles) => {
			queue = tiles;
			pump();
		};
		gl.moved = (nx, nz) => {
			at = { x: Math.round(nx), z: Math.round(nz) };
			onmove?.(nx, nz);
		};
		const ro = new ResizeObserver(() => gl?.resize(host.clientWidth, host.clientHeight));
		ro.observe(host);
		gl.resize(host.clientWidth, host.clientHeight);
		colors();
		gl.apply($state.snapshot(settings));
		gl.lookAt(x, z);
		goX = String(Math.round(x));
		goZ = String(Math.round(z));
		// Follow the editor's light and dark themes.
		const mo = new MutationObserver(colors);
		mo.observe(document.documentElement, { attributes: true, attributeFilter: ['class', 'data-theme', 'style'] });
		const mq = matchMedia('(prefers-color-scheme: dark)');
		mq.addEventListener('change', colors);
		return () => {
			ro.disconnect();
			mo.disconnect();
			mq.removeEventListener('change', colors);
			cancelAnimationFrame(travelling);
			gl?.dispose();
			gl = null;
		};
	});

	const looks: { id: VolumeLook; name: string; note: string }[] = [
		{ id: 'caves', name: 'Caves', note: 'Stone is a faint haze; air under solid ground is drawn by height, so caves stand out.' },
		{ id: 'shell', name: 'Shell', note: 'Every surface where density crosses the threshold, see-through and coloured by height.' },
		{ id: 'solid', name: 'Solid', note: 'The first surface is opaque. Cut into it to see cave cross-sections.' }
	];
</script>

<div class="absolute inset-0 flex min-h-0" data-volume-view>
	<!-- svelte-ignore a11y_no_noninteractive_tabindex, a11y_no_noninteractive_element_interactions -->
	<div
		class="relative min-w-0 flex-1 bg-canvas text-muted outline-none"
		bind:this={host}
		tabindex="0"
		role="application"
		aria-label="3D density view. Drag to orbit, right-drag or WASD to travel, scroll to zoom."
		onkeydown={keydown}
		onkeyup={keyup}
		onblur={() => held.clear()}
	>
		<canvas bind:this={canvas} class="absolute inset-0 h-full w-full touch-none" oncontextmenu={(e) => e.preventDefault()}></canvas>
		{#if failure}
			<div class="absolute inset-x-3 top-3 rounded border border-red-300 bg-red-50 px-3 py-2 text-xs text-red-800 dark:border-red-800 dark:bg-red-950 dark:text-red-200">
				The 3D view can't start: {failure}.
			</div>
		{:else if tileError}
			<div class="absolute inset-x-3 top-3 rounded border border-red-300 bg-red-50 px-3 py-2 text-xs text-red-800 dark:border-red-800 dark:bg-red-950 dark:text-red-200">
				Preview failed: {tileError}
			</div>
		{/if}
		<div class="pointer-events-none absolute bottom-2 left-2 flex flex-col gap-0.5 rounded bg-black/60 px-2 py-1 font-mono text-[11px] text-white">
			<span>x {at.x} · z {at.z}{loading ? ` · sampling ${loading} tile${loading === 1 ? '' : 's'}…` : ''}</span>
			<span class="text-white/70">drag: orbit · right-drag or WASD: travel (Shift: faster) · scroll: zoom</span>
		</div>
	</div>
	<aside class="flex w-56 shrink-0 flex-col gap-3 overflow-y-auto border-l border-line bg-panel p-3 text-xs text-ink" aria-label="3D view settings">
		<div class="flex flex-col gap-1">
			<div class="flex overflow-hidden rounded border border-line" role="group" aria-label="Look">
				{#each looks as l (l.id)}
					<button class="flex-1 px-1.5 py-0.5 {settings.look === l.id ? 'bg-accent text-white' : 'hover:bg-hover'}" aria-pressed={settings.look === l.id} onclick={() => set('look', l.id)}
						>{l.name}</button
					>
				{/each}
			</div>
			<p class="text-[11px] text-muted">{looks.find((l) => l.id === settings.look)?.note}</p>
		</div>
		<Slider label="Threshold" value={settings.threshold} min={-0.5} max={0.5} step={0.005} reset={0} digits={3} title="Density counts as solid above this" onchange={(v) => set('threshold', v)} />
		<Slider label="Cut from top" prefix="y" value={settings.cutY} min={-64} max={320} step={1} reset={defaultSettings.cutY} onchange={(v) => set('cutY', v)} />
		<Slider label="Cut from east" suffix="%" value={settings.cutX * 100} min={1} max={100} step={1} reset={100} onchange={(v) => set('cutX', v / 100)} />
		<Slider label="Cut from south" suffix="%" value={settings.cutZ * 100} min={1} max={100} step={1} reset={100} onchange={(v) => set('cutZ', v / 100)} />
		{#if settings.look === 'caves'}
			<Slider label="Stone haze" value={settings.haze} min={0} max={0.03} step={0.0005} reset={defaultSettings.haze} digits={4} onchange={(v) => set('haze', v)} />
			<Slider label="Cave walls" value={settings.caveOpacity} min={0.05} max={1} step={0.01} reset={defaultSettings.caveOpacity} digits={2} onchange={(v) => set('caveOpacity', v)} />
		{/if}
		<div class="flex flex-col gap-1">
			<span>Colour is height</span>
			<div class="h-2 rounded-sm" style="background: linear-gradient(90deg, rgb(115 31 140), rgb(219 69 56), rgb(237 173 51), rgb(64 168 115), rgb(89 148 235))"></div>
			<div class="flex justify-between font-mono text-[10px] text-muted"><span>-64</span><span>2</span><span>68</span><span>134</span><span>200</span></div>
		</div>
		<form
			class="flex items-center gap-1"
			onsubmit={(e) => {
				e.preventDefault();
				goTo();
			}}
		>
			<span>Go to</span>
			<input aria-label="Go to x" class="w-14 rounded border border-line bg-canvas px-1 py-0.5" placeholder="x" bind:value={goX} />
			<input aria-label="Go to z" class="w-14 rounded border border-line bg-canvas px-1 py-0.5" placeholder="z" bind:value={goZ} />
			<button class="rounded border border-line px-1.5 py-0.5 hover:bg-hover">Go</button>
		</form>
	</aside>
</div>
