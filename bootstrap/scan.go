package bootstrap

import (
	"bytes"
	"io"

	"github.com/mna/pigeon/ast"
)

// Scanner tokenizes an input source for the PEG grammar.
type Scanner struct {
	r    io.RuneReader
	errh func(ast.Pos, error)

	eof  bool
	cur  rune
	cpos ast.Pos
	cw   int

	tok bytes.Buffer
}

// Init initializes the scanner to read and tokenize text from r.
func (s *Scanner) Init(filename string, r io.Reader, errh func(ast.Pos, error)) {
	_ = "STUB: not implemented"
	return
}

// Scan returns the next token, along with a boolean indicating if EOF was
// reached (false means no more tokens).
func (s *Scanner) Scan() (Token, bool) { _ = "STUB: not implemented"; return *new(Token), false }

// move to first rune

// the first switch cases all position the scanner on the next rune
// by their calls to scan*

func (s *Scanner) scanIdentifier() string { _ = "STUB: not implemented"; return "" }

func (s *Scanner) scanComment() (tid, string) { _ = "STUB: not implemented"; return *new(tid), "" }

// initial '/' already consumed

func (s *Scanner) scanCode() string { _ = "STUB: not implemented"; return "" }

func (s *Scanner) scanEscape(quote rune) bool {
	_ = "STUB: not implemented"
	// scanEscape is always called as part of a greater token, so do not
	// reset s.tok, and write s.cur before calling s.read.
	return false
}

// unicode character class, only valid if quote is ']'

// unicode class name, read until '}'

// single letter class

func (s *Scanner) scanClass() string { _ = "STUB: not implemented"; return "" }

// opening '['

// \n not consumed

// can have an optional "i" ignore case suffix

func (s *Scanner) scanRawString() string { _ = "STUB: not implemented"; return "" }

// opening '`'

// can have an optional "i" ignore case suffix

func stripCR(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func (s *Scanner) scanString() string { _ = "STUB: not implemented"; return "" }

// opening '"'

// \n not consumed

// can have an optional "i" ignore case suffix

func (s *Scanner) scanChar() string { _ = "STUB: not implemented"; return "" }

// opening "'"

// must be followed by one char (which may be an escape) and a single
// quote, but read until we find that closing quote.

// \n not consumed

// can have an optional "i" ignore case suffix

func (s *Scanner) scanRuleDef() string { _ = "STUB: not implemented"; return "" }

// read advances the Scanner to the next rune.
func (s *Scanner) read() { _ = "STUB: not implemented"; return }

// newline is '\n' as in Go

// whitespace is the same as Go, except that it doesn't skip newlines,
// those are returned as tokens.
func (s *Scanner) skipWhitespace() { _ = "STUB: not implemented"; return }

func isRuleDefStart(r rune) bool { _ = "STUB: not implemented"; return false }

/* leftwards arrow */
/* long leftwards arrow */

// isLetter has the same definition as Go.
func isLetter(r rune) bool { _ = "STUB: not implemented"; return false }

// isDigit has the same definition as Go.
func isDigit(r rune) bool { _ = "STUB: not implemented"; return false }

func digitVal(r rune) int { _ = "STUB: not implemented"; return 0 }

// notify the handler of an error.
func (s *Scanner) error(p ast.Pos, err error) { _ = "STUB: not implemented"; return }

// helper to generate and notify of an error.
func (s *Scanner) errorf(f string, args ...any) { _ = "STUB: not implemented"; return }

// helper to generate and notify of an error at a specific position.
func (s *Scanner) errorpf(p ast.Pos, f string, args ...any) { _ = "STUB: not implemented"; return }

// notify a non-recoverable error that terminates the scanning.
func (s *Scanner) fatalError(err error) { _ = "STUB: not implemented"; return }

// convert the reader to a rune reader if required.
func runeReader(r io.Reader) io.RuneReader { _ = "STUB: not implemented"; return *new(io.RuneReader) }
