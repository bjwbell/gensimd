package codegen

import (
	"fmt"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types"
)

type identifier struct {
	f       *Function
	name    string
	typ     types.Type
	local   *varInfo
	param   *paramInfo
	cnst    *ssa.Const
	offset  int
	storage storer
	ptr     *identifier
}

type context struct {
	f   *Function
	loc ssa.Instruction
}

func (ident *identifier) String() string {
	return fmt.Sprintf("identifier{name: %v, typ: %v, local: %v, param: %v, cnst: %v, offset: %v}",
		ident.name, ident.typ, ident.local, ident.param, ident.cnst, ident.offset)
}

func (ident *identifier) size() uint {
	return sizeof(ident.typ)
}

func (ident *identifier) align() uint {
	return align(ident.typ)
}

// Addr returns the register and offset to access the backing memory of name. It also
// returns the size of name in bytes.
// For locals the register is the stack pointer (SP) and for params the register
// is the frame pointer (FP).
func (name *identifier) Addr() (reg register, offset int, size uint) {
	offset = name.offset
	size = name.size()
	if name.local != nil {
		reg = *getRegister(REG_SP)
	} else if name.param != nil {
		reg = *getRegister(REG_FP)
	} else {
		ice(fmt.Sprintf("identifier (%v) is not a local or param", name))
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

func (ident *identifier) newValue(reg *register, offset uint, size uint) string {
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
		fmt.Printf(ident.f.Indent+"New value %v (offset=%v, size=%v) -> %v (d=%v,p=%v)\n",
			ident.name, offset, size, reg.name, reg.dirty, parentName)
	}

	if reg.parent != ident.storage {
		asm = ident.storage.store(ctx, reg, chunk)
	} else {
		// nothing to do
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

func (ident *identifier) spillAllRegisters() string {
	return ident.spillRegisters(true)
}

func (ident *identifier) spillDirtyRegisters() string {
	return ident.spillRegisters(false)
}

func (ident *identifier) spillRegisters(all bool) string {
	ctx := context{ident.f, nil}
	return ident.storage.spillRegisters(ctx, all)
}

func (ident *identifier) spillRegister(r *register, force bool) string {
	ctx := context{ident.f, nil}
	return ident.storage.spillRegister(ctx, r, force)
}

func (name *identifier) IsSsaLocal() bool {
	return name.local != nil && name.local.info != nil
}

func (name *identifier) IsPointer() bool {
	_, ok := name.typ.(*types.Pointer)
	return ok
}

func (name *identifier) PointerUnderlyingType() types.Type {
	if !name.IsPointer() {
		ice(fmt.Sprintf("identifier (%v) not ptr type", name))
	}
	ptrType := name.typ.(*types.Pointer)
	return ptrType.Elem()
}

func (name *identifier) IsInteger() bool {
	if !isBasic(name.typ) {
		return false
	}
	t := name.typ.(*types.Basic)
	return t.Info()&types.IsInteger == types.IsInteger
}

type varInfo struct {
	name string
	// offset is from the stack pointer (SP)
	info *ssa.Alloc
}

func (v *varInfo) ssaName() string {
	return v.info.Name()
}

type paramInfo struct {
	name string
	// offset is from the frame pointer (FP)
	info  *ssa.Parameter
	extra interface{}
}

func (p *paramInfo) ssaName() string {
	return p.info.Name()
}
