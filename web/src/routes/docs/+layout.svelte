<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { guides } from '#lib/docs/docs.ts';
	import OnThisPage from '#lib/docs/OnThisPage.svelte';
	import type { LayoutProps } from './$types';

	let { children }: LayoutProps = $props();

	const links = [
		...guides.map((g) => ({ href: resolve('/docs/[slug]', { slug: g.slug }), title: g.title })),
		{ href: resolve('/docs/reference'), title: 'Reference' }
	];
	let open = $state(false);
	let main: HTMLElement | undefined = $state();
</script>

<div class="min-h-dvh bg-bg">
	<header class="sticky top-0 z-10 flex items-center gap-3 border-b border-line bg-panel px-4 py-2">
		<a class="font-semibold tracking-tight" href={resolve('/docs')}>MMS docs</a>
		<button class="ml-auto rounded px-2 py-1 text-sm hover:bg-hover md:hidden" onclick={() => (open = !open)} aria-expanded={open}>Contents</button>
		<a class="rounded px-2 py-1 text-sm text-accent hover:bg-hover md:ml-auto" href={resolve('/')}>Open the editor</a>
	</header>
	<div class="mx-auto flex max-w-7xl gap-8 px-4">
		<nav class="{open ? 'block' : 'hidden'} fixed inset-x-0 top-11 z-10 border-b border-line bg-panel p-3 md:sticky md:top-14 md:block md:h-fit md:w-52 md:shrink-0 md:border-0 md:bg-transparent md:p-0 md:pt-6" aria-label="Docs">
			<ul class="space-y-0.5 text-sm">
				{#each links as l (l.href)}
					<li>
						<a
							class="block rounded px-2 py-1.5 hover:bg-hover {page.url.pathname === l.href || (l.href === links[0].href && page.url.pathname === resolve('/docs')) ? 'bg-selected font-medium' : 'text-muted'}"
							href={l.href}
							onclick={() => (open = false)}>{l.title}</a
						>
					</li>
				{/each}
			</ul>
		</nav>
		<main class="min-w-0 flex-1 py-6" bind:this={main}>
			{@render children()}
		</main>
		<aside class="hidden xl:sticky xl:top-14 xl:block xl:h-fit xl:w-56 xl:shrink-0 xl:pt-6">
			<OnThisPage root={main} />
		</aside>
	</div>
</div>
