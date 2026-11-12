// Reading the density function spline at a place in a project file for
// the spline editor, and writing the editor's edits back. A spline is
// read from the text the server last saw, which the file may have moved
// past, so each read keeps what maps its ranges to the text now.
import { LSPPlugin } from '@codemirror/lsp-client';
import { ChangeSet, type ChangeSpec, type Text } from '@codemirror/state';
import type { EditorView, ViewUpdate } from '@codemirror/view';
import type { SplineAt, SplineEdit } from '@mms/preview';
import type { Engine } from './engine.svelte';
import type { ProjectFile } from './workspace';

/**
 * Where a spline was read: the server's text it was read from (doc), the
 * edits made to the file since (since), the text after them (now), and
 * the offset in doc it was read at.
 */
export interface SplineTrack {
	uri: string;
	doc: Text;
	since: ChangeSet;
	now: Text;
	at: number;
}

/** A file's text: the server's, the changes since, and the text now, plus its editor if it is open. */
function fileText(engine: Engine, uri: string): { view: EditorView | null; synced: Text; since: ChangeSet; now: Text } | null {
	const f = engine.workspace.getFile(uri) as ProjectFile | null;
	if (!f) return null;
	const view = f.getView();
	const plugin = view && LSPPlugin.get(view);
	if (view && plugin) return { view, synced: plugin.syncedDoc, since: plugin.unsyncedChanges, now: view.state.doc };
	if (f.pending) return { view: null, synced: f.pending.prevDoc, since: f.pending.changes, now: f.doc };
	return { view: null, synced: f.doc, since: ChangeSet.empty(f.doc.length), now: f.doc };
}

const same = (a: Text, b: Text) => a === b || a.eq(b);

/**
 * The spline at an offset in a file's text now, and where it was read.
 * Undefined when the text changed while the server answered, so the
 * answer is about text that is gone; a new analysis follows anyway.
 */
export async function readSpline(engine: Engine, uri: string, pos: number): Promise<{ spline: SplineAt | null; track: SplineTrack } | undefined> {
	const t = fileText(engine, uri);
	if (!t) return undefined;
	const at = Math.min(t.since.invertedDesc.mapPos(Math.min(pos, t.now.length)), t.synced.length);
	const line = t.synced.lineAt(at);
	const spline = await engine.spline(uri, { line: line.number - 1, character: at - line.from });
	const after = fileText(engine, uri);
	if (!after || !same(after.now, t.now)) return undefined;
	return { spline, track: { uri, doc: t.synced, since: t.since, now: t.now, at } };
}

/** Where a spline was read, as an offset in its file's text now, or null if the file changed in ways not followed. */
export function trackPos(engine: Engine, track: SplineTrack): number | null {
	const t = fileText(engine, track.uri);
	return t && same(t.now, track.now) ? track.since.mapPos(track.at) : null;
}

/** Follows an editor change to the file a spline was read from. */
export function followEdit(track: SplineTrack, uri: string, u: ViewUpdate) {
	if (!u.docChanged || track.uri !== uri || !same(track.now, u.startState.doc)) return;
	track.since = track.since.compose(u.changes);
	track.now = u.state.doc;
}

/**
 * Writes the spline editor's edits to the file. An open file takes them
 * through its editor, so they can be undone; one that isn't open (a
 * pinned spline's, with another file in the editor) takes them directly.
 */
export function applySplineEdits(engine: Engine, track: SplineTrack, edits: SplineEdit[]) {
	const t = fileText(engine, track.uri);
	if (!t || !same(t.now, track.now)) return;
	const doc = track.doc;
	const pos = (p: SplineEdit['range']['start']) => {
		const line = doc.line(Math.min(p.line + 1, doc.lines));
		return track.since.mapPos(Math.min(line.from + p.character, line.to));
	};
	const changes: ChangeSpec = edits.map((e) => ({ from: pos(e.range.start), to: pos(e.range.end), insert: e.newText }));
	if (t.view) {
		t.view.dispatch({ changes, userEvent: 'input' });
		return;
	}
	const set = ChangeSet.of(changes, t.now.length);
	engine.workspace.updateFile(track.uri, { changes: set });
	track.since = track.since.compose(set);
	track.now = set.apply(t.now);
}
