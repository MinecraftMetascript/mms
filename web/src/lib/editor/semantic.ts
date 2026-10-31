import { LSPPlugin } from '@codemirror/lsp-client';
import { RangeSetBuilder } from '@codemirror/state';
import { Decoration, ViewPlugin, type DecorationSet, type EditorView, type ViewUpdate } from '@codemirror/view';
import { tokenTypes } from '../engine/engine.svelte';

const marks = tokenTypes.map((t) => Decoration.mark({ class: `cm-sem-${t}` }));
const declMarks = tokenTypes.map((t) => Decoration.mark({ class: `cm-sem-${t} cm-sem-declaration` }));

/**
 * Semantic highlighting from textDocument/semanticTokens/full, refreshed
 * shortly after each edit. Decorations are mapped through edits while a
 * request is in flight so colours do not jump.
 */
export function semanticHighlighting(fetch: (uri: string) => Promise<number[]>, delay = 150) {
	return ViewPlugin.fromClass(
		class {
			decorations: DecorationSet = Decoration.none;
			timer: ReturnType<typeof setTimeout> | null = null;
			seq = 0;
			constructor(readonly view: EditorView) {
				this.schedule(10);
			}
			update(u: ViewUpdate) {
				if (u.docChanged) {
					this.decorations = this.decorations.map(u.changes);
					this.schedule(delay);
				}
			}
			schedule(ms: number) {
				if (this.timer) clearTimeout(this.timer);
				this.timer = setTimeout(() => this.run(), ms);
			}
			async run() {
				const plugin = LSPPlugin.get(this.view);
				if (!plugin) return;
				const seq = ++this.seq;
				let data: number[];
				try {
					await plugin.client.initializing;
					data = await fetch(plugin.uri);
				} catch {
					return;
				}
				if (seq !== this.seq || !LSPPlugin.get(this.view)) return;
				const doc = plugin.syncedDoc;
				const b = new RangeSetBuilder<Decoration>();
				let line = 0;
				let char = 0;
				for (let i = 0; i + 4 < data.length; i += 5) {
					line += data[i];
					char = data[i] === 0 ? char + data[i + 1] : data[i + 1];
					if (line >= doc.lines) break;
					const from = plugin.fromPosition({ line, character: char }, doc);
					const to = plugin.fromPosition({ line, character: char + data[i + 2] }, doc);
					const set = data[i + 4] & 1 ? declMarks : marks;
					const m = set[data[i + 3]];
					if (m && to > from) b.add(from, to, m);
				}
				// Map from the synced document to the editor's current one.
				this.decorations = b.finish().map(plugin.unsyncedChanges);
				this.view.dispatch({});
			}
			destroy() {
				if (this.timer) clearTimeout(this.timer);
				this.seq++;
			}
		},
		{ decorations: (v) => v.decorations }
	);
}
