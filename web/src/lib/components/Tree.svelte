<script lang="ts">
	import Tree from './Tree.svelte';
	import type { Node } from './tree';
	interface Props {
		nodes: Node[];
		selected: string | null;
		onselect: (path: string) => void;
		depth?: number;
	}
	let { nodes, selected, onselect, depth = 0 }: Props = $props();
	let collapsed = $state<Record<string, boolean>>({});
</script>

<ul class="text-[13px]">
	{#each nodes as n (n.path)}
		<li>
			{#if n.leaf}
				<button
					class="flex w-full items-center truncate rounded py-0.5 pr-2 text-left hover:bg-hover"
					class:bg-selected={selected === n.path}
					style:padding-left="{depth * 12 + 8}px"
					title={n.path}
					onclick={() => onselect(n.path)}>{n.name}</button
				>
			{:else}
				<button
					class="flex w-full items-center gap-1 truncate rounded py-0.5 pr-2 text-left text-muted hover:bg-hover"
					style:padding-left="{depth * 12 + 2}px"
					onclick={() => (collapsed[n.path] = !collapsed[n.path])}
				>
					<span class="inline-block w-3 text-[10px]">{collapsed[n.path] ? '▸' : '▾'}</span>{n.name}
				</button>
				{#if !collapsed[n.path]}
					<Tree nodes={n.children} {selected} {onselect} depth={depth + 1} />
				{/if}
			{/if}
		</li>
	{/each}
</ul>
