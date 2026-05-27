package ast

// A Visitor implements a Visit method, which is invoked for each Expression
// encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of Expression with the visitor w, followed by a call of w.Visit(nil).
type Visitor interface {
	Visit(expr Expression) (w Visitor)
}

// Walk traverses an AST in depth-first order: It starts by calling
// v.Visit(expr); Expression must not be nil. If the visitor w returned by
// v.Visit(expr) is not nil, Walk is invoked recursively with visitor
// w for each of the non-nil children of Expression, followed by a call of
// w.Visit(nil).
func Walk(v Visitor, expr Expression) { _ = "STUB: not implemented"; return }

// Nothing to do

// Nothing to do

// Nothing to do

// Nothing to do

// Nothing to do

// Nothing to do

// Nothing to do

type inspector func(Expression) bool

func (f inspector) Visit(expr Expression) Visitor { _ = "STUB: not implemented"; return *new(Visitor) }

// Inspect traverses an AST in depth-first order: It starts by calling
// f(expr); expr must not be nil. If f returns true, Inspect invokes f
// recursively for each of the non-nil children of expr, followed by a
// call of f(nil).
func Inspect(expr Expression, f func(Expression) bool) { _ = "STUB: not implemented"; return }
