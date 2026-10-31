<script lang="ts">
	import { basicSetup } from 'codemirror';
	import { EditorState, Compartment } from '@codemirror/state';
	import { EditorView, keymap } from '@codemirror/view';
	import { indentUnit } from '@codemirror/language';
	import { setDiagnostics } from '@codemirror/lint';
	import { LSPPlugin } from '@codemirror/lsp-client';
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
					{ key: 'Shift-Alt-f', run: formatCommand, preventDefault: true }
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
		const plugin = LSPPlugin.get(v);
		const ds = engine.diagnostics[uriOf(shown)];
		if (!plugin || !ds) return;
		const doc = v.state.doc;
		v.dispatch(
			setDiagnostics(
				v.state,
				ds.map((d) => ({
					from: Math.min(plugin.fromPosition(d.range.start, doc), doc.length),
					to: Math.min(plugin.fromPosition(d.range.end, doc), doc.length),
					severity: d.severity === 1 ? 'error' : d.severity === 2 ? 'warning' : 'info',
					message: d.message
				}))
			)
		);
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
