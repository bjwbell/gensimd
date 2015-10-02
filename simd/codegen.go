package simd

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"unsafe"
)

func (f *File) Assembly() (string, *Error) {
	// TODO
	if err := f.Valid(); err != nil {
		return "", &Error{errors.New(fmt.Sprintf("File validation failed, message: %v", err.Err)), err.Pos}
	}
	assembly := ""
	for _, decl := range f.ast.Decls {
		if asm, err := f.asmTopLevelDecl(decl); err != nil {
			return "", err
		} else {
			assembly = assembly + "\n" + asm
		}
	}

	return assembly, nil
}

func (f *File) asmTopLevelDecl(decl ast.Decl) (string, *Error) {
	asm := ""
	if funcDecl, ok := decl.(*ast.FuncDecl); ok {
		fn := Function{
			file:      f,
			name:      funcDecl.Name.Name,
			decl:      funcDecl,
			vars:      nil,
			initBlock: nil,
			unusedReg: nil,
			usedReg:   nil}
		fn.init()
		if fnAsm, err := fn.asm(); err != nil {
			return "", err
		} else {
			asm = asm + "\n" + fnAsm
		}
	}
	return asm, nil
}

func (fn *Function) asm() (string, *Error) {
	asm := fn.asmFunc()
	return asm, nil
}

func (fn *Function) asmFunc() string {
	fpSize := fn.varsSize()
	funcAsm := ""
	asm := fmt.Sprintf(`TEXT ·%v(SB),$%v-$%v
	%v
	RET`, fn.name, fn.argsSize(), fpSize, funcAsm)
	return asm
}

func (fn *Function) asmInitStmts(block *ast.BlockStmt) string {
	if fn.initBlock == nil {
		panic("ERROR fn.init == nil")
	} else {
		if len(fn.initBlock) == 0 {
			return ""
		}
		return ""
	}
}

func (fn *Function) varsSize() int {
	size := 0
	for _, v := range fn.vars {
		size += v.size
	}
	return size
}

func (fn *Function) init() {
	fn.initRegs()
	fn.initVarsInfo()
}

func (fn *Function) initVarsInfo() {
	size := 0
	if vars, initStmts, err := fn.file.validFuncBody(fn.decl.Body); err != nil {
		panic("Unexpected error in function body")
	} else {
		for _, stmt := range vars {
			decl := stmt.Decl
			d, _ := decl.(*ast.GenDecl)
			if d.Tok != token.VAR {
				panic("Unexpected decl token")
			}
			spec := d.Specs[0]
			vspec, _ := spec.(*ast.ValueSpec)
			name := vspec.Names[0]
			varName := name.Name
			s := fn.file.sizeof(fn.file.info.TypeOf(vspec.Type))
			fn.vars[varName] = varinfo{
				name:   varName,
				offset: size,
				size:   s}
			size += s
		}
		fn.initBlock = initStmts
	}
}

func (fn *Function) initRegs() {
	for _, r := range regnames {
		typ := IntReg
		size := IntRegSize
		if r[0] == 'X' {
			typ = FloatReg
			size = FloatRegSize
		}
		reg := register{r, typ, size}
		fn.unusedReg = append(fn.unusedReg, reg)
	}
}

func (fn *Function) allocReg(t registerType, size int) register {
	// TODO
	return register{"", IntReg, 0}
}

func (fn *Function) freeReg(reg register) {
	// TODO
	fmt.Println("freereg: ", reg)
}

// argsSize returns the size of the arguments in bytes
func (fn *Function) argsSize() int {
	if fn.decl.Type.Params == nil {
		return 0
	}
	size := 0
	for _, param := range fn.decl.Type.Params.List {
		t := fn.file.info.TypeOf(param.Type)
		size += fn.file.sizeof(t)
	}
	return size
}

var pointerSize = 8
var sliceSize = 16

func (f *File) sizeof(t types.Type) int {
	switch t.(type) {
	default:
		panic("Error unknown type in sizeof")
	case *types.Basic:
		basic, _ := t.(*types.Basic)
		return f.sizeBasic(basic)
	case *types.Pointer:
		return pointerSize
	case *types.Slice:
		return sliceSize
	case *types.Array:
		arr, _ := t.(*types.Array)
		return int(arr.Len()) * f.sizeof(arr.Elem())
	case *types.Named:
		named, _ := t.(*types.Named)
		tname := named.Obj()
		i := Int(0)
		ivar := &i
		simdInt := reflect.TypeOf(i)
		simdIntVar := reflect.TypeOf(ivar)
		var i4var Int4Var
		simdInt4Var := reflect.TypeOf(i4var)
		var i4 Int4
		simdInt4 := reflect.TypeOf(i4)
		switch tname.Name() {
		default:
			panic("Error unknown type in sizeof")
		case simdInt.Name():
			return int(unsafe.Sizeof(i))
		case simdInt4.Name():
			return int(unsafe.Sizeof(i4))
		case simdIntVar.Name():
			return int(unsafe.Sizeof(IntVar(nil)))
		case simdInt4Var.Name():
			return int(unsafe.Sizeof(Int4Var(nil)))
		}
	}
}

var intSize = 8
var uintSize = 8
var boolSize = 1

// sizeBasic return the size in bytes of a basic type
func (f *File) sizeBasic(b *types.Basic) int {
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
