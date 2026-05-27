package asmgoto

// Instruction is an assembler instruction
type Instruction interface {
	// Assemble returns an assembler instruction as string representation
	Assemble() string
}

// Noop is the no-operation assembler instruction
type Noop struct{}

// Assemble returns an assembler instruction as string representation
func (n Noop) Assemble() string {
	_ = "STUB: not implemented"

	// Jump is the jump assembler instruction
	return ""
}

type Jump struct {
	RelAddr int
}

// Assemble returns an assembler instruction as string representation
func (j Jump) Assemble() string { _ = "STUB: not implemented"; return "" }

type labelLookup map[string]label

type label struct {
	line  int
	jumps []unresolvedJump
}

type unresolvedJump struct {
	line int
	jump *Jump
}

func addLabel(c *current, name string) { _ = "STUB: not implemented"; return }

// Label not seen yet, add to labelLookup

// Label already seen

// Update position for later usage of Label

// Update all already known jumps to this Label with the correct relative jump distance

func addJump(c *current, name string) *Jump { _ = "STUB: not implemented"; return nil }

// Label not seen yet, create Label with invalid line = -1, add Jump to unresolvedJump

// Label already seen as target of an other Jump, add Jump to unresolvedJump

// Label already seen, calculate correct relative jump distance

func labelCheck(c *current) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Iterate through all Label, there must be no unresolved jumps
