import { blueprints, ownKinds, ref, resourceKinds } from '#lib/docs/refdata.ts';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => ({ ownKinds, resourceKinds, blueprints, total: ref.kinds.length + blueprints.length });
