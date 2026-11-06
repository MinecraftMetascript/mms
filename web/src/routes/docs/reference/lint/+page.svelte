<script lang="ts">
	import { resolve } from '$app/paths';
	import { refURL } from '#lib/docs/reference.ts';
	import Doc from '#lib/docs/Doc.svelte';

	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const base = resolve('/').replace(/\/$/, '');
</script>

<svelte:head><title>Lint rules · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>Lint rules</h1>
	<p>
		Lint rules point out style problems and likely mistakes; errors the checker reports are not lint. A project can change any rule's
		severity in <code>mms.json</code>, and <a href="{base}/docs/guides/lint">Lint rules</a> in the guides explains how, and how to silence one.
	</p>
	<table>
		<thead><tr><th>Rule</th><th>Default</th><th>Reports</th></tr></thead>
		<tbody>
			{#each data.rules as r (r.id)}
				<tr><td><a href={refURL.rule(base, r.id)}><code>{r.id}</code></a></td><td>{r.severity}</td><td>{#if r.doc}<Doc text={r.doc} />{/if}</td></tr>
			{/each}
		</tbody>
	</table>
</article>
