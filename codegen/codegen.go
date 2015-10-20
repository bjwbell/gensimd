package codegen

import (
	"errors"
	"fmt"
	"go/token"
	"math"
	"strconv"
	"strings"

	"golang.org/x/tools/go/types"

	"golang.org/x/tools/go/ssa"
)

type phiInfo struct {
	value ssa.Value
	phi   *ssa.Phi
}

type Function struct {
	// output function name
	outfn       string
	Indent      string
	ssa         *ssa.Function
	registers   map[string]bool // maps register to false if unused and true if used
	identifiers map[string]identifier
	// map from block index to the successor block indexes that need phi vars set
	phiInfo map[int]map[int][]phiInfo

	jmpLabels []string
}

type identifier struct {
	name   string
	typ    types.Type
	local  *varInfo
	param  *paramInfo
	cnst   *ssa.Const
	offset int
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
		internal(fmt.Sprintf("identifier (%v) is not a local or param", name))
	}
	return
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
		internal(fmt.Sprintf("identifier (%v) not ptr type", name))
	}
	ptrType := name.typ.(*types.Pointer)
	return ptrType.Elem()
}

func (name *identifier) IsArray() bool {
	_, ok := name.typ.(*types.Array)
	return ok
}

func (name *identifier) IsSlice() bool {
	_, ok := name.typ.(*types.Slice)
	return ok
}

func (name *identifier) IsBasic() bool {
	_, ok := name.typ.(*types.Basic)
	return ok
}

func (name *identifier) IsInteger() bool {
	if !name.IsBasic() {
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

type paramSlice struct {
	lenOffset int
}

type Error struct {
	Err error
	Pos token.Pos
}

func ErrorMsg(msg string) (string, *Error) {
	return "", ErrorMsg2(msg)
}

func ErrorMsg2(msg string) *Error {
	return &Error{Err: errors.New(msg), Pos: 0}
}

func CreateFunction(fn *ssa.Function, outfn string) (*Function, *Error) {
	if fn == nil {
		return nil, ErrorMsg2("Nil function passed in")
	}
	f := Function{ssa: fn, outfn: outfn}
	f.Indent = "        "
	f.init()
	return &f, nil
}

func AssemblyFilePreamble() string {
	return "// +build amd64\n\n"
}

func (f *Function) GoAssembly() (string, *Error) {
	return f.Func()
}

func (f *Function) Params() (string, *Error) {
	// offset in bytes from frame pointer (FP)
	offset := int(0)
	asm := ""
	for _, p := range f.ssa.Params {
		param := paramInfo{name: p.Name(), info: p}
		// TODO alloc reg based on other param types
		if _, ok := p.Type().(*types.Slice); ok {
			param.extra = paramSlice{lenOffset: offset + int(sizePtr())}
		} else if basic, ok := p.Type().(*types.Basic); ok {
			switch basic.Kind() {
			default:
				return ErrorMsg(fmt.Sprintf("Unsupported param type (%v)", basic))
			case types.Float32, types.Float64:
				break

				// supported param types
			case types.Bool, types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
				types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
				break
			}

		} else {

		}
		info := identifier{name: param.name, typ: param.info.Type(),
			local: nil, param: &param, offset: offset}
		f.identifiers[param.name] = info
		if info.align() > info.size() {
			offset += int(info.align())
		} else {
			offset += int(info.size())
		}
	}
	return asm, nil
}

func (f *Function) Func() (string, *Error) {

	params, err := f.Params()
	if err != nil {
		return params, err
	}

	zeroRetValue, err := f.ZeroRetValue()
	if err != nil {
		return params + zeroRetValue, err
	}

	zeroSsaLocals, err := f.ZeroSsaLocals()
	if err != nil {
		return params + zeroRetValue + zeroSsaLocals, err
	}

	if err := f.computePhi(); err != nil {
		return "", err
	}

	basicblocks, err := f.BasicBlocks()
	if err != nil {
		return params + zeroRetValue + zeroSsaLocals + basicblocks, err
	}

	zeroNonSsaLocals, err := f.ZeroNonSsaLocals()
	if err != nil {
		return zeroNonSsaLocals, err
	}

	frameSize := f.localsSize()
	frameSize = f.align(frameSize)
	argsSize := f.retOffset() + int(f.retAlign())
	asm := params
	asm += f.SetStackPointer()
	asm += zeroRetValue
	asm += zeroSsaLocals
	asm += zeroNonSsaLocals
	asm += basicblocks
	asm = f.fixupRets(asm)
	a := fmt.Sprintf("TEXT ·%v(SB),NOSPLIT,$%v-%v\n%v", f.outfname(), frameSize, argsSize, asm)
	return a, nil
}

func (f *Function) newJmpLabel() string {
	label := "lbl" + strconv.Itoa(len(f.jmpLabels)+1)
	f.jmpLabels = append(f.jmpLabels, label)
	return label
}

func (f *Function) align(size uint32) uint32 {
	// on amd64 stack size should be 8 byte aligned
	align := f.stackAlign()
	return size + (align - size%align)
}

// amd64 has 8 byte stack alignment
var stackAlignment uint32 = 8

func (f *Function) stackAlign() uint32 {
	return stackAlignment
}

func (f *Function) GoProto() (string, string) {
	pkgname := "package " + f.ssa.Package().Pkg.Name() + "\n"
	fnproto := "func " + f.outfname() + "(" + strings.TrimPrefix(f.ssa.Signature.String(), "func(")
	return pkgname, fnproto
}

func (f *Function) outfname() string {
	if f.outfn != "" {
		return f.outfn
	}
	return f.ssa.Name()
}

func (f *Function) ZeroSsaLocals() (string, *Error) {

	asm := f.Indent + "// BEGIN ZeroSsaLocals\n"
	offset := int(0)
	locals := f.ssa.Locals

	for _, local := range locals {
		if local.Heap {
			return ErrorMsg(fmt.Sprintf("Can't heap alloc local, name: %v", local.Name()))
		}
		sp := getRegister(REG_SP)

		//local values are always addresses, and have pointer types, so the type
		//of the allocated variable is actually
		//Type().Underlying().(*types.Pointer).Elem().
		typ := local.Type().Underlying().(*types.Pointer).Elem()
		size := sizeof(typ)
		localOffset := -(offset + int(size))
		asm += ZeroMemory(f.Indent, local.Name(), localOffset, size, sp)
		v := varInfo{name: local.Name(), info: local}
		info := identifier{name: v.name, typ: typ, local: &v, param: nil, offset: localOffset}
		f.identifiers[v.name] = info
		if info.align() > info.size() {
			offset += int(info.align())
		} else {
			offset += int(size)
		}
	}

	asm += f.Indent + "// END ZeroSsaLocals\n"
	return asm, nil
}

func (f *Function) AllocLocal(name string, typ types.Type) (identifier, *Error) {
	size := sizeof(typ)
	offset := int(size)
	if align(typ) > size {
		offset = int(align(typ))
	}
	v := varInfo{name: name, info: nil}
	info := identifier{
		name:   name,
		typ:    typ,
		param:  nil,
		local:  &v,
		offset: -int(f.localsSize()) - offset}

	f.identifiers[v.name] = info
	// zeroing the memory is done at the beginning of the function
	return info, nil
}

func (f *Function) ZeroNonSsaLocals() (string, *Error) {
	asm := f.Indent + "// BEGIN ZeroNonSsaLocals\n"
	for _, name := range f.identifiers {
		if name.local == nil || name.IsSsaLocal() {
			continue
		}
		sp := getRegister(REG_SP)
		asm += ZeroMemory(f.Indent, name.name, name.offset, name.size(), sp)
	}
	asm += f.Indent + "// END ZeroNonSsaLocals\n"
	return asm, nil
}

func (f *Function) ZeroRetValue() (string, *Error) {
	asm := f.Indent + "// BEGIN ZeroRetValue\n"
	asm += ZeroMemory(f.Indent, retName(), f.retOffset(), f.retSize(), getRegister(REG_FP))
	asm += f.Indent + "// END ZeroRetValue\n"
	return asm, nil
}

func (f *Function) BasicBlocks() (string, *Error) {
	asm := ""
	for i := 0; i < len(f.ssa.Blocks); i++ {
		a, err := f.BasicBlock(f.ssa.Blocks[i])
		asm += a
		if err != nil {
			return asm, err
		}
	}
	return asm, nil
}

func (f *Function) BasicBlock(block *ssa.BasicBlock) (string, *Error) {
	asm := "block" + strconv.Itoa(block.Index) + ":\n"
	for i := 0; i < len(block.Instrs); i++ {
		a, err := f.Instr(block.Instrs[i])
		asm += a
		if err != nil {
			return asm, err
		}
	}
	return asm, nil
}

func (f *Function) Instr(instr ssa.Instruction) (string, *Error) {

	if instr == nil {
		internal("nil instr")
	}
	asm := ""
	var err *Error

	errormsg := func(msg string) (string, *Error) {
		return "", &Error{Err: fmt.Errorf(msg), Pos: instr.Pos()}
	}

	switch instr := instr.(type) {
	default:
		asm = f.Indent + fmt.Sprintf("Unknown ssa instruction: %v\n", instr)
	case *ssa.Alloc:
		asm, err = f.AllocInstr(instr)
	case *ssa.BinOp:
		asm, err = f.BinOp(instr)
	case *ssa.Call:
		asm = f.Indent + fmt.Sprintf("ssa.Call: %v, name: %v\n", instr, instr.Name())
	case *ssa.ChangeInterface:
		asm, err = errormsg("converting interfaces unsupported")
	case *ssa.ChangeType:
		asm, err = errormsg("changing between types unsupported")
	case *ssa.Convert:
		asm, err = f.Convert(instr)
	case *ssa.Defer:
		asm, err = errormsg("defer unsupported")
	case *ssa.Extract:
		asm, err = errormsg("extracting tuple values unsupported")
	case *ssa.Field:
		asm, err = errormsg("field access unimplemented")
	case *ssa.FieldAddr:
		asm, err = errormsg("field access unimplemented")
	case *ssa.Go:
		asm, err = errormsg("go keyword unsupported")
	case *ssa.If:
		asm, err = f.If(instr)
	case *ssa.Index:
		asm, err = f.Index(instr)
	case *ssa.IndexAddr:
		asm, err = f.IndexAddr(instr)
	case *ssa.Jump:
		asm, err = f.Jump(instr)
	case *ssa.Lookup:
		asm, err = errormsg("maps unsupported")
	case *ssa.MakeChan:
		asm, err = errormsg("channels unsupported")
	case *ssa.MakeClosure:
		asm, err = errormsg("closures unsupported")
	case *ssa.MakeInterface, *ssa.MakeMap, *ssa.MakeSlice:
		asm, err = errormsg("make slice/map/interface unsupported")
	case *ssa.MapUpdate:
		asm, err = errormsg("map update unsupported")
	case *ssa.Next:
		asm, err = errormsg("map/string iterators unsupported")
	case *ssa.Panic:
		asm, err = errormsg("panic unimplemented")
	case *ssa.Phi:
		asm, err = f.Phi(instr)
	case *ssa.Range:
		asm, err = errormsg("range unsupported")
	case *ssa.Return:
		asm, err = f.Return(instr)
	case *ssa.Select, *ssa.RunDefers, *ssa.Send:
		asm, err = errormsg("select/send/defer unsupported")
	case *ssa.Slice:
		asm, err = f.Slice(instr)
	case *ssa.Store:
		asm, err = f.Store(instr)
	case *ssa.TypeAssert:
		asm, err = errormsg("type assert unsupported")
	case *ssa.UnOp:
		asm, err = f.UnOp(instr)
	}

	return asm, err
}

type fromto struct {
	from types.BasicKind
	to   types.BasicKind
}

func GetXmmVariant(t types.Type) XmmData {
	if isFloat32(t) {
		return XMM_F32
	} else if isFloat64(t) {
		return XMM_F64
	}
	return XMM_INVALID
}

func (f *Function) Slice(instr *ssa.Slice) (string, *Error) {
	return ErrorMsg("slice creation unsupported")
}

func (f *Function) Convert(instr *ssa.Convert) (string, *Error) {
	from := instr.X.Type()
	to := instr.Type()

	var fromType InstrOpType
	var toType InstrOpType
	var fromXmm XmmData
	var toXmm XmmData

	if isInteger(from) && isInteger(to) {
		fromType = INTEGER_OP
		toType = INTEGER_OP
		fromXmm = XMM_INVALID
		toXmm = XMM_INVALID
	} else if isInteger(from) && isFloat(to) {
		fromType = INTEGER_OP
		toType = XMM_OP
		fromXmm = XMM_INVALID
		toXmm = GetXmmVariant(to)
	} else if isFloat(from) && isInteger(to) {
		fromType = XMM_OP
		fromXmm = GetXmmVariant(from)
		toType = INTEGER_OP
		toXmm = XMM_INVALID
	} else if isFloat(from) && isFloat(to) {
		fromType = XMM_OP
		fromXmm = GetXmmVariant(from)
		toType = XMM_OP
		toXmm = GetXmmVariant(to)
	} else {
		fmtstr := "Cannot convert from (%v) to (%v)"
		msg := fmt.Sprintf(fmtstr, from.String(), to.String())
		return ErrorMsg(msg)
	}

	return f.ConvertFromTo(instr, fromType, toType, fromXmm, toXmm)
}

func (f *Function) AllocFromToRegs(instr *ssa.Convert) (from register, to register) {

	fromInfo := f.allocOnDemand(instr.X)
	toInfo := f.allocOnDemand(instr)
	if fromInfo == nil || toInfo == nil {
		internal("converting between types")
	}
	fromreg := f.allocReg(regType(fromInfo.typ), fromInfo.size())
	toreg := f.allocReg(regType(toInfo.typ), toInfo.size())

	return fromreg, toreg

}

func (f *Function) ConvertFromTo(instr *ssa.Convert, fromOpType, toOpType InstrOpType, fromXmm, toXmm XmmData) (string, *Error) {

	asm := ""
	from, to := f.AllocFromToRegs(instr)
	if a, err := f.LoadValueSimple(instr.X, &from); err != nil {
		return "", err
	} else {
		asm += a
	}
	tmp := f.allocReg(DATA_REG, sizeInt())

	fromType :=
		OpDataType{
			op:         fromOpType,
			InstrData:  InstrData{signed: signed(instr.X.Type()), size: f.sizeof(instr.X)},
			xmmvariant: fromXmm}
	toType :=
		OpDataType{
			op:         toOpType,
			InstrData:  InstrData{signed: signed(instr.Type()), size: f.sizeof(instr)},
			xmmvariant: toXmm}

	// round uint64 to nearest int64 before converting to float32/float64
	if isUint(instr.X.Type()) && f.sizeof(instr.X) == 8 && isFloat(instr.Type()) {
		var floatSize uint
		if isFloat32(instr.Type()) {
			floatSize = 32
		} else if isFloat64(instr.Type()) {
			floatSize = 64
		} else {
			internal("converting from uint64 to float")
		}

		asm += f.ConvertUint64ToFloat(&tmp, &from, &to, floatSize)
	} else {
		asm += ConvertOp(f.Indent, &from, fromType, &to, toType, &tmp)
	}
	toNameInfo := f.identifiers[instr.Name()]
	if a, err := f.StoreValue(&toNameInfo, &to); err != nil {
		return "", err
	} else {
		asm += a
	}
	f.freeReg(from)
	f.freeReg(to)
	f.freeReg(tmp)

	return asm, nil
}

func (f *Function) ConvertUint64ToFloat(tmp, regU64, regFloat *register, floatSize uint) string {

	asm := ""
	regI64 := f.allocReg(DATA_REG, 8)
	round := f.newJmpLabel()
	noround := f.newJmpLabel()
	end := f.newJmpLabel()
	cvt := ""
	add := ""
	if floatSize == 32 {
		cvt = "CVTSQ2SS"
		add = "ADDSS"
	} else {
		cvt = "CVTSQ2SD"
		add = "ADDSD"
	}
	str := f.Indent + "//      U64\n" +
		f.Indent + "CMPQ	%v, $-1\n" +
		f.Indent + "//      jmp to rounding\n" +
		f.Indent + "JEQ	%v\n" +
		f.Indent + "//     U64\n" +
		f.Indent + "CMPQ	%v, $-1\n" +
		f.Indent + "//      jmp to no rounding\n" +
		f.Indent + "JGE	%v\n" +
		f.Indent + "// rounding label\n" +
		f.Indent + "%v:\n" +
		f.Indent + "//     U64 I64\n" +
		f.Indent + "MOVQ	%v, %v\n" +
		f.Indent + "//         I64\n" +
		f.Indent + "SHRQ	$1, %v\n" +
		f.Indent + "//     U64 TMP\n" +
		f.Indent + "MOVQ	%v, %v\n" +
		f.Indent + "//         TMP\n" +
		f.Indent + "ANDL	$1, %v\n" +
		f.Indent + "//     TMP I64\n" +
		f.Indent + "ORQ	%v, %v\n" +
		f.Indent + "//CVT  I64 XMM\n" +
		f.Indent + "%v    %v, %v\n" +
		f.Indent + "//ADD XMM, XMM\n" +
		f.Indent + "%v    %v, %v\n" +
		f.Indent + "//    jmp to end\n" +
		f.Indent + "JMP   %v\n" +
		f.Indent + "//    jmp label for no rounding\n" +
		f.Indent + "%v:\n" +
		f.Indent + "//CVT U64 XMM\n" +
		f.Indent + "%v     %v, %v\n" +
		f.Indent + "//    jmp label for end\n" +
		f.Indent + "%v:\n"

	asm += fmt.Sprintf(str,
		regU64.name,
		round,
		regU64.name,
		noround,
		round,
		regU64.name, regI64.name,
		regI64.name,
		regU64.name, tmp.name,
		tmp.name,
		tmp.name, regI64.name,
		cvt, regI64.name, regFloat.name,
		add, regFloat.name, regFloat.name,
		end,
		noround,
		cvt, regU64.name, regFloat.name,
		end)

	return asm
}

func (f *Function) If(instr *ssa.If) (string, *Error) {
	asm := ""
	tblock, fblock := -1, -1
	if instr.Block() != nil && len(instr.Block().Succs) == 2 {
		tblock = instr.Block().Succs[0].Index
		fblock = instr.Block().Succs[1].Index

	}
	if tblock == -1 || fblock == -1 {
		internal("malformed CFG with if stmt")
	}

	if info, ok := f.identifiers[instr.Cond.Name()]; !ok {

		return ErrorMsg(fmt.Sprintf("If: unhandled case, cond (%v)", instr.Cond))
	} else {

		a, err := f.JumpPreamble(instr.Block().Index, fblock)
		if err != nil {
			return "", err
		}
		asm += a
		r, offset, size := info.Addr()
		asm += CmpMemImm32(f.Indent, info.name, int32(offset), &r, uint32(0), size)
		asm += f.Indent + "JEQ    " + "block" + strconv.Itoa(fblock) + "\n"
		a, err = f.JumpPreamble(instr.Block().Index, tblock)
		if err != nil {
			return "", err
		}
		asm += a
		asm += f.Indent + "JMP    " + "block" + strconv.Itoa(tblock) + "\n"

	}

	asm = f.Indent + fmt.Sprintf("// BEGIN ssa.If, %v\n", instr) + asm
	asm += f.Indent + fmt.Sprintf("// END ssa.If, %v\n", instr)

	return asm, nil
}

func (f *Function) JumpPreamble(blockIndex, jmpIndex int) (string, *Error) {
	asm := ""
	phiInfos := f.phiInfo[blockIndex][jmpIndex]
	for _, phiInfo := range phiInfos {
		store := ssa.Store{Addr: phiInfo.phi, Val: phiInfo.value}
		if a, err := f.Store(&store); err != nil {
			return asm, err
		} else {
			asm += a
		}
	}
	return asm, nil
}

func (f *Function) Jump(jmp *ssa.Jump) (string, *Error) {
	asm := ""
	block := -1
	if jmp.Block() != nil && len(jmp.Block().Succs) == 1 {
		block = jmp.Block().Succs[0].Index
	} else {
		internal("malformed CFG with jmp stmt")
	}
	a, err := f.JumpPreamble(jmp.Block().Index, block)
	if err != nil {
		return "", err
	}
	asm += a
	asm += f.Indent + "JMP block" + strconv.Itoa(block) + "\n"
	asm = f.Indent + "// BEGIN ssa.Jump\n" + asm
	asm += f.Indent + "// END ssa.Jump\n"
	return asm, nil
}

func (f *Function) computePhi() *Error {
	for i := 0; i < len(f.ssa.Blocks); i++ {
		if err := f.computeBasicBlockPhi(f.ssa.Blocks[i]); err != nil {
			return err
		}
	}
	return nil
}

func (f *Function) computeBasicBlockPhi(block *ssa.BasicBlock) *Error {
	for i := 0; i < len(block.Instrs); i++ {
		instr := block.Instrs[i]
		switch instr := instr.(type) {
		default:
			break
		case *ssa.Phi:
			if err := f.computePhiInstr(instr); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Function) computePhiInstr(phi *ssa.Phi) *Error {
	blockIndex := phi.Block().Index
	for i, edge := range phi.Edges {
		edgeBlock := -1
		if phi.Block() != nil && i < len(phi.Block().Preds) {
			edgeBlock = phi.Block().Preds[i].Index
		}
		if edgeBlock == -1 {
			internal("malformed CFG")
		}
		if _, ok := f.phiInfo[edgeBlock]; !ok {
			f.phiInfo[edgeBlock] = make(map[int][]phiInfo)
		}
		f.phiInfo[edgeBlock][blockIndex] = append(f.phiInfo[edgeBlock][blockIndex], phiInfo{value: edge, phi: phi})
	}
	return nil
}

func (f *Function) Phi(phi *ssa.Phi) (string, *Error) {

	if nameinfo := f.allocOnDemand(phi); nameinfo == nil {
		return ErrorMsg("Error in ssa.Phi allocation")
	}

	asm := f.Indent
	asm += fmt.Sprintf("// BEGIN ssa.Phi, name (%v), comment (%v), value (%v)\n", phi.Name(), phi.Comment, phi)
	asm += f.Indent + fmt.Sprintf("// END ssa.Phi, %v\n", phi)
	return asm, nil
}

var dummySpSize = uint32(math.MaxUint32)

func (f *Function) Return(ret *ssa.Return) (string, *Error) {
	asm := ResetStackPointer(f.Indent, dummySpSize)
	asm = f.Indent + "// BEGIN ssa.Return\n" + asm
	if a, err := f.CopyToRet(ret.Results); err != nil {
		return "", err
	} else {
		asm += a
	}
	asm += Ret(f.Indent)
	asm += f.Indent + "// END ssa.Return\n"
	return asm, nil
}

func (f *Function) CopyToRet(val []ssa.Value) (string, *Error) {
	if len(val) == 0 {
		return "", nil
	}
	if len(val) > 1 {
		return ErrorMsg("Multiple return values not supported")
	}
	retAddr :=
		identifier{
			name:   retName(),
			typ:    f.retType(),
			local:  nil,
			param:  f.retParam(),
			offset: f.retOffset()}

	return f.StoreValAddr(val[0], &retAddr)
}

func ResetStackPointer(indent string, size uint32) string {
	/*sp := getRegister(REG_SP)
	return AddImm32Reg(indent, size, sp)*/
	return ""
}

func (f *Function) fixupRets(asm string) string {
	old := ResetStackPointer(f.Indent, dummySpSize)
	new := ResetStackPointer(f.Indent, f.localsSize())
	return strings.Replace(asm, old, new, -1)
}

func (f *Function) SetStackPointer() string {
	/*sp := getRegister(REG_SP)
	asm := SubImm32Reg(f.Indent, uint32(f.localsSize()), sp)
	return asm*/
	return ""
}

func (f *Function) StoreValAddr(val ssa.Value, addr *identifier) (string, *Error) {

	if nameinfo := f.allocOnDemand(val); nameinfo == nil {
		internal("error in allocating local")
	}
	if addr.local == nil && addr.param == nil {
		internal(fmt.Sprintf("invalid addr \"%v\"", addr))
	}

	asm := ""
	asm += f.Indent + fmt.Sprintf("// BEGIN StoreValAddr addr name:%v, val name:%v\n", addr.name, val.Name()) + asm

	if isComplex(val.Type()) {
		return ErrorMsg("complex32/64 unsupported")
	}

	if isFloat(val.Type()) {

		valReg := f.allocReg(regType(val.Type()), f.sizeof(val))
		a, err := f.LoadValue(val, 0, f.sizeof(val), &valReg)
		if err != nil {
			return a, err
		}
		asm += a

		a, err = f.StoreValue(addr, &valReg)
		if err != nil {
			return a, err
		}
		asm += a
		f.freeReg(valReg)

	} else {

		size := f.sizeof(val)
		iterations := size
		datasize := 1

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
				internal(fmt.Sprintf("Size (%v) not multiple of sizeInt (%v)", size, sizeInt()))
			}
		}

		valReg := f.allocReg(DATA_REG, uint(datasize))

		for i := int(0); i < int(iterations); i++ {
			offset := i * datasize
			a, err := f.LoadValue(val, offset, uint(datasize), &valReg)
			if err != nil {
				return a, err
			}
			asm += a
			a, err = f.StoreReg(&valReg, addr, offset, uint(datasize))
			if err != nil {
				return a, err
			}
			asm += a
		}
		f.freeReg(valReg)

	}

	asm += f.Indent +
		fmt.Sprintf("// END StoreValAddr addr name:%v, val name:%v\n",
			addr.name, val.Name())
	return asm, nil
}

func (f *Function) Store(instr *ssa.Store) (string, *Error) {
	if nameinfo := f.allocOnDemand(instr.Addr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	addr, ok := f.identifiers[instr.Addr.Name()]
	if !ok {
		internal(fmt.Sprintf("couldnt store identifier \"%v\"", addr.name))
	}
	asm := f.Indent + fmt.Sprintf("// BEGIN Store, instr: %v\n", instr)
	a, err := f.StoreValAddr(instr.Val, &addr)
	asm = asm + a + f.Indent + fmt.Sprintf("// END Store, instr: %v\n", instr)
	return asm, err
}

func (f *Function) BinOp(instr *ssa.BinOp) (string, *Error) {

	if nameinfo := f.allocOnDemand(instr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	var regX, regY *register
	var regVal register
	size := f.sizeof(instr)
	xIsSigned := signed(instr.X.Type())

	// comparison op results are size 1 byte, but that's not supported
	if size == 1 {
		regVal = f.allocReg(regType(instr.Type()), 8*size)
	} else {
		regVal = f.allocReg(regType(instr.Type()), size)
	}

	asm, regX, regY, err := f.BinOpLoadXY(instr)

	if err != nil {
		return asm, err
	}

	size = f.sizeof(instr.X)

	switch instr.Op {
	default:
		internal(fmt.Sprintf("unknown op (%v)", instr.Op))
	case token.ADD, token.SUB, token.MUL, token.QUO, token.REM:
		instrdata := GetOpDataType(instr.Type())
		asm += ArithOp(f.Indent, instrdata, instr.Op, regX, regY, &regVal)
	case token.AND, token.OR, token.XOR, token.SHL, token.SHR, token.AND_NOT:
		asm += BitwiseOp(f.Indent, instr.Op, xIsSigned, regX, regY, &regVal, size)
	case token.EQL, token.NEQ, token.LEQ, token.GEQ, token.LSS, token.GTR:
		if size != f.sizeof(instr.Y) {
			internal("comparing two different size values")
		}
		instrdata := GetOpDataType(instr.X.Type())
		asm += CmpOp(f.Indent, instrdata, instr.Op, regX, regY, &regVal)
	}

	f.freeReg(*regX)
	f.freeReg(*regY)

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		internal(fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr))
	}

	a, err := f.StoreValue(&addr, &regVal)
	if err != nil {
		return asm, err
	} else {
		asm += a
	}
	f.freeReg(regVal)

	asm = fmt.Sprintf(f.Indent+"// BEGIN ssa.BinOp, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf(f.Indent+"// END ssa.BinOp, %v = %v\n", instr.Name(), instr)
	return asm, nil
}

func (f *Function) BinOpLoadXY(instr *ssa.BinOp) (asm string, x *register, y *register, err *Error) {

	if nameinfo := f.allocOnDemand(instr); nameinfo == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	if nameinfo := f.allocOnDemand(instr.X); nameinfo == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr.X))
	}
	if nameinfo := f.allocOnDemand(instr.Y); nameinfo == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr.Y))
	}

	xtmp := f.allocReg(regType(instr.X.Type()), f.sizeof(instr.X))
	x = &xtmp
	ytmp := f.allocReg(regType(instr.Y.Type()), f.sizeof(instr.Y))
	y = &ytmp
	asm = f.Indent + "// BEGIN BinOpLoadXY\n"

	if a, err := f.LoadValue(instr.X, 0, f.sizeof(instr.X), x); err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}

	if a, err := f.LoadValue(instr.Y, 0, f.sizeof(instr.Y), y); err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}

	asm += f.Indent + "// END BinOpLoadXY\n"
	return asm, x, y, nil
}

func (f *Function) sizeof(val ssa.Value) uint {
	if _, ok := val.(*ssa.Const); ok {
		return f.sizeofConst(val.(*ssa.Const))
	}
	info, ok := f.identifiers[val.Name()]
	if !ok {
		internal(fmt.Sprintf("unknown name (%v), value (%v)\n", val.Name(), val))
	}
	_, _, size := info.Addr()
	return size
}

func (f *Function) sizeofConst(cnst *ssa.Const) uint {
	return sizeof(cnst.Type())
}

func (f *Function) LoadValueSimple(val ssa.Value, reg *register) (string, *Error) {
	asm := f.Indent + fmt.Sprintf("// BEGIN LoadValueSimple, val: %v, reg %v\n", val, reg.name)
	a, err := f.LoadValue(val, 0, f.sizeof(val), reg)
	asm += a
	asm += f.Indent + fmt.Sprintf("// END LoadValueSimple, val: %v, reg %v\n", val, reg.name)
	return asm, err
}

func (f *Function) LoadValue(val ssa.Value, offset int, size uint, reg *register) (string, *Error) {
	if _, ok := val.(*ssa.Const); ok {
		return f.LoadConstValue(val.(*ssa.Const), reg)
	}
	ident, ok := f.identifiers[val.Name()]
	if !ok {
		internal(fmt.Sprintf("unknown name (%v), value (%v)\n", val.Name(), val))
	}

	r, roffset, rsize := ident.Addr()
	if rsize%size != 0 {
		msg := "ident (%v), size (%v) not a multiple of chunk size (%v) in loading"
		internal(fmt.Sprintf(msg, ident.name, ident.size(), size))
	}
	if size > 8 {
		msg := "ident (%v), loading more than 8 byte chunk"
		internal(fmt.Sprintf(msg, ident.name))
	}

	datatype := GetIntegerOpDataType(false, size)
	asm := f.Indent + fmt.Sprintf("// BEGIN LoadValue, val: %v, offset %v, size %v\n", val, offset, size)
	a := MovMemReg(f.Indent, datatype, ident.name, roffset+offset, &r, reg)
	asm += a
	asm += f.Indent + fmt.Sprintf("// END LoadValue, val: %v, offset %v, size %v\n", val, offset, size)
	return asm, nil
}

func (f *Function) StoreValue(ident *identifier, reg *register) (string, *Error) {
	if ident.size() > reg.size() {
		msgstr := "identifier, %v, size (%v) is greater than register size (%v)"
		internal(fmt.Sprintf(msgstr, ident.name, ident.size(), reg.size()))

	}
	return f.StoreReg(reg, ident, 0, ident.size())
}

func (f *Function) StoreReg(reg *register, ident *identifier, offset int, size uint) (string, *Error) {
	r, roffset, rsize := ident.Addr()
	if rsize%size != 0 {
		internal(fmt.Sprintf("storing identifier \"%v\"", ident.name))
	}
	if rsize == 0 {
		internal(fmt.Sprintf("identifier (%v) size is 0", ident.name))
	}
	asm := f.Indent + fmt.Sprintf("// BEGIN StoreReg, size (%v)\n", size)
	instrdata := GetIntegerOpDataType(false, size)
	asm += MovRegMem(f.Indent, instrdata, reg, ident.name, &r, offset+roffset)
	asm += f.Indent + fmt.Sprintf("// END StoreReg, size (%v)\n", size)
	return asm, nil
}

func (f *Function) LoadConstValue(cnst *ssa.Const, r *register) (string, *Error) {

	if isBool(cnst.Type()) {
		var val int8
		if cnst.Value.String() == "true" {
			val = 1
		}
		return MovImm8Reg(f.Indent, val, r), nil
	}
	if isFloat(cnst.Type()) {
		if r.typ != XMM_REG {
			internal("can't load float const into non xmm register")
		}
		tmp := f.allocReg(DATA_REG, 8)
		asm := ""
		if isFloat32(cnst.Type()) {
			asm = MovImmf32Reg(f.Indent, float32(cnst.Float64()), &tmp, r)
		} else {
			asm = MovImmf64Reg(f.Indent, cnst.Float64(), &tmp, r)
		}
		f.freeReg(tmp)
		return asm, nil

	}
	if isComplex(cnst.Type()) {
		internal("complex64/128 unsupported")
	}

	size := sizeof(cnst.Type())
	signed := signed(cnst.Type())
	var val int64
	if signed {
		val = cnst.Int64()
	} else {

		val = int64(cnst.Uint64())
	}
	return MovImmReg(f.Indent, val, size, r), nil
}

func (f *Function) UnOp(instr *ssa.UnOp) (string, *Error) {
	var err *Error
	asm := ""
	switch instr.Op {
	default:
		internal(fmt.Sprintf("unknown Op token (%v): \"%v\"", instr.Op, instr))
	case token.NOT: // logical negation
		asm, err = f.UnOpXor(instr, 1)
	case token.XOR: //bitwise negation
		asm, err = f.UnOpXor(instr, -1)
	case token.SUB: // arithmetic negation e.g. x=>-x
		asm, err = f.UnOpSub(instr)
	case token.MUL: //pointer indirection
		asm, err = f.UnOpPointer(instr)
	}
	asm = f.Indent + fmt.Sprintf("// BEGIN ssa.UnOp: %v = %v\n", instr.Name(), instr) + asm
	asm += f.Indent + fmt.Sprintf("// END ssa.UnOp: %v = %v\n", instr.Name(), instr)
	return asm, err

}

// bitwise negation
func (f *Function) UnOpXor(instr *ssa.UnOp, xorVal int32) (string, *Error) {

	if nameinfo := f.allocOnDemand(instr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	size := f.sizeof(instr)
	reg := f.allocReg(regType(instr.X.Type()), size)

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		msg := fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr)
		internal(msg)
	}

	asm := ZeroReg(f.Indent, &reg)

	asm, err := f.LoadValueSimple(instr.X, &reg)
	if err != nil {
		return asm, err
	}

	if size < 8 {
		asm += XorImm32Reg(f.Indent, xorVal, &reg, size)
	} else {
		asm += XorImm64Reg(f.Indent, int64(xorVal), &reg, size)
	}

	a, err := f.StoreValue(&addr, &reg)
	f.freeReg(reg)

	if err != nil {
		return asm, err
	} else {
		asm += a
	}

	asm = fmt.Sprintf(f.Indent+"// BEGIN ssa.UnOpNot, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf(f.Indent+"// END ssa.UnOpNot, %v = %v\n", instr.Name(), instr)
	return asm, nil
}

// arithmetic negation
func (f *Function) UnOpSub(instr *ssa.UnOp) (string, *Error) {

	if nameinfo := f.allocOnDemand(instr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	var regX register
	var regSubX register
	var regVal register

	regVal = f.allocReg(regType(instr.Type()), f.sizeof(instr))
	regX = f.allocReg(regType(instr.X.Type()), f.sizeof(instr.X))
	regSubX = f.allocReg(regType(instr.X.Type()), f.sizeof(instr.X))

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		msg := fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr)
		internal(msg)
	}

	asm, err := f.LoadValueSimple(instr.X, &regX)
	if err != nil {
		return asm, err
	}

	asm += ZeroReg(f.Indent, &regSubX)
	instrdata := GetOpDataType(instr.Type())
	asm += ArithOp(f.Indent, instrdata, token.SUB, &regSubX, &regX, &regVal)
	f.freeReg(regX)
	f.freeReg(regSubX)

	a, err := f.StoreValue(&addr, &regVal)
	if err != nil {
		return asm, err
	} else {
		asm += a
	}
	f.freeReg(regVal)

	asm = fmt.Sprintf(f.Indent+"// BEGIN ssa.UnOpSub, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf(f.Indent+"// END ssa.UnOpSub, %v = %v\n", instr.Name(), instr)
	return asm, nil

}

//pointer indirection
func (f *Function) UnOpPointer(instr *ssa.UnOp) (string, *Error) {
	assignment := f.allocOnDemand(instr)
	if assignment == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	xName := instr.X.Name()
	xInfo, okX := f.identifiers[xName]

	// TODO add complex64/128 support
	if isComplex(instr.Type()) || isComplex(instr.X.Type()) {
		return ErrorMsg("complex64/complex128 unimplemented")
	}
	if !okX {
		msgstr := "Unknown name for UnOp X (%v), instr \"(%v)\""
		internal(fmt.Sprintf(msgstr, instr.X, instr))
	}
	if xInfo.local == nil && xInfo.param == nil && !xInfo.IsPointer() {
		fmtstr := "in UnOp, X (%v) isn't a pointer, X.type (%v), instr \"(%v)\""
		msg := fmt.Sprintf(fmtstr, instr.X, instr.X.Type(), instr)
		internal(msg)
	}

	asm := ""

	xReg, xOffset, xSize := xInfo.Addr()
	aReg, aOffset, aSize := assignment.Addr()

	if xSize != sizePtr() {
		fmt.Printf("instr: %v\n", instr)
		internal(fmt.Sprintf("xSize (%v) != ptr size (%v)", xSize, sizePtr()))
	}

	size := aSize

	tmp1 := f.allocReg(DATA_REG, DataRegSize)
	tmp2 := f.allocReg(regType(instr.Type()), DataRegSize)
	instrdata := GetOpDataType(instr.Type())

	asm += MovMemIndirectMem(f.Indent, instrdata, xInfo.name, xOffset, &xReg, assignment.name, aOffset, &aReg, size, &tmp1, &tmp2)

	f.identifiers[assignment.name] = *assignment

	f.freeReg(tmp1)
	f.freeReg(tmp2)

	return asm, nil
}

func (f *Function) Index(instr *ssa.Index) (string, *Error) {
	if instr == nil {
		return ErrorMsg("nil instr")
	}
	asm := ""
	xInfo := f.identifiers[instr.X.Name()]
	assignment := f.allocOnDemand(instr)
	if assignment == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	xReg, xOffset, _ := xInfo.Addr()
	aReg, aOffset, _ := assignment.Addr()
	addrReg := f.allocReg(DATA_REG, sizePtr())
	idxReg := f.allocReg(DATA_REG, sizePtr())

	f.LoadValueSimple(instr.Index, &idxReg)

	asm += Lea(f.Indent, xInfo.name, xOffset, &xReg, &addrReg)

	instrdata := GetIntegerOpDataType(false, idxReg.width/8)
	asm += AddRegReg(f.Indent, instrdata, &idxReg, &addrReg)

	instrdata = GetOpDataType(assignment.typ)
	asm += MovRegMem(f.Indent, instrdata, &addrReg, assignment.name, &aReg, aOffset)

	f.freeReg(idxReg)
	f.freeReg(addrReg)

	f.identifiers[instr.Name()] = *assignment

	asm = f.Indent + fmt.Sprintf("// BEGIN ssa.Index: %v = %v\n", instr.Name(), instr) + asm
	asm += f.Indent + fmt.Sprintf("// END ssa.Index: %v = %v\n", instr.Name(), instr)

	return asm, nil
}

func (f *Function) IndexAddr(instr *ssa.IndexAddr) (string, *Error) {
	if instr == nil {
		return ErrorMsg("nil instr")

	}

	asm := ""
	xInfo := f.identifiers[instr.X.Name()]
	assignment := f.allocOnDemand(instr)

	xReg, xOffset, _ := xInfo.Addr()
	aReg, aOffset, _ := assignment.Addr()
	addrReg := f.allocReg(DATA_REG, sizePtr())
	idxReg := f.allocReg(DATA_REG, sizePtr())

	if a, err := f.LoadValueSimple(instr.Index, &idxReg); err == nil {
		asm += a
	} else {
		return "", err
	}
	asm += MulImm32RegReg(f.Indent, uint32(sizeofElem(xInfo.typ)), &idxReg, &idxReg)

	if xInfo.IsSlice() {
		// TODO: add bounds checking
		optypes := GetIntegerOpDataType(false, sizePtr())
		asm += MovMemReg(f.Indent, optypes, xInfo.name, xOffset, &xReg, &addrReg)
	} else if xInfo.IsArray() {
		asm += Lea(f.Indent, xInfo.name, xOffset, &xReg, &addrReg)
	} else {
		internal("indexing non-slice/array variable")
	}

	optypes := GetIntegerOpDataType(false, idxReg.size())

	asm += AddRegReg(f.Indent, optypes, &idxReg, &addrReg)

	optypes = GetIntegerOpDataType(false, assignment.size())
	asm += MovRegMem(f.Indent, optypes, &addrReg, assignment.name, &aReg, aOffset)

	f.freeReg(idxReg)
	f.freeReg(addrReg)

	f.identifiers[instr.Name()] = *assignment

	asm = f.Indent + fmt.Sprintf("// BEGIN ssa.IndexAddr: %v = %v\n", instr.Name(), instr) + asm
	asm += f.Indent + fmt.Sprintf("// END ssa.IndexAddr: %v = %v\n", instr.Name(), instr)
	return asm, nil
}

func (f *Function) AllocInstr(instr *ssa.Alloc) (string, *Error) {
	asm := ""
	if instr == nil {
		return ErrorMsg("AllocInstr: nil instr")

	}
	if instr.Heap {
		return ErrorMsg("AllocInstr: heap alloc")
	}

	//Alloc values are always addresses, and have pointer types, so the type
	//of the allocated variable is actually
	//Type().Underlying().(*types.Pointer).Elem().
	info := f.identifiers[instr.Name()]
	if info.local == nil {
		internal(fmt.Sprintf("expected %v to be a local variable", instr.Name()))
	}
	if _, ok := info.typ.(*types.Pointer); ok {
	} else {
	}
	f.identifiers[instr.Name()] = info
	return asm, nil
}

func (f *Function) Value(value ssa.Value, dstReg *register, dstVar *varInfo) string {
	if dstReg == nil && dstVar == nil {
		internal("destination register/var are nil!")
	}
	if dstReg != nil && dstVar != nil {
		internal("destination register/var are both non-nil!")
	}
	if dstReg != nil {
		// TODO
	}
	if dstVar != nil {
		// TODO
	}
	return ""
}

func (f *Function) localsSize() uint32 {
	size := uint32(0)
	for _, name := range f.identifiers {
		if name.local != nil {
			size += uint32(name.size())
		}
	}
	return size
}

func (f *Function) init() *Error {
	f.registers = make(map[string]bool)
	f.identifiers = make(map[string]identifier)
	f.phiInfo = make(map[int]map[int][]phiInfo)
	f.initRegs()
	return nil
}

func (f *Function) initRegs() {
	for _, r := range registers {
		f.registers[r.name] = false
	}
}

// size in bytes
func (f *Function) allocReg(t RegType, size uint) register {
	var reg register
	found := false
	for i := 0; i < len(registers); i++ {
		r := registers[i]
		if f.excludeReg(&r) {
			continue
		}
		used := f.registers[r.name]
		if used || r.typ != t {
			continue
		}
		// r.width is in bits so multiple size (which is in bytes) by 8
		for i := range r.datasizes {
			if r.datasizes[i] == size {
				reg = r
				found = true
				break
			}
		}
	}
	if found {
		f.registers[reg.name] = true
	} else {
		// any of the data registers can be used as an address register on x86_64
		if t == ADDR_REG {
			return f.allocReg(DATA_REG, size)
		} else {
			msgstr := "couldn't alloc register (t: %v, w: %v, s: %v)"
			internal(fmt.Sprintf(msgstr, t, size*8, size))
		}
	}
	return reg
}

func (f *Function) excludeReg(reg *register) bool {
	for _, r := range excludedRegisters {
		if r.name == reg.name {
			return true
		}
	}
	return false
}

// zeroReg returns the assembly for zeroing the passed in register
func (f *Function) zeroReg(r *register) string {
	return ZeroReg(f.Indent, r)
}

func (f *Function) freeReg(reg register) {
	f.registers[reg.name] = false
}

// paramsSize returns the size of the parameters in bytes
func (f *Function) paramsSize() uint {
	size := uint(0)
	for _, p := range f.ssa.Params {
		size += sizeof(p.Type())
	}
	return size
}

func retName() string {
	return "ret0"
}

// retType gives the return type
func (f *Function) retType() types.Type {
	results := f.ssa.Signature.Results()
	if results.Len() == 0 {
		return nil
	}
	if results.Len() > 1 {
		panic("Functions with more than one return value not supported")
	}
	return results.At(0).Type()
}

func (f *Function) retParam() *paramInfo {
	return &paramInfo{name: retName(), info: nil, extra: nil}
}

// retSize returns the size of the return value in bytes
func (f *Function) retSize() uint {
	size := sizeof(f.retType())
	return size
}

// retOffset returns the offset of the return value in bytes
func (f *Function) retOffset() int {
	align := f.retAlign()
	padding := align - f.paramsSize()%align
	if padding == align {
		padding = 0
	}
	return int(f.paramsSize() + padding)
}

// retAlign returns the byte alignment alignment for the return value
func (f *Function) retAlign() uint {
	align := align(f.retType())
	// TODO: fix, why always 8 bytes with go compiler?
	if align < 8 {
		align = 8
	}
	return align
}

func (f *Function) allocOnDemand(v ssa.Value) *identifier {

	if nameinfo, ok := f.identifiers[v.Name()]; ok {
		return &nameinfo
	}

	switch v := v.(type) {
	case *ssa.Const:
		nameinfo := identifier{name: v.Name(), typ: v.Type(), local: nil, param: nil, cnst: v}
		f.identifiers[v.Name()] = nameinfo
		return &nameinfo
	}

	local, err := f.AllocLocal(v.Name(), v.Type())
	if err != nil {
		return nil
	}

	f.identifiers[v.Name()] = local

	return &local
}

func internal(msg string) {
	panic(fmt.Sprintf("Internal error, \"%v\"", msg))
}
