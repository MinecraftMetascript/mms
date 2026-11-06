import { ownKinds, ref, resourceKinds } from '#lib/docs/refdata.ts';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => ({ ownKinds, resourceKinds, total: ref.kinds.length });
