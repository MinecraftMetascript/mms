<script lang="ts" module>
	export interface MenuItem {
		label: string;
		/** Shortcut hint shown on the right, such as "F12". */
		keys?: string;
		disabled?: boolean;
		run: () => void;
	}
	/** A list of items; null draws a separator. */
	export type MenuEntries = (MenuItem | null)[];
</script>

<script lang="ts">
	import { tick } from 'svelte';

	let open = $state(false);
	let x = $state(0);
	let y = $state(0);
	let entries = $state<MenuEntries>([]);
	let root: HTMLDivElement | undefined = $state();
	let onclose: (() => void) | undefined;

	/** Opens the menu at a viewport point. onclose runs when it closes without an action. */
	export async function show(at: { x: number; y: number }, items: MenuEntries, closed?: () => void) {
		entries = items;
		x = at.x;
		y = at.y;
		onclose = closed;
		open = true;
		await tick();
		if (!root) return;
		// Keep the menu inside the viewport.
		const r = root.getBoundingClientRect();
		if (r.right > innerWidth - 4) x = Math.max(4, innerWidth - r.width - 4);
		if (r.bottom > innerHeight - 4) y = Math.max(4, innerHeight - r.height - 4);
		items_()[0]?.focus();
	}

	function close(restore = true) {
		if (!open) return;
		open = false;
		if (restore) onclose?.();
	}

	function items_() {
		return [...(root?.querySelectorAll<HTMLButtonElement>('button:not(:disabled)') ?? [])];
	}

	function pick(item: MenuItem) {
		close(false);
		item.run();
	}

	function onkeydown(e: KeyboardEvent) {
		const list = items_();
		const i = list.indexOf(document.activeElement as HTMLButtonElement);
		if (e.key === 'ArrowDown' || e.key === 'ArrowUp') {
			e.preventDefault();
			const n = list.length;
			if (n) list[(i + (e.key === 'ArrowDown' ? 1 : n - 1) + (i < 0 ? 1 : 0)) % n].focus();
		} else if (e.key === 'Home' || e.key === 'End') {
			e.preventDefault();
			list[e.key === 'Home' ? 0 : list.length - 1]?.focus();
		} else if (e.key === 'Escape' || e.key === 'Tab') {
			e.preventDefault();
			close();
		}
	}
</script>

<svelte:window
	onpointerdown={(e) => open && root && !root.contains(e.target as Node) && close(false)}
	onblur={() => close(false)}
	onresize={() => close(false)}
/>

{#if open}
	<div
		bind:this={root}
		role="menu"
		tabindex="-1"
		aria-label="Editor actions"
		class="fixed z-[200] min-w-56 rounded-md border border-line bg-panel p-1 text-sm shadow-lg"
		style="left: {x}px; top: {y}px"
		{onkeydown}
		oncontextmenu={(e) => e.preventDefault()}
	>
		{#each entries as item, i (i)}
			{#if item}
				<button
					role="menuitem"
					class="flex w-full items-center justify-between gap-6 rounded px-2 py-1.5 text-left hover:bg-hover focus:bg-hover focus:outline-none disabled:opacity-50 disabled:hover:bg-transparent"
					disabled={item.disabled}
					onclick={() => pick(item)}
				>
					<span class="truncate">{item.label}</span>
					{#if item.keys}<kbd class="font-sans text-xs whitespace-nowrap text-muted">{item.keys}</kbd>{/if}
				</button>
			{:else}
				<div class="my-1 border-t border-line" role="separator"></div>
			{/if}
		{/each}
	</div>
{/if}
