import { llmsFull } from '#lib/docs/ai.ts';

// At the site's root, next to /llms.txt.
export const prerender = true;
export const GET = () => new Response(llmsFull(), { headers: { 'content-type': 'text/plain; charset=utf-8' } });
