import { counts, ref } from '#lib/docs/refdata.ts';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => ({ counts, groups: ref.groups.map((g) => ({ dispatcher: g.dispatcher, count: g.helpers.length })) });
