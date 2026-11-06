<script lang="ts">
	import { resolve } from '$app/paths';
	import { groupTitle, refURL } from '#lib/docs/reference.ts';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();

	const base = resolve('/').replace(/\/$/, '');
</script>

<svelte:head><title>Helpers · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>Helpers</h1>
	<p>
		A helper builds one typed object. <code>Name(a, b)</code> fills the parameters in the order listed, and <code>Name(a = 1)</code> names
		them. There are {data.counts.helpers} helpers, grouped by what they make.
	</p>
	<ul class="cards">
		{#each data.groups as g (g.dispatcher)}
			<li>
				<a href={refURL.group(base, g.dispatcher)}><strong>{groupTitle(g.dispatcher)}</strong><span>{g.count} {g.count === 1 ? 'helper' : 'helpers'}{#if g.dispatcher}, <code>{g.dispatcher}</code>{/if}</span></a>
			</li>
		{/each}
	</ul>
</article>
