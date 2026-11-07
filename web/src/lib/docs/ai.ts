// The reference in forms that AI tools and scripts read easily, served
// next to the docs: /docs/llms.txt (an index in the llms.txt
// convention), /docs/llms-full.txt (every page and the whole reference
// as Markdown) and /docs/reference.yaml (the generated reference as data).
import { docGroups, groupPages, referenceGroups, siteURL } from './site.ts';
import { ref, resourceKinds, ownKinds, helperCount, blueprints, forms, preludeTemplates, preludeTypes } from './refdata.ts';
import { groupId, groupTitle, type Param, type PreludeDecl, type PreludeField } from './reference.ts';

const intro =
	'MMS (Minecraft Metascript) is a language for Minecraft Java Edition data packs and resource packs. ' +
	'Declarations such as `LootTable`, `Recipe` and `Function` compile to the JSON files and `.mcfunction` files the game reads.';

const url = (path: string) => `${siteURL}${path}`;

/** /docs/llms.txt: what MMS is, and a link to every page. */
export function llmsIndex(): string {
	const out = [`# MMS`, '', `> ${intro}`, ''];
	for (const g of [...docGroups, ...referenceGroups]) {
		const ps = groupPages(g);
		if (!ps.length) continue;
		out.push(`## ${g.title}`, '');
		for (const p of ps) out.push(`- [${p.title}](${url(`/docs/${p.path}`)})${p.summary ? `: ${p.summary}` : ''}`);
		out.push('');
	}
	out.push(
		'## Reference',
		'',
		`- [Declaration kinds](${url('/docs/reference/kinds')}): the ${ref.kinds.length + blueprints.length} kinds a declaration can have, and the folder each writes to`,
		...blueprints.map((b) => `- [${b.name}](${url(`/docs/reference/kinds/${b.name}`)}): its fields, and what its name stands for`),
		`- [Templates and forms](${url('/docs/reference/prelude')}): the prelude's templates and types, and what each operator means by type`,
		`- [Helpers](${url('/docs/reference/helpers')}): ${helperCount} helpers in ${ref.groups.length} groups, each building one typed object`,
		`- [Types](${url('/docs/reference/types')}): named types that parameters refer to`,
		`- [Lint rules](${url('/docs/reference/lint')}): the ${ref.rules.length} lint rules and their default severities`,
		'',
		'## Optional',
		'',
		`- [Everything in one file](${url('/docs/llms-full.txt')}): every page and the whole reference as Markdown`,
		`- [Reference as YAML](${url('/docs/reference.yaml')}): kinds, helpers, types and lint rules as data`,
		''
	);
	return out.join('\n');
}

const fieldLine = (f: PreludeField) => paramLine({ ...f, type: f.type ?? 'any' });
const usesLines = (d: PreludeDecl) => (d.uses ?? []).map((u) => `  - where \`${u.type}\` goes${u.doc ? `: ${u.doc.replace(/\s*\n\s*/g, ' ')}` : ''}`);

const paramLine = (p: Param) =>
	`  - \`${p.name}\`: ${p.type}${p.default ? ` (default \`${p.default}\`)` : p.optional ? ' (optional)' : ''}${p.doc ? `. ${p.doc.replace(/\s*\n\s*/g, ' ')}` : ''}`;

/** /docs/llms-full.txt: every hand-written page, then the reference. */
export function llmsFull(): string {
	const out = [`# MMS`, '', `> ${intro}`, '', `Minecraft version: ${ref.version}.`, ''];
	for (const g of [...docGroups, ...referenceGroups]) {
		for (const p of groupPages(g)) {
			out.push(`<!-- ${url(`/docs/${p.path}`)} -->`, '', p.source.trim(), '');
		}
	}
	out.push('# Reference: declaration kinds', '');
	for (const k of ownKinds) out.push(`- \`${k.name}\`: ${k.doc ?? ''}`);
	for (const k of resourceKinds) out.push(`- \`${k.name}\`: writes \`${k.resources ? 'assets' : 'data'}/<namespace>/${k.folder}/\`${k.doc ? `. ${k.doc}` : ''}`);
	for (const b of blueprints) {
		out.push('', `## ${b.name}`, '', `A Blueprint in the prelude. ${b.doc ?? ''}`, '', 'Fields:', ...(b.fields ?? []).map(fieldLine));
		if (b.uses?.length) out.push('', 'Its name stands for:', ...usesLines(b));
	}
	out.push('', '# Reference: templates and forms', '');
	for (const d of [...preludeTemplates, ...preludeTypes]) {
		out.push(`### ${d.name}`, '', '```', d.kind === 'Type' ? `Type ${d.signature}` : d.signature, '```', '');
		if (d.doc) out.push(d.doc, '');
		if (d.kind === 'Template' && d.fields?.length) out.push(...d.fields.map(fieldLine), '');
	}
	out.push('## Forms of expression', '', 'What an operator or literal means depends on the type expected where it is written.', '');
	for (const d of forms) out.push(`- ${d.name}: ${d.doc ?? ''}`, ...usesLines(d));
	out.push('', '# Reference: helpers', '', 'A helper builds one typed object. `Name(a, b)` fills the parameters in order; `Name(a = 1)` names them.', '');
	for (const g of ref.groups) {
		out.push(`## ${groupTitle(g.dispatcher)}${g.dispatcher ? ` (\`${g.dispatcher}\`)` : ''}`, '');
		for (const h of g.helpers) {
			out.push(`### ${h.name}`, '', '```', h.signature, '```', '');
			if (h.case) out.push(`Makes \`${h.case}\`.`, '');
			if (h.doc) out.push(h.doc, '');
			if (h.params?.length) out.push(...h.params.map(paramLine), '');
			if (h.example) out.push('```mms', h.example, '```', '');
		}
	}
	out.push('# Reference: types', '');
	for (const t of ref.types) {
		if (t.values?.length) out.push(`- \`${t.name}\`: one of ${t.values.map((v) => `"${v.value}"`).join(', ')}`);
		else if (t.def?.length) out.push(`- \`${t.name}\`: ${t.def.map((p) => p.text).join('')}`);
		else if (t.fields?.length) out.push(`- \`${t.name}\`: a record with fields`, ...t.fields.map(paramLine));
	}
	out.push('', '# Reference: lint rules', '');
	for (const r of ref.rules) out.push(`- \`${r.id}\` (${r.severity} by default): ${r.doc ?? ''}`);
	out.push('');
	return out.join('\n');
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
