<script lang="ts">
	import { resolve } from '$app/paths';
	import Doc from '#lib/docs/Doc.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const r = $derived(data.rule);
	const base = resolve('/').replace(/\/$/, '');
</script>

<svelte:head><title>{r.id} · Lint rules · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1><code>{r.id}</code></h1>
	<p><span class="severity">{r.severity} by default</span></p>
	{#if r.doc}{#each r.doc.split(/\n\n+/) as para, i (i)}<p><Doc text={para} /></p>{/each}{/if}
	<h2 id="configure">Changing it</h2>
	<p>Set the rule's severity in <code>mms.json</code>, to <code>error</code>, <code>warning</code>, <code>info</code>, <code>hint</code> or <code>off</code>:</p>
	<pre><code>{`{\n  "lint": { "${r.id}": "off" }\n}`}</code></pre>
	<p>
		To silence it on one line, write <code>// mms-ignore {r.id}</code> on the line or the line above. See <a href="{base}/docs/guides/lint">Lint rules</a>
		in the guides for more.
	</p>
</article>
