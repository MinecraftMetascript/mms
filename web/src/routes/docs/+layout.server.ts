import { docsNav, referenceNav } from '#lib/docs/nav.ts';
import type { LayoutServerLoad } from './$types';

// The sidebars are built at build time, so the generated reference stays
// out of the pages' JavaScript.
export const load: LayoutServerLoad = () => ({ docsNav, referenceNav });
