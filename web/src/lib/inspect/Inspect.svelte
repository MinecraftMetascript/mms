<script lang="ts">
	import type { OutputFile } from '../engine/engine.svelte';
	import { inspectable, label, type Json } from './common';
	import { analyseLoot } from './loot';
	import { analyseRecipe } from './recipe';
	import { analyseBiome } from './biome';
	import LootView from './LootView.svelte';
	import RecipeView from './RecipeView.svelte';
	import BiomeView from './BiomeView.svelte';

	let { file }: { file: OutputFile } = $props();
	const info = $derived(inspectable(file.path));
	const json = $derived.by((): Json | undefined => {
		try {
			return JSON.parse(file.text);
		} catch {
			return undefined;
		}
	});
	const titles = { loot_table: 'Loot table', recipe: 'Recipe', biome: 'Biome' };
</script>

<div class="h-full overflow-auto p-3" aria-label="Inspector">
	{#if info && json !== undefined}
		<h2 class="mb-3 text-sm">
			<span class="font-semibold">{label(info.id)}</span>
			<span class="text-xs text-muted">· {titles[info.kind]} <span class="font-mono">{info.id}</span></span>
		</h2>
		{#if info.kind === 'loot_table'}
			<LootView pools={analyseLoot(json)} />
		{:else if info.kind === 'recipe'}
			{@const r = analyseRecipe(json)}
			{#if r}<RecipeView info={r} />{/if}
		{:else}
			{@const b = analyseBiome(json)}
			{#if b}<BiomeView info={b} />{/if}
		{/if}
	{/if}
</div>
