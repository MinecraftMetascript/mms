import { LSPPlugin } from '@codemirror/lsp-client';
import { RangeSetBuilder } from '@codemirror/state';
import { Decoration, ViewPlugin, WidgetType, type DecorationSet, type EditorView, type ViewUpdate } from '@codemirror/view';

/** An LSP inlay hint with a plain string label. */
export interface InlayHint {
	position: { line: number; character: number };
	label: string;
	tooltip?: string;
}

class HintWidget extends WidgetType {
	constructor(
		readonly label: string,
		readonly tooltip: string
	) {
		super();
	}
	eq(o: HintWidget) {
		return o.label === this.label && o.tooltip === this.tooltip;
	}
	toDOM() {
		const el = document.createElement('span');
		el.className = 'cm-inlay-hint';
		el.textContent = this.label;
		if (this.tooltip) el.title = this.tooltip;
		return el;
	}
	ignoreEvent() {
		return false;
	}
}

/**
 * Inlay hints from textDocument/inlayHint, such as the base item after a
 * custom item used as a plain item id. Refreshed shortly after each edit,
 * like semantic highlighting.
 */
export function inlayHints(fetch: (uri: string) => Promise<InlayHint[]>, delay = 200) {
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
				} else if (u.focusChanged && u.view.hasFocus) {
					// Another file may have changed what the hints say.
					this.schedule(10);
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
				let hints: InlayHint[];
				try {
					await plugin.client.initializing;
					hints = await fetch(plugin.uri);
				} catch {
					return;
				}
				if (seq !== this.seq || !LSPPlugin.get(this.view)) return;
				const doc = plugin.syncedDoc;
				const b = new RangeSetBuilder<Decoration>();
				for (const h of hints) {
					if (h.position.line >= doc.lines) continue;
					const at = plugin.fromPosition(h.position, doc);
					b.add(at, at, Decoration.widget({ widget: new HintWidget(h.label, h.tooltip ?? ''), side: 1 }));
				}
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
