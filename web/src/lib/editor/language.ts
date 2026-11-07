import { StreamLanguage, HighlightStyle, syntaxHighlighting } from '@codemirror/language';
import { tags as t } from '@lezer/highlight';

/**
 * A small tokenizer so comments, strings, numbers and commands are
 * coloured the moment a file opens. Semantic tokens from the server
 * refine identifiers (kinds, helpers, fields, references) afterwards.
 */
export const mmsLanguage = StreamLanguage.define<{ inBlock: boolean }>({
	name: 'mms',
	startState: () => ({ inBlock: false }),
	token(stream, state) {
		if (state.inBlock) {
			if (stream.skipTo('*/')) {
				stream.pos += 2;
				state.inBlock = false;
			} else stream.skipToEnd();
			return 'comment';
		}
		if (stream.sol()) {
			stream.eatSpace();
			if (stream.match(/^\/[a-z][^\n]*/)) return 'meta';
		}
		if (stream.eatSpace()) return null;
		if (stream.match('//')) {
			stream.skipToEnd();
			return 'comment';
		}
		if (stream.match('/*')) {
			state.inBlock = true;
			return 'comment';
		}
		if (stream.match(/^"(?:[^"\\]|\\.)*"?/) || stream.match(/^'(?:[^'\\]|\\.)*'?/)) return 'string';
		if (stream.match(/^\d+(\.\d+)?([eE][+-]?\d+)?/)) return 'number';
		if (stream.match(/^(true|false|const|let|extern|primary|if|else|while|for|in|return|break|continue|emit|check|as|setup|run|expect)\b/)) return 'keyword';
		if (stream.match(/^[a-z0-9_.-]+:[a-z0-9_./-]+/)) return 'atom';
		// Identifiers are left to semantic tokens, which know kinds from
		// helpers from references.
		if (stream.match(/^[A-Za-z_][A-Za-z0-9_]*/)) return null;
		if (stream.match(/^(==|!=|<=|>=|&&|\|\||\.\.|[-+*/%=<>!])/)) return 'operator';
		stream.next();
		return null;
	},
	languageData: { commentTokens: { line: '//', block: { open: '/*', close: '*/' } } }
});

export const mmsHighlight = syntaxHighlighting(
	HighlightStyle.define([
		{ tag: t.comment, class: 'cm-sem-comment' },
		{ tag: t.string, class: 'cm-sem-string' },
		{ tag: t.number, class: 'cm-sem-number' },
		{ tag: t.keyword, class: 'cm-sem-keyword' },
		{ tag: t.meta, class: 'cm-sem-macro' },
		{ tag: t.atom, class: 'cm-sem-enumMember' },
		{ tag: t.operator, class: 'cm-sem-operator' }
	])
);

/** Book markup (.mmsbook): headings, page breaks, style tags, links and ${...}. */
export const bookLanguage = StreamLanguage.define<{ heading: boolean }>({
	name: 'mmsbook',
	startState: () => ({ heading: false }),
	token(stream, state) {
		if (stream.sol()) {
			state.heading = false;
			if (stream.match(/^\s*---\s*$/)) return 'meta';
			if (stream.match('# ')) {
				state.heading = true;
				return 'keyword';
			}
		}
		if (stream.match(/^\\./)) return 'string';
		if (stream.match(/^\$\{[^}]*\}?/)) return 'atom';
		if (stream.match(/^\{[^}]*\}?/)) return 'keyword';
		if (stream.match(/^\[[^\]]*\]\([^)]*\)?/)) return 'string';
		if (stream.match(/^(\*\*|__|~~|\|\||\*)/)) return 'operator';
		stream.next();
		return state.heading ? 'heading' : null;
	}
});

export const bookHighlight = syntaxHighlighting(
	HighlightStyle.define([
		{ tag: t.meta, class: 'cm-sem-macro' },
		{ tag: t.keyword, class: 'cm-sem-keyword' },
		{ tag: t.atom, class: 'cm-sem-variable' },
		{ tag: t.string, class: 'cm-sem-string' },
		{ tag: t.operator, class: 'cm-sem-operator' },
		{ tag: t.heading, fontWeight: 'bold' }
	])
);
