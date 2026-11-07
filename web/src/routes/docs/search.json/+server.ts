import { searchIndex } from '#lib/docs/search.ts';

// The site search's index, fetched the first time search opens.
export const prerender = true;
export const GET = () => new Response(JSON.stringify(searchIndex()), { headers: { 'content-type': 'application/json' } });
