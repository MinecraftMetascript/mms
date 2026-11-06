import raw from './reference.json';
import type { Reference } from './reference.ts';

// The output of `mms reference`, generated into reference.json by
// scripts/build-wasm.sh.
export const ref = raw as Reference;

/** Resource kinds (a data or resource pack folder each), and MMS's own. */
export const resourceKinds = ref.kinds.filter((k) => k.folder);
export const ownKinds = ref.kinds.filter((k) => !k.folder);
export const helperCount = ref.groups.reduce((n, g) => n + g.helpers.length, 0);

/** The reference's size, for pages that describe it. */
export const counts = {
	version: ref.version,
	kinds: ref.kinds.length,
	helpers: helperCount,
	groups: ref.groups.length,
	types: ref.types.length,
	rules: ref.rules.length
};
