<script lang="ts">
	// The preview panel's page. The extension sends the generated pack and
	// the output at the cursor; the pane and pin button are the web
	// editor's own components from packages/preview.
	import { PinButton, PreviewPane, shows, type OutputFile } from '@mms/preview';
	import PreviewWorker from '@mms/preview/worker?worker&inline';
	import type { ToPanel, ToExtension } from '../src/messages';

	const vscode = acquireVsCodeApi<{ pinned: string | null }>();

	let output = $state<OutputFile[]>([]);
	let selected = $state<string | null>(null);
	let ready = $state(false);
	// A pinned output keeps the panel on one declaration while the cursor
	// moves elsewhere; it still redraws as the pack changes.
	let pinned = $state<string | null>(vscode.getState()?.pinned ?? null);
	const path = $derived(pinned ?? selected);
	const pinnable = $derived(shows(path));

	$effect(() => vscode.setState({ pinned }));

	function togglePin() {
		pinned = pinned ? null : path && shows(path) ? path : null;
	}

	window.addEventListener('message', (e: MessageEvent<ToPanel>) => {
		const m = e.data;
		if (m.type === 'output') {
			output = m.files;
			ready = true;
		} else if (m.type === 'cursor') {
			selected = m.path;
		}
	});
	vscode.postMessage({ type: 'ready' } satisfies ToExtension);

	const source = $derived(output.find((f) => f.path === path));
</script>

<div class="flex h-full flex-col bg-panel text-ink">
	{#if pinned || pinnable}
		<div class="flex items-center border-b border-line text-xs">
			{#if source?.uri && source.range}
				<button
					class="min-w-0 truncate px-3 py-1 text-muted hover:text-ink"
					title="Go to the declaration"
					onclick={() => vscode.postMessage({ type: 'reveal', uri: source.uri!, range: source.range! } satisfies ToExtension)}
					>{source.path}</button
				>
			{/if}
			<PinButton class="ml-auto" {pinned} onclick={togglePin} />
		</div>
	{/if}
	<div class="min-h-0 flex-1">
		{#if ready}
			<PreviewPane {output} {path} {pinned} onunpin={() => (pinned = null)} createWorker={() => new PreviewWorker()} />
		{:else}
			<p class="p-3 text-xs text-muted">Waiting for the MMS language server…</p>
		{/if}
	</div>
</div>
