package codegen

import (
	"fmt"

	"go/types"

	"golang.org/x/tools/go/ssa"
)

type identifier struct {
	f       *Function
	name    string
	typ     types.Type
	offset  int
	storage storer
	ptr     *identifier
	// offset is from the stack pointer (SP)
	local *ssa.Alloc
	// offset is from the frame pointer (FP)
	param    *ssa.Parameter
	cnst     *ssa.Const
	value    ssa.Value
	spilling bool
}

type context struct {
	f   *Function
	loc ssa.Instruction
}

func (ident *identifier) String() string {
	local := "nil"
	if ident.local != nil {
		local = ident.local.String()
	}
	param := "nil"
	if ident.param != nil {
		param = ident.param.String()
	}
	cnst := "nil"
	if ident.cnst != nil {
		cnst = ident.cnst.String()
	}
	return fmt.Sprintf("identifier{name: %v, typ: %v, local: %v, param: %v, cnst: %v, offset: %v}",
		ident.name, ident.typ, local, param, cnst, ident.offset)
}

func (ident *identifier) size() uint {
	return sizeof(ident.typ)
}

func (ident *identifier) align() uint {
	return align(ident.typ)
}

// Addr returns the register and offset to access the backing memory of ident. It also
// returns the size of ident in bytes.
// For locals the register is the stack pointer (SP) and for params the register
// is the frame pointer (FP).
func (ident *identifier) Addr() (reg register, offset int, size uint) {
	offset = ident.offset
	size = ident.size()
	if ident.isParam() || ident.isRetIdent() {
		reg = *getRegister(REG_FP)
	} else {
		reg = *getRegister(REG_SP)
	}
	return
}

func (ident *identifier) storageRegion() region {
	return region{offset: 0, size: ident.size()}
}

func (ident *identifier) initStorage(valid bool) {
	if ident.cnst != nil {
		cnst := &constant{
			aliaser: &aliaser{},
			parent:  ident,
			Const:   ident.cnst}
		cnst.aliaser.src = cnst
		ident.storage = cnst
		return
	} else {
		mem := &memory{
			parent:             ident,
			initializedRegions: nil,
			aliaser:            &aliaser{},
		}
		// parameter storage is already initialized
		if ident.param != nil {
			mem.initializedRegions = append(mem.initializedRegions, ident.storageRegion())
		}
		mem.aliaser.src = mem
		ident.storage = mem
		return
	}
}

func (ident *identifier) load(ctx context) (string, *register) {
	return ident.loadChunk(ctx, 0, ident.size())
}

func (ident *identifier) loadChunk(ctx context, offset uint, size uint) (string, *register) {
	return ident.storage.load(ctx, region{offset, size})
}

func (ident *identifier) loadConst(ctx context) (string, *register) {
	rgn := region{offset: 0, size: ident.storage.size()}
	return ident.storage.load(ctx, rgn)
}

func (ident *identifier) newValue(ctx context, reg *register, offset uint, size uint) string {
	// constants can't be modified
	if ident.cnst != nil {
		ice("cannot modify constant")
	}

	asm := ""
	chunk := region{offset, size}
	if ident.f.Trace {
		parentName := "nil"
		if reg.parent != nil {
			parentName = reg.parent.owner().name
		}
		fmt.Printf(ident.f.Indent+"New value %v (offset=%v, size=%v) -> %v (d=%v, p=%v)\n",
			ident.name, offset, size, reg.name, reg.dirty, parentName)
	}
	if ident.spilling {
		asm = ident.storage.storeAndSpill(ctx, reg, chunk)
	} else {
		asm = ident.storage.store(ctx, reg, chunk)
	}
	if ident.f.Trace {
		parentName := "nil"
		if reg.parent != nil {
			parentName = reg.parent.owner().name
		}
		fmt.Printf(ident.f.Indent+"New value {%v (offset=%v, size=%v) -> %v (d=%v, p=%v)}\n",
			ident.name, offset, size, reg.name, reg.dirty, parentName)
	}
	return asm
}

func (ident *identifier) storeAndSpill(reg *register, offset uint, size uint) string {
	// constants can't be modified
	if ident.cnst != nil {
		ice("cannot modify constant")
	}

	asm := ""
	chunk := region{offset, size}
	ctx := context{ident.f, nil}
	if ident.f.Trace {
		parentName := "nil"
		if reg.parent != nil {
			parentName = reg.parent.owner().name
		}
		fmt.Printf(ident.f.Indent+"New value %v (offset=%v, size=%v) -> %v (d=%v, p=%v)\n",
			ident.name, offset, size, reg.name, reg.dirty, parentName)
	}
	asm = ident.storage.storeAndSpill(ctx, reg, chunk)
	if ident.f.Trace {
		parentName := "nil"
		if reg.parent != nil {
			parentName = reg.parent.owner().name
		}
		fmt.Printf(ident.f.Indent+"New value {%v (offset=%v, size=%v) -> %v (d=%v, p=%v)}\n",
			ident.name, offset, size, reg.name, reg.dirty, parentName)
	}
	return asm
}

func (ident *identifier) spillAllRegisters(loc ssa.Instruction) string {
	return ident.spillRegisters(loc, true)
}

func (ident *identifier) spillDirtyRegisters(loc ssa.Instruction) string {
	return ident.spillRegisters(loc, false)
}

func (ident *identifier) spillRegisters(loc ssa.Instruction, all bool) string {
	ctx := context{ident.f, loc}
	return ident.storage.spillRegisters(ctx, all)
}

func (ident *identifier) spillRegister(r *register, force bool) string {
	ctx := context{ident.f, nil}
	return ident.storage.spillRegister(ctx, r, force)
}

func (ident *identifier) ssaInstr() ssa.Instruction {
	if ident.local != nil {
		return ident.local
	}
	if instr, ok := ident.ssaValue().(ssa.Instruction); ok {
		return instr
	}
	return nil
}

func (ident *identifier) ssaValue() ssa.Value {
	if ident.cnst != nil {
		return ident.cnst
	}
	if ident.local != nil {
		return ident.local
	}
	if ident.param != nil {
		return ident.param
	}
	return ident.value
}

func (ident *identifier) isPhi() bool {
	_, ok := ident.ssaValue().(*ssa.Phi)
	return ok
}

func (ident *identifier) isSsaLocal() bool {
	return ident.local != nil
}

func (ident *identifier) isParam() bool {
	return ident.param != nil
}

func (ident *identifier) isConst() bool {
	return ident.cnst != nil
}

func (ident *identifier) isRetIdent() bool {
	return ident.name == retName()
}

func (ident *identifier) isBlockLocal() bool {
	if ident.isSsaLocal() ||
		ident.isParam() ||
		ident.isPhi() ||
		ident.isRetIdent() {

		return false
	} else {
		// expensive computation
		return len(getBlocks(ident)) <= 1
	}
}

func (ident *identifier) isPointer() bool {
	_, ok := ident.typ.(*types.Pointer)
	return ok
}

func (ident *identifier) ptrUnderlyingType() types.Type {
	if !ident.isPointer() {
		ice(fmt.Sprintf("identifier (%v) not ptr type", ident))
	}
	ptrType := ident.typ.(*types.Pointer)
	return ptrType.Elem()
}

func (name *identifier) isInteger() bool {
	if !isBasic(name.typ) {
		return false
	}
	t := name.typ.(*types.Basic)
	return t.Info()&types.IsInteger == types.IsInteger
}
