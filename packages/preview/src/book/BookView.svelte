<script lang="ts">
	// A written book as the game shows it: the open book with the page's
	// text in the vanilla font, page turning, links and hover text, and
	// sample values for the scores and other parts filled in when the
	// book is first opened.
	import { onDestroy } from 'svelte';
	import { loadLang } from './assets';
	import { BookRenderer, backButton, forwardButton, glyphAt, screenH, screenW } from './render';
	import type { Json, Placed, Sample } from './text';
	import type { OutputFile } from '../types';

	interface Props {
		/** The book's item stack as JSON: {id, components}. */
		stack: string;
		/** The pack, for the project's translations and fonts. */
		output: OutputFile[];
		/** The page to open, from 0, such as the page at the cursor. */
		page?: number;
	}
	let { stack, output, page = 0 }: Props = $props();

	type Obj = { [k: string]: Json };
	const content = $derived.by(() => {
		try {
			const s = JSON.parse(stack) as Obj;
			const comps = (s.components ?? {}) as Obj;
			return (comps['minecraft:written_book_content'] ?? comps.written_book_content ?? null) as Obj | null;
		} catch {
			return null;
		}
	});
	// A page is a text component, or {raw, filtered} for filtered text.
	const pages = $derived.by((): Json[] => {
		const ps = content?.pages;
		if (!Array.isArray(ps)) return [];
		return ps.map((p) => (p && typeof p === 'object' && !Array.isArray(p) && 'raw' in p ? p.raw : p));
	});
	const rawTitle = $derived(content?.title);
	const title = $derived(typeof rawTitle === 'string' ? rawTitle : rawTitle && typeof rawTitle === 'object' && 'raw' in rawTitle ? String(rawTitle.raw) : '');
	const generations = ['Original', 'Copy of original', 'Copy of a copy', 'Tattered'];

	let index = $state(0);
	$effect(() => {
		index = page;
	});
	const current = $derived(Math.max(0, Math.min(index, pages.length - 1)));

	// Sample values, by sample key; kept while the panel is open.
	let samples = $state<Record<string, string>>({});
	let used = $state<Sample[]>([]);
	const sampleOf = (s: Sample) => samples[s.key] ?? (s.kind === 'score' ? '0' : s.kind === 'selector' ? 'Steve' : '');

	// The project's own translations and fonts, from its resource pack.
	const packFiles = $derived(new Map(output.filter((f) => f.path.startsWith('assets/')).map((f) => [f.path, f.text])));
	const parsed = new Map<string, unknown>();
	function packJson(path: string): unknown {
		const text = packFiles.get(path);
		if (text === undefined) return undefined;
		if (!parsed.has(text)) {
			try {
				parsed.set(text, JSON.parse(text));
			} catch {
				parsed.set(text, undefined);
			}
		}
		return parsed.get(text);
	}

	let canvas = $state<HTMLCanvasElement>();
	let renderer = $state<BookRenderer | null>(null);
	let error = $state('');
	let glyphs = $state<Placed[]>([]);
	let lines = $state(0);
	let hoverAt = $state<{ x: number; y: number } | null>(null);
	let tooltip = $state<Placed[] | null>(null);
	let action = $state('');
	let notes = $state<string[]>([]);
	let width = $state(0);

	// The vanilla language file plus the project's en_us lines.
	$effect(() => {
		const files = packFiles;
		let stale = false;
		loadLang()
			.then(async (vanilla) => {
				const lang = { ...vanilla };
				for (const [path] of files) if (/^assets\/[^/]+\/lang\/en_us\.json$/.test(path)) Object.assign(lang, packJson(path) ?? {});
				if (stale) return;
				const r = new BookRenderer(packJson, lang);
				await r.loadChrome();
				if (!stale) renderer = r;
			})
			.catch((e) => (error = `Couldn't load the game's book assets: ${e instanceof Error ? e.message : e}`));
		return () => (stale = true);
	});

	// Lay out the current page whenever it, a sample or the assets change.
	$effect(() => {
		const r = renderer;
		const pg = pages[current];
		const vals = { ...samples };
		if (!r || pg === undefined) {
			glyphs = [];
			return;
		}
		let stale = false;
		const res = { lang: r.lang, sample: (s: Sample) => vals[s.key] ?? sampleOf(s), used: new Map<string, Sample>() };
		r.prepare(pg, res).then((laid) => {
			if (stale) return;
			glyphs = laid.glyphs;
			lines = laid.glyphs.reduce((n, g) => Math.max(n, g.line + 1), 0);
			used = [...res.used.values()];
			notes = [...r.notes];
		});
		return () => (stale = true);
	});

	const scale = $derived(Math.max(1, Math.floor(Math.min(width, 4 * screenW) / screenW)));
	const dpr = typeof devicePixelRatio === 'number' ? devicePixelRatio : 1;

	function draw() {
		if (!canvas || !renderer) return;
		const ctx = canvas.getContext('2d');
		if (!ctx) return;
		renderer.draw(ctx, scale * dpr, { glyphs, index: current, count: Math.max(1, pages.length) }, hoverAt, tooltip);
	}
	$effect(() => {
		// Redraw on any change to what is shown.
		void [glyphs, hoverAt, tooltip, scale, renderer, current];
		draw();
	});
	// Scrambled text changes every frame, as in the game.
	const scrambled = $derived(glyphs.some((g) => g.piece.style.obfuscated && g.line < 14) || (tooltip?.some((g) => g.piece.style.obfuscated) ?? false));
	let frame = 0;
	$effect(() => {
		if (!scrambled) return;
		let last = 0;
		const tick = (t: number) => {
			if (t - last > 50) {
				last = t;
				draw();
			}
			frame = requestAnimationFrame(tick);
		};
		frame = requestAnimationFrame(tick);
		return () => cancelAnimationFrame(frame);
	});
	onDestroy(() => cancelAnimationFrame(frame));

	function point(e: MouseEvent) {
		const rect = canvas!.getBoundingClientRect();
		return { x: ((e.clientX - rect.left) / rect.width) * screenW, y: ((e.clientY - rect.top) / rect.height) * screenH };
	}
	const inside = (p: { x: number; y: number }, b: typeof forwardButton) => p.x >= b.x && p.x < b.x + b.w && p.y >= b.y && p.y < b.y + b.h;

	let hoverKey = '';
	function onmove(e: MouseEvent) {
		const p = point(e);
		hoverAt = p;
		const g = glyphAt(glyphs, p);
		const h = g?.piece.style.hover;
		const key = h === undefined ? '' : JSON.stringify(h);
		if (key === hoverKey) return;
		hoverKey = key;
		if (h === undefined || !renderer) {
			tooltip = null;
			return;
		}
		const r = renderer;
		r.prepare(h, { lang: r.lang, sample: sampleOf, used: new Map() }, true).then((laid) => {
			if (hoverKey === key) tooltip = laid.glyphs;
		});
	}
	function onleave() {
		hoverAt = null;
		hoverKey = '';
		tooltip = null;
	}
	function turn(to: number) {
		index = Math.max(0, Math.min(to, pages.length - 1));
		action = '';
	}
	function onclick(e: MouseEvent) {
		const p = point(e);
		if (current < pages.length - 1 && inside(p, forwardButton)) return turn(current + 1);
		if (current > 0 && inside(p, backButton)) return turn(current - 1);
		const click = glyphAt(glyphs, p)?.piece.style.click;
		if (!click) return;
		switch (click.action) {
			case 'change_page': {
				const n = Number(click.page);
				if (Number.isInteger(n) && n >= 1 && n <= pages.length) turn(n - 1);
				else action = `Page ${click.page} doesn't exist, so the link does nothing.`;
				break;
			}
			case 'run_command':
				action = `Runs ${click.command} as the reader (the game asks first if they can't run it).`;
				break;
			case 'suggest_command':
				action = `Puts "${click.command}" in the reader's chat box.`;
				break;
			case 'copy_to_clipboard':
				action = `Copies "${click.value}" to the clipboard.`;
				break;
			case 'open_url':
				action = `Opens ${click.url} after asking the reader.`;
				break;
			default:
				action = `Click action ${click.action}.`;
		}
	}
	function onkey(e: KeyboardEvent) {
		if (e.key === 'ArrowRight' || e.key === 'PageDown') turn(current + 1);
		else if (e.key === 'ArrowLeft' || e.key === 'PageUp') turn(current - 1);
		else return;
		e.preventDefault();
	}
</script>

<div class="flex h-full flex-col overflow-auto p-3" aria-label="Book preview">
	{#if content}
		<h2 class="mb-2 text-sm">
			<span class="font-semibold">{title || 'Untitled book'}</span>
			<span class="text-xs text-muted">
				{#if typeof content.author === 'string'}· by {content.author}{/if}
				· {generations[Number(content.generation ?? 0)] ?? generations[0]}
				· {pages.length} page{pages.length === 1 ? '' : 's'}
			</span>
		</h2>
		<div bind:clientWidth={width} class="w-full">
			<!-- svelte-ignore a11y_click_events_have_key_events, a11y_no_static_element_interactions -->
			<canvas
				bind:this={canvas}
				width={screenW * scale * dpr}
				height={screenH * scale * dpr}
				style="width: {screenW * scale}px; height: {screenH * scale}px; image-rendering: pixelated"
				class="mx-auto block"
				onmousemove={onmove}
				onmouseleave={onleave}
				{onclick}
			></canvas>
		</div>
		<div class="mt-2 flex items-center justify-center gap-3 text-xs" role="group" aria-label="Pages">
			<button class="rounded px-2 py-0.5 hover:bg-hover disabled:opacity-40" disabled={current === 0} onclick={() => turn(current - 1)} onkeydown={onkey}>‹ Previous</button>
			<span class="text-muted">Page {current + 1} of {Math.max(1, pages.length)}</span>
			<button class="rounded px-2 py-0.5 hover:bg-hover disabled:opacity-40" disabled={current >= pages.length - 1} onclick={() => turn(current + 1)} onkeydown={onkey}>Next ›</button>
		</div>
		{#if error}<p class="mt-2 text-xs text-red-500">{error}</p>{/if}
		{#if !renderer && !error}<p class="mt-2 text-xs text-muted">Loading the game's book textures and font…</p>{/if}
		{#if action}<p class="mt-2 text-xs" role="status">{action}</p>{/if}
		{#if lines > 14}
			<p class="mt-2 text-xs text-amber-600 dark:text-amber-400">
				This page takes {lines} lines; the game shows 14 and cuts off the rest.
			</p>
		{/if}
		{#if used.length}
			<fieldset class="mt-3 text-xs">
				<legend class="mb-1 text-muted">Filled in when the reader first opens the book. Try values:</legend>
				{#each used as s (s.key)}
					<label class="mb-1 flex items-center gap-2">
						<span class="w-40 truncate font-mono" title={s.label}>{s.label}</span>
						<input
							class="min-w-0 flex-1 rounded border border-line bg-bg px-1.5 py-0.5"
							value={sampleOf(s)}
							oninput={(e) => (samples = { ...samples, [s.key]: e.currentTarget.value })}
						/>
					</label>
				{/each}
			</fieldset>
		{/if}
		{#each notes as n (n)}<p class="mt-2 text-xs text-muted">{n}</p>{/each}
		<p class="mt-3 text-xs text-muted">Click the arrows in the book to turn pages, or links to follow them. Point at a link to see its hover text.</p>
	{:else}
		<p class="text-xs text-muted">This item has no written book content to show.</p>
	{/if}
</div>
