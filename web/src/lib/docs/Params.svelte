<script lang="ts">
	import Doc from './Doc.svelte';
	import TypeText from './TypeText.svelte';
	import type { Param } from './reference.ts';

	// Parameters or fields: name, type, default and documentation.
	let { params, defaults = true }: { params: Param[]; defaults?: boolean } = $props();
</script>

<dl class="params">
	{#each params as p (p.name)}
		<dt><code>{p.name}</code>{#if defaults && p.default}<span class="opt">default <code>{p.default}</code></span>{:else if p.optional}<span class="opt">optional</span>{/if}</dt>
		<dd>
			<TypeText parts={p.typeParts ?? [{ text: p.type }]} />
			{#if p.doc}<p><Doc text={p.doc} /></p>{/if}
		</dd>
	{/each}
</dl>
