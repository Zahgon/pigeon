package builder

import (
	"errors"
)

// ErrInvalidParameters is parameters error.
var ErrInvalidParameters = errors.New("invalid parameters passed to function")

func min(a1 int, a2 int) int { _ = "STUB: not implemented"; return 0 }

// StronglyConnectedComponents compute strongly сonnected сomponents of a graph.
// Tarjan's strongly connected components algorithm.
func StronglyConnectedComponents(
	vertices []string, edges map[string]map[string]struct{},
) []map[string]struct{} {
	_ = "STUB: not implemented"
	// Tarjan's strongly connected components algorithm
	return nil
}

func contains(s []string, e string) bool { _ = "STUB: not implemented"; return false }

func reduceGraph(
	graph map[string]map[string]struct{}, scc map[string]struct{},
) map[string]map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// FindCyclesInSCC find cycles in SCC emanating from start.
// Yields lists of the form ['A', 'B', 'C', 'A'], which means there's
// a path from A -> B -> C -> A.  The first item is always the start
// argument, but the last item may be another element, e.g.  ['A',
// 'B', 'C', 'B'] means there's a path from A to B and there's a
// cycle from B to C and back.
func FindCyclesInSCC(
	graph map[string]map[string]struct{}, scc map[string]struct{}, start string,
) ([][]string, error) {
	_ = "STUB: not implemented"
	// Basic input checks.
	return nil, nil
}

// Reduce the graph to nodes in the SCC.

// Recursive helper that yields cycles.

// TODO: Make this not quadratic.
