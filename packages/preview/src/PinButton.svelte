<script lang="ts">
	import { shows } from './paths';

	interface Props {
		/** The pinned output path, or null when the preview follows the cursor. */
		pinned: string | null;
		onclick: () => void;
		class?: string;
	}
	let { pinned, onclick, class: cls = '' }: Props = $props();
</script>

<button
	class="flex min-w-0 items-center gap-1 px-3 py-1 {pinned ? 'text-accent' : 'text-muted hover:text-ink'} {cls}"
	aria-pressed={pinned !== null}
	title={pinned ? 'Unpin: the preview follows the cursor again' : 'Pin: keep previewing this declaration while you edit others'}
	{onclick}
>
	<svg viewBox="0 0 16 16" class="h-3.5 w-3.5 shrink-0" fill={pinned ? 'currentColor' : 'none'} stroke="currentColor" stroke-width="1.4" aria-hidden="true">
		<path d="M6 1.5h4l-.5 4 2.5 2.5v1.5H4V8l2.5-2.5z" stroke-linejoin="round" />
		<path d="M8 9.5v5" stroke-linecap="round" />
	</svg>
	{#if pinned}
		<span class="truncate">Pinned {shows(pinned)?.id}</span>
	{:else}
		<span>Pin</span>
	{/if}
</button>
