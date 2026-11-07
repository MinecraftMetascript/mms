// The shape of `mms reference`, generated into reference.json by
// scripts/build-wasm.sh.
export interface Reference {
	version: string;
	kinds: { name: string; folder?: string; resources?: boolean; doc?: string }[];
	groups: Group[];
	types: TypeEntry[];
	rules: Rule[];
	prelude: PreludeDecl[];
}

/** A Template, Blueprint or Type that every project has, from the prelude. */
export interface PreludeDecl {
	kind: 'Blueprint' | 'Template' | 'Type';
	name: string;
	/** How it is written, such as `tell(target: Selector, message: Text): Commands`. */
	signature: string;
	doc?: string;
	/** A Blueprint's or Type's fields, or a Template's parameters. */
	fields?: PreludeField[];
	/** The first type of `uses`. */
	as?: string;
	/** Each `as T = value`: what the name, or a built-in form, means where a T goes. */
	uses?: { type: string; doc?: string }[];
	/** The compiler supplies its meaning: a built-in Template, or a form of expression such as `a..b`. */
	builtin?: boolean;
}

export interface PreludeField {
	name: string;
	type?: string;
	default?: string;
	optional?: boolean;
	doc?: string;
}

/** A lint rule and its default severity. */
export interface Rule {
	id: string;
	severity: 'error' | 'warning' | 'info' | 'hint';
	doc?: string;
}

export interface Group {
	dispatcher: string;
	helpers: Helper[];
}

export interface Helper {
	name: string;
	case?: string;
	signature: string;
	params?: Param[];
	/** Hand-written: what the helper does, as Markdown with `code` spans. */
	doc?: string;
	/** Hand-written: MMS source using the helper. */
	example?: string;
}

export interface Param {
	name: string;
	/** The readable type as text. */
	type: string;
	/** The same type in pieces, with named types linked. */
	typeParts?: TypePart[];
	optional?: boolean;
	/** The value MMS writes when the parameter is left out. */
	default?: string;
	doc?: string;
}

export interface TypePart {
	text: string;
	link?: { group?: string; helper?: string; type?: string };
}

/** A named type that no helper group explains, such as an enum. */
export interface TypeEntry {
	id: string;
	name: string;
	def?: TypePart[];
	values?: { value: string; doc?: string }[];
	fields?: Param[];
}

/** A group's heading: what its helpers build. */
export function groupTitle(dispatcher: string): string {
	if (!dispatcher) return 'Objects';
	return dispatcher
		.replace(/^minecraft:/, '')
		.replace(/^worldgen\//, '')
		.split('_')
		.map((w, i) => (i === 0 ? w[0].toUpperCase() + w.slice(1) : w))
		.join(' ');
}

export const groupId = (dispatcher: string) => (dispatcher ? dispatcher.replace(/^minecraft:/, '').replace(/\W/g, '-') : 'objects');
export const helperId = (dispatcher: string, name: string) => `${groupId(dispatcher)}-${name}`;
export const typeId = (id: string) => `type-${id.replace(/\W/g, '-')}`;

/** URLs of reference pages, under base (the site's base path). */
export const refURL = {
	group: (base: string, dispatcher: string) => `${base}/docs/reference/helpers/${groupId(dispatcher)}`,
	helper: (base: string, dispatcher: string, name: string) => `${refURL.group(base, dispatcher)}#${helperId(dispatcher, name)}`,
	type: (base: string, id: string) => `${base}/docs/reference/types#${typeId(id)}`,
	rule: (base: string, id: string) => `${base}/docs/reference/lint/${id}`,
	kind: (base: string, name: string) => `${base}/docs/reference/kinds/${name}`,
	prelude: (base: string, name: string) => `${base}/docs/reference/prelude#prelude-${name}`
};

/** The URL a type part links to, or "" for plain text. */
export function linkHref(base: string, link: TypePart['link']): string {
	if (!link) return '';
	if (link.type) return refURL.type(base, link.type);
	if (link.group === undefined) return '';
	return link.helper ? refURL.helper(base, link.group, link.helper) : refURL.group(base, link.group);
}

/** The call shape without types, such as `Noise(noise, xz_scale?, y_scale?)`. */
export const callShape = (h: Helper) => `${h.name}(${(h.params ?? []).map((p) => p.name + (p.optional ? '?' : '')).join(', ')})`;

const matches = (s: string, q: string) => s.toLowerCase().includes(q);

/** Types whose name matches q. */
export function filterTypes(types: TypeEntry[], q: string): TypeEntry[] {
	const s = q.trim().toLowerCase();
	return s ? types.filter((t) => matches(t.name, s)) : types;
}

/** Splits schema documentation into text and `code` parts. */
export function inlineCode(doc: string): { code: boolean; text: string }[] {
	return doc.split('`').map((text, i) => ({ code: i % 2 === 1, text }));
}
