// External scanner for tree-sitter-mms.
//
// A raw command is a line in a function body whose first non-blank
// character is '/' followed by a lowercase letter (MMSLexer.g4, COMMAND).
// Whether '/' starts a line is not something a regular lexer can see, so
// this scanner lexes the command name and the text between `${...}`
// interpolations.

#include "tree_sitter/parser.h"

enum TokenType {
  COMMAND_NAME,
  COMMAND_TEXT,
  ERROR_SENTINEL,
};

void *tree_sitter_mms_external_scanner_create(void) { return NULL; }
void tree_sitter_mms_external_scanner_destroy(void *payload) { (void)payload; }
unsigned tree_sitter_mms_external_scanner_serialize(void *payload, char *buffer) {
  (void)payload;
  (void)buffer;
  return 0;
}
void tree_sitter_mms_external_scanner_deserialize(void *payload, const char *buffer, unsigned length) {
  (void)payload;
  (void)buffer;
  (void)length;
}

static bool is_lower(int32_t c) { return c >= 'a' && c <= 'z'; }
static bool is_newline(int32_t c) { return c == '\n' || c == '\r'; }

// The rest of a command line, up to the next `${` or the end of the line.
static bool scan_command_text(TSLexer *lexer) {
  bool any = false;
  for (;;) {
    int32_t c = lexer->lookahead;
    if (lexer->eof(lexer) || is_newline(c)) break;
    if (c == '$') {
      lexer->mark_end(lexer);
      lexer->advance(lexer, false);
      if (lexer->lookahead == '{') {
        // Stop before `${`; the grammar lexes it.
        if (!any) return false;
        lexer->result_symbol = COMMAND_TEXT;
        return true;
      }
      any = true;
      continue;
    }
    lexer->advance(lexer, false);
    any = true;
  }
  if (!any) return false;
  lexer->mark_end(lexer);
  lexer->result_symbol = COMMAND_TEXT;
  return true;
}

bool tree_sitter_mms_external_scanner_scan(void *payload, TSLexer *lexer, const bool *valid_symbols) {
  (void)payload;
  if (valid_symbols[ERROR_SENTINEL]) return false;

  // Command text follows the name or a closing interpolation brace
  // directly, so it never skips whitespace first: the spaces are part of
  // the text.
  if (valid_symbols[COMMAND_TEXT] && !is_newline(lexer->lookahead)) {
    return scan_command_text(lexer);
  }

  if (!valid_symbols[COMMAND_NAME]) return false;

  bool line_start = lexer->get_column(lexer) == 0;
  while (lexer->lookahead == ' ' || lexer->lookahead == '\t' || is_newline(lexer->lookahead)) {
    if (is_newline(lexer->lookahead)) line_start = true;
    lexer->advance(lexer, true);
  }
  if (!line_start || lexer->lookahead != '/') return false;
  lexer->advance(lexer, false);
  if (!is_lower(lexer->lookahead)) return false;
  while (is_lower(lexer->lookahead) || lexer->lookahead == '_') lexer->advance(lexer, false);
  lexer->mark_end(lexer);
  lexer->result_symbol = COMMAND_NAME;
  return true;
}
