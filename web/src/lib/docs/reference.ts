// The shape of `mms reference`, generated into reference.json by
// scripts/build-wasm.sh.
export interface Reference {
	version: string;
	kinds: { name: string; folder?: string; doc?: string }[];
	groups: Group[];
}

export interface Group {
	dispatcher: string;
	helpers: Helper[];
}

export interface Helper {
	name: string;
	case?: string;
	signature: string;
	params?: { name: string; type: string; optional?: boolean; doc?: string }[];
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

/** Groups with only the helpers whose name, type or parameters match q. */
export function filterGroups(groups: Group[], q: string): Group[] {
	const s = q.trim().toLowerCase();
	if (!s) return groups;
	return groups
		.map((g) => ({
			...g,
			helpers: g.helpers.filter(
				(h) => h.name.toLowerCase().includes(s) || (h.case ?? '').includes(s) || groupTitle(g.dispatcher).toLowerCase().includes(s)
			)
		}))
		.filter((g) => g.helpers.length);
}

/** Splits schema documentation into text and `code` parts. */
export function inlineCode(doc: string): { code: boolean; text: string }[] {
	return doc.split('`').map((text, i) => ({ code: i % 2 === 1, text }));
}
