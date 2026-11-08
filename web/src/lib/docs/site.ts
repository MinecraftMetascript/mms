import { Marked } from 'marked';

// The website's hand-written pages: every Markdown file under site/ in the
// repository, bundled at build time. site/ mirrors the URLs, so
// site/guides/functions.md is /docs/guides/functions.
const raw = import.meta.glob('../../../../site/**/*.md', {
	query: '?raw',
	import: 'default',
	eager: true
}) as Record<string, string>;

/** A sidebar group of hand-written pages: one folder of site/. */
export interface Group {
	/** The folder under site/, such as "guides" or "reference/language". */
	dir: string;
	title: string;
	/** Pages in this order come first; any others follow by title. */
	order: string[];
}

// The Docs tab teaches; the Reference tab is for looking things up.
export const docGroups: Group[] = [
	{ dir: 'start', title: 'Start', order: ['getting-started', 'tutorial-dimension'] },
	{ dir: 'guides', title: 'Guides', order: ['resources', 'functions', 'lint'] },
	{ dir: 'concepts', title: 'Concepts', order: ['how-functions-compile'] }
];
export const referenceGroups: Group[] = [
	{
		dir: 'reference/language',
		title: 'Language',
		order: ['declarations', 'values', 'reuse', 'blueprints', 'functions', 'text', 'tests']
	}
];

export interface Page {
	/** The path under /docs, such as "guides/functions". */
	path: string;
	dir: string;
	slug: string;
	title: string;
	/** The first paragraph, as plain text. */
	summary: string;
	source: string;
}

const prefix = '../../../../site/';

export const pages: Page[] = Object.entries(raw)
	.map(([file, source]) => {
		const path = file.slice(prefix.length).replace(/\.md$/, '');
		const dir = path.includes('/') ? path.slice(0, path.lastIndexOf('/')) : '';
		const slug = path.slice(dir.length ? dir.length + 1 : 0);
		const title = /^#\s+(.+)$/m.exec(source)?.[1] ?? slug;
		return { path, dir, slug, title, summary: summaryOf(source), source };
	})
	.sort((a, b) => a.path.localeCompare(b.path));

export const page = (path: string) => pages.find((p) => p.path === path);

/** The pages of a group, in sidebar order. */
export function groupPages(g: Group): Page[] {
	const mine = pages.filter((p) => p.dir === g.dir);
	const rank = (p: Page) => {
		const i = g.order.indexOf(p.slug);
		return i < 0 ? g.order.length : i;
	};
	return mine.sort((a, b) => rank(a) - rank(b) || a.title.localeCompare(b.title));
}

/** Every hand-written page in reading order: Docs, then Reference. */
export const readingOrder = (): Page[] => [...docGroups, ...referenceGroups].flatMap(groupPages);

/** The first paragraph of a page, without Markdown. */
function summaryOf(md: string): string {
	const para = md
		.replace(/^#.*$/gm, '')
		.split(/\n\s*\n/)
		.map((s) => s.trim())
		.find((s) => s && !/^(```|\||[-*] |\d+\. |<)/.test(s));
	return plain(para ?? '');
}

/** Markdown inline text as plain text. */
export function plain(md: string): string {
	return md
		.replace(/!?\[([^\]]*)\]\([^)]*\)/g, '$1')
		.replace(/[*_`]/g, '')
		.replace(/\s+/g, ' ')
		.trim();
}

const site = 'https://minecraftmetascript.github.io/mms/';
export const siteURL = site.replace(/\/$/, '');
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
 * Rewrites a link in a page at `from` (a path under /docs): other pages
 * (`../guides/functions.md#calls`) to their URL, the site's own URLs to
 * this build, and other repository paths to GitHub.
 */
export function rewriteHref(href: string, base: string, from: string): string {
	if (href.startsWith(site)) return `${base}/${href.slice(site.length)}`;
	if (/^[a-z]+:|^#/.test(href)) return href;
	const url = new URL(href, `https://x/site/${from}`);
	const m = /^\/site\/(.+)\.md$/.exec(url.pathname);
	if (m && page(m[1])) return `${base}/docs/${m[1]}${url.hash}`;
	return repo + url.pathname.slice(1) + url.hash;
}

const escape = (s: string) =>
	s.replace(/[&<>"]/g, (c) => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;' })[c]!);

/** Renders a page's Markdown to HTML; base is the site's base path. */
export function render(p: Page, base: string): string {
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
				const to = rewriteHref(href, base, p.path);
				const ext = /^https?:/.test(to) && !to.startsWith(base + '/');
				const t = title ? ` title="${escape(title)}"` : '';
				return `<a href="${escape(to)}"${t}${ext ? ' rel="noopener"' : ''}>${inner}</a>`;
			}
		}
	});
	return marked.parse(p.source, { async: false }) as string;
}

/** The h2 and h3 sections of a page: heading, anchor and plain text. */
export function sections(p: Page): { title: string; id: string; text: string }[] {
	const out: { title: string; id: string; text: string }[] = [];
	let fence = false;
	let cur: { title: string; id: string; lines: string[] } | undefined;
	for (const line of p.source.split('\n')) {
		if (line.startsWith('```')) fence = !fence;
		const h = !fence && /^###?\s+(.+)$/.exec(line);
		if (h) {
			if (cur) out.push({ title: cur.title, id: cur.id, text: plain(cur.lines.join(' ')) });
			cur = { title: plain(h[1]), id: slugify(h[1]), lines: [] };
		} else if (cur && !fence) cur.lines.push(line);
	}
	if (cur) out.push({ title: cur.title, id: cur.id, text: plain(cur.lines.join(' ')) });
	return out;
}

/** A page as a card or list item: no source. */
export const card = (p: Page) => ({ path: p.path, title: p.title, summary: p.summary });
