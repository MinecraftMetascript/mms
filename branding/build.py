"""Builds the MMS brand SVGs and web/src/lib/components/Logo.svelte. See branding/README.md."""
import os, sys
from fontTools.ttLib import TTFont
from fontTools.pens.svgPathPen import SVGPathPen
from fontTools.pens.transformPen import TransformPen

FONTS, OUT = sys.argv[1], sys.argv[2]
os.makedirs(OUT, exist_ok=True)

LIGHT = dict(ink="#1b1f24", top="#f2b01e", s1="#2563eb", s2="#1a46c4", sub="#65707d")
DARK = dict(ink="#e6e9ed", top="#fbbf24", s1="#3b82f6", s2="#2459d6", sub="#8d97a3")
TILE = "#2563eb"

def f(n):
    return f"{n:.2f}".rstrip("0").rstrip(".")

def cube(cx, cy, w, h, top, left, right):
    t = w * 0.57735
    P = lambda pts: " ".join(f"{f(x)},{f(y)}" for x, y in pts)
    return (f'<path fill="{top}" d="M{P([(cx,cy-t),(cx+w,cy),(cx,cy+t),(cx-w,cy)])}Z"/>'
            f'<path fill="{left}" d="M{P([(cx-w,cy),(cx,cy+t),(cx,cy+t+h),(cx-w,cy+h)])}Z"/>'
            f'<path fill="{right}" d="M{P([(cx,cy+t),(cx+w,cy),(cx+w,cy+h),(cx,cy+t+h)])}Z"/>')

BR = "M15 9C9 9 10 14 10 20v6c0 4-2 6-6 6 4 0 6 2 6 6v6c0 6-1 11 5 11"
BRR = "M49 9c6 0 5 5 5 11v6c0 4 2 6 6 6-4 0-6 2-6 6v6c0 6 1 11-5 11"

def braces(stroke):
    return (f'<g fill="none" stroke="{stroke}" stroke-width="5" stroke-linecap="round" stroke-linejoin="round">'
            f'<path d="{BR}"/><path d="{BRR}"/></g>')

def mark_body(p, stroke=None):
    return braces(stroke or p["ink"]) + cube(32, 24, 13, 15, p["top"], p["s1"], p["s2"])

def svg(vb, body, title=None, extra=""):
    t = f"<title>{title}</title>" if title else ""
    return f'<svg xmlns="http://www.w3.org/2000/svg" viewBox="{vb}"{extra}>{t}{body}</svg>\n'

def text_path(font, text, size, x, baseline, tracking_em):
    gs = font.getGlyphSet()
    cmap = font.getBestCmap()
    upm = font["head"].unitsPerEm
    s = size / upm
    out, pen_x = [], x
    for i, ch in enumerate(text):
        g = cmap[ord(ch)]
        sp = SVGPathPen(gs)
        gs[g].draw(TransformPen(sp, (s, 0, 0, -s, pen_x, baseline)))
        out.append(sp.getCommands())
        pen_x += gs[g].width * s
        if i < len(text) - 1:
            pen_x += tracking_em * size
    return "".join(out), pen_x

bold = TTFont(f"{FONTS}/cp700.ttf")
medium = TTFont(f"{FONTS}/cp500.ttf")

# Marks
for mode, p in (("light", LIGHT), ("dark", DARK)):
    open(f"{OUT}/mark-{mode}.svg", "w").write(svg("0 0 64 64", mark_body(p), "MMS"))
# Adaptive mark: braces follow the colour scheme of whatever displays it.
adaptive = ('<style>.b{stroke:#1b1f24}@media (prefers-color-scheme:dark){.b{stroke:#e6e9ed}}</style>'
            '<g class="b" fill="none" stroke-width="5" stroke-linecap="round" stroke-linejoin="round">'
            f'<path d="{BR}"/><path d="{BRR}"/></g>' + cube(32, 24, 13, 15, LIGHT["top"], LIGHT["s1"], LIGHT["s2"]))
open(f"{OUT}/mark.svg", "w").write(svg("0 0 64 64", adaptive, "MMS"))

# Favicon / app icon: white braces and the cube on a blue tile, readable at 16px.
fav = (f'<rect width="16" height="16" rx="3.5" fill="{TILE}"/>'
       '<g fill="none" stroke="#fff" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round">'
       '<path d="M4.6 2.8C3 2.8 3.2 4.2 3.2 5.6v1c0 .8-.4 1.4-1.2 1.4.8 0 1.2.6 1.2 1.4v1c0 1.4-.2 2.8 1.4 2.8"/>'
       '<path d="M11.4 2.8c1.6 0 1.4 1.4 1.4 2.8v1c0 .8.4 1.4 1.2 1.4-.8 0-1.2.6-1.2 1.4v1c0 1.4.2 2.8-1.4 2.8"/></g>'
       + cube(8, 6.2, 3.6, 4.2, LIGHT["top"], "#e3ecff", "#b4ccff"))
open(f"{OUT}/favicon.svg", "w").write(svg("0 0 16 16", fav, "MMS"))

# Lockups: mark + "MMS" + "METASCRIPT", text as outlines.
main_d, main_end = text_path(bold, "MMS", 48, 94, 46, 0.04)
sub_d, sub_end = text_path(medium, "METASCRIPT", 13, 95, 67, 0.3)
width = f(max(main_end, sub_end) + 2)
for mode, p in (("light", LIGHT), ("dark", DARK)):
    body = (f'<g transform="scale(1.25)">{mark_body(p)}</g>'
            f'<path fill="{p["ink"]}" d="{main_d}"/><path fill="{p["sub"]}" d="{sub_d}"/>')
    open(f"{OUT}/logo-{mode}.svg", "w").write(svg(f"0 0 {width} 80", body, "MMS Metascript"))
print("width", width)

# Svelte component for app headers: the mark plus "MMS", braces and text in currentColor.
cap = bold["OS/2"].sCapHeight / bold["head"].unitsPerEm
size = 38
base = 32 + cap * size / 2
d, end = text_path(bold, "MMS", size, 74, base, 0.04)
comp = f'''<!-- Generated from branding/ (see branding/README.md). The MMS mark, optionally with the "MMS" wordmark. -->
<script lang="ts">
	let {{ wordmark = true, class: cls = '' }}: {{ wordmark?: boolean; class?: string }} = $props();
</script>

<svg class={{cls}} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 {{wordmark ? {f(end + 1)} : 64}} 64" role="img" aria-label="MMS">
	<g fill="none" stroke="currentColor" stroke-width="5" stroke-linecap="round" stroke-linejoin="round"><path d="{BR}" /><path d="{BRR}" /></g>
	{cube(32, 24, 13, 15, LIGHT["top"], LIGHT["s1"], LIGHT["s2"]).replace("/>", " />")}
	{{#if wordmark}}<path fill="currentColor" d="{d}" />{{/if}}
</svg>
'''
open(sys.argv[3], "w").write(comp)
