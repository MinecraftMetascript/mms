<script lang="ts">
	import Doc from './Doc.svelte';
	import Params from './Params.svelte';
	import { callShape, helperId, type Helper } from './reference.ts';

	// One helper: its call shape, documentation, parameters and example.
	let { dispatcher, helper: h }: { dispatcher: string; helper: Helper } = $props();
	const id = $derived(helperId(dispatcher, h.name));
</script>

<div class="entry" {id} tabindex="-1">
	<div class="entry-head">
		<a class="entry-name" href="#{id}">{h.name}</a>
		{#if h.case}<span class="entry-meta">makes <code>{h.case}</code></span>{/if}
	</div>
	{#if h.doc}{#each h.doc.split(/\n\n+/) as para, i (i)}<p class="entry-doc"><Doc text={para} /></p>{/each}{/if}
	<pre class="shape"><code>{callShape(h)}</code></pre>
	{#if h.params?.length}<Params params={h.params} />{/if}
	{#if h.example}<pre class="example"><code>{h.example}</code></pre>{/if}
</div>
