<script lang="ts">
	// The jigsaw simulation preview: the pieces a jigsaw structure places
	// for a layout seed, in 3D. Boxes show each piece coloured by depth with
	// its jigsaw blocks marked (green joined, red found nothing, grey an end
	// on purpose); Blocks shows the pieces' blocks as plain coloured cubes.
	// A slice cuts the view along x, y or z to see inside. Clicking a piece
	// shows where it came from and what each of its jigsaw blocks joined.
	import { onDestroy, onMount } from 'svelte';
	import { MAX_DEPTH, simulate, structureFiles, type Box, type NoneReason, type PlacedPiece, type Simulation } from './simulate';
	import { placeBlocks } from './blocks';
	import { blockMesh, boxMesh, clip, edgeLines, fade, hit, ray, VoxelView, type Camera, type Rgb, type Slice } from './voxel';
	import { blockColor } from '../preview/colors';
	import type { OutputAsset, OutputFile } from '../types';
	import { record, restore, useMemory } from '../memory';

	interface Props {
		/** The jigsaw structure's output path. */
		path: string;
		output: OutputFile[];
		assets: OutputAsset[];
		/** The pack's structure files as data: URLs, by pack path, for the block view. */
		structures?: Record<string, string>;
	}
	let { path, output, assets, structures = {} }: Props = $props();

	let seed = $state(0);
	let mode = $state<'boxes' | 'blocks'>('boxes');
	let axis = $state<'off' | 'x' | 'y' | 'z'>('off');
	let at = $state(0);
	let selected = $state<number | null>(null);

	/** A depth set here, in place of the structure's size; null follows the size. */
	let depth = $state<number | null>(null);

	const sim: Simulation | null = $derived(simulate(path, output, assets, { seed, ...(depth !== null ? { depth } : {}) }));
	const setDepth = (v: number) => {
		depth = Number.isFinite(v) ? Math.max(0, Math.min(MAX_DEPTH, Math.round(v))) : null;
	};
	$effect(() => {
		// A new layout drops the selection, which may not exist in it.
		void sim;
		selected = null;
	});

	const ns = $derived(path.match(/^data\/([^/]+)\//)?.[1] ?? '');
	const strip = (s: string) => (s.startsWith(`${ns}:`) ? s.slice(ns.length + 1) : s.replace(/^minecraft:/, ''));

	// The extent of every placed piece, for the camera and the slider.
	const extent = $derived.by((): Box => {
		const ps = sim?.pieces ?? [];
		if (!ps.length) return [[0, 0, 0], [0, 0, 0]];
		return [[0, 1, 2].map((i) => Math.min(...ps.map((p) => p.box[0][i]))) as Box[0], [0, 1, 2].map((i) => Math.max(...ps.map((p) => p.box[1][i]))) as Box[1]];
	});
	const axisIndex = $derived(axis === 'x' ? 0 : axis === 'y' ? 1 : 2) as 0 | 1 | 2;
	const slice = $derived<Slice | null>(axis === 'off' ? null : { axis: axisIndex, at });
	// A new axis starts the slice through the middle.
	let lastAxis = 'off';
	$effect(() => {
		if (axis !== lastAxis && axis !== 'off') at = Math.round((extent[0][axisIndex] + extent[1][axisIndex]) / 2);
		lastAxis = axis;
	});

	// One colour per depth, cycling.
	const DEPTH: Rgb[] = [
		[230, 159, 0],
		[86, 180, 233],
		[0, 158, 115],
		[204, 121, 167],
		[213, 94, 0],
		[0, 114, 178],
		[240, 228, 66],
		[150, 150, 150]
	];
	const depthColor = (d: number) => DEPTH[d % DEPTH.length];
	const JOINED: Rgb = [40, 170, 70];
	const NONE: Rgb = [220, 40, 40];
	const END: Rgb = [120, 120, 120];
	const ACCENT: Rgb = [37, 99, 235];

	const fileUrls = $derived.by(() => {
		const out = new Map<string, string>();
		for (const [id, a] of structureFiles(assets)) {
			const url = structures[a.path];
			if (url) out.set(id, url);
		}
		return out;
	});
	const placed = $derived(sim && mode === 'blocks' ? placeBlocks(sim, fileUrls) : null);

	const isEnd = (r: NoneReason) => r === 'empty' || r === 'empty-piece';
	const counts = $derived.by(() => {
		let none = 0;
		for (const p of sim?.pieces ?? []) for (const j of p.jigsaws) if (j.result && 'none' in j.result && !isEnd(j.result.none)) none++;
		return { none };
	});

	let canvas: HTMLCanvasElement | undefined = $state();
	let holder: HTMLDivElement | undefined = $state();
	let view: VoxelView | null = null;
	let error = $state<string | null>(null);
	let camera = $state<Camera>({ target: [0, 0, 0], yaw: Math.PI / 4, pitch: Math.PI / 5, distance: 60 });
	let framedFor: Simulation | null = null;

	// Each structure remembers its seed, view, slice, depth and camera (see
	// memory.ts). One shown for the first time starts from its own size,
	// framed by the camera.
	const memory = useMemory();
	let shownPath: string | null = null;
	/** The structure whose camera is in place, so its first layout isn't framed over it. */
	let framedPath: string | null = null;
	$effect.pre(() => {
		if (path === shownPath) return;
		shownPath = path;
		const saved = record(memory.current.get(`jigsaw-sim:${path}`));
		seed = typeof saved?.seed === 'number' && Number.isInteger(saved.seed) ? saved.seed : 0;
		mode = saved?.mode === 'blocks' ? 'blocks' : 'boxes';
		axis = (['x', 'y', 'z'] as const).find((a) => a === saved?.axis) ?? 'off';
		lastAxis = axis; // keeps the slice where it was
		at = typeof saved?.at === 'number' && Number.isFinite(saved.at) ? saved.at : 0;
		depth = typeof saved?.depth === 'number' && Number.isFinite(saved.depth) ? Math.max(0, Math.min(MAX_DEPTH, Math.round(saved.depth))) : null;
		const cam = saved?.camera ? restore<Camera>({ target: [0, 0, 0], yaw: camera.yaw, pitch: camera.pitch, distance: 60 }, saved.camera) : null;
		framedFor = null;
		if (cam && cam.target.length === 3 && cam.target.every(Number.isFinite) && cam.distance > 0) {
			camera = cam;
			framedPath = path; // stays where it was left, not framed again
		} else framedPath = null;
	});
	$effect(() => {
		const settings = { seed, mode, axis, at, depth, camera: $state.snapshot(camera) };
		if (path === shownPath) memory.current.set(`jigsaw-sim:${path}`, settings);
	});

	function frame() {
		const [lo, hi] = extent;
		const size = Math.max(hi[0] - lo[0], hi[1] - lo[1], hi[2] - lo[2]) + 1;
		camera = { target: [(lo[0] + hi[0] + 1) / 2, (lo[1] + hi[1] + 1) / 2, (lo[2] + hi[2] + 1) / 2], yaw: camera.yaw, pitch: camera.pitch, distance: Math.max(size * 1.6, 12) };
	}

	// Rebuild the meshes when what they show changes.
	const meshes = $derived.by(() => {
		if (!sim) return null;
		const tris: Float32Array[] = [];
		const lines: Float32Array[] = [];
		const pieces = sim.pieces.map((p, i) => ({ box: p.box, rgb: selected !== null && selected !== i ? fade(depthColor(p.depth)) : depthColor(p.depth) }));
		if (mode === 'blocks' && placed) {
			tris.push(blockMesh(placed.blocks, placed.names.map(blockColor), slice, selected));
			if (selected !== null) lines.push(edgeLines([{ box: sim.pieces[selected].box, rgb: ACCENT }], slice, 0.02));
		} else {
			tris.push(boxMesh(pieces, slice));
			const marks: { box: Box; rgb: Rgb }[] = [];
			for (const p of sim.pieces)
				for (const j of p.jigsaws) {
					if (!j.result && !j.joined) continue;
					const rgb = j.joined || (j.result && 'piece' in j.result) ? JOINED : j.result && 'none' in j.result && isEnd(j.result.none) ? END : NONE;
					marks.push({ box: [j.pos, j.pos], rgb });
				}
			// Jigsaw blocks as small cubes standing out of their piece's box.
			tris.push(boxMesh(marks, slice, 0.25));
			lines.push(edgeLines(pieces.map((p) => ({ box: p.box, rgb: [40, 40, 48] as Rgb })), slice, -0.06));
			if (selected !== null) lines.push(edgeLines([{ box: sim.pieces[selected].box, rgb: ACCENT }], slice, 0.02));
		}
		lines.push(edgeLines([{ box: sim.bounds, rgb: [160, 160, 170] }], null));
		const join = (xs: Float32Array[]) => {
			const out = new Float32Array(xs.reduce((n, x) => n + x.length, 0));
			let o = 0;
			for (const x of xs) {
				out.set(x, o);
				o += x.length;
			}
			return out;
		};
		return { tris: join(tris), lines: join(lines) };
	});

	let background: Rgb = [236, 239, 243];
	function readBackground() {
		if (!holder) return;
		const m = getComputedStyle(holder).backgroundColor.match(/\d+(\.\d+)?/g);
		if (m && m.length >= 3) background = [Number(m[0]), Number(m[1]), Number(m[2])];
	}

	let raf = 0;
	function redraw() {
		cancelAnimationFrame(raf);
		raf = requestAnimationFrame(() => view?.draw(camera, background));
	}

	$effect(() => {
		// Read meshes first, so the effect follows it even before the view exists.
		const m = meshes;
		if (!view || !m) return;
		view.set(m.tris, m.lines);
		redraw();
	});
	$effect(() => {
		void camera;
		redraw();
	});
	$effect(() => {
		if (sim && sim !== framedFor && sim.pieces.length) {
			// Frame the first layout; keep the camera when re-rolling, and
			// where it was left when the structure comes back.
			if (!framedFor && framedPath !== path) frame();
			framedFor = sim;
			framedPath = path;
		}
	});

	let resize: ResizeObserver | undefined;
	onMount(() => {
		if (!canvas || !holder) return;
		try {
			view = new VoxelView(canvas);
		} catch (e) {
			error = e instanceof Error ? e.message : String(e);
			return;
		}
		readBackground();
		resize = new ResizeObserver(() => {
			if (!canvas || !holder) return;
			const dpr = window.devicePixelRatio || 1;
			canvas.width = Math.max(1, Math.round(holder.clientWidth * dpr));
			canvas.height = Math.max(1, Math.round(holder.clientHeight * dpr));
			readBackground();
			redraw();
		});
		resize.observe(holder);
		if (meshes) view.set(meshes.tris, meshes.lines);
		redraw();
	});
	onDestroy(() => {
		cancelAnimationFrame(raf);
		resize?.disconnect();
		view?.dispose();
		view = null;
	});

	// Drag to turn, shift-drag (or right-drag) to pan, wheel to zoom, click to pick.
	let drag: { x: number; y: number; moved: boolean; pan: boolean } | null = null;
	function down(e: PointerEvent) {
		drag = { x: e.clientX, y: e.clientY, moved: false, pan: e.shiftKey || e.button === 2 };
		(e.currentTarget as HTMLElement).setPointerCapture(e.pointerId);
	}
	function move(e: PointerEvent) {
		if (!drag) return;
		const dx = e.clientX - drag.x, dy = e.clientY - drag.y;
		if (Math.abs(dx) + Math.abs(dy) > 3) drag.moved = true;
		if (!drag.moved) return;
		drag.x = e.clientX;
		drag.y = e.clientY;
		if (drag.pan) {
			const s = camera.distance / 500;
			const cy = Math.cos(camera.yaw), sy = Math.sin(camera.yaw);
			camera = { ...camera, target: [camera.target[0] - dx * s * cy - dy * s * sy * Math.sin(camera.pitch), camera.target[1] + dy * s * Math.cos(camera.pitch), camera.target[2] + dx * s * sy - dy * s * cy * Math.sin(camera.pitch)] };
		} else {
			camera = { ...camera, yaw: camera.yaw - dx * 0.01, pitch: Math.max(-1.5, Math.min(1.5, camera.pitch + dy * 0.01)) };
		}
	}
	function up(e: PointerEvent) {
		if (drag && !drag.moved && sim && holder) {
			const r = holder.getBoundingClientRect();
			const { from, dir } = ray(camera, e.clientX - r.left, e.clientY - r.top, r.width, r.height);
			let best: number | null = null, dist = Infinity;
			sim.pieces.forEach((p, i) => {
				const b = clip(p.box, slice);
				const t = b && hit(from, dir, b);
				if (t !== null && t !== undefined && t < dist) [best, dist] = [i, t];
			});
			selected = best;
		}
		drag = null;
	}
	function wheel(e: WheelEvent) {
		e.preventDefault();
		camera = { ...camera, distance: Math.max(2, Math.min(2000, camera.distance * Math.exp(e.deltaY * 0.001))) };
	}

	const piece = $derived(selected !== null && sim ? sim.pieces[selected] : null);
	const turned = (p: PlacedPiece) => ['not turned', 'turned 90° clockwise', 'turned 180°', 'turned 90° counterclockwise'][p.rotation];
	const sizeOf = (b: Box) => `${b[1][0] - b[0][0] + 1}×${b[1][1] - b[0][1] + 1}×${b[1][2] - b[0][2] + 1}`;
	const why: Record<NoneReason, string> = {
		empty: 'an end: its pool is minecraft:empty',
		'unknown-pool': 'nothing: its pool is not in the pack',
		depth: 'nothing: at the last depth, and no fallback piece fit',
		'no-fit': 'nothing: no piece could join or fit',
		'empty-piece': 'nothing: the pool drew the empty piece'
	};
</script>

<div class="flex h-full flex-col text-xs" aria-label="Jigsaw simulation">
	<div class="flex flex-wrap items-center gap-x-3 gap-y-1.5 border-b border-line px-3 py-2">
		<label class="flex items-center gap-1" title="The same layout seed always grows the same layout here. It is not a world seed: the game also seeds by chunk, and terrain is left out.">
			<span class="text-muted">Layout seed</span>
			<input class="w-20 rounded border border-line bg-panel px-1 py-0.5 font-mono" type="number" step="1" bind:value={seed} aria-label="Layout seed" />
		</label>
		<button type="button" class="rounded border border-line px-1.5 py-0.5 hover:bg-hover" onclick={() => (seed = Math.floor(Math.random() * 2 ** 31))}>Re-roll</button>
		{#if sim}
			<label class="flex items-center gap-1" title="How many pieces deep to grow from the start piece, 0 to {MAX_DEPTH} (the game's limit). Starts at the structure's size, {sim.size}; this only changes the preview, not the pack.">
				<span class="text-muted">Depth</span>
				<input class="w-12 rounded border border-line bg-panel px-1 py-0.5 font-mono" type="number" min="0" max={MAX_DEPTH} step="1" value={sim.depth} oninput={(e) => setDepth(e.currentTarget.valueAsNumber)} aria-label="Depth" />
				<input type="range" class="w-24" min="0" max={MAX_DEPTH} step="1" value={sim.depth} oninput={(e) => setDepth(e.currentTarget.valueAsNumber)} aria-label="Depth slider" />
			</label>
			{#if depth !== null && depth !== sim.size}
				<button type="button" class="rounded border border-line px-1.5 py-0.5 hover:bg-hover" title="Go back to the structure's own size" onclick={() => (depth = null)}>Size {sim.size}</button>
			{/if}
		{/if}
		<span class="flex overflow-hidden rounded border border-line" role="group" aria-label="Show">
			{#each [['boxes', 'Boxes'], ['blocks', 'Blocks']] as [id, label] (id)}
				<button type="button" class="px-1.5 py-0.5 {mode === id ? 'bg-selected font-medium' : 'hover:bg-hover'}" aria-pressed={mode === id} onclick={() => (mode = id as typeof mode)}>{label}</button>
			{/each}
		</span>
		<span class="flex items-center gap-1" role="group" aria-label="Slice">
			<span class="text-muted" title="Cut the view to see inside: keep only what is at or below the slider along an axis">Slice</span>
			<span class="flex overflow-hidden rounded border border-line">
				{#each ['off', 'x', 'y', 'z'] as id (id)}
					<button type="button" class="px-1.5 py-0.5 {axis === id ? 'bg-selected font-medium' : 'hover:bg-hover'}" aria-pressed={axis === id} onclick={() => (axis = id as typeof axis)}>{id === 'off' ? 'Off' : id}</button>
				{/each}
			</span>
			{#if axis !== 'off'}
				<input type="range" class="w-28" min={extent[0][axisIndex]} max={extent[1][axisIndex]} step="1" bind:value={at} aria-label="Slice position" />
				<span class="w-10 font-mono">{axis} ≤ {at}</span>
			{/if}
		</span>
		<button type="button" class="rounded border border-line px-1.5 py-0.5 hover:bg-hover" title="Fit the whole structure in view" onclick={frame}>Fit</button>
		{#if sim}
			<span class="ml-auto text-muted" aria-label="Simulation summary">
				{#if sim.failed}
					<span class="text-amber-600 dark:text-amber-400">{sim.failed}</span>
				{:else}
					{sim.pieces.length} {sim.pieces.length === 1 ? 'piece' : 'pieces'} · {sim.reached} deep (size {sim.size}{sim.depth !== sim.size ? `, grown to ${sim.depth}` : ''}){#if counts.none}
						· <span class="text-red-600 dark:text-red-400">{counts.none} jigsaw {counts.none === 1 ? 'block' : 'blocks'} found nothing</span>{/if}
				{/if}
			</span>
		{/if}
	</div>
	<div class="relative min-h-0 flex-1 bg-canvas" bind:this={holder}>
		<canvas
			bind:this={canvas}
			class="absolute inset-0 h-full w-full touch-none select-none"
			aria-label="Simulated structure"
			onpointerdown={down}
			onpointermove={move}
			onpointerup={up}
			onwheel={wheel}
			ondblclick={frame}
			oncontextmenu={(e) => e.preventDefault()}
		></canvas>
		{#if error}
			<p class="absolute inset-x-3 top-3 rounded border border-red-300 bg-red-50 px-3 py-2 text-red-800 dark:border-red-800 dark:bg-red-950 dark:text-red-200">The 3D view needs WebGL 2, which isn't available here ({error}).</p>
		{/if}
		{#if mode === 'blocks' && placed?.missing.length}
			<p class="absolute left-3 top-3 rounded bg-panel/90 px-2 py-1 text-muted">Blocks not read for {placed.missing.map(strip).join(', ')}.</p>
		{/if}
		<p class="pointer-events-none absolute bottom-2 left-3 text-[10px] text-muted">Drag to turn · shift-drag to pan · wheel to zoom · click a piece · double-click to fit</p>
		{#if sim?.notes.length}
			<ul class="absolute bottom-6 left-3 max-w-[60%] text-[10px] text-muted" aria-label="Simulation notes">
				{#each sim.notes as n (n)}<li>{n}</li>{/each}
			</ul>
		{/if}
		{#if piece}
			<div class="absolute right-2 top-2 max-h-[calc(100%-1rem)] w-72 overflow-auto rounded border border-line bg-panel p-2 shadow" aria-label="Simulated piece">
				<div class="flex items-baseline gap-2">
					<span class="inline-block h-2.5 w-2.5 shrink-0 rounded-sm" style="background: rgb({depthColor(piece.depth).join(',')})"></span>
					<span class="font-mono font-semibold">{strip(piece.id)}</span>
					<button type="button" class="ml-auto text-muted hover:text-ink" aria-label="Close" onclick={() => (selected = null)}>✕</button>
				</div>
				<p class="mt-1 text-muted">
					{piece.depth === 0 ? 'The start piece' : `Depth ${piece.depth}`}, from {strip(piece.pool)}; {turned(piece)}; {sizeOf(piece.box)} at {piece.box[0].join(' ')}.
					{#if piece.parent}Joined onto <button type="button" class="text-accent hover:underline" onclick={() => (selected = piece.parent!.piece)}>{strip(sim!.pieces[piece.parent.piece].id)}</button>.{/if}
				</p>
				{#if piece.jigsaws.length}
					<table class="mt-1.5 w-full">
						<tbody>
							{#each piece.jigsaws as j, i (i)}
								<tr class="border-t border-line align-top">
									<td class="py-0.5 pr-1 font-mono">{strip(j.name)}<span class="text-muted"> → {strip(j.target)}</span></td>
									<td class="py-0.5">
										{#if j.joined}<span class="text-muted">joined its parent</span>
										{:else if j.result && 'piece' in j.result}{@const to = j.result.piece}<button type="button" class="text-left text-accent hover:underline" onclick={() => (selected = to)}>{strip(sim!.pieces[to].id)}</button>
										{:else if j.result && 'none' in j.result}<span class={isEnd(j.result.none) ? 'text-muted' : 'text-red-600 dark:text-red-400'}>{why[j.result.none]}</span>
										{/if}
									</td>
								</tr>
							{/each}
						</tbody>
					</table>
				{/if}
			</div>
		{/if}
	</div>
</div>
