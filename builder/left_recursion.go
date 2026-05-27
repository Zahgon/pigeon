package builder

import (
	"errors"

	"github.com/mna/pigeon/ast"
)

var (
	// ErrNoLeader is no leader error.
	ErrNoLeader = errors.New(
		"SCC has no leadership candidate (no element is included in all cycles)")
	// ErrHaveLeftRecursion is recursion error.
	ErrHaveLeftRecursion = errors.New("grammar contains left recursion")
)

// PrepareGrammar evaluates parameters associated with left recursion.
func PrepareGrammar(grammar *ast.Grammar) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ComputeNullables evaluates nullable nodes.
func ComputeNullables(rules map[string]*ast.Rule) {
	_ = "STUB: not implemented"
	// Compute which rules in a grammar are nullable
	return
}

func findLeader(
	graph map[string]map[string]struct{}, scc map[string]struct{},
) (string, error) {
	_ = "STUB: not implemented"
	// Try to find a leader such that all cycles go through it.
	return "", nil
}

// Pick an arbitrary leader from the candidates.

// The only element.

// ComputeLeftRecursives evaluates left recursion.
func ComputeLeftRecursives(rules map[string]*ast.Rule) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// The only element.

// MakeFirstGraph compute the graph of left-invocations.
// There's an edge from A to B if A may invoke B at its initial position.
// Note that this requires the nullable flags to have been computed.
func MakeFirstGraph(rules map[string]*ast.Rule) map[string]map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}
