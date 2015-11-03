package codegen

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types"
)

type storage interface {
	String() string
	owner() *identifier
	newOwner(newOwner *identifier, newRegion region)
	ownerRegion() region
	newRegion(ownerRegion region)
	isValid() bool
	setValid(valid bool)
	copyto(f *Function, dst storage) string
	copyChunk(f *Function, chunk transfer, dst storage) string
}

type region struct {
	offset uint
	size   uint
}

type transfer struct {
	size      uint
	srcOffset uint
	dstOffset uint
}

type memory struct {
	stale   bool
	ownedby *identifier
	owns    region

	reg    *register
	offset int
	size   uint
}

type constant struct {
	ownedby *identifier
	owns    region
	*ssa.Const
}

func (r region) String() string {
	return fmt.Sprintf("{offset: %v, size: %v}", r.offset, r.size)
}

func (r region) min() uint {
	return r.offset
}

func (r region) max() uint {
	return r.offset + r.size
}

func (r region) contains(value uint) bool {
	return r.min() <= value && r.max() >= value
}

func (r region) overlap(other region) region {
	var min, max uint
	if r.min() > other.min() {
		min = r.min()
	} else {
		min = other.min()
	}
	if r.max() < other.max() {
		max = r.max()
	} else {
		max = other.max()
	}
	// no overlap
	if min >= max {
		return region{}
	} else {
		return region{offset: min, size: max - min}
	}
}

func (r *register) String() string {
	return fmt.Sprintf("register %v - %v (d=%v)", r.name, r.ownerRegion().String(), r.dirty)
}

func (r *register) OpDataType() OpDataType {
	if len(r.aliases()) == 0 {
		ice("cant load register, missing identifier")
	}
	var optype OpDataType
	if isIntegerSimd(r.owner().typ) || isSimd(r.owner().typ) || isSSE2(r.owner().typ) {
		optype = GetOpDataType(r.owner().typ)
	} else {
		optype = GetIntegerOpDataType(false, r.size())
	}
	return optype
}

func (r *register) owner() *identifier {
	return r.ownr
}

func (r *register) newOwner(owner *identifier, ownerRegion region) {
	r.ownr = owner
	r.ownrRegion = ownerRegion
}

func (r *register) ownerRegion() region {
	return r.ownrRegion
}

func (r *register) newRegion(newregion region) {
	r.ownrRegion = newregion
}

func (r *register) isValid() bool {
	return r.isValidAlias
}

func (r *register) setValid(valid bool) {
	r.isValidAlias = valid
}

func (r *register) copyto(f *Function, dst storage) string {
	chunk := transfer{size: r.OpDataType().size, srcOffset: 0, dstOffset: 0}
	return r.copyChunk(f, chunk, dst)
}

func (r *register) copyChunk(f *Function, chunk transfer, dst storage) string {
	if chunk.srcOffset != r.ownerRegion().offset {
		ice(fmt.Sprintf("cant transfer from register, with nonzero src offset (%v)", chunk.srcOffset))
	}
	if chunk.size > r.size() {
		msg := "cant copy chunk (size = %v) from register \"%v\" (size = %v)"
		ice(fmt.Sprintf(msg, chunk.size, r.name, r.size()))
	}
	if _, ok := dst.(*register); ok && chunk.dstOffset != 0 {
		ice("cant copy chunk to register, with nonzero dst offset ")
	}

	// nothing to copy
	if chunk.size == 0 {
		return ""
	}

	switch dst := dst.(type) {
	case *register:
		optype := r.OpDataType()
		optype.size = chunk.size
		//own := dst.owner()
		//validAlias := dst.isValidAlias
		//dst.ownr = nil
		return MovRegReg(optype, r, dst, false)
	case *memory:
		return r.save(chunk, dst)
	}
	// dst is not register or memory
	ice("invalid chunk copy destination")
	return ""
}

func (r *register) save(chunk transfer, dst *memory) string {
	optype := r.OpDataType()
	optype.size = chunk.size
	if dst.size < chunk.dstOffset+chunk.size {
		ice("cant save register chunk to memory")
	}
	dstOffset := dst.offset + int(chunk.dstOffset)
	return MovRegMem(optype, r, dst.name(), dst.reg, dstOffset)
}

func (m *memory) String() string {
	if m.ownedby == nil {
		ice("Unowned memory")
	}
	name := "memory " + m.ownedby.name + "+" + strconv.Itoa(m.offset) + "(" + m.reg.name + ")"
	name += " - " + m.ownerRegion().String()
	return strings.Replace(name, "+-", "-", -1)
}

func (m *memory) owner() *identifier {
	return m.ownedby
}

func (m *memory) newOwner(owner *identifier, ownerRegion region) {
	m.ownedby = owner
	m.owns = ownerRegion
}

func (m *memory) ownerRegion() region {
	return m.owns
}

func (m *memory) newRegion(newregion region) {
	m.owns = newregion
}

func (m *memory) name() string {
	return m.owner().name
}

func (m *memory) optype() OpDataType {
	return GetOpDataType(m.owner().typ)
}

func (m *memory) isValid() bool {
	return !m.stale
}

func (m *memory) setValid(valid bool) {
	if m.owner().f.Trace {
		fmt.Printf("%v (v=%v, old=%v)\n", m.String(), valid, !m.stale)
	}
	m.stale = !valid
}

func (m *memory) copyto(f *Function, dst storage) string {
	chunk := transfer{size: m.optype().size, srcOffset: 0, dstOffset: 0}
	return m.copyChunk(f, chunk, dst)
}

func (m *memory) copyChunk(f *Function, chunk transfer, dst storage) string {
	if chunk.size > m.size {
		msg := "cant transfer (size = %v) from \"%v\" (size = %v)"
		ice(fmt.Sprintf(msg, chunk.size, m.name(), m.size))
	}
	if _, ok := dst.(*register); ok && chunk.dstOffset != 0 {
		ice("cant transfer to register, with nonzero dst offset")
	}
	switch dst := dst.(type) {
	case *register:
		if isIntegerSimd(m.owner().typ) || isSimd(m.owner().typ) || isSSE2(m.owner().typ) {
			if chunk.size > XmmRegSize {
				ice("can't copy to register, size too large")
			}
			optype := m.optype()
			optype.size = chunk.size
			offset := m.offset + int(chunk.srcOffset)
			return MovMemReg(optype, m.name(), offset, m.reg, dst, false)
		} else {
			if chunk.size > DataRegSize {
				ice("can't copy to register, size too large")
			}
			optype := GetIntegerOpDataType(false, chunk.size)
			offset := m.offset + int(chunk.srcOffset)
			return MovMemReg(optype, m.name(), offset, m.reg, dst, false)
		}
	case *memory:
		return copyMemMem(m, dst, chunk)
	}
	// dst is not register or memory
	ice("invalid chunk copy destination")
	return ""
}

func copyMemMem(src, dst *memory, chunk transfer) string {
	size := chunk.size
	iterations := size
	datasize := uint(1)

	if size >= sizeBasic(types.Int64) {
		iterations = size / sizeBasic(types.Int64)
		datasize = 8
	} else if size >= sizeBasic(types.Int32) {
		iterations = size / sizeBasic(types.Int32)
		datasize = 4
	} else if size >= sizeBasic(types.Int16) {
		iterations = size / sizeBasic(types.Int16)
		datasize = 2
	}

	if size > sizeInt() {
		if size%sizeInt() != 0 {
			ice(fmt.Sprintf("Size (%v) not multiple of sizeInt (%v)", size, sizeInt()))
		}
	}
	f := src.owner().f
	asm := ""
	for i := uint(0); i < iterations; i++ {
		offset := i * datasize
		smallChunk := transfer{
			srcOffset: chunk.srcOffset + offset,
			dstOffset: chunk.dstOffset + offset,
			size:      datasize,
		}
		a, tmp := f.allocTempReg(DATA_REG, smallChunk.size)
		asm += a
		asm += copySmallMemMem(smallChunk, src, dst, tmp)
		f.freeReg(tmp)
	}
	return asm
}

func copySmallMemMem(chunk transfer, src, dst *memory, tmp *register) string {
	if dst.size < chunk.dstOffset+chunk.size {
		ice("cant copy chunk, chunk.dstOffset+chunk.size >= dst.size")
	}
	if src.size < chunk.srcOffset+chunk.size {
		ice("cant copy chunk, chunk.srcOffset+chunk.size >= src.size")
	}
	if chunk.size > DataRegSize {
		ice("trying to copy more one register sized piece of memory")
	}
	srcOffset := src.offset + int(chunk.srcOffset)
	dstOffset := dst.offset + int(chunk.dstOffset)
	optype := GetIntegerOpDataType(false, chunk.size)
	asm := MovMemMem(optype.op, src.name(), srcOffset, src.reg,
		dst.name(), dstOffset, dst.reg, chunk.size, tmp)
	return asm
}

func (cnst *constant) String() string {
	return "constant " + cnst.Const.String()
}

func (cnst *constant) owner() *identifier {
	return cnst.ownedby
}

func (cnst *constant) newOwner(owner *identifier, ownerRegion region) {
	if cnst.ownedby != owner {
		cnst.ownedby = owner
	}
}

func (cnst *constant) ownerRegion() region {
	return cnst.owns
}

func (cnst *constant) newRegion(newregion region) {
	cnst.owns = newregion
}

func (cnst *constant) isValid() bool {
	return true
}

func (cnst *constant) setValid(valid bool) {
	ice("unexpected, invalid constant")
}

func (cnst *constant) copyto(f *Function, dst storage) string {

	chunk := transfer{
		size:      GetOpDataType(cnst.owner().typ).size,
		srcOffset: 0,
		dstOffset: 0}

	return cnst.copyChunk(f, chunk, dst)
}

func (cnst *constant) copyChunk(f *Function, chunk transfer, dst storage) string {
	if chunk.srcOffset != 0 {
		ice(fmt.Sprintf("cant transfer from constant, with nonzero src offset (%v)", chunk.srcOffset))
	}
	if chunk.size != sizeof(cnst.owner().typ) {
		msg := "cant copy chunk (size = %v) from constant \"%v\" (size = %v)"
		ice(fmt.Sprintf(msg, chunk.size, cnst.owner().name, sizeof(cnst.owner().typ)))
	}
	if _, ok := dst.(*register); ok && chunk.dstOffset != 0 {
		ice("cant copy chunk to register, with nonzero dst offset ")
	}
	r, ok := dst.(*register)
	if !ok {
		ice("copying constant to memory not implemented")
	}
	if isBool(cnst.Type()) {
		var val int8
		if cnst.Value.String() == "true" {
			val = 1
		}
		return MovImm8Reg(val, r, false)
	}
	if isFloat(cnst.Type()) {
		if r.typ != XMM_REG {
			ice("can't load float const into non xmm register")
		}
		asm, tmp := f.allocTempReg(DATA_REG, 8)
		if isFloat32(cnst.Type()) {
			asm = MovImmf32Reg(float32(cnst.Float64()), tmp, r, false)
		} else {
			asm = MovImmf64Reg(cnst.Float64(), tmp, r, false)
		}
		f.freeReg(tmp)
		return asm

	}
	if isComplex(cnst.Type()) {
		ice("complex64/128 unsupported")
	}
	size := sizeof(cnst.Type())
	signed := signed(cnst.Type())
	var val int64
	if signed {
		val = cnst.Int64()
	} else {

		val = int64(cnst.Uint64())
	}
	return MovImmReg(val, size, r, false)
}
