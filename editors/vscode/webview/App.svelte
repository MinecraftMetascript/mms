<script lang="ts">
	// The preview panel's page. The extension sends the generated pack and
	// the declaration to show (picked with a Preview CodeLens); the pane is
	// the web editor's own component from packages/preview. The panel stays
	// on that declaration while you edit elsewhere.
	import { PreviewPane, type OutputFile } from '@mms/preview';
	import PreviewWorker from '@mms/preview/worker?worker&inline';
	import type { ToPanel, ToExtension } from '../src/messages';

	// The state is what VS Code hands back when it restores the panel.
	const vscode = acquireVsCodeApi<{ path: string | null }>();

	let output = $state<OutputFile[]>([]);
	let path = $state<string | null>(vscode.getState()?.path ?? null);
	let ready = $state(false);

	$effect(() => vscode.setState({ path }));

	window.addEventListener('message', (e: MessageEvent<ToPanel>) => {
		const m = e.data;
		if (m.type === 'output') {
			output = m.files;
			ready = true;
		} else if (m.type === 'target') {
			path = m.path;
		}
	});
	vscode.postMessage({ type: 'ready' } satisfies ToExtension);

	const source = $derived(output.find((f) => f.path === path));
</script>

<div class="flex h-full flex-col bg-panel text-ink">
	{#if source?.uri && source.range}
		<div class="flex items-center border-b border-line text-xs">
			<button
				class="min-w-0 truncate px-3 py-1 text-muted hover:text-ink"
				title="Go to the declaration"
				onclick={() => vscode.postMessage({ type: 'reveal', uri: source.uri!, range: source.range! } satisfies ToExtension)}
				>{source.path}</button
			>
		</div>
	{/if}
	<div class="min-h-0 flex-1">
		{#if ready}
			<PreviewPane
				{output}
				{path}
				pinned={path}
				unpinLabel="clear it"
				onunpin={() => vscode.postMessage({ type: 'clear' } satisfies ToExtension)}
				createWorker={() => new PreviewWorker()}
				hint="Click Preview above a noise, density function, noise settings, surface rule, dimension, loot table, recipe or biome to show it here."
			/>
		{:else}
			<p class="p-3 text-xs text-muted">Waiting for the MMS language server…</p>
		{/if}
	</div>
</div>
