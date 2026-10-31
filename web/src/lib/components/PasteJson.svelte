<script lang="ts">
	import type { Engine, ResourceKind } from '../engine/engine.svelte';

	interface Props {
		engine: Engine;
		/** The file the declaration goes into. */
		uri: string;
		oninsert: (text: string) => void;
	}
	let { engine, uri, oninsert }: Props = $props();

	let dialog: HTMLDialogElement;
	let kinds = $state<ResourceKind[]>([]);
	let kind = $state('');
	let name = $state('');
	let json = $state('');
	let error = $state('');
	let busy = $state(false);

	export async function open() {
		error = '';
		dialog.showModal();
		if (!kinds.length) {
			try {
				kinds = await engine.resourceKinds();
				if (!kind) kind = kinds.find((k) => k.kind === 'LootTable')?.kind ?? kinds[0]?.kind ?? '';
			} catch (e) {
				error = `The engine is not ready: ${e instanceof Error ? e.message : e}`;
			}
		}
	}

	const nameOK = $derived(/^([a-z0-9_.-]+:)?[a-z0-9_.-]+(\/[a-z0-9_.-]+)*$/.test(name.trim()));

	async function convert(e: SubmitEvent) {
		e.preventDefault();
		error = '';
		try {
			JSON.parse(json);
		} catch (err) {
			error = `That is not valid JSON: ${err instanceof Error ? err.message : err}`;
			return;
		}
		if (!kinds.some((k) => k.kind === kind) || !nameOK) return;
		busy = true;
		try {
			const text = await engine.convertJSON(uri, kind, name.trim(), json);
			oninsert(text);
			dialog.close();
			json = '';
			name = '';
		} catch (err) {
			error = err instanceof Error ? err.message : String(err);
		} finally {
			busy = false;
		}
	}
</script>

<dialog bind:this={dialog} class="m-auto w-[min(40rem,calc(100vw-2rem))] rounded-lg border border-line bg-panel p-0 text-ink shadow-xl backdrop:bg-black/40">
	<form class="flex flex-col gap-3 p-4" onsubmit={convert}>
		<h2 class="font-semibold">Paste JSON as MMS</h2>
		<p class="text-sm text-muted">Paste a data pack file. It becomes a declaration at the end of the open file.</p>
		<div class="flex flex-wrap gap-3">
			<label class="flex min-w-40 flex-1 flex-col gap-1 text-sm">
				Kind
				<select class="rounded border border-line bg-panel px-2 py-1" bind:value={kind} aria-label="Kind">
					{#each kinds as k (k.kind)}<option value={k.kind}>{k.kind} ({k.folder})</option>{/each}
				</select>
			</label>
			<label class="flex min-w-40 flex-1 flex-col gap-1 text-sm">
				Name
				<input class="rounded border bg-panel px-2 py-1 font-mono {name && !nameOK ? 'border-red-500' : 'border-line'}" bind:value={name} placeholder="blocks/ruby_ore" aria-label="Name" />
			</label>
		</div>
		<label class="flex flex-col gap-1 text-sm">
			JSON
			<textarea class="h-64 rounded border border-line bg-panel p-2 font-mono text-xs" bind:value={json} spellcheck="false" aria-label="JSON" placeholder={'{ "pools": [] }'}></textarea>
		</label>
		{#if error}<p class="text-sm text-red-600" role="alert">{error}</p>{/if}
		<div class="flex justify-end gap-2">
			<button type="button" class="rounded px-3 py-1 text-sm hover:bg-hover" onclick={() => dialog.close()}>Cancel</button>
			<button type="submit" class="rounded bg-accent px-3 py-1 text-sm font-medium text-white disabled:opacity-50" disabled={busy || !kind || !nameOK || !json.trim()}>Convert</button>
		</div>
	</form>
</dialog>
