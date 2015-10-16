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
	cnst   *ssa.Const
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
		panic(fmt.Sprintf("nameInfo (%v) not ptr type", name))
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
	asm := ""
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

func (f *Function) AllocLocal(name string, typ types.Type) (nameInfo, *Error) {
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
	//ZeroMemory(f.Indent, v.name, v.offset, v.size, sp)
	return info, nil
}

func (f *Function) ZeroNonSsaLocals() (string, *Error) {
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
		panic("Nil instr")
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
		asm, err = errormsg("slice creation unimplemented")
	case *ssa.Store:
		asm, err = f.Store(instr)
	case *ssa.TypeAssert:
		asm, err = errormsg("type assert unsupported")
	case *ssa.UnOp:
		asm, err = f.UnOp(instr)
	}

	return asm, err
}

func (f *Function) Convert(instr *ssa.Convert) (string, *Error) {
	return "", nil
}

func (f *Function) If(instr *ssa.If) (string, *Error) {
	asm := ""
	tblock, fblock := -1, -1
	if instr.Block() != nil && len(instr.Block().Succs) == 2 {
		tblock = instr.Block().Succs[0].Index
		fblock = instr.Block().Succs[1].Index

	}
	if tblock == -1 || fblock == -1 {
		panic("If: malformed CFG")
	}

	if info, ok := f.ssaNames[instr.Cond.Name()]; !ok {

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
		panic("Jump: malformed CFG")
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
			panic("computePhiInstr: malformed CFG")
		}
		if _, ok := f.phiInfo[edgeBlock]; !ok {
			f.phiInfo[edgeBlock] = make(map[int][]phiInfo)
		}
		f.phiInfo[edgeBlock][blockIndex] = append(f.phiInfo[edgeBlock][blockIndex], phiInfo{value: edge, phi: phi})
	}
	return nil
}

func (f *Function) Phi(phi *ssa.Phi) (string, *Error) {

	if nameinfo := f.allocValueOnDemand(phi); nameinfo == nil {
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
		nameInfo{
			name:   retName(),
			typ:    f.retType(),
			local:  nil,
			param:  f.retParam(),
			size:   f.retSize(),
			offset: f.retOffset(),
			align:  f.retAlign()}

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

func (f *Function) StoreValAddr(val ssa.Value, addr *nameInfo) (string, *Error) {

	if nameinfo := f.allocValueOnDemand(val); nameinfo == nil {
		return ErrorMsg("Error in allocating local")
	}
	if addr.local == nil && addr.param == nil {
		return ErrorMsg(fmt.Sprintf("Invalid addr \"%v\"", addr))
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

		a, err = f.StoreReg(&valReg, addr, 0)
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
			a, err := f.LoadValue(val, offset, uint(datasize), &valReg)
			if err != nil {
				return a, err
			}
			asm += a
			a, err = f.StoreReg(&valReg, addr, offset)
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
	if nameinfo := f.allocValueOnDemand(instr.Addr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	addr, ok := f.ssaNames[instr.Addr.Name()]
	if !ok {

		panic("Couldnt find instr.Addr in ssaNames")
	}
	return f.StoreValAddr(instr.Val, &addr)
}

func (f *Function) BinOp(instr *ssa.BinOp) (string, *Error) {

	if nameinfo := f.allocValueOnDemand(instr); nameinfo == nil {
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

	a, err := f.StoreReg(&regVal, &addr, 0)
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

	if nameinfo := f.allocValueOnDemand(instr); nameinfo == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	if nameinfo := f.allocValueOnDemand(instr.X); nameinfo == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr.X))
	}
	if nameinfo := f.allocValueOnDemand(instr.Y); nameinfo == nil {
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

func (f *Function) LoadValueSimple(val ssa.Value, reg *register) (string, *Error) {
	return f.LoadValue(val, 0, f.sizeof(val), reg)
}

func (f *Function) LoadValue(val ssa.Value, offset int, size uint, reg *register) (string, *Error) {
	if _, ok := val.(*ssa.Const); ok {
		return f.LoadConstValue(val.(*ssa.Const), reg)
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

func (f *Function) StoreReg(reg *register, addr *nameInfo, offset int) (string, *Error) {
	r, roffset, rsize := addr.Addr()
	if rsize > sizePtr() {
		panic(fmt.Sprintf("Greater than ptr sized (%v), addr (%v), name (%v)\n", rsize, *addr, addr.name))
	}
	if rsize == 0 {
		panic(fmt.Sprintf("size == 0 for addr (%v)", *addr))
	}
	asm := f.Indent + fmt.Sprintf("// BEGIN StoreReg, size (%v)\n", rsize)
	instrdata := GetInstrDataType(addr.typ)
	asm += MovRegMem(f.Indent, instrdata, reg, addr.name, &r, offset+roffset)
	asm += f.Indent + fmt.Sprintf("// END StoreReg, size (%v)\n", rsize)
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
		panic("complex64/128 unsupported")
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
		panic(fmt.Sprintf("Unknown Op token (%v): \"%v\"", instr.Op, instr))
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

	if nameinfo := f.allocValueOnDemand(instr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	size := f.sizeof(instr)
	reg := f.allocReg(regType(instr.X.Type()), size)

	addr, ok := f.ssaNames[instr.Name()]
	if !ok {
		panic(fmt.Sprintf("Unknown name (%v), instr (%v)\n", instr.Name(), instr))
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

	a, err := f.StoreReg(&reg, &addr, 0)
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

	if nameinfo := f.allocValueOnDemand(instr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
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

	asm, err := f.LoadValueSimple(instr.X, &regX)
	if err != nil {
		return asm, err
	}

	asm += ZeroReg(f.Indent, &regSubX)
	instrdata := GetInstrDataType(instr.Type())
	asm += ArithOp(f.Indent, instrdata, token.SUB, &regSubX, &regX, &regVal)
	f.freeReg(regX)
	f.freeReg(regSubX)

	a, err := f.StoreReg(&regVal, &addr, 0)
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
	assignment := f.allocValueOnDemand(instr)
	if assignment == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	xName := instr.X.Name()
	xInfo, okX := f.ssaNames[xName]

	// TODO add complex64/128 support
	if isComplex(instr.Type()) || isComplex(instr.X.Type()) {
		return ErrorMsg("complex64/complex128 unimplemented")
	}
	if !okX {
		panic(fmt.Sprintf("Unknown name for UnOp X (%v), instr \"(%v)\"", instr.X, instr))
	}
	if xInfo.local == nil && xInfo.param == nil && !xInfo.IsPointer() {
		fmtstr := "In UnOp, X (%v) isn't a pointer, X.type (%v), instr \"(%v)\""
		msg := fmt.Sprintf(fmtstr, instr.X, instr.X.Type(), instr)
		panic(msg)
	}

	asm := ""

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

	f.ssaNames[assignment.name] = *assignment

	f.freeReg(tmp1)
	f.freeReg(tmp2)

	return asm, nil
}

func (f *Function) Index(instr *ssa.Index) (string, *Error) {
	if instr == nil {
		return ErrorMsg("nil instr")
	}
	asm := ""
	xInfo := f.ssaNames[instr.X.Name()]
	assignment := f.allocValueOnDemand(instr)
	if assignment == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	xReg, xOffset, _ := xInfo.Addr()
	aReg, aOffset, _ := assignment.Addr()
	addrReg := f.allocReg(DATA_REG, sizePtr())
	idxReg := f.allocReg(DATA_REG, sizePtr())

	f.LoadValueSimple(instr.Index, &idxReg)

	asm += Lea(f.Indent, xInfo.name, xOffset, &xReg, &addrReg)

	instrdata := GetIntegerInstrDataType(false, idxReg.width/8)
	asm += AddRegReg(f.Indent, instrdata, &idxReg, &addrReg)

	instrdata = GetInstrDataType(assignment.typ)
	asm += MovRegMem(f.Indent, instrdata, &addrReg, assignment.name, &aReg, aOffset)

	f.freeReg(idxReg)
	f.freeReg(addrReg)

	f.ssaNames[instr.Name()] = *assignment

	asm = f.Indent + fmt.Sprintf("// BEGIN ssa.IndexAddr: %v = %v\n", instr.Name(), instr) + asm
	asm += f.Indent + fmt.Sprintf("// END ssa.IndexAddr: %v = %v\n", instr.Name(), instr)

	return asm, nil
}

func (f *Function) IndexAddr(instr *ssa.IndexAddr) (string, *Error) {
	if instr == nil {
		return ErrorMsg("nil instr")

	}

	asm := ""
	xInfo := f.ssaNames[instr.X.Name()]
	assignment := f.allocValueOnDemand(instr)

	xReg, xOffset, _ := xInfo.Addr()
	aReg, aOffset, _ := assignment.Addr()
	addrReg := f.allocReg(DATA_REG, sizePtr())
	idxReg := f.allocReg(DATA_REG, sizePtr())

	f.LoadValueSimple(instr.Index, &idxReg)

	asm += Lea(f.Indent, xInfo.name, xOffset, &xReg, &addrReg)

	instrdata := GetIntegerInstrDataType(false, idxReg.width/8)
	asm += AddRegReg(f.Indent, instrdata, &idxReg, &addrReg)

	instrdata = GetInstrDataType(assignment.typ)
	asm += MovRegMem(f.Indent, instrdata, &addrReg, assignment.name, &aReg, aOffset)

	f.freeReg(idxReg)
	f.freeReg(addrReg)

	f.ssaNames[instr.Name()] = *assignment

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

func (f *Function) Value(value ssa.Value, dstReg *register, dstVar *varInfo) string {
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

func (f *Function) allocValueOnDemand(v ssa.Value) *nameInfo {

	if nameinfo, ok := f.ssaNames[v.Name()]; ok {
		return &nameinfo
	}

	switch v := v.(type) {
	case *ssa.Const:
		nameinfo := nameInfo{name: v.Name(), typ: v.Type(), local: nil, param: nil, cnst: v}
		f.ssaNames[v.Name()] = nameinfo
		return &nameinfo
	}

	local, err := f.AllocLocal(v.Name(), v.Type())
	if err != nil {
		return nil
	}

	f.ssaNames[v.Name()] = local

	return &local
}
