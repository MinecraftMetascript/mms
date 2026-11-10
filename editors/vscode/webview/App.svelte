<script lang="ts">
	// The preview panel's page. The extension sends the generated pack and
	// the declaration to show (picked with a Preview CodeLens); the pane is
	// the web editor's own component from packages/preview. The panel stays
	// on that declaration while you edit elsewhere.
	import { itemPath, PreviewPane, type OutputFile, type OutputItem } from '@mms/preview';
	import PreviewWorker from '@mms/preview/worker?worker&inline';
	import type { OutputAsset, OutputText, SplineAt, ToPanel, ToExtension } from '../src/messages';

	// The state is what VS Code hands back when it restores the panel.
	const vscode = acquireVsCodeApi<{ path: string | null; page?: number }>();

	let output = $state<OutputFile[]>([]);
	let items = $state<OutputItem[]>([]);
	let texts = $state<OutputText[]>([]);
	let assets = $state<OutputAsset[]>([]);
	let textures = $state<Record<string, string>>({});
	let structures = $state<Record<string, string>>({});
	let path = $state<string | null>(vscode.getState()?.path ?? null);
	// The book page to open, from 0.
	let page = $state(vscode.getState()?.page ?? 0);
	let ready = $state(false);
	// The spline at the cursor in the last .mms editor, for the spline editor.
	let spline = $state<SplineAt | null>(null);
	let splineUri: string | null = null;
	// The document version the spline's ranges are in, sent back with each edit.
	let splineVersion = -1;

	$effect(() => vscode.setState({ path, page }));

	window.addEventListener('message', (e: MessageEvent<ToPanel>) => {
		const m = e.data;
		if (m.type === 'output') {
			output = m.files;
			items = m.items;
			texts = m.texts ?? [];
			assets = m.assets;
			textures = m.textures ?? {};
			structures = m.structures ?? {};
			ready = true;
		} else if (m.type === 'target') {
			path = m.path;
			page = m.page ?? 0;
		} else if (m.type === 'spline') {
			spline = m.spline;
			splineUri = m.uri;
			splineVersion = m.version;
		}
	});
	vscode.postMessage({ type: 'ready' } satisfies ToExtension);

	// The declaration shown: a file, or a custom item or book.
	const source = $derived.by(() => {
		const f = output.find((f) => f.path === path);
		if (f) return { label: f.path, uri: f.uri, range: f.range };
		const it = items.find((it) => itemPath(it.id) === path);
		if (it) return { label: it.id, uri: it.uri, range: it.range };
		const t = texts.find((t) => t.id === path);
		return t ? { label: t.kind === 'chat' ? 'chat message' : t.kind, uri: t.uri, range: t.range } : undefined;
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
				{texts}
				{assets}
				{textures}
				{structures}
				{page}
				{path}
				pinned={path}
				unpinLabel="clear it"
				onunpin={() => vscode.postMessage({ type: 'clear' } satisfies ToExtension)}
				createWorker={() => new PreviewWorker()}
				{spline}
				onsplineedit={(edits) => splineUri && vscode.postMessage({ type: 'splineEdit', uri: splineUri, version: splineVersion, edits } satisfies ToExtension)}
				hint="Click Preview above a noise, density function, noise settings, surface rule, dimension, loot table, recipe, biome, template pool, jigsaw structure or book to show it here, or Edit spline above a spline to graph it."
			/>
		{:else}
			<p class="p-3 text-xs text-muted">Waiting for the MMS language server…</p>
		{/if}
	</div>
</div>
