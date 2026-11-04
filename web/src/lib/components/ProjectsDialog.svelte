<script lang="ts">
	import { onMount } from 'svelte';
	import { starter, type Project } from '../project/project.svelte';
	import { onProjectsChanged, requestPersistence, type StorageState } from '../project/store';

	interface Props {
		project: Project;
		/** Called after another project opens, so the page can reset its view. */
		onswitch?: () => void;
	}
	let { project, onswitch }: Props = $props();

	let dialog: HTMLDialogElement;
	let renaming = $state<string | null>(null);
	let renameText = $state('');
	let storage = $state<StorageState | null>(null);
	let busy = $state(false);

	export async function open() {
		renaming = null;
		dialog.showModal();
		await project.refresh();
		storage = await requestPersistence(false);
	}

	onMount(() => onProjectsChanged(() => void project.refresh()));

	async function run(fn: () => Promise<unknown>, switches = true) {
		busy = true;
		try {
			await fn();
			if (switches) onswitch?.();
		} finally {
			busy = false;
		}
	}

	const openProject = (id: string) =>
		run(async () => {
			if (!(await project.open(id))) alert('That project is no longer saved.');
			dialog.close();
		});
	const create = () =>
		run(async () => {
			await project.create(structuredClone(starter));
			dialog.close();
		});
	const duplicate = (id: string) => run(() => project.duplicate(id).then(() => dialog.close()));
	function remove(id: string, name: string) {
		if (!confirm(`Delete the project ${name}? This cannot be undone; export it first to keep a copy.`)) return;
		void run(() => project.deleteProject(id), id === project.id);
	}
	function startRename(id: string, name: string) {
		renaming = id;
		renameText = name;
	}
	async function commitRename() {
		const id = renaming;
		renaming = null;
		if (id) await run(() => project.renameProject(id, renameText), false);
	}
	async function keep() {
		storage = await requestPersistence(true);
		if (!storage.persisted) alert('The browser declined. Projects are still saved, but it may clear them when the disk is low on space. Bookmarking the site or installing it can help; exporting a zip keeps a copy.');
	}

	const ago = (ms: number) => {
		const s = Math.round((Date.now() - ms) / 1000);
		if (s < 60) return 'just now';
		if (s < 3600) return `${Math.round(s / 60)} min ago`;
		if (s < 86400) return `${Math.round(s / 3600)} h ago`;
		return new Date(ms).toLocaleDateString();
	};
	const mb = (n: number) => (n / 1024 / 1024).toFixed(n < 10 * 1024 * 1024 ? 1 : 0) + ' MB';
</script>

<dialog bind:this={dialog} class="m-auto w-[min(34rem,calc(100vw-2rem))] rounded-lg border border-line bg-panel p-0 text-ink shadow-xl backdrop:bg-black/40" aria-labelledby="projects-title">
	<div class="flex flex-col gap-3 p-4">
		<div class="flex items-center justify-between gap-2">
			<h2 id="projects-title" class="font-semibold">Projects</h2>
			<button class="rounded bg-accent px-3 py-1 text-sm font-medium text-white disabled:opacity-50" disabled={busy} onclick={create}>New project</button>
		</div>
		<p class="text-sm text-muted">Projects are saved in this browser as you type. Export a zip from the Project menu to move one to another browser.</p>
		<ul class="flex max-h-[50vh] flex-col gap-1 overflow-y-auto" aria-label="Saved projects">
			{#each project.saved as p (p.id)}
				<li class="group flex items-center gap-2 rounded border px-2 py-1.5 {p.id === project.id ? 'border-accent' : 'border-line'}">
					{#if renaming === p.id}
						<form class="flex min-w-0 flex-1 gap-2" onsubmit={(e) => { e.preventDefault(); void commitRename(); }}>
							<!-- svelte-ignore a11y_autofocus -->
							<input class="min-w-0 flex-1 rounded border border-line bg-panel px-1 py-0.5 text-sm" bind:value={renameText} aria-label="Project name" autofocus onkeydown={(e) => { if (e.key === 'Escape') { e.preventDefault(); e.stopPropagation(); renaming = null; } }} />
							<button class="rounded px-2 text-sm text-accent hover:bg-hover" type="submit">Save</button>
						</form>
					{:else}
						<button class="min-w-0 flex-1 text-left disabled:cursor-default" disabled={busy} onclick={() => openProject(p.id)} aria-current={p.id === project.id ? 'true' : undefined}>
							<span class="block truncate text-sm font-medium">{p.name}</span>
							<span class="block text-xs text-muted">{p.id === project.id ? 'Open · ' : ''}{p.files} file{p.files === 1 ? '' : 's'} · {ago(p.updated)}</span>
						</button>
						<div class="flex shrink-0 gap-0.5 text-xs">
							<button class="rounded px-1.5 py-1 text-muted hover:bg-hover hover:text-ink" disabled={busy} onclick={() => startRename(p.id, p.name)} aria-label="Rename {p.name}">Rename</button>
							<button class="rounded px-1.5 py-1 text-muted hover:bg-hover hover:text-ink" disabled={busy} onclick={() => duplicate(p.id)} aria-label="Duplicate {p.name}">Duplicate</button>
							<button class="rounded px-1.5 py-1 text-muted hover:bg-hover hover:text-red-600" disabled={busy} onclick={() => remove(p.id, p.name)} aria-label="Delete {p.name}">Delete</button>
						</div>
					{/if}
				</li>
			{:else}
				<li class="text-sm text-muted">This browser cannot save projects (storage is blocked, as in some private windows).</li>
			{/each}
		</ul>
		<div class="flex flex-wrap items-center justify-between gap-2 border-t border-line pt-3 text-xs text-muted">
			{#if storage}
				<span>
					{storage.persisted ? 'Kept: the browser will not clear these to free space.' : 'The browser may clear these when the disk is low on space.'}
					{#if storage.usage !== undefined}{mb(storage.usage)} used.{/if}
				</span>
				{#if !storage.persisted}
					<button class="rounded px-2 py-1 text-accent hover:bg-hover" onclick={keep}>Keep them</button>
				{/if}
			{/if}
			<button class="ml-auto rounded px-3 py-1 text-sm text-ink hover:bg-hover" onclick={() => dialog.close()}>Close</button>
		</div>
	</div>
</dialog>
