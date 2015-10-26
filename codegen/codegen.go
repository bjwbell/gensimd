package codegen

import (
	"errors"
	"fmt"
	"go/token"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/bjwbell/gensimd/codegen/sse2"

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
	Indent      string
	identifiers map[string]identifier
	jmpLabels   []string
	outfn       string // output function name

	// map from block index to the successor block indexes that need phi vars set
	phiInfo map[int]map[int][]phiInfo

	// maps register to false if unused and true if used
	registers map[string]bool

	ssa *ssa.Function
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

func CreateFunction(fn *ssa.Function, outfn string, debug bool) (*Function, *Error) {
	if fn == nil {
		return nil, ErrorMsg2("Nil function passed in")
	}
	f := Function{ssa: fn, outfn: outfn, Debug: debug}
	f.Indent = "        "
	f.init()
	return &f, nil
}

func AssemblyFilePreamble() string {
	return "// +build amd64\n\n"
}

func (f *Function) GoAssembly() (string, *Error) {
	asm, err := f.Func()
	if !f.Debug {
		asm = StripDebugComments(asm, f.Indent)
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
		param := paramInfo{name: p.Name(), info: p}
		// TODO alloc reg based on other param types
		if _, ok := p.Type().(*types.Slice); ok {
			param.extra = paramSlice{lenOffset: offset + int(sizePtr())}
		} else if basic, ok := p.Type().(*types.Basic); ok {
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
	argsSize := f.retOffset() + int(f.retSize())
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

func (f *Function) GoProto() (string, string, string) {
	pkgname := "package " + f.ssa.Package().Pkg.Name() + "\n"
	imports := "import " + "\"github.com/bjwbell/gensimd/simd\"\n"
	sig := strings.TrimPrefix(f.ssa.Signature.String(), "func(")
	sig = strings.Replace(sig, "github.com/bjwbell/gensimd/", "", -1)
	fnproto := "func " + f.outfname() + "(" + sig
	return pkgname, imports, fnproto
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
	asm := fmt.Sprintf(f.Indent+"// BEGIN Builtin.Len: %v\n", call)
	callcommon := call.Common()
	arg := callcommon.Args[0]

	if callcommon.IsInvoke() {
		panic(internal("len is a function, not a method"))
	}
	if len(callcommon.Args) != 1 {
		panic(internal(fmt.Sprintf("too many args (%v) for len", len(callcommon.Args))))
	}

	ident := f.Ident(call.Value())
	if reflectType(ident.typ).Name() != "int" {
		panic(internal(fmt.Sprintf("len returns int not (%v)", reflectType(ident.typ).Name())))
	}

	if isArray(arg.Type()) {

		length := reflectType(arg.Type()).Len()
		if length >= math.MaxInt32 {
			panic(internal(fmt.Sprintf("array too large (%v), maximum (%v)", length, math.MaxInt32)))
		}
		reg := f.allocIdentReg(ident, sizeof(ident.typ), arg.Pos())
		asm += MovImm32Reg(f.Indent, int32(length), reg)
		a, err := f.StoreValue(ident, reg, arg.Pos())
		asm += a
		if err != nil {
			return asm, err
		}

	} else if isSlice(arg.Type()) {
		a, err := f.SliceLen(arg, ident)
		asm += a
		if err != nil {
			return asm, err
		}

	} else {
		panic(internal(fmt.Sprintf("bad type (%v) passed to len", arg.Type())))
	}
	asm += fmt.Sprintf(f.Indent+"// END Builtin.Len: %v\n", call)
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
	var a string
	var e *Error

	args := call.Common().Args
	x := f.Ident(args[0])
	y := f.Ident(args[1])
	result := f.Ident(call)
	name := call.Common().StaticCallee().Name()
	if simdinstr, ok := getSimdInstr(name); ok {
		if result.typ != x.typ {
			panic(internal(fmt.Sprintf("Simd variable type (%v) and op type (%v)  dont match", result.typ.String(), x.typ.String())))
		}
		optypes := GetOpDataType(x.typ)
		a, e := packedOp(f, simdinstr, optypes.xmmvariant, y, x, result, call.Pos())
		return a, e
	} else {
		intrinsic, ok := intrinsics[name]
		if !ok {
			panic(internal(fmt.Sprintf("Expected simd intrinsic got (%v)", name)))
		}
		a, e = intrinsic(f, x, y, result, call.Pos())
		return a, e
	}
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

func isSSE2Intrinsic(call *ssa.Call) (sse2.SSE2Instr, bool) {
	if call.Common() == nil || call.Common().StaticCallee() == nil {
		return sse2.INVALID, false
	}
	name := call.Common().StaticCallee().Name()
	return getSSE2(name)
}

func (f *Function) SSE2Intrinsic(call *ssa.Call, sse2Instr sse2.SSE2Instr) (string, *Error) {
	var a string
	var e *Error

	args := call.Common().Args
	x := f.Ident(args[0])
	y := f.Ident(args[1])
	result := f.Ident(call)
	a, e = sse2Op(f, sse2Instr, x, y, result, call.Pos())
	return a, e
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

func (f *Function) ConvertFromTo(instr *ssa.Convert, fromOpType, toOpType InstrOpType, fromXmm, toXmm XmmData) (string, *Error) {

	asm := ""

	a, from, err := f.LoadValueSimple(instr.X)
	if err != nil {
		return "", err
	} else {
		asm += a
	}

	to := f.allocIdentReg(f.Ident(instr), f.Ident(instr).size(), instr.Pos())

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

		asm += f.ConvertUint64ToFloat(&tmp, from, to, floatSize)
	} else {
		asm += ConvertOp(f.Indent, from, fromType, to, toType, &tmp)
	}
	toNameInfo := f.identifiers[instr.Name()]
	if a, err := f.StoreValue(&toNameInfo, to, instr.Pos()); err != nil {
		return "", err
	} else {
		asm += a
	}
	f.freeReg(*from)
	f.freeReg(*to)
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
	str := f.Indent + "//           U64\n" +
		f.Indent + "CMPQ	     %v, $-1\n" +
		f.Indent + "// jmp to rounding\n" +
		f.Indent + "JEQ	     %v\n" +
		f.Indent + "//           U64\n" +
		f.Indent + "CMPQ	     %v, $-1\n" +
		f.Indent + "// jmp to no rounding\n" +
		f.Indent + "JGE	     %v\n" +
		f.Indent + "// rounding label\n" +
		f.Indent + "%v:\n" +
		f.Indent + "//           U64 I64\n" +
		f.Indent + "MOVQ	     %v, %v\n" +
		f.Indent + "//           I64\n" +
		f.Indent + "SHRQ	      $1, %v\n" +
		f.Indent + "//           U64 TMP\n" +
		f.Indent + "MOVQ	     %v, %v\n" +
		f.Indent + "//               TMP\n" +
		f.Indent + "ANDL	     $1, %v\n" +
		f.Indent + "//           TMP I64\n" +
		f.Indent + "ORQ	     %v, %v\n" +
		f.Indent + "//CVT        I64 XMM\n" +
		f.Indent + "%-9v    %v, %v\n" +
		f.Indent + "//ADD        XMM, XMM\n" +
		f.Indent + "%-9v    %v, %v\n" +
		f.Indent + "// jmp to end\n" +
		f.Indent + "JMP          %v\n" +
		f.Indent + "// no rounding label\n" +
		f.Indent + "%v:\n" +
		f.Indent + "//CVT        U64 XMM\n" +
		f.Indent + "%-9v    %v, %v\n" +
		f.Indent + "// end label\n" +
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
		jeq := "JEQ"
		asm += f.Indent + fmt.Sprintf("%-9v    ", jeq) + "block" + strconv.Itoa(fblock) + "\n"
		a, err = f.JumpPreamble(instr.Block().Index, tblock)
		if err != nil {
			return "", err
		}
		asm += a
		jmp := "JMP"
		asm += f.Indent + fmt.Sprintf("%-9v    ", jmp) + "block" + strconv.Itoa(tblock) + "\n"

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

	if nameinfo := f.Ident(phi); nameinfo == nil {
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

	return f.StoreValAddr(val[0], &retAddr, val[0].Pos())
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

func (f *Function) StoreValAddr(val ssa.Value, addr *identifier, pos token.Pos) (string, *Error) {

	if nameinfo := f.Ident(val); nameinfo == nil {
		internal("error in allocating local")
	}
	if addr.local == nil && addr.param == nil {
		internal(fmt.Sprintf("invalid addr \"%v\"", addr))
	}

	asm := ""
	asm += f.Indent + fmt.Sprintf("// BEGIN StoreValAddr addr name:%v, val name:%v\n", addr.name, val.Name()) + asm

	if isComplex(val.Type()) {
		return ErrorMsg("complex32/64 unsupported")
	} else if isFloat(val.Type()) {

		a, valReg, err := f.LoadValue(val, 0, f.sizeof(val))
		if err != nil {
			return a, err
		}
		asm += a

		a, err = f.StoreValue(addr, valReg, pos)
		if err != nil {
			return a, err
		}
		asm += a
		f.freeReg(*valReg)

	} else if isSimd(val.Type()) || isSSE2(val.Type()) {

		a, valReg, err := f.LoadSimdValue(val)
		if err != nil {
			return a, err
		}
		asm += a
		a, err = f.StoreSimd(valReg, addr, pos)
		if err != nil {
			return a, err
		}
		asm += a
		f.freeReg(*valReg)

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

		for i := int(0); i < int(iterations); i++ {
			offset := i * datasize
			a, valReg, err := f.LoadValue(val, offset, uint(datasize))
			if err != nil {
				return a, err
			}
			asm += a
			a, err = f.StoreReg(valReg, addr, offset, uint(datasize))
			if err != nil {
				return a, err
			}
			asm += a
			f.freeReg(*valReg)
		}
	}

	asm += f.Indent +
		fmt.Sprintf("// END StoreValAddr addr name:%v, val name:%v\n",
			addr.name, val.Name())
	return asm, nil
}

func (f *Function) Store(instr *ssa.Store) (string, *Error) {
	if nameinfo := f.Ident(instr.Addr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot store value: %v", instr))
	}
	addr, ok := f.identifiers[instr.Addr.Name()]
	if !ok {
		internal(fmt.Sprintf("couldnt store identifier \"%v\"", addr.name))
	}
	asm := f.Indent + fmt.Sprintf("// BEGIN Store, instr: %v\n", instr)
	a, err := f.StoreValAddr(instr.Val, &addr, instr.Pos())
	asm = asm + a + f.Indent + fmt.Sprintf("// END Store, instr: %v\n", instr)
	return asm, err
}

func (f *Function) BinOp(instr *ssa.BinOp) (string, *Error) {

	nameinfo := f.Ident(instr)
	if nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	var regX, regY, regVal *register
	size := f.sizeof(instr)
	xIsSigned := signed(instr.X.Type())

	// comparison op results are size 1 byte, but that's not supported
	if size == 1 {
		regVal = f.allocIdentReg(nameinfo, 8*size, instr.Pos())
	} else {
		regVal = f.allocIdentReg(nameinfo, size, instr.Pos())
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
		optypes := GetOpDataType(instr.Type())
		asm += ArithOp(f.Indent, optypes, instr.Op, regX, regY, regVal)
	case token.AND, token.OR, token.XOR, token.SHL, token.SHR, token.AND_NOT:
		asm += BitwiseOp(f.Indent, instr.Op, xIsSigned, regX, regY, regVal, size)
	case token.EQL, token.NEQ, token.LEQ, token.GEQ, token.LSS, token.GTR:
		if size != f.sizeof(instr.Y) {
			internal("comparing two different size values")
		}
		optypes := GetOpDataType(instr.X.Type())
		asm += CmpOp(f.Indent, optypes, instr.Op, regX, regY, regVal)
	}

	f.freeReg(*regX)
	f.freeReg(*regY)

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		internal(fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr))
	}

	a, err := f.StoreValue(&addr, regVal, instr.Pos())
	if err != nil {
		return asm, err
	} else {
		asm += a
	}
	f.freeReg(*regVal)

	asm = fmt.Sprintf(f.Indent+"// BEGIN ssa.BinOp, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf(f.Indent+"// END ssa.BinOp, %v = %v\n", instr.Name(), instr)
	return asm, nil
}

func (f *Function) BinOpLoadXY(instr *ssa.BinOp) (asm string, x *register, y *register, err *Error) {

	if nameinfo := f.Ident(instr); nameinfo == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	if nameinfo := f.Ident(instr.X); nameinfo == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr.X))
	}
	if nameinfo := f.Ident(instr.Y); nameinfo == nil {
		return "", nil, nil, ErrorMsg2(fmt.Sprintf("Cannot alloc value: %v", instr.Y))
	}

	var a string

	asm = f.Indent + "// BEGIN BinOpLoadXY\n"

	a, x, err = f.LoadValue(instr.X, 0, f.sizeof(instr.X))
	if err != nil {
		return "", nil, nil, err
	} else {
		asm += a
	}

	a, y, err = f.LoadValue(instr.Y, 0, f.sizeof(instr.Y))
	if err != nil {
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

func (f *Function) SliceLen(slice ssa.Value, ident *identifier) (string, *Error) {
	if _, ok := slice.Type().(*types.Slice); !ok {
		panic(internal(fmt.Sprintf("getting len of slice, type should slice not (%v)", slice.Type().String())))
	}

	asm := f.Indent + fmt.Sprintf("// BEGIN SliceLen: slice (%v), ident (%v)\n", slice, *ident)

	a, reg, err := f.LoadValue(slice, sliceLenOffset(), sliceLenSize())
	asm += a
	if err != nil {
		return asm, err
	}

	a, err = f.StoreValue(ident, reg, slice.Pos())
	asm += a
	if err != nil {
		return asm, err
	}
	asm += f.Indent + fmt.Sprintf("// END SliceLen: slice (%v), ident (%v)\n", slice, *ident)
	return asm, nil
}

func (f *Function) LoadSimdValue(simdvalue ssa.Value) (string, *register, *Error) {
	asm := f.Indent + fmt.Sprintf("// BEGIN LoadSimdValue, simdvalue: %v\n", simdvalue)
	a, reg, err := f.LoadValue(simdvalue, 0, f.sizeof(simdvalue))
	asm += a
	asm += f.Indent + fmt.Sprintf("// END LoadSimdValue, simdvalue: %v, reg %v\n", simdvalue, reg.name)
	return asm, reg, err
}

func (f *Function) LoadSimd(ident *identifier, pos token.Pos) (string, *register, *Error) {
	asm := f.Indent + fmt.Sprintf("// BEGIN LoadSimd, ident: %v\n", ident.name)
	a, reg, err := f.LoadIdentSimple(ident, pos)
	asm += a
	if reg.typ != XMM_REG {
		xmmReg := f.allocReg(XMM_REG, 16)
		asm += MovRegReg(f.Indent, OpDataType{XMM_OP, InstrData{}, XMM_F128}, reg, &xmmReg)
		f.freeReg(*reg)
		reg = &xmmReg
	}

	asm += f.Indent + fmt.Sprintf("// END LoadSimd, ident: %v, reg %v\n", ident.name, reg.name)
	return asm, reg, err
}

func (f *Function) LoadSSE2(ident *identifier, pos token.Pos) (string, *register, *Error) {
	return f.LoadSimd(ident, pos)
}

func (f *Function) LoadIdentSimple(ident *identifier, pos token.Pos) (string, *register, *Error) {
	asm := f.Indent + fmt.Sprintf("// BEGIN LoadIdentSimple, ident: %v\n", ident.name)
	a, reg, err := f.LoadIdent(ident, 0, sizeof(ident.typ), pos)
	asm += a
	asm += f.Indent + fmt.Sprintf("// END LoadIdentSimple, ident: %v, reg %v\n", ident.name, reg.name)
	return asm, reg, err
}

func (f *Function) allocIdentReg(ident *identifier, size uint, pos token.Pos) *register {

	reg := f.allocReg(regType(ident.typ), size)
	if size > 8 && ((!isSimd(ident.typ) && !isSSE2(ident.typ)) || reg.typ != XMM_REG) {
		msg := "ident (%v), allocating more than 8 byte chunk"
		internal(fmt.Sprintf(msg, ident.name))
	}
	if isSimd(ident.typ) || isSSE2(ident.typ) || isIntegerSSE2(ident.typ) {
		return &reg
	}

	return &reg

}

func (f *Function) LoadIdent(ident *identifier, offset int, size uint, pos token.Pos) (string, *register, *Error) {
	r, roffset, rsize := ident.Addr()
	if rsize%size != 0 {
		msg := "ident (%v), size (%v) not a multiple of chunk size (%v)"
		internal(fmt.Sprintf(msg, ident.name, ident.size(), size))
	}
	reg := f.allocIdentReg(ident, size, pos)
	if size > 8 && ((!isSimd(ident.typ) && !isSSE2(ident.typ)) || reg.typ != XMM_REG) {
		msg := "ident (%v), loading more than 8 byte chunk"
		internal(fmt.Sprintf(msg, ident.name))
	}
	if isSimd(ident.typ) || isSSE2(ident.typ) || isIntegerSSE2(ident.typ) {
		optypes := GetOpDataType(ident.typ)
		asm := f.Indent + fmt.Sprintf("// BEGIN LoadIdent (SIMD), ident: %v, offset %v, size %v\n", ident.name, offset, size)
		a := ""
		if isIntegerSimd(ident.typ) || isIntegerSSE2(ident.typ) {
			a = MovIntegerSimdMemReg(f.Indent, optypes, ident.name, roffset+offset, &r, reg)
		} else {
			a = MovMemReg(f.Indent, optypes, ident.name, roffset+offset, &r, reg)
		}

		asm += a
		asm += f.Indent + fmt.Sprintf("// END LoadIdent (SIMD), ident: %v, offset %v, size %v\n", ident.name, offset, size)
		return asm, reg, nil
	}

	optypes := GetIntegerOpDataType(false, size)
	asm := f.Indent + fmt.Sprintf("// BEGIN LoadIdent, ident: %v, offset %v, size %v\n", ident.name, offset, size)
	a := MovMemReg(f.Indent, optypes, ident.name, roffset+offset, &r, reg)
	asm += a
	asm += f.Indent + fmt.Sprintf("// END LoadIdent, ident: %v, offset %v, size %v\n", ident.name, offset, size)
	return asm, reg, nil
}

func (f *Function) LoadValueSimple(val ssa.Value) (string, *register, *Error) {
	asm := f.Indent + fmt.Sprintf("// BEGIN LoadValueSimple, val: %v\n", val)
	a, reg, err := f.LoadValue(val, 0, f.sizeof(val))
	asm += a
	asm += f.Indent + fmt.Sprintf("// END LoadValueSimple, val: %v, reg %v\n", val, reg.name)
	return asm, reg, err
}

func (f *Function) LoadValue(val ssa.Value, offset int, size uint) (string, *register, *Error) {

	if _, ok := val.(*ssa.Const); ok {
		return f.LoadConstValue(val.(*ssa.Const))
	}
	ident, ok := f.identifiers[val.Name()]
	if !ok {
		internal(fmt.Sprintf("unknown name (%v), value (%v)\n", val.Name(), val))
	}

	asm := f.Indent + fmt.Sprintf("// BEGIN LoadValue, val: %v, offset %v, size %v\n", val, offset, size)
	a, reg, err := f.LoadIdent(&ident, offset, size, val.Pos())
	if err != nil {
		return "", nil, err
	}
	asm += a
	asm += f.Indent + fmt.Sprintf("// END LoadValue, val: %v, offset %v, size %v\n", val, offset, size)
	return asm, reg, nil
}

func (f *Function) StoreSimd(reg *register, ident *identifier, pos token.Pos) (string, *Error) {
	return f.StoreValue(ident, reg, pos)
}

func (f *Function) StoreSSE2(reg *register, ident *identifier, pos token.Pos) (string, *Error) {
	return f.StoreValue(ident, reg, pos)
}

func (f *Function) StoreValue(ident *identifier, reg *register, pos token.Pos) (string, *Error) {
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
	if isSimd(ident.typ) || isSSE2(ident.typ) {
		asm := f.Indent + fmt.Sprintf("// BEGIN StoreReg (SIMD), size (%v)\n", size)
		optypes := GetOpDataType(ident.typ)
		if isIntegerSimd(ident.typ) || isIntegerSSE2(ident.typ) {
			asm += MovIntegerSimdRegMem(f.Indent, optypes, reg, ident.name, &r, offset+roffset)
		} else {
			asm += MovRegMem(f.Indent, optypes, reg, ident.name, &r, offset+roffset)
		}
		asm += f.Indent + fmt.Sprintf("// END StoreReg (SIMD), size (%v)\n", size)
		return asm, nil

	} else {
		asm := f.Indent + fmt.Sprintf("// BEGIN StoreReg, size (%v)\n", size)
		optypes := GetIntegerOpDataType(false, size)
		asm += MovRegMem(f.Indent, optypes, reg, ident.name, &r, offset+roffset)
		asm += f.Indent + fmt.Sprintf("// END StoreReg, size (%v)\n", size)
		return asm, nil
	}
}

func (f *Function) LoadConstValue(cnst *ssa.Const) (string, *register, *Error) {

	if isBool(cnst.Type()) {
		r := f.allocReg(regType(cnst.Type()), 1)
		var val int8
		if cnst.Value.String() == "true" {
			val = 1
		}
		return MovImm8Reg(f.Indent, val, &r), &r, nil
	}
	if isFloat(cnst.Type()) {
		r := f.allocReg(regType(cnst.Type()), 1)
		if r.typ != XMM_REG {
			internal("can't load float const into non xmm register")
		}
		tmp := f.allocReg(DATA_REG, 8)
		asm := ""
		if isFloat32(cnst.Type()) {
			asm = MovImmf32Reg(f.Indent, float32(cnst.Float64()), &tmp, &r)
		} else {
			asm = MovImmf64Reg(f.Indent, cnst.Float64(), &tmp, &r)
		}
		f.freeReg(tmp)
		return asm, &r, nil

	}
	if isComplex(cnst.Type()) {
		internal("complex64/128 unsupported")
	}

	size := sizeof(cnst.Type())
	signed := signed(cnst.Type())
	r := f.allocReg(regType(cnst.Type()), size)
	var val int64
	if signed {
		val = cnst.Int64()
	} else {

		val = int64(cnst.Uint64())
	}
	return MovImmReg(f.Indent, val, size, &r), &r, nil
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

	if nameinfo := f.Ident(instr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		msg := fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr)
		internal(msg)
	}
	asm, reg, err := f.LoadValueSimple(instr.X)

	if err != nil {
		return asm, err
	}

	if reg.size() < 8 {
		asm += XorImm32Reg(f.Indent, xorVal, reg, reg.size())
	} else {
		asm += XorImm64Reg(f.Indent, int64(xorVal), reg, reg.size())
	}

	a, err := f.StoreValue(&addr, reg, instr.Pos())
	f.freeReg(*reg)

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

	if nameinfo := f.Ident(instr); nameinfo == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}
	var regX *register
	var regSubX register
	var regVal register

	regVal = f.allocReg(regType(instr.Type()), f.sizeof(instr))
	regSubX = f.allocReg(regType(instr.X.Type()), f.sizeof(instr.X))

	addr, ok := f.identifiers[instr.Name()]
	if !ok {
		msg := fmt.Sprintf("unknown name (%v), instr (%v)\n", instr.Name(), instr)
		internal(msg)
	}

	asm, regX, err := f.LoadValueSimple(instr.X)
	if err != nil {
		return asm, err
	}

	asm += ZeroReg(f.Indent, &regSubX)
	optypes := GetOpDataType(instr.Type())
	asm += ArithOp(f.Indent, optypes, token.SUB, &regSubX, regX, &regVal)
	f.freeReg(*regX)
	f.freeReg(regSubX)

	a, err := f.StoreValue(&addr, &regVal, instr.Pos())
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

//pointer indirection, in assignment such as "z = *x"
func (f *Function) UnOpPointer(instr *ssa.UnOp) (string, *Error) {
	assignment := f.Ident(instr)
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

	// ssa locals are always pointers (even though actually not)
	if xSize != sizePtr() && !xInfo.IsSsaLocal() {
		internal(fmt.Sprintf("xSize (%v) != ptr size (%v)", xSize, sizePtr()))
	}

	size := aSize

	optypes := GetOpDataType(instr.Type())

	if xInfo.IsSsaLocal() {
		var tmp register
		if isSimd(instr.Type()) || isSSE2(instr.Type()) {
			tmp = f.allocReg(XMM_REG, XmmRegSize)
		} else {
			tmp = f.allocReg(DATA_REG, DataRegSize)
		}
		asm += MovMemMem(
			f.Indent,
			xInfo.name,
			xOffset,
			&xReg,
			assignment.name,
			aOffset,
			&aReg,
			size,
			&tmp)
		f.freeReg(tmp)
	} else {
		var tmpData register
		if isSimd(instr.Type()) || isSSE2(instr.Type()) {
			tmpData = f.allocReg(XMM_REG, XmmRegSize)
		} else {
			tmpData = f.allocReg(regType(instr.Type()), DataRegSize)
		}
		tmpAddr := f.allocReg(DATA_REG, DataRegSize)

		asm += MovMemIndirectMem(
			f.Indent,
			optypes,
			xInfo.name,
			xOffset,
			&xReg,
			assignment.name,
			aOffset,
			&aReg,
			size,
			&tmpAddr,
			&tmpData)
		f.freeReg(tmpAddr)
		f.freeReg(tmpData)

	}
	f.identifiers[assignment.name] = *assignment

	asm = fmt.Sprintf(f.Indent+"// BEGIN ssa.UnOpPointer, %v = %v\n", instr.Name(), instr) + asm
	asm += fmt.Sprintf(f.Indent+"// END ssa.UnOpPointer, %v = %v\n", instr.Name(), instr)
	return asm, nil
}

func (f *Function) Index(instr *ssa.Index) (string, *Error) {
	if instr == nil {
		return ErrorMsg("nil instr")
	}
	asm := ""
	xInfo := f.identifiers[instr.X.Name()]
	assignment := f.Ident(instr)
	if assignment == nil {
		return ErrorMsg(fmt.Sprintf("Cannot alloc value: %v", instr))
	}

	xReg, xOffset, _ := xInfo.Addr()
	aReg, aOffset, _ := assignment.Addr()
	addrReg := f.allocReg(DATA_REG, sizePtr())
	//idxReg := f.allocReg(DATA_REG, sizePtr())

	a, idxReg, err := f.LoadValueSimple(instr.Index)
	if err != nil {
		return "", err
	}
	asm += a

	asm += Lea(f.Indent, xInfo.name, xOffset, &xReg, &addrReg)

	optypes := GetIntegerOpDataType(false, idxReg.width/8)
	asm += AddRegReg(f.Indent, optypes, idxReg, &addrReg)

	optypes = GetOpDataType(assignment.typ)
	asm += MovRegMem(f.Indent, optypes, &addrReg, assignment.name, &aReg, aOffset)

	f.freeReg(*idxReg)
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
	assignment := f.Ident(instr)

	xReg, xOffset, _ := xInfo.Addr()
	aReg, aOffset, _ := assignment.Addr()
	addrReg := f.allocReg(DATA_REG, sizePtr())
	//idxReg := f.allocReg(DATA_REG, sizePtr())

	a, idxReg, err := f.LoadValueSimple(instr.Index)
	if err != nil {
		return "", err
	}
	asm += a

	asm += MulImm32RegReg(f.Indent, uint32(sizeofElem(xInfo.typ)), idxReg, idxReg)

	if isSlice(xInfo.typ) {
		// TODO: add bounds checking
		optypes := GetIntegerOpDataType(false, sizePtr())
		asm += MovMemReg(f.Indent, optypes, xInfo.name, xOffset, &xReg, &addrReg)
	} else if isArray(xInfo.typ) {
		asm += Lea(f.Indent, xInfo.name, xOffset, &xReg, &addrReg)
	} else if isSimd(xInfo.typ) {
		asm += Lea(f.Indent, xInfo.name, xOffset, &xReg, &addrReg)
	} else {

		internal(fmt.Sprintf("indexing non-slice/array variable, type %v", xInfo.typ))
	}

	optypes := GetIntegerOpDataType(false, idxReg.size())

	asm += AddRegReg(f.Indent, optypes, idxReg, &addrReg)

	optypes = GetIntegerOpDataType(false, assignment.size())
	asm += MovRegMem(f.Indent, optypes, &addrReg, assignment.name, &aReg, aOffset)

	f.freeReg(*idxReg)
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
		msg := "Heap allocations are unsupported (are all print and log statements removed?), ssa variable: %v, type: %v"
		msgstr := fmt.Sprintf(msg, instr.Name(), instr.Type())
		return ErrorMsg(msgstr)
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

func internal(msg string) string {
	panic(fmt.Sprintf("Internal error, \"%v\"", msg))
}

func StripDebugComments(assembly, indent string) string {
	lines := strings.Split(assembly, "\n")
	stripped := ""
	begin := indent + "// BEGIN"
	end := indent + "// END"
	for _, line := range lines {
		// skip debug comments
		if strings.HasPrefix(line, begin) || strings.HasPrefix(line, end) {
			continue
		}
		stripped += line + "\n"
	}
	return stripped
}
