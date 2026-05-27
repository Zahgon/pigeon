package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	in := os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer func() { _ = f.Close() }()
		in = f
	}
	pn, err := ParseReader("", in)
	if err != nil {
		log.Fatal(err)
	}
	ret, err := pn.(ProgramNode).exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}

var lvalues = make(map[string]int)

// Statement is the smallest standalone element
type Statement interface {
	exec() error
}

// ProgramNode is a root node
type ProgramNode struct {
	statements StatementsNode
	ret        ReturnNode
}

func newProgramNode(stmts StatementsNode, ret ReturnNode) (ProgramNode, error) {
	_ = "STUB: not implemented"
	return *new(ProgramNode), nil
}

func (n ProgramNode) exec() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// StatementsNode is a list of statement
type StatementsNode struct {
	statements []Statement
}

func newStatementsNode(stmts any) (StatementsNode, error) {
	_ = "STUB: not implemented"
	return *new(StatementsNode), nil
}

func (n StatementsNode) exec() error { _ = "STUB: not implemented"; return nil }

// ReturnNode return value to the caller.
type ReturnNode struct {
	arg IdentifierNode
}

func newReturnNode(arg IdentifierNode) (ReturnNode, error) {
	_ = "STUB: not implemented"
	return *new(ReturnNode), nil
}

func (n ReturnNode) exec() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// IfNode controls conditional branching.
type IfNode struct {
	arg        LogicalExpressionNode
	statements StatementsNode
}

func newIfNode(arg LogicalExpressionNode, stmts StatementsNode) (IfNode, error) {
	_ = "STUB: not implemented"
	return *new(IfNode), nil
}

func (n IfNode) exec() error { _ = "STUB: not implemented"; return nil }

// AssignmentNode gives a value to a variable
type AssignmentNode struct {
	lvalue string
	rvalue AdditiveExpressionNode
}

func newAssignmentNode(lvalue IdentifierNode, rvalue AdditiveExpressionNode) (AssignmentNode, error) {
	_ = "STUB: not implemented"
	return *new(AssignmentNode), nil
}

func (n AssignmentNode) exec() error { _ = "STUB: not implemented"; return nil }

// LogicalExpressionNode is a logical expression
type LogicalExpressionNode struct {
	expr PrimaryExpressionNode
}

func newLogicalExpressionNode(expr PrimaryExpressionNode) (LogicalExpressionNode, error) {
	_ = "STUB: not implemented"
	return *new(LogicalExpressionNode), nil
}

func (n LogicalExpressionNode) exec() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// AdditiveExpressionNode is a additive expression
type AdditiveExpressionNode struct {
	arg1 any
	arg2 PrimaryExpressionNode
	op   string
}

func newAdditiveExpressionNode(arg PrimaryExpressionNode, rest any) (AdditiveExpressionNode, error) {
	_ = "STUB: not implemented"
	return *new(AdditiveExpressionNode), nil
}

func (n AdditiveExpressionNode) exec() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// PrimaryExpressionNode is a basic element
type PrimaryExpressionNode struct {
	arg any
}

func newPrimaryExpressionNode(arg any) (PrimaryExpressionNode, error) {
	_ = "STUB: not implemented"
	return *new(PrimaryExpressionNode), nil
}

func (n PrimaryExpressionNode) exec() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// IntegerNode is a integer number
type IntegerNode struct {
	val int
}

func newIntegerNode(val string) (IntegerNode, error) {
	_ = "STUB: not implemented"
	return *new(IntegerNode), nil
}

func (n IntegerNode) exec() (int, error) {
	_ = "STUB: not implemented"

	// IdentifierNode is a reference to variable
	return 0, nil
}

type IdentifierNode struct {
	val string
}

func newIdentifierNode(val string) (IdentifierNode, error) {
	_ = "STUB: not implemented"
	return *new(IdentifierNode), nil
}

func (n IdentifierNode) exec() (int, error) { _ = "STUB: not implemented"; return 0, nil }
