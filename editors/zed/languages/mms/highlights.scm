; Highlights for Minecraft Metascript. Capture names follow Zed's theme
; scopes, and as in Zed a later pattern wins over an earlier one, so the
; general rules come first. editors/zed/languages/mms/highlights.scm is a
; copy of this file (CI checks).

(identifier) @variable

; Declarations: `Kind name ...`. Kinds are PascalCase spec names.
(declaration kind: (identifier) @type)
(declaration name: (_) @constant)
((declaration kind: (identifier) @_kind name: (identifier) @function)
  (#eq? @_kind "Function"))

(parameter name: (identifier) @variable.parameter)
(type (_) @type)
(type (identifier) @type.builtin (#match? @type.builtin "^(int|bool|string|float|selector|any)$"))

(let_statement name: (identifier) @variable)
(for_statement name: (identifier) @variable)

; `key = value` in resource bodies and object literals.
(assignment left: (identifier) @property)
(assignment left: (string) @property)
(assignment left: (resource_location) @property)
(argument name: (identifier) @property)
(member_expression property: (identifier) @property)

; Calls: PascalCase builders and helpers, lowercase user functions.
(call_expression function: (identifier) @function)
(call_expression function: (member_expression property: (identifier) @function.method))
(block_call_expression function: (identifier) @function)
(block_call_expression function: (call_expression function: (identifier) @function))

; Raw commands.
(command_name) @function.builtin
(command_text) @string.special
(interpolation ["${" "}"] @punctuation.special)

[
  "let"
  "if"
  "else"
  "while"
  "for"
  "in"
  "return"
] @keyword
(break_statement) @keyword
(continue_statement) @keyword

(boolean) @boolean
(number) @number
(string) @string
(resource_location) @string.special.symbol
(tag) @tag
(selector) @variable.special
(coordinates ["~" "^"] @operator)

[
  "=" "+=" "-=" "*=" "/=" "%="
  "+" "-" "*" "/" "%"
  "==" "!=" "<" "<=" ">" ">="
  "&&" "||" "!"
  "->" ".."
] @operator

["(" ")" "[" "]" "{" "}"] @punctuation.bracket
["," "." ":"] @punctuation.delimiter

(line_comment) @comment
(block_comment) @comment
(doc_comment) @comment.doc

