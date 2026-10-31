import { setDiagnostics, type Action, type Diagnostic as LintDiagnostic } from '@codemirror/lint';
import { LSPPlugin } from '@codemirror/lsp-client';
import type { EditorView } from '@codemirror/view';
import type { Diagnostic, Fix, TextEdit } from '../engine/engine.svelte';

const severities = ['error', 'error', 'warning', 'info', 'hint'] as const;

type Pos = TextEdit['range']['start'];

/**
 * Maps a server position to the view. Diagnostics can describe a document
 * the client has since replaced, so out-of-range lines are clamped (the
 * next publish corrects them) rather than thrown on.
 */
function mapper(plugin: LSPPlugin, len: number) {
	const doc = plugin.syncedDoc;
	return (p: Pos) => {
		const off = p.line + 1 > doc.lines ? doc.length : plugin.fromPosition(p, doc);
		return Math.min(plugin.unsyncedChanges.mapPos(Math.min(off, doc.length)), len);
	};
}

/**
 * Applies LSP edits to a view. Positions are in the document the server
 * last saw, mapped through edits typed since.
 */
export function applyEdits(view: EditorView, edits: TextEdit[]) {
	const plugin = LSPPlugin.get(view);
	if (!plugin) return;
	const pos = mapper(plugin, view.state.doc.length);
	view.dispatch({
		changes: edits.map((e) => ({ from: pos(e.range.start), to: pos(e.range.end), insert: e.newText })),
		userEvent: 'input.fix'
	});
}

function actions(fixes: Fix[] | undefined): Action[] {
	return (fixes ?? []).map((f) => ({
		name: f.title,
		apply: (view: EditorView) => applyEdits(view, f.edits)
	}));
}

/** Shows server diagnostics, with their fixes as actions, in a view. */
export function showDiagnostics(view: EditorView, ds: Diagnostic[]) {
	const plugin = LSPPlugin.get(view);
	if (!plugin) return;
	const pos = mapper(plugin, view.state.doc.length);
	const out: LintDiagnostic[] = ds.map((d) => ({
		from: pos(d.range.start),
		to: pos(d.range.end),
		severity: severities[d.severity ?? 1] ?? 'error',
		message: d.message,
		source: d.code,
		markClass: d.tags?.includes(1) ? 'cm-unnecessary' : undefined,
		actions: actions(d.data?.fixes)
	}));
	view.dispatch(setDiagnostics(view.state, out));
}
