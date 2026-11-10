<script lang="ts">
	import { resolve } from '$app/paths';
	import TypeEntry from '#lib/docs/TypeEntry.svelte';
	import { filterTypes } from '#lib/docs/reference.ts';

	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	let q = $state('');
	const shown = $derived(filterTypes(data.types, q));
</script>

<svelte:head><title>{data.title} types · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>{data.title} types</h1>
	<p>
		Named types from this part of the game's schema. A field or typed const of one takes any of the values shown; types made by helpers link to them. See <a href={resolve('/docs/reference/types')}
			>all types</a
		>.
	</p>
	<input
		class="w-full rounded-md border border-line bg-panel px-3 py-2 text-sm outline-none focus:border-accent"
		type="search"
		placeholder="Filter by name"
		aria-label="Filter types"
		bind:value={q}
	/>
	{#each shown as t (t.id)}
		<TypeEntry type={t} />
	{:else}
		<p class="text-muted">Nothing matches “{q}”.</p>
	{/each}
</article>
