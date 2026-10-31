<script lang="ts">
	import { basicSetup } from 'codemirror';
	import { EditorState, Compartment } from '@codemirror/state';
	import { EditorView, keymap } from '@codemirror/view';
	import { indentWithTab } from '@codemirror/commands';
	import { acceptCompletion } from '@codemirror/autocomplete';
	import { indentUnit } from '@codemirror/language';
	import { applyEdits, showDiagnostics } from './diagnostics';
	import { onMount, untrack } from 'svelte';
	import type { Engine } from '../engine/engine.svelte';
	import { uriOf } from '../project/project.svelte';
	import { mmsHighlight, mmsLanguage } from './language';
	import { semanticHighlighting } from './semantic';
	import { forEachDiagnostic } from '@codemirror/lint';
	import { findReferences, jumpToDefinition } from '@codemirror/lsp-client';
	import { inlineRename, renameInline } from './rename';
	import ContextMenu, { type MenuEntries } from './ContextMenu.svelte';

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
	/** Formats the open file with the engine's formatter. */
	export function format() {
		if (view) formatCommand(view);
	}
	let shown = '';
	let menu: ContextMenu | undefined = $state();

	const mac = typeof navigator !== 'undefined' && /Mac|iP(hone|ad)/.test(navigator.platform);
	const mod = mac ? '⌘' : 'Ctrl';
	const alt = mac ? '⌥' : 'Alt';

	/** The right-click menu: language actions, then fixes for problems under the cursor. */
	function openMenu(v: EditorView, e: MouseEvent) {
		// Shift+right-click keeps the browser's own menu (copy, paste, spellcheck).
		if (e.shiftKey) return false;
		e.preventDefault();
		const at = v.posAtCoords({ x: e.clientX, y: e.clientY });
		const sel = v.state.selection.main;
		if (at != null && (at < sel.from || at > sel.to)) v.dispatch({ selection: { anchor: at } });
		const head = v.state.selection.main.head;
		const run = (cmd: (v: EditorView) => boolean) => () => {
			v.focus();
			cmd(v);
		};
		const fixes: MenuEntries = [];
		forEachDiagnostic(v.state, (d, from, to) => {
			if (head < from || head > to) return;
			for (const a of d.actions ?? []) fixes.push({ label: a.name, run: () => (v.focus(), a.apply(v, from, to)) });
		});
		const items: MenuEntries = [
			{ label: 'Go to definition', keys: 'F12', run: run(jumpToDefinition) },
			{ label: 'Find references', keys: 'Shift+F12', run: run(findReferences) },
			{ label: 'Rename symbol', keys: 'F2', run: run(renameInline) },
			null,
			{ label: 'Format document', keys: `Shift+${alt}+F`, run: run(formatCommand) },
			{ label: 'Fix all in this file', keys: `${mod}+${alt}+F`, run: run(fixAllCommand) }
		];
		if (fixes.length) items.push(null, ...fixes);
		menu?.show({ x: e.clientX, y: e.clientY }, items, () => v.focus());
		return true;
	}
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
				inlineRename(),
				keymap.of([
					{ key: 'Mod-s', run: formatCommand, preventDefault: true },
					{ key: 'Shift-Alt-f', run: formatCommand, preventDefault: true },
					{ key: 'Mod-Alt-f', run: fixAllCommand, preventDefault: true },
					// Tab accepts an open completion, otherwise indents; Shift+Tab unindents.
					// Escape then Tab still moves focus out of the editor.
					{ key: 'Tab', run: acceptCompletion },
					indentWithTab
				]),
				EditorView.domEventHandlers({ contextmenu: (e, v) => openMenu(v, e) }),
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
<ContextMenu bind:this={menu} />
