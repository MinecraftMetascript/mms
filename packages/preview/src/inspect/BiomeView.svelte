<script lang="ts">
	import { label, pct } from './common';
	import type { BiomeInfo } from './biome';

	let { info }: { info: BiomeInfo } = $props();
	const num = (n: number | null) => (n === null ? '?' : String(n));
</script>

<div class="flex flex-col gap-4 text-xs">
	<section>
		<h3 class="mb-1 font-semibold">Climate</h3>
		<dl class="grid grid-cols-[auto_1fr] gap-x-4 gap-y-1">
			<dt class="text-muted">Temperature</dt>
			<dd>{num(info.temperature)}{info.frozen ? ' (frozen)' : ''}</dd>
			<dt class="text-muted">Downfall</dt>
			<dd>{num(info.downfall)}</dd>
			<dt class="text-muted">Precipitation</dt>
			<dd>{info.precipitation}</dd>
		</dl>
	</section>
	{#if info.colors.length}
		<section>
			<h3 class="mb-1 font-semibold">Colours</h3>
			<ul class="grid grid-cols-[repeat(auto-fill,minmax(8rem,1fr))] gap-2">
				{#each info.colors as c (c.name)}
					<li class="flex items-center gap-2">
						<span class="h-6 w-6 shrink-0 rounded border border-line" style="background:{c.color}"></span>
						<span>
							<span class="block capitalize">{c.name}</span>
							<span class="font-mono text-muted" title={c.derived ? 'From the colormap; approximate' : undefined}>{c.color}{c.derived ? ' ≈' : ''}</span>
						</span>
					</li>
				{/each}
			</ul>
		</section>
	{/if}
	{#if info.spawns.length}
		<section>
			<h3 class="mb-1 font-semibold">Spawns</h3>
			{#each info.spawns as cat (cat.category)}
				<h4 class="mt-2 text-muted capitalize">{cat.category}</h4>
				<ul>
					{#each cat.mobs as m, i (i)}
						<li class="flex items-center gap-2 py-0.5">
							<span class="w-28 truncate" title={m.mob}>{label(m.mob)}</span>
							<span class="h-1.5 flex-1 rounded bg-hover"><span class="block h-full rounded bg-accent" style="width:{m.chance * 100}%"></span></span>
							<span class="w-10 text-right">{pct(m.chance)}</span>
							<span class="w-10 text-right text-muted" title="Group size">×{m.count}</span>
						</li>
					{/each}
				</ul>
			{/each}
		</section>
	{/if}
	{#if info.features.length}
		<section>
			<h3 class="mb-1 font-semibold">Features</h3>
			<ul class="grid grid-cols-2 gap-x-4">
				{#each info.features as f (f.step)}
					<li class="flex justify-between {f.count ? '' : 'text-muted'}"><span class="capitalize">{f.step}</span><span>{f.count}</span></li>
				{/each}
			</ul>
			<p class="mt-1 text-muted">{info.carvers} {info.carvers === 1 ? 'carver' : 'carvers'}</p>
		</section>
	{/if}
	{#if info.music.length}
		<section>
			<h3 class="mb-1 font-semibold">Music</h3>
			<p class="font-mono">{info.music.join(', ')}</p>
		</section>
	{/if}
</div>
