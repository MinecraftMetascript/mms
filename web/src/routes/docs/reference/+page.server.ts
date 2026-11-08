import { card, groupPages, referenceGroups } from '#lib/docs/site.ts';
import { counts } from '#lib/docs/refdata.ts';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => ({ counts, language: referenceGroups.flatMap(groupPages).map(card) });
