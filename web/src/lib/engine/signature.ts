// CodeMirror's LSP client draws signature help with the signature's own
// documentation only, never the active parameter's. Editors like VS Code
// show both, so the server keeps them apart; for the web editor this
// folds the active parameter's text into the signature's documentation,
// first, since it is what the person is typing.

interface MarkupContent {
	kind: string;
	value: string;
}
type Doc = string | MarkupContent | undefined;

interface ParameterInformation {
	label: string | [number, number];
	documentation?: Doc;
}
interface SignatureInformation {
	label: string;
	documentation?: Doc;
	parameters?: ParameterInformation[];
	activeParameter?: number;
}
export interface SignatureHelp {
	signatures: SignatureInformation[];
	activeSignature?: number;
	activeParameter?: number;
}

const text = (d: Doc) => (d == null ? '' : typeof d === 'string' ? d : d.value);

/** Returns help with the active parameter's documentation folded in. */
export function withParameterDoc(help: SignatureHelp): SignatureHelp {
	const sig = help.signatures?.[help.activeSignature ?? 0];
	if (!sig) return help;
	const n = sig.activeParameter ?? help.activeParameter;
	const param = n == null ? undefined : sig.parameters?.[n];
	const pdoc = text(param?.documentation);
	if (!param || !pdoc) return help;
	const name = typeof param.label === 'string' ? param.label.split(/[?:= ]/)[0] : sig.label.slice(...param.label);
	const value = [`**${name}**: ${pdoc}`, text(sig.documentation)].filter(Boolean).join('\n\n');
	return {
		...help,
		signatures: help.signatures.map((s) => (s === sig ? { ...s, documentation: { kind: 'markdown', value } } : s))
	};
}

/**
 * Wraps a transport's incoming messages so signature help responses carry
 * the active parameter's documentation. ids are those of signatureHelp
 * requests sent so far.
 */
export class SignatureShim {
	private ids = new Set<string | number>();

	/** Notes an outgoing message. */
	sent(message: string) {
		if (!message.includes('"textDocument/signatureHelp"')) return;
		try {
			const m = JSON.parse(message);
			if (m.method === 'textDocument/signatureHelp' && m.id != null) this.ids.add(m.id);
		} catch {
			// Not JSON: nothing to track.
		}
	}

	/** Rewrites an incoming message if it answers signature help. */
	received(message: string): string {
		if (!this.ids.size) return message;
		try {
			const m = JSON.parse(message);
			if (m.id == null || !this.ids.delete(m.id) || !m.result) return message;
			return JSON.stringify({ ...m, result: withParameterDoc(m.result) });
		} catch {
			return message;
		}
	}
}
