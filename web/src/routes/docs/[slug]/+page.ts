import { error } from '@sveltejs/kit';
import { resolve } from '$app/paths';
import { guide, guides, render } from '#lib/docs/docs.ts';
import type { EntryGenerator, PageLoad } from './$types';

export const entries: EntryGenerator = () => guides.map((g) => ({ slug: g.slug }));

export const load: PageLoad = ({ params }) => {
	const g = guide(params.slug);
	if (!g) error(404, 'No such guide');
	return { title: g.title, html: render(g.source, resolve('/').replace(/\/$/, '')) };
};
