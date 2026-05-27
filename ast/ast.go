// Package ast defines the abstract syntax tree for the PEG grammar.
//
// The parser generator's PEG grammar generates a tree using this package
// that is then converted by the builder to the simplified AST used in
// the generated parser.
package ast

// Pos represents a position in a source file.
type Pos struct {
	Filename string
	Line     int
	Col      int
	Off      int
}

// String returns the textual representation of a position.
func (p Pos) String() string { _ = "STUB: not implemented"; return "" }

// Grammar is the top-level node of the AST for the PEG grammar.
type Grammar struct {
	p     Pos
	Init  *CodeBlock
	Rules []*Rule
}

var _ Expression = (*Grammar)(nil)

// NewGrammar creates a new grammar at the specified position.
func NewGrammar(p Pos) *Grammar { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (g *Grammar) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (g *Grammar) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (g *Grammar) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (g *Grammar) IsNullable() bool { _ = "STUB: not implemented"; return false }

// InitialNames returns names of nodes with which an expression can begin.
func (g *Grammar) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// Rule represents a rule in the PEG grammar. It has a name, an optional
// display name to be used in error messages, and an expression.
type Rule struct {
	p           Pos
	Name        *Identifier
	DisplayName *StringLit
	Expr        Expression

	// Fields below to work with left recursion.
	Visited       bool
	Nullable      bool
	LeftRecursive bool
	Leader        bool
}

var _ Expression = (*Rule)(nil)

// NewRule creates a rule with at the specified position and with the
// specified name as identifier.
func NewRule(p Pos, name *Identifier) *Rule { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (r *Rule) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (r *Rule) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (r *Rule) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// A left-recursive rule is considered non-nullable.
	return false
}

// IsNullable returns the nullable attribute of the node.
func (r *Rule) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (r *Rule) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// Expression is the interface implemented by all expression types.
type Expression interface {
	Pos() Pos

	// for work with left recursion
	NullableVisit(rules map[string]*Rule) bool
	IsNullable() bool
	InitialNames() map[string]struct{}
}

// ChoiceExpr is an ordered sequence of expressions. The parser tries to
// match any of the alternatives in sequence and stops at the first one
// that matches.
type ChoiceExpr struct {
	p            Pos
	Alternatives []Expression

	Nullable bool
}

var _ Expression = (*ChoiceExpr)(nil)

// NewChoiceExpr creates a choice expression at the specified position.
func NewChoiceExpr(p Pos) *ChoiceExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (c *ChoiceExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (c *ChoiceExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (c *ChoiceExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (c *ChoiceExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (c *ChoiceExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// FailureLabel is an identifier, which can by thrown and recovered in a grammar.
type FailureLabel string

// RecoveryExpr is an ordered sequence of expressions. The parser tries to
// match any of the alternatives in sequence and stops at the first one
// that matches.
type RecoveryExpr struct {
	p           Pos
	Expr        Expression
	RecoverExpr Expression
	Labels      []FailureLabel

	Nullable bool
}

var _ Expression = (*RecoveryExpr)(nil)

// NewRecoveryExpr creates a choice expression at the specified position.
func NewRecoveryExpr(p Pos) *RecoveryExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (r *RecoveryExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (r *RecoveryExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (r *RecoveryExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (r *RecoveryExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (r *RecoveryExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// ActionExpr is an expression that has an associated block of code to
// execute when the expression matches.
type ActionExpr struct {
	p      Pos
	Expr   Expression
	Code   *CodeBlock
	FuncIx int

	Nullable bool
}

var _ Expression = (*ActionExpr)(nil)

// NewActionExpr creates a new action expression at the specified position.
func NewActionExpr(p Pos) *ActionExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (a *ActionExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (a *ActionExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (a *ActionExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (a *ActionExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (a *ActionExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// ThrowExpr is an expression that throws an FailureLabel to be caught by a
// RecoveryChoiceExpr.
type ThrowExpr struct {
	p     Pos
	Label string
}

var _ Expression = (*ThrowExpr)(nil)

// NewThrowExpr creates a new throw expression at the specified position.
func NewThrowExpr(p Pos) *ThrowExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (t *ThrowExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (t *ThrowExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (t *ThrowExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (t *ThrowExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (t *ThrowExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// SeqExpr is an ordered sequence of expressions, all of which must match
// if the SeqExpr is to be a match itself.
type SeqExpr struct {
	p     Pos
	Exprs []Expression

	Nullable bool
}

var _ Expression = (*SeqExpr)(nil)

// NewSeqExpr creates a new sequence expression at the specified position.
func NewSeqExpr(p Pos) *SeqExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (s *SeqExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (s *SeqExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (s *SeqExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (s *SeqExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (s *SeqExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// LabeledExpr is an expression that has an associated label. Code blocks
// can access the value of the expression using that label, that becomes
// a local variable in the code.
type LabeledExpr struct {
	p     Pos
	Label *Identifier
	Expr  Expression
}

var _ Expression = (*LabeledExpr)(nil)

// NewLabeledExpr creates a new labeled expression at the specified position.
func NewLabeledExpr(p Pos) *LabeledExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (l *LabeledExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (l *LabeledExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (l *LabeledExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (l *LabeledExpr) IsNullable() bool { _ = "STUB: not implemented"; return false }

// InitialNames returns names of nodes with which an expression can begin.
func (l *LabeledExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// AndExpr is a zero-length matcher that is considered a match if the
// expression it contains is a match.
type AndExpr struct {
	p    Pos
	Expr Expression
}

// NewAndExpr creates a new and (&) expression at the specified position.
func NewAndExpr(p Pos) *AndExpr { _ = "STUB: not implemented"; return nil }

var _ Expression = (*AndExpr)(nil)

// Pos returns the starting position of the node.
func (a *AndExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (a *AndExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (a *AndExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (a *AndExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (a *AndExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// NotExpr is a zero-length matcher that is considered a match if the
// expression it contains is not a match.
type NotExpr struct {
	p    Pos
	Expr Expression
}

var _ Expression = (*NotExpr)(nil)

// NewNotExpr creates a new not (!) expression at the specified position.
func NewNotExpr(p Pos) *NotExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (n *NotExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (n *NotExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (n *NotExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (n *NotExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (n *NotExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// ZeroOrOneExpr is an expression that can be matched zero or one time.
type ZeroOrOneExpr struct {
	p    Pos
	Expr Expression
}

var _ Expression = (*ZeroOrOneExpr)(nil)

// NewZeroOrOneExpr creates a new zero or one expression at the specified
// position.
func NewZeroOrOneExpr(p Pos) *ZeroOrOneExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (z *ZeroOrOneExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (z *ZeroOrOneExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (z *ZeroOrOneExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (z *ZeroOrOneExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (z *ZeroOrOneExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// ZeroOrMoreExpr is an expression that can be matched zero or more times.
type ZeroOrMoreExpr struct {
	p    Pos
	Expr Expression
}

var _ Expression = (*ZeroOrMoreExpr)(nil)

// NewZeroOrMoreExpr creates a new zero or more expression at the specified
// position.
func NewZeroOrMoreExpr(p Pos) *ZeroOrMoreExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (z *ZeroOrMoreExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (z *ZeroOrMoreExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (z *ZeroOrMoreExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (z *ZeroOrMoreExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (z *ZeroOrMoreExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// OneOrMoreExpr is an expression that can be matched one or more times.
type OneOrMoreExpr struct {
	p    Pos
	Expr Expression
}

var _ Expression = (*OneOrMoreExpr)(nil)

// NewOneOrMoreExpr creates a new one or more expression at the specified
// position.
func NewOneOrMoreExpr(p Pos) *OneOrMoreExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (o *OneOrMoreExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (o *OneOrMoreExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (o *OneOrMoreExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (o *OneOrMoreExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (o *OneOrMoreExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// RuleRefExpr is an expression that references a rule by name.
type RuleRefExpr struct {
	p    Pos
	Name *Identifier

	Nullable bool
}

var _ Expression = (*RuleRefExpr)(nil)

// NewRuleRefExpr creates a new rule reference expression at the specified
// position.
func NewRuleRefExpr(p Pos) *RuleRefExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (r *RuleRefExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (r *RuleRefExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (r *RuleRefExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// Token or unknown; never empty.

// IsNullable returns the nullable attribute of the node.
func (r *RuleRefExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (r *RuleRefExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// StateCodeExpr is an expression which can modify the internal state of the parser.
type StateCodeExpr struct {
	p      Pos
	Code   *CodeBlock
	FuncIx int
}

var _ Expression = (*StateCodeExpr)(nil)

// NewStateCodeExpr creates a new state (#) code expression at the specified
// position.
func NewStateCodeExpr(p Pos) *StateCodeExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (s *StateCodeExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (s *StateCodeExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (s *StateCodeExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (s *StateCodeExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (s *StateCodeExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// AndCodeExpr is a zero-length matcher that is considered a match if the
// code block returns true.
type AndCodeExpr struct {
	p      Pos
	Code   *CodeBlock
	FuncIx int
}

var _ Expression = (*AndCodeExpr)(nil)

// NewAndCodeExpr creates a new and (&) code expression at the specified
// position.
func NewAndCodeExpr(p Pos) *AndCodeExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (a *AndCodeExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (a *AndCodeExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (a *AndCodeExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (a *AndCodeExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (a *AndCodeExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// NotCodeExpr is a zero-length matcher that is considered a match if the
// code block returns false.
type NotCodeExpr struct {
	p      Pos
	Code   *CodeBlock
	FuncIx int
}

var _ Expression = (*NotCodeExpr)(nil)

// NewNotCodeExpr creates a new not (!) code expression at the specified
// position.
func NewNotCodeExpr(p Pos) *NotCodeExpr { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (n *NotCodeExpr) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (n *NotCodeExpr) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (n *NotCodeExpr) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (n *NotCodeExpr) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (n *NotCodeExpr) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// LitMatcher is a string literal matcher. The value to match may be a
// double-quoted string, a single-quoted single character, or a back-tick
// quoted raw string.
type LitMatcher struct {
	posValue   // can be str, rstr or char
	IgnoreCase bool
}

var _ Expression = (*LitMatcher)(nil)

// NewLitMatcher creates a new literal matcher at the specified position and
// with the specified value.
func NewLitMatcher(p Pos, v string) *LitMatcher { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (l *LitMatcher) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (l *LitMatcher) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (l *LitMatcher) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false

	// IsNullable returns the nullable attribute of the node.
}

func (l *LitMatcher) IsNullable() bool {
	_ = "STUB: not implemented"
	// The string token ” is considered empty.
	return false
}

// InitialNames returns names of nodes with which an expression can begin.
func (l *LitMatcher) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// CharClassMatcher is a character class matcher. The value to match must
// be one of the specified characters, in a range of characters, or in the
// Unicode classes of characters.
type CharClassMatcher struct {
	posValue
	IgnoreCase     bool
	Inverted       bool
	Chars          []rune
	Ranges         []rune // pairs of low/high range
	UnicodeClasses []string
}

var _ Expression = (*CharClassMatcher)(nil)

// NewCharClassMatcher creates a new character class matcher at the specified
// position and with the specified raw value. It parses the raw value into
// the list of characters, ranges and Unicode classes.
func NewCharClassMatcher(p Pos, raw string) *CharClassMatcher {
	_ = "STUB: not implemented"
	return nil
}

func (c *CharClassMatcher) parse() { _ = "STUB: not implemented"; return }

// "unquote" the character classes

// content of char class is necessarily valid, so escapes are correct

// extract ranges and chars

// start of range is the last Char added

// Pos returns the starting position of the node.
func (c *CharClassMatcher) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (c *CharClassMatcher) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (c *CharClassMatcher) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false

	// IsNullable returns the nullable attribute of the node.
}

func (c *CharClassMatcher) IsNullable() bool { _ = "STUB: not implemented"; return false }

// InitialNames returns names of nodes with which an expression can begin.
func (c *CharClassMatcher) InitialNames() map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// AnyMatcher is a matcher that matches any character except end-of-file.
type AnyMatcher struct {
	posValue
}

var _ Expression = (*AnyMatcher)(nil)

// NewAnyMatcher creates a new any matcher at the specified position. The
// value is provided for completeness' sake, but it is always the dot.
func NewAnyMatcher(p Pos, v string) *AnyMatcher { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (a *AnyMatcher) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (a *AnyMatcher) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (a *AnyMatcher) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"

	// IsNullable returns the nullable attribute of the node.
	return false
}

func (a *AnyMatcher) IsNullable() bool {
	_ = "STUB: not implemented"

	// InitialNames returns names of nodes with which an expression can begin.
	return false
}

func (a *AnyMatcher) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// CodeBlock represents a code block.
type CodeBlock struct {
	posValue
}

var _ Expression = (*CodeBlock)(nil)

// NewCodeBlock creates a new code block at the specified position and with
// the specified value. The value includes the outer braces.
func NewCodeBlock(p Pos, code string) *CodeBlock { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (c *CodeBlock) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (c *CodeBlock) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (c *CodeBlock) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (c *CodeBlock) IsNullable() bool { _ = "STUB: not implemented"; return false }

// InitialNames returns names of nodes with which an expression can begin.
func (c *CodeBlock) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// Identifier represents an identifier.
type Identifier struct {
	posValue
}

var _ Expression = (*Identifier)(nil)

// NewIdentifier creates a new identifier at the specified position and
// with the specified name.
func NewIdentifier(p Pos, name string) *Identifier { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (i *Identifier) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (i *Identifier) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (i *Identifier) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (i *Identifier) IsNullable() bool { _ = "STUB: not implemented"; return false }

// InitialNames returns names of nodes with which an expression can begin.
func (i *Identifier) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

// StringLit represents a string literal.
type StringLit struct {
	posValue
}

var _ Expression = (*StringLit)(nil)

// NewStringLit creates a new string literal at the specified position and
// with the specified value.
func NewStringLit(p Pos, val string) *StringLit { _ = "STUB: not implemented"; return nil }

// Pos returns the starting position of the node.
func (s *StringLit) Pos() Pos {
	_ = "STUB: not implemented"

	// String returns the textual representation of a node.
	return *new(Pos)
}

func (s *StringLit) String() string { _ = "STUB: not implemented"; return "" }

// NullableVisit recursively determines whether an object is nullable.
func (s *StringLit) NullableVisit(rules map[string]*Rule) bool {
	_ = "STUB: not implemented"
	return false
}

// IsNullable returns the nullable attribute of the node.
func (s *StringLit) IsNullable() bool { _ = "STUB: not implemented"; return false }

// InitialNames returns names of nodes with which an expression can begin.
func (s *StringLit) InitialNames() map[string]struct{} { _ = "STUB: not implemented"; return nil }

type posValue struct {
	p   Pos
	Val string
}
