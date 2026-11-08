<script lang="ts">
	// The preview pane both the web editor and VS Code show: an inspector
	// for loot tables, recipes and biomes, a connection graph for jigsaw
	// pools and structures, and a rendered
	// preview for worldgen, with a note when a pinned declaration has left
	// the pack.
	import BookView from './book/BookView.svelte';
	import ChatView from './chat/ChatView.svelte';
	import Inspect from './inspect/Inspect.svelte';
	import { itemContext } from './inspect/items';
	import { inspectable } from './inspect/common';
	import JigsawView from './jigsaw/JigsawView.svelte';
	import JigsawSim from './jigsaw/JigsawSim.svelte';
	import { jigsawGraph } from './jigsaw/graph';
	import { bookAt, itemPath, jigsawable, shows } from './paths';
	import Preview from './preview/Preview.svelte';
	import type { OutputAsset, OutputFile, OutputItem, OutputText } from './types';

	interface Props {
		output: OutputFile[];
		/** The output path to show: the pinned one, else the one at the cursor. */
		path: string | null;
		/** The pinned output path, if any. */
		pinned: string | null;
		onunpin: () => void;
		createWorker: () => Worker;
		/** What to say when there is nothing to preview. */
		hint?: string;
		/** The link that drops a missing pinned output. */
		unpinLabel?: string;
		/** The custom items and books; a book shows as its pages. */
		items?: OutputItem[];
		/** The messages text commands show; one shows as chat, a title or an action bar. */
		texts?: OutputText[];
		/** The book page to open, from 0, such as the one at the cursor. */
		page?: number;
		/** The files the pack copies; structure files carry their jigsaw blocks. */
		assets?: OutputAsset[];
		/**
		 * Image URLs (data: or webview URIs) of the pack's copied textures,
		 * by their path in the pack, so a custom item shows its icon.
		 */
		textures?: Record<string, string>;
		/** The pack's copied structure files as data: URLs, by their path in the pack, for the jigsaw simulation's blocks. */
		structures?: Record<string, string>;
	}
	let { output, path, pinned, onunpin, createWorker, hint, unpinLabel = 'unpin it', items = [], texts = [], page = 0, assets = [], textures = {}, structures = {} }: Props = $props();
	// A jigsaw structure shows its connections or a simulated layout.
	let jigsawTab = $state<'connections' | 'simulation'>('connections');
	const itemInfo = $derived(itemContext(output, items, textures));

	const file = $derived(output.find((f) => f.path === path) ?? null);
	const book = $derived(bookAt(path, items) ? (items.find((it) => itemPath(it.id) === path) ?? null) : null);
	// A pinned output that has left the pack (renamed, deleted or broken).
	const text = $derived(path?.startsWith('text:') ? (texts.find((t) => t.id === path) ?? null) : null);
	const missing = $derived(pinned !== null && !file && !book && !text);
	const jigsaw = $derived(file && jigsawable(file.path) ? jigsawGraph(file.path, output, assets) : null);
</script>

<div class="flex h-full flex-col">
	{#if missing}
		<p class="border-b border-line px-3 py-1.5 text-xs text-muted">
			{shows(pinned, items, texts)?.id ?? pinned?.replace(/^item:/, '')} isn't in the pack right now. It shows again once it builds, or
			<button class="text-accent hover:underline" onclick={onunpin}>{unpinLabel}</button>.
		</p>
	{/if}
	<div class="min-h-0 flex-1">
		{#if book}
			<BookView stack={book.text} {output} {page} />
		{:else if text}
			<ChatView {text} {output} />
		{:else if jigsaw && jigsaw.root.kind === 'structure' && file}
			<div class="flex h-full flex-col">
				<div class="flex gap-1 border-b border-line px-2 pt-1 text-xs" role="tablist" aria-label="Jigsaw preview">
					{#each [['connections', 'Connections'], ['simulation', 'Simulation']] as [id, label] (id)}
						<button
							type="button"
							role="tab"
							aria-selected={jigsawTab === id}
							class="-mb-px border-b-2 px-2 py-1 {jigsawTab === id ? 'border-accent font-medium' : 'border-transparent text-muted hover:text-ink'}"
							onclick={() => (jigsawTab = id as typeof jigsawTab)}>{label}</button
						>
					{/each}
				</div>
				<div class="min-h-0 flex-1">
					{#if jigsawTab === 'simulation'}
						<JigsawSim path={file.path} {output} {assets} {structures} />
					{:else}
						<JigsawView graph={jigsaw} />
					{/if}
				</div>
			</div>
		{:else if jigsaw}
			<JigsawView graph={jigsaw} />
		{:else if file && jigsawable(file.path)}
			<p class="p-3 text-xs text-muted">This structure isn't a jigsaw one, so it has no pieces to connect.</p>
		{:else if file && inspectable(file.path)}
			<Inspect {file} items={itemInfo} />
		{:else}
			<Preview {output} target={path} {createWorker} {hint} />
		{/if}
	</div>
</div>
