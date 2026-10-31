(declaration
  body: (block "{" (_)* @function.inside "}")) @function.around

(declaration
  value: (_) @function.inside) @function.around

(doc_comment) @comment.around
(line_comment) @comment.around
(block_comment) @comment.around
