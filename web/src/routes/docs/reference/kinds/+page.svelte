<script lang="ts">
	import Doc from '#lib/docs/Doc.svelte';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
</script>

<svelte:head><title>Declaration kinds · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>Declaration kinds</h1>
	<p>Every declaration starts with its kind: <code>LootTable gems {'{ ... }'}</code> or <code>ItemTag gems = [...]</code>. There are {data.total} of them.</p>

	<h2 id="mms">MMS kinds</h2>
	<p>These are MMS's own, for functions, text and other things the game has no single file for.</p>
	<table>
		<thead><tr><th>Kind</th><th>Does</th></tr></thead>
		<tbody>
			{#each data.ownKinds as k (k.name)}
				<tr id="kind-{k.name}"><td><code>{k.name}</code></td><td>{#if k.doc}<Doc text={k.doc} />{/if}</td></tr>
			{/each}
		</tbody>
	</table>

	<h2 id="resources">Resource kinds</h2>
	<p>
		Each of these makes one resource in the folder shown: in the data pack, or for <code>assets/</code> folders in the resource pack. Tags of
		other registries follow the same pattern, such as <code>DamageTypeTag</code>.
	</p>
	<table>
		<thead><tr><th>Kind</th><th>Folder</th></tr></thead>
		<tbody>
			{#each data.resourceKinds as k (k.name)}
				<tr id="kind-{k.name}"><td><code>{k.name}</code></td><td><code>{k.resources ? 'assets' : 'data'}/&lt;namespace&gt;/{k.folder}/</code>{#if k.doc}<br /><Doc text={k.doc} />{/if}</td></tr>
			{/each}
		</tbody>
	</table>
</article>
