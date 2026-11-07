<script lang="ts">
	import Doc from './Doc.svelte';
	import Params from './Params.svelte';
	import type { PreludeDecl } from './reference.ts';

	// One Template or Type of the prelude: how it is written, its doc, its
	// parameters or fields, and for a form of expression what it means
	// where each type is expected.
	let { decl: d }: { decl: PreludeDecl } = $props();
	const id = $derived(`prelude-${d.name}`);
	const params = $derived((d.fields ?? []).map((f) => ({ ...f, type: f.type ?? 'any' })));
</script>

<div class="entry" {id} tabindex="-1">
	<div class="entry-head">
		<a class="entry-name" href="#{id}">{d.name}</a>
		{#if d.kind === 'Template' && d.builtin}<span class="entry-meta">built in</span>{/if}
	</div>
	{#if d.doc}{#each d.doc.split(/\n\n+/) as para, i (i)}<p class="entry-doc"><Doc text={para} /></p>{/each}{/if}
	{#if !(d.kind === 'Type' && d.builtin)}<pre class="shape"><code>{d.kind === 'Type' ? 'Type ' : ''}{d.signature}</code></pre>{/if}
	{#if d.uses?.length}
		<dl class="params">
			{#each d.uses as u (u.type)}
				<dt>where <code>{u.type}</code> goes</dt>
				<dd>{#if u.doc}<p><Doc text={u.doc} /></p>{/if}</dd>
			{/each}
		</dl>
	{:else if params.length && !(d.kind === 'Type' && d.builtin)}
		<Params {params} />
	{/if}
</div>
