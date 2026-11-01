package syntax

// Incremental parsing. A file is split into chunks at top-level
// declaration boundaries; chunks whose text did not change are reused
// rather than parsed again, so an edit costs one declaration's parse.

// Chunk is one independently parsed piece of a file.
type Chunk struct {
	Text string
	Off  int   // rune offset in the file
	File *File // spans are absolute for Off
}

// Reparse parses src, reusing chunks from prev (the previous result's
// chunks; nil for a first parse). Moved maps each declaration taken from a
// reused chunk at a new offset to its old node, with the old-to-new node
// map; declarations reused in place keep their identity.
func Reparse(src string, prev []Chunk) (*File, []Chunk, map[*Decl]Moved) {
	runes := []rune(src)
	bounds := splitChunks(runes)
	byText := map[string][]Chunk{}
	for _, c := range prev {
		byText[c.Text] = append(byText[c.Text], c)
	}
	f := &File{Src: runes}
	chunks := make([]Chunk, 0, len(bounds))
	moved := map[*Decl]Moved{}
	for i, start := range bounds {
		end := len(runes)
		if i+1 < len(bounds) {
			end = bounds[i+1]
		}
		text := string(runes[start:end])
		var cf *File
		if olds := byText[text]; len(olds) > 0 {
			old := olds[0]
			byText[text] = olds[1:]
			if old.Off == start {
				cf = old.File
			} else {
				nodes := map[Node]Node{}
				cf = shiftFile(old.File, start-old.Off, nodes)
				for j, d := range cf.Decls {
					moved[d] = Moved{Old: old.File.Decls[j], Nodes: nodes, Delta: start - old.Off}
				}
			}
		} else {
			cf = Parse(text)
			cf = shiftFile(cf, start, nil)
		}
		chunks = append(chunks, Chunk{Text: text, Off: start, File: cf})
		f.Decls = append(f.Decls, cf.Decls...)
		f.Items = append(f.Items, cf.Items...)
		f.Tokens = append(f.Tokens, cf.Tokens...)
		f.Diags = append(f.Diags, cf.Diags...)
	}
	return f, chunks, moved
}

// Moved records a declaration reused at a new offset.
type Moved struct {
	Old   *Decl
	Nodes map[Node]Node // old node to new node
	Delta int
}

// splitChunks returns chunk start offsets: 0, then the start of each line
// that begins a top-level declaration (`Kind name` in column 0 outside any
// bracket, comment or command), moved up over comment lines directly above
// so doc comments stay with their declaration.
func splitChunks(src []rune) []int {
	bounds := []int{0}
	depth := 0
	inBlock := false
	// commentRun is the start of the run of comment-only lines ending at
	// the current line, or -1.
	commentRun := -1
	for i := 0; i < len(src); {
		lineStart := i
		topLevel := depth == 0 && !inBlock
		if topLevel && lineStart > 0 && startsDecl(src, i) {
			b := lineStart
			if commentRun >= 0 {
				b = commentRun
			}
			if b > bounds[len(bounds)-1] {
				bounds = append(bounds, b)
			}
		}
		// Scan the line.
		content := false
		comment := inBlock
		j := i
		for j < len(src) && (src[j] == ' ' || src[j] == '\t') {
			j++
		}
		if !inBlock && j+1 < len(src) && src[j] == '/' && src[j+1] >= 'a' && src[j+1] <= 'z' {
			// A command: verbatim to the end of the line.
			content = true
			for j < len(src) && src[j] != '\n' {
				j++
			}
		}
		for j < len(src) && src[j] != '\n' {
			r := src[j]
			switch {
			case inBlock:
				if r == '*' && j+1 < len(src) && src[j+1] == '/' {
					inBlock = false
					j++
				}
			case r == '/' && j+1 < len(src) && src[j+1] == '/':
				comment = true
				for j < len(src) && src[j] != '\n' {
					j++
				}
				continue
			case r == '/' && j+1 < len(src) && src[j+1] == '*':
				inBlock, comment = true, true
				j++
			case r == '"' || r == '\'':
				content = true
				for j++; j < len(src) && src[j] != r && src[j] != '\n'; j++ {
					if src[j] == '\\' {
						j++
					}
				}
				if j < len(src) && src[j] == '\n' {
					continue
				}
			case r == '{' || r == '(' || r == '[':
				content = true
				depth++
			case r == '}' || r == ')' || r == ']':
				content = true
				if depth > 0 {
					depth--
				}
			case r == ' ' || r == '\t' || r == '\r':
			default:
				content = true
			}
			j++
		}
		// A line inside a block comment continues a run of comment lines
		// but never starts one: the comment may have opened after code.
		if comment && !content && depth == 0 && (topLevel || commentRun >= 0) {
			if commentRun < 0 {
				commentRun = lineStart
			}
		} else {
			commentRun = -1
		}
		i = j + 1
	}
	return bounds
}

// startsDecl reports whether the line at i starts `Ident name` (the same
// test the parser uses to stop a missing value from swallowing the next
// declaration).
func startsDecl(src []rune, i int) bool {
	isStart := func(r rune) bool { return r == '_' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' }
	isChar := func(r rune) bool { return isStart(r) || r >= '0' && r <= '9' }
	if i >= len(src) || !isStart(src[i]) {
		return false
	}
	for i < len(src) && isChar(src[i]) {
		i++
	}
	if i >= len(src) || (src[i] != ' ' && src[i] != '\t') {
		return false
	}
	for i < len(src) && (src[i] == ' ' || src[i] == '\t') {
		i++
	}
	return i < len(src) && (isChar(src[i]) || src[i] == '.' || src[i] == '-')
}
