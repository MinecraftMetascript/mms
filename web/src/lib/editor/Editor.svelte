<script lang="ts">
	import { basicSetup } from 'codemirror';
	import { EditorState, Compartment } from '@codemirror/state';
	import { EditorView, closeHoverTooltips, keymap } from '@codemirror/view';
	import { indentWithTab } from '@codemirror/commands';
	import { acceptCompletion, startCompletion } from '@codemirror/autocomplete';
	import { indentUnit } from '@codemirror/language';
	import { applyEdits, showDiagnostics } from './diagnostics';
	import { onMount, untrack } from 'svelte';
	import type { Engine } from '../engine/engine.svelte';
	import { uriOf } from '../project/project.svelte';
	import { bookHighlight, bookLanguage, mmsHighlight, mmsLanguage } from './language';
	import { semanticHighlighting } from './semantic';
	import { inlayHints } from './inlay';
	import { forEachDiagnostic } from '@codemirror/lint';
	import { LSPPlugin, findReferences, jumpToDefinition } from '@codemirror/lsp-client';
	import { inlineRename, renameInline } from './rename';
	import ContextMenu, { type MenuEntries, type MenuItem } from './ContextMenu.svelte';
	import { fixAllKeys, formatKeys } from './shortcuts';

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

	/** The right-click menu: language actions, then fixes for problems under the cursor. */
	function openMenu(v: EditorView, e: MouseEvent) {
		// Shift+right-click keeps the browser's own menu (copy, paste, spellcheck).
		if (e.shiftKey) return false;
		e.preventDefault();
		const at = v.posAtCoords({ x: e.clientX, y: e.clientY });
		const sel = v.state.selection.main;
		if (at != null && (at < sel.from || at > sel.to)) v.dispatch({ selection: { anchor: at } });
		const fixes = fixesAt(v);
		refactorsAt(v).then((refactors) => {
			const items: MenuEntries = [
				{ label: 'Go to definition', keys: 'F12', run: run(v, jumpToDefinition) },
				{ label: 'Find references', keys: 'Shift+F12', run: run(v, findReferences) },
				{ label: 'Rename symbol', keys: 'F2', run: run(v, renameInline) },
				null,
				{ label: 'Format document', keys: formatKeys, run: run(v, formatCommand) },
				{ label: 'Fix all in this file', keys: fixAllKeys, run: run(v, fixAllCommand) }
			];
			if (fixes.length) items.push(null, ...fixes);
			if (refactors.length) items.push(null, ...refactors);
			menu?.show({ x: e.clientX, y: e.clientY }, items, () => v.focus());
		});
		return true;
	}
	const run = (v: EditorView, cmd: (v: EditorView) => boolean) => () => {
		v.focus();
		cmd(v);
	};

	/** Quick fixes for the problems under the cursor. */
	function fixesAt(v: EditorView): MenuItem[] {
		const head = v.state.selection.main.head;
		const fixes: MenuItem[] = [];
		forEachDiagnostic(v.state, (d, from, to) => {
			if (head < from || head > to) return;
			for (const a of d.actions ?? []) fixes.push({ label: a.name, run: () => (v.focus(), a.apply(v, from, to)) });
		});
		return fixes;
	}

	/**
	 * Refactorings for the selection, such as extracting an inline value to
	 * its own declaration. Applying one then opens the rename box on the
	 * new declaration's name, as VS Code does.
	 */
	async function refactorsAt(v: EditorView): Promise<MenuItem[]> {
		const plugin = LSPPlugin.get(v);
		if (!plugin) return [];
		const uri = uriOf(shown);
		const sel = v.state.selection.main;
		const range = { start: plugin.toPosition(sel.from), end: plugin.toPosition(sel.to) };
		let refactors;
		try {
			refactors = await engine.refactors(uri, range);
		} catch {
			return [];
		}
		// Edits typed while the request ran would shift the positions.
		if (v.state.selection.main.from !== sel.from || uriOf(shown) !== uri) return [];
		return refactors.map((r) => ({
			label: r.title,
			run: () => {
				v.focus();
				// The declaration is inserted at a line start ahead of the value
				// it replaces, so the first edit's start is where its text begins.
				const head = r.edits[0];
				const insertAt = head ? plugin.fromPosition(head.range.start) : -1;
				applyEdits(v, r.edits);
				const nameAt = r.name && head ? head.newText.indexOf(' ' + r.name) : -1;
				if (insertAt < 0 || nameAt < 0) return;
				const at = insertAt + nameAt + 1;
				v.dispatch({ selection: { anchor: at }, scrollIntoView: true });
				renameInline(v);
			}
		}));
	}

	/** Cmd/Ctrl+. : a menu of quick fixes and refactorings at the cursor, as in VS Code. */
	function quickFixCommand(v: EditorView) {
		const at = v.coordsAtPos(v.state.selection.main.head);
		if (!at) return false;
		// The lint tooltip under the cursor would cover the menu.
		v.dispatch({ effects: closeHoverTooltips });
		const fixes = fixesAt(v);
		refactorsAt(v).then((refactors) => {
			const items: MenuEntries = [...fixes];
			if (refactors.length) {
				if (items.length) items.push(null);
				items.push(...refactors);
			}
			if (!items.length) items.push({ label: 'No quick fixes here', disabled: true, run: () => {} });
			items.push(null, { label: 'Fix all in this file', keys: fixAllKeys, run: run(v, fixAllCommand) });
			menu?.show({ x: at.left, y: at.bottom + 2 }, items, () => v.focus());
		});
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
		const book = p.endsWith('.mmsbook');
		return EditorState.create({
			doc,
			extensions: [
				basicSetup,
				indentUnit.of('  '),
				EditorView.lineWrapping,
				EditorState.tabSize.of(2),
				...(book ? [bookLanguage, bookHighlight] : [mmsLanguage, mmsHighlight]),
				engine.client.plugin(uriOf(p), book ? 'mmsbook' : 'mms'),
				semanticHighlighting((uri) => engine.semanticTokens(uri)),
				inlayHints((uri) => engine.inlayHints(uri)),
				inlineRename(),
				keymap.of([
					{ key: 'Mod-s', run: formatCommand, preventDefault: true },
					// Option+Shift+F types "Ï" on macOS, so Format is Cmd+Shift+I there.
					// Cmd+Option+F (Fix all) is fine: with Cmd held, nothing is typed.
					{ key: 'Shift-Alt-f', mac: 'Shift-Mod-i', run: formatCommand, preventDefault: true },
					{ key: 'Mod-Alt-f', run: fixAllCommand, preventDefault: true },
					{ key: 'Mod-.', run: quickFixCommand, preventDefault: true },
					// Ctrl+Space (from basicSetup) is taken by the OS on macOS (input
					// source switching) and by some IMEs elsewhere, so offer the
					// VS Code's alternatives too: Cmd/Ctrl+I and Alt/Option+Space.
					{ key: 'Mod-i', run: startCompletion, preventDefault: true },
					{ key: 'Alt-Space', run: startCompletion, preventDefault: true },
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
		// Links in hovers (the reference) open beside the editor, not over it.
		const openLink = (e: MouseEvent) => {
			const a = (e.target as Element | null)?.closest?.('.cm-tooltip a[href]') as HTMLAnchorElement | null;
			if (!a || !/^https?:/.test(a.href)) return;
			e.preventDefault();
			window.open(a.href, '_blank', 'noopener');
		};
		host.addEventListener('click', openLink);
		restoreDiagnostics(view);
		onview?.(view);
		return () => {
			host.removeEventListener('click', openLink);
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
