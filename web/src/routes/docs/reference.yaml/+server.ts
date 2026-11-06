import { referenceYaml } from '#lib/docs/ai.ts';

export const prerender = true;
export const GET = () => new Response(referenceYaml(), { headers: { 'content-type': 'application/yaml; charset=utf-8' } });
