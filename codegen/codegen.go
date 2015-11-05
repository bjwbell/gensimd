package codegen

import (
	"errors"
	"fmt"
	"go/token"
	"math"
	"reflect"
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
	// if Debug is set, debug comments are included in assembly output
	Debug       bool
	PrintSpills bool
	Trace       bool
	Optimize    bool
	Indent      string
	identifiers map[string]*identifier
	jmpLabels   []string
	outfn       string // output function name

	// map from block index to the successor block indexes that need phi vars set
	phiInfo map[int]map[int][]phiInfo

	// maps register to false if unused and true if used
	registers []register

	ssa *ssa.Function
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

func CreateFunction(fn *ssa.Function, outfn string, debug bool, trace bool, optimize bool) (*Function, *Error) {
	if fn == nil {
		return nil, ErrorMsg2("Nil function passed in")
	}
	f := Function{ssa: fn, outfn: outfn, Debug: debug, Trace: trace, Optimize: optimize}
	f.Indent = "        "
	f.init()
	return &f, nil
}

func AssemblyFilePreamble() string {
	preamble := "// +build amd64 !noasm !appengine\n\n"
	preamble += "#include \"textflag.h\"\n\n"
	return preamble
}

func (f *Function) GoAssembly() (string, *Error) {
	asm, err := f.Func()
	if !f.Debug {
		asm = stripDebug(asm, f.Indent)
	}
	return asm, err
}

func (f *Function) Position(pos token.Pos) token.Position {
	return f.ssa.Prog.Fset.Position(pos)
}

func (f *Function) Params() (string, *Error) {
	// offset in bytes from frame pointer (FP)
	offset := int(0)
	asm := ""
	for _, p := range f.ssa.Params {
		param := p
		// TODO alloc reg based on other param types
		if basic, ok := p.Type().(*types.Basic); ok {
			switch basic.Kind() {
			default:
				err := ErrorMsg2(fmt.Sprintf("Unsupported param type (%v)", basic))
				err.Pos = p.Pos()
				return "", err

			case types.Float32, types.Float64:
				break

				// supported param types
			case types.Bool, types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
				types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
				break
			}

		} else {

		}
		ident := identifier{f: f, name: param.Name(), typ: param.Type(),
			local: nil, param: param, offset: offset, storage: nil}
		ident.initStorage(true)
		f.identifiers[param.Name()] = &ident
		if ident.align() > ident.size() {
			offset += int(ident.align())
		} else {
			offset += int(ident.size())
		}
	}
	return asm, nil
}

func (f *Function) Func() (string, *Error) {
	if f.Trace {
		fmt.Printf("TRACE FUNC - %v\n", f.ssa.Name())
		fmt.Println("TRACE PARAMS")
	}
	params, err := f.Params()
	if err != nil {
		return params, err
	}
	if f.Trace {
		fmt.Println("TRACE {PARAMS}")
		fmt.Println("TRACE ZeroValues")
	}
	zeroRetValue, err := f.ZeroRetValue()
	if err != nil {
		return params + zeroRetValue, err
	}
	zeroSsaLocals, err := f.ZeroSsaLocals()
	if err != nil {
		return params + zeroRetValue + zeroSsaLocals, err
	}
	if f.Trace {
		fmt.Println("TRACE {ZeroValues}")
		fmt.Println("TRACE ComputePhi")
	}
	if err := f.computePhi(); err != nil {
		return "", err
	}
	if f.Trace {
		fmt.Println("TRACE {ComputePhi}")
		fmt.Println("TRACE BasicBlocks")
	}
	basicblocks, err := f.BasicBlocks()
	if err != nil {
		return params + zeroRetValue + zeroSsaLocals + basicblocks, err
	}
	if f.Trace {
		fmt.Println("TRACE {BasicBlocks}")
	}
	frameSize := f.localIdentsSize()
	frameSize = f.align(frameSize)
	argsSize := f.retOffset() + int(f.retSize())
	asm := params
	asm += f.SetStackPointer()
	asm += zeroRetValue
	asm += zeroSsaLocals
	asm += basicblocks
	asm = f.fixupRets(asm)
	asm = addIndent(asm, f.Indent)
	a := fmt.Sprintf("TEXT ·%v(SB),$%v-%v\n%v", f.outfname(), frameSize, argsSize, asm)
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

func (f *Function) GoProto() (string, string, string) {
	pkgname := "package " + f.ssa.Package().Pkg.Name() + "\n"
	imports := "import " + "\"github.com/bjwbell/gensimd/simd\"\n"
	sig := strings.TrimPrefix(f.ssa.Signature.String(), "func(")
	// TODO: HACK!
	if !strings.Contains(sig, "github.com/bjwbell/gensimd") {
		imports = ""
	}
	sig = strings.Replace(sig, "github.com/bjwbell/gensimd/", "", -1)
	fnproto := "func " + f.outfname() + "(" + sig + "\n"
	return pkgname, imports, fnproto
}

func (f *Function) outfname() string {
	if f.outfn != "" {
		return f.outfn
	}
	return f.ssa.Name()
}

func (f *Function) ZeroSsaLocals() (string, *Error) {

	asm := "// BEGIN ZeroSsaLocals\n"
	offset := int(0)
	locals := f.ssa.Locals
	ctx := context{f, nil}
	for _, local := range locals {
		if local.Heap {
			err := ErrorMsg2(fmt.Sprintf("Can't heap alloc local, name: %v", local.Name()))
			err.Pos = local.Pos()
			return "", err
		}
		sp := getRegister(REG_SP)

		//local values are always addresses, and have pointer types, so the type
		//of the allocated variable is actually
		//Type().Underlying().(*types.Pointer).Elem().
		typ := local.Type().Underlying().(*types.Pointer).Elem()
		size := sizeof(typ)
		localOffset := -(offset + int(size))
		asm += ZeroMemory(ctx, local.Name(), localOffset, size, sp)
		ident := identifier{f: f, name: local.Name(), typ: typ, local: local, param: nil, offset: localOffset}
		ident.initStorage(false)
		f.identifiers[local.Name()] = &ident
		if ident.align() > ident.size() {
			offset += int(ident.align())
		} else {
			offset += int(size)
		}
	}

	asm += "// END ZeroSsaLocals\n"
	return asm, nil
}

func (f *Function) newIdent(v ssa.Value) (identifier, *Error) {
	name := v.Name()
	typ := v.Type()
	size := sizeof(typ)
	offset := int(size)
	if align(typ) > size {
		offset = int(align(typ))
	}
	ident := identifier{
		f:      f,
		name:   name,
		typ:    typ,
		param:  nil,
		local:  nil,
		value:  v,
		offset: -int(f.localIdentsSize()) - offset}
	ident.initStorage(false)
	f.identifiers[name] = &ident
	// zeroing the memory is done at the beginning of the function
	return ident, nil
}

func (f *Function) ZeroRetValue() (string, *Error) {
	ctx := context{f, nil}
	asm := "// BEGIN ZeroRetValue\n"
	asm += ZeroMemory(ctx, retName(), f.retOffset(), f.retSize(), getRegister(REG_FP))
	asm += "// END ZeroRetValue\n"
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
	lastInstr := block.Instrs[len(block.Instrs)-1]
	ctx := context{f, lastInstr}

	if _, err := f.spillNonBlockRegisters(ctx); err != nil {
		return "", err
	}
	return asm, nil
}

func (f *Function) Instr(instr ssa.Instruction) (string, *Error) {

	if instr == nil {
		ice("nil instr")
	}
	asm := ""
	var err *Error

	errormsg := func(msg string) (string, *Error) {
		return "", &Error{Err: fmt.Errorf(msg), Pos: instr.Pos()}
	}
	if f.Trace {
		if _, ok := instr.(*ssa.DebugRef); ok {
			// Nothing to do
		} else {
			if v, ok := instr.(ssa.Value); ok {
				fmt.Printf("TRACE %v = %v\n", v.Name(), v)
			} else {
				fmt.Printf("TRACE %v\n", instr)
			}
		}
	}
	switch instr := instr.(type) {
	default:
		err = &Error{Err: fmt.Errorf("Unknown ssa instruction (type:%v): %v\n", reflect.TypeOf(instr), instr), Pos: instr.Pos()}
	case *ssa.Alloc:
		asm, err = f.AllocInstr(instr)
	case *ssa.BinOp:
		asm, err = f.BinOp(instr)
	case *ssa.Call:
		asm, err = f.Call(instr)
	case *ssa.ChangeInterface:
		asm, err = errormsg("converting interfaces unsupported")
	case *ssa.ChangeType:
		asm, err = errormsg("changing between types unsupported")
	case *ssa.Convert:
		asm, err = f.Convert(instr)
	case *ssa.DebugRef:
		// Nothing to do
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

	if err != nil && !err.Pos.IsValid() {
		err.Pos = instr.Pos()
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

func (f *Function) Len(call *ssa.Call) (string, *Error) {
	asm := fmt.Sprintf("// BEGIN Builtin.Len: %v\n", call)
	callcommon := call.Common()
	arg := callcommon.Args[0]
	ctx := context{f, call}
	if callcommon.IsInvoke() {
		panic(ice("len is a function, not a method"))
	}
	if len(callcommon.Args) != 1 {
		panic(ice(fmt.Sprintf("too many args (%v) for len", len(callcommon.Args))))
	}

	ident := f.Ident(call.Value())
	if reflectType(ident.typ).Name() != "int" {
		panic(ice(fmt.Sprintf("len returns int not (%v)", reflectType(ident.typ).Name())))
	}

	if isArray(arg.Type()) {

		length := reflectType(arg.Type()).Len()
		if length >= math.MaxInt32 {
			panic(ice(fmt.Sprintf("array too large (%v), maximum (%v)", length, math.MaxInt32)))
		}
		a, reg := f.allocIdentReg(call, ident, sizeof(ident.typ))
		asm += a
		asm += MovImm32Reg(ctx, int32(length), reg, false)
		a, err := f.StoreValue(call, ident, reg)
		asm += a
		if err != nil {
			return asm, err
		}

	} else if isSlice(arg.Type()) {
		a, err := f.SliceLen(call, arg, ident)
		asm += a
		if err != nil {
			return asm, err
		}

	} else {
		panic(ice(fmt.Sprintf("bad type (%v) passed to len", arg.Type())))
	}
	asm += fmt.Sprintf("// END Builtin.Len: %v\n", call)
	return asm, nil
}

func (f *Function) Builtin(call *ssa.Call, builtin *ssa.Builtin) (string, *Error) {
	obj := builtin.Object()
	if builtin.Name() == "len" && obj.String() == "builtin len" {
		return f.Len(call)
	} else {
		return ErrorMsg(fmt.Sprintf("builtin (%v) not supported", builtin.Name()))
	}
}

func (f *Function) Call(call *ssa.Call) (string, *Error) {
	common := call.Common()
	funct := common.Value
	if builtin, ok := funct.(*ssa.Builtin); ok {
		return f.Builtin(call, builtin)
	}
	if isSimdIntrinsic(call) {
		return f.SimdIntrinsic(call)
	}
	if sse2instr, ok := isSSE2Intrinsic(call); ok {
		return f.SSE2Intrinsic(call, sse2instr)
	}
	name := "UNKNOWN FUNC NAME"
	if call.Common().Method != nil {
		name = call.Common().Method.Name()
	} else if call.Common().StaticCallee() != nil {
		name = call.Common().StaticCallee().Name()
	}
	msg := fmt.Sprintf("function calls are not supported, func name (%v), description (%v)",
		name, call.Common().Description())
	return ErrorMsg(msg)

}

func (f *Function) SimdIntrinsic(call *ssa.Call) (string, *Error) {
	var asm string
	var err *Error

	args := call.Common().Args
	x := f.Ident(args[0])
	y := f.Ident(args[1])
	result := f.Ident(call)
	name := call.Common().StaticCallee().Name()
	if simdinstr, ok := getSimdInstr(name); ok {
		if result.typ != x.typ {
			panic(ice(fmt.Sprintf("Simd variable type (%v) and op type (%v)  dont match", result.typ.String(), x.typ.String())))
		}
		optypes := GetOpDataType(x.typ)
		a, e := packedOp(f, call, simdinstr, optypes.xmmvariant, y, x, result)
		asm = a
		err = e
	} else {
		intrinsic, ok := intrinsics[name]
		if !ok {
			panic(ice(fmt.Sprintf("Expected simd intrinsic got (%v)", name)))
		}
		a, e := intrinsic(f, call, x, y, result)
		asm = a
		err = e
	}
	asm = fmt.Sprintf("// BEGIN SIMD Intrinsic %v\n", call) + asm +
		fmt.Sprintf("// END SIMD Intrinsic %v\n", call)
	return asm, err
}

func isSimdIntrinsic(call *ssa.Call) bool {
	if call.Common() == nil || call.Common().StaticCallee() == nil {
		return false
	}
	name := call.Common().StaticCallee().Name()
	if _, ok := getSimdInstr(name); ok {
		return ok
	} else {
		_, ok := intrinsics[name]
		return ok
	}
}

func isSSE2Intrinsic(call *ssa.Call) (Intrinsic, bool) {
	if call.Common() == nil || call.Common().StaticCallee() == nil {
		return Intrinsic{}, false
	}
	name := call.Common().StaticCallee().Name()
	return getSSE2(name)
}

func (f *Function) SSE2Intrinsic(call *ssa.Call, sse2intrinsic Intrinsic) (string, *Error) {
	args := call.Common().Args
	return sse2Intrinsic(f, call, call, sse2intrinsic, args), nil
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
		fromType = OP_DATA
		toType = OP_DATA
		fromXmm = XMM_INVALID
		toXmm = XMM_INVALID
	} else if isInteger(from) && isFloat(to) {
		fromType = OP_DATA
		toType = OP_XMM
		fromXmm = XMM_INVALID
		toXmm = GetXmmVariant(to)
	} else if isFloat(from) && isInteger(to) {
		fromType = OP_XMM
		fromXmm = GetXmmVariant(from)
		toType = OP_DATA
		toXmm = XMM_INVALID
	} else if isFloat(from) && isFloat(to) {
		fromType = OP_XMM
		fromXmm = GetXmmVariant(from)
		toType = OP_XMM
		toXmm = GetXmmVariant(to)
	} else {
		fmtstr := "Cannot convert from (%v) to (%v)"
		msg := fmt.Sprintf(fmtstr, from.String(), to.String())
		return ErrorMsg(msg)
	}

	return f.ConvertFromTo(instr, fromType, toType, fromXmm, toXmm)
}

func (f *Function) ConvertFromTo(instr *ssa.Convert, fromOpType, toOpType InstrOpType, fromXmm, toXmm XmmData) (string, *Error) {

	asm := ""
	ctx := context{f, instr}
	a, from, err := f.LoadValueSimple(instr, instr.X)
	if err != nil {
		return "", err
	} else {
		asm += a
	}

	a1, to := f.allocIdentReg(instr, f.Ident(instr), f.Ident(instr).size())
	asm += a1

	a, tmp := f.allocReg(instr, DATA_REG, sizeInt())
	asm += a

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
			ice("converting from uint64 to float")
		}

		asm += f.ConvertUint64ToFloat(instr, tmp, from, to, floatSize)
	} else {
		asm += ConvertOp(ctx, from, fromType, to, toType, tmp)
	}
	toIdent := f.identifiers[instr.Name()]
	if a, err := f.StoreValue(instr, toIdent, to); err != nil {
		return "", err
	} else {
		asm += a
	}
	f.freeReg(from)
	f.freeReg(to)
	f.freeReg(tmp)

	return asm, nil
}

func (f *Function) ConvertUint64ToFloat(loc ssa.Instruction, tmp, regU64, regFloat *register, floatSize uint) string {

	asm := ""
	a, regI64 := f.allocReg(loc, DATA_REG, 8)
	asm += a
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
	str := "//           U64\n" +
		"CMPQ	     %v, $-1\n" +
		"// jmp to rounding\n" +
		"JEQ	     %v\n" +
		"//           U64\n" +
		"CMPQ	     %v, $-1\n" +
		"// jmp to no rounding\n" +
		"JGE	     %v\n" +
		"// rounding label\n" +
		"%v:\n" +
		"//           U64 I64\n" +
		"MOVQ	     %v, %v\n" +
		"//           I64\n" +
		"SHRQ	      $1, %v\n" +
		"//           U64 TMP\n" +
		"MOVQ	     %v, %v\n" +
		"//               TMP\n" +
		"ANDL	     $1, %v\n" +
		"//           TMP I64\n" +
		"ORQ	     %v, %v\n" +
		"//CVT        I64 XMM\n" +
		"%-9v    %v, %v\n" +
		"//ADD        XMM, XMM\n" +
		"%-9v    %v, %v\n" +
		"// jmp to end\n" +
		"JMP          %v\n" +
		"// no rounding label\n" +
		"%v:\n" +
		"//CVT        U64 XMM\n" +
		"%-9v    %v, %v\n" +
		"// end label\n" +
		"%v:\n"

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

	f.freeReg(regI64)

	return asm
}

func (f *Function) If(instr *ssa.If) (string, *Error) {
	asm := ""
	ctx := context{f, instr}
	tblock, fblock := -1, -1
	if instr.Block() != nil && len(instr.Block().Succs) == 2 {
		tblock = instr.Block().Succs[0].Index
		fblock = instr.Block().Succs[1].Index

	}
	if tblock == -1 || fblock == -1 {
		ice("malformed CFG with if stmt")
	}

	cond, ok := f.identifiers[instr.Cond.Name()]
	if !ok {

		return ErrorMsg(fmt.Sprintf("If: unhandled case, cond (%v)", instr.Cond))
	}
	a, reg, err := f.LoadIdentSimple(instr, cond)
	if err != nil {
		return "", err
	}
	asm += a

	a, err = f.JumpPreamble(instr, instr.Block().Index, fblock)
	if err != nil {
		return "", err
	}
	asm += a
	asm += CmpRegImm32(ctx, reg, uint32(0), cond.size())
	f.freeReg(reg)

	asm += fmt.Sprintf("%-9v    ", JEQ) + "block" + strconv.Itoa(fblock) + "\n"
	a, err = f.JumpPreamble(instr, instr.Block().Index, tblock)
	if err != nil {
		return "", err
	}
	asm += a
	jmp := "JMP"
	asm += fmt.Sprintf("%-9v    ", jmp) + "block" + strconv.Itoa(tblock) + "\n"
	asm = fmt.Sprintf("// BEGIN ssa.If, %v\n", instr) + asm
	asm += fmt.Sprintf("// END ssa.If, %v\n", instr)

	return asm, nil
}

func (f *Function) JumpPreamble(loc ssa.Instruction, blockIndex, jmpIndex int) (string, *Error) {
	asm := ""
	phiInfos := f.phiInfo[blockIndex][jmpIndex]
	for _, phiInfo := range phiInfos {
		ident := f.Ident(phiInfo.phi)
		ident.spilling = true
		if a, err := f.StoreValAddr(loc, phiInfo.value, ident); err != nil {
			return a, err
		} else {
			asm += a
		}
		if a, err := f.spillAllIdent(ident, loc); err != nil {
			return "", err
		} else {
			asm += a
		}
		ident.spilling = false
	}

	if a, e := f.spillRegisters(context{f, loc}); e != nil {
		return a, e
	} else {
		asm += a
	}

	asm = fmt.Sprintf("// BEGIN JumpPreamble block%v -> block%v\n", blockIndex, jmpIndex) + asm +
		fmt.Sprintf("// END JumpPreamble block%v -> block%v\n", blockIndex, jmpIndex)

	return asm, nil
}

func (f *Function) Jump(jmp *ssa.Jump) (string, *Error) {
	asm := ""
	block := -1
	if jmp.Block() != nil && len(jmp.Block().Succs) == 1 {
		block = jmp.Block().Succs[0].Index
	} else {
		ice("malformed CFG with jmp stmt")
	}
	a, err := f.JumpPreamble(jmp, jmp.Block().Index, block)
	if err != nil {
		return "", err
	}
	asm += a
	asm += "JMP block" + strconv.Itoa(block) + "\n"
	asm = "// BEGIN ssa.Jump\n" + asm
	asm += "// END ssa.Jump\n"
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
			ice("malformed CFG")
		}
		if _, ok := f.phiInfo[edgeBlock]; !ok {
			f.phiInfo[edgeBlock] = make(map[int][]phiInfo)
		}
		f.phiInfo[edgeBlock][blockIndex] = append(f.phiInfo[edgeBlock][blockIndex], phiInfo{value: edge, phi: phi})
	}
	return nil
}

func (f *Function) Phi(phi *ssa.Phi) (string, *Error) {

	if ident := f.Ident(phi); ident == nil {
		return ErrorMsg("Error in ssa.Phi allocation")
	}

	asm := fmt.Sprintf("// BEGIN ssa.Phi, name (%v), comment (%v), value (%v)\n", phi.Name(), phi.Comment, phi)
	asm += fmt.Sprintf("// END ssa.Phi, %v\n", phi)
	return asm, nil
}

var dummySpSize = uint32(math.MaxUint32)

func (f *Function) Return(ret *ssa.Return) (string, *Error) {
	asm := ResetStackPointer(dummySpSize)
	asm = "// BEGIN ssa.Return\n" + asm
	retIdent := f.retIdent()
	retIdent.spilling = true
	if a, err := f.CopyToRet(ret, ret.Results); err != nil {
		return "", err
	} else {
		asm += a
	}
	if a, e := f.spillAllIdent(retIdent, ret); e != nil {
		return a, e
	} else {
		asm += a
	}
	retIdent.spilling = false
	asm += Ret()
	asm += "// END ssa.Return\n"
	return asm, nil
}

func (f *Function) retIdent() *identifier {
	if ident, ok := f.identifiers[retName()]; !ok {
		retIdent :=
			identifier{
				f:      f,
				name:   retName(),
				typ:    f.retType(),
				local:  nil,
				param:  nil,
				value:  nil,
				offset: f.retOffset()}
		retIdent.initStorage(true)
		f.identifiers[retIdent.name] = &retIdent
		return &retIdent

	} else {
		return ident
	}
}

func (f *Function) CopyToRet(ret *ssa.Return, val []ssa.Value) (string, *Error) {
	if len(val) == 0 {
		return "", nil
	}
	if len(val) > 1 {
		e := ErrorMsg2("Multiple return values not supported")
		e.Pos = val[1].Pos()
		return "", e
	}
	retIdent := f.retIdent()
	return f.StoreValAddr(ret, val[0], retIdent)
}

func ResetStackPointer(size uint32) string {
	/*sp := getRegister(REG_SP)
	return AddImm32Reg(indent, size, sp)*/
	return ""
}

func (f *Function) fixupRets(asm string) string {
	old := ResetStackPointer(dummySpSize)
	new := ResetStackPointer(f.localIdentsSize())
	return strings.Replace(asm, old, new, -1)
}

func (f *Function) SetStackPointer() string {
	/*sp := getRegister(REG_SP)
	asm := SubImm32Reg(uint32(f.localsSize()), sp)
	return asm*/
	return ""
}

func (f *Function) StoreValAddr(loc ssa.Instruction, val ssa.Value, addr *identifier) (string, *Error) {

	if ident := f.Ident(val); ident == nil {
		ice("error in allocating local")
	}
	if addr.isConst() {
		ice(fmt.Sprintf("invalid addr \"%v\"", addr))
	}

	asm := ""
	asm += fmt.Sprintf("// BEGIN StoreValAddr addr name:%v, val name:%v\n", addr.name, val.Name()) + asm

	if isComplex(val.Type()) {
		return ErrorMsg("complex32/64 unsupported")
	} else if isXmm(val.Type()) {
		a, valReg, err := f.LoadValue(loc, val, 0, f.sizeof(val))
		if err != nil {
			return a, err
		}
		asm += a

		a, err = f.StoreValue(loc, addr, valReg)
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
				ice(fmt.Sprintf("Size (%v) not multiple of sizeInt (%v)", size, sizeInt()))
			}
		}
		for i := 0; i < int(iterations); i++ {
			offset := uint(i * datasize)
			a, valReg, err := f.LoadValue(loc, val, offset, uint(datasize))
			if err != nil {
				return a, err
			}
			asm += a
			a, err = f.AssignRegIdent(loc, valReg, addr, offset, uint(datasize))
			if err != nil {
				return a, err
			}
			asm += a
			f.freeReg(valReg)
		}
	}
	asm += fmt.Sprintf("// END StoreValAddr addr name:%v, val name:%v\n", addr.name, val.Name())
	return asm, nil
}

func (f *Function) Store(instr *ssa.Store) (string, *Error) {
	if ident := f.Ident(instr.Addr); ident == nil {
		return ErrorMsg(fmt.Sprintf("Cannot store value: %v", instr))
	}
	addr, ok := f.identifiers[instr.Addr.Name()]
	if !ok {
		ice(fmt.Sprintf("couldnt store identifier \"%v\"", addr.name))
	}
	asm := fmt.Sprintf("// BEGIN Store %v\n", instr)
	a, err := f.StoreValAddr(instr, instr.Val, addr)
	asm = asm + a + fmt.Sprintf("// END Store %v\n", instr)
	return asm, err
}

func (f *Function) BinOp(instr *ssa.BinOp) (string, *Error) {
	ctx := context{f, instr}
	ident := f.Ident(instr)
	if ident == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	var regX, regY, regVal *register
	size := f.sizeof(instr)
	xIsSigned := signed(instr.X.Type())

	var a string
	// comparison op results are size 1 byte, but that's not supported
	if size == 1 {
		a, regVal = f.allocIdentReg(instr, ident, 8*size)
	} else {
		a, regVal = f.allocIdentReg(instr, ident, size)
	}
	asm := a
	asmload, regX, regY, err := f.BinOpLoadXY(instr)
	asm += asmload

	if err != nil {
		return asm, err
	}

	size = f.sizeof(instr.X)

	switch instr.Op {
	default:
		ice(fmt.Sprintf("unknown op (%v)", instr.Op))
	case token.ADD, token.SUB, token.MUL, token.QUO, token.REM:
		optypes := GetOpDataType(instr.Type())
		asm += ArithOp(ctx, optypes, instr.Op, regX, regY, regVal)
	case token.AND, token.OR, token.XOR, token.SHL, token.SHR, token.AND_NOT:
		asm += BitwiseOp(ctx, instr.Op, xIsSigned, regX, regY, regVal, size)
	case token.EQL, token.NEQ, token.LEQ, token.GEQ, token.LSS, token.GTR:
		if size != f.sizeof(instr.Y) {
			ice("comparing two different size values")
		}
		optypes := GetOpDataType(instr.X.Type())
		asm += CmpOp(ctx, optypes, instr.Op, regX, regY, regVal)
	}
	f.freeReg(regX)
	f.freeReg(regY)

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		ice(fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr))
	}

	a1, err := f.StoreValue(instr, addr, regVal)
	if err != nil {
		return asm, err
	} else {
		asm += a1
	}
	f.freeReg(regVal)

	asm = fmt.Sprintf("// BEGIN ssa.BinOp, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf("// END ssa.BinOp, %v = %v\n", instr.Name(), instr)
	return asm, nil
}

func (f *Function) BinOpLoadXY(instr *ssa.BinOp) (asm string, x *register, y *register, err *Error) {
	if isPointer(instr.Type()) {
		panic("ptr")
	}
	if ident := f.Ident(instr); ident == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	if ident := f.Ident(instr.X); ident == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr.X))
	}
	if ident := f.Ident(instr.Y); ident == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr.Y))
	}

	var a string
	asm = "// BEGIN BinOpLoadXY\n"
	if isPointer(instr.X.Type()) {
		panic("ptr")
	}
	a, x, err = f.LoadValue(instr, instr.X, 0, f.sizeof(instr.X))
	x.inUse = true
	if err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}
	if isPointer(f.Ident(instr.Y).typ) {
		panic("ptr")
	}

	a, y, err = f.LoadValue(instr, instr.Y, 0, f.sizeof(instr.Y))
	y.inUse = true
	if err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}

	asm += "// END BinOpLoadXY\n"
	return asm, x, y, nil
}

func (f *Function) sizeof(val ssa.Value) uint {
	if _, ok := val.(*ssa.Const); ok {
		return f.sizeofConst(val.(*ssa.Const))
	}
	info, ok := f.identifiers[val.Name()]
	if !ok {
		ice(fmt.Sprintf("unknown name (%v), value (%v)\n", val.Name(), val))
	}
	_, _, size := info.Addr()
	return size
}

func (f *Function) sizeofConst(cnst *ssa.Const) uint {
	return sizeof(cnst.Type())
}

func (f *Function) SliceLen(loc ssa.Instruction, slice ssa.Value, ident *identifier) (string, *Error) {
	if _, ok := slice.Type().(*types.Slice); !ok {
		panic(ice(fmt.Sprintf("getting len of slice, type should slice not (%v)", slice.Type().String())))
	}

	asm := fmt.Sprintf("// BEGIN SliceLen: slice (%v), ident (%v)\n", slice, ident.String())

	a, reg, err := f.LoadValue(loc, slice, sliceLenOffset(), sliceLenSize())
	asm += a
	if err != nil {
		return asm, err
	}

	a, err = f.StoreValue(loc, ident, reg)
	asm += a
	if err != nil {
		return asm, err
	}
	f.freeReg(reg)
	asm += fmt.Sprintf("// END SliceLen: slice (%v), ident (%v)\n", slice, *ident)
	return asm, nil
}

func (f *Function) LoadSimdValue(loc ssa.Instruction, simdvalue ssa.Value) (string, *register, *Error) {
	asm := fmt.Sprintf("// BEGIN LoadSimdValue, simdvalue: %v\n", simdvalue)
	a, reg, err := f.LoadValue(loc, simdvalue, 0, f.sizeof(simdvalue))
	asm += a
	asm += fmt.Sprintf("// END LoadSimdValue, simdvalue: %v, reg %v\n", simdvalue, reg.name)
	return asm, reg, err
}

func (f *Function) LoadSimd(loc ssa.Instruction, ident *identifier) (string, *register, *Error) {
	asm := fmt.Sprintf("// BEGIN LoadSimd, ident: %v\n", ident.name)
	a, reg, err := f.LoadIdentSimple(loc, ident)
	asm += a
	if reg.typ != XMM_REG {
		a, xmmReg := f.allocReg(loc, XMM_REG, 16)
		asm += a
		asm += MovRegReg(context{f, loc}, OpDataType{OP_XMM, InstrData{}, XMM_F128}, reg, xmmReg, false)
		f.freeReg(reg)
		reg = xmmReg
	}
	asm += fmt.Sprintf("// END LoadSimd, ident: %v, reg %v\n", ident.name, reg.name)
	return asm, reg, err
}

func (f *Function) LoadSSE2(loc ssa.Instruction, ident *identifier) (string, *register, *Error) {
	return f.LoadSimd(loc, ident)
}

func (f *Function) LoadIdentSimple(loc ssa.Instruction, ident *identifier) (string, *register, *Error) {
	asm := fmt.Sprintf("// BEGIN LoadIdentSimple, ident: %v\n", ident.name)
	a, reg, err := f.LoadIdent(loc, ident, 0, sizeof(ident.typ))
	asm += a
	asm += fmt.Sprintf("// END LoadIdentSimple, ident: %v, reg %v\n", ident.name, reg.name)
	return asm, reg, err
}

func (f *Function) allocIdentReg(loc ssa.Instruction, ident *identifier, size uint) (string, *register) {
	asm, reg := f.allocReg(loc, regType(ident.typ), size)
	if size > 8 && !isXmm(ident.typ) {
		msg := "ident (%v), allocating more than 8 byte chunk"
		ice(fmt.Sprintf(msg, ident.name))
	}
	return asm, reg
}

func (f *Function) LoadIdent(loc ssa.Instruction, ident *identifier, offset uint, size uint) (string, *register, *Error) {
	asm, reg := ident.loadChunk(context{f, loc}, offset, size)
	return asm, reg, nil
}

func (f *Function) LoadValueSimple(loc ssa.Instruction, val ssa.Value) (string, *register, *Error) {
	asm := fmt.Sprintf("// BEGIN LoadValueSimple, val: %v\n", val)
	a, reg, err := f.LoadValue(loc, val, 0, f.sizeof(val))
	asm += a
	asm += fmt.Sprintf("// END LoadValueSimple, val: %v, reg %v\n", val, reg.name)
	return asm, reg, err
}

func (f *Function) LoadValue(loc ssa.Instruction, val ssa.Value, offset uint, size uint) (string, *register, *Error) {

	ident := f.Ident(val)
	asm := fmt.Sprintf("// BEGIN LoadValue, val %v (= %v), offset %v, size %v\n", val.Name(), val, offset, size)
	a, reg, err := f.LoadIdent(loc, ident, offset, size)
	if err != nil {
		return "", nil, err
	}
	asm += a
	asm += fmt.Sprintf("// END LoadValue, val %v (= %v), offset %v, size %v\n", val.Name(), val, offset, size)
	return asm, reg, nil
}

func (f *Function) StoreSimd(loc ssa.Instruction, reg *register, ident *identifier) (string, *Error) {
	return f.StoreValue(loc, ident, reg)
}

func (f *Function) StoreSSE2(loc ssa.Instruction, reg *register, ident *identifier) (string, *Error) {
	return f.StoreValue(loc, ident, reg)
}

func (f *Function) StoreValue(loc ssa.Instruction, ident *identifier, reg *register) (string, *Error) {
	if ident.size() > reg.size() {
		msgstr := "identifier, %v, size (%v) is greater than register size (%v)"
		ice(fmt.Sprintf(msgstr, ident.name, ident.size(), reg.size()))

	}
	return f.AssignRegIdent(loc, reg, ident, 0, ident.size())
}

func (f *Function) AssignRegIdent(loc ssa.Instruction, reg *register, ident *identifier, offset uint, size uint) (string, *Error) {
	return ident.newValue(context{f, loc}, reg, offset, size), nil
}

func (f *Function) UnOp(instr *ssa.UnOp) (string, *Error) {
	var err *Error
	asm := ""
	switch instr.Op {
	default:
		ice(fmt.Sprintf("unknown Op token (%v): \"%v\"", instr.Op, instr))
	case token.NOT: // logical negation
		asm, err = f.UnOpXor(instr, 1)
	case token.XOR: //bitwise negation
		asm, err = f.UnOpXor(instr, -1)
	case token.SUB: // arithmetic negation e.g. x=>-x
		asm, err = f.UnOpSub(instr)
	case token.MUL: //pointer indirection
		asm, err = f.UnOpPointer(instr)
	}
	asm = fmt.Sprintf("// BEGIN ssa.UnOp: %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf("// END ssa.UnOp: %v = %v\n", instr.Name(), instr)
	return asm, err
}

// bitwise negation
func (f *Function) UnOpXor(instr *ssa.UnOp, xorVal int32) (string, *Error) {
	ctx := context{f, instr}
	if ident := f.Ident(instr); ident == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		msg := fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr)
		ice(msg)
	}

	asm, reg, err := f.LoadValueSimple(instr, instr.X)
	if err != nil {
		return asm, err
	}

	if reg.size() < 8 {
		asm += XorImm32Reg(ctx, xorVal, reg, reg.size(), true)
	} else {
		asm += XorImm64Reg(ctx, int64(xorVal), reg, reg.size(), true)
	}

	a, err := f.StoreValue(instr, addr, reg)
	f.freeReg(reg)

	if err != nil {
		return asm, err
	} else {
		asm += a
	}

	asm = fmt.Sprintf("// BEGIN ssa.UnOpNot, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf("// END ssa.UnOpNot, %v = %v\n", instr.Name(), instr)
	return asm, nil
}

// arithmetic negation
func (f *Function) UnOpSub(instr *ssa.UnOp) (string, *Error) {
	ctx := context{f, instr}
	if ident := f.Ident(instr); ident == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	var regSubX, regX, regVal *register
	var a1, a2 string
	a1, regVal = f.allocReg(instr, regType(instr.Type()), f.sizeof(instr))
	a2, regSubX = f.allocReg(instr, regType(instr.X.Type()), f.sizeof(instr.X))

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		msg := fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr)
		ice(msg)
	}

	a3, regX, err := f.LoadValueSimple(instr, instr.X)
	asm := a1 + a2 + a3
	if err != nil {
		return asm, err
	}

	asm += ZeroReg(ctx, regSubX)
	optypes := GetOpDataType(instr.Type())
	asm += ArithOp(ctx, optypes, token.SUB, regSubX, regX, regVal)
	f.freeReg(regX)
	f.freeReg(regSubX)

	a, err := f.StoreValue(instr, addr, regVal)
	if err != nil {
		return asm, err
	} else {
		asm += a
	}

	f.freeReg(regVal)
	asm = fmt.Sprintf("// BEGIN ssa.UnOpSub, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf("// END ssa.UnOpSub, %v = %v\n", instr.Name(), instr)
	return asm, nil

}

//pointer indirection, in assignment such as "z = *x"
func (f *Function) UnOpPointer(instr *ssa.UnOp) (string, *Error) {
	asm := ""
	assignment := f.Ident(instr)
	if assignment == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	xName := instr.X.Name()
	xInfo, okX := f.identifiers[xName]
	if !xInfo.isSsaLocal() && xInfo.param == nil && xInfo.ptr == nil {
		panic("unexpected nil ptr")
	} else if !xInfo.isSsaLocal() && xInfo.param == nil {
		asm += xInfo.ptr.spillAllRegisters(instr)
	}
	// TODO add complex64/128 support
	if isComplex(instr.Type()) || isComplex(instr.X.Type()) {
		return ErrorMsg("complex64/complex128 unimplemented")
	}
	if !okX {
		msgstr := "Unknown name for UnOp X (%v), instr \"(%v)\""
		ice(fmt.Sprintf(msgstr, instr.X, instr))
	}
	if xInfo.local == nil && xInfo.param == nil && !xInfo.isPointer() {
		fmtstr := "in UnOp, X (%v) isn't a pointer, X.type (%v), instr \"(%v)\""
		msg := fmt.Sprintf(fmtstr, instr.X, instr.X.Type(), instr)
		ice(msg)
	}
	_, _, size := assignment.Addr()
	if xInfo.isSsaLocal() {
		ctx := context{f, instr}
		a, reg := xInfo.load(ctx)
		asm += a
		asm += assignment.newValue(ctx, reg, 0, xInfo.size())
		f.freeReg(reg)
	} else {
		var tmpData *register
		var a1 string
		if isXmm(instr.Type()) {
			a1, tmpData = f.allocReg(instr, XMM_REG, XmmRegSize)
		} else {
			a1, tmpData = f.allocReg(instr, regType(instr.Type()), DataRegSize)
		}
		asm += a1
		var a2 string
		var tmpAddr *register
		a2, tmpAddr = f.allocReg(instr, DATA_REG, DataRegSize)
		asm += a2

		src, ok := xInfo.storage.(*memory)
		if !ok {
			ice("cannot dereference pointer to constant")
		}
		dst, ok := assignment.storage.(*memory)
		if !ok {
			ice("cannot modify constant")
		}
		dst.removeAliases()
		ctx := context{f, instr}
		a, srcReg := src.load(ctx, src.ownerRegion())
		asm += a
		aReg, aOffset, _ := assignment.Addr()
		asm += MovRegIndirectMem(ctx, dst.optype(), srcReg, assignment.name, aOffset, &aReg, size, tmpAddr, tmpData)
		dst.setInitialized(region{0, size})
		f.freeReg(srcReg)
		f.freeReg(tmpAddr)
		f.freeReg(tmpData)
	}

	asm = fmt.Sprintf("// BEGIN ssa.UnOpPointer, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf("// END ssa.UnOpPointer, %v = %v\n", instr.Name(), instr)
	return asm, nil
}

func (f *Function) Index(instr *ssa.Index) (string, *Error) {
	panic("index instruction not implemented")
	/*if instr == nil {
		return ErrorMsg("nil instr")
	}
	asm := ""
	xInfo := f.identifiers[instr.X.Name()]
	assignment := f.Ident(instr)
	if assignment == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	xReg, xOffset, _ := xInfo.Addr()
	a1, addr := f.allocIdentReg(instr, assignment, assignment.size())
	asm += a1
	a, idx, err := f.LoadValueSimple(instr, instr.Index)
	if err != nil {
		return "", err
	}
	asm += a

	asm += Lea(xInfo.name, xOffset, &xReg, addr)

	optypes := GetIntegerOpDataType(false, idx.width/8)
	asm += AddRegReg(optypes, idx, addr)

	if a, e := f.StoreValue(instr, assignment, addr); e != nil {
		return "", e
	} else {
		asm += a
	}

	f.freeReg(idx)
	f.freeReg(addr)

	asm = fmt.Sprintf("// BEGIN ssa.Index: %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf("// END ssa.Index: %v = %v\n", instr.Name(), instr)

	return asm, nil*/
}

func (f *Function) IndexAddr(instr *ssa.IndexAddr) (string, *Error) {
	ctx := context{f, instr}
	if instr == nil {
		return ErrorMsg("nil instr")

	}

	asm := ""
	xInfo := f.identifiers[instr.X.Name()]
	assignment := f.Ident(instr)
	assignment.ptr = xInfo
	if a, e := f.spillAllIdent(xInfo, instr); e != nil {
		return a, e
	} else {
		asm += a
	}

	xReg, xOffset, _ := xInfo.Addr()
	a1, addr := f.allocIdentReg(instr, assignment, assignment.size())
	asm += a1

	a, idx, err := f.LoadValueSimple(instr, instr.Index)
	if err != nil {
		return "", err
	}
	asm += a

	asm += MulImm32RegReg(ctx, uint32(sizeofElem(xInfo.typ)), idx, idx, true)

	if isSlice(xInfo.typ) {
		// TODO: add bounds checking
		optypes := GetIntegerOpDataType(false, sizePtr())
		asm += MovMemReg(ctx, optypes, xInfo.name, xOffset, &xReg, addr, false)
	} else if isArray(xInfo.typ) {
		asm += Lea(ctx, xInfo.name, xOffset, &xReg, addr, false)
		//assignment.aliases = append(assignment.aliases, xInfo)
	} else if isSimd(xInfo.typ) {
		asm += Lea(ctx, xInfo.name, xOffset, &xReg, addr, false)
		//assignment.aliases = append(assignment.aliases, xInfo)
	} else {
		ice(fmt.Sprintf("indexing non-slice/array variable, type %v", xInfo.typ))
	}

	optypes := GetIntegerOpDataType(false, idx.size())
	asm += AddRegReg(ctx, optypes, idx, addr, false)

	a, e := f.StoreValue(instr, assignment, addr)
	if e != nil {
		return "", e
	}
	asm += a

	f.freeReg(idx)
	f.freeReg(addr)

	asm = fmt.Sprintf("// BEGIN ssa.IndexAddr: %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf("// END ssa.IndexAddr: %v = %v\n", instr.Name(), instr)
	return asm, nil
}

func (f *Function) AllocInstr(instr *ssa.Alloc) (string, *Error) {
	asm := ""
	if instr == nil {
		return ErrorMsg("AllocInstr: nil instr")

	}
	if instr.Heap {
		msg := "Heap allocations are unsupported (are all print and log statements removed?), ssa variable: %v, type: %v"
		msgstr := fmt.Sprintf(msg, instr.Name(), instr.Type())
		return ErrorMsg(msgstr)
	}

	//Alloc values are always addresses, and have pointer types, so the type
	//of the allocated variable is actually
	//Type().Underlying().(*types.Pointer).Elem().
	info := f.identifiers[instr.Name()]
	if info.local == nil {
		ice(fmt.Sprintf("expected %v to be a local variable", instr.Name()))
	}
	if _, ok := info.typ.(*types.Pointer); ok {
	} else {
	}
	f.identifiers[instr.Name()] = info
	return asm, nil
}

func (f *Function) localIdentsSize() uint32 {
	size := uint32(0)
	for _, ident := range f.identifiers {
		if !ident.isConst() && !ident.isParam() && !ident.isRetIdent() {
			size += uint32(ident.size())
		}
	}
	return size
}

func (f *Function) init() *Error {
	for _, r := range registers {
		f.registers = append(f.registers, r)
	}
	f.identifiers = make(map[string]*identifier)
	f.phiInfo = make(map[int]map[int][]phiInfo)
	return nil
}

func (f *Function) spillDirtyIdent(ident *identifier, loc ssa.Instruction) (string, *Error) {
	return ident.spillDirtyRegisters(loc), nil
}

func (f *Function) spillAllIdent(ident *identifier, loc ssa.Instruction) (string, *Error) {
	return ident.spillAllRegisters(loc), nil
}

func (f *Function) spillRegisters(ctx context) (string, *Error) {
	asm := ""
	for i := 0; i < len(f.registers); i++ {
		r := &f.registers[i]
		if f.excludeReg(r) {
			continue
		}
		asm += r.spill(ctx)
	}
	return asm, nil
}

func (f *Function) spillNonBlockRegisters(ctx context) (string, *Error) {
	asm := ""
	for i := 0; i < len(f.registers); i++ {
		r := &f.registers[i]
		if f.excludeReg(r) {
			continue
		}
		blockLocal := false
		if r.parent != nil {
			blockLocal = r.parent.owner().isBlockLocal()
		}
		if !blockLocal {
			asm += r.spill(ctx)
		}
	}
	return asm, nil
}

// allocReg allocates a register, if none are available a victim register is spilled, loc is the instr location,
// where the allocation is taking place (this is used to choose the victim register when spilling),
// t is the register type eg XMM_REG, size is bytes.
func (f *Function) allocReg(loc ssa.Instruction, t RegType, size uint) (string, *register) {
	asm := ""
	reg := f.allocUnusedReg(t, size)
	if reg == nil {
		reg = f.chooseVictim(t, size)
		if reg == nil {
			for _, r := range f.registers {
				fmt.Println("R: ", r)
				fmt.Println("R.inUse: ", r.inUse)
				fmt.Println("R IS ALIAS: ", r.parent != nil)
			}

			ice(fmt.Sprintf("can't alloc register (t: %v, size: %v)", t, size))

		}
		if reg.inUse {
			ice(fmt.Sprintf("can't alloc register (t: %v, size: %v)", t, size))
		}
		if !f.regDead(reg) {
			parent := reg.parent
			a := reg.spill(context{f, loc})
			if a != "" {

				if f.PrintSpills || f.Trace {
					fmt.Printf("Spilling %v\n", reg.name)
				}
				asm = fmt.Sprintf("// BEGIN RegSpill %v, ident %v\n", reg.name, parent.owner().name)
				asm += a
				asm += fmt.Sprintf("// END RegSpill %v, ident %v\n", reg.name, parent.owner().name)
			}

		}

	}
	reg.inUse = true
	reg.dirty = false
	reg.parent = nil
	return asm, reg
}

func (f *Function) allocTempReg(t RegType, size uint) (string, *register) {
	reg := f.allocUnusedReg(t, size)
	if reg != nil {
		return "", reg
	}
	asm := ""
	reg = f.chooseVictim(t, size)
	if reg == nil {

		for _, r := range f.registers {
			fmt.Println("R: ", r)
			fmt.Println("R IS ALIAS: ", r.parent != nil)
		}

		ice(fmt.Sprintf("can't alloc register (t: %v, size: %v)", t, size))
	}
	if !f.regDead(reg) {
		parent := reg.parent
		a := reg.spill(context{f, nil})
		if a != "" {
			if f.PrintSpills {
				fmt.Printf("Spilling %v\n", reg.name)
			}
			asm = fmt.Sprintf("// BEGIN RegSpill %v, ident %v\n", reg.name, parent.owner().name)
			asm += a
			asm += fmt.Sprintf("// END RegSpill %v, ident %v\n", reg.name, parent.owner().name)
		}

	}

	reg.inUse = true
	reg.dirty = false
	reg.parent = nil
	return asm, reg
}

func (f *Function) allocUnusedReg(t RegType, size uint) *register {
	var reg *register
	for i := 0; i < len(f.registers); i++ {
		r := &f.registers[i]
		if f.excludeReg(r) {
			continue
		}
		used := r.inUse || r.parent != nil
		if used || r.typ != t {
			continue
		}
		for i := range r.datasizes {
			if r.datasizes[i] == size {
				reg = r
				break
			}
		}
	}
	if reg != nil {
		reg.inUse = true
		reg.dirty = false
		reg.parent = nil
	}
	return reg
}

func (f *Function) chooseVictim(t RegType, size uint) *register {
	var victim *register
	for i := 0; i < len(f.registers); i++ {
		r := &f.registers[i]
		if f.excludeReg(r) {
			continue
		}
		if r.inUse || r.typ != t {
			continue
		}
		sizeMatch := false
		for i := range r.datasizes {
			if r.datasizes[i] == size {
				sizeMatch = true
				break
			}
		}
		if !sizeMatch {
			continue
		}
		if victim == nil {
			victim = r
		} else {
			if f.compareVictims(r, victim) >= 1 {
				victim = r
			}
		}
	}
	if victim == nil {
		ice(fmt.Sprintf("no eligible registers to spill!\n"))
	}
	return victim
}

func (f *Function) compareVictims(vic1, vic2 *register) int {
	if !vic1.inUse && vic2.inUse {
		return 1
	} else if vic1.inUse && !vic2.inUse {
		return -1
	} else if vic1.parent == vic2.parent {
		return 0
	} else if vic1.parent == nil {
		return 1
	} else {
		return -1
	}
}

func (f *Function) regDead(r *register) bool {
	return !r.inUse && r.parent == nil
}

func (f *Function) excludeReg(reg *register) bool {
	for _, r := range excludedRegisters {
		if r.name == reg.name {
			return true
		}
	}
	return false
}

func (f *Function) instrIdxInBlock(instr ssa.Instruction) int {
	block := instr.Block()
	for i := 0; i < len(block.Instrs); i++ {
		if block.Instrs[i] == instr {
			return i
		}
	}
	panic(ice("couldn't find instruction location in basic block"))
}

func (f *Function) isAlive(loc ssa.Instruction, val ssa.Value) bool {
	block := loc.Block()
	instrIdx := f.instrIdxInBlock(loc)
	if val.Referrers() == nil {
		return false
	}

	for _, ref := range *val.Referrers() {
		bk := ref.Block()
		if bk.Index != block.Index {
			continue
		}
		for i, _ := range bk.Instrs {
			if i > instrIdx {
				return true
			}
		}
	}
	return false
}

// zeroReg returns the assembly for zeroing the passed in register
func (f *Function) zeroReg(r *register) string {
	ctx := context{f, nil}
	return ZeroReg(ctx, r)
}

func (f *Function) freeReg(reg *register) {
	reg.inUse = false
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

// retSize returns the size of the return value in bytes
func (f *Function) retSize() uint {
	size := sizeof(f.retType())
	return size
}

// retOffset returns the offset of the return value in bytes
func (f *Function) retOffset() int {
	align := f.retAlign()
	// TODO: FIX
	// HACK!!!
	if isSimd(f.retType()) || isSSE2(f.retType()) {
		align = 8
	}
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

func (f *Function) Ident(v ssa.Value) *identifier {

	if ident, ok := f.identifiers[v.Name()]; ok {
		if ident == nil {
			ice("nil ident")
		}
		return ident
	}

	switch v := v.(type) {
	case *ssa.Const:
		ident := identifier{f: f, name: v.Name(), typ: v.Type(), local: nil, param: nil, cnst: v}
		ident.initStorage(true)
		f.identifiers[v.Name()] = &ident
		return &ident
	}

	local, err := f.newIdent(v)
	if err != nil {
		return nil
	}

	f.identifiers[v.Name()] = &local

	return &local
}

type WrappedVal struct {
	ssa.Value
}

func (wv WrappedVal) Operands(rands []*ssa.Value) []*ssa.Value {
	return nil
}
