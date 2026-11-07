<script lang="ts">
	import { resolve } from '$app/paths';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const base = resolve('/').replace(/\/$/, '');
	const c = $derived(data.counts);
</script>

<svelte:head><title>MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>MMS docs</h1>
	<p>
		MMS is a language for Minecraft data packs and resource packs. The <strong>Docs</strong> teach it, from a first project to how it
		works underneath. The <strong>Reference</strong> is for looking things up.
	</p>
	<ul class="cards">
		<li><a href="{base}/docs/start/getting-started"><strong>Start here</strong><span>Install MMS and build a first data pack.</span></a></li>
		<li><a href={resolve('/docs/reference')}><strong>Reference</strong><span>{c.kinds} kinds, {c.helpers} helpers and {c.rules} lint rules, for Minecraft {c.version}.</span></a></li>
		<li><a href={resolve('/')}><strong>Open the editor</strong><span>Write MMS in the browser, with previews.</span></a></li>
	</ul>
	{#each data.groups as g (g.dir)}
		<h2 id={g.dir}>{g.title}</h2>
		<ul class="cards">
			{#each g.pages as p (p.path)}
				<li><a href="{base}/docs/{p.path}"><strong>{p.title}</strong>{#if p.summary}<span>{p.summary.length > 140 ? p.summary.slice(0, 137).replace(/\s+\S*$/, '') + '…' : p.summary}</span>{/if}</a></li>
			{/each}
		</ul>
	{/each}
</article>
