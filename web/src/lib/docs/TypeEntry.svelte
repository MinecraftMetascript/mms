<script lang="ts">
	import Doc from './Doc.svelte';
	import Params from './Params.svelte';
	import TypeText from './TypeText.svelte';
	import { typeId, type TypeEntry } from './reference.ts';

	// One named type: what it is, values that fit, an enum's values or a
	// record's fields, and where it is used.
	let { type: t }: { type: TypeEntry } = $props();
</script>

<div class="entry" id={typeId(t.id)} tabindex="-1">
	<div class="entry-head">
		<a class="entry-name" href="#{typeId(t.id)}">{t.name}</a>
		{#if t.values?.length}
			<span class="entry-meta">is one of</span>
		{:else if t.def?.length}
			<span class="entry-meta">is <TypeText parts={t.def} /></span>
		{/if}
	</div>
	{#if t.examples?.length}
		<p class="examples">For example {#each t.examples as x, i (x)}{#if i > 0}{i === t.examples.length - 1 ? ' or ' : ', '}{/if}<code>{x}</code>{/each}.</p>
	{/if}
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
	{#if t.fields?.length}<Params params={t.fields} defaults={false} />{/if}
	{#if t.includes?.length}
		<p class="used-by">Also has the fields of {#each t.includes as u, i (u.text)}{#if i > 0}, {/if}<TypeText parts={[u]} />{/each}.</p>
	{/if}
	{#if t.usedBy?.length}
		<p class="used-by">Used by {#each t.usedBy as u, i (u.text)}{#if i > 0}, {/if}<TypeText parts={[u]} />{/each}.</p>
	{/if}
</div>
