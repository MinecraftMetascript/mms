<script lang="ts">
	import Prose from '#lib/docs/Prose.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head><title>{data.title} · MMS docs</title></svelte:head>

<Prose html={data.html} />
