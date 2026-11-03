<script lang="ts">
	// The jigsaw connection preview: the pools a structure grows through,
	// left to right by depth. One labelled line per pool and jigsaw target
	// shows which pool grows into which; pointing at a piece shows its own
	// lines to the pieces that can join it, and lights those pieces. Dead
	// ends, pieces nothing can join, and pools past the structure's size
	// are marked.
	import { type JigsawGraph, type Link, type Piece } from './graph';
	import { jigsawMarkdown, jigsawMermaid, joinsText } from './export';

	let { graph }: { graph: JigsawGraph } = $props();

	// Layout, in px: one column per depth, pools stacked in each.
	const W = 212;
	const GAP = 72;
	const HEAD = 30;
	const ROW = 24;
	const PAD = 8;
	const TOP = 28;
	const LEFT = 12;

	// Ids in the minecraft namespace or the root's own are shown without it.
	const ns = $derived(graph.root.id.split(':')[0]);
	const strip = (s: string) => (s.startsWith(`${ns}:`) ? s.slice(ns.length + 1) : s.replace(/^minecraft:/, ''));
	const pct = (n: number) => (n >= 0.995 ? '100%' : n < 0.01 ? '<1%' : `${Math.round(n * 100)}%`);

	interface Box {
		x: number;
		y: number;
		h: number;
		/** The y of each piece row's middle. */
		rows: number[];
	}
	const layout = $derived.by(() => {
		const cols: number[] = [];
		const boxes: Box[] = graph.pools.map((pool) => {
			const y = cols[pool.depth] ?? TOP;
			const n = Math.max(pool.pieces.length, 1) + (pool.fallback ? 1 : 0);
			const h = HEAD + n * ROW + PAD;
			cols[pool.depth] = y + h + 20;
			return { x: LEFT + pool.depth * (W + GAP), y, h, rows: pool.pieces.map((_, i) => y + HEAD + i * ROW + ROW / 2) };
		});
		const depth = graph.pools.reduce((n, p) => Math.max(n, p.depth), 0);
		return { boxes, width: LEFT * 2 + (depth + 1) * W + depth * GAP, height: Math.max(...cols.filter((c) => c !== undefined), TOP + 60), depth };
	});

	interface Edge {
		key: string;
		from: [number, number];
		to: [number, number];
		/** The pieces at either end, as pool:piece; the head end may be a pool. */
		ends: string[];
		dead: boolean;
		maybe: boolean;
		fallback?: boolean;
		/** Back to the same or an earlier depth, where growth repeats. */
		back?: boolean;
		title: string;
	}
	// Piece to piece, shown for the piece in focus; dead ends always.
	const edges = $derived.by(() => {
		const out = new Map<string, Edge>();
		const { boxes } = layout;
		graph.pools.forEach((pool, pi) => {
			const b = boxes[pi];
			pool.pieces.forEach((p, qi) => {
				const from: [number, number] = [b.x + W, b.rows[qi]];
				for (const l of p.links) {
					const j = p.jigsaws![l.jigsaw];
					const name = `${strip(j.name)} → ${strip(j.target)} in ${strip(l.pool)}`;
					const heads: { at: [number, number]; end: string }[] = [];
					if (l.to !== undefined) {
						const tb = boxes[l.to];
						for (const ti of l.pieces) heads.push({ at: [tb.x, tb.rows[ti]], end: `${l.to}:${ti}` });
						if (l.fallback) {
							const fb = boxes[l.fallback.to];
							for (const ti of l.fallback.pieces) heads.push({ at: [fb.x, fb.rows[ti]], end: `${l.fallback.to}:${ti}` });
						} else if (!l.pieces.length) heads.push({ at: [tb.x, tb.y + HEAD / 2], end: `${l.to}` });
					}
					for (const h of heads) {
						const key = `${pi}:${qi}>${h.end}`;
						const e = out.get(key);
						if (e) {
							e.title += `\n${name}`;
							e.dead &&= !!l.dead;
							continue;
						}
						const back = graph.pools[Number(h.end.split(':')[0])].depth <= pool.depth;
						out.set(key, { key, from, to: h.at, back, ends: [`${pi}:${qi}`, h.end], dead: !!l.dead, maybe: (l.maybe && !l.pieces.length) || !!l.fallback, title: l.fallback ? `${name} (via fallback)` : name });
					}
				}
			});
			if (pool.fallbackTo !== undefined) {
				const tb = boxes[pool.fallbackTo];
				const y = b.y + b.h - PAD - ROW / 2;
				out.set(`${pi}>fb`, { key: `${pi}>fb`, from: [b.x + W, y], to: [tb.x, tb.y + HEAD / 2], ends: [`${pi}`, `${pool.fallbackTo}`], dead: false, maybe: false, fallback: true, title: `fallback of ${strip(pool.id)}` });
			}
		});
		return [...out.values()];
	});

	// Pool to pool: one line per source pool, pool it grows into and
	// target name, from the pieces with that jigsaw to the pieces that
	// match it. Lines back to an earlier depth (where growth repeats) are
	// left out; the column label says it repeats.
	interface Bundle {
		key: string;
		from: [number, number];
		to: [number, number];
		label: string;
		/** Source pieces and matched pieces, as pool:piece. */
		ends: Set<string>;
		maybe: boolean;
		title: string;
	}
	const bundles = $derived.by(() => {
		const acc = new Map<string, { from: number[]; to: number[]; pi: number; ti: number; target: string; ends: Set<string>; maybe: boolean; srcs: Set<string> }>();
		const { boxes } = layout;
		graph.pools.forEach((pool, pi) => {
			pool.pieces.forEach((p, qi) => {
				for (const l of p.links) {
					if (l.dead || l.to === undefined) continue;
					const target = p.jigsaws![l.jigsaw].target;
					const heads: [number, number[]][] = l.fallback ? [[l.fallback.to, l.fallback.pieces]] : [[l.to, l.pieces]];
					for (const [ti, pieces] of heads) {
						if (graph.pools[ti].depth <= pool.depth) continue;
						const key = `${pi}>${ti}|${target}`;
						let a = acc.get(key);
						if (!a) acc.set(key, (a = { from: [], to: [], pi, ti, target, ends: new Set(), maybe: false, srcs: new Set() }));
						if (!a.srcs.has(`${qi}`)) {
							a.srcs.add(`${qi}`);
							a.from.push(boxes[pi].rows[qi]);
						}
						a.ends.add(`${pi}:${qi}`);
						for (const t of pieces) {
							if (!a.ends.has(`${ti}:${t}`)) a.to.push(boxes[ti].rows[t]);
							a.ends.add(`${ti}:${t}`);
						}
						a.maybe ||= !!l.fallback || (l.maybe && !pieces.length);
					}
				}
			});
		});
		const mean = (ys: number[]) => ys.reduce((n, y) => n + y, 0) / ys.length;
		return [...acc.entries()].map(([key, a]): Bundle => {
			const fb = boxes[a.pi], tb = boxes[a.ti];
			const into = a.to.length ? a.to.length : 0;
			return {
				key,
				from: [fb.x + W, mean(a.from)],
				to: [tb.x, into ? mean(a.to) : tb.y + HEAD / 2],
				label: strip(a.target),
				ends: a.ends,
				maybe: a.maybe,
				title: `${a.srcs.size} ${a.srcs.size === 1 ? 'piece' : 'pieces'} of ${strip(graph.pools[a.pi].id)} join ${into || 'unknown'} ${into === 1 ? 'piece' : 'pieces'} of ${strip(graph.pools[a.ti].id)} at ${strip(a.target)}`
			};
		});
	});

	/** The earliest depth this pool's pieces join back into, or null. */
	function repeatsInto(pi: number): number | null {
		let d: number | null = null;
		for (const e of edges)
			if (e.back && !e.dead && e.ends[0].startsWith(`${pi}:`)) {
				const t = graph.pools[Number(e.ends[1].split(':')[0])].depth;
				d = d === null ? t : Math.min(d, t);
			}
		return d;
	}

	/** A line between two ports; one that runs backwards loops round. */
	function curve([x1, y1]: [number, number], [x2, y2]: [number, number]): string {
		if (x2 > x1 + 8) {
			const c = Math.max((x2 - x1) / 2, 30);
			return `M${x1},${y1} C${x1 + c},${y1} ${x2 - c},${y2} ${x2},${y2}`;
		}
		// Back to the same or an earlier column: out to the right, round
		// below both ends, and in from the left.
		const low = Math.max(y1, y2) + 18;
		return `M${x1},${y1} C${x1 + 70},${y1} ${x1 + 70},${low} ${x1},${low} L${x2},${low} C${x2 - 70},${low} ${x2 - 70},${y2} ${x2},${y2}`;
	}

	// The piece the pointer is over, or the one clicked, as pool:piece.
	let hover = $state<string | null>(null);
	let picked = $state<string | null>(null);
	const focus = $derived(hover ?? picked);
	const lit = (e: Edge) => focus !== null && (e.ends[0] === focus || e.ends[1] === focus || e.ends[1] === focus.split(':')[0]);
	// The pieces the focused piece joins or is joined by, to light up.
	const related = $derived.by(() => {
		const out = new Set<string>();
		if (focus) for (const e of edges) if (e.ends[0] === focus || e.ends[1] === focus) for (const k of e.ends) out.add(k);
		return out;
	});
	// Lines drawn: piece lines for the piece in focus, dead ends, and
	// each pool's line to its fallback. Lines back to an earlier depth,
	// where growth repeats, are not drawn: the pool's card says it
	// repeats, and the pieces they join still light up.
	const shown = $derived(edges.filter((e) => e.dead || (e.fallback && !e.ends[0].includes(':')) || (!e.back && focus !== null && (e.ends[0] === focus || e.ends[1] === focus))));
	$effect(() => {
		void graph;
		if (picked && !pieceAt(picked)) picked = null;
	});
	function pieceAt(key: string | null): { pool: number; piece: Piece } | null {
		if (!key) return null;
		const [a, b] = key.split(':').map(Number);
		const piece = graph.pools[a]?.pieces[b];
		return piece ? { pool: a, piece } : null;
	}
	const detail = $derived(pieceAt(picked));
	/** The pieces a jigsaw block can join, in words. */
	const joins = (l: Link) => joinsText(graph, detail!.piece, l, strip);

	// Copy the graph as text to share it: a Markdown report (with a
	// Mermaid chart) or just the Mermaid chart.
	let copied = $state<string | null>(null);
	let copyTimer: ReturnType<typeof setTimeout> | undefined;
	async function copy(what: 'markdown' | 'mermaid') {
		const text = what === 'markdown' ? jigsawMarkdown(graph) : jigsawMermaid(graph);
		let ok = false;
		try {
			await navigator.clipboard.writeText(text);
			ok = true;
		} catch {
			// Some hosts block the clipboard API; fall back to a selection.
			const area = document.createElement('textarea');
			area.value = text;
			area.style.position = 'fixed';
			area.style.opacity = '0';
			document.body.append(area);
			area.select();
			ok = document.execCommand('copy');
			area.remove();
		}
		copied = ok ? what : 'failed';
		clearTimeout(copyTimer);
		copyTimer = setTimeout(() => (copied = null), 2000);
	}
	const problems = $derived(graph.problems);
</script>

<div class="flex h-full flex-col text-xs" aria-label="Jigsaw connections">
	<div class="flex flex-wrap items-baseline gap-x-3 gap-y-1 border-b border-line px-3 py-2">
		<h2 class="text-sm">
			<span class="font-semibold">{strip(graph.root.id)}</span>
			<span class="text-muted">· {graph.root.kind === 'structure' ? 'Jigsaw structure' : 'Template pool'}</span>
		</h2>
		{#if graph.size !== undefined}<span class="text-muted" title="The structure's size: how many pieces deep it grows from the start piece">size {graph.size}</span>{/if}
		{#if graph.startName}<span class="text-muted">starts at <span class="font-mono">{strip(graph.startName)}</span></span>{/if}
		<span class="ml-auto flex flex-wrap gap-2" aria-label="Problems">
			{#if problems.deadEnds}<span class="text-red-600 dark:text-red-400">{problems.deadEnds} dead {problems.deadEnds === 1 ? 'end' : 'ends'}</span>{/if}
			{#if problems.unreachable}<span class="text-amber-600 dark:text-amber-400">{problems.unreachable} never joined</span>{/if}
			{#if problems.tooDeep}<span class="text-amber-600 dark:text-amber-400">{problems.tooDeep} past size</span>{/if}
			{#if problems.unknown}<span class="text-muted" title="Pieces that are not this project's structure files, so their jigsaw blocks are not read">{problems.unknown} not read</span>{/if}
			{#if !problems.deadEnds && !problems.unreachable && !problems.tooDeep}<span class="text-muted">No problems found</span>{/if}
		</span>
		<span class="flex gap-1" aria-label="Export">
			<button
				type="button"
				class="rounded border border-line px-1.5 py-0.5 hover:bg-hover"
				title="Copy every pool, piece and jigsaw block (position, facing, size, what it joins, dead ends) as Markdown with a Mermaid chart, to paste into an issue or a chat"
				onclick={() => copy('markdown')}>{copied === 'markdown' ? 'Copied' : 'Copy as Markdown'}</button
			>
			<button
				type="button"
				class="rounded border border-line px-1.5 py-0.5 hover:bg-hover"
				title="Copy the graph as a Mermaid flowchart"
				onclick={() => copy('mermaid')}>{copied === 'mermaid' ? 'Copied' : 'Copy Mermaid'}</button
			>
			{#if copied === 'failed'}<span class="text-red-600 dark:text-red-400">Couldn't copy</span>{/if}
		</span>
	</div>
	<div class="min-h-0 flex-1 overflow-auto bg-canvas">
		<div class="relative" style="width: {layout.width}px; height: {layout.height}px">
			{#each { length: layout.depth + 1 } as _, d (d)}
				{@const deep = graph.size !== undefined && d > graph.size}
				<div class="absolute top-1.5 text-[10px] uppercase tracking-wide {deep ? 'text-amber-600 dark:text-amber-400' : 'text-muted'}" style="left: {LEFT + d * (W + GAP)}px">
					{d === 0 ? 'Start' : `Depth ${d}`}{deep ? ' · past size' : ''}{d === layout.depth && graph.repeatsFrom !== undefined
						? ` · then as depth ${graph.repeatsFrom} again`
						: ''}
				</div>
			{/each}
			<svg class="pointer-events-none absolute inset-0" width={layout.width} height={layout.height} aria-hidden="true">
				<defs>
					<marker id="jig-arrow" viewBox="0 0 8 8" refX="7" refY="4" markerWidth="7" markerHeight="7" orient="auto-start-reverse">
						<path d="M0,0 L8,4 L0,8 z" fill="context-stroke" />
					</marker>
				</defs>
				{#each bundles as e (e.key)}
					<path
						d={curve(e.from, e.to)}
						fill="none"
						marker-end="url(#jig-arrow)"
						stroke-width="1.75"
						stroke-dasharray={e.maybe ? '4 3' : undefined}
						class="stroke-muted"
						opacity={focus ? 0.15 : 0.7}
					><title>{e.title}</title></path>
				{/each}
				{#each shown as e (e.key)}
					<path
						d={curve(e.from, e.to)}
						fill="none"
						marker-end="url(#jig-arrow)"
						stroke-width={lit(e) ? 2 : 1.25}
						stroke-dasharray={e.fallback || e.maybe ? '4 3' : undefined}
						class={e.dead ? 'stroke-red-500' : lit(e) ? 'stroke-accent' : 'stroke-muted'}
						opacity={lit(e) || e.dead ? 1 : focus ? 0.15 : 0.7}
					><title>{e.title}</title></path>
				{/each}
			</svg>
			{#if !focus}
				{#each bundles as e (e.key)}
					<span
						class="pointer-events-none absolute -translate-x-1/2 -translate-y-1/2 rounded border border-line bg-panel px-1 font-mono text-[10px] whitespace-nowrap text-muted"
						style="left: {(e.from[0] + e.to[0]) / 2}px; top: {(e.from[1] + e.to[1]) / 2}px">{e.label}</span
					>
				{/each}
			{/if}
			{#each graph.pools as pool, pi (`${pool.id}@${pool.depth}`)}
				{@const b = layout.boxes[pi]}
				{@const loops = repeatsInto(pi)}
				<section
					class="absolute rounded border bg-panel shadow-sm {pool.tooDeep ? 'border-amber-500/70' : !pool.declared && pool.id.split(':')[0] !== 'minecraft' ? 'border-red-500/80' : 'border-line'}"
					style="left: {b.x}px; top: {b.y}px; width: {W}px; height: {b.h}px"
					aria-label="Pool {pool.id}"
				>
					<header class="flex items-center gap-1.5 border-b border-line px-2" style="height: {HEAD}px">
						<span class="min-w-0 flex-1 truncate font-mono font-semibold" title={pool.id}>{strip(pool.id)}</span>
						{#if pi === 0}<span class="rounded bg-hover px-1 text-[10px] text-muted">start</span>{/if}
						{#if loops !== null}<span
								class="rounded bg-hover px-1 text-[10px] text-muted"
								title="Its jigsaw blocks also draw from pools at depth {loops}, so growth repeats from there; point at a piece to light the pieces it joins">↻ depth {loops}</span
							>{/if}
					</header>
					{#each pool.pieces as p, qi (qi)}
						{@const key = `${pi}:${qi}`}
						<button
							type="button"
							class="flex w-full items-center gap-1.5 px-2 text-left hover:bg-hover {picked === key ? 'bg-selected' : related.has(key) && focus !== key ? 'bg-accent/15' : ''} {p.unreachable ? 'text-amber-700 dark:text-amber-300' : ''}"
							style="height: {ROW}px"
							title={`${p.label ? `${p.label} (${p.id})` : p.id}${p.unreachable ? ': no jigsaw block that draws from this pool targets a name this piece has, so it is never placed' : ''}`}
							onpointerenter={() => (hover = key)}
							onpointerleave={() => (hover = null)}
							onclick={() => (picked = picked === key ? null : key)}
						>
							{#if p.label}<span class="min-w-0 flex-1 truncate">{p.label}</span>{:else}<span class="min-w-0 flex-1 truncate {p.kind === 'single' ? 'font-mono' : 'italic text-muted'}">{p.kind === 'single' ? strip(p.id) : p.id}</span>{/if}
							{#if p.links.some((l) => l.dead)}<span class="text-red-600 dark:text-red-400" title="A jigsaw block here joins nothing">●</span>{/if}
							<span class="shrink-0 text-[10px] text-muted" title="Weight {p.weight}">{pct(p.share)}</span>
							<span class="w-6 shrink-0 text-right text-[10px] text-muted" title={p.jigsaws ? `${p.jigsaws.length} jigsaw ${p.jigsaws.length === 1 ? 'block' : 'blocks'}` : 'Jigsaw blocks unknown: not a structure file of this project'}>{p.jigsaws ? `⧉${p.jigsaws.length}` : '?'}</span>
						</button>
					{:else}
						<p class="flex items-center px-2 {pool.declared || pool.id.split(':')[0] === 'minecraft' ? 'text-muted' : 'text-red-600 dark:text-red-400'}" style="height: {ROW}px">
							{#if pool.declared}No pieces{:else if pool.id.split(':')[0] === 'minecraft'}Vanilla pool, not read{:else}No TemplatePool declares this{/if}
						</p>
					{/each}
					{#if pool.fallback}
						<p class="flex items-center gap-1 px-2 text-[10px] text-muted" style="height: {ROW}px">
							fallback <span class="truncate font-mono" title={pool.fallback}>{strip(pool.fallback)}</span>
						</p>
					{/if}
				</section>
			{/each}
		</div>
	</div>
	<div class="max-h-[40%] shrink-0 overflow-auto border-t border-line px-3 py-2" aria-label="Piece details">
		{#if detail}
			{@const p = detail.piece}
			<h3 class="mb-1 font-semibold">
				{#if p.label}{p.label} <span class="font-normal text-muted">·</span> {/if}<span class="font-mono">{strip(p.id)}</span>
				<span class="font-normal text-muted">· in {strip(graph.pools[detail.pool].id)}, weight {p.weight} ({pct(p.share)}){p.size ? `, size ${p.size.join('×')}` : ''}</span>
			</h3>
			{#if p.unreachable}
				<p class="mb-1 text-amber-700 dark:text-amber-300">No jigsaw block that draws from this pool targets a name this piece has, so it is never placed.</p>
			{/if}
			{#if !p.jigsaws}
				<p class="text-muted">Its jigsaw blocks are unknown: it is not a structure file of this project.</p>
			{:else if !p.jigsaws.length}
				<p class="text-muted">It has no jigsaw blocks, so the structure ends here.</p>
			{:else}
				<table class="w-full border-collapse">
					<thead class="text-left text-muted">
						<tr><th class="py-0.5 font-normal">At</th><th class="font-normal" title="The way it points, then the way its top points">Facing</th><th class="font-normal">Name</th><th class="font-normal">Pool</th><th class="font-normal">Target</th><th class="font-normal">Joins</th></tr>
					</thead>
					<tbody>
						{#each p.jigsaws as j, ji (ji)}
							{@const l = p.links.find((l) => l.jigsaw === ji)}
							<tr class="border-t border-line align-top">
								<td class="py-1 pr-2 whitespace-nowrap">{j.pos.join(' ')}</td>
								<td class="pr-2 whitespace-nowrap" title={j.joint ? `joint: ${j.joint}` : undefined}>{j.orientation?.replace('_', ', top ') ?? '?'}</td>
								<td class="pr-2 font-mono">{strip(j.name)}</td>
								<td class="pr-2 font-mono">{strip(j.pool)}</td>
								<td class="pr-2 font-mono">{strip(j.target)}</td>
								<td class={l?.dead ? 'text-red-600 dark:text-red-400' : ''}>{l ? joins(l) : 'nothing: it only receives'}</td>
							</tr>
						{/each}
					</tbody>
				</table>
			{/if}
		{:else}
			<p class="text-muted">Lines join a pool to the pools it grows into, labelled with the jigsaw name they join at. Point at a piece to see the pieces it can join; click it to list its jigsaw blocks. Red lines join nothing.</p>
		{/if}
	</div>
</div>
