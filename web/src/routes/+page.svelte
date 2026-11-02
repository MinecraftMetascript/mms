<script lang="ts">
	import { onMount, tick } from 'svelte';
	import { asset, resolve } from '$app/paths';
	import { LSPPlugin } from '@codemirror/lsp-client';
	import { EditorSelection } from '@codemirror/state';
	import { EditorView } from '@codemirror/view';
	import Editor from '#lib/editor/Editor.svelte';
	import Sidebar from '#lib/components/Sidebar.svelte';
	import JsonView from '#lib/components/JsonView.svelte';
	import { Engine, type Diagnostic, type DocSymbol, type OutputFile } from '#lib/engine/engine.svelte.ts';
	import { Project, pathOf, uriOf, starter, type ProjectData } from '#lib/project/project.svelte.ts';
	import { examples } from '#lib/project/examples.ts';
	import { download, hasResourcePack, isDataPack, packZip, projectZip, readProjectZip, readZip, readZipAssets } from '#lib/project/zip.ts';
	import PasteJson from '#lib/components/PasteJson.svelte';
	import ProjectsDialog from '#lib/components/ProjectsDialog.svelte';
	import { requestPersistence } from '#lib/project/store.ts';
	import Dropdown from '#lib/components/Dropdown.svelte';
	import { PinButton, PreviewPane, shows } from '@mms/preview';
	import PreviewWorker from '@mms/preview/worker?worker';

	const project = new Project();
	let engine = $state<Engine | null>(null);
	let view: EditorView | null = null;
	let viewWaiters: ((v: EditorView) => void)[] = [];
	let loaded = $state(false);
	let selectedOutput = $state<string | null>(null);
	let rightTab = $state<'preview' | 'json'>('preview');
	let cursor = $state<{ uri: string; line: number } | null>(null);
	let theme = $state<'system' | 'light' | 'dark'>('system');
	const themes = { system: 'System', light: 'Light', dark: 'Dark' } as const;
	function nextTheme() {
		theme = theme === 'system' ? 'light' : theme === 'light' ? 'dark' : 'system';
	}
	let fileInput: HTMLInputElement;
	let pasteJson = $state<ReturnType<typeof PasteJson> | null>(null);
	let editor = $state<ReturnType<typeof Editor> | null>(null);
	let projects = $state<ReturnType<typeof ProjectsDialog> | null>(null);

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
		// Save before the tab goes away; the save timer may not get to run.
		const flush = () => void project.flush();
		const onHide = () => document.visibilityState === 'hidden' && flush();
		window.addEventListener('pagehide', flush);
		document.addEventListener('visibilitychange', onHide);
		return () => {
			window.removeEventListener('pagehide', flush);
			document.removeEventListener('visibilitychange', onHide);
			e.dispose();
		};
	});

	// Keep the server's file set and config in step with the project.
	function syncEngine() {
		if (!engine) return;
		engine.setConfig($state.snapshot(project.config));
		engine.workspace.setFiles(project.files.map((f) => ({ uri: uriOf(f.path), text: f.text })));
		engine.setAssets($state.snapshot(project.assets));
		engine.scheduleRefresh();
	}
	$effect(() => {
		// Track the set of paths (not their text, which the editor syncs).
		void project.files.map((f) => f.path).join('\n');
		void project.assets.map((f) => f.path + ':' + f.data.length).join('\n');
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
		// Files hoisted from inline values share their declaration's lines;
		// the declaration's own file is the one to show.
		return engine.output.find((f) => f.uri === c.uri && f.range && !f.implicit && within(f.range, c.line)) ?? null;
	});
	$effect(() => {
		if (cursorOutput) selectedOutput = cursorOutput.path;
	});
	const selectedFile = $derived(engine?.output.find((f) => f.path === selectedOutput) ?? null);

	// A pinned output keeps the Preview tab on one declaration while the
	// cursor moves elsewhere; it still redraws as the pack changes. The JSON
	// tab keeps following the cursor.
	let pinned = $state<string | null>(null);
	const previewPath = $derived(pinned ?? selectedOutput);
	const pinnable = $derived(shows(previewPath));
	function togglePin() {
		pinned = pinned ? null : previewPath && shows(previewPath) ? previewPath : null;
	}

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
		// A pinned preview stays on screen; otherwise show JSON for outputs with no preview.
		if (!pinned && !shows(path)) rightTab = 'json';
	}

	// Opens data as a new saved project; the open one stays saved.
	async function openAsNew(data: ProjectData) {
		await project.create(data);
		onswitch();
		// Asked on a user action, since Firefox shows a prompt for it.
		void requestPersistence(true);
	}
	// Another project opened: forget this one's preview state.
	function onswitch() {
		pinned = null;
		selectedOutput = null;
		cursor = null;
	}
	function openExample(id: string) {
		const ex = examples.find((e) => e.id === id);
		if (ex) void openAsNew(structuredClone(ex.data));
	}
	function newProject() {
		void openAsNew(structuredClone(starter));
	}
	async function importZip(e: Event) {
		const file = (e.target as HTMLInputElement).files?.[0];
		(e.target as HTMLInputElement).value = '';
		if (!file) return;
		try {
			const bytes = new Uint8Array(await file.arrayBuffer());
			const entries = readZip(bytes);
			if (isDataPack(entries)) {
				if (!engine) return;
				const res = await engine.importPack([...entries].map(([path, text]) => ({ path, text })));
				if (!res.files.length) throw new Error('the pack has no resources MMS can import');
				const skipped = res.skipped.length ? ` ${res.skipped.length} file(s) cannot be imported and will be left out.` : '';
				if (skipped && !confirm(`Import ${res.count} resource(s) from the data pack ${file.name}?${skipped}`)) return;
				const files = res.files.map((f) => ({ path: f.path, text: f.text }));
				await openAsNew({
					name: file.name.replace(/\.zip$/i, ''),
					config: { namespace: res.namespace, description: res.description || 'Imported with MMS' },
					files,
					active: files[0].path
				});
				return;
			}
			await openAsNew(readProjectZip(file.name, entries, readZipAssets(bytes)));
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
	// The game loads a resource pack from its own zip, like `mms build --zip`,
	// so it gets its own button (browsers block a second download per click).
	const resourcePack = $derived(!!engine && hasResourcePack(engine.output, engine.outputAssets));
	function exportPack(resources = false) {
		if (!engine) return;
		if (engine.hasErrors && !confirm('The project has errors, so the pack is incomplete. Download it anyway?')) return;
		const name = (project.name || 'pack') + (resources ? '-resources' : '');
		download(`${name}.zip`, packZip(engine.output, engine.outputAssets, $state.snapshot(project.assets), resources));
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

<div class="flex h-dvh flex-col">
	<header class="flex items-center gap-2 border-b border-line bg-panel px-3 py-2">
		<span class="font-semibold tracking-tight">MMS</span>
		<button class="hidden max-w-48 truncate rounded px-2 py-1 text-sm text-muted hover:bg-hover hover:text-ink sm:inline" onclick={() => projects?.open()} title="Switch, rename or delete saved projects" aria-label="{project.name} (switch project)">{project.name}</button>
		<div class="flex items-center sm:ml-2">
			<Dropdown label="Project">
				{#snippet children(close)}
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={() => { close(); newProject(); }}>New project</button>
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={() => { close(); projects?.open(); }}>Open project…</button>
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={() => { close(); fileInput.click(); }}>
						<span class="block">Import zip…</span>
						<span class="block text-xs text-muted">An MMS project or a data pack</span>
					</button>
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={() => { close(); download(`${project.name || 'project'}-source.zip`, projectZip(project.data())); }}>Export project zip</button>
				{/snippet}
			</Dropdown>
			<Dropdown label="Examples" panelClass="w-72">
				{#snippet children(close)}
					{#each examples as ex (ex.id)}
						<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover" onclick={() => { close(); openExample(ex.id); }}>
							<span class="block">{ex.id}</span>
							<span class="block text-xs text-muted">{ex.description}</span>
						</button>
					{/each}
				{/snippet}
			</Dropdown>
		</div>
		<ProjectsDialog bind:this={projects} {project} {onswitch} />
		<input type="file" accept=".zip" class="hidden" bind:this={fileInput} onchange={importZip} />
		{#if engine}
			<PasteJson bind:this={pasteJson} {engine} uri={uriOf(project.active)} oninsert={insertDecl} />
		{/if}
		<div class="ml-auto flex items-center gap-1 text-sm sm:gap-2">
			<span class="flex items-center gap-1.5 text-muted" role="status" title={statusText}>
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
			<a class="rounded px-1.5 py-1 hover:bg-hover sm:px-2" href={resolve('/docs')}>Docs</a>
			<button
				class="rounded p-1.5 hover:bg-hover"
				onclick={nextTheme}
				aria-label="Theme: {themes[theme]}"
				title="Theme: {themes[theme]} (click to change)"
				><svg class="size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
					{#if theme === 'light'}
						<circle cx="12" cy="12" r="4" /><path d="M12 2v2M12 20v2M4.9 4.9l1.4 1.4M17.7 17.7l1.4 1.4M2 12h2M20 12h2M4.9 19.1l1.4-1.4M17.7 6.3l1.4-1.4" />
					{:else if theme === 'dark'}
						<path d="M21 12.8A9 9 0 1 1 11.2 3a7 7 0 0 0 9.8 9.8z" />
					{:else}
						<circle cx="12" cy="12" r="9" /><path d="M12 3a9 9 0 0 1 0 18z" fill="currentColor" />
					{/if}
				</svg></button
			>
			{#if engine?.status === 'failed'}
				<button class="rounded px-2 py-1 text-accent hover:bg-hover" onclick={() => engine?.restart()} title={engine.statusDetail}>Restart</button>
			{/if}
			{#if resourcePack}
				<button class="rounded border border-accent px-3 py-1 font-medium text-accent hover:bg-hover" title="The textures, sounds, models and translations, as a resource pack; put it in a world as resourcepacks/resources.zip" onclick={() => exportPack(true)}><span class="sm:hidden">Resources</span><span class="hidden sm:inline">Download resources</span></button>
			{/if}
			<button class="rounded bg-accent px-3 py-1 font-medium text-white hover:opacity-90 disabled:opacity-50" disabled={!engine?.output.length} onclick={() => exportPack()}><span class="sm:hidden">Pack</span><span class="hidden sm:inline">Download pack</span></button>
		</div>
	</header>

	<main class="grid min-h-0 flex-1 grid-cols-1 grid-rows-[minmax(0,1fr)_minmax(0,1fr)] md:grid-cols-[13rem_minmax(0,1fr)_minmax(0,1fr)] md:grid-rows-1">
		<aside class="hidden min-h-0 border-r border-line bg-panel md:block">
			{#if engine}
				<Sidebar {project} {engine} {selectedOutput} onoutput={selectOutput} onsymbol={(uri: string, s: DocSymbol) => reveal(uri, s.selectionRange)} />
			{/if}
		</aside>
		<section class="flex min-h-0 flex-col border-b border-line md:border-r md:border-b-0">
			<div class="flex items-center gap-1 border-b border-line bg-panel px-2 py-1 text-sm">
				<select class="min-w-0 flex-1 rounded border border-line bg-panel px-1 py-0.5 md:hidden" bind:value={project.active} aria-label="File">
					{#each project.files as f (f.path)}<option value={f.path}>{f.path}</option>{/each}
				</select>
				<span class="hidden min-w-0 flex-1 truncate text-xs text-muted md:block" title={project.active}>{project.active}</span>
				<button
					class="rounded px-2 py-0.5 whitespace-nowrap hover:bg-hover disabled:opacity-50"
					disabled={!engine || !project.activeFile}
					onclick={() => pasteJson?.open()}
					title="Convert vanilla JSON to an MMS declaration and add it to this file"><span class="sm:hidden">JSON…</span><span class="hidden sm:inline">Paste JSON…</span></button
				>
				<button
					class="rounded px-2 py-0.5 whitespace-nowrap hover:bg-hover disabled:opacity-50"
					disabled={!editor || !project.activeFile}
					onclick={() => editor?.fixAll()}
					title="Apply every safe fix in this file (Ctrl/⌘+Alt+F). Ctrl/⌘+. shows the fixes at the cursor.">Fix all</button
				>
				<button
					class="rounded px-2 py-0.5 whitespace-nowrap hover:bg-hover disabled:opacity-50"
					disabled={!editor || !project.activeFile}
					onclick={() => editor?.format()}
					title="Format this file (Ctrl/⌘+S or Shift+Alt+F)">Format</button
				>
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
				{#if rightTab === 'preview' && (pinned || pinnable)}
					<PinButton class="ml-auto" {pinned} onclick={togglePin} />
				{/if}
			</div>
			<div class="min-h-0 flex-1">
				{#if engine}
					<div class="h-full" class:hidden={rightTab !== 'preview'}>
						<PreviewPane
							output={engine.output}
							path={previewPath}
							{pinned}
							onunpin={() => (pinned = null)}
							createWorker={() => new PreviewWorker()}
						/>
					</div>
					<div class="h-full" class:hidden={rightTab !== 'json'}>
						<JsonView file={selectedFile} onsource={(f: OutputFile) => f.uri && f.range && reveal(f.uri, f.range)} />
					</div>
				{/if}
			</div>
		</section>
	</main>
</div>
