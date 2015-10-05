package codegen

import (
	"errors"
	"fmt"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/types"

	"reflect"
	"unsafe"

	"github.com/bjwbell/gensimd/codegen/instructionsetxml"
	"github.com/bjwbell/gensimd/simd"

	"golang.org/x/tools/go/ssa"
)

type Function struct {
	Indent         string
	ssa            *ssa.Function
	instructionset *instructionsetxml.Instructionset
	locals         map[string]varInfo
	params         map[string]paramInfo
	registers      map[register]bool // maps register to false if unused and true if used
}

type varInfo struct {
	name string
	// offset from the stack base (SB)
	offset int
	size   int
	info   *ssa.Alloc
}

type paramInfo struct {
	name string
	// offset from the frame pointer (FP)
	offset int
	size   int
	info   *ssa.Parameter
	extra  interface{}
}

type paramSlice struct {
	offset      int
	lenOffset   int
	reg         register
	regValid    bool
	lenReg      register
	lenRegValid bool
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
	if err := f.init(); err != nil {
		return "", err
	}
	return f.asmFunc()
}

func (f *Function) initLocals() *Error {
	offset := 0
	locals := f.ssa.Locals
	for _, local := range locals {
		if local.Heap {
			return &Error{Err: errors.New(fmt.Sprintf("Can't heap alloc local, name: %v", local.Name())), Pos: local.Pos()}
		}
		size := sizeof(local.Type())
		v := varInfo{name: local.Name(), offset: offset, size: size, info: local}
		f.locals[v.name] = v
		offset += size
	}
	return nil
}

func memFn(name string, offset int) func() string {
	return func() string {
		return name + "+" +
			fmt.Sprintf("%v", offset) +
			"(FP)"
	}
}

func regFn(name string) func() string {
	return func() string {
		return name
	}
}

func (f *Function) asmParams() (string, *Error) {
	// offset in bytes
	offset := 0
	asm := ""
	for _, p := range f.ssa.Params {
		param := paramInfo{name: p.Name(), offset: offset, info: p, size: sizeof(p.Type())}
		// TODO alloc reg based on other param types
		if _, ok := p.Type().(*types.Slice); ok {
			opMem := Operand{
				Type:   OperandType(M64),
				Input:  true,
				Output: false,
				Value:  nil}
			opReg := Operand{
				Type:   OperandType(R64),
				Input:  false,
				Output: true,
				Value:  nil}
			ops := []*Operand{&opMem, &opReg}
			// TODO is sizeof data always pointer size?
			fmt.Println("pointerSize:", pointerSize)
			reg := f.allocReg(AddrReg, pointerSize)
			opMem.Value = memFn(param.name, offset)
			opReg.Value = regFn(reg.name)
			if a, err := InstAsm(f.instructionset, InstName(MOVQ), ops); err != nil {
				return "", &Error{err, p.Pos()}
			} else {
				asm += f.Indent + a + "\n"
			}

			// TODO is sizeof length data always pointer size?
			lenReg := f.allocReg(AddrReg, pointerSize)
			opMem.Value = memFn("len", offset+pointerSize)
			opReg.Value = regFn(lenReg.name)
			if a, err := InstAsm(f.instructionset, InstName(MOVQ), ops); err != nil {
				return "", &Error{err, p.Pos()}
			} else {
				asm += f.Indent + a + "\n"
			}
			param.extra = paramSlice{offset: offset, reg: reg, regValid: true, lenReg: lenReg, lenRegValid: true}
		} else {
			return "", &Error{Err: errors.New("Unsupported param type"), Pos: p.Pos()}
		}
		f.params[param.name] = param
		offset += param.size
	}
	return asm, nil
}

func (f *Function) asmFunc() (string, *Error) {
	fpSize := f.localsSize()
	funcAsm, err := f.asmParams()
	if err != nil {
		return "", err
	}
	basicblocksAsm, err2 := f.asmBasicBlocks()
	if err2 != nil {
		return "", err2
	}

	funcAsm += basicblocksAsm

	asm := fmt.Sprintf("TEXT ·%v(SB),NOSPLIT,$%v-$%v\n%v", f.ssa.Name(), fpSize, f.paramsSize(), funcAsm)
	return asm, nil
}

func (f *Function) asmBasicBlocks() (string, *Error) {
	asm := ""
	for i := 0; i < len(f.ssa.Blocks); i++ {
		asm += f.asmBasicBlock(f.ssa.Blocks[i])
	}
	return asm, nil
}

func (f *Function) asmBasicBlock(block *ssa.BasicBlock) string {
	asm := strconv.Itoa(block.Index) + ":\n"
	for i := 0; i < len(block.Instrs); i++ {

		asm += f.Indent + f.asmInstr(block.Instrs[i])
	}
	return asm
}

func (f *Function) asmInstr(instr ssa.Instruction) string {
	if instr == nil {
		panic("Nil instr in asmInstr")
	}
	// TODO
	asm := ""
	switch instr.(type) {
	default:
		asm += fmt.Sprintf("Unknown ssa instruction: %v", instr)
	case *ssa.Alloc:
		i := instr.(*ssa.Alloc)
		asm += fmt.Sprintf("ssa.Alloc: %v, name: %v", instr, i.Name())
	case *ssa.BinOp:
		i := instr.(*ssa.BinOp)
		asm += fmt.Sprintf("ssa.BinOp: %v, name: %v", instr, i.Name())
	case *ssa.Call:
		i := instr.(*ssa.Call)
		asm += fmt.Sprintf("ssa.Call: %v, name: %v", instr, i.Name())
	case *ssa.ChangeInterface:
		i := instr.(*ssa.ChangeInterface)
		asm += fmt.Sprintf("ssa.ChangeInterface: %v, name: %v", instr, i.Name())
	case *ssa.ChangeType:
		i := instr.(*ssa.ChangeType)
		asm += fmt.Sprintf("ssa.ChangeType: %v, name: %v", instr, i.Name())
	case *ssa.Convert:
		i := instr.(*ssa.Convert)
		asm += fmt.Sprintf("ssa.Convert: %v, name: %v", instr, i.Name())
	case *ssa.Defer:
		asm += fmt.Sprintf("ssa.Defer: %v", instr)
	case *ssa.Extract:
		i := instr.(*ssa.Extract)
		asm += fmt.Sprintf("ssa.Extra: %v, name: %v", instr, i.Name())
	case *ssa.Field:
		i := instr.(*ssa.Field)
		asm += fmt.Sprintf("ssa.Field: %v, name: %v", instr, i.Name())
	case *ssa.FieldAddr:
		i := instr.(*ssa.FieldAddr)
		asm += fmt.Sprintf("ssa.FieldAddr: %v, name: %v", instr, i.Name())
	case *ssa.Go:
		asm += fmt.Sprintf("ssa.Go: %v", instr)
	case *ssa.If:
		asm += fmt.Sprintf("ssa.If: %v", instr)
	case *ssa.Index:
		i := instr.(*ssa.Index)
		asm += fmt.Sprintf("ssa.Index: %v, name: %v", instr, i.Name())
	case *ssa.IndexAddr:
		i := instr.(*ssa.IndexAddr)
		asm += fmt.Sprintf("ssa.IndexAddr: %v, name: %v", i, i.Name())
	case *ssa.Jump:
		jmp := instr.(*ssa.Jump)
		asm += strings.Replace(jmp.String(), "jump", "JMP ", -1)
	case *ssa.Lookup:
		i := instr.(*ssa.Lookup)
		asm += fmt.Sprintf("ssa.Lookup: %v, name: %v", instr, i.Name())
	case *ssa.MakeChan:
		i := instr.(*ssa.MakeChan)
		asm += fmt.Sprintf("ssa.MakeChan: %v, name: %v", instr, i.Name())
	case *ssa.MakeClosure:
		i := instr.(*ssa.MakeClosure)
		asm += fmt.Sprintf("ssa.MakeClosure: %v, name: %v", instr, i.Name())
	case *ssa.MakeInterface:
		i := instr.(*ssa.MakeInterface)
		asm += fmt.Sprintf("ssa.MakeInterface: %v, name: %v", instr, i.Name())
	case *ssa.MakeMap:
		i := instr.(*ssa.MakeMap)
		asm += fmt.Sprintf("ssa.MakeMap: %v, name: %v", instr, i.Name())
	case *ssa.MakeSlice:
		i := instr.(*ssa.MakeSlice)
		asm += fmt.Sprintf("ssa.MakeSlice: %v, name: %v", instr, i.Name())
	case *ssa.MapUpdate:
		asm += fmt.Sprintf("ssa.MapUpdate: %v", instr)
	case *ssa.Next:
		i := instr.(*ssa.Next)
		asm += fmt.Sprintf("ssa.Next: %v, name: %v", instr, i.Name())
	case *ssa.Panic:
		asm += fmt.Sprintf("ssa.Panic: %v", instr)
	case *ssa.Phi:
		i := instr.(*ssa.Phi)
		asm += fmt.Sprintf("ssa.Phi: %v, name: %v", instr, i.Name())
	case *ssa.Range:
		i := instr.(*ssa.Range)
		asm += fmt.Sprintf("ssa.Range: %v, name: %v", instr, i.Name())
	case *ssa.Return:
		asm += fmt.Sprintf("ssa.Return: %v", instr)
	case *ssa.RunDefers:
		asm += fmt.Sprintf("ssa.RunDefers: %v", instr)
	case *ssa.Select:
		i := instr.(*ssa.Select)
		asm += fmt.Sprintf("ssa.Select: %v, name: %v", instr, i.Name())
	case *ssa.Send:
		asm += fmt.Sprintf("ssa.Send: %v", instr)
	case *ssa.Slice:
		i := instr.(*ssa.Slice)
		asm += fmt.Sprintf("ssa.Slice: %v, name: %v", instr, i.Name())
	case *ssa.Store:
		i := instr.(*ssa.Store)
		asm += fmt.Sprintf("ssa.Store: %v, addr: %v, val: %v", instr, i.Addr, i.Val)
	case *ssa.TypeAssert:
		i := instr.(*ssa.TypeAssert)
		asm += fmt.Sprintf("ssa.TypeAssert: %v, name: %v", instr, i.Name())
	case *ssa.UnOp:
		i := instr.(*ssa.UnOp)
		asm += fmt.Sprintf("ssa.UnOp: %v, name: %v", instr, i.Name())
	}
	return asm + "\n"
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

func (f *Function) localsSize() int {
	size := 0
	for _, v := range f.locals {
		size += v.size
	}
	return size
}

func (f *Function) init() *Error {
	f.locals = make(map[string]varInfo)
	f.params = make(map[string]paramInfo)
	f.registers = make(map[register]bool)
	f.initRegs()
	return f.initLocals()
}

func (f *Function) initRegs() {
	for _, r := range registers {
		f.registers[r] = false
	}
}

// width in bytes
func (f *Function) allocReg(t RegType, width int) register {
	var reg register
	found := false
	for i := 0; i < len(registers); i++ {
		r := registers[i]
		used := f.registers[r]
		// r.width is in bits so multiple width (which is in bytes) by 8
		if !used && r.typ == t && r.width == width*8 {
			reg = r
			found = true
			break
		}
	}
	if found {
		f.registers[reg] = true
	} else {
		// any of the data registers can be used as an address register on x86_64
		if t == AddrReg {
			return f.allocReg(DataReg, width)
		} else {
			panic(fmt.Sprintf("couldn't alloc register, type: %v, size: %v", t, width*8))
		}
	}
	return reg
}

func (f *Function) freeReg(reg register) {
	f.registers[reg] = false
}

// paramsSize returns the size of the parameters in bytes
func (f *Function) paramsSize() int {
	size := 0
	for _, p := range f.ssa.Params {
		size += sizeof(p.Type())
	}
	return size
}

var pointerSize = 8
var sliceSize = 24
var pointerSizeInBits = 8 * 8
var sliceSizeInBits = 24 * 8

func sizeof(t types.Type) int {
	switch t.(type) {
	default:
		panic("Error unknown type in sizeof")
	case *types.Basic:
		basic, _ := t.(*types.Basic)
		return sizeBasic(basic)
	case *types.Pointer:
		return pointerSize
	case *types.Slice:
		return sliceSize
	case *types.Array:
		arr, _ := t.(*types.Array)
		return int(arr.Len()) * sizeof(arr.Elem())
	case *types.Named:
		named, _ := t.(*types.Named)
		tname := named.Obj()
		i := simd.Int(0)
		simdInt := reflect.TypeOf(i)
		var i4 simd.Int4
		simdInt4 := reflect.TypeOf(i4)
		switch tname.Name() {
		default:
			panic("Error unknown type in sizeof")
		case simdInt.Name():
			return int(unsafe.Sizeof(i))
		case simdInt4.Name():
			return int(unsafe.Sizeof(i4))
		}
	}
}

var intSize = 8
var uintSize = 8
var boolSize = 1

// sizeBasic return the size in bytes of a basic type
func sizeBasic(b *types.Basic) int {
	switch b.Kind() {
	default:
		panic("Unknown basic type")
	case types.Bool:
		return 1
	case types.Int:
		return intSize
	case types.Int8:
		return 1
	case types.Int16:
		return 2
	case types.Int32:
		return 4
	case types.Int64:
		return 8
	case types.Uint:
		return uintSize
	case types.Uint8:
		return 1
	case types.Uint16:
		return 2
	case types.Uint32:
		return 4
	case types.Uint64:
		return 8
	case types.Float32:
		return 4
	case types.Float64:
		return 8
	}
}
