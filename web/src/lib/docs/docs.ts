import { Marked } from 'marked';

// The guides in docs/, bundled at build time. Order is the sidebar's.
const raw = import.meta.glob('../../../../docs/*.md', {
	query: '?raw',
	import: 'default',
	eager: true
}) as Record<string, string>;

const order = [
	'getting-started',
	'resources',
	'functions',
	'build-time',
	'lint',
	'tutorial-dimension'
];

export interface Guide {
	slug: string;
	title: string;
	source: string;
}

export const guides: Guide[] = order
	.map((slug) => {
		const source = raw[`../../../../docs/${slug}.md`] ?? '';
		const title = /^#\s+(.+)$/m.exec(source)?.[1] ?? slug;
		return { slug, title, source };
	})
	.filter((g) => g.source);

export const guide = (slug: string) => guides.find((g) => g.slug === slug);

const site = 'https://minecraftmetascript.github.io/mms/';
const repo = 'https://github.com/MinecraftMetascript/mms/blob/main/';

/** GitHub's heading anchors: lower case, punctuation dropped, spaces to dashes. */
export function slugify(text: string): string {
	return text
		.toLowerCase()
		.replace(/<[^>]+>/g, '')
		.replace(/[^\p{L}\p{N}\s_-]/gu, '')
		.trim()
		.replace(/\s/g, '-');
}

/**
 * Rewrites a link in a guide: other guides (`functions.md#calls`) to their
 * page, the site's own URLs to this build, and repository paths to GitHub.
 */
export function rewriteHref(href: string, base: string): string {
	if (href.startsWith(site)) return `${base}/${href.slice(site.length)}`;
	if (/^[a-z]+:|^#/.test(href)) return href;
	const m = /^([\w-]+)\.md(#.*)?$/.exec(href);
	if (m && guide(m[1])) return `${base}/docs/${m[1]}${m[2] ?? ''}`;
	return repo + new URL(href, 'https://x/docs/').pathname.slice(1);
}

const escape = (s: string) =>
	s.replace(/[&<>"]/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' })[c]!);

/** Renders a guide's Markdown to HTML; base is the site's base path. */
export function render(md: string, base: string): string {
	const marked = new Marked({
		gfm: true,
		renderer: {
			heading({ tokens, depth, text }) {
				const id = slugify(text);
				const inner = this.parser.parseInline(tokens);
				return `<h${depth} id="${id}"><a href="#${id}">${inner}</a></h${depth}>\n`;
			},
			link({ href, title, tokens }) {
				const inner = this.parser.parseInline(tokens);
				const to = rewriteHref(href, base);
				const ext = /^https?:/.test(to) && !to.startsWith(base + '/');
				const t = title ? ` title="${escape(title)}"` : '';
				return `<a href="${escape(to)}"${t}${ext ? ' rel="noopener"' : ''}>${inner}</a>`;
			}
		}
	});
	return marked.parse(md, { async: false }) as string;
}
