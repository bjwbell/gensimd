package simd

import (
	"errors"
	"fmt"
	"go/token"

	"golang.org/x/tools/go/types"

	"reflect"
	"unsafe"

	"github.com/bjwbell/gensimd/simd"

	"golang.org/x/tools/go/ssa"
)

type Function struct {
	ssa      *ssa.Function
	locals   map[string]varinfo
	params   map[string]paraminfo
	register map[register]bool // maps register to false if unused and true if used
}

type varinfo struct {
	name string
	// offset from the stack base (SB)
	offset int
	size   int
	info   *ssa.Alloc
}

type paraminfo struct {
	name string
	// offset from the frame pointer (FP)
	offset int
	size   int
	info   *ssa.Parameter
}

type Error struct {
	Err error
	Pos token.Pos
}

func (f *Function) GoAssembly() (string, *Error) {
	// TODO
	assembly := ""
	if err := f.initLocals(); err != nil {
		return "", err
	} else {

	}
	return assembly, nil
}

func (f *Function) initLocals() *Error {
	offset := 0
	locals := f.ssa.Locals
	for _, local := range locals {
		if local.Heap {
			return &Error{Err: errors.New(fmt.Sprintf("Can't heap alloc local, name: %v", local.Name())), Pos: local.Pos()}
		}
		size := sizeof(local.Type())
		v := varinfo{name: local.Name(), offset: offset, size: size, info: local}
		f.locals[v.name] = v
		offset += size
	}
	return nil
}

func (f *Function) asmParams() (string, *Error) {
	offset := 0
	for _, p := range f.ssa.Params {
		param := paraminfo{name: p.Name(), offset: offset, info: p, size: sizeof(p.Type())}
		f.params[param.name] = param
		offset += param.size
		// TODO alloc reg based on param type
	}
	return "", nil
}

func (f *Function) asmFunc() string {
	fpSize := f.paramsSize()
	funcAsm := ""
	asm := fmt.Sprintf(`TEXT ·%v(SB),$%v-$%v
	%v
	RET`, f.ssa.Name(), f.paramsSize(), fpSize, funcAsm)
	return asm
}

func (f *Function) asmValue(value ssa.Value, dstReg *register, dstVar *varinfo) string {
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

func (f *Function) init() {
	f.initRegs()
	f.initLocals()
}

func (f *Function) initRegs() {
	for _, r := range registers {
		f.register[r] = false
	}
}

func (f *Function) allocReg(t RegType, size int) register {
	var reg register
	found := false
	for r, used := range f.register {
		if !used && r.typ == t {
			reg = r
			found = true
			break
		}
	}
	if found {
		f.register[reg] = true
	} else {
		panic(fmt.Sprintf("couldn't alloc register, type: %v, size: %v", t, size))
	}
	return reg
}

func (f *Function) freeReg(reg register) {
	f.register[reg] = false
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
var sliceSize = 16

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
