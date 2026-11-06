package syntax

// CloneExpr returns a deep copy of x with fresh nodes and the same spans.
func CloneExpr(x Expr) Expr { return shifter{}.expr(x) }

// CloneExprAt returns a deep copy of x with every span set to sp: a value
// written elsewhere, copied in where sp caused it.
func CloneExprAt(x Expr, sp Span) Expr { return shifter{fix: &sp}.expr(x) }

// CloneStmt returns a deep copy of st with fresh nodes and the same spans.
func CloneStmt(st Stmt) Stmt { return shifter{}.stmt(st) }

// CloneStmtAt returns a deep copy of st with every span set to sp.
func CloneStmtAt(st Stmt, sp Span) Stmt { return shifter{fix: &sp}.stmt(st) }

// CloneBlock returns a deep copy of b with fresh nodes and the same spans.
func CloneBlock(b *Block) *Block { return shifter{}.block(b) }

// CloneDecl returns a deep copy of d with fresh nodes and the same spans.
func CloneDecl(d *Decl) *Decl { return shifter{}.decl(d) }

// CloneTypeAt returns a deep copy of t with every span set to sp.
func CloneTypeAt(t TypeExpr, sp Span) TypeExpr { return shifter{fix: &sp}.typ(t) }

// ParseExpr parses text as one expression, such as the inside of a
// `${...}` in a string or a command, with spans moved by off.
func ParseExpr(text string, off int) (Expr, []Diagnostic) {
	f := &File{}
	x := parseExpr(text, off, f)
	return x, f.Diags
}

// CloneType returns a deep copy of t with the same spans.
func CloneType(t TypeExpr) TypeExpr { return shifter{}.typ(t) }
