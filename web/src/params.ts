import { defineParams } from '@sveltejs/kit/params';

// /docs/<page>.md is the page as Markdown, for AI tools and scripts; the
// HTML routes take every other path, so the two never compete.
export const params = defineParams({
	markdown: (param) => (param.endsWith('.md') ? param : undefined),
	page: (param) => (param.endsWith('.md') ? undefined : param)
});
