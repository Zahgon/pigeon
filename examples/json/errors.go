package json

// ErrorLister is the public interface to access the inner errors
// included in a errList
type ErrorLister interface {
	Errors() []error
}

func (e errList) Errors() []error {
	_ = "STUB: not implemented"

	// ParserError is the public interface to errors of type parserError
	return nil
}

type ParserError interface {
	Error() string
	InnerError() error
	Pos() (int, int, int)
	Expected() []string
}

func (p *parserError) InnerError() error { _ = "STUB: not implemented"; return nil }

func (p *parserError) Pos() (line, col, offset int) { _ = "STUB: not implemented"; return 0, 0, 0 }

func (p *parserError) Expected() []string { _ = "STUB: not implemented"; return nil }
