import { LSPPlugin } from '@codemirror/lsp-client';
import { Prec, StateEffect, StateField, type Extension } from '@codemirror/state';
import { Decoration, EditorView, keymap, showTooltip, type Command, type DecorationSet, type Tooltip } from '@codemirror/view';

interface Range {
	start: { line: number; character: number };
	end: { line: number; character: number };
}
interface Location {
	uri: string;
	range: Range;
}

interface RenameState {
	from: number;
	to: number;
	word: string;
}

const open = StateEffect.define<RenameState>();
const close = StateEffect.define<null>();
/** Highlights the other occurrences in this file while the box is open. */
const marks = StateEffect.define<{ from: number; to: number }[]>();

const tooltips = new WeakMap<RenameState, Tooltip>();

const occurrence = Decoration.mark({ class: 'cm-rename-occurrence' });

const renameField = StateField.define<{ rename: RenameState | null; marks: DecorationSet }>({
	create: () => ({ rename: null, marks: Decoration.none }),
	update(value, tr) {
		let { rename, marks: deco } = value;
		for (const e of tr.effects) {
			if (e.is(open)) {
				rename = e.value;
				deco = Decoration.none;
			} else if (e.is(close)) {
				rename = null;
				deco = Decoration.none;
			} else if (e.is(marks)) {
				deco = Decoration.set(
					e.value.filter((r) => r.to > r.from).map((r) => occurrence.range(r.from, r.to)),
					true
				);
			}
		}
		// Editing the document under the box cancels it, like VS Code.
		if (tr.docChanged && rename && !tr.effects.some((e) => e.is(open))) {
			rename = null;
			deco = Decoration.none;
		}
		return rename === value.rename && deco === value.marks ? value : { rename, marks: deco };
	},
	provide: (f) => [
		showTooltip.compute([f], (state) => {
			const r = state.field(f).rename;
			if (!r) return null;
			// One tooltip per rename, so highlighting occurrences doesn't rebuild the box.
			let t = tooltips.get(r);
			if (!t) tooltips.set(r, (t = renameTooltip(r)));
			return t;
		}),
		EditorView.decorations.from(f, (v) => v.marks)
	]
});

function renameTooltip(r: RenameState): Tooltip {
	return {
		pos: r.from,
		above: false,
		strictSide: false,
		arrow: false,
		create(view) {
			const dom = document.createElement('div');
			dom.className = 'cm-rename';
			const input = document.createElement('input');
			input.className = 'cm-rename-input';
			input.value = r.word;
			input.spellcheck = false;
			input.setAttribute('aria-label', 'New name');
			const hint = document.createElement('div');
			hint.className = 'cm-rename-hint';
			hint.textContent = 'Enter to rename · Esc to cancel';
			dom.append(input, hint);

			const cancel = () => {
				view.dispatch({ effects: close.of(null) });
				view.focus();
			};
			input.addEventListener('keydown', (e) => {
				if (e.key === 'Enter') {
					e.preventDefault();
					const name = input.value.trim();
					if (!name || name === r.word) return cancel();
					hint.textContent = 'Renaming…';
					doRename(view, r.from, name).then(
						() => cancel(),
						(err: unknown) => {
							hint.textContent = err instanceof Error ? err.message : String(err);
							hint.classList.add('cm-rename-error');
						}
					);
				} else if (e.key === 'Escape') {
					e.preventDefault();
					e.stopPropagation();
					cancel();
				}
			});
			input.addEventListener('blur', () => {
				if (view.state.field(renameField, false)?.rename === r) view.dispatch({ effects: close.of(null) });
			});
			loadPreview(view, r, hint);
			return {
				dom,
				mount() {
					input.focus();
					input.select();
				}
			};
		}
	};
}

/** Counts the occurrences the rename will touch, and highlights the ones in this file. */
function loadPreview(view: EditorView, r: RenameState, hint: HTMLElement) {
	const plugin = LSPPlugin.get(view);
	if (!plugin) return;
	plugin.client.sync();
	plugin.client
		.request<object, Location[] | null>('textDocument/references', {
			textDocument: { uri: plugin.uri },
			position: plugin.toPosition(r.from),
			context: { includeDeclaration: true }
		})
		.then((locs) => {
			if (!locs?.length || view.state.field(renameField, false)?.rename !== r) return;
			const files = new Set(locs.map((l) => l.uri)).size;
			const here = locs
				.filter((l) => l.uri === plugin.uri)
				.map((l) => ({ from: plugin.fromPosition(l.range.start), to: plugin.fromPosition(l.range.end) }));
			view.dispatch({ effects: marks.of(here) });
			const n = locs.length;
			hint.textContent = `${n} occurrence${n === 1 ? '' : 's'}${files > 1 ? ` in ${files} files` : ''} · Enter to rename · Esc to cancel`;
		})
		.catch(() => {});
}

async function doRename(view: EditorView, pos: number, newName: string) {
	const plugin = LSPPlugin.get(view);
	if (!plugin) throw new Error('No language server');
	plugin.client.sync();
	await plugin.client.withMapping(async (mapping) => {
		const res = await plugin.client.request<object, { changes?: Record<string, { range: Range; newText: string }[]> } | null>(
			'textDocument/rename',
			{ newName, position: plugin.toPosition(pos), textDocument: { uri: plugin.uri } }
		);
		if (!res?.changes) throw new Error('Nothing to rename here');
		for (const uri in res.changes) {
			const edits = res.changes[uri];
			if (!edits.length || !plugin.client.workspace.getFile(uri)) continue;
			plugin.client.workspace.updateFile(uri, {
				changes: edits.map((c) => ({
					from: mapping.mapPosition(uri, c.range.start),
					to: mapping.mapPosition(uri, c.range.end),
					insert: c.newText
				})),
				userEvent: 'rename'
			});
		}
	});
}

/** Opens an inline rename box at the symbol under the cursor. */
export const renameInline: Command = (view) => {
	const range = view.state.wordAt(view.state.selection.main.head);
	const plugin = LSPPlugin.get(view);
	if (!range || !plugin) return false;
	view.dispatch({ effects: open.of({ from: range.from, to: range.to, word: view.state.sliceDoc(range.from, range.to) }) });
	return true;
};

const theme = EditorView.baseTheme({
	'&light .cm-tooltip:has(> .cm-rename), &dark .cm-tooltip:has(> .cm-rename)': {
		border: '1px solid var(--color-accent)',
		borderRadius: '6px',
		backgroundColor: 'var(--color-panel)',
		boxShadow: '0 4px 14px rgb(0 0 0 / 0.15)',
		overflow: 'hidden'
	},
	'.cm-rename': { display: 'flex', flexDirection: 'column', minWidth: '14rem' },
	'.cm-rename-input': {
		fontFamily: 'var(--font-mono, monospace)',
		fontSize: 'inherit',
		color: 'var(--color-ink)',
		background: 'transparent',
		border: 'none',
		outline: 'none',
		padding: '3px 6px'
	},
	'.cm-rename-hint': {
		fontFamily: 'var(--font-sans, system-ui, sans-serif)',
		fontSize: '11px',
		color: 'var(--color-muted)',
		borderTop: '1px solid var(--color-line)',
		padding: '2px 6px'
	},
	'.cm-rename-error': { color: '#dc2626' },
	'.cm-rename-occurrence': {
		backgroundColor: 'color-mix(in oklab, var(--color-accent) 18%, transparent)',
		outline: '1px solid color-mix(in oklab, var(--color-accent) 45%, transparent)',
		borderRadius: '2px'
	}
});

/** VS Code-style rename: F2 opens a box over the symbol instead of a panel. */
export function inlineRename(): Extension {
	return [renameField, theme, Prec.high(keymap.of([{ key: 'F2', run: renameInline, preventDefault: true }]))];
}
