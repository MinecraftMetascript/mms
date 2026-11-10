import { error } from '@sveltejs/kit';
import { markdownFor, markdownPaths } from '#lib/docs/ai.ts';
import type { EntryGenerator, RequestHandler } from './$types';

// Every docs page as Markdown at its URL plus .md, such as
// /docs/reference/kinds/JigsawStructure.md, for AI tools and scripts.
export const prerender = true;
export const entries: EntryGenerator = () => markdownPaths().map((p) => ({ path: `${p}.md` }));

export const GET: RequestHandler = ({ params }) => {
	const md = markdownFor(params.path.slice(0, -'.md'.length));
	if (md === undefined) error(404, 'No such page');
	return new Response(md, { headers: { 'content-type': 'text/markdown; charset=utf-8' } });
};
