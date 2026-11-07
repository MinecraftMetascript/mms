import { docGroups, groupPages, referenceGroups } from './site.ts';
import { blueprints, ref } from './refdata.ts';
import { groupId, groupTitle } from './reference.ts';

/** A sidebar link; `children` show while the link's section is open. */
export interface NavLink {
	path: string; // under the base, such as "/docs/guides/functions"
	title: string;
	children?: NavLink[];
}
export interface NavGroup {
	title: string;
	links: NavLink[];
}

/** The Docs tab's sidebar: Start, Guides and Concepts. */
export const docsNav: NavGroup[] = docGroups
	.map((g) => ({ title: g.title, links: groupPages(g).map((p) => ({ path: `/docs/${p.path}`, title: p.title })) }))
	.filter((g) => g.links.length);

/** The Reference tab's sidebar. */
export const referenceNav: NavGroup[] = [
	{ title: 'Reference', links: [{ path: '/docs/reference', title: 'Overview' }] },
	...referenceGroups
		.map((g) => ({ title: g.title, links: groupPages(g).map((p) => ({ path: `/docs/${p.path}`, title: p.title })) }))
		.filter((g) => g.links.length),
	{
		title: 'Generated',
		links: [
			{
				path: '/docs/reference/kinds',
				title: 'Declaration kinds',
				children: blueprints.map((b) => ({ path: `/docs/reference/kinds/${b.name}`, title: b.name }))
			},
			{
				path: '/docs/reference/helpers',
				title: 'Helpers',
				children: ref.groups.map((g) => ({ path: `/docs/reference/helpers/${groupId(g.dispatcher)}`, title: groupTitle(g.dispatcher) }))
			},
			{ path: '/docs/reference/prelude', title: 'Templates and forms' },
			{ path: '/docs/reference/types', title: 'Types' },
			{ path: '/docs/reference/lint', title: 'Lint rules', children: ref.rules.map((r) => ({ path: `/docs/reference/lint/${r.id}`, title: r.id })) }
		]
	}
];
