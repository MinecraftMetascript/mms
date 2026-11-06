<script lang="ts">
	import TypeEntry from '#lib/docs/TypeEntry.svelte';
	import { filterTypes } from '#lib/docs/reference.ts';

	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	let q = $state('');
	const shown = $derived(filterTypes(data.types, q));
</script>

<svelte:head><title>Types · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>Types</h1>
	<p>Named types that helper parameters refer to: enums, aliases and records. Types made by a helper group link straight to that group instead.</p>
	<input
		class="w-full rounded-md border border-line bg-panel px-3 py-2 text-sm outline-none focus:border-accent"
		type="search"
		placeholder="Filter by name, such as Axis"
		aria-label="Filter types"
		bind:value={q}
	/>
	{#each shown as t (t.id)}
		<TypeEntry type={t} />
	{:else}
		<p class="text-muted">Nothing matches “{q}”.</p>
	{/each}
</article>
