import { markdownPaths } from '#lib/docs/ai.ts';
import { docsNav, referenceNav } from '#lib/docs/nav.ts';
import type { LayoutServerLoad } from './$types';

const markdown = new Set(markdownPaths());

// The sidebars are built at build time, so the generated reference stays
// out of the pages' JavaScript. A page with a Markdown version gets its
// path under /docs, to link to it.
export const load: LayoutServerLoad = ({ url }) => {
	const path = /\/docs\/(.+?)\/?$/.exec(url.pathname)?.[1];
	return { docsNav, referenceNav, markdown: path && markdown.has(path) ? path : undefined };
};
