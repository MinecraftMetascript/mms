import { resolve } from '$app/paths';
import { guide, render } from '#lib/docs/docs.ts';
import type { PageLoad } from './$types';

// The docs home is the getting-started guide.
export const load: PageLoad = () => {
	const g = guide('getting-started')!;
	return { title: g.title, html: render(g.source, resolve('/').replace(/\/$/, '')) };
};
