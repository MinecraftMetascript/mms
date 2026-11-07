<script lang="ts">
	import { goto } from '$app/navigation';
	import Prose from '#lib/docs/Prose.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();

	// An old guide URL: keep the section anchor on the way to the new one.
	$effect(() => {
		if (data.redirect) goto(data.redirect + location.hash, { replaceState: true });
	});
</script>

<svelte:head>
	{#if data.redirect}
		<!-- Without JavaScript only: a refresh timer, once started, is not
		     stopped by client-side navigation. The effect above keeps the
		     section anchor. -->
		<noscript><meta http-equiv="refresh" content="0; url={data.redirect}" /></noscript>
		<link rel="canonical" href={data.redirect} />
	{:else}
		<title>{data.title} · MMS docs</title>
	{/if}
</svelte:head>

{#if data.redirect}
	<p>This page moved to <a class="text-accent" href={data.redirect}>{data.redirect}</a>.</p>
{:else}
	<Prose html={data.html!} />
	{#if data.prev || data.next}
		<nav class="mt-10 flex justify-between gap-4 border-t border-line pt-4 text-sm" aria-label="Previous and next">
			{#if data.prev}<a class="hover:underline" href={data.prev.href}><span class="block text-xs text-muted">Previous</span>{data.prev.title}</a>{:else}<span></span>{/if}
			{#if data.next}<a class="text-right hover:underline" href={data.next.href}><span class="block text-xs text-muted">Next</span>{data.next.title}</a>{/if}
		</nav>
	{/if}
{/if}
