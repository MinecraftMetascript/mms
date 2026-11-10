<script lang="ts">
	import Prose from '#lib/docs/Prose.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head><title>{data.title} · MMS docs</title></svelte:head>

<Prose html={data.html} />
{#if data.prev || data.next}
	<nav class="mt-10 flex justify-between gap-4 border-t border-line pt-4 text-sm" aria-label="Previous and next">
		{#if data.prev}<a class="hover:underline" href={data.prev.href}><span class="block text-xs text-muted">Previous</span>{data.prev.title}</a>{:else}<span></span>{/if}
		{#if data.next}<a class="text-right hover:underline" href={data.next.href}><span class="block text-xs text-muted">Next</span>{data.next.title}</a>{/if}
	</nav>
{/if}
