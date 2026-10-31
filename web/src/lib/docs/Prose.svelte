<script lang="ts">
	// Rendered guide HTML. The guides are this repository's own Markdown.
	let { html }: { html: string } = $props();
</script>

<!-- eslint-disable-next-line svelte/no-at-html-tags -->
<article class="prose-mms">{@html html}</article>
