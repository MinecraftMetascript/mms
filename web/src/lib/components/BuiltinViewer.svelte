<script lang="ts">
	// A read-only view of a file that comes with the language, where go to
	// definition leads: a prelude file (a built-in Template, Blueprint or
	// Type) or one of vanilla's own data pack JSON files (a vanilla noise,
	// loot table or tag).
	import { basicSetup } from 'codemirror';
	import { EditorState } from '@codemirror/state';
	import { EditorView } from '@codemirror/view';
	import { json } from '@codemirror/lang-json';
	import { tick } from 'svelte';
	import { mmsHighlight, mmsLanguage } from '#lib/editor/language.ts';

	let dialog: HTMLDialogElement;
	let host: HTMLDivElement;
	let title = $state('');
	let vanilla = $state(false);
	let view: EditorView | null = null;

	/** Shows the file and resolves to its view, for the caller to select a range in. */
	export async function show(uri: string, text: string): Promise<EditorView> {
		vanilla = uri.startsWith('mms-vanilla:');
		title = uri.replace(/^mms-(prelude|vanilla):\/*/, '');
		view?.destroy();
		dialog.showModal();
		await tick();
		view = new EditorView({
			state: EditorState.create({
				doc: text,
				extensions: [
					basicSetup,
					EditorState.readOnly.of(true),
					EditorView.editable.of(false),
					EditorView.lineWrapping,
					vanilla ? json() : [mmsLanguage, mmsHighlight]
				]
			}),
			parent: host
		});
		return view;
	}

	function close() {
		dialog.close();
	}
</script>

<dialog
	bind:this={dialog}
	class="m-auto h-[min(40rem,calc(100vh-2rem))] w-[min(56rem,calc(100vw-2rem))] rounded-lg border border-line bg-panel p-0 text-ink shadow-xl backdrop:bg-black/40"
	aria-labelledby="builtin-title"
	onclose={() => {
		view?.destroy();
		view = null;
	}}
>
	<div class="flex h-full flex-col">
		<div class="flex items-center gap-2 border-b border-line px-4 py-2 text-sm">
			<h2 id="builtin-title" class="min-w-0 truncate font-semibold">{vanilla ? 'Vanilla' : 'Prelude'}: <code>{title}</code></h2>
			<span class="shrink-0 text-xs text-muted">{vanilla ? "the game's own data, read-only" : 'built into the language, read-only'}</span>
			<button class="ml-auto shrink-0 rounded px-3 py-1 hover:bg-hover" onclick={close}>Close</button>
		</div>
		<div class="editor min-h-0 flex-1 overflow-auto" data-testid="builtin-view" bind:this={host}></div>
	</div>
</dialog>
