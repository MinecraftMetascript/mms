// The docs in forms that AI tools and scripts read easily: every page as
// Markdown at its own URL plus .md (/docs/guides/functions.md), /llms.txt
// (an index in the llms.txt convention, at the site's root where tools
// look), /llms-full.txt (every page in one file) and /docs/reference.yaml
// (the generated reference as data).
import { docGroups, groupPages, page, pages, referenceGroups, rewriteHref, siteURL } from './site.ts';
import { ref, resourceKinds, ownKinds, helperCount, blueprints, forms, preludeTemplates, preludeTypes, counts } from './refdata.ts';
import { groupId, groupTitle, type Group, type Param, type PreludeDecl, type PreludeField } from './reference.ts';

const intro =
	'MMS (Minecraft Metascript) is a language for Minecraft Java Edition data packs and resource packs. ' +
	'Declarations such as `LootTable`, `Recipe` and `Function` compile to the JSON files and `.mcfunction` files the game reads.';

const url = (path: string) => `${siteURL}${path}`;
/** The Markdown URL of a page under /docs, such as "guides/functions". */
export const mdURL = (path: string) => url(`/docs/${path}.md`);

/** /llms.txt: what MMS is, and a link to every page as Markdown. */
export function llmsIndex(): string {
	const out = [`# MMS`, '', `> ${intro}`, ''];
	out.push(`Every page below is Markdown. The same page as HTML is at its URL without \`.md\`.`, '');
	for (const g of [...docGroups, ...referenceGroups]) {
		const ps = groupPages(g);
		if (!ps.length) continue;
		out.push(`## ${g.title}`, '');
		for (const p of ps) out.push(`- [${p.title}](${mdURL(p.path)})${p.summary ? `: ${p.summary}` : ''}`);
		out.push('');
	}
	out.push(
		'## Reference',
		'',
		`- [Declaration kinds](${mdURL('reference/kinds')}): the ${counts.kinds} kinds a declaration can have, and the folder each writes to`,
		...blueprints.map((b) => `- [${b.name}](${mdURL(`reference/kinds/${b.name}`)}): its fields, and what its name stands for`),
		`- [Templates and forms](${mdURL('reference/prelude')}): the prelude's templates and types, and what each operator means by type`,
		`- [Helpers](${mdURL('reference/helpers')}): ${helperCount} helpers in ${ref.groups.length} groups, each building one typed object; each group has a page`,
		`- [Types](${mdURL('reference/types')}): named types that parameters refer to`,
		`- [Lint rules](${mdURL('reference/lint')}): the ${ref.rules.length} lint rules and their default severities; each rule has a page`,
		'',
		'## Optional',
		'',
		`- [Everything in one file](${url('/llms-full.txt')}): every page above, and every helper group and lint rule, as Markdown`,
		`- [Reference as YAML](${url('/docs/reference.yaml')}): kinds, helpers, types and lint rules as data`,
		''
	);
	return out.join('\n');
}

/** /llms-full.txt: every page that has a Markdown version, in reading order. */
export function llmsFull(): string {
	const out = [`# MMS`, '', `> ${intro}`, '', `Minecraft version: ${ref.version}.`, ''];
	for (const path of markdownPaths()) out.push(`<!-- ${url(`/docs/${path}`)} -->`, '', markdownFor(path)!.trim(), '');
	return out.join('\n');
}

/**
 * Every page under /docs that has a Markdown version, in reading order:
 * the hand-written pages, then the generated reference.
 */
export function markdownPaths(): string[] {
	return [
		...[...docGroups, ...referenceGroups].flatMap(groupPages).map((p) => p.path),
		// Pages under site/ that no sidebar group lists still have a URL.
		...pages.filter((p) => ![...docGroups, ...referenceGroups].some((g) => g.dir === p.dir)).map((p) => p.path),
		'reference',
		'reference/kinds',
		...blueprints.map((b) => `reference/kinds/${b.name}`),
		'reference/prelude',
		'reference/helpers',
		...ref.groups.map((g) => `reference/helpers/${groupId(g.dispatcher)}`),
		'reference/types',
		'reference/lint',
		...ref.rules.map((r) => `reference/lint/${r.id}`)
	];
}

let markdownSet: Set<string> | undefined;
const hasMarkdown = (path: string) => (markdownSet ??= new Set(markdownPaths())).has(path);

/** The Markdown of the page at path under /docs, or undefined if there is none. */
export function markdownFor(path: string): string | undefined {
	const p = page(path);
	if (p) return pageMarkdown(p.source, p.path);
	if (path === 'reference') return referenceIndex();
	if (path === 'reference/kinds') return kindsPage();
	if (path === 'reference/prelude') return preludePage();
	if (path === 'reference/helpers') return helpersIndex();
	if (path === 'reference/types') return typesPage();
	if (path === 'reference/lint') return lintIndex();
	const [, kind, name] = /^reference\/(kinds|helpers|lint)\/([^/]+)$/.exec(path) ?? [];
	if (kind === 'kinds') {
		const b = blueprints.find((b) => b.name === name);
		return b && blueprintPage(b);
	}
	if (kind === 'helpers') {
		const g = ref.groups.find((g) => groupId(g.dispatcher) === name);
		return g && helperGroupPage(g);
	}
	if (kind === 'lint') {
		const r = ref.rules.find((r) => r.id === name);
		return r && rulePage(r);
	}
}

/**
 * A hand-written page with its links made absolute: other pages to their
 * Markdown URL, and other repository paths to GitHub. Code is left alone.
 */
export function pageMarkdown(source: string, from: string): string {
	const link = (href: string) => {
		const to = rewriteHref(href, siteURL, from);
		const m = /^(.*\/docs\/[^#]+?)\/?(#.*)?$/.exec(to);
		return m && to.startsWith(siteURL + '/docs/') && hasMarkdown(m[1].slice(siteURL.length + 6)) ? `${m[1]}.md${m[2] ?? ''}` : to;
	};
	let fence = false;
	return source
		.split('\n')
		.map((line) => {
			if (/^\s*(```|~~~)/.test(line)) fence = !fence;
			if (fence || line.trimStart().startsWith('```')) return line;
			// Code spans first, so a link inside one is left alone.
			return line.replace(/(`+)[^`]*?\1|(\]\()([^)\s]+)/g, (all, _tick, open, href) => (open ? open + link(href) : all));
		})
		.join('\n');
}

const flat = (doc?: string) => (doc ?? '').replace(/[ \t]*\\\n/g, '\n').replace(/\s*\n\s*/g, ' ').trim();
const firstPara = (doc?: string) => flat((doc ?? '').split(/\n\n/)[0]);
const paras = (doc?: string) => (doc ? doc.split(/\n\n+/).map(flat).filter(Boolean) : []);

const paramLine = (p: Param) =>
	`- \`${p.name}\`: ${p.type}${p.default ? ` (default \`${p.default}\`)` : p.optional ? ' (optional)' : ''}${p.doc ? `. ${flat(p.doc)}` : ''}`;
const fieldLine = (f: PreludeField) => paramLine({ ...f, type: f.type ?? 'any' });
const usesLines = (d: PreludeDecl) => (d.uses ?? []).map((u) => `- where \`${u.type}\` goes${u.doc ? `: ${flat(u.doc)}` : ''}`);

function referenceIndex(): string {
	const language = referenceGroups.flatMap(groupPages);
	return [
		'# Reference',
		'',
		`Everything MMS knows about Minecraft ${counts.version}, for looking things up. The generated pages come from the same schema and rules the checker uses. To learn MMS, start with [Getting started](${mdURL('start/getting-started')}) instead.`,
		'',
		'## Language',
		'',
		...language.map((p) => `- [${p.title}](${mdURL(p.path)})${p.summary ? `: ${p.summary}` : ''}`),
		'',
		'## Generated',
		'',
		`- [Declaration kinds](${mdURL('reference/kinds')}): ${counts.kinds} kinds, and the folder each one writes to.`,
		`- [Helpers](${mdURL('reference/helpers')}): ${counts.helpers} helpers in ${counts.groups} groups, each building one typed object.`,
		`- [Templates and forms](${mdURL('reference/prelude')}): the prelude's templates and types, and what each operator means by type.`,
		`- [Types](${mdURL('reference/types')}): ${counts.types} named types that parameters refer to.`,
		`- [Lint rules](${mdURL('reference/lint')}): ${counts.rules} rules, their default severity and what they catch.`,
		'',
		'## For AI tools',
		'',
		`- [llms.txt](${url('/llms.txt')}): what MMS is, and a link to every page.`,
		`- [llms-full.txt](${url('/llms-full.txt')}): every page and the whole reference as one Markdown file.`,
		`- [reference.yaml](${url('/docs/reference.yaml')}): kinds, helpers, types and lint rules as data.`,
		''
	].join('\n');
}

function kindsPage(): string {
	return [
		'# Declaration kinds',
		'',
		`Every declaration starts with its kind: \`LootTable gems { ... }\` or \`ItemTag gems = [...]\`. There are ${counts.kinds} of them.`,
		'',
		'## MMS kinds',
		'',
		"These are MMS's own, for functions, text and other things the game has no single file for.",
		'',
		...ownKinds.map((k) => `- \`${k.name}\`${k.doc ? `: ${flat(k.doc)}` : ''}`),
		'',
		'## Blueprint kinds',
		'',
		'These are written in MMS, in the prelude. Each one writes several resources from one declaration, and has a page of its own.',
		'',
		...blueprints.map((b) => `- [\`${b.name}\`](${mdURL(`reference/kinds/${b.name}`)})${b.doc ? `: ${firstPara(b.doc)}` : ''}`),
		'',
		'## Resource kinds',
		'',
		'Each of these makes one resource in the folder shown: in the data pack, or for `assets/` folders in the resource pack. Tags of other registries follow the same pattern, such as `DamageTypeTag`.',
		'',
		...resourceKinds.map((k) => `- \`${k.name}\`: writes \`${k.resources ? 'assets' : 'data'}/<namespace>/${k.folder}/\`${k.doc ? `. ${flat(k.doc)}` : ''}`),
		''
	].join('\n');
}

function blueprintPage(b: PreludeDecl): string {
	const required = (b.fields ?? []).filter((f) => !f.optional);
	const shape = `${b.name} name {\n${required.map((f) => `  ${f.name} = ...`).join('\n')}${required.length ? '\n' : ''}}`;
	const out = [`# ${b.name}`, '', ...paras(b.doc).flatMap((p) => [p, '']), '```', shape, '```', ''];
	out.push(`\`${b.name}\` is a [Blueprint](${mdURL('reference/language/blueprints')}) in the prelude: it is written in MMS, and every project has it.`, '');
	out.push('## Fields', '', ...((b.fields ?? []).length ? b.fields!.map(fieldLine) : ['It has no fields.']), '');
	if (b.uses?.length) out.push('## What its name stands for', '', "Where one of these types is expected, the declaration's name means:", '', ...usesLines(b), '');
	return out.join('\n');
}

function preludeEntry(d: PreludeDecl): string[] {
	const out = [`### ${d.name}`, ''];
	if (d.kind === 'Template' && d.builtin) out.push('Built in.', '');
	for (const p of paras(d.doc)) out.push(p, '');
	if (!(d.kind === 'Type' && d.builtin)) out.push('```', d.kind === 'Type' ? `Type ${d.signature}` : d.signature, '```', '');
	if (d.uses?.length) out.push(...usesLines(d), '');
	else if (d.fields?.length && !(d.kind === 'Type' && d.builtin)) out.push(...d.fields.map(fieldLine), '');
	return out;
}

function preludePage(): string {
	return [
		'# Templates and forms',
		'',
		`Every project has the prelude: [Templates](${mdURL('reference/language/reuse')}#templates), Types and [Blueprint kinds](${mdURL('reference/kinds')}#blueprint-kinds) written in MMS, and the built-in templates the compiler supplies. Call a template like any helper.`,
		'',
		'## Templates',
		'',
		...preludeTemplates.flatMap(preludeEntry),
		'## Types',
		'',
		...preludeTypes.flatMap(preludeEntry),
		'## Forms of expression',
		'',
		`Operators and literals are forms of expression. What one means depends on the type expected where it is written, so \`1..3\` is a uniform number in a loot pool's rolls and bounds in a predicate. This is the [shorthand](${mdURL('reference/language/values')}#shorthand) the language builds in.`,
		'',
		...forms.flatMap(preludeEntry)
	].join('\n');
}

function helpersIndex(): string {
	return [
		'# Helpers',
		'',
		`A helper builds one typed object. \`Name(a, b)\` fills the parameters in the order listed, and \`Name(a = 1)\` names them. There are ${counts.helpers} helpers, grouped by what they make.`,
		'',
		...ref.groups.map(
			(g) =>
				`- [${groupTitle(g.dispatcher)}](${mdURL(`reference/helpers/${groupId(g.dispatcher)}`)}): ${g.helpers.length} ${g.helpers.length === 1 ? 'helper' : 'helpers'}${g.dispatcher ? `, \`${g.dispatcher}\`` : ''}`
		),
		''
	].join('\n');
}

function helperGroupPage(g: Group): string {
	const out = [`# ${groupTitle(g.dispatcher)}`, ''];
	out.push(
		g.dispatcher
			? `Helpers that build a \`${g.dispatcher}\` object, one for each of its types. Where a field takes one, write the helper's call; its name is the type it makes.`
			: 'Helpers for objects that are not chosen by a `type` field. Each builds the object it is named after.',
		''
	);
	for (const h of g.helpers) {
		out.push(`## ${h.name}`, '', '```', h.signature, '```', '');
		if (h.case) out.push(`Makes \`${h.case}\`.`, '');
		if (h.doc) out.push(h.doc, '');
		if (h.params?.length) out.push(...h.params.map(paramLine), '');
		if (h.example) out.push('```mms', h.example, '```', '');
	}
	return out.join('\n');
}

function typesPage(): string {
	const out = ['# Types', '', 'Named types that helper parameters refer to: enums, aliases and records.', ''];
	for (const t of ref.types) {
		if (t.values?.length) {
			if (t.values.some((v) => v.doc)) out.push(`- \`${t.name}\`: one of`, ...t.values.map((v) => `  - \`"${v.value}"\`${v.doc ? `: ${flat(v.doc)}` : ''}`));
			else out.push(`- \`${t.name}\`: one of ${t.values.map((v) => `\`"${v.value}"\``).join(', ')}`);
		} else if (t.def?.length) out.push(`- \`${t.name}\`: ${t.def.map((p) => p.text).join('')}`);
		else if (t.fields?.length) out.push(`- \`${t.name}\`: a record with fields`, ...t.fields.map((f) => '  ' + paramLine(f)));
	}
	out.push('');
	return out.join('\n');
}

function lintIndex(): string {
	return [
		'# Lint rules',
		'',
		`Lint rules point out style problems and likely mistakes; errors the checker reports are not lint. A project can change any rule's severity in \`mms.json\`, and [Lint rules](${mdURL('guides/lint')}) in the guides explains how, and how to silence one.`,
		'',
		...ref.rules.map((r) => `- [\`${r.id}\`](${mdURL(`reference/lint/${r.id}`)}) (${r.severity} by default)${r.doc ? `: ${flat(r.doc)}` : ''}`),
		''
	].join('\n');
}

function rulePage(r: (typeof ref.rules)[number]): string {
	return [
		`# ${r.id}`,
		'',
		`${r.severity[0].toUpperCase()}${r.severity.slice(1)} by default.`,
		'',
		...paras(r.doc).flatMap((p) => [p, '']),
		'## Changing it',
		'',
		"Set the rule's severity in `mms.json`, to `error`, `warning`, `info`, `hint` or `off`:",
		'',
		'```json',
		`{\n  "lint": { "${r.id}": "off" }\n}`,
		'```',
		'',
		`To silence it on one line, write \`// mms-ignore ${r.id}\` on the line or the line above. See [Lint rules](${mdURL('guides/lint')}) in the guides for more.`,
		''
	].join('\n');
}

/** /docs/reference.yaml: the generated reference as data. */
export function referenceYaml(): string {
	const params = (ps?: Param[]) =>
		ps?.map((p) => ({ name: p.name, type: p.type, optional: p.optional || undefined, default: p.default, doc: p.doc }));
	const data = {
		version: ref.version,
		docs: url('/docs'),
		kinds: ref.kinds.map((k) => ({
			name: k.name,
			folder: k.folder ? `${k.resources ? 'assets' : 'data'}/<namespace>/${k.folder}/` : undefined,
			doc: k.doc
		})),
		blueprints: blueprints.map((b) => ({
			name: b.name,
			url: url(`/docs/reference/kinds/${b.name}`),
			doc: b.doc,
			fields: b.fields?.map((f) => ({ name: f.name, type: f.type, optional: f.optional || undefined, default: f.default, doc: f.doc })),
			stands_for: b.uses?.map((u) => ({ type: u.type, doc: u.doc }))
		})),
		templates: preludeTemplates.map((d) => ({
			name: d.name,
			signature: d.signature,
			builtin: d.builtin || undefined,
			doc: d.doc,
			params: d.fields?.map((f) => ({ name: f.name, type: f.type, optional: f.optional || undefined, default: f.default, doc: f.doc }))
		})),
		prelude_types: preludeTypes.map((d) => ({ name: d.name, is: d.signature, doc: d.doc })),
		forms: forms.map((d) => ({ name: d.name, doc: d.doc, means: d.uses?.map((u) => ({ where: u.type, doc: u.doc })) })),
		helpers: ref.groups.map((g) => ({
			group: groupTitle(g.dispatcher),
			dispatcher: g.dispatcher || undefined,
			url: url(`/docs/reference/helpers/${groupId(g.dispatcher)}`),
			helpers: g.helpers.map((h) => ({ name: h.name, makes: h.case, signature: h.signature, doc: h.doc, params: params(h.params), example: h.example }))
		})),
		types: ref.types.map((t) => ({
			name: t.name,
			id: t.id,
			is: t.def?.map((p) => p.text).join(''),
			values: t.values?.map((v) => (v.doc ? { value: v.value, doc: v.doc } : v.value)),
			fields: params(t.fields)
		})),
		lint_rules: ref.rules.map((r) => ({ id: r.id, severity: r.severity, doc: r.doc }))
	};
	return `# The MMS reference for Minecraft ${ref.version}, generated by \`mms reference\`.\n${yaml(data, 0)}`;
}

/**
 * Plain data as YAML: maps, lists, strings, numbers and booleans. Strings
 * are always double-quoted (JSON's escapes are valid YAML), and undefined
 * values and empty lists are left out.
 */
export function yaml(v: unknown, indent: number): string {
	const pad = ' '.repeat(indent);
	if (Array.isArray(v)) {
		return v
			.map((item) => {
				if (isMap(item)) {
					const body = yaml(item, indent + 2);
					return `\n${pad}- ${body.trimStart()}`;
				}
				return `\n${pad}- ${scalar(item)}`;
			})
			.join('');
	}
	if (isMap(v)) {
		return Object.entries(v)
			.filter(([, x]) => x !== undefined && !(Array.isArray(x) && !x.length))
			.map(([k, x]) => (isMap(x) || Array.isArray(x) ? `\n${pad}${k}:${yaml(x, indent + 2)}` : `\n${pad}${k}: ${scalar(x)}`))
			.join('');
	}
	return scalar(v);
}

const isMap = (v: unknown): v is Record<string, unknown> => typeof v === 'object' && v !== null && !Array.isArray(v);
const scalar = (v: unknown) => (typeof v === 'string' ? JSON.stringify(v) : String(v));
