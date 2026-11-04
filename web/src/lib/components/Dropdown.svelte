<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		/** The button text; a ▾ follows it. */
		label: string;
		/** Extra classes for the panel, such as its width. */
		panelClass?: string;
		/** The menu's items. Call `close` from an item's handler. */
		children: Snippet<[close: () => void]>;
	}
	let { label, panelClass = 'w-64', children }: Props = $props();

	let open = $state(false);
	let root: HTMLDivElement | undefined = $state();
	let button: HTMLButtonElement | undefined = $state();

	const close = () => (open = false);

	// Close on a click outside the menu, or on Escape.
	function onWindowPointerDown(e: PointerEvent) {
		if (open && root && !root.contains(e.target as Node)) open = false;
	}
	function onWindowKeydown(e: KeyboardEvent) {
		if (open && e.key === 'Escape') {
			open = false;
			button?.focus();
		}
	}
</script>

<svelte:window onpointerdown={onWindowPointerDown} onkeydown={onWindowKeydown} />

<div class="relative" bind:this={root}>
	<button class="rounded px-1.5 py-1 text-sm whitespace-nowrap hover:bg-hover sm:px-2" bind:this={button} onclick={() => (open = !open)} aria-expanded={open}>{label} ▾</button>
	{#if open}
		<div class="absolute top-full left-0 z-20 mt-1 max-h-[80dvh] overflow-y-auto rounded-md border border-line bg-panel p-1 text-sm shadow-lg {panelClass}">
			{@render children(close)}
		</div>
	{/if}
</div>
