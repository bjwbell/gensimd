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
	aliases []storage
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

func (ident *identifier) addAlias(newAlias storage) bool {
	if ident.f.Trace {
		fmt.Printf("Alias: %v -> %v\n", ident.name, newAlias.String())

	}
	new := true
	for _, alias := range ident.aliases {
		if alias == newAlias {
			new = false
		}
		if ident.f.Trace {
			fmt.Printf("     : %v -> %v (v=%v)\n", ident.name, alias.String(), alias.isValid())
		}
	}
	if new {
		ident.aliases = append(ident.aliases, newAlias)
	}
	return new
}

func (ident *identifier) removeAlias(oldAlias storage) bool {
	aliases := []storage{}
	removed := false
	if ident.f.Trace {
		if r, ok := oldAlias.(*register); ok {
			fmt.Printf("Alias: %v -/ %v (d=%v)\n", ident.name, oldAlias.String(), r.dirty)
		} else {
			fmt.Printf("Alias: %v -/ %v\n", ident.name, oldAlias.String())
		}
	}
	for _, alias := range ident.aliases {
		if alias == oldAlias {
			removed = true
		} else {
			aliases = append(aliases, alias)
			if ident.f.Trace {
				fmt.Printf("     : %v -> %v (v=%v)\n", ident.name, alias.String(), alias.isValid())
			}
		}
	}
	if !removed {
		for _, alias := range ident.aliases {
			fmt.Println("ALIAS:", alias)

		}

		fmt.Println("IDENT.NAME:", ident.name)
		fmt.Println("REMOVEAL ALIAS:", oldAlias)
		panic("DIDN'T REMOVE ALIAS")
	}
	ident.aliases = aliases
	return removed
}

func compareStorage(x, y storage) int {
	if x.isValid() && !y.isValid() {
		return 1
	}
	if !x.isValid() && y.isValid() {
		return -1
	}
	_, xIsReg := x.(*register)
	_, yIsReg := y.(*register)

	if xIsReg && !yIsReg {
		return 1
	} else if xIsReg && yIsReg {
		return 0
	} else {
		return -1
	}
}

func (ident *identifier) initStorage(valid bool) {
	if ident.getBestStorage(0, ident.size()) != nil {
		return
	}

	if ident.cnst != nil {
		st := &constant{
			ownedby: ident,
			owns:    region{offset: 0, size: ident.size()},
			Const:   ident.cnst}
		ident.addAlias(st)
		return
	}

	r, roffset, rsize := ident.Addr()
	st := &memory{
		stale:   !valid,
		ownedby: ident,
		owns:    region{offset: 0, size: ident.size()},
		reg:     &r,
		offset:  roffset,
		size:    rsize}

	ident.addAlias(st)
}

func (ident *identifier) getBestStorage(offset uint, size uint) storage {
	var st storage

	for _, alias := range ident.aliases {
		if alias.ownerRegion().overlap(region{offset, size}).size == 0 {
			continue
		}
		if st == nil || compareStorage(st, alias) < 0 {
			st = alias
		}
	}
	return st
}

func (ident *identifier) load(ctx context) (string, *register) {
	return ident.loadChunk(ctx, 0, ident.size())
}

func (ident *identifier) loadChunk(ctx context, offset uint, size uint) (string, *register) {
	f := ctx.f
	loc := ctx.loc
	st := ident.getBestStorage(offset, size)

	if ident.cnst != nil && (st == nil || !st.isValid()) {
		return ident.loadConst(ctx)
	}
	asm := ""
	if st == nil || !st.isValid() {
		ice(fmt.Sprintf("nothing backing identifier (%v), offset %v, size %v, best found \"%v, valid %v\"", ident.name, offset, size, st.String(), st.isValid()))
	}

	chunk := transfer{srcOffset: uint(offset), dstOffset: 0, size: size}
	if reg, ok := st.(*register); ok {
		if !reg.isValid() {
			ice("unexpected stale register")
		}
		return "", reg
	}
	if !st.isValid() {
		ice("unexpected stale storage")
	}
	asm, reg := f.allocIdentReg(loc, ident, size)
	if size > 8 &&
		((!isSimd(ident.typ) && !isSSE2(ident.typ)) || reg.typ != XMM_REG) {

		msg := "ident (%v), loading more than 8 byte chunk"
		ice(fmt.Sprintf(msg, ident.name))
	}

	asm += fmt.Sprintf("// BEGIN LoadIdent, ident: %v, offset %v, size %v\n", ident.name, offset, size)
	asm += st.copyChunk(ctx.f, chunk, reg)
	asm += fmt.Sprintf("// END LoadIdent, ident: %v, offset %v, size %v\n", ident.name, offset, size)
	if offset < 0 {
		panic("unexpected, offset < 0")
	}
	reg.newOwner(ident, region{offset: uint(offset), size: size})
	reg.dirty = false
	reg.isValidAlias = true
	ident.addAlias(reg)
	return asm, reg
}

func (ident *identifier) loadConst(ctx context) (string, *register) {
	cnst := ident.cnst
	f := ctx.f
	loc := ctx.loc
	if isBool(cnst.Type()) {
		a, r := f.allocReg(loc, regType(cnst.Type()), 1)
		var val int8
		if cnst.Value.String() == "true" {
			val = 1
		}
		return a + MovImm8Reg(val, r, false), r
	}
	if isFloat(cnst.Type()) {
		a, r := f.allocReg(loc, regType(cnst.Type()), 1)
		if r.typ != XMM_REG {
			ice("can't load float const into non xmm register")
		}
		a2, tmp := f.allocReg(loc, DATA_REG, 8)
		asm := a + a2
		if isFloat32(cnst.Type()) {
			asm = MovImmf32Reg(float32(cnst.Float64()), tmp, r, false)
		} else {
			asm = MovImmf64Reg(cnst.Float64(), tmp, r, false)
		}
		f.freeReg(tmp)
		return asm, r

	}
	if isComplex(cnst.Type()) {
		ice("complex64/128 unsupported")
	}

	size := sizeof(cnst.Type())
	signed := signed(cnst.Type())
	a, r := f.allocReg(loc, regType(cnst.Type()), size)
	var val int64
	if signed {
		val = cnst.Int64()
	} else {

		val = int64(cnst.Uint64())
	}
	return a + MovImmReg(val, size, r, false), r
}

func (ident *identifier) newValue(reg *register, offset uint, size uint) string {
	// constants can't be modified
	if ident.cnst != nil {
		ice("cannot modify constant")
	}

	asm := ""
	chunk := region{offset, size}

	if ident.f.Trace {
		fmt.Printf("NEW VALUE %v (offset=%v, size=%v) -> %v (d=%v, v=%v)\n", ident.name, offset, size, reg.name, reg.dirty, reg.isValidAlias)
	}

	if reg.owner() != nil && reg.owner() != ident {
		dst := ident.getBestStorage(offset, size)
		copyChunk := transfer{srcOffset: reg.ownerRegion().offset, dstOffset: uint(offset), size: size}
		asm = reg.copyChunk(ident.f, copyChunk, dst)

		// HACK!, should only set valid if entire dst written to
		dst.setValid(true)

		for _, alias := range ident.aliases {
			if alias == reg || alias == dst {
				continue
			}
			if alias.ownerRegion().overlap(chunk).size > 0 {
				alias.setValid(false)
			}
		}

	} else {
		reg.setValid(true)
		if reg.owner() != ident {
			if reg.owner() != nil {
				owner := reg.owner()
				owner.removeAlias(reg)
			}
			reg.newOwner(ident, region{uint(offset), size})
			ident.addAlias(reg)
		}
		if reg.dirty {
			for _, alias := range ident.aliases {
				if alias == reg {
					continue
				}
				if alias.ownerRegion().overlap(chunk).size > 0 {
					alias.setValid(false)
				}
			}
		}
	}
	if ident.f.Trace {
		fmt.Printf("NEW VALUE %v (offset=%v, size=%v) -> %v (d=%v, v=%v)\n", ident.name, offset, size, reg.name, reg.dirty, reg.isValidAlias)
	}
	return asm
}

func (ident *identifier) dirtyRegions(exclude storage) []region {
	regions := []region{}
	for _, alias := range ident.aliases {
		if reg, ok := alias.(*register); ok {
			if reg.dirty && alias != exclude {
				regions = append(regions, reg.ownerRegion())
			}
		}
	}
	return regions
}

func (ident *identifier) spillAllRegisters() string {
	return ident.spillRegisters(true)
}

func (ident *identifier) spillDirtyRegisters() string {
	return ident.spillRegisters(false)
}

func (ident *identifier) spillRegisters(all bool) string {
	asm := ""
	for _, alias := range ident.aliases {
		if r, ok := alias.(*register); ok {
			asm += ident.spillRegister(r, all)
		}
	}
	return asm
}

func (ident *identifier) spillRegister(r *register, force bool) string {
	if r.owner().name != ident.name {
		ice("wrong owner for register")
	}

	asm := ""
	f := ident.f
	ident.removeAlias(r)

	if r.isValid() && (r.dirty || force) {
		for _, alias := range ident.aliases {
			overlap := r.ownerRegion().overlap(alias.ownerRegion())
			chunk := transfer{
				srcOffset: overlap.offset,
				dstOffset: overlap.offset,
				size:      overlap.size}
			if overlap.size != 0 {
				asm += r.copyChunk(f, chunk, alias)
			}
			dirtyRegions := ident.dirtyRegions(alias)
			dirtyOverlap := alias.ownerRegion()
			if len(dirtyRegions) == 0 {
				dirtyOverlap = region{}
			}
			for _, dirtyRegion := range dirtyRegions {
				dirtyOverlap = dirtyOverlap.overlap(dirtyRegion)
			}
			alias.setValid(dirtyOverlap.size == 0)
			if ident.f.Trace {
				fmt.Printf("Spill %v (d=%v) -> \"%v\" (valid %v), region %v\n", r.name, r.dirty, alias.String(), alias.isValid(), overlap)
			}

		}
	}
	r.newOwner(nil, region{})
	r.inUse = false
	return asm
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
