import { error } from '@sveltejs/kit';
import { ref, typeAreas } from '#lib/docs/refdata.ts';
import { typeAreaTitle } from '#lib/docs/reference.ts';
import type { EntryGenerator, PageServerLoad } from './$types';

// One page per part of the game's schema, such as worldgen.
export const entries: EntryGenerator = () => typeAreas.map((a) => ({ area: a.area }));

export const load: PageServerLoad = ({ params }) => {
	const types = ref.types.filter((t) => t.area === params.area);
	if (!types.length) error(404, 'No such types page');
	return { area: params.area, title: typeAreaTitle(params.area), types };
};
