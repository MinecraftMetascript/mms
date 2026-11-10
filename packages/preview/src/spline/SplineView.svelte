<script lang="ts">
	// The spline editor: the curve of the density function spline at the
	// cursor, worked out the way the game does. Dragging a point changes
	// its location and value, and dragging the handle of the selected
	// point changes its derivative; each drag is one text edit on the
	// number literals, so undo works as usual. A number written as a
	// const or an expression shows but cannot be dragged. A point whose
	// value is another spline is hollow: clicking it opens that spline,
	// and the sliders under the graph pick the nested splines' coordinates.
	import {
		addPoint,
		evaluate,
		formatNumber,
		nestedCoordinates,
		pointValue,
		removePoint,
		sameSpline,
		setNumber,
		splineAt,
		type Coordinates,
		type SplineAt,
		type SplinePointData,
		type TextEdit
	} from './spline';

	let { data, onedit }: { data: SplineAt; onedit?: (edits: TextEdit[]) => void } = $props();

	// The open spline, as point indices from the root. It follows the
	// cursor when the cursor moves to another spline, and stays where the
	// breadcrumb or a click put it while the source is edited.
	let path = $state<number[]>([]);
	let lastFocus = '';
	let lastRoot: SplineAt['root'] | undefined;
	$effect.pre(() => {
		const key = JSON.stringify(data.focus);
		if (key !== lastFocus || !sameSpline(lastRoot, data.root)) path = [...data.focus];
		lastFocus = key;
		lastRoot = data.root;
	});
	const open = $derived(splineAt(data.root, path));
	const spline = $derived(open.spline);
	// The crumbs back out: the root, then each point opened.
	const crumbs = $derived.by(() => {
		const out: { label: string; path: number[] }[] = [{ label: data.root.coordinate || 'spline', path: [] }];
		let s = data.root;
		open.path.forEach((i, k) => {
			s = s.points[i].spline!;
			out.push({ label: `${s.coordinate || 'spline'} (point ${i + 1})`, path: open.path.slice(0, k + 1) });
		});
		return out;
	});

	// The values the nested splines' coordinates take, by coordinate.
	let coords = $state<Coordinates>({});
	const nested = $derived(nestedCoordinates(spline));
	const coordinateValue = (name: string, min: number, max: number) => coords[name] ?? Math.min(Math.max(0, min), max);
	const at = $derived(Object.fromEntries(nested.map((c) => [c.name, coordinateValue(c.name, c.min, c.max)])));

	let selected = $state<number | null>(null);
	$effect.pre(() => {
		void spline;
		if (selected !== null && selected >= spline.points.length) selected = null;
	});

	// A drag in progress: the point or its derivative handle, and the
	// numbers it has reached. The graph keeps its scale until it ends.
	type Drag = { kind: 'point' | 'handle'; index: number; location: number; value: number; derivative: number; moved: boolean };
	let drag = $state<Drag | null>(null);
	// What a finished drag wrote, shown until the source catches up.
	let pending = $state<Drag | null>(null);
	$effect.pre(() => {
		void data;
		pending = null;
	});
	const shown = $derived(drag ?? pending);

	const loc = (i: number) => (shown && shown.index === i ? shown.location : (spline.points[i].location.value ?? 0));
	const val = (i: number) => (shown && shown.index === i ? shown.value : pointValue(spline.points[i], at));
	const der = (i: number) => (shown && shown.index === i ? shown.derivative : (spline.points[i].derivative.value ?? 0));
	// The spline with a dragged point's numbers in place.
	const live = $derived.by(() => {
		if (!shown) return spline;
		const points = spline.points.map((p, i): SplinePointData => {
			if (i !== shown.index) return p;
			const v = p.spline ? undefined : { ...p.value, value: shown.value };
			return { ...p, location: { ...p.location, value: shown.location }, derivative: { ...p.derivative, value: shown.derivative }, value: v };
		});
		return { ...spline, points };
	});

	// Layout, in px.
	let width = $state(400);
	let height = $state(260);
	const M = { left: 44, right: 12, top: 12, bottom: 28 };

	function domainOf() {
		const ls = spline.points.map((p) => p.location.value).filter((v): v is number => v !== null);
		let x0 = ls.length ? Math.min(...ls) : -1;
		let x1 = ls.length ? Math.max(...ls) : 1;
		const pad = Math.max((x1 - x0) * 0.1, 0.2);
		x0 -= pad;
		x1 += pad;
		const ys: number[] = [];
		for (let k = 0; k <= 120; k++) ys.push(evaluate(spline, x0 + ((x1 - x0) * k) / 120, at));
		spline.points.forEach((p) => ys.push(pointValue(p, at)));
		let y0 = Math.min(...ys);
		let y1 = Math.max(...ys);
		const ypad = Math.max((y1 - y0) * 0.12, 0.1);
		y0 -= ypad;
		y1 += ypad;
		return { x0, x1, y0, y1 };
	}
	let frozen = $state<ReturnType<typeof domainOf> | null>(null);
	const domain = $derived(frozen ?? domainOf());
	const sx = (x: number) => M.left + ((x - domain.x0) / (domain.x1 - domain.x0)) * (width - M.left - M.right);
	const sy = (y: number) => M.top + (1 - (y - domain.y0) / (domain.y1 - domain.y0)) * (height - M.top - M.bottom);
	const ix = (px: number) => domain.x0 + ((px - M.left) / (width - M.left - M.right)) * (domain.x1 - domain.x0);
	const iy = (py: number) => domain.y0 + (1 - (py - M.top) / (height - M.top - M.bottom)) * (domain.y1 - domain.y0);

	const curve = $derived.by(() => {
		const n = Math.max(2, Math.round(width / 2));
		const pts: string[] = [];
		for (let k = 0; k <= n; k++) {
			const x = domain.x0 + ((domain.x1 - domain.x0) * k) / n;
			pts.push(`${sx(x).toFixed(1)},${sy(evaluate(live, x, at)).toFixed(1)}`);
		}
		return 'M' + pts.join('L');
	});

	function ticks(a: number, b: number, count: number): number[] {
		const step0 = (b - a) / count;
		const mag = 10 ** Math.floor(Math.log10(step0));
		const step = [1, 2, 5, 10].map((m) => m * mag).find((s) => s >= step0) ?? 10 * mag;
		const out: number[] = [];
		for (let t = Math.ceil(a / step) * step; t <= b + 1e-9; t += step) out.push(Math.abs(t) < step / 1e6 ? 0 : t);
		return out;
	}
	const xticks = $derived(ticks(domain.x0, domain.x1, Math.max(2, Math.floor(width / 70))));
	const yticks = $derived(ticks(domain.y0, domain.y1, Math.max(2, Math.floor(height / 40))));

	// The derivative handle sits this many px from its point, along the tangent.
	const HANDLE = 36;
	function handleAt(i: number): [number, number] {
		const px = sx(loc(i));
		const py = sy(val(i));
		// The tangent's direction on screen.
		const dx = sx(loc(i) + 1) - sx(loc(i));
		const dy = sy(val(i) + der(i)) - sy(val(i));
		const len = Math.hypot(dx, dy) || 1;
		return [px + (dx / len) * HANDLE, py + (dy / len) * HANDLE];
	}

	const canMove = (p: SplinePointData) => !!p.location.range || (!p.spline && !!p.value?.range);

	let svg = $state<SVGSVGElement>();
	function local(e: PointerEvent): [number, number] {
		const r = svg!.getBoundingClientRect();
		return [e.clientX - r.left, e.clientY - r.top];
	}

	function startPoint(e: PointerEvent, i: number) {
		e.stopPropagation();
		selected = i;
		frozen = domain;
		drag = { kind: 'point', index: i, location: loc(i), value: val(i), derivative: der(i), moved: false };
		svg!.setPointerCapture(e.pointerId);
	}
	function startHandle(e: PointerEvent, i: number) {
		e.stopPropagation();
		frozen = domain;
		drag = { kind: 'handle', index: i, location: loc(i), value: val(i), derivative: der(i), moved: false };
		svg!.setPointerCapture(e.pointerId);
	}
	// Shift snaps to a grid of 0.05.
	const snap = (v: number, e: PointerEvent) => (e.shiftKey ? Math.round(v / 0.05) * 0.05 : v);

	let hover = $state<number | null>(null);
	function move(e: PointerEvent) {
		const [px, py] = local(e);
		hover = px >= M.left && px <= width - M.right ? ix(px) : null;
		if (!drag) return;
		const p = spline.points[drag.index];
		if (drag.kind === 'point') {
			if (!canMove(p)) return;
			const location = p.location.range ? snap(ix(px), e) : drag.location;
			const value = !p.spline && p.value?.range ? snap(iy(py), e) : drag.value;
			drag = { ...drag, location, value, moved: true };
		} else {
			const dx = ix(px) - drag.location;
			if (Math.abs(sx(drag.location + dx) - sx(drag.location)) < 4) return;
			const d = (iy(py) - drag.value) / dx;
			drag = { ...drag, derivative: e.shiftKey ? Math.round(d * 20) / 20 : d, moved: true };
		}
	}
	function end() {
		const d = drag;
		drag = null;
		frozen = null;
		if (!d) return;
		const p = spline.points[d.index];
		if (!d.moved) {
			// A click on a nested spline's point opens it.
			if (d.kind === 'point' && p.spline) {
				path = [...open.path, d.index];
				selected = null;
			}
			return;
		}
		const edits: TextEdit[] = [];
		const push = (x: TextEdit | null) => x && edits.push(x);
		if (d.kind === 'point') {
			if (d.location !== (p.location.value ?? 0)) push(setNumber(p.location, d.location));
			if (!p.spline && p.value && d.value !== (p.value.value ?? 0)) push(setNumber(p.value, d.value));
		} else {
			push(setNumber(p.derivative, d.derivative));
		}
		if (edits.length) {
			pending = d;
			onedit?.(edits);
		}
	}

	function add(e: MouseEvent) {
		if (!svg || !spline.editable || (e.target as Element).tagName === 'circle') return;
		const r = svg.getBoundingClientRect();
		const x = ix(e.clientX - r.left);
		const y = iy(e.clientY - r.top);
		const edit = addPoint(spline, x, y);
		if (edit) onedit?.([edit]);
	}
	function remove(i: number) {
		const edit = removePoint(spline.points[i]);
		if (edit) {
			selected = null;
			onedit?.([edit]);
		}
	}
	function key(e: KeyboardEvent) {
		if ((e.key === 'Delete' || e.key === 'Backspace') && selected !== null) {
			e.preventDefault();
			remove(selected);
		} else if (e.key === 'Escape') selected = null;
	}
	// The selected point's numbers, typed in.
	function typed(i: number, which: 'location' | 'value' | 'derivative', text: string) {
		const v = Number(text);
		const p = spline.points[i];
		const n = which === 'value' ? p.value : p[which];
		if (!n || !Number.isFinite(v)) return;
		const edit = setNumber(n, v, 6);
		if (edit) onedit?.([edit]);
	}
	const sel = $derived(selected !== null ? spline.points[selected] : null);
	const fields = $derived(
		sel
			? ([
					['location', 'Location', sel.location],
					['value', 'Value', sel.value],
					['derivative', 'Derivative', sel.derivative]
				] as const)
			: []
	);
</script>

<div class="flex h-full flex-col text-xs" data-testid="spline-view">
	<div class="flex flex-wrap items-center gap-1 border-b border-line px-3 py-1.5">
		{#each crumbs as c, k (k)}
			{#if k > 0}<span class="text-muted">›</span>{/if}
			{#if k === crumbs.length - 1}
				<span class="font-mono text-ink">{c.label}</span>
			{:else}
				<button class="font-mono text-accent hover:underline" onclick={() => ((path = c.path), (selected = null))}>{c.label}</button>
			{/if}
		{/each}
		{#if hover !== null}
			<span class="ml-auto font-mono text-muted" data-testid="spline-readout">{formatNumber(hover)} → {formatNumber(evaluate(live, hover, at))}</span>
		{/if}
	</div>
	<!-- svelte-ignore a11y_no_noninteractive_tabindex, a11y_no_noninteractive_element_interactions -->
	<div class="relative min-h-0 flex-1 outline-none" bind:clientWidth={width} bind:clientHeight={height} tabindex="0" role="application" aria-label="Spline graph" onkeydown={key}>
		<svg
			bind:this={svg}
			{width}
			{height}
			class="absolute inset-0 touch-none select-none"
			role="presentation"
			onpointermove={move}
			onpointerup={end}
			onpointercancel={end}
			onpointerleave={() => (hover = null)}
			ondblclick={add}
		>
			{#each xticks as t (t)}
				<line x1={sx(t)} x2={sx(t)} y1={M.top} y2={height - M.bottom} class="stroke-line" stroke-width={t === 0 ? 1.5 : 1} />
				<text x={sx(t)} y={height - M.bottom + 14} text-anchor="middle" class="fill-muted text-[10px]">{formatNumber(t, 2)}</text>
			{/each}
			{#each yticks as t (t)}
				<line x1={M.left} x2={width - M.right} y1={sy(t)} y2={sy(t)} class="stroke-line" stroke-width={t === 0 ? 1.5 : 1} />
				<text x={M.left - 6} y={sy(t) + 3} text-anchor="end" class="fill-muted text-[10px]">{formatNumber(t, 2)}</text>
			{/each}
			{#if hover !== null}
				<line x1={sx(hover)} x2={sx(hover)} y1={M.top} y2={height - M.bottom} class="stroke-muted" stroke-dasharray="3 3" />
			{/if}
			<path d={curve} fill="none" class="stroke-accent" stroke-width="2" />
			{#if selected !== null && selected < spline.points.length && spline.points[selected].derivative.range}
				{@const [hx, hy] = handleAt(selected)}
				{@const [ox, oy] = [2 * sx(loc(selected)) - hx, 2 * sy(val(selected)) - hy]}
				<line x1={ox} y1={oy} x2={hx} y2={hy} class="stroke-ink" stroke-opacity="0.6" />
				<circle
					cx={hx}
					cy={hy}
					r="5"
					class="cursor-pointer fill-panel stroke-ink"
					role="slider"
					aria-label="Derivative"
					aria-valuenow={der(selected)}
					tabindex="-1"
					data-testid="spline-handle"
					onpointerdown={(e) => startHandle(e, selected!)}
				/>
			{/if}
			{#each spline.points as p, i (i)}
				{@const movable = canMove(p)}
				<circle
					cx={sx(loc(i))}
					cy={sy(val(i))}
					r={selected === i ? 7 : 6}
					class="{p.spline ? 'fill-panel' : 'fill-accent'} stroke-accent {movable || p.spline ? 'cursor-pointer' : 'cursor-not-allowed'}"
					stroke-width={selected === i ? 3 : 2}
					stroke-dasharray={movable ? undefined : '2 2'}
					role="button"
					tabindex="-1"
					aria-label="Point {i + 1}"
					data-testid="spline-point"
					onpointerdown={(e) => startPoint(e, i)}
				>
					<title
						>{p.spline ? `A spline of ${p.spline.coordinate}: click to open it` : `Point ${i + 1}`}{movable ? '' : ' (written as an expression, so it cannot be dragged)'}</title
					>
				</circle>
			{/each}
		</svg>
		{#if spline.points.length === 0}
			<p class="absolute inset-x-0 top-1/2 text-center text-muted">This spline has no points. Double-click to add one.</p>
		{/if}
	</div>
	<div class="flex flex-col gap-1.5 border-t border-line px-3 py-2">
		{#if sel && selected !== null}
			{@const i = selected}
			<div class="flex flex-wrap items-center gap-x-3 gap-y-1" data-testid="spline-selected">
				<span class="font-medium">Point {i + 1}</span>
				{#each fields as [which, label, n] (which)}
					<label class="flex items-center gap-1">
						<span class="text-muted">{label}</span>
						{#if which === 'value' && sel.spline}
							<button class="font-mono text-accent hover:underline" onclick={() => ((path = [...open.path, i]), (selected = null))}>spline of {sel.spline.coordinate}</button>
						{:else}
							<input
								class="w-20 rounded border border-line bg-bg px-1 font-mono disabled:opacity-60"
								value={n?.value === null || n?.value === undefined ? '' : formatNumber(n.value, 6)}
								disabled={!n?.range}
								title={n?.range ? undefined : 'Written as an expression; edit it in the source'}
								onchange={(e) => typed(i, which, (e.target as HTMLInputElement).value)}
							/>
						{/if}
					</label>
				{/each}
				{#if sel.remove}
					<button class="ml-auto text-muted hover:text-ink" onclick={() => remove(i)}>Delete point</button>
				{/if}
			</div>
		{:else}
			<p class="text-muted">
				{spline.points.length} point{spline.points.length === 1 ? '' : 's'} on {spline.coordinate || 'the coordinate'}. Drag a point to move it, select one to drag its slope{spline.editable
					? ', and double-click the graph to add a point'
					: ''}. Shift snaps.
			</p>
		{/if}
		{#each nested as c (c.name)}
			<label class="flex items-center gap-2">
				<span class="w-40 truncate font-mono text-muted" title="The value of {c.name} for the nested splines">{c.name}</span>
				<input
					type="range"
					class="min-w-0 flex-1"
					min={c.min}
					max={c.max}
					step={(c.max - c.min) / 200 || 0.01}
					value={at[c.name]}
					oninput={(e) => (coords = { ...coords, [c.name]: Number((e.target as HTMLInputElement).value) })}
				/>
				<span class="w-12 text-right font-mono">{formatNumber(at[c.name], 2)}</span>
			</label>
		{/each}
	</div>
</div>
