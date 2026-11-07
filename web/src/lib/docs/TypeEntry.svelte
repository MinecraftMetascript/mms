<script lang="ts">
	import Doc from './Doc.svelte';
	import Params from './Params.svelte';
	import TypeText from './TypeText.svelte';
	import { typeId, type TypeEntry } from './reference.ts';

	// One named type: an enum's values, an alias or a record's fields.
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
</div>
