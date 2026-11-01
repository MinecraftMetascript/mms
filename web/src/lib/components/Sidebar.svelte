<script lang="ts">
	import type { DocSymbol, Diagnostic, Engine, OutputFile } from '../engine/engine.svelte';
	import { assetExtensions, isAssetPath, uriOf, type Project } from '../project/project.svelte';
	import { toBase64 } from '../project/base64';
	import Tree from './Tree.svelte';
	import { buildTree } from './tree';

	interface Props {
		project: Project;
		engine: Engine;
		selectedOutput: string | null;
		onoutput: (path: string) => void;
		onsymbol: (uri: string, s: DocSymbol) => void;
	}
	let { project, engine, selectedOutput, onoutput, onsymbol }: Props = $props();

	let tab = $state<'sources' | 'pack' | 'symbols'>('sources');
	let renaming = $state<string | null>(null);
	let renameText = $state('');

	const counts = (ds: Diagnostic[] | undefined) => {
		let errors = 0;
		let warnings = 0;
		for (const d of ds ?? []) {
			if ((d.severity ?? 1) === 1) errors++;
			else if (d.severity === 2) warnings++;
		}
		return { errors, warnings };
	};

	const packTree = $derived(buildTree(engine.output.map((f: OutputFile) => f.path)));

	function startRename(path: string) {
		renaming = path;
		renameText = path;
	}
	function finishRename() {
		if (renaming && renameText.trim() && renameText !== renaming) {
			if (!project.rename(renaming, renameText)) alert(`A file named ${renameText} already exists.`);
		}
		renaming = null;
	}
	function newFile() {
		const name = prompt('File name', 'new.mms');
		if (name) project.add(name);
	}
	function remove(path: string) {
		if (confirm(`Delete ${path}? This cannot be undone.`)) project.remove(path);
	}
	let assetInput = $state<HTMLInputElement>();
	async function addAssets(e: Event) {
		const input = e.target as HTMLInputElement;
		for (const file of input.files ?? []) {
			if (!isAssetPath(file.name)) continue;
			if (project.assets.some((a) => a.path === file.name) && !confirm(`Replace ${file.name}?`)) continue;
			project.setAsset(file.name, toBase64(new Uint8Array(await file.arrayBuffer())));
		}
		input.value = '';
	}
	function removeAsset(path: string) {
		if (confirm(`Delete ${path}? This cannot be undone.`)) project.removeAsset(path);
	}
	const symbolKind: Record<number, string> = { 3: 'ns', 12: 'fn', 13: 'obj', 18: 'tag', 23: 'res' };
</script>

<div class="flex h-full min-h-0 flex-col">
	<div class="flex border-b border-line text-xs" role="tablist">
		{#each [['sources', 'Sources'], ['pack', 'Pack'], ['symbols', 'Symbols']] as [id, label] (id)}
			<button
				role="tab"
				aria-selected={tab === id}
				class="flex-1 border-b-2 px-2 py-2 font-medium {tab === id ? 'border-accent text-ink' : 'border-transparent text-muted hover:text-ink'}"
				onclick={() => (tab = id as typeof tab)}>{label}</button
			>
		{/each}
	</div>
	<div class="min-h-0 flex-1 overflow-auto p-1.5">
		{#if tab === 'sources'}
			<ul class="text-[13px]">
				{#each project.files as f (f.path)}
					{@const c = counts(engine.diagnostics[uriOf(f.path)])}
					<li class="group flex items-center rounded hover:bg-hover" class:bg-selected={project.active === f.path}>
						{#if renaming === f.path}
							<!-- svelte-ignore a11y_autofocus -->
							<input
								class="m-0.5 w-full rounded border border-accent bg-panel px-1.5 py-0.5"
								bind:value={renameText}
								autofocus
								onblur={finishRename}
								onkeydown={(e) => {
									if (e.key === 'Enter') finishRename();
									if (e.key === 'Escape') renaming = null;
								}}
							/>
						{:else}
							<button class="flex min-w-0 flex-1 items-center gap-1.5 px-2 py-0.5 text-left" onclick={() => (project.active = f.path)} ondblclick={() => startRename(f.path)}>
								<span class="truncate">{f.path}</span>
								{#if c.errors}<span class="rounded-full bg-red-600 px-1.5 text-[10px] leading-4 text-white">{c.errors}</span>{/if}
								{#if !c.errors && c.warnings}<span class="rounded-full bg-amber-500 px-1.5 text-[10px] leading-4 text-white">{c.warnings}</span>{/if}
							</button>
							<button class="hidden px-1 text-muted group-hover:block hover:text-ink" title="Rename" aria-label="Rename {f.path}" onclick={() => startRename(f.path)}>✎</button>
							<button class="hidden px-1.5 text-muted group-hover:block hover:text-red-600" title="Delete" aria-label="Delete {f.path}" onclick={() => remove(f.path)}>×</button>
						{/if}
					</li>
				{/each}
			</ul>
			<button class="mt-1 w-full rounded px-2 py-1 text-left text-[13px] text-accent hover:bg-hover" onclick={newFile}>+ New file</button>
			{#if project.assets.length}
				<div class="px-2 pt-3 pb-0.5 text-[11px] font-medium tracking-wide text-muted uppercase">Assets</div>
				<ul class="text-[13px]">
					{#each project.assets as a (a.path)}
						<li class="group flex items-center rounded hover:bg-hover" title="Use it with Texture(&quot;{a.path}&quot;) or Sound(&quot;{a.path}&quot;)">
							<span class="flex min-w-0 flex-1 items-center gap-1.5 px-2 py-0.5">
								{#if a.path.endsWith('.png')}
									<img class="size-4 shrink-0 [image-rendering:pixelated]" src="data:image/png;base64,{a.data}" alt="" />
								{:else}
									<span class="w-4 shrink-0 text-center text-[10px] text-muted">{a.path.endsWith('.ogg') ? '♪' : '▦'}</span>
								{/if}
								<span class="truncate">{a.path}</span>
							</span>
							<button class="hidden px-1.5 text-muted group-hover:block hover:text-red-600" title="Delete" aria-label="Delete {a.path}" onclick={() => removeAsset(a.path)}>×</button>
						</li>
					{/each}
				</ul>
			{/if}
			<button class="w-full rounded px-2 py-1 text-left text-[13px] text-accent hover:bg-hover" onclick={() => assetInput?.click()}>+ Add texture or sound</button>
			<input type="file" multiple accept={assetExtensions.join(',')} class="hidden" bind:this={assetInput} onchange={addAssets} />
		{:else if tab === 'pack'}
			{#if engine.output.length}
				<Tree nodes={packTree} selected={selectedOutput} onselect={onoutput} />
			{:else}
				<p class="p-2 text-xs text-muted">Nothing generated yet.</p>
			{/if}
		{:else}
			{#each project.files as f (f.path)}
				{@const syms = engine.symbols[uriOf(f.path)] ?? []}
				{#if syms.length}
					<div class="px-2 pt-2 pb-0.5 text-[11px] font-medium tracking-wide text-muted uppercase">{f.path}</div>
					<ul class="text-[13px]">
						{#each syms as s, i (i)}
							<li>
								<button class="flex w-full items-center gap-2 rounded px-2 py-0.5 text-left hover:bg-hover" onclick={() => onsymbol(uriOf(f.path), s)}>
									<span class="w-6 shrink-0 font-mono text-[10px] text-muted">{symbolKind[s.kind] ?? ''}</span>
									<span class="truncate">{s.name}</span>
									<span class="ml-auto shrink-0 text-[11px] text-muted">{s.detail}</span>
								</button>
							</li>
						{/each}
					</ul>
				{/if}
			{/each}
		{/if}
	</div>
</div>
