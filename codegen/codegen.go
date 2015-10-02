package simd

import (
	"fmt"
	"go/token"

	"golang.org/x/tools/go/types"

	"reflect"
	"unsafe"

	"github.com/bjwbell/gensimd/simd"

	"golang.org/x/tools/go/ssa"
)

type Function struct {
	fn        *ssa.Function
	vars      map[string]varinfo
	unusedReg []register
	usedReg   []register
}

type varinfo struct {
	name   string
	offset int
	size   int
}

type Error struct {
	Err error
	Pos token.Pos
}

func (f *Function) GoAssembly() (string, *Error) {
	// TODO
	assembly := ""
	return assembly, nil
}

func (f *Function) asmFunc() string {
	fpSize := f.varsSize()
	funcAsm := ""
	asm := fmt.Sprintf(`TEXT ·%v(SB),$%v-$%v
	%v
	RET`, f.fn.Name(), f.argsSize(), fpSize, funcAsm)
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

func (f *Function) varsSize() int {
	size := 0
	for _, v := range f.vars {
		size += v.size
	}
	return size
}

func (f *Function) init() {
	f.initRegs()
	f.initVarsInfo()
}

func (f *Function) initVarsInfo() {
	// TODO

}

func (f *Function) initRegs() {
	for _, r := range regnames {
		typ := IntReg
		size := IntRegSize
		if r[0] == 'X' {
			typ = FloatReg
			size = FloatRegSize
		}
		reg := register{r, typ, size}
		f.unusedReg = append(f.unusedReg, reg)
	}
}

func (f *Function) allocReg(t registerType, size int) register {
	reg := register{typ: t, size: size, name: ""}
	return f.moveReg(&f.unusedReg, &f.usedReg, reg)
}

func (f *Function) freeReg(reg register) {
	f.moveReg(&f.usedReg, &f.unusedReg, reg)
}

func (f *Function) moveReg(src *[]register, dst *[]register, reg register) register {
	regs := []register{}
	var mreg *register
	for _, r := range *src {
		if r.typ == reg.typ && r.size == reg.size && (reg.name == "" || reg.name == r.name) {

			mreg = &r
		} else {
			regs = append(regs, r)
		}
	}
	src = &regs
	if mreg == nil {
		panic(fmt.Sprintf("Couldn't move register: type:%v, size:%v\n", reg.typ, reg.size))
	} else {
		*dst = append(*dst, *mreg)
		return *mreg
	}
}

// argsSize returns the size of the arguments in bytes
func (f *Function) argsSize() int {
	size := 0
	for _, p := range f.fn.Params {
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
