<script lang="ts">
	// The sections of the page, read from the h2 and h3 headings inside
	// `root`. It watches the DOM, so headings that come and go (a new page,
	// or the reference filter hiding groups) update the list. The section
	// being read is highlighted.
	let { root }: { root: HTMLElement | undefined } = $props();

	type Item = { id: string; text: string; depth: number };
	let items: Item[] = $state([]);
	let active = $state('');
	let list: HTMLElement | undefined = $state();

	const OFFSET = 140; // below the sticky header, filter and group heading

	function headings(): HTMLElement[] {
		return root ? [...root.querySelectorAll<HTMLElement>(':is(h2, h3)[id]')] : [];
	}

	function collect() {
		const next = headings().map((h) => ({ id: h.id, text: (h.dataset.toc ?? h.textContent ?? '').trim(), depth: h.tagName === 'H2' ? 2 : 3 }));
		const same = next.length === items.length && next.every((n, i) => n.id === items[i].id && n.text === items[i].text);
		if (!same) items = next;
		spy();
	}

	function spy() {
		let current = '';
		for (const h of headings()) {
			if (h.getBoundingClientRect().top > OFFSET) break;
			current = h.id;
		}
		if (!current && items.length) current = items[0].id;
		if (current !== active) active = current;
	}

	// Batch DOM and scroll changes into one update per frame.
	function batch(fn: () => void) {
		let queued = false;
		return () => {
			if (queued) return;
			queued = true;
			requestAnimationFrame(() => {
				queued = false;
				fn();
			});
		};
	}

	$effect(() => {
		if (!root) return;
		const onChange = batch(collect);
		const onScroll = batch(spy);
		const mo = new MutationObserver(onChange);
		mo.observe(root, { childList: true, subtree: true, characterData: true });
		addEventListener('scroll', onScroll, { passive: true });
		addEventListener('resize', onScroll, { passive: true });
		collect();
		return () => {
			mo.disconnect();
			removeEventListener('scroll', onScroll);
			removeEventListener('resize', onScroll);
		};
	});

	// Keep the highlighted entry visible in a long list without moving the page.
	$effect(() => {
		const a = active && list?.querySelector<HTMLElement>(`a[data-id="${CSS.escape(active)}"]`);
		if (!a || !list) return;
		const top = a.offsetTop - list.offsetTop;
		if (top < list.scrollTop || top + a.offsetHeight > list.scrollTop + list.clientHeight) {
			list.scrollTop = top - list.clientHeight / 3;
		}
	});
</script>

{#if items.length > 1}
	<nav class="toc" aria-label="On this page">
		<p class="toc-title">On this page</p>
		<ul bind:this={list}>
			{#each items as it (it.id)}
				<li class:sub={it.depth === 3}>
					<a href="#{it.id}" data-id={it.id} class:active={it.id === active} aria-current={it.id === active ? 'location' : undefined}>{it.text}</a>
				</li>
			{/each}
		</ul>
	</nav>
{/if}

<style>
	.toc {
		font-size: 13px;
	}
	.toc-title {
		font-weight: 600;
		margin: 0 0 0.5rem;
		padding-left: 0.75rem;
	}
	ul {
		max-height: calc(100dvh - 8rem);
		overflow-y: auto;
		border-left: 1px solid var(--c-line);
		position: relative;
	}
	li.sub a {
		padding-left: 1.5rem;
	}
	a {
		display: block;
		padding: 0.2rem 0.75rem;
		margin-left: -1px;
		border-left: 2px solid transparent;
		color: var(--c-muted);
		line-height: 1.4;
	}
	a:hover {
		color: var(--c-ink);
	}
	a.active {
		color: var(--c-accent);
		border-left-color: var(--c-accent);
		font-weight: 500;
	}
</style>
