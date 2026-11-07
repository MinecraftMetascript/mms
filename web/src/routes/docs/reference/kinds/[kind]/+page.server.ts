import { error } from '@sveltejs/kit';
import { blueprints } from '#lib/docs/refdata.ts';
import type { EntryGenerator, PageServerLoad } from './$types';

// One page per kind written in MMS in the prelude, such as CustomItem.
export const entries: EntryGenerator = () => blueprints.map((b) => ({ kind: b.name }));

export const load: PageServerLoad = ({ params }) => {
	const kind = blueprints.find((b) => b.name === params.kind);
	if (!kind) error(404, 'No such declaration kind');
	return { kind };
};
