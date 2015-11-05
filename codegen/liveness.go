package codegen

import "golang.org/x/tools/go/ssa"

func alive(ident *identifier, loc ssa.Instruction) bool {
	return aliveTest(ident, loc, false)
}

func aliveAfter(ident *identifier, loc ssa.Instruction) bool {
	return aliveTest(ident, loc, true)
}

func aliveTest(ident *identifier, loc ssa.Instruction, after bool) bool {
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
		// no ssa.Value, assume alive
		return true
	}
	start := false
	for _, i := range basicBlock.Instrs {
		if !after {
			if i == loc {
				start = true
			}
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
		if after {
			if i == loc {
				start = true
			}
		}
	}
	// no instruction at or after loc uses ident as an operand
	return false
}

func getBlocks(ident *identifier) []*ssa.BasicBlock {
	f := ident.f
	var blocks []*ssa.BasicBlock
	for _, b := range f.ssa.Blocks {
		if inBlock(ident, b) {
			blocks = append(blocks, b)
		}
	}
	return blocks
}

func inBlock(ident *identifier, b *ssa.BasicBlock) bool {
	for _, i := range b.Instrs {
		if ident.isRetIdent() {
			if _, ok := i.(*ssa.Return); ok {
				return true
			}
		}
		ops := i.Operands(nil)
		for _, op := range ops {
			if op != nil {
				if (*op).Name() == ident.name {
					return true
				}
			}
		}
	}
	return false
}
