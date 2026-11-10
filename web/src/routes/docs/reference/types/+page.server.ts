import { ref, typeAreas } from '#lib/docs/refdata.ts';
import type { PageServerLoad } from './$types';

// The types index: one page per part of the game's schema, and every
// type's name for the filter.
export const load: PageServerLoad = () => ({
	areas: typeAreas,
	names: ref.types.map((t) => ({ id: t.id, name: t.name, area: t.area })).sort((a, b) => a.name.localeCompare(b.name))
});
