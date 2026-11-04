<script lang="ts">
	// The preview panel's page. The extension sends the generated pack and
	// the declaration to show (picked with a Preview CodeLens); the pane is
	// the web editor's own component from packages/preview. The panel stays
	// on that declaration while you edit elsewhere.
	import { itemPath, PreviewPane, type OutputFile, type OutputItem } from '@mms/preview';
	import PreviewWorker from '@mms/preview/worker?worker&inline';
	import type { OutputAsset, ToPanel, ToExtension } from '../src/messages';

	// The state is what VS Code hands back when it restores the panel.
	const vscode = acquireVsCodeApi<{ path: string | null; page?: number }>();

	let output = $state<OutputFile[]>([]);
	let items = $state<OutputItem[]>([]);
	let assets = $state<OutputAsset[]>([]);
	let textures = $state<Record<string, string>>({});
	let path = $state<string | null>(vscode.getState()?.path ?? null);
	// The book page to open, from 0.
	let page = $state(vscode.getState()?.page ?? 0);
	let ready = $state(false);

	$effect(() => vscode.setState({ path, page }));

	window.addEventListener('message', (e: MessageEvent<ToPanel>) => {
		const m = e.data;
		if (m.type === 'output') {
			output = m.files;
			items = m.items;
			assets = m.assets;
			textures = m.textures ?? {};
			ready = true;
		} else if (m.type === 'target') {
			path = m.path;
			page = m.page ?? 0;
		}
	});
	vscode.postMessage({ type: 'ready' } satisfies ToExtension);

	// The declaration shown: a file, or a custom item or book.
	const source = $derived.by(() => {
		const f = output.find((f) => f.path === path);
		if (f) return { label: f.path, uri: f.uri, range: f.range };
		const it = items.find((it) => itemPath(it.id) === path);
		return it ? { label: it.id, uri: it.uri, range: it.range } : undefined;
	});
</script>

<div class="flex h-full flex-col bg-panel text-ink">
	{#if source?.uri && source.range}
		<div class="flex items-center border-b border-line text-xs">
			<button
				class="min-w-0 truncate px-3 py-1 text-muted hover:text-ink"
				title="Go to the declaration"
				onclick={() => vscode.postMessage({ type: 'reveal', uri: source.uri!, range: source.range! } satisfies ToExtension)}
				>{source.label}</button
			>
		</div>
	{/if}
	<div class="min-h-0 flex-1">
		{#if ready}
			<PreviewPane
				{output}
				{items}
				{assets}
				{textures}
				{page}
				{path}
				pinned={path}
				unpinLabel="clear it"
				onunpin={() => vscode.postMessage({ type: 'clear' } satisfies ToExtension)}
				createWorker={() => new PreviewWorker()}
				hint="Click Preview above a noise, density function, noise settings, surface rule, dimension, loot table, recipe, biome, template pool, jigsaw structure or book to show it here."
			/>
		{:else}
			<p class="p-3 text-xs text-muted">Waiting for the MMS language server…</p>
		{/if}
	</div>
</div>
