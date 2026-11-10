import { error } from '@sveltejs/kit';
import { ref } from '#lib/docs/refdata.ts';
import type { EntryGenerator, PageServerLoad } from './$types';

// One page per lint rule, so a diagnostic can link to its rule.
export const entries: EntryGenerator = () => ref.rules.map((r) => ({ rule: r.id }));

export const load: PageServerLoad = ({ params }) => {
	const rule = ref.rules.find((r) => r.id === params.rule);
	if (!rule) error(404, 'No such lint rule');
	return { rule };
};
