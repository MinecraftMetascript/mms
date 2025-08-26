package primitives

type Docable interface {
	Comment() *string
	SetComment(string)
}

type BaseDocable struct {
	comment *string
}

func (b *BaseDocable) Comment() *string          { return b.comment }
func (b *BaseDocable) SetComment(comment string) { b.comment = &comment }
