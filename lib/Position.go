package lib

import protocol "github.com/tliron/glsp/protocol_3_16"

func AddCols(p protocol.Position, cols int) protocol.Position {
	return protocol.Position{
		Line:      p.Line,
		Character: p.Character + protocol.UInteger(cols),
	}
}
