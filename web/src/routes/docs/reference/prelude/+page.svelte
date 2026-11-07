<script lang="ts">
	import { resolve } from '$app/paths';
	import PreludeEntry from '#lib/docs/PreludeEntry.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const base = resolve('/').replace(/\/$/, '');
</script>

<svelte:head><title>Templates and forms · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>Templates and forms</h1>
	<p>
		Every project has the prelude: <a href="{base}/docs/reference/language/reuse#templates">Templates</a>, Types and
		<a href="{base}/docs/reference/kinds#blueprints">Blueprint kinds</a> written in MMS, and the built-in templates the compiler supplies. Call a
		template like any helper.
	</p>

	<h2 id="templates">Templates</h2>
	{#each data.templates as d (d.name)}<PreludeEntry decl={d} />{/each}

	<h2 id="types">Types</h2>
	{#each data.types as d (d.name)}<PreludeEntry decl={d} />{/each}

	<h2 id="forms">Forms of expression</h2>
	<p>
		Operators and literals are forms of expression. What one means depends on the type expected where it is written, so <code>1..3</code> is a uniform
		number in a loot pool's rolls and bounds in a predicate. This is the <a href="{base}/docs/reference/language/values#shorthand">shorthand</a> the
		language builds in.
	</p>
	{#each data.forms as d (d.name)}<PreludeEntry decl={d} />{/each}
</article>
