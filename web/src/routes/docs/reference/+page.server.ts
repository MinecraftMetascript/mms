import { resolve } from '$app/paths';
import { card, groupPages, referenceGroups } from '#lib/docs/site.ts';
import { counts, ref } from '#lib/docs/refdata.ts';
import { groupId, helperId, refURL, typeId } from '#lib/docs/reference.ts';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => {
	const base = resolve('/').replace(/\/$/, '');
	// Where each anchor of the old one-page reference now lives, so links
	// from before the split (such as older editor hovers) still land.
	const moved: Record<string, string> = { kinds: `${base}/docs/reference/kinds`, helpers: `${base}/docs/reference/helpers`, types: `${base}/docs/reference/types` };
	for (const g of ref.groups) {
		moved[groupId(g.dispatcher)] = refURL.group(base, g.dispatcher);
		for (const h of g.helpers) moved[helperId(g.dispatcher, h.name)] = refURL.helper(base, g.dispatcher, h.name);
	}
	for (const t of ref.types) moved[typeId(t.id)] = refURL.type(base, t.id);
	return { counts, language: referenceGroups.flatMap(groupPages).map(card), moved };
};
