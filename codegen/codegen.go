package codegen

import (
	"errors"
	"fmt"
	"go/token"
	"math"
	"strconv"
	"strings"

	"golang.org/x/tools/go/types"

	"reflect"

	"github.com/bjwbell/gensimd/simd"

	"golang.org/x/tools/go/ssa"
)

type phiInfo struct {
	value ssa.Value
	phi   *ssa.Phi
}

type Function struct {
	// output function name
	outfn     string
	Indent    string
	ssa       *ssa.Function
	registers map[string]bool // maps register to false if unused and true if used
	ssaNames  map[string]nameInfo
	// map from block index to the successor block indexes that need phi vars set
	phiInfo map[int]map[int][]phiInfo
}

type nameInfo struct {
	name   string
	typ    types.Type
	local  *varInfo
	param  *paramInfo
	offset int
	size   uint
	align  uint
}

// Addr returns the register and offset to access the backing memory of name.
// For locals the register is the stack pointer (SP) and for params the register
// is the frame pointer (FP).
func (name *nameInfo) Addr() (reg register, offset int, size uint) {
	offset = name.offset
	size = name.size
	if name.local != nil {
		reg = *getRegister(REG_SP)
	} else if name.param != nil {
		reg = *getRegister(REG_FP)
	} else {
		panic(fmt.Sprintf("nameInfo (%v) is not a local or param", name))
	}
	return
}

func (name *nameInfo) IsSsaLocal() bool {
	return name.local != nil && name.local.info != nil
}

func (name *nameInfo) IsPointer() bool {
	_, ok := name.typ.(*types.Pointer)
	return ok
}

func (name *nameInfo) PointerUnderlyingType() types.Type {
	if !name.IsPointer() {
		panic(fmt.Sprintf("nameInfo (%v) not pointer type in PointerUnderlyingType", name))
	}
	ptrType := name.typ.(*types.Pointer)
	return ptrType.Elem()
}

func (name *nameInfo) IsArray() bool {
	_, ok := name.typ.(*types.Array)
	return ok
}

func (name *nameInfo) IsSlice() bool {
	_, ok := name.typ.(*types.Slice)
	return ok
}

func (name *nameInfo) IsBasic() bool {
	_, ok := name.typ.(*types.Basic)
	return ok
}

func (name *nameInfo) IsInteger() bool {
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

func CreateFunction(fn *ssa.Function, outfn string) (*Function, *Error) {
	if fn == nil {
		return nil, &Error{Err: errors.New("Nil function passed in")}
	}
	f := Function{ssa: fn, outfn: outfn}
	f.Indent = "        "
	f.init()
	return &f, nil
}

func (f *Function) GoAssembly() (string, *Error) {
	return f.asmFunc()
}

func memFn(name string, offset uint, regName string) func() string {
	return func() string {
		return fmt.Sprintf("%v+%v(%v)", name, offset, regName)
	}
}

func regFn(name string) func() string {
	return func() string {
		return name
	}
}

func (f *Function) asmParams() (string, *Error) {
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
				return "", &Error{Err: fmt.Errorf("Unsupported param type (%v)", basic), Pos: p.Pos()}
			case types.Float32, types.Float64:
				break

				// supported param types
			case types.Bool, types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
				types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
				break
			}

		} else {

		}
		info := nameInfo{name: param.name, typ: param.info.Type(),
			local: nil, param: &param, offset: offset, size: sizeof(p.Type()), align: align(p.Type())}
		f.ssaNames[param.name] = info
		if info.align > info.size {
			offset += int(info.align)
		} else {
			offset += int(info.size)
		}
	}
	return asm, nil
}

func (f *Function) asmFunc() (string, *Error) {

	params, err := f.asmParams()
	if err != nil {
		return params, err
	}

	zeroRetValue, err := f.asmZeroRetValue()
	if err != nil {
		return params + zeroRetValue, err
	}

	zeroSsaLocals, err := f.asmZeroSsaLocals()
	if err != nil {
		return params + zeroRetValue + zeroSsaLocals, err
	}

	if err := f.computePhi(); err != nil {
		return "", err
	}

	basicblocks, err := f.asmBasicBlocks()
	if err != nil {
		return params + zeroRetValue + zeroSsaLocals + basicblocks, err
	}

	zeroNonSsaLocals, err := f.asmZeroNonSsaLocals()
	if err != nil {
		return zeroNonSsaLocals, err
	}

	frameSize := f.localsSize()
	frameSize = f.align(frameSize)
	argsSize := f.retOffset() + int(f.retAlign())
	asm := params
	asm += f.asmSetStackPointer()
	asm += zeroRetValue
	asm += zeroSsaLocals
	asm += zeroNonSsaLocals
	asm += basicblocks
	asm = f.fixupRets(asm)
	a := fmt.Sprintf("TEXT ·%v(SB),NOSPLIT,$%v-%v\n%v", f.outfname(), frameSize, argsSize, asm)
	return a, nil
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

func (f *Function) asmZeroSsaLocals() (string, *Error) {
	asm := ""
	offset := int(0)
	locals := f.ssa.Locals
	for _, local := range locals {
		if local.Heap {

			msg := fmt.Errorf("Can't heap alloc local, name: %v", local.Name())
			return "", &Error{Err: msg, Pos: local.Pos()}
		}
		sp := getRegister(REG_SP)

		//local values are always addresses, and have pointer types, so the type
		//of the allocated variable is actually
		//Type().Underlying().(*types.Pointer).Elem().
		typ := local.Type().Underlying().(*types.Pointer).Elem()
		size := sizeof(typ)
		asm += ZeroMemory(f.Indent, local.Name(), offset, size, sp)
		v := varInfo{name: local.Name(), info: local}
		info := nameInfo{name: v.name, typ: typ, local: &v, param: nil, offset: offset, size: size, align: align(typ)}
		f.ssaNames[v.name] = info
		if info.align > info.size {
			offset += int(info.align)
		} else {
			offset += int(size)
		}
	}
	return asm, nil
}

func (f *Function) asmAllocLocal(name string, typ types.Type) (nameInfo, *Error) {
	size := sizeof(typ)
	offset := int(size)
	if align(typ) > size {
		offset = int(align(typ))
	}
	v := varInfo{name: name, info: nil}
	info := nameInfo{
		name:   name,
		typ:    typ,
		param:  nil,
		local:  &v,
		offset: -int(f.localsSize()) - offset,
		size:   size,
		align:  align(typ)}
	f.ssaNames[v.name] = info
	// zeroing the memory is done at the beginning of the function
	//asmZeroMemory(f.Indent, v.name, v.offset, v.size, sp)
	return info, nil
}

func (f *Function) asmZeroNonSsaLocals() (string, *Error) {
	asm := ""
	for _, name := range f.ssaNames {
		if name.local == nil || name.IsSsaLocal() {
			continue
		}
		sp := getRegister(REG_SP)
		asm += ZeroMemory(f.Indent, name.name, name.offset, name.size, sp)
	}
	return asm, nil
}

func (f *Function) asmZeroRetValue() (string, *Error) {
	asm := f.Indent + "// BEGIN asmZeroRetValue\n"
	asm += ZeroMemory(f.Indent, retName(), f.retOffset(), f.retSize(), getRegister(REG_FP))
	asm += f.Indent + "// END asmZeroRetValue\n"
	return asm, nil
}

func (f *Function) asmBasicBlocks() (string, *Error) {
	asm := ""
	for i := 0; i < len(f.ssa.Blocks); i++ {
		a, err := f.asmBasicBlock(f.ssa.Blocks[i])
		asm += a
		if err != nil {
			return asm, err
		}
	}
	return asm, nil
}

func (f *Function) asmBasicBlock(block *ssa.BasicBlock) (string, *Error) {
	asm := "block" + strconv.Itoa(block.Index) + ":\n"
	for i := 0; i < len(block.Instrs); i++ {
		a, err := f.asmInstr(block.Instrs[i])
		asm += a
		if err != nil {
			return asm, err
		}

	}
	return asm, nil
}

func (f *Function) asmInstr(instr ssa.Instruction) (string, *Error) {

	if instr == nil {
		panic("Nil instr")
	}
	asm := ""
	caseAsm := ""
	var caseErr *Error
	errormsg := func(msg string) (string, *Error) {
		return "", &Error{Err: fmt.Errorf(msg), Pos: instr.Pos()}
	}
	switch instr := instr.(type) {
	default:
		caseAsm = f.Indent + fmt.Sprintf("Unknown ssa instruction: %v\n", instr)
	case *ssa.Alloc:
		caseAsm, caseErr = f.asmAllocInstr(instr)
	case *ssa.BinOp:
		caseAsm, caseErr = f.asmBinOp(instr)
	case *ssa.Call:
		caseAsm = f.Indent + fmt.Sprintf("ssa.Call: %v, name: %v\n", instr, instr.Name())
	case *ssa.ChangeInterface:
		caseAsm, caseErr = errormsg("converting interfaces unsupported")
	case *ssa.ChangeType:
		caseAsm, caseErr = errormsg("changing between types unsupported")
	case *ssa.Convert:
		caseAsm, caseErr = errormsg("type conversion unimplemented")
	case *ssa.Defer:
		caseAsm, caseErr = errormsg("defer unsupported")
	case *ssa.Extract:
		caseAsm, caseErr = errormsg("extracting tuple values unsupported")
	case *ssa.Field:
		caseAsm, caseErr = errormsg("field access unimplemented")
	case *ssa.FieldAddr:
		caseAsm, caseErr = errormsg("field access unimplemented")
	case *ssa.Go:
		caseAsm, caseErr = errormsg("go keyword unsupported")
	case *ssa.If:
		caseAsm, caseErr = f.asmIf(instr)
	case *ssa.Index:
		caseAsm, caseErr = errormsg("index access unimplemented")
	case *ssa.IndexAddr:
		caseAsm, caseErr = f.asmIndexAddr(instr)
	case *ssa.Jump:
		caseAsm, caseErr = f.asmJump(instr)
	case *ssa.Lookup:
		caseAsm, caseErr = errormsg("maps unsupported")
	case *ssa.MakeChan:
		caseAsm, caseErr = errormsg("channels unsupported")
	case *ssa.MakeClosure:
		caseAsm, caseErr = errormsg("closures unsupported")
	case *ssa.MakeInterface, *ssa.MakeMap, *ssa.MakeSlice:
		caseAsm, caseErr = errormsg("make slice/map/interface unsupported")
	case *ssa.MapUpdate:
		caseAsm, caseErr = errormsg("map update unsupported")
	case *ssa.Next:
		caseAsm, caseErr = errormsg("map/string iterators unsupported")
	case *ssa.Panic:
		caseAsm, caseErr = errormsg("panic unimplemented")
	case *ssa.Phi:
		caseAsm, caseErr = f.asmPhi(instr)
	case *ssa.Range:
		caseAsm, caseErr = errormsg("range unsupported")
	case *ssa.Return:
		caseAsm, caseErr = f.asmReturn(instr)
	case *ssa.Select, *ssa.RunDefers, *ssa.Send:
		caseAsm, caseErr = errormsg("select/send/defer unsupported")
	case *ssa.Slice:
		caseAsm, caseErr = errormsg("slice creation unimplemented")
	case *ssa.Store:
		caseAsm, caseErr = f.asmStore(instr)
	case *ssa.TypeAssert:
		caseAsm, caseErr = errormsg("type assert unsupported")
	case *ssa.UnOp:
		caseAsm, caseErr = f.asmUnOp(instr)
	}

	if caseErr != nil {
		return caseAsm, caseErr
	} else {
		asm += caseAsm
	}

	return asm, nil
}

func (f *Function) asmIf(instr *ssa.If) (string, *Error) {
	asm := ""
	tblock, fblock := -1, -1
	if instr.Block() != nil && len(instr.Block().Succs) == 2 {
		tblock = instr.Block().Succs[0].Index
		fblock = instr.Block().Succs[1].Index

	}
	if tblock == -1 || fblock == -1 {
		panic("asmIf: malformed CFG")
	}
	if info, ok := f.ssaNames[instr.Cond.Name()]; !ok {
		err := fmt.Errorf("asmIf: unhandled case, cond (%v)", instr.Cond)
		return "", &Error{Err: err, Pos: instr.Pos()}
	} else {
		a, err := f.asmJumpPreamble(instr.Block().Index, fblock)
		if err != nil {
			return "", err
		}
		asm += a
		r, offset, size := info.Addr()
		asm += CmpMemImm32(f.Indent, info.name, int32(offset), &r, uint32(0), size)
		asm += f.Indent + "JEQ    " + "block" + strconv.Itoa(fblock) + "\n"
		a, err = f.asmJumpPreamble(instr.Block().Index, tblock)
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

func (f *Function) asmJumpPreamble(blockIndex, jmpIndex int) (string, *Error) {
	asm := ""
	phiInfos := f.phiInfo[blockIndex][jmpIndex]
	for _, phiInfo := range phiInfos {
		store := ssa.Store{Addr: phiInfo.phi, Val: phiInfo.value}
		if a, err := f.asmStore(&store); err != nil {
			return asm, err
		} else {
			asm += a
		}
	}
	return asm, nil
}

func (f *Function) asmJump(jmp *ssa.Jump) (string, *Error) {
	asm := ""
	block := -1
	if jmp.Block() != nil && len(jmp.Block().Succs) == 1 {
		block = jmp.Block().Succs[0].Index
	} else {
		panic("asmJump: malformed CFG")
	}
	a, err := f.asmJumpPreamble(jmp.Block().Index, block)
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
			panic("computePhiInstr: malformed CFG")
		}
		if _, ok := f.phiInfo[edgeBlock]; !ok {
			f.phiInfo[edgeBlock] = make(map[int][]phiInfo)
		}
		f.phiInfo[edgeBlock][blockIndex] = append(f.phiInfo[edgeBlock][blockIndex], phiInfo{value: edge, phi: phi})
	}
	return nil
}

func (f *Function) asmPhi(phi *ssa.Phi) (string, *Error) {
	if err := f.allocValueOnDemand(phi); err != nil {
		return "", err
	}
	asm := f.Indent
	asm += fmt.Sprintf("// BEGIN ssa.Phi, name (%v), comment (%v), value (%v)\n", phi.Name(), phi.Comment, phi)
	asm += f.Indent + fmt.Sprintf("// END ssa.Phi, %v\n", phi)
	return asm, nil
}

var dummySpSize = uint32(math.MaxUint32)

func (f *Function) asmReturn(ret *ssa.Return) (string, *Error) {
	asm := asmResetStackPointer(f.Indent, dummySpSize)
	asm = f.Indent + "// BEGIN ssa.Return\n" + asm
	if a, err := f.asmCopyToRet(ret.Results); err != nil {
		return "", err
	} else {
		asm += a
	}
	asm += Ret(f.Indent)
	asm += f.Indent + "// END ssa.Return\n"
	return asm, nil
}

func (f *Function) asmCopyToRet(val []ssa.Value) (string, *Error) {
	if len(val) == 0 {
		return "", nil
	}
	if len(val) > 1 {
		err := Error{
			Err: fmt.Errorf("Multiple return values not supported"),
			Pos: 0}
		return "", &err
	}
	retAddr := nameInfo{name: retName(), typ: f.retType(), local: nil, param: f.retParam(), size: f.retSize(), offset: f.retOffset(), align: f.retAlign()}
	return f.asmStoreValAddr(val[0], &retAddr)
}

func asmResetStackPointer(indent string, size uint32) string {
	/*sp := getRegister(REG_SP)
	return asmAddImm32Reg(indent, size, sp)*/
	return ""
}

func (f *Function) fixupRets(asm string) string {
	old := asmResetStackPointer(f.Indent, dummySpSize)
	new := asmResetStackPointer(f.Indent, f.localsSize())
	return strings.Replace(asm, old, new, -1)
}

func (f *Function) asmSetStackPointer() string {
	/*sp := getRegister(REG_SP)
	asm := asmSubImm32Reg(f.Indent, uint32(f.localsSize()), sp)
	return asm*/
	return ""
}

func (f *Function) asmStoreValAddr(val ssa.Value, addr *nameInfo) (string, *Error) {

	var err *Error
	if err = f.allocValueOnDemand(val); err != nil {
		return "", err
	}
	if addr.local == nil && addr.param == nil {
		msg := fmt.Errorf("Invalid addr \"%v\"", addr)
		return "", &Error{Err: msg, Pos: 0}
	}

	asm := ""
	asm += f.Indent + fmt.Sprintf("// BEGIN asmStoreValAddr addr name:%v, val name:%v\n", addr.name, val.Name()) + asm

	if isComplex(val.Type()) {
		return "", &Error{fmt.Errorf("complex32/64 unsupported"), val.Pos()}
	}

	if isFloat(val.Type()) {

		valReg := f.allocReg(regType(val.Type()), f.sizeof(val))
		a, err := f.asmLoadValue(val, 0, f.sizeof(val), &valReg)
		if err != nil {
			return a, err
		}
		asm += a

		a, err = f.asmStoreReg(&valReg, addr, 0)
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
				panic(fmt.Sprintf("Size (%v) not multiple of sizeInt (%v)", size, sizeInt()))
			}
		}

		valReg := f.allocReg(DATA_REG, uint(datasize))

		for i := int(0); i < int(iterations); i++ {
			offset := i * datasize
			a, err := f.asmLoadValue(val, offset, uint(datasize), &valReg)
			if err != nil {
				return a, err
			}
			asm += a
			a, err = f.asmStoreReg(&valReg, addr, offset)
			if err != nil {
				return a, err
			}
			asm += a
		}
		f.freeReg(valReg)

	}

	asm += f.Indent +
		fmt.Sprintf("// END asmStoreValAddr addr name:%v, val name:%v\n",
			addr.name, val.Name())
	return asm, nil
}

func (f *Function) asmStore(instr *ssa.Store) (string, *Error) {
	if err := f.allocValueOnDemand(instr.Addr); err != nil {
		return "", err
	}
	addr, ok := f.ssaNames[instr.Addr.Name()]
	if !ok {
		panic("Couldnt find instr.Addr in ssaNames")
	}
	return f.asmStoreValAddr(instr.Val, &addr)
}

func (f *Function) asmBinOp(instr *ssa.BinOp) (string, *Error) {
	if err := f.allocValueOnDemand(instr); err != nil {
		return "", err
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
	asm, regX, regY, err := f.asmBinOpLoadXY(instr)
	if err != nil {
		return asm, err
	}

	size = f.sizeof(instr.X)

	switch instr.Op {
	default:
		panic(fmt.Sprintf("Unknown op (%v)", instr.Op))
	case token.ADD, token.SUB, token.MUL, token.QUO, token.REM:
		instrdata := GetInstrDataType(instr.Type())
		asm += ArithOp(f.Indent, instrdata, instr.Op, regX, regY, &regVal)
	case token.AND, token.OR, token.XOR, token.SHL, token.SHR, token.AND_NOT:
		asm += BitwiseOp(f.Indent, instr.Op, xIsSigned, regX, regY, &regVal, size)
	case token.EQL, token.NEQ, token.LEQ, token.GEQ, token.LSS, token.GTR:
		if size != f.sizeof(instr.Y) {
			panic("Comparing two different size values")
		}
		instrdata := GetInstrDataType(instr.X.Type())
		asm += CmpOp(f.Indent, instrdata, instr.Op, regX, regY, &regVal)
	}
	f.freeReg(*regX)
	f.freeReg(*regY)

	addr, ok := f.ssaNames[instr.Name()]
	if !ok {
		panic(fmt.Sprintf("Unknown name (%v), instr (%v)\n", instr.Name(), instr))
	}

	a, err := f.asmStoreReg(&regVal, &addr, 0)
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

func (f *Function) asmBinOpLoadXY(instr *ssa.BinOp) (asm string, x *register, y *register, err *Error) {

	if err = f.allocValueOnDemand(instr); err != nil {
		return "", nil, nil, err
	}
	if err = f.allocValueOnDemand(instr.X); err != nil {
		return "", nil, nil, err
	}
	if err = f.allocValueOnDemand(instr.Y); err != nil {
		return "", nil, nil, err
	}

	xtmp := f.allocReg(regType(instr.X.Type()), f.sizeof(instr.X))
	x = &xtmp
	ytmp := f.allocReg(regType(instr.Y.Type()), f.sizeof(instr.Y))
	y = &ytmp
	asm = f.Indent + "// BEGIN asmBinOpLoadXY\n"

	if a, err := f.asmLoadValue(instr.X, 0, f.sizeof(instr.X), x); err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}

	if a, err := f.asmLoadValue(instr.Y, 0, f.sizeof(instr.Y), y); err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}

	asm += f.Indent + "// END asmBinOpLoadXY\n"
	return asm, x, y, nil
}

func (f *Function) sizeof(val ssa.Value) uint {
	if _, ok := val.(*ssa.Const); ok {
		return f.sizeofConst(val.(*ssa.Const))
	}
	info, ok := f.ssaNames[val.Name()]
	if !ok {
		panic(fmt.Sprintf("Unknown name (%v), value (%v)\n", val.Name(), val))
	}
	_, _, size := info.Addr()
	return size
}

func (f *Function) sizeofConst(cnst *ssa.Const) uint {
	return sizeof(cnst.Type())
}

func (f *Function) asmLoadValueSimple(val ssa.Value, reg *register) (string, *Error) {
	return f.asmLoadValue(val, 0, f.sizeof(val), reg)
}

func (f *Function) asmLoadValue(val ssa.Value, offset int, size uint, reg *register) (string, *Error) {
	if _, ok := val.(*ssa.Const); ok {
		return f.asmLoadConstValue(val.(*ssa.Const), reg)
	}
	info, ok := f.ssaNames[val.Name()]
	if !ok {
		panic(fmt.Sprintf("Unknown name (%v), value (%v)\n", val.Name(), val))
	}

	r, roffset, rsize := info.Addr()
	if rsize%size != 0 {
		panic(fmt.Sprintf("Size (%v) value not divisor of value (%v) size (%v), name (%v)\n", size, val, rsize, val.Name()))
	}
	if size > 8 {
		panic(fmt.Sprintf("Greater than 8 byte sized (%v) value, value (%v), name (%v)\n", size, val, val.Name()))
	}

	datatype := GetInstrDataType(val.Type())
	return MovMemReg(f.Indent, datatype, info.name, roffset+offset, &r, reg), nil
}

func (f *Function) asmStoreReg(reg *register, addr *nameInfo, offset int) (string, *Error) {
	r, roffset, rsize := addr.Addr()
	if rsize > 8 {
		panic(fmt.Sprintf("Greater than 8 byte sized (%v) value, addr (%v), name (%v)\n", rsize, *addr, addr.name))
	}
	if rsize == 0 {
		panic(fmt.Sprintf("size == 0 for addr (%v)", *addr))
	}
	asm := f.Indent + fmt.Sprintf("// BEGIN asmStoreReg, size (%v)\n", rsize)
	instrdata := GetInstrDataType(addr.typ)
	asm += MovRegMem(f.Indent, instrdata, reg, addr.name, &r, offset+roffset)
	asm += f.Indent + fmt.Sprintf("// END asmStoreReg, size (%v)\n", rsize)
	return asm, nil
}

func (f *Function) asmLoadConstValue(cnst *ssa.Const, r *register) (string, *Error) {

	if isBool(cnst.Type()) {
		var val int8
		if cnst.Value.String() == "true" {
			val = 1
		}
		return MovImm8Reg(f.Indent, val, r), nil
	}
	if isFloat(cnst.Type()) {
		if r.typ != XMM_REG {
			panic("Can't load float const into non xmm register")
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
		panic("Complex64/128 is unsupported ")
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

func (f *Function) asmUnOp(instr *ssa.UnOp) (string, *Error) {
	var err *Error
	asm := ""
	switch instr.Op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v): \"%v\"", instr.Op, instr))
	case token.NOT: // logical negation
		asm, err = f.asmUnOpXor(instr, 1)
	case token.XOR: //bitwise negation
		asm, err = f.asmUnOpXor(instr, -1)
	case token.SUB: // arithmetic negation e.g. x=>-x
		asm, err = f.asmUnOpSub(instr)
	case token.MUL: //pointer indirection
		asm, err = f.asmUnOpPointer(instr)
	}
	asm = f.Indent + fmt.Sprintf("// BEGIN ssa.UnOp: %v = %v\n", instr.Name(), instr) + asm
	asm += f.Indent + fmt.Sprintf("// END ssa.UnOp: %v = %v\n", instr.Name(), instr)
	return asm, err

}

// bitwise negation
func (f *Function) asmUnOpXor(instr *ssa.UnOp, xorVal int32) (string, *Error) {

	if err := f.allocValueOnDemand(instr); err != nil {
		return "", err
	}
	size := f.sizeof(instr)
	reg := f.allocReg(regType(instr.X.Type()), size)

	addr, ok := f.ssaNames[instr.Name()]
	if !ok {
		panic(fmt.Sprintf("Unknown name (%v), instr (%v)\n", instr.Name(), instr))
	}

	asm := ZeroReg(f.Indent, &reg)

	asm, err := f.asmLoadValueSimple(instr.X, &reg)
	if err != nil {
		return asm, err
	}

	if size < 8 {
		asm += XorImm32Reg(f.Indent, xorVal, &reg, size)
	} else {
		asm += XorImm64Reg(f.Indent, int64(xorVal), &reg, size)
	}

	a, err := f.asmStoreReg(&reg, &addr, 0)
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
func (f *Function) asmUnOpSub(instr *ssa.UnOp) (string, *Error) {

	if err := f.allocValueOnDemand(instr); err != nil {
		return "", err
	}
	var regX register
	var regSubX register
	var regVal register

	regVal = f.allocReg(regType(instr.Type()), f.sizeof(instr))
	regX = f.allocReg(regType(instr.X.Type()), f.sizeof(instr.X))
	regSubX = f.allocReg(regType(instr.X.Type()), f.sizeof(instr.X))

	addr, ok := f.ssaNames[instr.Name()]
	if !ok {
		panic(fmt.Sprintf("Unknown name (%v), instr (%v)\n", instr.Name(), instr))
	}

	asm, err := f.asmLoadValueSimple(instr.X, &regX)
	if err != nil {
		return asm, err
	}

	asm += ZeroReg(f.Indent, &regSubX)
	instrdata := GetInstrDataType(instr.Type())
	asm += ArithOp(f.Indent, instrdata, token.SUB, &regSubX, &regX, &regVal)
	f.freeReg(regX)
	f.freeReg(regSubX)

	a, err := f.asmStoreReg(&regVal, &addr, 0)
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
func (f *Function) asmUnOpPointer(instr *ssa.UnOp) (string, *Error) {
	assignment, ok := f.ssaNames[instr.Name()]
	xName := instr.X.Name()
	xInfo, okX := f.ssaNames[xName]
	// TODO add complex64/128 support
	if isComplex(instr.Type()) || isComplex(instr.X.Type()) {
		return "", &Error{fmt.Errorf("complex64/complex128 unimplemented"), instr.Pos()}
	}
	if !okX {
		panic(fmt.Sprintf("Unknown name for UnOp X (%v), instr \"(%v)\"", instr.X, instr))
	}
	if xInfo.local == nil && xInfo.param == nil && !xInfo.IsPointer() {
		panic(fmt.Sprintf("In UnOp, X (%v) isn't a pointer, X.type (%v), instr \"(%v)\"", instr.X, instr.X.Type(), instr))
	}
	asm := ""
	if !ok {
		info, err := f.asmAllocLocal(instr.Name(), instr.Type())
		if err != nil {
			panic(fmt.Sprintf("Err in UnOp X (%v), instr \"(%v)\", msg: \"%v\"", instr.X, instr, err))
		}
		assignment = info
	}

	xReg, xOffset, xSize := xInfo.Addr()
	aReg, aOffset, aSize := assignment.Addr()
	if xSize != sizePtr() {
		fmt.Printf("instr: %v\n", instr)
		panic(fmt.Sprintf("xSize (%v) != ptr size (%v)", xSize, sizePtr()))
	}

	size := aSize
	tmp1 := f.allocReg(DATA_REG, DataRegSize)
	tmp2 := f.allocReg(regType(instr.Type()), DataRegSize)
	instrdata := GetInstrDataType(instr.Type())
	asm += MovMemIndirectMem(f.Indent, instrdata, xInfo.name, xOffset, &xReg, assignment.name, aOffset, &aReg, size, &tmp1, &tmp2)
	f.ssaNames[assignment.name] = assignment
	f.freeReg(tmp1)
	f.freeReg(tmp2)
	return asm, nil
}

func (f *Function) asmIndexAddr(instr *ssa.IndexAddr) (string, *Error) {
	if instr == nil {
		return "", &Error{Err: errors.New("asmIndexAddr: nil instr"), Pos: instr.Pos()}

	}
	asm := ""
	xInfo := f.ssaNames[instr.X.Name()]

	assignment, ok := f.ssaNames[instr.Name()]
	if !ok {
		local, err := f.asmAllocLocal(instr.Name(), instr.Type())
		if err != nil {
			msg := fmt.Errorf("err in indexaddr op, msg:\"%v\"", err)
			return asm, &Error{Err: msg, Pos: instr.Pos()}
		}
		assignment = local
		f.ssaNames[instr.Name()] = assignment
	}

	xReg, xOffset, _ := xInfo.Addr()
	aReg, aOffset, _ := assignment.Addr()
	addrReg := f.allocReg(DATA_REG, sizePtr())
	idxReg := f.allocReg(DATA_REG, sizePtr())
	f.asmLoadValueSimple(instr.Index, &idxReg)
	asm += Lea(f.Indent, xInfo.name, xOffset, &xReg, &addrReg)
	instrdata := InstrDataType{INTEGER_OP, InstrData{signed: false, size: idxReg.width / 8}, XMM_INVALID}
	asm += AddRegReg(f.Indent, instrdata, &idxReg, &addrReg)
	instrdata = GetInstrDataType(assignment.typ)
	asm += MovRegMem(f.Indent, instrdata, &addrReg, assignment.name, &aReg, aOffset)
	f.freeReg(idxReg)
	f.freeReg(addrReg)
	f.ssaNames[instr.Name()] = assignment
	asm = f.Indent + fmt.Sprintf("// BEGIN ssa.IndexAddr: %v = %v\n", instr.Name(), instr) + asm
	asm += f.Indent + fmt.Sprintf("// END ssa.IndexAddr: %v = %v\n", instr.Name(), instr)
	return asm, nil
}

func (f *Function) asmAllocInstr(instr *ssa.Alloc) (string, *Error) {
	asm := ""
	if instr == nil {
		return "", &Error{Err: errors.New("asmAllocInstr: nil instr"), Pos: instr.Pos()}

	}
	if instr.Heap {
		return "", &Error{Err: errors.New("asmAllocInstr: heap alloc"), Pos: instr.Pos()}
	}

	//Alloc values are always addresses, and have pointer types, so the type
	//of the allocated variable is actually
	//Type().Underlying().(*types.Pointer).Elem().
	info := f.ssaNames[instr.Name()]
	if info.local == nil {
		panic(fmt.Sprintf("Expect %v to be a local variable", instr.Name()))
	}
	if _, ok := info.typ.(*types.Pointer); ok {
	} else {
	}
	f.ssaNames[instr.Name()] = info
	return asm, nil
}

func (f *Function) asmValue(value ssa.Value, dstReg *register, dstVar *varInfo) string {
	if dstReg == nil && dstVar == nil {
		panic("Both dstReg & dstVar are nil!")
	}
	if dstReg != nil && dstVar != nil {
		panic("Both dstReg & dstVar are non nil!")
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
	for _, name := range f.ssaNames {
		if name.local != nil {
			size += uint32(name.size)
		}
	}
	return size
}

func (f *Function) init() *Error {
	f.registers = make(map[string]bool)
	f.ssaNames = make(map[string]nameInfo)
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
			panic(fmt.Sprintf("couldn't alloc register, type: %v, width in bits: %v, size in bytes:%v", t, size*8, size))
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

func (f *Function) allocValueOnDemand(v ssa.Value) *Error {
	_, ok := f.ssaNames[v.Name()]
	if ok {
		return nil
	}
	switch v.(type) {
	case *ssa.Const:
		return nil
	}
	if !ok {
		local, err := f.asmAllocLocal(v.Name(), v.Type())
		if err != nil {
			msg := fmt.Errorf("err in allocValueOnDemand, msg:\"%v\"", err)
			return &Error{Err: msg, Pos: v.Pos()}
		}
		f.ssaNames[v.Name()] = local
	}
	return nil
}

var pointerSize = uint(8)
var sliceSize = uint(24)

type simdInfo struct {
	name     string
	size     uint
	elemSize uint
}

func simdReflect(t reflect.Type) simdInfo {
	elemSize := uint(0)
	if t.Kind() == reflect.Array {
		elemSize = uint(t.Elem().Size())
	}
	return simdInfo{t.Name(), uint(t.Size()), elemSize}
}

func simdTypes() []simdInfo {
	simdInt := reflect.TypeOf(simd.Int(0))
	simdInt4 := reflect.TypeOf(simd.Int4{})
	return []simdInfo{simdReflect(simdInt), simdReflect(simdInt4)}
}

func isSimd(t types.Type) bool {
	if t, ok := t.(*types.Named); ok {
		tname := t.Obj()
		for _, simdType := range simdTypes() {
			if tname.Name() == simdType.name {
				return true
			}
		}
	}
	return false
}

func simdTypeInfo(t types.Type) (simdInfo, error) {
	if !isSimd(t) {
		msg := fmt.Errorf("type (%v) is not simd type", t.String())
		return simdInfo{}, msg
	}
	named := t.(*types.Named)
	tname := named.Obj()
	for _, simdType := range simdTypes() {
		if tname.Name() == simdType.name {
			return simdType, nil
		}
	}
	msg := fmt.Errorf("type (%v) couldn't find simd type info", t.String())
	return simdInfo{}, msg
}

func simdHasElemSize(t types.Type) bool {
	if simdInfo, err := simdTypeInfo(t); err == nil {
		return simdInfo.elemSize > 0
	} else {
		msg := fmt.Sprintf("Error in simdHasElemSize, type (%v) is not simd", t.String())
		panic(msg)
	}
}

func simdElemSize(t types.Type) uint {
	if simdInfo, err := simdTypeInfo(t); err == nil {
		return simdInfo.elemSize
	} else {
		msg := fmt.Sprintf("Error in simdElemSize, type (%v) is not simd", t.String())
		panic(msg)
	}
}

func sizeofElem(t types.Type) uint {
	var e types.Type
	switch t := t.(type) {
	default:
		panic(fmt.Sprintf("t (%v) not an array or slice type\n", t.String()))
	case *types.Slice:
		e = t.Elem()
	case *types.Array:
		e = t.Elem()
	case *types.Named:
		if isSimd(t) && simdHasElemSize(t) {
			return simdElemSize(t)
		}
		panic(fmt.Sprintf("t (%v), isSimd (%v)\n", t.String(), isSimd(t)))
	}
	return sizeof(e)
}

func sizeof(t types.Type) uint {

	switch t := t.(type) {
	case *types.Tuple:
		// TODO: usage of reflect most likely wrong!
		// uint(reflect.TypeOf(t).Elem().Size())
		panic("Tuples are unsupported")
	case *types.Basic:
		return sizeBasic(t.Kind())
	case *types.Pointer:
		return sizePtr()
	case *types.Slice:
		return sliceSize
	case *types.Array:
		// TODO: calculation most likely wrong
		// return uint(t.Len()) * sizeof(t.Elem())
		panic("Arrays are unsupported")
	case *types.Named:
		if !isSimd(t) {
			panic("Named type is unsupported")
		}
		if info, err := simdTypeInfo(t); err != nil {
			panic(fmt.Sprintf("Error unknown type in sizeof err:\"%v\"", err))
		} else {
			return info.size
		}
	}
	panic(fmt.Sprintf("Error unknown type: %v", t))
}

func sizeInt() uint {
	return sizeBasic(types.Int)
}

func sizePtr() uint {
	return pointerSize
}

// sizeBasic return the size in bytes of a basic type
func sizeBasic(b types.BasicKind) uint {
	return uint(reflectBasic(b).Size())
}

func align(t types.Type) uint {

	switch t := t.(type) {
	case *types.Tuple:
		// TODO: usage of reflect most likely wrong!
		return uint(reflect.TypeOf(t).Elem().Size())
	case *types.Basic:
		return alignBasic(t.Kind())
	case *types.Pointer:
		return alignPtr()
	case *types.Slice:
		return alignSlice()
	case *types.Array:
		// TODO: most likely wrong
		return alignSlice()
	case *types.Named:
		panic(fmt.Sprintf("Error unknown named type in align:\"%v\"", t))
	}
	panic(fmt.Sprintf("Error unknown type (%v)", t))
}

func alignPtr() uint {
	return 8
}
func alignSlice() uint {
	return 8
}

func alignBasic(b types.BasicKind) uint {
	return uint(reflectBasic(b).Align())
}

func signed(t types.Type) bool {

	switch t := t.(type) {
	case *types.Basic:
		return signedBasic(t.Kind())
	}
	panic(fmt.Sprintf("Error unknown type: %v", t))
}

func signedBasic(b types.BasicKind) bool {
	switch b {
	case types.Bool:
		return false
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64:
		return true
	case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
		return false
	case types.Float32, types.Float64:
		return true
	}
	panic(fmt.Sprintf("Unknown basic type (%v)", b))
}

func isBool(t types.Type) bool {
	return isBasicKind(t, types.Bool)
}

func isFloat(t types.Type) bool {
	return isFloat32(t) || isFloat64(t)
}

func isFloat32(t types.Type) bool {
	return isBasicKind(t, types.Float32)
}

func isFloat64(t types.Type) bool {
	return isBasicKind(t, types.Float64)
}

func isComplex(t types.Type) bool {
	return isBasicKind(t, types.Complex64) || isBasicKind(t, types.Complex128)
}

func isBasicKind(t types.Type, basickind types.BasicKind) bool {
	if t, ok := t.(*types.Basic); ok {
		return t.Kind() == basickind
	}
	return false

}
func reflectType(t types.Type) reflect.Type {
	switch t := t.(type) {
	case *types.Tuple:
		// TODO
	case *types.Basic:
		return reflectBasic(t.Kind())
	case *types.Pointer:
		// TODO
	case *types.Slice:
		// TODO
	case *types.Array:
		// TODO
	case *types.Named:
		// TODO
	}
	panic(fmt.Sprintf("Error unknown type:\"%v\"", t))
}

func reflectBasic(b types.BasicKind) reflect.Type {
	switch b {
	default:
		panic("Unknown basic type")
	case types.Bool:
		return reflect.TypeOf(true)
	case types.Int:
		return reflect.TypeOf(int(1))
	case types.Int8:
		return reflect.TypeOf(int8(1))
	case types.Int16:
		return reflect.TypeOf(int16(1))
	case types.Int32:
		return reflect.TypeOf(int32(1))
	case types.Int64:
		return reflect.TypeOf(int64(1))
	case types.Uint:
		return reflect.TypeOf(uint(1))
	case types.Uint8:
		return reflect.TypeOf(uint8(1))
	case types.Uint16:
		return reflect.TypeOf(uint16(1))
	case types.Uint32:
		return reflect.TypeOf(uint32(1))
	case types.Uint64:
		return reflect.TypeOf(uint64(1))
	case types.Float32:
		return reflect.TypeOf(float32(1))
	case types.Float64:
		return reflect.TypeOf(float64(1))
	}
}

var bInstrData = InstrDataType{INTEGER_OP, InstrData{signed: false, size: 1}, XMM_INVALID}
var f32InstrData = InstrDataType{XMM_OP, InstrData{}, XMM_F32}
var f64InstrData = InstrDataType{XMM_OP, InstrData{}, XMM_F64}

func GetInstrDataType(t types.Type) InstrDataType {
	if isBool(t) {
		return bInstrData
	}
	if isFloat32(t) {
		return f32InstrData
	} else if isFloat64(t) {
		return f64InstrData
	}
	if isComplex(t) {
		panic("complex32/64 unsupported")
	}
	instrdata := InstrDataType{
		INTEGER_OP,
		InstrData{signed: signed(t), size: sizeof(t)},
		XMM_INVALID}
	return instrdata

}

func regType(t types.Type) RegType {
	if isFloat(t) {
		return XMM_REG
	}
	return DATA_REG
}
