<script lang="ts">
	// The World tab: the World and Structures pages VS Code shows in its
	// side bar (packages/preview), for a world folder on this computer. In
	// Chrome and Edge the page can write to a folder the user picks, so
	// it builds the pack into the world as `mms build --world` does (the
	// engine works out what to write) and keeps the folder for the next
	// visit. Other browsers can only read a picked folder, so the World
	// page offers the downloads instead, and Structures lists the files
	// as they were when picked.
	import WorldTool from '@mms/preview/world/WorldTool.svelte';
	import { importPath, nextView, packDimensions, structureCall, type WorldAction, type WorldFeatures, type WorldResult, type WorldState, type WorldStructureRow } from '@mms/preview/world/tool';
	import type { Engine } from '../engine/engine.svelte';
	import type { Project } from '../project/project.svelte';
	import { fromBase64, toBase64 } from '../project/base64';
	import { canWrite, FilesFolder, HandleFolder, isWorld, levelName, pickFolder, savedFolder, saveFolder, type WorldFolder } from '../world/folder';
	import { dimensionRows, findStructures, readPackDir, resetDimensions, resetFolder, writePackDir, type FoundStructure } from '../world/world';

	let {
		engine,
		project,
		onproblems,
		onpiece,
		resources,
		ondownload
	}: {
		engine: Engine;
		project: Project;
		/** Shows the first problem in the project. */
		onproblems: () => void;
		/** Shows a structure file of the project, by its path in the pack. */
		onpiece: (packPath: string) => void;
		/** Whether the project has a resource pack. */
		resources: boolean;
		/** Downloads the data pack, or the resource pack. */
		ondownload: (resources: boolean) => void;
	} = $props();

	const writable = canWrite();
	const ns = $derived(project.config.namespace);

	// The build world, and a separate world structures come from, if any.
	let world = $state<WorldFolder>();
	let pieces = $state<WorldFolder>();
	let names = $state<Record<'build' | 'structures', string | undefined>>({ build: undefined, structures: undefined });
	// Saved folders the browser wants the user to allow again this visit.
	let locked = $state<HandleFolder[]>([]);
	let building = $state(false);
	let result = $state<WorldResult>();
	let found = $state<FoundStructure[]>([]);
	let rows = $state<WorldStructureRow[]>([]);
	let dimensions = $state<WorldState['dimensions']>([]);
	let note = $state('');

	const resetKey = $derived(`mms-world-reset:${project.id}`);
	let resetOn = $state(false);
	$effect(() => {
		try {
			resetOn = localStorage.getItem(resetKey) === '1';
		} catch {
			resetOn = false;
		}
	});

	// The project's saved worlds, each time another project opens.
	$effect(() => {
		const id = project.id;
		world = pieces = undefined;
		locked = [];
		result = undefined;
		found = [];
		void (async () => {
			const b = await savedFolder(id, 'build');
			const s = await savedFolder(id, 'structures');
			if (id !== project.id) return;
			for (const f of [b, s]) if (f && !(await f.allowed(false))) locked.push(f);
			world = b;
			pieces = s;
		})();
	});
	const usable = (f: WorldFolder | undefined) => (f && !locked.includes(f as HandleFolder) ? f : undefined);
	const from = $derived(usable(pieces) ?? usable(world));

	$effect(() => {
		const f = world;
		names.build = f?.name;
		if (f && usable(f)) void levelName(f).then((n) => world === f && (names.build = n ?? f.name));
	});
	$effect(() => {
		const f = pieces;
		names.structures = f?.name;
		if (f && usable(f)) void levelName(f).then((n) => pieces === f && (names.structures = n ?? f.name));
	});

	/** Reads the worlds again: the pack's dimensions there, and the saved structures. */
	async function refresh() {
		const w = usable(world);
		const ids = packDimensions(engine.output.map((f) => f.path));
		const dims = w ? await dimensionRows(w, ids) : [];
		if (JSON.stringify(dims) !== JSON.stringify(dimensions)) dimensions = dims;
		const f = from;
		const list = f ? await findStructures(f) : [];
		// Only a change redraws the list.
		const sig = (l: FoundStructure[]) => l.map((s) => `${s.file}@${s.saved.getTime()}`).join('\n');
		if (sig(list) !== sig(found)) found = list;
	}
	$effect(() => {
		void engine.output;
		void from;
		void world;
		void refresh();
	});
	// New saves show up when the user comes back from the game, and while
	// the tab is open.
	$effect(() => {
		const t = setInterval(() => document.visibilityState === 'visible' && void refresh(), 4000);
		const back = () => void refresh();
		window.addEventListener('focus', back);
		return () => {
			clearInterval(t);
			window.removeEventListener('focus', back);
		};
	});

	// Each structure's bytes, as base64 to compare with the project's copy,
	// read again only when its file changed.
	const hashes = new Map<string, { key: string; data: string }>();
	async function bytes(s: FoundStructure): Promise<string> {
		const key = `${s.saved.getTime()}`;
		const got = hashes.get(s.file);
		if (got?.key === key) return got.data;
		const data = toBase64(await s.data());
		hashes.set(s.file, { key, data });
		return data;
	}
	const placed = (s: FoundStructure) => {
		const rel = importPath(s, ns, 'structures');
		return { rel, call: structureCall(rel, s) };
	};
	$effect(() => {
		const list = found;
		const assets = project.assets.map((a) => ({ ...a }));
		const unused = new Set(engine.unusedAssets.filter((p) => p.endsWith('.nbt')));
		void (async () => {
			const out: WorldStructureRow[] = [];
			for (const s of list) {
				const { rel, call } = placed(s);
				const mine = assets.find((a) => a.path === rel);
				const status = !mine ? 'new' : mine.data === (await bytes(s)) ? 'imported' : 'changed';
				out.push({
					file: s.file,
					id: `${s.ns}:${s.path}`,
					status,
					rel,
					call,
					saved: s.saved.toLocaleString(),
					savedMs: s.saved.getTime(),
					unused: status !== 'new' && unused.has(rel),
					packId: s.ns !== ns && ns ? `${ns}:${s.path}` : undefined
				});
			}
			if (list === found) rows = out;
		})();
	});
	const unusedFiles = $derived.by(() => {
		const saved = new Set(rows.map((r) => r.rel));
		return engine.unusedAssets.filter((p) => p.endsWith('.nbt') && !saved.has(p)).sort();
	});

	const describe = (f: WorldFolder, name: string | undefined) => ({ name: name ?? f.name, folder: name && name !== f.name ? f.name : undefined, path: f.name });
	const toolState = $derived<WorldState>({
		project: { name: project.name, ns, path: 'mms.json', canChoose: false },
		world: usable(world) ? describe(world!, names.build) : undefined,
		structureWorld: from ? { ...describe(from, from === pieces ? names.structures : names.build), separate: from === pieces } : undefined,
		building,
		buildOnSave: false,
		resetDimensions: resetOn,
		dimensions,
		result,
		structures: rows,
		unusedFiles,
		unusedPaused: engine.hasErrors
	});

	const features: WorldFeatures = { revealFile: 'Preview the piece', deleteFromWorld: writable };
	// Language server errors are plain { code, message } objects.
	const errText = (e: unknown) => (e && typeof e === 'object' && 'message' in e ? String(e.message) : String(e));
	const at = () => new Date().toLocaleTimeString(undefined, { hour: '2-digit', minute: '2-digit' });

	async function choose(purpose: 'build' | 'structures') {
		let f: WorldFolder | undefined;
		if (writable) f = await pickFolder(purpose === 'build' ? 'mms-world' : 'mms-structures');
		else if (purpose === 'structures') f = await pickFiles();
		if (!f) return;
		if (!(await isWorld(f))) {
			alert(`${f.name} has no level.dat, so it is not a world. Choose a world's own folder, in the game's saves folder.`);
			return;
		}
		if (purpose === 'build') {
			world = f;
			result = undefined;
		} else {
			pieces = f;
		}
		locked = locked.filter((l) => l.handle !== f.handle);
		await saveFolder(project.id, purpose, f);
	}

	// Other browsers read a folder through a file input.
	let filesInput = $state<HTMLInputElement>();
	let filesPicked: ((f: WorldFolder | undefined) => void) | undefined;
	function pickFiles(): Promise<WorldFolder | undefined> {
		return new Promise((resolve) => {
			filesPicked = resolve;
			filesInput!.value = '';
			filesInput!.click();
		});
	}

	async function allow(f: HandleFolder) {
		if (await f.allowed(true)) {
			locked = locked.filter((l) => l !== f);
			await refresh();
		}
	}

	async function build(resetIds?: string[]) {
		const w = usable(world);
		if (!w || building) return;
		if (!(w instanceof HandleFolder) || !(await w.allowed(true))) {
			result = { ok: false, message: 'Not allowed to write to the world', detail: 'Allow this page to edit the folder when the browser asks, then build again.', problems: false, at: at() };
			return;
		}
		building = true;
		try {
			if (!(await isWorld(w))) throw new Error(`${w.name} has no level.dat any more`);
			const dir = `datapacks/${ns}`;
			const before = await readPackDir(w, dir);
			const zipFile = await w.file('resourcepacks/resources.zip');
			const oldZip = zipFile ? toBase64(new Uint8Array(await zipFile.arrayBuffer())) : null;
			let plan;
			try {
				plan = await engine.worldInstall({ before, oldZip, reset: !!resetIds || resetOn, resetIds });
			} catch (e) {
				const message = errText(e);
				const problems = /has errors/.test(message);
				result = { ok: false, message: problems ? 'The project has errors; nothing written' : 'Build failed', detail: message, problems, at: at() };
				return;
			}
			// The browser cannot tell whether the game has the world open,
			// which would save the loaded chunks again over a reset.
			if (plan.reset.length && !confirm(`Reset ${plan.reset.join(', ')}? Leave the world in Minecraft first (Save and Quit to Title): the game would save its loaded chunks again. The old chunks are kept in the world's ${resetFolder} folder until the next reset.`)) {
				result = { ok: false, message: `Leave the world to reset ${plan.reset.join(', ')}`, detail: '', problems: false, worldOpen: true, at: at() };
				return;
			}
			await writePackDir(w, dir, plan.data);
			if (plan.zip !== undefined) await w.write('resourcepacks/resources.zip', fromBase64(plan.zip));
			const done = plan.reset.length ? await resetDimensions(w, plan.reset) : [];
			const line = plan.message.replace(/^next step: /, '');
			result = {
				ok: true,
				files: plan.data.length + (plan.zip !== undefined ? plan.resources : 0),
				at: at(),
				next: nextView(plan.next, line),
				resources: plan.zip !== undefined,
				reset: done.filter((d) => d.moved.length).map((d) => d.id),
				unvisited: done.filter((d) => !d.moved.length).map((d) => d.id)
			};
			if (plan.replaced?.length) note = `Overwrote ${plan.replaced.join(', ')} in the world's resources.zip.`;
		} catch (e) {
			result = { ok: false, message: 'Build failed', detail: errText(e), problems: false, at: at() };
		} finally {
			building = false;
			void refresh();
		}
	}

	async function importRows(list: WorldStructureRow[]) {
		for (const r of list) {
			const s = found.find((f) => f.file === r.file);
			if (s) project.setAsset(r.rel, await bytes(s));
		}
		note = list.length === 1 ? `Imported ${list[0].rel}. Use it in a template pool with ${list[0].call}.` : `Imported ${list.length} structures.`;
	}

	async function copy(text: string) {
		try {
			await navigator.clipboard.writeText(text);
		} catch {
			note = text;
		}
	}

	async function send(m: WorldAction) {
		note = '';
		const row = (file: string) => rows.find((r) => r.file === file);
		switch (m.type) {
			case 'run':
				switch (m.action) {
					case 'build':
						return build();
					case 'chooseWorld':
						return choose('build');
					case 'chooseStructureWorld':
						return choose('structures');
					case 'sameStructureWorld':
						pieces = undefined;
						return saveFolder(project.id, 'structures', undefined);
					case 'dismiss':
						result = undefined;
						return;
					case 'copyReload':
						return copy('/reload');
					case 'showProblems':
						return onproblems();
					case 'importWaiting':
						return importRows(rows.filter((r) => r.status !== 'imported'));
				}
				return;
			case 'resetDimensions':
				resetOn = m.on;
				try {
					localStorage.setItem(resetKey, m.on ? '1' : '0');
				} catch {
					/* storage unavailable: the choice lasts for this visit */
				}
				return;
			case 'resetDimension': {
				const vanilla = /^minecraft:/.test(m.id);
				if (vanilla && !confirm(`Reset ${m.id}? All of its terrain generates again, including what you built there.`)) return;
				return build([m.id]);
			}
			case 'import': {
				const r = row(m.file);
				return r && importRows([r]);
			}
			case 'importFiles':
				return importRows(m.files.map(row).filter((r): r is WorldStructureRow => !!r));
			case 'copyCall': {
				const r = row(m.file);
				return r && copy(r.call);
			}
			case 'deleteFromWorld': {
				const s = found.find((f) => f.file === m.file);
				const f = from;
				if (!s || !(f instanceof HandleFolder) || !(await f.allowed(true))) return;
				if (!confirm(`Delete ${s.ns}:${s.path} from ${f.name}? This deletes its saves in the world's generated folder. The copy imported into the project, if any, stays. If the world is open, the game may still use the structure until it is reopened.`)) return;
				for (const folder of ['structure', 'structures']) await f.remove(`generated/${s.ns}/${folder}/${s.path}.nbt`);
				return refresh();
			}
			case 'revealFile': {
				const a = engine.outputAssets.find((a) => a.src === m.rel);
				if (a) onpiece(a.path);
				else note = `${m.rel} is not in the pack: no StructureFile(...) call uses it.`;
				return;
			}
		}
	}
</script>

<input
	type="file"
	class="hidden"
	webkitdirectory
	bind:this={filesInput}
	onchange={(e) => {
		const list = e.currentTarget.files;
		filesPicked?.(list?.length ? new FilesFolder(list) : undefined);
	}}
	oncancel={() => filesPicked?.(undefined)}
/>

<div class="h-full overflow-auto text-sm" aria-label="World">
	{#each locked as f (f)}
		<div class="flex items-center gap-2 border-b border-line px-3 py-2">
			<span class="grow text-muted">The browser needs your OK to use {f.name} again.</span>
			<button class="rounded border border-line px-2 py-0.5 hover:bg-hover" onclick={() => allow(f)}>Allow</button>
		</div>
	{/each}
	{#if writable}
		<WorldTool page="world" state={toolState} {send} {features} />
	{:else}
		<section class="border-b border-line px-3 py-2">
			<p class="my-1">This browser cannot write to a folder, so it cannot build into a world. Download the pack and put it in the world's <code>datapacks</code> folder; put the resource pack, if there is one, in the world as <code>resourcepacks/resources.zip</code>. Chrome and Edge build straight into a world.</p>
			<div class="mt-2 flex gap-2">
				<button class="rounded bg-accent px-3 py-1 font-medium text-white hover:opacity-90" onclick={() => ondownload(false)}>Download pack</button>
				{#if resources}
					<button class="rounded border border-accent px-3 py-1 font-medium text-accent hover:bg-hover" onclick={() => ondownload(true)}>Download resources</button>
				{/if}
			</div>
		</section>
	{/if}
	<h2 class="px-3 pt-3 text-[11px] font-bold tracking-wide uppercase">Structures</h2>
	<WorldTool page="structures" state={toolState} {send} {features} />
	{#if note}<p class="px-3 py-2 text-xs text-muted" aria-live="polite">{note}</p>{/if}
</div>
