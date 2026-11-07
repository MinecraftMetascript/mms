import { forms, preludeTemplates, preludeTypes } from '#lib/docs/refdata.ts';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => ({ templates: preludeTemplates, types: preludeTypes, forms });
