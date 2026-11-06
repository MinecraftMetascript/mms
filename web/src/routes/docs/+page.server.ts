import { card, docGroups, groupPages } from '#lib/docs/site.ts';
import { counts } from '#lib/docs/refdata.ts';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => ({
	counts,
	groups: docGroups.map((g) => ({ dir: g.dir, title: g.title, pages: groupPages(g).map(card) })).filter((g) => g.pages.length)
});
