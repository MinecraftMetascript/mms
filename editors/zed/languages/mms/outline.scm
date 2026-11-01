; Every top-level declaration: `Noise ridges`, `Function add`,
; `LootTable blocks/ruby_ore`, and nameless ones such as `OnLoad`.
(declaration
  kind: (identifier) @context
  name: (_) @name) @item

(declaration
  kind: (identifier) @name
  !name) @item

; Build-time loops group the declarations they make.
(top_for
  "for" @context
  name: (identifier) @name) @item

; Doc comments sit right before declarations.
(doc_comment) @annotation
