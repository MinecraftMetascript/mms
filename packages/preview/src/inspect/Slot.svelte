<script lang="ts">
	import { label } from './common';
	import type { Slot } from './recipe';

	interface Props {
		slot: Slot;
		count?: number;
		big?: boolean;
	}
	let { slot, count = 1, big = false }: Props = $props();
	const title = $derived(slot?.map(label).join(' or ') ?? '');
	// Short text for the tile: the first item's name, and how many others fit.
	const text = $derived(slot ? label(slot[0]) : '');
</script>

<div
	class="relative flex items-center justify-center overflow-hidden rounded border border-line bg-canvas p-1 text-center leading-tight {big ? 'h-20 w-20 text-xs' : 'h-16 w-16 text-[10px]'}"
	{title}
	role="img"
	aria-label={title || 'empty slot'}
>
	{#if slot}
		<span class="line-clamp-3 break-words {slot[0].startsWith('#') ? 'text-accent' : ''}">{text}</span>
		{#if slot.length > 1}<span class="absolute top-0.5 left-1 text-[9px] text-muted">+{slot.length - 1}</span>{/if}
		{#if count > 1}<span class="absolute right-1 bottom-0.5 font-semibold">{count}</span>{/if}
	{/if}
</div>
