/**
 * Sanitises the HTML that lsp-client renders from the server's Markdown
 * (hover, completion and signature docs) before it goes into innerHTML.
 * Doc comments come from project files, and a project can be imported
 * from anyone, so only formatting tags survive, attributes are dropped
 * except `class` (the highlighting of code) and a link's href and title,
 * and links go only to http(s), # or a relative path.
 */

const allowed = new Set([
	'a', 'b', 'blockquote', 'br', 'code', 'del', 'div', 'em', 'h1', 'h2', 'h3', 'h4', 'h5', 'h6', 'hr', 'i',
	'li', 'ol', 'p', 'pre', 's', 'span', 'strong', 'sub', 'sup', 'table', 'tbody', 'td', 'th', 'thead', 'tr', 'ul',
]);
/** Elements dropped with everything inside them, not just unwrapped. */
const dropped = new Set(['script', 'style', 'iframe', 'object', 'embed', 'template', 'svg', 'math', 'noscript', 'form']);

export function safeHref(href: string): boolean {
	const h = href.trim();
	if (h.startsWith('#') || h.startsWith('/') || h.startsWith('./') || h.startsWith('../')) return true;
	if (/^https?:/i.test(h)) return true;
	// A relative path: no scheme before the first / ? or #.
	return !/^[^/?#]*:/.test(h);
}

function clean(node: Node) {
	for (const child of [...node.childNodes]) {
		if (child.nodeType === Node.TEXT_NODE) continue;
		if (child.nodeType !== Node.ELEMENT_NODE) {
			child.remove(); // comments, processing instructions
			continue;
		}
		const el = child as Element;
		const tag = el.tagName.toLowerCase();
		if (dropped.has(tag)) {
			el.remove();
			continue;
		}
		clean(el);
		if (!allowed.has(tag)) {
			el.replaceWith(...el.childNodes);
			continue;
		}
		for (const attr of [...el.attributes]) {
			const keep = attr.name === 'class' || (tag === 'a' && (attr.name === 'title' || (attr.name === 'href' && safeHref(attr.value))));
			if (!keep) el.removeAttribute(attr.name);
		}
	}
}

export function sanitizeHTML(html: string): string {
	const t = document.createElement('template');
	t.innerHTML = html;
	clean(t.content);
	const div = document.createElement('div');
	div.append(t.content);
	return div.innerHTML;
}
