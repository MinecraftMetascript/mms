import { llmsIndex } from '#lib/docs/ai.ts';

// At the site's root, where the llms.txt convention says tools look.
// /docs/llms.txt serves the same file.
export const prerender = true;
export const GET = () => new Response(llmsIndex(), { headers: { 'content-type': 'text/plain; charset=utf-8' } });
