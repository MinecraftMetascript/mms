<script lang="ts">
	import type { OutputFile } from '../engine/engine.svelte';

	interface Props {
		file: OutputFile | null;
		onsource: (f: OutputFile) => void;
	}
	let { file, onsource }: Props = $props();

	// Light JSON colouring: keys, strings, numbers, literals.
	function highlight(text: string): string {
		const esc = text.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
		return esc.replace(
			/("(?:[^"\\]|\\.)*")(\s*:)?|\b(-?\d+(?:\.\d+)?(?:[eE][+-]?\d+)?)\b|\b(true|false|null)\b/g,
			(m, str, colon, num, lit) => {
				if (str) return colon ? `<span class="j-key">${str}</span>${colon}` : `<span class="j-str">${str}</span>`;
				if (num) return `<span class="j-num">${num}</span>`;
				if (lit) return `<span class="j-lit">${lit}</span>`;
				return m;
			}
		);
	}
</script>

<div class="flex h-full min-h-0 flex-col">
	{#if file}
		<div class="flex items-center gap-2 border-b border-line px-3 py-1.5 text-xs">
			<span class="truncate font-mono text-muted" title={file.path}>{file.path}</span>
			{#if file.uri}
				<button class="ml-auto shrink-0 rounded px-1.5 py-0.5 text-accent hover:bg-hover" onclick={() => onsource(file)}>Show source</button>
			{/if}
		</div>
		<pre class="json min-h-0 flex-1 overflow-auto px-3 py-2 font-mono text-[12.5px] leading-5">{@html highlight(file.text)}</pre>
	{:else}
		<p class="p-4 text-sm text-muted">Select a generated file, or put the cursor in a declaration, to see its JSON.</p>
	{/if}
</div>
