<script lang="ts">
	import { getContext } from 'svelte';
	import type { Obj } from './common';
	import { customsOfBase, emptyItems, itemsKey, look, type ItemContext } from './items';
	import type { Slot } from './recipe';

	interface Props {
		slot: Slot;
		count?: number;
		big?: boolean;
		/** The stack's components, for a result: a custom item's name and icon. */
		components?: Obj;
	}
	let { slot, count = 1, big = false, components }: Props = $props();
	const items = getContext<() => ItemContext>(itemsKey) ?? (() => emptyItems);
	const looks = $derived(slot?.map((id, i) => look(id, i === 0 ? components : undefined, items())) ?? []);
	const first = $derived(looks[0]);
	const title = $derived.by(() => {
		if (!looks.length) return '';
		if (looks.length > 1) return looks.map((l) => l.title.split('\n')[0]).join(' or ');
		const l = looks[0];
		// An ingredient matches by id only, so custom items on that base fit too.
		const customs = components ? [] : customsOfBase(l.id, items());
		return customs.length ? `${l.title}\nAny ${l.name} fits, including the custom ${customs.length > 1 ? 'items' : 'item'} ${customs.join(', ')}` : l.title;
	});
</script>

<div
	class="relative flex items-center justify-center overflow-hidden rounded border border-line bg-canvas p-1 text-center leading-tight {big ? 'h-20 w-20 text-xs' : 'h-16 w-16 text-[10px]'}"
	{title}
	role="img"
	aria-label={looks.map((l) => l.name).join(' or ') || 'empty slot'}
>
	{#if first}
		{#if first.icon}
			<img class="{big ? 'size-12' : 'size-10'} [image-rendering:pixelated]" src={first.icon} alt="" />
		{:else}
			<span class="line-clamp-3 break-words {first.id.startsWith('#') ? 'text-accent' : ''}">{first.name}</span>
		{/if}
		{#if looks.length > 1}<span class="absolute top-0.5 left-1 text-[9px] text-muted">+{looks.length - 1}</span>{/if}
		{#if count > 1}<span class="absolute right-1 bottom-0.5 font-semibold">{count}</span>{/if}
	{/if}
</div>
