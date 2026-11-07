<script lang="ts">
	import { resolve } from '$app/paths';
	import { linkHref, type TypePart } from './reference.ts';

	// A readable type, with named types as links to their entries.
	let { parts }: { parts: TypePart[] } = $props();
	const base = resolve('/').replace(/\/$/, '');
</script>

<code class="type">{#each parts as p, i (i)}{#if linkHref(base, p.link)}<a href={linkHref(base, p.link)}>{p.text}</a>{:else}{p.text}{/if}{/each}</code>
