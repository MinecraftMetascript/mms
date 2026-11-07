<script lang="ts">
	import { resolve } from '$app/paths';
	import Doc from '#lib/docs/Doc.svelte';
	import Params from '#lib/docs/Params.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const k = $derived(data.kind);
	const base = resolve('/').replace(/\/$/, '');
	const params = $derived((k.fields ?? []).map((f) => ({ ...f, type: f.type ?? 'any' })));
	const required = $derived((k.fields ?? []).filter((f) => !f.optional));
	const shape = $derived(`${k.name} name {\n${required.map((f) => `  ${f.name} = ...`).join('\n')}${required.length ? '\n' : ''}}`);
</script>

<svelte:head><title>{k.name} · Declaration kinds · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1><code>{k.name}</code></h1>
	{#if k.doc}{#each k.doc.split(/\n\n+/) as para, i (i)}<p><Doc text={para} /></p>{/each}{/if}
	<pre class="shape"><code>{shape}</code></pre>
	<p>
		<code>{k.name}</code> is a <a href="{base}/docs/reference/language/blueprints">Blueprint</a> in the prelude: it is written in MMS, and every project has it.
	</p>

	<h2 id="fields">Fields</h2>
	{#if params.length}<Params {params} />{:else}<p>It has no fields.</p>{/if}

	{#if k.uses?.length}
		<h2 id="as">What its name stands for</h2>
		<p>Where one of these types is expected, the declaration's name means:</p>
		<dl class="params">
			{#each k.uses as u (u.type)}
				<dt><code>{u.type}</code></dt>
				<dd>{#if u.doc}<p><Doc text={u.doc} /></p>{/if}</dd>
			{/each}
		</dl>
	{/if}
</article>
