import { error } from '@sveltejs/kit';
import { resolve } from '$app/paths';
import { moved, page, pages, readingOrder, render } from '#lib/docs/site.ts';
import type { EntryGenerator, PageServerLoad } from './$types';

// Every hand-written page under site/, plus the guides' old URLs, which
// now redirect.
export const entries: EntryGenerator = () => [...pages.map((p) => ({ path: p.path })), ...Object.keys(moved).map((path) => ({ path }))];

export const load: PageServerLoad = ({ params }) => {
	const base = resolve('/').replace(/\/$/, '');
	if (moved[params.path]) return { redirect: `${base}/docs/${moved[params.path]}` };
	const p = page(params.path);
	if (!p) error(404, 'No such page');
	const order = readingOrder();
	const i = order.indexOf(p);
	const link = (q: (typeof order)[number] | undefined) => q && { href: `${base}/docs/${q.path}`, title: q.title };
	return {
		title: p.title,
		html: render(p, base),
		prev: i > 0 ? link(order[i - 1]) : undefined,
		next: i >= 0 ? link(order[i + 1]) : undefined
	};
};
