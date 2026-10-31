<script lang="ts">
	import raw from '#lib/docs/reference.json';
	import { filterGroups, groupId, groupTitle, inlineCode, type Reference } from '#lib/docs/reference.ts';

	const ref = raw as Reference;
	const resources = ref.kinds.filter((k) => k.folder);
	const own = ref.kinds.filter((k) => !k.folder);
	const total = ref.groups.reduce((n, g) => n + g.helpers.length, 0);

	let q = $state('');
	const shown = $derived(filterGroups(ref.groups, q));
</script>

<svelte:head><title>Reference · MMS docs</title></svelte:head>

<article class="prose-mms">
	<h1>Reference</h1>
	<p>
		Every declaration kind and helper for Minecraft {ref.version}, generated from the same schema the checker uses. In the editor, completion
		and hover show the same information where you type.
	</p>

	<h2 id="kinds">Kinds</h2>
	<p>MMS adds these kinds for functions and code:</p>
	<table>
		<thead><tr><th>Kind</th><th>Does</th></tr></thead>
		<tbody>
			{#each own as k (k.name)}
				<tr><td><code>{k.name}</code></td><td>{k.doc}</td></tr>
			{/each}
		</tbody>
	</table>
	<p>Each of these makes one resource in the data pack folder shown. Tags of other registries follow the same pattern, such as <code>DamageTypeTag</code>.</p>
	<table>
		<thead><tr><th>Kind</th><th>Folder</th></tr></thead>
		<tbody>
			{#each resources as k (k.name)}
				<tr><td><code>{k.name}</code></td><td><code>data/&lt;namespace&gt;/{k.folder}/</code></td></tr>
			{/each}
		</tbody>
	</table>

	<h2 id="helpers">Helpers</h2>
	<p>
		A helper builds one typed object. <code>Name(a, b)</code> fills the fields in the order shown, and <code>Name(field = value)</code> names them.
		Fields marked <code>?</code> are optional. There are {total} helpers, grouped by what they make.
	</p>
	<input
		class="mb-4 w-full rounded-md border border-line bg-panel px-3 py-2 text-sm outline-none focus:border-accent"
		type="search"
		placeholder="Filter helpers, such as noise or loot_function"
		aria-label="Filter helpers"
		bind:value={q}
	/>
	{#if !shown.length}
		<p class="text-muted">No helper matches “{q}”.</p>
	{/if}
	{#each shown as g (g.dispatcher)}
		<section aria-label={groupTitle(g.dispatcher)}>
			<h3 id={groupId(g.dispatcher)}>
				{groupTitle(g.dispatcher)}
				{#if g.dispatcher}<span class="font-normal text-muted"><code>{g.dispatcher}</code></span>{/if}
			</h3>
			<dl>
				{#each g.helpers as h (h.name)}
					<dt id="{groupId(g.dispatcher)}-{h.name}"><code>{h.signature}</code>{#if h.case}<span class="ml-2 text-xs text-muted">{h.case}</span>{/if}</dt>
					{#if h.params?.some((p) => p.doc)}
						<dd>
							<ul>
								{#each h.params.filter((p) => p.doc) as p (p.name)}
									<li><code>{p.name}</code>: {#each inlineCode(p.doc ?? '') as part, i (i)}{#if part.code}<code>{part.text}</code>{:else}{part.text}{/if}{/each}</li>
								{/each}
							</ul>
						</dd>
					{/if}
				{/each}
			</dl>
		</section>
	{/each}
</article>
