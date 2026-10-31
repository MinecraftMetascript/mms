<script lang="ts">
	import { label } from './common';
	import type { RecipeInfo } from './recipe';
	import Slot from './Slot.svelte';

	let { info }: { info: RecipeInfo } = $props();
	const grid = $derived.by(() => {
		const g = info.grid ?? [];
		return Array.from({ length: 3 }, (_, y) => Array.from({ length: 3 }, (_, x) => g[y]?.[x] ?? null));
	});
	const titles: Record<RecipeInfo['layout'], string> = {
		grid: 'Crafting',
		cooking: 'Cooking',
		stonecutting: 'Stonecutter',
		smithing: 'Smithing',
		brewing: 'Brewing',
		special: 'Special recipe'
	};
</script>

<div class="flex flex-col gap-4">
	<div class="text-xs text-muted">
		{titles[info.layout]} · <span class="font-mono">{info.type}</span>
		{#if info.shapeless}· shapeless{/if}
		{#if info.category}· {info.category}{/if}
		{#if info.group}· group <span class="font-mono">{info.group}</span>{/if}
	</div>
	<div class="flex flex-wrap items-center gap-4">
		{#if info.layout === 'grid'}
			<div class="grid grid-cols-3 gap-1" aria-label="Crafting grid">
				{#each grid as row, y (y)}{#each row as s, x (x)}<Slot slot={s} />{/each}{/each}
			</div>
		{:else if info.layout === 'cooking' || info.layout === 'stonecutting'}
			<Slot slot={info.input ?? null} />
		{:else if info.layout === 'smithing' && info.smithing}
			<div class="flex gap-1">
				<Slot slot={info.smithing.template} />
				<Slot slot={info.smithing.base} />
				<Slot slot={info.smithing.addition} />
			</div>
		{:else if info.layout === 'brewing' && info.brewing}
			<div class="flex flex-col items-center gap-1">
				<Slot slot={info.brewing.reagent} />
				<Slot slot={info.brewing.input} />
				{#if info.notes?.input}<span class="text-xs text-muted">{info.notes.input}</span>{/if}
			</div>
		{/if}
		{#if info.layout !== 'special'}<span class="text-2xl text-muted" aria-hidden="true">→</span>{/if}
		{#if info.result}
			<div class="flex flex-col items-center gap-1">
				<Slot slot={[info.result.id]} count={info.result.count} big />
				<span class="text-xs" aria-label="Result">{info.result.count > 1 ? `${info.result.count} × ` : ''}{label(info.result.id)}</span>
				{#if info.notes?.result}<span class="text-xs text-muted">{info.notes.result}</span>{/if}
			</div>
		{/if}
	</div>
	{#if info.layout === 'cooking'}
		<dl class="grid grid-cols-[auto_1fr] gap-x-4 gap-y-1 text-xs">
			<dt class="text-muted">Time</dt>
			<dd>{info.cookingTime} ticks ({((info.cookingTime ?? 0) / 20).toFixed(1)} s)</dd>
			<dt class="text-muted">Experience</dt>
			<dd>{info.experience}</dd>
		</dl>
	{:else if info.layout === 'special'}
		<p class="text-xs text-muted">This recipe type is handled by the game itself; there is no grid to show.</p>
	{/if}
</div>
