; Book markup parsed with the Markdown block grammar: `# heading` lines,
; `---` page breaks (which Markdown may read as a heading underline) and
; backslash escapes.
(atx_heading) @title
(atx_h1_marker) @punctuation.special
(thematic_break) @punctuation.special
(setext_h2_underline) @punctuation.special
(backslash_escape) @string.escape
