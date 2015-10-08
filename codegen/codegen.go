package codegen

import (
	"errors"
	"fmt"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/types"

	"reflect"

	"github.com/bjwbell/gensimd/codegen/instructionsetxml"
	"github.com/bjwbell/gensimd/simd"

	"golang.org/x/tools/go/ssa"
)

type Function struct {
	Indent         string
	ssa            *ssa.Function
	instructionset *instructionsetxml.Instructionset
	registers      map[string]bool // maps register to false if unused and true if used
	ssaNames       map[string]nameInfo
}

type nameInfo struct {
	name string
	typ  types.Type
	//reg   *register
	local *varInfo
	param *paramInfo
}

// RegAndOffset returns the register and offset to access the nameInfo memory.
// For locals the register is the stack pointer (SP) and for params the register
// is the frame pointer (FP).
func (name *nameInfo) MemRegOffsetSize() (reg register, offset uint, size uint) {
	if name.local != nil {

		reg = *getRegister(REG_SP)
		offset = name.local.offset
		size = name.local.size
	} else if name.param != nil {
		reg = *getRegister(REG_FP)
		offset = name.param.offset
		size = name.param.size
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
	// offset from the stack pointer (SP)
	offset uint
	size   uint
	info   *ssa.Alloc
	//reg    *register
}

/*func (v *varInfo) Reg() (*register, error) {
	if v.reg != nil {
		return v.reg, nil
	} else {
		return nil, errors.New("varInfo has no reg set")
	}
}*/

func (v *varInfo) ssaName() string {
	return v.info.Name()
}

type paramInfo struct {
	name string
	// offset from the frame pointer (FP)
	offset uint
	size   uint
	info   *ssa.Parameter
	extra  interface{}
}

/*func (p *paramInfo) Reg() (*register, error) {
	if p.extra != nil {
		return &p.extra.(*paramSlice).reg, nil
	} else {
		return nil, errors.New("param p has no register set")
	}
}*/

func (p *paramInfo) ssaName() string {
	return p.info.Name()
}

type paramSlice struct {
	//offset uint
	lenOffset uint
	/*reg       register
	lenReg    register*/
}

type Error struct {
	Err error
	Pos token.Pos
}

func CreateFunction(instructionsetPath string, fn *ssa.Function) (*Function, *Error) {
	if fn == nil {
		return nil, &Error{Err: errors.New("Nil function passed in")}
	}
	instructionset, err := instructionsetxml.LoadInstructionset(instructionsetPath)
	if err != nil {
		return nil, &Error{Err: err}
	}
	f := Function{ssa: fn, instructionset: instructionset}
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
	offset := uint(0)
	asm := ""
	//var reg *register
	for _, p := range f.ssa.Params {
		param := paramInfo{name: p.Name(), offset: offset, info: p, size: sizeof(p.Type())}
		// TODO alloc reg based on other param types
		if _, ok := p.Type().(*types.Slice); ok {
			/*opMem := Operand{Type: OperandType(M64), Input: true, Output: false, Value: memFn(param.name, offset, "FP")}
			r := f.allocReg(AddrReg, pointerSize)
			opReg := Operand{Type: OperandType(R64), Input: false, Output: true, Value: regFn(r.name)}
			ops := []*Operand{&opMem, &opReg}
			if a, err := InstrAsm(f.instructionset, GetInstrType(TMOV), ops); err != nil {
				return "", &Error{err, p.Pos()}
			} else {
				asm += f.Indent + a + "\n"
			}
			// TODO is sizeof length data always pointer size?
			lenReg := f.allocReg(AddrReg, pointerSize)
			opMem.Value = memFn(param.name+"len", offset+pointerSize, "FP")
			opReg.Value = regFn(lenReg.name)
			if a, err := InstrAsm(f.instructionset, GetInstrType(TMOV), ops); err != nil {
				return "", &Error{err, p.Pos()}
			} else {
				asm += f.Indent + a + "\n"
			}*/
			param.extra = paramSlice{lenOffset: offset + pointerSize} //,reg: r, lenReg: lenReg}
			//reg = &r
		} else if basic, ok := p.Type().(*types.Basic); ok && basic.Kind() == types.Int {
			/*opMem := Operand{Type: OperandType(M64), Input: true, Output: false, Value: memFn(param.name, offset, "FP")}
			r := f.allocReg(DataReg, intSize())
			opReg := Operand{Type: OperandType(R64), Input: false, Output: true, Value: regFn(r.name)}
			ops := []*Operand{&opMem, &opReg}
			if a, err := InstrAsm(f.instructionset, GetInstrType(TMOV), ops); err != nil {
				return "", &Error{err, p.Pos()}
			} else {
				asm += f.Indent + a + "\n"
			}
			reg = &r*/
		} else {
			return "", &Error{Err: errors.New("Unsupported param type"), Pos: p.Pos()}
		}
		f.ssaNames[param.name] = nameInfo{name: param.name, typ: param.info.Type(),
			//reg: reg,
			local: nil, param: &param}
		offset += param.size
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

	basicblocks, err := f.asmBasicBlocks()
	if err != nil {
		return params + zeroRetValue + zeroSsaLocals + basicblocks, err
	}

	zeroNonSsaLocals, err := f.asmZeroNonSsaLocals()
	if err != nil {
		return zeroNonSsaLocals, err
	}

	frameSize := f.localsSize()
	asm := params
	asm += zeroRetValue
	asm += zeroSsaLocals
	asm += zeroNonSsaLocals
	asm += basicblocks

	a := fmt.Sprintf("TEXT ·%v(SB),NOSPLIT,$%v-%v\n%v", f.ssa.Name(), frameSize, f.paramsSize()+f.retSize(), asm)
	return a, nil
}

func (f *Function) asmZeroSsaLocals() (string, *Error) {
	asm := ""
	offset := uint(0)
	locals := f.ssa.Locals
	for _, local := range locals {
		if local.Heap {
			msg := errors.New(fmt.Sprintf("Can't heap alloc local, name: %v", local.Name()))
			return "", &Error{Err: msg, Pos: local.Pos()}
		}
		sp := getRegister(REG_SP)

		//local values are always addresses, and have pointer types, so the type
		//of the allocated variable is actually
		//Type().Underlying().(*types.Pointer).Elem().
		typ := local.Type().Underlying().(*types.Pointer).Elem()
		size := sizeof(typ)
		asm += asmZeroMemory(f.Indent, local.Name(), offset, size, sp)
		v := varInfo{name: local.Name(), offset: offset, size: size, info: local}
		f.ssaNames[v.name] = nameInfo{name: v.name, typ: typ,
			//reg: nil,
			local: &v, param: nil}
		offset += size
	}
	return asm, nil
}

func (f *Function) asmAllocLocal(name string, typ types.Type) (local nameInfo, err *Error) {
	size := sizeof(typ)
	//single byte size not supported
	if size == 1 {
		size = 8
	}
	v := varInfo{name: name, offset: f.localsSize(), size: sizeof(typ), info: nil}
	info := nameInfo{name: name, typ: typ, param: nil, local: &v}
	f.ssaNames[v.name] = info
	// zeroing the memory is done at the beginning of the function
	//asmZeroMemory(f.Indent, v.name, v.offset, v.size, sp)
	local = info
	err = nil
	return
}

func (f *Function) asmZeroNonSsaLocals() (string, *Error) {
	asm := ""
	for _, name := range f.ssaNames {
		if name.local == nil || name.IsSsaLocal() {
			continue
		}
		sp := getRegister(REG_SP)
		// single byte size is not supported
		if name.local.size == 1 {
			name.local.size = 8
		}
		asm += asmZeroMemory(f.Indent, name.name, name.local.offset, name.local.size, sp)
	}
	return asm, nil
}

func (f *Function) asmZeroRetValue() (string, *Error) {
	asm := asmZeroMemory(f.Indent, "~ret", f.retOffset(), f.retSize(), getRegister(REG_FP))
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
	asm := strconv.Itoa(block.Index) + ":\n"
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
		panic("Nil instr in asmInstr")
	}
	asm := ""

	switch instr := instr.(type) {
	default:
		asm += f.Indent + fmt.Sprintf("Unknown ssa instruction: %v\n", instr)
	case *ssa.Alloc:
		if a, err := f.asmAllocInstr(instr); err != nil {
			//log.Fatal("Error in f.asmAllocInstr")
			return a, err
		} else {
			asm += a
		}
	case *ssa.BinOp:
		if a, err := f.asmBinOp(instr); err != nil {
			return a, err
		} else {
			asm += f.Indent + fmt.Sprintf("// ssa.BinOp: %v, name: %v\n", instr, instr.Name())
			asm += a
		}
	case *ssa.Call:
		asm += f.Indent + fmt.Sprintf("ssa.Call: %v, name: %v\n", instr, instr.Name())
	case *ssa.ChangeInterface:
		asm += f.Indent + fmt.Sprintf("ssa.ChangeInterface: %v, name: %v\n", instr, instr.Name())
	case *ssa.ChangeType:
		asm += f.Indent + fmt.Sprintf("ssa.ChangeType: %v, name: %v\n", instr, instr.Name())
	case *ssa.Convert:
		asm += f.Indent + fmt.Sprintf("ssa.Convert: %v, name: %v\n", instr, instr.Name())
	case *ssa.Defer:
		asm += f.Indent + fmt.Sprintf("ssa.Defer: %v\n", instr)
	case *ssa.Extract:
		asm += f.Indent + fmt.Sprintf("ssa.Extra: %v, name: %v\n", instr, instr.Name())
	case *ssa.Field:
		asm += f.Indent + fmt.Sprintf("ssa.Field: %v, name: %v\n", instr, instr.Name())
	case *ssa.FieldAddr:
		asm += f.Indent + fmt.Sprintf("ssa.FieldAddr: %v, name: %v\n", instr, instr.Name())
	case *ssa.Go:
		asm += f.Indent + fmt.Sprintf("ssa.Go: %v\n", instr)
	case *ssa.If:
		asm += f.Indent + fmt.Sprintf("ssa.If: %v\n", instr)
	case *ssa.Index:
		asm += f.Indent + fmt.Sprintf("ssa.Index: %v, name: %v\n", instr, instr.Name())
	case *ssa.IndexAddr:
		if a, err := f.asmIndexAddr(instr); err != nil {
			return a, err
		} else {
			asm += f.Indent + fmt.Sprintf("// ssa.IndexAddr: %v, name: %v\n", instr, instr.Name())
			asm += a
		}
	case *ssa.Jump:
		asm += f.Indent + strings.Replace(instr.String(), "jump", "JMP ", -1) + "\n"
	case *ssa.Lookup:
		asm += f.Indent + fmt.Sprintf("ssa.Lookup: %v, name: %v\n", instr, instr.Name())
	case *ssa.MakeChan:
		asm += f.Indent + fmt.Sprintf("ssa.MakeChan: %v, name: %v\n", instr, instr.Name())
	case *ssa.MakeClosure:
		asm += f.Indent + fmt.Sprintf("ssa.MakeClosure: %v, name: %v\n", instr, instr.Name())
	case *ssa.MakeInterface:
		asm += f.Indent + fmt.Sprintf("ssa.MakeInterface: %v, name: %v\n", instr, instr.Name())
	case *ssa.MakeMap:
		asm += f.Indent + fmt.Sprintf("ssa.MakeMap: %v, name: %v\n", instr, instr.Name())
	case *ssa.MakeSlice:
		asm += f.Indent + fmt.Sprintf("ssa.MakeSlice: %v, name: %v\n", instr, instr.Name())
	case *ssa.MapUpdate:
		asm += f.Indent + fmt.Sprintf("ssa.MapUpdate: %v\n", instr)
	case *ssa.Next:
		asm += f.Indent + fmt.Sprintf("ssa.Next: %v, name: %v\n", instr, instr.Name())
	case *ssa.Panic:
		asm += f.Indent + fmt.Sprintf("ssa.Panic: %v", instr) + "\n"
	case *ssa.Phi:
		asm += f.Indent + fmt.Sprintf("ssa.Phi: %v, name: %v\n", instr, instr.Name())
	case *ssa.Range:
		asm += f.Indent + fmt.Sprintf("ssa.Range: %v, name: %v\n", instr, instr.Name())
	case *ssa.Return:
		asm += f.Indent + fmt.Sprintf("ssa.Return: %v", instr) + "\n"
	case *ssa.RunDefers:
		asm += f.Indent + fmt.Sprintf("ssa.RunDefers: %v", instr) + "\n"
	case *ssa.Select:
		asm += f.Indent + fmt.Sprintf("ssa.Select: %v, name: %v\n", instr, instr.Name())
	case *ssa.Send:
		asm += f.Indent + fmt.Sprintf("ssa.Send: %v", instr) + "\n"
	case *ssa.Slice:
		asm += f.Indent + fmt.Sprintf("ssa.Slice: %v, name: %v\n", instr, instr.Name())
	case *ssa.Store:
		asm += f.Indent + fmt.Sprintf("ssa.Store: %v, addr: %v, val: %v\n", instr, instr.Addr, instr.Val)
	case *ssa.TypeAssert:
		asm += f.Indent + fmt.Sprintf("ssa.TypeAssert: %v, name: %v\n", instr, instr.Name())
	case *ssa.UnOp:
		if a, err := f.asmUnOp(instr); err != nil {
			return a, err
		} else {
			asm += f.Indent + fmt.Sprintf("// ssa.UnOp: %v, name: %v\n", instr, instr.Name())
			asm += a
		}
	}
	return asm, nil
}

func (f *Function) asmBinOp(instr *ssa.BinOp) (string, *Error) {
	if err := f.allocValueOnDemand(instr); err != nil {
		return "", err
	}
	var regX, regY *register
	var regVal register
	// comparison op results are size 1 byte, but that's not supported
	if f.sizeof(instr) == 1 {
		regVal = f.allocReg(DataReg, 8*f.sizeof(instr))
	} else {
		regVal = f.allocReg(DataReg, f.sizeof(instr))
	}
	asm, regX, regY, err := f.asmBinOpLoadXY(instr)
	if err != nil {
		return asm, err
	}
	switch instr.Op {
	default:
		panic(fmt.Sprintf("Unknown op (%v) in asmBinOp", instr.Op))
	case token.ADD, token.SUB, token.MUL, token.QUO, token.REM:
		asm += asmArithOp(f.Indent, instr.Op, regX, regY, &regVal)
	case token.AND, token.OR, token.XOR, token.SHL, token.SHR, token.AND_NOT:
		asm += asmBitwiseOp(f.Indent, instr.Op, regX, regY, &regVal)
	case token.EQL, token.NEQ, token.LEQ, token.GEQ, token.LSS, token.GTR:
		asm += asmCmpOp(f.Indent, instr.Op, regX, regY, &regVal)
	}
	f.freeReg(*regX)
	f.freeReg(*regY)
	f.freeReg(regVal)
	asm = fmt.Sprintf(f.Indent+"//instr %v\n", instr) + asm
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

	xtmp := f.allocReg(DataReg, f.sizeof(instr.X))
	x = &xtmp
	ytmp := f.allocReg(DataReg, f.sizeof(instr.Y))
	y = &ytmp
	asm = ""
	if a, err := f.asmLoadValue(instr.X, x); err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}
	if a, err := f.asmLoadValue(instr.Y, y); err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}
	return asm, x, y, nil
}

func (f *Function) sizeof(val ssa.Value) uint {
	if _, ok := val.(*ssa.Const); ok {
		return f.sizeofConst(val.(*ssa.Const))
	}
	info, ok := f.ssaNames[val.Name()]
	if !ok {
		panic(fmt.Sprintf("Unknown name (%v) in asmLoadValue, value (%v)\n", val.Name(), val))
	}
	_, _, size := info.MemRegOffsetSize()
	return size
}

func (f *Function) sizeofConst(cnst *ssa.Const) uint {
	return sizeof(cnst.Type())
}

func (f *Function) asmLoadValue(val ssa.Value, reg *register) (string, *Error) {
	if _, ok := val.(*ssa.Const); ok {
		return f.asmLoadConstValue(val.(*ssa.Const), reg)
	}
	info, ok := f.ssaNames[val.Name()]
	if !ok {
		panic(fmt.Sprintf("Unknown name (%v) in asmLoadValue, value (%v)\n", val.Name(), val))
	}
	// TODO handle non 64 bit values
	r, offset, size := info.MemRegOffsetSize()
	if size != 8 {
		panic(fmt.Sprintf("Non 64bit sized (%v) value in asmLoadValue, value (%v), name (%v)\n", size, val, val.Name()))
	}
	return asmMovMemReg(f.Indent, info.name, offset, &r, reg), nil
}

func (f *Function) asmLoadConstValue(cnst *ssa.Const, r *register) (string, *Error) {
	cnstValue := uint32(cnst.Uint64())
	return asmLoadImm32(f.Indent, cnstValue, r), nil
}

func (f *Function) asmUnOp(instr *ssa.UnOp) (string, *Error) {
	switch instr.Op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v) in asmUnOp: \"%v\"", instr.Op, instr))
	case token.NOT: // logical negation
		return f.asmUnOpNot(instr)
	case token.XOR: //bitwise complement
		return f.asmUnOpXor(instr)
	case token.SUB: // arithmetic negation
		return f.asmUnOpSub(instr)
	case token.MUL: //pointer indirection
		return f.asmUnOpPointer(instr)
	}
}

// logical negation
func (f *Function) asmUnOpNot(instr *ssa.UnOp) (string, *Error) {
	// TODO
	return fmt.Sprintf(f.Indent+"// instr %v\n", instr), nil
}

//bitwise complement
func (f *Function) asmUnOpXor(instr *ssa.UnOp) (string, *Error) {
	// TODO
	return fmt.Sprintf(f.Indent+"// instr %v\n", instr), nil
}

// arithmetic negation
func (f *Function) asmUnOpSub(instr *ssa.UnOp) (string, *Error) {
	// TODO
	return fmt.Sprintf(f.Indent+"// instr %v\n", instr), nil
}

//pointer indirection
func (f *Function) asmUnOpPointer(instr *ssa.UnOp) (string, *Error) {
	assignment, ok := f.ssaNames[instr.Name()]
	xName := instr.X.Name()
	xInfo, okX := f.ssaNames[xName]

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
		/*if xInfo.local == nil && xInfo.param == nil {
			assignment.typ = xInfo.PointerUnderlyingType()
		} else {
			assignment.typ = xInfo.typ
		}*/
	}
	xReg, xOffset, xSize := xInfo.MemRegOffsetSize()
	aReg, aOffset, aSize := assignment.MemRegOffsetSize()
	if xSize != aSize {
		panic("xSize := aSize in asmUnOpPointer")
	}
	size := aSize

	asm += asmMovMemMem(f.Indent, xInfo.name, xOffset, &xReg, assignment.name, aOffset, &aReg, size)
	f.ssaNames[assignment.name] = assignment
	return asm, nil
}

func (f *Function) asmIndexAddr(instr *ssa.IndexAddr) (string, *Error) {
	if instr == nil {
		return "", &Error{Err: errors.New("asmIndexAddr: nil instr"), Pos: instr.Pos()}

	}
	asm := ""
	constIndex := false
	paramIndex := false
	var cnst *ssa.Const
	var param *ssa.Parameter
	switch instr.Index.(type) {
	default:
	case *ssa.Const:
		constIndex = true
		cnst = instr.Index.(*ssa.Const)
	case *ssa.Parameter:
		paramIndex = true
		param = instr.Index.(*ssa.Parameter)
	}

	xInfo := f.ssaNames[instr.X.Name()]

	// TODO check if xInfo is pointer, array, struct, etc.
	//if xInfo.IsPointer() || xInfo.IsArray() {

	/*if xInfo.reg == nil {
		msg := fmt.Sprintf("nil xInfo.reg (%v) in indexaddr op", xInfo.name)
		return asm, &Error{Err: errors.New(msg), Pos: instr.Pos()}
	}*/

	assignment, ok := f.ssaNames[instr.Name()]
	if !ok {
		local, err := f.asmAllocLocal(instr.Name(), instr.Type())
		if err != nil {
			msg := fmt.Sprintf("err in indexaddr op, msg:\"%v\"", err)
			return asm, &Error{Err: errors.New(msg), Pos: instr.Pos()}
		}
		assignment = local
		f.ssaNames[instr.Name()] = assignment
	}

	if constIndex {
		tmpReg := f.allocReg(DataReg, pointerSize)
		size := uint(sizeofElem(xInfo.typ))
		idx := uint(cnst.Uint64())
		xReg, xOffset, _ := xInfo.MemRegOffsetSize()
		assignmentReg, assignmentOffset, _ := assignment.MemRegOffsetSize()
		asm += asmLea(f.Indent, xInfo.name, xOffset+idx*size, &xReg, &tmpReg)
		asm += asmMovRegMem(f.Indent, &tmpReg, assignment.name, &assignmentReg, assignmentOffset)
		f.freeReg(tmpReg)
	} else if paramIndex {
		p := f.ssaNames[param.Name()]
		tmpReg := f.allocReg(DataReg, pointerSize)
		tmp2Reg := f.allocReg(DataReg, pointerSize)
		xReg, xOffset, _ := xInfo.MemRegOffsetSize()
		pReg, pOffset, pSize := p.MemRegOffsetSize()
		if pSize != 8 {
			fmt.Println("instr:", instr)
			fmt.Println("pSize:", pSize)
			panic("Index size not 8 bytes in asmIndexAddr")
		}
		assignmentReg, assignmentOffset, _ := assignment.MemRegOffsetSize()
		asm += asmMovMemReg(f.Indent, p.name, pOffset, &pReg, &tmp2Reg)
		asm += asmLea(f.Indent, xInfo.name, xOffset, &xReg, &tmpReg)
		asm += asmAddRegReg(f.Indent, &tmpReg, &tmp2Reg)
		asm += asmMovRegMem(f.Indent, &tmp2Reg, assignment.name, &assignmentReg, assignmentOffset)
		f.freeReg(tmpReg)
		f.freeReg(tmp2Reg)

	} else {
		asm = fmt.Sprintf(f.Indent+"// instr:%v\n", instr)
	}
	f.ssaNames[instr.Name()] = assignment
	return asm, nil
}

func (f *Function) asmAllocInstr(instr *ssa.Alloc) (string, *Error) {
	asm := ""
	//var err error
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
		/*opMem := Operand{Type: OperandType(M64), Input: true, Output: false, Value: memFn(info.name, info.local.offset, "SP")}
		reg := f.allocReg(AddrReg, pointerSize)
		opReg := Operand{Type: OperandType(R64), Input: false, Output: true, Value: regFn(reg.name)}
		ops := []*Operand{&opMem, &opReg}
		info.reg = &reg
		if asm, err = InstrAsm(f.instructionset, GetInstrType(TMOV), ops); err != nil {
			return "", &Error{err, instr.Pos()}
		}
		comment := f.Indent + fmt.Sprintf("// %v = %v\n", info.name, reg.name)
		asm = comment + f.Indent + asm + "\n"*/
	} else {
		/*opMem := Operand{Type: OperandType(M), Input: true, Output: false, Value: memFn(info.name, info.local.offset, "SP")}
		reg := f.allocReg(AddrReg, pointerSize)
		//info.reg = &reg
		opReg := Operand{Type: OperandType(R64), Input: false, Output: true, Value: regFn(reg.name)}
		ops := []*Operand{&opMem, &opReg}
		if asm, err = InstrAsm(f.instructionset, GetInstrType(TLEA), ops); err != nil {
			return "", &Error{err, instr.Pos()}
		}
		comment := f.Indent + fmt.Sprintf("// &%v = %v\n", info.name, reg.name)
		asm = comment + f.Indent + asm + "\n"*/
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

func (f *Function) localsSize() uint {
	size := uint(0)
	for _, name := range f.ssaNames {
		if name.local != nil {
			size += sizeof(name.typ)
		}
	}
	return size
}

func (f *Function) init() *Error {
	f.registers = make(map[string]bool)
	f.ssaNames = make(map[string]nameInfo)
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
		// r.width is in bits so multiple size (which is in bytes) by 8
		if !used && r.typ == t && r.width == size*8 {
			reg = r
			found = true
			break
		}
	}
	if found {
		f.registers[reg.name] = true
	} else {
		// any of the data registers can be used as an address register on x86_64
		if t == AddrReg {
			return f.allocReg(DataReg, size)
		} else {
			panic(fmt.Sprintf("couldn't alloc register, type: %v, width in bits: %v", t, size*8))
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
	return asmZeroReg(f.Indent, r)
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

// retSize returns the size of the return value in bytes
func (f *Function) retSize() uint {
	results := f.ssa.Signature.Results()
	if results.Len() == 0 {
		return 0
	}
	if results.Len() > 1 {
		panic("Functions with more than one return value not supported")
	}
	size := sizeof(results)
	return size
}

// retOffset returns the offset of the return value in bytes
func (f *Function) retOffset() uint {
	return f.paramsSize()
}

func (f *Function) allocValueOnDemand(v ssa.Value) *Error {
	_, ok := f.ssaNames[v.Name()]
	if !ok {
		local, err := f.asmAllocLocal(v.Name(), v.Type())
		if err != nil {
			msg := fmt.Sprintf("err in allocValueOnDemand, msg:\"%v\"", err)
			return &Error{Err: errors.New(msg), Pos: v.Pos()}
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
		msg := fmt.Sprintf("type (%v) is not simd type", t.String())
		return simdInfo{}, errors.New(msg)
	}
	named := t.(*types.Named)
	tname := named.Obj()
	for _, simdType := range simdTypes() {
		if tname.Name() == simdType.name {
			return simdType, nil
		}
	}
	msg := fmt.Sprintf("type (%v) couldn't find simd type info", t.String())
	return simdInfo{}, errors.New(msg)
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
	default:
		fmt.Println("t:", t)
		panic("Error unknown type in sizeof")
	case *types.Tuple:
		// TODO: fix, usage of reflect is wrong!
		return uint(reflect.TypeOf(t).Elem().Size())
	case *types.Basic:
		return sizeBasic(t)
	case *types.Pointer:
		return pointerSize
	case *types.Slice:
		return sliceSize
	case *types.Array:
		// TODO: fix, calculation most likely wrong
		return uint(t.Len()) * sizeof(t.Elem())
	case *types.Named:
		if !isSimd(t) {

		}
		if info, err := simdTypeInfo(t); err != nil {
			panic(fmt.Sprintf("Error unknown type in sizeof err:\"%v\"", err))
		} else {
			return info.size
		}
	}
}

func intSize() uint {
	return uint(reflect.TypeOf(int(1)).Size())
}

func uintSize() uint {
	return uint(reflect.TypeOf(uint(1)).Size())
}

func boolSize() uint {
	return uint(reflect.TypeOf(true).Size())
}

func ptrSize() uint {
	return pointerSize
}

// sizeBasic return the size in bytes of a basic type
func sizeBasic(b *types.Basic) uint {
	switch b.Kind() {
	default:
		panic("Unknown basic type")
	case types.Bool:
		return uint(reflect.TypeOf(true).Size())
	case types.Int:
		return uint(reflect.TypeOf(int(1)).Size())
	case types.Int8:
		return uint(reflect.TypeOf(int8(1)).Size())
	case types.Int16:
		return uint(reflect.TypeOf(int16(1)).Size())
	case types.Int32:
		return uint(reflect.TypeOf(int32(1)).Size())
	case types.Int64:
		return uint(reflect.TypeOf(int64(1)).Size())
	case types.Uint:
		return uint(reflect.TypeOf(uint(1)).Size())
	case types.Uint8:
		return uint(reflect.TypeOf(uint8(1)).Size())
	case types.Uint16:
		return uint(reflect.TypeOf(uint16(1)).Size())
	case types.Uint32:
		return uint(reflect.TypeOf(uint32(1)).Size())
	case types.Uint64:
		return uint(reflect.TypeOf(uint64(1)).Size())
	case types.Float32:
		return uint(reflect.TypeOf(float32(1)).Size())
	case types.Float64:
		return uint(reflect.TypeOf(float64(1)).Size())
	}
}
