<script lang="ts">
	import { resolve } from '$app/paths';
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
	const base = resolve('/').replace(/\/$/, '');
	const c = $derived(data.counts);
</script>

<svelte:head><title>Reference · MMS docs</title></svelte:head>

<article class="prose-mms reference">
	<h1>Reference</h1>
	<p>
		Everything MMS knows about Minecraft {c.version}, for looking things up. The generated pages come from the same schema and rules the
		checker uses, so they match what the editor's completion and hover show. To learn MMS, start with the <a href="{base}/docs/start/getting-started">Docs</a>
		instead.
	</p>
	{#if data.language.length}
		<h2 id="language">Language</h2>
		<ul class="cards">
			{#each data.language as p (p.path)}
				<li><a href="{base}/docs/{p.path}"><strong>{p.title}</strong>{#if p.summary}<span>{p.summary}</span>{/if}</a></li>
			{/each}
		</ul>
	{/if}
	<h2 id="generated">Generated</h2>
	<ul class="cards">
		<li><a href={resolve('/docs/reference/kinds')}><strong>Declaration kinds</strong><span>{c.kinds} kinds, and the folder each one writes to.</span></a></li>
		<li><a href={resolve('/docs/reference/helpers')}><strong>Helpers</strong><span>{c.helpers} helpers in {c.groups} groups, each building one typed object.</span></a></li>
		<li><a href={resolve('/docs/reference/prelude')}><strong>Templates and forms</strong><span>The prelude's templates and types, and what each operator means by type.</span></a></li>
		<li><a href={resolve('/docs/reference/types')}><strong>Types</strong><span>{c.types} named types that parameters refer to.</span></a></li>
		<li><a href={resolve('/docs/reference/lint')}><strong>Lint rules</strong><span>{c.rules} rules, their default severity and what they catch.</span></a></li>
	</ul>
	<h2 id="for-ai-tools">For AI tools</h2>
	<p>The same reference in forms that language models and scripts read easily:</p>
	<ul>
		<li><a href="{base}/docs/llms.txt"><code>llms.txt</code></a>: what MMS is, and a link to every page.</li>
		<li><a href="{base}/docs/llms-full.txt"><code>llms-full.txt</code></a>: every page and the whole reference as one Markdown file.</li>
		<li><a href="{base}/docs/reference.yaml"><code>reference.yaml</code></a>: kinds, helpers, types and lint rules as data.</li>
	</ul>
</article>
