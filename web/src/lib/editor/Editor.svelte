<script lang="ts">
	import { basicSetup } from 'codemirror';
	import { EditorState, Compartment } from '@codemirror/state';
	import { EditorView, keymap } from '@codemirror/view';
	import { indentUnit } from '@codemirror/language';
	import { applyEdits, showDiagnostics } from './diagnostics';
	import { onMount, untrack } from 'svelte';
	import type { Engine } from '../engine/engine.svelte';
	import { uriOf } from '../project/project.svelte';
	import { mmsHighlight, mmsLanguage } from './language';
	import { semanticHighlighting } from './semantic';

	interface Props {
		engine: Engine;
		path: string;
		text: string;
		onchange: (path: string, text: string) => void;
		onview?: (view: EditorView) => void;
		/** Extra extensions, such as a cursor listener. */
		extra?: import('@codemirror/state').Extension;
	}
	let { engine, path, text, onchange, onview, extra = [] }: Props = $props();

	let host: HTMLDivElement;
	let view: EditorView | null = null;

	/** Applies every safe fix in the open file. */
	export function fixAll() {
		if (view) fixAllCommand(view);
	}
	let shown = '';
	const theme = new Compartment();

	function formatCommand(v: EditorView) {
		const uri = uriOf(shown);
		engine.format(uri).then((out) => {
			if (out != null && out !== v.state.doc.toString()) {
				v.dispatch({ changes: { from: 0, to: v.state.doc.length, insert: out }, userEvent: 'format' });
			}
		});
		return true;
	}

	function stateFor(p: string, doc: string) {
		return EditorState.create({
			doc,
			extensions: [
				basicSetup,
				indentUnit.of('  '),
				EditorView.lineWrapping,
				EditorState.tabSize.of(2),
				mmsLanguage,
				mmsHighlight,
				engine.client.plugin(uriOf(p), 'mms'),
				semanticHighlighting((uri) => engine.semanticTokens(uri)),
				keymap.of([
					{ key: 'Mod-s', run: formatCommand, preventDefault: true },
					{ key: 'Shift-Alt-f', run: formatCommand, preventDefault: true },
					{ key: 'Mod-Alt-f', run: fixAllCommand, preventDefault: true }
				]),
				theme.of([]),
				extra,
				EditorView.updateListener.of((u) => {
					if (u.docChanged) onchange(p, u.state.doc.toString());
				})
			]
		});
	}

	/** Shows diagnostics the server sent while another file was open. */
	function restoreDiagnostics(v: EditorView) {
		const ds = engine.diagnostics[uriOf(shown)];
		if (ds) showDiagnostics(v, ds);
	}

	function fixAllCommand(v: EditorView) {
		engine.fixAll(uriOf(shown)).then((edits) => {
			if (edits.length) applyEdits(v, edits);
		});
		return true;
	}

	function show(p: string, doc: string) {
		if (!view) return;
		engine.client.sync();
		shown = p;
		view.setState(stateFor(p, doc));
		restoreDiagnostics(view);
		onview?.(view);
	}

	onMount(() => {
		shown = path;
		view = new EditorView({ state: stateFor(path, text), parent: host });
		restoreDiagnostics(view);
		onview?.(view);
		return () => {
			view?.destroy();
			view = null;
		};
	});

	$effect(() => {
		const p = path;
		if (view && p !== shown) untrack(() => show(p, text));
	});
</script>

<div class="editor h-full min-h-0 overflow-hidden" bind:this={host}></div>
