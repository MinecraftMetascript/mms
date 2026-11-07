<script lang="ts">
	import { resolve } from '$app/paths';
	import { goto } from '$app/navigation';
	import type { Entry, Finder } from './find.ts';

	// The site search: Ctrl+K or / opens it. The index is fetched the first
	// time it opens.
	let open = $state(false);
	let q = $state('');
	let active = $state(0);
	let finder: Finder | undefined = $state();
	let failed = $state(false);
	let input: HTMLInputElement | undefined = $state();
	const base = resolve('/').replace(/\/$/, '');
	const results = $derived(finder ? finder.search(q) : []);

	async function show() {
		open = true;
		active = 0;
		queueMicrotask(() => input?.select());
		if (!finder && !failed) {
			try {
				// FlexSearch loads with the index, the first time search opens.
				const [{ Finder }, entries] = await Promise.all([import('./find.ts'), fetch(base + '/docs/search.json').then((r) => r.json() as Promise<Entry[]>)]);
				finder = new Finder(entries);
			} catch {
				failed = true;
			}
		}
	}

	function choose(e: Entry) {
		open = false;
		q = '';
		goto(base + e.path);
	}

	function onkey(e: KeyboardEvent) {
		const typing = e.target instanceof HTMLElement && e.target.closest('input, textarea, [contenteditable]');
		if ((e.key === 'k' && (e.metaKey || e.ctrlKey)) || (e.key === '/' && !typing && !open)) {
			e.preventDefault();
			show();
		}
	}

	function onInputKey(e: KeyboardEvent) {
		if (e.key === 'Escape') open = false;
		else if (e.key === 'ArrowDown') active = Math.min(active + 1, results.length - 1);
		else if (e.key === 'ArrowUp') active = Math.max(active - 1, 0);
		else if (e.key === 'Enter' && results[active]) choose(results[active]);
		else return;
		e.preventDefault();
	}
</script>

<svelte:window onkeydown={onkey} />

<button class="flex items-center gap-2 rounded border border-line px-2 py-1 text-sm text-muted hover:bg-hover" onclick={show}>
	Search<kbd class="hidden rounded border border-line px-1 text-[11px] sm:inline">Ctrl K</kbd>
</button>

{#if open}
	<!-- svelte-ignore a11y_click_events_have_key_events, a11y_no_static_element_interactions -->
	<div class="fixed inset-0 z-30 bg-black/30" onclick={() => (open = false)}></div>
	<div class="fixed inset-x-4 top-16 z-40 mx-auto max-w-xl rounded-lg border border-line bg-panel shadow-lg" role="dialog" aria-label="Search the docs">
		<input
			bind:this={input}
			bind:value={q}
			oninput={() => (active = 0)}
			onkeydown={onInputKey}
			class="w-full rounded-t-lg border-b border-line bg-panel px-3 py-2.5 outline-none"
			type="search"
			placeholder="Search guides, kinds, helpers, types and lint rules"
			aria-label="Search the docs"
		/>
		<ul class="max-h-[60vh] overflow-y-auto py-1" role="listbox" aria-label="Results">
			{#each results as r, i (r.path)}
				<li role="option" aria-selected={i === active}>
					<a class="block px-3 py-1.5 {i === active ? 'bg-selected' : ''}" href={base + r.path} onclick={(e) => { e.preventDefault(); choose(r); }} onmousemove={() => (active = i)}>
						<span class="font-medium">{r.title}</span>
						<span class="ml-2 text-xs text-muted">{r.where}</span>
					</a>
				</li>
			{:else}
				<li class="px-3 py-2 text-sm text-muted">{q.trim() ? (finder ? `Nothing matches “${q}”.` : failed ? 'Search could not load.' : 'Loading…') : 'Type to search.'}</li>
			{/each}
		</ul>
	</div>
{/if}
