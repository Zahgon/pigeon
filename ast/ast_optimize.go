package ast

type grammarOptimizer struct {
	rule            string
	protectedRules  map[string]struct{}
	rules           map[string]*Rule
	ruleUsesRules   map[string]map[string]struct{}
	ruleUsedByRules map[string]map[string]struct{}
	visitor         func(expr Expression) Visitor
	optimized       bool
}

func newGrammarOptimizer(protectedRules []string) *grammarOptimizer {
	_ = "STUB: not implemented"
	return nil
}

// Visit is a generic Visitor to be used with Walk
// The actual function, which should be used during Walk
// is held in ruleRefOptimizer.visitor.
func (r *grammarOptimizer) Visit(expr Expression) Visitor {
	_ = "STUB: not implemented"
	return *

	// init is a Visitor, which is used with the Walk function
	// The purpose of this function is to initialize the reference
	// maps rules, ruleUsesRules and ruleUsedByRules.
	new(Visitor)
}

func (r *grammarOptimizer) init(expr Expression) Visitor {
	switch expr := expr.(type) {
	case *Rule:
		// Keep track of current rule, which is processed
		r.rule = expr.Name.Val
		r.rules[expr.Name.Val] = expr
	case *RuleRefExpr:
		// Fill ruleUsesRules and ruleUsedByRules for every RuleRefExpr
		set(r.ruleUsesRules, r.rule, expr.Name.Val)
		set(r.ruleUsedByRules, expr.Name.Val, r.rule)
	}
	return r
}

// Add element to map of maps, initialize the inner map
// if necessary.
func set(m map[string]map[string]struct{}, src, dst string) { _ = "STUB: not implemented"; return }

// optimize is a Visitor, which is used with the Walk function
// The purpose of this function is to perform the actual optimizations.
// See Optimize for a detailed list of the performed optimizations.
func (r *grammarOptimizer) optimize(expr0 Expression) Visitor {
	_ = "STUB: not implemented"
	return *new(Visitor)
}

// Optimize choice nested in choice

// Combine sequence of single char LitMatcher to CharClassMatcher

// Combine two LitMatcher to CharClassMatcher
// "a" / "b" => [ab]

// Combine LitMatcher with CharClassMatcher
// "a" / [bc] => [abc]

// Combine CharClassMatcher with LitMatcher
// [ab] / "c" => [abc]

// Combine CharClassMatcher with CharClassMatcher
// [ab] / [cd] => [abcd]

// If one of the optimizations was applied, remove the second element from Alternatives

// Reset optimized at the start of each Walk.

// Remove Rule, if it is no longer used by any other Rule and it is not the first Rule.

// Compensate for the removed item

// Optimize nested sequences

// Combine sequence of LitMatcher

func (r *grammarOptimizer) optimizeRules(exprs []Expression) []Expression {
	_ = "STUB: not implemented"
	return nil
}

func (r *grammarOptimizer) optimizeRule(expr Expression) Expression {
	_ = "STUB: not implemented"
	// Optimize RuleRefExpr
	return *new(Expression)
}

// TODO: Check if reference exists, otherwise raise an error, which reference is missing!

// Remove Choices with only one Alternative left

// Remove Sequence with only one Expression

// cloneExpr takes an Expression and deep clones it (including all children)
// This is necessary because referenced Rules are denormalized and therefore
// have to become independent from their original Expression.
func cloneExpr(expr Expression) Expression { _ = "STUB: not implemented"; return *new(Expression) }

// cleanupCharClassMatcher is a Visitor, which is used with the Walk function
// The purpose of this function is to cleanup the redundancies created by the
// optimize Visitor. This includes to remove redundant entries in Chars, Ranges
// and UnicodeClasses of the given CharClassMatcher as well as regenerating the
// correct content for the Val field (string representation of the CharClassMatcher).
func (r *grammarOptimizer) cleanupCharClassMatcher(expr0 Expression) Visitor {
	_ = "STUB: not implemented"
	// We are only interested in nodes of type *CharClassMatcher
	return *new(Visitor)
}

// Remove redundancies in Chars

// Remove redundancies in Ranges

// Remove redundancies in UnicodeClasses

// Regenerate the content for Val

func escapeRune(r rune) string { _ = "STUB: not implemented"; return "" }

// Optimize walks a given grammar and optimizes the grammar in regards
// of parsing performance. This is done with several optimizations:
//   - removal of unreferenced rules
//   - replace rule references with a copy of the referenced Rule, if the
//     referenced rule it self has no references.
//   - resolve nested choice expressions
//   - resolve choice expressions with only one alternative
//   - resolve nested sequences expression
//   - resolve sequence expressions with only one element
//   - combine character class matcher and literal matcher, where possible
func Optimize(g *Grammar, alternateEntrypoints ...string) { _ = "STUB: not implemented"; return }
