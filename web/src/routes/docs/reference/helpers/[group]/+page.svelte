<script lang="ts">
	import HelperEntry from '#lib/docs/HelperEntry.svelte';
	import { groupTitle } from '#lib/docs/reference.ts';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const g = $derived(data.group);
</script>

<svelte:head><title>{groupTitle(g.dispatcher)} · Helpers · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>{groupTitle(g.dispatcher)}</h1>
	<p>
		{#if g.dispatcher}
			Helpers that build a <code>{g.dispatcher}</code> object, one for each of its types. Where a field takes one, write the helper's call; its
			name is the type it makes.
		{:else}
			Helpers for objects that are not chosen by a <code>type</code> field. Each builds the object it is named after.
		{/if}
		Click a type to see what it accepts.
	</p>
	{#each g.helpers as h (h.name)}
		<h2 class="sr-only" id="h-{h.name}" data-toc={h.name}>{h.name}</h2>
		<HelperEntry dispatcher={g.dispatcher} helper={h} />
	{/each}
</article>
