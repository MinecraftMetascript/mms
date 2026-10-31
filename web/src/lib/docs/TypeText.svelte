<script lang="ts">
	import { linkId, type TypePart } from './reference.ts';

	// A readable type, with named types as links to their entries.
	let { parts, onjump }: { parts: TypePart[]; onjump: (e: MouseEvent, id: string) => void } = $props();
</script>

<code class="type">{#each parts as p, i (i)}{#if linkId(p.link)}<a href="#{linkId(p.link)}" onclick={(e) => onjump(e, linkId(p.link))}>{p.text}</a>{:else}{p.text}{/if}{/each}</code>
