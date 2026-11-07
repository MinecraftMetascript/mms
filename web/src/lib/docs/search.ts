// The site search's index, built at build time and served as
// /docs/search.json. The browser side is find.ts.

import { docGroups, groupPages, referenceGroups, sections, type Group } from './site.ts';
import { ref } from './refdata.ts';
import { groupId, groupTitle, helperId, typeId } from './reference.ts';
import type { Entry } from './find.ts';

const pagesOf = (groups: Group[]) =>
	groups.flatMap((g) =>
		groupPages(g).flatMap((p) => [
			{ title: p.title, where: g.title, path: `/docs/${p.path}`, text: p.summary },
			...sections(p).map((s) => ({ title: s.title, where: `${g.title} · ${p.title}`, path: `/docs/${p.path}#${s.id}`, text: s.text.slice(0, 300) }))
		])
	);

/** Everything the site search finds, built at build time. */
export function searchIndex(): Entry[] {
	return [
		...pagesOf(docGroups),
		...pagesOf(referenceGroups),
		...ref.kinds.map((k) => ({
			title: k.name,
			where: 'Declaration kinds',
			path: `/docs/reference/kinds#kind-${k.name}`,
			text: k.doc ?? (k.folder ? `${k.resources ? 'assets' : 'data'}/<namespace>/${k.folder}/` : '')
		})),
		...ref.groups.flatMap((g) => [
			{ title: groupTitle(g.dispatcher), where: 'Helpers', path: `/docs/reference/helpers/${groupId(g.dispatcher)}`, text: `${g.helpers.length} helpers` },
			...g.helpers.map((h) => ({
				title: h.name,
				where: `Helpers · ${groupTitle(g.dispatcher)}`,
				path: `/docs/reference/helpers/${groupId(g.dispatcher)}#${helperId(g.dispatcher, h.name)}`,
				text: (h.doc ?? h.signature).slice(0, 300)
			}))
		]),
		...ref.types.map((t) => ({ title: t.name, where: 'Types', path: `/docs/reference/types#${typeId(t.id)}`, text: '' })),
		...ref.rules.map((r) => ({ title: r.id, where: 'Lint rules', path: `/docs/reference/lint/${r.id}`, text: r.doc ?? '' }))
	];
}
