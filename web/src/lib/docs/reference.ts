// The shape of `mms reference`, generated into reference.json by
// scripts/build-wasm.sh.
export interface Reference {
	version: string;
	kinds: { name: string; folder?: string; doc?: string }[];
	groups: Group[];
	types: TypeEntry[];
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
}

export interface Param {
	name: string;
	/** The readable type as text. */
	type: string;
	/** The same type in pieces, with named types linked. */
	typeParts?: TypePart[];
	optional?: boolean;
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

/** The anchor a type part links to, or "" for plain text. */
export function linkId(link: TypePart['link']): string {
	if (!link) return '';
	if (link.type) return typeId(link.type);
	if (link.group === undefined) return '';
	return link.helper ? helperId(link.group, link.helper) : groupId(link.group);
}

/** The call shape without types, such as `Noise(noise, xz_scale?, y_scale?)`. */
export const callShape = (h: Helper) => `${h.name}(${(h.params ?? []).map((p) => p.name + (p.optional ? '?' : '')).join(', ')})`;

const matches = (s: string, q: string) => s.toLowerCase().includes(q);

/**
 * Groups with only the helpers that match q: by name, the type they make,
 * their group, or a parameter's name or type.
 */
export function filterGroups(groups: Group[], q: string): Group[] {
	const s = q.trim().toLowerCase();
	if (!s) return groups;
	return groups
		.map((g) => {
			if (matches(groupTitle(g.dispatcher), s) || matches(g.dispatcher, s)) return g;
			return {
				...g,
				helpers: g.helpers.filter(
					(h) => matches(h.name, s) || matches(h.case ?? '', s) || (h.params ?? []).some((p) => matches(p.name, s) || matches(p.type, s))
				)
			};
		})
		.filter((g) => g.helpers.length);
}

/** Types whose name matches q. */
export function filterTypes(types: TypeEntry[], q: string): TypeEntry[] {
	const s = q.trim().toLowerCase();
	return s ? types.filter((t) => matches(t.name, s)) : types;
}

/** Splits schema documentation into text and `code` parts. */
export function inlineCode(doc: string): { code: boolean; text: string }[] {
	return doc.split('`').map((text, i) => ({ code: i % 2 === 1, text }));
}
