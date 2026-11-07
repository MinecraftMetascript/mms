import { llmsIndex } from '#lib/docs/ai.ts';

export const prerender = true;
export const GET = () => new Response(llmsIndex(), { headers: { 'content-type': 'text/plain; charset=utf-8' } });
