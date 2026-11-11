<script lang="ts">
	import { resolve } from '$app/paths';
	import { refURL, typeAreaTitle } from '#lib/docs/reference.ts';

	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const base = resolve('/').replace(/\/$/, '');
	let q = $state('');
	const shown = $derived.by(() => {
		const s = q.trim().toLowerCase();
		return s ? data.names.filter((t) => t.name.toLowerCase().includes(s)).slice(0, 200) : [];
	});
</script>

<svelte:head><title>Types · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>Types</h1>
	<p>
		Every type from the game's schema that MMS code can name, such as <code>ClimateParameter</code> in <code>const c: ClimateParameter = [-1, 1]</code> or a Blueprint field. Each entry says what
		the type is, gives values that fit, and lists where it is used. Hovering a type in the editor shows the same.
	</p>
	<input
		class="w-full rounded-md border border-line bg-panel px-3 py-2 text-sm outline-none focus:border-accent"
		type="search"
		placeholder="Find a type by name, such as ClimateParameter"
		aria-label="Find a type"
		bind:value={q}
	/>
	{#if q.trim()}
		<ul>
			{#each shown as t (t.id)}
				<li><a href={refURL.type(base, t.id, t.area)}><code>{t.name}</code></a> <span class="text-muted">· {typeAreaTitle(t.area)}</span></li>
			{:else}
				<li class="text-muted">Nothing matches “{q}”.</li>
			{/each}
		</ul>
	{:else}
		<ul>
			{#each data.areas as a (a.area)}
				<li><a href={refURL.types(base, a.area)}>{a.title}</a> <span class="text-muted">· {a.count} types</span></li>
			{/each}
		</ul>
	{/if}
</article>
