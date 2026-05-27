package bootstrap

import (
	"io"

	"github.com/mna/pigeon/ast"
)

type errList []error

func (e *errList) reset() { _ = "STUB: not implemented"; return }

func (e *errList) add(p ast.Pos, err error) { _ = "STUB: not implemented"; return }

func (e *errList) err() error { _ = "STUB: not implemented"; return nil }

func (e *errList) Error() string { _ = "STUB: not implemented"; return "" }

// Parser holds the state to parse the PEG grammar into
// an abstract syntax tree (AST).
type Parser struct {
	s   Scanner
	tok Token

	errs *errList
	dbg  bool
	pk   Token
}

func (p *Parser) in(s string) string { _ = "STUB: not implemented"; return "" }

func (p *Parser) out(s string) { _ = "STUB: not implemented"; return }

// NewParser creates a new Parser.
func NewParser() *Parser { _ = "STUB: not implemented"; return nil }

// Parse parses the data from the reader r and generates the AST
// or returns an error if it fails. The filename is used as information
// in the error messages.
func (p *Parser) Parse(filename string, r io.Reader) (*ast.Grammar, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) read() { _ = "STUB: not implemented"; return }

func (p *Parser) peek() Token { _ = "STUB: not implemented"; return *new(Token) }

func (p *Parser) skip(ids ...tid) { _ = "STUB: not implemented"; return }

func (p *Parser) grammar() *ast.Grammar { _ = "STUB: not implemented"; return nil }

// advance to the first token

func (p *Parser) expect(ids ...tid) bool { _ = "STUB: not implemented"; return false }

func (p *Parser) rule() *ast.Rule { _ = "STUB: not implemented"; return nil }

func (p *Parser) expression() ast.Expression {
	_ = "STUB: not implemented"
	return *new(ast.Expression)
}

// move after the slash

func (p *Parser) actionExpr() ast.Expression {
	_ = "STUB: not implemented"
	return *new(ast.Expression)
}

func (p *Parser) seqExpr() ast.Expression { _ = "STUB: not implemented"; return *new(ast.Expression) }

func (p *Parser) labeledExpr() ast.Expression {
	_ = "STUB: not implemented"
	return *new(ast.Expression)
}

func (p *Parser) prefixedExpr() ast.Expression {
	_ = "STUB: not implemented"
	return *new(ast.Expression)
}

func (p *Parser) suffixedExpr() ast.Expression {
	_ = "STUB: not implemented"
	return *new(ast.Expression)
}

func (p *Parser) primaryExpr() ast.Expression {
	_ = "STUB: not implemented"
	return *new(ast.Expression)
}

// literal matcher

// character class matcher

// any matcher

// rule reference expression

// expression in parenthesis

// if p.tok.id != eof && p.tok.id != eol && p.tok.id != semicolon {
// 	p.errs.add(p.tok.pos, fmt.Errorf("invalid token %s (%q) for primary expression", p.tok.id, p.tok.lit))
// }

func (p *Parser) ruleRefExpr() ast.Expression {
	_ = "STUB: not implemented"
	return *new(ast.Expression)
}
