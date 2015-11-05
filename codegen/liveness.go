package codegen

import "golang.org/x/tools/go/ssa"

func alive(ident *identifier, loc ssa.Instruction) bool {
	// assume all identifiers not local to a basic block are alive
	if !ident.isBlockLocal() {
		return true
	}
	if loc == nil {
		ice("invalid SSA instruction")
	}
	basicBlock := loc.Block()
	if basicBlock == nil {
		ice("cant get basic block for SSA instruction")
	}
	value := ident.ssaValue()
	if value == nil {
		// can't get ssa.Value, assume alive
		return true
	}
	start := false
	for _, i := range basicBlock.Instrs {
		if i == loc {
			start = true
		}
		if start {
			ops := i.Operands(nil)
			for _, op := range ops {
				// instruction at or after loc uses ident as an operand
				if op != nil && *op == value {
					return true
				}
			}
		}
	}
	// no instruction at or after loc uses ident as an operand
	return false
}
