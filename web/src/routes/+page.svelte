<script lang="ts">
	import { onMount, tick } from 'svelte';
	import { asset, resolve } from '$app/paths';
	import { LSPPlugin } from '@codemirror/lsp-client';
	import { EditorSelection } from '@codemirror/state';
	import { EditorView } from '@codemirror/view';
	import Editor from '#lib/editor/Editor.svelte';
	import Sidebar from '#lib/components/Sidebar.svelte';
	import JsonView from '#lib/components/JsonView.svelte';
	import Preview from '#lib/preview/Preview.svelte';
	import { Engine, type Diagnostic, type DocSymbol, type OutputFile } from '#lib/engine/engine.svelte.ts';
	import { Project, pathOf, uriOf, starter } from '#lib/project/project.svelte.ts';
	import { examples } from '#lib/project/examples.ts';
	import { download, isDataPack, packZip, projectZip, readProjectZip, readZip } from '#lib/project/zip.ts';
	import PasteJson from '#lib/components/PasteJson.svelte';
	import { previewable } from '#lib/preview/protocol.ts';
	import Inspect from '#lib/inspect/Inspect.svelte';
	import { inspectable } from '#lib/inspect/common.ts';

	const project = new Project();
	let engine = $state<Engine | null>(null);
	let view: EditorView | null = null;
	let viewWaiters: ((v: EditorView) => void)[] = [];
	let loaded = $state(false);
	let selectedOutput = $state<string | null>(null);
	let rightTab = $state<'preview' | 'json'>('preview');
	let cursor = $state<{ uri: string; line: number } | null>(null);
	let theme = $state<'system' | 'light' | 'dark'>('system');
	let menu = $state(false);
	let menuRoot: HTMLDivElement | undefined = $state();
	let menuButton: HTMLButtonElement | undefined = $state();

	// Close the Project menu on a click outside it, or on Escape.
	function onWindowPointerDown(e: PointerEvent) {
		if (menu && menuRoot && !menuRoot.contains(e.target as Node)) menu = false;
	}
	function onWindowKeydown(e: KeyboardEvent) {
		if (menu && e.key === 'Escape') {
			menu = false;
			menuButton?.focus();
		}
	}
	let fileInput: HTMLInputElement;
	let pasteJson = $state<ReturnType<typeof PasteJson> | null>(null);
	let editor = $state<ReturnType<typeof Editor> | null>(null);

	onMount(() => {
		try {
			theme = (localStorage.getItem('mms-theme') as typeof theme) ?? 'system';
		} catch {
			/* storage unavailable */
		}
		const e = new Engine(asset('mms-worker.js'), {
			onText: (uri, text) => project.setText(pathOf(uri), text),
			display: async (uri) => {
				const path = pathOf(uri);
				if (!project.files.some((f) => f.path === path)) return null;
				if (project.active === path && view) return view;
				const v = new Promise<EditorView>((resolve) => viewWaiters.push(resolve));
				project.active = path;
				return v;
			}
		});
		engine = e;
		project.load().then(() => {
			syncEngine();
			loaded = true;
		});
		return () => e.dispose();
	});

	// Keep the server's file set and config in step with the project.
	function syncEngine() {
		if (!engine) return;
		engine.setConfig($state.snapshot(project.config));
		engine.workspace.setFiles(project.files.map((f) => ({ uri: uriOf(f.path), text: f.text })));
		engine.scheduleRefresh();
	}
	$effect(() => {
		// Track the set of paths (not their text, which the editor syncs).
		void project.files.map((f) => f.path).join('\n');
		void project.config.namespace;
		void project.config.description;
		if (loaded) syncEngine();
	});

	$effect(() => {
		const t = theme;
		if (t === 'system') document.documentElement.removeAttribute('data-theme');
		else document.documentElement.setAttribute('data-theme', t);
		try {
			localStorage.setItem('mms-theme', t);
		} catch {
			/* ignore */
		}
	});

	function onview(v: EditorView) {
		view = v;
		for (const w of viewWaiters.splice(0)) w(v);
		trackCursor(v);
	}

	// Follow the cursor: the declaration under it picks the preview and JSON.
	function trackCursor(v: EditorView) {
		const line = v.state.doc.lineAt(v.state.selection.main.head).number - 1;
		cursor = { uri: uriOf(project.active), line };
	}
	const cursorExt = EditorView.updateListener.of((u) => {
		if (u.selectionSet || u.docChanged) trackCursor(u.view);
	});

	const within = (r: Diagnostic['range'], line: number) => r.start.line <= line && line <= r.end.line;
	const cursorOutput = $derived.by(() => {
		if (!engine || !cursor) return null;
		const c = cursor;
		return engine.output.find((f) => f.uri === c.uri && f.range && within(f.range, c.line)) ?? null;
	});
	$effect(() => {
		if (cursorOutput) selectedOutput = cursorOutput.path;
	});
	const selectedFile = $derived(engine?.output.find((f) => f.path === selectedOutput) ?? null);

	async function reveal(uri: string, range: Diagnostic['range']) {
		const path = pathOf(uri);
		if (project.active !== path) {
			const v = new Promise<EditorView>((resolve) => viewWaiters.push(resolve));
			project.active = path;
			await v;
		}
		await tick();
		const v = view;
		const plugin = v && LSPPlugin.get(v);
		if (!v || !plugin) return;
		const doc = v.state.doc;
		const from = Math.min(plugin.fromPosition(range.start, doc), doc.length);
		const to = Math.min(plugin.fromPosition(range.end, doc), doc.length);
		v.dispatch({ selection: EditorSelection.range(from, to), effects: EditorView.scrollIntoView(from, { y: 'center' }) });
		v.focus();
	}

	function selectOutput(path: string) {
		selectedOutput = path;
		rightTab = previewable(path) || inspectable(path) ? rightTab : 'json';
	}

	function openExample(id: string) {
		const ex = examples.find((e) => e.id === id);
		menu = false;
		if (!ex) return;
		if (!confirm(`Open the ${ex.id} example? It replaces the current project; export it first to keep a copy.`)) return;
		project.replace(structuredClone(ex.data));
	}
	function newProject() {
		menu = false;
		if (confirm('Start a new project? It replaces the current one; export it first to keep a copy.')) project.replace(structuredClone(starter));
	}
	async function importZip(e: Event) {
		const file = (e.target as HTMLInputElement).files?.[0];
		(e.target as HTMLInputElement).value = '';
		if (!file) return;
		try {
			const entries = readZip(new Uint8Array(await file.arrayBuffer()));
			if (isDataPack(entries)) {
				if (!engine) return;
				const res = await engine.importPack([...entries].map(([path, text]) => ({ path, text })));
				if (!res.files.length) throw new Error('the pack has no resources MMS can import');
				const skipped = res.skipped.length ? ` ${res.skipped.length} file(s) cannot be imported and will be left out.` : '';
				if (!confirm(`Import ${res.count} resource(s) from the data pack ${file.name}?${skipped} It replaces the current project.`)) return;
				const files = res.files.map((f) => ({ path: f.path, text: f.text }));
				project.replace({
					name: file.name.replace(/\.zip$/i, ''),
					config: { namespace: res.namespace, description: res.description || 'Imported with MMS' },
					files,
					active: files[0].path
				});
				return;
			}
			const data = readProjectZip(file.name, entries);
			if (confirm(`Import ${data.files.length} file(s) from ${file.name}? It replaces the current project.`)) project.replace(data);
		} catch (err) {
			alert(`Could not import ${file.name}: ${err instanceof Error ? err.message : err}`);
		}
	}
	// Appends a declaration to the open file and selects it.
	async function insertDecl(text: string) {
		const f = project.activeFile;
		if (!f) return;
		const sep = f.text === '' ? '' : f.text.endsWith('\n\n') ? '' : f.text.endsWith('\n') ? '\n' : '\n\n';
		const v = view;
		if (v) {
			const from = v.state.doc.length;
			v.dispatch({ changes: { from, insert: sep + text }, selection: EditorSelection.cursor(from + sep.length), scrollIntoView: true });
			v.focus();
		} else {
			project.setText(f.path, f.text + sep + text);
		}
	}
	function exportPack() {
		if (!engine) return;
		if (engine.hasErrors && !confirm('The project has errors, so the pack is incomplete. Download it anyway?')) return;
		download(`${project.name || 'pack'}.zip`, packZip(engine.output));
	}

	const totals = $derived.by(() => {
		let errors = 0;
		let warnings = 0;
		for (const ds of Object.values(engine?.diagnostics ?? {})) {
			for (const d of ds) {
				if ((d.severity ?? 1) === 1) errors++;
				else if (d.severity === 2) warnings++;
			}
		}
		return { errors, warnings };
	});
	const statusText = $derived(
		!engine
			? 'Loading'
			: engine.status === 'ready'
				? totals.errors
					? `${totals.errors} error${totals.errors === 1 ? '' : 's'}`
					: totals.warnings
						? `${totals.warnings} warning${totals.warnings === 1 ? '' : 's'}`
						: 'No problems'
				: engine.status === 'starting'
					? 'Starting engine…'
					: engine.status === 'restarting'
						? 'Restarting engine…'
						: 'Engine stopped'
	);
</script>

<svelte:window onpointerdown={onWindowPointerDown} onkeydown={onWindowKeydown} />

<div class="flex h-dvh flex-col">
	<header class="flex items-center gap-3 border-b border-line bg-panel px-3 py-2">
		<span class="font-semibold tracking-tight">MMS</span>
		<span class="hidden text-sm text-muted sm:inline">{project.name}</span>
		<div class="relative ml-2" bind:this={menuRoot}>
			<button class="rounded px-2 py-1 text-sm whitespace-nowrap hover:bg-hover" bind:this={menuButton} onclick={() => (menu = !menu)} aria-expanded={menu}>Project ▾</button>
			{#if menu}
				<div class="absolute top-full left-0 z-20 mt-1 w-64 rounded-md border border-line bg-panel p-1 text-sm shadow-lg">
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={newProject}>New project</button>
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={() => { menu = false; fileInput.click(); }}>
						<span class="block">Import zip…</span>
						<span class="block text-xs text-muted">An MMS project or a data pack</span>
					</button>
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" disabled={!project.activeFile} onclick={() => { menu = false; pasteJson?.open(); }}>Paste JSON as MMS…</button>
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" disabled={!project.activeFile} onclick={() => { menu = false; editor?.fixAll(); }}>
						<span class="block">Fix all in this file</span>
						<span class="block text-xs text-muted">Applies safe fixes (Ctrl/⌘+Alt+F)</span>
					</button>
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={() => { menu = false; download(`${project.name || 'project'}-source.zip`, projectZip(project.data())); }}>Export project zip</button>
					<div class="my-1 border-t border-line"></div>
					<div class="px-2 pt-1 pb-0.5 text-[11px] font-medium tracking-wide text-muted uppercase">Examples</div>
					{#each examples as ex (ex.id)}
						<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={() => openExample(ex.id)}>
							<span class="block">{ex.id}</span>
							<span class="block text-xs text-muted">{ex.description}</span>
						</button>
					{/each}
					<div class="my-1 border-t border-line"></div>
					<label class="flex items-center justify-between px-2 py-1.5">
						Theme
						<select class="rounded border border-line bg-panel px-1 py-0.5 text-sm" bind:value={theme}>
							<option value="system">System</option>
							<option value="light">Light</option>
							<option value="dark">Dark</option>
						</select>
					</label>
				</div>
			{/if}
		</div>
		<button
			class="rounded px-2 py-1 text-sm whitespace-nowrap hover:bg-hover disabled:opacity-50"
			disabled={!editor || !project.activeFile}
			onclick={() => editor?.format()}
			title="Format this file (Ctrl/⌘+S or Shift+Alt+F)">Format</button
		>
		<input type="file" accept=".zip" class="hidden" bind:this={fileInput} onchange={importZip} />
		{#if engine}
			<PasteJson bind:this={pasteJson} {engine} uri={uriOf(project.active)} oninsert={insertDecl} />
		{/if}
		<div class="ml-auto flex items-center gap-2 text-sm">
			<span class="flex items-center gap-1.5 text-muted" role="status">
				<span
					class="inline-block size-2 rounded-full {engine?.status === 'ready'
						? totals.errors
							? 'bg-red-500'
							: 'bg-emerald-500'
						: engine?.status === 'failed'
							? 'bg-red-500'
							: 'bg-amber-400'}"
				></span>
				<span class="hidden sm:inline">{statusText}</span>
			</span>
			<a class="rounded px-2 py-1 hover:bg-hover" href={resolve('/docs')}>Docs</a>
			{#if engine?.status === 'failed'}
				<button class="rounded px-2 py-1 text-accent hover:bg-hover" onclick={() => engine?.restart()} title={engine.statusDetail}>Restart</button>
			{/if}
			<button class="rounded bg-accent px-3 py-1 font-medium text-white hover:opacity-90 disabled:opacity-50" disabled={!engine?.output.length} onclick={exportPack}><span class="sm:hidden">Pack</span><span class="hidden sm:inline">Download pack</span></button>
		</div>
	</header>

	<main class="grid min-h-0 flex-1 grid-cols-1 grid-rows-[minmax(0,1fr)_minmax(0,1fr)] md:grid-cols-[13rem_minmax(0,1fr)_minmax(0,1fr)] md:grid-rows-1">
		<aside class="hidden min-h-0 border-r border-line bg-panel md:block">
			{#if engine}
				<Sidebar {project} {engine} {selectedOutput} onoutput={selectOutput} onsymbol={(uri: string, s: DocSymbol) => reveal(uri, s.selectionRange)} />
			{/if}
		</aside>
		<section class="flex min-h-0 flex-col border-b border-line md:border-r md:border-b-0">
			<div class="flex items-center gap-2 border-b border-line bg-panel px-2 py-1 md:hidden">
				<select class="min-w-0 flex-1 rounded border border-line bg-panel px-1 py-0.5 text-sm" bind:value={project.active} aria-label="File">
					{#each project.files as f (f.path)}<option value={f.path}>{f.path}</option>{/each}
				</select>
				<span class="text-xs text-muted">{statusText}</span>
			</div>
			<div class="min-h-0 flex-1">
			{#if engine && loaded && project.activeFile}
				<Editor
					bind:this={editor}
					{engine}
					path={project.active}
					text={project.activeFile.text}
					onchange={(p: string, t: string) => project.setText(p, t)}
					onview={(v: EditorView) => {
						onview(v);
						v.dispatch({ effects: [] });
					}}
					extra={cursorExt}
				/>
			{:else if loaded}
				<p class="p-4 text-sm text-muted">No files. Add one from the sidebar.</p>
			{/if}
			</div>
		</section>
		<section class="flex min-h-0 flex-col bg-panel">
			<div class="flex border-b border-line text-xs" role="tablist">
				{#each [['preview', 'Preview'], ['json', 'JSON']] as [id, label] (id)}
					<button
						role="tab"
						aria-selected={rightTab === id}
						class="border-b-2 px-3 py-2 font-medium {rightTab === id ? 'border-accent text-ink' : 'border-transparent text-muted hover:text-ink'}"
						onclick={() => (rightTab = id as typeof rightTab)}>{label}</button
					>
				{/each}
			</div>
			<div class="min-h-0 flex-1">
				{#if engine}
					<div class="h-full" class:hidden={rightTab !== 'preview'}>
						{#if selectedFile && inspectable(selectedFile.path)}
							<Inspect file={selectedFile} />
						{:else}
							<Preview output={engine.output} target={selectedOutput} />
						{/if}
					</div>
					<div class="h-full" class:hidden={rightTab !== 'json'}>
						<JsonView file={selectedFile} onsource={(f: OutputFile) => f.uri && f.range && reveal(f.uri, f.range)} />
					</div>
				{/if}
			</div>
		</section>
	</main>
</div>
