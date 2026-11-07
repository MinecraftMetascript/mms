import raw from './reference.json';
import type { Reference } from './reference.ts';

// The output of `mms reference`, generated into reference.json by
// scripts/build-wasm.sh.
export const ref = raw as Reference;

/** Resource kinds (a data or resource pack folder each), and MMS's own. */
export const resourceKinds = ref.kinds.filter((k) => k.folder);
export const ownKinds = ref.kinds.filter((k) => !k.folder);
/** Kinds written in MMS in the prelude, such as CustomItem: each has a page. */
export const blueprints = ref.prelude.filter((d) => d.kind === 'Blueprint');
/** The prelude's Templates and named Types, which projects call and use. */
export const preludeTemplates = ref.prelude.filter((d) => d.kind === 'Template');
export const preludeTypes = ref.prelude.filter((d) => d.kind === 'Type' && !d.builtin);
/** The forms of expression, such as `a..b` and `!a`, and what each means by type. */
export const forms = ref.prelude.filter((d) => d.kind === 'Type' && d.builtin);
export const helperCount = ref.groups.reduce((n, g) => n + g.helpers.length, 0);

/** The reference's size, for pages that describe it. */
export const counts = {
	version: ref.version,
	kinds: ref.kinds.length + blueprints.length,
	helpers: helperCount,
	groups: ref.groups.length,
	types: ref.types.length,
	rules: ref.rules.length
};
