package ast

type Location struct {
	// Line 0-indexed
	Line int `json:"line"`
	// Column 0-indexed
	Column int `json:"column"`
	// Index Character index from the beginning of the file
	Index int `json:"index"`
}

type SourceLocation struct {
	Start Location
	Stop  Location
}

func (sl SourceLocation) Intersects(other SourceLocation) bool {
	return sl.ContainsLocation(other.Start) || sl.ContainsLocation(other.Stop)
}

func (sl SourceLocation) Contains(other SourceLocation) bool {
	return sl.ContainsLocation(other.Start) && sl.ContainsLocation(other.Stop)
}

func (sl SourceLocation) ContainsLocation(l Location) bool {
	lineContained := sl.Start.Line <= l.Line && sl.Stop.Line >= l.Line
	columnContained := sl.Start.Column <= l.Column && l.Column >= l.Column
	return lineContained && columnContained
}

type Diagnostic struct {
	Message  string `json:"message"`
	Severity string `json:"severity"`
	Source   string `json:"source"`
	Filename string `json:"filename"`
	Location SourceLocation
}
