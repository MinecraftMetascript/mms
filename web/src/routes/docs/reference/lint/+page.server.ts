import { ref } from '#lib/docs/refdata.ts';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = () => ({ rules: ref.rules });
