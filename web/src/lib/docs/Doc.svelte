<script lang="ts">
	import { inlineCode } from './reference.ts';

	// Schema documentation, with `code` spans. A trailing backslash is a
	// Markdown line break.
	let { text }: { text: string } = $props();
</script>

{#each inlineCode(text.replace(/[ \t]*\\\n/g, '\n')) as part, i (i)}{#if part.code}<code>{part.text}</code>{:else}{part.text}{/if}{/each}
