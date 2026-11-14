<script lang="ts">
	import Logo from '#lib/components/Logo.svelte';
	import { onMount, tick, untrack } from 'svelte';
	import { asset, resolve } from '$app/paths';
	import { LSPPlugin } from '@codemirror/lsp-client';
	import { EditorSelection } from '@codemirror/state';
	import { EditorView } from '@codemirror/view';
	import Editor from '#lib/editor/Editor.svelte';
	import Sidebar from '#lib/components/Sidebar.svelte';
	import JsonView from '#lib/components/JsonView.svelte';
	import TestsView from '#lib/components/TestsView.svelte';
	import JigsawPanel from '#lib/components/JigsawPanel.svelte';
	import WorldPanel from '#lib/components/WorldPanel.svelte';
	import { Engine, itemKey, outputKey, type Diagnostic, type DocSymbol, type Expansion, type OutputFile } from '#lib/engine/engine.svelte.ts';
	import { Project, previewMemoryKey, pathOf, uriOf, starter, type ProjectData } from '#lib/project/project.svelte.ts';
	import { examples } from '#lib/project/examples.ts';
	import { download, hasResourcePack, isDataPack, packZip, projectZip, readProjectZip, readZip } from '#lib/project/zip.ts';
	import PasteJson from '#lib/components/PasteJson.svelte';
	import ProjectsDialog from '#lib/components/ProjectsDialog.svelte';
	import AboutDialog from '#lib/components/AboutDialog.svelte';
	import BuiltinViewer from '#lib/components/BuiltinViewer.svelte';
	import { requestPersistence } from '#lib/project/store.ts';
	import { applySplineEdits, followEdit, readSpline, trackPos, type SplineTrack } from '#lib/engine/spline.ts';
	import Dropdown from '#lib/components/Dropdown.svelte';
	import { fixAllKeys, formatKeys, mod, quickFixKeys } from '#lib/editor/shortcuts.ts';
	import { bookPageAt, PinButton, PreviewPane, shows, storedMemory, type SplineAt, type SplineEdit } from '@mms/preview';
	import PreviewWorker from '@mms/preview/worker?worker';

	const project = new Project();
	// What each preview was left showing, per declaration, kept in this browser for each project.
	const previewMemory = $derived(storedMemory(previewMemoryKey(project.id)));
	let engine = $state<Engine | null>(null);
	let view: EditorView | null = null;
	let viewWaiters: ((v: EditorView) => void)[] = [];
	let loaded = $state(false);
	let selectedOutput = $state<string | null>(null);
	let rightTab = $state<'preview' | 'json' | 'expanded' | 'tests' | 'jigsaw' | 'world'>('preview');
	let cursor = $state<{ uri: string; line: number; character: number } | null>(null);
	let theme = $state<'system' | 'light' | 'dark'>('system');
	const themes = { system: 'System', light: 'Light', dark: 'Dark' } as const;
	function nextTheme() {
		theme = theme === 'system' ? 'light' : theme === 'light' ? 'dark' : 'system';
	}
	let fileInput: HTMLInputElement;
	let pasteJson = $state<ReturnType<typeof PasteJson> | null>(null);
	let editor = $state<ReturnType<typeof Editor> | null>(null);
	let projects = $state<ReturnType<typeof ProjectsDialog> | null>(null);
	let about = $state<ReturnType<typeof AboutDialog> | null>(null);
	let builtinViewer = $state<ReturnType<typeof BuiltinViewer> | null>(null);

	onMount(() => {
		try {
			theme = (localStorage.getItem('mms-theme') as typeof theme) ?? 'system';
		} catch {
			/* storage unavailable */
		}
		const e = new Engine(asset('mms-worker.js'), {
			onText: (uri, text) => project.setText(pathOf(uri), text),
			display: async (uri): Promise<EditorView | null> => {
				if (uri.startsWith('mms-prelude:') || uri.startsWith('mms-vanilla:')) {
					const text: string | null = await e.builtinFile(uri);
					return text === null || !builtinViewer ? null : builtinViewer.show(uri, text);
				}
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
		const head = v.state.selection.main.head;
		const l = v.state.doc.lineAt(head);
		cursor = { uri: uriOf(project.active), line: l.number - 1, character: head - l.from };
	}
	const cursorExt = EditorView.updateListener.of((u) => {
		if (u.docChanged) {
			const uri = uriOf(project.active);
			if (cursorTrack) followEdit(cursorTrack, uri, u);
			for (const t of pinnedTracks.values()) followEdit(t, uri, u);
		}
		if (u.selectionSet || u.docChanged) trackCursor(u.view);
	});

	const within = (r: Diagnostic['range'], line: number) => r.start.line <= line && line <= r.end.line;
	const cursorOutput = $derived.by(() => {
		if (!engine || !cursor) return null;
		const c = cursor;
		// Files hoisted from inline values share their declaration's lines;
		// the declaration's own file is the one to show.
		// A custom item shows the item stack its name stands for: it may
		// write no file at all, and its files are only its look and name.
		// A book's .mmsbook file shows the book that reads it.
		const it = engine.outputItems.find((it) => (it.uri === c.uri && within(it.range, c.line)) || bookPageAt(it.pages, c.uri, c.line) !== null);
		if (it) return itemKey(it);
		// A tell, title or actionbar line, or a .mmstext file one reads,
		// shows its message.
		const t = engine.outputTexts.find((t) => t.uri === c.uri && within(t.range, c.line));
		if (t) return t.id;
		const f = engine.output.find((f) => f.uri === c.uri && f.range && !f.implicit && within(f.range, c.line));
		return f ? outputKey(f) : null;
	});
	$effect(() => {
		if (cursorOutput) selectedOutput = cursorOutput;
	});
	const selectedFile = $derived.by((): OutputFile | null => {
		if (!engine || !selectedOutput) return null;
		const f = engine.output.find((f) => outputKey(f) === selectedOutput);
		if (f) return f;
		const it = engine.outputItems.find((it) => itemKey(it) === selectedOutput);
		if (it) return { path: it.id, text: it.text, uri: it.uri, range: it.range, item: it };
		const t = engine.outputTexts.find((t) => t.id === selectedOutput);
		return t ? { path: t.kind, text: JSON.stringify(JSON.parse(t.text), null, 2), uri: t.uri, range: t.range } : null;
	});

	// The Preview tab shows one pane, or two split side by side or
	// stacked. A pane follows the cursor, or is pinned: it keeps showing one
	// declaration, and the spline it showed, while the cursor moves
	// elsewhere, and still redraws as the pack changes. The JSON tab keeps
	// following the cursor.
	interface Pane {
		id: number;
		pinned: boolean;
		/** The pinned output; null for a spline pinned outside of one. */
		path: string | null;
		/** The pinned spline, read again after each analysis. */
		spline: SplineAt | null;
		/** The book page to open. */
		page: number;
	}
	const newPane = (id: number): Pane => ({ id, pinned: false, path: null, spline: null, page: 0 });
	let panes = $state<Pane[]>([newPane(0)]);
	let nextPaneId = 1;
	let splitDir = $state<'row' | 'column'>('column');
	// The first pane's share of the split, in percent.
	let splitAt = $state(50);
	let splitBox = $state<HTMLElement | null>(null);
	// The Tests tab shows how many tests fail, so a failure is seen from the other tabs.
	const testsLabel = $derived.by(() => {
		const failed = engine?.tests.filter((t) => !t.pass).length ?? 0;
		return failed ? `Tests (${failed} failed)` : 'Tests';
	});
	$effect(() => {
		if (rightTab === 'tests' && !engine?.hasTests) rightTab = 'preview';
	});
	// The Jigsaw tab, for a project with jigsaw structures: template pools,
	// jigsaw holes or structure files with jigsaw blocks.
	const hasJigsaw = $derived(
		!!engine &&
			(engine.outputHoles.length > 0 ||
				engine.output.some((f) => f.path.includes('/worldgen/template_pool/')) ||
				engine.outputAssets.some((a) => a.jigsaws?.length))
	);
	$effect(() => {
		if (rightTab === 'jigsaw' && !hasJigsaw) rightTab = 'preview';
	});
	// Goes to the first error, for a world build that stopped on one.
	function firstProblem() {
		for (const [uri, ds] of Object.entries(engine?.diagnostics ?? {})) {
			const d = ds.find((d) => (d.severity ?? 1) === 1);
			if (d) return void reveal(uri, d.range);
		}
	}
	// Adds text to a source file through its editor, so it can be undone,
	// and selects what was added.
	async function insertAt(uri: string, line: number | undefined, text: string, select: string) {
		await reveal(uri, { start: { line: 0, character: 0 }, end: { line: 0, character: 0 } });
		const v = view;
		if (!v) return;
		const doc = v.state.doc;
		const at = line !== undefined && line < doc.lines ? doc.line(line + 1).from : doc.length;
		let insert = text;
		if (at === doc.length && doc.length > 0 && !doc.sliceString(doc.length - 1).endsWith('\n')) insert = '\n' + insert;
		v.dispatch({ changes: { from: at, insert }, userEvent: 'input' });
		const from = v.state.doc.sliceString(0).indexOf(select, at);
		if (from >= 0) v.dispatch({ selection: EditorSelection.range(from, from + select.length), effects: EditorView.scrollIntoView(from, { y: 'center' }) });
		v.focus();
	}
	// The Expanded tab shows what the Template call or Blueprint instance
	// at the cursor writes, refreshed as the cursor moves and the code changes.
	let expansion = $state<Expansion | null>(null);
	$effect(() => {
		const c = cursor;
		void engine?.output; // a new analysis
		if (rightTab !== 'expanded' || !engine || !c) return;
		let live = true;
		void engine.expand(c.uri, { line: c.line, character: c.character }).then((x) => {
			if (live) expansion = x;
		});
		return () => {
			live = false;
		};
	});
	// The density function spline at the cursor opens the spline editor in
	// the Preview tab; its edits go to the source through the editor, so
	// they can be undone. cursorTrack, and pinnedTracks for each pinned
	// pane's spline by pane id, map the spline's ranges to the text now.
	// They change in place as the text changes, so are not $state.
	let cursorSpline = $state<SplineAt | null>(null);
	let cursorTrack: SplineTrack | null = null;
	const pinnedTracks = new Map<number, SplineTrack>();
	const following = $derived(panes.some((p) => !p.pinned));
	$effect(() => {
		const c = cursor;
		void engine?.output; // a new analysis
		const v = view;
		if (rightTab !== 'preview' || !following || !engine || !c || !v) {
			cursorSpline = null;
			return;
		}
		let live = true;
		void readSpline(engine, c.uri, v.state.selection.main.head).then((r) => {
			if (!live || !r) return;
			cursorSpline = r.spline;
			cursorTrack = r.track;
		});
		return () => {
			live = false;
		};
	});
	// A pinned spline is read again, where it was pinned, after each analysis.
	let pinnedReads = 0;
	const latestRead = new Map<number, number>();
	$effect(() => {
		void engine?.output; // a new analysis
		const e = engine;
		if (rightTab !== 'preview' || !e) return;
		for (const p of panes) {
			if (!p.pinned) continue;
			const id = p.id;
			untrack(() => {
				const track = pinnedTracks.get(id);
				const pos = track && trackPos(e, track);
				if (!track || pos === null || pos === undefined) return;
				const seq = ++pinnedReads;
				latestRead.set(id, seq);
				void readSpline(e, track.uri, pos).then((r) => {
					const pane = panes.find((x) => x.id === id);
					if (!r || !pane?.pinned || latestRead.get(id) !== seq || !pinnedTracks.has(id)) return;
					pinnedTracks.set(id, r.track);
					pane.spline = r.spline;
				});
			});
		}
	});
	function editSpline(p: Pane, edits: SplineEdit[]) {
		const track = p.pinned ? pinnedTracks.get(p.id) : cursorTrack;
		if (engine && track) applySplineEdits(engine, track, edits);
	}
	const panePath = (p: Pane) => (p.pinned ? p.path : selectedOutput);
	const paneSpline = (p: Pane) => (p.pinned ? p.spline : cursorSpline);
	const pinnable = $derived(!!cursorSpline || !!shows(selectedOutput, engine?.outputItems, engine?.outputTexts));
	function pin(p: Pane) {
		if (!pinnable) return;
		p.pinned = true;
		p.path = shows(selectedOutput, engine?.outputItems, engine?.outputTexts) ? selectedOutput : null;
		p.spline = cursorSpline;
		if (cursorSpline && cursorTrack) pinnedTracks.set(p.id, { ...cursorTrack });
	}
	function unpin(p: Pane) {
		p.pinned = false;
		p.path = null;
		p.spline = null;
		pinnedTracks.delete(p.id);
	}
	const togglePin = (p: Pane) => (p.pinned ? unpin(p) : pin(p));
	const pinLabel = (p: Pane) => (p.pinned && !p.path ? 'Pinned spline' : undefined);
	// A second pane shows what the first does, pinned, so the first can go
	// on following the cursor; next to a pinned first pane it follows the cursor.
	function split() {
		const p = newPane(nextPaneId++);
		panes.push(p);
		const added = panes[panes.length - 1];
		if (!panes[0].pinned) pin(added);
	}
	function closePane(p: Pane) {
		pinnedTracks.delete(p.id);
		panes = panes.filter((x) => x.id !== p.id);
	}
	function resizeSplit(e: PointerEvent) {
		const handle = e.currentTarget as HTMLElement;
		handle.setPointerCapture(e.pointerId);
		const move = (m: PointerEvent) => {
			const r = splitBox?.getBoundingClientRect();
			if (!r) return;
			const at = splitDir === 'row' ? (m.clientX - r.left) / r.width : (m.clientY - r.top) / r.height;
			splitAt = Math.round(Math.min(85, Math.max(15, at * 100)));
		};
		const up = () => {
			handle.removeEventListener('pointermove', move);
			handle.removeEventListener('pointerup', up);
		};
		handle.addEventListener('pointermove', move);
		handle.addEventListener('pointerup', up);
	}
	// A book opens on the page the cursor is in, and stays there while
	// the cursor is outside its text.
	$effect(() => {
		for (const p of panes) {
			const path = panePath(p);
			const it = engine?.outputItems.find((it) => itemKey(it) === path);
			const at = it && cursor ? bookPageAt(it.pages, cursor.uri, cursor.line) : null;
			if (at !== null && untrack(() => p.page) !== at) p.page = at;
		}
	});

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
		if (!panes.some((p) => p.pinned) && !shows(path, engine?.outputItems, engine?.outputTexts)) rightTab = 'json';
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
		for (const p of panes) unpin(p);
		cursorTrack = null;
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
			const { texts: entries, assets } = readZip(bytes);
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
			await openAsNew(readProjectZip(file.name, entries, assets));
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
	// The pack's copied textures as data URLs, by pack path, so the recipe
	// inspector can draw a custom item's icon.
	const textures = $derived.by(() => {
		const out: Record<string, string> = {};
		if (!engine) return out;
		for (const a of engine.outputAssets) {
			if (!/^assets\/[^/]+\/textures\/.+\.png$/.test(a.path)) continue;
			const f = project.assets.find((x) => x.path === a.src);
			if (f) out[a.path] = `data:image/png;base64,${f.data}`;
		}
		return out;
	});
	// The pack's structure files as data URLs, by pack path, so the jigsaw
	// simulation can draw their blocks.
	const structures = $derived.by(() => {
		const out: Record<string, string> = {};
		if (!engine) return out;
		for (const a of engine.outputAssets) {
			if (!a.path.endsWith('.nbt')) continue;
			const f = project.assets.find((x) => x.path === a.src);
			if (f) out[a.path] = `data:application/octet-stream;base64,${f.data}`;
		}
		return out;
	});
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
		<span class="flex shrink-0 items-center"><Logo class="h-6 w-auto sm:hidden" wordmark={false} /><Logo class="hidden h-6 w-auto sm:block" /></span>
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
					<!-- Phones have no room for the About button in the top bar. -->
					<button class="block w-full rounded px-2 py-1.5 text-left hover:bg-hover sm:hidden" onclick={() => { close(); about?.open(); }}>About and credits</button>
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
		<AboutDialog bind:this={about} />
		<BuiltinViewer bind:this={builtinViewer} />
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
			{#if engine?.configError}
				<span class="rounded px-1.5 py-0.5 text-sm text-red-600" role="alert" title={engine.configError}>Settings not applied</span>
			{:else if engine?.configWarnings.length}
				<span class="rounded px-1.5 py-0.5 text-sm text-amber-600" role="status" title={engine.configWarnings.join('\n')}>Check settings</span>
			{/if}
			{#if project.saveError}
				<span class="rounded px-1.5 py-0.5 text-sm text-red-600" role="alert" title={project.saveError}>Not saved</span>
			{/if}
			<a class="rounded px-1.5 py-1 hover:bg-hover sm:px-2" href={resolve('/docs')}>Docs</a>
			<button class="hidden rounded p-1.5 hover:bg-hover sm:inline-block" onclick={() => about?.open()} aria-label="About and credits" title="About MMS, and credits: previews by Deepslate"
				><svg class="size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><circle cx="12" cy="12" r="9" /><path d="M12 11v5M12 8h.01" /></svg></button
			>
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
					title="Apply every safe fix in this file ({fixAllKeys}). {quickFixKeys} shows the fixes at the cursor.">Fix all</button
				>
				<button
					class="rounded px-2 py-0.5 whitespace-nowrap hover:bg-hover disabled:opacity-50"
					disabled={!editor || !project.activeFile}
					onclick={() => editor?.format()}
					title="Format this file ({mod}+S or {formatKeys})">Format</button
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
				{#each [['preview', 'Preview'], ['json', 'JSON'], ['expanded', 'Expanded'], ...(engine?.hasTests ? [['tests', testsLabel]] : []), ...(hasJigsaw ? [['jigsaw', 'Jigsaw']] : []), ['world', 'World']] as [id, label] (id)}
					<button
						role="tab"
						aria-selected={rightTab === id}
						class="border-b-2 px-3 py-2 font-medium {rightTab === id ? 'border-accent text-ink' : 'border-transparent text-muted hover:text-ink'}"
						onclick={() => (rightTab = id as typeof rightTab)}>{label}</button
					>
				{/each}
				{#if rightTab === 'preview'}
					<div class="ml-auto flex min-w-0 items-center">
						{#if panes.length === 1 && (panes[0].pinned || pinnable)}
							<PinButton pinned={panes[0].pinned ? (panes[0].path ?? '') : null} label={pinLabel(panes[0])} onclick={() => togglePin(panes[0])} />
						{/if}
						{#if panes.length === 1}
							<button class="shrink-0 px-2 py-1 text-muted hover:text-ink" onclick={split} aria-label="Split the preview" title="Split: show two previews at once, such as a pinned one and the one at the cursor">
								<svg viewBox="0 0 16 16" class="h-3.5 w-3.5" fill="none" stroke="currentColor" stroke-width="1.4" aria-hidden="true">
									<rect x="1.5" y="2.5" width="13" height="11" rx="1.5" />
									<path d="M8 2.5v11" />
								</svg>
							</button>
						{:else}
							<button
								class="shrink-0 px-2 py-1 text-muted hover:text-ink"
								onclick={() => (splitDir = splitDir === 'row' ? 'column' : 'row')}
								aria-label={splitDir === 'row' ? 'Stack the previews' : 'Put the previews side by side'}
								title={splitDir === 'row' ? 'Stack the previews' : 'Put the previews side by side'}
							>
								<svg viewBox="0 0 16 16" class="h-3.5 w-3.5" fill="none" stroke="currentColor" stroke-width="1.4" aria-hidden="true">
									<rect x="1.5" y="2.5" width="13" height="11" rx="1.5" />
									<path d={splitDir === 'row' ? 'M1.5 8h13' : 'M8 2.5v11'} />
								</svg>
							</button>
						{/if}
					</div>
				{/if}
			</div>
			<div class="min-h-0 flex-1">
				{#if engine}
					<div class="flex h-full min-h-0 {splitDir === 'row' ? 'flex-row' : 'flex-col'}" class:hidden={rightTab !== 'preview'} bind:this={splitBox}>
						{#each panes as p, i (p.id)}
							{#if i > 0}
								<!-- svelte-ignore a11y_no_static_element_interactions -->
								<div
									role="separator"
									aria-orientation={splitDir === 'row' ? 'vertical' : 'horizontal'}
									aria-label="Resize the previews"
									class="shrink-0 touch-none bg-line hover:bg-accent {splitDir === 'row' ? 'w-1 cursor-col-resize' : 'h-1 cursor-row-resize'}"
									onpointerdown={resizeSplit}
								></div>
							{/if}
							<div class="flex min-h-0 min-w-0 basis-0 flex-col" style:flex-grow={panes.length > 1 ? (i === 0 ? splitAt : 100 - splitAt) : 1} data-testid="preview-pane">
								{#if panes.length > 1}
									<div class="flex items-center border-b border-line text-xs">
										{#if p.pinned || pinnable}
											<PinButton pinned={p.pinned ? (p.path ?? '') : null} label={pinLabel(p)} onclick={() => togglePin(p)} />
										{/if}
										<button class="ml-auto shrink-0 px-2 py-1 text-muted hover:text-ink" onclick={() => closePane(p)} aria-label="Close this preview" title="Close this preview">
											<svg viewBox="0 0 16 16" class="h-3.5 w-3.5" fill="none" stroke="currentColor" stroke-width="1.4" stroke-linecap="round" aria-hidden="true">
												<path d="M4 4l8 8M12 4l-8 8" />
											</svg>
										</button>
									</div>
								{/if}
								<div class="min-h-0 flex-1">
									<PreviewPane
										output={engine.output}
										items={engine.outputItems}
										texts={engine.outputTexts}
										assets={engine.outputAssets}
										{textures}
										{structures}
										page={p.page}
										path={panePath(p)}
										pinned={p.pinned ? p.path : null}
										onunpin={() => unpin(p)}
										createWorker={() => new PreviewWorker()}
										spline={paneSpline(p)}
										onsplineedit={(edits: SplineEdit[]) => editSpline(p, edits)}
										memory={previewMemory}
									/>
								</div>
							</div>
						{/each}
					</div>
					<div class="h-full" class:hidden={rightTab !== 'json'}>
						<JsonView file={selectedFile} onsource={(f: OutputFile) => f.uri && f.range && reveal(f.uri, f.range)} />
					</div>
					{#if rightTab === 'expanded'}
						<div class="flex h-full min-h-0 flex-col text-sm" aria-label="Expanded">
							{#if expansion}
								<p class="border-b border-line px-3 py-1.5 text-xs text-muted">What <span class="font-mono text-ink">{expansion.title}</span> writes</p>
								<pre class="min-h-0 flex-1 overflow-auto px-3 py-2 font-mono text-[12.5px] leading-5" data-testid="expanded-text">{expansion.text || '(nothing)'}</pre>
							{:else}
								<p class="p-4 text-sm text-muted">Put the cursor in a Template call or a Blueprint declaration, such as a CustomItem, to see the code it writes.</p>
							{/if}
						</div>
					{/if}
					{#if rightTab === 'jigsaw'}
						<JigsawPanel
							{engine}
							projectId={project.id}
							projectName={project.name}
							ns={project.config.namespace}
							onreveal={reveal}
							onpiece={(path) => {
								selectedOutput = path;
								rightTab = 'preview';
							}}
							oninsert={insertAt}
						/>
					{/if}
					{#if rightTab === 'world'}
						<WorldPanel
							{engine}
							{project}
							onproblems={firstProblem}
							onpiece={(path) => {
								selectedOutput = path;
								rightTab = 'preview';
							}}
							resources={resourcePack}
							ondownload={exportPack}
						/>
					{/if}
					{#if rightTab === 'tests'}
						<div class="h-full">
							<TestsView tests={engine.tests} error={engine.testsError} onsource={reveal} />
						</div>
					{/if}
				{/if}
			</div>
		</section>
	</main>
</div>
