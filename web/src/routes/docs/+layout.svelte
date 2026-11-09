<script lang="ts">
	import { asset, resolve } from '$app/paths';
	import { page } from '$app/state';
	import Logo from '#lib/components/Logo.svelte';
	import OnThisPage from '#lib/docs/OnThisPage.svelte';
	import Search from '#lib/docs/Search.svelte';
	import type { NavLink } from '#lib/docs/nav.ts';
	import { credits } from '#lib/credits.ts';
	import '#lib/docs/reference.css';
	import type { LayoutProps } from './$types';

	let { children, data }: LayoutProps = $props();

	const base = resolve('/').replace(/\/$/, '');
	const here = $derived(page.url.pathname.replace(/\/$/, ''));
	const tab = $derived(/\/docs\/reference(\/|$)/.test(here) ? 'reference' : 'docs');
	const nav = $derived(tab === 'reference' ? data.referenceNav : data.docsNav);
	const current = (l: NavLink) => here === base + l.path;
	const within = (l: NavLink) => current(l) || here.startsWith(base + l.path + '/');
	let open = $state(false);
	let main: HTMLElement | undefined = $state();
</script>

<div class="min-h-dvh bg-bg">
	<header class="sticky top-0 z-20 flex items-center gap-1 border-b border-line bg-panel px-4 py-2 sm:gap-3">
		<a class="mr-2 flex items-center gap-2 font-semibold tracking-tight" href={resolve('/docs')}><Logo class="h-6 w-auto" /></a>
		<nav class="flex gap-1 text-sm" aria-label="Sections">
			<a class="rounded px-2 py-1 hover:bg-hover {tab === 'docs' ? 'bg-selected font-medium' : 'text-muted'}" href={resolve('/docs')}>Docs</a>
			<a class="rounded px-2 py-1 hover:bg-hover {tab === 'reference' ? 'bg-selected font-medium' : 'text-muted'}" href={resolve('/docs/reference')}>Reference</a>
		</nav>
		<div class="ml-auto flex items-center gap-2">
			<Search />
			<button class="rounded px-2 py-1 text-sm hover:bg-hover md:hidden" onclick={() => (open = !open)} aria-expanded={open}>Contents</button>
			<a class="hidden rounded px-2 py-1 text-sm text-accent hover:bg-hover sm:inline" href={resolve('/')}>Open the editor</a>
		</div>
	</header>
	<div class="mx-auto flex max-w-7xl gap-8 px-4">
		<nav
			class="{open ? 'block' : 'hidden'} fixed inset-x-0 top-11 z-10 max-h-[calc(100dvh-2.75rem)] overflow-y-auto border-b border-line bg-panel p-3 md:sticky md:top-14 md:block md:max-h-[calc(100dvh-3.5rem)] md:w-56 md:shrink-0 md:border-0 md:bg-transparent md:p-0 md:pt-6 md:pb-6"
			aria-label={tab === 'reference' ? 'Reference' : 'Docs'}
		>
			{#each nav as g (g.title)}
				<h2 class="mt-4 mb-1 px-2 text-xs font-semibold tracking-wide text-muted uppercase first:mt-0">{g.title}</h2>
				<ul class="space-y-0.5 text-sm">
					{#each g.links as l (l.path)}
						<li>
							<a class="block rounded px-2 py-1.5 hover:bg-hover {current(l) ? 'bg-selected font-medium' : within(l) ? 'font-medium' : 'text-muted'}" href={base + l.path} aria-current={current(l) ? 'page' : undefined} onclick={() => (open = false)}>{l.title}</a>
							{#if l.children && within(l)}
								<ul class="my-0.5 ml-2 space-y-0.5 border-l border-line pl-2">
									{#each l.children as c (c.path)}
										<li><a class="block rounded px-2 py-1 hover:bg-hover {current(c) ? 'bg-selected font-medium' : 'text-muted'}" href={base + c.path} aria-current={current(c) ? 'page' : undefined} onclick={() => (open = false)}>{c.title}</a></li>
									{/each}
								</ul>
							{/if}
						</li>
					{/each}
				</ul>
			{/each}
		</nav>
		<main class="min-w-0 flex-1 py-6" bind:this={main}>
			{@render children()}
		</main>
		<aside class="hidden xl:sticky xl:top-14 xl:block xl:h-fit xl:w-56 xl:shrink-0 xl:pt-6">
			<OnThisPage root={main} />
		</aside>
	</div>
	<footer class="border-t border-line px-4 py-4 text-xs text-muted">
		<p class="mx-auto max-w-7xl">
			Built on
			{#each credits as c, i (c.name)}<a class="text-accent hover:underline" href={c.url} target="_blank" rel="noopener">{c.name}</a> by <a class="hover:underline" href={c.byUrl} target="_blank" rel="noopener">{c.by}</a>{i < credits.length - 2 ? ', ' : i === credits.length - 2 ? ' and ' : '. '}{/each}
			Previews are rendered with Deepslate.
			<a class="text-accent hover:underline" href={asset('third-party-licenses.txt')} target="_blank">Third-party licenses</a>.
			For AI tools: <a class="text-accent hover:underline" href="{base}/llms.txt">llms.txt</a>, <a class="text-accent hover:underline" href="{base}/llms-full.txt">llms-full.txt</a>, <a class="text-accent hover:underline" href="{base}/docs/reference.yaml">reference.yaml</a>.
		</p>
	</footer>
</div>
