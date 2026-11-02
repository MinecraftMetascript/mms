<script lang="ts">
	// Credits for the upstream work the editor is built on, Deepslate first.
	import { asset } from '$app/paths';
	import { credits } from '../credits';

	let dialog: HTMLDialogElement;

	export function open() {
		dialog.showModal();
	}
</script>

<dialog bind:this={dialog} class="m-auto w-[min(34rem,calc(100vw-2rem))] rounded-lg border border-line bg-panel p-0 text-ink shadow-xl backdrop:bg-black/40" aria-labelledby="about-title">
	<div class="flex flex-col gap-3 p-4 text-sm">
		<h2 id="about-title" class="font-semibold">About MMS</h2>
		<p class="text-muted">
			Minecraft Metascript turns one source into a whole data pack. It stands on other people's work, with thanks to:
		</p>
		<ul class="flex flex-col gap-2.5" aria-label="Credits">
			{#each credits as c (c.name)}
				<li>
					<a class="font-medium text-accent hover:underline" href={c.url} target="_blank" rel="noopener">{c.name}</a>
					<span class="text-muted">by</span>
					<a class="hover:underline" href={c.byUrl} target="_blank" rel="noopener">{c.by}</a>{#if c.license}<span class="text-muted">{' · '}{c.license}</span>{/if}
					<span class="block text-xs text-muted">{c.what}</span>
				</li>
			{/each}
		</ul>
		<p class="text-xs text-muted">
			Minecraft is a trademark of Mojang. MMS is not affiliated with Mojang or Microsoft.
			<a class="text-accent hover:underline" href={asset('third-party-licenses.txt')} target="_blank">Third-party licenses</a> ·
			<a class="text-accent hover:underline" href="https://github.com/MinecraftMetascript/mms" target="_blank" rel="noopener">Source on GitHub</a>
		</p>
		<div class="flex justify-end">
			<button class="rounded px-3 py-1 text-ink hover:bg-hover" onclick={() => dialog.close()}>Close</button>
		</div>
	</div>
</dialog>
