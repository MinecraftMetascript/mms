/**
 * Tree-sitter grammar for Minecraft Metascript (.mms).
 *
 * It mirrors engine/syntax/grammar/MMSParser.g4 closely enough for
 * highlighting, brackets, indentation and the outline in editors such as
 * Zed. The ANTLR grammar stays the source of truth: the engine never uses
 * this one, and it is more lenient where that makes recovery easier
 * (list elements need no separator, for example).
 *
 * @file Minecraft Metascript grammar for tree-sitter
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

const PREC = {
  block_call: 0,
  for: 1,
  arrow: 1,
  or: 2,
  and: 3,
  equal: 4,
  compare: 5,
  range: 6,
  add: 7,
  mul: 8,
  unary: 9,
  postfix: 10,
};

const commaSep = (rule) => optional(seq(rule, repeat(seq(',', rule)), optional(',')));

module.exports = grammar({
  name: 'mms',

  // Commands are only commands at the start of a line, which a regular
  // lexer cannot see, so src/scanner.c lexes them (see MMSLexer.g4,
  // COMMAND). `_error_sentinel` is never used by a rule; it tells the
  // scanner that tree-sitter is in error recovery.
  externals: ($) => [$.command_name, $.command_text, $._error_sentinel],

  extras: ($) => [/\s/, $.line_comment, $.block_comment, $.doc_comment],

  word: ($) => $.identifier,

  rules: {
    source_file: ($) => repeat($.declaration),

    // Noise ridges = Noise(-7)
    // LootTable blocks/ruby_ore { pools = [ ... ] }
    // Global deaths: int = 0
    // Function add(a: int, b: int): int { return a + b }
    // OnLoad { ... }
    declaration: ($) =>
      choice(
        prec.right(
          seq(
            field('kind', $.identifier),
            field('name', choice($.identifier, $.resource_location)),
            optional(
              choice(
                $.type_annotation,
                seq($.parameters, optional($.type_annotation)),
              ),
            ),
            optional(choice(seq('=', optional(field('value', $._expression))), field('body', $.block))),
          ),
        ),
        seq(field('kind', $.identifier), optional($.parameters), field('body', $.block)),
      ),

    type_annotation: ($) => seq(':', field('type', $.type)),
    type: ($) => choice($.identifier, $.resource_location),

    parameters: ($) => seq('(', commaSep($.parameter), ')'),
    parameter: ($) => seq(field('name', $.identifier), optional($.type_annotation)),

    // A block holds fields in a resource body or an object literal, and
    // statements in a function body; the checker decides which.
    block: ($) => seq('{', repeat(seq($._statement, optional(','))), '}'),

    _statement: ($) =>
      choice(
        $.let_statement,
        $.if_statement,
        $.while_statement,
        $.for_statement,
        $.return_statement,
        $.break_statement,
        $.continue_statement,
        $.command,
        $.assignment,
        $.compound_assignment,
        $.expression_statement,
      ),

    let_statement: ($) =>
      prec.right(
        seq(
          'let',
          field('name', $.identifier),
          optional($.type_annotation),
          optional(seq('=', field('value', $._expression))),
        ),
      ),

    if_statement: ($) =>
      prec.right(
        seq(
          'if',
          '(',
          field('condition', $._expression),
          ')',
          field('consequence', $.block),
          optional(seq('else', field('alternative', choice($.if_statement, $.block)))),
        ),
      ),

    while_statement: ($) => seq('while', '(', field('condition', $._expression), ')', field('body', $.block)),

    for_statement: ($) =>
      prec(PREC.for, seq('for', field('name', $.identifier), 'in', field('iterable', $._expression), field('body', $.block))),

    return_statement: ($) => prec.right(seq('return', optional($._expression))),
    break_statement: (_) => 'break',
    continue_statement: (_) => 'continue',

    // /say fib(10) is ${fib(10)}
    command: ($) =>
      prec.right(seq($.command_name, repeat(choice($.command_text, $.interpolation)))),
    interpolation: ($) => seq('${', $._expression, '}'),

    // `key = value` in a resource body or object literal; a variable,
    // member or index assignment in a function body.
    assignment: ($) =>
      prec(-1, seq(field('left', $._expression), '=', field('right', $._expression))),

    compound_assignment: ($) =>
      prec(-1,
        seq(
          field('target', $._expression),
          field('operator', choice('+=', '-=', '*=', '/=', '%=')),
          field('value', $._expression),
        ),
      ),

    expression_statement: ($) => prec(-1, $._expression),

    _expression: ($) =>
      choice(
        $.unary_expression,
        $.binary_expression,
        $.range_expression,
        $.arrow_expression,
        $.member_expression,
        $.call_expression,
        $.index_expression,
        $.block_call_expression,
        $._primary,
      ),

    unary_expression: ($) =>
      prec(PREC.unary, seq(field('operator', choice('!', '-', '+')), field('operand', $._expression))),

    binary_expression: ($) => {
      const table = [
        [PREC.mul, choice('*', '/', '%')],
        [PREC.add, choice('+', '-')],
        [PREC.compare, choice('<', '<=', '>', '>=')],
        [PREC.equal, choice('==', '!=')],
        [PREC.and, '&&'],
        [PREC.or, '||'],
      ];
      return choice(
        ...table.map(([p, op]) =>
          prec.left(
            /** @type {number} */ (p),
            seq(field('left', $._expression), field('operator', /** @type {any} */ (op)), field('right', $._expression)),
          ),
        ),
      );
    },

    range_expression: ($) =>
      choice(
        prec.right(PREC.range, seq(field('start', $._expression), '..', optional(field('end', $._expression)))),
        prec.right(PREC.range, seq('..', field('end', $._expression))),
      ),

    arrow_expression: ($) =>
      prec.right(PREC.arrow, seq(field('left', $._expression), '->', field('right', $._expression))),

    member_expression: ($) =>
      prec.right(PREC.postfix, seq(field('object', $._expression), '.', optional(field('property', $.identifier)))),

    call_expression: ($) =>
      prec(PREC.postfix, seq(field('function', $._expression), field('arguments', $.arguments))),

    index_expression: ($) =>
      prec(PREC.postfix, seq(field('object', $._expression), '[', commaSep($.argument), ']')),

    // Builder(1) { field = value }
    // Block calls bind loosest, so `for n in 0..count { ... }` reads the
    // braces as the loop body (for_statement wins the tie). Grouping like
    // `a + B() { }` is wrong but harmless for highlighting.
    block_call_expression: ($) =>
      prec.left(PREC.block_call, seq(field('function', $._expression), field('body', $.block))),

    arguments: ($) => seq('(', commaSep($.argument), ')'),
    argument: ($) => seq(optional(seq(field('name', $.identifier), '=')), field('value', $._expression)),

    _primary: ($) =>
      choice(
        $.number,
        $.string,
        $.boolean,
        $.identifier,
        $.resource_location,
        $.tag,
        $.selector,
        $.coordinates,
        $.parenthesized_expression,
        $.list,
        $.block,
      ),

    parenthesized_expression: ($) => seq('(', $._expression, ')'),

    // Elements may be separated by a comma or a newline (MMSParser.g4 checks
    // the newline; this grammar just makes the comma optional).
    list: ($) => seq('[', repeat(prec(-1, seq($._expression, optional(',')))), ']'),

    tag: ($) => seq('#', choice($.identifier, $.resource_location)),
    selector: ($) => seq('@', $.identifier),
    coordinates: ($) => prec.right(repeat1(seq(choice('~', '^'), optional(seq(optional('-'), $.number))))),

    boolean: (_) => choice('true', 'false'),

    number: (_) => /\d+(\.\d+)?([eE][+-]?\d+)?/,

    string: (_) =>
      token(
        choice(
          seq('"', repeat(choice(/[^"\\\r\n]/, /\\./)), '"'),
          seq("'", repeat(choice(/[^'\\\r\n]/, /\\./)), "'"),
        ),
      ),

    // minecraft:stone, blocks/ruby_ore, example:terrain/noise_0
    resource_location: (_) =>
      token(
        choice(
          /[a-z0-9_.\-]+:[a-z0-9_.\/\-]+/,
          /[a-zA-Z_][a-zA-Z0-9_]*(\/[a-z0-9_.\/\-]+)+/,
        ),
      ),

    identifier: (_) => /[a-zA-Z_][a-zA-Z0-9_]*/,

    doc_comment: (_) => token(prec(1, seq('/**', /[^*]*\*+([^/*][^*]*\*+)*/, '/'))),
    block_comment: (_) => token(seq('/*', /[^*]*\*+([^/*][^*]*\*+)*/, '/')),
    line_comment: (_) => token(seq('//', /[^\r\n]*/)),
  },
});
