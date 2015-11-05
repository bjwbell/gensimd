package codegen

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types"
)

type storer interface {
	String() string
	owner() *identifier
	ownerRegion() region
	load(ctx context, chunk region) (string, *register)
	store(ctx context, r *register, chunk region) string
	storeAndSpill(ctx context, r *register, chunk region) string
	spillRegister(ctx context, r *register, force bool) string
	spillRegisters(ctx context, force bool) string
	addAlias(ctx context, newAlias alias) bool
	removeAlias(ctx context, r *register) bool
	size() uint
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

// alias between src and dst
type alias struct {
	dst *register
	region
}

type aliaser struct {
	src     storer
	aliases []alias
}

type memory struct {
	parent *identifier
	*aliaser
	initializedRegions []region
}

type constant struct {
	parent *identifier
	*ssa.Const
	*aliaser
}

func (a alias) String() string {
	return a.dst.name + " {offset:" + strconv.Itoa(int(a.offset)) + ", size: " + strconv.Itoa(int(a.size)) + "}"
}

// check looks for duplicate aliases and panics if any are found
func (a *aliaser) check() {
	for i := range a.aliases {
		for j := range a.aliases {
			if i == j {
				continue
			}
			a1 := a.aliases[i]
			a2 := a.aliases[j]
			if a1.dst == a2.dst {
				ice("duplicate aliases")
			}
			if a1.overlap(a2.region).size != 0 {
				ice("overlapping aliases")
			}
		}
	}
}

func (a *aliaser) addAlias(ctx context, newAlias alias) bool {
	// ice on duplicate/overlapping aliases
	a.check()
	newAlias.dst.parent = a.src
	identName := a.src.owner().name
	if ctx.f.Trace {
		fmt.Printf(ctx.f.Indent+"Alias: %v -> %v\n", identName, newAlias.String())

	}
	new := true
	for _, alias := range a.aliases {
		if alias.dst == newAlias.dst {
			new = false
		}
		if ctx.f.Trace {
			fmt.Printf(ctx.f.Indent+"     : %v -> %v\n", identName, alias.dst.name)
		}
	}
	if new {
		a.aliases = append(a.aliases, newAlias)
	}
	// ice on duplicate/overlapping aliases
	a.check()
	return new
}

func (a *aliaser) removeAlias(ctx context, r *register) bool {
	// ice on duplicate/overlapping aliases
	a.check()
	if r.parent != a.src {
		ice(fmt.Sprintf("cannot remove register (%v) alias", r))
	}
	aliases := []alias{}
	removed := false
	identName := a.src.owner().name
	if ctx.f.Trace {
		fmt.Printf(ctx.f.Indent+"Alias: %v -/ %v (d=%v)\n", identName, r.name, r.dirty)

	}
	for _, alias := range a.aliases {
		if alias.dst == r {
			removed = true
		} else {
			aliases = append(aliases, alias)
			if ctx.f.Trace {
				fmt.Printf(ctx.f.Indent+"     : %v -> %v\n", identName, alias.String())
			}
		}
	}
	if !removed {
		fmt.Println("IDENT.NAME:", identName)
		fmt.Println("REMOVEAL ALIAS:", r.name)
		panic("DIDN'T REMOVE ALIAS")
	}
	a.aliases = aliases
	r.parent = nil
	ctx.f.freeReg(r)
	// ice on duplicate/overlapping aliases
	a.check()
	return removed
}

func (a *aliaser) removeAliases() bool {
	// ice on duplicate/overlapping aliases
	a.check()

	hasAliases := len(a.aliases) > 0
	for _, alias := range a.aliases {
		parent := alias.dst.parent
		if parent != a.src {
			ice("invalid alias")
		}
		alias.dst.parent = nil
		alias.dst.dirty = false
	}
	a.aliases = nil
	return hasAliases
}

func (a *aliaser) isLoaded(region region) bool {
	// ice on duplicate/overlapping aliases
	a.check()
	return a.fetch(region) != nil
}

func (a *aliaser) fetch(region region) *register {
	// ice on duplicate/overlapping aliases
	a.check()
	var r *register
	for _, alias := range a.aliases {
		if alias.offset == region.offset {
			if alias.size == region.size {
				r = alias.dst
				break
			} else {
				ice("alias size doesnt match region size")
			}
		}
	}
	// ice on duplicate/overlapping aliases
	a.check()
	return r
}

func (a *aliaser) getAlias(r *register) *alias {
	// ice on duplicate/overlapping aliases
	a.check()
	var regAlias *alias
	for _, alias := range a.aliases {
		if alias.dst == r {
			regAlias = &alias
			break
		}
	}
	// ice on duplicate/overlapping aliases
	a.check()
	return regAlias
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

func (r region) overlapRegions(regions []region) region {
	// ice if regions are not disjoint
	checkDisjoint(regions)
	overlaps := []region{}
	for _, region := range regions {
		if overlap := r.overlap(region); overlap.size > 0 {
			overlaps = append(overlaps, overlap)
		}
	}
	overlaps = mergeRegions(overlaps)
	if len(overlaps) > 1 {
		ice("unexpected, disjoint region overlaps")
		return region{}
	} else if len(overlaps) == 1 {
		return overlaps[0]
	} else {
		return region{}
	}
}

func mergeRegions(regions []region) []region {
	// ice if regions are not disjoint
	checkDisjoint(regions)
	var merged []region
	for _, region := range regions {
		merged = mergeRegion(merged, region)
	}
	// ice if regions are not disjoint
	checkDisjoint(merged)
	return merged
}

func mergeRegion(merged []region, newRegion region) []region {
	// ice if regions are not disjoint
	checkDisjoint(merged)
	if newRegion.overlapRegions(merged).size > 0 {
		ice("non-disjoint regions")
	}
	newMerged := []region{}
	if len(merged) == 0 {
		newMerged = append(newMerged, newRegion)
		return newMerged
	}
	for i := range merged {
		region := merged[i]
		if newRegion.max() == merged[i].min() {
			region.size += newRegion.size
			region.offset = newRegion.offset
		} else if newRegion.min() == merged[i].max() {
			region.size += newRegion.size
		}
		newMerged = append(newMerged, region)
	}
	// ice if regions are not disjoint
	checkDisjoint(newMerged)
	return newMerged
}

func checkDisjoint(regions []region) {
	if !disjoint(regions) {
		ice("non-disjoint regions")
	}
}

func disjoint(regions []region) bool {
	for i := range regions {
		for j := range regions {
			if i == j {
				continue
			}
			if regions[i].overlap(regions[j]).size > 0 {
				return false
			}
		}
	}
	return true
}

func (r *register) String() string {
	parent := "NONE"
	if r.parent != nil {
		parent = r.parent.owner().name
	}
	return fmt.Sprintf("register %v - %v (d=%v)", r.name, parent, r.dirty)
}

func (r *register) OpDataType() OpDataType {
	if r.parent == nil {
		ice("cant load register, missing identifier")
	}
	var optype OpDataType
	typ := r.parent.owner().typ
	if isXmm(typ) {
		optype = GetOpDataType(typ)
	} else {
		optype = GetIntegerOpDataType(false, r.size())
	}
	return optype
}

func (r *register) save(ctx context, chunk region, dst *memory) string {
	optype := r.OpDataType()
	optype.size = chunk.size
	if dst.size() < chunk.offset+chunk.size {
		ice("cant save register chunk to memory")
	}
	dstOffset := dst.offset() + int(chunk.offset)
	return MovRegMem(ctx, optype, r, dst.name(), dst.reg(), dstOffset)
}

func (m *memory) reg() *register {
	r, _, _ := m.Addr()
	return &r
}

func (m *memory) offset() int {
	_, offset, _ := m.Addr()
	return offset
}

func (m *memory) size() uint {
	_, _, size := m.Addr()
	return size
}

func (m *memory) String() string {
	if m.parent == nil {
		ice("Unowned memory")
	}
	name := "memory " + m.parent.name + "+" + strconv.Itoa(m.offset()) + "(" + m.reg().name + ")"
	name += " - " + m.ownerRegion().String()
	return strings.Replace(name, "+-", "-", -1)
}

func (m *memory) owner() *identifier {
	return m.parent
}

func (m *memory) ownerRegion() region {
	return m.parent.storageRegion()
}

func (m *memory) Addr() (register, int, uint) {
	return m.parent.Addr()
}

func (m *memory) name() string {
	return m.owner().name
}

func (m *memory) optype() OpDataType {
	if isXmm(m.owner().typ) {
		return GetOpDataType(m.owner().typ)
	} else {
		size := m.owner().size()
		if size > DataRegSize {
			size = DataRegSize
		}
		return GetIntegerOpDataType(false, size)
	}
}

func (m *memory) isInitialized(rgn region) bool {
	initialized := rgn.overlapRegions(m.initializedRegions)
	if initialized.size == 0 {
		return false
	} else {
		if initialized.size == rgn.size && initialized.offset != rgn.offset {
			ice("invalid initialized memory region")
		}
		return rgn.size == initialized.size
	}
}

func (m *memory) setInitialized(rgn region) {
	// region already initialized
	if m.isInitialized(rgn) {
		return
	} else {
		cpy := []region{}
		for _, r := range m.initializedRegions {
			cpy = append(cpy, r)
		}
		cpy = append(cpy, rgn)
		if !disjoint(cpy) {
			fmt.Println("rgn: ", rgn)

			fmt.Println("overlap: ", rgn.overlapRegions(m.initializedRegions))
			for _, r := range m.initializedRegions {
				fmt.Println("init rgn: ", r)
				fmt.Println("over: ", rgn.overlap(r))
			}
		}
		m.initializedRegions = mergeRegions(append(m.initializedRegions, rgn))
	}
}

func (m *memory) load(ctx context, chunk region) (string, *register) {
	if r := m.fetch(chunk); r != nil {
		r.inUse = true
		return "", r
	} else {
		return m.loadNew(ctx, chunk)
	}
}

func (m *memory) loadNew(ctx context, chunk region) (string, *register) {
	asm, r := ctx.f.allocIdentReg(ctx.loc, m.owner(), chunk.size)
	m.addAlias(ctx, alias{r, chunk})
	if isXmm(m.owner().typ) {
		if chunk.size > XmmRegSize {
			ice("can't copy to register, size too large")
		}
	} else {
		if chunk.size > DataRegSize {
			ice("can't copy to register, size too large")
		}
	}
	optype := m.optype()
	optype.size = chunk.size
	offset := m.offset() + int(chunk.offset)
	asm += MovMemReg(ctx, optype, m.name(), offset, m.reg(), r, false)
	r.dirty = false
	return asm, r
}

func (m *memory) store(ctx context, r *register, chunk region) string {
	return m.storeMem(ctx, r, chunk, false)
}

func (m *memory) storeMem(ctx context, r *register, chunk region, forceMem bool) string {
	if r.parent == m {
		if m.fetch(chunk) != r {
			ice("invalid register aliasing")
		}
		return ""
	}
	if dst := m.fetch(chunk); dst != nil {
		if !forceMem {
			return MovRegReg(ctx, m.optype(), r, dst, false)
		} else {
			m.removeAlias(ctx, dst)
		}
	}

	if !forceMem {
		// can repurpose r to alias m
		if unassignRegister(r, ctx.loc) {
			m.addAlias(ctx, alias{dst: r, region: chunk})
			return ""
		}
		f := m.owner().f
		if newReg := f.allocUnusedReg(regType(m.owner().typ), chunk.size); newReg != nil {
			m.addAlias(ctx, alias{dst: newReg, region: chunk})
			return MovRegReg(ctx, m.optype(), r, newReg, false)
		}
	}
	m.setInitialized(chunk)
	return r.save(ctx, chunk, m)
}

func unassignRegister(r *register, loc ssa.Instruction) bool {
	// owner missing
	if r.parent == nil {
		return true
	}
	owner := r.parent.owner()
	if owner == nil {
		ice("all storage should have owner")
	}
	if owner.f.Optimize && !aliveAfter(owner, loc) {
		if !r.parent.removeAlias(context{owner.f, loc}, r) {
			ice("couldnt unassign register from identifier")
		}
		return true
	}
	return false
}

func (m *memory) storeAndSpill(ctx context, r *register, chunk region) string {
	asm := m.storeMem(ctx, r, chunk, true)
	asm += m.spillRegisters(ctx, false)
	return asm
}

func (m *memory) spillRegister(ctx context, r *register, force bool) string {
	regAlias := m.getAlias(r)
	if regAlias == nil {
		ice("cannot spill register")
	}
	asm := ""

	if ctx.f.Optimize && !alive(m.owner(), ctx.loc) {
		// case 0 - do nothing, ident is dead
		if ctx.f.Trace {
			ident := m.owner()
			fmt.Printf(ident.f.Indent+"not spilling %v, %v dead\n",
				r.name, ident.name)
		}

	} else if r.dirty || force || !m.isInitialized(regAlias.region) {
		// case 1 - dirty == true or force == true or memory isn't initialized
		asm = r.save(ctx, regAlias.region, m)
		m.setInitialized(regAlias.region)
	} else {
		// case 2 - dirty == false and force == false, nothing to do
	}
	m.removeAlias(ctx, r)
	return asm
}

func (m *memory) spillRegisters(ctx context, force bool) string {
	asm := ""
	for _, alias := range m.aliases {
		asm += m.spillRegister(ctx, alias.dst, force)
	}
	return asm
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
	ctx := context{src.parent.f, nil}
	for i := uint(0); i < iterations; i++ {
		offset := i * datasize
		smallChunk := transfer{
			srcOffset: chunk.srcOffset + offset,
			dstOffset: chunk.dstOffset + offset,
			size:      datasize,
		}
		if r := src.fetch(region{smallChunk.srcOffset, datasize}); r != nil {
			asm += MovRegMem(ctx, src.optype(), r, dst.name(), dst.reg(), dst.offset()+int(smallChunk.dstOffset))
		} else {
			a, tmp := f.allocTempReg(DATA_REG, smallChunk.size)
			asm += a
			asm += copySmallMemMem(smallChunk, src, dst, tmp)
			f.freeReg(tmp)
		}
	}
	return asm
}

func copySmallMemMem(chunk transfer, src, dst *memory, tmp *register) string {
	if dst.size() < chunk.dstOffset+chunk.size {
		ice("cant copy chunk, chunk.dstOffset+chunk.size >= dst.size")
	}
	if src.size() < chunk.srcOffset+chunk.size {
		ice("cant copy chunk, chunk.srcOffset+chunk.size >= src.size")
	}
	if chunk.size > DataRegSize {
		ice("trying to copy more one register sized piece of memory")
	}
	srcOffset := src.offset() + int(chunk.srcOffset)
	dstOffset := dst.offset() + int(chunk.dstOffset)
	optype := GetIntegerOpDataType(false, chunk.size)
	ctx := context{src.parent.f, nil}
	asm := MovMemMem(ctx, optype.op, src.name(), srcOffset, src.reg(),
		dst.name(), dstOffset, dst.reg(), chunk.size, tmp)
	return asm
}

func (cnst *constant) String() string {
	return "constant " + cnst.Const.String()
}

func (cnst *constant) owner() *identifier {
	return cnst.parent
}

func (cnst *constant) ownerRegion() region {
	return cnst.parent.storageRegion()
}

func (cnst *constant) isValid() bool {
	return true
}

func (cnst *constant) setValid(valid bool) {
	ice("unexpected, invalid constant")
}

func (cnst *constant) size() uint {
	if isBool(cnst.Type()) {
		if sizeof(cnst.Type()) != 8 {
			fmt.Println("SIZEOF CONST BOOLEAN != 8")
			return 8
		}
		return sizeof(cnst.Type())
	}
	if isFloat(cnst.Type()) {
		if isFloat32(cnst.Type()) {
			if sizeof(cnst.Type()) != 4 {
				fmt.Println("SIZEOF CONST float32 != 4")
				return 4
			}
			return sizeof(cnst.Type())
		} else {
			if sizeof(cnst.Type()) != 8 {
				fmt.Println("SIZEOF CONST float64 != 8")
				return 8
			}
			return sizeof(cnst.Type())
		}

	}
	if isComplex(cnst.Type()) {
		ice("complex64/128 unsupported")
	}
	return sizeof(cnst.Type())
}

func (cnst *constant) load(ctx context, chunk region) (string, *register) {
	if r := cnst.fetch(chunk); r != nil {
		r.inUse = true
		return "", r
	}
	return cnst.loadNew(ctx, chunk)
}

func (cnst *constant) loadNew(ctx context, chunk region) (string, *register) {
	asm, r := ctx.f.allocIdentReg(ctx.loc, cnst.owner(), chunk.size)
	cnst.addAlias(ctx, alias{r, chunk})
	if isBool(cnst.Type()) {
		var val int8
		if cnst.Value.String() == "true" {
			val = 1
		}
		asm += MovImm8Reg(ctx, val, r, false)
	} else if isFloat(cnst.Type()) {
		if r.typ != XMM_REG {
			ice("can't load float const into non xmm register")
		}
		a, tmp := ctx.f.allocTempReg(DATA_REG, 8)
		asm += a
		if isFloat32(cnst.Type()) {
			a = MovImmf32Reg(ctx, float32(cnst.Float64()), tmp, r, false)
			asm += a
		} else {
			a = MovImmf64Reg(ctx, cnst.Float64(), tmp, r, false)
			asm += a
		}
		ctx.f.freeReg(tmp)

	} else if isComplex(cnst.Type()) {
		ice("complex64/128 unsupported")
	} else {
		size := sizeof(cnst.Type())
		signed := signed(cnst.Type())
		var val int64
		if signed {
			val = cnst.Int64()
		} else {
			val = int64(cnst.Uint64())
		}
		asm += MovImmReg(ctx, val, size, r, false)
	}
	r.dirty = false
	return asm, r

}

func (cnst *constant) store(ctx context, r *register, chunk region) string {
	ice("cannot modify constants")
	return ""
}

func (cnst *constant) storeAndSpill(ctx context, r *register, chunk region) string {
	ice("cannot modify constants")
	return ""
}

func (cnst *constant) spillRegister(ctx context, r *register, force bool) string {
	if r.dirty || force {
		ice(fmt.Sprintf("cannot write modified register (%v) to constant (dirty=%v, f=%v)", r.name, r.dirty, force))
	}
	if !cnst.removeAlias(ctx, r) {
		ice("cannot remove register alias")
	}
	return ""
}

func (cnst *constant) spillRegisters(ctx context, force bool) string {
	asm := ""
	for _, alias := range cnst.aliases {
		asm += cnst.spillRegister(ctx, alias.dst, force)
	}
	if asm != "" {
		ice("Assembly code generated, while spilling constant registers")
	}
	return asm
}
