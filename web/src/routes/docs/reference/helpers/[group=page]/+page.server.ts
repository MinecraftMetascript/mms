import { error } from '@sveltejs/kit';
import { ref } from '#lib/docs/refdata.ts';
import { groupId } from '#lib/docs/reference.ts';
import type { EntryGenerator, PageServerLoad } from './$types';

// One page per helper group.
export const entries: EntryGenerator = () => ref.groups.map((g) => ({ group: groupId(g.dispatcher) }));

export const load: PageServerLoad = ({ params }) => {
	const group = ref.groups.find((g) => groupId(g.dispatcher) === params.group);
	if (!group) error(404, 'No such helper group');
	return { group };
};
