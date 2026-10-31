<script lang="ts">
	import { tick } from 'svelte';
	import raw from '#lib/docs/reference.json';
	import Doc from '#lib/docs/Doc.svelte';
	import TypeText from '#lib/docs/TypeText.svelte';
	import { callShape, filterGroups, filterTypes, groupId, groupTitle, helperId, typeId, type Reference } from '#lib/docs/reference.ts';

	const ref = raw as Reference;
	const resources = ref.kinds.filter((k) => k.folder);
	const own = ref.kinds.filter((k) => !k.folder);
	const total = ref.groups.reduce((n, g) => n + g.helpers.length, 0);

	let q = $state('');
	const shown = $derived(filterGroups(ref.groups, q));
	const shownTypes = $derived(filterTypes(ref.types, q));

	// Following a link clears the filter first, so the target is on the page.
	async function jump(e: MouseEvent, id: string) {
		if (e.metaKey || e.ctrlKey || e.shiftKey || e.button !== 0) return;
		e.preventDefault();
		if (q) {
			q = '';
			await tick();
		}
		history.pushState(null, '', '#' + id);
		document.getElementById(id)?.scrollIntoView();
		document.getElementById(id)?.focus({ preventScroll: true });
	}
</script>

<svelte:head><title>Reference · MMS docs</title></svelte:head>

<article class="prose-mms reference">
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
	<p>Each of these makes one resource in the folder shown: in the data pack, or for <code>assets/</code> folders in the resource pack. Tags of other registries follow the same pattern, such as <code>DamageTypeTag</code>.</p>
	<table>
		<thead><tr><th>Kind</th><th>Folder</th></tr></thead>
		<tbody>
			{#each resources as k (k.name)}
				<tr><td><code>{k.name}</code></td><td><code>{k.resources ? 'assets' : 'data'}/&lt;namespace&gt;/{k.folder}/</code></td></tr>
			{/each}
		</tbody>
	</table>

	<h2 id="helpers">Helpers</h2>
	<p>
		A helper builds one typed object. <code>Name(a, b)</code> fills the parameters in the order listed, and <code>Name(a = 1)</code> names
		them. There are {total} helpers, grouped by what they make. Click a type to jump to it.
	</p>
	<div class="filter">
		<input
			class="w-full rounded-md border border-line bg-panel px-3 py-2 text-sm outline-none focus:border-accent"
			type="search"
			placeholder="Filter by helper, parameter or type, such as noise or BlockState"
			aria-label="Filter helpers"
			bind:value={q}
		/>
	</div>
	<nav class="toc" aria-label="Helper groups">
		{#each shown as g (g.dispatcher)}
			<a href="#{groupId(g.dispatcher)}">{groupTitle(g.dispatcher)} <span>{g.helpers.length}</span></a>
		{/each}
		{#if shownTypes.length}<a href="#types">Types <span>{shownTypes.length}</span></a>{/if}
	</nav>
	{#if !shown.length && !shownTypes.length}
		<p class="text-muted">Nothing matches “{q}”.</p>
	{/if}

	{#each shown as g (g.dispatcher)}
		<section aria-label={groupTitle(g.dispatcher)}>
			<h3 id={groupId(g.dispatcher)} data-toc={groupTitle(g.dispatcher)} tabindex="-1">
				<a href="#{groupId(g.dispatcher)}">{groupTitle(g.dispatcher)}</a>
				{#if g.dispatcher}<code class="ml-1 text-xs font-normal">{g.dispatcher}</code>{/if}
			</h3>
			{#each g.helpers as h (h.name)}
				{@const id = helperId(g.dispatcher, h.name)}
				<div class="entry" {id} tabindex="-1">
					<div class="entry-head">
						<a class="entry-name" href="#{id}">{h.name}</a>
						{#if h.case}<span class="entry-meta">makes <code>{h.case}</code></span>{/if}
					</div>
					<pre class="shape"><code>{callShape(h)}</code></pre>
					{#if h.params?.length}
						<dl class="params">
							{#each h.params as p (p.name)}
								<dt><code>{p.name}</code>{#if p.optional}<span class="opt">optional</span>{/if}</dt>
								<dd>
									<TypeText parts={p.typeParts ?? [{ text: p.type }]} onjump={jump} />
									{#if p.doc}<p><Doc text={p.doc} /></p>{/if}
								</dd>
							{/each}
						</dl>
					{/if}
				</div>
			{/each}
		</section>
	{/each}

	{#if shownTypes.length}
		<h2 id="types">Types</h2>
		<p>Named types that parameters refer to. Types made by a helper group link straight to that group instead.</p>
		{#each shownTypes as t (t.id)}
			<div class="entry" id={typeId(t.id)} tabindex="-1">
				<div class="entry-head">
					<a class="entry-name" href="#{typeId(t.id)}">{t.name}</a>
					{#if t.values?.length}
						<span class="entry-meta">is one of</span>
					{:else if t.def?.length}
						<span class="entry-meta">is <TypeText parts={t.def} onjump={jump} /></span>
					{/if}
				</div>
				{#if t.values?.length && !t.values.some((v) => v.doc)}
					<p class="values">{#each t.values as v (v.value)}<code>"{v.value}"</code> {/each}</p>
				{:else if t.values?.length}
					<dl class="params">
						{#each t.values as v (v.value)}
							<dt><code>"{v.value}"</code></dt>
							<dd>{#if v.doc}<Doc text={v.doc} />{/if}</dd>
						{/each}
					</dl>
				{/if}
				{#if t.fields?.length}
					<dl class="params">
						{#each t.fields as p (p.name)}
							<dt><code>{p.name}</code>{#if p.optional}<span class="opt">optional</span>{/if}</dt>
							<dd>
								<TypeText parts={p.typeParts ?? [{ text: p.type }]} onjump={jump} />
								{#if p.doc}<p><Doc text={p.doc} /></p>{/if}
							</dd>
						{/each}
					</dl>
				{/if}
			</div>
		{/each}
	{/if}
</article>

<style>
	.reference :global(.filter) {
		position: sticky;
		top: 2.75rem;
		z-index: 5;
		padding: 0.5rem 0;
		background: var(--c-bg);
	}
	.toc {
		display: flex;
		flex-wrap: wrap;
		gap: 0.25rem 0.4rem;
		margin: 0.25rem 0 1.25rem;
		font-size: 13px;
	}
	.toc a {
		border: 1px solid var(--c-line);
		border-radius: 999px;
		padding: 0.1rem 0.6rem;
		color: var(--c-ink);
		text-decoration: none;
	}
	.toc a:hover {
		background: var(--c-hover);
	}
	.toc span {
		color: var(--c-muted);
	}
	h3 {
		position: sticky;
		top: 6.1rem;
		z-index: 4;
		background: var(--c-bg);
		padding: 0.35rem 0;
		margin-top: 2rem;
		border-bottom: 1px solid var(--c-line);
	}
	.entry {
		border: 1px solid var(--c-line);
		border-radius: 8px;
		background: var(--c-panel);
		padding: 0.75rem 1rem;
		margin: 0.75rem 0;
		scroll-margin-top: 9.5rem;
		outline: none;
	}
	.entry:target,
	.entry:focus {
		border-color: var(--c-accent);
		box-shadow: 0 0 0 1px var(--c-accent);
	}
	h3:focus {
		outline: none;
	}
	.entry-head {
		display: flex;
		flex-wrap: wrap;
		align-items: baseline;
		gap: 0.25rem 0.75rem;
	}
	.entry-name {
		font-family: var(--font-mono);
		font-weight: 650;
		font-size: 1rem;
		color: var(--c-ink);
		text-decoration: none;
	}
	.entry-name:hover {
		text-decoration: underline;
	}
	.entry-meta {
		color: var(--c-muted);
		font-size: 13px;
	}
	.shape {
		margin: 0.4rem 0 0.6rem;
		padding: 0.3rem 0.6rem;
		font-size: 12.5px;
		white-space: pre-wrap;
	}
	.params {
		display: grid;
		grid-template-columns: minmax(7rem, max-content) 1fr;
		gap: 0;
		margin: 0;
		font-size: 14px;
	}
	.params dt,
	.params dd {
		padding: 0.4rem 0;
		border-top: 1px solid var(--c-line);
		margin: 0;
	}
	.params dt {
		padding-right: 1rem;
	}
	.params dd :global(p) {
		margin: 0.15rem 0 0;
		color: var(--c-muted);
		line-height: 1.5;
	}
	.params code {
		background: none;
		padding: 0;
	}
	.opt {
		display: block;
		font-size: 11px;
		color: var(--c-muted);
	}
	:global(.reference code.type) {
		background: none;
		padding: 0;
		font-size: 13px;
	}
	:global(.reference code.type a) {
		text-decoration: none;
		border-bottom: 1px dotted currentColor;
	}
	:global(.reference code.type a:hover) {
		border-bottom-style: solid;
	}
	.values {
		margin: 0.5rem 0 0;
		display: flex;
		flex-wrap: wrap;
		gap: 0.3rem;
	}
	.params dd :global(p) {
		white-space: pre-line;
	}
	@media (max-width: 640px) {
		.params {
			grid-template-columns: 1fr;
		}
		.params dd {
			border-top: 0;
			padding-top: 0;
		}
		.opt {
			display: inline;
			margin-left: 0.5rem;
		}
		h3 {
			top: 6.1rem;
		}
	}
</style>
