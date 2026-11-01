<script lang="ts">
	// The preview pane both the web editor and VS Code show: an inspector
	// for loot tables, recipes and biomes, and a rendered preview for
	// worldgen, with a note when a pinned declaration has left the pack.
	import Inspect from './inspect/Inspect.svelte';
	import { inspectable } from './inspect/common';
	import { shows } from './pin';
	import Preview from './preview/Preview.svelte';
	import type { OutputFile } from './types';

	interface Props {
		output: OutputFile[];
		/** The output path to show: the pinned one, else the one at the cursor. */
		path: string | null;
		/** The pinned output path, if any. */
		pinned: string | null;
		onunpin: () => void;
		createWorker: () => Worker;
	}
	let { output, path, pinned, onunpin, createWorker }: Props = $props();

	const file = $derived(output.find((f) => f.path === path) ?? null);
	// A pinned output that has left the pack (renamed, deleted or broken).
	const missing = $derived(pinned !== null && !file);
</script>

<div class="flex h-full flex-col">
	{#if missing}
		<p class="border-b border-line px-3 py-1.5 text-xs text-muted">
			{shows(pinned)?.id} isn't in the pack right now. It shows again once it builds, or
			<button class="text-accent hover:underline" onclick={onunpin}>unpin it</button>.
		</p>
	{/if}
	<div class="min-h-0 flex-1">
		{#if file && inspectable(file.path)}
			<Inspect {file} />
		{:else}
			<Preview {output} target={path} {createWorker} />
		{/if}
	</div>
</div>
